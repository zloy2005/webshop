package common

import (
	"github.com/spf13/pflag"
)

func MapAll(f *pflag.FlagSet, name string, errorHandling pflag.ErrorHandling, fn func(*pflag.Flag)) *pflag.FlagSet {
	fNew := pflag.NewFlagSet(name, errorHandling)
	f.VisitAll(func(flag *pflag.Flag) {
		fn(flag)
		fNew.AddFlag(flag)
	})
	return fNew
}

func MapWithPrefix(f *pflag.FlagSet, name string, errorHandling pflag.ErrorHandling, prefix string) *pflag.FlagSet {
	return MapAll(f, name, errorHandling, func(flag *pflag.Flag) {
		if prefix != "" {
			flag.Name = prefix + "." + flag.Name
		}
	})
}

func MapWithHidden(f *pflag.FlagSet, name string, errorHandling pflag.ErrorHandling, hidden bool) *pflag.FlagSet {
	return MapAll(f, name, errorHandling, func(flag *pflag.Flag) {
		flag.Hidden = hidden
	})
}
