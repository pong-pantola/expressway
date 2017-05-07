OLD:
https://f188c7a4c9c64941b0e593ce52a54707-vp0.us.blockchain.ibm.com:5001

NEW:
https://0b1ca0591fee4d7d882a5e691b780d1d-vp0.us.blockchain.ibm.com:5002


curl -i -X POST -H "Content-Type:application/json" https://0b1ca0591fee4d7d882a5e691b780d1d-vp0.us.blockchain.ibm.com:5002/registrar -d '{ "enrollId": "user_type1_3", "enrollSecret": "8dc9fff386" }'

curl -i -X POST -H "Content-Type:application/json" https://0b1ca0591fee4d7d882a5e691b780d1d-vp0.us.blockchain.ibm.com:5002/chaincode -d '{ "jsonrpc": "2.0", "method": "deploy", "params": { "type":1, "chaincodeID":{"path":"https://github.com/pong-pantola/expressway" }, "ctorMsg": { "function":"createTable", "args":[] }, "secureContext": "user_type1_3"}, "id": 1 }'

curl -i -X POST -H "Content-Type:application/json" https://0b1ca0591fee4d7d882a5e691b780d1d-vp0.us.blockchain.ibm.com:5002/chaincode -d '{ "jsonrpc": "2.0", "method": "query", "params": { "type":1, "chaincodeID":{ "name":"fe6d56e077bcac40f46db02413b88e226c5dafbceea98628b566d2fb2241ce90d2d1b25c500150d46451bfd116fbb62b005a043c9623d9fd10da4ce4a0aec2c0" }, "ctorMsg": {"function":"getAllVehicleBalance", "args":[] }, "secureContext": "user_type1_3" }, "id": 5 }'


curl -i -X POST -H "Content-Type:application/json" https://0b1ca0591fee4d7d882a5e691b780d1d-vp0.us.blockchain.ibm.com:5002/chaincode -d '{ "jsonrpc": "2.0", "method": "query", "params": { "type":1, "chaincodeID":{ "name":"c01032f0e2bfd365f2f8f3e0e103b331c98453d28fc3e9fe8d136413a1ff84294c6d0e20e2e932d5e6b1340ae9d66eb64ca63aa5867a4be01ee23528a4681807" }, "ctorMsg": {"function":"getAllVehicleViolation", "args":[] }, "secureContext": "user_type1_3" }, "id": 5 }'


curl -i -X POST -H "Content-Type:application/json" https://0b1ca0591fee4d7d882a5e691b780d1d-vp0.us.blockchain.ibm.com:5002/chaincode -d '{ "jsonrpc": "2.0", "method": "invoke", "params": { "type":1, "chaincodeID":{ "name":"c01032f0e2bfd365f2f8f3e0e103b331c98453d28fc3e9fe8d136413a1ff84294c6d0e20e2e932d5e6b1340ae9d66eb64ca63aa5867a4be01ee23528a4681807" }, "ctorMsg": {"function":"deleteAllVehicleViolation", "args":["AAA"] }, "secureContext": "user_type1_3" }, "id": 5 }'




OLD:



fe6d56e077bcac40f46db02413b88e226c5dafbceea98628b566d2fb2241ce90d2d1b25c500150d46451bfd116fbb62b005a043c9623d9fd10da4ce4a0aec2c0

NEW:

53aa257a55f950d1dd6a7381c04f86052c8f9ddc0a00f5223aa4ac8f5caf9b25f46c8102aa226f3ad316cfaf55b4aa2f4e63813abcb868b980dac2aa650c2ba0

INITIALIZE DEMO:
================

https://nr-expressway.mybluemix.net/initializeDemo

The endpoint above will do the following:
https://nr-expressway.mybluemix.net/deleteAllVehicleBalance
https://nr-expressway.mybluemix.net/insertVehicleBalance?PlateNo=ABC-123&Balance=30000
https://nr-expressway.mybluemix.net/insertVehicleBalance?PlateNo=XYZ-456&Balance=40000


https://nr-expressway.mybluemix.net/deleteAllExpresswayBalance
https://nr-expressway.mybluemix.net/insertExpresswayBalance?ExpresswayID=SLEX&Balance=0
https://nr-expressway.mybluemix.net/insertExpresswayBalance?ExpresswayID=NLEX&Balance=0
https://nr-expressway.mybluemix.net/insertExpresswayBalance?ExpresswayID=CAVITEX&Balance=0


https://nr-expressway.mybluemix.net/deleteAllExpresswayUsage


https://nr-expressway.mybluemix.net/deleteAllVehicleViolation





ENTRY:
======
https://nr-expressway.mybluemix.net/entryExpresswayUsage?PlateNo=ABC-123&ExpresswayID=SLEX&EntryTollGateID=MUNTINLUPA-SKYWAY


EXIT:
=====
https://nr-expressway.mybluemix.net/exitExpresswayUsage?PlateNo=ABC-123&ExpresswayID=SLEX&ExitTollGateID=NAIA3-SKYWAY

Sample Output:
{"PlateNo":"ABC-123","ExpresswayID":"SLEX","EntryDateTime":"2017-05-06 14-16-47","ExitDateTime":"2017-05-06 14-17-06"}




GET A PARTICULAR TOLL FEE:
==========================
https://nr-expressway.mybluemix.net/getExpresswayUsage?PlateNo=ABC-123&ExpresswayID=SLEX&EntryDateTime=2017-05-06%2014-16-47

Note:
-->Parameters will come from output of EXIT
-->Make sure to replace the space (" ") in the datetime parameters to %20


GET A PARTICULAR VIOLATION:
===========================
https://nr-expressway.mybluemix.net/getAllVehicleViolationByTimestamp?PlateNo=ABC-123&EntryDateTime=2017-05-06%2014-16-47&ExitDateTime=2017-05-06%2014-17-06

Note:
-->Parameters will come from output of EXIT
-->Make sure to replace the space (" ") in the datetime parameters to %20
 



OTHER GET (FOR DEBUGGING ONLY, NOT NEEDED IN ACTUAL APPLICATION):
=================================================================
https://nr-expressway.mybluemix.net/getVehicleBalance?PlateNo=ABC-123
https://nr-expressway.mybluemix.net/getAllVehicleBalance


https://nr-expressway.mybluemix.net/getExpresswayBalance?ExpresswayID=SLEX
https://nr-expressway.mybluemix.net/getAllExpresswayBalance


https://nr-expressway.mybluemix.net/getAllExpresswayUsage

https://nr-expressway.mybluemix.net/getAllVehicleViolation


SIMULATE IOT
============
https://nr-expressway.mybluemix.net/simulateIOT?PlateNo=ABC-123&speed=85&weather=rainy


SEND SMS
========
https://nr-expressway.mybluemix.net/sendSMS?message=THISISYOURMESSAGE


