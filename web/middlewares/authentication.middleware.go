package middlewares

import (
	"errors"
	"fmt"
	"kurdi-go/web/resources"
	"kurdi-go/web/utils"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthenticationMiddleware(a *fiber.App) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {

		if len(c.Get("Authorization")) == 0 {
			utils.Logger().Info(errors.New("unauthorized access").Error())
			response := resources.UnAuthorized("GENERAL.API_TOKEN_REQUIRED")
			c.Status(response.GetStatus()).JSON(response.GetData())
			return
		}
		tokenString := strings.Replace(c.Get("Authorization"), "Bearer ", "", -1)
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				utils.Logger().Info("Unexpected signing method")
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
			return hmacSampleSecret, nil
		})

		if token == nil {
			response := resources.UnAuthorized("GENERAL.INVALID_API_TOKEN")
			c.Status(response.GetStatus()).JSON(response.GetData())
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			UserId, _ := claims["sub"]

			if int(UserId.(float64)) != UserId {
				utils.Logger().Info(errors.New(", Invalid API token").Error())
				response := resources.UnAuthorized("INVALID_API_TOKEN")
				c.Status(response.GetStatus()).JSON(response.GetData())
				return
			}

			// Get nowUnixTime time.
			nowUnixTime := time.Now().Unix()

			// Get claims from JWT.
			claims, err := utils.ExtractTokenMetadata(c)
			if err != nil {
				// Return status 500 and JWT parse error.
				response := resources.UnAuthorized("LOGOUT_EXPIRATION_TOKEN")
				c.Status(response.GetStatus()).JSON(response.GetData())
			}

			// Set expiration time from JWT data of current book.
			expires := claims.Expires

			// Checking, if now time greather than expiration from JWT.
			if nowUnixTime > expires {
				// Return status 401 and unauthorized error message.
				response := resources.UnAuthorized("LOGOUT_EXPIRATION_TOKEN")
				c.Status(response.GetStatus()).JSON(response.GetData())
			}
			c.Set("user_id", strconv.FormatFloat(UserId.(float64), 'E', -1, 64))
		}
		c.Next()
	}
}
