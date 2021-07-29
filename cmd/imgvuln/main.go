package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adeniyistephen/imgvuln/pkg/params"
	"github.com/adeniyistephen/imgvuln/pkg/cmd"
	"github.com/adeniyistephen/imgvuln/pkg/imgvuln"
	"k8s.io/klog"

	// Load all known auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

var (
	// These variables are populated by GoReleaser via ldflags
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// main is the entrypoint of the imgvuln CLI executable command.
func main() {
	//parse arguments
	params.ParseArguments()

	defer klog.Flush()
	klog.InitFlags(nil)

	if err := cmd.Run(imgvuln.BuildInfo{
		Version:    version,
		Commit:     commit,
		Date:       date,
		Executable: executable(os.Args),
	}, os.Args, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func executable(args []string) string {
	if strings.HasPrefix(filepath.Base(args[0]), "kubectl-") {
		return "kubectl imgvuln"
	}
	return "imgvuln"
}