import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table } from 'antd';
import { ExclamationCircleOutlined, ArrowLeftOutlined } from '@ant-design/icons';
import './index.less'
import group from "../../../../api/group";
import Marking from "../../../../api/marking";
import * as Util from "../../../../util/Util";
const { Option } = Select;
export default class index extends Component {

    state = {

    }
    componentDidMount() {

    }


    render() {
        return (
            <DocumentTitle title='用户管理-用户管理'>
            <div className="export-page" data-component="export-page">

            </div>
          </DocumentTitle>
        )
    }

}