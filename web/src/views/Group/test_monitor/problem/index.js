import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Select, Table} from "antd";
import "./index.less";
import group from "../../../../api/group";
const {Option} = Select;
export default class index extends Component {

  componentDidMount() {
    this.questionList();
  }
    columns = [
      {
        title: "试卷号",
        width: 150,
        dataIndex: "TestId",
      },
      {
        title: "阅卷人一账号",
        width: 150,
        dataIndex: "ExaminerId",
      },
      {
        title: "阅卷人一名称",
        width: 150,
        dataIndex: "ExaminerName",
      },
      {
        title: "问题原因",
        width: 150,
        dataIndex: "ProblemType",
      },
    ]

    // 选择区
    state = {
      questionList: [],
      tableData: [],
      count: 0,
      questionIndex: 0,
    }

    questionList = () => {
      group.questionList({adminId: "1", subjectName: JSON.parse(localStorage.getItem("userInfo")).SubjectName})
        .then((res) => {
          if (res.data.status === "10000") {
            this.setState({
              questionList: res.data.data.questionsList,
            });
            console.log(res.data.data.questionsList);
            this.tableData(res.data.data.questionsList[0].QuestionId);
          }
        })
        .catch((e) => {
          console.log(e);
        });
    }
    tableData = (questionId) => {
      group.problemList({supervisorId: "2", questionId: questionId})
        .then((res) => {
          if (res.data.status === "10000") {
            let tableData = [];
            for (let i = 0; i < res.data.data.ProblemUnderCorrectedPaperVOList.length; i++) {
              let item = res.data.data.ProblemUnderCorrectedPaperVOList[i];
              let mes;
              if (item.ProblemType === 0) {
                mes = "图片不清";
              }else if(item.ProblemType === 1) {
                mes = "答错题目";
              }else{
                mes = item.ProblemMes;
              }
              tableData.push({
                TestId: item.TestId,
                ExaminerId: item.ExaminerId,
                ExaminerName: item.ExaminerName,
                ProblemType: mes,
              });
            }
            this.setState({
              tableData,
              count: res.data.data.count,
            });
          }
        })
        .catch((e) => {
          console.log(e);
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
      this.setState({
        questionIndex: index,
      });
      this.tableData(this.state.questionList[index].QuestionId);
    }
    // 选择区

    paperMark =() => {
      this.props.history.push("/home/group/markTasks/2/" + this.state.questionList[this.state.questionIndex].QuestionId);
    }
    render() {
      return (
        <DocumentTitle title="阅卷系统-问题卷">
          <div className="group-problem-page" data-component="group-problem-page">
            <div className="search-container">
              <div className="question-select">
                            题目选择：<Select
                  showSearch
                  style={{width: 120}}
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
              <div className="paper-num">
                            问题卷数：{this.state.count}
              </div>
              <div className="paper-mark" onClick={() => {this.paperMark();}}>
                            评阅所有问题卷
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
