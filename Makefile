all:
	@echo "[AWS S3] Compiling modules"

	@cd cli/ && go build -o ../bin/cli
	@echo "[AWS S3] The CLI was compiled"

	@cd server/ && go build -o ../bin/server
	@echo "[AWS S3] The server was compiled"

	@echo "[AWS S3] Done!!"