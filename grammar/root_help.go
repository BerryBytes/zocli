package grammar

import "github.com/MakeNowJust/heredoc"

var RootHelp = heredoc.Docf(`
	A CLI tool for 01cloud.com complete documentation is available at
	%[1]shttps://docs.01cloud.io/services/cli/%[1]s Using the CLI tool requires
	an 01cloud.com account, and using this tool you can manage your
	01cloud.com account and resources.
`, `"`)
