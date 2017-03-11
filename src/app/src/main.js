import App from './components/App';
import * as config from './config';

ReactDOM.render(<App clientId={config.AUTH0_CLIENT_ID} domain={config.AUTH0_DOMAIN} />,
  document.getElementById('app'));