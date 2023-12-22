package repo

import (
	"context"
	"mall/internal/entity"
)

// dao:mysql、redis、rpc
// 用户数据：来源于mysql？还是redis？（缓存），还是rpc？（商城规模较大，服务拆分：接入层、用户中心、订单中心、商品中心）

// Repository 仓库；数据存储的地方
// 面向接口编程
type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (rowId int, err error)
	GetUserByAccount(ctx context.Context, account, pwd string) (user entity.User, err error)
	FindUserById(ctx context.Context, userId int) (user entity.User, err error)
}
