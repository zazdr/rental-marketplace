.PHONY: ls install

ls:
	@mise ls

install:
	@mise install
	@curl -s -o ./app/ui/static/js/htmx.min.js \
		https://cdn.jsdelivr.net/npm/htmx.org@2.0.10/dist/htmx.min.js
