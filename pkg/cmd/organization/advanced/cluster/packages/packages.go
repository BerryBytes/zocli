// File: packages.go
// Description: This file contains the install, unsinstall and status
// command for packages
package packages

import (
	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/spf13/cobra"
)

// Opts
// Description: options for packages command
type Opts struct {
	F *factory.Factory
	I Interface

	PackageName    string
	ClusterID      int
	PackagesStatus map[string]bool

	TrueStatus map[string]bool
	FalsStatus map[string]bool

	PackageGroupToInstall string
	PackageConfig         *api.PackageConfig

	ShowTemplates bool
}

// NewPackagesCommand
// @Description: return packages command
// @param f
// @return packagesCmd
// Description: returns the command which are used to manage packages
// on a cluster
func NewPackagesCommand(f *factory.Factory) *cobra.Command {
	o := Opts{F: f}
	o.I = NewInterface(&o)
	packagesCmd := &cobra.Command{
		Use:     "packages",
		Aliases: []string{"p", "pkg", "pkgs", "pk", "package", "pack", "pa", "pac", "pacs", "packs"},
		Short:   "manage packages on cluster",
		Long:    grammar.PackagesHelp,
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
		DisableFlagsInUseLine: true,
	}

	packagesCmd.AddCommand(NewPackageInstallCommand(&o))
	packagesCmd.AddCommand(NewPackageUnInstallCommand(&o))
	packagesCmd.AddCommand(NewPackageStatusCommand(&o))
	return packagesCmd
}
