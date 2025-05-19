package control

import (
	"fmt"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/libs/dto"
	"github.com/curtis0505/bridge/libs/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

func (p *ControlHandler) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.logger.Info("event", "Authorization")
		claims, err := p.GetLoginUserInfo(c)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, types.NewResponseHeader(http.StatusUnauthorized, err))
			return
		}

		adminInfo, err := service.GetRegistry().AdminService().FindAdminInfo(c, bson.M{"email": claims.User})
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, types.NewResponseHeader(http.StatusUnauthorized, err))
			return
		}

		// SUPERUSER ONLY
		if adminInfo.Grade < 0xffffffff {
			c.Abort()
			c.JSON(http.StatusUnauthorized, types.NewResponseHeader(http.StatusUnauthorized, fmt.Errorf("permission required")))
			return
		}

		p.logger.Info("event", "Authorization", "email", claims.User, "admin", adminInfo.Username)

		c.Set("admin", adminInfo)
		c.Next()
	}
}

func (p *ControlHandler) GetLoginUserInfo(ctx *gin.Context) (*dto.AdminUserClaims, error) {
	token := p.GetAccessToken(ctx)
	claims := &dto.AdminUserClaims{}

	if _, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(clientState), nil
	}); err != nil {
		return nil, err
	}

	return claims, nil
}

func (p *ControlHandler) GetAccessToken(ctx *gin.Context) (token string) {
	bearToken := ctx.Request.Header.Get("Authorization")
	slice := strings.Split(bearToken, " ")
	if len(slice) > 1 {
		token = slice[1]
	}
	return
}

func (p *ControlHandler) GetAdminInfo(ctx *gin.Context) (dto.AdminInfo, error) {
	if account, ok := ctx.Get("admin"); ok {
		return account.(dto.AdminInfo), nil
	} else {
		return dto.AdminInfo{}, fmt.Errorf("not exist account")
	}
}
