build:
	npx tailwindcss -i ./assets/global.css -o ./public/global.css --minify
	go build

clean:
	rm -f ./public/global.css
	# rm -f ./rymi.dev

css:
	npx tailwindcss -i ./assets/global.css -o ./public/global.css --watch --minify

run:
	concurrently "air main.go" "make css" "browser-sync start --proxy "127.0.0.1:8080" --files "components/*.templ" --reloadDelay 3000 --no-open --no-notify"