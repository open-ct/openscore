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

}
export default group