package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingView represents the view for ping responses.
type PingView struct{}

// RenderPing renders the ping response.
func (v *PingView) RenderPing(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
