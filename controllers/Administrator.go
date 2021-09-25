package controllers

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/freetype"
	"github.com/xuri/excelize/v2"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"openscore/models"
	"openscore/requests"
	"os"
	"strconv"
	"strings"
	"time"

	"openscore/responses"
)

var (
	dpi      = flag.Float64("dpi", 200, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "font/simhei.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 12, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

/**
1.生成图片
 */
func   UploadPic(name string,text  []string)(src string) {

	flag.Parse()

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize the context.
	fg, bg := image.Black, image.White
	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	if *wonb {
		fg, bg = image.White, image.Black
		ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	}
	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	// Draw the guidelines.
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFixed(*size)>>6))
	for _, s := range text {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(*size * *spacing)
	}

	// Save that RGBA image to disk.
	name = name+".png"
	newPath:="./img/"+name

	outFile, err := os.Create(newPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}


	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Wrote out.png OK.")
	return name
}




/**
2.试卷导入
 */

func (c *AdminApiController) ReadExcel(){
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	defer c.ServeJSON()
	//var requestBody requests.ReadExcel
	var resp Response
	var  err error

	//err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	//if err!=nil {
	//	log.Println(err)
	//	resp = Response{"10001","cannot unmarshal",err}
	//	c.Data["json"] = resp
	//	return
	//}
	//supervisorId := requestBody.SupervisorId
	// filePath := requestBody.FilePath
	//bytes := requestBody.Excel
	//r := requests.ReadExcelBytes{}
	_, header, err := c.GetFile("excel")
	err = err
	if err != nil {
		log.Println(err)
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	//bytes := r.Excel

	//----------------------------------------------------
	//bytes, err := os.ReadFile(filePath)
	//
	//
	//file, err := os.Create("excelFile")
	//if err!=nil {
	//	log.Println(err)
	//	resp = Response{"30000","excel 表导入错误",err}
	//	c.Data["json"] = resp
	//	return
	//}
	//_, err = file.Write(bytes)
	//if err!=nil {
	//	log.Println(err)
	//	resp = Response{"30000","excel 表导入错误",err}
	//	c.Data["json"] = resp
	//	return
	//}

	f, err := excelize.OpenFile(header.Filename)
	if err != nil {
		log.Println(err)
		resp = Response{"30000","excel 表导入错误",err}
		c.Data["json"] = resp
		return
	}


	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet2")
	if err != nil {
		log.Println(err)
		resp = Response{"30000","excel 表导入错误",err}
		c.Data["json"] = resp
		return
	}
	

	for i:=1;i<len(rows);i++ {
		for j:=1;j<len(rows[i]);j++ {

			if i>=1&&j>=3 {
				//准备数据
			    testIdStr:=rows[i][0]
			    testId, _ := strconv.ParseInt(testIdStr, 10, 64)
			    questionIds := strings.Split(rows[0][j], "-")
			    questionIdStr:=questionIds[0]
			    questionId, _ := strconv.ParseInt(questionIdStr, 10, 64)
			    questionDetailIdStr:=questionIds[3]
				questionDetailId, _ := strconv.ParseInt(questionDetailIdStr, 10, 64)
				name:=rows[i][2]
				//填充数据
				var testPaperInfo  models.TestPaperInfo
				var testPaper models.TestPaper

				testPaperInfo.Question_detail_id=questionDetailId
				s:=rows[i][j]
				split := strings.Split(s, "\n")
				src := UploadPic(rows[i][0]+rows[0][j], split)
			    testPaperInfo.Pic_src=src
				//查看大题试卷是否已经导入
				has,err := testPaper.GetTestPaper(testId)
				if err!=nil {
					log.Println(err)
				}

				//导入大题试卷
				if !has {
					testPaper.Test_id=testId
					testPaper.Question_id=questionId
					testPaper.Candidate=name
					err = testPaper.Insert()
					if err != nil {
						log.Println(err)
						resp = Response{"30001","试卷大题导入错误",err}
						c.Data["json"] = resp
						return
					}
				}
				//导入小题试卷
				testPaperInfo.Test_id=testId
				err = testPaperInfo.Insert()
				if err != nil {
					log.Println(err)
					resp = Response{"30002","试卷小题导错误",err}
					c.Data["json"] = resp
					return
				}

			}

		}
		
	}
	//获取选项名 存导入试卷数
	for k:=3;k<len(rows[0]);k++ {
		questionIds := strings.Split(rows[0][k], "-")
		questionIdStr:=questionIds[0]
		questionId, _ := strconv.ParseInt(questionIdStr, 10, 64)
	  var topic models.Topic
		topic.Question_id=questionId
		topic.Import_number=int64(len(rows)-1)
		err = topic.Update()
		if err != nil {
			log.Println(err)
			resp = Response{"30003","大题导入试卷数更新错误",err}
			c.Data["json"] = resp
			return
		}
	}

	//err = file.Close()
	if err!=nil {
		log.Println(err)
	}
	err = os.Remove("excelFile")
	if err!=nil {
		log.Println(err)
	}
	//------------------------------------------------
	data := make(map[string]interface{})
	data["data"] =nil
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp


}


/**
3.大题列表
 */


func (c *AdminApiController) QuestionBySubList() {
	defer c.ServeJSON()
	var requestBody requests.QuestionBySubList
	var resp Response
	var  err error

	err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err!=nil {
		log.Println(err)
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	//supervisorId := requestBody.SupervisorId
	subjectName := requestBody.SubjectName
	//----------------------------------------------------
	//获取大题列表
	topics  := make([]models.Topic,0)
	err = models.FindTopicBySubNameList(&topics,subjectName)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30004","获取大题列表错误  ",err}
		c.Data["json"] = resp
		return
	}

	var questions = make([]responses.QuestionBySubListVO,len(topics))
	for i := 0; i < len(topics); i++ {

		questions[i].QuestionId=topics[i].Question_id
		questions[i].QuestionName=topics[i].Question_name

	}

	//----------------------------------------------------
	data := make(map[string]interface{})
	data["questionsList"] =questions
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}




/**
4.试卷参数导入
 */

func (c *AdminApiController) InsertTopic(){

		defer c.ServeJSON()
		var requestBody requests.AddTopic
		var resp Response
		var  err error

		err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
		if err!=nil {
			log.Println(err)
			resp = Response{"10001","cannot unmarshal",err}
			c.Data["json"] = resp
			return
		}
	//adminId := requestBody.AdminId
	topicName := requestBody.TopicName
	scoreType := requestBody.ScoreType
	score := requestBody.Score
	standardError := requestBody.Error
	subjectName := requestBody.SubjectName
	details := requestBody.TopicDetails

	//----------------------------------------------------
		//添加subject
		var subject models.Subject
		subject.SubjectName=subjectName
	    flag ,err:= models.GetSubjectBySubjectName(&subject,subjectName)
		if err!=nil {
		log.Println(err)
		}
	    subjectId := subject.SubjectId
	if !flag {
		err, subjectId = models.InsertSubject(&subject)
		if err!=nil {
			log.Println(err)
			resp  = Response{"30005","科目导入错误  ",err}
			c.Data["json"] = resp
			return
		}
	}
	//添加topic
		var topic models.Topic
		topic.Question_name=topicName
		topic.Score_type=scoreType
		topic.Question_score=score
		topic.Standard_error=standardError
		topic.Subject_name=subjectName
		topic.Import_time=time.Now()
		topic.Subject_Id=subjectId

		err, questionId:= models.InsertTopic(&topic)
		if err!=nil {
			log.Println(err)
			resp  = Response{"30006"," 大题参数导入错误 ",err}
			c.Data["json"] = resp
			return
		}


	var  addTopicVO responses.AddTopicVO
	var addTopicDetailVOList = make([]responses.AddTopicDetailVO,len(details))

		for i := 0; i < len(details); i++ {
			var subTopic models.SubTopic
			subTopic.Question_detail_name=details[i].TopicDetailName
			subTopic.Question_detail_score=details[i].DetailScore
			subTopic.Score_type=details[i].DetailScoreTypes
			subTopic.Question_id=questionId
			err,questionDetailId:= models.InsertSubTopic(&subTopic)
			if err!=nil {
				log.Println(err)
				resp  = Response{"30007","小题参数导入错误  ",err}
				c.Data["json"] = resp
				return
			}
			addTopicDetailVOList[i].QuestionDetailId=questionDetailId
		}
 		addTopicVO.QuestionId=questionId
 		addTopicVO.QuestionDetailIds=addTopicDetailVOList
		//----------------------------------------------------
		data := make(map[string]interface{})
		data["addTopicVO"] =addTopicVO
		resp = Response{"10000", "OK", data}
		c.Data["json"] = resp

}

/**
 5.科目选择
 */

func (c *AdminApiController) SubjectList(){

	defer c.ServeJSON()
	var requestBody requests.SubjectList
	var resp Response
	var  err error

	err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err!=nil {
		log.Println(err)
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	//supervisorId := requestBody.SupervisorId
	//----------------------------------------------------
	//获取科目列表
	subjects  := make([]models.Subject,0)
	err = models.FindSubjectList(&subjects)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30008","科目列表获取错误  ",err}
		c.Data["json"] = resp
		return
	}

	var subjectVOList = make([]responses.SubjectListVO,len(subjects))
	for i := 0; i < len(subjects); i++ {

		subjectVOList[i].SubjectName=subjects[i].SubjectName
		subjectVOList[i].SubjectId=subjects[i].SubjectId
	}

	//----------------------------------------------------
	data := make(map[string]interface{})
	data["subjectVOList"] =subjectVOList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}




/**
6.试卷分配界面
 */
func (c *AdminApiController) DistributionInfo(){

	defer c.ServeJSON()
	var requestBody requests.DistributionInfo
	var resp Response
	var  err error

	err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err!=nil {
		log.Println(err)
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	//supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId

	//----------------------------------------------------
	//标注输出
	var   distributionInfoVO responses.DistributionInfoVO
	//获取试卷导入数量
	var topic models.Topic
	topic.Question_id=questionId
	err = topic.GetTopic(questionId)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30009","获取试卷导入数量错误  ",err}
		c.Data["json"] = resp
		return
	}
	importNumber:=topic.Import_number
	distributionInfoVO.ImportTestNumber =importNumber
	//获取试卷未分配数量
	//查询相应试卷
	papers := make([]models.TestPaper, 0)
	err = models.FindUnDistributeTest(questionId, &papers)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30012","试卷分配异常，无法获取未分配试卷 ",err}
		c.Data["json"] = resp
		return
	}
	distributionInfoVO.LeftTestNumber=len(papers)
	//获取在线人数
	var onlineNumber ,err1= models.CountOnlineNumber()
	if err1!=nil {
		log.Println(err)
		resp  = Response{"30010","获取可分配人数错误  ",err}
		c.Data["json"] = resp
		return
	}
	distributionInfoVO.OnlineNumber=onlineNumber




	//----------------------------------------------------
	data := make(map[string]interface{})
	data["distributionInfoVO"] =distributionInfoVO
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}
/**
7.试卷分配
 */
