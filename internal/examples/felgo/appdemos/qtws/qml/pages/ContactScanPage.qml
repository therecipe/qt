import Felgo 3.0
import QtQuick 2.0
import QtMultimedia 5.5
import QZXing 2.3 // barcode scanning
import "../common"

Page {
  id: page
  title: "Add Contact"
  backgroundColor: "black"
  rightBarItem: ActivityIndicatorBarItem {
    animating: page.busy
    visible: animating
  }

  // shows activity indicator in navigation bar and scanner UI
  property bool busy: false

  // holds value of scanned QR tag
  property string lastScannedTag: ""
  signal tagFound(string tag)

  // initialize scanner when page is created
  Component.onCompleted: initializeScanner()

  // when tag was found, try loading contact
  onTagFound: {
    page.busy = true
    captureText.visible = false
    captureZone.color = "green"

    logic.loadContact(tag, function(data) {
      // success handler
      page.busy = false

      // add time of scan and scanned tag
      var time = Date.now() // UTC milliseconds timestamp
      data["scan_timestamp"] = time
      data["scan_tag"] = tag

      amplitude.logEvent("Card Scan Successful",{"tag" : tag, "data": data, "name": data.name})

      // check if contact already exists
      var contactAlreadyExists = dataModel.contacts && dataModel.contacts[tag]

      // add/overwrite contact in storage
      logic.addContact(tag, data)

      // show contact, also show dialog if existing contact was updated
      if(!contactAlreadyExists) {
        page.navigationStack.pop(getPage(navigationStack.depth - 2)) // removes content scan page
        page.navigationStack.push(Qt.resolvedUrl("ContactDetailPage.qml"), { contactId: tag, contact: data }) // show contact detail
      }
      else {
        NativeDialog.confirm("Contact Exists", "You already have this contact in your list, your contact was updated.", function(accept) {
          page.navigationStack.pop(getPage(navigationStack.depth - 2)) // removes content scan page
          page.navigationStack.push(Qt.resolvedUrl("ContactDetailPage.qml"), { contactId: tag, contact: data }) // show contact detail
        }, false)
      }
    }, function() {
      // error handler
      page.busy = false
      captureZone.color = "red"
      NativeDialog.confirm("Failed to fetch data","Please scan a valid code and make sure your device has internet access.", function() {
        initializeScanner()
      }, false)
    })
  }

  // item for scanner, allows to show / remove camera live capturing as needed
  Item {
    anchors.fill: parent

    // initialize capture zone size
    Component.onCompleted: {
      captureZone.width = Qt.binding(function() { return videoOutput.contentRect.width / 2 })
      captureZone.height = Qt.binding(function() { return videoOutput.contentRect.height / 2 })
    }

    // Camera
    Camera {
      id:camera
      focus {
        focusMode: CameraFocus.FocusContinuous
        focusPointMode: CameraFocus.FocusPointAuto
      }
    }

    // Live Video Output
    VideoOutput {
      id: videoOutput
      source: camera
      anchors.fill: parent
      autoOrientation: true
      fillMode: VideoOutput.PreserveAspectCrop
      filters: [ zxingFilter ]
      MouseArea {
        anchors.fill: parent
        onClicked: {
          camera.focus.customFocusPoint = Qt.point(mouse.x / width,  mouse.y / height);
          camera.focus.focusMode = CameraFocus.FocusMacro;
          camera.focus.focusPointMode = CameraFocus.FocusPointCustom;
        }
      }
    }

    // Barcode Scanner Video Filter
    QZXingFilter {
      id: zxingFilter
      // setup capture area
      captureRect: {
        videoOutput.contentRect;
        videoOutput.sourceRect;
        return videoOutput.mapRectToSource(videoOutput.mapNormalizedRectToItem(Qt.rect(
                                                                                 0.25, 0.25, 0.5, 0.5
                                                                                 )));
      }
      // set up decoder
      decoder {
        enabledDecoders: QZXing.DecoderFormat_QR_CODE
        onTagFound: {
          console.log(tag + " | " + decoder.foundedFormat() + " | " + decoder.charSet());
          zxingFilter.active = false

          // trigger tagfound signal for scanned tag
          page.lastScannedTag = tag // store scanned tag
          page.tagFound(page.lastScannedTag)
        }
        tryHarder: false
      }
    }
  } // QR Scanner Item

  // Scanner UI
  Item {
    id: scannerUI
    anchors.fill: parent
    z: 1

    // Capture Zone Overlay
    Rectangle {
      id: captureZone
      color: "red"
      opacity: 0.2
      anchors.centerIn: parent
    }

    // Capture Text Overlay
    AppText {
      id: captureText
      anchors.centerIn: parent
      color: "white"
      text: "[place QR code\nin this area]"
      font.pixelSize: sp(12)
      horizontalAlignment: AppText.AlignHCenter
      font.bold: true
    }

    // Activity Indicator Overlay
    AppActivityIndicator {
      animating: page.busy
      visible: animating
      anchors.centerIn: parent
      color: "white"
    }
  } // Scanner UI

  // initializeScanner - resets UI and creates camers live scanner item
  function initializeScanner() {
    // reset UI
    captureText.visible = true
    captureZone.color = "red"
    zxingFilter.active = true
  }
}
