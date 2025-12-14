package handlers

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/lexsos/home-proxy/internal/utils/logging"
)

// handleTunnel обрабатывает HTTPS CONNECT запросы (TLS/SSL)
func handleTunnel(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	logger := logging.LogFromContext(ctx)
	destConn, err := net.DialTimeout("tcp", req.Host, 10*time.Second)
	if err != nil {
		http.Error(w, "Failed to connect to destination", http.StatusServiceUnavailable)
		return
	}
	defer destConn.Close()
	// Ответить 200 клиенту (говорит, что соединение установлено)
	w.WriteHeader(http.StatusOK)

	// Получить подложенное соединение с клиентом
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		logger.Warnf("Hijack error: %v", err)
		return
	}
	defer clientConn.Close()

	// Прокидывать данные между клиентом и сервером (bidirectional copy)
	go io.Copy(clientConn, destConn)
	io.Copy(destConn, clientConn)
	logger.Info("Close connection")
}