func (c *AdminApiController) Distribution(){


	defer c.ServeJSON()
	var requestBody requests.Distribution
	var resp Response
	var  err error

	err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err!=nil {
		log.Println(err)
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	//supervisorId := requestBody.SupervisorId
	questionId := requestBody.QuestionId
	testNumber := requestBody.TestNumber
	userNumber := requestBody.UserNumber
	//----------------------------------------------------


	//是否需要二次阅卷
	var topic models.Topic
	topic.Question_id=questionId
	err = topic.GetTopic(questionId)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30011","试卷分配异常,无法获取试卷批改次数 ",err}
		c.Data["json"] = resp
		return
	}
	score_type := topic.Score_type

	//查询相应试卷
	papers := make([]models.TestPaper, 0)
	err = models.FindUnDistributeTest(questionId, &papers)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30012","试卷分配异常，无法获取未分配试卷 ",err}
		c.Data["json"] = resp
		return
	}
	testPapers := cutTest(papers, testNumber)
	//查找在线且未分配试卷的人
	usersList := make([]models.User, 0)
	err = models.FindUsers(&usersList)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30013","试卷分配异常，无法获取可分配阅卷员 ",err}
		c.Data["json"] = resp
		return
	}
	users := cutUser(usersList, userNumber)
	//第一次分配试卷
	countUser:=make([]int,userNumber)
	var ii int
	for i:=0 ;i<len(testPapers);{
		ii=i
		for j:=0 ;j<len(users);j++ {
			if testNumber==0{
				break
			}else {
			//修改testPaper改为已分配
			testPapers[ii].Correcting_status=1
				err := testPapers[ii].Update()
				if err!=nil {
					log.Println(err)
					resp  = Response{"30014","试卷第一次分配异常，无法更改试卷状态 ",err}
					c.Data["json"] = resp
					return
				}

				//添加试卷未批改记录
			var underCorrectedPaper  models.UnderCorrectedPaper
			underCorrectedPaper.Test_id=testPapers[ii].Test_id
			underCorrectedPaper.Question_id=testPapers[ii].Question_id
			underCorrectedPaper.Test_question_type=1
			underCorrectedPaper.User_id=users[j].User_id
				err = underCorrectedPaper.Save()
				if err!=nil {
					log.Println(err)
					resp  = Response{"30015","试卷第一次分配异常，无法生成待批改试卷 ",err}
					c.Data["json"] = resp
					return
				}
				countUser[j]=countUser[j]+1
			testNumber--
			ii++
			}
		}
		i=i+userNumber
	}
	//二次阅卷
	if score_type==1 {
		testNumber=len(testPapers)
		revers(users)
		var ii int
		for i:=0 ;i<len(testPapers);{
			ii=i
			for j:=0 ;j<len(users);j++ {
				if testNumber==0{
					break
				}else {
					//修改testPaper改为已分配
					testPapers[ii].Correcting_status=1
					err := testPapers[ii].Update()
					if err!=nil {
						log.Println(err)
						resp  = Response{"30016","试卷第二次分配异常，无法更改试卷状态 ",err}
						c.Data["json"] = resp
						return
					}

					//添加试卷未批改记录
					var underCorrectedPaper  models.UnderCorrectedPaper
					underCorrectedPaper.Test_id=testPapers[ii].Test_id
					underCorrectedPaper.Question_id=testPapers[ii].Question_id
					underCorrectedPaper.Test_question_type=2
					underCorrectedPaper.User_id=users[j].User_id
					err = underCorrectedPaper.Save()
					if err!=nil {
						log.Println(err)
						resp  = Response{"30017","试卷第二次分配异常，无法更改试卷状态 ",err}
						c.Data["json"] = resp
						return
					}
					countUser[j]=countUser[j]+1
					testNumber--
					ii++
				}
			}
			i=i+userNumber
		}

	}

	for i:=0;i<userNumber;i++{
		//添加试卷分配表
		var paperDistribution	 models.PaperDistribution
		paperDistribution.Test_distribution_number=int64(countUser[i])
		paperDistribution.User_id=users[i].User_id
		paperDistribution.Question_id=questionId
		err := paperDistribution.Save()
		if err!=nil {
			log.Println(err)
			resp  = Response{"30018","试卷分配异常，试卷分配添加异常 ",err}
			c.Data["json"] = resp
			return
		}

		//修改user变为已分配
		var user models.User
		user.IsDistribute=1
		user.QuestionId=questionId
		err = user.Update()
		if err!=nil {
			log.Println(err)
			resp  = Response{"30019","试卷分配异常，用户分配状态更新失败 ",err}
			c.Data["json"] = resp
			return
		}

	}

	//----------------------------------------------------
	data := make(map[string]interface{})
	data["data"] =nil
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}
/**
8.图片显示
 */
