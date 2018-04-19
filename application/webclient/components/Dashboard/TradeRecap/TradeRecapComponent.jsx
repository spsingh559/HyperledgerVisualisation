import React from 'react';
import Paper from 'material-ui/Paper';
import {Row,Col} from 'react-bootstrap';

// import TradeRecapChartComponent from './TradeRecapChartComponent.jsx';


const style={
  paperStyle:{
    width: "272px",
    height: "153px",
    borderRadius: "6px",
    backgroundColor:" #ffffff",
    border: "solid 1px #d5d5d5"
  },
  paperStyleChainCode:{
    width: "272px",
    height: "153px",
    marginLeft:"150px",
    borderRadius: "6px",
    backgroundColor:" #ffffff",
    border: "solid 1px #d5d5d5"
  }
}

// import ChannelCount from './'

export default class TradeStatusAndRecapComponent extends React.Component {
  
  

  render() {
      return (
        
          <div style={{width:"650px", height:"153px",marginTop:"50px"}} >
         <Row>
             <Col xs={4}> 
             <div style={style.paperStyle} zDepth={2} >
           <div style={{marginLeft:"20px"}}>
           <center>
<h3>Number of Channel</h3>

        <br />
        
       <h3>1 </h3>
       </center>
       
       
            </div>
           </div>
             </Col>
             <Col xs={4}>
             <div style={style.paperStyleChainCode}  >
           <div style={{marginLeft:"20px"}}>
           <center>
<h3>No. of Chaincode</h3>

        <br />
        
       <h3>4 </h3>
       </center>
       
       
            </div>
           </div>
             </Col>
             </Row>
           </div>
          
      )
    }
  }



