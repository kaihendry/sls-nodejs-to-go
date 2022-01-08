deploy:
	cd world && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	sls deploy

justgofunction:
	cd world && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	serverless deploy --function world
