import React, {Component} from "react";
import {Avatar, Dropdown, Layout, Menu} from "antd";
import DocumentTitle from "react-document-title";
import {Link, Redirect, Route, Switch} from "react-router-dom";
import * as Icon from "@ant-design/icons";
import {LogoutOutlined, SettingOutlined} from "@ant-design/icons";
import * as Setting from "../../Setting";
import * as AccountBackend from "../../backend/AccountBackend";
import Context from "../../util/Context";

import MarkTasks from "../Mark/MarkTasks";
import Answer from "../Mark/Answer";
import Sample from "../Mark/Sample";
import Review from "../Mark/Review";

import all from "../Group/mark_monitor/all";
import average from "../Group/mark_monitor/average";
import score from "../Group/mark_monitor/score";
import self from "../Group/mark_monitor/self";
import standard from "../Group/mark_monitor/standard";
import teacher from "../Group/mark_monitor/teacher";

import arbitration from "../Group/test_monitor/arbitration";
import marking from "../Group/test_monitor/marking";
import problem from "../Group/test_monitor/problem";
import markTasks from "../Group/test_monitor/markTasks";

import question from "../Manage/paper_manage/question";
import paper from "../Manage/paper_manage/paper";
import allot from "../Manage/paper_manage/allot";
import paperManage from "../Manage/paper_manage/manage";
import detail from "../Manage/paper_manage/detail";
import userManage from "../Manage/user_manage/user";

import userDetail from "../Train/detail";
import evaluate from "../Train/evaluate";
import grouping from "../Train/grouping";

import menuList from "../../menu/menuTab.js";
import normalLogin from "../Login/normaluser";

import "./index.less";

const {Header, Sider, Content} = Layout;
const {SubMenu} = Menu;

export default class index extends Component {
    state = {
      account: {},
      openKeys: [],
      selectedKeys: [],
      role: "",
    };
    permissionList = menuList

    componentDidMount() {
      this.getAccount();
      const pathname = this.props.location.pathname;
      const rank = pathname.split("/").slice(-2).reverse();
      switch (rank.length) {
      case 1:
        this.setState({
          openKeys: rank,
        });
        break;
      case 2:
        this.setState({
          selectedKeys: rank,
          openKeys: [rank[1]],
        });
        break;
      }
    }

    componentDidUpdate() {
      if (typeof this.props.location.state !== "undefined" && this.state.userInfo === null) {
        this.setState({account: this.props.location.state.userInfo, role: "user"});
      }
    }

    getAccount() {
      AccountBackend.getAccount()
        .then((res) => {
          if (res.status === "ok") {
            if (res.data.type === "normal-user") {
              this.setState({
                account: res.data,
                role: "admin",
              });
            } else {
              this.setState({
                account: res.data,
                role: "user",
              });
            }

          }
        });
    }

    logout() {
      AccountBackend.signout()
        .then((res) => {
          if (res.status === "ok") {
            this.setState({
              account: {},
              role: "",
            });

            Setting.showMessage("success", "Successfully logged out, redirected to homepage");

            Setting.goToLink("/");
          } else {
            Setting.showMessage("error", `Logout failed: ${res.msg}`);
          }
        });
      this.props.history.push("/home");
    }

    handleRightDropdownClick(e) {
      if (e.key === "0") {
        Setting.openLink(Setting.getMyProfileUrl(this.state.account));
      } else if (e.key === "1") {
        this.logout();
      }
    }

    renderAvatar() {
      if (this.state.role === "admin") {
        if (this.state.account.avatar === "") {
          return (
            <Avatar style={{backgroundColor: Setting.getAvatarColor(this.state.account.name), verticalAlign: "middle"}} size="large">
              {Setting.getShortName(this.state.account.name)}
            </Avatar>
          );
        } else {
          return (
            <Avatar src={this.state.account.avatar} style={{verticalAlign: "middle"}} size="large">
              {Setting.getShortName(this.state.account.name)}
            </Avatar>
          );
        }
      } else {
        return (
          <Avatar style={{verticalAlign: "middle"}} size="large">
          </Avatar>);
      }

    }

