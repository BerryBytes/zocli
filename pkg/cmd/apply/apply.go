package apply

import (
	"os"

	"github.com/berrybytes/zocli/grammar"
	"github.com/berrybytes/zocli/pkg/cmd/apply/requests"
	"github.com/berrybytes/zocli/pkg/middleware"
	"github.com/berrybytes/zocli/pkg/utils/factory"
	manifestprocessor "github.com/berrybytes/zocli/pkg/utils/manifestProcessor"
	"github.com/berrybytes/zocli/pkg/utils/manifestProcessor/models"
	"github.com/spf13/cobra"
)

type Opts struct {
	F         *factory.Factory
	File      string
	Request   requests.Interface
	Processor manifestprocessor.Interface
}

func NewApplyManifestCommand(f *factory.Factory) *cobra.Command {
	o := &Opts{F: f}
	o.Processor = manifestprocessor.New(f)
	o.Request = requests.NewInterface(f)
	manifest := &cobra.Command{
		Use:     "apply",
		Aliases: []string{"app", "appl", "ap"},
		Short:   "apply manifest file",
		Long:    grammar.ManifestHelp,
		Example: grammar.ManifestExample,
		GroupID: "basic",
		Run:     o.ManifestRunner,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				o.File = args[0]
			}
			middleware.LoggedIn(o.F)
		},
		DisableFlagsInUseLine: true,
	}

	manifest.Flags().StringVarP(&o.File, "file", "f", "", "manifest file")
	return manifest
}

// ManifestRunner
//
// First checks if the file name is supplied other than
// '-' and '$' symbol, as those two are commonly used in
// cli for piping the commands to the next.
// And if not then process the file by calling the
// GetFileData function which is present in manifestprocessor module
// If no any file information is present then, the command checks
// if data was provided through pipe command
func (o *Opts) ManifestRunner(cmd *cobra.Command, _ []string) {
	o.F.IO.StartProgressIndicator()
	if o.File != "" && o.File != "-" && o.File != "$" {
		o.MatchModel(o.Processor.GetFileData(o.File))
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		o.F.IO.StopProgressIndicator()
		o.F.Printer.Fatal(9, err.Error())
		return
	}
	if (info.Mode() & os.ModeCharDevice) != os.ModeCharDevice {
		o.MatchModel(o.Processor.ParseFromStdIn())
	}

	o.F.IO.StopProgressIndicator()
	_ = cmd.Help()
}

func (o *Opts) MatchModel(data interface{}, model string) {
	switch model {
	case "member":
		val, ok := data.(models.OrganizationMember)
		if ok {
			o.Request.CreateMember(&val)
		}
	case "project":
		val, ok := data.(models.Project)
		if ok {
			o.Request.CreateProject(&val)
		}
	case "organization":
		val, ok := data.(models.Organization)
		if ok {
			o.Request.CreateOrganization(&val)
		}
	case "environment":
		val, ok := data.(models.Env)
		if ok {
			o.Request.CreateEnv(&val)
		}
	case "application":
		val, ok := data.(models.App)
		if ok {
			o.Request.CreateApp(&val)
		}
	}
}
