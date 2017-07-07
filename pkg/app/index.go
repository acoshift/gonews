package app

import (
	"net/http"

	"github.com/acoshift/gonews/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
