import React from 'react';
import { Link } from 'react-router';
import Paper from 'material-ui/Paper';
import {Row,Col} from 'react-bootstrap';

const style={
  paperStyle:{
    width: "272px",
    height: "153px",
    marginLeft:"100px",
    borderRadius: "6px",
    backgroundColor:" #ffffff",
    border: "solid 1px #d5d5d5"
  },
  labelStyle:{
    width: "120px",
    height: "17px",
    fontFamily: "Helvetica",
    fontSize: "14px",
    fontWeight: "bold",
    fontStyle: "normal",
    fontStretch: "normal",
    lineHeight: "normal",
    letterSpacing: "normal",
    textAlign: "left",
    color: "#333333"
  },
  labelStyleNumber:{
    width: "67px",
    height: "58px",
    fontFamily: "Helvetica",
    fontSize: "48px",
    fontWeight: "300",
    fontStyle: "normal",
    fontStretch: "normal",
    lineHeight: "normal",
    letterSpacing: "normal",
    textAlign: "left",
    color: " #4a4a4a"
  },
  labelStyleSubText:{
    width: "38px",
    height: "19px",
    fontFamily: "Helvetica",
    fontSize: "16px",
    fontWeight: "bold",
    fontStyle: "normal",
    fontStretch: "normal",
    lineHeight: "normal",
    letterSpacing: "normal",
    textAlign: "left",
    color: " #333333"
  },
 
  sellText:{
    width: "52px",
    height: "19px",
    fontFamily: "Roboto",
    fontSize: "16px",
    fontWeight: "normal",
    fontStyle: "normal",
    fontStretch: "normal",
    lineHeight: "normal",
    letterSpacing: "normal",
    textAlign: "left",
    color: " #333333"
  },
  buyText:{
    width: "52px",
    height: "19px",
    fontFamily: "Roboto",
    fontSize: "16px",
    fontWeight: "normal",
    fontStyle: "normal",
    fontStretch: "normal",
    lineHeight: "normal",
    letterSpacing: "normal",
    textAlign: "left",
    color: " #333333"
  },
  sellRectange:{
    float:"left",
    width: "20px",
  height: "20px",
  backgroundColor: "#ffd403"
  },
  buyRectange:{
    float:"left",
    width: "20px",
  height: "20px",
  backgroundColor:"#52d05a"
  },
  
}

export default class ConfirmTradeComponent extends React.Component {
  
  render() {
      return (
        
           <div style={style.paperStyle} zDepth={2} >
           <div style={{marginLeft:"20px"}}>
        <span style={style.labelStyle}>Confirmed Trades </span>
        <br />
        <br />
        <span style={style.labelStyleNumber} >25 </span>
        <span style={style.labelStyleSubText}>Total </span>
        <br/>
        <Row>
        <Col xs={4} >
        <span style={style.sellRectange}  >
              </span>
              <span style={style.sellText} > 16 Sell </span>
      </Col>
      <Col xs={4}>
      <span style={style.buyRectange}>
              </span>
              <span style={style.buyText}> 14 Buy </span>  
      </Col>
                       
            </Row>
            </div>
           </div>
         
      )
    }
  }



