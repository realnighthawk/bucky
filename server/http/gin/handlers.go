package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthHandler(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusOK)
	_, err := ctx.Writer.Write([]byte("ok"))
	if err != nil {
		ctx.JSON(500, err)
		return
	}
}

func pingHandler(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusOK)
	_, err := ctx.Writer.Write([]byte("pong"))
	if err != nil {
		ctx.JSON(500, err)
		return
	}
}

func (srv *Server) statusHandler(ctx *gin.Context) {
	ctx.JSON(200, srv)
}
