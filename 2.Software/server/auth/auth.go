package auth

var AUTH = "123456"

func Auth(token string) bool {
	return token == AUTH
}