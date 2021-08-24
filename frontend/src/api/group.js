import axios from 'axios'
function getServerUrl() {
  const hostname = window.location.hostname
  if (hostname === 'localhost') {
    return `http://${hostname}:8080/openct`
  }
  return '/openct'
}
axios.defaults.baseURL = getServerUrl();
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
axios.defaults.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
const group = {

  userInfo(data) {
    return axios.post('/marking/supervisor/user/info', data)
  },

  questionList(data) {
    return axios.post('/marking/supervisor/question/list', data)
  },

  teacherMonitor(data) {
    return axios.post('/marking/supervisor/teacher/monitoring', data)
  },

  scoreMonitor(data) {
    return axios.post('/marking/supervisor/score/distribution', data)
  },

  selfTeacher(data) {
    return axios.post('/marking/supervisor/question/teacher/list', data)
  },

  selfMonitor(data) {
    return axios.post('/marking/supervisor/self/score', data)
  },



}
export default group