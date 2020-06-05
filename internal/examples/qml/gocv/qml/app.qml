import QtQuick 2.0
import QtQuick.Controls 2.13

import CustomModule 1.0

Item {
	width: 300; height: 200

	GoCV {
		anchors.fill: parent

		Button {
			onClicked: parent.detect = !parent.detect
			text: parent.detect ? "Stop detecting" : "Start detecting"
		}
	}
}
