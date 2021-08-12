import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'

export default class index extends Component {


  render() {

    return (
      <DocumentTitle title="阅卷系统-答案">
        <div className="answer-tasks-page" data-component="answer-tasks-page">
          <div className="answer-paper">

          </div>
          <div className="answer-score">

          </div>
        </div>
      </DocumentTitle>
    )
  }

}
