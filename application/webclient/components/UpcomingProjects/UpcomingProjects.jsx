import React from 'react';

import {Grid, Row, Col} from 'react-bootstrap';

import {Card, CardActions, CardHeader, CardText} from 'material-ui/Card';
import FlatButton from 'material-ui/FlatButton';

import UpcomingProjectsMain from './UpcomingProjectsMain';

const projectData=[
    {
        _id:"23423",
        orgName:"Org1",
        orgDescription:"Org1 belongs to Manufacturing Sector",
        peers:["peer0.org1","peer1.org1","peer2.org1"],
        orderer:"orderer0.org1",
        endorsingPeer:["peer0.org1","peer0.org2"],
    },
    {
        _id:"23424",
        orgName:"Org2",
        orgDescription:"Org2 belongs to Transport Sector",
        peers:["peer0.org2","peer1.org2","peer2.org2"],
        orderer:"orderer0.org2",
        endorsingPeer:["peer0.org1","peer0.org2"],
    },
    {
        _id:"23425",
        orgName:"Org3",
        orgDescription:"Org1 belongs to Retail Sector",
        peers:["peer0.org3","peer1.org3","peer2.org3"],
        orderer:"orderer0.org3",
        endorsingPeer:["peer0.org1","peer0.org3"],
    },

]
export default class UpcomingProjects extends React.Component{

    render(){
        return(
        <div style={{marginTop:"65px"}} className="upcomingBackground">
        <Grid>
            <center>
            <h2> Organisation Details </h2>
            </center>
            <UpcomingProjectsMain projectData={projectData}/>
            </Grid>
            </div>
        )
            
    }
}