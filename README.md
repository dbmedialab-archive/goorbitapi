GO OrbitAPI client
==================

[![Build Status](https://travis-ci.org/dbmedialab/goorbitapi.svg)](https://travis-ci.org/dbmedialab/goorbitapi) [![Coverage Status](https://coveralls.io/repos/dbmedialab/goorbitapi/badge.png)](https://coveralls.io/r/dbmedialab/goorbitapi) [![GoDoc](https://godoc.org/github.com/dbmedialab/goorbitapi?status.png)](https://godoc.org/github.com/dbmedialab/goorbitapi)

Go client for accessing the Orbit API - http://orbitapi.com/

Currently only Concept Tagging call has its own data type and API call.  All other calls can be made against the API, but these need to be done through Post() or Get().

See the `cmd/` directory for sample code:

## Quick start

```
$ cd cmd
$ vim sample.php
	Edit line 11 and add your API key
$ go run sample.php
```

### Text query: entity information from Wikipedia

```
$ cd cmd
$ ORBIT_API_KEY=YOUR-KEY go run textquery.go -text "Jeg liker politikk sa Solberg til Dagbladet." -sentences 3
```