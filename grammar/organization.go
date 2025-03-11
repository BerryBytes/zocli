package grammar

import "github.com/MakeNowJust/heredoc"

var OrganizationHelp = heredoc.Doc(`
	Manage and maintain different portions of organization
`)

var OrganizationGetHelp = heredoc.Doc(`
	Get orgranizations details

	NOTE: Your default organization i.e. of id 0 will not be
	shown, as it is used as default if no any context is set.
`)

var OrganizationGetExample = heredoc.Doc(`
	zocli org get
`)

var OrganizationUseDefaultHelp = heredoc.Doc(`
	use organization as default for performing tasks
	NOTE: you must have access to the organization.
`)

var OrganizationDeleteHelp = heredoc.Doc(`
	Delete the orgranization in which your active token reside on.
	NOTE: you must have access to the organization.
`)

var OrganizationMembersHelp = heredoc.Doc(`
	Manage and maintain different portions of organization members
`)

var OrganizationMembersGetHelp = heredoc.Doc(`
	Get orgranization members details

	NOTE: Your default organization members i.e. of id 0 will not be
	shown, as it is used as default if no any context is set.
`)

var OrganizationMemberGetExample = heredoc.Doc(`
	zocli member get
`)

var OrganizationMemberDeleteHelp = heredoc.Doc(`
	Delete the orgranization member in which your active token reside on.
	NOTE: you must have access to the organization.
`)

var OrganizationMemberGetHelp = heredoc.Doc(`
	Get orgranizations members details

	NOTE: Your default organization members i.e. of id 0 will not be
	shown, as it is used as default if no any context is set.
`)

var ClusterHelp = heredoc.Doc(`
	Manage the clusters on organizations.

	Note: You must be admin to use this command.
`)

var ClusterImportHelp = heredoc.Doc(`
	Import clusters from different providers on organizations.
	You can provide two files for this command, in which one
	contains the information regarding the cluster creation and
	next one will be the actual file provided by your provider.

	Note: You must be admin to use this command.
`)

var ClusterGetHelp = heredoc.Doc(`
	List clusters imported or created on organization.

	Note: You must be admin to use this command.
`)

var PackagesHelp = heredoc.Doc(`
	Manage the packages on clusters.
`)

var PackagesInstallHelp = heredoc.Doc(`
	Install packages on clusters.
	Standard templates can also be applied to install packages.
	To show the templates and exit use --show-templates flag.
`)

var PackagesUnInstallHelp = heredoc.Doc(`
	UnInstall packages on clusters.
`)

var PackagesStatusHelp = heredoc.Doc(`
	Get the status for the packages or single package.
`)
