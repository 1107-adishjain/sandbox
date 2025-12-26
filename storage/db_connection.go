package storage
import(
	"gorm.io/gorm"
	"fmt"
	"gorm.io/driver/postgres"
	"github.com/1107-adishjain/sandbox/config"
)

func NewConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn:= fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s connect_timeout=10 sslmode=prefer", 
	cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPassword)
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		return nil,err
	}
	return db, nil
}

func CloseConnection(db *gorm.DB) error{
	sqlDB, err := db.DB()
	if err != nil{
		return err
	}
	return sqlDB.Close()
}