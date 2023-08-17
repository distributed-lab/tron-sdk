package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateTronAddress(t *testing.T) {

	t.Run("invalid", func(t *testing.T) {
		err := validateTronAddress([]byte("dsadsaasd"))
		assert.Error(t, err)
		assert.Equal(t, "should be a valid tron address", err.Error())
	})
}

func TestValidateTronAddres_invalid(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		err := validateTronAddress([]byte("TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs"))
		assert.NoError(t, err)
	})
}
