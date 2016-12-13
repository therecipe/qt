import QtQuick 2.0

Rectangle
{
	property string addr: "localhost:8833"
	property var hc

	width: 320
	height: 240

	Text
	{
    id: textView
		anchors.centerIn: parent
		text: "click here!"

		Connections
		{
			target: HelloClientFactory

			onInfo: textView.text = msg
			onError: textView.text = "Failed to " + err + " for " + addr + "\nWith error message: " + msg + "\n"
			onSuccess: textView.text = "Received: \"" + msg + "\"\n"
		}
  }

	MouseArea
	{
		anchors.fill: parent
		onClicked: hc.sayHello("world")
	}

	Component.onCompleted: hc = HelloClientFactory.newClient(addr)
	Component.onDestruction: hc.shutdown()
}
