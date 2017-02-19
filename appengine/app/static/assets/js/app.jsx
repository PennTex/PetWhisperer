var App = React.createClass({
  componentWillMount: function() {
    this.setupAjax();
    this.createLock();
    this.setState({idToken: this.getIdToken()})
  },
  createLock: function() {
    this.lock = new Auth0Lock(this.props.clientId, this.props.domain);
  },
  setupAjax: function() {
    $.ajaxSetup({
      'beforeSend': function(xhr) {
        if (localStorage.getItem('userToken')) {
          xhr.setRequestHeader('Authorization',
                'Bearer ' + localStorage.getItem('userToken'));
        }
      }
    });
  },
  getIdToken: function() {
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
  },
  render: function() {
    if (this.state.idToken) {
      return (<LoggedIn lock={this.lock} idToken={this.state.idToken} />);
    } else {
      return (<Home lock={this.lock} />);
    }
  }
});

var Home = React.createClass({
  componentDidMount: function() {
    this.props.lock.show({
      closable: false
    });
  },

  render: function() {
    return null;
  }
});

var LoggedIn = React.createClass({
  logout : function(){
    localStorage.removeItem('userToken');
    this.props.lock.logout({
      client_id: AUTH0_CLIENT_ID,
      returnTo:'http://localhost:8080'})
  },

  getInitialState: function() {
    return {
      profile: null,
      pets: null
    }
  },

  componentDidMount: function() {
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
  },

  render: function() {
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
});

var Pet = React.createClass({
  render : function(){
    return(
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
})

ReactDOM.render(<App clientId={AUTH0_CLIENT_ID} domain={AUTH0_DOMAIN} />,
  document.getElementById('app'));
