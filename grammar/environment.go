package grammar

import "github.com/MakeNowJust/heredoc"

var EnvironmentCommandHelp = heredoc.Doc(`
	Environment command helps you to move around Environment eco system
	for any project for 01cloud.
`)

var EnvironmentOverviewHelp = heredoc.Doc(`
	Get the usage overview of any of the environment.
	Note: You must have access to the environment
`)

var EnvironmentGetHelp = heredoc.Doc(`
	Get the environments residing on projects.
	NOTE: If any application is found on the context then, it will be used,
	if nothing is supplied
`)

var EnvironmentStopHelp = heredoc.Doc(`
	Stop any environment residing on projects.
	NOTE: If any application is found on the context then, it will be used,
	if nothing is supplied
`)

var EnvironmentStartHelp = heredoc.Doc(`
	Start any environment residing on projects.
	NOTE: If any application is found on the context then, it will be used,
	if nothing is supplied
`)

var EnvironmentDeleteHelp = heredoc.Doc(`
	Delete any environment residing on projects.
	NOTE: If any application is found on the context then, it will be used,
	if nothing is supplied
`)

var EnvironmentRenameHelp = heredoc.Doc(`
	Rename any environment residing on projects.
	NOTE: If any application is found on the context then, it will be used,
	if nothing is supplied
`)
