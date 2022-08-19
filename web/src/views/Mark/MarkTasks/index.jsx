import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Button, Input, Modal, Radio, Select, Space, message} from "antd";
import {ExclamationCircleOutlined} from "@ant-design/icons";
import * as Settings from "../../../Setting";
import "./index.less";
import * as Util from "../../../util/Util";
import Marking from "../../../api/marking";
const {Option} = Select;

export default class index extends Component {

  userId = 1

  state = {
    problemVisible: false,
    problemValue: 1,
    inpValu: "",
    papers: [],
    currentPaper: {},
    testLength: 0,
    selectId: [],
    selectScore: [],
    subTopic: [],
    markScore: [],
  };

  componentDidMount() {
    if(localStorage.getItem("account") === "1" || localStorage.getItem("account") === "2") {
      this.getAllPaper();
    }
  }

  // 总试卷获取 
  getAllPaper = () => {
    Marking.testList({userId: this.userId})
      .then((res) => {
        if (res.data.status === "ok") {
          let papers = [...res.data.data.TestIds];
          this.setState(
            {
              papers,
            }
          );
          this.getCurrentPaper();
        }else if(res.data.status === "10003") {
          if (res.data.msg === "there is no paper to correct") {
            // let {user_type} = JSON.parse(this.state.account.tag)
            // if(user_type==="2")
            message.warning("没有试卷待批改");
          }
        }
      })
      .catch((e) => {
        Settings.showMessage("error", e);
      });
  }

  // 当前试卷
  getCurrentPaper = () => {
    Marking.testDisplay({userId: this.userId, testId: this.state.papers[0]})
      .then((res) => {
        if (res.data.status === "10000") {
          let currentPaper = res.data.data;
          let subTopic = res.data.data.subTopic;
          let testLength = res.data.data.subTopic.length;
          let markScore = [];
          for (let i = 0; i < subTopic.length; i++) {
            markScore.push(subTopic[i].score_type.split("-"));
          }
          this.setState({
            currentPaper,
            subTopic,
            markScore,
            testLength,
          });
        }
      })
      .catch((e) => {
        Settings.showMessage("error", e);
      });
  }

  // 打分展示
  imgScore = (item) => {
    let index;
    for (let i = 0; i < this.state.selectId.length; i++) {
      if (item === this.state.selectId[i]) {
        index = i;
      }
    }
    return this.state.selectScore[index];
  }

  // 阅卷区
  showTest = () => {
    let testPaper = null;
    if (this.state.currentPaper.testInfos !== undefined) {
      testPaper = this.state.currentPaper.testInfos.map((item, index) => {
        return <div className="test-question-img" key={index} data-content-before={this.imgScore(item.test_detail_id)}>
          <img src={item.pic_src} alt="加载失败" />
        </div>;
      });
    }
    return testPaper;
  }

  render() {
    return (
      <DocumentTitle title="阅卷系统-评卷">
        <div className="mark-tasks-page" data-component="mark-tasks-page">
          <div className="mark-paper">
            {
              this.showTest()
            }
          </div>
          <div className="mark-score">
            {
              this.renderScoreDropDown()
            }
            {
              this.renderPush()
            }
          </div>
        </div>
      </DocumentTitle>
    );
  }

  // 评分区
  selectBox = (index) => {
    let selectList;
    if (this.state.markScore.length !== 0) {
      selectList = this.state.markScore[index].map((item, i) => {
        return <Option key={i} value={item} label={item}>{item}</Option>;
      });
    } else {
      return null;
    }
    return selectList;
  }

  showSelect = () => {
    let scoreSelect = null;
    if (this.state.currentPaper.testInfos !== undefined) {
      scoreSelect = this.state.currentPaper.subTopic.map((item, index) => {
        return <div key={index} className="score-select">
          {item.question_detail_name}：<Select
            showSearch
            key={index}
            placeholder="请选择分数"
            style={{width: 120}}
            onSelect={this.select.bind(this, item.test_detail_id)}
            optionFilterProp="label"
            filterOption={(input, option) =>
              option.label.indexOf(input) >= 0
            }
            filterSort={(optionA, optionB) =>
              optionA.label.localeCompare(optionB.label)
            }
            autoFocus={index === 0}
          >
            {
              this.selectBox(index)
            // this.selectBox(item.question_detail_score)
            }
          </Select>
        </div>;
      });
    }

    return scoreSelect;
  }
  select = (item, value) => {
    if (this.state.selectId.length < this.state.testLength) {
      this.setState({
        selectId: [...this.state.selectId, item],
        selectScore: [...this.state.selectScore, value],
      });
    } else {
      let reviseSelectNo = 0;
      let newSelectScore = [];
      for (let i = 0; i < this.state.selectId.length; i++) {
        if (this.state.selectId[i] === item) {
          reviseSelectNo = i;
        }
      }
      this.state.selectScore.map((e, index) => {
        if (index === reviseSelectNo) {
          newSelectScore.push(value);
        } else {
          newSelectScore.push(e);
        }
      });
      this.setState({
        selectScore: newSelectScore,
      });
    }
  }
  renderScoreDropDown() {

    return (
      <div className="score-container">
        {
          this.showSelect()
        }
      </div>
    );
  }

