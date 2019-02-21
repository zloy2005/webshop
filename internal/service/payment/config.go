package payment

import (
	"github.com/spf13/pflag"
	"github.com/zloy2005/webshop/internal/common"
)

type Config struct {
	Merchant string
	Key      string
	Currency string
}

func (c *Config) Flags(name, prefix string) *pflag.FlagSet {
	f := pflag.NewFlagSet(name, pflag.PanicOnError)
	f.StringVar(&c.Merchant, "merchant", "", "")
	f.StringVar(&c.Key, "key", "", "")
	return common.MapWithPrefix(f, name, pflag.PanicOnError, prefix)
}
