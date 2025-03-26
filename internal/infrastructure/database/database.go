package database

import (
	"airplane/internal/components/logger"
	"airplane/internal/config"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

func newDB(in digIn) *DB {
	return &DB{
		config: in.Config.Database,
		logger: in.Logger.SysLogger.Named("database"),
	}
}

// DB defines db connections.
type DB struct {
	*gorm.DB

	config *config.DatabaseConfig

	logger   logger.ILogger
	connLock *sync.Mutex
}

func (d *DB) Run(ctx context.Context, stop context.CancelFunc) error {
	d.initialize(ctx, d.config)
	d.logger.Info(ctx, "Database is running")
	return nil
}

func (d *DB) Shutdown(ctx context.Context) error {
	d.finalize()
	d.logger.Info(ctx, "Database is shutdown")
	return nil
}

// Initialize initializes models.
// It only creates the connection instance, doesn't reset or migrate anything.
func (d *DB) initialize(ctx context.Context, cfg *config.DatabaseConfig) {

	if d.connLock == nil {
		d.connLock = &sync.Mutex{}
	}
	d.connLock.Lock()
	defer d.connLock.Unlock()

	if d.DB == nil {
		d.logger.Debug(ctx, "Initializing database ...")
		d.DB = d.dialDB()
	}

	d.logger.Debug(ctx, "Done")
}

// Finalize closes the database.
func (d *DB) finalize() {
	d.connLock.Lock()
	defer d.connLock.Unlock()

	if d.DB != nil {
		sql, err := d.DB.DB()
		if err != nil {
			panic(err)
		}
		sql.Close()
		d.DB = nil
	}
}

func (d *DB) dialDB() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	db, err = d.connect()
	if err != nil {
		d.logger.Panic(nil, err)
	}
	return db
}

func (d *DB) connect() (db *gorm.DB, err error) {
	db, err = gorm.Open(d.config.Open(), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return
	}

	// Set database parameters.
	sql, err := db.DB()
	if err != nil {
		return nil, err
	}

	sql.SetMaxIdleConns(d.config.MaxIdleConns)
	sql.SetMaxOpenConns(d.config.MaxOpenConns)
	sql.SetConnMaxLifetime(time.Duration(d.config.MaxLifetime) * time.Second)

	return
}
