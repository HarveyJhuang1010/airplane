```text
◆ 連線配置
與datasource配置相同，當Application啟動時，flyway將會進行連線，並執行migration
‣ spring.flyway.url
‣ spring.flyway.user
‣ spring.flyway.password

◆ 檔案存取位置
可接受多個參數(請用逗號分隔)，指定flyway讀取執行migration的路徑
‣ spring.flyway.locations

◆ 歷史紀錄表定義
Flyway歷史紀錄表預設命名為flyway_schema_history，若有需要可以更改名稱
‣ spring.flyway.table

◆ 是否執行起始版號
當資料庫不為空，是否要執行起始版本，並建立歷史紀錄表，預設為false，
如果並非在專案一開始就導入flyway，就需要設定為true
‣ spring.flyway.baseline-on-migrate

◆ 起始版本設定
設定migration的起始版號
‣ spring.flyway.baseline-version

◆ 執行migration是否允許無序執行
‣ spring.flyway.out-of-order

◆ 執行migration是否自動驗證
‣ spring.flyway.validate-on-migrate

◆ 是否啟用Flyway
‣ spring.flyway.enabled

```


flyway.yml文件配置：
```text
 enabled: true #是否开启flyway，默认true.
 baseline-on-migrate: true #当迁移时发现目标schema非空，而且带有没有元数据的表时，是否自动执行基准迁移，默认false.
 clean-on-validation-error: true #当发现校验错误时是否自动调用clean，默认false.
 sql-migration-prefix: V #迁移文件的前缀，默认为V
 sql-migration-suffixes: .sql #移脚本的后缀，默认为.sql
 locations: -classpath:db/migration/mysql #迁移脚本的位置，默认db/migration.可以多个，可以给每个环境使用不同位置
 ignore-missing-migrations: true #忽略缺失的升级脚本验证
 flyway.baseline-description对执行迁移时基准版本的描述.
 flyway.baseline-on-migrate当迁移时发现目标schema非空，而且带有没有元数据的表时，是否自动执行基准迁移，默认false.
 flyway.baseline-version开始执行基准迁移时对现有的schema的版本打标签，默认值为1.
 flyway.check-location检查迁移脚本的位置是否存在，默认false.
 flyway.clean-on-validation-error当发现校验错误时是否自动调用clean，默认false.
 flyway.enabled是否开启flywary，默认true.
 flyway.encoding设置迁移时的编码，默认UTF-8.
 flyway.ignore-failed-future-migration当读取元数据表时是否忽略错误的迁移，默认false.
 flyway.init-sqls当初始化好连接时要执行的SQL.
 flyway.locations迁移脚本的位置，默认db/migration.
 flyway.out-of-order是否允许无序的迁移，默认false.
 flyway.password目标数据库的密码.
 flyway.placeholder-prefix设置每个placeholder的前缀，默认${.
 flyway.placeholder-replacementplaceholders是否要被替换，默认true.
 flyway.placeholder-suffix设置每个placeholder的后缀，默认}.
 flyway.placeholders.[placeholder name]设置placeholder的value
 flyway.schemas设定需要flywary迁移的schema，大小写敏感，默认为连接默认的schema.
 flyway.sql-migration-prefix迁移文件的前缀，默认为V.
 flyway.sql-migration-separator迁移脚本的文件名分隔符，默认__
 flyway.sql-migration-suffix迁移脚本的后缀，默认为.sql
 flyway.tableflyway使用的元数据表名，默认为schema_version
 flyway.target迁移时使用的目标版本，默认为latest version
 flyway.url迁移时使用的JDBC URL，如果没有指定的话，将使用配置的主数据源
 flyway.user迁移数据库的用户名
 flyway.validate-on-migrate迁移时是否校验，默认为true.
```

