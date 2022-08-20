import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import * as Settings from "../../../Setting";
import "./index.less";
import Marking from "../../../api/marking";

export default class index extends Component {

  userId = "1"
  state = {
    papers: [],
    keyTest: [],
  };
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
          this.getAnswer();
        }
      })
      .catch((e) => {
        Settings.showMessage("error", e);
      });
  }

  getAnswer = () => {
    Marking.testAnswer({userId: this.userId, testId: this.state.papers[0]})
      .then((res) => {
        if (res.data.status === "10000") {
          this.setState({
            keyTest: res.data.data.Pics,
          });
        }
      })
      .catch((e) => {
        Settings.showMessage("error", e);
      });
  }
  componentDidMount() {
    this.getAllPaper();
  }

  // 答案区
  showTest = () => {
    let testPaper = null;
    if (this.state.keyTest !== undefined) {
      testPaper = this.state.keyTest.map((item, index) => {
        return <div key={index} className="test-question-img">
          <img src={"data:image/jpg;base64," + item} alt="加载失败" />
        </div>;
      });
    }
    return testPaper;
  }

  render() {
    return (
      <DocumentTitle title="阅卷系统-答案">
        <div className="answer-tasks-page" data-component="answer-tasks-page">
          <div className="answer-paper">
            {
              this.showTest()
            }
          </div>
          <div className="answer-score">

          </div>
        </div>
      </DocumentTitle>
    );
  }

}
