import React from 'react';
import { Link } from 'react-router';
import Axios from 'axios';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
// import {Grid} from 'react-bootstrap';
import {Grid,Row,Col} from 'react-bootstrap';
import TradePortalComponent from './Dashboard/TradePortalComponent.jsx';
import TradeStatusAndRecap from './Dashboard/TradeStatusAndRecap/TradeStatusAndRecapComponent.jsx';
import Notification from './Dashboard/Notification.jsx';
import ParcelStatusComponent from './Dashboard/ParcelStatusComponent.jsx';
import TradeSummaryPublished from './Dashboard/TradeSummaryPublished.jsx';
export default class Home extends React.Component {
  
  render() {
  
      return (
        <Grid style={{marginTop:"90px"}}>  
        
            
         
			<Row >
			<Col xs={2}><TradePortalComponent />
      </Col>
      <Col xs={7}><TradeStatusAndRecap />
      </Col>
      <Col xs={3}><Notification />
      </Col>
			</Row>
      <Row>
        <Col xs={6}> <ParcelStatusComponent />
        </Col>
        <Col xs={6}> <TradeSummaryPublished />
        </Col>
        </Row>        
          
         
          </Grid>
      )
    }
  }



