package midleware
import (
	"TaskDo/app/api/schemas"
	"TaskDo/app/core"
	"net/http"
	"github.com/gin-gonic/gin"
)


func AuthenticateMiddleware(c *gin.Context) {
	tokenstring := c.GetHeader("Authorization")
	if tokenstring == "" {
		schemas.ResponseJSON(c, http.StatusBadRequest, "Token is required", nil)
		c.Abort()
		return
	}
	token := tokenstring[7:]
	claims, err := security.VarifyToken(token)
	if err != nil {
		schemas.ResponseJSON(c, http.StatusUnauthorized, "Invalid or expired token", err)
		c.Abort()
		return
	}
	c.Set("UserID", claims["sub"])
	c.Next()
}