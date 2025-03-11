package grammar

import "github.com/MakeNowJust/heredoc"

var ApplicationHelp = heredoc.Doc(`
	All application related commands that can be used to manage the apps.
`)

var ApplicationGetExample = heredoc.Doc(`
	To retrieve all application in a project use,
	zocli application get

	Or to get a single application use,
	zocli application get <:id>
`)

var ApplicationRenameHelp = heredoc.Doc(`
	Rename any application on a project.
	NOTE: You must have necessary permissions.
`)

var ApplicationDeleteHelp = heredoc.Doc(`
	Delete any application on a project.
	NOTE: You must have necessary permissions.
`)

var ApplicationGetHelp = heredoc.Doc(`
	Get details about any applications residing in a project.
`)

var ApplicationUseDefaultExample = heredoc.Doc(`
	zocli app use <:appID>
`)

var ApplicationUseDefaultHelp = heredoc.Doc(`
	Set Default application for the current context.
`)

var ApplicationGetDefaultHelp = heredoc.Doc(`
	Get the default app which is currently in use from the
	configuration file.
	For this, only the active orgranization's default application will
	be retrieved and printed.
`)

var ApplicationDeleteDefaultHelp = heredoc.Doc(`
	Remove the default app which is currently in use from the
	configuration file.
	For this, only the active orgranization's default app will
	be removed and none other profiles.
`)
