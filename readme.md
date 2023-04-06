# 用的核心组件

- beego
  - https://github.com/go-xorm/xorm
  - https://pkg.go.dev/github.com/beego/beego#readme-more-info-at-beego-me-http-beego-me
- xrom  
  - https://xorm.io/docs/chapter-01/readme/


# 目录安排

- /pkg ： 可以被外部应用使用的代码库
- /build : 编译后的可执行文件
- /cmd : 程序入口
- /api : 当前项目对外提供的各种不同类型的 API 接口定义文件
- /worker ：任务
    - crontab : 定时任务
    - queue : 循环任务
    - command : 命令
- /extenal : 内部访问
  - /apps : 多服务, 有需要可以把external多目录部署，和api一样
- /internal ：业务代码
    - /biz : 处理请求
    - /core ： 核心，数据库、缓存等操作
    - /base : 公共依赖，比如后台配置
    - /data ： 数据库实体
    - /event ：事件，适用于快速、需要队列化的小任务（需要全局顺序用kafka）