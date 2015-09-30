package vine

type popularWrapper struct {
	Data    popularPage `json:"data"`
	Success bool        `json:"success"`
	Error   string      `json:"error"`
}

type userWrapper struct {
	Data    *User  `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

//User is a collection of a Vine user's profile information.
type User struct {
	Username          string   `json:"username"`
	FollowerCount     int64    `json:"followerCount"`
	Verified          int      `json:"verified"`
	VanityUrls        []string `json:"vanityUrls"`
	LoopCount         int64    `json:"loopCount"`
	LoopVelocity      float64  `json:"loopVelocity"`
	AvatarUrl         string   `json:"avatarUrl"`
	AuthoredPostCount int64    `json:"authoredPostCount"`
	UserId            int64    `json:"userId"`
	UserIdStr         string   `json:"userIdStr"`
	PostCount         int64    `json:"postCount"`
	ProfileBackground string   `json:"profileBackground"`
	LikeCount         int64    `json:"likeCount"`
	Private           int      `json:"private"`
	Location          string   `json:"location"`
	FollowingCount    int64    `json:"followingCount"`
	ExplicitContent   int      `json:"explicitContent"`
	Description       string   `json:"description"`
}

type popularPage struct {
	AnchorStr string           `json:"anchorStr"`
	Records   []*PopularRecord `json:"records"`
	NextPage  int              `json:"nextPage"`
	Size      int              `json:"size"`
}

//PopularRecord is a small user record meant for unique user identification
//only...for now. This is bound to change in the future as the package grows
//to accommodate for the remaining Vine API endpoints.
type PopularRecord struct {
	UserId    int64  `json:"userId"`
	UserIdStr string `json:"userIdStr"`
}
