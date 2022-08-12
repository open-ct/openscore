export default [
  {
    key: "mark-tasks",
    userPermission: "阅卷员",
    menu_name: "评卷",
    menu_url: "/home/mark-tasks",
    icon: "FormOutlined",
    chidPermissions: [
    ],
  },
  {
    key: "answer",
    userPermission: "阅卷员",
    menu_name: "答案",
    menu_url: "/home/answer",
    icon: "CheckCircleOutlined",
    menu_id: 1,
    chidPermissions: [
    ],
  },
  {
    key: "sample",
    userPermission: "阅卷员",
    menu_name: "样卷",
    menu_url: "/home/sample",
    icon: "ProfileOutlined",
    chidPermissions: [
    ],
  },
  {
    key: "review",
    userPermission: "阅卷员",
    menu_name: "回评",
    menu_url: "/home/review",
    icon: "HighlightOutlined",
    chidPermissions: [
    ],
  },
  // {
  //     key: "selfMark",
  //     userPermission: "阅卷员",
  //     menu_name: "自评",
  //     menu_url: "/home/selfMark",
  //     icon: "SolutionOutlined",
  //     chidPermissions: [
  //     ]
  // },
  {
    key: "mark_monitor",
    userPermission: "组长",
    menu_name: "评卷监控",
    menu_url: "/home/markMonitor",
    icon: "FormOutlined",
    chidPermissions: [
      {
        key: "teacher",
        userPermission: "组长",
        menu_name: "教师监控",
        menu_url: "/home/markMonitor/teacher",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "score",
        userPermission: "组长",
        menu_name: "分值分布",
        menu_url: "/home/markMonitor/score",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "self",
        userPermission: "组长",
        menu_name: "自评监控",
        menu_url: "/home/markMonitor/self",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "average",
        userPermission: "组长",
        menu_name: "平均分监控",
        menu_url: "/home/markMonitor/average",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "standard",
        userPermission: "组长",
        menu_name: "标准差监控",
        menu_url: "/home/markMonitor/standard",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "all",
        userPermission: "组长",
        menu_name: "总进度监控",
        menu_url: "/home/markMonitor/all",
        icon: "",
        chidPermissions: [
        ],
      },
    ],
  },
  {
    key: "test_management",
    userPermission: "组长",
    menu_name: "试卷管理",
    menu_url: "/home/testMonitor",
    icon: "TableOutlined",
    chidPermissions: [
      {
        key: "arbitration",
        userPermission: "组长",
        menu_name: "仲裁卷",
        menu_url: "/home/group/arbitration",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "problem",
        userPermission: "组长",
        menu_name: "问题卷",
        menu_url: "/home/group/problem",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "marking",
        userPermission: "组长",
        menu_name: "自评卷",
        menu_url: "/home/group/marking",
        icon: "",
        chidPermissions: [
        ],
      },
    ],
  },
  {
    key: "group_user_management",
    userPermission: "组长",
    menu_name: "用户管理",
    menu_url: "/home/group/userMonitor",
    icon: "UserOutlined",
    chidPermissions: [
    ],
  },

  {
    key: "paper_management",
    userPermission: "管理员",
    menu_name: "试卷管理",
    menu_url: "/home/management/paper",
    icon: "FormOutlined",
    chidPermissions: [
      {
        key: "question",
        userPermission: "管理员",
        menu_name: "题目设置",
        menu_url: "/home/management/question",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "paper",
        userPermission: "管理员",
        menu_name: "试卷导入",
        menu_url: "/home/management/paper",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "paper_allot",
        userPermission: "管理员",
        menu_name: "试卷分配",
        menu_url: "/home/management/paper_allot",
        icon: "",
        chidPermissions: [
        ],
      },
      {
        key: "paper_manage",
        userPermission: "管理员",
        menu_name: "试卷管理",
        menu_url: "/home/management/paper_manage",
        icon: "",
        chidPermissions: [
        ],
      },
    ],
  },
  {
    key: "user_management",
    userPermission: "管理员",
    menu_name: "用户管理",
    menu_url: "/home/management/user",
    icon: "UserOutlined",
    chidPermissions: [
      // {
      //   key: "user_export",
      //   userPermission: "管理员",
      //   menu_name: "导入导出用户",
      //   menu_url: "/home/management/user/user_export",
      //   icon: "",
      //   chidPermissions: [
      //   ],
      // },
      {
        key: "user_manage",
        userPermission: "管理员",
        menu_name: "用户管理",
        menu_url: "/home/management/user_manage/user_manage",
        icon: "",
        chidPermissions: [
        ],
      },
    ],
  },
];
// 生成左边菜单
// export function bindMenu(menulist) {
//   let MenuList = menulist.map((item) => {
//     if (item.chidPermissions.length === 0) {  //没有子菜单
//       return <Menu.Item key={item.key} icon={React.createElement(Icon[item.icon])} onClick={() => this.add(item.menu_name, item.menu_url, item.menu_id, menu_imgClass)} ><Link to={item.menu_url}>{item.menu_name}</Link></Menu.Item>
//     }
//     else {
//       return <SubMenu key={item.key} icon={<UserOutlined />} title={item.menu_name}>
//         {this.bindMenu(item.chidPermissions)}
//       </SubMenu>
//     }

//   })
//   return MenuList
// }
// componentDidMount() {
//   console.log("will mount")
//   // console.log(JSON.parse(this.props.user.user.permissionList))
//   // let menuList = window.sessionStorage.getItem('user')?(JSON.parse(window.sessionStorage.getItem('user'))):[];
//   // console.log('类相关',typeof menuList);
//   let leftMenu = this.bindMenu(this.props.user.user.permissionList);
//   this.setState({
//     leftMenu:leftMenu
//   })
//   let routerList = this.bindRouter(this.props.user.user.permissionList)
//   console.log('routerList',routerList);
//   this.setState({
//     routerList:routerList
//   })
// }

// 动态生成路由
// export function bindRouter(list) {
//   let routerList = list.map((item) => {
//     if (item.chidPermissions.length === 0) {
//       return <Route key={item.key} path={item.menu_url} component={loadable(() => import(`./${item.component_path}`))}></Route>
//     } else {
//       return <Route key={item.key} path={item.menu_url} render={() => {
//         let componentName = loadable(() => import(`./${item.component_path}`));
//         return <componentName>
//           {this.bindRouter(item.chidPermissions)}
//         </componentName>
//       }}>
//       </Route>
//     }
//   })
//   return routerList
// }
