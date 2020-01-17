import QtQuick 2.0
import QtQuick.Controls 1.4
import QtGraphicalEffects 1.0
import Felgo 3.0
import QtMultimedia 5.5
import "../common"
import "../controls"
import "../qmlvideofx/qml"

Page {
  id: page
  // get the total height of status bar and navigation bar
  readonly property real statusHeight: nativeUtils.safeAreaInsets.top > Theme.statusBarHeight ? nativeUtils.safeAreaInsets.top : Theme.statusBarHeight
  readonly property real barHeight: dp(Theme.navigationBar.height) + statusHeight

  // navigation bar is 100 percent translucent, the page then also fills the whole screen
  // this allows us to display a custom navigation bar background for this page
  navigationBarTranslucency: 1.0
  useSafeArea: false // do not automatically align page content with safe area

  // custom navigation bar background that shows an image
  Rectangle {
    id: background
    width: parent.width
    height: barHeight
    color: Theme.navigationBar.backgroundColor

    // add the image
    Image {
      id: bgImage
      source: "../../assets/felgo-logo.png"
      anchors.fill: parent
      fillMode: Image.PreserveAspectCrop
      opacity: 0.15 + 0.15 * (blurRadius / 64)

      // the strength of the blur is based on the flickable scroll position
      property real blurRadius: Math.max(0, Math.min(64, pageFlickable.contentY * 0.25))

      // apply blur effect
      layer.enabled: true
      layer.effect: FastBlur {
        radius: bgImage.blurRadius
      }
    }
  } // Rectangle

  Item {
    id: flickArea
    anchors.top: background.bottom
    anchors.bottom: parent.bottom
    anchors.left: parent.safeArea.left
    anchors.right: parent.safeArea.right

    // flickable
    AppFlickable {
      id: pageFlickable
      anchors.fill: parent
      contentHeight: content.height + dp(Theme.navigationBar.defaultBarItemPadding) * 2

      // show/hide embedded video during scrolling to opimize scroll performance
      property real yBeginning: 0
      property real yDelta: pageFlickable.contentY - pageFlickable.yBeginning
      property real videoContentY: 0
      Component.onCompleted: pageFlickable.yBeginning = pageFlickable.contentY
      onHeightChanged: videoContentY = appsVideo.parent.mapToItem(pageFlickable.contentItem, 0 , 0).y
      onContentHeightChanged: videoContentY = appsVideo.parent.mapToItem(pageFlickable.contentItem, 0 , 0).y

      // add abckground to flickable to also have background within shader effect
      Rectangle {
        anchors.fill: parent
        color: page.backgroundColor
      }

      // remove focus from controls if clicked outside
      MouseArea {
        anchors.fill: parent
        onClicked: pageFlickable.forceActiveFocus()
      }

      // content
      Column  {
        id: content
        width: parent.width

        // Custom Shaders
        SectionHeader { icon: IconType.photo; text: "Custom Shader Effects" }
        SectionDescription { text: "Use Felgo for eye-catching UI animations, graphical effects, and custom shader effects." }
        SectionContent { contentItem: AppButton {
            text: fxControls.visible ? "Hide Shader Example" : "Open Shader Example"
            anchors.horizontalCenter: parent.horizontalCenter
            flat: false
            onClicked: {
              if(!fxControls.visible) {
                appsVideo.play()
                fxControls.show()
              }
              else {
                if(appsVideo.started) {
                  appsVideo.pause()
                  appsVideo.pausedByUser = true
                }
                fxControls.hide()
              }
            }
          }
        }

        // embedded apps video
        Rectangle {
          color: "black"
          width: parent.width
          height: width / 1280 * 720

          SectionContent {
            anchors.centerIn: parent
            color: "transparent"
            contentItem: Column {
              width: parent.width
              anchors.centerIn: parent
              spacing: dp(10)

              AppImage {
                source: "../../assets/playicon.png"
                width: dp(75)
                height: width / sourceSize.width * sourceSize.height
                anchors.horizontalCenter: parent.horizontalCenter
              }
              AppText {
                width: parent.width * 0.75
                text: "Tap to play video for live video shader effects"
                horizontalAlignment: AppText.AlignHCenter
                anchors.horizontalCenter: parent.horizontalCenter
                wrapMode: AppText.WordWrap
                font.pixelSize: sp(16)
                color: "white"
              }
            }
          }

          Video {
            id: appsVideo
            anchors.fill: parent
            source: isWindows ? "https://www.felgo.com/resources/apps/Felgo_mobile.wmv" : "../../assets/multimedia/Felgo_mobile.mp4"
            autoLoad: true
            property bool started: false
            property bool pausedByUser: false
            onPlaying: { pausedByUser = false; started = true }
            onStopped: { pausedByUser = false; started = false }

            // hide and stop pause if not within visible view to improve performance
            visible: (pageFlickable.yDelta + pageFlickable.height >= pageFlickable.videoContentY) && (pageFlickable.yDelta <= pageFlickable.videoContentY + height)
            onVisibleChanged:if(!visible && started) pause(); else if(visible && started && !pausedByUser) play();

            // play / pause video on click
            MouseArea {
              anchors.fill: parent
              onClicked: {
                if(!appsVideo.started) {
                  appsVideo.play()
                  fxControls.show()
                }
                else {
                  if(appsVideo.pausedByUser) appsVideo.play(); else { appsVideo.pause(); appsVideo.pausedByUser = true }
                }
              }
            }
          }
        }

        // Animations
        SectionSpacer { }
        SectionHeader { icon: IconType.exchange; text: "Powerful Animations" }
        SectionDescription { text: "Easily add animations for the properties of an element. Smooth transitions to the target value are automatically created." }
        SectionContent { contentItem: Column {
            width: parent.width
            spacing: dp(10)

            // color animation
            Rectangle {
              id: colorRect
              width: parent.width
              anchors.horizontalCenter: parent.horizontalCenter
              radius: height * 0.5
              height: dp(50)
              color: "green"

              AppText { text: "Color Change"; anchors.centerIn: parent; font.pixelSize: sp(11); color: "white" }
              PropertyAnimation {
                target: colorRect
                property: "color"
                to: "orange"
                running: true
                duration: 1000
                onStopped: {
                  to = (colorRect.color.r !== 0) ? "green" : "orange"
                  start()
                }
              }
            } // color animation

            // movement
            Rectangle {
              id: moveCircle
              width: dp(75)
              height: dp(50)
              radius: width * 0.5
              color: "#01a9e2"
              x: leftX
              property real leftX: 0
              property real rightX: parent.width - moveCircle.width

              AppText { text: "Movement"; anchors.centerIn: parent; font.pixelSize: sp(11); color: "white" }
              PropertyAnimation {
                target: moveCircle
                property: "x"
                to: moveCircle.rightX
                running: true
                duration: 1500
                onStopped: {
                  to = (moveCircle.x > moveCircle.leftX) ? moveCircle.leftX : moveCircle.rightX
                  start()
                }
                easing.type: Easing.OutBounce
              }
            }

            // rotation
            Row {
              anchors.horizontalCenter: parent.horizontalCenter
              height: dp(50)
              spacing: dp(Theme.navigationBar.defaultBarItemPadding)
              AppImage {
                id: rotateImage
                width: height
                height: parent.height
                source: "../../assets/felgo-logo.png"
                RotationAnimation {
                  target: rotateImage
                  running: true
                  from: 0
                  to: 360
                  duration: 2000
                  loops: Animation.Infinite
                }
              }

              AppText {
                anchors.verticalCenter: parent.verticalCenter
                text: "Rotation"
                font.pixelSize: sp(11)
              }
            }


          } // animation column
        } // Animations SectionContent

        // combined animations and easing
        SectionSpacer { }
        SectionHeader { text: "Combined Animations and Easing" }
        SectionDescription { text: "Combine multiple types to create cool animations." }
        SectionContent {
          contentItem: Item {
            width: animationImage.width + animationText.width
            height: dp(50)
            anchors.horizontalCenter: parent.horizontalCenter

            AppImage {
              id: animationImage
              width: dp(50)
              height: dp(50)
              source: "../../assets/felgo-logo.png"
              opacity: 0
            }
            AppText {
              id: animationText
              anchors.leftMargin: dp(10)
              anchors.left: animationImage.right
              width: dp(180)
              text: "Create beautiful animations in no time."
              y: animationImage.y + animationImage.height
              opacity: 0
            }

            SequentialAnimation {
              running: true
              loops: Animation.Infinite

              // logo fade in
              ParallelAnimation {
                id: logoFadeIn
                readonly property int duration: 1500
                NumberAnimation {
                  target: animationImage
                  property: "opacity"
                  to: 1
                  duration: logoFadeIn.duration * 0.75
                }
                NumberAnimation {
                  target: animationImage
                  property: "scale"
                  from: 1.5
                  to: 1
                  duration: logoFadeIn.duration
                  easing.type: Easing.OutBounce
                }
                RotationAnimation {
                  target: animationImage
                  from: 180
                  to: 360
                  duration: logoFadeIn.duration * 0.5
                }
              }
              // text fade in
              ParallelAnimation {
                id: textFadeIn
                readonly property int duration: 1000
                NumberAnimation {
                  target: animationText
                  property: "opacity"
                  to: 1
                  duration: textFadeIn.duration * 0.5
                }
                NumberAnimation {
                  target: animationText
                  property: "y"
                  from: animationImage.y + animationImage.height
                  to: (animationImage.height - animationText.height) * 0.5
                  duration: textFadeIn.duration * 0.5
                  easing.type: Easing.OutExpo
                }
                ColorAnimation {
                  target: animationText
                  property: "color"
                  from: "black"
                  to: Theme.tintColor
                  duration: textFadeIn.duration
                }
              }
              PauseAnimation { duration: 1000 }
              // all fade out
              ParallelAnimation {
                id: allFadeOut
                readonly property int duration: 250
                NumberAnimation {
                  target: animationImage
                  property: "opacity"
                  to: 0
                  duration: allFadeOut.duration
                }
                NumberAnimation {
                  target: animationText
                  property: "opacity"
                  to: 0
                  duration: allFadeOut.duration
                }
                NumberAnimation {
                  target: animationImage
                  property: "scale"
                  to: 1.5
                  easing.type: Easing.InExpo
                  duration: allFadeOut.duration
                }
                NumberAnimation {
                  target: animationText
                  property: "y"
                  to: animationImage.y - animationText.height
                  easing.type: Easing.InExpo
                  duration: allFadeOut.duration
                }
              }
              PauseAnimation { duration: 1500 }
            }
          }
        } // Combined Animation SectionContent

        SectionDescription { text: "Different easing settings allow to customize the interpolation of animated properties. Click \"show\" to open a dialog that uses the selected setting." }
        SectionContent { contentItem: Row {
            anchors.centerIn: parent
            spacing: dp(10)

            CustomComboBox {
              id: easingComboBox
              implicitWidth: dp(100) + 30
              model: ["OutBounce", "OutBack", "Linear", "OutCubic"]
              anchors.verticalCenter: parent.verticalCenter
            }
            AppButton { text: "Show"; onClicked: easingDialog.open() }
          }
        } // Easing SectionContent

        // transitions
        SectionSpacer { }
        SectionHeader { text: "Transition Effects" }
        SectionDescription { text: "Use animations to automatically add a transition effect when your view changes." }
        SectionContent {
          verticalMargin: 0
          horizontalMargin: 0
          contentItem: Column {
            width: parent.width
            spacing: dp(10)

            ListView {
              id: appListView
              width: parent.width
              height: dummyModel.count * dp(Theme.listItem.minimumHeight) + dp(1)
              model: ListModel { id: dummyModel }
              delegate: SimpleRow {
                id: row
                text: title
                style.showDisclosure: false
                property bool isSubItem: title.substring(0, 4) === "Sub-"

                // button to remove from list
                IconButton {
                  id: closeButton
                  icon: IconType.close
                  anchors.right: parent.right
                  anchors.verticalCenter: parent.verticalCenter
                  onClicked: {
                    for(var i = 0; i < dummyModel.count; i++) {
                      if(dummyModel.get(i).title === title) {
                        removeTransition.targetY = row.y - dp(20)
                        dummyModel.remove(i, 1)
                        break
                      }
                    }
                  }
                }

                // button add item
                IconButton {
                  visible: !row.isSubItem
                  icon: IconType.pluscircle
                  anchors.right: closeButton.left
                  anchors.verticalCenter: parent.verticalCenter
                  onClicked: {
                    for(var i = 0; i < dummyModel.count; i++) {
                      if(dummyModel.get(i).title === title) {
                        var subTitle = "Sub-"+title
                        var j = 1
                        while(i+j < dummyModel.count && dummyModel.get(i+j).title === subTitle) {
                          j++
                          subTitle = "Sub-"+subTitle
                        }
                        dummyModel.insert(i+j, { title: subTitle })
                        break
                      }
                    }
                  }
                }
              }

              // set up transition animations
              add: Transition {
                ParallelAnimation {
                  NumberAnimation { property: "opacity"; from: 0; to: 1; duration: 200 }
                  NumberAnimation { property: "x"; from: -width; duration: 600; easing.type: Easing.OutBack }
                }
              }
              addDisplaced: Transition {
                NumberAnimation { properties: "x,y"; duration: 100 }
              }
              remove: Transition {
                id: removeTransition
                property real targetY: 0
                ParallelAnimation {
                  NumberAnimation { property: "opacity"; to: 0; duration: 400 }
                  NumberAnimation { properties: "scale"; to: 0; duration: 400 }
                  NumberAnimation { id: yAnim; property: "y"; to: removeTransition.targetY; duration: 400 }
                }
              }
              removeDisplaced: Transition {
                NumberAnimation { properties: "x,y"; duration: 400 }
              }

              // fill with dummy data
              Component.onCompleted: { reset() }
              function reset() {
                dummyModel.clear()
                for(var i=0; i < 3; i++)
                  dummyModel.insert(i, { title: "Item "+i })
              }
            }
            AppButton {
              horizontalMargin: 0; verticalMargin: 0
              text: "Reset"; anchors.horizontalCenter: parent.horizontalCenter; onClicked: appListView.reset() }
          }
        } // Transition SectionContent

        // Graphical Effects
        SectionSpacer { }
        SectionHeader { icon: IconType.bolt; text: "Graphical Effects" }
        SectionDescription { text: "Many different graphical effects are available and can be applied to any UI element.\n
The displace effect displayed below adds a horizontal and vertical pixel offset based on the Felgo logo. The properties of such graphical effects can also be animated." }
        SectionContent { contentItem: AppImage {
            id: displacementImage
            height: dp(150)
            width: height / sourceSize.height * sourceSize.width
            source: "../../assets/devices.png"
            anchors.horizontalCenter: parent.horizontalCenter
            layer.enabled: true
            layer.effect: Displace {
              id: displaceEffect
              displacement: 0.2
              displacementSource: Rectangle {
                width: displacementImage.width
                height: displacementImage.height
                color: Qt.rgba(0.5,0.5,0,1)

                AppImage {
                  id: displaceLogo
                  anchors.fill: parent
                  fillMode: AppImage.PreserveAspectFit
                  source: "../../assets/felgo-logo.png"
                }

                PropertyAnimation {
                  target: displaceEffect
                  property: "displacement"
                  to: 0.3
                  running: true
                  duration: 1500
                  onStopped: {
                    to *= -1
                    start()
                  }
                  easing.type: Easing.OutBounce
                }
              } // displacement source
            } // Displace effect
          } // Displacement Image
        } // Displace Effect SectionContent

        // Blend
        SectionSpacer { }
        SectionHeader { text: "Blend Effect" }
        SectionDescription { text: "Choose between different blend modes when layering items.\n
Note: Not all available effect modes are shown. This also applies to the color and style effects below." }
        SectionContent { contentItem: Column {
            width: parent.width
            spacing: dp(Theme.navigationBar.defaultBarItemPadding)

            Row {
              spacing: parent.spacing
              anchors.horizontalCenter: parent.horizontalCenter
              height: blendComboBox.height

              AppText { text: "Blend Mode: "; anchors.verticalCenter: parent.verticalCenter }
              CustomComboBox {
                id: blendComboBox
                implicitWidth: dp(110) + 30
                model: ["Normal", "Addition", "Subtract", "Multiply", "Divide", "Exclusion", "DarkerColor", "LighterColor"]
                anchors.verticalCenter: parent.verticalCenter
              }
            }

            AppImage {
              id: blendImage
              height: dp(150)
              width: height / sourceSize.height * sourceSize.width
              source: "../../assets/devices.png"
              anchors.horizontalCenter: parent.horizontalCenter
              layer.enabled: true
              layer.effect: Blend {
                mode: blendComboBox.currentText
                foregroundSource: AppImage {
                  width: blendImage.width
                  height: blendImage.height
                  fillMode: AppImage.PreserveAspectFit
                  source: "../../assets/felgo-logo.png"
                }
              } // Blend Effect
            } // Blend Image
          }
        } // Blend Effect SectionContent

        // Color
        SectionSpacer { }
        SectionHeader { text: "Color Adjustment" }
        SectionDescription { text: "Easily change color, contrast or saturation of UI elements with ready-to-use effects." }
        SectionContent { contentItem: Column {
            width: parent.width
            spacing: dp(Theme.navigationBar.defaultBarItemPadding)
            Row {
              spacing: parent.spacing
              anchors.horizontalCenter: parent.horizontalCenter

              AppText { text: "Apply:"; anchors.verticalCenter: parent.verticalCenter }
              CustomComboBox {
                id: colorComboBox
                implicitWidth: dp(160) + 30
                model: ["BrightnessContrast", "ColorOverlay", "Colorize", "HueSaturation"]
                anchors.verticalCenter: parent.verticalCenter
                onCurrentIndexChanged: {
                  // apply new effect
                  switch(currentIndex) {
                  case 0: colorEffectImage.layer.effect = brightnessContrastEffect; break
                  case 1: colorEffectImage.layer.effect = colorOverlayEffect; break
                  case 2: colorEffectImage.layer.effect = colorizeEffect; break
                  case 3: colorEffectImage.layer.effect = hueSaturationEffect; break
                  }
                }
              }
            }
            AppImage {
              id: colorEffectImage
              height: dp(150)
              width: height / sourceSize.height * sourceSize.width
              source: "../../assets/devices.png"
              anchors.horizontalCenter: parent.horizontalCenter
              layer.enabled: true
              layer.effect: brightnessContrastEffect

              Component {
                id: brightnessContrastEffect
                BrightnessContrast { brightness: -0.5; contrast: 0.5 }
              }
              Component {
                id: colorOverlayEffect
                ColorOverlay { color: Qt.rgba(0,0,1,0.25) }
              }
              Component {
                id: colorizeEffect
                Colorize { hue: 0.7; saturation: 0.5 }
              }
              Component {
                id: hueSaturationEffect
                HueSaturation { hue: 0.3; saturation: 0.5 }
              }
            }
          }
        } // Color Effect SectionContent

        // Style
        SectionSpacer { }
        SectionHeader { text: "Improved Style" }
        SectionDescription { text: "Quickly improve the style of your app for example by adding a shadow, glow or blur effect. A dynamic blur effect is also used for the navigation bar of this page." }
        SectionContent { contentItem: Column {
            width: parent.width
            spacing: dp(Theme.navigationBar.defaultBarItemPadding)
            Row {
              spacing: parent.spacing
              anchors.horizontalCenter: parent.horizontalCenter

              AppText { text: "Add: "; anchors.verticalCenter: parent.verticalCenter }
              CustomComboBox {
                id: styleComboBox
                implicitWidth: dp(120) + 30
                model: ["DropShadow", "FastBlur", "RadialBlur", "Glow"]
                anchors.verticalCenter: parent.verticalCenter
                onCurrentIndexChanged: {
                  // apply new effect
                  switch(currentIndex) {
                  case 0: styleEffectImage.layer.effect = dropShadowEffect; break
                  case 1: styleEffectImage.layer.effect = fastBlurEffect; break
                  case 2: styleEffectImage.layer.effect = radialBlurEffect; break
                  case 3: styleEffectImage.layer.effect = glowEffect; break
                  }
                }
              }
            }
            AppImage {
              id: styleEffectImage
              height: dp(150)
              width: height / sourceSize.height * sourceSize.width
              source: "../../assets/felgo-logo.png"
              anchors.horizontalCenter: parent.horizontalCenter
              layer.enabled: true
              layer.effect: dropShadowEffect

              Component {
                id: dropShadowEffect
                DropShadow { horizontalOffset: 6; verticalOffset: 6
                  radius: 12.0; samples: 17; color: "#80000000" }
              }
              Component {
                id: fastBlurEffect
                FastBlur { radius: dp(8) }
              }
              Component {
                id: radialBlurEffect
                RadialBlur { samples: 24; angle: 30 }
              }
              Component {
                id: glowEffect
                Glow { radius: 16; samples: 17; color: Qt.rgba(1,1,0,0.5) }
              }
            }
          }
        } // Style Effect SectionContent

      } // column
    } // flickable

    // scroll indicator
    ScrollIndicator {
      flickable: pageFlickable
      z: 1
    }

    Dialog {
      id: easingDialog
      title: "Animation Easing"
      positiveActionLabel: "Cool!"
      negativeAction: false
      onAccepted: close()
      scaleAnimation.easing.type: Easing[easingComboBox.currentText]
      scaleAnimation.duration: easingComboBox.currentIndex === 0 ? 500 : 300

      AppText {
        y: dp(50)
        width: parent.width - 2 * dp(Theme.navigationBar.defaultBarItemPadding)
        anchors.horizontalCenter: parent.horizontalCenter
        horizontalAlignment: AppText.AlignHCenter
        text: "This dialog was opened with easing type \""+easingComboBox.currentText+"\"."
      }
    }
  }

  // logo will be visible behind page when using page effect ;-)
  Column {
    anchors.centerIn: flickArea
    width: flickArea.width
    spacing: dp(10)
    opacity: 0.5

    Row {
      anchors.horizontalCenter: parent.horizontalCenter
      width: parent.width * 0.75
      spacing: dp(10)

      AppImage {
        width: parent.width - hiddenText.width
        height: width / sourceSize.width * sourceSize.height
        source: "../../assets/felgo-logo.png"
        fillMode: AppImage.PreserveAspectFit
        anchors.verticalCenter: parent.verticalCenter
      }
      AppText {
        id: hiddenText
        width: parent.width * 0.75
        text: "Congratulations, you found this hidden message!"
        font.pixelSize: dp(10)
        wrapMode: Text.WordWrap
        anchors.verticalCenter: parent.verticalCenter
      }
    }
    AppText {
      text: "Shaders are awesome. ;-)"
      font.pixelSize: dp(12)
      anchors.horizontalCenter: parent.horizontalCenter
    }
  }

  // apply shader fx effect
  VideoFXEffect {
    id: fxEffect
    anchors.fill: flickArea
    divider.anchors.bottomMargin: fxControls.parameterPanelHeight
    sourceForShaderEffect.sourceItem: flickArea
    gripSize: fxControls.gripSize
  }

  // shader fx controls
  VideoFXControls {
    id: fxControls
    anchors.topMargin: background.height - scaledMargin
    targetEffect: fxEffect
    visible: false
    onEffectLoaded: {
      if(name === "No effect")
        fxControls.visible = false
    }

    function show() {
      fxControls.visible = true
      fxEffect.divider.value = 0.5
      fxControls.loadEffect("Emboss", "EffectEmboss.qml")
    }

    function hide() {
      fxControls.visible = false
      fxControls.loadEffect("No effect", "EffectPassThrough.qml")
    }
  }
}
