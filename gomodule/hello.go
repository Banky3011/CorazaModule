package gomodule

import (
	"C"
	"log"
	"strings"

	"github.com/corazawaf/coraza/v3"
)

type Request struct {
	RemoteAddr  string
	Path        string
	Port        int
	Query       string
	HTTPVersion string
	Headers     string
	Body        string
}

func MyWaf(req Request) int {
	waf, err := coraza.NewWAF(coraza.NewWAFConfig().
		WithDirectivesFromFile("coraza.conf").
		WithDirectivesFromFile("coreruleset/rules/*.conf").
		WithDirectivesFromFile("coreruleset/crs-setup.conf.example"))

	if err != nil {
		log.Fatalf("Error creating WAF: %v", err)
		return 500
	}

	tx := waf.NewTransaction()
	defer func() {
		tx.ProcessLogging()
		log.Println("Transaction closed successfully")
		tx.Close()
	}()

	tx.ProcessConnection(req.RemoteAddr, req.Port, "127.0.0.1", 8081)

	tx.ProcessURI(req.Path, "?", req.HTTPVersion)

	headersMap := make(map[string]string)
	headers := strings.Split(req.Headers, ",")
	for _, header := range headers {
		parts := strings.Split(header, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headersMap[key] = value
		}
	}

	for name, value := range headersMap {
		tx.AddRequestHeader(name, value)
	}

	if it := tx.ProcessRequestHeaders(); it != nil {
		log.Printf("(Phase 1) Transaction was interrupted with status %d\n", it.Status)
		log.Printf("Interrupted Request Details:\n")
		log.Printf("Remote Address: %s\n", req.RemoteAddr)
		log.Printf("Request URI: %s\n", req.Path)
		log.Printf("HTTP Version: %s\n", req.HTTPVersion)
		log.Printf("Request Headers:\n%s\n", req.Headers)
		log.Printf("Request Body:\n%s\n", req.Body)
		return it.Status
	}

	if it, err := tx.ProcessRequestBody(); it != nil {
		log.Printf("Transaction was interrupted in phase 2 with status %d\n", it.Status)
		log.Print(err)
		return it.Status
	}

	log.Printf("Request Allowed")
	return 200
}
