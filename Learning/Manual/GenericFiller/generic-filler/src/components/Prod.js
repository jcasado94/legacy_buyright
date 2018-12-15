import React, { Component } from 'react';

import ProdTitle from './ProdTitle'
import ProdImage from './ProdImage'
import TagInput from './TagInput';

import messages from './messages'
import ProdUsages from './ProdUsages';

class Prod extends Component {

  constructor() {
    super();

    this.state = {
      prodTags: [],
      prodTagsSuggestions: [
        {id: 'one', text: 'One'},
        {id: 'two', text: 'Two'},
        {id: 'three', text: 'Three'},
        {id: 'four', text: 'Four'},
        {id: 'five', text: 'Five'},
      ],
      prodUsages: [
        'one',
        'two'
      ],
      allUsages: {
        'one': 'One',
        'two': 'Two',
        'three': 'Three',
        'four': 'Four',
        'five': 'Five'
      }
    }

    this.handleTagAddition = this.handleTagAddition.bind(this)
    this.handleTagDelete = this.handleTagDelete.bind(this);
    this.handleUsageAddition = this.handleUsageAddition.bind(this);
    this.handleUsageDelete = this.handleUsageDelete.bind(this);

  }

	render() {
		return (
      <div className="Prod">
        <ProdTitle />
        <ProdImage />
        <div className="prod-properties">
          <TagInput
            onTagAddition={this.handleTagAddition}
            onTagDelete={this.handleTagDelete}
            title={messages.TAGS_TITLE}
            placeholder={messages.TAGS_PLACEHOLDER}
            tagList={this.state.prodTags}
            tagSuggestions={this.state.prodTagsSuggestions}
          />
          <ProdUsages
            allUsages={this.state.allUsages}
            prodUsages={this.state.prodUsages}
          />
        </div>
      </div>
		);
  }

  handleTagAddition(tag) {
    this.setState({
      prodTags: listAddition(tag, this.state.prodTags)
    })
  }

  handleTagDelete(i) {
    this.setState({
      prodTags: listDelete(i, this.state.prodTags)
    })
  }

  handleUsageAddition(tag) {
    this.setState({
      prodUsages: listAddition(tag, this.state.prodUsages)
    })
  }

  handleUsageDelete(i) {
    this.setState({
      prodUsages: listDelete(i, this.state.prodUsages)
    })
  }

}

const listAddition = (tag, list) => {
  return [...list, tag]
}

const listDelete = (i, list) => {
  return list.filter((tag, index) => index !== i)
}


export default Prod;
