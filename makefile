run:
	direnv allow .  && export FUNCTION_TARGET=threads && go run cmd/main.go