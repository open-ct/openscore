import React from "react";
import {BrowserRouter, Route} from "react-router-dom";
import "./App.less";
import AuthCallback from "./AuthCallback";
import Home from "./views/Home";
import * as Setting from "./Setting";
import * as Conf from "./Conf";
import normalLogin from "./views/Login/normaluser";

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
          <Route path="/normalLogin" component={normalLogin} />
          <Route path="/home" component={Home} />
        </BrowserRouter>
      </div>
    );
  }
}

export default App;
