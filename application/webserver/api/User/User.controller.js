var controller={};
var UserModel=require('./User.Model');

controller.loginData=function(req,res){
  UserModel.find({userName:req.body.userName,userPassword:req.body.userPassword}).exec(function(err,data){
          console.log(data);
           if(data){
           res.send({message:'success'});

           }
           else{
           res.send({message:'fail'});
           }

       });
}

controller.signUpData=function(req,res){
  UserModel.create(req.body,function(err,data){
          if(err) { console.log('server error in post'+err); }
          else{
           res.json({message:data});
          }
        });
}

exports = module.exports = controller;
