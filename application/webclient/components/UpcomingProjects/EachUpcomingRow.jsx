import React from 'react';
import { Row,Col } from 'react-bootstrap';
import Avatar from 'material-ui/Avatar';
import Chip from 'material-ui/Chip';
import {Card, CardActions, CardHeader, CardText} from 'material-ui/Card';
import FlatButton from 'material-ui/FlatButton';

const style={
    projectBrief:{
    borderRadius: "6px",
    border: "solid 1px #d5d5d5"
    }
}
export default class EachUpcomingRow extends React.Component{

   

    render(){
        let arr=[];
        this.props.data.peers.forEach((data)=>{
            arr.push(<Chip
               style={{marginTop:"10px"}}
              
              >
                <Avatar src="./../../images/nodeIcon.jpg" />
                {data}
              </Chip>)
        })

        let endorseArr=[];
        this.props.data.endorsingPeer.forEach((data)=>{
            endorseArr.push(<Chip
               style={{marginTop:"10px"}}
              
              >
                <Avatar src="./../../images/nodeIcon.jpg" />
                {data}
              </Chip>)
        })


        return(
           
            <Card style={{marginTop:"30px"}}>
    <CardHeader
      title={this.props.data.orgName}
      subtitle={this.props.data.orgDescription}
      actAsExpander={true}
      showExpandableButton={true}
    />
    
    <CardText expandable={true}>
      <Row>
          <Col xs={8}>
          <h3> Peers List</h3>
          {arr}
          </Col>
          <Col xs={4} style={style.projectBrief}>
          
         
          <br />
           <h4>Orderer : {this.props.data.orderer} </h4>
           <h4> Endorsing Peers List </h4>
           { endorseArr}
           
          </Col>
          </Row>
    </CardText>
  </Card>
        )
    }
}