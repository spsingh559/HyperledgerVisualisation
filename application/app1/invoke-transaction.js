/**
 * Copyright 2017 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */
'use strict';
var path = require('path');
var fs = require('fs');
var util = require('util');
var hfc = require('fabric-client');
var Peer = require('fabric-client/lib/Peer.js');
var helper = require('./helper.js');
var logger = helper.getLogger('invoke-chaincode');
var EventHub = require('fabric-client/lib/EventHub.js');
var ORGS = hfc.getConfigSetting('network-config');

var invokeChaincode = function(peerNames, channelName, chaincodeName, fcn, args, username, org,socket) {
	logger.debug(util.format('\n============ invoke transaction on organization %s ============\n', org));
	var client = helper.getClientForOrg(org);
	logger.debug(util.format('\n============ client %s ============\n', client));
	console.log(client);
	var channel = helper.getChannelForOrg(org);
	logger.debug(util.format('\n============ chennel %s ============\n', channel));
	console.log(channel);
	var targets = (peerNames) ? helper.newPeers(peerNames, org) : undefined;
	logger.debug(util.format('\n============ targets %s ============\n', targets));
	console.log(targets)
	var tx_id = null;
	// client='gins';
	// var someValue = helper.getChannelForOrg(org);
	// console.log("got the value", someValue);
	// var obj={
	// 	name:'shyam'
	// }
	// if(typeof channel === "object") {
	// 	console.log("channel is an object", channel);
	// 	socket.emit('pranjul',obj);
	// }
	// else {
	// 	console.log("channel is not an object", typeof channel);
	// 	socket.emit('pranjul', "Shyaaaam");
	// }
	
	// console.log("after emit");
	// logger.debug(util.format('\n============ after socket %s ============\n', targets));
	

	return helper.getRegisteredUsers(username, org).then((user) => {
		tx_id = client.newTransactionID();
		logger.debug(util.format('Sending transaction "%j"', tx_id));
		// send proposal to endorser
		var request = {
			chaincodeId: chaincodeName,
			fcn: fcn,
			args: args,
			chainId: channelName,
			txId: tx_id
		};

		if (targets)
			request.targets = targets;
		console.log('----------------------Send tx proposal-------------------');
		console.log(channel.sendTransactionProposal(request));
		return channel.sendTransactionProposal(request);
	}, (err) => {
		logger.error('Failed to enroll user \'' + username + '\'. ' + err);
		throw new Error('Failed to enroll user \'' + username + '\'. ' + err);
	}).then((results) => {
		console.log('----------------------Proposal Response-------------------');
		console.log(results);
		var proposalResponses = results[0];
		var proposal = results[1];
		var all_good = true;
		for (var i in proposalResponses) {
			let one_good = false;
			if (proposalResponses && proposalResponses[i].response &&
				proposalResponses[i].response.status === 200) {
				one_good = true;
				logger.info('transaction proposal was good');
			} else {
				logger.error('transaction proposal was bad');
			}
			all_good = all_good & one_good;
		}
		if (all_good) {

			console.log('All good');
			console.log(all_good);
			logger.debug(util.format(
				'Successfully sent Proposal and received ProposalResponse: Status - %s, message - "%s", metadata - "%s", endorsement signature: %s',
				proposalResponses[0].response.status, proposalResponses[0].response.message,
				proposalResponses[0].response.payload, proposalResponses[0].endorsement
				.signature));
			var request = {
				proposalResponses: proposalResponses,
				proposal: proposal
			};
			// set the transaction listener and set a timeout of 30sec
			// if the transaction did not get committed within the timeout period,
			// fail the test
			var transactionID = tx_id.getTransactionID();
			console.log(' -------------------generated tx id is---------------------');
			console.log(transactionID)
			var eventPromises = [];

			if (!peerNames) {
				peerNames = channel.getPeers().map(function(peer) {
					return peer.getName();
				});
			}

			var eventhubs = helper.newEventHubs(peerNames, org);

			console.log('---------------eventhubs---------------------------');
			console.log(eventhubs);
			for (let key in eventhubs) {
				let eh = eventhubs[key];
				eh.connect();

				let txPromise = new Promise((resolve, reject) => {
					let handle = setTimeout(() => {
						eh.disconnect();
						reject();
					}, 30000);

					eh.registerTxEvent(transactionID, (tx, code) => {
						clearTimeout(handle);
						eh.unregisterTxEvent(transactionID);
						eh.disconnect();

						if (code !== 'VALID') {
							console.log(' -----------------tx valid------------------------');
							logger.error(
								'The SSLNG transaction was invalid, code = ' + code);
							reject();
						} else {
							console.log('------------------------tx commited on peer-----------------')
							logger.info(
								'The SSLNG transaction has been committed on peer ' +
								eh._ep._endpoint.addr);
							resolve();
						}
					});
				});
				eventPromises.push(txPromise);
			};
			var sendPromise = channel.sendTransaction(request);
			console.log('-----------------sendPromise-------------------');
			console.log(sendPromise);
			return Promise.all([sendPromise].concat(eventPromises)).then((results) => {
				logger.debug(' event promise all complete and testing complete');
				// socket.emit('pranjul', "client");
				return results[0]; // the first returned value is from the 'sendPromise' which is from the 'sendTransaction()' call
			}).catch((err) => {
				logger.error(
					'Failed to send transaction and get notifications within the timeout period.'
				);
				return 'Failed to send transaction and get notifications within the timeout period.';
			});
		} else {

			console.log('-----------------Propsal Response not good-------------------');
			logger.error(
				'Failed to send Proposal or receive valid response. Response null or status is not 200. exiting...'
			);
			return 'Failed to send Proposal or receive valid response. Response null or status is not 200. exiting...';
		}
	}, (err) => {
		logger.error('Failed to send proposal due to error: ' + err.stack ? err.stack :
			err);
		return 'Failed to send proposal due to error: ' + err.stack ? err.stack :
			err;
	}).then((response) => {
		console.log('----------------response -------------------');
		console.log(response);
		if (response.status === 'SUCCESS') {
			logger.info('Successfully sent transaction to the orderer.');
			// var someValue = helper.getChannelForOrg(org);
			// console.log("got the value", someValue);
			// socket.emit('pranjul',someValue);
			// console.log("after emit");
			return tx_id.getTransactionID();
		} else {
			logger.error('Failed to order the transaction. Error code: ' + response.status);
			return 'Failed to order the transaction. Error code: ' + response.status;
		}
	}, (err) => {
		logger.error('Failed to send transaction due to error: ' + err.stack ? err
			.stack : err);
		return 'Failed to send transaction due to error: ' + err.stack ? err.stack :
			err;
	});
};

exports.invokeChaincode = invokeChaincode;
