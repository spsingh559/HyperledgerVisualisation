import React from 'react';
import { Link } from 'react-router';
import Paper from 'material-ui/Paper';
import {Row,Col} from 'react-bootstrap';

const style={
  paperStyle:{
    width: "272px",
    height: "153px",
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
  
  state={
    secondsElapsed: 0,
    secondsElapsedBuy:0,
    secondsElapsedSell:0
  }

  componentDidMount=()=>{
    this.intervalTrack = setInterval(() => this.tick(), 500);
    this.intervalTrackSell = setInterval(() => this.tickSell(), 800);
    this.intervalTrackBuy = setInterval(() => this.tickBuy(), 1000);

  }
  static get contextTypes() {
    return {
      router: React.PropTypes.object.isRequired
    }
  }

  // tick=()=> {
  //   if(this.props.totalTrade==0){
  //  this.setState({secondsElapsed:0});
  //  clearInterval(this.intervalTrack);
  //   }else{
  //   this.setState((prevState) => ({
  //     secondsElapsed: prevState.secondsElapsed + 1,

  //   }));

  //   if(this.state.secondsElapsed>=this.props.totalTrade){
  //     clearInterval(this.intervalTrack);
  //   }
  // }
  // }

  // tickBuy=()=> {
  //   if(this.props.buyCount==0){
  //     this.setState({secondsElapsedBuy:0});
  //     clearInterval(this.intervalTrackBuy);
  //      }else{
  //   this.setState((prevState) => ({
  //     secondsElapsedBuy: prevState.secondsElapsedBuy + 1,

  //   }));

  //   if(this.state.secondsElapsedBuy>=this.props.buyCount){
  //     clearInterval(this.intervalTrackBuy);
  //   }
  // }
  // }

  // tickSell=()=> {
  //   if(this.props.sellCount==0){
  //     this.setState({secondsElapsedSell:0});
     
  //       clearInterval(this.intervalTrackSell);
      
  //      }else{
  //   this.setState((prevState) => ({
  //     secondsElapsedSell: prevState.secondsElapsedSell + 1,

  //   }));

  //   if(this.state.secondsElapsedSell>=this.props.sellCount){
  //     clearInterval(this.intervalTrackSell);
  //   }
  // }
  // }

  // navigateCT=()=>{
  //   this.context.router.push('/confirmTrade');
  // }


  render() {
      return (
        
           <div style={style.paperStyle} zDepth={2} >
           <div style={{marginLeft:"20px"}}>
           <center>
<h3>Participating Organisation</h3>

        <br />
        
       <h3>{this.props.totalTrade} </h3>
       </center>
       
       
            </div>
           </div>
         
      )
    }
  }



