pragma Singleton
import Felgo 3.0
import QtQuick 2.0

Item {

  readonly property string licenseKey: ""

  readonly property int gameId: 406
  readonly property string gameSecret: "qtws2017github"

  readonly property string appKey: "<your-mp-appkey>"
  readonly property string pushKey: "<your-mp-pushkey>"

  readonly property string facebookAppId: "<your-fb-appid>"
  readonly property string amplitudeApiKey: "<your-amplitude-apikey>"

  readonly property string feedbackSecret: "<feedback-api-secret>"
  readonly property string feedbackUrl: "<feedback-api-url>"
}
