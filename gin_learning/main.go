package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	r := gin.Default()
	r.POST("/ping", func(c *gin.Context) {
		body:= struct {
			Message string `json:"message"`
		}{}
		err := c.BindJSON(&body)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(body.Message))
		c.JSON(http.StatusOK, gin.H{
			"message": body.Message,
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}