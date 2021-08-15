import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Table, Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'
import Marking from "../../api/marking";

export default class index extends Component {
  userId = "1"
  state = {
    reviewVisible : true,
    reviewList : []
  }
  componentDidMount() {
    this.getReviewList();
  }
  getReviewList = () => {
    Marking.testReview({ userId: this.userId})
    .then((res) => {
      if (res.data.status == "10000") {
        let reviewList = []
        for (let i = 0; i < res.data.data.records.length; i++) {
          reviewList.push({
            order: i+1,
            test_id: res.data.data.records[i].Test_id,
            score: res.data.data.records[i].Score
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
      <DocumentTitle title="阅卷系统-回评">
        <div className="answer-tasks-page" data-component="answer-tasks-page">
          <div className="answer-paper">
            
          </div>
          <div className="answer-score">

          </div>
          {
            this.reviewModal()
          }
        </div>
      </DocumentTitle>
    )

  }
  handleOk = () => {
    this.setState({
      reviewVisible: false,
    });
  };

  handleCancel = () => {
    this.setState({
      reviewVisible: false,
    });
  };
  reviewModal() {
    const { problemValue } = this.state;
    return (
      <Modal
        title="回评列表"
        visible={this.state.reviewVisible}
        onOk={this.handleOk}
        onCancel={this.handleCancel}
        width={800}
      >
        {
          this.reviewTable()
        }
      </Modal>
    )
  }
  reviewTable() {
    return (
      <Table columns={this.columns} dataSource={this.state.reviewList} />
    ) 
  }
}