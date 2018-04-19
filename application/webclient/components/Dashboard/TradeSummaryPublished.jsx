import React from 'react';
import { Link } from 'react-router';
import Paper from 'material-ui/Paper';
import {
  Table,
  TableBody,
  TableFooter,
  TableHeader,
  TableHeaderColumn,
  TableRow,
  TableRowColumn,
} from 'material-ui/Table';



const style={
  paperStyle:{
    width: "560px",
    height: "280px",
    borderRadius: "6px",
    marginTop:"30px",
    marginBottom:"20px",
    backgroundColor:" #ffffff",
    padding:"0px 10px",
    border: "solid 1px #d5d5d5"
  }
 
}

export default class TradeSummaryPublished extends React.Component {
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
      return (
        <div>
         <div style={style.paperStyle} zDepth={2} >
        <h4><b>Trade Summary Published</b></h4>
        <Table onRowSelection={this.handleRowSelection} className="tradeTable">
                <TableHeader  displaySelectAll={this.state.showCheckboxes}
                      adjustForCheckbox={this.state.showCheckboxes}>
                  <TableRow>
                    <TableHeaderColumn>Commodity</TableHeaderColumn>
                    <TableHeaderColumn>Date/Time</TableHeaderColumn>
                    <TableHeaderColumn>Trade ID</TableHeaderColumn>
                    <TableHeaderColumn>Volume</TableHeaderColumn>
                    
                  </TableRow>
                </TableHeader>
                <TableBody displayRowCheckbox={this.state.showCheckboxes}>
              {/* <TableBody>
                        {
                            this.props.pendingParcelData.map((item, i) => {
                                return (                                    
                                    <ParcelPendingDataRow data={item} key={i} />
                                );
                            })
                        }           
                      </TableBody>*/}
                      <TableRow >
            <TableRowColumn style={{width:"10%"}}>Crude Oil </TableRowColumn>
            <TableRowColumn style={{width:"10%"}}>20-01-20118 :10:30:00 PM IST</TableRowColumn>
            <TableRowColumn style={{width:"10%"}}>TR1232134234</TableRowColumn>
            <TableRowColumn style={{width:"10%"}}>10000 BL</TableRowColumn>
            </TableRow>
            <TableRow >
            <TableRowColumn style={{width:"10%"}}>Crude Oil </TableRowColumn>
            <TableRowColumn style={{width:"10%"}}>22-01-2018 :18:35:00 PM IST</TableRowColumn>
            <TableRowColumn style={{width:"10%"}}>TR12321342356</TableRowColumn>
            <TableRowColumn style={{width:"10%"}}>10000 BL</TableRowColumn>
            </TableRow>
            </TableBody>
                </Table>
           </div>
          </div>
      )
    }
  }



