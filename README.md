# Leboncoin Fizz Buzz like technical test repo

[![Build Status](https://api.travis-ci.com/monkeydioude/lbc-test.svg?branch=master)](https://app.travis-ci.com/github/monkeydioude/lbc-test)

This REST API exposes 2 endpoints:

- POST /fizz-buzz
- GET /fizz-buzz/stats

### POST /fizz-buzz

Takes 5 query parameters:

- int1
- int2
- limit
- str1
- str2

ex: /fizz-buzz?int1=2&int2=3&limit=10&str1=a&str2=b

Core endpoint of the test. For details see below.

### GET /fizz-buzz/stats

Takes no parameter.

Display query parameters and number of hits of the most used request to POST /fizz-buzz.

### Tests and Run

- To run tests: `go test -v ./...`
- To run the project: `make` or `cd cmd/lbc-test && go install && $GOBIN/lbc-test`

2 environment vars can be provided to `lbc-test` binary:

- `PORT` sets which port lbc-test web server should listen to. Default is `8084`
- `DB_PATH` sets location of the binary db file for the `/fizz-buzz/stats` endpoint. Default is `./`

### Details of the test:

```
Exercise: Write a simple fizz-buzz REST server.

"The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"".
The output would look like this: ""1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...""."

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters: three integers Int1, Int2 and Limit, and two strings Str1 and Str2.
- Returns a list of strings with numbers from 1 to Limit, where: all multiples of Int1 are replaced by Str1, all multiples of Int2 are replaced by Str2, all multiples of Int1 and Int2 are replaced by Str1Str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request
```

### Choices for this project:

As a technical preface for a live interview, I present this test as a proof of my base knowledge of Go language. With that in mind I chose to restrict myself and handle this test using Go's standard library only, with the exception of only one external package, used in the optional part of this test.

#### 1/ Directory structure

I chose to use [golang standards project layout](https://github.com/golang-standards/project-layout), a classic for separating internal and externally re-usable libraries.

#### 2/ Libraries

This project use only standard library, with the exception of 1 external package. `go.etcd.io/bbolt` is a binary file based, key/value storage, kind of DB providing a very low level and robust API, with a lock system, made for speed.

#### 3/ Technical choices

Following my take on the test, I decided to show my knowledge of Go by building almost everything from scratch.

Serving with the basic `net/http` server and handling requests using `http.Handler` interface is, to me, a good way to show Middleware concept and understanding of Go's interface.

Building a router -as an externally reusable package- handling `*http.Request` strenghtens this demonstration of `net/http` package understanding and shows that I am capable of building a simple HTTP mux over the fairly complex concept of `http.Handler` middlewares. I could have used `gorilla/mux` or `gin-gonic/gin`, especially `gin`, which I am a bit fond of its performances and simplicity (the router I built kind of look like gin's syntax).
