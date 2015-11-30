# FreeTree ![Status](https://img.shields.io/badge/status-stable-green.svg?style=plastic) [![Build Status](http://img.shields.io/travis/teh-cmc/freetree.svg?style=plastic)](https://travis-ci.org/teh-cmc/freetree) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=plastic)](http://godoc.org/github.com/teh-cmc/freetree)

Generic binary search tree with zero GC overhead, for golang. Built on [mmm](https://github.com/teh-cmc/mmm).

## What you should know

`FreeTree` was developed mainly as a proof-of-concept for [mmm](https://github.com/teh-cmc/mmm); that is, to demonstrate how you can use `mmm` to avoid GC overhead in "pointer-heavy" Go software, without modifying nor complexifying your original design (i.e. entirely redesigning your software to avoid the use of pointers, which often leads to overly complex and less maintainable code).

Although I do use it for some big immutable caches of mine, `FreeTree`'s API is quite incomplete and could certainly be better designed; especially during initialization where a lot of unnecessary copying could probably be avoided.
Feel free to improve it :)

## Install

```bash
go get -u github.com/teh-cmc/freetree
```

## Example

Here's a simple example of usage (code [here](examples/simple.go)):

```Go
```

## Demonstration

Complete code for the following demonstration is available [here](experiment/experiment.go).

All of the results shown below were computed using a DELL XPS 15-9530 (i7-4712HQ@2.30GHz).

## License ![License](https://img.shields.io/badge/license-MIT-blue.svg?style=plastic)

The MIT License (MIT) - see LICENSE for more details

Copyright (c) 2015	Clement 'cmc' Rey	<cr.rey.clement@gmail.com>
