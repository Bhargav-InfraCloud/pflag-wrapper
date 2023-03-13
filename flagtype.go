package pflagwrapper

import (
	"time"
)

type FlagType interface {
	string | bool | int | time.Duration
}

type flag[T FlagType] struct {
	target     *T
	name       string
	shorthand  string
	defaultVal T
	desc       string
}

func NewFlag[T FlagType](target *T, name, shorthand, description string, defaultVal T) *flag[T] {
	return &flag[T]{
		target:     target,
		name:       name,
		shorthand:  shorthand,
		defaultVal: defaultVal,
		desc:       description,
	}
}

type FlagDetails interface {
	addToFlagSet(fs *FlagSet)
	flagName() string
	isDefaultEmpty() bool
	placeholder() string
}

func (f *flag[T]) flagName() string {
	return f.name
}

func (f *flag[T]) isDefaultEmpty() bool {
	switch any(f.defaultVal).(type) {
	case string:
		defaultVal := any(f.defaultVal).(string)
		return defaultVal == ""
	case bool:
		defaultVal := any(f.defaultVal).(bool)
		return !defaultVal
	case int:
		defaultVal := any(f.defaultVal).(int)
		return defaultVal == 0
	case time.Duration:
		defaultVal := any(f.defaultVal).(time.Duration)
		return defaultVal == 0
	}

	return false
}

func (f *flag[T]) placeholder() string {
	switch any(f.defaultVal).(type) {
	case string:
		return `string`
	case bool:
		return ``
	case int:
		return `int`
	case time.Duration:
		return `duration`
	}

	return ``
}

func (f *flag[T]) addToFlagSet(fs *FlagSet) {
	switch any(f.defaultVal).(type) {
	case string:
		defaultVal := any(f.defaultVal).(string)
		target := any(f.target).(*string)

		if f.shorthand == "" {
			fs.pflagSet.StringVar(target, f.name, defaultVal, f.desc)

			return
		}

		fs.pflagSet.StringVarP(target, f.name, f.shorthand, defaultVal, f.desc)
	case bool:
		defaultVal := any(f.defaultVal).(bool)
		target := any(f.target).(*bool)

		if f.shorthand == "" {
			fs.pflagSet.BoolVar(target, f.name, defaultVal, f.desc)

			return
		}

		fs.pflagSet.BoolVarP(target, f.name, f.shorthand, defaultVal, f.desc)
	case int:
		defaultVal := any(f.defaultVal).(int)
		target := any(f.target).(*int)

		if f.shorthand == "" {
			fs.pflagSet.IntVar(target, f.name, defaultVal, f.desc)

			return
		}

		fs.pflagSet.IntVarP(target, f.name, f.shorthand, defaultVal, f.desc)
	case time.Duration:
		defaultVal := any(f.defaultVal).(time.Duration)
		target := any(f.target).(*time.Duration)

		if f.shorthand == "" {
			fs.pflagSet.DurationVar(target, f.name, defaultVal, f.desc)

			return
		}

		fs.pflagSet.DurationVarP(target, f.name, f.shorthand, defaultVal, f.desc)
	}
}

// func (f *flag[T]) usage() {
// 	pflag.
// 	switch any(f.defaultVal).(type) {
// 	case string:
// 		defaultVal := any(f.defaultVal).(string)

// 		if f.shorthand == "" {
// 			pflag.String(f.name, defaultVal, f.desc)

// 			return
// 		}

// 		pflag.StringP(f.name, f.shorthand, defaultVal, f.desc)
// 	case bool:
// 		defaultVal := any(f.defaultVal).(bool)

// 		if f.shorthand == "" {
// 			pflag.Bool(f.name, defaultVal, f.desc)

// 			return
// 		}

// 		pflag.BoolP(f.name, f.shorthand, defaultVal, f.desc)
// 	case int:
// 		defaultVal := any(f.defaultVal).(int)

// 		if f.shorthand == "" {
// 			pflag.Int(f.name, defaultVal, f.desc)

// 			return
// 		}

// 		pflag.IntP(f.name, f.shorthand, defaultVal, f.desc)
// 	case time.Duration:
// 		defaultVal := any(f.defaultVal).(time.Duration)

// 		if f.shorthand == "" {
// 			pflag.Duration(f.name, defaultVal, f.desc)

// 			return
// 		}

// 		pflag.DurationP(f.name, f.shorthand, defaultVal, f.desc)
// 	}
// }
