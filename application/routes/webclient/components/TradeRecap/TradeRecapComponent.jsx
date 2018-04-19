import React from 'react';
import { Link } from 'react-router';
import {Grid,Row,Col} from 'react-bootstrap';

import TradeRecapFilterComponent from './TradeRecapFilterComponent';
import SecondaryNavigation from "./SecondaryNavigation";
import TradeTableComponent from "./TradeTableComponent"

export default class TradeRecapComponent extends React.Component {

  render() {
    return (
      <Row className="trContainer show-grid" style={{marginLeft:"0", marginRight:"0"}}>
        <Col md={3} className="trLeftCol">
          <TradeRecapFilterComponent />
        </Col>
        <Col md={9} style={{paddingLeft:"0", paddingRight:"0"}}>
          <SecondaryNavigation />   
          <TradeTableComponent headingText="Pending for approval" number="03" />
          <TradeTableComponent headingText="Trade request sent" number="04" />
        </Col>
      </Row>
    )
  }
}
