import {Select, Table} from "antd";
import React from "react";

import "./index.less";

const columns = [
  {
    title: "A组",
    dataIndex: "group",
  },
  {
    title: "一致率",
    dataIndex: "concordance",
  },
  {
    title: "1",
    dataIndex: "test_id",
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

const Evaluate = () => {

  return (
    <div className="evaluate-page">
      <div className="search-container">
        <Select placeholder="组别选择"></Select >
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
};

export default Evaluate;
