package validators

import (
	"regexp"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/pkg/errors"
)

var addressRegexp = regexp.MustCompile("^T([a-km-zA-HJ-NP-Z1-9]{33})$")

// validateTronAddress returns nil as error if validation passed successfuly.
func validateTronAddress(address address.Address) error {
	match, err := regexp.MatchString(addressRegexp.String(), string(address[:]))
	if err != nil {
		return errors.Wrap(err, "failed to match regexp")
	}

	if !match {
		return errors.New("should be a valid tron address")
	}

	return nil
}
