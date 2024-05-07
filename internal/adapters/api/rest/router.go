package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Mux struct {
// 	mux *http.ServeMux
// }

// func NewMux() *Mux {

// 	linkHandler := NewLinkHandler()

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", linkHandler.Link)
// 	return &Mux{
// 		mux: mux,
// 	}
// }

// func (m *Mux) ListenAndServe(addr string) error {
// 	if m.mux == nil {
// 		return fmt.Errorf("mux is not initialized")
// 	}
// 	return http.ListenAndServe(addr, m.mux)
// }

type Router struct {
	linkHandler *linkHandler
}

func NewRouter(linkHandler *linkHandler) *Router {
	return &Router{
		linkHandler: linkHandler,
	}
}

func (r Router) Router() *gin.Engine {
	engine := gin.New()
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"applicationErrorCode": http.StatusText(http.StatusNotFound), "message": "page not found"})
	})
	engine.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"applicationErrorCode": http.StatusText(http.StatusMethodNotAllowed), "message": "method not allowed"})
	})
	engine.POST("/", func(ctx *gin.Context) {
		r.linkHandler.Short(ctx.Writer, ctx.Request)
	})
	engine.GET("/:id", func(ctx *gin.Context) {
		r.linkHandler.RedirectOriginal(ctx.Writer, ctx.Request)
	})
	return engine
}
