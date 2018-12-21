import React, { Component } from 'react';
import '../styles/App.css';
import '../styles/reactTags.css';

import Prod from './Prod'

class App extends Component {
	render() {
		return (
			<div className="App">
				<div className="App-main">
					<Prod/>
				</div>
			</div>
		);
	}
}

export default App;
