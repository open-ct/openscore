import React, {Component} from "react";
import {Avatar, Dropdown, Layout, Menu} from "antd";
import DocumentTitle from "react-document-title";
import {Link, Redirect, Route, Switch} from "react-router-dom";
import * as Icon from "@ant-design/icons";
import {LogoutOutlined, SettingOutlined} from "@ant-design/icons";
import * as Setting from "../../Setting";
import * as AccountBackend from "../../backend/AccountBackend";

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

import menuList from "../../menu/menuTab.js";
import normalLogin from "../Login/normaluser";
import logoUrl from "../../asset/images/OpenCT_Logo.png";
import group from "../../api/group";
import "./index.less";

const {Header, Sider, Content} = Layout;
const {SubMenu} = Menu;

export default class index extends Component {
    state = {
      account: null,
      openKeys: [],
      selectedKeys: [],
      userInfo: {},
    };
    permissionList = menuList

    componentDidMount() {
      this.getAccount();
      this.userInfo();
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
          openKeys: rank[1],
        });
        break;
      }

    }
    userInfo = () => {
      group.userInfo({supervisorId: "1"})
        .then((res) => {
          if (res.data.status === "10000") {
            this.setState({
              userInfo: res.data.data.userInfo,
            });
            localStorage.setItem("userInfo", JSON.stringify(res.data.data.userInfo));
          }
        })
        .catch((e) => {
          Setting.showMessage("error", e);
        });
    }
    getAccount() {
      AccountBackend.getAccount()
        .then((res) => {
          this.setState({
            account: res.data,
          });
          localStorage.setItem("account", JSON.stringify(this.state.account));
        });
    }

    logout() {
      this.setState({
        expired: false,
        submitted: false,
      });

      AccountBackend.signout()
        .then((res) => {
          localStorage.setItem("account", "");
          if (res.status === "ok") {
            this.setState({
              account: null,
            });

            Setting.showMessage("success", "Successfully logged out, redirected to homepage");

            Setting.goToLink("/");
          } else {
            Setting.showMessage("error", `Logout failed: ${res.msg}`);
          }
        });
    }

    handleRightDropdownClick(e) {
      if (e.key === "0") {
        Setting.openLink(Setting.getMyProfileUrl(this.state.account));
      } else if (e.key === "1") {
        this.logout();
      }
    }

    renderAvatar() {
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
      if (this.state.account === undefined || this.state.account === null) {
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
      if(this.state.account) {
        return menulist.map((item) => {
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
      console.log(selectedKeys.keyPath);
    }

    render() {
      const {openKeys, selectedKeys} = this.state;
      console.log(openKeys, selectedKeys);
      return (
        <DocumentTitle title="阅卷系统">
          <Layout className="home-page" data-component="home-page">
            <Header>
              <div className="header-box">
                <div className="header-logo">
                  <img src={logoUrl} alt="" />
                  <span className="header-title">OpenCT在线阅卷系统</span>
                </div>

                <div className="header-info">
                  <span className="header-teacher">教师：小屋</span>
                  <span className="header-teacher">任务：正评卷</span>
                  <span className="header-teacher">题目：第一题</span>
                  <span className="header-teacher">评卷数量：201</span>
                  <span className="header-teacher">平均速度：6.5秒/份</span>
                  <span className="header-teacher">当前密号：2008886</span>
                </div>
                {
                  this.renderAccount()
                }
              </div>

            </Header>
            <Layout className="container">
              <Sider>
                <Menu
                  onOpenChange={this.onOpenChange.bind(this)}
                  style={{width: 200, height: "100%"}}
                  openKeys={openKeys}
                  selectedKeys={selectedKeys}
                  onClick={this.onClick}
                  mode="inline"
                >
                  {
                    this.bindMenu(this.permissionList)
                  }

                </Menu>
              </Sider>
              <Content>
                <Switch>
                  {this.permissionList[0].userPermission === "阅卷员" ? <>
                    {this.openKeys === [] ? <Redirect from="/home" to="/home/mark-tasks" exact></Redirect> : null}
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
                    <Route path="/home/userManagement/user/user_manage" component={userManage} exact></Route>

                    <Route path="/home/normaluser" component={normalLogin} exact></Route>
                  </>

                    : null
                  }
                </Switch>
              </Content>
            </Layout>
          </Layout>
        </DocumentTitle>
      );
    }
}
