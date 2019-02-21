package server

import (
	"github.com/spf13/pflag"

	"github.com/zloy2005/webshop/internal/service/payment"
	"github.com/zloy2005/webshop/internal/storage"
	"github.com/zloy2005/webshop/web"
)

type Config struct {
	Storage    storage.Config
	Server     web.Config
	Payment    payment.Config
	ConfigFile string `yaml:"-"`
}

func (c *Config) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("webshop", pflag.PanicOnError)
	f.AddFlagSet(c.Storage.Flags("Storage", "storage"))
	f.AddFlagSet(c.Payment.Flags("Payment", "payment"))
	f.AddFlagSet(c.Server.Flags())
	f.StringVar(&c.ConfigFile, "config", "", "")
	return f
}
