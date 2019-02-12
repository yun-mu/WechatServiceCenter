package constant

const (

	/****************************************** table name ****************************************/

	TableUser = "user"

	/****************************************** user ****************************************/

	UserUnFollowStatus = 0
	UserFollowStatus   = 5

	/****************************************** redis ****************************************/

	RedisDefaultExpire     = 3600 * 24 * 7 // 7天
	RedisDefaultRandExpire = 3600 * 24     // 1天

	RedisUserInfo = "user:info:%s" // format: user:info:<unionid>
)
