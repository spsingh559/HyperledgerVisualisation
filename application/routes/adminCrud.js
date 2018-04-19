module.exports= function(app) {


'use strict';
var log4js = require('log4js');
var logger = log4js.getLogger('SampleWebApp');
//var express = require('express');
var session = require('express-session');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
var http = require('http');
var util = require('util');
//var app = express();
var expressJWT = require('express-jwt');
var jwt = require('jsonwebtoken');
var bearerToken = require('express-bearer-token');
var cors = require('cors');
// require('./productCrud.js');
var MongoClient = require('mongodb').MongoClient;
var ObjectIdVar = require('mongodb').ObjectID;
require('../config.js');
var hfc = require('fabric-client');
var dbUrl = 'mongodb://localhost:27017/sslng';
var helper = require('../app/helper.js');
var channels = require('../app/create-channel.js');
var join = require('../app/join-channel.js');
var install = require('../app/install-chaincode.js');
var instantiate = require('../app/instantiate-chaincode.js');
var invoke = require('../app/invoke-transaction.js');
var query = require('../app/query.js');
var host = process.env.HOST || hfc.getConfigSetting('host');
var port = process.env.PORT || hfc.getConfigSetting('port');
var username;
var orgName;
var multipart = require('connect-multiparty');
var multipartMiddleware = multipart();
///////////////////////////////////////////////////////////////////////////////
///////////////////////// REST ENDPOINTS START HERE ///////////////////////////
///////////////////////////////////////////////////////////////////////////////
// Register and enroll user
/*app.post('/users', function(req, res) {
	 username = req.body.username;
	 orgName = req.body.orgName;
	 module.exports.orgName = orgName;

	logger.debug('End point : /users');
	logger.debug('User name : ' + username);
	logger.debug('Org name  : ' + orgName);
	if (!username) {
		res.json(getErrorMessage('\'username\''));
		return;
	}
	if (!orgName) {
		res.json(getErrorMessage('\'orgName\''));
		return;
	}
	var token = jwt.sign({
		exp: Math.floor(Date.now() / 1000) + parseInt(hfc.getConfigSetting('jwt_expiretime')),
		// username: username,
		// orgName: orgName
	}, app.get('secret'));
	helper.getRegisteredUsers(username, orgName, true).then(function(response) {
		if (response && typeof response !== 'string') {
			response.token = token;
			res.json(response);
		} else {
			res.json({
				success: false,
				message: response
			});
		}
	});
});*/

app.post('/userLogin', multipartMiddleware, function(req, res, next) {

		console.log(req.body);

	 MongoClient.connect(dbUrl, function(err, db) {
        console.log("connected to DB gfgf")
        if (err) throw err;
        var collection = db.collection('myuser');

        collection.findOne({

            username: req.body.username
        },

function(err, result) {
            if (result != null) {
                console.log(result.password);
console.log(req.body.password);
                if (result.password == req.body.password) {
                    req.session.username = result.username;
                    req.session.role = result.role;
                    req.session.orgName = result.orgName;
                     console.log("userLogin " + " role of user ="+ result.role + " username = "+result.username);
                    console.log("password  matched");
                    if (result.role == "Uniper") {
                    	console.log("inside role = " +result.username);
                    	console.log("inside role = " +result.orgName);
                        req.session.username = result.username;
                        req.session.orgName = result.orgName;
                        username = req.session.username;
                        orgName = req.session.orgName;
                        logger.debug('End point : /userLogin');
	logger.debug('User name : ' + username);
	logger.debug('Org name  : ' + orgName);
	if (!username) {
		res.json(getErrorMessage('\'username\''));
		return;
	}
	if (!orgName) {
		res.json(getErrorMessage('\'orgName\''));
		return;
	}
	// var token = jwt.sign({
	// 	exp: Math.floor(Date.now() / 1000) + parseInt(hfc.getConfigSetting('jwt_expiretime')),
  //
	// }, app.get('secret'));

  //     helper.getRegisteredUsers(username, orgName, true).then(function(response) {
  //                       	console.log("getRegisteredUsers "+  username);
	// 	if (response && typeof response !== 'string' ) {
	// 		response.token = token;
  //
	// 		 res.redirect('/uniper_index#/uniper');
	// 		 console.log(response);
	// 	} else {
	// 		res.json({
	// 			success: false,
	// 			message: response
	// 		});
	// 	}
	// });
 res.redirect('/uniper_index#/uniper');
                    }

	else  if (result.role == "Doctor") {
                    	console.log("inside role = " +result.username);
                    	console.log("inside role = " +result.orgName);
                        req.session.username = result.username;
                        req.session.orgName = result.orgName;
                        username = req.session.username;
                        orgName = req.session.orgName;
                        logger.debug('End point : /userLogin');
	logger.debug('User name : ' + username);
	logger.debug('Org name  : ' + orgName);
	if (!username) {
		res.json(getErrorMessage('\'username\''));
		return;
	}
	if (!orgName) {
		res.json(getErrorMessage('\'orgName\''));
		return;
	}
	// var token = jwt.sign({
	// 	exp: Math.floor(Date.now() / 1000) + parseInt(hfc.getConfigSetting('jwt_expiretime')),
  //
	// }, app.get('secret'));

  /*    helper.getRegisteredUsers(username, orgName, true).then(function(response) {
                        	console.log("getRegisteredUsers "+  username);
		if (response && typeof response !== 'string' ) {
			response.token = token;

			 res.redirect('/Doctor_index#/Doctor_Dashboard');
			 console.log(response);
		} else {
			res.json({
				success: false,
				message: response
			});
		}
	});*/

                    }

                    db.close();
                } else {
                    console.log("password do not match");
                    res.redirect('/login?valid=y');

                }
            } else {
                console.log("username do not match");

                res.redirect('/login?valid=y');

                console.log(err);
            }
        });
    });

});
// Create Channel
app.post('/channels', function(req, res) {
	logger.info('<<<<<<<<<<<<<<<<< C R E A T E  C H A N N E L >>>>>>>>>>>>>>>>>');
	logger.debug('End point : /channels');
	var channelName = req.body.channelName;
	var channelConfigPath = req.body.channelConfigPath;
	logger.debug('Channel name : ' + channelName);
	logger.debug('channelConfigPath : ' + channelConfigPath); //../artifacts/channel/mychannel.tx
	if (!channelName) {
		res.json(getErrorMessage('\'channelName\''));
		return;
	}
	if (!channelConfigPath) {
		res.json(getErrorMessage('\'channelConfigPath\''));
		return;
	}

	channels.createChannel(channelName, channelConfigPath, req.username, req.orgname)
	.then(function(message) {
		res.send(message);
	});
});
// Join Channel
app.post('/channels/:channelName/peers', function(req, res) {
	logger.info('<<<<<<<<<<<<<<<<< J O I N  C H A N N E L >>>>>>>>>>>>>>>>>');
	var channelName = req.params.channelName;
	var peers = req.body.peers;
	logger.debug('channelName : ' + channelName);
	logger.debug('peers : ' + peers);
	if (!channelName) {
		res.json(getErrorMessage('\'channelName\''));
		return;
	}
	if (!peers || peers.length == 0) {
		res.json(getErrorMessage('\'peers\''));
		return;
	}

	join.joinChannel(channelName, peers, req.username, req.orgname)
	.then(function(message) {
		res.send(message);
	});
});
// Install chaincode on target peers
app.post('/chaincodes', function(req, res) {
	logger.debug('==================== INSTALL CHAINCODE ==================');
	var peers = req.body.peers;
	var chaincodeName = req.body.chaincodeName;
	var chaincodePath = req.body.chaincodePath;
	var chaincodeVersion = req.body.chaincodeVersion;
	logger.debug('peers : ' + peers); // target peers list
	logger.debug('chaincodeName : ' + chaincodeName);
	logger.debug('chaincodePath  : ' + chaincodePath);
	logger.debug('chaincodeVersion  : ' + chaincodeVersion);
	if (!peers || peers.length == 0) {
		res.json(getErrorMessage('\'peers\''));
		return;
	}
	if (!chaincodeName) {
		res.json(getErrorMessage('\'chaincodeName\''));
		return;
	}
	if (!chaincodePath) {
		res.json(getErrorMessage('\'chaincodePath\''));
		return;
	}
	if (!chaincodeVersion) {
		res.json(getErrorMessage('\'chaincodeVersion\''));
		return;
	}

	install.installChaincode(peers, chaincodeName, chaincodePath, chaincodeVersion, req.username, req.orgname)
	.then(function(message) {
		res.send(message);
	});
});
// Instantiate chaincode on target peers
app.post('/channels/:channelName/chaincodes', function(req, res) {
	logger.debug('==================== INSTANTIATE CHAINCODE ==================');
	var chaincodeName = req.body.chaincodeName;
	var chaincodeVersion = req.body.chaincodeVersion;
	var channelName = req.params.channelName;
	var fcn = req.body.fcn;
	var args = req.body.args;
	logger.debug('channelName  : ' + channelName);
	logger.debug('chaincodeName : ' + chaincodeName);
	logger.debug('chaincodeVersion  : ' + chaincodeVersion);
	logger.debug('fcn  : ' + fcn);
	logger.debug('args  : ' + args);
	if (!chaincodeName) {
		res.json(getErrorMessage('\'chaincodeName\''));
		return;
	}
	if (!chaincodeVersion) {
		res.json(getErrorMessage('\'chaincodeVersion\''));
		return;
	}
	if (!channelName) {
		res.json(getErrorMessage('\'channelName\''));
		return;
	}
	if (!args) {
		res.json(getErrorMessage('\'args\''));
		return;
	}
	instantiate.instantiateChaincode(channelName, chaincodeName, chaincodeVersion, fcn, args, req.username, req.orgname)
	.then(function(message) {
		res.send(message);
	});
});
console.log("admin "+username);
console.log("adminOrg "+orgName);
// module.exports.username = username;
// module.exports.orgName = orgName;
module.exports.test = "test";
module.exports.test1 = "test1";

}
