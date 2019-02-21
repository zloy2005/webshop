package web

import (
	"github.com/spf13/pflag"
)

type Config struct {
	Port string
}

func (c *Config) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("server", pflag.PanicOnError)
	f.StringVar(&c.Port, "port", "8080", "")
	return f
}
