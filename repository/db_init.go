package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const preCount = 5 //预取出记录数
const Ngork = "https://bed6-121-22-29-122.jp.ngrok.io/video/"

var (
	UserIndexMap  map[string]DbUser
	Count         int64 //数据库总记录数
	PreTmpVideos  []*DbVideo
	PreVideos     []Video
	videoSequence int64
	preVideoCount = 30
	UserTmpVideos []*DbVideo
	UserVideos    []Video
	//FavoriteTmpVideos []*DbVideo
	FavoriteVideos []Video
	DbComments     []*DbComment
	Comments       []Comment
	Follows        []User
	Followers      []User
	//Favorite
)

type VideoList []*DbVideo

func (x VideoList) Len() int {
	return len(x)
}

func (x VideoList) Less(i, j int) bool {
	return x[j].CreateTime < x[i].CreateTime
}

func (x VideoList) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

var db, err = gorm.Open(mysql.Open("root:zlh123456@tcp(127.0.0.1:3306)/douyin"))

func init() {
	if err != nil {
		fmt.Println(err)
	}
	var Usertmp []DbUser
	db.AutoMigrate(&DbUser{})

	//预先从数据库取出preCount条记录放进内存，算是做预备缓存
	result := db.Table("db_users").Limit(preCount).Find(&Usertmp)

	if result.Error != nil {
		fmt.Println("find records error")
	}
	db.AutoMigrate(&DbVideo{})
	db.AutoMigrate(&DbFavoriteList{})
	UserVideos = []Video{}
	FavoriteVideos = []Video{}
	PreVideos = []Video{}
	PreTmpVideos = []*DbVideo{}
	DbComments = []*DbComment{}
	Comments = []Comment{}
	//UserIndexMap = DbUser{}

	//db.Create(DbVideo{Token: "zhanglinhui123456", PlayUrl: "guanyanglin.mp4", CreateTime: time.Now().Unix(), Id: 1})
	//initPreVideo()
	initUserCount()
	initUserIndexMap(Usertmp)
}

func initPreVideo() {
	//NewVideoDaoInstance().QueryVideoByCreateTimeFromDb()
}

func initUserCount() {
	NewUserDaoInstance().QueryUserCountFromDb()
}

func initUserIndexMap(users []DbUser) error {
	userTmpMap := make(map[string]DbUser)
	for _, v := range users {
		userTmpMap[v.Token] = v
	}
	UserIndexMap = userTmpMap
	return nil
}
