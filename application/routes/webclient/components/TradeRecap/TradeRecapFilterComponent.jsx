import React, { Component } from 'react';
import { Link } from 'react-router';

import {Grid,Row,Col} from 'react-bootstrap';

import TextField from 'material-ui/TextField';
import Checkbox from 'material-ui/Checkbox';
import ActionFavorite from 'material-ui/svg-icons/action/favorite';
import ActionFavoriteBorder from 'material-ui/svg-icons/action/favorite-border';
import Visibility from 'material-ui/svg-icons/action/visibility';
import VisibilityOff from 'material-ui/svg-icons/action/visibility-off';
import Slider from 'material-ui/Slider';

const styles = {
    block: {
      maxWidth: 250,
    },
    checkbox: {
      marginBottom: 16,
    },
};

class TradeRecapFilterComponent extends Component{
    state = {
        checked: false,
      }
    
      updateCheck() {
        this.setState((oldState) => {
          return {
            checked: !oldState.checked,
          };
        });
      }
    
    render() {
        return (
            <div>
                <h1 style={{marginBottom:'40px'}}>Filter By</h1>
                <ul className="trSwitch">
                <li><button className="selected"> ALL</button> </li>
                <li><button> BUY</button></li>
                <li><button> SELL</button></li>
                </ul>
                <form>
                    <div className="trFilterType">
                        <span className="trFilterLabel">commodity</span>
                        <TextField id="commodity" />
                        <div>                        
                            <Checkbox label="Brent" 
                                style={styles.checkbox} 
                                style={{marginBottom:'0px', fontWeight:'100'}}
                                inputStyle={{border:"1px solid #d5d5d5"}}/>
                            <Checkbox label="Tapis" style={styles.checkbox} style={{marginBottom:'0px'}} />
                            <Checkbox label="Natgas" style={styles.checkbox} style={{marginBottom:'0px'}} />
                            <Checkbox label="Wti" style={styles.checkbox} style={{marginBottom:'0px'}} />
                            <Checkbox label="Jcc" style={styles.checkbox} style={{marginBottom:'0px'}} />
                        </div>
                    </div>
                    <div className="trFilterType">
                        <span className="trFilterLabel">partners</span>
                        <div>                        
                            <Checkbox label="ABN-AMRO" style={styles.checkbox} style={{marginBottom:'0px', fontWeight:'100'}} />
                            <Checkbox label="BP" style={styles.checkbox} style={{marginBottom:'0px'}} />
                            <Checkbox label="ING" style={styles.checkbox} style={{marginBottom:'0px'}} />
                            <Checkbox label="SHELL" style={styles.checkbox} style={{marginBottom:'0px'}} />
                            <Checkbox label="STATE-OIL" style={styles.checkbox} style={{marginBottom:'0px'}} />
                            <Checkbox label="SOCIETE-GENERAL" style={styles.checkbox} style={{marginBottom:'0px'}} />
                        </div>
                    </div>
                    <div className="trFilterType">
                        <span className="trFilterLabel">Volume (in BBL)</span>
                        <div>                        
                            <Slider/>
                        </div>
                    </div>
                    <div className="trFilterType nobb">
                        <span className="trFilterLabel">Price (in m)</span>
                        <div>                        
                            <Slider />
                        </div>
                    </div>
                </form>
            </div>
        );
    }
}

export default TradeRecapFilterComponent;