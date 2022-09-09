import React from "react";
import {Select, Table} from "antd";

import "./index.less";

const columns = [
  {
    title: "A组",
    dataIndex: "group",
  },
  {
    title: "操作",
    dataIndex: "",
    render: () => <a style={{width: 50}}>删除</a>,
  },
];
const data = [];

for (let i = 0; i < 46; i++) {
  data.push({
    key: i,
    group: `Edward King ${i}`,
    concordance: 32,
    test_id: `London, Park Lane no. ${i}`,
  });
}

export default function Detail() {
  return (
    <div className="detail-page">
      <div className="search-container">
          学科选择：
        <Select
          style={{width: 120, marginRight: 50}}
          optionFilterProp="label"
        >
        </Select>
          大题选择：
        <Select
          style={{width: 120, marginRight: 50}}
          optionFilterProp="label">
        </Select>
          考场选择：
        <Select
          style={{width: 120}}
          optionFilterProp="label">
        </Select>
      </div>
      <div className="display-container">
        <Table
          pagination={{position: ["bottomCenter"]}}
          columns={columns}
          dataSource={data}
        />
      </div>
    </div>
  );
}
