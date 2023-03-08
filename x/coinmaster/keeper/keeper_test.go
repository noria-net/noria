package keeper_test

import (
	"testing"

	"github.com/noria-net/noria/x/coinmaster/keeper"
	"github.com/stretchr/testify/assert"
)

func Test_IsDenomWhiteListed(t *testing.T) {
	var result bool
	var denomToAllow = "allow_me"

	// List is empty
	result = keeper.IsDenomWhiteListed([]string{}, denomToAllow)
	assert.True(t, result)

	// List has one empty entry
	result = keeper.IsDenomWhiteListed([]string{""}, denomToAllow)
	assert.True(t, result)

	// List has one empty entry and another with an undefined value
	result = keeper.IsDenomWhiteListed([]string{"", "undefined"}, denomToAllow)
	assert.False(t, result)

	// List has one entry containing the allowed value
	result = keeper.IsDenomWhiteListed([]string{denomToAllow}, denomToAllow)
	assert.True(t, result)
}
