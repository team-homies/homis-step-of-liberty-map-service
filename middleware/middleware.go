package middleware

import (
	"context"
	"errors"
	"main/app/grpc/proto/auth"
	"main/config"
	"main/constant/common"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// auth 토큰검증 공통 middleware
func AuthVerificationMiddleware(c *fiber.Ctx) error {
	// 0. grpc 연결 맺기
	var address string
	if viper.GetString(config.GRPC_AUTH_HOST) == "localhost" {
		address = viper.GetString(config.GRPC_AUTH_PORT)
	} else {
		address = viper.GetString(config.GRPC_AUTH_HOST) + viper.GetString(config.GRPC_AUTH_PORT)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	authClient := auth.NewAuthServiceClient(conn)

	// 1. header에서 token 추출
	tokenString := c.Get("X-Authorization")

	// 2. auth 토큰검증 grpc 호출
	// 3. response (id, message, ischecked)
	res, err := authClient.GetAuthVerification(context.Background(), &auth.AuthRequest{
		Token: tokenString,
	})
	if err != nil {
		return err
	}
	// 3-1. ischecked false면 return
	if !res.IsAuth {
		return errors.New("invalid token")
	}

	// 3-2. ischecked true면 locals에 id 저장
	c.Locals(common.LOCALS_USER_ID, res.UserId)

	// 4. c.Next()
	return c.Next()
}
