import React from 'react';

export default class App extends React.Component {
  componentWillMount() {
    this.setupAjax();
    this.createLock();
    this.setState({idToken: this.getIdToken()})
  }

  createLock() {
    this.lock = new Auth0Lock(this.props.clientId, this.props.domain);
  }

  setupAjax() {
    $.ajaxSetup({
      'beforeSend': function(xhr) {
        if (localStorage.getItem('userToken')) {
          xhr.setRequestHeader('Authorization',
                'Bearer ' + localStorage.getItem('userToken'));
        }
      }
    });
  }

  getIdToken() {
    var idToken = localStorage.getItem('userToken');
    var authHash = this.lock.parseHash(window.location.hash);
    if (!idToken && authHash) {
      if (authHash.id_token) {
        idToken = authHash.id_token
        localStorage.setItem('userToken', authHash.id_token);
        window.location.replace("http://localhost:8080");
      }
      if (authHash.error) {
        console.log("Error signing in", authHash);
      }
    }
    return idToken;
  }

  render() {
    if (this.state.idToken) {
      return (<LoggedIn lock={this.lock} idToken={this.state.idToken} />);
    } else {
      return (<Home lock={this.lock} />);
    }
  }
}