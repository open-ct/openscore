import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined, ArrowLeftOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";
import Manage from "../../../../api/manage";
import Marking from "../../../../api/marking";
import * as Util from "../../../../util/Util";
const { Option } = Select;
export default class index extends Component {

    adminId = "1"

    state = {
        subjectList: [],
        questionList: [],
        ImportTestNumber: undefined,
        OnlineNumber: undefined,
        testNumber: undefined,
        userNumber: undefined,
        subjectValue: undefined,
        questionValue: undefined,
        loading: false,
        ScoreType:undefined
    }

    componentDidMount() {
        this.subjectList()
    }

    subjectList = () => {
        Manage.subjectList({ adminId: this.adminId })
            .then((res) => {
                if (res.data.status === "10000") {
                    this.setState({
                        subjectList: res.data.data.subjectVOList
                    })
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    getSubjectOption = () => {
        let subjectOption
        if (this.state.subjectList.length) {
            subjectOption = this.state.subjectList.map(item => {
                return <Option key={item.SubjectId} value={item.SubjectName}>{item.SubjectName}</Option>
            })
        } else {
            return null
        }
        return subjectOption
    }
    subjectSelect = (e) => {
        this.setState({
            subjectValue: e,
            questionValue: undefined
        })
        Manage.questionList({ adminId: this.adminId, subjectName: e })
            .then((res) => {
                if (res.data.status === "10000") {
                    this.setState({
                        questionList: res.data.data.questionsList,
                        ImportTestNumber: undefined,
                        OnlineNumber: undefined,
                        ScoreType:undefined,
                        testNumber: undefined,
                        userNumber: undefined
                    })
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    getQuestionOption = () => {
        let questionOption
        if (this.state.questionList.length) {
            questionOption = this.state.questionList.map(item => {
                return <Option key={item.QuestionId} value={item.QuestionId}>{item.QuestionName}</Option>
            })
        } else {
            return null
        }
        return questionOption
    }
    questionSelect = (e) => {
        this.setState({
            questionValue: e
        })
        Manage.distributeInfo({ adminId: this.adminId, questionId: e })
            .then((res) => {
                if (res.data.status === "10000") {
                    this.setState({
                        ImportTestNumber: res.data.data.distributionInfoVO.ImportTestNumber,
                        OnlineNumber: res.data.data.distributionInfoVO.OnlineNumber,
                        ScoreType:res.data.data.distributionInfoVO.ScoreType,
                    })
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    distributePaper = () => {
        console.log(this.state.questionValue, this.state.testNumber, this.state.userNumber)
        if (this.state.questionValue != undefined && this.state.testNumber && this.state.userNumber) {
            Modal.confirm({
                title: '????????????',
                icon: <ExclamationCircleOutlined />,
                content: '',
                okText: '??????',
                cancelText: '??????',
                onOk: () => {
                    this.setState({
                        loading: true
                    })
                    Manage.distributePaper({ adminId: this.adminId, questionId: this.state.questionValue, testNumber: Number(this.state.testNumber), userNumber: Number(this.state.userNumber) })
                        .then((res) => {
                            if (res.data.status === "10000") {
                                this.setState({
                                    loading: false,
                                    questionList: [],
                                    ImportTestNumber: undefined,
                                    OnlineNumber: undefined,
                                    testNumber: undefined,
                                    userNumber: undefined,
                                    questionValue: undefined,
                                    ScoreType:undefined
                                })
                                message.success('?????????????????????')
                            }
                        })
                        .catch((e) => {
                            console.log(e)
                        })
                }
            });
        } else {
            message.warning('???????????????????????????????????????')
        }
    }
    goToDetail = () => {
        if (this.state.subjectValue) {
            this.props.history.push({pathname:"/home/management/detailTable",query:{subjectName:this.state.subjectValue}})
        }else {
            message.warning('?????????????????????')
        }   
    }
    ScoreType = () => {console.log('111111')
        if (this.state.ScoreType === 1) {
            return '???'
        } else if (this.state.ScoreType === 1) {
            return '???'
        }else {
            return null
        }
    }
    render() {
        return (
            <DocumentTitle title='????????????-????????????'>
                <div className="allot-page" data-component="allot-page">
                    <div className="subject-setting">
                        <div className="setting-header">????????????</div>
                        <div className="setting-box">
                            <div className="setting-input">
                                <div className="setting-item">
                                    ???????????????<Select
                                        style={{ width: 120 }}
                                        placeholder="???????????????"
                                        onSelect={(e) => { this.subjectSelect(e) }}
                                        value={this.state.subjectValue}
                                    >
                                        {
                                            this.getSubjectOption()
                                        }
                                    </Select>
                                </div>
                                <div className="setting-item">
                                    ???????????????<Select
                                        style={{ width: 120 }}
                                        placeholder="???????????????"
                                        onSelect={(e) => { this.questionSelect(e) }}
                                        value={this.state.questionValue}
                                    >
                                        {
                                            this.getQuestionOption()
                                        }
                                    </Select>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="question-setting">
                        <div className="setting-header">????????????</div>
                        <div className="setting-box">
                            <div className="setting-input">
                                <div className="setting-item">
                                    ????????????????????????{this.state.OnlineNumber}
                                </div>
                                <div className="setting-item">
                                    ?????????????????????{this.state.ImportTestNumber}
                                </div>
                                <div className="setting-item">
                                    ???????????????????????????{this.ScoreType()}
                                </div>
                            </div>
                            <div className="setting-input" style={{ marginTop: 24 }}>
                                <div className="setting-item">
                                    ?????????????????????<Input placeholder="??????????????????" value={this.state.userNumber} style={{ width: 120 }} onChange={e => {
                                        this.setState({
                                            userNumber: e.target.value
                                        })
                                    }} />
                                </div>
                                <div className="setting-item">
                                    ????????????????????????<Input placeholder="????????????????????????" value={this.state.testNumber} style={{ width: 120 }} onChange={e => {
                                        this.setState({
                                            testNumber: e.target.value
                                        })
                                    }} />
                                </div>
                            </div>
                        </div>
                    </div>
                    <Button type="primary" onClick={() => { this.distributePaper() }} loading={this.state.loading}>??????</Button>
                    <Button type="default" style={{ marginLeft: '20px' }} onClick={() => { this.goToDetail() }}>????????????</Button>
                </div>
            </DocumentTitle>
        )
    }

}