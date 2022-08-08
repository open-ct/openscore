import React, {Component} from "react";
import DocumentTitle from "react-document-title";
import {Table} from "antd";
import {ArrowLeftOutlined} from "@ant-design/icons";
import "./index.less";
import Manage from "../../../../api/manage";
export default class index extends Component {
    adminId = "1"
    state = {
      title: "",
      subjectName: "",
      detailList: [],
    }

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
          ];

          let data;
          data = this.state.detailList[index].SubTopicVOList;
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
          </div>
        </DocumentTitle>
      );
    }

}
