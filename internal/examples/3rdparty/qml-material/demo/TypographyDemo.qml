import QtQuick 2.4
import QtQuick.Layouts 1.1
import Material 0.2
import Material.ListItems 0.1 as ListItem

Item {

    implicitHeight: column.height

    property var styles: [
        "display4",
        "display3",
        "display2",
        "display1",
        "headline",
        "title",
        "subheading",
        "body2",
        "body1",
        "caption",
        "button"
    ]

    ColumnLayout {
        id: column
        spacing: dp(16)

        ListItem.Subheader {
            text: "Font Weights"
        }

        Label {
            font.family: "Roboto"
            font.weight: Font.Light
            text: "Roboto Light"
            font.pixelSize: dp(34)

            anchors {
                left: parent.left
                margins: dp(16)
            }
        }

        Label {
            font.family: "Roboto"
            text: "Roboto Regular"
            font.pixelSize: dp(34)

            anchors {
                left: parent.left
                margins: dp(16)
            }
        }

        Label {
            font.family: "Roboto"
            font.weight: Font.DemiBold
            text: "Roboto Medium"
            font.pixelSize: dp(34)

            anchors {
                left: parent.left
                margins: dp(16)
            }
        }

        Label {
            font.family: "Roboto"
            font.weight: Font.Bold
            text: "Roboto Bold"
            font.pixelSize: dp(34)

            anchors {
                left: parent.left
                margins: dp(16)
            }
        }

        ListItem.Subheader {
            text: "Label Styles"
        }

        Repeater {
            model: styles
            delegate: Row {
                anchors {
                    left: parent.left
                    margins: dp(16)
                }

                Label {
                    text: modelData
                    width: dp(100)
                }

                Label {
                    style: modelData
                    text: {
                        var text = fontInfo["font"].substring(0,1).toUpperCase() + fontInfo["font"].substring(1)

                        if (style == "button")
                            text += " (ALL CAPS)"

                        text += " " + fontInfo["size"] + "sp"

                        if (fontInfo.size_desktop) {
                            text += " (Device), " + fontInfo["font"].substring(0,1).toUpperCase() + fontInfo["font"].substring(1)

                            if (style == "button")
                                text += " (ALL CAPS)"

                            text += " " + fontInfo["size_desktop"] + "sp (Desktop)"
                        }

                        return text
                    }
                }
            }
        }
    }
}
