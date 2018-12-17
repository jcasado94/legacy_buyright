import React, { Component } from 'react';

import ProdTitle from './ProdTitle'
import ProdImage from './ProdImage'
import TagInput from './TagInput';

import { TAGS_TITLE, TAGS_PLACEHOLDER} from '../global/global'
import ProdUsagesContainer from '../containers/ProdUsagesContainer';

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
      ]
    }

    this.handleTagAddition = this.handleTagAddition.bind(this)
    this.handleTagDelete = this.handleTagDelete.bind(this);

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
            title={TAGS_TITLE}
            placeholder={TAGS_PLACEHOLDER}
            tagList={this.state.prodTags}
            tagSuggestions={this.state.prodTagsSuggestions}
          />
          <ProdUsagesContainer />
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

}

const listAddition = (tag, list) => {
  return [...list, tag]
}

const listDelete = (i, list) => {
  return list.filter((tag, index) => index !== i)
}


export default Prod;
