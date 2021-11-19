package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UseStatic(engine *gin.Engine) {
	engine.StaticFS("/static", http.Dir("./static"))
}
