#!/usr/bin/env sh

migrate create -ext sql -dir /migrations -seq $@
