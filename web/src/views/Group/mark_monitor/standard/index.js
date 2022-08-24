import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Select, Table} from "antd";
import * as Settings from "../../../../Setting";
import "./index.less";
import group from "../../../../api/group";
import Manage from "../../../../api/manage";
const {Option} = Select;
export default class index extends Component {

  componentDidMount() {
    this.questionList();
  }

    // 选择区
    state = {
      questionList: [],
      tableData: [],
      fullScore: undefined,
      subjectList: [],
    }

    questionList = () => {
      Manage.subjectList().then((res) => {
        this.setState({subjectList: res.data.data.subjectVOList});
      })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
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
    tableData = (questionId) => {
      group.standardMonitor({supervisorId: "2", questionId: questionId})
        .then((res) => {
          if (res.data.status === "10000") {
            let tableData = [];
            for (let i = 0; i < res.data.data.ScoreDeviationVOList.length; i++) {
              let item = res.data.data.ScoreDeviationVOList[i];
              tableData.push({
                UserName: item.UserName,
                Deviation: item.DeviationScore,
              });
            }
            this.setState({
              tableData,
              fullScore: res.data.data.fullScore,
            });
          }
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }
    columns = [
      {
        title: "教师",
        width: 150,
        dataIndex: "UserName",
      },
      {
        title: "标准差",
        width: 180,
        dataIndex: "Deviation",
      },
    ]
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
        <DocumentTitle title="阅卷系统-标准差监控">
          <div className="standard-monitor-page" data-component="standard-monitor-page">
            <div className="search-container">
              <div className="question-select">
                题目选择：<Select
                  showSearch
                  style={{width: 120, marginRight: 70}}
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
              </div>
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
              <div className="question-score">
                满分：{this.state.fullScore}
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
