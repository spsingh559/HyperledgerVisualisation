import React from 'react';
import CircularProgress from 'material-ui/CircularProgress';
import {Grid,Row,Col} from 'react-bootstrap';
import {
    Step,
    Stepper,
    StepLabel,
  } from 'material-ui/Stepper';

const style={
    styleDiv:{
        width: "250px",
        height: "70px",
        borderRadius: "6px",
        border: "solid 1px #d5d5d5"
    },
    styleDiv1:{
        width: "250px",
        height: "70px",
        borderRadius: "6px",
        marginLeft:"20px",
        border: "solid 1px #d5d5d5"
    }
}

export default class Transaction extends React.Component{
    state={
        txView:false,
        basicData:{
            peerNames:"peerNames",
            channelName:"channelName",
            chaincodeName:"chaincodeName",
            org:"org"
        },
        data:[{
            name:"Tx Initiated",
         time:"2:04"
        },{
            name:"Tx Initiated",
            time:"2:04"
        },{
        name:"Tx Initiated",
        time:"2:04"
        }
    ] ,
    dynamicValue:0       
    }

    static get contextTypes() {
        return {
          router: React.PropTypes.object.isRequired,
          socket: React.PropTypes.object.isRequired
        }
      }
 
      
    componentDidMount=()=>{
        console.log('component');
        this.context.socket.on('txInit',(msg)=>{
            // alert('tx init');

           this.setState({txView:true});
          });

          this.context.socket.on('basicInfo',(msg)=>{
              console.log(msg);
           this.setState({basicData:JSON.parse(msg)});
          });


        this.context.socket.on('sendTxProposal',(msg)=>{
            // alert('sendTxProposal');
            console.log('sendTxProposal');
            console.log(msg);
        //  this.setState({basicData:JSON.parse(msg)});
        });

        this.context.socket.on('sendTxProposalDecline',(msg)=>{
            console.log('sendTxProposalDecline');
            console.log(msg);
        //  this.setState({basicData:JSON.parse(msg)});
        });

        this.context.socket.on('proposalResponse',(msg)=>{
            console.log('proposalResponse');
            console.log(msg);
        //  this.setState({basicData:JSON.parse(msg)});
        });

        this.context.socket.on('proposalReceived',(msg)=>{
            console.log('proposalReceived');
            console.log(msg);
        //  this.setState({basicData:JSON.parse(msg)});
        });

        this.context.socket.on('CommitmentStatus',(msg)=>{
            console.log('CommitmentStatus');
            console.log(msg);
        //  this.setState({basicData:JSON.parse(msg)});
        });

        this.context.socket.on('finalStatus',(msg)=>{
            console.log('finalStatus');
            console.log(msg);
        //  this.setState({basicData:JSON.parse(msg)});
        });



    }

    
    render(){
if(this.state.txView==false){
    return(
        <div  style ={{marginTop:"90px"}}>
            <center>
                <h1> Wait ! </h1><br />
                <h2>
                    We are detecting any transaction initiated by Client 
                    </h2>
                    <br />
                    <CircularProgress />
                    </center>
            </div>
    )
}else{
    // console.log(this.state.basicData);
let arr=[];
   
    this.state.data.forEach((data,i)=>{
      // console.log(data);
      // console.log(i);

     
      arr.push(<Step key={i}><StepLabel>{data.name} <br /> {data.time}</StepLabel></Step>);
    
    })

    return(
        <div style ={{marginTop:"90px"}}>
            <center>
                <h2> Transaction flow of Org1 </h2>
                </center>
                <Grid>
        
                    <Row>
                        <Col xs={3} style={style.styleDiv}>
                        <b>
                        Channel Name <br /><br />
                        {this.state.basicData.channelName}
                        </b>
                        </Col>
                        <Col xs={3} style={style.styleDiv1}>
                        <b>
                            Peer Name <br /> <br />{this.state.basicData.peerNames}
                            </b>
                        </Col>
                        <Col xs={3} style={style.styleDiv1}>
                        <b>
                        Chain code Name <br /> <br />{this.state.basicData.chaincodeName}
                        </b>
                        </Col>
                        <Col xs={3} style={style.styleDiv1}>
                        <b>
                        Organisation Name <br /> <br />{this.state.basicData.org}
                        </b>
                        </Col>
                        </Row>
                        <Row>
                        <Stepper >

          {arr}

        </Stepper>
                            </Row>
                    </Grid>
            </div>
    )

}
        
    }
}