build:
	docker build . -t miasma:dev
run: build
	@echo ""
	@echo "---"
	@echo ""
	docker run -i --env-file .env -p 3000:3000 -v "$(shell pwd)"/data:/data/miasma -v /var/run/docker.sock:/var/run/docker.sock miasma:dev
watch:
	@modd
swagger:
	swagger generate server -t internal/server/gen -f ./api/swagger.yml --exclude-main -A miasma

#  Aliases
b: build
r: run
w: watch
s: swagger