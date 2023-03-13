package pflagwrapper

import "github.com/spf13/pflag"

type FlagSet struct {
	pflagSet *pflag.FlagSet
	name     string
}

func NewFlagSet(name string) *FlagSet {
	return &FlagSet{
		pflagSet: pflag.NewFlagSet(name, pflag.ContinueOnError),
		name:     name,
	}
}
