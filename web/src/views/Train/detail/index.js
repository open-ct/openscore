import React, {useEffect, useState} from "react";
import {Button, Popconfirm, Select, Table} from "antd";
import Manage from "../../../api/manage";
import "./index.less";

export default function Detail() {
  const [group, setGroup] = useState([]);
  const [data, setData] = useState([]);
  const columns = [
    {
      title: "试卷ID",
      dataIndex: "test_id",
    },
    {
      title: "操作",
      dataIndex: "",
      render: (_, record) =>
        <Popconfirm
          title={`Sure to delete user: ${record.account} ?`}
          onConfirm={() => {
            Manage.deletePaperFromGroup({...record}).then((res) => {
              setData(data.filter((d) => d.test_id !== record.test_id));
            });
          }}
        >
          <Button style={{marginBottom: "10px", marginRight: "10px"}} type="danger">{"删除"}</Button>
        </Popconfirm>,
    },
  ];

  useEffect(() => {
    Manage.getListPaperGroups().then((res) => {
      if (res.data.data.groups !== null) {
        setGroup(res.data.data.groups);
      }
    });
  }, []);

  const handleGroupChange = (e) => {

    group.forEach((g) => {
      if (g.group_id === e) {
        const newData = g.papers.map((p) => {
          return {test_id: p.test_id, group_id: g.group_id};
        });
        const newColums = [...columns];
        newColums[0].title = g.group_name;
        setData(newData);
      }
    });
  };
  return (
    <div className="detail-page">
      <div className="search-container">
          组别选择：
        <Select
          style={{width: 120, marginRight: 50}}
          optionFilterProp="label"
          onChange={handleGroupChange}
        >
          {group.map((g) => <Select.Option key={g.group_id} value={g.group_id}>{g.group_name}</Select.Option>)}
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
