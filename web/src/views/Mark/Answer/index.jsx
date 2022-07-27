import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'
import Marking from "../../../api/marking";

export default class index extends Component {

  userId = 1
  state = {
    papers: [],
    keyTest: [],
  };
  getAllPaper = () => {
    Marking.testList({ userId: this.userId })
    .then((res) => {
      console.log(res)
      if (res.data.status == "10000") {
        let papers = [...res.data.data.TestIds]
        this.setState(
          {
            papers ,
          }
        )
        this.getAnswer();
      }
    })
    .catch((e) => {
      console.log(e)
    })
  }

  getAnswer = () => {
    console.log('====')
    Marking.testAnswer({ userId: this.userId, testId:  this.state.papers[0]})
    .then((res) => {
      console.log(res)
      if (res.data.status == "10000") {
        this.setState({
          keyTest: res.data.data.Pics
        })
      }
    })
    .catch((e) => {
      console.log(e)
    })
  }
  componentDidMount() {
    this.getAllPaper();
  }


  // 答案区
  showTest = () => {
    let testPaper = null;
    if (this.state.keyTest != undefined || this.state.keyTest != null) {
      testPaper = this.state.keyTest.map((item) => {
        return <div className="test-question-img">
        <img src={'data:image/jpg;base64,'+item} alt="加载失败" />
        </div>
      })
    }
    return testPaper
  }

  render() {
    return (
      <DocumentTitle title="阅卷系统-答案">
        <div className="answer-tasks-page" data-component="answer-tasks-page">
          <div className="answer-paper">
            {
              this.showTest()
            }
          </div>
          <div className="answer-score">

          </div>
        </div>
      </DocumentTitle>
    )
  }

}
