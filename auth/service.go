package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("STARTUP_CROWD_FUNDING")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signToken, err
	}

	return signToken, nil

}
