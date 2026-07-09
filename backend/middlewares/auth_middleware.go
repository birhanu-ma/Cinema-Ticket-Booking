package middlewares

import (
    "errors"           
    "net/http"         
    "strings"          
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5" 
	"cinema-ticket-booking/backend/utils"

)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authorization header required",
            })
            c.Abort()
            return
        }
        
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid authorization header format",
            })
            c.Abort()
            return
        }
        
        tokenString := parts[1]
        claims := &utils.Claims{}
        token, err := jwt.ParseWithClaims(
            tokenString,
            claims,
            func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, jwt.ErrTokenSignatureInvalid
                }
                return utils.JwtSecret, nil
            },
        )
        
      if err != nil || !token.Valid {
    println("JWT Validation Debug Error:", err.Error()) 
    
    c.JSON(http.StatusUnauthorized, gin.H{
        "error": "Invalid or expired token",
    })
    c.Abort()
    return 
}
        c.Set("User_id", claims.UserID)
        c.Set("email", claims.Email)
        c.Set("role", claims.Role)
        c.Next()
    }
}

func RequireRole(allowedRoles ...string) gin.HandlerFunc{
	return func(c *gin.Context){
		role, exists:=c.Get("role")
		if !exists{
			c.JSON(http.StatusForbidden, gin.H{
				"error":"No role information found",
			})
			c.Abort()
			return
		}
		userRole:=role.(string)
		for _, allowedRole:=range allowedRoles{
			if userRole==allowedRole{
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{
			"error":"Insufficient permissions",
		})
		c.Abort()
	}
}

func getCurrentUser(c *gin.Context)(uint, string,error){
	userID, exists:=c.Get("User_id")
	if !exists{
		return 0,"", errors.New("user not authenticated")
	}
	email, exists :=c.Get("email")
	if !exists{
		return 0, "", errors.New("email not found in context")
	}
	return userID.(uint), email.(string), nil
}
