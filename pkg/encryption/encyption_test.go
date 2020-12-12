package encryption

import (
  "github.com/stretchr/testify/assert"
  "strings"
  "testing"
)

func Test_Encrypt(t *testing.T) {
  const key = "super_secret"
  const scrt = "Shhhh! I am a super secret"

  t.Run("should encrypt a secret", func(t *testing.T) {
    entxt, err := Encrypt(scrt, key)
    if err != nil {
      t.Errorf("failed to encrypt: %v", err)
    }

    assert.Len(t, strings.Split(entxt, ":"), 2)

    detxt, err := Decrypt(entxt, key)
    if err != nil {
      t.Errorf("failed to decrypt: %v", err)
    }

    assert.Equal(t, scrt, detxt, "expected secret: %s, got: %s", scrt, detxt)
  })
}

func Test_Decrypt(t *testing.T) {
  const key = "super_secret"
  const scrt = "Shhhh! I am a super secret"

  t.Run("should fail if the encrypted secret is missing the nonce", func(t *testing.T) {
    entxt, err := Encrypt(scrt, key)
    if err != nil {
      t.Errorf("failed to encrypt: %v", err)
    }

    entxt = strings.Split(entxt, ":")[1]

    _, err = Decrypt(entxt, key)

    assert.NotNil(t, err, "expected error to not be nil")
  })

  t.Run("should decrypt a secret", func(t *testing.T) {
    entxt, err := Encrypt(scrt, key)
    if err != nil {
      t.Errorf("failed to encrypt: %v", err)
    }

    detxt, err := Decrypt(entxt, key)
    if err != nil {
      t.Errorf("failed to decrypt: %v", err)
    }

    assert.Equal(t, scrt, detxt, "expected secret: %s, got: %s", scrt, detxt)
  })
}
