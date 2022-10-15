package userService

import (
	"douyin/dao"
	commentImp "douyin/dao/commnetImp"
	"douyin/dao/favListImp"
	"douyin/dao/favouriteImp"
	"douyin/dao/feedImp"
	"douyin/dao/followImp"
	"douyin/dao/infoImp"
	"douyin/dao/publishImp"
	"douyin/dao/relationImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
	"fmt"
	"mime/multipart"
	"time"
)

func Feed(latestTime int64, token string, response *response.Feed) {
	//用户登录状态
	//解析token拿到用户信息
	user, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Token Encryption failed"
		return
	}
	//token无误
	//声明视频流需要到数据库里拿的module,去数据库拿值
	var data []module.VideoWithAuthor
	var message string
	if latestTime > 0 {
		//限制时间戳
		message = feedImp.Feed2(latestTime, &data)
	} else {
		//没有限制时间戳
		message = feedImp.Feed1(&data)
	}
	if message != "" {
		//拿data过程有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data无误拿到
	if len(data) < 1 {
		response.StatusCode = -1
		response.StatusMsg = "没有更多视频了"
		return
	}
	//根据userId查用户对data里的视频是否喜欢
	var flag bool
	var isFav []bool
	for i := 0; i < len(data); i++ {
		flag, message = feedImp.Feed3(user.UserId, data[i].VideoId)
		if message != "" {
			break
		}
		isFav = append(isFav, flag)
	}
	if message != "" {
		//查是否喜欢过程中有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//根据userid查用户对data里视频的作者是否关注
	var isFol []bool
	for i := 0; i < len(data); i++ {
		flag, message = feedImp.Feed4(data[i].AuthorId, user.UserId)
		if message != "" {
			break
		}
		isFol = append(isFol, flag)
	}
	if message != "" {
		//查是否喜欢过程中有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data,isFav,isFol无误拿到,装填response
	var videoTemp module.Video
	for i := 0; i < len(data); i++ {
		videoTemp.Id = data[i].VideoId
		videoTemp.Author.Id = data[i].UserId
		videoTemp.Author.Name = data[i].Username
		videoTemp.Author.IsFollow = isFol[i]
		videoTemp.Author.FollowCount = data[i].FollowCount
		videoTemp.Author.FollowerCount = data[i].FollowerCount
		videoTemp.CommentCount = data[i].ComCount
		videoTemp.FavoriteCount = data[i].FavCount
		videoTemp.CoverUrl = data[i].CoverUrl
		videoTemp.IsFavorite = isFav[i]
		videoTemp.PlayUrl = data[i].PlayUrl
		videoTemp.VideoTitle = data[i].VideoTitle
		videoTemp.Author.Signature = data[i].Signature
		videoTemp.Author.BackgroundImage = data[i].BackgroundImage
		videoTemp.Author.Avatar = data[i].Avatar
		response.List = append(response.List, videoTemp)
	}
	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.NextTime = data[len(data)-1].UploadDate
}
func PublishAction(data *multipart.FileHeader, token string, title string, response *response.PublishAction) {
	//根据token解析出来的内容判断内容是否存在
	tmp, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Encryption failed"
		return
	}
	usertable := new(module.UserTable)
	exist := publishImp.QueryUserId(tmp.UserId, usertable)
	if exist != nil {
		response.StatusCode = -1
		response.StatusMsg = exist.Error()
		return
	}
	fileContent, _ := data.Open()

	play_url := tools.GetPlayUrl(data.Filename, fileContent)
	cover_url := tools.GetCoverUrl(data.Filename)
	err = publishImp.InsertData(tmp.UserId, play_url, cover_url, title)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}
	// 成功的返回值
	response.StatusCode = 0
	response.StatusMsg = "successful"
	return
}
func Register(username string, password string, response *response.Register) {
	//创建用户
	u := module.UserTable{
		UserId:          0,
		Username:        username,
		Password:        password,
		Signature:       "欢迎使用抖声APP",
		Avatar:          "https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20Avatar/3.jpeg",
		BackgroundImage: "https://yygh-lamo.oss-cn-beijing.aliyuncs.com/User%20background/defaultBackGround.png",
		FollowerCount:   0,
		FollowCount:     0,
	}
	if err := dao.Db.Create(&u).Error; err != nil {
		response.StatusCode = -1
		response.StatusMsg = "The username already exists"
		return
	}
	x := module.UserTable{}
	if err := dao.Db.Order("user_id desc").First(&x); err != nil {
		response.StatusCode = -1
		response.StatusMsg = "意外错误"
	}
	token, _ := tools.GenerateToken(x.UserId, username, password)
	fmt.Printf("%v\n", x)
	response.Token = token
	response.UserId = x.UserId
	response.StatusCode = 0
	response.StatusMsg = "successful"
	return
}

