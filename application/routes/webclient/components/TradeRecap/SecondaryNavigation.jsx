import React, { Component } from 'react';
import { Link } from 'react-router';
import {Grid,Row,Col} from 'react-bootstrap';

class SecondaryNavigation extends Component{

    draftClickHandler = (e) => {     
        console.log(e)
    }

    amendClickHandler = (e) => {     
        console.log(e)
    }

    rejectedClickHandler = (e) => {     
       console.log(e)
    }

    render() {
        return(
            <div className="trSecNav">
                <a className="trLink selected" onTouchTap={() => this.draftClickHandler(event)}>Draft</a>
                <a className="trLink" onTouchTap={() => this.amendClickHandler(event)}>Amend</a>
                <a className="trLink" onTouchTap={() => this.rejectedClickHandler(event)}>Rejected</a>
            </div>
        );
    }
}

export default SecondaryNavigation;