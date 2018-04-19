import React from 'react';
import { Link } from 'react-router';
import Paper from 'material-ui/Paper';
import NotificationFistLChild from './Notification/NotificationFistLChild.jsx';

const style={
  paperStyle:{
  width: "300px",
  height: "476px",
  marginLeft:"100px",
  borderRadius: "6px",
  backgroundColor: "#ffffff",
  border: "solid 1px #d5d5d5",
  overflowY:"auto"
  }
}

export default class Notification extends React.Component {

  state={
    notificationData:[
      { _id:"1274230895",
      notificationText:"User is enrolled for Org 1",
      timeStamp:"03 Jan 2017 | 1:56 PM",
     
    },
    {
      _id:"1274230897",
      notificationText:"Channel is created by admin user",
      timeStamp:"03 Jan 2017 | 2:00 PM",
     
    },
      {
      _id:"1274230894",
      notificationText:"Chaincode is deployed on channel Mychannel",
      timeStamp:"03 Jan 2017 | 2:00 PM",
    },
    
    
    {
      _id:"1274230898",
      notificationText:"Peer1 has join the Mychannel",
      timeStamp:"03 Jan 2017 | 2:00 PM",
    },
   




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



