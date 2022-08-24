import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Select, Table} from "antd";
import * as Settings from "../../../../Setting";
import "./index.less";
import group from "../../../../api/group";
import Manage from "../../../../api/manage";
const {Option} = Select;
export default class index extends Component {

    supervisorId = "2"

    state = {
      userInfo: {},
      questionList: [],
      tableData: [],
      subjectList: [],
    }

    columns = [
      {
        title: "老师",
        width: 120,
        dataIndex: "UserName",
      },
      {
        title: "评卷数量",
        width: 120,
        dataIndex: "TestDistributionNumber",
      },
      {
        title: "阅卷完成数",
        width: 120,
        dataIndex: "TestSuccessNumber",
      },
      {
        title: "阅卷失败数量",
        width: 120,
        dataIndex: "TestProblemNumber",
      },
      {
        title: "未评数量",
        width: 120,
        dataIndex: "TestRemainingNumber",
      },
      {
        title: "评卷速度（秒/份）",
        width: 180,
        dataIndex: "MarkingSpeed",
      },
      {
        title: "预计时间（小时）",
        width: 120,
        dataIndex: "PredictTime",
      },
      {
        title: "平均分",
        width: 120,
        dataIndex: "AverageScore",
      },
      {
        title: "有效度",
        width: 120,
        dataIndex: "Validity",
      },
      {
        title: "标准差",
        width: 120,
        dataIndex: "StandardDeviation",
      },
      {
        title: "在线情况",
        width: 120,
        dataIndex: "IsOnline",
      },

    ]

    questionList = () => {
      Manage.subjectList().then((res) => {
        this.setState({subjectList: res.data.data.subjectVOList});
      })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }
    tableData = (questionId) => {
      group.teacherMonitor({supervisorId: "2", questionId: questionId})
        .then((res) => {
          if (res.data.status === "10000") {
            let tableData = [];
            for (let i = 0; i < res.data.data.teacherMonitoringList.length; i++) {
              let item = res.data.data.teacherMonitoringList[i];
              tableData.push({
                UserName: item.UserName,
                TestDistributionNumber: item.TestDistributionNumber,
                TestSuccessNumber: item.TestSuccessNumber,
                TestRemainingNumber: item.TestRemainingNumber,
                TestProblemNumber: item.TestProblemNumber,
                MarkingSpeed: item.MarkingSpeed,
                AverageScore: item.AverageScore,
                Validity: item.Validity,
                StandardDeviation: item.StandardDeviation,
                PredictTime: item.PredictTime,
                IsOnline: item.IsOnline === 1 ? "在线" : "离线",
              });
            }
            this.setState({
              tableData,
            });
          }
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }
    componentDidMount() {
      this.questionList();
    }

    // 题目选择区
    selectBox = () => {
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

    select = (e) => {
      let index;
      for (let i = 0; i < this.state.questionList.length; i++) {
        if (this.state.questionList[i].QuestionName === e) {
          index = i;
        }
      }
      this.tableData(this.state.questionList[index].QuestionId);
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
        <DocumentTitle title="阅卷系统-教师监控">
          <div className="teacher-monitor-page" data-component="teacher-monitor-page">
            <div className="search-container">
              <div className="question-select">
                  题目选择：<Select
                  showSearch
                  style={{width: 120, marginRight: 50}}
                  optionFilterProp="label"
                  onSelect={(e) => {this.select(e);}}
                  filterOption={(input, option) =>
                    option.label.indexOf(input) >= 0
                  }
                  filterSort={(optionA, optionB) =>
                    optionA.label.localeCompare(optionB.label)
                  }
                  placeholder={this.state.questionList.length > 0 ? this.state.questionList[0].QuestionName : null}
                  defaultValue={this.state.questionList.length > 0 ? this.state.questionList[0].QuestionName : null}
                >
                  {
                    this.selectBox()
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
          </div>
        </DocumentTitle>
      );
    }

}
