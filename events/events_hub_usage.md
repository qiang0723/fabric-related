
##Test chaincode events

`docker exec -it Vp0 bash`:
```bash
$ cd  /go/src/github.com/hyperledger/fabric/examples/events/block-listener
$ go build
$ ./block-listener
```

Vp0 will launch a listener, and wait for events

```bash
root@vp0:/go/src/github.com/hyperledger/fabric/examples/events/block-listener# ./block-listener 
Event Address: 0.0.0.0:7053

Received block
--------------
Transaction:
	[type:CHAINCODE_DEPLOY chaincodeID:"\n?github.com/hyperledger/fabric/examples/chaincode/go/eventsender\022\200\0011eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9" payload:"\n\341\001\010\001\022\304\001\n?github.com/hyperledger/fabric/examples/chaincode/go/eventsender\022\200\0011eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9\032\026\n\004init\n\001a\n\003100\n\001b\n\003200" txid:"1eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9" timestamp:<seconds:1489569033 nanos:606882736 > ]

Received block
--------------
Transaction:
	[type:CHAINCODE_INVOKE chaincodeID:"\022\200\0011eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9" payload:"\n\234\001\010\001\022\203\001\022\200\0011eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9\032\022\n\006invoke\n\001a\n\001b\n\00210" txid:"992b2fe7-bd90-4d12-bd91-e94f1aae4349" timestamp:<seconds:1489569776 nanos:69171498 > ]


Received block
--------------
Transaction:
	[type:CHAINCODE_INVOKE chaincodeID:"\022\200\0011eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9" payload:"\n\224\001\010\001\022\203\001\022\200\0011eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9\032\n\n\005query\n\001a" txid:"a9654054-23c5-4333-ac54-135f98772c6c" timestamp:<seconds:1489569795 nanos:64624895 > ]
```

`docker exec -it Vp1 bash`:
```bash
$ export CCPATH=github.com/hyperledger/fabric/examples/chaincode/go/eventsender
```
Send a deploy transaction:
```bash
$ CORE_PEER_ADDRESS=0.0.0.0:7051 peer chaincode deploy  -l golang  -p $CCPATH   -c '{"Function":"init", "Args": ["a","100", "b", "200"]}'
```

Send a invoke transacton

```bash
$ export CCID=1eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9 

$ CORE_PEER_ADDRESS=0.0.0.0:7051  peer chaincode invoke -n $CCID  -c '{"Function":"invoke", "Args": ["a","b","10"]}'

$ CORE_PEER_ADDRESS=0.0.0.0:7051  peer chaincode invoke -n $CCID -c '{"Function": "query", "Args": ["a"]}'
```

### block-listener  listening chaincode 

Vp0 :
```bash
export CCID=1ee...
./block-listener -events-from-chaincode=$CCID
```

Vp1, Test chaincode event
```bash
root@vp1:/go/src/github.com/hyperledger/fabric# CORE_PEER_ADDRESS=0.0.0.0:7051  peer chaincode invoke -n $CCID  -c '{"Function":"greet", "Args":["hello","world"]}'
Output:
11:04:00.564 [logging] LoggingInit -> DEBU 001 Setting default logging level to DEBUG for command 'chaincode'
11:04:00.567 [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 002 Successfully invoked transaction: chaincodeSpec:<type:GOLANG chaincodeID:<name:"1eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9" > ctorMsg:<args:"greet" args:"hello" args:"world" > > (2c621596-5a65-4f75-aca9-0d3ef9ecc2eb)
11:04:00.571 [main] main -> INFO 003 Exiting.....
```

Vp0 output:
```bash
Received block
--------------
Transaction:
	[type:CHAINCODE_INVOKE chaincodeID:"\022\200\0011eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9" payload:"\n\237\001\010\001\022\203\001\022\200\0011eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9\032\025\n\005greet\n\005hello\n\005world" txid:"2c621596-5a65-4f75-aca9-0d3ef9ecc2eb" timestamp:<seconds:1489575840 nanos:566925340 > ]


Received chaincode event
------------------------
Chaincode Event:&{chaincodeID:"1eecca2f9029d466e2e0bd648985a0c671a1d1d5e4f880e240cb8c03eba2a3fc99e2888e85e5da04630b03c55bfaba7f43bb39a117774196de9b51647eac05b9" txID:"2c621596-5a65-4f75-aca9-0d3ef9ecc2eb" eventName:"evtsender" payload:"Event 6,hello,world" }
```

