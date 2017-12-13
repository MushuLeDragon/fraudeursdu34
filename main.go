package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var apiURI = "localhost:3000"

// App is the FraudeursDu3-4 application
type App struct {
	Router *gin.Engine
}

// Init initializes and sets up the routes
func (app *App) initializeRouter() {
	app.Router = gin.Default()

	app.Router.StaticFile("/", "front/dist/index.html")
	app.Router.Static("/static", "front/dist/static")

	api := app.Router.Group("/api")

	trains := api.Group("/trains")
	{
		trains.GET("", app.reverseProxy(apiURI))
		trains.GET("/:id", app.reverseProxy(apiURI))
	}

	reservations := api.Group("/reservations")
	{
		reservations.GET("", app.reverseProxy(apiURI))
		reservations.POST("", app.reverseProxy(apiURI))
		reservations.DELETE("/:id", app.reverseProxy(apiURI))
	}
}

func (app *App) reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Scheme: "http",
			Host:   target,
		})

		proxy.Director = func(r *http.Request) {
			r.Host = target
			r.URL.Host = r.Host
			r.URL.Scheme = "http"
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// capture signals to a channel
	c := make(chan os.Signal, 2)

	// Both interruption signals are to be caught
	// It will prevent from quitting
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Goroutine to asynchronously check the signals
	go func() {
		// Browse signals
		<-c

		fmt.Println("Exiting !")

		// Actually exit
		os.Exit(1)
	}()

	// Read the env variable $PORT
	port := os.Getenv("PORT")

	// Default value for $PORT
	if port == "" {
		port = "8880"
	}

	api := os.Getenv("API_URL")

	if api != "" {
		apiURI = api
	}

	app := App{}

	app.initializeRouter()

	// Notifying start
	fmt.Printf("Listening at :%v\n", port)

	// Start the http server
	app.Router.Run(":" + port)
}
