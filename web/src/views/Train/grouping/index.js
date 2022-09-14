import React, {useEffect, useState} from "react";
import {
  Button,
  Form,
  Input,
  Select,
  Table,
  message
} from "antd";

import "./index.less";
import Manage from "../../../api/manage";

export default function Grouping(props) {

  const [subjectList, setSubjectList] = useState([]);
  const [questionList, setQuestionList] = useState([]);
  const [schoolList, setSchoolList] = useState([]);
  const [studentList, setStudentsList] = useState([]);
  const [selectList, setSelectList] = useState([]);
  const [paperList, setPaperList] = useState([]);

  useEffect(() => {
    Manage.subjectList().then((res) => {
      if (res.data.data.subjectVOList !== null) {
        setSubjectList(res.data.data.subjectVOList);
      }
    });
  }, []);

  useEffect(() => {
    Manage.getSchoolsList().then((res) => {
      if (res.data.data !== null) {
        setSchoolList(res.data.data);
      }
    });
  }, []);

  useEffect(() => {
    message.warn("打分时用\"-\"分隔分数，例 7-4-13");
  }, []);

  const rowSelection = {

    onChange: (_, selectedRows) => {
      if (selectedRows.length === 0) {
        setSelectList([]);
        setPaperList([]);
        return;
      }
      const newSelectList = studentList.filter((student) => {
        for (let select of selectedRows) {
          if (select.ticket_id === student.ticket_id) {
            return true;
          }
        }
      });

      setSelectList(newSelectList);
      Promise.all(selectedRows.map((row) =>
        Manage.getListTestPaperInfo({test_id: row.test_id})))
        .then((res) => {
          const list = res.map((r) => r.data.data);
          setPaperList(list);
        });

    },

    getCheckboxProps: (record) => ({
      disabled: record.name === "Disabled User",
      // Column configuration not to be checked
      name: record.name,
    }),
  };

  const columns = [
    {
      title: "准考证号",
      dataIndex: "ticket_id",
      render: (text) => <a>{text}</a>,
    },
    {
      title: "打分",
      dataIndex: "point",
      render: (_, record) => <Input onChange={(e) => {
        const newList = [...studentList];
        newList[record.key] = {...newList[record.key], point: e.target.value};
        const newSelectList = selectList.map((select) => {
          if (select.ticket_id === record.ticket_id) {
            select.point = e.target.value;
          }
          return select;
        });
        setSelectList(newSelectList);
        setStudentsList(newList);
      }} style={{width: 50}} />,
    },
  ];

  const handleFormChange = (e, fileds) => {
    if (e.length == 0) {
      return;
    }
    switch (e[0].name[0]) {
    case "subject":
      Manage.questionInfo(e[0].value).then((res) => {
        if(res.data.data.topicVOList !== null) {
          setQuestionList(res.data.data.topicVOList);
        }
      });
      break;
    case "question_id":
      Manage.getListTestPapersByQuestionId({question_id: e[0].value, school: fileds[2].value, ticket_id: fileds[4].value}).then((res) => {
        if (res.data.data === null) {
          setStudentsList([]);
          return;
        }
        const stuList = res.data.data.map((stu, i) => ({key: i, test_id: stu.test_id, ticket_id: stu.ticket_id}));
        setStudentsList(stuList);
      });
      break;
    case "group_name":
      break;
    case "school":
      if (!fileds[1].value) {
        return;
      }
      Manage.getListTestPapersByQuestionId({school: e[0].value, question_id: fileds[1].value, ticket_id: fileds[4].value}).then((res) => {
        if (res.data.data === null) {
          setStudentsList([]);
          return;
        }
        const stuList = res.data.data.map((stu, i) => ({key: i, test_id: stu.test_id, ticket_id: stu.ticket_id}));
        setStudentsList(stuList);
      });
      break;
    case "ticket_id":
      if (!fileds[1].value) {
        return;
      }
      Manage.getListTestPapersByQuestionId({school: fileds[2].value, question_id: fileds[1].value}).then((res) => {
        if (res.data.data === null) {
          setStudentsList([]);
          return;
        }
        const stuList = res.data.data
          .map((stu, i) => ({key: i, test_id: stu.test_id, ticket_id: stu.ticket_id}))
          .filter((stu) => stu.ticket_id.includes(e[0].value));
        setStudentsList(stuList);
      });
      break;
    }
  };

  const onFinish = (values) => {
    if (typeof values.group_name === "undefined") {
      message.error("请填写组别名称!");
      return;
    } else if (typeof values.question_id === "undefined") {
      message.error("请填写大题号!");
      return;
    } else if (selectList.length === 0) {
      message.error("至少选择一位同学的题目!");
      return;
    }
    for (let stu of selectList) {
      if (typeof stu.point === "undefined") {
        message.error("请记住给所有选中的试卷打分！");
        return;
      }

      paperList.forEach((paper) => {
        if (paper[0].ticket_id === stu.ticket_id && paper.length !== stu.point.split("-").length) {
          message.error("请注意分数填写形式！");
          return;
        }
      });
    }

    const {group_name, question_id} = values;
    const papers = selectList.map((select) => ({
      test_id: select.test_id,
      scores: select.point.split("-").map(p => +p),
    }));

    Manage.teachingPaperGrouping({question_id, group_name, papers}).then((res) => {
      message.success("分组成功!");
      setTimeout(() => {
        location.reload();
      }, 1000);
    });
  };

  return (
    <div className="grouping-page">
      <div className="subject-list">
        {
          paperList.length === 0 ?
            <h1>请先选择学生试卷</h1> :
            paperList.map((paper, i) =>
              <div key={i}>
                <h3>学生学号:{paper[0].ticket_id}</h3>
                <div className="paper-list">
                  {
                    paper.map((question, i) =>
                      <div className="paper" key={i}>
                        <img src={question.pic_src} alt="题目" />
                      </div>)
                  }
                </div>
              </div>)}
      </div>
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
          onFieldsChange={handleFormChange}
        >
          <Form.Item label="学科" name="subject">
            <Select>
              {subjectList.map((sub) => <Select.Option key={sub.SubjectId} value={sub.SubjectName}>{sub.SubjectName}</Select.Option>)}
            </Select>
          </Form.Item>
          <Form.Item label="大题号" name="question_id">
            <Select>
              {questionList.map((question) => <Select.Option key={question.TopicId} value={question.TopicId}>{question.TopicId}</Select.Option>)}
            </Select>
          </Form.Item>
          <Form.Item label="考场" name="school">
            <Select>
              {schoolList.map((school, i) => <Select.Option key={i} value={school}>{school}</Select.Option>)}
            </Select>
          </Form.Item>
          <Form.Item label="分组名称" name="group_name">
            <Input />
          </Form.Item>
          <Form.Item label="准考证号" name="ticket_id">
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
            dataSource={studentList}
            pagination={{defaultPageSize: 5}}
          />
        </div>
      </div>
    </div>
  );
}
