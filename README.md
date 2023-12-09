# NOTICE
This is a fork from 'facebookexperimental/go-media-webtransport-server' and 'facebookexperimental/webcodecs-capture-play' <br/>
to put them in one repository and make the client (webcodecs-capture-play) run using Gin with the certificates.<br/>
Also, updated the scripts to generate certificates from mkcert.<br/>
This way I can study the code without running browser with specific arguments and I will be able to test it from firefox. <br/>
Unfortunately, current version of Firefox does not support MediaStreamTrackProcessor <br/>

## Installation
Install go 1.18 to compile
Install mkcert from https://github.com/FiloSottile/mkcert
And run
```bash
go mod tidy
bash scripts/create-server-certs.sh

go build -o client/client ./client/main.go
go build -o server/server ./server/main.go
```

## Running
First increase the UDP buffer size for WebTransport to work
```bash
sudo sysctl -w net.core.rmem_max=2500000
sudo sysctl -w net.core.wmem_max=2500000
```

Open a terminal and run the server
```bash
cd server
./server
```

Open another terminal and run the client
```bash
cd client
./client
```

## Using
Open Google's Chrome and two tabs </br>
The first tab visit https://localhost:8080/encoder <br/>
The second tab visit https://localhost:8080/player <br/>
Copy the StreamID from the first tab to the second tab <br/>
Press start from the first tab (allow camera) and then on the second tab and you will start seeing the feed from the first tab going to the second tab <br/>