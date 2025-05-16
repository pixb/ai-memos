package store

import (
	"sync"

	"github.com/pixb/memos/server/profile"
)

// 提供对数据库访问的对象
type Store struct {
	Profile               *profile.Profile
	driver                Driver
	workspaceSettingCache sync.Map // map[string] *storepb.WorkspaceSetting
	userCache             sync.Map // map[int]*User
	userSettingCache      sync.Map // map[string]*storepb.UserSetting
	idpCache              sync.Map // map[int]*storepb.IdentityProvider
}

// New 创建一个新的 Store 对象
func New(driver Driver, profile *profile.Profile) *Store {
	return &Store{
		Profile: profile,
		driver:  driver,
	}
}

func (s *Store) GetDriver() Driver {
	return s.driver
}

// Close 关闭数据库连接
func (s *Store) Close() error {
	return s.driver.Close()
}
