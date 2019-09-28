import Felgo 3.0
import QtQuick 2.5
import QtQuick.Controls 2.0 as Quick2
import Qt.labs.settings 1.0 as QtLabs
import QtQuick.Controls.Material 2.0
import "../controls"
import "../common"


Page {
  id: controlsPage

  // configure custom rightBarItem for navigation bar of widgets page
  rightBarItem:  NavigationBarRow {
    id: rightNavBarRow

    TextButtonBarItem {
      text: Theme.tintColor !== Qt.rgba(0.25,0.25,0.25,1) ? "Grey" : "Color"
      textItem.font.pixelSize: sp(16)
      title: text
      showItem: Theme.isAndroid ? showItemNever : showItemAlways

      // switch theme color when clicked
      property color prevColor
      onClicked: {
        if(text.toLowerCase() === "grey") {
          prevColor = Theme.tintColor
          Theme.colors.tintColor = Qt.rgba(0.25,0.25,0.25,1)
        }
        else
          Theme.colors.tintColor = prevColor
      }
    }

    IconButtonBarItem {
      onClicked: NativeDialog.confirm("Controls", "This page shows different controls you can use in your app.",
                                      function(ok){}, false)
      title: "Info"
      icon: IconType.info
    }
  }

  // define quick controls colors
  Material.accent: Theme.tintColor

  AppTabBar {
    id: controlsTabBar
    AppTabButton { text: "iOS" }
    AppTabButton { text: system.isPlatform(1) ? "Custom" : "Android" }
    onCurrentIndexChanged: {
      var target = currentIndex == 0 ? "ios" : (currentIndex == 1 ? "android" : "windows")
      if(target !== Theme.platform)
        Theme.platform = target
    }
    Component.onCompleted: currentIndex = Theme.isIos ? 0 : Theme.isAndroid ? 1 : 2
    Connections {
      target: Theme
      onPlatformChanged: {
        var idx = 2
        if(Theme.platform === "ios")
          idx = 0
        else if(Theme.platform === "android")
          idx = 1
        if(controlsTabBar.currentIndex !== idx)
          controlsTabBar.currentIndex = idx
      }
    }
  }

  // page item container
  Item {
    anchors.top: controlsTabBar.bottom
    anchors.bottom: parent.bottom
    width: parent.width

    AppFlickable {
      id: scroll
      anchors.fill: parent
      contentHeight: content.height

      // remove focus from controls if clicked outside
      MouseArea {
        anchors.fill: parent
        onClicked: scroll.forceActiveFocus()
      }

      // content
      Column  {
        id: content
        width: parent.width
        anchors.horizontalCenter: parent.horizontalCenter

        // remove the header because it looks strange if the controls are above the search bar. also, the info that these are controls is already in the tab name
        //SectionHeader { image: "../../assets/felgo-logo.png"; text: "Felgo Controls" }

        SearchBar {
          id: searchBar
          iosAlternateStyle: true
          onAccepted: {
            NativeDialog.confirm("Search", "This is a search bar with a platform-specific look.", function(){}, false)
          }
        }

        SectionDescription { text: "Felgo offers a pre-defined style for Android and iOS based on the app theme. Toggle the switch below to change the platform theme." }

        // arrange the components based on their usefulness & how well they look native (eg AppCheckBox does not really look native thus move it further down)
        // switch
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "AppSwitch"; horizontalAlignment: Text.AlignHCenter }
            Row {
              anchors.horizontalCenter: parent.horizontalCenter
              spacing: parent.spacing
              AppText {
                anchors.verticalCenter: parent.verticalCenter
                text: system.isPlatform(1) ? "Custom Theme:" : "iOS Theme:"
              }
              AppSwitch {
                anchors.verticalCenter: parent.verticalCenter
                checked: Theme.isAndroid && system.isPlatform(1) || Theme.isIos && !system.isPlatform(1)
                updateChecked: false
                onToggled: {
                  if(system.isPlatform(1))
                    Theme.platform = checked ? "ios" : "android"
                  else
                    Theme.platform = checked ? "android" : "ios"
                }
              }
            }
          }
        }

        // button
        SectionContent {
          contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "AppButton"; horizontalAlignment: Text.AlignHCenter }
            Row { anchors.horizontalCenter: parent.horizontalCenter
              AppButton { id: normButton; verticalMargin: 0; text: "Normal"; flat: false; onClicked: text = text === "Normal" ? "Clicked" : "Normal" }
              AppButton { id: flatButton; verticalMargin: 0; text: "Flat"; flat: true; onClicked: text = text === "Flat" ? "Clicked" : "Flat" }
            }
          }
          color: Theme.secondaryBackgroundColor
        }

        // slider
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "AppSlider"; horizontalAlignment: Text.AlignHCenter }
            AppSlider { implicitWidth: dp(200); anchors.horizontalCenter: parent.horizontalCenter }
          }
        }

        // range slider
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "AppRangeSlider"; horizontalAlignment: Text.AlignHCenter }
            AppRangeSlider { implicitWidth: dp(200); anchors.horizontalCenter: parent.horizontalCenter }
          }
          color: Theme.secondaryBackgroundColor
        }


        // checkbox
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "AppCheckBox"; horizontalAlignment: Text.AlignHCenter }
            AppCheckBox { text: !checked ? "Check Me!" : "Checked"; anchors.horizontalCenter: parent.horizontalCenter }
          }
        }

        // text field
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "AppTextField"; horizontalAlignment: Text.AlignHCenter }
            AppTextField {  implicitWidth: dp(200); placeholderText: "Enter text ..."; anchors.horizontalCenter: parent.horizontalCenter }
          }
          color: Theme.secondaryBackgroundColor
        }

        // text edit
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "AppTextEdit"; horizontalAlignment: Text.AlignHCenter }
            CustomAppTextEdit {  anchors.horizontalCenter: parent.horizontalCenter }
          }
        }

        // busy indicator
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "BusyIndicator"; horizontalAlignment: Text.AlignHCenter }
            Quick2.BusyIndicator {
              anchors.horizontalCenter: parent.horizontalCenter
              running: true; implicitWidth: dp(48); implicitHeight: width
              visible: !Theme.isIos
            }
            AppImage {
              id: iosActivity
              anchors.horizontalCenter: parent.horizontalCenter
              source: "../../assets/iosactivity.png"
              width: dp(28)
              height: dp(28)
              smooth: true
              antialiasing: true

              Timer {
                running: true
                repeat: true
                interval: 100
                onTriggered: {
                  iosActivity.rotation += 30
                  if(iosActivity.rotation === 360)
                    iosActivity.rotation = 0
                }
              }
              visible: Theme.isIos
            }
          }
          color: Theme.secondaryBackgroundColor
        }

        // time picker dialog
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "Time or Date Picker (only iOS style)"; horizontalAlignment: Text.AlignHCenter }
            Row {
              anchors.horizontalCenter: parent.horizontalCenter
              spacing: parent.spacing
              AppText {
                text: "Time: "+leadingZero(timePickerDialog.time.hour)+":"+leadingZero(timePickerDialog.time.minute)
                anchors.verticalCenter: parent.verticalCenter
                function leadingZero(number) {
                  return ('00' + number).slice(-2)
                }
              }
              AppButton { flat: false; text: "Set"; onClicked: timePickerDialog.open(); anchors.verticalCenter: parent.verticalCenter }
            }
          }
        } // time picker

        // FloatingActionButton
        SectionContent { contentItem: Column {
            width: parent.width
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "ActionButton (only Android style)"; horizontalAlignment: Text.AlignHCenter }
            FloatingActionButton {
              anchors.horizontalCenter: parent.horizontalCenter
              icon: IconType.star
              onClicked: icon === IconType.star ? icon = IconType.envelope : icon = IconType.star
              anchors.bottom: undefined
              anchors.right: undefined
              visible: true
            }
          }
          color: Theme.secondaryBackgroundColor
        }


        // quick controls 2
        SectionSpacer { color: Theme.secondaryBackgroundColor }
        SectionHeader { icon: IconType.book; text: "Qt Quick Controls 2" }
        SectionDescription { text: "Choose between three different styles, including Material-design." }

        // combobox
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "ComboBox"; horizontalAlignment: Text.AlignHCenter }
            CustomComboBox { anchors.horizontalCenter: parent.horizontalCenter;  model: ["First", "Second", "Third"] }
          }
        }

        // dial
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "Dial"; horizontalAlignment: Text.AlignHCenter }
            Quick2.Dial {
              anchors.horizontalCenter: parent.horizontalCenter
              implicitWidth: dp(96); implicitHeight: implicitWidth
              padding: 0
              AppText { text: Math.round(parent.position * 100); font.pixelSize: sp(10); anchors.centerIn: parent }
            }
          }
          color: Theme.secondaryBackgroundColor
        }

        // progress bar
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "ProgressBar"; horizontalAlignment: Text.AlignHCenter }
            Quick2.ProgressBar {
              anchors.horizontalCenter: parent.horizontalCenter
              implicitWidth: 200; indeterminate: true
              scale: dp(0.85)
            }
            Quick2.ProgressBar {
              id: progressBar
              anchors.horizontalCenter: parent.horizontalCenter
              implicitWidth: 200
              scale: dp(0.85)
              PropertyAnimation {
                target: progressBar; property: "value"; from: 0; to: 1; duration: 3000;
                running: true; loops: Animation.Infinite
              }
            }
          }
        }

        // radio button
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "RadioButton"; horizontalAlignment: Text.AlignHCenter }
            Column {
              id: radioColumn
              anchors.horizontalCenter: parent.horizontalCenter
              CustomRadioButton { checked: true; text: "One" }
              CustomRadioButton { text: "Two" }
              CustomRadioButton { text: "Three" }
            }
          }
          color: Theme.secondaryBackgroundColor
        }

        // spin box
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "SpinBox"; horizontalAlignment: Text.AlignHCenter }
            // wrapper item for correct layout of scaled spinbox
            Item {
              anchors.horizontalCenter: parent.horizontalCenter
              implicitWidth: spinBox.implicitWidth * spinBox.scale
              implicitHeight: spinBox.implicitHeight * spinBox.scale
              Quick2.SpinBox {
                id: spinBox
                scale: dp(0.85) // density-dependent scaling
                from: 0; to: 100; stepSize: 1
                editable: true
                anchors.centerIn: parent
              }
            }
          }
        }

        // tool tip
        SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "ToolTip"; horizontalAlignment: Text.AlignHCenter }
            CustomToolTipButton {
              anchors.horizontalCenter: parent.horizontalCenter
              text: "Press and Hold"
              toolTip: "A simple ToolTip"
            }
          }
          color: Theme.secondaryBackgroundColor
        }

        // tumbler
        /* SectionContent { contentItem: Column {
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: dp(10)
            AppText { width: parent.width; font.pixelSize: sp(12); text: "Tumbler"; horizontalAlignment: Text.AlignHCenter }
            Row {
              anchors.horizontalCenter: parent.horizontalCenter
              id: dateRow
              spacing: dp(2)
              property var curDate: new Date()

              CustomTumbler{
                id: dateTumbler
                model: 31
                processDelegateText: parent.twoDigits
              }
              CustomTumbler{
                id: monthTumbler
                model: 12
                processDelegateText: parent.twoDigits
              }
              CustomTumbler{
                id: yearTumbler
                model: 100
                // display text for current year and [model] years back
                property int offset: dateRow.curDate.getFullYear() - (model * 0.75)
                processDelegateText: function(text) {
                  return offset + parseInt(text)
                }
              }

              // initialize tumblers
              Component.onCompleted: {
                dateTumbler.currentIndex = dateRow.curDate.getDate() - 1
                monthTumbler.currentIndex = dateRow.curDate.getMonth()
                yearTumbler.currentIndex = dateRow.curDate.getFullYear() - yearTumbler.offset
              }

              // adds + 1 to value a leading zero if required
              function twoDigits(value) {
                return ("0" + (parseInt(value) + 1)).slice(-2)
              }
            }
          }
        } // tumbler */
      } // content grid
    } // flickable

    // time picker dialog
    TimePickerDialog { id: timePickerDialog }

    // declare radio button group
    Quick2.ButtonGroup {
      buttons: radioColumn.children
    }

    // add scroll indicator
    ScrollIndicator {
      flickable: scroll
    }
  } // item
} // page
