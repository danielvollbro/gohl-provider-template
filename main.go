package main

import (
	"context"

	api "github.com/danielvollbro/gohl-api"
)

// TODO: Change provider name here
const (
	ProviderID   = "provider-template"
	ProviderName = "Template Provider"
)

func main() {
	// 1. Config Parsing
	// TODO: Read in environment variables you need
	// someConfig := os.Getenv("GOHL_CONFIG_MYVAR")

	// 2. Setup (Clients etc)
	ctx := context.Background()

	// 3. Run Analysis
	checks := runChecks(ctx)

	// 4. Report
	report := api.ScanReport{
		PluginID: ProviderID,
		Checks:   checks,
	}
	api.PrintReport(report)
}

func runChecks(ctx context.Context) []api.CheckResult {
	var checks []api.CheckResult

	checks = append(checks, api.CheckResult{
		ID:          "TMP-001",
		Name:        "Example Check",
		Description: "This is a template check",
		Passed:      true,
		Score:       10,
		MaxScore:    10,
		Remediation: "Nothing to do.",
	})

	return checks
}