func (c *AdminApiController) Pic() {
	defer c.ServeJSON()
	var requestBody requests.ReadFile
	var resp Response
	var  err error
	err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err!=nil {
		log.Println(err)
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	   //supervisorId := requestBody.SupervisorId
	   //获取图片名
	   picName := requestBody.PicName
       //获取图片地址（win版）
	   src:="C:\\Users\\chen\\go\\src\\open-ct\\img\\"+picName
	   //linux版（）
       // var src := "/usr/workspace/src/open-ct/"+name

     //-------------------------------------
	bytes, err := os.ReadFile(src)
	encoding := base64.StdEncoding.EncodeToString(bytes)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30020","图片显示异常 ",err}
		c.Data["json"] = resp
		return
	}
	//c.Ctx.Output.Header("Content-Type", "image/jpeg")
	//c.Ctx.Output.Header("Content-Length", strconv.Itoa(len(data)))
	//c.Ctx.WriteString(string(data))
	//c.Ctx.ResponseWriter.WriteHeader(200)
	//----------------------------
	data := make(map[string]interface{})
	data["encoding"] =encoding
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}

/**
数组转置函数
 */
func revers(users []models.User)  {
	for i:=0 ;i<len(users)/2;i++ {
		var temp models.User
		temp=users[i]
		users[i]=users[len(users)-i-1]
		users[len(users)-i-1]=temp
	}
}
/**
截断数组函数
 */
