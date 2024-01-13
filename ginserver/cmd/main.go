package main

import (
	"context"
	"entgo.io/ent/dialect"
	"errors"
	"flag"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	mtsgr "github.com/MehmetTalhaSeker/mts-gr"
	"github.com/MehmetTalhaSeker/mts-gr/ent"
	"github.com/MehmetTalhaSeker/mts-gr/ent/migrate"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var address string

	flag.StringVar(&address, "address", ":8081", "HTTP Server Address")
	flag.Parse()

	errC, err := run(address)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run(address string) (<-chan error, error) {
	entclient, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	if err := entclient.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	srv, err := newServer(serverConfig{
		Address:   address,
		EntClient: entclient,
	})

	if err != nil {
		return nil, errors.New("server creation failed")
	}

	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			entclient.Close()
			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(ctxTimeout); err != nil { //nolint: contextcheck
			errC <- err
		}

		fmt.Println("Shutdown completed")
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}

type serverConfig struct {
	Address   string
	EntClient *ent.Client
}

func newServer(conf serverConfig) (*http.Server, error) {
	router := gin.Default()
	router.POST("/query", graphqlHandler(conf.EntClient))
	router.GET("/", playgroundHandler())

	return &http.Server{
		Handler:           router,
		Addr:              conf.Address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}, nil
}

func graphqlHandler(client *ent.Client) gin.HandlerFunc {
	h := handler.NewDefaultServer(mtsgr.NewSchema(client))

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
