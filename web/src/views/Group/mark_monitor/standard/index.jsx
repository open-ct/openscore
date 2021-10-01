import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";
const { Option } = Select;
export default class index extends Component {

    componentDidMount() {
        this.questionList();
    }

    // 选择区
    state = {
        questionList: [],
        tableData: [],
        fullScore: undefined
    }

    questionList = () => {
        group.questionList({ adminId: "1",subjectName: JSON.parse(localStorage.getItem('userInfo')).SubjectName})
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
    tableData = (questionId) => {
        group.standardMonitor({ supervisorId: "2", questionId: questionId })
            .then((res) => {
                if (res.data.status == "10000") {
                    let tableData = [];
                    for (let i = 0; i < res.data.data.ScoreDeviationVOList.length; i++) {
                        let item = res.data.data.ScoreDeviationVOList[i]
                        tableData.push({
                            UserName: item.UserName,
                            Deviation: item.DeviationScore,
                        })
                    }
                    this.setState({
                        tableData,
                        fullScore: res.data.data.fullScore
                    })
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    columns = [
        {
            title: '教师',
            width: 150,
            dataIndex: 'UserName',
        },
        {
            title: '标准差',
            width: 180,
            dataIndex: 'Deviation',
        }
    ]
    render() {
        return (
            <DocumentTitle title="阅卷系统-标准差监控">
                <div className="standard-monitor-page" data-component="standard-monitor-page">
                    <div className="search-container">
                        <div className="question-select">
                            题目选择：<Select
                                showSearch
                                style={{ width: 120 }}
                                optionFilterProp="label"
                                onSelect={(e) => { this.select(e) }}
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
                        <div className="question-score">
                            {/* 满分：{this.state.fullScore} */}
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
