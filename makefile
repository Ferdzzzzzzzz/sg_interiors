dev:
	@go run main.go --dev

style:
	@tailwindcss -c ./tailwind.config.js -i ./in.css -o ./assets/styles/main.css --watch

build-style:
	@tailwindcss -c ./tailwind.config.js -i ./in.css -o ./assets/styles/main.css --minify

gen:
	@templ generate