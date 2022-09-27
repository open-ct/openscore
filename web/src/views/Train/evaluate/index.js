import {Button, Popconfirm, Select, Table} from "antd";
import React, {useEffect, useState} from "react";
import Manage from "../../../api/manage";

import "./index.less";

const Evaluate = () => {

  const [group, setGroupList] = useState([]);
  const [data, setData] = useState([]);
  const [select, setSelect] = useState();
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
      render: (_, records) => <p dangerouslySetInnerHTML={{__html: records.scores}}></p>
      ,
    },
    {
      title: "操作",
      dataIndex: "",
      render: (_, record) => {
        return record.account === "管理员" ? null :
          record.is_qualified ? <a>已合格</a> :
            <Popconfirm
              title={"确认要让他合格吗"}
              onConfirm={() => {
                Manage.updateUserQualified({account: record.account}).then(() => {
                  handleSelectChange(select);
                });
              }}
            >
              <Button style={{marginBottom: "10px", marginRight: "10px"}} type="primary">合格</Button>
            </Popconfirm>;
      },
    },
  ];

  useEffect(() => {
    Manage.getListPaperGroups().then((res) => {
      if (res.data.data.groups !== null) {
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
            let newScores = teacher.scores.map((s, i) => {
              return s === scores[i] ? s : `<span id='red'>${s}</span>`;
            });
            newData.push({
              key: teacher.teacher_account,
              account: teacher.teacher_account,
              concordance_rate: teacher.concordance_rate,
              scores: newScores.join("-"),
              is_qualified: teacher.is_qualified,
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
        setSelect(group_id);
      }

    });
  };

  return (
    <div className="evaluate-page">
      <div className="search-container">
        <Select placeholder="组别选择" value={select} onChange={handleSelectChange}>
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
