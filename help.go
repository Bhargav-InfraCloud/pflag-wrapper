package pflagwrapper

import (
	"fmt"
	"strings"
)

func formatUsage(fs *FlagSet, flags ...FlagDetails) string {
	var usage strings.Builder

	for _, flagDetails := range flags {
		var (
			name        = flagDetails.flagName()
			flag        = fs.pflagSet.Lookup(name)
			short       = flag.Shorthand
			desc        = flag.Usage
			defaultVal  = flag.DefValue
			placeholder = flagDetails.placeholder()
		)

		if short != "" {
			usage.WriteString(fmt.Sprintf("-%s,\t", short))
		} else {
			usage.WriteString(strings.Repeat(` `, 4) + "\t")
		}

		usage.WriteString(fmt.Sprintf("--%s\t", name))

		if placeholder != "" {
			usage.WriteString(fmt.Sprintf(" %s\t", placeholder))
		}

		usage.WriteString(fmt.Sprintf("%s\t", desc))

		if !flagDetails.isDefaultEmpty() {
			usage.WriteString(fmt.Sprintf(" (default: %s)", defaultVal))
		}

		usage.WriteRune('\n')
	}

	return usage.String()
}
