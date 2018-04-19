import React from 'react';
import { Link } from 'react-router';
import Axios from 'axios';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import Paper from 'material-ui/Paper';
import {Grid} from 'react-bootstrap';




export default class Login extends React.Component {
  constructor(props){
    super(props);
    this.state={
      username:'',
      password:'',
     
      
    }
    
  }
  handleUsername(e){
    this.setState({ username: e.target.value });
  }
  handlePassword(e){
    this.setState({ password: e.target.value });
  }

  
  static get contextTypes() {
    return {
      router: React.PropTypes.object.isRequired
    }
  }
 


  loginClick=()=>{
console.log('---------login clicked-----------');
this.context.router.push('/');

//     var data={
//       username:this.state.username,
//       password:this.state.password
//     }

//     Axios({
//       method:'post',
//       url:'/userLogin',
//       data:data
//     })
//     .then((data1) => {
//       console.log('Login details connected to server for post');
//       console.log(data1);


// if(data1.data.success==true){
//   console.log('------login success-------------------');
//   console.log(data1.data);
//   console.log('---------session variable started------------');
//   sessionStorage.setItem('userLoginDetails',JSON.stringify(data1.data));
//   if(data1.data.role=="Trader"){
//     console.log('---------Uniper--------');
  
//     this.context.router.push('/');
//   }else if(data1.data.role=="Operator"){
//     this.context.router.push('/parcelHome');
//   }
//   else if(data1.data.role=="Inspector"){
//     this.context.router.push('/inspector');
//   }else if(data1.data.role=="Shipping"){
//     this.context.router.push('/NewRequest');
//   }else if(data1.data.role=="Agent"){
//     this.context.router.push('/agent');
//   }


// }else{
//   alert('---------Enter Valid Credentials----------------');
// }

     

// })
//     .catch((error) => {

//       console.log(error);
//       console.log(error+"error in Login data for post");
//     });
  

  }
  render() {
  
  
      return (

        <div className="background">
        <center>
          <div style={{marginTop:"150px", color:"white"}} >
        {/* <Paper style={{height:'350px',width:'500px',marginTop:'150px', opacity: "0.5"}} zDepth={5} > */}
      


        <h2 style={{marginTop: '10px'}}>
       Blockchain Visual App Login
        </h2>


        <TextField
      hintText="User Name"
      floatingLabelText="Enter User Name"
      onChange = {(event,newValue) => this.setState({username:newValue})}
    /><br />
    <TextField
      hintText=" Password"
      type="password"
      onChange = {(event,newValue) => this.setState({password:newValue})}
      floatingLabelText="Enter Password"
    /><br />
        <div style={{marginTop:"50px"}}>
        <RaisedButton label="Login" primary={true}  onTouchTap={this.loginClick}/>
      
        </div>
        {/* </Paper> */}
        </div>
        </center>
    </div>

         
          )
      
    }
  }
