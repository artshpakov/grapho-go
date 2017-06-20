export const authors = (state, action) => {
  switch (action.type) {

    case 'ADD_AUTHOR':
      let data    = _.isArray(action.data) ? action.data : [action.data];
      let authors = data.map(person => _.extend({}, person, {active: true}))
      state.items = authors;
      return state;

    case 'SELECT_AUTHOR':
      state.currentAuthor = action.data;
      return state;

    default:
      return state || [];

  }
}

export const addAuthorAction = authors => ({type: 'ADD_AUTHOR', data: authors });
export const selectAuthorAction = author => ({type: 'SELECT_AUTHOR', data: author });
