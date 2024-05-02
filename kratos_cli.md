#### 创建一个微服务
```
    kratos new app/name --nomod
```

#### 添加proto文件
```
    kratos proto add api/helloworld/v1/demo.proto
```

#### 生成proto代码
```
    kratos proto client api/helloworld/v1/demo.proto
    
    kratos proto server api/helloworld/v1/demo.proto -t internal/service  // -t 指定生成目录
```

#### 修改 conf.proto， 重新生成 conf.pb.go
```
    protoc --go_out=. -I=. conf.proto
```

#### 命令运行
```
    kratos run 
```