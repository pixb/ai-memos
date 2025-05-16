# Proto Requirements

# 开发流程

- 创建proto文件夹
- 确定系统安装了`protobuf_buf`工具
==buf.yaml==
- 在`proto`文件夹，执行`buf mod init`生成`buf.yaml`文件
- 在`buf.yaml`中`version: v2`第二版本
- 在`buf.yaml`添加`deps` `- buf.build/googleapis/googleapis`是`Google protobuf标准API依赖`
- 在`buf.yaml`的`lint`,`use` `- BASIC`: 基础的`lint`检查
- 在`buf.yaml`的`lint`,`except`排除规则
	- `- ENUM_VALUE_PREFIX`: 枚举命名约束。
	- `- FIELD_NOT_REQUIRED`：默认`optional`.
	- `- PACKAGE_DIRECTORY_MATCH`: 目录必须与包名同样目录。
	- `- PACKAGE_NO_IMPORT_CYCLE`：循环导入依赖。
	- `- PACKAGE_VERSION_SUFFIX`：包名后缀必须包含版本。
- 在`buf.yaml`包含`disallow_comment_ignores: true`: 禁止使用行内注释。
- 在`buf.yaml`中添加`breaking`: 检测是否存在破坏性`API`
- `breaking` `use:` `- FILE`: 文件级别破坏性检测
- `breaking` `except:` 破坏性排除
	- `- EXTENSION_NO_DELETE`: 禁止删除 proto2 语法中的扩展字段
	- `- FIELD_SAME_DEFAULT`: 字段默认值改变检查。

==buf.gen.yaml==文件是 Buf 工具用于**代码生成**的配置文件
- `version: v2`: 格式版本，v2` 是当前推荐的版本。
- `managed:` `enabled: true` 管理模式，统一处理生成代码模式。
- `disabled:` 禁用选项
	- `- file_option: go_package`
		- `module: buf.build/googleapis/googleapis`
		- 设置`buf`不要管理来自 `buf.build/googleapis/googleapis` 的 `go_package` 选项。
	- `file_option: go_package_prefix`
		- `value: github.com/usermemos/memos/proto/gen`
		- 这确保了所有生成的 Go 代码都统一存放在项目预期的父包路径下。
- `plugins:` 生成代码的插件。

```yaml
plugins:
  - remote: buf.build/protocolbuffers/go
    out: gen
    opt: paths=source_relative
  - remote: buf.build/grpc/go
    out: gen
    opt: paths=source_relative
  - remote: buf.build/grpc-ecosystem/gateway
    out: gen
    opt: paths=source_relative
  - remote: buf.build/grpc-ecosystem/openapiv2
    out: gen
    opt: output_format=yaml,allow_merge=true
  - remote: buf.build/community/stephenh-ts-proto
    out: ../web/src/types/proto
    opt:
      - env=browser
      - useOptionals=messages
      - outputServices=generic-definitions
      - outputJsonMethods=false
      - useExactTypes=false
      - esModuleInterop=true
      - stringEnums=true
```