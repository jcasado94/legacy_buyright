import React, { Component } from 'react';

import ProdTitle from './ProdTitle'
import ProdImage from './ProdImage'

import ProdUsagesContainer from '../containers/ProdUsagesContainer';

class Prod extends Component {

	render() {
		return (
      <div className="Prod">
        <ProdTitle />
        <ProdImage />
        <div className="prod-properties">
          <ProdUsagesContainer />
        </div>
      </div>
		);
  }

}


export default Prod;
