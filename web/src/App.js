import React from "react";
import { Route, BrowserRouter } from "react-router-dom";
import './App.less';
import AuthCallback from "./AuthCallback";
import Home from './views/Home';
import * as Setting from "./Setting";
import * as Conf from "./Conf";

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
    };

    Setting.initServerUrl();
    Setting.initCasdoorSdk(Conf.AuthConfig);
  }

  render() {
    return (
      <div className="App">
        <BrowserRouter>
          <Route path="/" component={Home} />
          <Route exact path="/callback" component={AuthCallback} />
        </BrowserRouter>
      </div>
    );
  }
}

export default App;
