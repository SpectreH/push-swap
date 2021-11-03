#!/bin/bash
go build -o checker check/main.go | go build -o push-swap pushswap/main.go