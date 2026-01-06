package main

import (
	"log"
	"net/http"
	"os"

	httpHandlers "webhooks-service/internal/http"
	"webhooks-service/internal/store"
)

func main() {
	mux := http.NewServeMux()
	var storeImpl store.WebhookStore

	webhookHandler := httpHandlers.NewWebhookHandler(
		storeImpl,
		[]byte(os.Getenv("WEBHOOK_SECRET")),
	)
	mux.HandleFunc("/webhooks/", webhookHandler.HandleWebhook)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}