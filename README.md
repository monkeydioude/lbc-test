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

Core endpoint of the test. For details see below.

### GET /fizz-buzz/stats
Takes no parameter.

Display query parameters and number of hits of the most used request to POST /fizz-buzz.


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