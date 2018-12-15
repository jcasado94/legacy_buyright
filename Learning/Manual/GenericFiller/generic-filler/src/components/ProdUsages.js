import React, { Component } from 'react';
import Usage from './Usage';

class ProdUsages extends Component {

  constructor() {
    super();

    this.handleAddition = this.handleAddition.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
  }

  handleAddition(usage) {
    this.props.onUsageAddition(usage);
  }

  handleDelete(i) {
    this.props.onUsageDelete(i);
  }

  render() {
    return(
      <div className="Prod-usages">
        <h3>Usages</h3>
        <ol>
          {this.props.prodUsages.map(
            usageId => <Usage key={usageId} allUsages={this.props.allUsages}/>
          )}
        </ol>
      </div>
    );
  }

}

export default ProdUsages;