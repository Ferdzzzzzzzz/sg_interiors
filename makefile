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

gen:
	@tailwindcss -c ./tailwind.config.js -i ./in.css -o ./assets/styles/main.css
	@templ generate

# ==================================================================================================
# PROD

# This is the tailscale dns for this vm
rockup-1="rockup-1"

init-service:
	@echo "DEPLOYING SERVICE"
	@echo "=========================="
	@env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./bin/linux_arm64/sg_interiors ./main.go
	rsync -P ./prod/sg.service ubuntu@${rockup-1}:~
	rsync -P ./bin/linux_arm64/sg_interiors ubuntu@${rockup-1}:~
	ssh -t ubuntu@${rockup-1} 'sudo mv ~/sg.service /etc/systemd/system/ && sudo systemctl enable sg && sudo systemctl restart sg'

deploy:
	@echo "=========================="
	@echo "DEPLOYING PROD"
	@echo "=========================="
	@echo ""
	@echo ""
	@echo "BUILDING BINARY"
	@echo "========================="
	@echo "OS\tlinux"
	@echo "ARCH\tarm64"
	@env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./bin/linux_arm64/sg_interiors ./main.go
	@echo ""
	@echo "TRANSFER"
	@echo "=========================="
	@rsync -rP --delete ./bin/linux_arm64/sg_interiors ubuntu@${rockup-1}:~
	@echo ""
	@echo "RESTARTING"
	@echo "=========================="
	@echo "Restarting systemd service..."
	@ssh -t ubuntu@${rockup-1} 'sudo systemctl restart sg'
	@echo ""
	@echo "STATUS"
	@echo "=========================="
	@echo "complete"


add-sg-service:
	@echo "DEPLOYING SERVICE"
	@echo "=========================="
	@env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./bin/linux_arm64/sg_interiors ./main.go
	rsync -P ./prod/sg.service ubuntu@${rockup-1}:~
	rsync -P ./bin/linux_arm64/sg_interiors ubuntu@${rockup-1}:~
	ssh -t ubuntu@${rockup-1} 'sudo mv ~/sg.service /etc/systemd/system/ && sudo systemctl enable sg && sudo systemctl restart sg'