  renderPush() {

    return (
      <div className="push-container">
        <div>
          <Button type="primary" style={{width: 60}} className="push-submit" onClick={() => {
            this.showWarning(1);
          }}>提交</Button>
        </div>
        <div className="push-paper" >
          <Button className="push-problem" style={{width: 60}} onClick={() => {
            this.showWarning(2);
          }}>问题卷</Button>
          <Button className="push-excellent" style={{width: 60}} onClick={() => {
            this.showWarning(3);
          }}>优秀卷</Button>
        </div>
        {
          this.problemModal()
        }
      </div>
    );
  }

  showWarning = (value) => {
    let title = "";
    let okContent = "提交后系统将记录该试卷";
    switch (value) {
    case 1:
      title = "请确认是否提交该试卷";
      break;
    case 2:
      title = "请确认是否提交该问题卷";
      break;
    case 3:
      title = "请确认是否提交该优秀卷";
      break;
    }

    if (value === 1) {
      let selectOrdered = [];
      for (let i = 0; i < this.state.subTopic.length; i++) {
        for (let j = 0; j < this.state.selectId.length; j++) {
          if (this.state.selectId[j] === this.state.subTopic[i].test_detail_id) {
            selectOrdered.push(this.state.selectScore[j]);
          }
        }
      }
      let content = "该试卷评分记录为：";
      for (let i = 0; i < selectOrdered.length; i++) {
        content += selectOrdered[i] + "\n";
      }
      okContent = content + "提交后系统将记录该试卷";
    }
    Modal.confirm({
      title: title,
      icon: <ExclamationCircleOutlined />,
      content: okContent,
      okText: "确认",
      cancelText: "取消",
      onOk: () => {
        let Qustion_detail_id = Util.getTextByJs(this.state.selectId);
        let Question_detail_score = Util.getTextByJs(this.state.selectScore);
        if (value === 1) {
          if (this.state.selectScore.length < this.state.testLength) {
            message.warning("请将分数打全");
          } else {
            Marking.testPoint({
              userId: this.userId,
              testId: this.state.currentPaper.testId,
              scores: Question_detail_score,
              testDetailId: Qustion_detail_id,
            })
              .then((res) => {
                if (res.data.status === "10000") {
                  this.setState({
                    selectId: [],
                    selectScore: [],
                    currentPaper: {},
                  });
                  this.getAllPaper();
                  message.success("打分成功");
                }
              })
              .catch((e) => {
                Settings.showMessage("error", e);
              });
          }
        } else if (value === 3) {
          console.log("3");
        } else {
          console.log("2", this);
          this.setState({
            problemVisible: true,
          });
        }
      },
    });
  }

  handleOk = () => {
    if (this.state.problemValue === 3) {
      Marking.testProblem({
        userId: this.userId,
        testId: this.state.currentPaper.testId,
        problemType: this.state.problemValue,
        problemMessage: this.state.inpValu,
      })
        .then(() => {
          this.setState({
            selectId: [],
            selectScore: [],
            currentPaper: {},
          });
          this.getAllPaper();
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    } else {
      Marking.testProblem({
        userId: this.userId,
        testId: this.state.currentPaper.testId,
        problemType: this.state.problemValue,
      })
        .then(() => {
          this.setState({
            selectId: [],
            selectScore: [],
            currentPaper: {},
          });
          this.getAllPaper();
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }
    this.setState({
      problemVisible: false,
    });
  };

  handleCancel = () => {
    this.setState({
      problemVisible: false,
    });
  };
  onRidioChange = e => {
    this.setState({
      problemValue: e.target.value,
    });
  }
  handelChange(e) {
    this.setState({
      inpValu: e.target.value,
    });
  }
  problemModal() {
    const {problemValue} = this.state;
    return (
      <Modal autoFocus="true"
        title="请选择问题卷类型"
        visible={this.state.problemVisible}
        onOk={this.handleOk}
        onCancel={this.handleCancel}
      >
        <Radio.Group onChange={this.onRidioChange} value={problemValue}>
          <Space direction="vertical">
            <Radio value={1}>图像不清</Radio>
            <Radio value={2}>答错题目</Radio>
            <Radio value={3}>其他错误</Radio>
          </Space>
        </Radio.Group>
        <Input placeholder="请输入问题" onChange={this.handelChange.bind(this)} style={{marginTop: 10}} />
      </Modal>
    );
  }
}
