package gomodule

import (
	"github.com/corazawaf/coraza/v3"
    "github.com/gin-gonic/gin"
	"log"
	// "net/http"
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
		tx := waf.NewTransaction()
		defer func() {
			tx.ProcessLogging()
			log.Println("Transaction closed successfully")
			tx.Close()
		}()
	
		tx.ProcessConnection("127.0.0.1", 8080, "127.0.0.1", 12345)
	
		if it := tx.ProcessRequestHeaders(); it != nil {
			log.Printf("Transaction was interrupted with status %d\n", it.Status)
			c.AbortWithStatus(it.Status)
			return
		} else {
			log.Printf("Request Allowed")
		}
	
		c.Next()
	}
}

func RunServer() {
	r := gin.Default()
	r.Use(MyWaf())
	//r.GET("/", showAllowed)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func main() {
	RunServer()
}