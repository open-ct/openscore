import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined, ArrowLeftOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";

export default class index extends Component {

    state = {
        type: '',
        paperButton: 1
    }
    componentDidMount() {
        if (this.props.match.params.type === "1") {
            this.setState({
                type: '仲裁卷'
            })
        } else if (this.props.match.params.type === "2") {
            this.setState({
                type: '问题卷'
            })
        } else {
            this.setState({
                type: '正评卷'
            })
        }
    }

    // 选择区

    // 提交区
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
            //   this.problemModal()
            }
          </div>
        );
      }
    render() {
        return (
            <DocumentTitle title={'试卷管理-' + this.state.type}>
                <div className="group-markTaks-page" data-component="group-markTaks-page">
                    <div className="search-container">
                        <div className="goBack" onClick={() => { window.history.back(-1) }}>
                            <ArrowLeftOutlined />
                            <span className="paperDetail">试卷详情</span>
                            <span className="goBackPre">返回仲裁卷</span>
                        </div>
                        <div className="currentPaperNo">
                            试卷号：11110
                        </div>
                        <div className="paperNum">
                            {this.state.type}数：30
                        </div>
                        <div className="buttonList">
                            <Button type="primary" style={{ width: 60 }} onClick={()=>{this.paperDisplay(1)}}>评卷</Button>
                            <Button type="default" style={{ width: 60 }} onClick={()=>{this.paperDisplay(2)}}>答案</Button>
                            <Button type="default" style={{ width: 60 }} onClick={()=>{this.paperDisplay(3)}}>样卷</Button>
                            <Button type="default" style={{ width: 32 }} onClick={()=>{this.paperDisplay(4)}}>…</Button>
                        </div>
                    </div>
                    <div className="paperPage">
                        <div className="mark-paper">

                        </div>
                        <div className="mark-score">
                            {
                                this.renderPush()
                            }
                        </div>
                    </div>
                </div>
            </DocumentTitle>
        )
    }




}
