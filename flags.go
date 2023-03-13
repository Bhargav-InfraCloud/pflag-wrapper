package pflagwrapper

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/spf13/pflag"
)

var ErrHelp = pflag.ErrHelp

type Manager struct {
	BaseFlagSet *pflag.FlagSet
	Usage       []string

	help *bool
}

func NewManager(flagSetName string) *Manager {
	fs := pflag.NewFlagSet(flagSetName, pflag.ContinueOnError)
	return &Manager{
		BaseFlagSet: fs,
		help:        fs.BoolP("help", "h", false, "Print help"),
	}
}

func (m *Manager) AddFlagSet(fs *FlagSet, flags ...FlagDetails) *Manager {
	for _, flag := range flags {
		flag.addToFlagSet(fs)
	}

	m.BaseFlagSet.AddFlagSet(fs.pflagSet)
	m.Usage = append(m.Usage, formatUsage(fs, flags...))

	return m
}

func (m *Manager) Parse(args []string) error {
	err := m.BaseFlagSet.Parse(args)
	if err != nil {
		return fmt.Errorf("failed to parse input flags: %w", err)
	}

	if *m.help {
		m.printUsage(os.Stdout)

		return ErrHelp
	}

	return nil
}

func (m *Manager) printUsage(out io.Writer) {
	w := tabwriter.NewWriter(out, 0, 30, 10, '\t', tabwriter.TabIndent)

	fmt.Fprintf(w, "%s\n", m.Usage)

	w.Flush()
}
