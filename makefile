dev:
	@tailwindcss -c ./tailwind.config.js -i ./in.css -o ./assets/styles/main.css
	@templ generate
	@echo "--------------------------------------------------------"
	@echo "Starting Server"
	@echo ""
	@go run main.go --dev

style:
	@tailwindcss -c ./tailwind.config.js -i ./in.css -o ./assets/styles/main.css --watch

build:
	@tailwindcss -c ./tailwind.config.js -i ./in.css -o ./assets/styles/main.css --minify
	@templ generate