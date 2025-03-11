package grammar

import "github.com/MakeNowJust/heredoc"

var AuthHelp = heredoc.Doc(`
	This command help you for authentication purposes.
	Choose any of the subcommands to proceed on.
`)

var WarnCreds = heredoc.Doc(`
	Warning: It is not recommended to use email and password.
`)

var AuthExamples = heredoc.Doc(`
	zocli auth --help
	zocli auth login -t
	zocli auth login -c
`)

var StatusHelp = heredoc.Docf(`
	Verifies and displays information about your authentication state.

	This command will test your authentication state using the saved token,
	and tries to fetch basic information.
`)

var LoginExample = heredoc.Doc(`
	zocli auth login ( for interactive mode )

	zocli auth login -b
	zocli auth login -t
	zocli auth login -c
`)

var LoginHelp = heredoc.Docf(`
			Authenticate with 01Cloud-host

			The default authentication mode is a token based authentication flow.
			You can also provide token on standard input using %[1]s--token%[1]s
			flag.

			Alternatively, you can also use email and password to authenticate
			with 01Cloud-host. You can also provide email and password on
			standard input using %[1]sinteractive mode%[1]s and providing %[1]s--cred%[1]s
			flag.

			Alternatively, zocli will use the authentication token found in
			environment variables. This method is most suitable for
			"headless" use of zocli such as in automation. See
			%[1]szocli help environment%[1]s for more info.
`, `"`)

var LogOutHelp = heredoc.Docf(`
	Remove authentication from this device.

	This command removes the authentication configuration on this device,
	and also invalidates the session for the token.
`)
