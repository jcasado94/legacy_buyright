import { combineReducers } from 'redux'
import {
  ADD_USAGE,
  REMOVE_USAGE
} from '../actions/actions'

const defaultUsagesState = {
  prodUsages: []
}

const usages = (state = defaultUsagesState, action) => {
  switch (action.type) {
    case ADD_USAGE:
      return {
        prodUsages: [
          ...state.prodUsages, 
          action.usageId
        ]
      }
    case REMOVE_USAGE:
      return {
        prodUsages: [
          ...state.prodUsages.slice(0, action.i),
          ...state.prodUsages.slice(action.i+1)
        ]
      }
    default:
      return state;
  }
}

const app = combineReducers({
  usages
});

export default app;