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
    width: "574px",
    height: "280px",
    borderRadius: "6px",
    marginTop:"30px",
    marginBottom:"20px",
    padding:"0px 10px",
    backgroundColor:" #ffffff",
    border: "solid 1px #d5d5d5"
  }
 
}


export default class ParcelStatusComponent extends React.Component {
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
         <h4><b>Parcel Status</b></h4>
        <Table onRowSelection={this.handleRowSelection} className="tradeTable">
                <TableHeader  displaySelectAll={this.state.showCheckboxes}
                      adjustForCheckbox={this.state.showCheckboxes}>
                  <TableRow>
                    <TableHeaderColumn>Parcel</TableHeaderColumn>
                    <TableHeaderColumn style={{paddingLeft:"0px",paddingRight:"0px"}}>Awaiting Nomination</TableHeaderColumn>
                    <TableHeaderColumn>Scheduled</TableHeaderColumn>
                    <TableHeaderColumn>Shipped/Enroute</TableHeaderColumn>
                    <TableHeaderColumn>Completed</TableHeaderColumn>
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
            <TableRowColumn style={{width:"20%"}}>TR032182830</TableRowColumn>
            <TableRowColumn colspan="2"><div className="progress" style={{marginBottom:"0px"}}>
  <div className="progress-bar progress-bar-warning" role="progressbar" aria-valuenow="40" aria-valuemin="0" aria-valuemax="100" style={{width: "100%"}}>
   
  </div>
</div></TableRowColumn>
            
            </TableRow>
            <TableRow >
            <TableRowColumn style={{width:"20%"}}>TR032395050</TableRowColumn>
            <TableRowColumn colspan="2"><div className="progress" style={{marginBottom:"0px"}}>
  <div className="progress-bar progress-bar-info" role="progressbar" aria-valuenow="40" aria-valuemin="0" aria-valuemax="100" style={{width: "50%"}}>
   
  </div>
</div></TableRowColumn>
            
            </TableRow>
            <TableRow >
            <TableRowColumn style={{width:"20%"}}>TR03239459052</TableRowColumn>
            <TableRowColumn colspan="2"><div className="progress" style={{marginBottom:"0px"}}>
  <div className="progress-bar progress-bar-success" role="progressbar" aria-valuenow="40" aria-valuemin="0" aria-valuemax="100" style={{width: "90%"}}>
    
  </div>
</div></TableRowColumn>
            
            </TableRow>
            <TableRow >
            <TableRowColumn style={{width:"20%"}}>TR03453452</TableRowColumn>
            <TableRowColumn colspan="2"><div className="progress" style={{marginBottom:"0px"}}>
  <div className="progress-bar progress-bar-danger" role="progressbar" aria-valuenow="40" aria-valuemin="0" aria-valuemax="100" style={{width: "100%"}}>
    
  </div>
</div></TableRowColumn>
            
            </TableRow>
         
            </TableBody>
                </Table>
           </div>
          </div>
      )
    }
  }



