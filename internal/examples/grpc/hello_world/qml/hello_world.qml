import QtQuick 2.0

Rectangle
{
	width: 320
	height: 240

	Text
	{
    id: textView
		anchors.centerIn: parent
		text: "click here!"
  }

	MouseArea
	{
		anchors.fill: parent

		onClicked:
		{
			var addr = "localhost:8833"

			textView.text = "Trying " + addr + "\n"

			var hc = HelloClientFactory.newClient(addr)

			if (hc == null)
			{
				textView.text = "Failed to create client for " + addr + "\n"
				return
			}

			var result = hc.sayHello("world")
			if (result.length == 2)
			{
				textView.text = "Failed to receive response for " + addr + "\nWith error message: " + result[1] + "\n"
				return
			}
			textView.text = "Received: \"" + result[0] + "\"\n"

			var exit = hc.shutdown()
			if (exit.length > 0)
			{
				textView.text = "Failed to shutdown client for " + addr + "\nWith error message: " + exit + "\n"
			}
		}
	}
}
