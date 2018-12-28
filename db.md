# 投票活动原型系统

## 活动类

### 1. 活动表[atv_activity_info]

| 字段名         | 类型           | 默认值 | 为空 |   约束    | 描述                                |
| -------------- | -------------- | :----: | :--: | :-------: | :---------------------------------- |
| atv_id         | number(10)     |        |  否  | PK,CS,SEQ,DI | 活动编号                            |
| appid          | varchar2(64)   |        |  否  |    CS     | appid                               |
| title          | varchar2(128)  |        |  否  |   CUQS,DN | 活动标题                            |
| tag            | varchar2(32)   |        |  否  |    CUS    | 活动标识                            |
| url            | varchar2(128)  |        |  否  |    CS     | 活动链接                            |
| expire_start   | date           |        |  否  |    CUS    | 开始时间                            |
| expire_end     | date           |        |  否  |    CUS    | 结束时间                            |
| member_count   | number(10)     |        |  否  |    CS     | 参赛人数                            |
| voted_count    | number(10)     |        |  否  |    CS     | 投票总数                            |
| involved_count | number(10)     |        |  否  |    CS     | 投票人数                            |
| visit_count    | number(10)     |        |  否  |    CS     | 访问总数                            |
| gift_count     | number(10)     |        |  否  |    CQS    | 创建时间                            |
| update_time    | date           |        |  否  |    CS     | 更新时间                            |
| voting_rule    | number(10)     |        |  否  |    CUS    | 投票规则(1:每日可投　 2:只能投一票) |
| operator       | varchar2(32)   |        |  否  |    CS     | 操作人                              |
| status         | number(2)      |   0    |  否  |    CUS    | 状态(1：禁用 0：启用)             |
| music          | varchar2(256)  |        |  是  |    CUS    | 背景音乐                            |
| rule           | varchar2(2048) |        |  是  |    CUS    | 活动规则                            |
| intro          | nclob          |        |  是  |    CUS    | 活动介绍                            |
| prize          | nclob          |        |  是  |    CUS    | 奖品明细                            |

### 2. 活动宣传页[atv_banner_page]

| 字段名       | 类型           | 默认值  | 为空 |   约束                       | 描述                         |
| ----------- | ------------- | :----: | :--: | :-------------------------: | :---------------------------|
| atv_page_id | number(10)    |        |  否  | PK,CS,SEQ                   | 页面编号                     |
| atv_id      | number(10)    |        |  否  | CQS,DD(atv_activity_info)   | 活动编号                     |
| page_index  | number(2)     |        |  否  |    CUS                      | 页索引                       |
| bg_img_url  | varchar2(128) |        |  是  |    CUS                      | 背景图片                     |
| content     | nclob         |        |  是  |    CUS                      | 页面内容                     |
| is_show     | number(1)     |        |  否  |    CUS                      | 是否显示 (0:显示　 1:不显示)   |

