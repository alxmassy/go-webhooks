package http

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"webhooks-service/internal/store"
)

type WebhookHandler struct {
	store  store.WebhookStore
	secret []byte
}

func NewWebhookHandler(store store.WebhookStore, secret []byte) *WebhookHandler {
	return &WebhookHandler{
		store:  store,
		secret: secret,
	}
}

func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	log.Println("webhook handler hit:", r.URL.Path)

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	eventIDHeader := r.Header.Get("X-Event-ID")
	if eventIDHeader == "" {
		http.Error(w, "Missing X-Event-ID header", http.StatusBadRequest)
		return
	}

	sigHeader := r.Header.Get("X-Signature")
	if sigHeader == "" {
		http.Error(w, "Missing X-Signature header", http.StatusBadRequest)
		return
	}

	if !verifySignature(h.secret, body, sigHeader) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func verifySignature(secret, payload []byte, signature string) bool {
	mac := hmac.New(sha256.New, secret)
	mac.Write(payload)
	expectedMAC := mac.Sum(nil)

	providedMAC, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	return hmac.Equal(providedMAC, expectedMAC)
}
