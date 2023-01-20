FILES=*.go

client: utils.go client.go
	go build $(FILES)
