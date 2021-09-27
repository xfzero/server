package base

type AdminUser struct {
	Id            uint64
	Username      string
	LastLoginTime uint64
	Group         []uint32
}

type UserInfo struct {
	Id            uint64
	Username      string
	LastLoginTime uint64
	Group         []uint32
	Access        []uint32
	Token         string
}
