import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";
import { Link } from 'react-router-dom'
import echarts from 'echarts/lib/echarts'
import ReactEcharts from 'echarts-for-react';
const { Option } = Select;
export default class index extends Component {

    supervisorId = '2'

    state = {
        questionList: [],
        teacherList: [],
        tableData: [],
        QuestionId: undefined,
        selfScoreRecordVOList:[]
    }
    columns = [
        {
            title: '序号（试卷编码）',
            width: 150,
            dataIndex: 'order',
        },
        {
            title: '上次评分',
            width: 150,
            dataIndex: 'last',
        },
        {
            title: '本次评分',
            width: 150,
            dataIndex: 'current',
        },
        {
            title: '操作',
            width: 150,
            dataIndex: 'operation',
            render: (text, record, index) => <Link to={{pathname: '/home/group/marking', state:{TestId: this.state.selfScoreRecordVOList[index].TestId}}}>自评</Link>
        },

    ]

    componentDidMount() {
        this.questionList();
    }
    
    getOption = ()=>{ 
        let X_data=[]
        let Y1_data=[]
        let Y2_data=[]
        for (let i = 0; i < this.state.tableData.length; i++) {
            X_data.push(this.state.tableData[i].order)
        }
        for (let i = 0; i < this.state.tableData.length; i++) {
            Y1_data.push(this.state.tableData[i].last)
        }
        for (let i = 0; i < this.state.tableData.length; i++) {
            Y2_data.push(this.state.tableData[i].current)
        }
        let option = {
            xAxis: {
                name: '序号',
                data: X_data
            },
            yAxis: {
                name: '分数'
            },
            series: [{
                name: '分数',
                type: 'bar',
                data: Y1_data
            },
            {
                name: '分数',
                type: 'bar',
                data: Y2_data
            }]
        };
        return option;
    };
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
                return <Option key={i} value={item.QuestionName} label={item.QuestionName}>{item.QuestionName}</Option>
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
                return <Option key={item.UserName} value={item.UserName} label={item.UserName}>{item.UserName}</Option>
            })
        } else {
            return null
        }
        return selectList
    }
    selectTeacher = (e) => {
        console.log(e)
        let index
        for (let i = 0; i < this.state.teacherList.length; i++) {
            if (this.state.teacherList[i].UserName === e) {
                index = i
            }
        }
        this.tableData(this.state.teacherList[index].UserId)
    }
    selectQuestion = (e) => {
        console.log(e)
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
                    let selfScoreRecordVOList = res.data.data.selfScoreRecordVOList
                    for (let i = 0; i < res.data.data.selfScoreRecordVOList.length; i++) {
                        let item = res.data.data.selfScoreRecordVOList[i]
                        tableData.push({
                            order : `${i+1}(${item.TestId})`,
                            last : item.Score,
                            current : item.SelfScore
                         })
                    }
                    this.setState({
                        tableData,selfScoreRecordVOList
                    })
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
                            columns={this.columns}
                            dataSource={this.state.tableData}
                        />
                    </div>
                    <ReactEcharts option={this.getOption()} style={{width:762,height:300}}/>
                </div>
            </DocumentTitle>
        )
    }




}
