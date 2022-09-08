import React from "react";
import {
  Button,
  Form,
  Input,
  Select,
  Table
} from "antd";

import "./index.less";

const columns = [
  {
    title: "准考证号",
    dataIndex: "name",
    render: (text) => <a>{text}</a>,
  },
  {
    title: "打分",
    dataIndex: "",
    render: () => <Input style={{width: 50}} />,
  },
];
const data = [
  {
    key: "1",
    name: "20202",
  },
  {
    key: "2",
    name: "20312",
  },
  {
    key: "3",
    name: "20731",
  },
  {
    key: "4",
    name: "20731",
  }, {
    key: "5",
    name: "20731",
  }, {
    key: "6",
    name: "20731",
  }, {
    key: "7",
    name: "20731",
  }, {
    key: "8",
    name: "20731",
  },
  {
    key: "9",
    name: "20731",
  }, {
    key: "10",
    name: "20731",
  }, {
    key: "11",
    name: "20731",
  },
];
const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, "selectedRows: ", selectedRows);
  },
  getCheckboxProps: (record) => ({
    disabled: record.name === "Disabled User",
    // Column configuration not to be checked
    name: record.name,
  }),
};

export default function Grouping() {

  const onFinish = (values) => {
    console.log(values);
  };

  return (
    <div className="grouping-page">
      <div className="subject-list"></div>
      <div className="grouping-setting">
        <Form
          labelCol={{
            span: 6,
            offset: 2,
          }}
          wrapperCol={{
            span: 12,
          }}
          style={{
            marginTop: 10,
          }}
          layout="horizontal"
          labelAlign="left"
          onFinish={onFinish}
        >
          <Form.Item label="学科" name="subject">
            <Select>
              <Select.Option value="demo">Demo</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item label="大题号" name="question">
            <Select>
              <Select.Option value="demo">Demo</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item label="考场号" name="position">
            <Select>
              <Select.Option value="demo">Demo</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item label="试卷分组" name="grouping">
            <Select>
              <Select.Option value="demo">Demo</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item label="准考证号" name="student_id">
            <Input />
          </Form.Item>
          <Form.Item style={{position: "absolute", bottom: 120, left: "50%", transform: "translateX(-50%)"}} wrapperCol={{offset: 8, span: 16}}>
            <Button htmlType="submit" type="primary">确认分组</Button>
          </Form.Item>
        </Form>
        <div className="student-table">
          <Table
            style={{width: "80%"}}
            size="small"
            rowSelection={{
              type: "checkbox",
              ...rowSelection,
            }}
            columns={columns}
            dataSource={data}
            pagination={{defaultPageSize: 5}}
          />
        </div>
      </div>
    </div>
  );
}
