import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Select, Table} from "antd";
import * as Settings from "../../../../Setting";
import "./index.less";
import group from "../../../../api/group";
import ReactEcharts from "echarts-for-react";
import Manage from "../../../../api/manage";

const {Option} = Select;
export default class index extends Component {

    supervisorId = "2"

    state = {
      questionList: [],
      teacherList: [],
      tableData: [],
      QuestionId: undefined,
      selfScoreRecordVOList: [],
      subjectList: [],
    }
    columns = [
      {
        title: "序号（试卷编码）",
        width: 150,
        dataIndex: "order",
      },
      {
        title: "评分",
        width: 150,
        dataIndex: "Score",
      },
      {
        title: "自评",
        width: 150,
        dataIndex: "SelfScore",
      },
      {
        title: "实际误差",
        width: 150,
        dataIndex: "Error",
      },
      {
        title: "标准误差",
        width: 150,
        dataIndex: "StandardError",
      },
      {
        title: "是否合格",
        width: 150,
        dataIndex: "IsQualified",
      },

    ]

    // componentDidMount() {
    //   this.questionList();
    // }

    getOption = () => {
      let X_data = [];
      let Y1_data = [];
      let Y2_data = [];
      for (let i = 0; i < this.state.tableData.length; i++) {
        X_data.push(this.state.tableData[i].order);
      }
      for (let i = 0; i < this.state.tableData.length; i++) {
        Y1_data.push(this.state.tableData[i].Score);
      }
      for (let i = 0; i < this.state.tableData.length; i++) {
        Y2_data.push(this.state.tableData[i].SelfScore);
      }
      return {
        xAxis: {
          name: "序号",
          data: X_data,
        },
        yAxis: {
          name: "分数",
        },
        series: [{
          name: "分数",
          type: "bar",
          data: Y1_data,
        },
        {
          name: "分数",
          type: "bar",
          data: Y2_data,
        }],
      };
    };
    questionList = () => {
      Manage.subjectList().then((res) => {
        this.setState({subjectList: res.data.data.subjectVOList});
      })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }
    teacherList = (questionId) => {
      group.selfTeacher({supervisorId: "2", questionId})
        .then((res) => {
          if (res.data.status === "10000") {
            this.setState({
              teacherList: res.data.data.teacherVOList,
            });
            this.tableData(res.data.data.teacherVOList[0].UserId);
          }
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }
    // 题目选择区
    selectQuestionBox = () => {
      let selectList;
      if (this.state.questionList.length !== 0) {
        selectList = this.state.questionList.map((item, i) => {
          return <Option key={i} value={item.QuestionName} label={item.QuestionName}>{item.QuestionName}</Option>;
        });
      } else {
        return null;
      }
      return selectList;
    }
    selectTeacherBox = () => {
      let selectList;
      if (this.state.teacherList.length !== 0) {
        selectList = this.state.teacherList.map((item) => {
          return <Option key={item.UserName} value={item.UserName} label={item.UserName}>{item.UserName}</Option>;
        });
      } else {
        return null;
      }
      return selectList;
    }
    selectTeacher = (e) => {
      Settings.showMessage("error", e);
      let index;
      for (let i = 0; i < this.state.teacherList.length; i++) {
        if (this.state.teacherList[i].UserName === e) {
          index = i;
        }
      }
      this.tableData(this.state.teacherList[index].UserId);
    }
    selectQuestion = (e) => {
      Settings.showMessage("error", e);
      let index;
      for (let i = 0; i < this.state.questionList.length; i++) {
        if (this.state.questionList[i].QuestionName === e) {
          index = i;
        }
      }
      this.setState({
        QuestionId: this.state.questionList[index].QuestionId,
      });
      this.teacherList(this.state.questionList[index].QuestionId);
    }

    tableData = (examinerId) => {
      group.selfMonitor({supervisorId: "2", examinerId})
        .then((res) => {
          if (res.data.status === "10000") {
            let tableData = [];
            let selfScoreRecordVOList = res.data.data.selfScoreRecordVOList;
            for (let i = 0; i < res.data.data.selfScoreRecordVOList.length; i++) {
              let item = res.data.data.selfScoreRecordVOList[i];
              tableData.push({
                order: `${i + 1}(${item.TestId})`,
                Score: item.Score,
                SelfScore: item.SelfScore,
                Error: item.Error,
                IsQualified: item.IsQualified === 0 ? "不合格" : "合格",
                StandardError: item.StandardError,
              });
            }
            this.setState({
              tableData, selfScoreRecordVOList,
            });
          }
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }

    onSelectsub = (e) => {
      group.questionList({subjectName: e})
        .then((res) => {
          if (res.data.status === "10000") {
            this.setState({
              questionList: res.data.data.questionsList,
            });
            if (res.data.data.questionsList.length > 0) {this.tableData(res.data.data.questionsList[0].QuestionId);}
          }
        });
    }

    selectSubject = () => {
      return this.state.subjectList.map((item, i) => {
        return <Option key={i} value={item.SubjectName} label={item.SubjectName}>{item.SubjectName}</Option>;
      });
    }

    render() {
      return (
        <DocumentTitle title="阅卷系统-自评监控">
          <div className="self-monitor-page" data-component="self-monitor-page">
            <div className="search-container">
              <div className="question-select">
                  题目选择：<Select
                  key="question"
                  showSearch
                  style={{width: 120}}
                  optionFilterProp="label"
                  onSelect={(e) => {this.selectQuestion(e);}}
                  filterOption={(input, option) =>
                    option.label.indexOf(input) >= 0
                  }
                  filterSort={(optionA, optionB) =>
                    optionA.label.localeCompare(optionB.label)
                  }
                  placeholder={this.state.questionList.length > 0 ? this.state.questionList[0].QuestionName : null}
                  initialValue={this.state.questionList.length > 0 ? this.state.questionList[0].QuestionName : null}
                >
                  {
                    this.selectQuestionBox()
                  }
                </Select>
              </div>
              <div className="teacher-select">
                  教师选择：<Select
                  key="teacher"
                  showSearch
                  style={{width: 120, marginRight: 70}}
                  optionFilterProp="label"
                  onSelect={(e) => {this.selectTeacher(e);}}
                  filterOption={(input, option) =>
                    option.label.indexOf(input) >= 0
                  }
                  filterSort={(optionA, optionB) =>
                    optionA.label.localeCompare(optionB.label)
                  }
                  placeholder={this.state.teacherList.length > 0 ? this.state.teacherList[0].UserName : null}
                  initialValue={this.state.teacherList.length > 0 ? this.state.teacherList[0].UserName : null}
                >
                  {
                    this.selectTeacherBox()
                  }
                </Select>
                科目选择：<Select
                  style={{width: 120}}
                  optionFilterProp="label"
                  onSelect={(e) => {this.onSelectsub(e);}}
                  filterOption={(input, option) =>
                    option.label.indexOf(input) >= 0
                  }
                  filterSort={(optionA, optionB) =>
                    optionA.label.localeCompare(optionB.label)
                  }>
                  {
                    this.selectSubject()
                  }
                </Select>
              </div>
            </div>
            <div className="display-container">
              <Table
                pagination={{position: ["bottomCenter"]}}
                columns={this.columns}
                dataSource={this.state.tableData}
              />
            </div>
            <ReactEcharts option={this.getOption()} style={{width: 762, height: 300}} />
          </div>
        </DocumentTitle>
      );
    }

}
