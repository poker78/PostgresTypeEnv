package PostgresTypeEnv

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// تابع برای اتصال به دیتابیس PostgreSQL با استفاده از جنریک
func ConnectPostgres[T any](config ConfigEnv[T]) *gorm.DB {
	// ساخت کانکشن استرینگ با استفاده از پارامترهای موجود در config
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.PostgresUrl, config.PostgresPort, config.PostgresUsername, config.PostgresPassword, config.PostgresDb, config.SslMode, config.TimeZone)

	// باز کردن اتصال با استفاده از Gorm
	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// بررسی خطاهای اتصال به دیتابیس
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Postgres: %v", err))
	}

	// دریافت شی دیتابیس از Gorm برای تنظیمات بیشتر
	sqlDb, err := dbClient.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get database object from Gorm: %v", err))
	}

	// بررسی ارتباط با دیتابیس
	err = sqlDb.Ping()
	if err != nil {
		panic(fmt.Sprintf("Database ping failed: %v", err))
	}

	// تنظیم حداکثر تعداد کانکشن‌های idle
	sqlDb.SetMaxIdleConns(config.SetMaxIdleConns)
	// تنظیم حداکثر تعداد کانکشن‌های باز
	sqlDb.SetMaxOpenConns(config.SetMaxOpenConns)
	// تنظیم طول عمر هر کانکشن
	sqlDb.SetConnMaxLifetime(time.Duration(config.SetConnMaxLifetime) * time.Hour)

	return dbClient
}
