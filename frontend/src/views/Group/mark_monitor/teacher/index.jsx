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
        userInfo: {},
        questionList: [],
        tableData: []
    }
    
    columns = [
        {
            title: '老师',
            width: 120,
            dataIndex: 'UserName',
        },
        {
            title: '评卷数量',
            width: 120,
            dataIndex: 'TestDistributionNumber',
        },
        {
            title: '正评数量',
            width: 120,
            dataIndex: 'TestSuccessNumber',
        },
        {
            title: '无效数量',
            width: 120,
            dataIndex: 'TestProblemNumber',
        },
        {
            title: '未评数量',
            width: 120,
            dataIndex: 'TestRemainingNumber',
        },
        {
            title: '评卷速度（秒/份）',
            width: 180,
            dataIndex: 'MarkingSpeed',
        },
        {
            title: '平均分',
            width: 120,
            dataIndex: 'AverageScore',
        },
        {
            title: '有效度',
            width: 120,
            dataIndex: 'Validity',
        },
        {
            title: '标准差',
            width: 120,
            dataIndex: 'StandardDeviation',
        },
        {
            title: '自评指数',
            width: 120,
            dataIndex: 'EvaluationIndex',
        },
        {
            title: '在线情况',
            width: 120,
            dataIndex: 'OnlineTime',
        },

    ]

    userInfo = () => {
        group.userInfo({ supervisorId: this.supervisorId })
            .then((res) => {
                if (res.data.status == "10000") {
                    this.setState({
                        userInfo: res.data.userInfo
                    })
                    console.log(res.data.data.userInfo)
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }

    questionList = () => {
        group.questionList({ supervisorId: "2" })
            .then((res) => {
                if (res.data.status == "10000") {
                    this.setState({
                        questionList: res.data.data.questionsList
                    })
                    console.log(res.data.data.questionsList)
                    this.tableData(res.data.data.questionsList[0].QuestionId)
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    tableData = (questionId) => {
        group.teacherMonitor({ supervisorId: "2", questionId: questionId })
            .then((res) => {
                if (res.data.status == "10000") {
                    let tableData = [];
                    for (let i = 0; i < res.data.data.teacherMonitoringList.length; i++) {
                        let item = res.data.data.teacherMonitoringList[i]
                        tableData.push({
                            UserName: item.UserName,
                            TestDistributionNumber: item.TestDistributionNumber,
                            TestSuccessNumber: item.TestSuccessNumber,
                            TestRemainingNumber: item.TestRemainingNumber,
                            TestProblemNumber: item.TestProblemNumber,
                            MarkingSpeed: item.MarkingSpeed,
                            AverageScore: item.AverageScore,
                            Validity: item.Validity,
                            StandardDeviation: item.StandardDeviation,
                            EvaluationIndex: item.EvaluationIndex,
                            OnlineTime: item.OnlineTime = 0 ? "在线" : "离线"
                        })
                    }
                    this.setState({
                        tableData
                    })
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    componentDidMount() {
        this.userInfo();
        this.questionList();
    }


    // 题目选择区
    selectBox = () => {
        let selectList
        if (this.state.questionList.length != 0) {
            selectList = this.state.questionList.map((item, i) => {
                return <Option key={i} value={item.QuestionName} label={item.QuestionName}>{item.QuestionName}</Option>
            })
        } else {
            return null
        }
        return selectList
    }
    select = (e) => {
        let index
        for (let i = 0; i < this.state.questionList.length; i++) {
            if (this.state.questionList[i].QuestionName === e) {
                index = i
            }
        }
        this.tableData(this.state.questionList[index].QuestionId)
    }
    render() {
        return (
            <DocumentTitle title="阅卷系统-教师监控">
                <div className="teacher-monitor-page" data-component="teacher-monitor-page">
                    <div className="search-container">
                        <div className="question-select">
                            题目选择：<Select
                                showSearch
                                style={{ width: 120 }}
                                optionFilterProp="label"
                                onSelect={(e)=>{this.select(e)}}
                                filterOption={(input, option) =>
                                    option.label.indexOf(input) >= 0
                                }
                                filterSort={(optionA, optionB) =>
                                    optionA.label.localeCompare(optionB.label)
                                }
                                placeholder={this.state.questionList.length > 0 ? this.state.questionList[0].QuestionName : null}
                                defaultValue={this.state.questionList.length > 0 ? this.state.questionList[0].QuestionName : null}
                            >
                                {
                                    this.selectBox()
                                }

                            </Select>
                        </div>
                    </div>
                    <div className="display-container">
                        <Table
                            pagination={{ position: ['bottomCenter'] }}
                            columns={this.columns}
                            dataSource={this.state.tableData}
                        />
                    </div>
                </div>
            </DocumentTitle>
        )
    }




}
