

使用gorm  gin 框架实现 数据库调用与http服务 \
调用 pinying库实现 中文 拼音匹配 \
search层与sercice层分离 实现解耦 \
地址"github.com/Lofanmi/pinyin-golang/pinyin" \
通过mysql模糊查询实现前缀匹配 \
~~~
dao.SqlSession.Table("record").Select("title").
Where("title LIKE ?", keyword+"%").Find(&recordlist)
匹配关键字
dao.SqlSession.Table("record").Select("id").Last(&record)
保存搜索词记录
~~~

## config
数据库配置信息
## src
### common
返回结构体
### dao
连接mysql
### model
records 表 \
id  uint primary key \
title  varchar(128) 
### routes
创建路由
### search
search接口 调用sercice层
### service
与数据库交互 \
查找记录以及存储历史记录

