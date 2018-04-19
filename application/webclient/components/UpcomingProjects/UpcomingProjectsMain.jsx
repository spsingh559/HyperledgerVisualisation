import React from 'react';
import EachUpcomingRow from './EachUpcomingRow';

export default class UpcomingProjectsMain extends React.Component{

    render(){
      
        let newData= this.props.projectData.map((data,i)=>{
            return(
                <EachUpcomingRow 
                key={data._id}
                data={data}
                i={i}
                />
            )
        })
        return(
            <div>
                {newData}
                </div>
        )
    }
}