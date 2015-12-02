package main
import "github.com/gin-gonic/gin"

func main() {
	rtr := gin.Default()
	rtr.Static("/files", "./files")
	rtr.Run(":8020")
}
