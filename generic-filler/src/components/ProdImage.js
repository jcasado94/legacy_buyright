import React, { Component } from 'react';

class ProdImage extends Component {
	
	constructor() {
		super();
		this.state = {
			src: "https://www.compraonline.bonpreuesclat.cat/images/28431/ad712a26-9f70-4ff7-98c8-726af03ca400/500x500.jpg"
		}
	}
	
  render() {
    return (
			<img className="Prod-image" alt="" src={this.state.src}></img>
    );
  }
}

export default ProdImage;
