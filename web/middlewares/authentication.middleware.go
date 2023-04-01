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

func AuthenticationMiddleware() func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {

		if len(ctx.Get("Authorization")) == 0 {
			utils.Logger().Info(errors.New("unauthorized access").Error())
			response := resources.UnAuthorized("GENERAL.API_TOKEN_REQUIRED")
			ctx.Status(response.GetStatus()).JSON(response.GetData())
			return errors.New("unauthorized access")
		}

		token, err := extractTokenFromAutoriztionHeader(ctx.Get("Authorization"))
		if err == nil {
			utils.Logger().Info(err.Error())
			response := resources.UnAuthorized("GENERAL.INVALID_API_TOKEN")
			ctx.Status(response.GetStatus()).JSON(response.GetData())
			return nil
		}
		if token == nil {
			response := resources.UnAuthorized("GENERAL.INVALID_API_TOKEN")
			ctx.Status(response.GetStatus()).JSON(response.GetData())
			return nil
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {

			isTokenExpired, err := isTokenExpired(ctx)

			if err != nil {
				utils.Logger().Info(err.Error())
				response := resources.UnAuthorized("INVALID_API_TOKEN")
				ctx.Status(response.GetStatus()).JSON(response.GetData())
				return nil
			}
			if isTokenExpired {
				// Return status 401 and unauthorized error message.
				response := resources.UnAuthorized("LOGOUT_EXPIRATION_TOKEN")
				ctx.Status(response.GetStatus()).JSON(response.GetData())
			}

			if ok := setTokenClaimsInTheContext(ctx, claims); !ok {
				utils.Logger().Info(errors.New(", Invalid API token").Error())
				response := resources.UnAuthorized("INVALID_API_TOKEN")
				ctx.Status(response.GetStatus()).JSON(response.GetData())
				return nil
			}
		}
		ctx.Next()
		return nil
	}
}

func extractTokenFromAutoriztionHeader(authorizationHeaderValue string) (*jwt.Token, error) {

	tokenString := strings.Replace(authorizationHeaderValue, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if token.Header["alg"] != "HS256" {
			utils.Logger().Info("Unexpected signing method")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			utils.Logger().Info("Unexpected signing method")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func isTokenExpired(ctx *fiber.Ctx) (isTokenEXpired bool, err error) {

	// Get nowUnixTime time.
	nowUnixTime := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(ctx)

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if nowUnixTime > expires {
		isTokenEXpired = true
	}
	return
}

func setTokenClaimsInTheContext(ctx *fiber.Ctx, claims jwt.MapClaims) (ok bool) {
	UserId, hasUserId := claims["user_id"]
	if !hasUserId {
		ok = hasUserId
		return
	}
	ctx.Set("user_id", strconv.FormatFloat(UserId.(float64), 'E', -1, 64))
	ok = true
	return
}