// PublishList all users have same publish video list
func PublishList(token string, userId string, response *response.PublishList) {
	// 查询该作者的所有video
	var videoList []*module.VideoTable
	if err := publishImp.QueryVideoByUserId(userId, &videoList); err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询作者信息
	author := new(module.UserTable)
	if err := publishImp.QueryAuthorByUserId(userId, author); err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询用户是否关注了此作者
	// 先把token里面的userId解析出来，不同于作者的userId
	userClaims, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}
	// 在根据用户id和作者id查询是否关注
	var follow *module.FollowTable
	isFollow, err := publishImp.IsFollow(userClaims.UserId, userId, follow)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询该用户是否给video点赞
	var fav []*module.FavTable
	isFavList, err := publishImp.IsFavorite(userClaims.UserId, videoList, fav)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 该作者和其所有的 video 成功查询后，填装response
	authorResp := new(module.User)
	authorResp.Id = author.UserId
	authorResp.Name = author.Username
	authorResp.FollowCount = author.FollowCount
	authorResp.FollowerCount = author.FollowerCount
	authorResp.IsFollow = isFollow

	videoListResp := make([]module.Video, len(videoList))
	for i := 0; i < len(videoList); i++ {
		videoListResp[i].Id = videoList[i].VideoId
		videoListResp[i].Author = *authorResp
		videoListResp[i].PlayUrl = videoList[i].PlayUrl
		videoListResp[i].CoverUrl = videoList[i].CoverUrl
		videoListResp[i].FavoriteCount = videoList[i].FavCount
		videoListResp[i].CommentCount = videoList[i].ComCount
		videoListResp[i].IsFavorite = isFavList[i]
		videoListResp[i].VideoTitle = videoList[i].VideoTitle
	}

	response.StatusCode = 0
	response.StatusMsg = "success"
	response.VideoList = videoListResp
	return
}
func UserInfo(token string, userId string, response *response.UserInfo) {
	//查询用户信息
	// 先把token里面的userId解析出来，不同于作者的userId
	userClaims, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}
	// 在根据用户id和要查询的userid查询是否关注（也就是用户id是否关注了要查询的userid）
	var follow *module.FollowTable
	isFollow, _ := infoImp.IsFollow(userClaims.UserId, userId, follow)

	userTable := new(module.UserTable)
	err = infoImp.SelectAuthorByUserId(userId, userTable)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "The user information does not exist"
		return
	}
	response.StatusCode = 0
	response.Id = userTable.UserId
	response.Name = userTable.Username
	response.IsFollow = isFollow
	response.FollowCount = userTable.FollowCount
	response.FollowerCount = userTable.FollowerCount
	response.Signature = userTable.Signature
	response.BackgroundImage = userTable.BackgroundImage
	response.Avatar = userTable.Avatar
	response.StatusMsg = "successful"
	return
}
func UserFav(userId int64, videoId int64, actionType int64, response *response.Favourite) {
	//根据actionType进行点赞服务或者取消点赞服务
	if actionType == 1 {
		//点赞
		//将点赞记录同步更新到数据库
		mes := favouriteImp.Insert(userId, videoId)
		if mes != "" {
			response.StatusCode = -1
			response.StatusMsg = mes
			return
		}
		response.StatusCode = 0
		response.StatusMsg = "点赞成功"
		return
	}
	if actionType == 2 {
		//取消点赞
		//将点赞记录从数据库里同步删除
		mes := favouriteImp.Delete(userId, videoId)
		if mes != "" {
			response.StatusCode = -1
			response.StatusMsg = mes
			return
		}
		response.StatusCode = 0
		response.StatusMsg = "取消点赞成功"
		return
	}
	//actionType意外的值错误
	response.StatusCode = -1
	response.StatusMsg = "ActionType value is invalid"
	return
}
func UserFol(followId int64, followerId int64, actionType int64, response *response.Follow) {
	//根据actionType提供关注服务或者取消关注服务
	if actionType == 1 {
		//关注
		//将关注记录同步更新到数据库
		mes := followImp.Insert(followId, followerId)
		if mes != "" {
			response.StatusCode = -1
			response.StatusMsg = mes
			return
		}
		response.StatusCode = 0
		response.StatusMsg = "关注成功"
		return
	}
	if actionType == 2 {
		//取消关注
		//将关注记录从数据库里同步删除
		mes := followImp.Delete(followId, followerId)
		if mes != "" {
			response.StatusCode = -1
			response.StatusMsg = mes
			return
		}
		response.StatusCode = 0
		response.StatusMsg = "取消关注成功"
		return
	}
	//actionType意外的值错误
	response.StatusCode = -1
	response.StatusMsg = "ActionType value is invalid"
	return
}
func FavList(userId int64, token string, response *response.FavouriteList) {
	//解析token拿visitorId
	user, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Token Encryption failed"
		return
	}
	//token无误
	//声明点赞列表和数据库对接的module,去数据库拿值
	var data []module.UserLikeVideoList
	message := favListImp.GetVideoList(userId, &data)
	if message != "" {
		//拿data过程有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data无误拿到
	if len(data) < 1 {
		response.StatusCode = -1
		response.StatusMsg = "还没有点赞过视频"
		return
	}
	//根据visitorId查用户对data里的视频是否喜欢
	var flag bool
	var isFav []bool
	for i := 0; i < len(data); i++ {
		flag, message = favListImp.IsFav(user.UserId, data[i].VideoId)
		if message != "" {
			break
		}
		isFav = append(isFav, flag)
	}
	if message != "" {
		//查是否喜欢过程中有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//根据userid查用户对data里视频的作者是否关注
	var isFol []bool
	for i := 0; i < len(data); i++ {
		flag, message = favListImp.IsFollow(data[i].AuthorId, user.UserId)
		if message != "" {
			break
		}
		isFol = append(isFol, flag)
	}
	if message != "" {
		//查是否喜欢过程中有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data,isFav,isFol无误拿到,装填response
	var videoTemp module.Video
	for i := 0; i < len(data); i++ {
		videoTemp.Id = data[i].VideoId
		videoTemp.Author.Id = data[i].UserId
		videoTemp.Author.Name = data[i].Username
		videoTemp.Author.IsFollow = isFol[i]
		videoTemp.Author.FollowCount = data[i].FollowCount
		videoTemp.Author.FollowerCount = data[i].FollowerCount
		videoTemp.CommentCount = data[i].ComCount
		videoTemp.FavoriteCount = data[i].FavCount
		videoTemp.CoverUrl = data[i].CoverUrl
		videoTemp.IsFavorite = isFav[i]
		videoTemp.PlayUrl = data[i].PlayUrl
		videoTemp.VideoTitle = data[i].VideoTitle
		videoTemp.Author.Signature = data[i].Signature
		videoTemp.Author.BackgroundImage = data[i].BackgroundImage
		videoTemp.Author.Avatar = data[i].Avatar
		response.List = append(response.List, videoTemp)
	}
	response.StatusCode = 0
	response.StatusMsg = "successful"
}
func FollowList(token string, userId int64, response *response.FollowList) {
	// 根据userId查询出该用户的关注列表
	var userList []module.UserTable
	if err := relationImp.QueryUserById(userId, &userList); err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询当前用户是否关注了列表中的用户
	isFolList := make([]bool, len(userList))
	if token != "" {
		userClaims, err := tools.AnalyseToken(token)
		if err != nil {
			response.StatusCode = -1
			response.StatusMsg = err.Error()
			return
		}
		isFolList, err = relationImp.IsFollow(userClaims.UserId, userList)
		if err != nil {
			response.StatusCode = -1
			response.StatusMsg = err.Error()
			return
		}
	}

	userListResp := make([]module.User, len(userList))
	for i := 0; i < len(userList); i++ {
		userListResp[i].Id = userList[i].UserId
		userListResp[i].Name = userList[i].Username
		userListResp[i].FollowCount = userList[i].FollowCount
		userListResp[i].FollowerCount = userList[i].FollowerCount
		userListResp[i].IsFollow = isFolList[i]
		userListResp[i].Signature = userList[i].Signature
		userListResp[i].Avatar = userList[i].Avatar
		userListResp[i].BackgroundImage = userList[i].BackgroundImage
	}

	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.UserList = userListResp
}

func FollowerList(token string, userId int64, response *response.FollowList) {
	// 根据userId查询出该用户的粉丝列表
	var userList []module.UserTable
	if err := relationImp.QueryFollwerListUserById(userId, &userList); err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	// 查询当前用户是否关注了列表中的用户
	isFolList := make([]bool, len(userList))
	if token != "" {
		userClaims, err := tools.AnalyseToken(token)
		if err != nil {
			response.StatusCode = -1
			response.StatusMsg = err.Error()
			return
		}
		isFolList, err = relationImp.IsFollwer(userClaims.UserId, userList)
		if err != nil {
			response.StatusCode = -1
			response.StatusMsg = err.Error()
			return
		}
	}
	userListResp := make([]module.User, len(userList))
	for i := 0; i < len(userList); i++ {
		userListResp[i].Id = userList[i].UserId
		userListResp[i].Name = userList[i].Username
		userListResp[i].FollowCount = userList[i].FollowCount
		userListResp[i].FollowerCount = userList[i].FollowerCount
		userListResp[i].IsFollow = isFolList[i]
		userListResp[i].Signature = userList[i].Signature
		userListResp[i].BackgroundImage = userList[i].BackgroundImage
		userListResp[i].Avatar = userList[i].Avatar
	}

	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.UserList = userListResp
}

// add comment
func AddComment(token string, videoId int64, commentText string, response *response.CommentActionResponse) {
	tmp, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Encryption failed"
		return
	}
	usertable := new(module.UserTable)
	exist := commentImp.QueryUserId(tmp.UserId, usertable)
	if exist != nil {
		response.StatusCode = -1
		response.StatusMsg = " user is not login"
		return
	}
	createtime := time.Now().Format("2006-01-02")

	// insert CommentMsg
	//err = commentImp.InsertCommentMsg(videoId, userId, commentText, time)
	err = commentImp.InsertCommentMsg(videoId, tmp.UserId, commentText, createtime)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}

	CommentTable := new(module.CommentTable)
	// Query CommentMsg
	err = commentImp.QueryCommentMsgRes(videoId, tmp.UserId, createtime, CommentTable)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
	}

	nusertable := new(module.UserTable)
	err = commentImp.QueryUserId(tmp.UserId, nusertable)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
	}

	response.StatusCode = 0
	response.StatusMsg = "successful"
	response.CommentId = CommentTable.CommentId
	response.User.Id = usertable.UserId
	response.Name = nusertable.Username
	response.FollowCount = nusertable.FollowCount
	response.FollowerCount = nusertable.FollowerCount
	response.Content = CommentTable.Content
	response.CreateDate = createtime
	return
}

