import { createStore, combineReducers } from 'redux';
import { authors } from './reducers/authors'

export default createStore(combineReducers({ authors }));
