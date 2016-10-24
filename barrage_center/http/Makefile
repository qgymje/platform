PROJECT_NAME=barrage_center_http

pack: clean
	godep save

clean:
	@echo "removing pkgs.." && go clean

http:
	fswatch

rpc:
	fswatch

build:
	env GOOS=linux GOARCH=amd64 go build -o $(PROJECT_NAME)

sync:
	scp $(PROJECT_NAME) root@node1:~/platform/$(PROJECT_NAME)
	scp ./configs/* root@node1:~/platform/$(PROJECT_NAME)/configs/
