package main

import (
	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
)

type XYSeriesIODevice struct {
	core.QIODevice

	//TODO: custom constructors
	//_ func(series *charts.QXYSeries) `constructor:"init"`

	m_series *charts.QXYSeries
}

func (d *XYSeriesIODevice) init(series *charts.QXYSeries) {
	d.m_series = series

	d.ConnectReadData(d.readData)
	d.ConnectWriteData(d.writeData)
}

func (d *XYSeriesIODevice) readData(data *string, maxSize int64) int64 {
	return -1
}

func (d *XYSeriesIODevice) writeData(data []byte, maxSize int64) int64 {
	var rang = 2000
	var oldPoints = d.m_series.PointsVector()
	var points []*core.QPointF
	var resolution int64 = 4

	if len(oldPoints) < rang {
		points = d.m_series.PointsVector()
	} else {
		for i := int64(maxSize / resolution); i < int64(len(oldPoints)); i++ {
			points = append(points, core.NewQPointF3(float64(i-maxSize/resolution), oldPoints[i].Y()))
		}
	}

	var size = int64(len(points))
	for k := int64(0); k < int64(maxSize/resolution); k++ {
		var y = float64(data[resolution*k]-128) / 128.0
		if y > 1 {
			y -= 2
		}
		points = append(points, core.NewQPointF3(float64(k+size), y))
	}

	d.m_series.Replace5(points)
	return maxSize
}
