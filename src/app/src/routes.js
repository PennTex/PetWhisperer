import React from 'react'
import { Route, IndexRoute } from 'react-router'
import App from './components/App';
import Home from './components/Home';
import LoggedIn from './components/LoggedIn';
import Pet from './components/Pet';

const routes = (
  <Route path="/" component={App}>
    <IndexRoute component={Home}/>
  </Route>
);

export default routes;