import React from 'react';
import BarChart from 'react-bar-chart';

const margin = {top: 10, right: 20, bottom: 30, left: 50};
export default class TradeRecapChartComponent extends React.Component{

    render(){

        return(
            <div style={{width: '50%',marginTop:'20px'}} className="BarChartCSS"className="BarChartCSS">
               
          <BarChart ylabel='Trade Recap'
          
            width={200}
            height={300}
            margin={margin}
            data={this.props.tradeRecapData}
          />
                </div>
        )
    }
}