    renderRightDropdown() {
      const menu = (
        <Menu onClick={this.handleRightDropdownClick.bind(this)}>
          <Menu.Item key="0">
            <SettingOutlined />
                    My Account
          </Menu.Item>
          <Menu.Item key="1">
            <LogoutOutlined />
                    Logout
          </Menu.Item>
        </Menu>
      );

      return (
        <Dropdown key="200" overlay={menu} >
          <a className="ant-dropdown-link" href="#" style={{float: "right", marginLeft: "50px"}}>
            {
              this.renderAvatar()
            }
          </a>
        </Dropdown>
      );
    }

    renderAccount() {
      if (this.state.role === "") {
        return (
          <>
            <a href={Setting.getSigninUrl()} style={{color: "#ffffff", marginLeft: "50px"}}>
                管理员登录
            </a>
            <Link
              to={"/home/normaluser"} style={{color: "#ffffff", marginLeft: "50px"}}>
                组长/阅卷老师登录
            </Link>
          </>
        );
      } else {
        return (
          this.renderRightDropdown()
        );
      }

    }

    bindMenu = (menulist) => {
      if (this.state.role === "admin") {
        return menulist.filter((item) => {
          return item.userPermission === "管理员";
        }).map((item) => {
          if (item.chidPermissions.length === 0) {  // 没有子菜单
            return <Menu.Item key={item.key}
              icon={item.icon ? React.createElement(Icon[item.icon]) : null}><Link
                to={item.menu_url}>{item.menu_name}</Link></Menu.Item>;
          } else {
            return <SubMenu key={item.key} title={item.menu_name} icon={React.createElement(Icon[item.icon])}>
              {this.bindMenu(item.chidPermissions)}
            </SubMenu>;
          }
        });
      } else if(this.state.role === "user") {
        if (this.state.account.type === "normal") {
          return menulist.filter((item) => {
            return item.userPermission === "阅卷员";
          }).map((item) => {
            if (item.chidPermissions.length === 0) {  // 没有子菜单
              return <Menu.Item key={item.key}
                icon={item.icon ? React.createElement(Icon[item.icon]) : null}><Link
                  to={item.menu_url}>{item.menu_name}</Link></Menu.Item>;
            } else {
              return <SubMenu key={item.key} title={item.menu_name} icon={React.createElement(Icon[item.icon])}>
                {this.bindMenu(item.chidPermissions)}
              </SubMenu>;
            }
          });
        } else {
          return menulist.filter((item) => {
            return item.userPermission === "组长";
          }).map((item) => {
            if (item.chidPermissions.length === 0) {  // 没有子菜单
              return <Menu.Item key={item.key}
                icon={item.icon ? React.createElement(Icon[item.icon]) : null}><Link
                  to={item.menu_url}>{item.menu_name}</Link></Menu.Item>;
            } else {
              return <SubMenu key={item.key} title={item.menu_name} icon={React.createElement(Icon[item.icon])}>
                {this.bindMenu(item.chidPermissions)}
              </SubMenu>;
            }
          });
        }
      } else {return null;}

    }

    onOpenChange = (openKeys) => {
      if (openKeys.length === 1 || openKeys.length === 0) {
        this.setState({
          openKeys,
        });
        return;
      }
      const latestOpenKey = openKeys[openKeys.length - 1];
      if (latestOpenKey.includes(openKeys[0])) {
        this.setState({
          openKeys,
        });
      } else {
        this.setState({
          openKeys: [latestOpenKey],
        });
      }
    }

    onClick = (selectedKeys) => {
      this.setState({selectedKeys: selectedKeys.keyPath});
    }

