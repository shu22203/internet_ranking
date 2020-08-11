package ir

import (
	"github.com/shu22203/internet_ranking/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHoldInternetRanking(t *testing.T) {
	user := user.NewUser()
	ir := HoldInternetRanking(user)

	t.Run("開催したuserはIRのadminになる", func(t *testing.T) {
		assert.Contains(t, ir.AdminIds(), user.Id)
	})

	t.Run("開催したuserはIRのownerになる", func(t *testing.T) {
		assert.Equal(t, ir.OwnerId(), user.Id)
	})
}
