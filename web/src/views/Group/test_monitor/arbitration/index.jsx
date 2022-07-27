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
        count: 0,
        questionIndex:0
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
    columns = [
        {
            title: '试卷号',
            width: 150,
            dataIndex: 'TestId',
        },
        {
            title: '阅卷人一账号',
            width: 150,
            dataIndex: 'ExaminerFirstId',
        },
        {
            title: '阅卷人一名称',
            width: 150,
            dataIndex: 'ExaminerFirstName',
        },
        {
            title: '阅卷人一分数',
            width: 150,
            dataIndex: 'ExaminerFirstScore',
        },
        {
            title: '阅卷人二账号',
            width: 150,
            dataIndex: 'ExaminerSecondId',
        },
        {
            title: '阅卷人二名称',
            width: 150,
            dataIndex: 'ExaminerSecondName',
        },
        {
            title: '阅卷人二分数',
            width: 150,
            dataIndex: 'ExaminerSecondScore',
        },
        {
            title: '阅卷人三账号',
            width: 150,
            dataIndex: 'ExaminerThirdId',
        },
        {
            title: '阅卷人三名称',
            width: 150,
            dataIndex: 'ExaminerThirdName',
        },
        {
            title: '阅卷人三分数',
            width: 150,
            dataIndex: 'ExaminerThirdScore',
        },
        {
            title: '标准误差',
            width: 150,
            dataIndex: 'StandardError',
        },
        {
            title: '实际误差',
            width: 150,
            dataIndex: 'PracticeError',
        },

    ]

    tableData = (questionId) => {
        group.arbitramentList({ supervisorId: 2, questionId: questionId })
            .then((res) => {
                if (res.data.status == "10000") {
                    let tableData = [];
                    for (let i = 0; i < res.data.data.arbitramentTestVOList.length; i++) {
                        let item = res.data.data.arbitramentTestVOList[i]
                        tableData.push(item)
                    }
                    this.setState({
                        tableData,
                        count: res.data.data.count
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
        this.setState({
            questionIndex:index
        })
        this.tableData(this.state.questionList[index].QuestionId)
    }
    // 选择区

    paperMark =() => {
        this.props.history.push('/home/group/markTasks/1/'+this.state.questionList[this.state.questionIndex].QuestionId)
    }
    render() {
        return (
            <DocumentTitle title="阅卷系统-仲裁卷">
                <div className="group-arbitration-page" data-component="group-arbitration-page">
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
                        <div className="paper-num">
                            待仲裁数：{this.state.count}
                        </div>
                        <div className="paper-mark" onClick={() => {this.paperMark()}}>
                            评阅所有仲裁卷
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
