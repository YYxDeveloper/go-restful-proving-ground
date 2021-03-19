package main
import (
	"github.com/YYxDeveloper/go-restful-proving-ground/app"
	"github.com/gin-gonic/gin"
	"net/http"

)
func main() {
	//apiServer := gin.New()
	//apiServer.GET("/hello", HelloWorld)
	//
	//apiServer.Run(":3388")
	app.StartApplication()
}

func HelloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, "Hello World")
}
