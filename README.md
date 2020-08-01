# Github-Changelog

## Description

Github-Changelog is a CLI tool that helps creating changelogs. It searches for merged pull requests of a repository,
reads their body and looks for a specified start string ( default: "Changelog:").

## Installation

### Windows 

- Download and unzip the package containing the binary
- The binary is ready for usage 

### Linux

- Download and unzip the package containing the binary
- Make the binary executable by using `chmod +x changelog`
- Sometimes it is required to change the permission `chmod 755 changelog`

## Usage

The general usage is:

`main [global options] command [command options] [arguments...]`

### Create

Create a new changelog.

`main create [command options] [arguments...]`

#### Required Options

- `--out value, -o value`         Path to output file
- `--repo value, -r value`        The repository name
- `--owner value, -O value`       The owner of the repository
- `--oauth value, -X value`       Personal access token used to perform actions ( Required for private repositories)

#### Other options

- `--sprint value, -s value`      Length of the sprint in days e.g. '-s 7' means last 7 days (default: 7)
- `--since value, -S value`       Start date of the sprint in the format 2006-01-02T15:04:05
- `--version value, -v value`     A version string can be provided that will appear in the heading of the changelog file
- `--tag value, -t value`         The string that tags the changelog texts

## Authentication

Besides using the `-oauth` flag to provide an access token it is possible to add the token inside a .env file.
To achieve that rename the `.env.example` file that comes with the binary to `.env` and adjust value of `OAUTH_TOKEN=<your_token_here>`. Using the .env file AND the `-oauth` flag, the flag value will be used inside the .env value.