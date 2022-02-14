package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(EncodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

// janganbuatdisini
var SECRET_KEY = []byte("rahasiaya")

// agar bisa dipanggil ditempat lain
func NewService() *jwtService{
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// Payload
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// buat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// signature
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil{
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(EncodedToken string) (*jwt.Token, error){
	// cek apakah secret key sama
	token, err := jwt.Parse(EncodedToken, func(t *jwt.Token) (interface{}, error) {
		// cek apakah method yg digunakan sama
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil{
		return token, err
	}

	return token, nil
}