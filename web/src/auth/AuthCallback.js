/*
 * @Author: Junlang
 * @Date: 2021-07-22 18:39:51
 * @LastEditTime: 2021-07-23 15:48:38
 * @LastEditors: Junlang
 * @FilePath: /casbin-oa/web/src/auth/AuthCallback.js
 */
// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from "react";
import {Button, Result, Spin} from "antd";
import {withRouter} from "react-router-dom";
import * as AccountBackend from "./AccountBackend";
import * as Settings from "../Setting";

class AuthCallback extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      msg: null,
    };
  }

  componentDidMount() {
    this.login();
  }

  login() {
    console.log(this.props);
    const params = new URLSearchParams(this.props.location.search);
    console.log(params.get("code"));
    AccountBackend.login(params.get("code"), params.get("state"))
      .then((res) => {
        console.log(res);
        if (res.status === "ok") {
          Settings.showMessage("success", "Logged in successfully");
          Settings.goToLink("/home");
        } else {
          this.setState({
            msg: res.msg,
          });
        }
      });
  }

  render() {
    return (
      <div style={{textAlign: "center"}}>
        {
          (this.state.msg === null) ? (
            <Spin size="large" tip="Signing in..." style={{paddingTop: "10%"}} />
          ) : (
            <div style={{display: "inline"}}>
              <Result
                status="error"
                title="Login Error"
                subTitle={this.state.msg}
                extra={[
                  <Button type="primary" key="details">
                    Details
                  </Button>,
                  <Button key="help">Help</Button>,
                ]}
              >
              </Result>
            </div>
          )
        }
      </div>
    );
  }
}

export default withRouter(AuthCallback);
