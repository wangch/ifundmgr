// Copyright 2015 iCloudFund. All Rights Reserved.

package main

// icloudfund.com ==> 官网 + 网关
// 官网 说明页面和链接: ICC/USD 锚定值  ICC 发行数量 法币发行数量 如果购买和赎回 ICC 如何存款和取款法币 各种说明帮助和链接
// ripple.txt + federation API + quote API
// 和Ripple 网关不同的是 : 网关接收用户发送的 ICC 时, 作为ICC赎回处理

// ifundmgr.com ==> 网关平台
// 用于法币的发行 存款取款的申请，确认，审批，保存，查询等
// 用于ICC的发行和回收 发行回收的申请，确认，审批，查询等
// 监控网关钱包的交易

// 存款申请审批机制：
// 	a, 存款人通过客服指导存款
// 	b, 客服填写存款申请记录 存款人姓名, ICC钱包, 银行卡号, 币种, 数额, 转账银行
// 	c, 到账后财务审批此申请记录, 确认
// 	d, 总监确认
//  e, 会计通过客户端(gateway hotwallet 账号)发送存款人金额, 确认

// 取款申请审批机制：
// 	a, 用户在客户端发起取款申请——通过客户端取款功能, 填写存款人姓名, 银行卡号, 开户行地址, 金额, 币种, 联系信息等; 发送给网关gateway
// 	b, 财务对取款申请进行审批
//  c, 总监审核
// 	d, 会计根据根据此申请, 进行转账操作

// 官网发行ICC机制:
//     ICC锚锭美元价格, 并且始终不变, 比如1 ICC == 1 USD
// 	a, 客户转账美元购买ICC（使用其他货币, 根据美元汇率计算？）
// 	b, 客服填写存款申请记录 存款人姓名, ICC钱包, 银行卡号, 币种, 数额, 转账银行
// 	c, 财务审批确认
// 	d, 总监确认
//  e, 会计通过gateway发送存款人金额, 确认

// 官网回收ICC机制:
//  a, 用户填写ICC赎回申请: 在客户端取款页面, 币种填写 ICC （或者实现类似取款的赎回页面, 只允许用户发送ICC到指定的账户）
// 	b, 财务通过监控ICC赎回账号, 确认取款申请（可以使用gateway一个钱包地址）
//  c, 总监确认
// 	d, 会计进行转账操作, 确认

//

// ===========

// 系统的用户是预生成的，所以无需注册：管理员预先生成几个账户和密码分发给系统响应的工作人员，并要求更改密码；
// 更改密码需要提供旧密码
// 客服只能提交存款申请
// 财务只能对提交的申请进行确认，和在取款和回收转账完成后提交OK ==> 列出存款申请表单，列出总监确认的取款表单
// 总监能够查看所有记录，包括已完成和未完成的记录，但是不能代替财务神奇和确认
// 对于存款，总监确认后，记录入库保存；对于取款，财务确认后，记录入库保存

// 财务人员可以根据存款内容查询记录

// 数据库表 3个：
// 1, 人员表
// 2, 请求表 ==> deposit withdrawal issue redeem
// 3, 记录表 ==> deposit withdrawal issue redeem
// 4, log 日志表

// 存款取款以及ICC发行和回收需要指定手续费，并计入记录

// ===========

// 数据库：mysql

// ===========

// 对账功能:
// 网关账户(cold wallet)货币currency余额(负值) == 数据库中 currency 记录 sum(deposit) - sum(withdrawal) == 财务账户 currency货币余额
// ICC 网关账户 ICC发行数量 == 数据库 ICC 记录 sum(issue) - sum(redeem) == 对应财务账户美元余额 * ICC/USD

/*
ip: 74.82.1.163 和 74.82.1.164
username: CHWang
password:  Wang366478
*/

// m01 1qaz2wsx iLyHPoNvsdN7LFFvpy8GGkDJGV1xo3V5We
// m02 1qaz2wsx iMxKojv7vNYyca7YdsVBSRKCmvGciHUNap
// m03 1qaz2wsx iLwUZfEo8pB9VTxzDtjwJBBuiWuVpLxz9m

// 测试
// 1, 存款港元 到 m01
// 2, 取款港元
// 3, 发行ICC
// 4, 回收ICC

// TODO
// 1, 统计数据
// 2, 修改密码
// 3, 填写表单时, 金额和费率允许小数
// 4, 界面优化
