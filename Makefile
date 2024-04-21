run:
	templ generate
	npx tailwindcss -i ./assets/input.css -o ./assets/output.css
	npm run build
	go run .

createdb:
	docker run -d --name mongodb -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=secret mongo

dropdb:
	docker stop mongodb
	docker rm mongodb