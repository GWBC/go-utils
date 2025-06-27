package utils

import "github.com/pquerna/otp/totp"

type TOTP struct {
}

func (t *TOTP) Gen(issuer string, accountName string, period uint) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
		Period:      period,
	})

	if err != nil {
		return "", "", err
	}

	return key.Secret(), key.URL(), nil
}

func (t *TOTP) Check(secret string, key string) bool {
	return totp.Validate(key, secret)
}
