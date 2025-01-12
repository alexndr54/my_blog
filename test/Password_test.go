package main

import (
	"blog/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	teks := "demo"
	var PassswordEncrypt string

	t.Run("CreateHash", func(t *testing.T) {
		PasswordEnc, err := helper.EncryptPassword(teks)
		assert.Nil(t, err)
		PassswordEncrypt = PasswordEnc
		t.Log("Password Hashed: ", PasswordEnc)
	})

	t.Run("VerifyPassword", func(t *testing.T) {
		Verif := helper.VerifyPassword(PassswordEncrypt, teks)
		assert.True(t, Verif)
		t.Log("Passsword Verify: ", Verif)
	})

}

func TestVerifyPassword(t *testing.T) {
	PasswordHashed := "$2a$10$Umt/AguxazFp2wWmqW90FeYtLE4kalEdtlak2g7VyJ632IbqE2GHK"
	PasswordAsli := "sialexsofficial@gmail.com"
	Verify := helper.VerifyPassword(PasswordHashed, PasswordAsli)
	assert.True(t, Verify)
	t.Log("Passsword Verify: ", Verify)
}
