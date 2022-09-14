import {Select, Table} from "antd";
import React, {useEffect, useState} from "react";
import Manage from "../../../api/manage";

import "./index.less";

const Evaluate = () => {

  const [group, setGroupList] = useState([]);
  const [data, setData] = useState([]);
  const columns = [
    {
      title: "用户",
      dataIndex: "account",
    },
    {
      title: "一致率",
      dataIndex: "concordance_rate",
    },
    {
      title: "题目分数",
      dataIndex: "scores",
    },
    {
      title: "操作",
      dataIndex: "",
      render: (_, record) => <a
        onClick={() => {
          console.log(record);
        }}
        style={{width: 50}}>
          合格
      </a>,
    },
  ];

  useEffect(() => {
    Manage.getListPaperGroups().then((res) => {
      if (res.data.data !== null) {
        setGroupList(res.data.data.groups);
      }
    });
  }, []);

  const handleSelectChange = (group_id) => {
    Manage.getListGroupGrades({group_id}).then((res) => {

      if (res.data.data !== null) {
        const {scores, teacher_grades} = res.data.data;
        const scoreStr = scores.join("-");
        const newData = [];
        newData.push({
          scores: scoreStr,
          account: "管理员",
          concordance_rate: "/",
        });
        if (teacher_grades !== null) {
          teacher_grades.map((teacher) => {
            newData.push({
              key: teacher.teacher_account,
              account: teacher.teacher_account,
              concordance_rate: teacher.concordance_rate,
              scores: teacher.scores.join("-"),
            });
          });
        } else {
          newData.push(
            {
              account: "暂未有人完成评卷",
            }
          );
        }
        setData(newData);
      }

    });
  };

  // for (let i = 0; i < 46; i++) {
  //   data.push({
  //     key: i,
  //     group: `Edward King ${i}`,
  //     concordance: 32,
  //     test_id: `London, Park Lane no. ${i}`,
  //   });
  // }

  return (
    <div className="evaluate-page">
      <div className="search-container">
        <Select placeholder="组别选择" onChange={handleSelectChange}>
          {group.map((g) => <Select.Option key={g.group_id} value={g.group_id}>{g.group_name}</Select.Option>)}
        </Select >
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
