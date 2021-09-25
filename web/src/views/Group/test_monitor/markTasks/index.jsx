import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined, ArrowLeftOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";
import Marking from "../../../../api/marking";
import * as Util from "../../../../util/Util";
const { Option } = Select;
export default class index extends Component {

    state = {
        type: '',
        paperButton: 1,
        count: 0,
        currentTestId: undefined,
        selectId: [],
        selectScore: [],
        problemVisible: false,
        problemValue: 1,
        inpValu: '',
        currentPaper: {},
        subTopic: [],
        testLength: 0,
        markScore: [],
        keyTest: [],
        sampleList: [],
        samplePaper: [],
        QuestionId: undefined,
    }

    componentDidMount() {
        console.log(this.props.match.params)
        this.setState({
            QuestionId: parseInt(this.props.match.params.QuestionId)
        })
        if (this.props.match.params.type === "1") {
            this.setState({
                type: '仲裁卷'
            })
            this.getArbitrationTestId();
        } else if (this.props.match.params.type === "2") {
            this.setState({
                type: '问题卷'
            })
            this.getProblemTestId()
        } else {
            this.setState({
                type: '自评卷'
            })
            this.getSelfTestId()
        }
    }

    // 仲裁卷
    getArbitrationTestId = () => {
        group.arbitrationTestId({ supervisorId: "2",questionId:parseInt(this.props.match.params.QuestionId) })
            .then((res) => {
                if (res.data.status == "10000") {
                    let currentTestId = res.data.data.arbitramentUnmarkListVOList[0].TestId
                    this.setState({
                        count: res.data.data.arbitramentUnmarkListVOList.length,
                        currentTestId
                    })
                    this.getCurrentPaper();
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }

    // 问题卷
    getProblemTestId = () => {
        group.problemTestId({ supervisorId: "2",questionId:parseInt(this.props.match.params.QuestionId) })
            .then((res) => {
                if (res.data.status == "10000") {
                    let currentTestId = res.data.data.ProblemUnmarkVOList[0].TestId
                    this.setState({
                        count: res.data.data.ProblemUnmarkVOList.length,
                        currentTestId
                    })
                    this.getCurrentPaper();
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }

    // 自评卷
    getSelfTestId = () => {
        group.selfTestId({ supervisorId: "2",questionId:parseInt(this.props.match.params.QuestionId) })
            .then((res) => {
                if (res.data.status == "10000") {
                    let currentTestId = res.data.data.selfUnmarkVOList[0].TestId
                    this.setState({
                        count: res.data.data.selfUnmarkVOList.length,
                        currentTestId
                    })
                    this.getCurrentPaper();
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    // 当前试卷
    getCurrentPaper = () => {
        Marking.testDisplay({ userId: '2', testId: this.state.currentTestId })
            .then((res) => {
                if (res.data.status == "10000") {
                    let currentPaper = res.data.data
                    let subTopic = res.data.data.subTopic
                    let testLength = res.data.data.subTopic.length
                    let markScore = []
                    for (let i = 0; i < subTopic.length; i++) {
                        markScore.push(subTopic[i].score_type.split('-'))
                    }
                    this.setState({
                        currentPaper,
                        subTopic,
                        markScore,
                        testLength
                    })
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    // 选择评分区
    showSelect = () => {
        let scoreSelect = null;
        if (this.state.currentPaper.testInfos != undefined) {
            scoreSelect = this.state.currentPaper.subTopic.map((item, index) => {
                return <div className="score-select">
                    {item.question_detail_name}：<Select
                        showSearch
                        key={index}
                        placeholder="请选择分数"
                        style={{ width: 120 }}
                        onSelect={this.select.bind(this, item.test_detail_id)}
                        optionFilterProp="label"
                        filterOption={(input, option) =>
                            option.label.indexOf(input) >= 0
                        }
                        filterSort={(optionA, optionB) =>
                            optionA.label.localeCompare(optionB.label)
                        }
                    >
                        {
                            this.selectBox(index)
                            // this.selectBox(item.question_detail_score)
                        }
                    </Select>
                </div>
            })
        }

        return scoreSelect
    }
    select = (item, value) => {
        console.log(item)
        if (this.state.selectId.length < this.state.testLength) {
            this.setState({
                selectId: [...this.state.selectId, item],
                selectScore: [...this.state.selectScore, value]
            })
        } else {
            let reviseSelectNo = 0;
            let newSelectScore = [];
            for (let i = 0; i < this.state.selectId.length; i++) {
                if (this.state.selectId[i] == item) {
                    reviseSelectNo = i
                }
            }
            this.state.selectScore.map((e, index) => {
                if (index == reviseSelectNo) {
                    newSelectScore.push(value)
                } else {
                    newSelectScore.push(e)
                }
            })
            this.setState({
                selectScore: newSelectScore
            })
        }
        console.log(this.state.selectId, this.state.selectScore)
    }
    selectBox = (index) => {
        let selectList
        if (this.state.markScore.length != 0) {
            selectList = this.state.markScore[index].map((item, i) => {
                return <Option key={i} value={item} label={item}>{item}</Option>
            })
        } else {
            return null
        }
        return selectList
    }

    renderScoreDropDown() {

        return (
            <div className="score-container">
                {
                    this.showSelect()
                }
            </div>
        );
    }

    // 提交区
    renderPush() {

        return (
            <div className="push-container">
                <div>
                    <Button type="primary" style={{ width: 60 }} className="push-submit" onClick={() => {
                        this.showWarning(1)
                    }}>提交</Button>
                </div>
                {
                    this.state.type !== "自评卷" ? <div className="push-paper" >
                        <Button className="push-problem" style={{ width: 60 }} onClick={() => {
                            this.showWarning(2)
                        }}>问题卷</Button>
                        <Button className="push-excellent" style={{ width: 60 }} onClick={() => {
                            this.showWarning(3)
                        }}>优秀卷</Button>
                        {
                            this.problemModal()
                        }
                    </div> : null
                }
            </div>
        );
    }
    showWarning = (value) => {
        let title = '';
        let okContent = '提交后系统将记录该试卷';
        switch (value) {
            case 1:
                title = '请确认是否提交该试卷';
                break;
            case 2:
                title = '请确认是否提交该问题卷';
                break;
            case 3:
                title = '请确认是否提交该优秀卷';
                break;
        }

        if (value == 1) {
            let selectOrdered = []
            for (let i = 0; i < this.state.subTopic.length; i++) {
                for (let j = 0; j < this.state.selectId.length; j++) {
                    if (this.state.selectId[j] == this.state.subTopic[i].test_detail_id) {
                        selectOrdered.push(this.state.selectScore[j])
                    }
                }
            }
            let content = '该试卷评分记录为：';
            for (let i = 0; i < selectOrdered.length; i++) {
                content += selectOrdered[i] + '\n'
            }
            okContent = content + '提交后系统将记录该试卷';
        }
        Modal.confirm({
            title: title,
            icon: <ExclamationCircleOutlined />,
            content: okContent,
            okText: '确认',
            cancelText: '取消',
            onOk: () => {
                let Qustion_detail_id = Util.getTextByJs(this.state.selectId);
                let Question_detail_score = Util.getTextByJs(this.state.selectScore);
                if (value == 1) {
                    if (this.state.selectScore.length < this.state.testLength) {
                        message.warning('请将分数打全')
                    } else {
                        group.MonitorPoint({
                            supervisorId: '2',
                            testId: this.state.currentTestId,
                            scores: Question_detail_score,
                            testDetailId: Qustion_detail_id
                        })
                            .then((res) => {
                                if (res.data.status == "10000") {
                                    this.setState({
                                        selectId: [],
                                        selectScore: [],
                                        currentPaper: {},
                                        count: undefined,
                                        currentTestId: undefined,
                                    })
                                    if (this.state.type === '仲裁卷') {
                                        this.getArbitrationTestId();
                                    } else if (this.state.type === '问题卷') {
                                        this.getProblemTestId()
                                    } else {
                                        this.getSelfTestId()
                                    }
                                }
                            })
                            .catch((e) => {
                                console.log(e)
                            })
                    }
                } else if (value == 3) {
                    console.log('3')
                } else {
                    console.log('2', this)
                    this.setState({
                        problemVisible: true,
                    });
                }
            }
        });
    }
    handleOk = () => {
        if (this.state.problemValue == 3) {
            Marking.testProblem({
                userId: this.userId,
                testId: this.state.currentTestId,
                problemType: this.state.problemValue,
                problemMessage: this.state.inpValu
            })
                .then((res) => {
                    this.setState({
                        selectId: [],
                        selectScore: [],
                        currentPaper: {},
                        count: undefined,
                        currentTestId: undefined,
                    })
                    if (this.state.type === '仲裁卷') {
                        this.getArbitrationTestId();
                    } else if (this.state.type === '问题卷') {

                        this.getProblemTestId()
                    } else {
                    }
                })
                .catch((e) => {
                    console.log(e)
                })
        } else {
            Marking.testProblem({
                userId: this.userId,
                testId: this.state.currentTestId,
                problemType: this.state.problemValue
            })
                .then((res) => {
                    this.setState({
                        selectId: [],
                        selectScore: [],
                        currentPaper: {},
                        count: undefined,
                        currentTestId: undefined,
                    })
                    if (this.state.type === '仲裁卷') {
                        this.getArbitrationTestId();
                    } else if (this.state.type === '问题卷') {

                        this.getProblemTestId()
                    } else {
                    }
                })
                .catch((e) => {
                    console.log(e)
                })
        }

        this.setState({
            problemVisible: false,
        });
    };

    handleCancel = () => {
        console.log('Clicked cancel button');
        this.setState({
            problemVisible: false,
        });
    };
    handelChange(e) {
        this.setState({
            inpValu: e.target.value
        })
    }
    onRidioChange = e => {
        console.log('radio checked', e.target.value);
        this.setState({
            problemValue: e.target.value,
        });
    }
    problemModal() {
        const { problemValue } = this.state;
        return (
            <Modal
                title="请选择问题卷类型"
                visible={this.state.problemVisible}
                onOk={this.handleOk}
                onCancel={this.handleCancel}
            >
                <Radio.Group onChange={this.onRidioChange} value={problemValue}>
                    <Space direction="vertical">
                        <Radio value={1}>图像不清</Radio>
                        <Radio value={2}>答错题目</Radio>
                        <Radio value={3}>其他错误</Radio>
                    </Space>
                </Radio.Group>
                <Input placeholder="请输入问题" onChange={this.handelChange.bind(this)} style={{ marginTop: 10 }} />
            </Modal>
        )
    }
    // 试卷区
    showTest = () => {
        let testPaper = null;
        if (this.state.paperButton === 1) {
            if (this.state.currentPaper.testInfos != undefined) {
                testPaper = this.state.currentPaper.testInfos.map((item) => {
                    return <div className="test-question-img" data-content-before={this.imgScore(item.test_detail_id)}>
                        <img src={'data:image/jpg;base64,' + item.picCode} alt="加载失败" />
                    </div>
                })
            }
        } else if (this.state.paperButton === 2) {
            if (this.state.keyTest != undefined || this.state.keyTest != null) {
                testPaper = this.state.keyTest.map((item) => {
                    return <div className="answer-question-img">
                        <img src={'data:image/jpg;base64,' + item} alt="加载失败" />
                    </div>
                })
            }
        } else if (this.state.paperButton === 3) {
            if (this.state.samplePaper !== undefined || this.state.samplePaper !== null) {
                testPaper = this.state.samplePaper.map((item) => {
                    return <img src={'data:image/jpg;base64,' + item.picCode} alt="加载失败" className="answer-question-img" />
                })
            }
        }
        return testPaper
    }
    // 打分展示
    imgScore = (item) => {
        let index
        for (let i = 0; i < this.state.selectId.length; i++) {
            if (item == this.state.selectId[i]) {
                index = i
            }
        }
        console.log(this.state.selectScore[index])
        return this.state.selectScore[index]
    }

    // 功能切换
    paperDisplay = (func) => {
        this.setState({
            paperButton: func
        })
        if (func === 1) {
            if (this.props.match.params.type === "1") {
                this.getArbitrationTestId();
            } else if (this.props.match.params.type === "2") {
                this.getProblemTestId()
            } else {
            }
        } else if (func === 2) {
            this.getAnswerPaper()
        } else if (func === 3) {
            this.getSampleList()
        }
    }
    // 答案
    getAnswerPaper = () => {
        Marking.testAnswer({ userId: '1', testId: 1 })
            .then((res) => {
                console.log(res)
                if (res.data.status == "10000") {
                    this.setState({
                        keyTest: res.data.data.keyTest
                    })
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    // 样卷
    getSampleList = () => {
        Marking.testExampleList({ userId: '1', testId: 1 })
            .then((res) => {
                if (res.data.status === "10000") {
                    let sampleList = []
                    for (let i = 0; i < res.data.data.exampleTestPapers.length; i++) {
                        sampleList.push({
                            order: i,
                            question_id: res.data.data.exampleTestPapers[i].test_id,
                            question_name: res.data.data.exampleTestPapers[i].candidate,
                            score: res.data.data.exampleTestPapers[i].final_score,
                        })
                    }
                    this.setState({
                        sampleList: sampleList
                    })
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }

    columns = [
        {
            title: '序号',
            width: 90,
            dataIndex: 'order',
        },
        {
            title: '题号',
            width: 90,
            dataIndex: 'question_id',
        },
        {
            title: '题名',
            width: 90,
            dataIndex: 'question_name',
        },
        {
            title: '分数',
            width: 90,
            dataIndex: 'score',
        },

    ];
    sampleTable() {
        return (
            <Table columns={this.columns}
                dataSource={this.state.sampleList}
                scroll={{ x: 400 }}
                pagination={{ position: ['bottomCenter'] }}
                onRow={(record) => ({
                    onClick: () => {
                        this.selectRow(record);
                    },
                })}
            />
        )
    }
    selectRow = (record) => {
        console.log(record.order)
        Marking.testDetail({ userId: '1', exampleTestId: record.question_id })
            .then((res) => {
                if (res.data.status == "10000") {
                    this.setState({
                        samplePaper: res.data.data.test[record.order]
                    })
                    console.log(this.state.samplePaper)
                }
            })
            .catch((e) => {
                console.log(e)
            })
    }
    render() {
        return (
            <DocumentTitle title={'试卷管理-' + this.state.type}>
                <div className="group-markTaks-page" data-component="group-markTaks-page">
                    <div className="search-container">
                        <div className="goBack" onClick={() => { window.history.back(-1) }}>
                            <ArrowLeftOutlined />
                            <span className="paperDetail">试卷详情</span>
                            <span className="goBackPre">返回{this.state.type}</span>
                        </div>
                        <div className="currentPaperNo">
                            试卷号：{this.state.currentTestId}
                        </div>
                        <div className="paperNum">
                            {this.state.type}数：{this.state.count}
                        </div>
                        <div className="buttonList">
                            <Button type={this.state.paperButton === 1 ? "primary" : "default"} style={{ width: 60 }} onClick={() => { this.paperDisplay(1) }}>评卷</Button>
                            <Button type={this.state.paperButton === 2 ? "primary" : "default"} style={{ width: 60 }} onClick={() => { this.paperDisplay(2) }}>答案</Button>
                            <Button type={this.state.paperButton === 3 ? "primary" : "default"} style={{ width: 60 }} onClick={() => { this.paperDisplay(3) }}>样卷</Button>
                            <Button type="default" style={{ width: 32 }}>…</Button>
                        </div>
                    </div>
                    <div className="paperPage">
                        <div className="mark-paper">
                            {
                                this.showTest()
                            }
                        </div>
                        <div className="mark-score">
                            {this.state.paperButton === 1 ?
                                this.renderScoreDropDown() : null
                            }
                            {this.state.paperButton === 1 ?
                                this.renderPush() : null
                            }
                            {this.state.paperButton === 3 ?
                                this.sampleTable() : null
                            }
                        </div>
                    </div>
                </div>
            </DocumentTitle>
        )
    }




}
