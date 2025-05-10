#!/bin/bash

# 创建 build 目录
mkdir -p build

# 判断操作系统类型
if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" || "$OSTYPE" == "win32" ]]; then
  # Windows 系统
  echo "Building for Windows..."
  go build -o build/memos.exe bin/memos/main.go
else
  # 非 Windows 系统 (Linux, macOS, etc.)
  echo "Building for $(uname -s)..."
  go build -o build/memos bin/memos/main.go
fi

# 检查编译是否成功
if [ $? -eq 0 ]; then
  echo "Build successful!"
  echo "You can now run the server using: ./build/memos --mode dev"
else
  echo "Build failed."
fi