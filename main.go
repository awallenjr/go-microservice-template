package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from your Go microservice!")
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from your Go microservice!",
		})
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("anthony-allen.com"),
		Cache:      autocert.DirCache("cache/"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
