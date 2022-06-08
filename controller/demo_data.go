package controller

import "simple-demo/repository"

var DemoVideos = []repository.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://006b-121-22-29-122.jp.ngrok.io/video/guanyanglin.mp4",
		CoverUrl:      "/5.jpeg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []repository.Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = repository.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var DemoUser1 = repository.User{
	Id:            2,
	Name:          "TestUser1",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
