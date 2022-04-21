package jwtUsecase

import (
	"api_ticket/repository"
	"api_ticket/usecase"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtUsecase struct {
	ticketRepo repository.TicketRepositoryInterface
}

func GetJwtUsecase(ticketRepository repository.TicketRepositoryInterface) usecase.JwtUsecase {
	return &jwtUsecase{
		ticketRepo: ticketRepository,
	}
}

type CustomClaim struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

func (jwtuc *jwtUsecase) GenerateToken(userId string) (string, error) {
	claim := CustomClaim{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
	// 	UserID: userId,
	// 	StandardClaims: jwt.StandardClaims{
	// 		IssuedAt:  time.Now().Unix(),
	// 		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
	// 		Issuer:    os.Getenv("APP_NAME"),
	// 	},
	// })
	// return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (jwtuc *jwtUsecase) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	// claims := &CustomClaim{}
	// jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte(os.Getenv("SECRET_KEY")), nil
	// })
	// return claims.UserID
}

func (jwtuc *jwtUsecase) ValidateTokenAndGetUserId(token string) (string, error) {

	validatedToken, _ := jwtuc.ValidateToken(token)
	// if err != nil {
	// 	return "", err
	// }

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to claim token")
	}
	if claims["exp"].(float64) <= float64(time.Now().Unix()) {
		return "", errors.New("token was expired")
	}
	fmt.Println(int64(claims["exp"].(float64)), time.Now().Unix())
	return claims["user_id"].(string), nil
	//return validatedToken, nil
}
