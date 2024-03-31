package gomodule

import (
	"github.com/corazawaf/coraza/v3"
    "github.com/gin-gonic/gin"
	"log"
	"fmt"
	// "net/http"
)

type Request struct {
	ip string
	port int
} 

func MyWaf(req Request) gin.HandlerFunc {

	println("In MyWaf Func")
	println(req.ip)


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
	
		tx.ProcessConnection(req.ip, req.port, "127.0.0.1", 12345)
    
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

func RunServer(req Request) {
	r := gin.Default()
	//log.Printf(req.ip)
	fmt.Printf("%v\n", req)
	r.Use(MyWaf(req))


	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func main() {

	req := Request{
		ip: "127.0.0.1",
		port: 8080,
	}

	RunServer(req)
}