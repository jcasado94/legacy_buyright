import React, { Component } from 'react';
import classnames from 'classnames';

class Usage extends Component {

  constructor(props) {
    super(props);

    this.handleUsageClick = this.handleUsageClick.bind(this);
  }

  handleUsageClick() {    
    this.props.onUsageClick(this.props.usageId, this.props.pos);
  }

  render() {
    return(
      <button 
        className={classnames('usage-button', {'active': this.props.pos !== -1})}
        onClick={this.handleUsageClick}>
        <div className="usage-button-number">{this.props.pos !== -1 && this.props.pos + 1}</div>
  	    <div className="usage-button-text">{this.props.name}</div>
      </button>
    );
  }

}


export default Usage;