import React from "react";
import * as Auth from "../../../auth/Auth";
export const Nav00DataSource = {
  wrapper: {className: "header0 home-page-wrapper"},
  page: {className: "home-page l680pjqpd6-editor_css"},
  logo: {
    className: "header0-logo",
    children: "",
  },
  Menu: {
    className: "header0-menu",
    children: [
      {
        name: "item0",
        className: "header0-item",
        children: {
          href: "#",
          children: [
            {
              children: (
                <span>
                  <span>
                    <p>OpenCT官网</p>
                  </span>
                </span>
              ),
              name: "text",
            },
          ],
        },
      },
      {
        name: "item1",
        className: "header0-item",
        children: {
          href: "#",
          children: [
            {
              children: (
                <span>
                  <span>
                    <p>常见问题</p>
                  </span>
                </span>
              ),
              name: "text",
            },
          ],
        },
      },
      {
        name: "item2",
        className: "header0-item",
        children: {
          href: "#",
          children: [
            {
              children: (
                <span>
                  <span>
                    <p>联系我们</p>
                  </span>
                </span>
              ),
              name: "text",
            },
          ],
        },
      },
    ],
  },
  mobileMenu: {className: "header0-mobile-menu"},
};
export const Content00DataSource = {
  wrapper: {
    className: "home-page-wrapper content0-wrapper l6804odqyna-editor_css",
  },
  page: {className: "home-page content0"},
  titleWrapper: {
    className: "title-wrapper",
    children: [
      {
        name: "title",
        children: (
          <span>
            <p>请选择你的角色登录</p>
          </span>
        ),
      },
    ],
  },
  childWrapper: {
    className: "content0-block-wrapper l6816tw8z8-editor_css",
    children: [
      {
        name: "block0",
        className: "content0-block",
        md: 8,
        xs: 24,
        children: {
          className: "content0-block-item",
          children: [
            {
              name: "image",
              className: "content0-block-icon",
              children:
                  "https://zos.alipayobjects.com/rmsportal/WBnVOjtIlGWbzyQivuyq.png",
            },
            {
              name: "title",
              className: "content0-block-title",
              children: (
                <span>
                  <span>
                    <p>管理员</p>
                  </span>
                </span>
              ),
            },
            {
              name: "button",
              className: "",
              children: {
                children: (
                  <span>
                    <p>点击登录</p>
                  </span>
                ),
                href: Auth.getAuthorizeUrl(),
                type: "default",
                className: "l68186w9ejc-editor_css",
              },
            },
            {
              name: "content",
              children: (
                <span>
                  <span>
                    <span>
                      <span>
                        <p>管理阅卷人员</p>
                        <p>监控阅卷进度</p>
                        <p>导入导出试卷</p>
                      </span>
                    </span>
                  </span>
                </span>
              ),
            },
          ],
        },
      },
      {
        name: "block1",
        className: "content0-block",
        md: 8,
        xs: 24,
        children: {
          className: "content0-block-item",
          children: [
            {
              name: "image",
              className: "content0-block-icon",
              children:
                  "https://zos.alipayobjects.com/rmsportal/YPMsLQuCEXtuEkmXTTdk.png",
            },
            {
              name: "title",
              className: "content0-block-title",
              children: (
                <span>
                  <span>
                    <p>阅卷组长</p>
                  </span>
                </span>
              ),
            },
            {
              name: "button",
              className: "",
              children: {
                children: (
                  <span>
                    <p>点击登录</p>
                  </span>
                ),
                href: "/login",
                type: "default",
                className: "l68149n8dtp-editor_css",
              },
            },
            {
              name: "content",
              children: (
                <span>
                  <p>把控组内阅卷</p>
                  <p>评阅特别标注试卷</p>
                </span>
              ),
              className: "l680wca8pp6-editor_css",
            },
          ],
        },
      },
      {
        name: "block2",
        className: "content0-block",
        md: 8,
        xs: 24,
        children: {
          className: "content0-block-item",
          children: [
            {
              name: "image",
              className: "content0-block-icon",
              children:
                  "https://zos.alipayobjects.com/rmsportal/EkXWVvAaFJKCzhMmQYiX.png",
            },
            {
              name: "title",
              className: "content0-block-title",
              children: (
                <span>
                  <span>
                    <span>
                      <p>阅卷员</p>
                    </span>
                  </span>
                </span>
              ),
            },
            {
              name: "button",
              className: "",
              children: {
                children: (
                  <span>
                    <p>点击登录</p>
                  </span>
                ),
                href: "/login",
                type: "default",
                className: "l681j0unnah-editor_css",
              },
            },
            {
              name: "content",
              children: (
                <span>
                  <p>在线批改试卷</p>
                </span>
              ),
            },
          ],
        },
      },
    ],
  },
};
