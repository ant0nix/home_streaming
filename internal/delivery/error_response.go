package delivery

import (
	"log"

	"github.com/gin-gonic/gin"
)

func newErrorResponce(c *gin.Context, mess string, statuscode int) {
	log.Println(mess)
	c.AbortWithStatusJSON(statuscode, mess)
}
