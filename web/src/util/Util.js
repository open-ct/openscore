export function getTextByJs(arr) {
  var str = "";
  for (var i = 0; i < arr.length; i++) {
    str += arr[i] + "-";
  }
  // 去掉最后一个逗号(如果不需要去掉，就不用写)

  if (str.length > 0) {

    str = str.substr(0, str.length - 1);

  }
  return str;
}
export function getRandom() {  // 获取随机数id
  let date = Date.now();
  let rund = Math.ceil(Math.random() * 1000);
  let id = date + "" + rund;
  return id;
}
