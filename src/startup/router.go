package startup

import (
	"github.com/gin-gonic/gin"
)

//NewRouter creates a gin instance and
// returns it.
func NewRouter() *gin.Engine {
	router := gin.New()
	return router
}
