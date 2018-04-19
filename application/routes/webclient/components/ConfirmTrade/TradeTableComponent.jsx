import React, { Component } from 'react';
import TradeDataRow from './TradeDataRow';

import {
    Table,
    TableBody,
    TableFooter,
    TableHeader,
    TableHeaderColumn,
    TableRow,
    TableRowColumn,
  } from 'material-ui/Table';


const data = [
    {
        date: 'date',
        tradeID: '1',
        type: 'buy',
        counterParty: 'STAT-OIL',
        product:'brent',
        quantity: '550k bbl',
        location: 'ROTTERDAM PORT',
        incoTerm: 'fob',
        confirmedTrade:{
            deliveryFromDate:"1 FEB 2018",
            deliveryToDate:"2 MAR 2018",
            laycanFromDate:"2 MAR 2018",
            laycanToDate:"2 MAR 2018",
            qualityApi:"30.8",
            qualitySul:"0.50",
            tolerence:"5%",
            priceType:"INDEX",
            index:"0.52",
            priceUom:"5%",
            totalAssociate:"$57,000",
            totalFees:"$7.2 MIN"
        }
    },
    {
        date:'date',
        tradeID: '2',
        type: 'buy',
        counterParty: 'STAT-OIL',
        product:'brent',
        quantity: '550k bbl',
        location: 'ROTTERDAM PORT',
        incoTerm: 'fob',
        confirmedTrade:{
            deliveryFromDate:"1 FEB 2018",
            deliveryToDate:"2 MAR 2018",
            laycanFromDate:"2 MAR 2018",
            laycanToDate:"2 MAR 2018",
            qualityApi:"30.8",
            qualitySul:"0.50",
            tolerence:"5%",
            priceType:"INDEX",
            index:"0.52",
            priceUom:"5%",
            totalAssociate:"$57,000",
            totalFees:"$7.2 MIN"
        }
    },
    {
        date: 'date',
        tradeID: '3',
        type: 'buy',
        counterParty: 'STAT-OIL',
        product:'brent',
        quantity: '550k bbl',
        location: 'ROTTERDAM PORT',
        incoTerm: 'fob',
        confirmedTrade:{
            deliveryFromDate:"1 FEB 2018",
            deliveryToDate:"2 MAR 2018",
            laycanFromDate:"2 MAR 2018",
            laycanToDate:"2 MAR 2018",
            qualityApi:"30.8",
            qualitySul:"0.50",
            tolerence:"5%",
            priceType:"INDEX",
            index:"0.52",
            priceUom:"5%",
            totalAssociate:"$57,000",
            totalFees:"$7.200 MIN"
        }
    },
];

class DraftTradeTableComponent extends Component {

    state = {
        fixedHeader: false,
        fixedFooter: false,
        stripedRows: true,
        showRowHover: true,
        selectable: true,
        multiSelectable: false,
        enableSelectAll: false,
        deselectOnClickaway: true,
        showCheckboxes: false,
        height: '300px',
       
      };
    
      isSelected = (index) => {
        return this.state.selected.indexOf(index) !== -1;
      };
    
      handleRowSelection = (selectedRows) => {
        this.setState({
          selected: selectedRows,
        });
    };
    
    render() {
        return(
            <div className="trTableContainer">
                <h2>
                    {this.props.headingText}
                    <span>{this.props.number}</span>
                </h2>                
                <Table onRowSelection={this.handleRowSelection}>
                <TableHeader  displaySelectAll={this.state.showCheckboxes}
            adjustForCheckbox={this.state.showCheckboxes}
>
                  <TableRow>
                    <TableHeaderColumn>Date</TableHeaderColumn>
                    <TableHeaderColumn>Trade ID</TableHeaderColumn>
                    <TableHeaderColumn>Type</TableHeaderColumn>
                    <TableHeaderColumn>Counter Party</TableHeaderColumn>
                    <TableHeaderColumn>Product</TableHeaderColumn>
                    <TableHeaderColumn>Quantity</TableHeaderColumn>
                    <TableHeaderColumn>Location</TableHeaderColumn>
                    <TableHeaderColumn>Inco Term</TableHeaderColumn>
                  </TableRow>
                </TableHeader>
                <TableBody>
                        {
                            data.map((item, i) => {
                                return (                                    
                                    <TradeDataRow data={item} key={i} />
                                );
                            })
                        }           
                   </TableBody>
                </Table>
            </div>
        );
    }
}

export default DraftTradeTableComponent;