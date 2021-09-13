import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined, ArrowLeftOutlined, DeleteOutlined } from '@ant-design/icons';
import './index.less';
import group from "../../../../api/group";
import Marking from "../../../../api/marking";
import Manage from "../../../../api/manage";
import * as Util from "../../../../util/Util";
const { Option } = Select;
export default class index extends Component {

    // constructor(props) {
    //     super(props);
    //     // create a ref to store the textInput DOM element
    //     this.myInput_name = React.createRef();
    //     this.myInput_score = React.createRef();
    //     this.myInput_type = React.createRef();
    // }
    adminId = "1"
    state = {
        questionList: [

        ],
        questionNo: 0,
        subjectName: "语文",
        topicName: undefined,
        score: undefined,
        error: undefined,
        scoreType: 0,
        topicDetails: [
            {
                topicDetailName: undefined,
                detailScore: undefined,
                DetailScoreTypes: undefined
            }
        ],
        loading: false,
        QuestionId: undefined,
        QuestionDetailIds: [],
        visible: false
    }

    componentDidMount() {
        let questionList = [...this.state.questionList]
        questionList.push(
            <div className="question-input" key={0}>
                <div className="question-item">
                    小题名：<Input placeholder="请输入小题名" style={{ width: 120 }} onChange={e => {
                        let topicDetails = [...this.state.topicDetails]
                        topicDetails[0].topicDetailName = e.target.value
                        this.setState({
                            topicDetails
                        })

                    }} />
                </div>
                <div className="question-item">
                    小题满分：<Input placeholder="请输入分数" style={{ width: 120 }} onChange={e => {
                        let topicDetails = [...this.state.topicDetails]
                        topicDetails[0].detailScore = Number(e.target.value)
                        this.setState({
                            topicDetails
                        })

                    }} />
                </div>
                <div className="question-item">
                    分数分布：<Input placeholder="请输入分布" style={{ width: 120 }} onChange={e => {
                        let topicDetails = [...this.state.topicDetails]
                        topicDetails[0].DetailScoreTypes = e.target.value
                        this.setState({
                            topicDetails
                        })

                    }} />
                </div>
            </div>
        )
        this.setState({
            questionList
        })
    }

    questionAdd = () => {
        let questionList = [...this.state.questionList]
        let questionNo = this.state.questionNo + 1
        let questionDetail = {
            topicDetailName: undefined,
            detailScore: undefined,
            DetailScoreTypes: undefined
        }
        this.setState({
            topicDetails: [...this.state.topicDetails, questionDetail]
        })
        questionList.push(
            <div className="question-input" key={questionNo}>
                <div className="question-item">
                    小题名：<Input placeholder="请输入小题名" style={{ width: 120 }} onChange={e => {
                        let topicDetails = [...this.state.topicDetails]
                        let addNo
                        for (let i = 0; i < this.state.questionList.length; i++) {
                            if (this.state.questionList[i].key === questionNo.toString()) {
                                addNo = i
                            }
                        }
                        topicDetails[addNo]['topicDetailName'] = e.target.value
                        console.log(topicDetails)
                        this.setState({
                            topicDetails
                        })
                    }} />
                </div>
                <div className="question-item">
                    小题满分：<Input placeholder="请输入分数" style={{ width: 120 }} onChange={e => {
                        let topicDetails = [...this.state.topicDetails]
                        let addNo
                        for (let i = 0; i < this.state.questionList.length; i++) {
                            if (this.state.questionList[i].key === questionNo.toString()) {
                                addNo = i
                            }
                        }
                        topicDetails[addNo]['detailScore'] = Number(e.target.value)
                        this.setState({
                            topicDetails
                        })
                    }} />
                </div>
                <div className="question-item">
                    分数分布：<Input placeholder="请输入分布" style={{ width: 120 }} onChange={e => {
                        let topicDetails = [...this.state.topicDetails]
                        let addNo
                        for (let i = 0; i < this.state.questionList.length; i++) {
                            if (this.state.questionList[i].key === questionNo.toString()) {
                                addNo = i
                            }
                        }
                        topicDetails[addNo]['DetailScoreTypes'] = e.target.value
                        this.setState({
                            topicDetails
                        })
                    }} />
                </div>
                <div className="question-item">
                    <DeleteOutlined style={{ fontSize: '20px', color: '#1890FF', height: '100%', marginTop: '5px', cursor: 'pointer' }} onClick={() => { this.questionDel(questionNo) }} />
                </div>
            </div>
        )
        this.setState({
            questionList,
            questionNo
        })
    }
    questionDel = (questionNo) => {
        let deleteNo
        let topicDetails = [...this.state.topicDetails]
        for (let i = 0; i < this.state.questionList.length; i++) {
            console.log(this.state.questionList[i].key)
            if (this.state.questionList[i].key === questionNo.toString()) {
                deleteNo = i
                topicDetails.splice(i,1)
            }
        }
        console.log(topicDetails)
        this.setState({
            questionList: this.state.questionList.filter((item, i) => i !== deleteNo),
            topicDetails
        })
    }
    confirmQuestionId = () => {
        let question_flag = true;
        let detail_flag = true;
        if (this.state.subjectName === undefined || this.state.topicName === undefined || this.state.scoreType === undefined || this.state.score === undefined || this.state.error === undefined) {
            question_flag = false
        }
        if (!question_flag) {
            message.warning('请将题目设置填写完整')
        }
        let topicDetails = this.state.topicDetails.filter(item => Object.keys(item).length !== 0)
        for (let i = 0; i < topicDetails.length; i++) {
            let item = topicDetails[i];
            if (item.topicDetailName === undefined || item.detailScore === undefined || item.DetailScoreTypes === undefined) {
                detail_flag = false;
            }
        }
        if (!detail_flag) {
            message.warning('请将小题设置填写完整')
        }
        if (question_flag && detail_flag) {
            Modal.confirm({
                title: '确认提交题目设置',
                icon: <ExclamationCircleOutlined />,
                content: '',
                okText: '确认',
                cancelText: '取消',
                onOk: () => {
                    this.setState({
                        loading: true
                    })
                    Manage.questionImport({
                        adminId: this.adminId,
                        topicName: this.state.topicName,
                        subjectName: this.state.subjectName,
                        score: Number(this.state.score),
                        error: Number(this.state.error),
                        scoreType: this.state.scoreType,
                        topicDetails
                    })
                        .then((res) => {
                            if (res.data.status === "10000") {
                                this.setState({
                                    visible: true,
                                    QuestionId: res.data.data.addTopicVO.QuestionId,
                                    QuestionDetailIds: res.data.data.addTopicVO.QuestionDetailIds,
                                })
                            }
                        })
                        .catch((e) => {
                            console.log(e)
                        })
                }
            });
        }
    }
    hideModal = () => {
        this.setState({
            questionNo: -1,
            subjectName: "语文",
            topicName: undefined,
            score: undefined,
            error: undefined,
            scoreType: 0,
            loading: false,
            QuestionId: undefined,
            QuestionDetailIds: [],
            visible: false,
        })
        this.setState({
            topicDetails: [
            ],
            questionList: [
            ]
        }
            // this.myInput_name.current.state.value = undefined;
            // this.myInput_score.current.state.value = undefined;
            // this.myInput_type.current.state.value = undefined;
        )
        message.success("题目设置成功！")
    }
    cancelModal = () => {
        this.setState({
            visible: false
        })
    }
    questionIdModal = () => {

        return (
            <>
                <Modal
                    title="请确认"
                    visible={this.state.visible}
                    onOk={this.hideModal}
                    onCancel={this.cancelModal}
                    okText="确认"
                    cancelText="取消"
                    maskClosable={false}
                >
                    <p>大题ID号：{this.state.QuestionId}</p>
                    {
                        this.state.QuestionDetailIds.map((item, i) => {
                            return <p key={i}>小题ID号：{item.QuestionDetailId}</p>
                        })
                    }
                </Modal>
            </>
        );
    }

