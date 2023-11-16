# IP地址地理位置匹配工具 🌍

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/charSLee013/IPGeoCIDR/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/charSLee013/IPGeoCIDR)](https://goreportcard.com/report/github.com/charSLee013/IPGeoCIDR)

该项目是一个用于查询IP地址地理位置的工具，它可以根据提供的CIDR或CIDR文件，查询其中每个IP地址所属的国家，并将匹配的IP地址保存到输出文件中.

## 功能特点 ✨

- 支持输入CIDR或CIDR文件进行IP地址解析
- 可以使用正则表达式过滤指定国家的IP地址
- 使用qqwry数据库进行IP地址的查询
- 输出结果保存到指定的文件中

## 快速开始 🚀

### 编译项目

1. 克隆项目到本地：
```
git clone https://github.com/charSLee013/IPGeoCIDR.git
```

2. 进入项目目录：
```
cd your-project
```

3. 编译项目：
```
go build .
```

### 使用示例 📝

#### 查询单个CIDR的IP地址

```
go run . -cidr "192.168.0.0/24" -output "geo_ips.txt"
```

#### 查询CIDR文件中的IP地址

```
go run . -cidr "cidr.txt" -output "geo_ips.txt"
```

#### 指定国家的正则表达式

```
go run . -cidr "192.168.0.0/24" -country "缅甸" -output "geo_ips.txt"
```

#### 指定多个国家的正则表达式

```
go run . -cidr "192.168.0.0/24" -country "(日本|韩国|新加坡)" -output "geo_ips.txt"
```

#### 指定Cloudflare CDN IP的国家
```shell
# 下载cloudflare ip range
wget -O ips.txt https://www.cloudflare.com/ips-v4/#
go run . -cidr ips.txt -country "(日本|韩国|新加坡)" -output "geo_ips.txt"
```

## 常见问题 ❓

1. **如何添加自定义qqwry数据库？**

   默认情况下，程序将使用内置的qqwry数据库。如果您想使用自己的数据库，可以在程序中修改`db_path`变量，将其指向您自己的qqwry数据库文件.
   或者直接将其置于二进制同位置下的同名文件 `qqwry.dat`即可

2. **如何解析IPv6地址？**

   该工具当前仅支持IPv4地址的解析，暂不支持IPv6地址.

3. **如何导出结果到CSV文件？**

   目前该工具仅支持输出为文本文件（每行一个IP地址）。如果您需要导出为CSV文件，可以在输出文件后缀名为`.csv`，然后使用逗号分隔每个字段.

## 许可证 📄

该项目基于 MIT 许可证。详细信息请参阅 [LICENSE](https://github.com/charSLee013/IPGeoCIDR/blob/main/LICENSE) 文件.
