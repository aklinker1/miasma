build:
	docker build . -t miasma:dev
run: build
	@echo ""
	@echo "---"
	@echo ""
	docker run -i --env-file .env -p 3000:3000 -v "$(shell pwd)"/data:/data/miasma miasma:dev
watch:
	@modd
swagger:
	swagger generate server -t internal/gen -f ./api/swagger.yml --exclude-main -A miasma
