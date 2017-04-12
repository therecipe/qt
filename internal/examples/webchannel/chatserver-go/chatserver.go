package main

import (
	"sort"

	"github.com/therecipe/qt/core"
)

type ChatServer struct {
	core.QObject

	_ func() `constructor:"init"`

	_ []string `property:"userList"`

	_ func(userName string) bool  `slot:"login"`
	_ func(userName string) bool  `slot:"logout"`
	_ func(user, msg string) bool `slot:"sendMessage"`
	_ func(user string)           `slot:"keepAliveResponse"`
	_ func()                      `slot:"sendKeepAlive"`
	_ func()                      `slot:"checkKeepAliveResponses"`

	_ func(time, user, msg string) `signal:"newMessage"`
	_ func()                       `signal:"keepAlive"`
	_ func()                       `signal:"userCountChanged"`

	m_userList            []string
	m_stillAliveUsers     []string
	m_keepAliveCheckTimer *core.QTimer
}

func (s *ChatServer) init() {

	t := core.NewQTimer(s)
	t.ConnectTimeout(s.sendKeepAlive)
	t.Start(10000)

	s.m_keepAliveCheckTimer = core.NewQTimer(s)
	s.m_keepAliveCheckTimer.SetSingleShot(true)
	s.m_keepAliveCheckTimer.SetInterval(2000)
	s.m_keepAliveCheckTimer.ConnectTimeout(s.checkKeepAliveResponses)

	s.ConnectLogin(s.login)
	s.ConnectLogout(s.logout)
	s.ConnectSendMessage(s.sendMessage)
	s.ConnectSendKeepAlive(s.sendKeepAlive)
	s.ConnectCheckKeepAliveResponses(s.checkKeepAliveResponses)
	s.ConnectKeepAliveResponse(s.keepAliveResponse)
}

func (s *ChatServer) login(userName string) bool {
	if s.m_keepAliveCheckTimer.IsActive() {
		s.m_keepAliveCheckTimer.Stop()
		s.m_stillAliveUsers = nil
	}

	for _, name := range s.m_userList {
		if name == userName {
			return false
		}
	}

	println("User logged in:", userName)
	s.m_userList = append(s.m_userList, userName)
	sort.Stable(sort.StringSlice(s.m_userList))
	s.UserListChanged(s.m_userList)
	s.UserCountChanged()
	return true
}

func (s *ChatServer) logout(userName string) bool {
	for i, name := range s.m_userList {
		if name == userName {
			s.m_userList = append(s.m_userList[:i], s.m_userList[i+1:]...)
			s.UserListChanged(s.m_userList)
			s.UserCountChanged()
			return true
		}
	}
	return false
}

func (s *ChatServer) sendMessage(user, msg string) bool {
	for _, name := range s.m_userList {
		if name == user {
			s.NewMessage(core.QTime_CurrentTime().ToString("HH:mm:ss"), user, msg)
			return true
		}
	}
	return false
}

func (s *ChatServer) sendKeepAlive() {
	s.KeepAlive()
	s.m_keepAliveCheckTimer.Start2()
}

func (s *ChatServer) checkKeepAliveResponses() {
	println("Keep Alive Check", s.m_stillAliveUsers)
	s.m_userList = s.m_stillAliveUsers
	s.m_stillAliveUsers = nil
	sort.Stable(sort.StringSlice(s.m_userList))
	s.UserListChanged(s.m_userList)
}

func (s *ChatServer) keepAliveResponse(user string) {
	s.m_stillAliveUsers = append(s.m_stillAliveUsers, user)
}
