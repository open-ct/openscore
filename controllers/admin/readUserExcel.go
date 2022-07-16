package admin

// 导入用户
// func (c *AdminApiController) ReadUserExcel() {
// 	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
// 	defer c.ServeJSON()
//
// 	file, header, err := c.GetFile("excel")
//
// 	if err != nil {
// 		log.Println(err)
// 		c.Data["json"] = Response{Status: "10001", Msg: "cannot unmarshal", Data: err}
// 		return
// 	}
// 	tempFile, err := os.Create(header.Filename)
// 	io.Copy(tempFile, file)
// 	f, err := excelize.OpenFile(header.Filename)
// 	if err != nil {
// 		log.Println(err)
// 		c.Data["json"] = Response{Status: "30000", Msg: "excel 表导入错误", Data: err}
// 		return
// 	}
//
// 	// Get all the rows in the Sheet1.
// 	rows, err := f.GetRows("Sheet1")
// 	if err != nil {
// 		log.Println(err)
// 		c.Data["json"] = Response{Status: "30000", Msg: "excel 表导入错误", Data: err}
// 		return
// 	}
//
// 	for _, r := range rows[1:] {
// 		row := make([]string, len(rows[0]))
// 		copy(row, r)
// 		var user model.User
// 		user.UserName = row[0]
// 		user.ExaminerCount = row[1]
// 		user.Password = row[2]
// 		user.IdCard = row[3]
// 		user.Address = row[4]
// 		user.Tel = row[5]
// 		user.Email = row[6]
// 		userType, _ := strconv.Atoi(row[7])
// 		user.UserType = int64(userType)
// 		user.SubjectName = row[8]
// 		if err := user.Insert(); err != nil {
// 			log.Println(err)
// 			c.Data["json"] = Response{Status: "30001", Msg: "用户导入错误", Data: err}
// 			return
// 		}
//
// 	}
//
// 	err = tempFile.Close()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	err = os.Remove(header.Filename)
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	// ------------------------------------------------
// 	data := make(map[string]interface{})
// 	data["data"] = nil
// 	c.Data["json"] = Response{Status: "10000", Msg: "OK", Data: data}
//
// }
