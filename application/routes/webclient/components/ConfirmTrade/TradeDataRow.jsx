import React from 'react';
import {TableRow,TableRowColumn,} from 'material-ui/Table';


import {Collapse,Well, Grid, Col, Row} from 'react-bootstrap';
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
//   import TradeSubChildComponent from './TradeSubChildComponent.jsx';
  class TradeDataRow extends React.Component{

    state={
        openRow:false
    }

    openCollapse=()=>{
        this.setState({openRow:!this.state.openRow});
        console.log('----------row clicked');
        
    }
render(){

    return (
       <div>
        <TableRow onTouchTap={this.openCollapse}>
            <TableRowColumn >{this.props.data.date}
            
             </TableRowColumn>
            <TableRowColumn>{this.props.data.tradeID} </TableRowColumn>
            <TableRowColumn> {this.props.data.type} </TableRowColumn>
            <TableRowColumn> {this.props.data.counterParty} </TableRowColumn>
            <TableRowColumn> {this.props.data.product} </TableRowColumn>
            <TableRowColumn> {this.props.data.quantity} </TableRowColumn>
            <TableRowColumn> {this.props.data.location} </TableRowColumn>
            <TableRowColumn> {this.props.data.incoTerm} </TableRowColumn>
            
        </TableRow>
        {/* <TableRow >    
        <TableRowColumn colSpan="8"> */}
        
        {/* <Collapse in={this.state.openRow}>
        <Row> */}
        {/* <Col md={3} style={style.dateRangeStyle}>
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
       </Row> */}
       <Collapse in={this.state.openRow} style={{position:"relative", left:"0"}}>
                   <div>
                   <div className="statuscards">
                        {/* card 1 */}
                        <div className="statuscard completed">
                            <div className="thumb">
                               
                            </div>
                            <h3>The Princess
                            <span>XYZ Shipping Co.</span>
                            <span>1M BBLS</span>    
                            </h3>
                            <div className="icon">
                               <i className="fa fa-check-circle fa-3" aria-hidden="true"></i>
                            </div>
                            <div className="date">
                                28 Feb 2018 - 5th March 2018
                            </div>
                        </div>
                        {/* card 2 */}
                        <div className="statuscard pending">
                            <div className="thumb">

                            </div>
                            <h3>Inspector
                                <span>XYZ Shipping Co.</span>
                                <span>1M BBLS</span>    
                            </h3>
                            <div className="icon">
                                <i className="fa fa-check-circle fa-3" aria-hidden="true"></i>
                            </div>
                            <div className="date">
                                28 Feb 2018 - 5th March 2018
                            </div>
                        </div>
                        {/* card 3 */}
                        <div className="statuscard notstarted">
                            <div className="thumb">

                            </div>
                            <h3>Agent
                                <span>XYZ Shipping Co.</span>
                                <span>1M BBLS</span>    
                            </h3>
                            <div className="icon">
                                <i className="fa fa-check-circle fa-3" aria-hidden="true"></i>
                            </div>
                            <div className="date">
                                28 Feb 2018 - 5th March 2018
                            </div>
                        </div>
                    </div>
                   </div>
                </Collapse>
      
    {/* </Collapse> */}
    
    {/* </TableRowColumn>
    </TableRow> */}
    </div>
    );
}
    
}

export default TradeDataRow;