import React, {Component} from "react";
import "./index.less";
import {Button, Col, Form, Input, Modal, Popconfirm, Row, Table} from "antd";
import * as Settings from "../../../../Setting";
import Manage from "../../../../api/manage";

export default class index extends Component {
    state = {
    }
    componentDidMount() {
      this.GetUsersList();
    }

    constructor(props) {
      super(props);
      this.state = {
        classes: props,
        users: null,
        subjectList: [],
        organizationName: props.match.params.organizationName,
        form_status: null,
        visible: false,
      };
    }
    formRef = React.createRef();
    UNSAFE_componentWillMount() {
    }
    addUser = () => {
      this.setState({
        visible: true,
        form_status: "add",
      });
    }

    //  取消按钮的点击事
    handleCancel = () => {
      // 点击取消按钮触发的事件
      this.setState({
        visible: false,
      });
    }
    GetUsersList = () => {
      Manage.listUsers({})
        .then((res) => {
          if (res.data.status === "ok" && res.data.data) {
            this.setState({
              users: res.data.data,
            });
          } else {
            this.setState({
              users: [],
            });
          }
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
      Manage.subjectList().then((res) => {
        let subject_namet_list = [];
        res.data.data.subjectVOList.forEach(el => {
          subject_namet_list.push({"text": el.SubjectName, "value": el.SubjectName});
        });
        this.setState({subjectList: subject_namet_list});
      })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }

    submit(data) {
      if (this.state.form_status === "add") {
        Manage.createUser(data)
          .then((res) => {
            if (res.data.status === "ok") {
              this.setState({
                users: res.data.data,
                visible: false,
              });
              this.formRef.current.resetFields();
              this.GetUsersList();
            }
          })
          .catch((e) => {
            Settings.showMessage("error", e);
          });
      } else if (this.state.form_status === "edit") {
        Manage.updateUser(data)
          .then((res) => {
            if (res.data.status === "ok") {
              this.setState({
                users: res.data.data,
                visible: false,
              });
              this.formRef.current.resetFields();
              this.GetUsersList();
            }
          })
          .catch((e) => {
            Settings.showMessage("error", e);
          });
      }
    }

    deleteUser(data) {
      Manage.deleteUser({"account": data})
        .then((res) => {
          if (res.data.status === "ok") {
            this.setState({
              users: res.data.data,
              visible: false,
            });
            this.GetUsersList();
          }
        })
        .catch((e) => {
          Settings.showMessage("error", e);
        });
    }
    editUser(data) {
      this.setState({
        visible: true,
        form_status: "edit",
      });
      setTimeout(() => {
        this.formRef.current.setFieldsValue({
          account: data.account,
          password: data.password,
          user_type: parseInt(data.user_type),
          subject_name: data.subject_name,
        });
      });
    }
    handleTempRead(record) {
      Manage.updateUser({...record, is_attempt: !record.is_attempt}).then((res) => {
        location.reload();
      });
    }
    renderTable(users) {
      const columns = [
        {
          title: "学科",
          dataIndex: "subject_name",
          key: "subject_name",
          width: "100px",
          filters: this.state.subjectList,
          filterSearch: true,
          onFilter: (value, record) => record.subject_name.startsWith(value),
          sorter: (a, b) => a.subject_name.localeCompare(b.subject_name),
        },
        {
          title: "用户类型",
          dataIndex: "user_type",
          key: "user_type",
          width: "100px",
          filters: [
            {
              text: "阅卷员",
              value: "normal",
            },
            {
              text: "组长",
              value: "supervisor",
            },
          ],
          filterSearch: true,
          onFilter: (value, record) => record.user_type.startsWith(value),
          sorter: (a, b) => a.user_type.localeCompare(b.user_type),
          render: (text) => {
            if (text === "normal") {
              return "阅卷员";
            } else if (text === "supervisor") {
              return "组长";
            } else {
              return "用户类型错误";
            }
          },
        },
        {
          title: "用户姓名",
          dataIndex: "account",
          key: "account",
          width: "100px",
          sorter: (a, b) => a.account.localeCompare(b.account),
        },
        {
          title: "密码",
          dataIndex: "password",
          key: "password",
          width: "100px",
          sorter: (a, b) => a.password.localeCompare(b.password),
        },
        {
          title: "操作",
          dataIndex: "",
          key: "op",
          width: "110px",
          fixed: "right",
          render: (text, record) => {
            return (
              <div>
                <Button style={{marginTop: "10px", marginBottom: "10px", marginRight: "10px"}} type="primary"
                  onClick={() => this.editUser(record)}>{("编辑")}</Button>
                <Popconfirm
                  title={`Sure to delete user: ${record.account} ?`}
                  onConfirm={() => this.deleteUser(record.account)}
                >
                  <Button style={{marginBottom: "10px", marginRight: "10px"}} type="danger">{"删除"}</Button>
                </Popconfirm>
                <Popconfirm
                  title={record.is_attempt ? "确认取消他的试评资格吗" : "确认要让他参与试评吗"}
                  onConfirm={() => {this.handleTempRead(record);}}
                >
                  <Button style={{marginBottom: "10px"}} type={record.is_attempt ? "" : "primary"}>
                    {record.is_attempt ? "试评中" : "试评"}
                  </Button>
                </Popconfirm>
              </div>
            );
          },
        },
      ];
      const rowSelection = {
        onChange: (selectedRowKeys, selectedRows) => {
          console.log(`selectedRowKeys: ${selectedRowKeys}`, "selectedRows: ", selectedRows);
        },
        getCheckboxProps: (record) => ({
          disabled: record.name === "Disabled User",
          // Column configuration not to be checked
          name: record.name,
        }),
      };
      return (
        <div>
          <Table columns={columns} dataSource={users} rowKey="user_id" size="middle" bordered pagination={{pageSize: 100}}
            title={() => (
              <div>
                {("用户管理")}&nbsp;&nbsp;&nbsp;&nbsp;
              </div>
            )}
            rowSelection={{
              type: "checkbox",
              ...rowSelection,
            }}
            loading={users === null}
          />
        </div>
      );
    }

    render() {
      return (
        <div>
          <Row style={{width: "100%"}}>
            <Col span={1}>
            </Col>
            <Col span={22}>
              {
                this.renderTable(this.state.users)
              }
            </Col>
            <Col span={1}>
            </Col>
          </Row>
          <Modal title={this.state.form_status === "add" ? "添加用户" : "修改用户"}
            visible={this.state.visible}   // 设置默认隐藏
            onCancel={this.handleCancel}  // 点击取消按钮，对话框消失
            footer={false}   // 对话框的底部
          >
            <Form
              onFinish={this.submit.bind(this)}
              className="invitecode"
              ref={this.formRef}>
              <Form.Item label="用户名称" name="account" >
                <Input disabled /></Form.Item>
              <Form.Item label="用户密码" name="password">
                <Input /></Form.Item>
              <Form.Item label="用户类型" name="user_type" type="number">
                <Input /></Form.Item>
              <Form.Item label="学科" name="subject_name" >
                <Input disabled /></Form.Item>
              <Form.Item >
                <Button type="primary" htmlType="submit"> 保存 </Button>
              </Form.Item>
            </Form>
          </Modal>
        </div>
      );
    }
}
