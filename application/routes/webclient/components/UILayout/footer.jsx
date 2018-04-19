import React from 'react';

const style = {
    backgroundColor: "rgb(0, 188, 212)",
    borderTop: "1px solid #E7E7E7",
    textAlign: "center",
    padding: "20px",
    position: "fixed",
    left: "0",
    bottom: "0",
    height: "30px",
    width: "100%",
};

const phantom = {
  display: 'block',
  padding: '20px',
  height: '100%',
  width: '100%',
}
export default class Footer extends React.Component{

	render(){

		return(

			<div>
                <div style={phantom} />
                <div style={style}>
                </div>
            </div>
			);
	}
}