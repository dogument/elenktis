package app

import (
	"net/http"

	"github.com/dogument/elenktis"
	"github.com/gorilla/mux"
)

type elenktisRunner struct {
}

// Run runs the Elenktis server
func (elenktisRunner) Run(cfg *elenktis.Config) error {
	router := mux.NewRouter()
	router.Path("/v1/log").Methods(http.MethodPost).HandlerFunc(createAuditLog(cfg))
	err := http.ListenAndServe(":8080", router)
	return err
}

func New() elenktis.Runner {
	return elenktisRunner{}
}
