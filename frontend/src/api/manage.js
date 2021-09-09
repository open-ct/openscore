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
const Manage = {



  questionImport(data) {
    return axios.post('/marking/admin/insertTopic', data)
  },

  subjectList(data) {
    return axios.post('/marking/admin/subjectList', data)
  },

  questionList(data) {
    return axios.post('/marking/admin/questionBySubList', data)
  },

  distibuteInfo(data) {
    return axios.post('/marking/admin/distribution/info', data)
  },
  distibutePaper(data) {
    return axios.post('/marking/admin/distribution', data)
  },

}
export default Manage