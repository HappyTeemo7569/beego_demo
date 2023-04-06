package models

import (
	"time"
)

type Cart struct {
	Id         int64     `xorm:"pk autoincr comment('购物车id') BIGINT(20)"`
	Userid     int64     `xorm:"not null default 0 comment('用户id') index BIGINT(20)"`
	Proid      int64     `xorm:"not null default 0 comment('商品id') index BIGINT(20)"`
	Quantity   int       `xorm:"not null default 0 comment('数量') INT(11)"`
	Checked    int       `xorm:"not null default 0 comment('是否选择,1=已勾选,0=未勾选') INT(11)"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

type Category struct {
	Id         int       `xorm:"not null pk autoincr comment('分类id') SMALLINT(6)"`
	Parentid   int       `xorm:"not null default 0 comment('父类别id当id=0时说明是根节点,一级类别') SMALLINT(6)"`
	Name       string    `xorm:"not null default '' comment('类别名称') VARCHAR(50)"`
	Status     int       `xorm:"not null default 1 comment('类别状态1-正常,2-已废弃') TINYINT(4)"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}

type Orderitem struct {
	Id           int64     `xorm:"pk autoincr comment('订单子表id') BIGINT(20)"`
	OrderId      string    `xorm:"not null default '' comment('订单id') index VARCHAR(64)"`
	UserId       int64     `xorm:"not null default 0 comment('用户id') index BIGINT(20)"`
	ProductId    int64     `xorm:"not null default 0 comment('商品id') index BIGINT(20)"`
	ProductName  string    `xorm:"not null default '' comment('商品名称') VARCHAR(100)"`
	ProductImage string    `xorm:"not null default '' comment('商品图片地址') VARCHAR(500)"`
	CurrentPrice string    `xorm:"not null default 0.00 comment('生成订单时的商品单价，单位是元,保留两位小数') DECIMAL(20,2)"`
	Quantity     int       `xorm:"not null default 0 comment('商品数量') INT(10)"`
	TotalPrice   string    `xorm:"not null default 0.00 comment('商品总价,单位是元,保留两位小数') DECIMAL(20,2)"`
	CreateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

type Orders struct {
	Id          string    `xorm:"not null pk default '' comment('订单id') VARCHAR(64)"`
	Userid      int64     `xorm:"not null default 0 comment('用户id') index BIGINT(20)"`
	Shoppingid  int64     `xorm:"not null default 0 comment('收货信息表id') BIGINT(20)"`
	Payment     string    `xorm:"default 0.00 comment('实际付款金额,单位是元,保留两位小数') DECIMAL(20,2)"`
	Paymenttype int       `xorm:"not null default 1 comment('支付类型,1-在线支付') TINYINT(4)"`
	Postage     int       `xorm:"not null default 0 comment('运费,单位是元') INT(10)"`
	Status      int       `xorm:"not null default 10 comment('订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭') SMALLINT(6)"`
	CreateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

type Payinfo struct {
	Id             int64     `xorm:"pk autoincr comment('支付信息表id') BIGINT(20)"`
	Orderid        string    `xorm:"not null default '' comment('订单id') index VARCHAR(64)"`
	Userid         int64     `xorm:"not null default 0 comment('用户id') index BIGINT(20)"`
	Payplatform    int       `xorm:"not null default 0 comment('支付平台:1-支付宝,2-微信') TINYINT(4)"`
	Platformnumber string    `xorm:"not null default '' comment('支付流水号') VARCHAR(200)"`
	Platformstatus string    `xorm:"not null default '' comment('支付状态') VARCHAR(20)"`
	CreateTime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

type Product struct {
	Id         int64     `xorm:"pk autoincr comment('商品id') BIGINT(20)"`
	Cateid     int       `xorm:"not null default 0 comment('类别Id') index SMALLINT(6)"`
	Name       string    `xorm:"not null default '' comment('商品名称') VARCHAR(100)"`
	Subtitle   string    `xorm:"not null default '' comment('商品副标题') VARCHAR(200)"`
	Images     string    `xorm:"not null default '' comment('图片地址,逗号分隔') VARCHAR(1024)"`
	Detail     string    `xorm:"not null default '' comment('商品详情') VARCHAR(1024)"`
	Price      string    `xorm:"not null default 0.00 comment('价格,单位-元保留两位小数') DECIMAL(20,2)"`
	Stock      int       `xorm:"not null default 0 comment('库存数量') INT(11)"`
	Status     int       `xorm:"not null default 1 comment('商品状态.1-在售 2-下架 3-删除') INT(6)"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') index DATETIME"`
}

type Reply struct {
	Id            int64     `xorm:"pk autoincr comment('评论表id') BIGINT(20)"`
	Business      string    `xorm:"not null default '' comment('评论业务类型') VARCHAR(64)"`
	Targetid      int64     `xorm:"not null default 0 comment('评论目标id') index BIGINT(20)"`
	ReplyUserid   int64     `xorm:"not null default 0 comment('回复用户id') BIGINT(20)"`
	BeReplyUserid int64     `xorm:"not null default 0 comment('被回复用户id') BIGINT(20)"`
	Parentid      int64     `xorm:"not null default 0 comment('父评论id') BIGINT(20)"`
	Content       string    `xorm:"not null default '' comment('评论内容') VARCHAR(255)"`
	Image         string    `xorm:"not null default '' comment('评论图片') VARCHAR(255)"`
	CreateTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

type Shipping struct {
	Id               int64     `xorm:"pk autoincr comment('收货信息表id') BIGINT(20)"`
	Orderid          string    `xorm:"not null default '' comment('订单id') index VARCHAR(64)"`
	Userid           int64     `xorm:"not null default 0 comment('用户id') index BIGINT(20)"`
	ReceiverName     string    `xorm:"not null default '' comment('收货姓名') VARCHAR(20)"`
	ReceiverPhone    string    `xorm:"not null default '' comment('收货固定电话') VARCHAR(20)"`
	ReceiverMobile   string    `xorm:"not null default '' comment('收货移动电话') VARCHAR(20)"`
	ReceiverProvince string    `xorm:"not null default '' comment('省份') VARCHAR(20)"`
	ReceiverCity     string    `xorm:"not null default '' comment('城市') VARCHAR(20)"`
	ReceiverDistrict string    `xorm:"not null default '' comment('区/县') VARCHAR(20)"`
	ReceiverAddress  string    `xorm:"not null default '' comment('详细地址') VARCHAR(200)"`
	CreateTime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

type User struct {
	Id         int64     `xorm:"pk autoincr comment('用户ID') BIGINT(20)"`
	Username   string    `xorm:"not null default '' comment('用户名') unique VARCHAR(50)"`
	Password   string    `xorm:"not null default '' comment('用户密码，MD5加密') VARCHAR(50)"`
	Phone      string    `xorm:"not null default '' comment('手机号') unique VARCHAR(20)"`
	Question   string    `xorm:"not null default '' comment('找回密码问题') VARCHAR(100)"`
	Answer     string    `xorm:"not null default '' comment('找回密码答案') VARCHAR(100)"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') index TIMESTAMP"`
}

type UserReceiveAddress struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	Uid           int64     `xorm:"not null default 0 comment('用户id') index BIGINT(20)"`
	Name          string    `xorm:"not null default '' comment('收货人名称') VARCHAR(64)"`
	Phone         string    `xorm:"not null default '' comment('手机号') VARCHAR(20)"`
	IsDefault     int       `xorm:"not null default 0 comment('是否为默认地址') TINYINT(1)"`
	PostCode      string    `xorm:"not null default '' comment('邮政编码') VARCHAR(100)"`
	Province      string    `xorm:"not null default '' comment('省份/直辖市') VARCHAR(100)"`
	City          string    `xorm:"not null default '' comment('城市') VARCHAR(100)"`
	Region        string    `xorm:"not null default '' comment('区') VARCHAR(100)"`
	DetailAddress string    `xorm:"not null default '' comment('详细地址(街道)') VARCHAR(128)"`
	IsDelete      int       `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
	CreateTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('数据创建时间[禁止在代码中赋值]') TIMESTAMP"`
	UpdateTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('数据更新时间[禁止在代码中赋值]') TIMESTAMP"`
}