// delete comment
func DeleteComment(token string, commentId int64, response *response.CommentActionResponse) {
	_, err := tools.AnalyseToken(token)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Encryption failed"
		return
	}
	CommentTable := new(module.CommentTable)
	// 目前的删除策略是根据评论id来进行删除(当前情况默认考虑评论唯一)
	exist := commentImp.DeleteCommentImp(commentId, CommentTable)
	if exist != nil {
		response.StatusCode = -1
		response.StatusMsg = exist.Error()
		return
	}
}

func ComList(videoId int64, response *response.CommentList) {
	var data []module.CommentTable
	err := commentImp.GetCommentList(videoId, &data)
	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = err.Error()
		return
	}
	//装填response
	var commentTemp module.Comment
	for i := 0; i < len(data); i++ {
		commentTemp.Id = data[i].CommentId
		commentTemp.Content = data[i].Content
		commentTemp.CreateDate = data[i].CreateDate
		//查询评论对应的用户
		var user module.UserTable
		userId := data[i].ComUserId
		if err := dao.Db.Where("user_id = ?", userId).Find(&user).Error; err != nil {
			response.StatusCode = -1
			response.StatusMsg = "The username already exists"
			return
		}
		commentTemp.User.Id = user.UserId
		commentTemp.User.Name = user.Username
		commentTemp.User.FollowCount = user.FollowCount
		commentTemp.User.FollowerCount = user.FollowerCount
		commentTemp.User.Avatar = user.Avatar
		response.List = append(response.List, commentTemp)
	}
	response.StatusCode = 0
	response.StatusMsg = "successful"
}
