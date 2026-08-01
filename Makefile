.PHONY: ls install

ls:
	@mise ls

HTMX        = https://cdn.jsdelivr.net/npm/htmx.org@2.0.10/dist/htmx.min.js
HYPERSCRIPT = https://cdn.jsdelivr.net/npm/hyperscript.org@0.9.93/dist/_hyperscript.min.js

install:
	@mise install
	@mkdir -p ./app/ui/static/js/
	@curl -s -o ./app/ui/static/js/htmx.min.js $(HTMX)
	@curl -s -o ./app/ui/static/js/hyperscript.min.js $(HYPERSCRIPT)
