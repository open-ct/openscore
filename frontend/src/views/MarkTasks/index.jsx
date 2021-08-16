import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import Zmage from 'react-zmage'

import './index.less'

import * as Util from "../../util/Util";
import Marking from "../../api/marking";
const { Option } = Select;
export default class index extends Component {

  // state = { warningVisible: false };
  // showModal = () => {
  //   this.setState({
  //     warningVisible: true,
  //   });
  // };

  // hideModal = () => {
  //   this.setState({
  //     warningVisible: false,
  //   });
  // };
  userId = "1"
  state = {
    problemVisible: false,
    problemValue: 1,
    papers: [],
    currentPaper: {},
    currentPaperNum: 0,
    testLength: 0,
    selectId: [],
    selectScore: [],
    subTopic: []
  };

  componentDidMount() {
    this.getAllPaper();
  }
  // 总试卷获取 
  getAllPaper = () => {
    Marking.testList({ userId: this.userId })
      .then((res) => {
        if (res.data.status == "10000") {
          let papers = [...res.data.data.papers]
          this.setState(
            {
              papers ,
              testLength: res.data.data.papers.length
            }
          )
          this.getCurrentPaper();
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }
  // 当前试卷
  getCurrentPaper = () => {
    Marking.testDisplay({ userId: this.userId, testId: this.state.papers[0].Test_id.toString() })
      .then((res) => {
        if (res.data.status == "10000") {
          let currentPaper = res.data.data
          let subTopic =res.data.data.subTopic
          this.setState({
            currentPaper,
            subTopic
          })
        }
      })
      .catch((e) => {
        console.log(e)
      })
  }
  // 阅卷区
  showTest = () => {
    let testPaper = null;
    if (this.state.currentPaper.picSrcs != undefined) {
      testPaper = this.state.currentPaper.picSrcs.map((item) => {
        return <img src={item.Pic_src} alt="加载失败" className="test-question-img"/>
      })
    }

    return testPaper
  }
  render() {
    return (
      <DocumentTitle title="阅卷系统-评卷">
        <div className="mark-tasks-page" data-component="mark-tasks-page">
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
  // 评分区
  selectBox = (data) => {
    let selectArr = [];
    for (let i = 0; i <parseInt(data) +1; i++) {
      let selectOpt = (
        <Option key={i} value={i}>{i}</Option>
      )
      selectArr.push(selectOpt)
    }
    return selectArr
  }

  showSelect = () => {
    let scoreSelect = null;
    if (this.state.currentPaper.picSrcs != undefined) {
      scoreSelect = this.state.currentPaper.subTopic.map((item,index) => {
        return <div className="score-select">
          {item.Question_detail_name}：<Select key={index}  placeholder="请选择分数" style={{ width: 120 }} onSelect={this.select.bind(this,item.Question_detail_id)}>
            {
              this.selectBox(item.Question_detail_score)
            }
          </Select>
        </div>
      })
    }

    return scoreSelect
  }
  select = (item,value) => {
      this.setState({
        selectId :  [...this.state.selectId,item],
        selectScore :  [...this.state.selectScore,value]
      })   
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
    Modal.confirm({
      title: title,
      icon: <ExclamationCircleOutlined />,
      content: '提交后系统将记录改试卷',
      okText: '确认',
      cancelText: '取消',
      onOk: () => {
        let Qustion_detail_id = Util.getTextByJs(this.state.selectId);
        let Question_detail_score = Util.getTextByJs(this.state.selectScore);
        console.log()
        if (value == 1) {
          Marking.testPoint({ 
            userId: this.userId, 
            testId: this.state.papers[0].Test_id.toString(),
            scores: Question_detail_score, 
            testDetailId: Qustion_detail_id
          })
          .then((res) => {
              this.setState({
                selectId: [],
                selectScore: []
              })
              this.getAllPaper();           
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
  handleOk = () => {
    Marking.testProblem({ 
      userId: this.userId, 
      testId: this.state.papers[0].Test_id.toString(),
      problemType: this.state.problemValue.toString()
    })
    .then((res) => {
        this.setState({
          selectId: [],
          selectScore: []
        })      
        this.getAllPaper();  
        this.forceUpdate();
    })
    .catch((e) => {
      console.log(e)
    })
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
        <Input placeholder="请输入问题" style={{ marginTop: 10 }} />
      </Modal>
    )
  }
}
