package dbcopy

import (
	"fmt"

	"gorm.io/gorm"
)

func CopySqlData[model any](
	srcDB, dstDB *gorm.DB,
	filter func(*gorm.DB) *gorm.DB,
	mode Mode,
	backupTableName string,
) error {
	var src []model
	if err := srcDB.Scopes(filter).Find(&src).Error; err != nil {
		return fmt.Errorf("read src failed: %w", err)
	}

	if len(src) == 0 {
		return fmt.Errorf("no data to copy")
	}

	if backupTableName != "" {
		if err := BackupSqlData[model](dstDB, filter, backupTableName); err != nil {
			return err
		}
	}

	return dstDB.Transaction(func(tx *gorm.DB) error {
		switch mode {
		case ModeBasic:
			var dst []model
			if err := tx.Scopes(filter).Find(&dst).Error; err != nil {
				return fmt.Errorf("read destination failed: %w", err)
			}
			if len(dst) > 0 {
				return fmt.Errorf("destination already has data")
			}

		case ModeReplace:
			var empty model
			if err := tx.Scopes(filter).Delete(&empty).Error; err != nil {
				return fmt.Errorf("failed to delete from destination: %w", err)
			}

		case ModeAppend:
			// no pre-check, just insert

		default:
			return fmt.Errorf("unsupported copy mode: %v", mode)
		}

		if err := tx.CreateInBatches(&src, BatchSize).Error; err != nil {
			return fmt.Errorf("failed to insert data: %w", err)
		}

		return nil
	})
}

func BackupSqlData[model any](
	db *gorm.DB,
	filter func(*gorm.DB) *gorm.DB,
	backupTableName string,
) error {
	var records []model
	if err := db.Scopes(filter).Find(&records).Error; err != nil {
		return fmt.Errorf("failed to fetch records for backup: %w", err)
	}
	if len(records) == 0 {
		return nil
	}

	backupDB := db.Table(backupTableName)

	if !db.Migrator().HasTable(backupTableName) {
		schemaModel := new(model)
		if err := backupDB.Migrator().CreateTable(schemaModel); err != nil {
			return fmt.Errorf("failed to create backup table: %w", err)
		}
	} else {
		return fmt.Errorf("backup table already exists")
	}

	if err := backupDB.CreateInBatches(&records, 500).Error; err != nil {
		return fmt.Errorf("failed to insert records into backup table: %w", err)
	}

	return nil
}
