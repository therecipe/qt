package main

import (
	"github.com/therecipe/qt/androidextras"
	"github.com/therecipe/qt/core"
)

type NotificationClient struct {
	core.QObject

	_ func() `constructor:"init"`

	_ string `property:"notification"`

	_ func(string) `slot:"updateAndroidNotification"`
}

func (c *NotificationClient) init() {
	c.ConnectNotificationChanged(c.updateAndroidNotification)
}

func (c *NotificationClient) updateAndroidNotification(n string) {

	var err = androidextras.QAndroidJniObject_CallStaticMethodVoid2Caught(
		"org/qtproject/example/notification/NotificationClient",
		"notify",
		"(Ljava/lang/String;)V",
		n,
	)

	if err != nil {
		println(err.Error())
	}
}
