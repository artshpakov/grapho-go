import store from './../../store';
import AuthorProfile from './../author_profile/author_profile.jsx';
import Post from './../post/post.jsx';


export default class PostsList extends React.Component {

  render() {
    let author = store.getState().authors.currentAuthor;

    if (author) {
      let posts = author.posts.map(post => <Post content={post} key={post.slug} />);

      return (
        <div>
          <AuthorProfile />
          { posts }
        </div>
      )
    } else {
      return <div/>
    };
  }

}
