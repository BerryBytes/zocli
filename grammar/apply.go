package grammar

import "github.com/MakeNowJust/heredoc"

var ManifestHelp = heredoc.Doc(`
	Apply manifest file to any of the kind supplied
`)

var ManifestExample = heredoc.Doc(`
	cat sample.yml | zocli apply
	zocli apply -f sample.yaml
	zocli apply sample.yaml
`)
