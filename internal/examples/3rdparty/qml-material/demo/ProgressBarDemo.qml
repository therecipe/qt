import QtQuick 2.4
import QtQuick.Layouts 1.1
import Material 0.2

Item {
    implicitHeight: grid.implicitHeight + dp(40)
    GridLayout {
        id: grid
        anchors.centerIn: parent
        columns: 2
        columnSpacing: dp(20)
        rowSpacing: dp(20)

        Label {
            text: "Determinate"
        }

        ProgressBar {
            Layout.fillWidth: true
            color: theme.accentColor

            SequentialAnimation on value {
                running: true
                loops: NumberAnimation.Infinite

                NumberAnimation {
                    duration: 3000
                    from: 0
                    to: 1
                }

                PauseAnimation { duration: 1000 } // This puts a bit of time between the loop
            }
        }

        Label {
            text: "Indeterminate"
        }

        ProgressBar {
            Layout.fillWidth: true
            color: theme.accentColor
            indeterminate: true
        }

        Label {
            text: "Default"
        }

        ProgressCircle {
            Layout.alignment: Qt.AlignCenter
        }

        Label {
            text: "Custom Color"
        }

        ProgressCircle {
            Layout.alignment: Qt.AlignCenter
            color: "#E91E63"
        }

        Label {
            text: "Cyclic Colors"
        }

        ProgressCircle {
            id: cyclicColorProgress
            Layout.alignment: Qt.AlignCenter
            SequentialAnimation {
                running: true
                loops: Animation.Infinite

                ColorAnimation {
                    from: "red"
                    to: "blue"
                    target: cyclicColorProgress
                    properties: "color"
                    easing.type: Easing.InOutQuad
                    duration: 2400
                }

                ColorAnimation {
                    from: "blue"
                    to: "green"
                    target: cyclicColorProgress
                    properties: "color"
                    easing.type: Easing.InOutQuad
                    duration: 1560
                }

                ColorAnimation {
                    from: "green"
                    to: "#FFCC00"
                    target: cyclicColorProgress
                    properties: "color"
                    easing.type: Easing.InOutQuad
                    duration:  840
                }

                ColorAnimation {
                    from: "#FFCC00"
                    to: "red"
                    target: cyclicColorProgress
                    properties: "color"
                    easing.type: Easing.InOutQuad
                    duration:  1200
                }
            }
        }

        Label {
            text: "Custom Size"
        }

        ProgressCircle {
            Layout.alignment: Qt.AlignCenter
            width: dp(64)
            height: dp(64)
        }

        Label {
            text: "Custom Size + Thickness"
        }

        ProgressCircle {
            Layout.alignment: Qt.AlignCenter
            width: dp(64)
            height: dp(64)
            dashThickness: dp(8)
        }

        Label {
            text: "Determinate Value"
        }

        ProgressCircle {
            id: determinateProgress
            Layout.alignment: Qt.AlignCenter
            width: dp(64)
            height: dp(64)
            indeterminate: false
            minimumValue: 0
            maximumValue: 100

            SequentialAnimation on value {
                running: true
                loops: NumberAnimation.Infinite

                NumberAnimation {
                    duration: 3000
                    from: determinateProgress.minimumValue
                    to: determinateProgress.maximumValue
                }

                PauseAnimation { duration: 1000 }
            }

            Label {
                anchors.centerIn: parent
                text: Math.round(determinateProgress.value) + "%"
            }
        }
    }
}
