{
  description = "solve";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-24.05";
  };
  outputs = { self, nixpkgs }:
    let

      # to work with older version of flakes
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";

      # Generate a user-friendly version number.
      version = builtins.substring 0 8 lastModifiedDate;

      # System types to support.
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in
    {

      # Provide some binary packages for selected system types.
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          solve = pkgs.buildGoModule {
            pname = "solve";
            inherit version;
            # In 'nix develop', we don't need a copy of the source tree
            # in the Nix store.
            src = nixpkgs.lib.cleanSource ./.;

            vendorHash = "sha256-sVVjWx3Ya0APuKsFtWKeOwUwExSymv8Ey9a0KcMMmZ0=";

            # Define the build inputs
            buildInputs = [ pkgs.go_1_21 ];
            nativeBuildInputs = [ pkgs.patchelf ];

            buildPhase = ''
              go build -o $out/build github.com/skip-mev/cns/cmd/cns_server
              go build -o $out/build github.com/skip-mev/cns/cmd/indexer
              go build -o $out/build github.com/skip-mev/cns/cmd/transaction_tracking
            '';

            # Use patchelf to fix rpath
            installPhase = ''
              mkdir -p $out/bin
              cp ./build/* $out/bin/
              for bin in $out/bin/*; do
                patchelf --set-rpath $out/lib $bin
              done
            '';

            # Will skip tests during build
            checkPhase = ''
            '';
          };
        });

      devShells = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [ go_1_21 gopls gotools go-tools jq buf postgresql_16 protobuf];
            shellHook = ''
              export PATH=${pkgs.writeShellScriptBin "tools" ''
                make -C tools local
              ''}/bin:$PATH
            '';
          };
        });

      # The default package for 'nix build'. This makes sense if the
      defaultPackage = forAllSystems (system: self.packages.${system}.solve);
    };
}
