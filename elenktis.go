package elenktis

// Runner is the interface that defines the runner operation
type Runner interface {
	Run(*Config) error
}
