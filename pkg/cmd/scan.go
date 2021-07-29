package cmd

import (
	"github.com/adeniyistephen/imgvuln/pkg/imgvuln"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewScanCmd(buildInfo imgvuln.BuildInfo, cf *genericclioptions.ConfigFlags) *cobra.Command {
	scanCmd := &cobra.Command{
		Use:     "scan",
		Aliases: []string{"generate"},
		Short:   "Manage security weakness identification tools",
	}
	scanCmd.AddCommand(NewScanVulnerabilityReportsCmd(buildInfo, cf))
	scanCmd.AddCommand(NewScanPolicyReportsCmd(buildInfo, cf))
	return scanCmd
}