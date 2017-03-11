import React from 'react';
import * as config from '../config';
import Pet from './Pet';

export default class LoggedIn extends React.Component {
  constructor(props) {
    super(props);
    
    this.state = {
      profile: null,
      pets: null
    }
  }

  logout() {
    console.log(this.props);
    localStorage.removeItem('userToken');
    this.props.lock.logout({
      client_id: config.AUTH0_CLIENT_ID,
      returnTo: 'http://localhost:8080'
    })
  }

  componentDidMount() {
    this.props.lock.getProfile(this.props.idToken, function (err, profile) {
      if (err) {
        console.log("Error loading the Profile", err);
      }
      this.setState({ profile: profile });
    }.bind(this));

    this.serverRequest = $.get(config.PET_WHISPERER_API_BASE + '/pets', function (result) {
      console.log(result)
      this.setState({
        pets: result,
      });
    }.bind(this));
  }

  render() {    
    if (this.state.profile) {
      let pets = '';

      if (this.state.profile.pets) {
        console.log(this.state.profile.pets);
        pets = this.state.profile.pets.map(function (pet, i) {
          return <Pet key={i} pet={pet} />
        })
      }

      return (
        <div className="col-lg-12">
          <span className="pull-right">{this.state.profile.nickname} <a href="#" onClick={this.logout.bind(this)}>Log out</a></span>
          <h2>Welcome to Pet Whisperer</h2>
          <p>Below you'll find all your pets, cool.</p>
          <div className="row">
            {pets}
          </div>
        </div>);
    } else {
      return (<div>Loading...</div>);
    }
  }
}