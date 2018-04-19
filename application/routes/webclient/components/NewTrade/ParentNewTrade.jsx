import React from 'react';
import NewTradeComponent from './NewTradeComponent.jsx';

export default class ParentNewTrade extends React.Component{

    submitNewTrade=(obj)=>{
console.log('-----------data reach to parent');
console.log(obj);


    }
    render(){
        return(
            <div>
                <NewTradeComponent submitNewTrade={this.submitNewTrade}/>
                </div>
        )
    }
}