package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Meeting{})
	DB.AutoMigrate(&Relation{})
	DB.AutoMigrate(&Check{})
}
