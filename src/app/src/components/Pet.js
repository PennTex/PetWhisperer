import React from 'react';

export default class Pet extends React.Component {
  render() {
    return (
      <div className="col-xs-4">
        <div className="panel panel-default">
          <div className="panel-heading">{this.props.pet.name} </div>
          <div className="panel-body">
            Id: {this.props.pet}
          </div>
          <div className="panel-footer">
            {this.props.pet} is a cool pet dude
        </div>
        </div>
      </div>);
  }
}