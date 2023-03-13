package pflagwrapper_test

import (
	"testing"
	"time"

	pflagwrapper "github.com/Bhargav-InfraCloud/pflag-wrapper"
)

func TestManager_Parse(t *testing.T) {
	var (
		token, org string
		timeout    time.Duration
	)

	args := []string{"-t", "abc", "-o", "infracloudio"}
	manager := pflagwrapper.NewManager(`test`).
		AddFlagSet(
			pflagwrapper.NewFlagSet(`secrets`),
			pflagwrapper.NewFlag(&token, "token", "t", "Bearer token for GitHub authentication", ""),
		).
		AddFlagSet(
			pflagwrapper.NewFlagSet(`environments`),
			pflagwrapper.NewFlag(&org, "org", "o", "Organization name to fetch the repos list", ""),
			pflagwrapper.NewFlag(&timeout, "timeout", "", "HTTP timeout", 30*time.Second),
		)

	err := manager.Parse(args)
	if err != nil {
		t.Errorf("parsing failed: %v", err)
	}

	// for _, line := range manager.Usage {
	// 	fmt.Println(line)
	// }
}

func TestManager_Parse_Help(t *testing.T) {
	var (
		token, org string
		timeout    time.Duration
	)

	args := []string{"--help"}
	manager := pflagwrapper.NewManager(`test`).
		AddFlagSet(
			pflagwrapper.NewFlagSet(`secrets`),
			pflagwrapper.NewFlag(&token, "token", "t", "Bearer token for GitHub authentication", ""),
		).
		AddFlagSet(
			pflagwrapper.NewFlagSet(`environments`),
			pflagwrapper.NewFlag(&org, "org", "o", "Organization name to fetch the repos list", ""),
			pflagwrapper.NewFlag(&timeout, "timeout", "", "HTTP timeout", 30*time.Second),
		)

	err := manager.Parse(args)
	if err != nil {
		t.Errorf("parsing failed: %v", err)
	}
}
