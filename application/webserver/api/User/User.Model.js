var express = require('express');
var mongoose = require('mongoose');
mongoose.connect('mongodb://localhost:27017/TimeManagmentApp');

var userSchema = new mongoose.Schema({
    _id:Date,
    userName:String,
    userPassword:String,
    userContactNumber:String
});

var userModel = mongoose.model('logindetails', userSchema);

exports = module.exports = userModel;
