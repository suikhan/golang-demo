数据库升级脚本设计
================

1、字符显示
----------

**用途：** 用于打印文字到屏幕中

```powershell
@ECHO("数据库升级工具 v1.0.0.1")
@ECHO("===================================")
@ECHO("电子病历系统数据库升级工具 Nueqsoft   ")
@ECHO("Design by suikhan@126.com")
@ECHO("")

```

2、数据输入
----------

**用途：** 用于从屏幕输入文字到变量中

```powershell
@INPUT("请输入EMR登录用户：") => EMR_USER
@INPUT("请输入EMR登录密码：") => EMR_PASSWORD
@INPUT("请输入EMR服务器IP：") => EMR_SERVERIP

@INPUT("请输入HIS登录用户：") => HIS_USER
@INPUT("请输入HIS登录密码：") => HIS_PASSWORD
@INPUT("请输入HIS服务器IP：") => HIS_SERVERIP

```

3、进度说明
----------

**用途：** 用于在进度条上面的Label中提示信息

```powershell
@HINT("正在执行....操作")

```

4、SQL代码段
------------

**用途：** 用与标识SQL代码段

```powershell
<sql>

</sql>
```

5、表格定义段
------------

```powershell
<table:user("用户表")>
usercode,varchar2(20),N
username,varchar2(200),N
password,varchar2(20),Y,"123456"
isvalid,char(1),N,"1"
</table>
```

6、表格数据段
------------

**用途：** 用于插入数据库中已有表的初始数据

```powershell
<data:user(usercode, username, password, isvalid)>
0000,超级管理员,1,1
0001,张三,1,1
</data>
```

7、执行脚本文件
--------------

**用途：** 调用多个脚本文件

```powershell
@EXECUTE("A.DUS")

```

8、执行查询语句
--------------

**用途：** 执行查询语句并返回结果

```powershell
@QUERY("SELECT * FROM ...") => HISVER

```

9、分支语句
----------

```powershell
@IF($HISVER == "5.1")
{
    @EXECUTE("A.DUS")
}
@ELSE
{
    @EXECUTE("B.DUS")
}
```