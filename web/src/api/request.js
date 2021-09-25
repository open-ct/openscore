import axios from 'axios'
import { message } from 'antd'


function getServerUrl() {
  const hostname = window.location.hostname
  if (hostname === 'localhost') {
    return `http://${hostname}:8080/openct`
  }
  return '/openct'
}
// axios.defaults.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
const request = axios.create({
  baseURL: getServerUrl(),
  timeout: 5000,
  withCredentials: true,
  crossDomain: true,

})
request.interceptors.response.use(
  config => {
    if (config.status !== 10000) {
      message.error(config.msg)
      setTimeout(() => {
        window.location.href = '/'
      }, 1000)
    }
    return config
  }, error => {
    return Promise.reject(error)
  }
)

export default request
