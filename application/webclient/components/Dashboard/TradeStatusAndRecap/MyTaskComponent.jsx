import React from 'react';
import { Link } from 'react-router';
import Paper from 'material-ui/Paper';
const style={
  paperStyle:{
    width: "272px",
    height: "153px",
    marginLeft:"150px",
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
  }
}

export default class MyTaskComponent extends React.Component {

  // state={
  //   secondsElapsed: 0,
  //   secondsElapsedBuy:0,
  //   secondsElapsedSell:0
  // }

  // componentDidMount=()=>{
  //   this.intervalTrack = setInterval(() => this.tick(), 1000);
  
  //   this.intervalTrackBuy = setInterval(() => this.tickBuy(), 800);

  // }
  // tick=()=> {
  //   this.setState((prevState) => ({
  //     secondsElapsed: prevState.secondsElapsed + 1,

  //   }));

  //   if(this.state.secondsElapsed>=6){
  //     clearInterval(this.intervalTrack);
  //   }
  // }

  // tickBuy=()=> {
  //   this.setState((prevState) => ({
  //     secondsElapsedBuy: prevState.secondsElapsedBuy + 1,

  //   }));

  //   if(this.state.secondsElapsedBuy>=4){
  //     clearInterval(this.intervalTrackBuy);
  //   }
  // }
  static get contextTypes() {
    return {
      router: React.PropTypes.object.isRequired
    }
  }

  navigateRecap=()=>{
    this.context.router.push('/tradeRecap');
  }

  
  render() {
      return (
       
         <div style={style.paperStyle} >
         <div style={{marginLeft:"20px"}}>
        <span style={style.labelStyle}>Node List</span>
       
        <br />
        <span style={style.labelStyleNumber} >{this.props.pendingCount}</span>
        <span style={style.labelStyleSubText}>Peers Count </span>
        <br/>
        <span style={style.labelStyleNumber} >{this.props.correctionPendingCount} </span>
        <span style={style.labelStyleSubText}>Orderes Count  </span>
        <br/>
        
            </div>
           </div>
         
      )
    }
  }



