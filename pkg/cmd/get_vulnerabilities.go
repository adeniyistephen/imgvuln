package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/adeniyistephen/imgvuln/pkg/apis/imgvulnsecurity/v1alpha1"
	"github.com/adeniyistephen/imgvuln/pkg/imgvuln"
	"github.com/adeniyistephen/imgvuln/pkg/vulnerabilityreport"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewGetVulnerabilitiesCmd(executable string, cf *genericclioptions.ConfigFlags, outWriter io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Aliases: []string{"vulns", "vuln"},
		Use:     "vulnerabilities (NAME | TYPE/NAME)",
		Short:   "Get vulnerabilities report",
		Long: `Get vulnerabilities report for the specified workload
TYPE is a Kubernetes workload. Shortcuts and API groups will be resolved, e.g. 'po' or 'deployments.apps'.
NAME is the name of a particular Kubernetes workload.
`,
		Example: fmt.Sprintf(`  # Get vulnerabilities for a Deployment with the specified name
  %[1]s get vulnerabilities.aquasecurity.github.io deploy/nginx
  # Get vulnerabilities for a Deployment with the specified name in the specified namespace
  %[1]s get vulnerabilities deploy/nginx -n staging`, executable),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			kubeConfig, err := cf.ToRESTConfig()
			if err != nil {
				return err
			}
			scheme := imgvuln.NewScheme()
			kubeClient, err := client.New(kubeConfig, client.Options{Scheme: scheme})
			if err != nil {
				return err
			}
			ns, _, err := cf.ToRawKubeConfigLoader().Namespace()
			if err != nil {
				return err
			}
			mapper, err := cf.ToRESTMapper()
			if err != nil {
				return err
			}
			workload, _, err := WorkloadFromArgs(mapper, ns, args)
			if err != nil {
				return err
			}

			reader := vulnerabilityreport.NewReadWriter(kubeClient)
			items, err := reader.FindByOwnerInHierarchy(ctx, workload)
			if err != nil {
				return fmt.Errorf("list vulnerability reports: %v", err)
			}
			if len(items) == 0 {
				fmt.Fprintf(outWriter, "No reports found in %s namespace.\n", workload.Namespace)
				return nil
			}

			format := cmd.Flag("output").Value.String()
			printer, err := genericclioptions.NewPrintFlags("").
				WithTypeSetter(imgvuln.NewScheme()).
				WithDefaultOutput(format).
				ToPrinter()
			if err != nil {
				return fmt.Errorf("create printer: %v", err)
			}

			if err := printer.PrintObj(&v1alpha1.VulnerabilityReportList{Items: items}, outWriter); err != nil {
				return fmt.Errorf("print vulnerability reports: %v", err)
			}

			return nil
		},
	}

	return cmd
}