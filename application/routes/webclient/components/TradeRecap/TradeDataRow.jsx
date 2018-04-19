import React from 'react';
import {
    Table,
    TableRow,
    TableBody,
    TableHeader,
    TableRowColumn,
  } from 'material-ui/Table';

  import {Collapse,Well,Row,Col} from 'react-bootstrap';

  const style={
    dateRangeStyle:{
        borderRight:"1px solid #ccc",
        marginTop:"5px"
    },
    dateRangeHeader:{
        marginTop: "0",
        color:"#666666"
    },
    paddingZero:{
        padding:"2px 0px"
    }
}

  class TradeDataRow extends React.Component{
    state={
        openRow:false,
        showCheckboxes: false
    }

    openCollapse=()=>{
        this.setState({openRow:!this.state.openRow});
        console.log('----------row clicked');
        
    }
render(){

    return (
       <div style={{width:"auto"}}>

            <TableRow onTouchTap={this.openCollapse} style={{width:"100%"}} >
                <TableRowColumn >{this.props.data.date}</TableRowColumn>
                <TableRowColumn>{this.props.data.tradeID} </TableRowColumn>
                <TableRowColumn> {this.props.data.type} </TableRowColumn>
                <TableRowColumn> {this.props.data.counterParty} </TableRowColumn>
                <TableRowColumn> {this.props.data.product} </TableRowColumn>
                <TableRowColumn> {this.props.data.quantity} </TableRowColumn>
                <TableRowColumn> {this.props.data.location} </TableRowColumn>
                <TableRowColumn> {this.props.data.incoTerm} </TableRowColumn>
            </TableRow>
           
           
            <Collapse in={this.state.openRow} style={{position:"relative", left:"0"}}>
            <Row>
        <Col md={3} style={style.dateRangeStyle}>
           <h6 style={style.dateRangeHeader}>Date Range</h6>
          
           <Col md={4} style={style.paddingZero}>
           <span>Delivery :</span>
           </Col>
           <Col md={8}  style={style.paddingZero}><span>{this.props.data.confirmedTrade.deliveryFromDate}</span><span>-</span>
           <span>{this.props.data.confirmedTrade.deliveryToDate}</span>
           </Col>
           
           <Col md={3} style={style.paddingZero}><span>Laycan :</span></Col>
           <Col md={9} style={style.paddingZero}><span>{this.props.data.confirmedTrade.laycanFromDate}</span><span>-</span>
           <span>{this.props.data.confirmedTrade.laycanToDate}</span>
           </Col>
        </Col>
        <Col md={3} style={style.dateRangeStyle}>
        <h6 style={style.dateRangeHeader}>Quality</h6>
      
        <Col md={7} style={style.paddingZero}><span>Quality (API) :</span></Col>
        <Col md={5} style={style.paddingZero}>
        <span>{this.props.data.confirmedTrade.qualityApi}</span></Col>
        <Col md={7} style={style.paddingZero}><span>Quality (SUL) :</span>
        </Col>
        <Col md={5} style={style.paddingZero}>
        <span>{this.props.data.confirmedTrade.qualitySul}</span>
        </Col>
       
        <Col md={7} style={style.paddingZero}>
        <span>Tolerence :</span>
        </Col>
        <Col md={5} style={style.paddingZero}><span>{this.props.data.confirmedTrade.tolerence}</span>
        </Col>
        </Col>
        <Col md={6} style={{marginTop:"5px"}}>
        <h6 style={style.dateRangeHeader}>Price</h6>
        <Col md={4}>
        
        <Col md={7} style={style.paddingZero}>
        <span>Price Type :</span>
        </Col>
        <Col md={5} style={style.paddingZero}><span>{this.props.data.confirmedTrade.priceType}</span>
        </Col>
        <Col md={7} style={style.paddingZero}><span>Index :</span>
        </Col>
        <Col md={5} style={style.paddingZero}><span>{this.props.data.confirmedTrade.index}</span>
        </Col>
        
        <Col md={7} style={style.paddingZero}><span>Price+UoM :</span>
        </Col>
        <Col md={5} style={style.paddingZero}><span>{this.props.data.confirmedTrade.priceUom}</span>
        </Col>
        </Col>
        <Col md={8}>
        <Col md={7} style={style.paddingZero}><span>Total Associate Fee :</span>
        </Col>
        <Col md={5} style={style.paddingZero}><span><b>{this.props.data.confirmedTrade.totalAssociate}</b></span>
        
        </Col>
        <Col md={7} style={style.paddingZero}><span>Total Fee :</span>
        </Col>
        <Col md={5} style={style.paddingZero}><span><b>{this.props.data.confirmedTrade.totalFees}</b></span>
        </Col>
        </Col>
        </Col>
        <Col md={12} >
        <div className="pull-left"><h2>Current Status: Vessel Identified</h2></div>
        <div className="pull-right"><h2>View Details</h2></div>
        </Col>
       </Row>
                </Collapse>
            
        </div>
    );
  }   
}

export default TradeDataRow;