import React from 'react';
import { Link } from 'react-router';
import Paper from 'material-ui/Paper';

const style={
  paperStyle:{
    width: "574px",
    height: "280px",
    borderRadius: "6px",
    marginTop:"30px",
    marginBottom:"20px",
    backgroundColor:" #ffffff",
    border: "solid 1px #d5d5d5"
  }
 
}


export default class ParcelStatusComponent extends React.Component {
  
  render() {
      return (
        <div>
         <div style={style.paperStyle} zDepth={2} >
        ParcelStatusComponent
           </div>
          </div>
      )
    }
  }



