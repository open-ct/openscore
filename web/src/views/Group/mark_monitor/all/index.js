import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Progress, Select, Table} from "antd";
import "./index.less";
import group from "../../../../api/group";
const {Option} = Select;
export default class index extends Component {

    supervisorId = "1"

    state = {
      questionList: [],
      tableData: [],
    }

    columns = [
      {
        title: "基本情况",
        children: [
          {
            title: "题号",
            width: 90,
            dataIndex: "QuestionId",
            fixed: "left",
          },
          {
            title: "题名",
            width: 90,
            dataIndex: "QuestionName",
            fixed: "left",
          },
          {
            title: "任务总量",
            width: 90,
            dataIndex: "ImportNumber",
            fixed: "left",
          },
          {
            title: "分配人数",
            width: 90,
            dataIndex: "DistributionUserNumber",
          },
          {
            title: "完成量",
            width: 90,
            dataIndex: "FinishNumber",
          },
          {
            title: "完成率",
            width: 90,
            dataIndex: "FinishRate",
            render: (text) => {
              return (<Progress type="circle" percent={text * 100} width={40} />);
            },
          },
          {
            title: "状态",
            width: 90,
            dataIndex: "IsAllFinished",
          },
        ],
      },
      {
        title: "进度预测（正评）",
        children: [
          {
            title: "平均分",
            width: 90,
            dataIndex: "AverageScore",
          },
          {
            title: "在线平均分",
            width: 90,
            dataIndex: "OnlineAverageScore",
          },
          {
            title: "预计时间",
            width: 90,
            dataIndex: "PredictTime",
          },
          {
            title: "在线预计时间",
            width: 90,
            dataIndex: "OnlinePredictTime",
          },
          {
            title: "平均速度",
            width: 90,
            dataIndex: "AverageSpeed",
          },
          {
            title: "在线平均速度",
            width: 90,
            dataIndex: "OnlineAverageSpeed",
          },
          {
            title: "自评指数",
            width: 90,
            dataIndex: "SelfScoreRate",
          },
        ],
      },
      {
        title: "一评情况",
        children: [
          {
            title: "一评完成量",
            width: 90,
            dataIndex: "FirstFinishedNumber",
          },
          {
            title: "一评完成率",
            width: 90,
            dataIndex: "FirstFinishedRate",
            render: (text) => {
              return (<Progress type="circle" percent={(text * 100).toFixed(1)} width={40} />);
            },
          },
          // {
          //     title: '一评未完成量',
          //     width: 90,
          //     dataIndex: 'FirstUnfinishedNumber',
          // },
          // {
          //     title: '一评未完成率',
          //     width: 90,
          //     dataIndex: 'FirstUnfinishedRate',
          // },
          {
            title: "状态",
            width: 90,
            dataIndex: "IsFirstFinished",
          },
        ],
      },
      {
        title: "二评情况",
        children: [
          {
            title: "二评完成量",
            width: 90,
            dataIndex: "SecondFinishedNumber",
          },
          {
            title: "二评完成率",
            width: 90,
            dataIndex: "SecondFinishedRate",
            render: (text) => {
              return (<Progress type="circle" percent={(text * 100).toFixed(1)} width={40} />);
            },
          },
          // {
          //     title: '二评未完成量',
          //     width: 90,
          //     dataIndex: 'SecondUnfinishedNumber',
          // },
          // {
          //     title: '二评未完成率',
          //     width: 90,
          //     dataIndex: 'SecondUnfinishedRate',
          // },
          {
            title: "状态",
            width: 90,
            dataIndex: "IsSecondFinished",
          },
        ],
      },
      {
        title: "三评情况",
        children: [
          {
            title: "三评完成量",
            width: 90,
            dataIndex: "ThirdFinishedNumber",
          },
          {
            title: "三评完成率",
            width: 90,
            dataIndex: "ThirdFinishedRate",
            render: (text) => {
              return (<Progress type="circle" percent={(text * 100).toFixed(1)} width={40} />);
            },
          },
          // {
          //     title: '三评未完成量',
          //     width: 90,
          //     dataIndex: 'ThirdUnfinishedNumber',
          // },
          // {
          //     title: '三评未完成率',
          //     width: 90,
          //     dataIndex: 'ThirdUnfinishedRate',
          // },
          {
            title: "状态",
            width: 90,
            dataIndex: "IsThirdFinished",
          },
        ],
      },
      {
        title: "仲裁卷",
        children: [
          {
            title: "产生量",
            width: 90,
            dataIndex: "ArbitramentNumber",
          },
          {
            title: "产生率",
            width: 90,
            dataIndex: "ArbitramentRate",
          },
          {
            title: "完成量",
            width: 90,
            dataIndex: "ArbitramentFinishedNumber",
          },
          {
            title: "完成率",
            width: 90,
            dataIndex: "ArbitramentFinishedRate",
            render: (text) => {
              return (<Progress type="circle" percent={(text * 100).toFixed(1)} width={40} />);
            },
          },
          // {
          //     title: '未完成量',
          //     width: 90,
          //     dataIndex: 'ArbitramentUnfinishedNumber',
          // },
          // {
          //     title: '未完成率',
          //     width: 90,
          //     dataIndex: 'ArbitramentUnfinishedRate',
          // },
          {
            title: "状态",
            width: 90,
            dataIndex: "IsArbitramentFinished",
          },

        ],
      },
      {
        title: "问题卷（回收已分配问题卷）",
        children: [
          {
            title: "产生量",
            width: 90,
            dataIndex: "ProblemNumber",
          },
          {
            title: "产生率",
            width: 90,
            dataIndex: "ProblemRate",
          },
          {
            title: "完成量",
            width: 90,
            dataIndex: "ProblemFinishedNumber",
          },
          {
            title: "完成率",
            width: 90,
            dataIndex: "ProblemUnfinishedRate",
            render: (text) => {
              return (<Progress type="circle" percent={(text * 100).toFixed(1)} width={40} />);
            },
          },
          // {
          //     title: '未完成量',
          //     width: 90,
          //     dataIndex: 'ProblemUnfinishedNumber',
          // },
          // {
          //     title: '未完成率',
          //     width: 90,
          //     dataIndex: 'ArbitramentUnfinishedRate',
          // },
          {
            title: "状态",
            width: 90,
            dataIndex: "IsProblemFinished",
          },
        ],
      },
    ]
    questionList = () => {
      group.questionList({adminId: "1", subjectName: JSON.parse(localStorage.getItem("userInfo")).SubjectName})
        .then((res) => {
          if (res.data.status === "10000") {
            this.setState({
              questionList: res.data.data.questionsList,
            });
            console.log(res.data.data.questionsList);
            // this.tableData(res.data.data.questionsList[0].QuestionId)
          }
        })
        .catch((e) => {
          console.log(e);
        });
    }

