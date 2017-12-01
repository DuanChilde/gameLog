### 环境
1. 确保安装bee工具以及beego环境
2. 编译安装protobuf编译器protoc `https://github.com/google/protobuf/releases`,将protoc命令加入环境变量
3. 下载golang的protobuf依赖 `go get github.com/golang/protobuf/protoc-gen-go` `go get github.com/golang/protobuf/proto`

### 启动
1. 根目录下执行`bee generate docs`，生成swagger数据（如果不需要swagger文档，可不执行这一步
2. 执行`bee run watchall true`，打开对应链接即可访问
