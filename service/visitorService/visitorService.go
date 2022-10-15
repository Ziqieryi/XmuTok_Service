package visitorService

import (
	"douyin/dao"
	commentImp "douyin/dao/commnetImp"
	"douyin/dao/favListImp"
	"douyin/dao/feedImp"
	"douyin/dao/loginImp"
	"douyin/module"
	"douyin/module/jsonModule/response"
	"douyin/tools"
)

func Login(username string, password string, response *response.Login) {

	data := new(module.UserTable)
	message := loginImp.Login(&username, &password, data)
	if message != "" {
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	// use tool class generate token
	token, err := tools.GenerateToken(data.UserId, username, password)

	if err != nil {
		response.StatusCode = -1
		response.StatusMsg = "Token Coding failure"
		return
	}
	//一切无误,装填信息
	response.Token = token
	response.UserId = data.UserId
	response.StatusCode = 0
	response.StatusMsg = "successful"
	return
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
		response.List = append(response.List, commentTemp)
	}
	response.StatusCode = 0
	response.StatusMsg = "successful"
}
func Feed(latestTime int64, response *response.Feed) {
	//游客登录状态
	//声明Table去数据库拿值,装填进response
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
		//有异常,装填response
		response.StatusCode = -1
		response.StatusMsg = message
		return
	} else {
		var videoTemp module.Video
		for i := 0; i < len(data); i++ {
			videoTemp.Id = data[i].VideoId
			videoTemp.Author.Id = data[i].UserId
			videoTemp.Author.Name = data[i].Username
			videoTemp.Author.IsFollow = false
			videoTemp.Author.FollowCount = data[i].FollowCount
			videoTemp.Author.FollowerCount = data[i].FollowerCount
			videoTemp.CommentCount = data[i].ComCount
			videoTemp.FavoriteCount = data[i].FavCount
			videoTemp.CoverUrl = data[i].CoverUrl
			videoTemp.IsFavorite = false
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
}
func FavList(userId int64, response *response.FavouriteList) {
	//游客登录状态
	//声明点赞列表和数据库对接的module,去数据库拿值
	var data []module.UserLikeVideoList
	message := favListImp.GetVideoList(userId, &data)
	if message != "" {
		//拿data过程有异常
		response.StatusCode = -1
		response.StatusMsg = message
		return
	}
	//data无误拿到,装填response
	var videoTemp module.Video
	for i := 0; i < len(data); i++ {
		videoTemp.Id = data[i].VideoId
		videoTemp.Author.Id = data[i].UserId
		videoTemp.Author.Name = data[i].Username
		videoTemp.Author.IsFollow = false
		videoTemp.Author.FollowCount = data[i].FollowCount
		videoTemp.Author.FollowerCount = data[i].FollowerCount
		videoTemp.CommentCount = data[i].ComCount
		videoTemp.FavoriteCount = data[i].FavCount
		videoTemp.CoverUrl = data[i].CoverUrl
		videoTemp.IsFavorite = false
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