    render() {
      const {openKeys, selectedKeys, account} = this.state;
      return (
        <Context.Provider value={account}>
          <DocumentTitle title="阅卷系统">
            <Layout className="home-page" data-component="home-page">
              <Header>
                <div className="header-box">
                  <div className="header-logo">
                    <img src="https://score-cdn.open-ct.com/OpenCT_Logo.png" alt="" />
                    <span className="header-title">OpenCT在线阅卷系统</span>
                  </div>

                  <div className="header-info">
                    {/* <span className="header-teacher">教师：小屋</span>
                    <span className="header-teacher">任务：正评卷</span>
                    <span className="header-teacher">题目：第一题</span>
                    <span className="header-teacher">评卷数量：201</span>
                    <span className="header-teacher">平均速度：6.5秒/份</span>
                    <span className="header-teacher">当前密号：2008886</span> */}
                  </div>
                  {
                    this.renderAccount()
                  }
                </div>

              </Header>
              <Layout className="container">
                {
                  this.state.role === "" && this.props.location.pathname === "/" ?
                    <Content>
                      <div className="content-notlogin">
                        <img src="https://score-cdn.open-ct.com/openct.png" alt="openCT_POST" />
                        <p>OpenScore阅卷系统</p>
                      </div>
                    </Content>
                    : <>
                      <Sider>
                        <Menu
                          onOpenChange={this.onOpenChange.bind(this)}
                          style={{width: 200, height: "100%"}}
                          selectedKeys={selectedKeys}
                          openKeys={openKeys}
                          onClick={this.onClick}
                          mode="inline"
                        >
                          {
                            this.bindMenu(this.permissionList)
                          }

                        </Menu>
                      </Sider>
                    </>
                }
                <Content>
                  <Switch>
                    <Route path="/home/normaluser" component={normalLogin} exact></Route>
                    {JSON.stringify(this.state.account) == "{}" ? <Redirect from="/" to="/"></Redirect> : null}
                    <Route path="/home/mark-tasks" component={MarkTasks} exact></Route>
                    <Route path="/home/answer" component={Answer} exact></Route>
                    <Route path="/home/sample" component={Sample} exact></Route>
                    <Route path="/home/review" component={Review} exact></Route>
                    {/* <Route path="/home/selfMark" component={SelfMark} exact></Route> */}

                    <Route path="/home/allMarkMonitor/all" component={all} exact></Route>
                    <Route path="/home/allMarkMonitor/average" component={average} exact></Route>
                    <Route path="/home/allMarkMonitor/score" component={score} exact></Route>
                    <Route path="/home/allMarkMonitor/self" component={self} exact></Route>
                    <Route path="/home/allMarkMonitor/standard" component={standard} exact></Route>
                    <Route path="/home/allMarkMonitor/teacher" component={teacher} exact></Route>

                    <Route path="/home/group/arbitration" component={arbitration} exact></Route>
                    <Route path="/home/group/marking" component={marking}></Route>
                    <Route path="/home/group/problem" component={problem} exact></Route>
                    <Route path="/home/group/markTasks/:type/:QuestionId" component={markTasks} exact></Route>

                    <Route path="/home/paperManagement/question" component={question} exact></Route>
                    <Route path="/home/paperManagement/paper" component={paper}></Route>
                    <Route path="/home/paperManagement/paper_allot" component={allot} exact></Route>
                    <Route path="/home/userManagement/paper_manage" component={paperManage} exact></Route>
                    <Route path="/home/userManagement/detailTable" component={detail} exact></Route>
                    <Route path="/home/userManagement/userManage" component={userManage} exact></Route>

                    <Route path="/home/userTrain/grouping" component={grouping} exact></Route>
                    <Route path="/home/userTrain/detail" component={userDetail} exact></Route>
                    <Route path="/home/userTrain/evaluate" component={evaluate} exact></Route>
                  </Switch>
                </Content>
              </Layout>
            </Layout>
          </DocumentTitle>
        </Context.Provider>
      );
    }
}
