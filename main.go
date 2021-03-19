package main
import (
	"github.com/gin-gonic/gin"
	"net/http"

)
func main() {
	//apiServer := gin.New()
	//apiServer.GET("/hello", HelloWorld)
	//
	//apiServer.Run(":3388")
}

func HelloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, "Hello World")
}
