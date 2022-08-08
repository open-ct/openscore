import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Button, Upload, message} from "antd";
import {UploadOutlined} from "@ant-design/icons";
import "./index.less";
export default class index extends Component {

  state = {

  }

  componentDidMount() {

  }

  render() {
    const props_1 = {
      name: "excel",
      action: "http://localhost:8080/openct/marking/admin/readExcel",
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
    const props_2 = {
      name: "excel",
      action: "http://localhost:8080/openct/marking/admin/readExapmleExcel",
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
      action: "http://localhost:8080/openct/marking/admin/readAnswerExcel",
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
      <DocumentTitle title="试卷管理-导入导出试卷">
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
      </DocumentTitle>
    );
  }

}
