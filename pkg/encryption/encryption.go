package encryption

import (
  "crypto"
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "encoding/hex"
  "errors"
  "fmt"
  "io"
  "strings"
)

func hashWithSha256(k string) []byte {
  hsh := crypto.SHA256.New()

  hsh.Write([]byte(k))

  return hsh.Sum(nil)
}

func Decrypt(entxt string, k string) (string, error) {
  enkey := hashWithSha256(k)
  spltxt := strings.Split(entxt, ":")
  if len(spltxt) < 2 || len(spltxt) > 2 {
    return "", errors.New("decryption failed: malformed encrypted text")
  }

  nonce, _ := hex.DecodeString(spltxt[0])
  cphrtxt, _ := hex.DecodeString(spltxt[1])
  block, err := aes.NewCipher(enkey)
  if err != nil {
    return "", err
  }

  aesgcm, err := cipher.NewGCM(block)
  if err != nil {
    return "", err
  }

  plntxt, err := aesgcm.Open(nil, nonce, cphrtxt, nil)
  if err != nil {
    return "", err
  }

  return string(plntxt), nil
}

func Encrypt(plntxt string, k string) (string, error) {
  nonce := make([]byte, 12) // Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
  if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
    return "", err
  }

  enkey := hashWithSha256(k)
  block, err := aes.NewCipher(enkey)
  if err != nil {
    return "", err
  }

  aesgcm, err := cipher.NewGCM(block)
  if err != nil {
    return "", err
  }

  cphrtxt := aesgcm.Seal(nil, nonce, []byte(plntxt), nil)

  return fmt.Sprintf("%s:%s", hex.EncodeToString(nonce), hex.EncodeToString(cphrtxt)), nil
}
