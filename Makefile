build_lambda:
	GOOS=linux GOARCH=amd64 go build -o lambdaBuild/bootstrap ./lambdaSrc
	zip lambdaBuild/function.zip lambdaBuild/bootstrap