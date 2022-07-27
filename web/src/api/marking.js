// import request from './request'
// import qs from 'qs';
import axios from 'axios'
// function getServerUrl() {
//     const hostname = window.location.hostname
//     if (hostname === 'localhost') {
//       return `http://${hostname}:8080/openct`
//     }
//     return '/openct'
//   }
// axios.defaults.baseURL = getServerUrl();
// axios.defaults.withCredentials=true;
// // axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
// axios.defaults.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'; 
const Marking = {
    //   testList(data) {
    //     return request({
    //       url: '/marking/score/test/list',
    //       method: 'post',
    //       data: data,
    //     })
    //   },
    // testList(userId) {
    //     return fetch(`http://localhost:8080/openct/marking/score/test/list`, {
    //         method: "post",
    //         credentials: "include",
    //         params: qs.stringify({userId:userId}),
    //     });
    // }
    testList(data) {
        return  axios.post('/marking/score/test/list', data)
    },
    testDisplay(data) {
        return  axios.post('/marking/score/test/display', data)
    },
    testPoint(data) {
        return  axios.post('/marking/score/test/point', data)
    },
    testProblem(data) {
        return  axios.post('/marking/score/test/problem', data)
    },
    testAnswer(data) {
        return  axios.post('/marking/score/test/answer', data)
    },
    testDetail(data) {
        return  axios.post('/marking/score/test/example/detail', data)
    },
    testExampleList(data) {
        return  axios.post('/marking/score/test/example/list', data)
    },
    testReview(data) {
        return  axios.post('/marking/score/test/review', data)
    },
    testReviewPoint(data) {
        return  axios.post('/marking/score/test/review/point', data)
    },
    testSelf(data) {
        return  axios.post('/marking/score/self/list', data)
    },
    testSelfScore(data) {
        return  axios.post('/marking/score/self/point', data)
    },
    
}


export default Marking