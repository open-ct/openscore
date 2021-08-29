import React, { Component } from 'react'
import { Layout, Menu } from 'antd';
import DocumentTitle from 'react-document-title'
import { Switch, Route, Redirect, history } from 'react-router-dom'
import { CheckCircleOutlined, FormOutlined, HighlightOutlined, ProfileOutlined } from '@ant-design/icons';
import * as Icon from '@ant-design/icons'
import { Avatar, Badge, Button, Col, Dropdown, Row } from "antd";
import { Link } from "react-router-dom";
import { BellOutlined, LogoutOutlined, SettingOutlined } from '@ant-design/icons';
import * as Auth from "../../auth/Auth";
import * as Setting from "../../Setting";
import * as Conf from "../../Conf";
import * as AccountBackend from "../../backend/AccountBackend";

import loadable from 'react-loadable';


import MarkTasks from '../Mark/MarkTasks'
import Answer from '../Mark/Answer'
import Sample from '../Mark/Sample'
import Review from '../Mark/Review'

import all from '../Group/mark_monitor/all'
import average from '../Group/mark_monitor/average'
import score from '../Group/mark_monitor/score'
import self from '../Group/mark_monitor/self'
import standard from '../Group/mark_monitor/standard'
import teacher from '../Group/mark_monitor/teacher'

import arbitration from '../Group/test_monitor/arbitration'
import marking from '../Group/test_monitor/marking'
import problem from '../Group/test_monitor/problem'
import markTasks from '../Group/test_monitor/markTasks'

import './index.less'
const { Header, Sider, Content } = Layout;
const { SubMenu } = Menu

export default class index extends Component {

    state = {
        account: null,
        current: "home"
    };
    permissionList = [
        {
            key: "mark-tasks",
            userPermission: "阅卷员",
            menu_name: "评卷",
            menu_url: "/home/mark-tasks",
            icon: "FormOutlined",
            chidPermissions: [
            ]
        },
        {
            key: "answer",
            userPermission: "阅卷员",
            menu_name: "答案",
            menu_url: "/home/answer",
            icon: "CheckCircleOutlined",
            menu_id: 1,
            chidPermissions: [
            ]
        },
        {
            key: "sample",
            userPermission: "阅卷员",
            menu_name: "样卷",
            menu_url: "/home/sample",
            icon: "ProfileOutlined",
            chidPermissions: [
            ]
        },
        {
            key: "review",
            userPermission: "阅卷员",
            menu_name: "回评",
            menu_url: "/home/review",
            icon: "HighlightOutlined",
            chidPermissions: [
            ]
        },
        {
            key: "mark_monitor",
            userPermission: "组长",
            menu_name: "评卷监控",
            menu_url: "/home/markMonitor",
            icon: "FormOutlined",
            chidPermissions: [
                {
                    key: "teacher",
                    userPermission: "组长",
                    menu_name: "教师监控",
                    menu_url: "/home/markMonitor/teacher",
                    icon: "",
                    chidPermissions: [
                    ]
                },
                {
                    key: "score",
                    userPermission: "组长",
                    menu_name: "分值分布",
                    menu_url: "/home/markMonitor/score",
                    icon: "",
                    chidPermissions: [
                    ]
                },
                {
                    key: "self",
                    userPermission: "组长",
                    menu_name: "自评监控",
                    menu_url: "/home/markMonitor/self",
                    icon: "",
                    chidPermissions: [
                    ]
                },
                {
                    key: "average",
                    userPermission: "组长",
                    menu_name: "平均分监控",
                    menu_url: "/home/markMonitor/average",
                    icon: "",
                    chidPermissions: [
                    ]
                },
                {
                    key: "standard",
                    userPermission: "组长",
                    menu_name: "标准差监控",
                    menu_url: "/home/markMonitor/standard",
                    icon: "",
                    chidPermissions: [
                    ]
                },
                {
                    key: "all",
                    userPermission: "组长",
                    menu_name: "总进度监控",
                    menu_url: "/home/markMonitor/all",
                    icon: "",
                    chidPermissions: [
                    ]
                },
            ]
        },
        {
            key: "test_management",
            userPermission: "组长",
            menu_name: "试卷管理",
            menu_url: "/home/testMonitor",
            icon: "TableOutlined",
            chidPermissions: [
                {
                    key: "arbitration",
                    userPermission: "组长",
                    menu_name: "仲裁卷",
                    menu_url: "/home/group/arbitration",
                    icon: "",
                    chidPermissions: [
                    ]
                },
                {
                    key: "problem",
                    userPermission: "组长",
                    menu_name: "问题卷",
                    menu_url: "/home/group/problem",
                    icon: "",
                    chidPermissions: [
                    ]
                },
                {
                    key: "marking",
                    userPermission: "组长",
                    menu_name: "自评卷",
                    menu_url: "/home/group/marking",
                    icon: "",
                    chidPermissions: [
                    ]
                },
            ]
        },
        {
            key: "user_management",
            userPermission: "组长",
            menu_name: "用户管理",
            menu_url: "/home/group/userMonitor",
            icon: "UserOutlined",
            chidPermissions: [
            ]
        },
    ]

