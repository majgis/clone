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

### -bash option

A boolean option to not execute operations, only write bash commands to STDOUT.

The motivation for outputing bash is to enable automatically
changing directory to the cloned repository.

To accomplish this task, use the following function:

    gittk(){
        eval "$(/usr/local/bin/gittk -bash $@)"
    }


To use only the clone command, use this function:

    clone(){
        eval "$(/usr/local/bin/gittk -bash clone $@)"
    }

Example of bash commands output:

    $ gittk clone https://github.com/majgis/ngify.git

    mkdir -p /home/mjackson/projects/github.com/majgis/ngify \
        && cd /home/mjackson/projects/github.com/majgis/ngify \
        && git clone https://github.com/majgis/ngify.git

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

