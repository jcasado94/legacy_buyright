import React, { Component } from 'react';
import { WithContext as ReactTags } from 'react-tag-input';

const placeholder = 'Select or create a tag'

class ProdTags extends Component {
	
	constructor() {
    super();
    
		this.state = {
      tags: [],
      suggestions: [
        { id: 'two', text: 'Two' },
        { id: 'three', text: 'Three' }
      ]
    }

    this.handleAddition = this.handleAddition.bind(this);
    this.handleDelete = this.handleDelete.bind(this);

  }
  
  handleDelete(i) {
    const { tags } = this.state;
    this.setState({
      tags: tags.filter((tag, index) => index !== i)
    })
  }

  handleAddition(tag) {
    this.setState(state => ({ tags: [...state.tags, tag] }));
  }

  render() {
    const { tags, suggestions } = this.state;
    return (
      <div className="react-tags">
			  <h3>Tags</h3>
        <ReactTags 
          tags={tags}
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

export default ProdTags;
