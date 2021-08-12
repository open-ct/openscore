import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'

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
  state = { 
    problemVisible: false,
    problemValue: 1, 
  };
  render() {

    return (
      <DocumentTitle title="阅卷系统-评卷">
        <div className="mark-tasks-page" data-component="mark-tasks-page">
          <div className="mark-paper">

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

  handleChange = (value) => {
    console.log(`selected ${value}`);
  }
  renderScoreDropDown() {
    const { Option } = Select;
    return (
      <div className="score-container">
        <div className="score-select">
          16题&nbsp;&nbsp;(3)&nbsp;&nbsp;：<Select placeholder="请选择分数" style={{ width: 120 }} onChange={this.handleChange}>
            <Option value="0">0</Option>
            <Option value="1">1</Option>
            <Option value="2">2</Option>
          </Select>
        </div>
        <div className="score-select">
          17题&nbsp;&nbsp;(3)&nbsp;&nbsp;：<Select placeholder="请选择分数" style={{ width: 120 }} onChange={this.handleChange}>
            <Option value="0">0</Option>
            <Option value="1">1</Option>
            <Option value="2">2</Option>
          </Select>
        </div>
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
        if (value == 1) {
          console.log('1');
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
    console.log('problem')
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
        <Input placeholder="请输入问题" style={{ marginTop : 10}} />
      </Modal>
    )
  }
}
