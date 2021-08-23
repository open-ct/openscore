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


    render() {
        return (
            <DocumentTitle title="阅卷系统-总体进度">
                <div className="all-monitor-page" data-component="all-monitor-page">
                    <div className="search-container">
                        <div className="question-select">
                            题目选择：<Select
                                style={{ width: 120 }}>


                            </Select>
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
