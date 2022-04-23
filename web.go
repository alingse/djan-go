package djan

import (
	"net/http"
)

func (d *djanAdmin) NewRest(prefix string) http.Handler {
	return &adminRest{
		djanAdmin: d,
		prefix:    prefix,
	}
}

type adminRest struct {
	*djanAdmin
	prefix string
}

func (d *adminRest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(r.URL.String()))
}
