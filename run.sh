#!/bin/bash

# shellcheck disable=SC2046
export $(grep -v '^#' .env | xargs )

go run main.go