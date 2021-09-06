import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined, ArrowLeftOutlined, DeleteOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";
import Marking from "../../../../api/marking";
import * as Util from "../../../../util/Util";
const { Option } = Select;
export default class index extends Component {

    state = {
        questionList: [],
        questionNo: 0

    }
    componentDidMount() {
        this.setState({
            questionList: [
                <div className="question-input">
                    <div className="question-item">
                        小题名：<Input placeholder="请输入小题名" style={{ width: 120 }} />
                    </div>
                    <div className="question-item">
                        小题满分：<Input placeholder="请输入分数" style={{ width: 120 }} />
                    </div>
                    <div className="question-item">
                        分数分布：<Input placeholder="请输入分布" style={{ width: 120 }} />
                    </div>
                </div>
            ]
        })
    }

    questionAdd = () => {
        let questionList = [...this.state.questionList]
        let questionNo = this.state.questionNo
        questionList.push(
            <div className="question-input" key={questionNo}>
                <div className="question-item">
                    小题名：<Input placeholder="请输入小题名" style={{ width: 120 }} />
                </div>
                <div className="question-item">
                    小题满分：<Input placeholder="请输入分数" style={{ width: 120 }} />
                </div>
                <div className="question-item">
                    分数分布：<Input placeholder="请输入分布" style={{ width: 120 }} />
                </div>
                <div className="question-item">
                    <DeleteOutlined style={{ fontSize: '20px', color: '#1890FF', height: '100%', marginTop: '5px', cursor: 'pointer' }} onClick={() => { this.questionDel(questionNo) }} />
                </div>
            </div>
        )
        console.log(questionList)
        this.setState({
            questionList,
            questionNo: this.state.questionNo + 1
        })
    }
    questionDel = (questionNo) => {
        let deleteNo
        for (let i = 1; i < this.state.questionList.length; i++) {
            console.log(this.state.questionList[i].key)
            if (this.state.questionList[i].key === questionNo.toString()) {
                deleteNo = i
            }
        }
        this.setState({
            questionList: this.state.questionList.filter((item, i) => i !== deleteNo)
        })
    }
    render() {
            return(
            <DocumentTitle title = '试卷管理-导入题目' >
                    <div className="question-page" data-component="question-page">
                        <div className="subject-setting">
                            <div className="setting-header">题目设置</div>
                            <div className="setting-box">
                                <div className="setting-input">
                                    <div className="subject-item">
                                        科目选择：    <Select defaultValue="语文" style={{ width: 120 }} onSelect={(e) => { this.select(e) }}>
                                            <Option value="jack">语文</Option>
                                            <Option value="lucy">数学</Option>
                                            <Option value="Yiminghe">英语</Option>
                                        </Select>
                                    </div>
                                    <div className="subject-item">
                                        大题名：<Input placeholder="请输入大题名" style={{ width: 120 }} />
                                    </div>
                                    <div className="subject-item">
                                        分数满分：<Input placeholder="请输入分数" style={{ width: 120 }} />
                                    </div>
                                    <div className="subject-item">
                                        标准误差：<Input placeholder="请输入误差" style={{ width: 120 }} />
                                    </div>
                                </div>
                                <div className="setting-second">
                                    是否需要二次阅卷：<Radio.Group onChange={e => { this.onChange(e) }}>
                                        <Radio value={1}>是</Radio>
                                        <Radio value={2}>否</Radio>
                                    </Radio.Group>
                                </div>
                            </div>
                        </div>
                        <div className="question-setting">
                            <div className="setting-header">小题设置</div>
                            <div className="setting-box">
                                {this.state.questionList}
                                <div className="question-add" onClick={() => this.questionAdd(this.state.questionNo)}>+</div>
                            </div>
                        </div>
                    </div>
            </DocumentTitle>
        )
    }

}
