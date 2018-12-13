import React, { Component } from 'react';
import { WithContext as ReactTags } from 'react-tag-input';

const placeholder = 'Select or create a usage'

class ProdUsages extends Component {
	
	constructor() {
    super();
    
		this.state = {
      usages: [],
      suggestions: [
        { id: 'two', text: 'Two' },
        { id: 'twoo', text: 'Twoo' },
        { id: 'twooo', text: 'Twooo' },
        { id: 'twooooo', text: 'Twoooo' }
      ]
    }

    this.handleAddition = this.handleAddition.bind(this);
    this.handleDelete = this.handleDelete.bind(this);

  }
  
  handleDelete(i) {
    const { usages } = this.state;
    this.setState({
      usages: usages.filter((tag, index) => index !== i)
    })
  }

  handleAddition(tag) {
    this.setState(state => ({ usages: [...state.usages, tag] }));
  }

  render() {
    const { usages, suggestions } = this.state;
    return (
      <div className="react-tags">
			  <h3>Usages</h3>
        <ReactTags 
          tags={usages}
          suggestions={suggestions}
          handleAddition={this.handleAddition}
          handleDelete={this.handleDelete}
          allowDragDrop={false}
          placeholder={placeholder}
          inline={false}
        />
      </div>
    );
  }
}

export default ProdUsages;
