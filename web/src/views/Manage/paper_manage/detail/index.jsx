import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {ArrowLeftOutlined} from "@ant-design/icons";
import {Button, Form, Input, Modal, Popconfirm, Table} from "antd";
import "./index.less";
import Manage from "../../../../api/manage";
export default class index extends Component {
  adminId = "1"
  state = {
    title: "",
    subjectName: "",
    detailList: [],
    smallForm_visible: false,
    smallForm_status: null,
    form_visible: false,
  }
  smallForm = React.createRef();
  formRef = React.createRef();
  componentDidMount() {
    console.log(this.props.location.query);
    if (this.props.location.query) {
      this.setState({
        title: "试卷分配",
        subjectName: this.props.location.query.subjectName,
      });
      this.paperDetail();
    } else {
      this.setState({
        title: "题目设置",
      });
      this.questionDetail();
    }
  }
  questionDetail = () => {
    Manage.questionInfo({adminId: this.adminId})
      .then((res) => {
        if (res.data.status === "10000") {
          this.setState({
            detailList: res.data.data.topicVOList,
          });
        }
      })
      .catch((e) => {
        console.log(e);
      });
  }

  paperDetail = () => {
    Manage.paperInfo({adminId: this.adminId, subjectName: this.props.location.query.subjectName})
      .then((res) => {
        if (res.data.status === "10000") {
          this.setState({
            detailList: res.data.data.distributionRecordList,
          });
        }
      })
      .catch((e) => {
        console.log(e);
      });
  }
  addQuestionDetail(data) {
    this.setState({
      smallForm_visible: true,
      smallForm_status: "add",
    });
    setTimeout(() => {
      console.log(data);
      this.smallForm.current.setFieldsValue({
        question_detail_id: data.TopicId,
      });
    });
  }
  deleteQuestionDetail(data) {
    console.log("data", data);
    Manage.deleteSmallQuestion({"question_detail_id": data})
      .then((res) => {
        if (res.data.status === "ok") {
          this.setState({
            users: res.data.data,
            smallForm_visible: false,
          });
          this.questionDetail();
        }
      })
      .catch((e) => {
        console.log(e);
      });
  }
  editQuestionDetail(data) {
    this.setState({
      smallForm_visible: true,
      smallForm_status: "edit",
    });
    setTimeout(() => {
      this.smallForm.current.setFieldsValue({
        question_detail_id: data.SubTopicId,
        question_detail_name: data.SubTopicName,
        question_detail_score: parseInt(data.Score),
        score_type: data.ScoreDistribution,
      });
    });
  }
  smallForm_submit(data) {
    data.question_detail_score = parseInt(data.question_detail_score);
    data.question_id = data.question_detail_id;
    if (this.state.smallForm_status === "add") {
      Manage.createSmallQuestion(data)
        .then((res) => {
          if (res.data.status === "ok") {
            this.setState({
              users: res.data.data,
              smallForm_visible: false,
            });
            this.smallForm.current.resetFields();
            this.questionDetail();
          }
        })
        .catch((e) => {
          console.log(e);
        });
    } else if (this.state.smallForm_status === "edit") {
      Manage.updateSmallQuestion(data)
        .then((res) => {
          if (res.data.status === "ok") {
            this.setState({
              users: res.data.data,
              smallForm_visible: false,
            });
            this.smallForm.current.resetFields();
            this.questionDetail();
          }
        })
        .catch((e) => {
          console.log(e);
        });
    }
  }
  updateQuestion(data) {
    this.setState({
      form_visible: true,
      form_status: "edit",
    });
    setTimeout(() => {
      this.formRef.current.setFieldsValue({
        question_id: data.TopicId,
        question_name: data.TopicName,
        standard_error: data.StandardError,
        // question_score: parseInt(data.Score),
        question_score_type: parseInt(data.ScoreType),
      });
    });
  }
  deleteQuestion(data) {
    Manage.deleteQuestion({"question_id": data})
      .then((res) => {
        if (res.data.status === "ok") {
          this.setState({
            users: res.data.data,
            visible: false,
          });
          this.questionDetail();
        }
      })
      .catch((e) => {
        console.log(e);
      });
  }
  submit(data) {
    data.standard_error = parseInt(data.standard_error);
    // data.question_score = parseInt(data.question_score)
    data.score_type = parseInt(data.question_score_type);
    Manage.updateQuestion(data)
      .then((res) => {
        if (res.data.status === "ok") {
          this.setState({
            users: res.data.data,
            form_visible: false,
          });
          this.formRef.current.resetFields();
          this.questionDetail();
        }
      })
      .catch((e) => {
        console.log(e);
      });
  }
  //  取消按钮的点击事
  handleCancel = () => {
    // 点击取消按钮触发的事件
    this.setState({
      smallForm_visible: false,
    });
  }
  handleCancel2 = () => {
    // 点击取消按钮触发的事件
    this.setState({
      form_visible: false,
    });
  }
  detailTable = () => {
    let columns;
    if (this.state.subjectName) {
      columns = [
        {
          title: "ID",
          width: 90,
          dataIndex: "TopicId",
        },
        {
          title: "题号",
          width: 90,
          dataIndex: "TopicName",
        },
        {
          title: "试卷导入总数",
          width: 90,
          dataIndex: "ImportNumber",
        },
        {
          title: "分配试卷数",
          width: 90,
          dataIndex: "DistributionTestNumber",
        },
        {
          title: "分配人数",
          width: 90,
          dataIndex: "DistributionUserNumber",
        },

      ];
      return (
        <Table columns={columns}
          dataSource={this.state.detailList}
          pagination={{position: ["bottomCenter"]}}
        />
      );
    } else {
      const expandedRowRender = (record, index) => {
        const columns = [
          {
            title: "小题ID",
            width: 90,
            dataIndex: "SubTopicId",
          },
          {
            title: "小题名",
            width: 90,
            dataIndex: "SubTopicName",
          },
          {
            title: "小题满分",
            width: 90,
            dataIndex: "Score",
          },
          {
            title: "分数分布",
            width: 90,
            dataIndex: "ScoreDistribution",
          },
          {
            title: "操作",
            dataIndex: "",
            key: "op",
            width: "10px",
            fixed: "right",
            render: (text, record) => {
              return (
                <div>
                  <Button style={{marginTop: "10px", marginBottom: "10px", marginRight: "10px"}} type="primary"
                    onClick={() => this.editQuestionDetail(record)}>{("编辑")}</Button>
                  <Popconfirm
                    title={"Sure delete ?"}
                    onConfirm={() => this.deleteQuestionDetail(record.SubTopicId)}
                  >
                    <Button style={{marginBottom: "10px"}} type="danger">{"删除"}</Button>
                  </Popconfirm>
                </div>
              );
            },
          },
        ];

        let data ;
        data = this.state.detailList[index].SubTopicVOList;
        data.map((el) => {
          el.question_id = this.state.detailList[index].TopicId;
        });
        return <Table columns={columns} dataSource={data} pagination={false} />;
      };
      columns = [
        {
          title: "科目",
          width: 90,
          dataIndex: "SubjectName",
        },
        {
          title: "大题ID",
          width: 90,
          dataIndex: "TopicId",
        },
        {
          title: "大题名",
          width: 90,
          dataIndex: "TopicName",
        },
        {
          title: "满分",
          width: 90,
          dataIndex: "Score",
        },
        {
          title: "标准误差",
          width: 90,
          dataIndex: "StandardError",
        },
        {
          title: "是否二次阅卷",
          width: 90,
          dataIndex: "ScoreType",
        },
        {
          title: "添加时间",
          width: 90,
          dataIndex: "ImportTime",
        },
        {
          title: "操作",
          dataIndex: "",
          key: "op",
          width: "10px",
          fixed: "right",
          render: (text, record) => {
            return (
              <div>
                <Button style={{marginTop: "10px", marginBottom: "10px", marginRight: "10px"}} type="primary"
                  onClick={() => this.addQuestionDetail(record)}>{("新增小题")}</Button>
                <Button style={{marginTop: "10px", marginBottom: "10px", marginRight: "10px"}} type="primary"
                  onClick={() => this.updateQuestion(record)}>{("编辑")}</Button>
                <Popconfirm
                  title={`Sure to delete user: ${record.name} ?`}
                  onConfirm={() => this.deleteQuestion(record.TopicId)}
                >
                  <Button style={{marginBottom: "10px"}} type="danger">{"删除"}</Button>
                </Popconfirm>
              </div>
            );
          },
        },

      ];
      return (
        <Table
          className="components-table-demo-nested"
          columns={columns}
          expandable={{expandedRowRender}}
          rowKey={record => record.TopicId}
          dataSource={this.state.detailList}
          pagination={{position: ["bottomCenter"]}}
        />
      );
    }
  }
  render() {

    return (
      <DocumentTitle title={this.state.title + "-详情"}>
        <div className="detail-page" data-component="detail-page">
          <div className="search-container">
            <div className="goBack" onClick={() => {window.history.back(-1);}}>
              <ArrowLeftOutlined />
              <span className="paperDetail">{this.state.title + "-详情"}</span>
              <span className="goBackPre">返回{this.state.title}</span>
            </div>
            {
              this.state.subjectName ? <div className="subjectName" style={{marginLeft: "56px"}}> 科目：{this.state.subjectName}</div> : null
            }

          </div>
          <div className="paperPage">
            {
              this.detailTable()
            }
          </div>
          {/* 小题弹窗 */}
          <Modal title={this.state.smallForm_status === "add" ? "添加小题" : "修改小题"}
            visible={this.state.smallForm_visible}   // 设置默认隐藏
            onCancel={this.handleCancel}  // 点击取消按钮，对话框消失
            footer={false}   // 对话框的底部
          >
            <Form
              onFinish={this.smallForm_submit.bind(this)}
              className="invitecode"
              ref={this.smallForm}>
              <Form.Item label="小题名称" name="question_detail_name">
                <Input /></Form.Item>
              <Form.Item label="小题总分" name="question_detail_score">
                <Input /></Form.Item>
              <Form.Item label="小题分数分布" name="score_type" type="number">
                <Input /></Form.Item>
              <Form.Item name="question_detail_id" type="number">
                <div></div>
              </Form.Item>
              <Form.Item >
                <Button type="primary" htmlType="smallForm_submit"> 保存 </Button>
              </Form.Item>
            </Form>
          </Modal>
          {/* 大题弹窗 */}
          <Modal title="修改大题"
            visible={this.state.form_visible}   // 设置默认隐藏
            onCancel={this.handleCancel2}  // 点击取消按钮，对话框消失
            footer={false}   // 对话框的底部
          >
            <Form
              onFinish={this.submit.bind(this)}
              className="invitecode"
              ref={this.formRef}>
              <Form.Item label="大题名称" name="question_name">
                <Input /></Form.Item>
              {/* < Form.Item label="总分" name="question_score">
                                <Input /></Form.Item> */}
              <Form.Item label="标准误差" name="standard_error" type="number">
                <Input /></Form.Item>
              <Form.Item label="是否二次阅卷" name="question_score_type" type="number">
                <Input /></Form.Item>
              <Form.Item name="question_id"><div></div></Form.Item>
              <Form.Item >
                <Button type="primary" htmlType="submit"> 保存 </Button>
              </Form.Item>
            </Form>
          </Modal>
        </div>
      </DocumentTitle>
    );
  }

}
