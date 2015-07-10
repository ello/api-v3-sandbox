all: pretty test

clean: 
	rm -rf node_modules

pretty:
	@node ./node_modules/js-beautify/js/bin/js-beautify.js -j -r -q *.js
	@node ./node_modules/js-beautify/js/bin/js-beautify.js -j -r -q **/*.js

server: 
	@bin/start

server-w: 
	@bin/start -w

test: 
	@node node_modules/mocha/bin/mocha -t 250 --reporter spec --check-leaks

test-w:
	@node node_modules/mocha/bin/mocha -t 250 --reporter min --check-leaks -w

.PHONY: test
