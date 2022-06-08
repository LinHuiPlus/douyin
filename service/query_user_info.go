package service

import (
	"fmt"
	"simple-demo/repository"
	"sync"
)

type UserInfo struct {
	User *repository.DbUser
}

func QueryUserInfo(token string) (*UserInfo, error) {
	return NewQueryUserInfoFlow(token).Do()
}

func NewQueryUserInfoFlow(token string) *QueryUserInfoFlow {
	return &QueryUserInfoFlow{
		userToken: token,
		userInfo:  nil,
		user:      nil,
	}
}

type QueryUserInfoFlow struct {
	userToken string
	userInfo  *UserInfo
	user      *repository.DbUser
}

func (f *QueryUserInfoFlow) Do() (*UserInfo, error) {
	if err := f.queryDbInfo(); err != nil {
		return nil, err
	}

	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packUserInfo(); err != nil {
		return nil, err
	}
	return f.userInfo, nil
}

func (f *QueryUserInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(1)
	//ans := true
	go func() {
		if user, exist := repository.NewUserDaoInstance().QueryUsersByTokenFromMemory(f.userToken); exist {
			f.user = &user
		}
		defer wg.Done()
	}()
	wg.Wait()
	if f.user == nil {
		return fmt.Errorf("user doesn't exist")
	}
	return nil
}

func (f *QueryUserInfoFlow) queryDbInfo() error {
	return repository.NewUserDaoInstance().QueryUserByTokenFromDb(f.userToken)
}

func (f *QueryUserInfoFlow) packUserInfo() error {
	f.userInfo = &UserInfo{
		User: f.user,
	}
	return nil
}
