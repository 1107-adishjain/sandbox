package storage
import(
	"gorm.io/gorm"
	"github.com/1107-adishjain/sandbox/config"
)

func NewConnection(cfg *config.Config) (*gorm.DB, error) {

	db,err := gorm.Open(nil, &gorm.Config{})
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