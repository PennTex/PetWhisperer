const AUTH0_CLIENT_ID='nNeQx06JEjFktQLEpVCM48lsMsbB3739'; 
const AUTH0_DOMAIN='pet-whisperer.auth0.com'; 
const AUTH0_CALLBACK_URL='http://localhost:8080';
const PET_WHISPERER_API_BASE="http://localhost:8081";

import App from './components/App';

ReactDOM.render(<App clientId={AUTH0_CLIENT_ID} domain={AUTH0_DOMAIN} />,
  document.getElementById('app'));