    tableData = () => {
      group.allMonitor({supervisorId: "2", subject: JSON.parse(localStorage.getItem("userInfo")).SubjectName})
        .then((res) => {
          if (res.data.status === "10000") {
            let tableData = [];
            for (let i = 0; i < res.data.data.scoreProgressVOList.length; i++) {
              let item = res.data.data.scoreProgressVOList[i];
              tableData.push(item);
            }
            this.setState({
              tableData,
            });
          }
        })
        .catch((e) => {
          console.log(e);
        });
    }
    componentDidMount() {
      this.questionList();
      this.tableData();
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
      console.log(selectList);
      return selectList;
    }
    select = (e) => {
      let index;
      for (let i = 0; i < this.state.questionList.length; i++) {
        if (this.state.questionList[i].QuestionName === e) {
          // eslint-disable-next-line unused-imports/no-unused-vars
          index = i;
        }
      }
      // this.tableData(this.state.questionList[index].QuestionId)
    }

    render() {
      return (
        <DocumentTitle title="阅卷系统-总体进度">
          <div className="all-monitor-page" data-component="all-monitor-page">
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
            </div>
            <div className="display-container">
              <Table
                pagination={{position: ["bottomCenter"]}}
                columns={this.columns}
                dataSource={this.state.tableData}
                scroll={{x: 5000}}
              />
            </div>
          </div>
        </DocumentTitle>
      );
    }

}
