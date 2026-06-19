all:
	@cd cli/ && go build -o ../bin/aws-cli
	@cd server/ && go build -o ../bin/aws-server