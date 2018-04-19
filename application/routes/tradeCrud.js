module.exports = function(app) {

var log4js = require('log4js');
var logger = log4js.getLogger('Trade');
var helper = require('../app/helper.js');
var invoke = require('../app/invoke-transaction.js');
var query = require('../app/query.js');
var mymodule = require('./adminCrud.js');
var chaincodeName = "Trades"; // trade chaincode name
var channelName = "mychannel"; //channel name used in forceField
var peers;
var username = "pranjul";
var orgName = "uniper";
var test = mymodule.test;
var Client = require('fabric-client'); // Include the Hyperledger-SDK from npm
console.log('------------------------------------------------------'+test);
console.log(mymodule.test1);
console.log('-----------------------------------------------------------'+orgName+'---------------------------');

// syntax
//eventHub.registerChaincodeEvent(<chaincode ID>, <event name>, <callback>);

//Create trade
app.post('/forceField/trade', function(req, res) {
	console.log('------------forceField inside-----------------------------------------------'+orgName+'---------------------------');
	var fcn = req.body.fcn;
	var args = req.body.args;
	console.log(req);
	peers = ["peer1"];
    if (!chaincodeName) {
		res.json(getErrorMessage('\'chaincodeName\''));
		return;
	}
	if (!channelName) {
		res.json(getErrorMessage('\'channelName\''));
		return;
	}
	if (!fcn) {
		res.json(getErrorMessage('\'fcn\''));
		return;
	}
	if (!args) {
		res.json(getErrorMessage('\'args\''));
		return;
	}
    if (!args) {
            res.json(getErrorMessage('\'args\''));
            return;
    }
	console.log("logargs");
	console.log(args);
	console.log("username"+ " " + username + " " + "orgname" + " " +orgName)
	
    invoke.invokeChaincode(peers, channelName, chaincodeName, fcn, args, username, orgName)
    .then(function(message) {
        res.send(message);
    });
});
//Get trade By ID
app.get('/forceField/trade/:TRID', function(req, res) {
    var fcn = "gettradeByID";
    var args  =  JSON.parse(req.params.TRID);
    var peer = req.query.peer;
	if (!chaincodeName) {
		res.json(getErrorMessage('\'chaincodeName\''));
		return;
	}
	if (!channelName) {
		res.json(getErrorMessage('\'channelName\''));
		return;
	}
	if (!fcn) {
		res.json(getErrorMessage('\'fcn\''));
		return;
	}
	if (!args) {
		res.json(getErrorMessage('\'args\''));
		return;
	}
    if (!args) {
            res.json(getErrorMessage('\'args\''));
            return;
    }
	query.queryChaincode(peer, channelName, chaincodeName, args, fcn, req.username, req.orgname)
    .then(function(message) {
        res.send(message);
    });
});