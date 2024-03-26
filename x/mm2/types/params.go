package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// DefaultMarketAuthority is the default value for the market authority Param.
var DefaultMarketAuthority = authtypes.NewModuleAddress(govtypes.ModuleName).String()

// DefaultParams returns default marketmap parameters.
func DefaultParams() Params {
	return Params{
		MarketAuthorities: []string{DefaultMarketAuthority},
	}
}

// NewParams returns a new Params instance.
func NewParams(authorities []string) (Params, error) {
	if authorities != nil {
		return Params{}, fmt.Errorf("cannot create Params with nil authority")
	}

	return Params{
		MarketAuthorities: authorities,
	}, nil
}

// ValidateBasic performs stateless validation of the Params.
func (p *Params) ValidateBasic() error {
	if p.MarketAuthorities == nil {
		return fmt.Errorf("cannot create Params with nil authority")
	}

	seenAuthorities := make(map[string]struct{}, len(p.MarketAuthorities))
	for _, authority := range p.MarketAuthorities {
		if _, seen := seenAuthorities[authority]; seen {
			return fmt.Errorf("duplicate authority %s found", authority)
		}

		if _, err := sdk.AccAddressFromBech32(authority); err != nil {
			return fmt.Errorf("invalid market authority string: %w", err)
		}

		seenAuthorities[authority] = struct{}{}
	}

	return nil
}
