# Stack-RPC

**1.0.0版本开发中，预计11月底发布，敬请期待！**

Stack-RPC旨在为中国开发者提供通用的分布式服务微服务开发库（比如配置管理、服务发现、熔断降级、路由、服务代理、安全、主从选举等）。基于Stack，开发者可以快速投入自身的业务开发中，只需要极少的学习成本。Stack适用于中小规模的开发场景，她可以轻易在桌面电脑、服务器、容器集群中搭建分布式服务。

项目于2020年11月2日正式成立，第一版本基于Go-Micro的1.18修改。

## 特性

Stack-RPC 同时即提供轻量的开发库，同时也提供对应高级别的扩展库，为大家带来开箱即用的开发体验。

支持的特性主要有：

- 分布式配置
- 服务注册与发现
- 服务路由
- 远程服务调用
- 负载均衡
- 链路中断与降级
- 分布式锁[todo]
- 主从选举[todo]
- 分布式广播

## 开始使用

我们为一直为大家准备持续开发、更新、愈加丰富的文档与资料：[MicroHQ](https://microhq.cn/docs/stack-rpc/introduce-cn)

## 与Go-Micro的差异

- 取消Micro工具集，以插件形式集成到Stack-RPC中
  - Web控制台插件
  - 网关插件
  - 注册中心插件
  - 配置中心插件
- 多类型服务同时部署
  - 支持RPC与HTTP同时暴露
- 增强日志特性
  - 支持动态修改日志级别
  - 支持日志自定义目录存储
  - 支持按级别存储不同文件
  - 支持按周期压缩日志文件
  - 支持按大小压缩日志文件
- 增强配置特性
  - 增加默认配置命名空间
  - 定义默认配置文件存储目录
  - 支持Apollo配置中心
- 删除不要的组件
  - Cloudflare

## 鸣谢

- 感谢Go-Micro库，提供优秀的扩展性极强的原始框架，Stack-RPC作为衍生版本，受益颇多，同时Go-Micro的肄业也给Stack-RPC创造了生命
- 感谢Spring-Cloud，作为使用最广泛的开源分布式开发库，我们参考了她许多优秀的设计与文档
- 感谢各位Go-Micro的历史提交者，他们的代码永远运行在大家的内存中
- 感谢各位支持**StackLabs**中国发展的贡献者们

## 贡献者

### 主要维护者

<div style="border: 1px solid black">
    <table border="0">
      <tr>
        <td align="center"><a href="https://github.com/printfcoder"><img src="https://avatars0.githubusercontent.com/u/20906540?s=96&v=4" width="50px;" alt=""/><br /><sub><b>Printfcoder</b></sub></a></td>
        <td align="center"><a href="https://github.com/hb-chen"><img src="https://avatars2.githubusercontent.com/u/730866?s=96&v=4" width="50px;" alt=""/><br /><sub><b>hb-chen</b></sub></a></td>
        <td align="center"><a href="https://github.com/Allenxuxu"><img src="https://avatars2.githubusercontent.com/u/20566897?s=96&v=4" width="50px;" alt=""/><br /><sub><b>Allenxuxu</b></sub></a></td>
      </tr>
    </table>
</div>

### 顾问

<div style="border: 1px solid black">
    <table border="0">
      <tr>
          <td align="center"><a href="https://github.com/crazybber"><img src="https://avatars0.githubusercontent.com/u/3401462?s=96&v=4" width="50px;" alt=""/><br /><sub><b>crazybber</b></sub></a></td>
          <td align="center"><a href="https://github.com/kevwan"><img src="https://avatars2.githubusercontent.com/u/1918356?s=96&v=4" width="50px;" alt=""/><br /><sub><b>Kevin Wan</b></sub></a></td>
        </tr>
    </table>
</div>

### 社区主要维护者

<div style="border: 1px solid black">
    <table border="0">
      <tr>
        <td align="center"><a href="https://github.com/printfcoder"><img src="https://avatars0.githubusercontent.com/u/20906540?s=96&v=4" width="50px;" alt=""/><br /><sub><b>Printfcoder</b></sub></a></td>
        <td align="center"><a href="https://github.com/hb-chen"><img src="https://avatars2.githubusercontent.com/u/730866?s=96&v=4" width="50px;" alt=""/><br /><sub><b>hb-chen</b></sub></a></td>
        <td align="center"><a href="https://github.com/Allenxuxu"><img src="https://avatars2.githubusercontent.com/u/20566897?s=96&v=4" width="50px;" alt=""/><br /><sub><b>Allenxuxu</b></sub></a></td>
        <td align="center"><a href="https://github.com/crazybber"><img src="https://avatars0.githubusercontent.com/u/3401462?s=96&v=4" width="50px;" alt=""/><br /><sub><b>crazybber</b></sub></a></td>
      </tr>
    </table>
</div>