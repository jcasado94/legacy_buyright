import React, { Component } from 'react';
import { WithContext as ReactTags } from 'react-tag-input';


class TagInput extends Component {
	
	constructor() {
    super();

    this.handleAddition = this.handleAddition.bind(this);
    this.handleDelete = this.handleDelete.bind(this);

  }
  
  handleDelete(i) {
    this.props.onTagDelete(i);
  }

  handleAddition(tag) {
    this.props.onTagAddition(tag);
  }

  render() {
    return (
      <div className="react-tags">
			  <h3>{this.props.title}</h3>
        <ReactTags 
          tags={this.props.tagList}
          suggestions={this.props.tagSuggestions}
          handleAddition={this.handleAddition}
          handleDelete={this.handleDelete}
          allowDragDrop={false}
          placeholder={this.props.placeholder}
          inline={false}
          minQueryLength={0}
        />
      </div>
    );
  }
}

export default TagInput;
