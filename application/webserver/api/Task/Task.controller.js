var controller={};
var TaskModel=require('./Task.Model');

// controller.LoginData=function(req,res){
//   var d= req.body.ID;
//   var e=req.body.password;
//   console.log("userId: "+d+" password: "+e);
//   //var flag=0;
//   empapplications.findOne({ID:d,password:e}).exec(function(err,data){
//           console.log(data);
//            if(data){
//            res.send({message:'success'});

//            }
//            else{
//            res.send({message:'fail'});
//            }

//        });

//   }


// exports = module.exports = controller;

controller.postData=function(req,res){
  TaskModel.create(req.body,function(err,data){
          if(err) { console.log('server error in post'+err); }
          else{
           res.json({message:data});
          }
        });

}

controller.getTaskDetail=function(req,res){
  console.log(req.params.timeStamp);
  TaskModel.find({taskDate:req.params.timeStamp}).exec(function(err,data){
 if(err) { console.log('server error in get'+err); }
          else{
           res.json({message:data});
           console.log('result of get operation is');
           console.log(data);
          }
        })

}

controller.getSearchResult=function(req,res){
  console.log(req.params.searchName);
  TaskModel.find({taskName:req.params.searchName}).exec(function(err,data){
 if(err) { console.log('server error in get'+err); }
          else{
           res.json({message:data});
           console.log('result of get operation is');
           console.log(data);
          }
        })

}

controller.removeTask=function(req,res){
  console.log(req.params._id);
  TaskModel.remove({_id:req.params._id},function(err, data){
            if(err)
            console.log('server error in remove'+ err);
            else
              res.json({message:data});
  });
}

controller.updateTask=function(req,res){
    console.log('Task api connected for update');
    console.log(req.body);
    // var data=req.body._id;
    // console.log(data);

    TaskModel.update({_id:req.body._id},req.body,function(err,data){
      if(err){console.log('server error in update of CellSummary'+ err);}
      else{
        res.json({message:data});
      }
    })
};


exports = module.exports = controller;