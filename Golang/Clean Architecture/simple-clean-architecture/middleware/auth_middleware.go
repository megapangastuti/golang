package middleware

import (
	"net/http"
	"strings"

	"simple-clean-architecture/model"

	"simple-clean-architecture/utils/service"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization" binding:"required"`
}

// Versi tanpa menggunakan role
// func AuthMiddleware(ctx *gin.Context) {
// 	var aH authHeader
// 	if err := ctx.ShouldBindHeader(&aH); err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
// 		return
// 	}

// 	token := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", 1)
// 	tokenClaim, err := service.VerifyToken(token)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
// 		return
// 	}

// 	ctx.Set("user", model.UserCredential{Id: tokenClaim.ID, Role: tokenClaim.Role})
// 	ctx.Next()
// }

// Versi dengan pengecekan role
func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var aH authHeader
		if err := ctx.ShouldBindHeader(&aH); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		token := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", 1)
		tokenClaim, err := service.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		ctx.Set("user", model.UserCredential{Id: tokenClaim.ID, Role: tokenClaim.Role})
		validRole := false
		for _, role := range roles {
			if role == tokenClaim.Role {
				validRole = true
				break
			}
		}

		if !validRole {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden Resource"})
			return
		}
		ctx.Next()
	}
}
