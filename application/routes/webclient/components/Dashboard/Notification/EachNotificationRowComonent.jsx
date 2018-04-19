import React from 'react';
import {List, ListItem} from 'material-ui/List';
import Divider from 'material-ui/Divider';

const style={
    actionNameStyle:{
    width: "56px",
  height: "17px",
  fontFamily: "Helvetica",
  fontSize: "14px",
  fontWeight: "bold",
  fontStyle: "normal",
  fontStretch: "normal",
  lineHeight: "normal",
  letterSpacing: "normal",
  textAlign: "left",
  color: "#1f497d",
    },
    primaryTextStyle:{
        width: "235px",
  height: "38px",
  fontFamily: "Roboto",
  fontSize: "16px",
  fontWeight: "normal",
  fontStyle: "normal",
  fontStretch: "normal",
  lineHeight: "normal",
  letterSpacing: "normal",
  textAlign: "left",
  color: " #333333",
       
    }
}

export default class EachNotificationRowComonent extends React.Component{

    render(){

        return(
            <div>
            <ListItem
         
          primaryText={this.props.notificationText}
          
          secondaryText={
            <p>
              <span style={{color: "darkBlack"}}>{this.props.timeStamp} </span>
            <span style={style.actionNameStyle}>{this.props.action?this.props.actionName:null}</span>
            </p>
          }
          secondaryTextLines={2}
        />
        <Divider inset={true} />
        </div>
        )
    }
}