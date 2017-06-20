import { render } from 'react-dom';

import store from './store'
import Authors from './components/authors/authors.jsx';
import PostsList from './components/posts_list/posts_list.jsx';


const run = () => {
  render(<Authors />, document.getElementById('authors'));
  render(<PostsList />, document.getElementById('posts'));
}
run();
store.subscribe(run);
