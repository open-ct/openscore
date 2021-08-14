import { Route, BrowserRouter, Switch } from "react-router-dom";

import './App.less';
import AuthCallback from "./auth/AuthCallback";
import Home from './views/Home'

function App() {
  return (
      <div className="App">
        <BrowserRouter>
          <Route path="/" component={Home} />
          <Route exact path="/callback" component={AuthCallback} />
        </BrowserRouter>
      </div>
  );
}

export default App;