    goToDetail = () => {
        console.log('123')
        this.props.history.push("/home/management/detailTable")
    }

    render() {
        return (
            <DocumentTitle title='试卷管理-题目设置' >
                <div className="question-page" data-component="question-page">
                    <div className="subject-setting">
                        <div className="setting-header">题目设置</div>
                        <div className="setting-box">
                            <div className="setting-input">
                                <div className="subject-item">
                                    科目选择：    <Select
                                        defaultValue={this.state.subjectName}
                                        style={{ width: 120 }}
                                        onSelect={(e) => {
                                            this.setState({
                                                subjectName: e
                                            })
                                        }}>
                                        <Option value="语文">语文</Option>
                                        <Option value="数学">数学</Option>
                                        <Option value="英语">英语</Option>
                                    </Select>
                                </div>
                                <div className="subject-item">
                                    大题名：<Input placeholder="请输入大题名" value={this.state.topicName} style={{ width: 120 }} onChange={e => {
                                        this.setState({
                                            topicName: e.target.value
                                        })
                                    }} />
                                </div>
                                <div className="subject-item">
                                    分数满分：<Input placeholder="请输入分数" value={this.state.score} style={{ width: 120 }} onChange={e => {
                                        this.setState({
                                            score: e.target.value
                                        })
                                    }} />
                                </div>
                                <div className="subject-item">
                                    标准误差：<Input placeholder="请输入误差" value={this.state.error} style={{ width: 120 }} onChange={e => {
                                        this.setState({
                                            error: e.target.value
                                        })
                                    }} />
                                </div>
                            </div>
                            <div className="setting-second">
                                是否需要二次阅卷：<Radio.Group
                                    onChange={e => {
                                        this.setState({
                                            scoreType: e.target.value
                                        })
                                    }}
                                    defaultValue={this.state.scoreType}>
                                    <Radio value={0}>否</Radio>
                                    <Radio value={1}>是</Radio>
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
                    {
                        this.questionIdModal()
                    }
                    <Button type="primary" onClick={() => { this.confirmQuestionId() }} loading={this.state.loading}>确认</Button>
                    <Button type="default" style={{ marginLeft: '20px' }} onClick={() => { this.goToDetail() }}>查看详情</Button>
                    <Button onClick={()=>{console.log(this.state.topicDetails,this.state.questionList)}}>...</Button>
                </div>
            </DocumentTitle>
        )
    }

}
