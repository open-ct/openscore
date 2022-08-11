import {BrowserRouter, Route} from "react-router-dom";
import "./App.less";
import AuthCallback from "./auth/AuthCallback";
import Home from "./views/Home";
import Login from "./views/Login/chooserole";
import normalLogin from "./views/Login/normaluser";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Route path="/" exact component={Login} />
        <Route path="/login" component={normalLogin} />
        <Route path="/callback" component={AuthCallback} />
        <Route path="/home" component={Home} />
      </BrowserRouter>
    </div>
  );
}

export default App;
