import React, { Component } from 'react';
import Usage from './Usage';
import { USAGES_TITLE } from '../global/global';

class ProdUsages extends Component {

  render() {
    return(
      <div className="Prod-usages">
        <h3>{USAGES_TITLE}</h3>
        {this.props.usages.map(({usageId, name, pos, key}) => {
          return(
          <Usage
            usageId={usageId}
            name={name}
            pos={pos}
            key={key}
            onUsageClick={this.props.onUsageClick}
          />)
        }
          
        )}
      </div>
    );
  }

}

export default ProdUsages;