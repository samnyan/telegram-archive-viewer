package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/rs/zerolog"
	"github.com/samnyan/go-telegram-archive/server/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Orm *gorm.DB
}

func Initialize(config *config.Config) *DB {
	l := NewDatabaseLogger()

	orm, err := gorm.Open(sqlite.Open(config.Database.Dsn), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		panic(fmt.Errorf("error when init gorm: %s", err))
	}
	return &DB{
		Orm: orm,
	}
}

type databaseLogger struct {
	log *zerolog.Logger
}

func NewDatabaseLogger() *databaseLogger {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &databaseLogger{
		log: &l,
	}
}

func (l *databaseLogger) LogMode(level logger.LogLevel) logger.Interface {
	var ll zerolog.Logger
	switch level {
	case logger.Silent:
		ll = l.log.Level(zerolog.Disabled)
	case logger.Error:
		ll = l.log.Level(zerolog.ErrorLevel)
	case logger.Warn:
		ll = l.log.Level(zerolog.WarnLevel)
	case logger.Info:
		ll = l.log.Level(zerolog.InfoLevel)
	}
	l.log = &ll
	return l
}

func (l *databaseLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if e := l.log.Info(); e.Enabled() {
		e.Ctx(ctx).Any("data", data).Msg(msg)
	}
}

func (l *databaseLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if e := l.log.Warn(); e.Enabled() {
		e.Ctx(ctx).Any("data", data).Msg(msg)
	}
}

func (l *databaseLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if e := l.log.Error(); e.Enabled() {
		e.Ctx(ctx).Any("data", data).Msg(msg)
	}
}

func (l *databaseLogger) Trace(ctx context.Context, begin time.Time, f func() (sql string, rowsAffected int64), err error) {

	var event *zerolog.Event

	if err != nil {
		event = l.log.Debug()
	} else {
		event = l.log.Trace()
	}

	zl := event.Ctx(ctx)

	var dur_key string

	switch zerolog.DurationFieldUnit {
	case time.Nanosecond:
		dur_key = "elapsed_ns"
	case time.Microsecond:
		dur_key = "elapsed_us"
	case time.Millisecond:
		dur_key = "elapsed_ms"
	case time.Second:
		dur_key = "elapsed"
	case time.Minute:
		dur_key = "elapsed_min"
	case time.Hour:
		dur_key = "elapsed_hr"
	default:
		zl.Interface("zerolog.DurationFieldUnit", zerolog.DurationFieldUnit).Msg("gormzerolog encountered a mysterious, unknown value for DurationFieldUnit")
		dur_key = "elapsed_"
	}

	event.Dur(dur_key, time.Since(begin))

	sql, rows := f()
	if sql != "" {
		event.Str("sql", sql)
	}
	if rows > -1 {
		event.Int64("rows", rows)
	}

	event.Send()

}
