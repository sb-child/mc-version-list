# mc-version-list

the version list of Minecraft (source: <https://zh.minecraft.wiki>)

Minecraft 版本列表, 来源: <https://zh.minecraft.wiki>

## BDS Version List (基岩版专用服务器 - 版本及协议版本列表)

### Json structure (Json 结构)

```jsonc
{
  // array of versions, from older to newer
  // 版本数组, 从旧到新
  "versions": [
    "1_9_0",
    "1_10_0",
    "1_20_51",
    "1_20_60_20",
    "1_20_60_21",
    "1_20_60_22",
    "1_20_60_23"
  ],
  // dict of versions
  // 版本字典
  "v": {
    "1_10_0": {
      // version 版本号
      "version": "1.10.0.7",
      // simple version 简化版本号
      "simple_version": "1.10.0",
      // article name 条目名
      "page": "基岩版1.10.0",
      // release version 是否为正式版
      "released": true,
      // protocol version 协议版本
      // -1 means unknown, -2 means currently not found
      // -1: 未知, -2: 目前未找到
      "protocol": 340,
      // major 主版本号
      "major_v": 1,
      // minor 次版本号
      "minor_v": 10,
      // build 构建版本号
      "build_v": 0,
      // revision 修订版本号
      "rev_v": 7
    },
    "1_9_0": {
      "version": "1.9.0.15",
      "simple_version": "1.9.0",
      "page": "基岩版1.9.0",
      "released": true,
      "protocol": 332,
      "major_v": 1,
      "minor_v": 9,
      "build_v": 0,
      "rev_v": 15
    }
  },
  // timestamp of this generation 此json的生成时间戳(UTC时间)
  "update_timestamp_ms": 1702752551478
}
```

### links (列表链接)

「sync every hour」(每小时同步一次)

- [「github gist」查看此 gist](https://gist.github.com/sb-child/b5a133b71e0a66b26360dff030906bc6)
- [「Raw Link」原始 json](https://gist.github.com/sb-child/b5a133b71e0a66b26360dff030906bc6/raw/mcvl-bds.json)
- [「Raw Link (ghproxy)」原始 json (中国境内专供)](https://mirror.ghproxy.com/https://gist.githubusercontent.com/sb-child/b5a133b71e0a66b26360dff030906bc6/raw/mcvl-bds.json)

## 本地使用

1. (手动编译) 拉取本 repo 的代码, 安装 `goframe cli`, 执行 `gf build` 编译代码, mcvl 程序会输出在 `bin/[系统架构]` 目录下

1. (或下载预编译版)

```bash
# linux
wget -O mcvl https://github.com/sb-child/mc-version-list/releases/latest/download/mcvl-linux-amd64 && chmod a+x mcvl
# windows
下载 https://github.com/sb-child/mc-version-list/releases/latest/download/mcvl-win-amd64.exe
```

1. 执行 `./mcvl generate` 将结果直接输出, 或者 `./mcvl generate -f=文件名` 将结果写入文件
