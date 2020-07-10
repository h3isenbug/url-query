package url

import (
	"errors"
	"github.com/google/uuid"
	"github.com/h3isenbug/url-common/pkg/services/log"
	"github.com/h3isenbug/url-query/handlers"
	"github.com/h3isenbug/url-query/repositories"
	"github.com/h3isenbug/url-query/services/url"
	"net/http"
)

type QueryHandler interface {
	GetLongURL(w http.ResponseWriter, r *http.Request)
}

type QueryHandlerV1 struct {
	urlService url.Service
	log        log.LogService
}

func NewQueryHandlerV1(urlService url.Service, log log.LogService) QueryHandler {
	return &QueryHandlerV1{urlService: urlService, log: log}
}

func (handler QueryHandlerV1) GetLongURL(w http.ResponseWriter, r *http.Request) {
	var shortPath = handlers.GetURLParams(r)["short_path"]
	var userAgent = r.Header.Get("User-Agent")
	var etag = r.Header.Get("If-None-Match")
	if len(etag) == 0 {
		id, _ := uuid.NewUUID()
		etag = id.String()
	}

	longURL, err := handler.urlService.GetLongURL(shortPath, etag, userAgent)
	if errors.Is(err, repositories.ErrNotFound) {
		handlers.SendError(w, http.StatusNotFound)
		return
	}
	if err != nil {
		handlers.SendError(w, http.StatusInternalServerError)
		handler.log.Error("could not get longURL of %s: %s", shortPath, err.Error())
		return
	}

	w.Header().Set("ETag", etag)
	w.Header().Set("Location", longURL)
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("redirecting..."))
}
