/**
 * @param {string} str
 * @return {number}
 */
var myAtoi = function (str) {
  var sign = "+"; //正負號
  var numReg = /^[0-9]/;

  //將前面的空白字串去掉,  ex. "^^^^-122a" -> "-122a" (^代表空白)
  var i = 0;
  while (str.charAt(i) == " " && i < str.length) {
    i++;
  }
  str = str.slice(i);

  //處理正負號, ex. "-122a" -> "122a"
  if (str.startsWith("+")) {
    str = str.slice(1);
  } else if (str.startsWith("-")) {
    str = str.slice(1);
    sign = "-";
  }

  //處理後的字串不是數字開頭,代表字串不合法
  if (!numReg.test(str)) return 0;

  //出現不是數字的符號就中斷
  var j = 0;
  while (j < str.length && numReg.test(str[j])) {
    j++;
  }
  // 截出開頭的數字字串, ex. '122a' -> '122'
  str = str.substr(0, j);

  //字串轉成int
  var value = parseInt(sign + str);
  if (value > 2147483647) {
    return 2147483647;
  }
  if (value < -2147483648) {
    return -2147483648;
  }
  return value;
};
