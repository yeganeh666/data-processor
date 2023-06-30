package envext

import (
	"errors"
	"strings"
)

var ErrInvalidSecretFile = errors.New("invalid secret file, expected 2 string separated by space")

// Credential encapsulates username and password
type Credential struct {
	Username string
	Password string
}

func (c *Credential) UnmarshalText(data []byte) error {
	username, password, err := unmarshalText(data)
	if err != nil {
		return err
	}
	c.Username = username
	c.Password = password
	return nil
}

// APICredential is like `Credential`, except that it uses
// the combination of AccessKey/SecretKey
type APICredential struct {
	AccessKey string
	SecretKey string
}

func (c *APICredential) UnmarshalText(data []byte) error {
	accessKey, secret, err := unmarshalText(data)
	if err != nil {
		return err
	}
	c.AccessKey = accessKey
	c.SecretKey = secret
	return nil
}

func unmarshalText(data []byte) (string, string, error) {
	dump := strings.Split(strings.TrimSpace(string(data)), " ")
	if len(dump) < 2 {
		return "", "", ErrInvalidSecretFile
	}
	return strings.TrimSpace(dump[0]), strings.TrimSpace(dump[len(dump)-1]), nil
}
