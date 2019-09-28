import Felgo 3.0
import QtQuick 2.7

Item {
  id: dataModel

  // property to configure target dispatcher / logic
  property alias dispatcher: logicConnection.target

  readonly property alias weatherAvailable: _.weatherAvailable
  readonly property alias weatherFromCache: _.weatherFromCache

  readonly property alias weatherData: _.weatherData

  signal loadJsonDataFailed(string error)

  Connections {
    id: logicConnection

    // action 1 - loadJsonData
    onLoadJsonData: {
      // Build query URL
      var params = "q=" + weatherSearchBar.text + "&units=metric&appid=" + app.weatherServiceAppId
      console.log("Query URL: " + "http://api.openweathermap.org/data/2.5/weather?" + params)

      HttpRequest.get("http://api.openweathermap.org/data/2.5/weather?" + params)
      .then(function(res) {
          console.log("DONE: " + res.text)
          var parsedWeather = res.text ? JSON.parse(res.text) : null

          if (parsedWeather && parsedWeather.cod === 200) {
            _.updateFromJson(parsedWeather)
          } else {
             if (parsedWeather && parsedWeather.message) {
              // Received a response, but the server reported the request was not successful
              loadJsonDataFailed(parsedWeather.message)
            } else {
              // All other cases - print the HTTP response status code / message
              loadJsonDataFailed("Request error: " + res.status + " / " + res.text)
            }
          }
      })
      .catch(function(err) {
        loadJsonDataFailed("Request error: "+err.message)
      })
    }

    // action 2 - clearData
    onClearData: {
      // Reset all data stored in the model and the cache
      _.setModelData(false, "", undefined, undefined, "", "", false)
    }
  }

  Storage {
    id: weatherLocalStorage

    Component.onCompleted: {
      // After the storage has been initialized, check if any weather data is cached.
      // If yes, load it into our model.
      _.loadModelFromStorage()
    }
  }

  // private
  Item {
    id: _

    property bool weatherAvailable: false
    property bool weatherFromCache: false

    property var weatherData: []

    function updateFromJson(parsedWeatherJson) {
      // Use the new parsed JSON file to update the model and the cache
      setModelData(true,
                   parsedWeatherJson.name + ", " + parsedWeatherJson.sys.country,
                   new Date(),    // Note: date.now() and new Date() are different - new Date() returns a QML Date object!
                   parsedWeatherJson.main.temp,
                   parsedWeatherJson.weather[0].main,
                   "http://openweathermap.org/img/w/" + parsedWeatherJson.weather[0].icon + ".png",
                   false)
    }

    function setModelData(weatherAvailable, weatherForCity, weatherDate, weatherTemp, weatherCondition, weatherIconUrl, weatherFromCache) {
      _.weatherData = {
        'weatherForCity': weatherForCity,
        'weatherDate': weatherDate,
        'weatherTemp': weatherTemp,
        'weatherCondition': weatherCondition,
        'weatherIconUrl': weatherIconUrl
      }
      console.log("Saved weather to dataModel")
      console.log(JSON.stringify(_.weatherData))

      _.weatherAvailable = weatherAvailable
      _.weatherFromCache = weatherFromCache
      saveModelToStorage()
    }

    function saveModelToStorage() {
      weatherLocalStorage.setValue("weatherData", _.weatherData)
    }

    function loadModelFromStorage() {
      console.log("Loading model from storage...")
      var savedWeatherData = weatherLocalStorage.getValue("weatherData")
      if (savedWeatherData) {
        console.log("Found data in cache!")
        _.weatherData = savedWeatherData
        _.weatherAvailable = true
        _.weatherFromCache = true
        console.log(JSON.stringify(_.weatherData))
      }
    }
  }
}
