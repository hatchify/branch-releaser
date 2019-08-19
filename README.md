# Branch releaser
Branch releaser will aid in releasing branches up the chain

## Installation
Installing by compilation is very straight forward. The following dependencies are required:
- Go
- Git

### Method 1 - Download and install
```bash
curl -s https://raw.githubusercontent.com/hatchify/branch-releaser/master/install | bash -s
```

### Method 2 - Bash script
```bash
#!/bin/bash
go get -u github.com/hatchify/branch-releaser;
```

### Options

| Option                 | Description                         |
| :--------------------  | :---------------------------------- |
| `--source branch`      | Source branch to merge from         |
| `--destination branch` | Destination branch to merge into    |
| `--r`                  | Whether or not the current request is for the current directory or its children  |

### Example

```shell script
$ branch-releaser --source beta --destination master --r 
```

#### Additional

Sometimes you should export PATH variable to your current bash environment to allow golang executable file work from any place of your bash

```shell script
export PATH="$PATH:$GOPATH/bin"
```