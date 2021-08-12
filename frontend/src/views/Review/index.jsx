import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Table, Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'

export default class index extends Component {
  state = {
    reviewVisible : true
  }
  columns = [
    {
      title: '序号',
      dataIndex: 'order',
      key: 'order',
      render: text => <a>{text}</a>,
    },
    {
      title: '1',
      dataIndex: 'age',
      key: 'age',
    },
    {
      title: '2',
      dataIndex: 'address',
      key: 'address',
    },
    {
      title: '3',
      key: 'tags',
      dataIndex: 'tags',
    },
    {
      title: '4',
      key: 'action',
      render: (text, record) => (
        <Space size="middle">
          <a>Invite {record.name}</a>
          <a>Delete</a>
        </Space>
      ),
    },
  ];
  
  data = [
    {
      key: '1',
      order: '试卷密号',
      age: 32,
      address: 'New York No. 1 Lake Park',
      tags: ['nice', 'developer'],
    },
    {
      key: '2',
      order: '分数',
      age: 42,
      address: 'London No. 1 Lake Park',
      tags: ['loser'],
    },
    {
      key: '3',
      order: 'Joe Black',
      age: 32,
      address: 'Sidney No. 1 Lake Park',
      tags: ['cool', 'teacher'],
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
    console.log('回评')
    this.setState({
      reviewVisible: false,
    });
  };

  handleCancel = () => {
    console.log('Clicked cancel button');
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
      <Table columns={this.columns} dataSource={this.data} />
    ) 
  }
}