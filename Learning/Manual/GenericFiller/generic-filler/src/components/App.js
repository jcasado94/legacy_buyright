import React, { Component } from 'react';
import '../styles/App.css';
import '../styles/reactTags.css';

import ProdTitle from './ProdTitle'
import ProdImage from './ProdImage'
import ProdTags from './ProdTags';
import ProdUsages from './ProdUsages';

class App extends Component {
	render() {
		return (
			<div className="App">
				<div className="App-main">
					<ProdTitle />
					<ProdImage />
					<div className="prod-properties">
						<ProdTags />
						<ProdUsages />
					</div>
				</div>
			</div>
		);
	}
}

export default App;
