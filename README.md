# gittk

A git toolkit with a focus on managing a multitude of repos.

## Usage

### View usage

Write usage to stdout

    $ gittk

### clone

Currently only supports github https and ssh URIs.

    $ gittk clone <GIT_URI>

The default projects directory is `~/projects`.  You can change this default with the
`GITTK_PATH` environment variable.

    $ GITTK_PATH=~/go/src clone <GIT_URI>

## Example tree

    projects
    └── github.com
        ├── majgis
        │   ├── change-log
        │   ├── gittk
        │   └── ngify
        ├── shuhei
        │   └── pelo
        └── willjk
            └── okta-saml-express

