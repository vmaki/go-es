package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

type migrationFunc func(gorm.Migrator, *sql.DB)

var migrationFiles []MigrationFile

type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

// Add 新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}

// getMigrationFile 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(name string) MigrationFile {
	for _, mfile := range migrationFiles {
		if name == mfile.FileName {
			return mfile
		}
	}
	return MigrationFile{}
}

// isNotMigrated 判断迁移是否已执行
func (m MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == m.FileName {
			return false
		}
	}

	return true
}
