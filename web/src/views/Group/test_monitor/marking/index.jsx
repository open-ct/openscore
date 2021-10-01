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
            title: '用户Id',
            width: 150,
            dataIndex: 'Userid',
        },
        {
            title: '用户名',
            width: 150,
            dataIndex: 'Name',
        },
        {
            title: '试卷号',
            width: 150,
            dataIndex: 'TestId',
        },
        {
            title: '上次评分',
            width: 150,
            dataIndex: 'Score',
        },
        {
            title: '本次评分',
            width: 150,
            dataIndex: 'SelfScore',
        },

    ]

    tableData = (questionId) => {
        group.selfMarkList({ supervisorId: "2", questionId: questionId })
            .then((res) => {
                if (res.data.status == "10000") {
                    let tableData = [];
                    for (let i = 0; i < res.data.data.selfMarkVOList.length; i++) {
                        let item = res.data.data.selfMarkVOList[i]
                        tableData.push(item)
                    }
                    this.setState({
                        tableData,
                        count: res.data.data.selfMarkVOList.length
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
        this.props.history.push('/home/group/markTasks/3/'+this.state.questionList[this.state.questionIndex].QuestionId)
    }
    render() {
        return (
            <DocumentTitle title="阅卷系统-自评卷">
                <div className="group-marking-page" data-component="group-marking-page">
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
                            待自评数：{this.state.count}
                        </div>
                        <div className="paper-mark" onClick={() => {this.paperMark()}}>
                            评阅所有自评卷
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
