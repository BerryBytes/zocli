package cluster

import (
	"bytes"
	"mime/multipart"
	"strings"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/fs"
	"github.com/berrybytes/zocli/pkg/utils/requester/defaults"
	"github.com/berrybytes/zocli/pkg/utils/terminal"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func NewClusterImportCommand(o *Opts) *cobra.Command {
	importCluster := &cobra.Command{
		Use:     "import",
		Aliases: []string{"i", "im", "imp", "impo", "impor", "copy"},
		Short:   "cluster import from providers",
		Long:    grammar.ClusterImportHelp,
		GroupID: "basic",
		PreRun: func(cmd *cobra.Command, args []string) {
			middleware.LoggedIn(o.F)

			// verify if the user is on default organization i.e 0,
			// if so then do not proceed.
			// NOTE: the server does not have any validation for this, so
			// CLI will handle this
			if o.F.Config.ActiveContext.OrganizationID == 0 {
				o.F.Printer.Fatal(1, "cannot import cluster on default Organization")
				return
			}
		},
		Run:                   o.ImportRunner,
		DisableFlagsInUseLine: true,
	}

	importCluster.Flags().StringVarP(&o.ProviderName, "provider", "N", "", "name for the provider")
	importCluster.Flags().StringVarP(&o.ClusterName, "name", "n", "", "name for the cluster")
	importCluster.Flags().StringVarP(&o.Region, "region", "r", "", "region of the cluster")
	importCluster.Flags().StringVarP(&o.Zone, "zone", "z", "", "zone of the cluster")
	importCluster.Flags().StringArrayVarP(&o.Labels, "labels", "l", []string{}, "labels for the cluster")

	importCluster.Flags().StringVarP(&o.FileName, "mainfile", "f", "", "cluster creation data FILE name")
	importCluster.Flags().StringVarP(&o.ProviderFile, "providerfile", "p", "", "actual provider FILE file")
	importCluster.Flags().SortFlags = false
	return importCluster
}

func (o *Opts) checkAllValues() {
	term := terminal.New(o.F.UserOS)
	if o.ProviderName == "" {
		o.F.Printer.Print("Enter Provider Name : ")
		o.ProviderName, _ = term.ReadInput()
		o.ProviderName = strings.ReplaceAll(o.ProviderName, "\n", "")
	}
	if o.ClusterName == "" {
		o.F.Printer.Print("Enter Cluster Name : ")
		o.ClusterName, _ = term.ReadInput()
		o.ClusterName = strings.ReplaceAll(o.ClusterName, "\n", "")
	}
	if o.Region == "" {
		o.F.Printer.Print("Enter Region Name : ")
		o.Region, _ = term.ReadInput()
		o.Region = strings.ReplaceAll(o.Region, "\n", "")
	}
	if o.Zone == "" {
		o.F.Printer.Print("Enter Zone Name : ")
		o.Zone, _ = term.ReadInput()
		o.Zone = strings.ReplaceAll(o.Zone, "\n", "")
	}
}

func (o *Opts) ImportRunner(cmd *cobra.Command, _ []string) {
	if o.ProviderFile != "" {
		data, err := fs.LoadFile(o.ProviderFile)
		if err != nil {
			o.F.Printer.Fatal(10, err.Error())
		}
		o.ProviderFileData = data
	} else {
		o.F.Printer.Fatal(1, "Provider File must be provided")
		return
	}
	var newImport ClusterImport
	if o.FileName != "" {
		o.F.IO.StartProgressIndicator()
		data, err := fs.LoadFile(o.FileName)
		if err != nil {
			o.F.Printer.Fatal(10, err.Error())
		}
		err = yaml.Unmarshal(data, &newImport)
		if err != nil {
			o.F.Printer.Fatal(9, err.Error())
		}
	} else {
		o.checkAllValues()
		newImport = ClusterImport{ClusterName: o.ClusterName, ProviderName: o.ProviderName, Region: o.Region, Zone: o.Zone}
	}

	o.createCluster(&newImport)
	_ = cmd.Help()
}

func (o *Opts) createCluster(newImport *ClusterImport) {
	newImport.FileData = string(o.ProviderFileData)
	err := newImport.Validate()
	if err != nil {
		o.F.IO.StopProgressIndicator()
		o.F.Printer.Fatal(1, err)
		return
	}
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	_ = w.WriteField("provider", "other")
	_ = w.WriteField("provider_name", newImport.ProviderName)
	_ = w.WriteField("name", newImport.ClusterName)
	_ = w.WriteField("region", newImport.Region)
	_ = w.WriteField("zone", newImport.Zone)
	_ = w.WriteField("labels", strings.Join(newImport.Labels, ","))
	_ = w.WriteField("file", newImport.FileData)
	_ = w.Close()
	o.F.IO.StopProgressIndicator()
	headers := o.F.GetAuth()
	headers["Content-Type"] = w.FormDataContentType()
	reqConfig := defaults.CreateCluster(o.F, map[string]interface{}{"headers": headers, "body": body.Bytes()})
	reqConfig.Request()
	o.F.Printer.Print(o.F.IO.ColorScheme().SuccessIcon())
	o.F.Printer.Printf(" Successfully created cluster")
	o.F.Printer.Exit(0)
}