func cutTest(oldData []models.TestPaper, n int) (newData[]models.TestPaper ) {
	 newData1 := make([]models.TestPaper, n)
	for i:=0 ;i<n;i++ {
		newData1[i]=oldData[i]
	}
	return newData1
}
func cutUser(oldData []models.User, n int) (newData[]models.User) {
	 newData1 := make([]models.User, n)
	for i:=0 ;i<n;i++ {
		newData1[i]=oldData[i]
	}
	return newData1
}

/**
9.大题展示列表
*/


func (c *AdminApiController) TopicList() {
	defer c.ServeJSON()
	var requestBody requests.TopicList
	var resp Response
	var  err error

	err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err!=nil {
		log.Println(err)
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	//supervisorId := requestBody.SupervisorId

	//----------------------------------------------------
	//获取大题列表
	topics  := make([]models.Topic,0)
	err = models.FindTopicList(&topics)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30021","获取大题参数设置记录表失败  ",err}
		c.Data["json"] = resp
		return
	}

	var topicVOList = make([]responses.TopicVO,len(topics))
	for i := 0; i < len(topics); i++ {

		topicVOList[i].SubjectName=topics[i].Subject_name
		topicVOList[i].TopicName=topics[i].Question_name
		topicVOList[i].Score=topics[i].Question_score
		topicVOList[i].StandardError=topics[i].Standard_error
		topicVOList[i].ScoreType=topics[i].Score_type
		topicVOList[i].TopicId=topics[i].Question_id
		topicVOList[i].ImportTime=topics[i].Import_time

		subTopics := make([]models.SubTopic, 0)
		models.FindSubTopicsByQuestionId(topics[i].Question_id,&subTopics)
		if err!=nil {
			log.Println(err)
			resp  = Response{"30022","获取小题参数设置记录表失败  ",err}
			c.Data["json"] = resp
			return
		}
		subTopicVOS := make([]responses.SubTopicVO, len(subTopics))
		for j:=0;j<len(subTopics);j++ {
			subTopicVOS[j].SubTopicId=subTopics[j].Question_detail_id
			subTopicVOS[j].SubTopicName=subTopics[j].Question_detail_name
			subTopicVOS[j].Score=subTopics[j].Question_detail_score
			subTopicVOS[j].ScoreDistribution=subTopics[j].Score_type
		}
		topicVOList[i].SubTopicVOList=subTopicVOS
	}

	//----------------------------------------------------
	data := make(map[string]interface{})
	data["topicVOList"] =topicVOList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}


