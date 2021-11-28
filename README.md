## gaia
主机批量初始化和部署配置服务器，它拥有如下特征：
- 目前主要基于 ssh 公钥认证对目标主机进行命令的执行和文件的上传下载。
- 插件驱动，插件分为任务性插件和服务性插件。
    - 任务性插件：用于执行或者组合为剧本再执行。
    - 服务性插件：将动态注册路由到服务器，提供 HTTP 服务。最常见的比如 仓库服务。
- 基于 clean architecture 思想进行编写，进行业务与框架的分离，能轻松切换存储层，WEB层的框架。


## TODO
### core
- 节点的添加，获取与删除。

### plugin
......