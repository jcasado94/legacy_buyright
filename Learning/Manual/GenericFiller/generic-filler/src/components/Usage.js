import React, { Component } from 'react';

class Usage extends Component {

  constructor() {
    super();
  }

  render() {
    return(
      <li className="Usage">
        <select className="Usage-select">
          {Object.keys(this.props.allUsages).map(
            usageId => <option key={usageId} value={usageId}>{this.props.allUsages[usageId]}</option>
          )}
        </select>
      </li>
    );
  }

}


export default Usage;