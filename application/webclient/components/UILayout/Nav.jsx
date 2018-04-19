import React from 'react';
import ReactDOM from 'react-dom';
import AppBar from 'material-ui/AppBar';
import Drawer from 'material-ui/Drawer';
import MenuItem from 'material-ui/MenuItem';
import FontIcon from 'material-ui/FontIcon';
import FlatButton from 'material-ui/FlatButton';
import {Link} from 'react-router';
import RaisedButton from 'material-ui/RaisedButton';
import Avatar from 'material-ui/Avatar';
import FileFolder from 'material-ui/svg-icons/file/folder';
import List from 'material-ui/List/List';
import ListItem from 'material-ui/List/ListItem';
const style = {
  labelStyle: {
      width: 'auto',
      height: '22px',
      family: 'Helvetica',
      size: '18px',
      weight: 'bold',
      style: 'normal',
      stretch: 'normal',
      height: 'normal',
      spacing: 'normal',
      align: 'left',
      color: '#ffffff',
      textTransform: 'lowercase'
     },
     labelStyle1: {
      width: 'auto',
      height: '22px',
      family: 'Helvetica',
      size: '18px',
      marginLeft:"-100px",
      weight: 'bold',
      style: 'normal',
      stretch: 'normal',
      height: 'normal',
      spacing: 'normal',
      align: 'left',
      color: '#ffffff',
      textTransform: 'lowercase'
     },
  buttonBorder:{
    width: '167px',
    height: '48px',
    radius: '6px',
    margin: '8px',
    border: 'solid 1px #ffffff',
    color:'#FFFFFF'
  }
} ;
import {
  blue300,
} from 'material-ui/styles/colors';



export default class Nav extends React.Component{
	state={
		openDrawer:false
  };
  static get contextTypes() {
    return {
      router: React.PropTypes.object.isRequired
    }
  }

	handleClose = () => this.setState({openDrawer: false});
  handleToggle = () => this.setState({openDrawer: !this.state.openDrawer});

  newTradeNavigation=()=>{
    this.context.router.push('/newTrade');
  }
  dashboardNavigation=()=>{
    this.context.router.push('/');
  }
  tradeRecapNavigation=()=>{
    this.context.router.push('/orgDetail');
  }
  confirmTradeNavigation=()=>{
    this.context.router.push('/transaction');
  }
  confirmParcelNavigation=()=>{
    this.context.router.push("/confirmParcel");
  }
  pendingParcelNavigation=()=>{
    this.context.router.push("/pendingParcel");
  }
  createParcelNavigation=()=>{
    this.context.router.push("/parcelHome");
  }
  handleLogout=()=>{
    sessionStorage.removeItem('userLoginDetails');
    this.context.router.push("/login");
  }
  inspectorNavigation=()=>{
    this.context.router.push("/inspector");
  }
  agentNavigation=()=>{
    this.context.router.push("/agent");
  }
  newrequestNavigation=()=>{
    this.context.router.push("/NewRequest");
  }
  ongoingNavigation=()=>{
    this.context.router.push("/Ongoing");
  }
	render(){
    
		return(
			<div>
			 <AppBar
             title="Blockchain Visual App"
             iconClassNameRight="muidocs-icon-navigation-expand-more"
             onLeftIconButtonTouchTap={this.handleToggle}
             style={{position: "fixed",top:'0',backgroundColor: '#1f497d'}}
            >
           <FlatButton style={style.buttonBorder} label="Dashboard" onTouchTap={this.dashboardNavigation} />
           <FlatButton style={style.buttonBorder} label="Organisation Detail" onTouchTap={this.tradeRecapNavigation} />
           <FlatButton style={style.buttonBorder} label="Transaction Flow" onTouchTap={this.confirmTradeNavigation} />
           <FlatButton style={style.buttonBorder} label="Health Checkup" onTouchTap={this.newTradeNavigation}>
            </FlatButton>
          
           </AppBar>

           <Drawer
          docked={false}
          width={200}
          open={this.state.openDrawer}
          onRequestChange={(openDrawer) => this.setState({openDrawer})}
          >

        
          <MenuItem onTouchTap={this.handleClose}>
           <Link to="/"> Home </Link>
          </MenuItem>
        	<MenuItem onTouchTap={this.handleClose}>
          <Link to ="/orgDetail">Organisation Detail</Link>
          </MenuItem>
          <MenuItem onTouchTap={this.handleClose}>
          <Link to ="/transaction">Transaction Flow</Link>
          </MenuItem>
          <MenuItem onTouchTap={this.handleClose}>
          <Link to ="/confirmTrade">Health Checkup </Link>
          </MenuItem>
          <MenuItem onTouchTap={this.handleClose}>
          <Link to ="/network">Network </Link>
          </MenuItem>
          <MenuItem onTouchTap={this.handleLogout}>
          Logout
          </MenuItem>
        </Drawer>
        </div>
      )
    }
}
