package controllers

import "github.com/gin-gonic/gin"

func HashPassword(password string) string {
	return ""
}

func VerifyPssword(userPassword string, givenPassword string) (bool, string) {
	return false, ""

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {}

}

func ProductViwerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {}

}
func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {}

}
func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {}

}
