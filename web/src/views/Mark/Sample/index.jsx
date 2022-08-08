import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Table} from "antd";
import "./index.less";
import Marking from "../../../api/marking";

export default class index extends Component {

  userId = "1"
  state = {
    papers: [],
    sampleList: [],
    samplePaper: [],
  };

  componentDidMount() {
    this.getAllPaper();
  }

  // 总试卷获取 
  getAllPaper = () => {
    Marking.testList({userId: this.userId})
      .then((res) => {
        if (res.data.status === "10000") {
          let papers = [...res.data.data.TestIds];
          this.setState(
            {
              papers,
            }
          );
          this.getSampleList();
        }
      })
      .catch((e) => {
        console.log(e);
      });
  }
  getSampleList = () => {
    Marking.testExampleList({userId: this.userId, testId: this.state.papers[0]})
      .then((res) => {
        if (res.data.status === "10000") {
          let sampleList = [];
          for (let i = 0; i < res.data.data.exampleTestPapers.length; i++) {
            sampleList.push({
              order: i,
              question_id: res.data.data.exampleTestPapers[i].test_id,
              question_name: res.data.data.exampleTestPapers[i].candidate,
              score: res.data.data.exampleTestPapers[i].final_score,
            });
          }
          this.setState({
            sampleList: sampleList,
          });
        }
      })
      .catch((e) => {
        console.log(e);
      });
  }

  columns = [
    {
      title: "序号",
      width: 90,
      dataIndex: "order",
    },
    {
      title: "题号",
      width: 90,
      dataIndex: "question_id",
    },
    {
      title: "题名",
      width: 90,
      dataIndex: "question_name",
    },
    {
      title: "分数",
      width: 90,
      dataIndex: "score",
    },

  ];

  showTest = () => {
    let testPaper = null;
    if (this.state.samplePaper !== undefined || this.state.samplePaper !== null) {
      testPaper = this.state.samplePaper.map((item) => {
        return <img src={"data:image/jpg;base64," + item.pic_src} alt="加载失败" className="test-question-img" />;
      });
    }

    return testPaper;
  }

  render() {

    return (
      <DocumentTitle title="阅卷系统-样卷">
        <div className="answer-tasks-page" data-component="answer-tasks-page">
          <div className="answer-paper">
            {
              this.showTest()
            }
          </div>
          <div className="answer-score">
            {
              this.sampleTable()
            }
          </div>
        </div>
      </DocumentTitle>
    );
  }
  sampleTable() {
    return (
      <Table columns={this.columns}
        dataSource={this.state.sampleList}
        scroll={{x: 400}}
        pagination={{position: ["bottomCenter"]}}
        onRow={(record) => ({
          onClick: () => {
            this.selectRow(record);
          },
        })}
      />
    );
  }
  selectRow = (record) => {
    console.log(record.order);
    Marking.testDetail({userId: this.userId, exampleTestId: record.question_id})
      .then((res) => {
        if (res.data.status == "10000") {
          this.setState({
            samplePaper: res.data.data.test[record.order],
          });
          console.log(this.state.samplePaper);
        }
      })
      .catch((e) => {
        console.log(e);
      });
  }

}
