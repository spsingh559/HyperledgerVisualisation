module.exports = function(app) {

var log4js = require('log4js');
var logger = log4js.getLogger('Product');
var helper = require('../app/helper.js');
var invoke = require('../app/invoke-transaction.js');
var query = require('../app/query.js');
var mymodule = require('./adminCrud.js');
var chaincodeName = "Products"; // Product chaincode name
var channelName = "mychannel"; //channel name used in SSLNG
var peers;
var username = "uniper1";
var orgName = "org1";
var test = mymodule.test;
console.log('------------------------------------------------------'+test);
console.log(mymodule.test1);

console.log('-----------------------------------------------------------'+orgName+'---------------------------');

//Create Product
app.post('/sslng/product', function(req, res) {
	console.log('------------sslng inside-----------------------------------------------'+orgName+'---------------------------');
	var fcn = req.body.fcn;
	 var args = req.body.args;
	 console.log(req);
	peers = ["peer1"];
// var args = [req.body.prid,
// 	 req.body.supplier_name,
// 	 req.body.source_location_city,
// 	 req.body.source_location_country,
// 	 req.body.volume,
// 	  req.body.supplier_cost,
// 		 req.body.supplier_load_date,
// 		 req.body.tranporter_name,
// 		 req.body.transporter_cost,
// 		  req.body.container_load_date,
// 			 req.body.load_port,
// 			 req.body.container_discharge_date,
// 			 req.body.discharge_port,
// 			  req.body.destination_location_country,
// 				req.body.destination_location_city,
// 				req.body.destination_date,
// 				req.body.created_by,
// 				req.body.updated_by,
// 				req.body.last_update_timestamp];
//var args=["PR1","sup1","EC","IND","100","1000","13/12/2017","tr1","1000", "14/12/2017","MU","15/12/2017","BY","BG","RGT","16/2017","cus1","cus1","12/12/2017"]
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
console.log("username"+ " " + username + " " + "orgname" + "hguuyuuguyfuyfuyt" + " " +orgName)
        invoke.invokeChaincode(peers, channelName, chaincodeName, fcn, args, username, orgName)
        .then(function(message) {
                res.send(message);
        });


});
//Update Product
app.post('/sslng/product/:PRID', function(req, res) {
	var fcn = "updateProduct";
	// var args = req.body.args;
	var peers = req.body.peers;
	var args = [prid, req.body.vesselName, req.body.supplier_name, req.body.source_location_city, req.body.source_location_country, req.body.volume, req.body.supplier_cost, req.body.supplier_load_date, req.body.tranporter_name, req.body.transporter_cost, req.body.container_load_date, req.body.load_port, req.body.container_discharge_date, req.body.discharge_port, req.body.destination_location_country, req.body.destination_location_city,req.body.destination_date,req.body.product_status,req.body.created_by,req.body.created_by,req.body.updated_by,req.body.last_update_timestamp];
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

        invoke.invokeChaincode(peers, channelName, chaincodeName, fcn, args, req.username, req.orgname)
        .then(function(message) {
                res.send(message);
        });


});

//Get Product By ID
app.get('/sslng/product/:PRID', function(req, res) {
        var fcn = "getProductByID";
        var args  =  JSON.parse(req.params.PRID);
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
// Get All Product
app.get('/ssnlg/products', function(req, res) {
	logger.debug('==================== QUERY BY CHAINCODE ==================');

	var args = req.query.args;
	var fcn = "getAllProducts";
	var peer = req.query.peer;

	logger.debug('channelName : ' + channelName);
	logger.debug('chaincodeName : ' + chaincodeName);
	logger.debug('fcn : ' + fcn);
	logger.debug('args : ' + args);

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
	args = args.replace(/'/g, '"');
	args = JSON.parse(args);
	logger.debug(args);

	query.queryChaincode(peer, channelName, chaincodeName, args, fcn, req.username, req.orgname)
	.then(function(message) {
		res.send(message);
	});
});

// for testing only - can be removed
app.post('/channels/:channelName/chaincodes/:chaincodeName', function(req, res) {
	logger.debug('==================== INVOKE ON CHAINCODE ==================');
	var fcn = "move";
	var args = req.body.args;
	logger.debug('channelName  : ' + channelName);
	logger.debug('chaincodeName : ' + chaincodeName);
	logger.debug('fcn  : ' + fcn);
	logger.debug('args  : ' + args);
	if (!args) {
		res.json(getErrorMessage('\'args\''));
		return;
	}

	invoke.invokeChaincode(peers, channelName, chaincodeName, fcn, args, req.username, req.orgname)
	.then(function(message) {
		res.send(message);
	});
});
// for testing - can be removed
app.get('/channels/:channelName/chaincodes/:chaincodeName', function(req, res) {
	logger.debug('==================== QUERY BY CHAINCODE ==================');
	var channelName = req.params.channelName;
	var chaincodeName = req.params.chaincodeName;
	let args = req.query.args;
	let fcn = req.query.fcn;
	let peer = req.query.peer;

	logger.debug('channelName : ' + channelName);
	logger.debug('chaincodeName : ' + chaincodeName);
	logger.debug('fcn : ' + fcn);
	logger.debug('args : ' + args);

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
	args = args.replace(/'/g, '"');
	args = JSON.parse(args);
	logger.debug(args);

	query.queryChaincode(peer, channelName, chaincodeName, args, fcn, req.username, req.orgname)
	.then(function(message) {
		res.send(message);
	});
});



}
