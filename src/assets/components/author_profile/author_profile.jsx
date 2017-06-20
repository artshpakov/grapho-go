import store from './../../store'


export default class AuthorProfile extends React.Component {

  render() {
    let author = store.getState().authors.currentAuthor;

    return (
      <div>
        <h1>{ author.name }</h1>
        <ul className="list-inline">
          <li>{ author.profile.location }</li>
          <li>{ author.email }</li>
        </ul>
      </div>
    );
  }

}
