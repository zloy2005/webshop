package storage

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/zloy2005/webshop/internal/common"
)

type Config struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Dbname   string
}

func (c *Config) Flags(name, prefix string) *pflag.FlagSet {
	f := pflag.NewFlagSet(name, pflag.PanicOnError)
	f.StringVar(&c.Host, "host", "127.0.0.1", "")
	f.Uint16Var(&c.Port, "port", 5432, "")
	f.StringVar(&c.User, "user", "root", "")
	f.StringVar(&c.Password, "password", "", "[secret]")
	f.StringVar(&c.Dbname, "dbname", "db", "")
	return common.MapWithPrefix(f, name, pflag.PanicOnError, prefix)
}

func (c *Config) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Dbname)
}
