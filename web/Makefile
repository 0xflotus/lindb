export PATH := node_modules/.bin:$(PATH)
export SHELL := /bin/bash

web_build:
	yarn install
	npm run build

web_clean:
	rm -rf web/build \
    rm -rf web/node_modules