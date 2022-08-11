import React, {Component} from "react";
import {Dropdown, Layout, Menu} from "antd";
import DocumentTitle from "react-document-title";
import {Link, Redirect, Route, Switch} from "react-router-dom";
import * as Icon from "@ant-design/icons";
import {LogoutOutlined, SettingOutlined} from "@ant-design/icons";
import * as Auth from "../../auth/Auth";
import * as Setting from "../../Setting";
import * as AccountBackend from "../../auth/AccountBackend";

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
import user_manage from "../Manage/user_manage/user_manage";

import menuList from "../../menu/menuTab.js";

import logoUrl from "../../asset/images/OpenCT_Logo.png";
import group from "../../api/group";
import "./index.less";
const {Header, Sider, Content} = Layout;
const {SubMenu} = Menu;

export default class index extends Component {

    state = {
      account: null,
      current: "home",

    };
    permissionList = menuList

    componentDidMount() {
      console.log(this.props);
      this.getAccount();
      this.userInfo();
      setTimeout(() => {
        console.log(this.state);
      }, 5000);
      let defaultpage = "/home/mark-tasks";
      console.log(defaultpage);
      this.setState(
        {current: defaultpage.substring(defaultpage.lastIndexOf("/") + 1, defaultpage.length)}
      );

      this.props.history.listen((e) => {
        let test = e.pathname;
        let text = test.substring(test.lastIndexOf("/") + 1, test.length);
        this.setState({
          current: text,
        });
      });
    }

    userInfo = () => {
      group.userInfo({supervisorId: 1})
        .then((res) => {
          if (res.data.status === "10000") {
            this.setState({
              userInfo: res.data.data.userInfo,
            });
            console.log(res.data.data.userInfo);
            localStorage.setItem("userInfo", JSON.stringify(res.data.data.userInfo));
          }
        })
        .catch((e) => {
          console.log(e);
        });
      const userInfo = {
        userName: "headman01",
        subjectName: "science",
      };
      localStorage.setItem("userInfo", JSON.stringify(userInfo));
    }
    getAccount() {
      if(localStorage.getItem("account") === "") {
        AccountBackend.getAccount()
          .then((res) => {
            this.setState({
              account: res.data,
            });
            localStorage.setItem("account", JSON.stringify(this.state.account));
          });
      } else{
        this.setState({
          account: {
            user_type: localStorage.getItem("account"),
            token: localStorage.getItem("token"),
          },
        });
      }

    }

    logout() {
      this.setState({
        expired: false,
        submitted: false,
      });

      AccountBackend.logout()
        .then((res) => {
          localStorage.setItem("account", "");
          if (res.status === "ok") {
            this.setState({
              account: null,
            });

            Setting.showMessage("success", "Successfully logged out, redirected to loginpage");

            Setting.goToLink("/");
          } else {
            Setting.showMessage("error", `Logout failed: ${res.msg}`);
          }
        });
    }

