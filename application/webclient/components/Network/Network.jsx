import React from 'react';

import {Map, InfoWindow, Marker, GoogleApiWrapper,Polyline} from 'google-maps-react';


    const style = {
        width: '100%',
        height: '100%'
      }

      const triangleCoords = [
        {lat: 25.774, lng: -80.190},
        {lat: 18.466, lng: -66.118},
        {lat: 32.321, lng: -64.757},
        {lat: 25.774, lng: -80.190}
      ];
    


export class Network extends React.Component {

    // onMouseoverMarker =(props, marker, e) =>{
    //     console.log('mouse over');
    //     // ..
    //     alert('Org1');
    //   }
    state = {
        showingInfoWindow: false,
        activeMarker: {},
        selectedPlace: {},
      };

      onMarkerClick = (props, marker, e) =>
    this.setState({
      selectedPlace: props,
      activeMarker: marker,
      showingInfoWindow: true
    });

  onMapClicked = (props) => {
    if (this.state.showingInfoWindow) {
      this.setState({
        showingInfoWindow: false,
        activeMarker: null
      })
    }
  };

  render() {
    return (
        

    
 <Map
google={this.props.google}
style={style}
className={'map'}
onClick={this.onMapClicked}
initialCenter={{
  lat: 13.633559,
  lng: 15.368832	
}}
zoom={3}
// onClick={this.onMapClicked}
>


          {/* <Polyline
          paths={triangleCoords}
          strokeColor="#0000FF"
          strokeOpacity={0.8}
          strokeWeight={2} /> */}

      <Marker
      onMouseover={this.onMarkerClick}
        name={'Organisation 1'}
        position={{lat: 38.889931, lng:  -77.009003}} />
        <Marker
        onMouseover={this.onMarkerClick}
        name={'Organisation 2'}
        position={{lat:51.508530, lng: -0.076132	}} />

         <Marker
         onMouseover={this.onMarkerClick}
        name={'Organisation 3'}
        position={{lat:35.685360, lng: 139.753372		}} />
       
      <Marker
      onMouseover={this.onMarkerClick}
        name={'Organisation 4'}
        position={{lat:-26.195246	, lng:28.034088		}}
        />

        <InfoWindow
          marker={this.state.activeMarker}
          visible={this.state.showingInfoWindow}>
            <div>
              <h1>{this.state.selectedPlace.name}</h1>
            </div>
        </InfoWindow>

    </Map> 

    
    );
  }
}

export default GoogleApiWrapper({
  apiKey: ('AIzaSyBx4_JsgDtIkHh8lMrjIOP0TTIYyJsmPlU')
})(Network)