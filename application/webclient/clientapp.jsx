import React from 'react';
import ReactDOM from 'react-dom';
// import {Router, Route, hashHistory} from 'react-router';

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
// <Route path="/dashboard" component={Dashboard} />
		// <Route path="/settings" component={Settings} />
		// <Route path="/analytics" component={Analytic} />
		// <IndexRoute component={MainApp} />

import injectTapEventPlugin from 'react-tap-event-plugin';
injectTapEventPlugin();

import Home from './components/Home.jsx';
import ParentComponent from './components/UILayout/Parent.jsx';
// import ParentNewTrade from './components/NewTrade/ParentNewTrade.jsx';
// import TradeRecapComponent from './components/TradeRecap/TradeRecapComponent.jsx';
// import ConfirmTradeComponent from './components/ConfirmTrade/ConfirmTradeComponent.jsx'
import Login from './components/Login.jsx';
// import parcelHome from './components/Parcel/parcelHome.jsx';
// import confirmParcel from './components/Parcel/ConfirmParcel/confirmParcel.jsx';
// import pendingParcel from './components/Parcel/PendingParcel/pendingParcel.jsx';
// // import createParcel from './components/Parcel/createParcel.jsx';
// // import AmendParentTrade from './components/NewTrade/AmendParentTrade.jsx';
// import InspectorComponent from './components/Inspector/InspectorComponent.jsx';
// import AgentComponent from './components/Agent/AgentComponent.jsx';
// import ParentOngoing from './components/Shipping/Ongoing/ParentOngoing.jsx';
// import ParentNewRequest from './components/Shipping/NewRequest/ParentNewRequest.jsx';
import  UpcomingProjects from './components/UpcomingProjects/UpcomingProjects';
import ContextComponent from './context.jsx';
import {Route, Router, IndexRoute, hashHistory} from 'react-router';
import Transaction from './components/Transaction/Transaction';
import Network from './components/Network/Network';

ReactDOM.render(
	<ContextComponent>
	<MuiThemeProvider>
	<Router history ={hashHistory} >
	<Route path="/login" component={Login} />
	<Route path="/" component={ParentComponent}>
	{/* <Route path="/newTrade" component={ParentNewTrade} />
	<Route path="/tradeRecap" component={TradeRecapComponent} />
	<Route path="/confirmTrade" component={ConfirmTradeComponent} />
	<Route path="/parcelHome" component={parcelHome} />
	<Route path="/confirmParcel" component={confirmParcel} />
	<Route path="/pendingParcel" component={pendingParcel} /> */}
	<Route path="/transaction" component={Transaction} />
	<Route path="/network" component={Network} />
	{/* <Route path="/editTradePage/:trid" component={AmendParentTrade}  /> */}
	{/* <Route path="/inspector" component={InspectorComponent} />
	<Route path="/Ongoing" component={ParentOngoing} />
	<Route path="/NewRequest" component={ParentNewRequest}  />	
	<Route path="/agent" component={AgentComponent}  /> */}
	<Route path="/orgDetail" component={UpcomingProjects} />
		<IndexRoute component={Home} />
		</Route>
	</Router>
	</MuiThemeProvider>
	</ContextComponent>,
 document.getElementById('mountapp'));
