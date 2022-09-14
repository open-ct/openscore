import axios from "axios";
import * as Settings from "../Setting";
function getServerUrl() {
  const hostname = window.location.hostname;
  if (hostname === "localhost") {
    return `http://${hostname}:8080/openct`;
  }
  return "/openct";
}
axios.defaults.baseURL = getServerUrl();
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
axios.defaults.headers["Content-Type"] = "application/x-www-form-urlencoded;charset=UTF-8";
const Manage = {

  questionImport(data) {
    return axios.post("/marking/admin/insertTopic", data);
  },

  subjectList(data) {
    return axios.post("/marking/admin/subjectList", data);
  },

  questionList(data) {
    return axios.post("/marking/admin/questionBySubList", data);
  },

  distributeInfo(data) {
    return axios.post("/marking/admin/distribution/info", data);
  },

  distributePaper(data) {
    return axios.post("/marking/admin/distribution", data);
  },

  questionInfo(data) {
    return axios.post("/marking/admin/topicList", data);
  },

  paperInfo(data) {
    return axios.post("/marking/admin/DistributionRecord", data);
  },

  listUsers(data) {
    return axios.post("/marking/admin/listUsers", data);
  },
  createUser(data) {
    return axios.post("/marking/admin/createUser", data);
  },
  deleteUser(data) {
    return axios.post("/marking/admin/deleteUser", data);
  },
  updateUser(data) {
    return axios.post("/marking/admin/updateUser", data);
  },
  createSmallQuestion(data) {
    return axios.post("/marking/admin/createSmallQuestion", data);
  },
  deleteSmallQuestion(data) {
    return axios.post("/marking/admin/deleteSmallQuestion", data);
  },
  updateSmallQuestion(data) {
    return axios.post("/marking/admin/updateSmallQuestion", data);
  },
  deleteQuestion(data) {
    return axios.post("/marking/admin/deleteQuestion", data);
  },
  updateQuestion(data) {
    return axios.post("/marking/admin/updateQuestion", data);
  },
  subjectAllot(data) {
    axios.request({
      url: "/marking/admin/writeUserExcel",
      headers: {
        "Content-Type": "application/json", // 重要
        "accept": "application/octet-stream", // 重要
      },
      method: "POST",
      data: data,
      responseType: "blob", // 重要
    }).then(function(response) {
      let data = response.data;
      let url = URL.createObjectURL(data);// 重要
      let link = document.createElement("a");
      link.href = url;
      link.download = "用户导出.xlsx";// 重要--决定下载文件名
      link.click();
      link.remove();
    }).catch(function(e) {Settings.showMessage("error", e);});
  },
  updateUserQualified(data) {
    return axios.post("/marking/admin/updateUserQualified", data);
  },
  getListGroupGrades(data) {
    return axios.post("/marking/admin/listGroupGrades", data);
  },
  deletePaperFromGroup(data) {
    return axios.post("/marking/admin/deletePaperFromGroup", data);
  },
  getListTestPapersByQuestionId(data) {
    return axios.post("/marking/admin/listTestPapersByQuestionId", data);
  },
  getListPaperGroups(data) {
    return axios.post("/marking/admin/listPaperGroups", data);
  },
  teachingPaperGrouping(data) {
    return axios.post("/marking/admin/teachingPaperGrouping", data);
  },
  getSchoolsList(data) {
    return axios.post("/marking/admin/listSchools", data);
  },
  getListTestPaperInfo(data) {
    return axios.post("/marking/admin/listTestPaperInfo", data);
  },
};
export default Manage;
