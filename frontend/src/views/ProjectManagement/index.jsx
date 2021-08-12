import React, { Component } from 'react'
import DocumentTitle from 'react-document-title'
import { Layout } from 'antd';
import './index.less'

const { Header, Content } = Layout;

export default class index extends Component {
    render() {
        return (
            <DocumentTitle title="项目管理">
                <Layout className="project-management-page" data-component="project-management-page">
                    <Header>
                        <div className="header-container">
                            <span className="title">Landing</span>
                            <div className="right-box">
                                
                            </div>
                        </div>
                    </Header>
                    <Content>Content</Content>
                </Layout>
            </DocumentTitle>
        )
    }
}
