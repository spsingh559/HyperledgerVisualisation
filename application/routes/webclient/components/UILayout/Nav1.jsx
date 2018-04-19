import React from 'react';
import ReactDOM from 'react-dom';
import AppBar from 'material-ui/AppBar';
import Drawer from 'material-ui/Drawer';
import MenuItem from 'material-ui/MenuItem';
import FontIcon from 'material-ui/FontIcon';
import {Link} from 'react-router';



export default class Nav extends React.Component{

	state={
		openDrawer:false
	};

	handleClose = () => this.setState({openDrawer: false});

	handleToggle = () => this.setState({openDrawer: !this.state.openDrawer});


	render(){

		return(
			<div>
			 <AppBar
      
             title="Force Field App"
             iconClassNameRight="muidocs-icon-navigation-expand-more"
             onLeftIconButtonTouchTap={this.handleToggle}
             style={{position: "fixed",top:'0'}}
           />
           <Drawer
          docked={false}
          width={200}
          open={this.state.openDrawer}
          onRequestChange={(openDrawer) => this.setState({openDrawer})}
        >

        <img src='http://downloadicons.net/sites/default/files/man-logo-icon-61344.png' 
        style={{width:'180px',height:'200px'}}/>
        <center style={{fontWeight:'bold',fontSize:16}}>Hello Buddy</center>
          <MenuItem onTouchTap={this.handleClose}>
           <Link to="/"> Home </Link>
          </MenuItem>
        
					<MenuItem onTouchTap={this.handleClose}>
          <Link to ="/newTrade">New Trade</Link>
          </MenuItem>
          <MenuItem onTouchTap={this.handleClose}>
          <Link to ="/tradeRecap">Trade Recap</Link>
          </MenuItem>
          <MenuItem onTouchTap={this.handleClose}>
          <Link to ="/confirmTrade">Confirm Trade</Link>
          </MenuItem>
        </Drawer>
        </div>
			)
	}
}