/**
DistributionRecord
 */
func (c *AdminApiController) DistributionRecord() {
	defer c.ServeJSON()
	var requestBody requests.DistributionRecord
	var resp Response
	var  err error

	err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err!=nil {
		log.Println(err)
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	//supervisorId := requestBody.SupervisorId
	subjectName := requestBody.SubjectName
	//----------------------------------------------------
	//获取大题列表
	topics  := make([]models.Topic,0)
	err = models.FindTopicBySubNameList(&topics,subjectName)
	if err!=nil {
		log.Println(err)
		resp  = Response{"30023","获取试卷分配记录表失败  ",err}
		c.Data["json"] = resp
		return
	}

	var distributionRecordList = make([]responses.DistributionRecordVO,len(topics))
	for i := 0; i < len(topics); i++ {

		distributionRecordList[i].TopicId=topics[i].Question_id
		distributionRecordList[i].TopicName=topics[i].Question_name
		distributionRecordList[i].ImportNumber=topics[i].Import_number
		distributionTestNumber,err :=models.CountTestDistributionNumberByQuestionId(topics[i].Question_id)
		if err!=nil {
			log.Println(err)
			resp  = Response{"30024","获取试卷分配记录表失败，统计试卷已分配数失败  ",err}
			c.Data["json"] = resp
			return
		}
		distributionUserNumber ,err:=models.CountUserDistributionNumberByQuestionId(topics[i].Question_id)
		if err!=nil {
			log.Println(err)
			resp  = Response{"30025","获取试卷分配记录表失败，统计用户已分配数失败  ",err}
			c.Data["json"] = resp
			return
		}
		distributionRecordList[i].DistributionUserNumber=distributionUserNumber
		distributionRecordList[i].DistributionTestNumber=distributionTestNumber

	}

	//----------------------------------------------------
	data := make(map[string]interface{})
	data["distributionRecordList"] =distributionRecordList
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp
}



/**
试卷删除
*/

func (c *AdminApiController) DeleteTest(){

	defer c.ServeJSON()
	var requestBody requests.DeleteTest
	var resp Response
	var  err error

	err =json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err!=nil {
		log.Println(err)
		resp = Response{"10001","cannot unmarshal",err}
		c.Data["json"] = resp
		return
	}
	//adminId := requestBody.AdminId
	questionId := requestBody.QuestionId

	//----------------------------------------------------
	count, err := models.CountUnScoreTestNumberByQuestionId(questionId)
	if count==0 {
		models.DeleteAllTest(questionId)
		subTopics:=make([]models.SubTopic,0)
		models.FindSubTopicsByQuestionId(questionId,&subTopics)
		for j:=0 ;j<len(subTopics);j++ {
			subTopic:= subTopics[j]
		testPaperInfos :=make([]models.TestPaperInfo,0)
		models.FindTestPaperInfoByQuestionDetailId(subTopic.Question_detail_id,&testPaperInfos)
			for k:=0;k<len(testPaperInfos);k++ {
				//imgSrc :=testPaperInfos[k].Pic_src
				//删除图片
				testPaperInfos[k].Delete()
			}
		}

	}else {
		resp  = Response{"30030","试卷未批改完不能删除  ",err}
		c.Data["json"] = resp
		return
	}

	//----------------------------------------------------
	data := make(map[string]interface{})
	data["data"] =nil
	resp = Response{"10000", "OK", data}
	c.Data["json"] = resp

}