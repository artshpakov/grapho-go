import React from 'react';

import './post.sass'


export default class Post extends React.Component {

  render() {
    let content = this.props.content;

    return (
      <div className="post"  key={content.slug}>
        <h3>{ content.title }</h3>
        <p><small>{ content.created_at }</small></p>
        { content.text.slice(0, 350) }
        <hr/>
      </div>
    );
  }

}
