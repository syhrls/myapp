package utils

import (
	"reflect"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RegisterGlobalUUIDHook sets a global hook to generate UUID for any model with ID field of type uuid.UUID
func RegisterGlobalUUIDHook(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("global:generate_uuid", func(tx *gorm.DB) {
		if tx.Statement.Schema != nil {
			idField := tx.Statement.Schema.LookUpField("ID")
			if idField != nil && idField.FieldType == reflect.TypeOf(uuid.UUID{}) {
				val, ok := idField.ValueOf(tx.Statement.Context, tx.Statement.ReflectValue)
				if ok && val == uuid.Nil {
					idField.Set(tx.Statement.Context, tx.Statement.ReflectValue, uuid.New())
				}
			}
		}
	})
}
