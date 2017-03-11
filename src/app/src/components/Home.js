import React from 'react';

export default class Home extends React.Component {
  componentDidMount() {
    this.props.lock.show({
      closable: false
    });
  }

  render() {
    return null;
  }
}