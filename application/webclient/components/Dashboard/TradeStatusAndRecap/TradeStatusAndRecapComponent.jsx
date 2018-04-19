import React from 'react';
// import { Link } from 'react-router';
import {Grid,Row,Col} from 'react-bootstrap';
import ConfirmTradeComponent from './ConfirmTradeComponent.jsx';
import MyTaskComponent from './MyTaskComponent.jsx';
import TradeRecapComponent from '../TradeRecap/TradeRecapComponent.jsx';
export default class TradeStatusAndRecapComponent extends React.Component {
  
 

  render() {
      return (
        <div>
         <Row>
             <Col xs={4}> <ConfirmTradeComponent  buyCount={this.props.buyCount} sellCount={this.props.sellCount} totalTrade={this.props.totalTrade}/>
             </Col>
             <Col xs={4}><MyTaskComponent correctionPendingCount={this.props.correctionPendingCount}  pendingCount={this.props.pendingCount}/>
             </Col>
             </Row>
             <Row>
                 <Col xs={8}><TradeRecapComponent />
                 </Col>
                 </Row>
          </div>
      )
    }
  }



