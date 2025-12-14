package handlers

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/lexsos/home-proxy/internal/auth"
	"github.com/lexsos/home-proxy/internal/filters"
	"github.com/lexsos/home-proxy/internal/utils/logging"
	"github.com/lexsos/home-proxy/internal/utils/request"
	"github.com/lexsos/home-proxy/internal/utils/response"
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
	reqId := request.GetOrGenId(r)
	ctx, logger := logging.WithFields(r.Context(), log.Fields{"src": r.RemoteAddr, "dst": r.Host, "reqId": reqId})
	account, err := proxy.authenticator.GetUser(ctx, r)
	if err != nil {
		logger.Warnf("Auth fail: %v", err)
		response.RequireAuth(w)
		return
	}
	if account == nil {
		response.RequireAuth(w)
		logger.Info("Account not found")
		return
	}
	ctx, logger = logging.WithField(ctx, "user", account.Login)

	dstDomain := request.GetDstDomain(r)
	allow, err := proxy.reqFilter.HasAccess(ctx, account.ProfileSlug, dstDomain)
	if err != nil {
		logger.Errorf("Filters fail: %v", err)
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
		handleTunnel(ctx, w, r)
	} else {
		handleHTTPProxy(w, r)
	}
}
