package djan

import (
	"net/http"
)

func (d *djanAdmin) NewHandler(prefix string) http.Handler {
	return &adminWeb{
		djanAdmin: d,
		prefix:    prefix,
	}
}

type adminWeb struct {
	*djanAdmin
	prefix string
}

func (d *adminWeb) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(r.URL.String()))
}
