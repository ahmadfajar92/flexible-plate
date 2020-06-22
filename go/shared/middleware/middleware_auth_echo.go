package middleware

import (
	"fmt"
	"net/http"
	"scaffold/shared"
	"scaffold/shared/interfaces"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// JWTData struct
type JWTData struct {
	Admin       bool   `json:"adm"`
	DeviceID    string `json:"did"`
	DeviceLogin string `json:"dli"`
	jwt.StandardClaims
}

// Authorization func
func Authorization(application interfaces.Application) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			token := strings.Split(authorization, " ")

			if len(token) < 2 || strings.ToLower(token[0]) != `bearer` {
				response := shared.JSONResponse(
					http.StatusUnauthorized,
					"Invalid Token",
					false,
					nil,
				)
				return c.JSON(response.Code, response)
			}

			_, err := jwtDecode(token[1])
			if err != nil {
				response := shared.JSONResponse(
					http.StatusUnauthorized,
					err.Error(),
					false,
					nil,
				)

				return c.JSON(response.Code, response)
			}

			return next(c)
		}
	}
}

func jwtDecode(tokenStr string) (*JWTData, error) {
	token, _ := jwt.ParseWithClaims(tokenStr, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
		t, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return t, nil
	})

	claims, _ := token.Claims.(*JWTData)

	return claims, nil
}
