<h1 align="center">Welcome to go-lastfm 👋 *** Work In Progress ***</h1>
<p>
  <p align="center">
  <a href="https://github.com/tsirysndr/go-lastfm/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/go-lastfm.svg" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/go-lastfm">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/go-lastfm">
  <img alt="GitHub closed pull requests" src="https://img.shields.io/github/issues-pr-closed-raw/tsirysndr/go-lastfm">
  <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/tsirysndr/go-lastfm">
  <img alt="GitHub issues" src="https://img.shields.io/github/issues/tsirysndr/go-lastfm">
  <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/go-lastfm">
  <a href="https://github.com/tsirysndr/go-lastfm/blob/master/LICENSE">
    <img alt="License: BSD" src="https://img.shields.io/badge/license-BSD-green.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/tsiry_sndr">
    <img alt="Twitter: tsiry_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" target="_blank" />
  </a>
</p>
</p>

go-lastfm is a Go client library for accessing the [Last.fm API](http://www.last.fm/api/)

## Install

```sh
go get github.com/tsirynsdr/go-lastfm
```

## Usage

Import the package into your project.

```Go
import "github.com/tsirynsdr/go-lastfm"
```

Construct a new Last.fm client, then use the various services on the client to access different parts of the Last.fm API. For example:

```Go
client := lastfm.NewClient("<YOUR API KEY>")
```

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!
