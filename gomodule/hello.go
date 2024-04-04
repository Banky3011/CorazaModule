package gomodule

import (
	"log"
	"sync"

	"github.com/corazawaf/coraza/v3"
	"github.com/gin-gonic/gin"
	// "net/http"
)

var (
	serverRunning bool
	serverLock    sync.Mutex
)

func MyWaf() gin.HandlerFunc {
	waf, err := coraza.NewWAF(coraza.NewWAFConfig().
		WithDirectivesFromFile("coraza.conf").
		WithDirectivesFromFile("coreruleset/rules/*.conf").
		WithDirectivesFromFile("coreruleset/crs-setup.conf.example"),
	)
	if err != nil {
		log.Fatalf("Error creating WAF: %v", err)
	}

	return func(c *gin.Context) {

		clientIp := c.ClientIP()

		log.Println("=== Client IP ===")
		log.Println(clientIp)
		log.Println("=================")

		//userPort := c.Request.URL.Port()

		tx := waf.NewTransaction()
		defer func() {
			tx.ProcessLogging()
			log.Println("Transaction closed successfully")
			tx.Close()
		}()

		tx.ProcessConnection("127.0.0.1", 55555, "127.0.0.1", 8081)

		//tx.ProcessURI("")
		tx.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")
		tx.AddRequestHeader("Content-Type", "text/plain")

		if it := tx.ProcessRequestHeaders(); it != nil {
			log.Printf("Transaction was interrupted with status %d\n", it.Status)
			c.AbortWithStatus(it.Status)
			return
		} else {
			log.Printf("Request Allowed")
		}

		if typeInterupt, it := tx.ProcessRequestBody(); typeInterupt != nil {
			log.Printf("Transaction was interrupted with status %d\n", typeInterupt.Status)
			log.Println(it.Error())
			return
		}

		c.Next()
	}
}

func RunServer() {
	serverLock.Lock()
	defer serverLock.Unlock()

	if serverRunning {
		log.Println("Server is already running")
		return
	}

	r := gin.Default()
	r.Use(MyWaf())

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	serverRunning = true
	log.Println("Server started")
}

func StopServer() {
	serverLock.Lock()
	defer serverLock.Unlock()

	if !serverRunning {
		log.Println("Server is not running")
		return
	}

	serverRunning = false
	log.Println("Server stopped")
}

func is_server_running() bool {
	serverLock.Lock()
	defer serverLock.Unlock()
	return serverRunning
}

func main() {
	RunServer()
}