    componentDidMount() {
        console.log(this.props)
        this.getAccount();
        setTimeout(() => {
            console.log(this.state)
        }, 5000);
        let moren = this.props.location.pathname;
        console.log(moren)
        this.setState(
            { current: moren.substring(moren.lastIndexOf("/") + 1, moren.length) }
        )
        
        this.props.history.listen((e) => {
            let test = e.pathname
            let text = test.substring(test.lastIndexOf("/") + 1, test.length)
            this.setState({
                current: text
            })
        })
    }

    getAccount() {
        AccountBackend.getAccount()
            .then((res) => {
                this.setState({
                    account: res.data,
                });
                localStorage.setItem("account", JSON.stringify(this.state.account))
            })
    }

    logout() {
        this.setState({
            expired: false,
            submitted: false,
        });

        AccountBackend.logout()
            .then((res) => {
                localStorage.setItem("account", "")
                if (res.status === 'ok') {
                    this.setState({
                        account: null
                    });

                    Setting.showMessage("success", `Successfully logged out, redirected to homepage`);

                    Setting.goToLink("/");
                } else {
                    Setting.showMessage("error", `Logout failed: ${res.msg}`);
                }
            });
    }

    handleRightDropdownClick(e) {
        if (e.key === '0') {
            Setting.openLink(Auth.getMyProfileUrl(this.state.account));
        } else if (e.key === '1') {
            this.logout();
        }
    }

    renderAvatar() {
        if (this.state.account.avatar === "") {
            return (
                <Avatar style={{ backgroundColor: Setting.getAvatarColor(this.state.account.name), verticalAlign: 'middle' }} size="large">
                    {Setting.getShortName(this.state.account.name)}
                </Avatar>
            )
        } else {
            return (
                <Avatar src={this.state.account.avatar} style={{ verticalAlign: 'middle' }} size="large">
                    {Setting.getShortName(this.state.account.name)}
                </Avatar>
            )
        }
    }

    renderRightDropdown() {
        const menu = (
            <Menu onClick={this.handleRightDropdownClick.bind(this)}>
                <Menu.Item key='0'>
                    <SettingOutlined />
                    My Account
                </Menu.Item>
                <Menu.Item key='1'>
                    <LogoutOutlined />
                    Logout
                </Menu.Item>
            </Menu>
        );

        return (
            <Dropdown key="200" overlay={menu} >
                <a className="ant-dropdown-link" href="#" style={{ float: 'right', marginLeft: '50px' }}>
                    {
                        this.renderAvatar()
                    }
                </a>
            </Dropdown>
        )
    }

    renderAccount() {
        if (this.state.account === undefined || this.state.account === null) {
            return (
                <a href={Auth.getAuthorizeUrl()} style={{ color: '#ffffff', marginLeft: '50px' }}>
                    登录
                </a>
            );
        } else {
            return (
                this.renderRightDropdown()
            )
        }
    }


    bindMenu = (menulist) => {
        let MenuList = menulist.map((item) => {
            if (item.chidPermissions.length === 0) {  //没有子菜单
                return <Menu.Item key={item.key} icon={item.icon ? React.createElement(Icon[item.icon]) : null}  ><Link to={item.menu_url}>{item.menu_name}</Link></Menu.Item>
            }
            else {
                return <SubMenu key={item.key} title={item.menu_name} icon={React.createElement(Icon[item.icon])}>
                    {this.bindMenu(item.chidPermissions)}
                </SubMenu>
            }

        })
        return MenuList
    }

    bindRouter = (menulist) => {
        let routerList = menulist.map((item) => {
            if (item.chidPermissions.length === 0) {
                return <Route key={item.key} path={item.menu_url} component={MarkTasks}></Route>
            }
        })
        return routerList
    }


    render() {
        return (
            <DocumentTitle title="阅卷系统">
                <Layout className="home-page" data-component="home-page">
                    <Header>
                        <div className="header-box">
                            <span className="header-title">OpenCT在线阅卷系统</span>
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
                                style={{ width: 200, height: '100%' }}
                                defaultSelectedKeys={['mark-tasks']}
                                mode="inline"
                                selectedKeys={[this.state.current]}
                                onClick={this.funcChange}
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

                                    <Route path="/home/markMonitor/all" component={all} exact></Route>
                                    <Route path="/home/markMonitor/average" component={average} exact></Route>
                                    <Route path="/home/markMonitor/score" component={score} exact></Route>
                                    <Route path="/home/markMonitor/self" component={self} exact></Route>
                                    <Route path="/home/markMonitor/standard" component={standard} exact></Route>
                                    <Route path="/home/markMonitor/teacher" component={teacher} exact></Route>

                                    <Route path="/home/group/arbitration" component={arbitration} exact></Route>
                                    <Route path="/home/group/marking/:TestId" component={marking}></Route>
                                    <Route path="/home/group/problem" component={problem} exact></Route>
                                    <Route path="/home/group/markTasks/:type" component={markTasks} exact></Route>                              
                                </>

                                    : null
                                }
                            </Switch>
                        </Content>
                    </Layout>
                </Layout>
            </DocumentTitle>
        )
    }
}
