import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Button, Input, Select} from "antd";
import * as Settings from "../../../../Setting";
import "./index.less";
import Manage from "../../../../api/manage";
const {Option} = Select;

export default class index extends Component {

    adminId = "1"

    state = {
      subjectList: [],  // 获取学科列表
      questionList: [], // 获取对应学科的大题列表
      allotNum: [],      // 记录对应大题的人数
      headmanNum: 0, // 组长数量
      subjectValue: undefined,
      loading: false,
    }

    componentDidMount() {
      this.subjectList();
    }

    subjectList = () => {
      Manage.subjectList({adminId: this.adminId})
        .then((res) => {
          if (res.data.status === "10000") {
            this.setState({
              subjectList: res.data.data.subjectVOList,
            });
          }
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }
    getSubjectOption = () => {
      let subjectOption;
      if (this.state.subjectList.length) {
        subjectOption = this.state.subjectList.map(item => {
          return <Option key={item.SubjectId} value={item.SubjectName}>{item.SubjectName}</Option>;
        });
      } else {
        return null;
      }
      return subjectOption;
    }

    subjectSelect = (e) => {
      this.setState({
        subjectValue: e,
      });
      Manage.questionList({adminId: this.adminId, subjectName: e})
        .then((res) => {
          if (res.data.status === "10000") {
            this.setState({
              questionList: res.data.data.questionsList,
              allotNum: new Array(res.data.data.questionsList.length).fill(0),
            });
          }
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }

    getQuestionAllot = () => {
      let subjectOption;

      if (this.state.questionList.length) {
        subjectOption = this.state.questionList.map((item, index) => {
          return <div className="setting-box" key={index}>
            {item.QuestionName} ：
            <Input style={{width: 200}} placeholder="请输入本题需要的人数"
              onChange={e => {
                let nowAllot = this.state.allotNum;
                nowAllot[index] = e.target.value;
                this.setState({
                  allotNum: nowAllot,
                });
              }}
            />
          </div>;
        });
      } else {
        return null;
      }
      return subjectOption;
    }

    // goToDetail = () => {
    //   if (this.state.subjectValue) {
    //     this.props.history.push({pathname: "/home/userManagement/detailTable", query: {subjectName: this.state.subjectValue}});
    //   }else {
    //     message.warning("请先选择科目！");
    //   }
    // }

    sendForm = () => {
      let listNum = this.state.questionList.length;
      let list = [];
      for(let i = 0;i < listNum;i++) {
        let item = {
          id: this.state.questionList[i].QuestionId,
          num: Number(this.state.allotNum[i]),
        };
        list.push(item);
      }

      let subjectName = this.state.subjectValue;
      const data = {
        subject_name: subjectName,
        supervisor_number: Number(this.state.headmanNum),
        list: list,
      };
      Manage.subjectAllot(data);
    }

    render() {
      return (
        <DocumentTitle title="试卷管理-试卷分配">
          <div className="allot-page" data-component="allot-page">
            <div className="subject-setting">
              <div className="setting-header">学科设置</div>
              <div className="setting-box">
                <div className="setting-input">
                  <div className="setting-item">
                                    科目选择：<Select
                      style={{width: 120}}
                      placeholder="请选择科目"
                      onSelect={(e) => {this.subjectSelect(e);}}
                      value={this.state.subjectValue}
                    >
                      {this.getSubjectOption()}
                    </Select>
                  </div>
                  <div className="setting-item">
                                    组长人数：
                    <Input style={{width: 200}}
                      placeholder="请输入本学科需要的组长数"
                      onChange={e => {
                        this.setState({
                          headmanNum: e.target.value,
                        });
                      }}
                    />
                  </div>
                </div>
              </div>
            </div>
            <div className="question-setting">
              <div className="setting-header">人数分配</div>
              {this.getQuestionAllot()}
            </div>
            <Button type="primary" onClick={() => {this.sendForm();}} loading={this.state.loading}>确认</Button>
            {/* <Button type="default" style={{marginLeft: "20px"}} onClick={() => {this.goToDetail();}}>查看详情</Button> */}
          </div>

        </DocumentTitle>
      );
    }

}
