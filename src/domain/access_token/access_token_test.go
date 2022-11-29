package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test to check constants have not been modified
func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime)
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token shoould not be expired")
	assert.Equal(t, "", at.AccessToken, "new access toekn should not have defined access token id")
	assert.True(t, at.UserID == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should not be expired")
}
