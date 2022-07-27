import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table, Upload } from 'antd';
import { ExclamationCircleOutlined, ArrowLeftOutlined, UploadOutlined } from '@ant-design/icons';
import './index.less'
import * as Setting from "../../../../Setting";
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
        const props_1 = {
            name: 'excel',
            action: Setting.ServerUrl+'/openct/marking/admin/readUserExcel',
            headers: {
              authorization: 'authorization-text',
            },
            onChange(info) {
              if (info.file.status !== 'uploading') {
                console.log(info.file, info.fileList);
              }
              if (info.file.status === 'done') {
                message.success(`${info.file.name} file uploaded successfully`);
              } else if (info.file.status === 'error') {
                message.error(`${info.file.name} file upload failed.`);
              }
            },
          }
        return (
            <DocumentTitle title='用户管理-用户导入导出'>
                <div className="question-page" data-component="question-page">
                    <Upload {...props_1}>
                        <Button icon={<UploadOutlined />}  style={{marginRight: 24}}>导入用户</Button>
                    </Upload>
                </div>
            </DocumentTitle>
        )
    }

}