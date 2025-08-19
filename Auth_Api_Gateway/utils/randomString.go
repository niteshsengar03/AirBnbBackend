package utils

import (
    "crypto/rand"
    "encoding/hex"
)

// GenerateRandomToken creates a secure random string of length nBytes*2
func GenerateRandomToken(nBytes int) (string, error) {
    b := make([]byte, nBytes)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(b), nil
}
