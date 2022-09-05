import {Button, Form, Input} from "antd";
import React, {Component} from "react";
import "./index.less";
import * as Setting from "../../../Setting";
import axios from "axios";
import * as Settings from "../../../Setting";
import group from "../../../api/group";

export async function login(account, password) {
  return await new Promise((resolve) => {
    axios.post(`${Setting.ServerUrl}/login`, {
      account: account,
      password: password,
    }).then(res => {
      resolve(res.data);
    }).catch(error => {
      Settings.showMessage("error", error);
    });
  });
}

export default class normalLogin extends Component {
    loginUser = () => {}
    onFinish = (values) => {
      group.UserLogin(values
      ).then(res => {
        if (res.data.status === "ok") {
          Settings.showMessage("success", "Logged in successfully");
          window.location.href = `http://${location.host}`;
        }
      });
    };

    render() {
      return (
        <div className="bg">
          <div className="login_card">
            <Form
              name="normal_login"
              className="login-form"
              initialValues={{remember: true}}
              onFinish={this.onFinish}
            >
              <Form.Item
                label="用户名"
                name="account"
                rules={[{required: true, message: "Please input your username!"}]}
              >
                <Input />
              </Form.Item>

              <Form.Item
                label="密码"
                name="password"
                rules={[{required: true, message: "Please input your password!"}]}
              >
                <Input.Password />
              </Form.Item>

              <Form.Item>
                <Button type="primary" htmlType="submit" block>
                                登录
                </Button>
              </Form.Item>
            </Form>
          </div>
        </div>
      );
    }
}
