# go-fill
[![LICENSE](https://img.shields.io/github/license/mashape/apistatus.svg?style=flat-square&label=License)](https://github.com/czasg/go-fill/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/czasg/go-fill/branch/main/graph/badge.svg?token=OkiSH6DMqf)](https://codecov.io/gh/czasg/go-fill)
[![GitHub Stars](https://img.shields.io/github/stars/czasg/go-fill.svg?style=flat-square&label=Stars&logo=github)](https://github.com/czasg/go-fill/stargazers)

## 背景
记得第一次接触go，当时在处理一段结构体解析上耗费了大量的精力。我对此感到很困惑，因为这在python中应该是很简单的一件事情。  
特别是`json.Marshal`时，当结构体中存在大量空值时，结果集中往往会出现大量nil属性，给前端同事带来不小的麻烦。   

go-fill 实现对结构体空值的填充，填充来源包括**零值、默认值、环境变量**等属性。

## 目标
1、填充环境变量
- [x] Env Value

2、填充默认值
- [x] Default Value

3、填充零值
- [x] Zero Value

## 标签说明
|tag|comment|
|---|---|
|fill:"fieldName"|字段名，若未指定，则默认是原字段大写。|
|fill:",default=value"|当字段最终判定为零值时，设定的默认值。|
|fill:",require"|当字段最终判定为零值时，返回错误。require在default之后生效。|
|fill:",empty"|设置当前字段名为空字符串.|
|fill:",sep=_"|嵌套结构体时，设定连接符，默认是下划线 "_"。|

## 使用
1、填充环境变量
```go
// 依赖
import "github.com/czasg/go-fill"
// 准备结构体
type Config struct {
	Host     string `fill:"HOST,default=localhost"`
	Port     int    `fill:"PORT,default=5432"`
	User     string `fill:"USER,default=root"`
	Password string `fill:"PASSWORD,default=root"`
}
// 初始化
cfg := Config{}
// 填充环境变量
_ = fill.FillEnv(&cfg)
_ = fill.Fill(&cfg, fill.OptEnv)
```

## Demo
```go
package main

import (
	"fmt"
	"github.com/czasg/go-fill"
	"os"
)

type Config struct {
	Postgres `env:"PG"`
	Redis    `env:"RDS"`
}

type Redis struct {
	Addr     string
	Password string
	DB       int
}

type Postgres struct {
	Addr     string
	User     string
	Password string
	Database string
}

func main() {
	_ = os.Setenv("PG_ADDR", "PG_ADDR")
	_ = os.Setenv("PG_USER", "PG_USER")
	_ = os.Setenv("PG_PASSWORD", "PG_PASSWORD")
	_ = os.Setenv("PG_DATABASE", "PG_DATABASE")
	_ = os.Setenv("RDS_ADDR", "RDS_ADDR")
	_ = os.Setenv("RDS_PASSWORD", "RDS_PASSWORD")
	_ = os.Setenv("RDS_DB", "RDS_DB")
	cfg := Config{}
	_ = fill.FillEnv(&cfg)
	fmt.Println(cfg)
}
```