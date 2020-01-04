<h1 align="center">Welcome to go-lastfm 👋 *** Work In Progress ***</h1>
<p>
  <p align="center">
  <a href="https://github.com/tsirysndr/go-lastfm/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/go-lastfm.svg" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/go-lastfm">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/go-lastfm">
  <a href="https://github.com/tsirysndr/go-lastfm/issues?q=is%3Apr+is%3Aclosed">
    <img alt="GitHub closed pull requests" src="https://img.shields.io/github/issues-pr-closed-raw/tsirysndr/go-lastfm">
  </a>
  <a href="https://github.com/tsirysndr/go-lastfm/pulls">
    <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/tsirysndr/go-lastfm">
  </a>
  <a href="https://github.com/tsirysndr/go-lastfm/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/tsirysndr/go-lastfm">
  </a>
  <a href="https://github.com/tsirysndr/go-lastfm/graphs/contributors">
    <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/go-lastfm">
  </a>
  <a href="https://github.com/tsirysndr/go-lastfm/blob/master/LICENSE">
    <img alt="License: BSD" src="https://img.shields.io/badge/license-BSD-green.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/tsiry_sndr">
    <img alt="Twitter: tsiry_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" target="_blank" />
  </a>
</p>
</p>

go-lastfm is a Go client library for accessing the [Last.fm API](http://www.last.fm/api/)

## 🚚 Install

```sh
go get github.com/tsirynsdr/go-lastfm
```

## 🚀 Usage

Import the package into your project.

```Go
import "github.com/tsirynsdr/go-lastfm"
```

Construct a new Last.fm client, then use the various services on the client to access different parts of the Last.fm API. For example:

```Go
client := lastfm.NewClient("<YOUR API KEY>")
res, _ := client.Artist.GetInfo("Travis Scott")
info, _ := json.Marshal(res)
fmt.Println(string(info))
```

## ✨ Coverage

Currently the following services are supported:

- [ ] Tag an album using a list of user supplied tag
- [ ] Get the metadata and tracklist for an album on Last.fm using the album name or a musicbrainz id
- [ ] Get the tags applied by an individual user to an album on Last.fm
- [ ] Get the top tags for an album on Last.fm, ordered by popularity
- [ ] Remove a user's tag from an album
- [ ] Search for an album by name. Returns album matches sorted by relevance
- [ ] Tag an artist with one or more user supplied tags
- [ ] Use the last.fm corrections data to check whether the supplied artist has a correction to a canonical artist
- [x] Get the metadata for an artist. Includes biography, truncated at 300 characters
- [ ] Get all the artists similar to this artist
- [ ] Get the tags applied by an individual user to an artist on Last.fm
- [ ] Get the top albums for an artist on Last.fm, ordered by popularity
- [ ] Get the top tags for an artist on Last.fm, ordered by popularity
- [ ] Get the top tracks by an artist on Last.fm, ordered by popularity
- [ ] Remove a user's tag from an artist
- [x] Search for an artist by name. Returns artist matches sorted by relevance
- [ ] Get the top artists chart
- [ ] Get the top tracks chart 
- [ ] Get the most popular artists on Last.fm by country
- [ ] Get the most popular tracks on Last.fm last week by country
- [ ] Get a paginated list of all the artists in a user's library, with play counts and tag counts
- [ ] Get the metadata for a tag
- [ ] Search for tags similar to this one. Returns tags ranked by similarity, based on listening data
- [ ] Get the top albums tagged by this tag, ordered by tag count
- [ ] Get the top artists tagged by this tag, ordered by tag count
- [ ] Fetches the top global tags on Last.fm, sorted by popularity (number of times used)
- [ ] Get the top tracks tagged by this tag, ordered by tag count
- [ ] Get a list of available charts for this tag, expressed as date ranges which can be sent to the chart services
- [ ] Tag an album using a list of user supplied tags
- [ ] Use the last.fm corrections data to check whether the supplied track has a correction to a canonical track
- [ ] Get the metadata for a track on Last.fm using the artist/track name or a musicbrainz id
- [ ] Get the similar tracks for this track on Last.fm, based on listening data
- [ ] Get the tags applied by an individual user to a track on Last.fm
- [ ] Get the top tags for this track on Last.fm, ordered by tag count
- [ ] Love a track for a user profile
- [ ] Remove a user's tag from a track
- [ ] Track.scrobble
- [ ] Track.search
- [ ] Track.unlove
- [ ] Track.updateNowPlaying
- [ ] Get a list of the user's friends on Last.fm
- [ ] Get information about a user profile
- [ ] Get the last 50 tracks loved by a user
- [ ] Get the user's personal tags
- [ ] Get a list of the recent tracks listened to by this user
- [ ] Get the top albums listened to by a user
- [ ] Get the top artists listened to by a user
- [ ] Get the top tags used by this user
- [ ] Get the top tracks listened to by a user
- [ ] Get an album chart for a user profile, for a given date range
- [ ] Get an artist chart for a user profile, for a given date range
- [ ] Get a list of available charts for this user, expressed as date ranges which can be sent to the chart services
- [ ] Get a track chart for a user profile, for a given date range


## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!
