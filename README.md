[![Go Report Card](https://goreportcard.com/badge/github.com/langered/gonedrive)](https://goreportcard.com/report/github.com/langered/gonedrive)
[![Coverage Status](https://coveralls.io/repos/github/langered/gonedrive/badge.svg?branch=master)](https://coveralls.io/github/langered/gonedrive?branch=master)
[![Actions Status](https://github.com/langered/gonedrive/workflows/Build%20&%20Test%20gonedrive/badge.svg)](https://github.com/langered/gonedrive/actions)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/langered/gonedrive/blob/master/LICENSE)

# Gonedrive
A CLI to use OneDrive in the terminal.

#Installation
## From source
Build Gonedrive from the source files:

```
git clone https://github.com/langered/gonedrive.git
cd gonedrive
make build
```

# Usage
At first you have to authenticate yourself to microsoft in order to get access to the data you stored in your OneDrive Account.
Run:

`gonedrive login`

This will redirect you to the login page of Microsoft. After a successful login, Gonedrive receives the token and stores it in your config file.
The default location for the config file is in your homedirectory `~/.gonedrive.yml`
The received token is valid for 1 hour.

`gonedrive --help` list all currently available commands.

```
A CLI to interact with items stored in OneDrive

Usage:
  gonedrive [command]

Available Commands:
  get         Get the content of a given file as stdout
  help        Help about any command
  list        List items under given path
  login       Login to OneDrive
  upload      Upload a stdin to onedrive by into the given file
  version     Shows current used version

Flags:
      --config string   config file (default is $HOME/.gonedrive.yaml)
  -h, --help            help for gonedrive

Use "gonedrive [command] --help" for more information about a command.
```

## Commands
In the following, short examples are provided for each command.

### Get
Get the content of a file as stdout.

Example:

`gonedrive get Dir1/Dir2/MyFile.txt`

### Upload

Upload the stdin to a file in OneDrive. This will create or overwrite the given file. The directory 

`gonedrive upload Dir1/Dir2/MyFile.txt "My content"`

You can also upload content of a file like this:

`gonedrive upload Dir1/Dir2/MyFile.txt "$(cat <path-to-file>)"`

Please keep in mind that only small files up to 4 MB can get uploaded until now.

### List
Lists the files and directory under a given path. If no path is given, it will list the items under root.

```
gonedrive list Dir1/Dir2
[MyFile.txt SecondFile.yml Another-Dir]
```

## TODO

There is a lot to come to Gonedrive.

* Upload files directly instead of using the stdin
* Support large file uploads
* Get a file and store it directly into a file
* Delete remote files

I am also open for any feature requests, ideas and improvements.


## Using Docker

This Repository also contains docker-images which you can find [here](https://github.com/langered/gonedrive/packages).

To get the latest image, which is built from master, run:

`docker pull docker.pkg.github.com/langered/gonedrive/gonedrive:latest`

When using the Dockerfile, you have to mount the local config-file to the docker
image.

Example:

`docker run -v $HOME/.gonedrive.yml:/.gonedrive.yml gonedrive [command]`
