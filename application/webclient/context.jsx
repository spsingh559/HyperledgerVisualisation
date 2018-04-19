import React from 'react';
// import restUrl from './restUrl';


export default class ContextComponent extends React.Component {
  getChildContext() {
    return {
      socket:io('http://13.127.234.53:8080')

    }
  }

  render(){
    return (this.props.children);
  }
}

ContextComponent.childContextTypes = {
  socket: React.PropTypes.object.isRequired
};
