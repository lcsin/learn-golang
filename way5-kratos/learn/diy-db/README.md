# 使用自定义的 ORM 框架
主要修改`internal\data\data.go`文件

# 以使用`sqlx`为例
1. 引入 sqlx 和 mysql 驱动：
```go
_ "github.com/go-sql-driver/mysql"
"github.com/jmoiron/sqlx"
```
2. 修改`data.Data`注入`sqlx.DB`
```go
type Data struct {
    // TODO wrapped database client
    db  *sqlx.DB
    log *log.Helper
}
```
3. 创建 DB Provider 
```go
// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDB)

// NewDB .
    func NewDB(c *conf.Data, logger log.Logger) *sqlx.DB {
    log := log.NewHelper(log.With(logger, "module", "data/sqlx"))
    
    db, err := sqlx.Connect(c.Database.Driver, c.Database.Source)
    if err != nil {
    log.Fatal("connect DB failed, err: %v\n", err)
    }
    
    return db
}

// NewData .
func NewData(db *sqlx.DB, logger log.Logger) (*Data, func(), error) {
    log := log.NewHelper(log.With(logger, "module", "data"))
    
    cleanup := func() {
    log.Info("closing the data resources")
    }
    
    return &Data{
        db:  db,
        log: log,
    }, cleanup, nil
}
```
4. 在 cmd 目录使用`wire`命令重新生成`initApp()`函数
