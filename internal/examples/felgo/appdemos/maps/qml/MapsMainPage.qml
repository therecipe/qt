import Felgo 3.0
import QtQuick 2.0
import QtQuick.XmlListModel 2.0
import QtLocation 5.5
import QtPositioning 5.5
import Qt.labs.settings 1.0


Page {

  // Custom colors
  readonly property color blueLightColor: Qt.rgba(0, 147/256.0, 201/256.0, 1)
  // Generic text blue color
  readonly property color blueDarkColor: Qt.rgba(69/256.0, 85/256.0, 105/256.0, 1)
  // Background grey
  readonly property color greyBackgroundColor: Qt.rgba(238/256.0, 238/256.0, 238/256.0, 1)
  // Grey line color
  readonly property color greyLineColor: Qt.rgba(221/256.0, 221/256.0, 221/256.0, 1)

  // Model
  XmlListModel {
    id: stationsModel

    source: "http://dynamisch.citybikewien.at/citybike_xml.php"
    query: "/stations/station"

    XmlRole { name: "internalId"; query: "internal_id/number()" }
    XmlRole { name: "name"; query: "name/string()" }
    XmlRole { name: "description"; query: "description/string()" }

    XmlRole { name: "latitude"; query: "latitude/number()" }
    XmlRole { name: "longitude"; query: "longitude/number()" }

    XmlRole { name: "availability"; query: "status/string()" }
    XmlRole { name: "boxes"; query: "boxes/string()" }
    XmlRole { name: "freeBoxes"; query: "free_boxes/number()" }
    XmlRole { name: "freeBikes"; query: "free_bikes/number()" }

    onStatusChanged: {
      if (status === XmlListModel.Ready) {
        favsModel.update()
      }
    }
  }

  // Settings
  Settings {
    id: settings

    property var favorites
    onFavoritesChanged: {
      console.debug("Updated favorites", JSON.stringify(favorites))
      favsModel.update()
    }

    Component.onCompleted: {
      if (favorites === undefined) {
        // Add a first favorite for demonstration purposes
        favorites = [ 1021 ]
      }
    }
  }

  ListModel {
    id: favsModel

    function update() {
      clear()

      // Nearest from GPS coords
      updateNearest()

      // Add favs
      for (var i = 0; i < stationsModel.count; i++) {
        var station = stationsModel.get(i)
        if (settings.favorites.indexOf(station.internalId) !== -1) {
          append({
                   internalId: station.internalId,
                   name: station.name,
                   freeBoxes: station.freeBoxes,
                   freeBikes: station.freeBikes,
                   favorited: settings.favorites.indexOf(station.internalId) !== -1
                 })
        }
      }
    }

    function updateNearest() {
      if (!map.userPositionAvailable)
        return

      // Compare current user location with all other stations
      var currentDistance = -1
      var currentIndex = -1

      for (var i = 0; i < stationsModel.count; i++) {
        var station = stationsModel.get(i)

        var distance = map.userPosition.coordinate.distanceTo(QtPositioning.coordinate(station.latitude, station.longitude))
        if (currentDistance === -1 || distance < currentDistance) {
          currentDistance = distance
          currentIndex = i
        }
      }

      // Get station
      if (currentIndex !== -1) {
        var station = stationsModel.get(currentIndex)

        set(0, {
              internalId: station.internalId,
              name: station.name,
              freeBoxes: station.freeBoxes,
              freeBikes: station.freeBikes,
              favorited: settings.favorites.indexOf(station.internalId) !== -1
            })
      }
    }
  }


  NavigationStack {
    splitView: false

    Page {
      id: page

      navigationBarHidden: true

      property bool displayFreeBikes: true
      property int selectedIndex: -1

      onSelectedIndexChanged: {
        var selectedStation = stationsModel.get(selectedIndex)
        if (selectedStation) {
          currentStationView.stationName = selectedStation.name
          currentStationView.stationBikes = selectedStation.freeBikes
          currentStationView.stationBoxes = selectedStation.freeBoxes
          currentStationView.stationFavorited = settings.favorites.indexOf(selectedStation.internalId) !== -1
        }
      }

      // Header
      Rectangle {
        id: header

        width: parent.width
        height: dp(210)
        // Bring in front of map
        z: 2

        color: greyBackgroundColor

        PageControl {
          id: pageControl

          anchors.horizontalCenter: parent.horizontalCenter
          anchors.top: parent.top
          anchors.topMargin: Theme.statusBarHeight + dp(10)

          opacity: page.selectedIndex >= 0 ? 0 : 1

          pageIcons: map.userPositionAvailable ? ({ 0: IconType.locationarrow }) : ({})

          tintColor: greyLineColor
          activeTintColor: blueLightColor

          onPageSelected: {
            innerList.scrollToPage(index)
          }

          pages: favsModel.count
        }

        ListView {
          id: innerList

          onContentXChanged: {
            currentIndex = Math.round(contentX / width)
            pageControl.currentPage = currentIndex
          }

          anchors.fill: parent
          anchors.topMargin: pageControl.y + pageControl.height + dp(10)

          visible: !currentStationView.visible

          model: favsModel
          orientation: ListView.Horizontal
          snapMode: ListView.SnapOneItem
          highlightFollowsCurrentItem: true

          delegate: StationView {
            width: innerList.width
            height: innerList.height

            stationName: name
            stationBikes: freeBikes
            stationBoxes: freeBoxes
            stationFavorited: favorited

            // We don't use the property signal change handler as this one is also emitted when
            // changing the property in code
            onFavoritedPressed: {
              if (settings.favorites.indexOf(internalId) === -1) {
                settings.favorites.push(internalId)
              }
              else {
                settings.favorites.splice(settings.favorites.indexOf(internalId), 1)
              }

              settings.favoritesChanged()
            }
          }

          function scrollToPage(index) {
            // TODO: Animate
            innerList.positionViewAtIndex(index, ListView.SnapPosition)
          }
        }

        StationView {
          id: currentStationView
          anchors.fill: innerList

          visible: page.selectedIndex >= 0

          backVisible: true

          onBackPressed: {
            page.selectedIndex = -1
          }

          // We don't use the property signal change handler as this one is also emitted when
          // changing the property in code
          onFavoritedPressed: {
            var internalId = stationsModel.get(page.selectedIndex).internalId

            if (settings.favorites.indexOf(internalId) === -1) {
              settings.favorites.push(internalId)
            }
            else {
              settings.favorites.splice(settings.favorites.indexOf(internalId), 1)
            }

            settings.favoritesChanged()
          }
        }

        // Drop a shadow on bottom of header
        Rectangle {
          anchors.top: parent.bottom
          width: parent.width
          height: dp(5)

          gradient: Gradient {
            GradientStop { position: 0.0; color: "#33000000" }
            GradientStop { position: 1.0; color: "transparent" }
          }
        }

      }

      AppMap {
        id: map

        anchors.top: header.bottom
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.bottom: parent.bottom

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

        // Defaults to Vienna, AT
        center: QtPositioning.coordinate(48.208417, 16.372472)
        zoomLevel: 15
        Component.onCompleted: {
          map.center = QtPositioning.coordinate(48.208417, 16.372472)
          map.zoomLevel = 15
        }

        onMapClicked: {
          // Clicked on map, remove current selection
          page.selectedIndex = -1
        }

        onUserPositionChanged: {
          favsModel.updateNearest()
        }

        // Station markers
        MapItemView {
          model: stationsModel

          delegate: MapQuickItem {
            coordinate: QtPositioning.coordinate(latitude, longitude)

            anchorPoint.x: image.width * 0.5
            anchorPoint.y: image.height

            sourceItem: AppImage {
              id: image

              width: dp(40)
              height: dp(34)

              source: {
                // Inactive
                if (availability !== "aktiv") {
                  return "../assets/pin-grey.png"
                }

                var freeItems = page.displayFreeBikes ? freeBikes : freeBoxes

                if (freeItems === 0) {
                  return "../assets/pin-red.png"
                }
                else if (freeItems <= 2) {
                  return "../assets/pin-orange.png"
                }
                else {
                  return "../assets/pin-green.png"
                }
              }

              MouseArea {
                anchors.fill: parent
                onClicked: {
                  page.selectedIndex = index
                }
              }
            }
          }
        }
      }

      IconButton {
        icon: IconType.locationarrow
        anchors.left: parent.left
        anchors.bottom: parent.bottom
        anchors.leftMargin: dp(10)
        anchors.bottomMargin: dp(10)

        enabled: map.userPositionAvailable

        size: dp(26)

        onClicked: {
          map.zoomToUserPosition()
        }
      }
    }
  }
}
