package gomodule

import (
	"C"
	"io"
	"log"
	"strings"

	"github.com/corazawaf/coraza/v3"
)

type Request struct {
	RemoteAddr        string
	Path              string
	Port              int
	Query             string
	HTTPVersion       string
	Method            string
	Headers           string
	Body              string
	HeaderHost        string
	HeaderUserAgent   string
	HeaderContentType *string
}

func CorazaModule(req Request) int {
	waf, err := coraza.NewWAF(coraza.NewWAFConfig().
		WithDirectivesFromFile("coreruleset/crs-setup.conf").
		WithDirectivesFromFile("coreruleset/modsecurity.conf").
		WithDirectivesFromFile("coreruleset/rules/*.conf"))

	if err != nil {
		log.Fatalf("Error creating WAF: %v", err)
		return 500
	}

	log.Printf(req.HeaderUserAgent)

	tx := waf.NewTransaction()
	defer func() {
		tx.ProcessLogging()
		log.Println("Transaction closed successfully")
		tx.Close()
	}()

	tx.ProcessConnection(req.RemoteAddr, req.Port, "172.29.122.57", 80)

	tx.ProcessURI(req.Path, req.Method, req.HTTPVersion)

	tx.SetServerName(req.HeaderHost)

	tx.AddRequestHeader("host", req.HeaderHost)
	tx.AddRequestHeader("user-agent", req.HeaderUserAgent)
	tx.AddRequestHeader("method", req.Method)

	if req.HeaderContentType != nil {
		tx.AddRequestHeader("content-type", *req.HeaderContentType)
	}

	if it := tx.ProcessRequestHeaders(); it != nil {
		log.Printf("Transaction was interrupted with status %d\n", it.Status)
		return it.Status
	}

	if tx.IsRequestBodyAccessible() {
		if req.Body != "" {
			bodyReader := strings.NewReader(req.Body)
			it, _, err := tx.ReadRequestBodyFrom(bodyReader)
			if err != nil {
				log.Printf("Failed to append request body: %v", err)
				return 500
			}

			if it != nil {
				log.Printf("Transaction was interrupted with status %d\n", it.Status)
				return it.Status
			}

			rbr, err := tx.RequestBodyReader()
			if err != nil {
				log.Printf("Failed to get the request body: %v", err)
				return 500
			}

			var remainingBody strings.Builder
			_, err = io.Copy(&remainingBody, rbr)
			if err != nil {
				log.Printf("Failed to read the remaining request body: %v", err)
				return 500
			}

			req.Body = remainingBody.String()
			log.Printf("---- req.Body ----")
			log.Printf(req.Body)
			log.Printf("------------------")
		}
	}

	if it, err := tx.ProcessRequestBody(); it != nil {
		log.Printf("Transaction was interrupted with status %d\n", it.Status)
		log.Print(err)
		return it.Status
	}

	log.Printf("Request Allowed")
	return 200
}
