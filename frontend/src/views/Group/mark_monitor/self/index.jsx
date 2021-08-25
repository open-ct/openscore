import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";
const { Option } = Select;
export default class index extends Component {

    supervisorId = '2'

    state = {
        questionList: [],
        teacherList: [],
        tableData: [],
        QuestionId: undefined
    }
    componentDidMount() {
        this.questionList();
    }
    questionList = () => {
        group.questionList({ supervisorId: "2" })
            .then((res) => {
                if (res.data.status == "10000") {
                    this.setState({
                        questionList: res.data.data.questionsList
                    })
                    this.teacherList(res.data.data.questionsList[0].QuestionId);
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    teacherList = (questionId) => {
        group.selfTeacher({ supervisorId: "2", questionId })
            .then((res) => {
                if (res.data.status == "10000") {
                    console.log(res.data)
                    this.setState({
                        teacherList: res.data.data.teacherVOList
                    })
                    this.tableData(res.data.data.teacherVOList[0].UserId)
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    // 题目选择区
    selectQuestionBox = () => {
        let selectList
        if (this.state.questionList.length != 0) {
            selectList = this.state.questionList.map((item, i) => {
                return <Option key={i} value={item.QuestionName} label={item.QuestionName}>{item.UserName}</Option>
            })
        } else {
            return null
        }
        return selectList
    }
    selectTeacherBox = () => {
        let selectList
        if (this.state.teacherList.length != 0) {
            selectList = this.state.teacherList.map((item, i) => {
                return <Option key={item.UserName} value={item.UserName} label={item.UserName}>{item.QuestionName}</Option>
            })
        } else {
            return null
        }
        return selectList
    }
    selectTeacher = (e) => {
        let index
        for (let i = 0; i < this.state.teacherList.length; i++) {
            if (this.state.teacherList[i].UserName === e) {
                index = i
            }
        }
        this.tableData(this.state.teacherList[index].UserId)
    }
    selectQuestion = (e) => {
        let index
        for (let i = 0; i < this.state.questionList.length; i++) {
            if (this.state.questionList[i].QuestionName === e) {
                index = i
            }
        }
        this.setState({
            QuestionId: this.state.questionList[index].QuestionId
        })
        this.teacherList(this.state.questionList[index].QuestionId)
    }
    tableData = (examinerId) => {
        group.selfMonitor({ supervisorId: "2", examinerId })
            .then((res) => {
                if (res.data.status == "10000") {
                    let tableData = [];
                    console.log(res.data)
                    // for (let i = 0; i < res.data.data.teacherMonitoringList.length; i++) {
                    //     let item = res.data.data.teacherMonitoringList[i]
                    //     tableData.push({
                    //         UserName: item.UserName,
                    //         TestDistributionNumber: item.TestDistributionNumber,
                    //         TestSuccessNumber: item.TestSuccessNumber,
                    //         TestRemainingNumber: item.TestRemainingNumber,
                    //         TestProblemNumber: item.TestProblemNumber,
                    //         MarkingSpeed: item.MarkingSpeed,
                    //         AverageScore: item.AverageScore,
                    //         Validity: item.Validity,
                    //         StandardDeviation: item.StandardDeviation,
                    //         EvaluationIndex: item.EvaluationIndex,
                    //         OnlineTime: item.OnlineTime = 0 ? "在线" : "离线"
                    //     })
                    // }
                    // this.setState({
                    //     tableData
                    // })
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }


    render() {
        return (
            <DocumentTitle title="阅卷系统-自评监控">
                <div className="self-monitor-page" data-component="self-monitor-page">
                    <div className="search-container">
                        <div className="question-select">
                            题目选择：<Select
                                key="question"
                                showSearch
                                style={{ width: 120 }}
                                optionFilterProp="label"
                                onSelect={(e) => { this.selectQuestion(e) }}
                                filterOption={(input, option) =>
                                    option.label.indexOf(input) >= 0
                                }
                                filterSort={(optionA, optionB) =>
                                    optionA.label.localeCompare(optionB.label)
                                }
                                placeholder={this.state.questionList.length > 0 ? this.state.questionList[0].QuestionName : null}
                                initialValue={this.state.questionList.length > 0 ? this.state.questionList[0].QuestionName : null}
                            >
                                {
                                    this.selectQuestionBox()
                                }
                            </Select>
                        </div>
                        <div className="teacher-select">
                            教师选择：<Select
                                key="teacher"
                                showSearch
                                style={{ width: 120 }}
                                optionFilterProp="label"
                                onSelect={(e) => { this.selectTeacher(e) }}
                                filterOption={(input, option) =>
                                    option.label.indexOf(input) >= 0
                                }
                                filterSort={(optionA, optionB) =>
                                    optionA.label.localeCompare(optionB.label)
                                }
                                placeholder={this.state.teacherList.length > 0 ? this.state.teacherList[0].UserName : null}
                                initialValue={this.state.teacherList.length > 0 ? this.state.teacherList[0].UserName : null}
                            >
                                {
                                    this.selectTeacherBox()
                                }
                            </Select>
                        </div>
                    </div>
                    <div className="display-container">
                        <Table
                            pagination={{ position: ['bottomCenter'] }}

                        />
                    </div>
                </div>
            </DocumentTitle>
        )
    }




}
