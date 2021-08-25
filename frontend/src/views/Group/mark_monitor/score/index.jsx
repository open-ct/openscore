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
        tableData: [],
        columns : [
            {
                title: '分数',
                width: 120,
                dataIndex: 'Score',
            },
        ],
        tableData:[
            {Score: '教师'}
        ]
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
                    console.log(res.data.data.questionsList)
                    this.tableData(res.data.data.questionsList[0].QuestionId)
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }

    tableData = (questionId) => {
        group.scoreMonitor({ supervisorId: "2", questionId: questionId })
            .then((res) => {
                if (res.data.status == "10000") {
                    console.log(res.data)
                    let columns = [{
                        title: '分数',
                        width: 120,
                        dataIndex: 'Score',
                    }]
                    let tableData = [{
                        Score: '教师'
                    }];

 
                    for (let i = 0; i < res.data.data.scoreDistributionList.length; i++) {
                        let item = res.data.data.scoreDistributionList[i]
                        tableData[0]["score_"+item.Score] = item.Rate
                        columns.push(
                            {
                                title: item.Score + 1,
                                width: 180,
                                dataIndex: "score_"+item.Score 
                            }
                        )
                    }
                    console.log(tableData,columns)
                    this.setState({
                        tableData,columns
                    })
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

    render() {
        return (
            <DocumentTitle title="阅卷系统-分值分布">
                <div className="score-monitor-page" data-component="score-monitor-page">
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
                    </div>
                    <div className="display-container">
                        <Table
                            
                            pagination={{ position: ['bottomCenter'] }}
                            columns={this.state.columns}
                            dataSource={this.state.tableData}
                        />
                    </div>
                </div>
            </DocumentTitle>
        )
    }




}
