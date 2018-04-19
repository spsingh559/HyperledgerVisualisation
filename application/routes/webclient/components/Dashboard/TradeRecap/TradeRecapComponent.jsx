import React from 'react';
import Paper from 'material-ui/Paper';

import TradeRecapChartComponent from './TradeRecapChartComponent.jsx';


const style={
  paperStyle:{
    width: "574px",
    height: "290px",
    borderRadius: "6px",
    marginLeft:"100px",
    marginTop:"30px",
    backgroundColor:" #ffffff",
    border: "solid 1px #d5d5d5"
  }
}


export default class TradeStatusAndRecapComponent extends React.Component {
  
  state={
    tradeRecapData:[
      
      {
        text:"Mercuria",
        value:4,
        value:8
      },
      {
        text:"Statoil",
        value:4,
        value:6
      },
      {
        text:"BP",
        value:5,
        value:6
      }
    ]
  }

  render() {
      return (
        
          <div style={style.paperStyle} zDepth={2} >
         <TradeRecapChartComponent tradeRecapData={this.state.tradeRecapData} />
           </div>
          
      )
    }
  }



