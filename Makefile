.PHONY: unocss

unocss:
	unocss "**/*.templ" -o assets/css/input.css --config uno.config.ts