## Test chaincode

vp0:
```bash
$ go build
$ ./block-listener
```

vp1 send valid deploy/create transaction to chaincode_example02
```bash
deploy:
$ export CCPATH=github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02
$ CORE_PEER_ADDRESS=0.0.0.0:7051 peer chaincode deploy -p $CCPATH -c '{"Function":"init", "Args": ["a","100", "b", "200"]}'

output:
11:25:03.450 [logging] LoggingInit -> DEBU 001 Setting default logging level to DEBUG for command 'chaincode'
11:25:06.989 [chaincodeCmd] chaincodeDeploy -> INFO 002 Deploy result: type:GOLANG chaincodeID:<path:"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02" name:"ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539" > ctorMsg:<args:"init" args:"a" args:"100" args:"b" args:"200" > 
Deploy chaincode: ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539
```

Now Vp0 output:
```bash
Received block
--------------
Transaction:
	[type:CHAINCODE_DEPLOY chaincodeID:"\nGgithub.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02\022\200\001ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539" payload:"\n\351\001\010\001\022\314\001\nGgithub.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02\022\200\001ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539\032\026\n\004init\n\001a\n\003100\n\001b\n\003200" txid:"ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539" timestamp:<seconds:1489577106 nanos:776382645 > ]

```
Vp1 input:
```bash
invoke:
$ export CCID=ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539
$ CORE_PEER_ADDRESS=0.0.0.0:7051 peer chaincode invoke -n $CCID -c '{"Function":"invoke", "Args": ["a","b","10"]}'

output:
11:28:26.367 [logging] LoggingInit -> DEBU 001 Setting default logging level to DEBUG for command 'chaincode'
11:28:26.374 [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 002 Successfully invoked transaction: chaincodeSpec:<type:GOLANG chaincodeID:<name:"ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539" > ctorMsg:<args:"invoke" args:"a" args:"b" args:"10" > > (5a58e9c8-5c16-4a2f-9d8a-10e203e437e4)
11:28:26.374 [main] main -> INFO 003 Exiting.....
```

Now Vp0 output:
```bash
Received block
--------------
Transaction:
	[type:CHAINCODE_INVOKE chaincodeID:"\022\200\001ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539" payload:"\n\234\001\010\001\022\203\001\022\200\001ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539\032\022\n\006invoke\n\001a\n\001b\n\00210" txid:"5a58e9c8-5c16-4a2f-9d8a-10e203e437e4" timestamp:<seconds:1489577306 nanos:373503591 > ]
```

### block-listener listen chaincode

Restart block-listener with -events-from-chaincode.
 
 Vp0 listening chaincode:

```bash
 $ ./block-listener -events-from-chaincode=$CCID
```
 
 Vp1 send a invoke transaction
 ```bash
 $ CORE_PEER_ADDRESS=0.0.0.0:7051 peer chaincode invoke -n $CCID -c '{"Function":"invoke", "Args": ["a","b","10"]}'

output:
11:35:53.990 [logging] LoggingInit -> DEBU 001 Setting default logging level to DEBUG for command 'chaincode'
11:35:53.996 [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 002 Successfully invoked transaction: chaincodeSpec:<type:GOLANG chaincodeID:<name:"ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539" > ctorMsg:<args:"invoke" args:"a" args:"b" args:"10" > > (9720eb57-4aa1-4408-a752-e358edd7d3bd)
11:35:53.996 [main] main -> INFO 003 Exiting.....
```
 
 Now Vp0 output:
 ```bash
 Received block
--------------
Transaction:
	[type:CHAINCODE_INVOKE chaincodeID:"\022\200\001ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539" payload:"\n\234\001\010\001\022\203\001\022\200\001ee5b24a1f17c356dd5f6e37307922e39ddba12e5d2e203ed93401d7d05eb0dd194fb9070549c5dc31eb63f4e654dbd5a1d86cbb30c48e3ab1812590cd0f78539\032\022\n\006invoke\n\001a\n\001b\n\00210" txid:"9720eb57-4aa1-4408-a752-e358edd7d3bd" timestamp:<seconds:1489577753 nanos:993991367 > ]


 ```
 
 ## Analytics
 
 Before we use batch data for analyze
 
 logstash -> elasticsearch ->spark
 
 Now we use real-time data
 
 logspout -> kafka -> sparkstreaming -> db(?)
 