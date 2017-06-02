import QtQuick 2.4
import Material 0.2
import Material.Extras 0.1

Item {

    Column {
        anchors.centerIn: parent
        spacing: dp(20)

        Button {
            anchors.horizontalCenter: parent.horizontalCenter

            text: "Push subpage"
            elevation: 1
            onClicked: pageStack.push(Qt.resolvedUrl("SubPage.qml"))
        }

        Button {
            anchors.horizontalCenter: parent.horizontalCenter

            text: "Push subpage with sidebar"
            elevation: 1
            onClicked: pageStack.push(Qt.resolvedUrl("SidebarPage.qml"))
        }

        Image {
            id: image

            anchors.horizontalCenter: parent.horizontalCenter

            source: Qt.resolvedUrl("images/balloon.jpg")
            width: dp(400)
            height: dp(250)

            Ink {
                anchors.fill: parent

                onClicked: overlayView.open(image)
            }
        }

        Label {
            anchors.horizontalCenter: parent.horizontalCenter

            style: "subheading"
            color: Theme.light.subTextColor
            text: "Tap to edit picture"
            font.italic: true
        }
    }

    OverlayView {
        id: overlayView

        width: dp(800)
        height: dp(500)

        Image {
            id: contentImage
            source: Qt.resolvedUrl("images/balloon.jpg")
            anchors.fill: parent
        }

        Row {
            anchors {
                top: parent.top
                right: parent.right
                rightMargin: dp(16)
            }
            height: dp(48)
            opacity: overlayView.transitionOpacity

            spacing: dp(24)

            Repeater {
                model: ["content/add", "image/edit", "action/delete"]

                delegate: IconButton {
                    id: iconAction

                    iconName: modelData

                    color: Theme.dark.iconColor
                    size: iconName == "content/add" ? dp(27) : dp(24)
                    anchors.verticalCenter: parent.verticalCenter
                }
            }
        }
    }
}
