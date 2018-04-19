import React from 'react';
import { Link } from 'react-router';
import {Grid,Row,Col} from 'react-bootstrap';

import TradeRecapFilterComponent from './TradeRecapFilterComponent';

import TradeTableComponent from "./TradeTableComponent"

export default class ConfirmTradeComponent extends React.Component {
  
  render() {
      return (
        <div style={{marginTop:"65px"}}>
        <Row className="trContainer show-grid" style={{marginLeft:"0", marginRight:"0"}}>
        <Col md={3} className="trLeftCol">
          <TradeRecapFilterComponent />
        </Col>
        <Col md={9} style={{paddingLeft:"0", paddingRight:"0"}}>
            
          <TradeTableComponent headingText="Confirmed Trades" number="04" />
          
        </Col>
      </Row>
          </div>
      )
    }
  }



