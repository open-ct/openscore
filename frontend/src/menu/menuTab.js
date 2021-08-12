
export let permissionList = [
  {
    key: "mark-tasks",
    userPermission: "阅卷员",
    menu_name: "评卷",
    menu_url: "/home/mark-tasks",
    component_path: "../MarkTasks",
    icon: "FormOutlined",
    chidPermissions: [
    ]
  },
  {
    key: "answer",
    userPermission: "阅卷员",
    menu_name: "答案",
    menu_url: "/home/answer",
    component_path: "../Answer",
    icon: "CheckCircleOutlined",
    menu_id: 1,
    chidPermissions: [
    ]
  },
  {
    key: "review",
    userPermission: "阅卷员",
    menu_name: "样卷",
    menu_url: "/home/sample",
    component_path: "../Sample",
    icon: "ProfileOutlined",
    chidPermissions: [
    ]
  },
  {
    key: review,
    userPermission: "阅卷员",
    menu_name: "回评",
    menu_url: "/home/review",
    component_path: "../Review",
    icon: "HighlightOutlined",
    chidPermissions: [
    ]
  },
]
// 生成左边菜单
export function bindMenu(menulist) {
  let MenuList = menulist.map((item) => {
    if (item.chidPermissions.length === 0) {  //没有子菜单
      return <Menu.Item key={item.key} icon={React.createElement(Icon[item.icon])} onClick={() => this.add(item.menu_name, item.menu_url, item.menu_id, menu_imgClass)} ><Link to={item.menu_url}>{item.menu_name}</Link></Menu.Item>
    }
    else {
      return <SubMenu key={item.key} icon={<UserOutlined />} title={item.menu_name}>
        {this.bindMenu(item.chidPermissions)}
      </SubMenu>
    }

  })
  return MenuList
}
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
export function bindRouter(list) {
  let routerList = list.map((item) => {
    if (item.chidPermissions.length === 0) {
      return <Route key={item.key} path={item.menu_url} component={loadable(() => import(`./${item.component_path}`))}></Route>
    } else {
      return <Route key={item.key} path={item.menu_url} render={() => {
        let componentName = loadable(() => import(`./${item.component_path}`));
        return <componentName>
          {this.bindRouter(item.chidPermissions)}
        </componentName>
      }}>
      </Route>
    }
  })
  return routerList
}


