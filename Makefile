build:
	docker build . -t aklinker1/miasma:dev
run: build
	@echo ""
	@echo "---"
	@echo ""
	docker run -i --env-file .env -p 3000:3000 -v "$(shell pwd)"/data:/data/miasma -v /var/run/docker.sock:/var/run/docker.sock aklinker1/miasma:dev
watch:
	@modd
swagger:
	swagger generate server -t internal/server/gen -f ./api/swagger.yml --exclude-main -A miasma
publish:
	docker login
	docker buildx build \
		--push \
		--platform linux/arm/v7,linux/arm64/v8,linux/amd64 \
		--tag aklinker1/miasma:nightly .

#  Aliases
b: build
r: run
w: watch
s: swagger
p: publish
