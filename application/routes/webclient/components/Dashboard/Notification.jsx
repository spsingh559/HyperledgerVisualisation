import React from 'react';
import { Link } from 'react-router';
import Paper from 'material-ui/Paper';
import NotificationFistLChild from './Notification/NotificationFistLChild.jsx';

const style={
  paperStyle:{
  width: "auto",
  height: "476px",
  borderRadius: "6px",
  backgroundColor: "#ffffff",
  border: "solid 1px #d5d5d5",
  overflowY:"auto"
  }
}

export default class Notification extends React.Component {

  state={
    notificationData:[
      {
      _id:"1274230894",
      notificationText:"TR011 is confirmed and waiting for Your Approval",
      timeStamp:"03 Jan 2017 | 2:00 PM",
      action:true,
      actionName:"Review"
    },
    { _id:"1274230895",
      notificationText:"TR032 is complete",
      timeStamp:"03 Jan 2017 | 1:56 PM",
      action:false,
      actionName:"Review"
    },
    {
      _id:"1274230897",
      notificationText:"TR094 status changed to ship",
      timeStamp:"03 Jan 2017 | 2:00 PM",
      action:true,
      actionName:"Track"
    },
    {
      _id:"1274230898",
      notificationText:"TR094 status changed to ship",
      timeStamp:"03 Jan 2017 | 2:00 PM",
      action:true,
      actionName:"Track"
    },
    {
      _id:"1274230899",
      notificationText:"TR094 status changed to ship",
      timeStamp:"03 Jan 2017 | 2:00 PM",
      action:true,
      actionName:"Track"
    },
    {
      _id:"1274230100",
      notificationText:"TR094 status changed to ship",
      timeStamp:"03 Jan 2017 | 2:00 PM",
      action:true,
      actionName:"Track"
    },
    {
      _id:"1274230101",
      notificationText:"TR094 status changed to ship",
      timeStamp:"03 Jan 2017 | 2:00 PM",
      action:true,
      actionName:"Track"
    }




    ]
      
    
  }
  
  render() {
      return (
        
          <div style={style.paperStyle} >
       <NotificationFistLChild notificationData={this.state.notificationData} />
           </div>
         
      )
    }
  }



