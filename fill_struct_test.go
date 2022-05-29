package fill

import (
    "fmt"
    "os"
    "testing"
    "time"
)

type TestStructFillEnv struct {
    TestStructFillEnv1 TestStructFillEnv1 `fill:"TestStructFillEnv1"`
    TestStructFillEnv2 TestStructFillEnv2 `fill:"TestStructFillEnv2"`
}

type TestStructFillEnv1 struct {
    A                  string
    B                  int
    TestStructFillEnv2 TestStructFillEnv2 `fill:"xxx,sep=--"`
}

type TestStructFillEnv2 struct {
    A string
    B int
}

func TestStructFill(t *testing.T) {
    assert := assertWrap(t)
    {
        nowUnix := int(time.Now().Unix())
        randomA := fmt.Sprintf("%d", nowUnix)
        _ = os.Setenv("TestStructFillEnv1_A", randomA)
        _ = os.Setenv("TestStructFillEnv1_B", randomA)
        _ = os.Setenv("TestStructFillEnv1--xxx_A", randomA)
        _ = os.Setenv("TestStructFillEnv1--xxx_B", randomA)
        _ = os.Setenv("TestStructFillEnv2_A", randomA)
        _ = os.Setenv("TestStructFillEnv2_B", randomA)
        test := TestStructFillEnv{}
        err := Fill(&test, OptEnv)
        assert("test fill env", test.TestStructFillEnv1.A, randomA)
        assert("test fill env", test.TestStructFillEnv1.B, nowUnix)
        assert("test fill env", test.TestStructFillEnv1.TestStructFillEnv2.A, randomA)
        assert("test fill env", test.TestStructFillEnv1.TestStructFillEnv2.B, nowUnix)
        assert("test fill env", test.TestStructFillEnv2.A, randomA)
        assert("test fill env", test.TestStructFillEnv2.B, nowUnix)
        assert("test fill env", err, nil)
    }
    {
        _ = os.Setenv("RPC", "rpc")
        _ = os.Setenv("RPC_ADDR", "rpc")
        _ = os.Setenv("RPC_USER_NAME", "rpc")
        _ = os.Setenv("mysql_ADDR", "mysql")
        _ = os.Setenv("mysql_PASSWORD", "mysql")
        _ = os.Setenv("mysql_DB", "mysql")
        _ = os.Setenv("mysql_NAME", "mysqlname")
        _ = os.Setenv("RPC_USER_NAME", "rpc")
        _ = os.Setenv("RDS_ADDR", "redis")
        _ = os.Setenv("RDS_PASSWORD", "redis")
        _ = os.Setenv("RDS_DB", "1")
        _ = os.Setenv("RDS-UU_NAME", "redis")
        _ = os.Setenv("PG__ADDR", "postgres")
        _ = os.Setenv("PG__PASSWORD", "postgres")
        _ = os.Setenv("PG__DB", "postgres")
        _ = os.Setenv("PG_USER_NAME", "postgres")
        test := StructConfig{}
        err := Fill(&test, OptEnv)
        assert("fill config", test.RPC.Addr, "rpc")
        assert("fill config", test.RPC.Name, "rpc")
        assert("fill config", test.RPC.User.Name, "rpc")
        assert("fill config", test.MySQL.Addr, "mysql")
        assert("fill config", test.MySQL.Password, "mysql")
        assert("fill config", test.MySQL.DB, "mysql")
        assert("fill config", test.MySQL.User.Name, "mysqlname")
        assert("fill config", test.Redis.Addr, "redis")
        assert("fill config", test.Redis.Password, "redis")
        assert("fill config", test.Redis.DB, 1)
        assert("fill config", test.Redis.User.Name, "redis")
        assert("fill config", test.Postgres.Addr, "postgres")
        assert("fill config", test.Postgres.Password, "postgres")
        assert("fill config", test.Postgres.DB, "postgres")
        assert("fill config", test.Postgres.User.Name, "postgres")
        assert("fill config", err, nil)
    }
}

type StructConfig struct {
    RPC
    MySQL    `fill:"mysql"`
    Redis    `fill:"RDS"`
    Postgres `fill:"PG"`
}

type RPC struct {
    Addr string
    Name string `fill:",empty"`
    User
}

type MySQL struct {
    Addr     string
    Password string
    DB       string
    User     `fill:",empty"`
}

type Redis struct {
    Addr     string
    Password string
    DB       int
    User     `fill:"UU,sep=-"`
}

type Postgres struct {
    Addr     string `fill:"ADDR,sep=__"`
    Password string `fill:"PASSWORD,sep=__"`
    DB       string `fill:"DB,sep=__"`
    User
}

type User struct {
    Name string
}
