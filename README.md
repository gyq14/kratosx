# kratosx

> 本项目来源：[https://github.com/limes-cloud/kratosx](https://github.com/limes-cloud/kratosx)

## 项目简介

kratosx 是基于 [Go](https://golang.org/) 语言和 [Kratos](https://go-kratos.dev/) 微服务框架的增强工具集，旨在提升 Go 微服务开发的效率和体验，提供一站式的项目脚手架、代码生成、服务治理、配置管理、中间件、常用库等能力。

## 主要功能

- **项目脚手架**：一键生成标准化微服务项目结构，支持多种布局和模板。
- **Proto 代码生成**：支持 proto 文件的服务端、客户端、API 模板、错误码等多种代码生成。
- **命令行工具**：丰富的 CLI 工具，涵盖项目创建、运行、升级、变更日志、Web 工具等。
- **中间件与库**：内置认证、JWT、限流、链路追踪、日志、数据库、缓存、邮箱、验证码、签名、监控等常用中间件和库。
- **多数据库支持**：支持 MySQL、PostgreSQL、SQLServer、ClickHouse 等。
- **服务治理**：内置注册发现、健康检查、Prometheus 监控、pprof 性能分析等。
- **配置热加载**：支持配置文件热加载与动态变更。

## 为什么要选择使用 Protocol Buffers？

- **效率高**：Protobuf 设计出来就是为了高效，序列化后的数据小，处理速度快。
- **强类型**：由于定义时需要指定数据类型，因此序列化和反序列化时能够确保类型安全。
- **兼容性好**：Protobuf 被设计为向前和向后兼容，即使数据结构发生变化，新旧版本的数据定义也能够互相理解。
- **跨平台和语言无关**：Protobuf 支持多种编程语言，能够轻松实现不同语言或平台间的数据交换。
- **自动化代码生成**：Protobuf 编译器能够自动生成不同编程语言的数据访问类，减少手动编码。
- **清晰的结构定义**：Protobuf 使得数据结构在 .proto 文件中一目了然，便于管理和维护。

## 安装 Protocol Buffers

### 下载
1. 访问 Protocol Buffers 的 GitHub 仓库页面：[protobuf](https://github.com/protocolbuffers/protobuf)
2. 点击“Releases”标签，找到最新的稳定版本。
3. 根据你的操作系统下载对应的预编译二进制文件或源代码包。


### 安装
- **Windows**：
  - 解压下载的压缩包。
  - 找到 bin 目录，将 protoc.exe 文件的路径添加到系统环境变量的 PATH 中。
- **macOS / Linux**：
  - 解压下载的压缩包。
  - 在终端中，使用 sudo 将 protoc 文件复制到 /usr/local/bin/ 目录下。
    ```bash
    sudo cp protoc /usr/local/bin/protoc
    ```
  - 确保 /usr/local/bin 在你的 PATH 环境变量中。

### 验证是否安装成功
```bash
protoc --version
```

## 安装 kratosx 及常用插件

kratosx cli 的主要功能是进行项目创建、proto 代码生成（grpc、http、error 代码等）。此 cli 工具在 kratos 原 cli 工具上做了一些调整和功能的增加，你可以使用 kratosx cli 更加高效地开发你的项目。

执行安装命令：
```bash
go install github.com/gyq14/kratosx/cmd/kratosx@latest 
go install github.com/gyq14/kratosx/kratosx/cmd/protoc-gen-go-httpx@latest 
go install github.com/gyq14/kratosx/kratosx/cmd/protoc-gen-go-errorsx@latest 
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest 
go install github.com/envoyproxy/protoc-gen-validate@latest
```

如果出现执行失败，检测一下自己网络是否可以访问 github。

各插件作用说明：
- **protoc-gen-go-httpx**：生成服务中的 http 端代码
- **protoc-gen-go-errorsx**：用于生成全局 error
- **protoc-gen-go**：生成 pb 代码
- **protoc-gen-go-grpc**：生成 grpc 代码
- **protoc-gen-openapi**：生成 openapi.yaml
- **protoc-gen-validate**：生成验证器，用于参数校验

### 验证是否完成安装
```bash
kratosx -v
```

## 命令行工具用法

### kratosx 主命令

```bash
kratosx [子命令] [参数]
```

#### 常用子命令

- `new`：创建微服务项目模板
  - 示例：`kratosx new helloworld`
  - 支持参数：`--repo-url`（模板仓库）、`--branch`（分支）、`--timeout`（超时）、`--nomod`（保留 go.mod）
- `proto`：Proto 相关代码生成
  - `add`：添加 proto API 模板
    - 示例：`kratosx proto add helloworld/v1/hello.proto`
  - `client`：生成 proto 客户端代码
    - 示例：`kratosx proto client api/helloworld.proto`
    - 支持参数：`--proto_path`、`--out_path`
  - `server`：生成 proto 服务端实现
    - 示例：`kratosx proto server api/helloworld.proto --target-dir=internal/service`
    - 支持参数：`--target-dir`
- `run`：运行项目
  - 示例：`kratosx run`
  - 支持参数：`--work`（指定工作目录）
- `upgrade`：升级 kratosx 及相关插件
  - 示例：`kratosx upgrade`
- `changelog`：获取 kratosx 变更日志或指定版本的 release 信息
  - 示例：`kratosx changelog dev` 或 `kratosx changelog v2.7.3`
  - 支持参数：`--repo-url`（仓库地址）
- `webutil`：启动 Web 工具服务
  - 示例：`kratosx webutil 8080`

## kratosx 的两种使用方式

**方式一：本地编译输出**
你可以将 kratosx 工具本身编译输出到本地 out 目录，便于本地管理和操作。例如：
- 编译输出 kratosx
  ```bash
  go build -o ./out/kratosx ./cmd/kratosx
  ./out/kratosx [子命令] [参数]
  ```
- 编译输出 protoc-gen-go-httpx：
  ```bash
  go build -o ./out/protoc-gen-go-httpx ./cmd/protoc-gen-go-httpx
  ./out/protoc-gen-go-httpx [参数]
  ```
- 编译输出 protoc-gen-go-errorsx：
  ```bash
  go build -o ./out/protoc-gen-go-errorsx ./cmd/protoc-gen-go-errorsx
  ./out/protoc-gen-go-errorsx [参数]
  ```
这样可以在 out 目录下本地管理和调用所有相关工具，无需全局安装。

**方式二：全局安装使用**
推荐通过 go install 命令将 kratosx 及相关插件安装到 GOPATH/bin 或 PATH 路径下，实现全局调用：
```bash
go install github.com/gyq14/kratosx/cmd/kratosx@latest
go install github.com/gyq14/kratosx/cmd/protoc-gen-go-httpx@latest
go install github.com/gyq14/kratosx/cmd/protoc-gen-go-errorsx@latest

kratosx [子命令] [参数]
```
安装后可在任意目录下直接使用 kratosx 命令进行项目创建、代码生成等操作。

## 典型使用流程

1. **创建新项目**
   ```bash
   kratosx new myservice
   cd myservice
   go generate ./...
   go build -o ./bin/ ./...
   ./bin/myservice -conf ./configs
   ```

2. **添加 proto API**
   ```bash
   kratosx proto add helloworld/v1/hello.proto
   ```
3. **生成 proto 客户端/服务端代码**
   ```bash
   kratosx proto client api/helloworld.proto
   kratosx proto server api/helloworld.proto --target-dir=internal/service
   ```
4. **运行服务**
   ```bash
   kratosx run
   ```
5. **升级工具链**
   ```bash
   kratosx upgrade
   ```


## License

本项目遵循 MIT 协议。

---

> 本项目源码来源于 [https://github.com/limes-cloud/kratosx](https://github.com/limes-cloud/kratosx)
