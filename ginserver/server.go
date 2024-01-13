package ginserver

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	mtsgr "github.com/MehmetTalhaSeker/mts-gr"
	"github.com/MehmetTalhaSeker/mts-gr/ent"
	"github.com/MehmetTalhaSeker/mts-gr/metastore"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewServer(ec *ent.Client, mt metastore.MetaStore) *http.Server {
	router := gin.Default()
	router.POST("/query", graphqlHandler(ec, mt))
	router.GET("/", playgroundHandler())

	return &http.Server{
		Handler:           router,
		Addr:              "localhost:8081",
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
}

func graphqlHandler(client *ent.Client, mt metastore.MetaStore) gin.HandlerFunc {
	h := handler.NewDefaultServer(mtsgr.NewSchema(client, mt))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("Analyzify", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
