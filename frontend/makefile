# Makefile for Frontend Application
# This makefile automates common tasks for the frontend application located in the 'app' directory.

APP_DIR := app
NPM := npm

.PHONY: help install dev run build clean preview

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^##' $(MAKEFILE_LIST) | sed -e 's/## //g' | column -t -s ':'

## install: Install dependencies
install:
	cd $(APP_DIR) && $(NPM) install

## dev: Start the development server (alias for run)
dev: run

## run: Start the development server
run:
	cd $(APP_DIR) && $(NPM) run dev

## build: Build the application for production
build:
	cd $(APP_DIR) && $(NPM) run build

## preview: Preview the production build locally
preview:
	cd $(APP_DIR) && $(NPM) run preview

## clean: Remove build artifacts and dependencies
clean:
	rm -rf $(APP_DIR)/dist
	rm -rf $(APP_DIR)/node_modules