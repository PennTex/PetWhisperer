import React from 'react';

export default class LoggedIn extends React.Component {
  logout(){
    localStorage.removeItem('userToken');
    this.props.lock.logout({
      client_id: AUTH0_CLIENT_ID,
      returnTo:'http://localhost:8080'})
  }

  getInitialState() {
    return {
      profile: null,
      pets: null
    }
  }

  componentDidMount() {
    this.props.lock.getProfile(this.props.idToken, function (err, profile) {
      if (err) {
        console.log("Error loading the Profile", err);
      }
      this.setState({profile: profile});
    }.bind(this));

    this.serverRequest = $.get(PET_WHISPERER_API_BASE + '/pets', function (result) {
      console.log(result)
      this.setState({
        pets: result,
      });
    }.bind(this));
  }

  render() {
    if (this.state.profile) {
      return (
        <div className="col-lg-12">
          <span className="pull-right">{this.state.profile.nickname} <a href="#" onClick={this.logout}>Log out</a></span>
          <h2>Welcome to Pet Whisperer</h2>
          <p>Below you'll find all your pets, cool.</p>
          <div className="row">
          {this.state.pets.map(function(pet, i){
            return <Pet key={i} pet={pet} />
          })}
          </div>
        </div>);
    } else {
      return (<div>Loading...</div>);
    }
  }
}