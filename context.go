package elenktis

import "context"

type ElenktisContext struct {
	context.Context
	Config *Config
}
