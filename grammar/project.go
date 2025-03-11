package grammar

import "github.com/MakeNowJust/heredoc"

var ProjectHelp = heredoc.Doc(`
	Project command helps you to move around project eco system for zocli.
`)

var ProjectGetHelp = heredoc.Doc(`
	Get the list of projects that you have access to.
	You can use this command to get list of projects,
	and even get detail about a single project.
`)

var ProjectGetExample = heredoc.Doc(`
	zocli project get <:id>
	zocli project get [-i id | -n name]
	zocli project get // to fetch all list
`)

var ProjectEnableHelp = heredoc.Doc(`
	Enable projects which has been disabled.
	Note: You must have access to the project
`)

var ProjectDisableHelp = heredoc.Doc(`
	Disable projects which is in enabled state.
	Note: You must have access to the project
`)

var ProjectRenameHelp = heredoc.Doc(`
	Rename projects which is in enabled state.
	Note: You must have access to the project
`)

var ProjectDeleteHelp = heredoc.Doc(`
	Delete project from your console.
	Note: You must have access to the project
`)

var ProjectOverviewHelp = heredoc.Doc(`
	Get the usage overview of any of the projects.
	Note: You must have access to the project
`)

var ProjectVarsHelp = heredoc.Doc(`
	List, update, add or delete any variables in the project scope
	Note: You must have access to the project
`)

var ProjectVarsGetHelp = heredoc.Doc(`
	List variables in the project scope
	Note: You must have access to the project
`)

var ProjectVarsDeleteHelp = heredoc.Doc(`
	Delete variables in the project scope
	Note: You must have access to the project
`)

var ProjectVarsUpdateHelp = heredoc.Doc(`
	Update project scoped variables
	Note: You must have access to the project
`)

var ProjectVarsUpdateExample = heredoc.Doc(`
	zocli project vars update <:projectId>
`)

var ProjectPermissionHelp = heredoc.Doc(`
	List, update and delete user permissions on projects.
	Note: You must have access to the project
`)

var ProjectPermissionListHelp = heredoc.Doc(`
	List user permissions on projects.
	Note: You must have access to the project
`)

var ProjectPermissionAddHelp = heredoc.Doc(`
	Add user permission to project.
	Note: You must have access to the project
`)

var ProjectPermissionListExamples = heredoc.Doc(`
	zocli project settings permission list <:id>
	zocli project settings permission list [-i id | -n name]
`)

var ProjectPermissionUpdateHelp = heredoc.Doc(`
	Update users permissions on projects.
	Note: You must have access and necessary permissions on the project
`)

var ProjectPermissionDeleteHelp = heredoc.Doc(`
	Delete users permissions on projects.
	Note: You must have access and necessary permissions on the project
`)

var ProjectLoadbalancerHelp = heredoc.Doc(`
	List, update and delete loadbalancers on projects.
	Note: You must have access to the project
`)

var ProjectLoadbalancerCreateHelp = heredoc.Doc(`
	Create loadbalancers on projects.
	Note: You must have access to the project
`)

var ProjectLoadbalancerListHelp = heredoc.Doc(`
	List loadbalancers on projects.
	Note: You must have access to the project
`)

var ProjectLoadbalancerDeleteHelp = heredoc.Doc(`
	Delete loadbalancers on projects.
	Note: You must have access to the project
`)

var ProjectUseDefaultHelp = heredoc.Doc(`
	Set any of the project to be used as a default one, which will help
	you to work with other commands without supplying project ID.

	Every organization context are saved as separate, which will help
	to work with many organization by maintaining different contexts.
`)

var ProjectUseDefaultExample = heredoc.Doc(`
	zocli project use <:projectID/:projectName>
`)

var ProjectGetDefaultHelp = heredoc.Doc(`
	Get the default project which is currently in use from the
	configuration file.
	For this, only the active orgranization's default project will
	be retrieved and printed.
`)

var ProjectDeleteDefaultHelp = heredoc.Doc(`
	Remove the default project which is currently in use from the
	configuration file.
	For this, only the active orgranization's default project will
	be removed and none other profiles.
`)
