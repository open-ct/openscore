import React, { Component, useState } from 'react'
import DocumentTitle from 'react-document-title'
import { Modal, Dropdown, Button, message, Space, Tooltip, Select, Radio, Input, Table, Upload } from 'antd';
import { ExclamationCircleOutlined, ArrowLeftOutlined, UploadOutlined } from '@ant-design/icons';
import './index.less'
import * as Setting from "../../../../Setting";
import axios from 'axios'
// function getServerUrl() {
//   const hostname = window.location.hostname
//   if (hostname === 'localhost') {
//     return `http://${hostname}:8080/openct`
//   }
//   return '/openct'
// }
// axios.defaults.baseURL = getServerUrl();
// axios.defaults.withCredentials=true;
// // axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
// axios.defaults.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
const { Option } = Select;
export default class index extends Component {

  state = {

  }

  componentDidMount() {

  }
  paperExport = () => {
    axios.request({
      url:"/marking/supervisor/writeScoreExcel",
      headers:{
        "Content-Type": "application/json",//重要
        "accept": "application/octet-stream",//重要
      },
      method:"POST",
      data:"",
      params:"",
      responseType: 'blob'//重要
    }).then(function (response) {
      var data=response.data;
      var url = URL.createObjectURL(data);//重要
      let link = document.createElement('a');
      link.href = url;
      link.download = "试卷导出.xlsx";//重要--决定下载文件名
      link.click();
      link.remove();
    }).catch(function (e) {console.log(e)})
    
}

  render() {
    const props_1 = {
      name: 'excel',
      action: Setting.ServerUrl+'/marking/admin/readExcel',
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
    const props_2 = {
      name: 'excel',
      action: Setting.ServerUrl+'/marking/admin/readExampleExcel',
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
    const props_3 = {
      name: 'excel',
      action: Setting.ServerUrl+'/marking/admin/readAnswerExcel',
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
      <DocumentTitle title='试卷管理-导入试卷'>
        <div className="export-page" data-component="export-page">
          <Upload {...props_1}>
            <Button icon={<UploadOutlined />}  style={{marginRight: 24}}>导入试卷</Button>
          </Upload>
          <Upload {...props_2}>
            <Button icon={<UploadOutlined />} style={{marginRight: 24}}>导入样卷</Button>
          </Upload>
          <Upload {...props_3}>
            <Button icon={<UploadOutlined />} style={{marginRight: 24}}>导入答案</Button>
          </Upload>
          <div>
            <Button onClick={this.paperExport} style={{marginRight: 24}}>导出成绩</Button>
          </div>
        </div>
      </DocumentTitle>
    )
  }

}