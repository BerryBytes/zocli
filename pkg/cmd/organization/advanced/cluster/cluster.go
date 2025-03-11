package cluster

import (
	"errors"
	"regexp"

	"github.com/berrybytes/zocli/api"
	"github.com/berrybytes/zocli/grammar"
	packages "github.com/berrybytes/zocli/pkg/cmd/organization/advanced/cluster/packages"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	"github.com/spf13/cobra"
)

type Opts struct {
	F *factory.Factory
	// actual data for creating cluster
	FileName string
	// provider file path
	ProviderFile string
	// actual data of the provider either json or yaml any thing will be
	// saved as byte, and will be sent to the server
	ProviderFileData []byte

	ProviderName string
	ClusterName  string
	Region       string
	Zone         string
	Labels       []string

	Output   string
	Clusters []api.ClusterList
}

// NOTE: This is actively developed for only Other Provider
type ClusterImport struct {
	ProviderName string `yaml:"provider_name" json:"provider_name"`
	ClusterName  string `yaml:"name" json:"name"`

	Region string `yaml:"region" json:"region"`
	Zone   string `yaml:"zone" json:"zone"`

	Registry string   `yaml:"-"`
	Labels   []string `yaml:"labels" json:"labels"`

	FileData string `yaml:"file" json:"file"`
}

func (c *ClusterImport) Validate() error {
	matched, err := regexp.MatchString("^[a-z0-9-]+$", c.ClusterName)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("Cluster Name is compulsory and is allowed only lowercase characters, numbers and hyphen")
	}

	matched, err = regexp.MatchString("^[a-z0-9-]+$", c.ProviderName)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("Provider Name is compulsory and is allowed only lowercase characters, numbers and hyphen")
	}

	matched, err = regexp.MatchString("^[a-z0-9-]+$", c.Region)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("Region Name is compulsory and is allowed only lowercase characters, numbers and hyphen")
	}

	return nil
}

func NewClusterCommand(f *factory.Factory) *cobra.Command {
	o := Opts{F: f}
	cluster := &cobra.Command{
		Use:     "cluster",
		Aliases: []string{"c", "cl", "clu", "clus", "clust", "cluste"},
		Short:   "cluster commands",
		Long:    grammar.ClusterHelp,
		GroupID: "admin",
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	cluster.AddGroup(&cobra.Group{
		ID:    "basic",
		Title: "Basic commands to maintain cluster",
	})

	cluster.AddCommand(NewClusterImportCommand(&o))
	cluster.AddCommand(NewClusterGetCommand(&o))
	cluster.AddCommand(packages.NewPackagesCommand(o.F))
	return cluster
}
