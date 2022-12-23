hrd
-

[![Build Status](https://github.com/thepkg/hrd/actions/workflows/ci.yml/badge.svg)](https://github.com/thepkg/hrd/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cn007b/hrd)](https://goreportcard.com/report/github.com/cn007b/hrd)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/a25d02fd43d34750911152b7a0d66ad6)](https://www.codacy.com/gh/cn007b/hrd/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cn007b/hrd&amp;utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/1d9d3d6acf37cde6e37e/maintainability)](https://codeclimate.com/github/cn007b/hrd/maintainability)

HRD - HTTP Request Dump, super simple `golang` web application
which dumps received HTTP request data back into response, which might be helpful for testing and debugging purposes.

## Usage

Pull docker image:
````sh
# basic app
docker pull ghcr.io/thepkg/hrd:v1.1.4

# with ENV dump
docker pull ghcr.io/thepkg/hrd:v1.1.5
````

Run docker container:
````sh
docker run -it -p 8080:8080 ghcr.io/thepkg/hrd:v1.1.5
````
