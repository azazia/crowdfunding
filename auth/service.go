package auth

import "github.com/dgrijalva/jwt-go"

type service interface {
	GenerateToken(userID int) (string, error)
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