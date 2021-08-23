import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";

export default class index extends Component {


    componentDidMount() {

    }

    // 选择区

    paperMark =() => {
        this.props.history.push('/home/group/markTasks/3')
    }
    render() {
        return (
            <DocumentTitle title="阅卷系统-正评卷">
                <div className="group-marking-page" data-component="group-marking-page">
                    <div className="search-container">
                        <div className="question-select">
                            题目选择：<Select
                                style={{ width: 120 }}>
                            </Select>
                        </div>
                        <div className="paper-num">
                            正评卷数：10
                        </div>
                        <div className="paper-mark" onClick={() => {this.paperMark()}}>
                            正评卷查询
                        </div>    
                    </div>
                    <div className="display-container">
                        <Table 
                            pagination={{ position: ['bottomCenter'] }}

                        />
                    </div>
                </div>
            </DocumentTitle>
        )
    }




}
