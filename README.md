# ZOCLI

**ZOCLI** is the official Command Line Interface (CLI) for managing and fulfilling tasks on the **01cloud** ecosystem. It empowers developers and DevOps professionals to perform operations efficiently without the need for a web interface.

ZOCLI leverages the powerful [Cobra](https://github.com/spf13/cobra) framework to build a robust and user-friendly command-line interface.

With ZOCLI, you can easily perform all the necessary operations available in our web application without relying on a web interface. It offers a seamless experience for developers, system administrators, or anyone who prefers working with the command line.

### Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)

### Features

- **Authentication & Configuration**: Securely log in and manage your session with various authentication methods, including Single Sign-On (SSO), Access Tokens, and Basic Credentials.

- **Resource Management**: Create, read, update, and delete resources such as applications, projects, and organizations directly from the command line.

- **Environment Management**: Manage environments associated with your applications, including starting, stopping, and retrieving details.

- **Context Management**: Easily switch between different contexts (projects and applications) to streamline your workflow.

- **Advanced Permissions**: Manage user permissions and variables within projects, ensuring that the right users have the right access.

- **Output Formatting**: Retrieve data in various formats (JSON, YAML) for integration with other tools and automation scripts.

- **Help & Aliases**: Access comprehensive help menus for all commands and utilize aliases for quicker command execution.

## Getting Started

To get started, install ZOCLI and authenticate using:

### Prerequisites

Ensure you have the following prerequisites before installing ZOCLI:
- [X] Supported OS: Linux, macOS, Windows
- [X] curl installed

### Installation

Check out the latest releases at [GitHub Releases](https://github.com/berrybytes/zocli/releases).

You can also install ZOCLI directly using the following command:

1. To install the latest version.
```bash
curl -sS https://raw.githubusercontent.com/BerryBytes/zocli/main/install.sh | bash
```

2. To install a specific version (e.g., `v0.0.2`):
```bash
curl -sS https://raw.githubusercontent.com/BerryBytes/zocli/main/install.sh | bash -s -- v0.0.2
```

### Usage

Start with `zocli --help` to get started. Here is an example of logging in:
```bash
zocli login
```

**NOTE: Manuals for using this application will be updated in the near future.**

### Commands
For detailed command information, you can browse the [Zocli Documentation](https://docs.01cloud.io/services/cli/quickstart_cli/).

### Examples
```bash
zocli login
```
```bash
zocli create project
```
```bash
zocli create app
```

### Contributing

We welcome contributions! Please see our [contributing guidelines](CONTRIBUTING.md) for more details.

## Releasing
To trigger a release, push a commit to `main` with `[release]` in the commit message (e.g., `git commit -m "Add feature [release]"`). The workflow will auto-increment the version, tag it, and create a draft release.
