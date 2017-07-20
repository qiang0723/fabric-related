### Prepare

```bash
$ go get gopkg.in/olivere/elastic.v5
```

> **Note:** 
Get this package may need webproxy. 
And this client correspond to elasticsearch-5.0.0

#### Config webproxy(optional)

1. In container, config /etc/resolv.conf just like host
2. Set http_proxy with authenticate

```bash
export https_proxy=https://bjwqwang%40cn.ibm.com:passwd@webproxy.watson.ibm.com:8080
export http_proxy=http://bjwqwang%40cn.ibm.com:passwd@webproxy.watson.ibm.com:8080
```
3. Test whether the configuration is valid
```bash
curl -XGET www.google.com
```

### Boot MVE 

### Git clone [fabric-docker-compose](https://github.com/yeasy/docker-compose-files.git)

```bash
$ cd ~/docker-compose/docker-compose-files/hyperledger/1.0
$ sudo docker-compose up
```
### Running program events-es to listen events

At terminal `fabric-peer0`:

```bash
$ docker exec -it fabric-peer0 bash
$ cd example/events/block-listener
$ go get gopkg.in/olivere/elastic.v5
$ go build events-es.go
$ ./events-es
```

### Send events

At terminal `fabric-cli`:

```bash
$ docker exec -it fabric-cli bash
$ peer chaincode  install -v 1.0 -n test_cc -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02 -c '{"Args":["init","a","100","b","200"]}' -o orderer0:7050
$ peer chaincode  instantiate -v 1.0 -n test_cc -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02 -c '{"Args":["init","a","100","b","200"]}' -o orderer0:7050
$ peer chaincode query -n test_cc -c '{"Args":["query","a"]}' -o orderer0:7050
$ peer chaincode invoke -n test_cc -c '{"Args":["invoke","a","b","10"]}' -o orderer0:7050
```

> **Note:**
If you want to see the events about create/join channel, should start MVE with `docker-compose-new-channel`


### Curl ES to check resault

```bash
curl -XGET http://172.16.3.9:9200/_cat/indices?v
curl -XGET http://172.16.3.9:9200/eventhub/_search?

```