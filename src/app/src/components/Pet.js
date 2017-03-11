import React from 'react';

export default class Pet extends React.Component {
  render() {
    return (
      <div className="col-xs-4">
        <div className="panel panel-default">
          <div className="panel-heading">{this.props.pet.name} </div>
          <div className="panel-body">
            Created: {this.props.pet.created_at}
          </div>
          <div className="panel-footer">
            {this.props.pet.name} is a cool pet dude
        </div>
        </div>
      </div>);
  }
}