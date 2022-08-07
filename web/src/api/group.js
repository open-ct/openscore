import axios from "axios";
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
const group = {

  userInfo(data) {
    return axios.post("/marking/supervisor/user/info", data);
  },

  questionList(data) {
    return axios.post("/marking/admin/questionBySubList", data);
  },

  teacherMonitor(data) {
    return axios.post("/marking/supervisor/teacher/monitoring", data);
  },

  scoreMonitor(data) {
    return axios.post("/marking/supervisor/score/distribution", data);
  },

  selfTeacher(data) {
    return axios.post("/marking/supervisor/question/teacher/list", data);
  },

  selfMonitor(data) {
    return axios.post("/marking/supervisor/self/score", data);
  },

  averageMonitor(data) {
    return axios.post("/marking/supervisor/average/score", data);
  },

  standardMonitor(data) {
    return axios.post("/marking/supervisor/score/deviation ", data);
  },

  allMonitor(data) {
    return axios.post("/marking/supervisor/score/progress ", data);
  },

  arbitramentList(data) {
    return axios.post("/marking/supervisor/arbitrament/list ", data);
  },

  problemList(data) {
    return axios.post("/marking/supervisor/problem/list ", data);
  },

  selfMarkList(data) {
    return axios.post("/marking/supervisor/selfMark/list ", data);
  },

  MonitorPoint(data) {
    return axios.post("/marking/supervisor/point", data);
  },

  problemTestId(data) {
    return axios.post("/marking/supervisor/problem/unmark/list", data);
  },

  arbitrationTestId(data) {
    return axios.post("/marking/supervisor/arbitrament/unmark/list", data);
  },

  selfTestId(data) {
    return axios.post("/marking/supervisor/selfMark/unmark/list", data);
  },

};
export default group;
