package cluster

import (
	"fmt"
	"time"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/utils"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	table "github.com/berrybytes/zocli/pkg/utils/tableprinter"
	"github.com/spf13/cobra"
)

// NewClusterGetCommand
//
// @param o
// @return getCluster
// Description: initializes a new command for getting clusters and returns it
func NewClusterGetCommand(o *Opts) *cobra.Command {
	getCluster := &cobra.Command{
		Use:                   "get",
		Aliases:               []string{"g", "ge", "list", "lis", "retrieve"},
		Short:                 "list all the clusters available on the organization",
		Long:                  grammar.ClusterGetHelp,
		GroupID:               "basic",
		Run:                   o.GetRunner,
		DisableFlagsInUseLine: true,
	}

	getCluster.Flags().StringVarP(&o.Output, "out", "o", "", "output in json or yaml")
	return getCluster
}

// GetRunner
//
// @param cmd
// @param args
// @return err
// Description: root runner for getting clusters
func (o *Opts) GetRunner(_ *cobra.Command, _ []string) {
	reqConfig := defaults.GetCluster(o.F, map[string]interface{}{"headers": o.F.GetAuth()})
	reqConfig.Method = "GET"
	reqConfig.Body = nil
	res := reqConfig.Request()
	err := utils.ConvertType(res.Data, &o.Clusters)
	if err != nil {
		o.F.Printer.Fatalf(9, "cannot unmarchal data")
		return
	}
	o.printClustersTable()
}

// printClustersTable
//
// Description: prints the clusters in a table
func (o *Opts) printClustersTable() {
	tableprinter := table.New(o.F, 0)
	color := o.F.IO.ColorScheme().ColorFromString("blue")
	tableprinter.HeaderRow(color, "ID", "Name", "Provider", "Region", "Active", "Created At")
	for _, cluster := range o.Clusters {
		tableprinter.AddField(fmt.Sprint(cluster.ID))
		tableprinter.AddField(cluster.Cluster.Name)
		tableprinter.AddField(cluster.ProviderName)
		tableprinter.AddField(cluster.Region)
		tableprinter.AddField(fmt.Sprint(cluster.Active))
		tableprinter.AddField(table.RelativeTimeAgo(time.Now(), cluster.Createdat))
		tableprinter.EndRow()
	}
	tableprinter.Print()
}
