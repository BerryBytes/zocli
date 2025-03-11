package grammar

import "github.com/MakeNowJust/heredoc"

var Version = heredoc.Docf(`
		zocli version: 0.0.1

		To update to the latest version, run:
			$ zocli update

		For more information about zocli, visit %[1]shttps://docs.01cloud.com/zocli%[1]s
`, `"`)
