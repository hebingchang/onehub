package middlewares

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"onehub/models"
)

type authHeader struct {
	Token   string    `header:"Token"`
}

func AdminMiddleware(c *gin.Context) {
	h := authHeader{}

	if viper.GetString("RPCSecret") == "" {
		c.Next()
	} else {
		if err := c.ShouldBindHeader(&h); err == nil {
			token := h.Token
			if token == "" {
				token = c.Param("token")
			}
			hash := sha256.New()
			hash.Write([]byte(viper.GetString("RPCSecret")))
			bs := hash.Sum(nil)
			secret := fmt.Sprintf("%x", bs)
			if token == secret {
				c.Next()
			} else {
				c.JSON(401, models.ErrorJson{
					Success: false,
					Message: "Unauthorized",
				})
				c.Abort()
			}
		} else {
			c.JSON(401, models.ErrorJson{
				Success: false,
				Message: "Unauthorized",
			})
			c.Abort()
		}
	}
}

