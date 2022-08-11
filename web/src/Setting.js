import {message} from "antd";
import axios from "axios";
import Sdk from "casdoor-js-sdk";

function getServerUrl() {
  const hostname = window.location.hostname;
  if (hostname === "localhost") {
    return `http://${hostname}:8080`;
  }
  return "";
}
// import {isMobile as isMobileDevice} from "react-device-detect";
export let ServerUrl2 = getServerUrl();
export let ServerUrl = ServerUrl2 + "/openct";
// export let ServerUrl = "http://47.106.83.3:8080"+'/openct';
// ServerUrl2="http://47.106.83.3:8080";
axios.defaults.baseURL = ServerUrl;
axios.defaults.withCredentials = true;
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
axios.defaults.headers["Content-Type"] = "application/x-www-form-urlencoded;charset=UTF-8";

export let CasdoorSdk;
export function initCasdoorSdk(config) {
  CasdoorSdk = new Sdk(config);
}

export function parseJson(s) {
  if (s === "") {
    return null;
  } else {
    return JSON.parse(s);
  }
}

export function myParseInt(i) {
  const res = parseInt(i);
  return isNaN(res) ? 0 : res;
}

export function openLink(link) {
  // this.props.history.push(link);
  const w = window.open("about:blank");
  w.location.href = link;
}

export function goToLink(link) {
  window.location.href = link;
}

export function showMessage(type, text) {
  if (type === "") {
    return "";
  } else if (type === "success") {
    message.success(text);
  } else if (type === "error") {
    message.error(text);
  } else if (type === "warn") {
    message.warn(text);
  }
}
