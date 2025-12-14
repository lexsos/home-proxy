package handlers

import (
	"io"
	"net"
	"net/http"
)

func handleHTTPProxy(w http.ResponseWriter, req *http.Request) {
	// Подменяем схему и адрес для http.Transport
	req.RequestURI = ""
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		req.Header.Set("X-Forwarded-For", clientIP)
	}

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// Копируем заголовки и тело ответа
	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
