[![Go Report Card](https://goreportcard.com/badge/github.com/langered/gonedrive)](https://goreportcard.com/report/github.com/langered/gonedrive)
[![Coverage Status](https://coveralls.io/repos/github/langered/gonedrive/badge.svg?branch=master)](https://coveralls.io/github/langered/gonedrive?branch=master)
[![Actions Status](https://github.com/langered/gonedrive/workflows/Build%20&%20Test%20gonedrive/badge.svg)](https://github.com/langered/gonedrive/actions)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/langered/gonedrive/blob/master/LICENSE)

# Gonedrive
A CLI to use OneDrive in the terminal.

# Installation

You can use the binary in the release or build a new binary from source.

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

## Documentation
Each command is documented in the [doc folder](https://github.com/langered/gonedrive/tree/master/doc)

## Using Docker

This Repository also contains docker-images which you can find [here](https://github.com/langered/gonedrive/packages).

To get the latest image, which is built from master, run:

`docker pull docker.pkg.github.com/langered/gonedrive/gonedrive:latest`

When using the Dockerfile, you have to mount the local config-file to the docker
image.

Example:

`docker run -v $HOME/.gonedrive.yml:/.gonedrive.yml gonedrive [command]`
