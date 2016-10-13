pack: clean
	godep save

clean:
	@echo "removing pkgs.." && go clean

http:
	fswatch

rpc:
	fswatch

build: clean
	rm -rf ./rpc
	env GOOS=linux GOARCH=amd64 go build -o rpc 

docker: build
	docker build -t coupon_center/rpc .

