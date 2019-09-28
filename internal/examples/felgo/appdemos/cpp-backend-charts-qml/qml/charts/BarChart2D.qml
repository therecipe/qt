import QtQuick 2.0
import QtCharts 2.2

Item {
  anchors.fill: parent

  ChartView {
    id: chartView
    width: parent.width
    height: parent.height * 0.5

    BarSeries {
      id: mySeries

      axisX: BarCategoryAxis {
        categories: {
          var arr = []
          var used = {}
          for(var i = 0; i < chartData.length; i++) {
            var year = chartData[i].year
            if(!used[year]) {
              arr.push(year)
              used[year] = true
            }
          }
          return arr
        }
      }
      axisY: ValueAxis {
        min: 0
        max: 50000
      }

      Connections {
        target: app
        onChartDataChanged: mySeries.udpateBarSeries()
      }
      Component.onCompleted: mySeries.udpateBarSeries()

      function udpateBarSeries() {
        mySeries.clear()
        var expenses = {}
        var maxExpense = 0
        for(var i = 0; i < chartData.length; i++) {
          var model = chartData[i]
          if(expenses[model.city] === undefined)
            expenses[model.city] = []

          expenses[model.city].push(parseInt(model.expenses))
          if(model.expenses > maxExpense)
            maxExpense = model.expenses
        }

        mySeries.axisY.max = maxExpense
        for(var city in expenses) {
          mySeries.append(city, expenses[city])
        }
      }
    }
  }
} // ChartViewContainer
