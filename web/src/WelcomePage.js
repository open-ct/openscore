/*
 * @Author: Junlang
 * @Date: 2021-07-23 16:26:29
 * @LastEditTime: 2021-07-23 16:35:19
 * @LastEditors: Junlang
 * @FilePath: /openscore/web/src/WelcomePage.js
 */
import React from "react";
import {Button, Col, Modal, Rate, Row, Switch, Table, Tag, Tooltip} from 'antd';
import {CheckCircleOutlined, SyncOutlined, CloseCircleOutlined, ExclamationCircleOutlined, MinusCircleOutlined} from '@ant-design/icons';
import {getUserProfileUrl} from "./auth/Auth";
import * as Conf from "./Conf";
import moment from "moment";
import * as Setting from "./Setting";
import * as Auth from "./auth/Auth";


class WelcomePage extends React.Component {
  constructor(props) {
    super(props);
    const programName = props.match.params.programName !== undefined ? props.match.params.programName : Conf.DefaultProgramName;
    this.state = {
      programName: programName,
  
    };
  }

  

  componentWillMount() {
   
  }

  render() {
    return (
      <div>
        <Row style={{width: "100%"}}>
          <div>Hello OpenScore</div>
        </Row>
      </div>
    );
  }
}

export default WelcomePage;
