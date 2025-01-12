package helper

import (
	"blog/entity"
	"golang.org/x/crypto/bcrypt"
	"html/template"
)

func ShowValueMessage(ValidationValue map[string]string, key string) string {
	return ValidationValue[key]
}
func ShowMessage(Status entity.StatusResponse) template.HTML {
	//if Status == nil {
	//	return template.HTML("")
	//}

	htmlRes := "<div class=\"alert alert-" + Status.Label + "\">" + Status.Message + "</div>"
	return template.HTML(htmlRes)
}

func ShowFailedMessage(s map[string]entity.ValidationList, key string) template.HTML {
	html := "<div style=\"color:red\">" + s[key].Message + "</div>"
	return template.HTML(html)
}

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false // password tidak cocok
	}
	return true // password cocok
}
