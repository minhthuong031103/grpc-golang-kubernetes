package jwt

import "github.com/kataras/jwt"

type JWT struct {
	SecretKey string
}

func NewJWT(secretKey string) *JWT {
	return &JWT{
		SecretKey: secretKey,
	}
}

func (j *JWT) GenerateToken(id string) (string, error) {
	user := map[string]string{
		"id": id,
	}
	token, err := jwt.Sign(jwt.HS256, []byte(j.SecretKey), user)
	return string(token), err
}

func (j *JWT) ValidateToken(token string) (map[string]interface{}, error) {
	verifiedToken, err := jwt.Verify(jwt.HS256, []byte(j.SecretKey), []byte(token))
	if err != nil {
		return nil, err
	}
	var claims map[string]interface{}
	verifiedToken.Claims(&claims)
	return claims, err
}
