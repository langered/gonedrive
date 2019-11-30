[![Go Report Card](https://goreportcard.com/badge/github.com/langered/gonedrive)](https://goreportcard.com/report/github.com/langered/gonedrive)
[![Coverage Status](https://coveralls.io/repos/github/langered/gonedrive/badge.svg?branch=master)](https://coveralls.io/github/langered/gonedrive?branch=master)
[![Actions Status](https://github.com/langered/gonedrive/workflows/Build%20&%20Test%20gonedrive/badge.svg)](https://github.com/langered/gonedrive/actions)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/langered/gonedrive/blob/master/LICENSE)

# Using Docker

When using the Dockerfile, you have to mount the local config-file to the docker
image.

This can be done with the following command, when the config file is stored in
the home directory:

`docker run -v $HOME/.gonedrive.yml:/.gonedrive.yml gonedrive [command]`
