// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"cool-admin-go/internal/service/internal/dao/internal"
)

// spaceTypeDao is the data access object for table space_type.
// You can define custom methods on it to extend its functionality as you wish.
type spaceTypeDao struct {
	*internal.SpaceTypeDao
}

var (
	// SpaceType is globally public accessible object for table space_type operations.
	SpaceType = spaceTypeDao{
		internal.NewSpaceTypeDao(),
	}
)

// Fill with you ideas below.
