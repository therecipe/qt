package main

import (
	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
)

type Widget struct {
	widgets.QWidget

	_ func() `constructor:"init"`

	m_device     *XYSeriesIODevice
	m_chart      *charts.QChart
	m_series     *charts.QLineSeries
	m_audioInput *multimedia.QAudioInput
}

func (w *Widget) init() {
	w.m_chart = charts.NewQChart(nil, 0)
	var chartView = charts.NewQChartView2(w.m_chart, w)
	chartView.SetMinimumSize2(800, 600)
	w.m_series = charts.NewQLineSeries(nil)
	w.m_chart.AddSeries(w.m_series)
	var axisX = charts.NewQValueAxis(nil)
	axisX.SetRange(0, 2000)
	axisX.SetLabelFormat("%g")
	axisX.SetTitleText("Samples")
	var axisY = charts.NewQValueAxis(nil)
	axisY.SetRange(-1, 1)
	axisY.SetTitleText("Audio level")
	w.m_chart.SetAxisX(axisX, w.m_series)
	w.m_chart.SetAxisY(axisY, w.m_series)
	w.m_chart.Legend().Hide()
	w.m_chart.SetTitle("Data from the microphone")

	var mainLayout = widgets.NewQVBoxLayout()
	mainLayout.AddWidget(chartView, 0, 0)
	w.SetLayout(mainLayout)

	var formatAudio = multimedia.NewQAudioFormat()
	formatAudio.SetSampleRate(8000)
	formatAudio.SetChannelCount(1)
	formatAudio.SetSampleSize(8)
	formatAudio.SetCodec("audio/pcm")
	formatAudio.SetByteOrder(multimedia.QAudioFormat__LittleEndian)
	formatAudio.SetSampleType(multimedia.QAudioFormat__UnSignedInt)

	var inputDevices = multimedia.QAudioDeviceInfo_DefaultInputDevice()
	w.m_audioInput = multimedia.NewQAudioInput2(inputDevices, formatAudio, w)

	w.m_device = NewXYSeriesIODevice2(w)
	w.m_device.init(w.m_series.QXYSeries_PTR())
	w.m_device.Open(core.QIODevice__WriteOnly)

	w.m_audioInput.Start(w.m_device)

	w.ConnectDestroyed(w.destroyed)
}

func (w *Widget) destroyed(*core.QObject) {
	w.m_audioInput.Stop()
	w.m_device.Close()
}
