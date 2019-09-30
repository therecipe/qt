import QtQuick 2.0
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4
import Felgo 3.0
import "../common"

// feedback window to contact Felgo
Dialog {
  id: feedback
  negativeAction: true
  negativeActionLabel: "Close"
  positiveActionLabel: "Send"
  autoSize: true
  outsideTouchable: false

  property string originalHintText: feedbackInput.visible ?
                                      "Your feedback helps us to improve this app"
                                    : "You can also add your email address so we can reply to you."

  onIsOpenChanged: {
    feedbackInput.visible = true
    feedbackInput.text = feedbackInput.placeHolder
  }

  onCanceled: {
    logic.setFeedBackSent(true)
    feedback.close()
  }
  onAccepted: {
    // close the feedback window and send the feedback
    // don't continue if the email was incorrect
    if (!emailInput.checkEmail() || !feedbackInput.checkFeedback()){
      return
    }

    if (feedbackInput.visible){
      feedbackInput.visible = false
      hintText.text = originalHintText
    }

    // check if there has been feedback and send it to Felgo
    else if (feedbackInput.text){
      amplitude.logEvent("Send Feedback")

      // send the enteredText and the optional email to Felgo
      //nativeUtils.sendEmail("support@felgo.com", gameTitle + " Feedback", "What do you think about" + gameTitle + "? What do you like, what are you missing?\nPlease add your feedback here:\n\n")

      var feedbackContent = feedbackInput.text + "\n\n" +
          "\nApp Starts: " + dataModel.localAppStarts +
          "\nApp VersionCode: " + system.appVersionCode +
          "\nPlatform: " + Qt.platform.os +
          "\nosType: " + system.osType +
          "\nosVersion: " + system.osVersion +
          "\nDeviceModel: " + nativeUtils.deviceModel() +
          "\nFelgo Version: " + system.vplayVersion +
          "\nUDID: " + system.UDID +
          "\nUser Id: " + socialViewItem.gameNetworkItem.user.userId +
          "\nUser Name: " +socialViewItem.gameNetworkItem.userName

      console.debug("Feedback: " + feedbackContent + "; email: " + emailInput.text)
      sendFeedback(feedbackContent, emailInput.text)

      logic.setFeedBackSent(true)
      feedback.close()
    } else {
      hintText.text = "Please enter your opinion!"
    }
  }

  Item {
    id: contentArea
    width: parent.width
    height: content.height

    // catch the mouse clicks to cancel focus if touched outside
    MouseArea {
      width: app.width
      height: app.height
      anchors.centerIn: parent
      onClicked: {
        emailInput.focus = false
        feedbackInput.focus = false
      }
    }

    Column {
      id: content
      width: parent.width
      spacing: dp(10)

      // spacer
      Item {
        width: parent.width
        height: parent.spacing
      }

      // feedback header
      Text {
        id: headerText
        horizontalAlignment: Text.AlignHCenter
        anchors.horizontalCenter: parent.horizontalCenter
        text: "Send Feedback"
        color: Theme.textColor
        font.pixelSize: sp(18)
        width: parent.width * 0.8
        wrapMode: Text.Wrap
      }

      // feedback note
      Text {
        id: hintText
        horizontalAlignment: Text.AlignHCenter
        anchors.horizontalCenter: parent.horizontalCenter
        text: originalHintText
        color: Theme.textColor
        font.pixelSize: sp(9)
        width: parent.width * 0.8
        wrapMode: Text.Wrap
      }

      // TextInput line with validator
      TextField {
        id: emailInput
        anchors.horizontalCenter: parent.horizontalCenter
        width: parent.width * 0.8
        visible: !feedbackInput.visible

        horizontalAlignment: Text.AlignHCenter
        font.pixelSize: sp(10)
        maximumLength: 200
        placeholderText: focus ? "" : "Your email (optional)"
        inputMethodHints: Qt.ImhNoPredictiveText
        validator: RegExpValidator{regExp: /^[a-zA-Z0-9äöüßÄÖÜ;,:._'#+*~@€<>|?ß=()/&%!°^" -]+$/}

        // TextFieldStyle formatting the background of emailInput
        style: TextFieldStyle {
          textColor: Theme.textColor
          background: Rectangle {
            radius: height
            color: Theme.secondaryBackgroundColor
            anchors.margins: -dp(2)
          }
        }

        // disable and reset the inputField when closed
        onVisibleChanged: {
          readOnly = visible ? false : true
          if (!visible) focus = false
          text = ""
        }

        onAccepted: {
          checkEmail()
        }

        // check whether the email input is correct or not
        // has to be either empty or 4+ symbols with one of them being @
        function checkEmail(){
          if (!text || (text.match(/[@]/i) && text.length >= 4)){
            // @ found
            hintText.text = originalHintText
            return true
          } else {
            hintText.text = "Invalid email address!"
            return false
          }
        }
      }

      // multiple TextInput lines with validator
      TextArea {
        id: feedbackInput
        anchors.horizontalCenter: parent.horizontalCenter
        width: parent.width * 0.8
        height: emailInput.height * 4
        wrapMode: Text.WrapAtWordBoundaryOrAnywhere
        text: placeHolder

        horizontalAlignment: Text.AlignHCenter
        font.pixelSize: sp(10)
        inputMethodHints: Qt.ImhNoPredictiveText
        style: TextAreaStyle {backgroundColor: "transparent"; textColor: "black" }

        focus: false

        property string placeHolder: "Click here to add your feedback!"

        onFocusChanged: {
          console.debug("Focus changed: " + focus + ", Text: " + text)
          if (focus && text == placeHolder){
            feedbackInput.remove(0, text.length)
          } else if (!focus && text == ""){
            feedbackInput.append(placeHolder)
          }
        }

        // disable and reset the inputField when closed
        onVisibleChanged: {
          feedbackInput.readOnly = visible ? false : true
          if(!visible) {
            feedbackInput.focus = false
          }
        }

        // check whether the email input is correct or not
        // has to be 3+ symbols
        function checkFeedback(){
          if (text && actualInput(text) && text != placeHolder){
            hintText.text = originalHintText
            return true
          } else {
            hintText.text = "Please enter your feedback!"
            return false
          }
        }

        // check whether the feedback contains three or more actual characters or only spaces
        function actualInput(inputString){
          // remove spaces and breaks
          var trimmedString = inputString.replace(/^\s*/, "").replace(/\s*$/, "")
          if (trimmedString.length >= 3){
            // the trimmed string is more than 3 characters long
            return true
          } else {
            // the trimmed string is less than 3 characters long
            return false
          }
        }
      }

      // spacer
      Item {
        width: parent.width
        height: parent.spacing
      }
    }

  } // content area

  // sendFeedback - uses XMLHttpRequest object to send the feedback to the Felgo servers
  function sendFeedback(feedback, email) {
    var feedbackSecret = AppSettings.feedbackSecret

    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
      if (xhr.readyState === XMLHttpRequest.DONE) {
        console.debug("Successfully sent feedback to Felgo, response:", xhr.responseText)
      }
    }
    xhr.open("POST", AppSettings.feedbackUrl, true)
    xhr.setRequestHeader("Content-Type", "application/json")
    xhr.setRequestHeader("Accept", "application/json")
    var send = { "shared_secret": feedbackSecret, "subject": "QtWS 2017 Feedback", "message": feedback, "name": "", "from": email }
    console.debug("sending this feedback request:", JSON.stringify(send))
    xhr.send(JSON.stringify(send))
  }
}
