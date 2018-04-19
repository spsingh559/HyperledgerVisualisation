import React from 'react';
import { Link } from 'react-router';
import Axios from 'axios';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import CircularProgress from 'material-ui/CircularProgress';
// import {Grid} from 'react-bootstrap';
import {Grid,Row,Col} from 'react-bootstrap';
// import TradePortalComponent from './Dashboard/TradePortalComponent.jsx';
import TradeStatusAndRecap from './Dashboard/TradeStatusAndRecap/TradeStatusAndRecapComponent.jsx';
import Notification from './Dashboard/Notification.jsx';
// import ParcelStatusComponent from './Dashboard/ParcelStatusComponent.jsx';
// import TradeSummaryPublished from './Dashboard/TradeSummaryPublished.jsx';


export default class Home extends React.Component {

  state={
    buyCount:0,
    sellCount:0,
    animating:false,
    tradeData:[],
    pendingCount:0,
    correctionPendingCount:0,
    secondsElapsedBuy:4
  }

  static get contextTypes() {
    return {
      router: React.PropTypes.object.isRequired,
      socket: React.PropTypes.object.isRequired
    }
  }
 
  closeActivityIndicator = () =>setTimeout(() =>this.setState({
    animating: true }), 4000);

    componentDidMount=()=>{
      // this.intervalTrack = setInterval(() => this.tick(), 500);
      // this.intervalTrackSell = setInterval(() => this.tickSell(), 800);
      // this.intervalTrackBuy = setInterval(() => this.tickBuy(), 1000);
      this.closeActivityIndicator();

      this.context.socket.on('pranjul',(msg)=>{
        // setTimeout(this.myFunction(msg), 2000);
        // console.log('socket initiated here for JOb Inititation------------------');
        //   console.log(msg);
        alert(msg);
        console.log(msg);
      });

      // this.contextTypes.socket.on('pranjul',function(msg){
      //  alert(msg);
      // });
  
    }

    // tickBuy=()=> {
     
    //   this.setState((prevState) => ({
    //     secondsElapsedBuy: prevState.secondsElapsedBuy - 1,
  
    //   }));
  
    //   if(this.state.secondsElapsedBuy>=4){
    //     clearInterval(this.intervalTrackBuy);
    //   }
    
    // }
  
  render() {
    // let retrievedUserDetails= JSON.parse(sessionStorage.getItem('userLoginDetails'));
  
    if(this.state.animating==false){
      return(
      <div style={{marginTop:"90px",width:"700px",height:"500px"}}>
      <center>
     
       
        
        <h3> Blockchain Visualisation Dashboard is Getting Ready </h3>
        <br />
      
    <img src="../images/landingPage.gif" width="300px" height="200px" />  
      
      </center>
      
    </div>
      )
    }else{

      return (
       
        <Grid style={{marginTop:"90px"}}>  
        
            
         
			<Row >
			{/* <Col xs={2}><TradePortalComponent />
      </Col> */}
      <Col xs={7}><TradeStatusAndRecap buyCount={4} sellCount={5} totalTrade={9} 
      correctionPendingCount={2}
      pendingCount={6}/>
      </Col>
      <Col xs={3}><Notification />
      </Col>
			</Row>
           
          
         
          </Grid>
      )
    }
    }
  }



