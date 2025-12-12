package handlers

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/auth"
	"github.com/lexsos/home-proxy/filters"
	"github.com/lexsos/home-proxy/logging"
	"github.com/lexsos/home-proxy/request"
	"github.com/lexsos/home-proxy/response"
)

type HttpProxyHandler struct {
	authenticator auth.HttpAuthenticator
	reqFilter     *filters.RequestFilter
}

func NewProxyHandler(authenticator auth.HttpAuthenticator, reqFilter *filters.RequestFilter) *HttpProxyHandler {
	return &HttpProxyHandler{
		authenticator: authenticator,
		reqFilter:     reqFilter,
	}
}

func (proxy *HttpProxyHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, logger := logging.WithFields(r.Context(), log.Fields{"src": r.RemoteAddr, "dst": r.Host})
	account, err := proxy.authenticator.GetUser(r)
	if err != nil {
		logger.Warn("auth fail: ", err)
		response.RequireAuth(w)
		return
	}
	if account == nil {
		response.RequireAuth(w)
		logger.Info("Fail auth")
		return
	}
	_, logger = logging.WithField(ctx, "user", account.Login)

	dstDomain := request.GetDstDomain(r)
	allow, err := proxy.reqFilter.HasAccess(account.ProfileSlug, dstDomain)
	if err != nil {
		logger.Error("Filters fail:", err)
		response.InternalError(w)
		return
	}
	if !allow {
		response.DomainForbidden(w, dstDomain)
		logger.Info("Deny access")
		return
	}

	logger.Info("New connection")
	if strings.ToLower(r.Method) == "connect" {
		handleTunnel(w, r)
	} else {
		handleHTTPProxy(w, r)
	}
}
