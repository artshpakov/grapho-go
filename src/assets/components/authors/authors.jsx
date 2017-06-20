import store from './../../store'
import { addAuthorAction, selectAuthorAction } from './../../reducers/authors'

import './authors.sass';


const AUTHORS_QUERY = `{
  users {
    id
    name
    email
    posts {
      title
      text
      slug
      created_at
    }
    profile {
      location
    }
  }
}`;


export default class Authors extends React.Component {

  load() {
    $.get("/graph?q="+AUTHORS_QUERY, (reply) => {
      store.dispatch(addAuthorAction(reply.data.users));
    })
  }

  componentDidMount() {
    this.serverRequest = this.load()
  }

  componentWillUnmount() {
    this.serverRequest.abort();
  }

  pickAuthor(author) {
    this.currentAuthor = author;
    store.dispatch(selectAuthorAction(author));
  }

  render() {
    let authors = _(store.getState().authors.items).where({active: true}).map(person => (
      <li key={person.id} className={ person == this.currentAuthor ? 'active' : '' }>
        <a href="javascript:;" onClick={this.pickAuthor.bind(this, person)}>{ person.name }</a>
      </li>
    ));

    return (
      <ul className="list-unstyled">{authors}</ul>
    );
  }

}
