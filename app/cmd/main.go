package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vranystepan/k8s-training/internal/hcregistry"
	"github.com/vranystepan/k8s-training/pkg/logging"
)

var (
	letters    = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	repository = []string{}
)

func main() {
	// obtain configuration values from env
	configBreak := getenvBool("CONFIG_BREAK", false)
	configSleep := getenvBool("CONFIG_SLEEP", false)

	// perform special "operational" case - break at the very
	// beggining of the application so it's gonna block the
	// rolling update I hope
	if configBreak {
		os.Exit(1)
	}

	// perform special "operational" case - sleep for one minute
	// to simulate some startup tasks like migration, generation
	// etc.
	if configSleep {
		time.Sleep(time.Minute * 1)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(logging.Config))
	e.Use(middleware.Recover())

	// Application handlers
	e.GET("/", hello)
	e.GET("/fill", fill)

	// Healthchecks
	e.GET("/_health/ready", ready)
	e.GET("/_health/alive", alive)

	// Handlers to set probes state
	e.GET("/_health/set/notalive", setNotAlive)
	e.GET("/_health/set/alive", setAlive)
	e.GET("/_health/set/notready", setNotReady)
	e.GET("/_health/set/ready", setReady)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handlers
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func ready(c echo.Context) error {
	if hcregistry.Ready {
		return c.String(http.StatusOK, "ready")
	}
	return c.String(http.StatusInternalServerError, "not ready")
}

func alive(c echo.Context) error {
	if hcregistry.Alive {
		return c.String(http.StatusOK, "alive")
	}
	return c.String(http.StatusInternalServerError, "not alive")
}

// manipulate the internal state of healthchceck
func setNotAlive(c echo.Context) error {
	hcregistry.Alive = false
	c.Response().Header().Set("pod", os.Getenv("HOSTNAME"))
	return c.String(http.StatusOK, "ok")
}

func setAlive(c echo.Context) error {
	hcregistry.Alive = true
	c.Response().Header().Set("pod", os.Getenv("HOSTNAME"))
	return c.String(http.StatusOK, "ok")
}

func setNotReady(c echo.Context) error {
	hcregistry.Ready = false
	c.Response().Header().Set("pod", os.Getenv("HOSTNAME"))
	return c.String(http.StatusOK, "ok")
}

func setReady(c echo.Context) error {
	hcregistry.Ready = true
	c.Response().Header().Set("pod", os.Getenv("HOSTNAME"))
	return c.String(http.StatusOK, "ok")
}

func fill(c echo.Context) error {
	// add some garbage to the memory
	for i := 1; i <= 10000; i++ {
		repository = append(repository, randSeq(1024))
	}

	// return some information so students know how they're progressing
	size := uintptr(len(repository)) * reflect.TypeOf(repository).Elem().Size()
	c.Response().Header().Set("len", fmt.Sprintf("%d", len(repository)))
	c.Response().Header().Set("size", fmt.Sprintf("%d", size))
	c.Response().Header().Set("pod", os.Getenv("HOSTNAME"))

	return c.String(http.StatusOK, "ok")
}

func getenvBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return boolValue
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
