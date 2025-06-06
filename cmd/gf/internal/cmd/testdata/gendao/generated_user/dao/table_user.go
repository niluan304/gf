// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"for-gendao-test/pkg/dao/internal"
)

// tableUserDao is the data access object for the table table_user.
// You can define custom methods on it to extend its functionality as needed.
type tableUserDao struct {
	*internal.TableUserDao
}

var (
	// TableUser is a globally accessible object for table table_user operations.
	TableUser = tableUserDao{internal.NewTableUserDao()}
)

// Add your custom methods and functionality below.
