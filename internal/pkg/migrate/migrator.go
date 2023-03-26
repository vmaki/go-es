package migrate

import (
	"go-es/global"
	"go-es/internal/pkg/console"
	"go-es/internal/pkg/database"
	"go-es/internal/pkg/filex"
	"gorm.io/gorm"
	"os"
)

// Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       global.GDB,
		Migrator: global.GDB.Migrator(),
	}

	// migrations 不存在的话就创建它
	migrator.createMigrationsTable()

	return migrator
}

// 创建 migrations 表
func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}

	// 不存在才创建
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

// Up 执行所有未迁移过的文件
func (migrator *Migrator) Up() {
	// 读取所有迁移文件，确保按照时间排序
	migrateFiles := migrator.readAllMigrationFiles()

	// 获取当前批次的值
	batch := migrator.getBatch()

	// 获取所有迁移数据
	var migrations []Migration
	migrator.DB.Find(&migrations)

	// 可以通过此值来判断数据库是否已是最新
	isNew := false

	// 对迁移文件进行遍历，如果没有执行过，就执行 up 回调
	for _, mFile := range migrateFiles {
		// 对比文件名称，看是否已经运行过
		if mFile.isNotMigrated(migrations) {
			migrator.runUpMigration(mFile, batch)
			isNew = true
		}
	}

	if !isNew {
		console.Success("database is up to date.")
	}
}

// 获取当前这个批次的值
func (migrator *Migrator) getBatch() int {
	batch := 1

	// 取最后执行的一条迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	// 如果有值的话，加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}

	return batch
}

// 从文件目录读取文件，保证正确的时间排序
func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	// 读取 database/migrations/ 目录下的所有文件
	// 默认是会按照文件名称进行排序
	files, err := os.ReadDir(migrator.Folder)
	if err != nil {
		console.Exit(err.Error())
	}

	var migrateFiles []MigrationFile
	for _, f := range files {
		// 去除文件后缀 .go
		fileName := filex.FileNameWithoutExtension(f.Name())

		// 通过迁移文件的名称获取『MigrationFile』对象
		f := getMigrationFile(fileName)

		// 加个判断，确保迁移文件可用，再放进 migrateFiles 数组中
		if len(f.FileName) > 0 {
			migrateFiles = append(migrateFiles, f)
		}
	}

	// 返回排序好的『MigrationFile』数组
	return migrateFiles
}

// 执行迁移，执行迁移的 up 方法
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	// 执行 up 区块的 SQL
	if mfile.Up != nil {
		// 友好提示
		console.Warning("migrating " + mfile.FileName)

		// 执行 up 方法
		mfile.Up(global.GDB.Migrator(), database.SqlDB)

		// 提示已迁移了哪个文件
		console.Success("migrated " + mfile.FileName)
	}

	// 入库
	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	if err != nil {
		console.Exit(err.Error())
	}
}

// Rollback 回滚上一个操作
func (migrator *Migrator) Rollback() {
	// 获取最后一批次的迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	var migrations []Migration
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)

	// 回滚最后一批次的迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

// 回退迁移，按照倒序执行迁移的 down 方法
func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {
	// 标记是否真的有执行了迁移回退的操作
	flag := false

	for _, _migration := range migrations {
		// 友好提示
		console.Warning("rollback " + _migration.Migration)

		// 执行迁移文件的 down 方法
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(global.GDB.Migrator(), database.SqlDB)
		}

		flag = true

		// 回退成功了就删除掉这条记录
		migrator.DB.Delete(&_migration)

		// 打印运行状态
		console.Success("finish " + mfile.FileName)
	}

	return flag
}

// Reset 回滚所有迁移
func (migrator *Migrator) Reset() {
	var migrations []Migration

	// 按照倒序读取所有迁移文件
	migrator.DB.Order("id DESC").Find(&migrations)

	// 回滚所有迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to reset.")
	}
}

// Refresh 回滚所有迁移，并运行所有迁移
func (migrator *Migrator) Refresh() {
	// 回滚所有迁移
	migrator.Reset()

	// 再次执行所有迁移
	migrator.Up()
}

// Fresh Drop 所有的表并重新运行所有迁移
func (migrator *Migrator) Fresh() {
	// 获取数据库名称，用以提示
	dbname := database.CurrentDatabase()

	// 删除所有表
	err := database.DeleteAllTables()
	if err != nil {
		console.Exit(err.Error())
	}

	console.Success("clearup database " + dbname)

	// 重新创建 migrates 表
	migrator.createMigrationsTable()
	console.Success("[migrations] table created.")

	// 重新调用 up 命令
	migrator.Up()
}
