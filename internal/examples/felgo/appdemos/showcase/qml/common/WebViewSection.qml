import Felgo 3.0
import QtQuick 2.5
import QtWebView 1.1

// content
Column  {
  id: webViewSection
  width: parent.width
  anchors.horizontalCenter: parent.horizontalCenter

  // component for web view page
  property Component webViewPageComponent: Page {
    id: webViewPage
    title: webView.title
    property alias url: webView.url

    rightBarItem: ActivityIndicatorBarItem {
      visible: webView.loading
    }

    WebView {
      id: webView
      url: Qt.resolvedUrl("http://www.google.com")
      anchors.fill: parent
    }
  }

  // web view
  SectionSpacer { }
  SectionHeader { icon: IconType.code; text: "Embedded Web Content" }
  SectionDescription { text: "Use a WebView to render web content directly in your app." }
  SectionContent { contentItem: Row {
      anchors.horizontalCenter: parent.horizontalCenter
      AppButton { flat: false; text: "Browse"; onClicked: navigationStack.push(webViewPageComponent) }
    }
  }
} // content column

