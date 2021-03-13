package auditor

import (
	"context"
	"log"

	"github.com/dogument/elenktis"
)

type elenktisAuditor struct {
}

func (elenktisAuditor) Create(ctx context.Context, alog elenktis.AuditLog) error {
	log.Println("created the audit log")
	return nil
}

func New(cfg *elenktis.Config) elenktis.Auditor {
	return elenktisAuditor{}
}
