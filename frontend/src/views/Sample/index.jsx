import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Table, Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'

export default class index extends Component {
  columns = [
    {
      title: '序号',
      width: 90,
      dataIndex: 'order',
      key: 'order',

      render: text => <a>{text}</a>,
    },
    {
      title: '小题',
      width: 90,
      dataIndex: 'age',
      key: 'age',
    },
    {
      title: '密号',
      width: 90,
      dataIndex: 'age',
      key: 'age',

    },
    {
      title: '专家',
      width: 90,
      dataIndex: 'address',
      key: 'address',
    },
    {
      title: '老师',
      key: 'tags',
      dataIndex: 'tags',
      width: 90,
    },
    {
      title: '差异',
      key: 'action',
      width: 90,
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
      address: '2030',
      tags: ['nice', 'developer'],
    },
    {
      key: '2',
      order: '分数',
      age: 42,
      address: '2030',
      tags: ['loser'],
    },

  ];

  render() {

    return (
      <DocumentTitle title="阅卷系统-样卷">
        <div className="answer-tasks-page" data-component="answer-tasks-page">
          <div className="answer-paper">

          </div>
          <div className="answer-score">
            {
              this.sampleTable()
            }
          </div>
        </div>
      </DocumentTitle>
    )
  }
  sampleTable() {
    return(
      <Table columns={this.columns} 
      dataSource={this.data} 
      scroll={{x:400}}
      pagination={{ position: ['bottomCenter'] }} />
    )
  }

}