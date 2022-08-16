import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Button, Upload, message} from "antd";
import {DownloadOutlined, UploadOutlined} from "@ant-design/icons";
import "./index.less";
import * as Setting from "../../../../Setting";
import axios from "axios";

export default class index extends Component {

  state = {

  }

  componentDidMount() {

  }

  paperExport = () => {
    axios.request({
      url: "/marking/supervisor/writeScoreExcel",
      headers: {
        "Content-Type": "application/json", // 重要
        "accept": "application/octet-stream", // 重要
      },
      method: "POST",
      data: "",
      params: "",
      responseType: "blob", // 重要
    }).then(function(response) {
      var data = response.data;
      var url = URL.createObjectURL(data);// 重要
      let link = document.createElement("a");
      link.href = url;
      link.download = "试卷成绩.xlsx";// 重要--决定下载文件名
      link.click();
      link.remove();
    }).catch(function(e) {console.log(e);});

  }

  render() {
    const props_1 = {
      name: "excel",
      action: Setting.ServerUrl + "/openct/marking/admin/readExcel",
      headers: {
        authorization: null,
      },
      onChange(info) {
        if (info.file.status !== "uploading") {
          console.log(info.file, info.fileList);
        }
        if (info.file.status === "done") {
          message.success(`${info.file.name} file uploaded successfully`);
        } else if (info.file.status === "error") {
          message.error(`${info.file.name} file upload failed.`);
        }
      },
    };
    const props_2 = {
      name: "excel",
      action: Setting.ServerUrl + "/openct/marking/admin/readExapmleExcel",
      headers: {
        authorization: "authorization-text",
      },
      onChange(info) {
        if (info.file.status !== "uploading") {
          console.log(info.file, info.fileList);
        }
        if (info.file.status === "done") {
          message.success(`${info.file.name} file uploaded successfully`);
        } else if (info.file.status === "error") {
          message.error(`${info.file.name} file upload failed.`);
        }
      },
    };
    const props_3 = {
      name: "excel",
      action: Setting.ServerUrl + "/openct/marking/admin/readAnswerExcel",
      headers: {
        authorization: "authorization-text",
      },
      onChange(info) {
        if (info.file.status !== "uploading") {
          console.log(info.file, info.fileList);
        }
        if (info.file.status === "done") {
          message.success(`${info.file.name} file uploaded successfully`);
        } else if (info.file.status === "error") {
          message.error(`${info.file.name} file upload failed.`);
        }
      },
    };
    return (
      <DocumentTitle title="试卷管理-导入试卷">
        <div className="export-page" data-component="export-page">
          <div className="export-page" data-component="export-page">
            <Upload {...props_1}>
              <Button icon={<UploadOutlined />} style={{marginRight: 24}}>导入试卷</Button>
            </Upload>
            <Upload {...props_2}>
              <Button icon={<UploadOutlined />} style={{marginRight: 24}}>导入样卷</Button>
            </Upload>
            <Upload {...props_3}>
              <Button icon={<UploadOutlined />} style={{marginRight: 24}}>导入答案</Button>
            </Upload>
          </div>
          <div className="export-page" data-component="export-page">
            <Button icon={<DownloadOutlined />} onClick={this.paperExport} style={{marginRight: 24}}>导出成绩</Button>
          </div>
        </div>
      </DocumentTitle>
    );
  }

}
