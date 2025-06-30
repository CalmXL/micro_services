# 密码 password

## 密码安全问题

- 明文密码历史 md5(password) => 32 位随机字符串的东西 => 入库
- 彩虹表 => 就是保存结果 反向推导

## 盐值加密
- md5(md5(password) + salt)