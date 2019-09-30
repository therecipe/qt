import Felgo 3.0
import QtQuick 2.0
import "../common"
import QtQuick.Layouts 1.1

FlickablePage {
  title: "Settings"
  rightBarItem: ActivityIndicatorBarItem { opacity: dataModel.loading ? 1 : 0 }

  flickable.contentWidth: width
  flickable.contentHeight: content.height + content.y * 2

  ColumnLayout {
    id: content
    y: spacing * 2
    width: parent.width
    spacing: dp(10)

    // cache / data buttons
    AppButton {
      Layout.alignment: Qt.AlignHCenter
      text: "Update Conference Data"
      enabled: !dataModel.loading
      onClicked: {
        amplitude.logEvent("Update Conference Data")
        logic.loadData()
      }
      verticalMargin: 0
    }

    AppButton {
      Layout.alignment: Qt.AlignHCenter
      text: "Clear Cached Data"
      enabled: false//dataModel.loaded && !dataModel.loading
      visible: false
      onClicked: {
        amplitude.logEvent("Clear Cache")
        logic.clearCache()
      }
      verticalMargin: 0
    }

    // separator
    Item {
      Layout.fillWidth: true
      Layout.preferredHeight: parent.spacing * 2
      Rectangle {
        width: parent.width * 0.4
        height: px(1)
        color: Theme.tintColor
        anchors.centerIn: parent
      }
    }

    // enable / disable notifications
    Row {
      id: notificationRow
      Layout.alignment: Qt.AlignHCenter
      spacing: parent.spacing

      AppText {
        anchors.verticalCenter: parent.verticalCenter
        text: "Session Reminder:"
        wrapMode: Text.NoWrap
      }

      AppSwitch {
        anchors.verticalCenter: parent.verticalCenter
        checked: dataModel.notificationsEnabled
        updateChecked: false
        onToggled: {
          amplitude.logEvent(!checked ? "Check Session Reminder" : "Uncheck Session Reminder")
          logic.setNotificationsEnabled(!checked)
        }
      } // AppSwitch
    }

    AppText {
      Layout.preferredWidth: parent.width - 2 * dp(Theme.navigationBar.defaultBarItemPadding)
      Layout.maximumWidth: parent.width - 2 * dp(Theme.navigationBar.defaultBarItemPadding)
      Layout.alignment: Qt.AlignHCenter
      text: "Sends a local push notification 10 minutes before a favorited session is starting."
      color: Theme.secondaryTextColor
      wrapMode: Text.WordWrap
      horizontalAlignment: Text.AlignHCenter
      font.pixelSize: sp(14)
    }

    // separator
    Item {
      Layout.fillWidth: true
      Layout.preferredHeight: parent.spacing * 2
      Rectangle {
        width: parent.width * 0.4
        height: px(1)
        color: Theme.tintColor
        anchors.centerIn: parent
      }
    }

    // style switch
    Row {
      Layout.alignment: Qt.AlignHCenter
      spacing: parent.spacing

      AppText {
        anchors.verticalCenter: parent.verticalCenter
        text: "Theme: "
        wrapMode: Text.NoWrap
      }

      AppButton {
        anchors.verticalCenter: parent.verticalCenter
        property string target: Theme.platform !== "ios" ? "iOS" : "Android"
        text: system.isPlatform(System.IOS) && target == "Android" ? "Custom" : target
        onClicked: {
          amplitude.logEvent("Change Theme", {"theme" : target, "platform" : (system.isPlatform(System.IOS) ? "iOS" : "Android")})
          Theme.platform = target.toLowerCase()
        }
        flat: false
        verticalMargin: 0
      }

      AppButton {
        text: "Reset"
        onClicked: {
          amplitude.logEvent("Reset Theme", {"oldTheme" : (Theme.platform == "ios" ? "iOS" : "Android")})
          Theme.reset()
        }
        flat: true
        verticalMargin: 0
      }
    }

    // separator
    Item {
      Layout.fillWidth: true
      Layout.preferredHeight: parent.spacing * 2
      Rectangle {
        width: parent.width * 0.4
        height: px(1)
        color: Theme.tintColor
        anchors.centerIn: parent
      }
    }

    // tint color
    AppText {
      Layout.alignment: Qt.AlignHCenter
      text: "Tint Color"
      wrapMode: Text.NoWrap
    }

    Flow {
      Layout.maximumWidth: parent.width
      Layout.alignment: Qt.AlignHCenter
      spacing: parent.spacing

      // default color
      ColorButton {
        color: Theme.isIos ? "#007aff" : "#3f5ab5"
        onClicked: {
          Theme.colors.tintColor = Qt.binding(function() { return Theme.isIos ? "#007aff" : (Theme.isAndroid ? "#3f51b5" : "#01a9e2")})
          amplitude.logEvent("Change Tint Color", {"color" : Theme.colors.tintColor, "index" : 0})
        }
      }

      // other colors
      Repeater {
        model: ["red", "green", "blue", "orange", "violet"]
        ColorButton {
          color: modelData
          onClicked: {
            amplitude.logEvent("Change Tint Color", {"color" : color, "index" : index+1})
            Theme.colors.tintColor = color
          }
        }
      }
    }

    // text color
    AppText {
      Layout.alignment: Qt.AlignHCenter
      text: "Text Color"
      wrapMode: Text.NoWrap
    }

    Flow {
      Layout.alignment: Qt.AlignHCenter
      Layout.maximumWidth: parent.width
      spacing: parent.spacing

      Repeater {
        model: [
          ["black", "#6C6C6C"],
          ["white", "#BFBFBF"],
          ["red", "#FF7D00"],
          ["green", "#007D4B"],
          ["blue", "#007DFF"]          // 7D00FF
        ]
        ColorButton {
          color: modelData[0]
          referenceColor: Theme.colors.textColor
          onClicked: {
            amplitude.logEvent("Change Text Color", {"color" : modelData[0], "index" : index})
            Theme.colors.textColor = modelData[0]
            Theme.colors.secondaryTextColor = modelData[1]
            if(index === 0)
              Theme.colors.disclosureColor = "#C5C5CA"
            else
              Theme.colors.disclosureColor = modelData[1]
          }
        }
      }
    }

    // background color
    AppText {
      Layout.alignment: Qt.AlignHCenter
      text: "Background Color"
      wrapMode: Text.NoWrap
    }

    Flow {
      Layout.alignment: Qt.AlignHCenter
      Layout.maximumWidth: parent.width
      spacing: parent.spacing

      Repeater {
        model: [
          ["white", "#EFEFF4", "#D9D9D9"],
          ["lightgrey", "#EFEFF4", "#BFBFBF"],
          ["#666666", "#0D0D16", "#404040"],
          ["black", "#0D0D16", "#262626"],
        ]
        ColorButton {
          color: modelData[0]
          referenceColor: Theme.colors.backgroundColor
          onClicked: {
            amplitude.logEvent("Change Background Color", {"color" : modelData[0], "index" : index})
            Theme.colors.backgroundColor = modelData[0]
            Theme.listItem.backgroundColor = modelData[0]
            if(modelData[1] === "")
              return
            Theme.colors.secondaryBackgroundColor = modelData[1]
            Theme.colors.selectedBackgroundColor = modelData[2]
          }
        }
      }
    }

    // separator
    Item {
      Layout.fillWidth: true
      Layout.preferredHeight: parent.spacing * 2
      Rectangle {
        width: parent.width * 0.4
        height: px(1)
        color: Theme.tintColor
        anchors.centerIn: parent
      }
    }

    // tab bar
    Row {
      Layout.alignment: Qt.AlignHCenter
      spacing: parent.spacing

      AppText {
        anchors.verticalCenter: parent.verticalCenter
        text: "Dark Tabs:"
        wrapMode: Text.NoWrap
      }

      AppSwitch {
        anchors.verticalCenter: parent.verticalCenter
        checked: Theme.tabBar.backgroundColor == "#080808"
        updateChecked: false
        onToggled: {
          amplitude.logEvent(!checked ? "Check Dark Tabs" : "Uncheck Dark Tabs")
          if(!checked) {
            Theme.tabBar.titleColor = Qt.binding(function() { return Theme.isAndroid ? Theme.navigationBar.itemColor : Theme.tintColor })
            Theme.tabBar.titleOffColor = Qt.binding(function() { return Theme.isAndroid ? Qt.darker(Theme.navigationBar.titleColor, 1.5) : "#616161" })
            Theme.tabBar.backgroundColor = "#080808"

            Theme.navigationBar.backgroundColor = "#080808"
            Theme.navigationBar.titleColor = Qt.binding(function() {return Theme.isIos ? "#fff" : Theme.isAndroid ? "#fff" : "#f8f8f8" })
            Theme.navigationBar.itemColor = Qt.binding(function() {return Theme.isIos ? Theme.tintColor : Theme.tintColor })

            Theme.colors.statusBarStyle = Theme.colors.statusBarStyleWhite
          }
          else {
            // default setting
            Theme.tabBar.titleColor = Qt.binding(function() { return Theme.tintColor })
            Theme.tabBar.titleOffColor = Qt.binding(function() { return Theme.secondaryTextColor })
            Theme.tabBar.backgroundColor = Qt.binding(function() { return Theme.isAndroid ? app.secondaryTintColor : "#f8f8f8" })

            Theme.navigationBar.backgroundColor = Qt.binding(function() { return Theme.isAndroid ? app.secondaryTintColor : "#f8f8f8" })
            Theme.navigationBar.titleColor = Qt.binding(function() {return Theme.isIos ? "#000" : Theme.isAndroid ? "#fff" : "#f8f8f8" })
            Theme.navigationBar.itemColor = Qt.binding(function() {return Theme.isIos ? Theme.tintColor : Theme.navigationBar.titleColor })

            Theme.colors.statusBarStyle = Qt.binding(function() { return Theme.isAndroid ? Theme.colors.statusBarStyleWhite : Theme.colors.statusBarStyleBlack })
          }
        }
      } // AppSwitch
    } // Row
  } // GridLayout
} // Page
