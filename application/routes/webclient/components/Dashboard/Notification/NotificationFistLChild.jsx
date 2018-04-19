import React from 'react';
import {List, ListItem} from 'material-ui/List';
import Subheader from 'material-ui/Subheader';

const style={
    listStyle:{
    width: "270px",
  height: "93px",
  backgroundColor: "#ebf0f0"
    }
}

import EachNotificationRowComonent from './EachNotificationRowComonent.jsx';

export default class NotificationFistLChild extends React.Component{

    render(){
console.log('----------Data reach to first L Child------------');
console.log(this.props.notificationData);

    let newData = this.props.notificationData.map((data)=>{
 return(
 <EachNotificationRowComonent
 key={data._id}
  notificationText={data.notificationText}
  timeStamp={data.timeStamp}
  action={data.action}
  actionName={data.actionName}
 />
 )
});

        return(
            <List >
                 <Subheader>Notifications</Subheader>
                {newData}
                </List>
        )
    }
}