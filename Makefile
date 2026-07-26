.PHONY: dev build css generate dist clean install run help setup

# ── Setup ──────────────────────────────────────────────────

setup: install
	@mkdir -p assets/css assets/js assets/images assets/fonts
	@mkdir -p cmd view/layout view/homepage model handler utils builder
	@mkdir -p .github/workflows
	@echo "Project structure created."
	@echo "Run 'make install' if not already done, then 'make dev' to start."

# ── Development ────────────────────────────────────────────

dev: install
	air

# ── Build ──────────────────────────────────────────────────

build: css generate
	go build -o ./tmp/main.exe ./cmd

# ── CSS ────────────────────────────────────────────────────

css:
	unocss "**/*.templ" -o assets/css/input.css --config uno.config.ts

# ── Templates ──────────────────────────────────────────────

generate:
	templ generate

# ── Static Site ────────────────────────────────────────────

dist: build
	go run ./cmd/generate

# ── Run ────────────────────────────────────────────────────

run: build
	./tmp/main.exe

# ── Dependencies ───────────────────────────────────────────

install:
	go mod tidy
	@command -v templ >/dev/null 2>&1 || go install github.com/a-h/templ/cmd/templ@latest
	@command -v unocss >/dev/null 2>&1 || npm install -g unocss
	@command -v air >/dev/null 2>&1 || go install github.com/air-verse/air@latest
	@echo "Downloading Alpine.js CSP build..."
	@curl -sL https://cdn.jsdelivr.net/npm/@alpinejs/csp/dist/cdn.min.js -o assets/js/cdn.min.js 2>/dev/null || echo "Alpine.js: skip (curl not available or already exists)"
	@echo "Dependencies installed."

# ── Clean ──────────────────────────────────────────────────

clean:
	rm -rf tmp dist

# ── Help ──────────────────────────────────────────────────

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  setup      Create project structure + install all dependencies"
	@echo "  dev        Start development server with live reload (air)"
	@echo "  build      Generate CSS + templates, compile Go binary"
	@echo "  css        Generate UnoCSS output (assets/css/input.css)"
	@echo "  generate   Generate Go code from .templ files"
	@echo "  dist       Build static site for deployment (output: dist/)"
	@echo "  run        Build and run the server once"
	@echo "  install    Install Go + Node dependencies + Alpine.js"
	@echo "  clean      Remove build artifacts (tmp/, dist/)"
	@echo "  help       Show this help"
