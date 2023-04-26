package util

import (
	"errors"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SplitStringIntoAddresses(str string) ([]sdk.AccAddress, error) {
	regex := (`^([a-z0-9]{40,80},)*[a-z0-9]{40,80}$`)

	if match, err := regexp.MatchString(regex, str); !match || err != nil {
		return nil, errors.New("invalid addresses")
	}

	var addresses = []sdk.AccAddress{}

	for _, addr := range strings.Split(str, ",") {
		validAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, validAddr)
	}
	return addresses, nil
}

func SplitStringIntoDenoms(str string) ([]string, error) {
	regex := `^([a-zA-Z0-9]{3,64},)*[a-zA-Z0-9]{3,64}$`

	if match, err := regexp.MatchString(regex, str); !match || err != nil {
		return nil, errors.New("invalid denoms")
	}

	return strings.Split(str, ","), nil
}
