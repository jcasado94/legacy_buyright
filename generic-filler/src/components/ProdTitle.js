import React, { Component } from 'react';

class ProdTitle extends Component {
	
	constructor() {
		super();
		this.state = {
			name: "Prod title"
		}
	}
	
  render() {
    return (
      <div className="Prod-title">
				<h1>{this.state.name}</h1>
			</div>
    );
  }
}

export default ProdTitle;
