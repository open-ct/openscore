import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Table, Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input } from 'antd';
import { ExclamationCircleOutlined, } from '@ant-design/icons';
import './index.less'
import Marking from "../../../api/marking";
import * as Util from "../../../util/Util";
const { Option } = Select;
export default class index extends Component {
  
  userId = "1"

  state = {
    reviewVisible: true,
    reviewList: [],
    problemVisible: false,
    inpValu: '',
    problemValue: 1,
    currentPaper: {},
    currentPaperNum: 0,
    testLength: 0,
    selectId: [],
    selectScore: [],
    subTopic: [],
    markScore: [],
  }

  componentDidMount() {
    this.getReviewList();
  }

  getReviewList = () => {
    Marking.testReview({ userId: this.userId })
      .then((res) => {
        if (res.data.status == "10000") {
          let reviewList = []
          for (let i = 0; i < res.data.data.testId.length; i++) {
            reviewList.push({
              order: i + 1,
              test_id: res.data.data.testId[i],
              score: res.data.data.score[i]
            })
          }
          this.setState({
            reviewList
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
      dataIndex: 'order',
    },
    {
      title: '试卷号',
      dataIndex: 'test_id',
    },
    {
      title: '分数',
      dataIndex: 'score',
    },
  ];

  render() {

    return (
      <DocumentTitle title="阅卷系统-自评">
        <div className="selfMark-tasks-page" data-component="selfMark-tasks-page">
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
  handleOk = (value) => {
    if (value == 1) {
      this.setState({
        reviewVisible: false,
      });
    } else {
      if (this.state.problemValue == 3) {
        Marking.testProblem({
          userId: this.userId,
          testId: this.state.currentPaper.testId,
          problemType: this.state.problemValue,
          problemMessage: this.state.inpValu
        })
          .then((res) => {
            this.setState({
              selectId: [],
              selectScore: [],
              reviewVisible: true
            })
            this.getReviewList();
          })
          .catch((e) => {
            console.log(e)
          })
      } else {
        Marking.testProblem({
          userId: this.userId,
          testId: this.state.currentPaper.testId,
          problemType: this.state.problemValue
        })
          .then((res) => {
            this.setState({
              selectId: [],
              selectScore: [],
              reviewVisible: true
            })
            this.getReviewList();
          })
          .catch((e) => {
            console.log(e)
          })
      }
      this.setState({
        problemVisible: false,
      });
    }

  };

  handleCancel = (value) => {
    if (value == 1) {
      this.setState({
        reviewVisible: false,
      });
    } else {
      this.setState({
        problemVisible: false,
      });
    }
  };
  reviewModal() {
    return (
      <Modal
        title="回评列表"
        visible={this.state.reviewVisible}
        onOk={() => this.handleOk(1)}
        onCancel={() => { this.handleCancel(1) }}
        width={800}
        maskClosable={false}
      >
        {
          this.reviewTable()
        }
      </Modal>
    )
  }
  reviewTable() {
    return (
      <Table
        columns={this.columns}
        pagination={false}
        dataSource={this.state.reviewList}
        onRow={(record) => ({
          onClick: () => {
            this.selectRow(record);
          },
        })}
      />
    )
  }
  selectRow = (record) => {
    this.getCurrentPaper(record.test_id)
    this.setState({
      reviewVisible: false
    })
  }
  // 当前试卷
  getCurrentPaper = (test_id) => {
    Marking.testDisplay({ userId: this.userId, testId: test_id })
      .then((res) => {
        if (res.data.status == "10000") {
          let currentPaper = res.data.data
          let subTopic = res.data.data.subTopic
          let testInfos = res.data.data.testInfos
          let testLength = res.data.data.subTopic.length
          let markScore = []
          let selectId = []
          let selectScore = []
          for (let i = 0; i < subTopic.length; i++) {
            markScore.push(subTopic[i].score_type.split('-'))
            selectId.push(testInfos[i].test_detail_id)
            // 判断用户改卷
            let userScore
            if (this.userId == testInfos[i].examiner_first_id) {
              userScore = testInfos[i].examiner_first_score
            } else if (this.userId == testInfos[i].examiner_second_id) {
              userScore = testInfos[i].examiner_second_score
            } else if (this.userId == testInfos[i].examiner_second_id) {
              userScore = testInfos[i].examiner_third_score
            }
            selectScore.push(userScore)
          }
          this.setState({
            currentPaper,
            subTopic,
            markScore,
            selectId,
            selectScore,
            testLength
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
          <img src={'data:image/jpg;base64,' + item.picCode} alt="加载失败" />
        </div>
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
            {
              this.reviewModal()
            }
          </div>
        </div>
      </DocumentTitle>
    )
  }
  // 评分区

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
          }}>自评</Button>
        </div>
        {/* <div className="push-paper" >
          <Button className="push-problem" style={{ width: 60 }} onClick={() => {
            this.showWarning(2)
          }}>问题卷</Button>
          <Button className="push-excellent" style={{ width: 60 }} onClick={() => {
            this.showWarning(3)
          }}>优秀卷</Button>
        </div> */}
        {
        //   this.problemModal()
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
          Marking.testReviewPoint({
            userId: this.userId,
            testId: this.state.currentPaper.testId,
            scores: Question_detail_score,
            testDetailId: Qustion_detail_id
          })
            .then((res) => {
              if (res.data.status == "10000") {
                this.setState({
                  selectId: [],
                  selectScore: [],
                  currentPaper: {},
                  reviewVisible: true
                })
                message.success('回评成功');
                this.props.history.push('/home/mark-tasks')
              }
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

//   handelChange(e) {
//     this.setState({
//       inpValu: e.target.value
//     })
//   }
//   onRidioChange = e => {
//     console.log('radio checked', e.target.value);
//     this.setState({
//       problemValue: e.target.value,
//     });
//   }
//   problemModal() {
//     const { problemValue } = this.state;
//     return (
//       <Modal
//         title="请选择问题卷类型"
//         visible={this.state.problemVisible}
//         onOk={() => this.handleOk(2)}
//         onCancel={() => { this.handleCancel(2) }}
//       >
//         <Radio.Group onChange={this.onRidioChange} value={problemValue}>
//           <Space direction="vertical">
//             <Radio value={1}>图像不清</Radio>
//             <Radio value={2}>答错题目</Radio>
//             <Radio value={3}>其他错误</Radio>
//           </Space>
//         </Radio.Group>
//         <Input placeholder="请输入问题" onChange={this.handelChange.bind(this)} style={{ marginTop: 10 }} />
//       </Modal>
//     )
//   }

}