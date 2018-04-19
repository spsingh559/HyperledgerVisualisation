var express = require('express');
var mongoose = require('mongoose');
mongoose.connect('mongodb://localhost:27017/TimeManagmentApp');

var taskSchema = new mongoose.Schema({

    _id:Date,
    taskName:String,
    taskDate:String,
    taskTime:String,
    taskStatus:String,
    taskPriority:String
});

var taskModel = mongoose.model('taskcollections', taskSchema);

exports = module.exports = taskModel;
