import Felgo 3.0
import QtQuick 2.5
import QtLocation 5.5
import QtPositioning 5.5
import QtQuick.Controls 2.0 as Quick2
import QtMultimedia 5.5
import QtGraphicalEffects 1.0
import "../controls"
import "../common"

FlickablePage {
  id: componentsPage

  property var video: null // embbedded video if currently playing
  property var alphaVideo: null // alphaVideo if currently playing

  flickable.contentHeight: content.height

  property real yBeginning: 0
  property real yDelta: flickable.contentY - yBeginning
  property real videoContentY: 0
  Component.onCompleted: yBeginning = flickable.contentY
  onHeightChanged: videoContentY = video !== null ? video.mapToItem(flickable.contentItem, 0 , 0).y : 0

  // remove focus from controls if clicked outside
  MouseArea {
    anchors.fill: parent
    onClicked: flickable.forceActiveFocus()
  }

  // content
  Column  {
    id: content
    width: parent.width
    anchors.horizontalCenter: parent.horizontalCenter

    SectionHeader { icon: IconType.paintbrush; text: "User Interface Styling" }
    SectionDescription { text: "You can customize the styling of all UI components and controls, or use the default native look & feel of your platform.\n
Change the tint color or switch the platform theme below. All UI elements will then update their look." }
    SectionContent { contentItem: Column {
        width: parent.width
        spacing: dp(Theme.navigationBar.defaultBarItemPadding)

        Row {
          id: tintColorRow
          anchors.horizontalCenter: parent.horizontalCenter
          spacing: parent.spacing
          property color defaultColor: Theme.isIos ? "#007aff" : (Theme.isAndroid ? "#3f51b5" : "#01a9e2") // default is felgo blue
          property int currentIndex: 0

          Connections {
            target: Theme
            onPlatformChanged: if(tintColorRow.currentIndex === 0) Theme.colors.tintColor = tintColorRow.defaultColor
          }

          AppText {
            text: "Tint:"
            anchors.verticalCenter: parent.verticalCenter
          }

          Repeater {
            model: [tintColorRow.defaultColor, "#FF3B30", "#4CD964", "#FF9500"]

            Rectangle {
              color: modelData

              width: dp(48)
              height: dp(48)
              MouseArea {
                anchors.fill: parent
                onClicked: { Theme.colors.tintColor =  color; tintColorRow.currentIndex = index }
              }
            }
          }
        } // tint row

        Row {
          anchors.horizontalCenter: parent.horizontalCenter
          AppText {
            text: "Theme: "
            anchors.verticalCenter: parent.verticalCenter
          }

          IconButton {
            anchors.verticalCenter: parent.verticalCenter
            icon: IconType.apple
            size: dp(24)
            color: Theme.isIos ? Theme.tintColor : Theme.secondaryTextColor
            onClicked: Theme.platform = "ios"
          }

          IconButton {
            id: androidIcon
            anchors.verticalCenter: parent.verticalCenter
            icon: IconType.android
            color: Theme.isAndroid ? Theme.tintColor : Theme.secondaryTextColor
            size: dp(24)
            visible: !system.isPlatform(1) // 1 == iOS
            onClicked: Theme.platform = "android"
          }

          AppButton {
            text: "Custom"
            anchors.top: parent.top
            anchors.topMargin: Theme.isAndroid ? dp(2) : dp(1)
            flat: true
            verticalPadding: 0
            borderColor: Theme.isIos ? Theme.secondaryTextColor : Theme.tintColor
            textColor: Theme.isIos ? borderColor : Theme.appButton.flatTextColor
            visible: !androidIcon.visible
            onClicked: Theme.platform = "android"
          }
        }

      } // column
    } // theming section

    // Native Features
    SectionSpacer { height: dp(5) }
    SectionHeader { icon: IconType.mobile; text: "Native Features" }
    SectionDescription { text: "Use native features like opening dialogs or access device information with a single code base & cross-platform." }
    SectionContent { contentItem: Column {
        width: parent.width
        spacing: dp(10)
        Row {
          anchors.horizontalCenter: parent.horizontalCenter
          AppButton { text: "Alert Dialog"; onClicked: nativeUtils.displayAlertDialog("Alert", "Be careful!", "I will", "Don't worry!"); verticalMargin: 0; flat: false }
          AppButton { text: "Input Dialog"; onClicked:NativeDialog.inputText("Input", "This is a native input dialog."); verticalMargin: 0; flat: false }
        }
        Row {
          anchors.horizontalCenter: parent.horizontalCenter
          AppButton { text: "Alert Sheet"; onClicked: nativeUtils.displayAlertSheet("Choose", ["Pizza", "Pasta", "Steak"], true); verticalMargin: 0; flat: false }
          AppButton { text: "Share Dialog"; onClicked: nativeUtils.share("Sharing is caring!", "https://www.felgo.com"); verticalMargin: 0; flat: false }
        }
      }
    }

    // open app / share
    SectionSpacer {}
    SectionHeader { text: "Open Apps, Send E-Mail and More" }
    SectionDescription { text: "Start other apps, open an url in the browser or let the user send emails. Take a photo from camera or pick an image from gallery." }

    SectionContent { contentItem: Column {
        width: parent.width
        spacing: dp(16)

        Row {
          anchors.horizontalCenter: parent.horizontalCenter
          CustomComboBox {
            id: comboBox2
            anchors.verticalCenter: parent.verticalCenter
            implicitWidth: dp(120) + 30
            model: ["Open App", "Open Url", "Send Mail"]
          }
          AppButton {
            text: "Open"
            anchors.verticalCenter: parent.verticalCenter
            onClicked: {
              switch(comboBox2.currentIndex) {
              case 0:
                NativeDialog.confirm(
                      "Open Facebook",
                      "The Facebook App will be opened. Proceed?",
                      function(proceed) {
                        if(!proceed) return
                        var launchParam = Theme.isIos ? "fb://profile" : "com.facebook.katana"
                        if(!nativeUtils.openApp(launchParam)) {
                          NativeDialog.confirm("Facebook could not be launched.","", function(){}, false)
                        }
                      })
                break
              case 1:
                NativeDialog.confirm(
                      "Open Url",
                      "The Felgo website will be opened. Proceed?",
                      function(proceed) {
                        if(!proceed) return
                        nativeUtils.openUrl("https://www.felgo.com")
                      })
                break
              case 2:
                NativeDialog.confirm(
                      "Send Mail",
                      "The mail app will be opened with a new draft. Proceed?",
                      function(proceed) {
                        if(!proceed) return
                        nativeUtils.sendEmail("", "I just wanted to ...", "wish you a cool day! ;)")
                      })
                break
              }
            }
          }
        } // open, share, ...

        // other native features
        Row {
          anchors.horizontalCenter: parent.horizontalCenter
          CustomComboBox {
            id: comboBox3
            anchors.verticalCenter: parent.verticalCenter
            implicitWidth: dp(120) + 30
            model: [ "Store Contact", "Vibrate" ]
          }
          AppButton {
            text: "Do It"
            anchors.verticalCenter: parent.verticalCenter
            onClicked: {
              switch(comboBox3.currentIndex) {
              case 0:
                var vCard = "BEGIN:VCARD
VERSION:2.1
N:Felgo;Office;;;
FN:Felgo
TEL;WORK:0123456789
EMAIL;WORK:help@felgo.com
ADR;WORK:;;Kolonitzgasse;Wien;;1030;Austria
ORG:Felgo
URL:https://www.felgo.com
END:VCARD";
                var success = nativeUtils.storeContacts(vCard)
                if(Theme.isIos) {
                  // handle success/error on iOS to give feedback to user
                  if(success)
                    NativeDialog.confirm("Contact stored.", "", function() {}, false)
                  else
                    NativeDialog.confirm("Could not store contact",
                                         "The app does not have permission to access the AddressBook, please allow access in the device settings.",
                                         function() {}, false)
                }
                else if (Theme.isAndroid) {
                  // only react to error on Android, if successful the device will open the vcard data file
                  if(!success) {
                    NativeDialog.confirm("Could not store contact",
                                         "Error when trying to store the vCard to the user documents folder.",
                                         function() {}, false)
                  }
                }
                else {
                  // platform not supported
                  NativeDialog.confirm("Could not store contact", "Storing contacts is only possible on iOS and Android.", function() {}, false)
                }
                break
                // case 0: store contact end
              case 1:
                nativeUtils.vibrate(); break
              }
            }
          }
        } // general native features

        Row {
          anchors.horizontalCenter: parent.horizontalCenter
          spacing: dp(10)
          AppText {
            anchors.verticalCenter: parent.verticalCenter
            text: "Statusbar visible:"
          }

          AppSwitch {
            anchors.verticalCenter: parent.verticalCenter
            checked: Theme.statusBarStyle === Theme.colors.statusBarStyleBlack
            updateChecked: false
            onToggled: Theme.colors.statusBarStyle = checked ? Theme.colors.statusBarStyleHidden : Theme.colors.statusBarStyleBlack
          }
        }

        // image/camera picker
        Row {
          anchors.horizontalCenter: parent.horizontalCenter
          height: userImage.height
          spacing: dp(10)
          width: userImage.width + spacing + choosePhotoText.width
          UserImage {
            id: userImage
            property string iconFontName: Theme.iconFont.name
            width: dp(72)
            height: width

            placeholderImage: "\uf007" // user
            source: ""

            editable: true
            editBackgroundColor: Theme.tintColor

            property bool shownEditPhotoDialog: false

            onEditClicked: {
              // We do not have camera feature on desktop yet, so just show file dialog
              if (system.desktopPlatform) {
                nativeUtils.displayImagePicker(qsTr("Choose Image"))
              }
              else {
                // Probably better use a QML styled dialog?
                shownEditPhotoDialog = true
                nativeUtils.displayAlertSheet("", ["Choose Photo", "Take Photo", "Reset Photo"], true)
              }
            }

            Connections {
              target: nativeUtils
              onAlertSheetFinished: {
                if (userImage.shownEditPhotoDialog) {
                  if (index == 0)
                    nativeUtils.displayImagePicker(qsTr("Choose Image")) // Choose image
                  else if (index == 1)
                    nativeUtils.displayCameraPicker("Take Photo") // Take from Camera
                  else if (index == 2)
                    userImage.source = "" // Reset
                  userImage.shownEditPhotoDialog = false
                }
              }

              onImagePickerFinished: {
                console.debug("Image picker finished with path:", path)
                if(accepted)
                  userImage.source = Qt.resolvedUrl(path)
              }

              onCameraPickerFinished: {
                console.debug("Camera picker finished with path:", path)
                if(accepted)
                  userImage.source = Qt.resolvedUrl(path)
              }
            }
          } // User Image
          AppText {
            id: choosePhotoText
            width: dp(130)
            text: "click to choose image or take photo"
            anchors.verticalCenter: parent.verticalCenter
            font.pixelSize: dp(12)
            wrapMode: Text.WordWrap
          }
        } // image camer picker
      } // Column
    } // Native SectionContent

    SectionSpacer { }
    SectionHeader { text: "Device Info" }
    SectionContent { contentItem: Row {
        width: parent.width
        Column {
          id: info1
          spacing: dp(5)
          width: (parent.width - parent.spacing) * 0.5
          AppText { text: "<b>OS Type:</b> "+system.osType; font.pixelSize: sp(11) }
          AppText { text: "<b>OS Version:</b> "+system.osVersion; font.pixelSize: sp(11) }
          AppText { text: "<b>Device Model:</b> "+nativeUtils.deviceModel(); font.pixelSize: sp(11) }
          AppText { text: "<b>Device Online:</b> "+system.isOnline; font.pixelSize: sp(11) }
          AppText { text: "<b>MAC Address:</b> "+system.macAddress; font.pixelSize: sp(11)
            width: parent.width; wrapMode: Text.WrapAtWordBoundaryOrAnywhere }
          AppText { text: "<b>Unique Device Id:</b> "+system.UDID; font.pixelSize: sp(11);
            width: parent.width; wrapMode: Text.WrapAtWordBoundaryOrAnywhere }
        }
        Column {
          id: info2
          spacing: dp(5)
          width: info1.width
          AppText { text: "<b>Orientation:</b> "+(Theme.isPortrait ? "portrait" : "landscape"); font.pixelSize: sp(11) }
          AppText { text: "<b>Statusbar Height:</b> "+Theme.statusBarHeight; font.pixelSize: sp(11) }
          AppText { text: "<b>App Identifier:</b> "+system.appIdentifier; font.pixelSize: sp(11)
            width: parent.width; wrapMode: Text.WrapAtWordBoundaryOrAnywhere }
          AppText { text: "<b>App Version:</b> "+system.appVersionCode; font.pixelSize: sp(11) }
          AppText { text: "<b>Qt Version:</b> "+system.qtVersion; font.pixelSize: sp(11) }
          AppText { text: "<b>Felgo Version:</b> "+system.vplayVersion; font.pixelSize: sp(11) }
        }
      }
    }

    // app map
    SectionSpacer {}
    SectionHeader { icon: IconType.mapmarker; text: "Maps and Positioning" }
    SectionDescription { text: "Embed maps with custom markers, routing and retrieve the user position." }
    Item {
      width: parent.width
      height: map.height + positionRow.height

      AppMap {
        id: map
        width: parent.width
        height: dp(200)
        showUserPosition: true

        plugin: Plugin {
          name: "mapbox"
          // configure your own map_id and access_token here
          parameters: [  PluginParameter {
              name: "mapbox.mapping.map_id"
              value: "mapbox.streets"
            },
            PluginParameter {
              name: "mapbox.access_token"
              value: "pk.eyJ1IjoiZ3R2cGxheSIsImEiOiJjaWZ0Y2pkM2cwMXZqdWVsenJhcGZ3ZDl5In0.6xMVtyc0CkYNYup76iMVNQ"
            },
            PluginParameter {
              name: "mapbox.mapping.highdpi_tiles"
              value: true
            }]
        }

        // set initial map coordinate and zoom level
        center: userPositionAvailable ? userPosition.coordinate : QtPositioning.coordinate(48.208417, 16.372472)
        zoomLevel: 13
        Component.onCompleted: {
          map.center = userPositionAvailable ? userPosition.coordinate : QtPositioning.coordinate(48.208417, 16.372472)
          map.zoomLevel = 13
        }

        // once we successfully received the location, we zoom to the user position
        onUserPositionAvailableChanged: {
          if(userPositionAvailable && userPosition.horizontalAccuracyValid && map.center !== userPosition.coordinate) {
            zoomToUserPosition()
          }
        }

        // add marker item to map
        MapQuickItem {
          // overlay will be placed below user position (if available)
          coordinate: parent.userPositionAvailable ? parent.userPosition.coordinate : parent.center

          // the anchor point specifies the point of the sourceItem that will be placed at the given coordinate
          anchorPoint: Qt.point(sourceItem.width/2, -dp(35))

          // place above user position and accuracy items
          z: 3

          // source item holds the actual item that is displayed on the map
          sourceItem: Rectangle {
            width: dp(150)
            height: dp(50)
            color: "white"
            radius: width * 0.05
            border.color: Theme.tintColor
            border.width: dp(1)
            layer.effect: DropShadow { horizontalOffset: 2; verticalOffset: 2
              radius: 16.0; samples: 16; color: "#80000000" }
            layer.enabled: true

            AppText {
              text: "This is a custom marker!"
              anchors.centerIn: parent
              font.pixelSize: sp(10)
            }
          }
        } // MapQuickItem
      }

      Row {
        id: positionRow
        anchors.bottom: parent.bottom
        anchors.horizontalCenter: parent.horizontalCenter
        spacing: dp(10)
        width: posButton.width + spacing + posText.contentWidth
        AppButton { id: posButton; text: "Zoom to Position"; onClicked: map.zoomToUserPosition() }
        AppText {
          id: posText
          text: "<b>Lat:</b> "+(map.userPositionAvailable ? map.userPosition.coordinate.latitude.toFixed(8) : "Not available")
                + "<br/><b>Lng:</b> "+(map.userPositionAvailable ? map.userPosition.coordinate.longitude.toFixed(8) : "Not available")
          font.pixelSize: sp(11)
          anchors.verticalCenter: parent.verticalCenter
        }
      }
    }

    // multimedia
    SectionSpacer {}
    SectionHeader { icon: IconType.videocamera; text: "Multimedia Components" }
    SectionDescription { text: "Show videos or play sound anywhere in your app. You can even play a video with alpha channel for advanced visual effects." }
    Rectangle {
      id: embeddedVideoContainer
      width: parent.width
      height: width * 0.75
      color: componentsPage.alphaVideo !== null ? "white" : "black"
      Component.onCompleted: playEmbeddedVideoTimer.start()
    }
    SectionContent { contentItem: Column {
        width: parent.width
        spacing: dp(10)
        Row {
          anchors.horizontalCenter: parent.horizontalCenter
          CustomComboBox {
            id: comboBox4
            implicitWidth: dp(150) + 30
            model: [ "Transparent Video", "Music", "Sound" ]
          }
          AppButton {
            id: playButton
            text: {
              if((comboBox4.currentIndex === 1 && music.playbackState === Audio.PlayingState)
                  || (comboBox4.currentIndex === 2 && sound.playing))
                return "Stop"
              else
                return "Play"
            }
            onClicked: {
              switch(comboBox4.currentIndex) {
              case 0:
                // pause and hide embedded video when playing overlay video
                // playing two videos at once didn't work for Samsung Android phone
                if(componentsPage.video !== null) {
                  componentsPage.video.pause()
                  componentsPage.video.visible = false
                }
                // do nothing if overlay videois already playing
                if(componentsPage.alphaVideo === null)
                  componentsPage.alphaVideo = alphaVideoComponent.createObject(componentsPage)
                break
              case 1:
                if(playButton.text !== "Stop")
                  music.play()
                else
                  music.stop()
                break
              case 2:
                if(playButton.text !== "Stop")
                  sound.play()
                else
                  sound.stop()
                break
              }
            }
          }
        }
      }
    } // multimedia section

    // Page Control
    SectionSpacer { }
    SectionHeader { text: "Convenience Controls" }
    SectionDescription { text: "Use a PageControl and a SwipeView for horizontal scrolling of content." }
    SectionSpacer { }
    Item {
      width: parent.width
      height: imgSwipeView.height + pageControl.height

      Quick2.SwipeView {
        id: imgSwipeView
        width: parent.width
        height: width / 1.5
        anchors.bottom: parent.bottom
        Item {  AppImage { source: "../../assets/devices.png"; anchors.fill: parent; fillMode: Image.PreserveAspectFit } }
        Item {  AppImage { source: "../../assets/onu.png"; anchors.fill: parent; fillMode: Image.PreserveAspectFit } }
        Item {  AppImage { source: "../../assets/squaby.png"; anchors.fill: parent; fillMode: Image.PreserveAspectFit } }
        Item {  AppImage { source: "../../assets/stack.png"; anchors.fill: parent; fillMode: Image.PreserveAspectFit } }
      }

      PageControl {
        id: pageControl
        height: implicitHeight + dp(5)
        pages: 4
        currentPage: imgSwipeView.currentIndex
        clickableIndicator: true
        spacing: dp(10)
        onPageSelected: imgSwipeView.currentIndex = index
      }
    } // page control / convenvience controls

    // Picture Viewer
    SectionDescription { text: "The PictureViewer shows an image in fullscreen." }
    SectionContent { contentItem: Column {
        spacing: dp(10)
        anchors.horizontalCenter: parent.horizontalCenter
        AppText { text: "click to open PictureViewer: "; font.pixelSize: dp(12); anchors.horizontalCenter: parent.horizontalCenter }
        AppImage {
          anchors.horizontalCenter: parent.horizontalCenter
          width: parent.parent.width * 0.35
          fillMode: AppImage.PreserveAspectFit
          source: "../../assets/felgo-logo.png"
          MouseArea {
            anchors.fill: parent
            onClicked: PictureViewer.show(getApplication(), parent.source)
          }
        }
      }
    } // convenience controls section

    // web view (we use a loader here to not use webview on desktop at all to reduce size of Felgo Sample Launcher desktop application)
    Loader {
      id: webViewLoader
      width: parent.width
      height: webViewLoader.item ? webViewLoader.item.implicitHeight : 0
      source: system.desktopPlatform ? "" : Qt.resolvedUrl("../common/WebViewSection.qml")
    }
  } // content column

  // media items
  Audio {
    id: music
    source: "../../assets/multimedia/coconut-land.wav"
    loops: 0
    autoLoad: true
  }

  SoundEffect {
    id: sound
    source: "../../assets/multimedia/collected.wav"
    loops: 0
  }

  // component for embedded video
  Component {
    id: videoComponent
    Video {
      anchors.centerIn: parent
      anchors.fill: parent
      source: isWindows ? "../../assets/multimedia/FelgoAnimation.wmv" : "../../assets/multimedia/FelgoAnimation.mp4"
      autoPlay: false

      // only show and play video if not dragging an video is visible in view
      property bool allowPlayback: !flickable.dragging && !flickable.flicking && inVisibleArea
      property bool inVisibleArea: (yDelta + flickable.height >= videoContentY) && (yDelta <= videoContentY + height)
      onAllowPlaybackChanged: if(allowPlayback) playTimer.start(); else playTimer.stop()
      visible: playbackState === MediaPlayer.PlayingState

      Timer {
        id: playTimer
        interval: 250
        onTriggered: parent.play()
      }

      Component.onCompleted: {
        videoContentY = embeddedVideoContainer.mapToItem(flickable.contentItem, 0 , 0).y
        if(!allowPlayback) stop(); else playTimer.start();
        stopped.connect(function() { if(allowPlayback) playTimer.start() }) // replay when finished
      }
    }
  } // Video

  // Timer to create/start embedded video
  // creating new video after previously destroying another one doesn't work immediately
  Timer {
    id: playEmbeddedVideoTimer
    interval: 2000
    onTriggered: componentsPage.video = videoComponent.createObject(embeddedVideoContainer)
  }

  // component for alpha video
  Component {
    id: alphaVideoComponent
    Rectangle {
      id: alphaVideoItem
      anchors.fill: parent
      color: Qt.rgba(0,0,0,0.5)

      AlphaVideo {
        anchors.centerIn: parent
        anchors.fill: parent
        source: isWindows ? "../../assets/multimedia/AlphaAnimation.wmv" : "../../assets/multimedia/AlphaAnimation.mp4"
        autoPlay: true

        // destroy video after playback
        property bool destroyWhenVisible: false
        onStatusChanged: {
          if(status === MediaPlayer.EndOfMedia) {
            if(!visible)
              destroyWhenVisible = true
            else
              destroyAlphaVideo()
          }
        }
        onVisibleChanged: if(visible && destroyWhenVisible) destroyAlphaVideo()

        function destroyAlphaVideo() {
          alphaVideoItem.destroy()
          componentsPage.alphaVideo = null
          if(componentsPage.video !== null) {
            componentsPage.video.destroy()
            playEmbeddedVideoTimer.start()
          }
        }
      } // AlphaVideo
    }
  } // AlphaVideo Component
} // page
