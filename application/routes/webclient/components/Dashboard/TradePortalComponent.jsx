import React from 'react';
import { Link } from 'react-router';
import Paper from 'material-ui/Paper';
import {List, ListItem} from 'material-ui/List';
import Subheader from 'material-ui/Subheader';
import Divider from 'material-ui/Divider';


import {makeSelectable} from 'material-ui/List';
import Avatar from 'material-ui/Avatar';
import FileFolder from 'material-ui/svg-icons/file/folder';
import ActionInfo from 'material-ui/svg-icons/action/info';

const style={
  paperStyle:{
    width: "272px",
  height: "476px",
  borderRadius: "6px",
  backgroundColor: "#ffffff",
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
}
const shellTrader={
  boxsize:{
      width: '272px',
      height: '476px',
      bordeRadius: '6px',
      backgroundColor: '#ffffff',
      border: 'solid 1px #d5d5d5',
      marginLeft: '20px',
      marginBottom: '20px'
  },
  tradePortfolio:{
      fontFamily: 'Helvetica',
      fontSize: '14px',
      fontWeight: 'bold',
      fontStyle: 'normal',
      fontStretch: 'normal',
      lineHeight: 'normal',
      letterSpacing: 'normal',
      textAlign: 'left',
      color: '#333333'
  },
  topBox:{  
    width: '270px',
    height: '85',
    border:'solid 1px #d5d5d5'
  },
  priceTrade:{  
    width: '270px',
    height: '85px',
    backgroundColor: '#e5e8e8'
  },
  leftBrent:{
    fontFamily: 'Helvetica',
    fontSize: '20px',
    fontWeight: 'bold',
    fontStyle: 'normal',
    fontStretch: 'normal',
    lineHeight: 'normal',
    letterSpacing: 'normal',
    textAlign: 'left',
    color: '#1f497d',
    paddingLeft:'15px'
  }, 
  rightBBL:{
      fontFamily: 'Roboto',
      fontSize: '20px',
      fontWeight: 'bold',
      fontStyle: 'normal',
      fontStretch: 'normal',
      lineHeight: 'normal',
      letterSpacing: 'normal',
      textAlign: 'right',
      color: '#333333',
      marginLeft:'60px'
  },
  priceTradeBBLS:{  
    width: '270px',
    height: '85px',
    backgroundColor: '#ebf0f0'
  },
  KBBLS:{
    fontFamily: 'Helvetica',
    fontSize: '17px',
    fontWeight: 'bold',
    fontStyle: 'normal',
    fontStretch: 'normal',
    lineHeight: 'normal',
    letterSpacing: 'normal',
    textAlign: 'left',
    color: '#333333',
    paddingLeft:'15px'
  },
  
  KBBLStext: {
    fontFamily: 'Helvetica',
    fontStyle: 'normal',
    fontStretch: 'normal',
    lineHeight: 'normal',
    letterSpacing: 'normal',
    textAlign: 'left',
    color: '#333333',
    paddingLeft:'15px',
    fontSize: '24px',
    fontWeight: 'normal'
    },
    LONG: {
      fontFamily: 'Helvetica',
      fontSize: '17px',
      fontWeight: 'normal',
      fontStyle: 'normal',
      fontStretch: 'normal',
      lineHeight: 'normal',
      letterSpacing: 'normal',
      textAlign: 'right',
      color: '#417505',
      float: 'right',
      marginRight: '15px'
    },
    LONGRED: {
      fontFamily: 'Helvetica',
      fontSize: '17px',
      fontWeight: 'normal',
      fontStyle: 'normal',
      fontStretch: 'normal',
      lineHeight: 'normal',
      letterSpacing: 'normal',
      textAlign: 'right',
      color: '#d0011b',
      float: 'right',
      marginRight: '25px'
    },
    AvailableCreditFor: {
      fontFamily: 'Helvetica',
      fontSize: '16px',
      fontWeight: 'bold',
      fontStyle: 'normal',
      fontStretch: 'normal',
      lineHeight: 'normal',
      letterSpacing: 'normal',
      textAlign: 'left',
      color: '#333333'
    },
    mn:{
      fontFamily: 'Helvetica',
      fontSize: '24px',
      fontWeight: 'normal',
      fontStyle: 'normal',
      fontStretch: 'normal',
      lineHeight: 'normal',
      letterSpacing: 'normal',
      color:'#417505',
      float: 'right'
    },
    creditUtilization:{
      fontFamily: 'Helvetica',
      fontSize: '16px',
      fontWeight: 'bold',
      fontStyle: 'normal',
      fontStretch: 'normal',
      lineHeight: 'normal',
      letterSpacing: 'normal',
      color: '#000000'
    },
    leftText:{
      fontFamily: 'Helvetica',
      fontSize: '16px',
      fontWeight: '300',
      fontStyle: 'normal',
      fontStretch: 'normal',
      lineHeight: 'normal',
      letterSpacing: 'normal',
      textAlign: 'left',
      color: '#333333',
      paddingLeft:'15px'
    },
    textRight:{
      fontFamily: 'Helvetica',
      fontSize: '16px',
      fontWeight: 'normal',
      fontStyle: 'normal',
      fontStretch: 'normal',
      lineHeight: 'normal',
      letterSpacing: 'normal',
      color: '#417505',
      float:'right',
      marginRight: '15px'
    }
}
export default class TradePortalComponent extends React.Component {
  
  render() {
      return (
       
        <div >
        <div style={shellTrader.boxsize}> 
        <div style={shellTrader.topBox}>
        <Subheader style={{fontWeight: 'bold',color: '#333333'}}>Trade Portfolio
               <img src="images/shell.jpg" style={{  width: '53px',height: '49px',float:'right',marginTop: '20px'}}  /></Subheader>
               <Subheader style={{fontWeight: 'bold',color: '#333333',fontSize: '24px'}}>Shell Trader </Subheader>
        </div>
        <div style={shellTrader.priceTrade}>
        <Subheader style={{fontWeight: 'bold',color: '#333333'}}>AVG. PRICE TRADED ON
        </Subheader>
        <span style={shellTrader.leftBrent}>BRENT &or;</span>
        <span style={shellTrader.rightBBL}>&and; $52/BBL</span>      
        </div>   
        <div style={shellTrader.priceTradeBBLS}>
        <p style={{paddingTop: '15px'}}>
          <span style={shellTrader.KBBLS}>500</span><span style={shellTrader.BBLSText}>K BBLS</span> 
          <span style={shellTrader.LONG}>LONG</span> 
        </p>        
        <p>
        <span style={shellTrader.KBBLS}>600</span><span style={shellTrader.BBLSText}>K BBLS</span> 
        <span style={shellTrader.LONGRED}>SHORT</span> 
        </p>
     
        </div>   
        <div>
        <Subheader style={{fontWeight: 'bold',color: '#333333'}}>Available Credit</Subheader>
        <span style={shellTrader.leftText}>For jan,2018</span>
        <span style={shellTrader.mn}>$7.2mm</span>
        </div>  
        <div>
        <h5 style={{fontWeight: 'bold',color: '#333333 ',marginLeft:'15px',paddingTop: '10px'}}>Credit Utilization (Last 3 Months)</h5>
        <p><span style={shellTrader.leftText}>Dec,2017</span><span style={shellTrader.textRight}>$7.2mm</span></p>
        <p><span style={shellTrader.leftText} >Dec,2017</span><span style={shellTrader.textRight} >$7.2mm</span></p>
        <p><span style={shellTrader.leftText}>Dec,2017</span><span style={shellTrader.textRight}>$7.2mm</span></p>
        </div>                     
       </div>
          </div>
          
      )
    }
  }



