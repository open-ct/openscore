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
// import './App.less';
// import React, { Component } from 'react'
// import {Switch,Route,Redirect,BrowserRouter as Router} from 'react-router-dom'

// import Login from './views/Login'
// import Home from './views/Home'
// import ProjectManagement from './views/ProjectManagement'

// export default class App extends Component {
//   render() {
//     return (
//       <div id="App">
//         <Router>
//           <Switch>
//             <Redirect from="/" to="/home" exact></Redirect>
//             <Route path="/project-management" component={ProjectManagement} exact></Route>
//             <Route path="/login" component={Login} exact></Route>
//             <Route path="/home" component={Home}></Route>
//           </Switch>
//         </Router>
//       </div>
//     )
//   }
// }

