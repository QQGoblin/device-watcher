# 说明

该项目用于学习kubernetes控制器的开发流程。

## 1. 工程目录结构

- cmd：应用程序入口
- config
    - crds：controller-gen 生成的CRD文件
- hack：构建，代码生成等相关脚本
- pkg：代码目录
    - apis: CRD定义
    - client: go-client 代码
    
工程创建过程如下：

- 编写apis目录中相应CRD的Go Struct，并添加代码生成器相关注释
- 生成 Client、DeepCopy 代码
- 编写控制器逻辑

# Kubernetes 相关知识

## 1. GVR

GVKs 或 GVRs 指的是：GroupVersionKind 或 GroupVersionResource，本工程中：

- Group: 通过apis包中的各个doc.go文件的**+groupName**注释指定，即apis包的结构为：apis/<group>/<version>/...
- Kind: Go Struct对象的名称，对应go文件名称为 xxx_types.go

## 2. Controller 类的结构

参考：github.com/QQGoblin/device-watcher/pkg/controller/nic_controller.go

## 3. 初始化 Controller


	