    handleRightDropdownClick(e) {
      if (e.key === "0") {
        Setting.openLink(Auth.getMyProfileUrl(this.state.account));
      } else if (e.key === "1") {
        this.logout();
      }
    }
    /*
        renderAvatar() {
            if (this.state.account.avatar === "") {
                return (
                    <Avatar style={{ backgroundColor: Util.getAvatarColor(this.state.account.name), verticalAlign: 'middle' }} size="large">
                        {Util.getShortName(this.state.account.name)}
                    </Avatar>
                )
            } else {
                return (
                    <Avatar src={this.state.account.avatar} style={{ verticalAlign: 'middle' }} size="large">
                        {Util.getShortName(this.state.account.name)}
                    </Avatar>
                )
            }
        }
    
     */

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
                    用户
          </a>
        </Dropdown>
      );
    }

    renderAccount() {
      if (this.state.account === undefined || this.state.account === null) {
        return "";
      } else {
        return (
          this.renderRightDropdown()
        );
      }
    }

    bindMenu = (menulist) => {
      if(this.state.account) {
        // let user_type ="3";
        // if(this.state.account.owner ===null||this.state.account.owner ===undefined)
        //     user_type=this.state.account.user_type;
        // else {
        //     user_type = "3";
        // }
        let menu;
        if(this.state.account.user_type === "2") {
          menu = menulist.filter(item => item.userPermission === "阅卷员");
        }else if(this.state.account.user_type === "1") {
          menu = menulist.filter(item => item.userPermission === "组长");
        }else if(this.state.account.user_type) {
          menu = menulist.filter(item => item.userPermission === "管理员");
        }else{
          menu = menulist.filter(item => item.userPermission === "管理员");
        }
        return menu.map((item) => {
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
      console.log(openKeys);
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

    render() {
      const {openKeys} = this.state;
      return (
        <DocumentTitle title="阅卷系统">
          <Layout data-component="home-page">
            <Header>
              <div className="header-box">
                <div className="header-logo">
                  <img src={logoUrl} alt="" />
                  <span className="header-title">OpenCT在线阅卷系统</span>
                </div>

                {/* <div className="header-info">
                                <span className="header-teacher">教师：小屋</span>
                                <span className="header-teacher">任务：正评卷</span>
                                <span className="header-teacher">题目：第一题</span>
                                <span className="header-teacher">评卷数量：201</span>
                                <span className="header-teacher">平均速度：6.5秒/份</span>
                                <span className="header-teacher">当前密号：2008886</span>
                            </div> */}
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
                  defaultOpenKeys={this.state.openKeys}
                  defaultSelectedKeys={["mark-tasks"]}
                  selectedKeys={[this.state.current]}
                  mode="inline"
                >
                  {/* <Menu.Item key="mark-tasks" icon={<FormOutlined />}>评卷</Menu.Item>
                                <Menu.Item key="answer" icon={<CheckCircleOutlined />}>答案</Menu.Item>
                                <Menu.Item key="sample" icon={<ProfileOutlined />}>样卷</Menu.Item>
                                <Menu.Item key="review" icon={<HighlightOutlined />}>回评</Menu.Item> */
                    this.bindMenu(this.permissionList)
                  }

                </Menu>
              </Sider>
              <Content>
                <Switch>
                  {this.permissionList[0].userPermission === "阅卷员" ? <>
                    <Redirect from="/home" to="/home/mark-tasks" exact></Redirect>
                    <Route path="/home/mark-tasks" component={MarkTasks} exact></Route>
                    <Route path="/home/answer" component={Answer} exact></Route>
                    <Route path="/home/sample" component={Sample} exact></Route>
                    <Route path="/home/review" component={Review} exact></Route>
                    {/* <Route path="/home/selfMark" component={SelfMark} exact></Route> */}

                    <Route path="/home/markMonitor/all" component={all} exact></Route>
                    <Route path="/home/markMonitor/average" component={average} exact></Route>
                    <Route path="/home/markMonitor/score" component={score} exact></Route>
                    <Route path="/home/markMonitor/self" component={self} exact></Route>
                    <Route path="/home/markMonitor/standard" component={standard} exact></Route>
                    <Route path="/home/markMonitor/teacher" component={teacher} exact></Route>

                    <Route path="/home/group/arbitration" component={arbitration} exact></Route>
                    <Route path="/home/group/marking" component={marking}></Route>
                    <Route path="/home/group/problem" component={problem} exact></Route>
                    <Route path="/home/group/markTasks/:type/:QuestionId" component={markTasks}
                      exact></Route>

                    <Route path="/home/management/question" component={question} exact></Route>
                    <Route path="/home/management/paper" component={paper}></Route>
                    <Route path="/home/management/paper_allot" component={allot} exact></Route>
                    <Route path="/home/management/paper_manage" component={paperManage}
                      exact></Route>
                    <Route path="/home/management/detailTable" component={detail} exact></Route>
                    <Route path="/home/management/user/user_manage" component={user_manage} exact></Route>
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
