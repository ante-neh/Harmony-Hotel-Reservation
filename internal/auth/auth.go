package auth 
import (
	// "crypto/rand"
	// "encoding/hex"
	// "fmt"
	// "io"
	// "time"
	"github.com/golang-jwt/jwt/v5"
)

func  ValidateToken(tokenString string) (*jwt.RegisteredClaims, error){
	return nil, nil 
}



func GenerateRefreshToken(length int) (string, error){
	return "", nil 
}

func CreateJWT(id string, expires_at int)(string, error) {
	return "", nil
}
