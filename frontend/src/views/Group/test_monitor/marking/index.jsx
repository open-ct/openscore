import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";
import Marking from "../../../../api/marking";
import * as Util from "../../../../util/Util";
const { Option } = Select;
export default class index extends Component {
    userId = "1"
    supervisorId = "2"
    state = {
        problemVisible: false,
        problemValue: 1,
        inpValu: '',
        currentPaper: {},
        currentPaperNum: 0,
        testLength: 0,
        selectId: [],
        selectScore: [],
        subTopic: [],
        markScore: [],
        TestId:undefined
    }
    componentDidMount() {
        if (this.props.location.state) {
            this.setState({
                TestId: this.props.location.state.TestId
            })
            this.getCurrentPaper(this.props.location.state.TestId)
        }
    }

    handleOk = (value) => {
        if (this.state.problemValue == 2) {
            Marking.testProblem({
                userId: this.userId,
                testId: this.state.currentPaper.testId,
                problemType: this.state.problemValue,
                problemMessage :this.state.inpValu
            })
                .then((res) => {
                    this.setState({
                        selectId: [],
                        selectScore: [],
                    })
                    this.props.history.push('/home/markMonitor/self')
                })
                .catch((e) => {
                    console.log(e)
                })
        }else{
            Marking.testProblem({
                userId: this.userId,
                testId: this.state.currentPaper.testId,
                problemType: this.state.problemValue
            })
                .then((res) => {
                    this.setState({
                        selectId: [],
                        selectScore: [],
                    })
                    this.props.history.push('/home/markMonitor/self')
                })
                .catch((e) => {
                    console.log(e)
                })
        }

        this.setState({
            problemVisible: false,
        });


    };

    handleCancel = (value) => {
        this.setState({
            problemVisible: false,
        });
    };
    // 当前试卷
    getCurrentPaper = (test_id) => {
        Marking.testDisplay({ userId: '2', testId: parseInt(test_id) })
            .then((res) => {
                if (res.data.status == "10000") {
                    let currentPaper = res.data.data
                    let subTopic = res.data.data.subTopic
                    let markScore = []
                    for (let i = 0; i < subTopic.length; i++) {
                        markScore.push(subTopic[i].score_type.split('-'))
                    }
                    this.setState({
                        currentPaper,
                        subTopic,
                        markScore,
                    })
                }
            })
            .catch((e) => {
                console.log(e)
            })
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
    // 阅卷区
    showTest = () => {
        let testPaper = null;
        if (this.state.currentPaper.testInfos != undefined) {
            testPaper = this.state.currentPaper.testInfos.map((item) => {
                return <div className="test-question-img" data-content-before={this.imgScore(item.test_detail_id)}>
                    <img src={'data:image/jpg;base64,'+item.picCode} alt="加载失败" />
                </div>
            })
        }

        return testPaper
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
                        defaultValue={this.state.selectScore[index]}
                    >
                        {
                            this.selectBox(index)
                        }
                    </Select>
                </div>
            })
        }

        return scoreSelect
    }
    select = (item, value) => {
        console.log(item)
        if (this.state.selectId.length < 3) {
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
    renderScoreDropDown() {

        return (
            <div className="score-container">
                {
                    this.showSelect()
                }
            </div>
        );
    }
    renderPush() {

        return (
            <div className="push-container">
                <div>
                    <Button type="primary" style={{ width: 60 }} className="push-submit" onClick={() => {
                        this.showWarning(1)
                    }}>提交</Button>
                </div>
                <div className="push-paper" >
                    <Button className="push-problem" style={{ width: 60 }} onClick={() => {
                        this.showWarning(2)
                    }}>问题卷</Button>
                    <Button className="push-excellent" style={{ width: 60 }} onClick={() => {
                        this.showWarning(3)
                    }}>优秀卷</Button>
                </div>
                {
                    this.problemModal()
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
                    group.MonitorPoint({
                        supervisorId: '2',
                        testId: this.state.TestId,
                        scores: Question_detail_score,
                        testDetailId: Qustion_detail_id
                    })
                        .then((res) => {
                            this.setState({
                                selectId: [],
                                selectScore: [],
                                currentPaper: {},
                            })
                            message.success('提交成功');
                            this.props.history.push('/home/markMonitor/self')
                        })
                        .catch((e) => {
                            console.log(e)
                        })
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



    onRidioChange = e => {
        console.log('radio checked', e.target.value);
        this.setState({
            problemValue: e.target.value,
        });
    }
    handelChange(e) {
        this.setState({
          inpValu: e.target.value
        })
      }
    problemModal() {
        const { problemValue } = this.state;
        return (
            <Modal
                title="请选择问题卷类型"
                visible={this.state.problemVisible}
                onOk={() => this.handleOk(2)}
                onCancel={() => { this.handleCancel(2) }}
            >
                <Radio.Group onChange={this.onRidioChange} value={problemValue}>
                    <Space direction="vertical">
                        <Radio value={1}>图像不清</Radio>
                        <Radio value={2}>答错题目</Radio>
                        <Radio value={3}>其他错误</Radio>
                    </Space>
                </Radio.Group>
                <Input placeholder="请输入问题"  onChange={this.handelChange.bind(this)} style={{ marginTop: 10 }} />
            </Modal>
        )
    }

    render() {
        return (
            <DocumentTitle title="阅卷系统-自评卷">
                <div className="group-marking-page" data-component="group-marking-page">
                    <div className="mark-paper">
                        {
                            this.showTest()
                        }
                    </div>
                    <div className="mark-score">
                        {
                            this.renderScoreDropDown()
                        }
                        {
                            this.renderPush()
                        }
                    </div>
                </div>
            </DocumentTitle>
        )
    }




}
