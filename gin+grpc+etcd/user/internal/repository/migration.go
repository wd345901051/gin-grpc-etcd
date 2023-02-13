package repository

import "fmt"

// 迁移函数
func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&User{},
		)
	if err != nil {
		fmt.Println("migration err", err)
	}
}
