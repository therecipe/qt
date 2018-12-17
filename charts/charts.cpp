// +build !minimal

#define protected public
#define private public

#include "charts.h"
#include "_cgo_export.h"

#include <QAbstractAxis>
#include <QAbstractBarSeries>
#include <QAbstractItemModel>
#include <QAbstractSeries>
#include <QAction>
#include <QActionEvent>
#include <QAreaLegendMarker>
#include <QAreaSeries>
#include <QBarCategoryAxis>
#include <QBarLegendMarker>
#include <QBarSeries>
#include <QBarSet>
#include <QBoxPlotLegendMarker>
#include <QBoxPlotSeries>
#include <QBoxSet>
#include <QBrush>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QCandlestickLegendMarker>
#include <QCandlestickModelMapper>
#include <QCandlestickSeries>
#include <QCandlestickSet>
#include <QCategoryAxis>
#include <QChart>
#include <QChartView>
#include <QChildEvent>
#include <QCloseEvent>
#include <QColor>
#include <QContextMenuEvent>
#include <QDBusPendingCallWatcher>
#include <QDateTime>
#include <QDateTimeAxis>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEasingCurve>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFocusEvent>
#include <QFont>
#include <QGraphicsItem>
#include <QGraphicsObject>
#include <QGraphicsScene>
#include <QGraphicsSceneContextMenuEvent>
#include <QGraphicsSceneDragDropEvent>
#include <QGraphicsSceneHoverEvent>
#include <QGraphicsSceneMouseEvent>
#include <QGraphicsSceneMoveEvent>
#include <QGraphicsSceneResizeEvent>
#include <QGraphicsSceneWheelEvent>
#include <QGraphicsTransform>
#include <QGraphicsWidget>
#include <QHBarModelMapper>
#include <QHBoxPlotModelMapper>
#include <QHCandlestickModelMapper>
#include <QHPieModelMapper>
#include <QHXYModelMapper>
#include <QHideEvent>
#include <QHorizontalBarSeries>
#include <QHorizontalPercentBarSeries>
#include <QHorizontalStackedBarSeries>
#include <QIcon>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLayout>
#include <QLegend>
#include <QLegendMarker>
#include <QLineSeries>
#include <QLocale>
#include <QLogValueAxis>
#include <QMargins>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPdfWriter>
#include <QPen>
#include <QPercentBarSeries>
#include <QPieLegendMarker>
#include <QPieSeries>
#include <QPieSlice>
#include <QPoint>
#include <QPointF>
#include <QPolarChart>
#include <QQuickItem>
#include <QRadioData>
#include <QRect>
#include <QRectF>
#include <QResizeEvent>
#include <QScatterSeries>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QSplineSeries>
#include <QStackedBarSeries>
#include <QString>
#include <QStyleOption>
#include <QStyleOptionGraphicsItem>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QVBarModelMapper>
#include <QVBoxPlotModelMapper>
#include <QVCandlestickModelMapper>
#include <QVPieModelMapper>
#include <QVXYModelMapper>
#include <QValueAxis>
#include <QVariant>
#include <QVector>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>
#include <QXYLegendMarker>
#include <QXYSeries>

typedef QtCharts::QAbstractAxis QAbstractAxis;
typedef QtCharts::QAbstractBarSeries QAbstractBarSeries;
typedef QtCharts::QAbstractSeries QAbstractSeries;
typedef QtCharts::QAreaLegendMarker QAreaLegendMarker;
typedef QtCharts::QAreaSeries QAreaSeries;
typedef QtCharts::QBarCategoryAxis QBarCategoryAxis;
typedef QtCharts::QBarLegendMarker QBarLegendMarker;
typedef QtCharts::QBarSeries QBarSeries;
typedef QtCharts::QBarSet QBarSet;
typedef QtCharts::QBoxPlotLegendMarker QBoxPlotLegendMarker;
typedef QtCharts::QBoxPlotSeries QBoxPlotSeries;
typedef QtCharts::QBoxSet QBoxSet;
typedef QtCharts::QCandlestickLegendMarker QCandlestickLegendMarker;
typedef QtCharts::QCandlestickModelMapper QCandlestickModelMapper;
typedef QtCharts::QCandlestickSeries QCandlestickSeries;
typedef QtCharts::QCandlestickSet QCandlestickSet;
typedef QtCharts::QCategoryAxis QCategoryAxis;
typedef QtCharts::QChart QChart;
typedef QtCharts::QChartView QChartView;
typedef QtCharts::QDateTimeAxis QDateTimeAxis;
typedef QtCharts::QHBarModelMapper QHBarModelMapper;
typedef QtCharts::QHBoxPlotModelMapper QHBoxPlotModelMapper;
typedef QtCharts::QHCandlestickModelMapper QHCandlestickModelMapper;
typedef QtCharts::QHPieModelMapper QHPieModelMapper;
typedef QtCharts::QHXYModelMapper QHXYModelMapper;
typedef QtCharts::QHorizontalBarSeries QHorizontalBarSeries;
typedef QtCharts::QHorizontalPercentBarSeries QHorizontalPercentBarSeries;
typedef QtCharts::QHorizontalStackedBarSeries QHorizontalStackedBarSeries;
typedef QtCharts::QLegend QLegend;
typedef QtCharts::QLegendMarker QLegendMarker;
typedef QtCharts::QLineSeries QLineSeries;
typedef QtCharts::QLogValueAxis QLogValueAxis;
typedef QtCharts::QPercentBarSeries QPercentBarSeries;
typedef QtCharts::QPieLegendMarker QPieLegendMarker;
typedef QtCharts::QPieSeries QPieSeries;
typedef QtCharts::QPieSlice QPieSlice;
typedef QtCharts::QPolarChart QPolarChart;
typedef QtCharts::QScatterSeries QScatterSeries;
typedef QtCharts::QSplineSeries QSplineSeries;
typedef QtCharts::QStackedBarSeries QStackedBarSeries;
typedef QtCharts::QVBarModelMapper QVBarModelMapper;
typedef QtCharts::QVBoxPlotModelMapper QVBoxPlotModelMapper;
typedef QtCharts::QVCandlestickModelMapper QVCandlestickModelMapper;
typedef QtCharts::QVPieModelMapper QVPieModelMapper;
typedef QtCharts::QVXYModelMapper QVXYModelMapper;
typedef QtCharts::QValueAxis QValueAxis;
typedef QtCharts::QXYLegendMarker QXYLegendMarker;
typedef QtCharts::QXYSeries QXYSeries;

class MyQAbstractAxis: public QAbstractAxis
{
public:
	void Signal_ColorChanged(QColor color) { callbackQAbstractAxis_ColorChanged(this, new QColor(color)); };
	void Signal_GridLineColorChanged(const QColor & color) { callbackQAbstractAxis_GridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_GridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_GridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_GridVisibleChanged(bool visible) { callbackQAbstractAxis_GridVisibleChanged(this, visible); };
	void Signal_LabelsAngleChanged(int angle) { callbackQAbstractAxis_LabelsAngleChanged(this, angle); };
	void Signal_LabelsBrushChanged(const QBrush & brush) { callbackQAbstractAxis_LabelsBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_LabelsColorChanged(QColor color) { callbackQAbstractAxis_LabelsColorChanged(this, new QColor(color)); };
	void Signal_LabelsFontChanged(const QFont & font) { callbackQAbstractAxis_LabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_LabelsVisibleChanged(bool visible) { callbackQAbstractAxis_LabelsVisibleChanged(this, visible); };
	void Signal_LinePenChanged(const QPen & pen) { callbackQAbstractAxis_LinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_LineVisibleChanged(bool visible) { callbackQAbstractAxis_LineVisibleChanged(this, visible); };
	void Signal_MinorGridLineColorChanged(const QColor & color) { callbackQAbstractAxis_MinorGridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_MinorGridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_MinorGridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_MinorGridVisibleChanged(bool visible) { callbackQAbstractAxis_MinorGridVisibleChanged(this, visible); };
	void Signal_ReverseChanged(bool reverse) { callbackQAbstractAxis_ReverseChanged(this, reverse); };
	void Signal_ShadesBorderColorChanged(QColor color) { callbackQAbstractAxis_ShadesBorderColorChanged(this, new QColor(color)); };
	void Signal_ShadesBrushChanged(const QBrush & brush) { callbackQAbstractAxis_ShadesBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_ShadesColorChanged(QColor color) { callbackQAbstractAxis_ShadesColorChanged(this, new QColor(color)); };
	void Signal_ShadesPenChanged(const QPen & pen) { callbackQAbstractAxis_ShadesPenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_ShadesVisibleChanged(bool visible) { callbackQAbstractAxis_ShadesVisibleChanged(this, visible); };
	void Signal_TitleBrushChanged(const QBrush & brush) { callbackQAbstractAxis_TitleBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_TitleFontChanged(const QFont & font) { callbackQAbstractAxis_TitleFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_TitleTextChanged(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtCharts_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQAbstractAxis_TitleTextChanged(this, textPacked); };
	void Signal_TitleVisibleChanged(bool visible) { callbackQAbstractAxis_TitleVisibleChanged(this, visible); };
	void Signal_VisibleChanged(bool visible) { callbackQAbstractAxis_VisibleChanged(this, visible); };
	 ~MyQAbstractAxis() { callbackQAbstractAxis_DestroyQAbstractAxis(this); };
	QAbstractAxis::AxisType type() const { return static_cast<QAbstractAxis::AxisType>(callbackQAbstractAxis_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractAxis_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQAbstractAxis_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractAxis_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractAxis_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractAxis_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractAxis_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractAxis_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractAxis_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractAxis_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAbstractAxis*)

int QAbstractAxis_QAbstractAxis_QRegisterMetaType(){qRegisterMetaType<QAbstractAxis*>(); return qRegisterMetaType<MyQAbstractAxis*>();}

void* QAbstractAxis_GridLineColor(void* ptr)
{
	return new QColor(static_cast<QAbstractAxis*>(ptr)->gridLineColor());
}

void* QAbstractAxis_MinorGridLineColor(void* ptr)
{
	return new QColor(static_cast<QAbstractAxis*>(ptr)->minorGridLineColor());
}

struct QtCharts_PackedString QAbstractAxis_QAbstractAxis_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t51d483 = QAbstractAxis::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t51d483.prepend("WHITESPACE").constData()+10), t51d483.size()-10 }; });
}

struct QtCharts_PackedString QAbstractAxis_QAbstractAxis_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t96fa0d = QAbstractAxis::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t96fa0d.prepend("WHITESPACE").constData()+10), t96fa0d.size()-10 }; });
}

void QAbstractAxis_ConnectColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(QColor)>(&QAbstractAxis::colorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(QColor)>(&MyQAbstractAxis::Signal_ColorChanged));
}

void QAbstractAxis_DisconnectColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(QColor)>(&QAbstractAxis::colorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(QColor)>(&MyQAbstractAxis::Signal_ColorChanged));
}

void QAbstractAxis_ColorChanged(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->colorChanged(*static_cast<QColor*>(color));
}

void QAbstractAxis_ConnectGridLineColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QColor &)>(&QAbstractAxis::gridLineColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QColor &)>(&MyQAbstractAxis::Signal_GridLineColorChanged));
}

void QAbstractAxis_DisconnectGridLineColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QColor &)>(&QAbstractAxis::gridLineColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QColor &)>(&MyQAbstractAxis::Signal_GridLineColorChanged));
}

void QAbstractAxis_GridLineColorChanged(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->gridLineColorChanged(*static_cast<QColor*>(color));
}

void QAbstractAxis_ConnectGridLinePenChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QPen &)>(&QAbstractAxis::gridLinePenChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QPen &)>(&MyQAbstractAxis::Signal_GridLinePenChanged));
}

void QAbstractAxis_DisconnectGridLinePenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QPen &)>(&QAbstractAxis::gridLinePenChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QPen &)>(&MyQAbstractAxis::Signal_GridLinePenChanged));
}

void QAbstractAxis_GridLinePenChanged(void* ptr, void* pen)
{
	static_cast<QAbstractAxis*>(ptr)->gridLinePenChanged(*static_cast<QPen*>(pen));
}

void QAbstractAxis_ConnectGridVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::gridVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_GridVisibleChanged));
}

void QAbstractAxis_DisconnectGridVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::gridVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_GridVisibleChanged));
}

void QAbstractAxis_GridVisibleChanged(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->gridVisibleChanged(visible != 0);
}

void QAbstractAxis_Hide(void* ptr)
{
	static_cast<QAbstractAxis*>(ptr)->hide();
}

void QAbstractAxis_ConnectLabelsAngleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(int)>(&QAbstractAxis::labelsAngleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(int)>(&MyQAbstractAxis::Signal_LabelsAngleChanged));
}

void QAbstractAxis_DisconnectLabelsAngleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(int)>(&QAbstractAxis::labelsAngleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(int)>(&MyQAbstractAxis::Signal_LabelsAngleChanged));
}

void QAbstractAxis_LabelsAngleChanged(void* ptr, int angle)
{
	static_cast<QAbstractAxis*>(ptr)->labelsAngleChanged(angle);
}

void QAbstractAxis_ConnectLabelsBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QBrush &)>(&QAbstractAxis::labelsBrushChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QBrush &)>(&MyQAbstractAxis::Signal_LabelsBrushChanged));
}

void QAbstractAxis_DisconnectLabelsBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QBrush &)>(&QAbstractAxis::labelsBrushChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QBrush &)>(&MyQAbstractAxis::Signal_LabelsBrushChanged));
}

void QAbstractAxis_LabelsBrushChanged(void* ptr, void* brush)
{
	static_cast<QAbstractAxis*>(ptr)->labelsBrushChanged(*static_cast<QBrush*>(brush));
}

void QAbstractAxis_ConnectLabelsColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(QColor)>(&QAbstractAxis::labelsColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(QColor)>(&MyQAbstractAxis::Signal_LabelsColorChanged));
}

void QAbstractAxis_DisconnectLabelsColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(QColor)>(&QAbstractAxis::labelsColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(QColor)>(&MyQAbstractAxis::Signal_LabelsColorChanged));
}

void QAbstractAxis_LabelsColorChanged(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->labelsColorChanged(*static_cast<QColor*>(color));
}

void QAbstractAxis_ConnectLabelsFontChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QFont &)>(&QAbstractAxis::labelsFontChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QFont &)>(&MyQAbstractAxis::Signal_LabelsFontChanged));
}

void QAbstractAxis_DisconnectLabelsFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QFont &)>(&QAbstractAxis::labelsFontChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QFont &)>(&MyQAbstractAxis::Signal_LabelsFontChanged));
}

void QAbstractAxis_LabelsFontChanged(void* ptr, void* font)
{
	static_cast<QAbstractAxis*>(ptr)->labelsFontChanged(*static_cast<QFont*>(font));
}

void QAbstractAxis_ConnectLabelsVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::labelsVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_LabelsVisibleChanged));
}

void QAbstractAxis_DisconnectLabelsVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::labelsVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_LabelsVisibleChanged));
}

void QAbstractAxis_LabelsVisibleChanged(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->labelsVisibleChanged(visible != 0);
}

void QAbstractAxis_ConnectLinePenChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QPen &)>(&QAbstractAxis::linePenChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QPen &)>(&MyQAbstractAxis::Signal_LinePenChanged));
}

void QAbstractAxis_DisconnectLinePenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QPen &)>(&QAbstractAxis::linePenChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QPen &)>(&MyQAbstractAxis::Signal_LinePenChanged));
}

void QAbstractAxis_LinePenChanged(void* ptr, void* pen)
{
	static_cast<QAbstractAxis*>(ptr)->linePenChanged(*static_cast<QPen*>(pen));
}

void QAbstractAxis_ConnectLineVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::lineVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_LineVisibleChanged));
}

void QAbstractAxis_DisconnectLineVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::lineVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_LineVisibleChanged));
}

void QAbstractAxis_LineVisibleChanged(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->lineVisibleChanged(visible != 0);
}

void QAbstractAxis_ConnectMinorGridLineColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QColor &)>(&QAbstractAxis::minorGridLineColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QColor &)>(&MyQAbstractAxis::Signal_MinorGridLineColorChanged));
}

void QAbstractAxis_DisconnectMinorGridLineColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QColor &)>(&QAbstractAxis::minorGridLineColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QColor &)>(&MyQAbstractAxis::Signal_MinorGridLineColorChanged));
}

void QAbstractAxis_MinorGridLineColorChanged(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->minorGridLineColorChanged(*static_cast<QColor*>(color));
}

void QAbstractAxis_ConnectMinorGridLinePenChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QPen &)>(&QAbstractAxis::minorGridLinePenChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QPen &)>(&MyQAbstractAxis::Signal_MinorGridLinePenChanged));
}

void QAbstractAxis_DisconnectMinorGridLinePenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QPen &)>(&QAbstractAxis::minorGridLinePenChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QPen &)>(&MyQAbstractAxis::Signal_MinorGridLinePenChanged));
}

void QAbstractAxis_MinorGridLinePenChanged(void* ptr, void* pen)
{
	static_cast<QAbstractAxis*>(ptr)->minorGridLinePenChanged(*static_cast<QPen*>(pen));
}

void QAbstractAxis_ConnectMinorGridVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::minorGridVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_MinorGridVisibleChanged));
}

void QAbstractAxis_DisconnectMinorGridVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::minorGridVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_MinorGridVisibleChanged));
}

void QAbstractAxis_MinorGridVisibleChanged(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->minorGridVisibleChanged(visible != 0);
}

void QAbstractAxis_ConnectReverseChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::reverseChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_ReverseChanged));
}

void QAbstractAxis_DisconnectReverseChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::reverseChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_ReverseChanged));
}

void QAbstractAxis_ReverseChanged(void* ptr, char reverse)
{
	static_cast<QAbstractAxis*>(ptr)->reverseChanged(reverse != 0);
}

void QAbstractAxis_SetGridLineColor(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->setGridLineColor(*static_cast<QColor*>(color));
}

void QAbstractAxis_SetGridLinePen(void* ptr, void* pen)
{
	static_cast<QAbstractAxis*>(ptr)->setGridLinePen(*static_cast<QPen*>(pen));
}

void QAbstractAxis_SetGridLineVisible(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->setGridLineVisible(visible != 0);
}

void QAbstractAxis_SetLabelsAngle(void* ptr, int angle)
{
	static_cast<QAbstractAxis*>(ptr)->setLabelsAngle(angle);
}

void QAbstractAxis_SetLabelsBrush(void* ptr, void* brush)
{
	static_cast<QAbstractAxis*>(ptr)->setLabelsBrush(*static_cast<QBrush*>(brush));
}

void QAbstractAxis_SetLabelsColor(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->setLabelsColor(*static_cast<QColor*>(color));
}

void QAbstractAxis_SetLabelsFont(void* ptr, void* font)
{
	static_cast<QAbstractAxis*>(ptr)->setLabelsFont(*static_cast<QFont*>(font));
}

void QAbstractAxis_SetLabelsVisible(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->setLabelsVisible(visible != 0);
}

void QAbstractAxis_SetLinePen(void* ptr, void* pen)
{
	static_cast<QAbstractAxis*>(ptr)->setLinePen(*static_cast<QPen*>(pen));
}

void QAbstractAxis_SetLinePenColor(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->setLinePenColor(*static_cast<QColor*>(color));
}

void QAbstractAxis_SetLineVisible(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->setLineVisible(visible != 0);
}

void QAbstractAxis_SetMax(void* ptr, void* max)
{
	static_cast<QAbstractAxis*>(ptr)->setMax(*static_cast<QVariant*>(max));
}

void QAbstractAxis_SetMin(void* ptr, void* min)
{
	static_cast<QAbstractAxis*>(ptr)->setMin(*static_cast<QVariant*>(min));
}

void QAbstractAxis_SetMinorGridLineColor(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->setMinorGridLineColor(*static_cast<QColor*>(color));
}

void QAbstractAxis_SetMinorGridLinePen(void* ptr, void* pen)
{
	static_cast<QAbstractAxis*>(ptr)->setMinorGridLinePen(*static_cast<QPen*>(pen));
}

void QAbstractAxis_SetMinorGridLineVisible(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->setMinorGridLineVisible(visible != 0);
}

void QAbstractAxis_SetRange(void* ptr, void* min, void* max)
{
	static_cast<QAbstractAxis*>(ptr)->setRange(*static_cast<QVariant*>(min), *static_cast<QVariant*>(max));
}

void QAbstractAxis_SetReverse(void* ptr, char reverse)
{
	static_cast<QAbstractAxis*>(ptr)->setReverse(reverse != 0);
}

void QAbstractAxis_SetShadesBorderColor(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->setShadesBorderColor(*static_cast<QColor*>(color));
}

void QAbstractAxis_SetShadesBrush(void* ptr, void* brush)
{
	static_cast<QAbstractAxis*>(ptr)->setShadesBrush(*static_cast<QBrush*>(brush));
}

void QAbstractAxis_SetShadesColor(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->setShadesColor(*static_cast<QColor*>(color));
}

void QAbstractAxis_SetShadesPen(void* ptr, void* pen)
{
	static_cast<QAbstractAxis*>(ptr)->setShadesPen(*static_cast<QPen*>(pen));
}

void QAbstractAxis_SetShadesVisible(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->setShadesVisible(visible != 0);
}

void QAbstractAxis_SetTitleBrush(void* ptr, void* brush)
{
	static_cast<QAbstractAxis*>(ptr)->setTitleBrush(*static_cast<QBrush*>(brush));
}

void QAbstractAxis_SetTitleFont(void* ptr, void* font)
{
	static_cast<QAbstractAxis*>(ptr)->setTitleFont(*static_cast<QFont*>(font));
}

void QAbstractAxis_SetTitleText(void* ptr, struct QtCharts_PackedString title)
{
	static_cast<QAbstractAxis*>(ptr)->setTitleText(QString::fromUtf8(title.data, title.len));
}

void QAbstractAxis_SetTitleVisible(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->setTitleVisible(visible != 0);
}

void QAbstractAxis_SetVisible(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->setVisible(visible != 0);
}

void QAbstractAxis_ConnectShadesBorderColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(QColor)>(&QAbstractAxis::shadesBorderColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(QColor)>(&MyQAbstractAxis::Signal_ShadesBorderColorChanged));
}

void QAbstractAxis_DisconnectShadesBorderColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(QColor)>(&QAbstractAxis::shadesBorderColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(QColor)>(&MyQAbstractAxis::Signal_ShadesBorderColorChanged));
}

void QAbstractAxis_ShadesBorderColorChanged(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->shadesBorderColorChanged(*static_cast<QColor*>(color));
}

void QAbstractAxis_ConnectShadesBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QBrush &)>(&QAbstractAxis::shadesBrushChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QBrush &)>(&MyQAbstractAxis::Signal_ShadesBrushChanged));
}

void QAbstractAxis_DisconnectShadesBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QBrush &)>(&QAbstractAxis::shadesBrushChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QBrush &)>(&MyQAbstractAxis::Signal_ShadesBrushChanged));
}

void QAbstractAxis_ShadesBrushChanged(void* ptr, void* brush)
{
	static_cast<QAbstractAxis*>(ptr)->shadesBrushChanged(*static_cast<QBrush*>(brush));
}

void QAbstractAxis_ConnectShadesColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(QColor)>(&QAbstractAxis::shadesColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(QColor)>(&MyQAbstractAxis::Signal_ShadesColorChanged));
}

void QAbstractAxis_DisconnectShadesColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(QColor)>(&QAbstractAxis::shadesColorChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(QColor)>(&MyQAbstractAxis::Signal_ShadesColorChanged));
}

void QAbstractAxis_ShadesColorChanged(void* ptr, void* color)
{
	static_cast<QAbstractAxis*>(ptr)->shadesColorChanged(*static_cast<QColor*>(color));
}

void QAbstractAxis_ConnectShadesPenChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QPen &)>(&QAbstractAxis::shadesPenChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QPen &)>(&MyQAbstractAxis::Signal_ShadesPenChanged));
}

void QAbstractAxis_DisconnectShadesPenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QPen &)>(&QAbstractAxis::shadesPenChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QPen &)>(&MyQAbstractAxis::Signal_ShadesPenChanged));
}

void QAbstractAxis_ShadesPenChanged(void* ptr, void* pen)
{
	static_cast<QAbstractAxis*>(ptr)->shadesPenChanged(*static_cast<QPen*>(pen));
}

void QAbstractAxis_ConnectShadesVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::shadesVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_ShadesVisibleChanged));
}

void QAbstractAxis_DisconnectShadesVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::shadesVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_ShadesVisibleChanged));
}

void QAbstractAxis_ShadesVisibleChanged(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->shadesVisibleChanged(visible != 0);
}

void QAbstractAxis_Show(void* ptr)
{
	static_cast<QAbstractAxis*>(ptr)->show();
}

void QAbstractAxis_ConnectTitleBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QBrush &)>(&QAbstractAxis::titleBrushChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QBrush &)>(&MyQAbstractAxis::Signal_TitleBrushChanged));
}

void QAbstractAxis_DisconnectTitleBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QBrush &)>(&QAbstractAxis::titleBrushChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QBrush &)>(&MyQAbstractAxis::Signal_TitleBrushChanged));
}

void QAbstractAxis_TitleBrushChanged(void* ptr, void* brush)
{
	static_cast<QAbstractAxis*>(ptr)->titleBrushChanged(*static_cast<QBrush*>(brush));
}

void QAbstractAxis_ConnectTitleFontChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QFont &)>(&QAbstractAxis::titleFontChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QFont &)>(&MyQAbstractAxis::Signal_TitleFontChanged));
}

void QAbstractAxis_DisconnectTitleFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QFont &)>(&QAbstractAxis::titleFontChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QFont &)>(&MyQAbstractAxis::Signal_TitleFontChanged));
}

void QAbstractAxis_TitleFontChanged(void* ptr, void* font)
{
	static_cast<QAbstractAxis*>(ptr)->titleFontChanged(*static_cast<QFont*>(font));
}

void QAbstractAxis_ConnectTitleTextChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QString &)>(&QAbstractAxis::titleTextChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QString &)>(&MyQAbstractAxis::Signal_TitleTextChanged));
}

void QAbstractAxis_DisconnectTitleTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(const QString &)>(&QAbstractAxis::titleTextChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(const QString &)>(&MyQAbstractAxis::Signal_TitleTextChanged));
}

void QAbstractAxis_TitleTextChanged(void* ptr, struct QtCharts_PackedString text)
{
	static_cast<QAbstractAxis*>(ptr)->titleTextChanged(QString::fromUtf8(text.data, text.len));
}

void QAbstractAxis_ConnectTitleVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::titleVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_TitleVisibleChanged));
}

void QAbstractAxis_DisconnectTitleVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::titleVisibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_TitleVisibleChanged));
}

void QAbstractAxis_TitleVisibleChanged(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->titleVisibleChanged(visible != 0);
}

void QAbstractAxis_ConnectVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::visibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_VisibleChanged));
}

void QAbstractAxis_DisconnectVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractAxis*>(ptr), static_cast<void (QAbstractAxis::*)(bool)>(&QAbstractAxis::visibleChanged), static_cast<MyQAbstractAxis*>(ptr), static_cast<void (MyQAbstractAxis::*)(bool)>(&MyQAbstractAxis::Signal_VisibleChanged));
}

void QAbstractAxis_VisibleChanged(void* ptr, char visible)
{
	static_cast<QAbstractAxis*>(ptr)->visibleChanged(visible != 0);
}

void QAbstractAxis_DestroyQAbstractAxis(void* ptr)
{
	static_cast<QAbstractAxis*>(ptr)->~QAbstractAxis();
}

void QAbstractAxis_DestroyQAbstractAxisDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QAbstractAxis_Type(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->type();
}

void* QAbstractAxis_LabelsBrush(void* ptr)
{
	return new QBrush(static_cast<QAbstractAxis*>(ptr)->labelsBrush());
}

void* QAbstractAxis_ShadesBrush(void* ptr)
{
	return new QBrush(static_cast<QAbstractAxis*>(ptr)->shadesBrush());
}

void* QAbstractAxis_TitleBrush(void* ptr)
{
	return new QBrush(static_cast<QAbstractAxis*>(ptr)->titleBrush());
}

void* QAbstractAxis_LabelsColor(void* ptr)
{
	return new QColor(static_cast<QAbstractAxis*>(ptr)->labelsColor());
}

void* QAbstractAxis_LinePenColor(void* ptr)
{
	return new QColor(static_cast<QAbstractAxis*>(ptr)->linePenColor());
}

void* QAbstractAxis_ShadesBorderColor(void* ptr)
{
	return new QColor(static_cast<QAbstractAxis*>(ptr)->shadesBorderColor());
}

void* QAbstractAxis_ShadesColor(void* ptr)
{
	return new QColor(static_cast<QAbstractAxis*>(ptr)->shadesColor());
}

void* QAbstractAxis_LabelsFont(void* ptr)
{
	return new QFont(static_cast<QAbstractAxis*>(ptr)->labelsFont());
}

void* QAbstractAxis_TitleFont(void* ptr)
{
	return new QFont(static_cast<QAbstractAxis*>(ptr)->titleFont());
}

void* QAbstractAxis_GridLinePen(void* ptr)
{
	return new QPen(static_cast<QAbstractAxis*>(ptr)->gridLinePen());
}

void* QAbstractAxis_LinePen(void* ptr)
{
	return new QPen(static_cast<QAbstractAxis*>(ptr)->linePen());
}

void* QAbstractAxis_MinorGridLinePen(void* ptr)
{
	return new QPen(static_cast<QAbstractAxis*>(ptr)->minorGridLinePen());
}

void* QAbstractAxis_ShadesPen(void* ptr)
{
	return new QPen(static_cast<QAbstractAxis*>(ptr)->shadesPen());
}

struct QtCharts_PackedString QAbstractAxis_TitleText(void* ptr)
{
	return ({ QByteArray t6dcec2 = static_cast<QAbstractAxis*>(ptr)->titleText().toUtf8(); QtCharts_PackedString { const_cast<char*>(t6dcec2.prepend("WHITESPACE").constData()+10), t6dcec2.size()-10 }; });
}

long long QAbstractAxis_Alignment(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->alignment();
}

long long QAbstractAxis_Orientation(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->orientation();
}

char QAbstractAxis_IsGridLineVisible(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->isGridLineVisible();
}

char QAbstractAxis_IsLineVisible(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->isLineVisible();
}

char QAbstractAxis_IsMinorGridLineVisible(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->isMinorGridLineVisible();
}

char QAbstractAxis_IsReverse(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->isReverse();
}

char QAbstractAxis_IsTitleVisible(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->isTitleVisible();
}

char QAbstractAxis_IsVisible(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->isVisible();
}

char QAbstractAxis_LabelsVisible(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->labelsVisible();
}

char QAbstractAxis_ShadesVisible(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->shadesVisible();
}

void* QAbstractAxis_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::metaObject());
	} else if (dynamic_cast<QValueAxis*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QValueAxis*>(ptr)->QValueAxis::metaObject());
	} else if (dynamic_cast<QLogValueAxis*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::metaObject());
	} else if (dynamic_cast<QDateTimeAxis*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::metaObject());
	} else if (dynamic_cast<QBarCategoryAxis*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QAbstractAxis*>(ptr)->QAbstractAxis::metaObject());
	}
}

int QAbstractAxis_LabelsAngle(void* ptr)
{
	return static_cast<QAbstractAxis*>(ptr)->labelsAngle();
}

void* QAbstractAxis___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractAxis___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAbstractAxis___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAbstractAxis___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractAxis___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractAxis___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractAxis___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractAxis___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractAxis___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractAxis___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractAxis___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractAxis___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractAxis___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractAxis___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractAxis___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QAbstractAxis_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QValueAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QValueAxis*>(ptr)->QValueAxis::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QLogValueAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QDateTimeAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QBarCategoryAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QAbstractAxis*>(ptr)->QAbstractAxis::event(static_cast<QEvent*>(e));
	}
}

char QAbstractAxis_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QValueAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QValueAxis*>(ptr)->QValueAxis::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLogValueAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDateTimeAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBarCategoryAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QAbstractAxis*>(ptr)->QAbstractAxis::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QAbstractAxis_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QValueAxis*>(ptr)->QValueAxis::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QLogValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QDateTimeAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QBarCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QAbstractAxis*>(ptr)->QAbstractAxis::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QAbstractAxis_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QValueAxis*>(ptr)->QValueAxis::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLogValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDateTimeAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBarCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractAxis*>(ptr)->QAbstractAxis::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractAxis_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QValueAxis*>(ptr)->QValueAxis::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLogValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDateTimeAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBarCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QAbstractAxis*>(ptr)->QAbstractAxis::customEvent(static_cast<QEvent*>(event));
	}
}

void QAbstractAxis_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::deleteLater();
	} else if (dynamic_cast<QValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QValueAxis*>(ptr)->QValueAxis::deleteLater();
	} else if (dynamic_cast<QLogValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::deleteLater();
	} else if (dynamic_cast<QDateTimeAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::deleteLater();
	} else if (dynamic_cast<QBarCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::deleteLater();
	} else {
		static_cast<QAbstractAxis*>(ptr)->QAbstractAxis::deleteLater();
	}
}

void QAbstractAxis_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QValueAxis*>(ptr)->QValueAxis::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLogValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDateTimeAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBarCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractAxis*>(ptr)->QAbstractAxis::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractAxis_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QValueAxis*>(ptr)->QValueAxis::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QLogValueAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QDateTimeAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QBarCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QAbstractAxis*>(ptr)->QAbstractAxis::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

class MyQAbstractBarSeries: public QAbstractBarSeries
{
public:
	void Signal_BarsetsAdded(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsAdded(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BarsetsRemoved(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsRemoved(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Clicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_Clicked(this, index, barset); };
	void Signal_CountChanged() { callbackQAbstractBarSeries_CountChanged(this); };
	void Signal_DoubleClicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_DoubleClicked(this, index, barset); };
	void Signal_Hovered(bool status, int index, QBarSet * barset) { callbackQAbstractBarSeries_Hovered(this, status, index, barset); };
	void Signal_LabelsAngleChanged(qreal angle) { callbackQAbstractBarSeries_LabelsAngleChanged(this, angle); };
	void Signal_LabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQAbstractBarSeries_LabelsFormatChanged(this, formatPacked); };
	void Signal_LabelsPositionChanged(QAbstractBarSeries::LabelsPosition position) { callbackQAbstractBarSeries_LabelsPositionChanged(this, position); };
	void Signal_LabelsPrecisionChanged(int precision) { callbackQAbstractBarSeries_LabelsPrecisionChanged(this, precision); };
	void Signal_LabelsVisibleChanged() { callbackQAbstractBarSeries_LabelsVisibleChanged(this); };
	void Signal_Pressed(int index, QBarSet * barset) { callbackQAbstractBarSeries_Pressed(this, index, barset); };
	void Signal_Released(int index, QBarSet * barset) { callbackQAbstractBarSeries_Released(this, index, barset); };
	 ~MyQAbstractBarSeries() { callbackQAbstractBarSeries_DestroyQAbstractBarSeries(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQAbstractBarSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAbstractBarSeries*)

int QAbstractBarSeries_QAbstractBarSeries_QRegisterMetaType(){qRegisterMetaType<QAbstractBarSeries*>(); return qRegisterMetaType<MyQAbstractBarSeries*>();}

char QAbstractBarSeries_Append(void* ptr, void* set)
{
	return static_cast<QAbstractBarSeries*>(ptr)->append(static_cast<QBarSet*>(set));
}

char QAbstractBarSeries_Append2(void* ptr, void* sets)
{
	return static_cast<QAbstractBarSeries*>(ptr)->append(({ QList<QBarSet *>* tmpP = static_cast<QList<QBarSet *>*>(sets); QList<QBarSet *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

char QAbstractBarSeries_Insert(void* ptr, int index, void* set)
{
	return static_cast<QAbstractBarSeries*>(ptr)->insert(index, static_cast<QBarSet*>(set));
}

char QAbstractBarSeries_Remove(void* ptr, void* set)
{
	return static_cast<QAbstractBarSeries*>(ptr)->remove(static_cast<QBarSet*>(set));
}

char QAbstractBarSeries_Take(void* ptr, void* set)
{
	return static_cast<QAbstractBarSeries*>(ptr)->take(static_cast<QBarSet*>(set));
}

void QAbstractBarSeries_ConnectBarsetsAdded(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(QList<QBarSet *>)>(&QAbstractBarSeries::barsetsAdded), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(QList<QBarSet *>)>(&MyQAbstractBarSeries::Signal_BarsetsAdded));
}

void QAbstractBarSeries_DisconnectBarsetsAdded(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(QList<QBarSet *>)>(&QAbstractBarSeries::barsetsAdded), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(QList<QBarSet *>)>(&MyQAbstractBarSeries::Signal_BarsetsAdded));
}

void QAbstractBarSeries_BarsetsAdded(void* ptr, void* sets)
{
	static_cast<QAbstractBarSeries*>(ptr)->barsetsAdded(({ QList<QBarSet *>* tmpP = static_cast<QList<QBarSet *>*>(sets); QList<QBarSet *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QAbstractBarSeries_ConnectBarsetsRemoved(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(QList<QBarSet *>)>(&QAbstractBarSeries::barsetsRemoved), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(QList<QBarSet *>)>(&MyQAbstractBarSeries::Signal_BarsetsRemoved));
}

void QAbstractBarSeries_DisconnectBarsetsRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(QList<QBarSet *>)>(&QAbstractBarSeries::barsetsRemoved), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(QList<QBarSet *>)>(&MyQAbstractBarSeries::Signal_BarsetsRemoved));
}

void QAbstractBarSeries_BarsetsRemoved(void* ptr, void* sets)
{
	static_cast<QAbstractBarSeries*>(ptr)->barsetsRemoved(({ QList<QBarSet *>* tmpP = static_cast<QList<QBarSet *>*>(sets); QList<QBarSet *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QAbstractBarSeries_Clear(void* ptr)
{
	static_cast<QAbstractBarSeries*>(ptr)->clear();
}

void QAbstractBarSeries_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int, QBarSet *)>(&QAbstractBarSeries::clicked), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int, QBarSet *)>(&MyQAbstractBarSeries::Signal_Clicked));
}

void QAbstractBarSeries_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int, QBarSet *)>(&QAbstractBarSeries::clicked), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int, QBarSet *)>(&MyQAbstractBarSeries::Signal_Clicked));
}

void QAbstractBarSeries_Clicked(void* ptr, int index, void* barset)
{
	static_cast<QAbstractBarSeries*>(ptr)->clicked(index, static_cast<QBarSet*>(barset));
}

void QAbstractBarSeries_ConnectCountChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)()>(&QAbstractBarSeries::countChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)()>(&MyQAbstractBarSeries::Signal_CountChanged));
}

void QAbstractBarSeries_DisconnectCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)()>(&QAbstractBarSeries::countChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)()>(&MyQAbstractBarSeries::Signal_CountChanged));
}

void QAbstractBarSeries_CountChanged(void* ptr)
{
	static_cast<QAbstractBarSeries*>(ptr)->countChanged();
}

void QAbstractBarSeries_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int, QBarSet *)>(&QAbstractBarSeries::doubleClicked), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int, QBarSet *)>(&MyQAbstractBarSeries::Signal_DoubleClicked));
}

void QAbstractBarSeries_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int, QBarSet *)>(&QAbstractBarSeries::doubleClicked), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int, QBarSet *)>(&MyQAbstractBarSeries::Signal_DoubleClicked));
}

void QAbstractBarSeries_DoubleClicked(void* ptr, int index, void* barset)
{
	static_cast<QAbstractBarSeries*>(ptr)->doubleClicked(index, static_cast<QBarSet*>(barset));
}

void QAbstractBarSeries_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(bool, int, QBarSet *)>(&QAbstractBarSeries::hovered), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(bool, int, QBarSet *)>(&MyQAbstractBarSeries::Signal_Hovered));
}

void QAbstractBarSeries_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(bool, int, QBarSet *)>(&QAbstractBarSeries::hovered), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(bool, int, QBarSet *)>(&MyQAbstractBarSeries::Signal_Hovered));
}

void QAbstractBarSeries_Hovered(void* ptr, char status, int index, void* barset)
{
	static_cast<QAbstractBarSeries*>(ptr)->hovered(status != 0, index, static_cast<QBarSet*>(barset));
}

void QAbstractBarSeries_ConnectLabelsAngleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(qreal)>(&QAbstractBarSeries::labelsAngleChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(qreal)>(&MyQAbstractBarSeries::Signal_LabelsAngleChanged));
}

void QAbstractBarSeries_DisconnectLabelsAngleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(qreal)>(&QAbstractBarSeries::labelsAngleChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(qreal)>(&MyQAbstractBarSeries::Signal_LabelsAngleChanged));
}

void QAbstractBarSeries_LabelsAngleChanged(void* ptr, double angle)
{
	static_cast<QAbstractBarSeries*>(ptr)->labelsAngleChanged(angle);
}

void QAbstractBarSeries_ConnectLabelsFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(const QString &)>(&QAbstractBarSeries::labelsFormatChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(const QString &)>(&MyQAbstractBarSeries::Signal_LabelsFormatChanged));
}

void QAbstractBarSeries_DisconnectLabelsFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(const QString &)>(&QAbstractBarSeries::labelsFormatChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(const QString &)>(&MyQAbstractBarSeries::Signal_LabelsFormatChanged));
}

void QAbstractBarSeries_LabelsFormatChanged(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QAbstractBarSeries*>(ptr)->labelsFormatChanged(QString::fromUtf8(format.data, format.len));
}

void QAbstractBarSeries_ConnectLabelsPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(QAbstractBarSeries::LabelsPosition)>(&QAbstractBarSeries::labelsPositionChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(QAbstractBarSeries::LabelsPosition)>(&MyQAbstractBarSeries::Signal_LabelsPositionChanged));
}

void QAbstractBarSeries_DisconnectLabelsPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(QAbstractBarSeries::LabelsPosition)>(&QAbstractBarSeries::labelsPositionChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(QAbstractBarSeries::LabelsPosition)>(&MyQAbstractBarSeries::Signal_LabelsPositionChanged));
}

void QAbstractBarSeries_LabelsPositionChanged(void* ptr, long long position)
{
	static_cast<QAbstractBarSeries*>(ptr)->labelsPositionChanged(static_cast<QAbstractBarSeries::LabelsPosition>(position));
}

void QAbstractBarSeries_ConnectLabelsPrecisionChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int)>(&QAbstractBarSeries::labelsPrecisionChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int)>(&MyQAbstractBarSeries::Signal_LabelsPrecisionChanged));
}

void QAbstractBarSeries_DisconnectLabelsPrecisionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int)>(&QAbstractBarSeries::labelsPrecisionChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int)>(&MyQAbstractBarSeries::Signal_LabelsPrecisionChanged));
}

void QAbstractBarSeries_LabelsPrecisionChanged(void* ptr, int precision)
{
	static_cast<QAbstractBarSeries*>(ptr)->labelsPrecisionChanged(precision);
}

void QAbstractBarSeries_ConnectLabelsVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)()>(&QAbstractBarSeries::labelsVisibleChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)()>(&MyQAbstractBarSeries::Signal_LabelsVisibleChanged));
}

void QAbstractBarSeries_DisconnectLabelsVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)()>(&QAbstractBarSeries::labelsVisibleChanged), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)()>(&MyQAbstractBarSeries::Signal_LabelsVisibleChanged));
}

void QAbstractBarSeries_LabelsVisibleChanged(void* ptr)
{
	static_cast<QAbstractBarSeries*>(ptr)->labelsVisibleChanged();
}

void QAbstractBarSeries_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int, QBarSet *)>(&QAbstractBarSeries::pressed), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int, QBarSet *)>(&MyQAbstractBarSeries::Signal_Pressed));
}

void QAbstractBarSeries_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int, QBarSet *)>(&QAbstractBarSeries::pressed), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int, QBarSet *)>(&MyQAbstractBarSeries::Signal_Pressed));
}

void QAbstractBarSeries_Pressed(void* ptr, int index, void* barset)
{
	static_cast<QAbstractBarSeries*>(ptr)->pressed(index, static_cast<QBarSet*>(barset));
}

void QAbstractBarSeries_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int, QBarSet *)>(&QAbstractBarSeries::released), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int, QBarSet *)>(&MyQAbstractBarSeries::Signal_Released));
}

void QAbstractBarSeries_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractBarSeries*>(ptr), static_cast<void (QAbstractBarSeries::*)(int, QBarSet *)>(&QAbstractBarSeries::released), static_cast<MyQAbstractBarSeries*>(ptr), static_cast<void (MyQAbstractBarSeries::*)(int, QBarSet *)>(&MyQAbstractBarSeries::Signal_Released));
}

void QAbstractBarSeries_Released(void* ptr, int index, void* barset)
{
	static_cast<QAbstractBarSeries*>(ptr)->released(index, static_cast<QBarSet*>(barset));
}

void QAbstractBarSeries_SetBarWidth(void* ptr, double width)
{
	static_cast<QAbstractBarSeries*>(ptr)->setBarWidth(width);
}

void QAbstractBarSeries_SetLabelsAngle(void* ptr, double angle)
{
	static_cast<QAbstractBarSeries*>(ptr)->setLabelsAngle(angle);
}

void QAbstractBarSeries_SetLabelsFormat(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QAbstractBarSeries*>(ptr)->setLabelsFormat(QString::fromUtf8(format.data, format.len));
}

void QAbstractBarSeries_SetLabelsPosition(void* ptr, long long position)
{
	static_cast<QAbstractBarSeries*>(ptr)->setLabelsPosition(static_cast<QAbstractBarSeries::LabelsPosition>(position));
}

void QAbstractBarSeries_SetLabelsPrecision(void* ptr, int precision)
{
	static_cast<QAbstractBarSeries*>(ptr)->setLabelsPrecision(precision);
}

void QAbstractBarSeries_SetLabelsVisible(void* ptr, char visible)
{
	static_cast<QAbstractBarSeries*>(ptr)->setLabelsVisible(visible != 0);
}

void QAbstractBarSeries_DestroyQAbstractBarSeries(void* ptr)
{
	static_cast<QAbstractBarSeries*>(ptr)->~QAbstractBarSeries();
}

void QAbstractBarSeries_DestroyQAbstractBarSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QAbstractBarSeries_LabelsPosition(void* ptr)
{
	return static_cast<QAbstractBarSeries*>(ptr)->labelsPosition();
}

struct QtCharts_PackedList QAbstractBarSeries_BarSets(void* ptr)
{
	return ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(static_cast<QAbstractBarSeries*>(ptr)->barSets()); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtCharts_PackedString QAbstractBarSeries_LabelsFormat(void* ptr)
{
	return ({ QByteArray tefa1eb = static_cast<QAbstractBarSeries*>(ptr)->labelsFormat().toUtf8(); QtCharts_PackedString { const_cast<char*>(tefa1eb.prepend("WHITESPACE").constData()+10), tefa1eb.size()-10 }; });
}

char QAbstractBarSeries_IsLabelsVisible(void* ptr)
{
	return static_cast<QAbstractBarSeries*>(ptr)->isLabelsVisible();
}

int QAbstractBarSeries_Count(void* ptr)
{
	return static_cast<QAbstractBarSeries*>(ptr)->count();
}

int QAbstractBarSeries_LabelsPrecision(void* ptr)
{
	return static_cast<QAbstractBarSeries*>(ptr)->labelsPrecision();
}

double QAbstractBarSeries_BarWidth(void* ptr)
{
	return static_cast<QAbstractBarSeries*>(ptr)->barWidth();
}

double QAbstractBarSeries_LabelsAngle(void* ptr)
{
	return static_cast<QAbstractBarSeries*>(ptr)->labelsAngle();
}

void* QAbstractBarSeries___append_sets_atList2(void* ptr, int i)
{
	return ({QBarSet * tmp = static_cast<QList<QBarSet *>*>(ptr)->at(i); if (i == static_cast<QList<QBarSet *>*>(ptr)->size()-1) { static_cast<QList<QBarSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractBarSeries___append_sets_setList2(void* ptr, void* i)
{
	static_cast<QList<QBarSet *>*>(ptr)->append(static_cast<QBarSet*>(i));
}

void* QAbstractBarSeries___append_sets_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QBarSet *>();
}

void* QAbstractBarSeries___barsetsAdded_sets_atList(void* ptr, int i)
{
	return ({QBarSet * tmp = static_cast<QList<QBarSet *>*>(ptr)->at(i); if (i == static_cast<QList<QBarSet *>*>(ptr)->size()-1) { static_cast<QList<QBarSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractBarSeries___barsetsAdded_sets_setList(void* ptr, void* i)
{
	static_cast<QList<QBarSet *>*>(ptr)->append(static_cast<QBarSet*>(i));
}

void* QAbstractBarSeries___barsetsAdded_sets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QBarSet *>();
}

void* QAbstractBarSeries___barsetsRemoved_sets_atList(void* ptr, int i)
{
	return ({QBarSet * tmp = static_cast<QList<QBarSet *>*>(ptr)->at(i); if (i == static_cast<QList<QBarSet *>*>(ptr)->size()-1) { static_cast<QList<QBarSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractBarSeries___barsetsRemoved_sets_setList(void* ptr, void* i)
{
	static_cast<QList<QBarSet *>*>(ptr)->append(static_cast<QBarSet*>(i));
}

void* QAbstractBarSeries___barsetsRemoved_sets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QBarSet *>();
}

void* QAbstractBarSeries___barSets_atList(void* ptr, int i)
{
	return ({QBarSet * tmp = static_cast<QList<QBarSet *>*>(ptr)->at(i); if (i == static_cast<QList<QBarSet *>*>(ptr)->size()-1) { static_cast<QList<QBarSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractBarSeries___barSets_setList(void* ptr, void* i)
{
	static_cast<QList<QBarSet *>*>(ptr)->append(static_cast<QBarSet*>(i));
}

void* QAbstractBarSeries___barSets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QBarSet *>();
}

long long QAbstractBarSeries_Type(void* ptr)
{
	return static_cast<QAbstractBarSeries*>(ptr)->type();
}

long long QAbstractBarSeries_TypeDefault(void* ptr)
{
	if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::type();
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::type();
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::type();
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::type();
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::type();
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBarSeries*>(ptr)->QBarSeries::type();
	} else {
	
	}
}

class MyQAbstractSeries: public QAbstractSeries
{
public:
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	 ~MyQAbstractSeries() { callbackQAbstractSeries_DestroyQAbstractSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQAbstractSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAbstractSeries*)

int QAbstractSeries_QAbstractSeries_QRegisterMetaType(){qRegisterMetaType<QAbstractSeries*>(); return qRegisterMetaType<MyQAbstractSeries*>();}

struct QtCharts_PackedList QAbstractSeries_AttachedAxes(void* ptr)
{
	return ({ QList<QAbstractAxis *>* tmpValue = new QList<QAbstractAxis *>(static_cast<QAbstractSeries*>(ptr)->attachedAxes()); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtCharts_PackedString QAbstractSeries_QAbstractSeries_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t43d83a = QAbstractSeries::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t43d83a.prepend("WHITESPACE").constData()+10), t43d83a.size()-10 }; });
}

struct QtCharts_PackedString QAbstractSeries_QAbstractSeries_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tbc095b = QAbstractSeries::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(tbc095b.prepend("WHITESPACE").constData()+10), tbc095b.size()-10 }; });
}

char QAbstractSeries_AttachAxis(void* ptr, void* axis)
{
	return static_cast<QAbstractSeries*>(ptr)->attachAxis(static_cast<QAbstractAxis*>(axis));
}

char QAbstractSeries_DetachAxis(void* ptr, void* axis)
{
	return static_cast<QAbstractSeries*>(ptr)->detachAxis(static_cast<QAbstractAxis*>(axis));
}

void QAbstractSeries_Hide(void* ptr)
{
	static_cast<QAbstractSeries*>(ptr)->hide();
}

void QAbstractSeries_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractSeries*>(ptr), static_cast<void (QAbstractSeries::*)()>(&QAbstractSeries::nameChanged), static_cast<MyQAbstractSeries*>(ptr), static_cast<void (MyQAbstractSeries::*)()>(&MyQAbstractSeries::Signal_NameChanged));
}

void QAbstractSeries_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSeries*>(ptr), static_cast<void (QAbstractSeries::*)()>(&QAbstractSeries::nameChanged), static_cast<MyQAbstractSeries*>(ptr), static_cast<void (MyQAbstractSeries::*)()>(&MyQAbstractSeries::Signal_NameChanged));
}

void QAbstractSeries_NameChanged(void* ptr)
{
	static_cast<QAbstractSeries*>(ptr)->nameChanged();
}

void QAbstractSeries_ConnectOpacityChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractSeries*>(ptr), static_cast<void (QAbstractSeries::*)()>(&QAbstractSeries::opacityChanged), static_cast<MyQAbstractSeries*>(ptr), static_cast<void (MyQAbstractSeries::*)()>(&MyQAbstractSeries::Signal_OpacityChanged));
}

void QAbstractSeries_DisconnectOpacityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSeries*>(ptr), static_cast<void (QAbstractSeries::*)()>(&QAbstractSeries::opacityChanged), static_cast<MyQAbstractSeries*>(ptr), static_cast<void (MyQAbstractSeries::*)()>(&MyQAbstractSeries::Signal_OpacityChanged));
}

void QAbstractSeries_OpacityChanged(void* ptr)
{
	static_cast<QAbstractSeries*>(ptr)->opacityChanged();
}

void QAbstractSeries_SetName(void* ptr, struct QtCharts_PackedString name)
{
	static_cast<QAbstractSeries*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QAbstractSeries_SetOpacity(void* ptr, double opacity)
{
	static_cast<QAbstractSeries*>(ptr)->setOpacity(opacity);
}

void QAbstractSeries_SetUseOpenGL(void* ptr, char enable)
{
	static_cast<QAbstractSeries*>(ptr)->setUseOpenGL(enable != 0);
}

void QAbstractSeries_SetVisible(void* ptr, char visible)
{
	static_cast<QAbstractSeries*>(ptr)->setVisible(visible != 0);
}

void QAbstractSeries_Show(void* ptr)
{
	static_cast<QAbstractSeries*>(ptr)->show();
}

void QAbstractSeries_ConnectUseOpenGLChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractSeries*>(ptr), static_cast<void (QAbstractSeries::*)()>(&QAbstractSeries::useOpenGLChanged), static_cast<MyQAbstractSeries*>(ptr), static_cast<void (MyQAbstractSeries::*)()>(&MyQAbstractSeries::Signal_UseOpenGLChanged));
}

void QAbstractSeries_DisconnectUseOpenGLChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSeries*>(ptr), static_cast<void (QAbstractSeries::*)()>(&QAbstractSeries::useOpenGLChanged), static_cast<MyQAbstractSeries*>(ptr), static_cast<void (MyQAbstractSeries::*)()>(&MyQAbstractSeries::Signal_UseOpenGLChanged));
}

void QAbstractSeries_UseOpenGLChanged(void* ptr)
{
	static_cast<QAbstractSeries*>(ptr)->useOpenGLChanged();
}

void QAbstractSeries_ConnectVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractSeries*>(ptr), static_cast<void (QAbstractSeries::*)()>(&QAbstractSeries::visibleChanged), static_cast<MyQAbstractSeries*>(ptr), static_cast<void (MyQAbstractSeries::*)()>(&MyQAbstractSeries::Signal_VisibleChanged));
}

void QAbstractSeries_DisconnectVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractSeries*>(ptr), static_cast<void (QAbstractSeries::*)()>(&QAbstractSeries::visibleChanged), static_cast<MyQAbstractSeries*>(ptr), static_cast<void (MyQAbstractSeries::*)()>(&MyQAbstractSeries::Signal_VisibleChanged));
}

void QAbstractSeries_VisibleChanged(void* ptr)
{
	static_cast<QAbstractSeries*>(ptr)->visibleChanged();
}

void QAbstractSeries_DestroyQAbstractSeries(void* ptr)
{
	static_cast<QAbstractSeries*>(ptr)->~QAbstractSeries();
}

void QAbstractSeries_DestroyQAbstractSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QAbstractSeries_Type(void* ptr)
{
	return static_cast<QAbstractSeries*>(ptr)->type();
}

void* QAbstractSeries_Chart(void* ptr)
{
	return static_cast<QAbstractSeries*>(ptr)->chart();
}

struct QtCharts_PackedString QAbstractSeries_Name(void* ptr)
{
	return ({ QByteArray tb9025a = static_cast<QAbstractSeries*>(ptr)->name().toUtf8(); QtCharts_PackedString { const_cast<char*>(tb9025a.prepend("WHITESPACE").constData()+10), tb9025a.size()-10 }; });
}

char QAbstractSeries_IsVisible(void* ptr)
{
	return static_cast<QAbstractSeries*>(ptr)->isVisible();
}

char QAbstractSeries_UseOpenGL(void* ptr)
{
	return static_cast<QAbstractSeries*>(ptr)->useOpenGL();
}

void* QAbstractSeries_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QScatterSeries*>(ptr)->QScatterSeries::metaObject());
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QSplineSeries*>(ptr)->QSplineSeries::metaObject());
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QLineSeries*>(ptr)->QLineSeries::metaObject());
	} else if (dynamic_cast<QXYSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QXYSeries*>(ptr)->QXYSeries::metaObject());
	} else if (dynamic_cast<QPieSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QPieSeries*>(ptr)->QPieSeries::metaObject());
	} else if (dynamic_cast<QCandlestickSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::metaObject());
	} else if (dynamic_cast<QBoxPlotSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::metaObject());
	} else if (dynamic_cast<QAreaSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAreaSeries*>(ptr)->QAreaSeries::metaObject());
	} else if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::metaObject());
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::metaObject());
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::metaObject());
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::metaObject());
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::metaObject());
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QBarSeries*>(ptr)->QBarSeries::metaObject());
	} else if (dynamic_cast<QAbstractBarSeries*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAbstractBarSeries*>(ptr)->QAbstractBarSeries::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QAbstractSeries*>(ptr)->QAbstractSeries::metaObject());
	}
}

double QAbstractSeries_Opacity(void* ptr)
{
	return static_cast<QAbstractSeries*>(ptr)->opacity();
}

void* QAbstractSeries___attachedAxes_atList(void* ptr, int i)
{
	return ({QAbstractAxis * tmp = static_cast<QList<QAbstractAxis *>*>(ptr)->at(i); if (i == static_cast<QList<QAbstractAxis *>*>(ptr)->size()-1) { static_cast<QList<QAbstractAxis *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractSeries___attachedAxes_setList(void* ptr, void* i)
{
	static_cast<QList<QAbstractAxis *>*>(ptr)->append(static_cast<QAbstractAxis*>(i));
}

void* QAbstractSeries___attachedAxes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAbstractAxis *>();
}

void* QAbstractSeries___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractSeries___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAbstractSeries___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAbstractSeries___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractSeries___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractSeries___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractSeries___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractSeries___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractSeries___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractSeries___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractSeries___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractSeries___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractSeries___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractSeries___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAbstractSeries___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QAbstractSeries_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScatterSeries*>(ptr)->QScatterSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplineSeries*>(ptr)->QSplineSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLineSeries*>(ptr)->QLineSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QXYSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QXYSeries*>(ptr)->QXYSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QPieSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPieSeries*>(ptr)->QPieSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QCandlestickSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QBoxPlotSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAreaSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAreaSeries*>(ptr)->QAreaSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBarSeries*>(ptr)->QBarSeries::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAbstractBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractBarSeries*>(ptr)->QAbstractBarSeries::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QAbstractSeries*>(ptr)->QAbstractSeries::event(static_cast<QEvent*>(e));
	}
}

char QAbstractSeries_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScatterSeries*>(ptr)->QScatterSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplineSeries*>(ptr)->QSplineSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLineSeries*>(ptr)->QLineSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QXYSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QXYSeries*>(ptr)->QXYSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPieSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPieSeries*>(ptr)->QPieSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCandlestickSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBoxPlotSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAreaSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAreaSeries*>(ptr)->QAreaSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBarSeries*>(ptr)->QBarSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractBarSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractBarSeries*>(ptr)->QAbstractBarSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QAbstractSeries*>(ptr)->QAbstractSeries::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QAbstractSeries_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QScatterSeries*>(ptr)->QScatterSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplineSeries*>(ptr)->QSplineSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineSeries*>(ptr)->QLineSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QXYSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYSeries*>(ptr)->QXYSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QPieSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieSeries*>(ptr)->QPieSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QCandlestickSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QBoxPlotSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAreaSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaSeries*>(ptr)->QAreaSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarSeries*>(ptr)->QBarSeries::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAbstractBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractBarSeries*>(ptr)->QAbstractBarSeries::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QAbstractSeries*>(ptr)->QAbstractSeries::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QAbstractSeries_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QScatterSeries*>(ptr)->QScatterSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplineSeries*>(ptr)->QSplineSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineSeries*>(ptr)->QLineSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QXYSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYSeries*>(ptr)->QXYSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPieSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieSeries*>(ptr)->QPieSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCandlestickSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBoxPlotSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAreaSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaSeries*>(ptr)->QAreaSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarSeries*>(ptr)->QBarSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractBarSeries*>(ptr)->QAbstractBarSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractSeries*>(ptr)->QAbstractSeries::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractSeries_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QScatterSeries*>(ptr)->QScatterSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplineSeries*>(ptr)->QSplineSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineSeries*>(ptr)->QLineSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QXYSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYSeries*>(ptr)->QXYSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPieSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieSeries*>(ptr)->QPieSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCandlestickSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBoxPlotSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAreaSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaSeries*>(ptr)->QAreaSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarSeries*>(ptr)->QBarSeries::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractBarSeries*>(ptr)->QAbstractBarSeries::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QAbstractSeries*>(ptr)->QAbstractSeries::customEvent(static_cast<QEvent*>(event));
	}
}

void QAbstractSeries_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QScatterSeries*>(ptr)->QScatterSeries::deleteLater();
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplineSeries*>(ptr)->QSplineSeries::deleteLater();
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineSeries*>(ptr)->QLineSeries::deleteLater();
	} else if (dynamic_cast<QXYSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYSeries*>(ptr)->QXYSeries::deleteLater();
	} else if (dynamic_cast<QPieSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieSeries*>(ptr)->QPieSeries::deleteLater();
	} else if (dynamic_cast<QCandlestickSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::deleteLater();
	} else if (dynamic_cast<QBoxPlotSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::deleteLater();
	} else if (dynamic_cast<QAreaSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaSeries*>(ptr)->QAreaSeries::deleteLater();
	} else if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::deleteLater();
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::deleteLater();
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::deleteLater();
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::deleteLater();
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::deleteLater();
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarSeries*>(ptr)->QBarSeries::deleteLater();
	} else if (dynamic_cast<QAbstractBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractBarSeries*>(ptr)->QAbstractBarSeries::deleteLater();
	} else {
		static_cast<QAbstractSeries*>(ptr)->QAbstractSeries::deleteLater();
	}
}

void QAbstractSeries_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QScatterSeries*>(ptr)->QScatterSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplineSeries*>(ptr)->QSplineSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineSeries*>(ptr)->QLineSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QXYSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYSeries*>(ptr)->QXYSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPieSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieSeries*>(ptr)->QPieSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCandlestickSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBoxPlotSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAreaSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaSeries*>(ptr)->QAreaSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarSeries*>(ptr)->QBarSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractBarSeries*>(ptr)->QAbstractBarSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QAbstractSeries*>(ptr)->QAbstractSeries::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QAbstractSeries_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QScatterSeries*>(ptr)->QScatterSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplineSeries*>(ptr)->QSplineSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineSeries*>(ptr)->QLineSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QXYSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYSeries*>(ptr)->QXYSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QPieSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieSeries*>(ptr)->QPieSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QCandlestickSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QBoxPlotSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAreaSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaSeries*>(ptr)->QAreaSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHorizontalStackedBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHorizontalPercentBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHorizontalBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarSeries*>(ptr)->QBarSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAbstractBarSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractBarSeries*>(ptr)->QAbstractBarSeries::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QAbstractSeries*>(ptr)->QAbstractSeries::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

class MyQAreaLegendMarker: public QAreaLegendMarker
{
public:
	QAreaSeries * series() { return static_cast<QAreaSeries*>(callbackQAreaLegendMarker_Series(this)); };
	QLegendMarker::LegendMarkerType type() { return static_cast<QLegendMarker::LegendMarkerType>(callbackQAreaLegendMarker_Type(this)); };
	 ~MyQAreaLegendMarker() { callbackQAreaLegendMarker_DestroyQAreaLegendMarker(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLegendMarker_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BrushChanged() { callbackQLegendMarker_BrushChanged(this); };
	void Signal_Clicked() { callbackQLegendMarker_Clicked(this); };
	void Signal_FontChanged() { callbackQLegendMarker_FontChanged(this); };
	void Signal_Hovered(bool status) { callbackQLegendMarker_Hovered(this, status); };
	void Signal_LabelBrushChanged() { callbackQLegendMarker_LabelBrushChanged(this); };
	void Signal_LabelChanged() { callbackQLegendMarker_LabelChanged(this); };
	void Signal_PenChanged() { callbackQLegendMarker_PenChanged(this); };
	void Signal_ShapeChanged() { callbackQLegendMarker_ShapeChanged(this); };
	void Signal_VisibleChanged() { callbackQLegendMarker_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQLegendMarker_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLegendMarker_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLegendMarker_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLegendMarker_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLegendMarker_CustomEvent(this, event); };
	void deleteLater() { callbackQLegendMarker_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLegendMarker_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLegendMarker_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLegendMarker_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLegendMarker_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAreaLegendMarker*)

int QAreaLegendMarker_QAreaLegendMarker_QRegisterMetaType(){qRegisterMetaType<QAreaLegendMarker*>(); return qRegisterMetaType<MyQAreaLegendMarker*>();}

void* QAreaLegendMarker_Series(void* ptr)
{
	return static_cast<QAreaLegendMarker*>(ptr)->series();
}

void* QAreaLegendMarker_SeriesDefault(void* ptr)
{
		return static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::series();
}

long long QAreaLegendMarker_Type(void* ptr)
{
	return static_cast<QAreaLegendMarker*>(ptr)->type();
}

long long QAreaLegendMarker_TypeDefault(void* ptr)
{
		return static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::type();
}

void QAreaLegendMarker_DestroyQAreaLegendMarker(void* ptr)
{
	static_cast<QAreaLegendMarker*>(ptr)->~QAreaLegendMarker();
}

void QAreaLegendMarker_DestroyQAreaLegendMarkerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQAreaSeries: public QAreaSeries
{
public:
	MyQAreaSeries(QLineSeries *upperSeries, QLineSeries *lowerSeries = Q_NULLPTR) : QAreaSeries(upperSeries, lowerSeries) {QAreaSeries_QAreaSeries_QRegisterMetaType();};
	MyQAreaSeries(QObject *parent = Q_NULLPTR) : QAreaSeries(parent) {QAreaSeries_QAreaSeries_QRegisterMetaType();};
	void Signal_BorderColorChanged(QColor color) { callbackQAreaSeries_BorderColorChanged(this, new QColor(color)); };
	void Signal_Clicked(const QPointF & point) { callbackQAreaSeries_Clicked(this, const_cast<QPointF*>(&point)); };
	void Signal_ColorChanged(QColor color) { callbackQAreaSeries_ColorChanged(this, new QColor(color)); };
	void Signal_DoubleClicked(const QPointF & point) { callbackQAreaSeries_DoubleClicked(this, const_cast<QPointF*>(&point)); };
	void Signal_Hovered(const QPointF & point, bool state) { callbackQAreaSeries_Hovered(this, const_cast<QPointF*>(&point), state); };
	void Signal_PointLabelsClippingChanged(bool clipping) { callbackQAreaSeries_PointLabelsClippingChanged(this, clipping); };
	void Signal_PointLabelsColorChanged(const QColor & color) { callbackQAreaSeries_PointLabelsColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_PointLabelsFontChanged(const QFont & font) { callbackQAreaSeries_PointLabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_PointLabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQAreaSeries_PointLabelsFormatChanged(this, formatPacked); };
	void Signal_PointLabelsVisibilityChanged(bool visible) { callbackQAreaSeries_PointLabelsVisibilityChanged(this, visible); };
	void Signal_Pressed(const QPointF & point) { callbackQAreaSeries_Pressed(this, const_cast<QPointF*>(&point)); };
	void Signal_Released(const QPointF & point) { callbackQAreaSeries_Released(this, const_cast<QPointF*>(&point)); };
	 ~MyQAreaSeries() { callbackQAreaSeries_DestroyQAreaSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQAreaSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAreaSeries*)

int QAreaSeries_QAreaSeries_QRegisterMetaType(){qRegisterMetaType<QAreaSeries*>(); return qRegisterMetaType<MyQAreaSeries*>();}

void* QAreaSeries_NewQAreaSeries2(void* upperSeries, void* lowerSeries)
{
	return new MyQAreaSeries(static_cast<QLineSeries*>(upperSeries), static_cast<QLineSeries*>(lowerSeries));
}

void* QAreaSeries_NewQAreaSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAreaSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQAreaSeries(static_cast<QObject*>(parent));
	}
}

void QAreaSeries_ConnectBorderColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(QColor)>(&QAreaSeries::borderColorChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(QColor)>(&MyQAreaSeries::Signal_BorderColorChanged));
}

void QAreaSeries_DisconnectBorderColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(QColor)>(&QAreaSeries::borderColorChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(QColor)>(&MyQAreaSeries::Signal_BorderColorChanged));
}

void QAreaSeries_BorderColorChanged(void* ptr, void* color)
{
	static_cast<QAreaSeries*>(ptr)->borderColorChanged(*static_cast<QColor*>(color));
}

void QAreaSeries_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &)>(&QAreaSeries::clicked), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &)>(&MyQAreaSeries::Signal_Clicked));
}

void QAreaSeries_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &)>(&QAreaSeries::clicked), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &)>(&MyQAreaSeries::Signal_Clicked));
}

void QAreaSeries_Clicked(void* ptr, void* point)
{
	static_cast<QAreaSeries*>(ptr)->clicked(*static_cast<QPointF*>(point));
}

void QAreaSeries_ConnectColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(QColor)>(&QAreaSeries::colorChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(QColor)>(&MyQAreaSeries::Signal_ColorChanged));
}

void QAreaSeries_DisconnectColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(QColor)>(&QAreaSeries::colorChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(QColor)>(&MyQAreaSeries::Signal_ColorChanged));
}

void QAreaSeries_ColorChanged(void* ptr, void* color)
{
	static_cast<QAreaSeries*>(ptr)->colorChanged(*static_cast<QColor*>(color));
}

void QAreaSeries_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &)>(&QAreaSeries::doubleClicked), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &)>(&MyQAreaSeries::Signal_DoubleClicked));
}

void QAreaSeries_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &)>(&QAreaSeries::doubleClicked), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &)>(&MyQAreaSeries::Signal_DoubleClicked));
}

void QAreaSeries_DoubleClicked(void* ptr, void* point)
{
	static_cast<QAreaSeries*>(ptr)->doubleClicked(*static_cast<QPointF*>(point));
}

void QAreaSeries_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &, bool)>(&QAreaSeries::hovered), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &, bool)>(&MyQAreaSeries::Signal_Hovered));
}

void QAreaSeries_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &, bool)>(&QAreaSeries::hovered), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &, bool)>(&MyQAreaSeries::Signal_Hovered));
}

void QAreaSeries_Hovered(void* ptr, void* point, char state)
{
	static_cast<QAreaSeries*>(ptr)->hovered(*static_cast<QPointF*>(point), state != 0);
}

void QAreaSeries_ConnectPointLabelsClippingChanged(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(bool)>(&QAreaSeries::pointLabelsClippingChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(bool)>(&MyQAreaSeries::Signal_PointLabelsClippingChanged));
}

void QAreaSeries_DisconnectPointLabelsClippingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(bool)>(&QAreaSeries::pointLabelsClippingChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(bool)>(&MyQAreaSeries::Signal_PointLabelsClippingChanged));
}

void QAreaSeries_PointLabelsClippingChanged(void* ptr, char clipping)
{
	static_cast<QAreaSeries*>(ptr)->pointLabelsClippingChanged(clipping != 0);
}

void QAreaSeries_ConnectPointLabelsColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QColor &)>(&QAreaSeries::pointLabelsColorChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QColor &)>(&MyQAreaSeries::Signal_PointLabelsColorChanged));
}

void QAreaSeries_DisconnectPointLabelsColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QColor &)>(&QAreaSeries::pointLabelsColorChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QColor &)>(&MyQAreaSeries::Signal_PointLabelsColorChanged));
}

void QAreaSeries_PointLabelsColorChanged(void* ptr, void* color)
{
	static_cast<QAreaSeries*>(ptr)->pointLabelsColorChanged(*static_cast<QColor*>(color));
}

void QAreaSeries_ConnectPointLabelsFontChanged(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QFont &)>(&QAreaSeries::pointLabelsFontChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QFont &)>(&MyQAreaSeries::Signal_PointLabelsFontChanged));
}

void QAreaSeries_DisconnectPointLabelsFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QFont &)>(&QAreaSeries::pointLabelsFontChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QFont &)>(&MyQAreaSeries::Signal_PointLabelsFontChanged));
}

void QAreaSeries_PointLabelsFontChanged(void* ptr, void* font)
{
	static_cast<QAreaSeries*>(ptr)->pointLabelsFontChanged(*static_cast<QFont*>(font));
}

void QAreaSeries_ConnectPointLabelsFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QString &)>(&QAreaSeries::pointLabelsFormatChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QString &)>(&MyQAreaSeries::Signal_PointLabelsFormatChanged));
}

void QAreaSeries_DisconnectPointLabelsFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QString &)>(&QAreaSeries::pointLabelsFormatChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QString &)>(&MyQAreaSeries::Signal_PointLabelsFormatChanged));
}

void QAreaSeries_PointLabelsFormatChanged(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QAreaSeries*>(ptr)->pointLabelsFormatChanged(QString::fromUtf8(format.data, format.len));
}

void QAreaSeries_ConnectPointLabelsVisibilityChanged(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(bool)>(&QAreaSeries::pointLabelsVisibilityChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(bool)>(&MyQAreaSeries::Signal_PointLabelsVisibilityChanged));
}

void QAreaSeries_DisconnectPointLabelsVisibilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(bool)>(&QAreaSeries::pointLabelsVisibilityChanged), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(bool)>(&MyQAreaSeries::Signal_PointLabelsVisibilityChanged));
}

void QAreaSeries_PointLabelsVisibilityChanged(void* ptr, char visible)
{
	static_cast<QAreaSeries*>(ptr)->pointLabelsVisibilityChanged(visible != 0);
}

void QAreaSeries_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &)>(&QAreaSeries::pressed), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &)>(&MyQAreaSeries::Signal_Pressed));
}

void QAreaSeries_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &)>(&QAreaSeries::pressed), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &)>(&MyQAreaSeries::Signal_Pressed));
}

void QAreaSeries_Pressed(void* ptr, void* point)
{
	static_cast<QAreaSeries*>(ptr)->pressed(*static_cast<QPointF*>(point));
}

void QAreaSeries_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &)>(&QAreaSeries::released), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &)>(&MyQAreaSeries::Signal_Released));
}

void QAreaSeries_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QAreaSeries*>(ptr), static_cast<void (QAreaSeries::*)(const QPointF &)>(&QAreaSeries::released), static_cast<MyQAreaSeries*>(ptr), static_cast<void (MyQAreaSeries::*)(const QPointF &)>(&MyQAreaSeries::Signal_Released));
}

void QAreaSeries_Released(void* ptr, void* point)
{
	static_cast<QAreaSeries*>(ptr)->released(*static_cast<QPointF*>(point));
}

void QAreaSeries_SetBorderColor(void* ptr, void* color)
{
	static_cast<QAreaSeries*>(ptr)->setBorderColor(*static_cast<QColor*>(color));
}

void QAreaSeries_SetBrush(void* ptr, void* brush)
{
	static_cast<QAreaSeries*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QAreaSeries_SetColor(void* ptr, void* color)
{
	static_cast<QAreaSeries*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QAreaSeries_SetLowerSeries(void* ptr, void* series)
{
	static_cast<QAreaSeries*>(ptr)->setLowerSeries(static_cast<QLineSeries*>(series));
}

void QAreaSeries_SetPen(void* ptr, void* pen)
{
	static_cast<QAreaSeries*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QAreaSeries_SetPointLabelsClipping(void* ptr, char enabled)
{
	static_cast<QAreaSeries*>(ptr)->setPointLabelsClipping(enabled != 0);
}

void QAreaSeries_SetPointLabelsColor(void* ptr, void* color)
{
	static_cast<QAreaSeries*>(ptr)->setPointLabelsColor(*static_cast<QColor*>(color));
}

void QAreaSeries_SetPointLabelsFont(void* ptr, void* font)
{
	static_cast<QAreaSeries*>(ptr)->setPointLabelsFont(*static_cast<QFont*>(font));
}

void QAreaSeries_SetPointLabelsFormat(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QAreaSeries*>(ptr)->setPointLabelsFormat(QString::fromUtf8(format.data, format.len));
}

void QAreaSeries_SetPointLabelsVisible(void* ptr, char visible)
{
	static_cast<QAreaSeries*>(ptr)->setPointLabelsVisible(visible != 0);
}

void QAreaSeries_SetPointsVisible(void* ptr, char visible)
{
	static_cast<QAreaSeries*>(ptr)->setPointsVisible(visible != 0);
}

void QAreaSeries_SetUpperSeries(void* ptr, void* series)
{
	static_cast<QAreaSeries*>(ptr)->setUpperSeries(static_cast<QLineSeries*>(series));
}

void QAreaSeries_DestroyQAreaSeries(void* ptr)
{
	static_cast<QAreaSeries*>(ptr)->~QAreaSeries();
}

void QAreaSeries_DestroyQAreaSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QAreaSeries_Type(void* ptr)
{
	return static_cast<QAreaSeries*>(ptr)->type();
}

long long QAreaSeries_TypeDefault(void* ptr)
{
		return static_cast<QAreaSeries*>(ptr)->QAreaSeries::type();
}

void* QAreaSeries_Brush(void* ptr)
{
	return new QBrush(static_cast<QAreaSeries*>(ptr)->brush());
}

void* QAreaSeries_BorderColor(void* ptr)
{
	return new QColor(static_cast<QAreaSeries*>(ptr)->borderColor());
}

void* QAreaSeries_Color(void* ptr)
{
	return new QColor(static_cast<QAreaSeries*>(ptr)->color());
}

void* QAreaSeries_PointLabelsColor(void* ptr)
{
	return new QColor(static_cast<QAreaSeries*>(ptr)->pointLabelsColor());
}

void* QAreaSeries_PointLabelsFont(void* ptr)
{
	return new QFont(static_cast<QAreaSeries*>(ptr)->pointLabelsFont());
}

void* QAreaSeries_LowerSeries(void* ptr)
{
	return static_cast<QAreaSeries*>(ptr)->lowerSeries();
}

void* QAreaSeries_UpperSeries(void* ptr)
{
	return static_cast<QAreaSeries*>(ptr)->upperSeries();
}

void* QAreaSeries_Pen(void* ptr)
{
	return new QPen(static_cast<QAreaSeries*>(ptr)->pen());
}

struct QtCharts_PackedString QAreaSeries_PointLabelsFormat(void* ptr)
{
	return ({ QByteArray t607311 = static_cast<QAreaSeries*>(ptr)->pointLabelsFormat().toUtf8(); QtCharts_PackedString { const_cast<char*>(t607311.prepend("WHITESPACE").constData()+10), t607311.size()-10 }; });
}

char QAreaSeries_PointLabelsClipping(void* ptr)
{
	return static_cast<QAreaSeries*>(ptr)->pointLabelsClipping();
}

char QAreaSeries_PointLabelsVisible(void* ptr)
{
	return static_cast<QAreaSeries*>(ptr)->pointLabelsVisible();
}

char QAreaSeries_PointsVisible(void* ptr)
{
	return static_cast<QAreaSeries*>(ptr)->pointsVisible();
}

class MyQBarCategoryAxis: public QBarCategoryAxis
{
public:
	MyQBarCategoryAxis(QObject *parent = Q_NULLPTR) : QBarCategoryAxis(parent) {QBarCategoryAxis_QBarCategoryAxis_QRegisterMetaType();};
	void Signal_CategoriesChanged() { callbackQBarCategoryAxis_CategoriesChanged(this); };
	void Signal_CountChanged() { callbackQBarCategoryAxis_CountChanged(this); };
	void Signal_MaxChanged(const QString & max) { QByteArray t070602 = max.toUtf8(); QtCharts_PackedString maxPacked = { const_cast<char*>(t070602.prepend("WHITESPACE").constData()+10), t070602.size()-10 };callbackQBarCategoryAxis_MaxChanged(this, maxPacked); };
	void Signal_MinChanged(const QString & min) { QByteArray tb6c935 = min.toUtf8(); QtCharts_PackedString minPacked = { const_cast<char*>(tb6c935.prepend("WHITESPACE").constData()+10), tb6c935.size()-10 };callbackQBarCategoryAxis_MinChanged(this, minPacked); };
	void Signal_RangeChanged(const QString & min, const QString & max) { QByteArray tb6c935 = min.toUtf8(); QtCharts_PackedString minPacked = { const_cast<char*>(tb6c935.prepend("WHITESPACE").constData()+10), tb6c935.size()-10 };QByteArray t070602 = max.toUtf8(); QtCharts_PackedString maxPacked = { const_cast<char*>(t070602.prepend("WHITESPACE").constData()+10), t070602.size()-10 };callbackQBarCategoryAxis_RangeChanged(this, minPacked, maxPacked); };
	 ~MyQBarCategoryAxis() { callbackQBarCategoryAxis_DestroyQBarCategoryAxis(this); };
	QAbstractAxis::AxisType type() const { return static_cast<QAbstractAxis::AxisType>(callbackQBarCategoryAxis_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractAxis_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ColorChanged(QColor color) { callbackQAbstractAxis_ColorChanged(this, new QColor(color)); };
	void Signal_GridLineColorChanged(const QColor & color) { callbackQAbstractAxis_GridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_GridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_GridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_GridVisibleChanged(bool visible) { callbackQAbstractAxis_GridVisibleChanged(this, visible); };
	void Signal_LabelsAngleChanged(int angle) { callbackQAbstractAxis_LabelsAngleChanged(this, angle); };
	void Signal_LabelsBrushChanged(const QBrush & brush) { callbackQAbstractAxis_LabelsBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_LabelsColorChanged(QColor color) { callbackQAbstractAxis_LabelsColorChanged(this, new QColor(color)); };
	void Signal_LabelsFontChanged(const QFont & font) { callbackQAbstractAxis_LabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_LabelsVisibleChanged(bool visible) { callbackQAbstractAxis_LabelsVisibleChanged(this, visible); };
	void Signal_LinePenChanged(const QPen & pen) { callbackQAbstractAxis_LinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_LineVisibleChanged(bool visible) { callbackQAbstractAxis_LineVisibleChanged(this, visible); };
	void Signal_MinorGridLineColorChanged(const QColor & color) { callbackQAbstractAxis_MinorGridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_MinorGridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_MinorGridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_MinorGridVisibleChanged(bool visible) { callbackQAbstractAxis_MinorGridVisibleChanged(this, visible); };
	void Signal_ReverseChanged(bool reverse) { callbackQAbstractAxis_ReverseChanged(this, reverse); };
	void Signal_ShadesBorderColorChanged(QColor color) { callbackQAbstractAxis_ShadesBorderColorChanged(this, new QColor(color)); };
	void Signal_ShadesBrushChanged(const QBrush & brush) { callbackQAbstractAxis_ShadesBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_ShadesColorChanged(QColor color) { callbackQAbstractAxis_ShadesColorChanged(this, new QColor(color)); };
	void Signal_ShadesPenChanged(const QPen & pen) { callbackQAbstractAxis_ShadesPenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_ShadesVisibleChanged(bool visible) { callbackQAbstractAxis_ShadesVisibleChanged(this, visible); };
	void Signal_TitleBrushChanged(const QBrush & brush) { callbackQAbstractAxis_TitleBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_TitleFontChanged(const QFont & font) { callbackQAbstractAxis_TitleFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_TitleTextChanged(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtCharts_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQAbstractAxis_TitleTextChanged(this, textPacked); };
	void Signal_TitleVisibleChanged(bool visible) { callbackQAbstractAxis_TitleVisibleChanged(this, visible); };
	void Signal_VisibleChanged(bool visible) { callbackQAbstractAxis_VisibleChanged(this, visible); };
	bool event(QEvent * e) { return callbackQAbstractAxis_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractAxis_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractAxis_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractAxis_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractAxis_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractAxis_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractAxis_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractAxis_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQBarCategoryAxis*)

int QBarCategoryAxis_QBarCategoryAxis_QRegisterMetaType(){qRegisterMetaType<QBarCategoryAxis*>(); return qRegisterMetaType<MyQBarCategoryAxis*>();}

void* QBarCategoryAxis_NewQBarCategoryAxis(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBarCategoryAxis(static_cast<QWindow*>(parent));
	} else {
		return new MyQBarCategoryAxis(static_cast<QObject*>(parent));
	}
}

struct QtCharts_PackedString QBarCategoryAxis_Categories(void* ptr)
{
	return ({ QByteArray tcc0bf4 = static_cast<QBarCategoryAxis*>(ptr)->categories().join("|").toUtf8(); QtCharts_PackedString { const_cast<char*>(tcc0bf4.prepend("WHITESPACE").constData()+10), tcc0bf4.size()-10 }; });
}

void QBarCategoryAxis_Append2(void* ptr, struct QtCharts_PackedString category)
{
	static_cast<QBarCategoryAxis*>(ptr)->append(QString::fromUtf8(category.data, category.len));
}

void QBarCategoryAxis_Append(void* ptr, struct QtCharts_PackedString categories)
{
	static_cast<QBarCategoryAxis*>(ptr)->append(QString::fromUtf8(categories.data, categories.len).split("|", QString::SkipEmptyParts));
}

void QBarCategoryAxis_ConnectCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)()>(&QBarCategoryAxis::categoriesChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)()>(&MyQBarCategoryAxis::Signal_CategoriesChanged));
}

void QBarCategoryAxis_DisconnectCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)()>(&QBarCategoryAxis::categoriesChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)()>(&MyQBarCategoryAxis::Signal_CategoriesChanged));
}

void QBarCategoryAxis_CategoriesChanged(void* ptr)
{
	static_cast<QBarCategoryAxis*>(ptr)->categoriesChanged();
}

void QBarCategoryAxis_Clear(void* ptr)
{
	static_cast<QBarCategoryAxis*>(ptr)->clear();
}

void QBarCategoryAxis_ConnectCountChanged(void* ptr)
{
	QObject::connect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)()>(&QBarCategoryAxis::countChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)()>(&MyQBarCategoryAxis::Signal_CountChanged));
}

void QBarCategoryAxis_DisconnectCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)()>(&QBarCategoryAxis::countChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)()>(&MyQBarCategoryAxis::Signal_CountChanged));
}

void QBarCategoryAxis_CountChanged(void* ptr)
{
	static_cast<QBarCategoryAxis*>(ptr)->countChanged();
}

void QBarCategoryAxis_Insert(void* ptr, int index, struct QtCharts_PackedString category)
{
	static_cast<QBarCategoryAxis*>(ptr)->insert(index, QString::fromUtf8(category.data, category.len));
}

void QBarCategoryAxis_ConnectMaxChanged(void* ptr)
{
	QObject::connect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)(const QString &)>(&QBarCategoryAxis::maxChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)(const QString &)>(&MyQBarCategoryAxis::Signal_MaxChanged));
}

void QBarCategoryAxis_DisconnectMaxChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)(const QString &)>(&QBarCategoryAxis::maxChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)(const QString &)>(&MyQBarCategoryAxis::Signal_MaxChanged));
}

void QBarCategoryAxis_MaxChanged(void* ptr, struct QtCharts_PackedString max)
{
	static_cast<QBarCategoryAxis*>(ptr)->maxChanged(QString::fromUtf8(max.data, max.len));
}

void QBarCategoryAxis_ConnectMinChanged(void* ptr)
{
	QObject::connect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)(const QString &)>(&QBarCategoryAxis::minChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)(const QString &)>(&MyQBarCategoryAxis::Signal_MinChanged));
}

void QBarCategoryAxis_DisconnectMinChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)(const QString &)>(&QBarCategoryAxis::minChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)(const QString &)>(&MyQBarCategoryAxis::Signal_MinChanged));
}

void QBarCategoryAxis_MinChanged(void* ptr, struct QtCharts_PackedString min)
{
	static_cast<QBarCategoryAxis*>(ptr)->minChanged(QString::fromUtf8(min.data, min.len));
}

void QBarCategoryAxis_ConnectRangeChanged(void* ptr)
{
	QObject::connect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)(const QString &, const QString &)>(&QBarCategoryAxis::rangeChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)(const QString &, const QString &)>(&MyQBarCategoryAxis::Signal_RangeChanged));
}

void QBarCategoryAxis_DisconnectRangeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarCategoryAxis*>(ptr), static_cast<void (QBarCategoryAxis::*)(const QString &, const QString &)>(&QBarCategoryAxis::rangeChanged), static_cast<MyQBarCategoryAxis*>(ptr), static_cast<void (MyQBarCategoryAxis::*)(const QString &, const QString &)>(&MyQBarCategoryAxis::Signal_RangeChanged));
}

void QBarCategoryAxis_RangeChanged(void* ptr, struct QtCharts_PackedString min, struct QtCharts_PackedString max)
{
	static_cast<QBarCategoryAxis*>(ptr)->rangeChanged(QString::fromUtf8(min.data, min.len), QString::fromUtf8(max.data, max.len));
}

void QBarCategoryAxis_Remove(void* ptr, struct QtCharts_PackedString category)
{
	static_cast<QBarCategoryAxis*>(ptr)->remove(QString::fromUtf8(category.data, category.len));
}

void QBarCategoryAxis_Replace(void* ptr, struct QtCharts_PackedString oldCategory, struct QtCharts_PackedString newCategory)
{
	static_cast<QBarCategoryAxis*>(ptr)->replace(QString::fromUtf8(oldCategory.data, oldCategory.len), QString::fromUtf8(newCategory.data, newCategory.len));
}

void QBarCategoryAxis_SetCategories(void* ptr, struct QtCharts_PackedString categories)
{
	static_cast<QBarCategoryAxis*>(ptr)->setCategories(QString::fromUtf8(categories.data, categories.len).split("|", QString::SkipEmptyParts));
}

void QBarCategoryAxis_SetMax(void* ptr, struct QtCharts_PackedString max)
{
	static_cast<QBarCategoryAxis*>(ptr)->setMax(QString::fromUtf8(max.data, max.len));
}

void QBarCategoryAxis_SetMin(void* ptr, struct QtCharts_PackedString min)
{
	static_cast<QBarCategoryAxis*>(ptr)->setMin(QString::fromUtf8(min.data, min.len));
}

void QBarCategoryAxis_SetRange(void* ptr, struct QtCharts_PackedString minCategory, struct QtCharts_PackedString maxCategory)
{
	static_cast<QBarCategoryAxis*>(ptr)->setRange(QString::fromUtf8(minCategory.data, minCategory.len), QString::fromUtf8(maxCategory.data, maxCategory.len));
}

void QBarCategoryAxis_DestroyQBarCategoryAxis(void* ptr)
{
	static_cast<QBarCategoryAxis*>(ptr)->~QBarCategoryAxis();
}

void QBarCategoryAxis_DestroyQBarCategoryAxisDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QBarCategoryAxis_Type(void* ptr)
{
	return static_cast<QBarCategoryAxis*>(ptr)->type();
}

long long QBarCategoryAxis_TypeDefault(void* ptr)
{
		return static_cast<QBarCategoryAxis*>(ptr)->QBarCategoryAxis::type();
}

struct QtCharts_PackedString QBarCategoryAxis_At(void* ptr, int index)
{
	return ({ QByteArray tf55150 = static_cast<QBarCategoryAxis*>(ptr)->at(index).toUtf8(); QtCharts_PackedString { const_cast<char*>(tf55150.prepend("WHITESPACE").constData()+10), tf55150.size()-10 }; });
}

struct QtCharts_PackedString QBarCategoryAxis_Max(void* ptr)
{
	return ({ QByteArray tba51dd = static_cast<QBarCategoryAxis*>(ptr)->max().toUtf8(); QtCharts_PackedString { const_cast<char*>(tba51dd.prepend("WHITESPACE").constData()+10), tba51dd.size()-10 }; });
}

struct QtCharts_PackedString QBarCategoryAxis_Min(void* ptr)
{
	return ({ QByteArray tbfa107 = static_cast<QBarCategoryAxis*>(ptr)->min().toUtf8(); QtCharts_PackedString { const_cast<char*>(tbfa107.prepend("WHITESPACE").constData()+10), tbfa107.size()-10 }; });
}

int QBarCategoryAxis_Count(void* ptr)
{
	return static_cast<QBarCategoryAxis*>(ptr)->count();
}

class MyQBarLegendMarker: public QBarLegendMarker
{
public:
	QAbstractBarSeries * series() { return static_cast<QAbstractBarSeries*>(callbackQBarLegendMarker_Series(this)); };
	QLegendMarker::LegendMarkerType type() { return static_cast<QLegendMarker::LegendMarkerType>(callbackQBarLegendMarker_Type(this)); };
	 ~MyQBarLegendMarker() { callbackQBarLegendMarker_DestroyQBarLegendMarker(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLegendMarker_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BrushChanged() { callbackQLegendMarker_BrushChanged(this); };
	void Signal_Clicked() { callbackQLegendMarker_Clicked(this); };
	void Signal_FontChanged() { callbackQLegendMarker_FontChanged(this); };
	void Signal_Hovered(bool status) { callbackQLegendMarker_Hovered(this, status); };
	void Signal_LabelBrushChanged() { callbackQLegendMarker_LabelBrushChanged(this); };
	void Signal_LabelChanged() { callbackQLegendMarker_LabelChanged(this); };
	void Signal_PenChanged() { callbackQLegendMarker_PenChanged(this); };
	void Signal_ShapeChanged() { callbackQLegendMarker_ShapeChanged(this); };
	void Signal_VisibleChanged() { callbackQLegendMarker_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQLegendMarker_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLegendMarker_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLegendMarker_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLegendMarker_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLegendMarker_CustomEvent(this, event); };
	void deleteLater() { callbackQLegendMarker_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLegendMarker_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLegendMarker_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLegendMarker_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLegendMarker_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQBarLegendMarker*)

int QBarLegendMarker_QBarLegendMarker_QRegisterMetaType(){qRegisterMetaType<QBarLegendMarker*>(); return qRegisterMetaType<MyQBarLegendMarker*>();}

void* QBarLegendMarker_Series(void* ptr)
{
	return static_cast<QBarLegendMarker*>(ptr)->series();
}

void* QBarLegendMarker_SeriesDefault(void* ptr)
{
		return static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::series();
}

void* QBarLegendMarker_Barset(void* ptr)
{
	return static_cast<QBarLegendMarker*>(ptr)->barset();
}

long long QBarLegendMarker_Type(void* ptr)
{
	return static_cast<QBarLegendMarker*>(ptr)->type();
}

long long QBarLegendMarker_TypeDefault(void* ptr)
{
		return static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::type();
}

void QBarLegendMarker_DestroyQBarLegendMarker(void* ptr)
{
	static_cast<QBarLegendMarker*>(ptr)->~QBarLegendMarker();
}

void QBarLegendMarker_DestroyQBarLegendMarkerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQBarSeries: public QBarSeries
{
public:
	MyQBarSeries(QObject *parent = Q_NULLPTR) : QBarSeries(parent) {QBarSeries_QBarSeries_QRegisterMetaType();};
	 ~MyQBarSeries() { callbackQBarSeries_DestroyQBarSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQBarSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BarsetsAdded(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsAdded(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BarsetsRemoved(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsRemoved(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Clicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_Clicked(this, index, barset); };
	void Signal_CountChanged() { callbackQAbstractBarSeries_CountChanged(this); };
	void Signal_DoubleClicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_DoubleClicked(this, index, barset); };
	void Signal_Hovered(bool status, int index, QBarSet * barset) { callbackQAbstractBarSeries_Hovered(this, status, index, barset); };
	void Signal_LabelsAngleChanged(qreal angle) { callbackQAbstractBarSeries_LabelsAngleChanged(this, angle); };
	void Signal_LabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQAbstractBarSeries_LabelsFormatChanged(this, formatPacked); };
	void Signal_LabelsPositionChanged(QAbstractBarSeries::LabelsPosition position) { callbackQAbstractBarSeries_LabelsPositionChanged(this, position); };
	void Signal_LabelsPrecisionChanged(int precision) { callbackQAbstractBarSeries_LabelsPrecisionChanged(this, precision); };
	void Signal_LabelsVisibleChanged() { callbackQAbstractBarSeries_LabelsVisibleChanged(this); };
	void Signal_Pressed(int index, QBarSet * barset) { callbackQAbstractBarSeries_Pressed(this, index, barset); };
	void Signal_Released(int index, QBarSet * barset) { callbackQAbstractBarSeries_Released(this, index, barset); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQBarSeries*)

int QBarSeries_QBarSeries_QRegisterMetaType(){qRegisterMetaType<QBarSeries*>(); return qRegisterMetaType<MyQBarSeries*>();}

void* QBarSeries_NewQBarSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBarSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQBarSeries(static_cast<QObject*>(parent));
	}
}

void QBarSeries_DestroyQBarSeries(void* ptr)
{
	static_cast<QBarSeries*>(ptr)->~QBarSeries();
}

void QBarSeries_DestroyQBarSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QBarSeries_Type(void* ptr)
{
	return static_cast<QBarSeries*>(ptr)->type();
}

long long QBarSeries_TypeDefault(void* ptr)
{
		return static_cast<QBarSeries*>(ptr)->QBarSeries::type();
}

class MyQBarSet: public QBarSet
{
public:
	MyQBarSet(const QString label, QObject *parent = Q_NULLPTR) : QBarSet(label, parent) {QBarSet_QBarSet_QRegisterMetaType();};
	void Signal_BorderColorChanged(QColor color) { callbackQBarSet_BorderColorChanged(this, new QColor(color)); };
	void Signal_BrushChanged() { callbackQBarSet_BrushChanged(this); };
	void Signal_Clicked(int index) { callbackQBarSet_Clicked(this, index); };
	void Signal_ColorChanged(QColor color) { callbackQBarSet_ColorChanged(this, new QColor(color)); };
	void Signal_DoubleClicked(int index) { callbackQBarSet_DoubleClicked(this, index); };
	void Signal_Hovered(bool status, int index) { callbackQBarSet_Hovered(this, status, index); };
	void Signal_LabelBrushChanged() { callbackQBarSet_LabelBrushChanged(this); };
	void Signal_LabelChanged() { callbackQBarSet_LabelChanged(this); };
	void Signal_LabelColorChanged(QColor color) { callbackQBarSet_LabelColorChanged(this, new QColor(color)); };
	void Signal_LabelFontChanged() { callbackQBarSet_LabelFontChanged(this); };
	void Signal_PenChanged() { callbackQBarSet_PenChanged(this); };
	void Signal_Pressed(int index) { callbackQBarSet_Pressed(this, index); };
	void Signal_Released(int index) { callbackQBarSet_Released(this, index); };
	void Signal_ValueChanged(int index) { callbackQBarSet_ValueChanged(this, index); };
	void Signal_ValuesAdded(int index, int count) { callbackQBarSet_ValuesAdded(this, index, count); };
	void Signal_ValuesRemoved(int index, int count) { callbackQBarSet_ValuesRemoved(this, index, count); };
	 ~MyQBarSet() { callbackQBarSet_DestroyQBarSet(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQBarSet_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQBarSet_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQBarSet_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQBarSet_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQBarSet_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQBarSet_CustomEvent(this, event); };
	void deleteLater() { callbackQBarSet_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQBarSet_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQBarSet_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQBarSet_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQBarSet_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQBarSet*)

int QBarSet_QBarSet_QRegisterMetaType(){qRegisterMetaType<QBarSet*>(); return qRegisterMetaType<MyQBarSet*>();}

void* QBarSet_NewQBarSet(struct QtCharts_PackedString label, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQBarSet(QString::fromUtf8(label.data, label.len), static_cast<QObject*>(parent));
	}
}

void* QBarSet_BorderColor(void* ptr)
{
	return new QColor(static_cast<QBarSet*>(ptr)->borderColor());
}

void* QBarSet_Color(void* ptr)
{
	return new QColor(static_cast<QBarSet*>(ptr)->color());
}

void* QBarSet_LabelColor(void* ptr)
{
	return new QColor(static_cast<QBarSet*>(ptr)->labelColor());
}

struct QtCharts_PackedString QBarSet_QBarSet_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t994acf = QBarSet::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t994acf.prepend("WHITESPACE").constData()+10), t994acf.size()-10 }; });
}

struct QtCharts_PackedString QBarSet_QBarSet_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tb413d1 = QBarSet::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(tb413d1.prepend("WHITESPACE").constData()+10), tb413d1.size()-10 }; });
}

void QBarSet_Append2(void* ptr, void* values)
{
	static_cast<QBarSet*>(ptr)->append(*static_cast<QList<qreal>*>(values));
}

void QBarSet_Append(void* ptr, double value)
{
	static_cast<QBarSet*>(ptr)->append(value);
}

void QBarSet_ConnectBorderColorChanged(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(QColor)>(&QBarSet::borderColorChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(QColor)>(&MyQBarSet::Signal_BorderColorChanged));
}

void QBarSet_DisconnectBorderColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(QColor)>(&QBarSet::borderColorChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(QColor)>(&MyQBarSet::Signal_BorderColorChanged));
}

void QBarSet_BorderColorChanged(void* ptr, void* color)
{
	static_cast<QBarSet*>(ptr)->borderColorChanged(*static_cast<QColor*>(color));
}

void QBarSet_ConnectBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::brushChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_BrushChanged));
}

void QBarSet_DisconnectBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::brushChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_BrushChanged));
}

void QBarSet_BrushChanged(void* ptr)
{
	static_cast<QBarSet*>(ptr)->brushChanged();
}

void QBarSet_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::clicked), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_Clicked));
}

void QBarSet_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::clicked), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_Clicked));
}

void QBarSet_Clicked(void* ptr, int index)
{
	static_cast<QBarSet*>(ptr)->clicked(index);
}

void QBarSet_ConnectColorChanged(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(QColor)>(&QBarSet::colorChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(QColor)>(&MyQBarSet::Signal_ColorChanged));
}

void QBarSet_DisconnectColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(QColor)>(&QBarSet::colorChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(QColor)>(&MyQBarSet::Signal_ColorChanged));
}

void QBarSet_ColorChanged(void* ptr, void* color)
{
	static_cast<QBarSet*>(ptr)->colorChanged(*static_cast<QColor*>(color));
}

void QBarSet_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::doubleClicked), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_DoubleClicked));
}

void QBarSet_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::doubleClicked), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_DoubleClicked));
}

void QBarSet_DoubleClicked(void* ptr, int index)
{
	static_cast<QBarSet*>(ptr)->doubleClicked(index);
}

void QBarSet_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(bool, int)>(&QBarSet::hovered), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(bool, int)>(&MyQBarSet::Signal_Hovered));
}

void QBarSet_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(bool, int)>(&QBarSet::hovered), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(bool, int)>(&MyQBarSet::Signal_Hovered));
}

void QBarSet_Hovered(void* ptr, char status, int index)
{
	static_cast<QBarSet*>(ptr)->hovered(status != 0, index);
}

void QBarSet_Insert(void* ptr, int index, double value)
{
	static_cast<QBarSet*>(ptr)->insert(index, value);
}

void QBarSet_ConnectLabelBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::labelBrushChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_LabelBrushChanged));
}

void QBarSet_DisconnectLabelBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::labelBrushChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_LabelBrushChanged));
}

void QBarSet_LabelBrushChanged(void* ptr)
{
	static_cast<QBarSet*>(ptr)->labelBrushChanged();
}

void QBarSet_ConnectLabelChanged(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::labelChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_LabelChanged));
}

void QBarSet_DisconnectLabelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::labelChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_LabelChanged));
}

void QBarSet_LabelChanged(void* ptr)
{
	static_cast<QBarSet*>(ptr)->labelChanged();
}

void QBarSet_ConnectLabelColorChanged(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(QColor)>(&QBarSet::labelColorChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(QColor)>(&MyQBarSet::Signal_LabelColorChanged));
}

void QBarSet_DisconnectLabelColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(QColor)>(&QBarSet::labelColorChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(QColor)>(&MyQBarSet::Signal_LabelColorChanged));
}

void QBarSet_LabelColorChanged(void* ptr, void* color)
{
	static_cast<QBarSet*>(ptr)->labelColorChanged(*static_cast<QColor*>(color));
}

void QBarSet_ConnectLabelFontChanged(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::labelFontChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_LabelFontChanged));
}

void QBarSet_DisconnectLabelFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::labelFontChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_LabelFontChanged));
}

void QBarSet_LabelFontChanged(void* ptr)
{
	static_cast<QBarSet*>(ptr)->labelFontChanged();
}

void QBarSet_ConnectPenChanged(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::penChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_PenChanged));
}

void QBarSet_DisconnectPenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)()>(&QBarSet::penChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)()>(&MyQBarSet::Signal_PenChanged));
}

void QBarSet_PenChanged(void* ptr)
{
	static_cast<QBarSet*>(ptr)->penChanged();
}

void QBarSet_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::pressed), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_Pressed));
}

void QBarSet_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::pressed), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_Pressed));
}

void QBarSet_Pressed(void* ptr, int index)
{
	static_cast<QBarSet*>(ptr)->pressed(index);
}

void QBarSet_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::released), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_Released));
}

void QBarSet_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::released), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_Released));
}

void QBarSet_Released(void* ptr, int index)
{
	static_cast<QBarSet*>(ptr)->released(index);
}

void QBarSet_Remove(void* ptr, int index, int count)
{
	static_cast<QBarSet*>(ptr)->remove(index, count);
}

void QBarSet_Replace(void* ptr, int index, double value)
{
	static_cast<QBarSet*>(ptr)->replace(index, value);
}

void QBarSet_SetBorderColor(void* ptr, void* color)
{
	static_cast<QBarSet*>(ptr)->setBorderColor(*static_cast<QColor*>(color));
}

void QBarSet_SetBrush(void* ptr, void* brush)
{
	static_cast<QBarSet*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QBarSet_SetColor(void* ptr, void* color)
{
	static_cast<QBarSet*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QBarSet_SetLabel(void* ptr, struct QtCharts_PackedString label)
{
	static_cast<QBarSet*>(ptr)->setLabel(QString::fromUtf8(label.data, label.len));
}

void QBarSet_SetLabelBrush(void* ptr, void* brush)
{
	static_cast<QBarSet*>(ptr)->setLabelBrush(*static_cast<QBrush*>(brush));
}

void QBarSet_SetLabelColor(void* ptr, void* color)
{
	static_cast<QBarSet*>(ptr)->setLabelColor(*static_cast<QColor*>(color));
}

void QBarSet_SetLabelFont(void* ptr, void* font)
{
	static_cast<QBarSet*>(ptr)->setLabelFont(*static_cast<QFont*>(font));
}

void QBarSet_SetPen(void* ptr, void* pen)
{
	static_cast<QBarSet*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QBarSet_ConnectValueChanged(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::valueChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_ValueChanged));
}

void QBarSet_DisconnectValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int)>(&QBarSet::valueChanged), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int)>(&MyQBarSet::Signal_ValueChanged));
}

void QBarSet_ValueChanged(void* ptr, int index)
{
	static_cast<QBarSet*>(ptr)->valueChanged(index);
}

void QBarSet_ConnectValuesAdded(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int, int)>(&QBarSet::valuesAdded), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int, int)>(&MyQBarSet::Signal_ValuesAdded));
}

void QBarSet_DisconnectValuesAdded(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int, int)>(&QBarSet::valuesAdded), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int, int)>(&MyQBarSet::Signal_ValuesAdded));
}

void QBarSet_ValuesAdded(void* ptr, int index, int count)
{
	static_cast<QBarSet*>(ptr)->valuesAdded(index, count);
}

void QBarSet_ConnectValuesRemoved(void* ptr)
{
	QObject::connect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int, int)>(&QBarSet::valuesRemoved), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int, int)>(&MyQBarSet::Signal_ValuesRemoved));
}

void QBarSet_DisconnectValuesRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QBarSet*>(ptr), static_cast<void (QBarSet::*)(int, int)>(&QBarSet::valuesRemoved), static_cast<MyQBarSet*>(ptr), static_cast<void (MyQBarSet::*)(int, int)>(&MyQBarSet::Signal_ValuesRemoved));
}

void QBarSet_ValuesRemoved(void* ptr, int index, int count)
{
	static_cast<QBarSet*>(ptr)->valuesRemoved(index, count);
}

void QBarSet_DestroyQBarSet(void* ptr)
{
	static_cast<QBarSet*>(ptr)->~QBarSet();
}

void QBarSet_DestroyQBarSetDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QBarSet_Brush(void* ptr)
{
	return new QBrush(static_cast<QBarSet*>(ptr)->brush());
}

void* QBarSet_LabelBrush(void* ptr)
{
	return new QBrush(static_cast<QBarSet*>(ptr)->labelBrush());
}

void* QBarSet_LabelFont(void* ptr)
{
	return new QFont(static_cast<QBarSet*>(ptr)->labelFont());
}

void* QBarSet_Pen(void* ptr)
{
	return new QPen(static_cast<QBarSet*>(ptr)->pen());
}

struct QtCharts_PackedString QBarSet_Label(void* ptr)
{
	return ({ QByteArray t266af2 = static_cast<QBarSet*>(ptr)->label().toUtf8(); QtCharts_PackedString { const_cast<char*>(t266af2.prepend("WHITESPACE").constData()+10), t266af2.size()-10 }; });
}

void* QBarSet_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QBarSet*>(ptr)->QBarSet::metaObject());
}

int QBarSet_Count(void* ptr)
{
	return static_cast<QBarSet*>(ptr)->count();
}

double QBarSet_At(void* ptr, int index)
{
	return static_cast<QBarSet*>(ptr)->at(index);
}

double QBarSet_Sum(void* ptr)
{
	return static_cast<QBarSet*>(ptr)->sum();
}

double QBarSet___append_values_atList2(void* ptr, int i)
{
	return ({qreal tmp = static_cast<QList<qreal>*>(ptr)->at(i); if (i == static_cast<QList<qreal>*>(ptr)->size()-1) { static_cast<QList<qreal>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBarSet___append_values_setList2(void* ptr, double i)
{
	static_cast<QList<qreal>*>(ptr)->append(i);
}

void* QBarSet___append_values_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qreal>();
}

void* QBarSet___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QBarSet___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QBarSet___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QBarSet___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBarSet___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QBarSet___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QBarSet___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBarSet___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QBarSet___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QBarSet___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBarSet___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QBarSet___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QBarSet___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBarSet___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QBarSet___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QBarSet_EventDefault(void* ptr, void* e)
{
		return static_cast<QBarSet*>(ptr)->QBarSet::event(static_cast<QEvent*>(e));
}

char QBarSet_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QBarSet*>(ptr)->QBarSet::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QBarSet_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QBarSet*>(ptr)->QBarSet::childEvent(static_cast<QChildEvent*>(event));
}

void QBarSet_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QBarSet*>(ptr)->QBarSet::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBarSet_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QBarSet*>(ptr)->QBarSet::customEvent(static_cast<QEvent*>(event));
}

void QBarSet_DeleteLaterDefault(void* ptr)
{
		static_cast<QBarSet*>(ptr)->QBarSet::deleteLater();
}

void QBarSet_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QBarSet*>(ptr)->QBarSet::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBarSet_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QBarSet*>(ptr)->QBarSet::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQBoxPlotLegendMarker: public QBoxPlotLegendMarker
{
public:
	QBoxPlotSeries * series() { return static_cast<QBoxPlotSeries*>(callbackQBoxPlotLegendMarker_Series(this)); };
	QLegendMarker::LegendMarkerType type() { return static_cast<QLegendMarker::LegendMarkerType>(callbackQBoxPlotLegendMarker_Type(this)); };
	 ~MyQBoxPlotLegendMarker() { callbackQBoxPlotLegendMarker_DestroyQBoxPlotLegendMarker(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLegendMarker_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BrushChanged() { callbackQLegendMarker_BrushChanged(this); };
	void Signal_Clicked() { callbackQLegendMarker_Clicked(this); };
	void Signal_FontChanged() { callbackQLegendMarker_FontChanged(this); };
	void Signal_Hovered(bool status) { callbackQLegendMarker_Hovered(this, status); };
	void Signal_LabelBrushChanged() { callbackQLegendMarker_LabelBrushChanged(this); };
	void Signal_LabelChanged() { callbackQLegendMarker_LabelChanged(this); };
	void Signal_PenChanged() { callbackQLegendMarker_PenChanged(this); };
	void Signal_ShapeChanged() { callbackQLegendMarker_ShapeChanged(this); };
	void Signal_VisibleChanged() { callbackQLegendMarker_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQLegendMarker_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLegendMarker_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLegendMarker_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLegendMarker_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLegendMarker_CustomEvent(this, event); };
	void deleteLater() { callbackQLegendMarker_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLegendMarker_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLegendMarker_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLegendMarker_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLegendMarker_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQBoxPlotLegendMarker*)

int QBoxPlotLegendMarker_QBoxPlotLegendMarker_QRegisterMetaType(){qRegisterMetaType<QBoxPlotLegendMarker*>(); return qRegisterMetaType<MyQBoxPlotLegendMarker*>();}

void* QBoxPlotLegendMarker_Series(void* ptr)
{
	return static_cast<QBoxPlotLegendMarker*>(ptr)->series();
}

void* QBoxPlotLegendMarker_SeriesDefault(void* ptr)
{
		return static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::series();
}

long long QBoxPlotLegendMarker_Type(void* ptr)
{
	return static_cast<QBoxPlotLegendMarker*>(ptr)->type();
}

long long QBoxPlotLegendMarker_TypeDefault(void* ptr)
{
		return static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::type();
}

void QBoxPlotLegendMarker_DestroyQBoxPlotLegendMarker(void* ptr)
{
	static_cast<QBoxPlotLegendMarker*>(ptr)->~QBoxPlotLegendMarker();
}

void QBoxPlotLegendMarker_DestroyQBoxPlotLegendMarkerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQBoxPlotSeries: public QBoxPlotSeries
{
public:
	MyQBoxPlotSeries(QObject *parent = Q_NULLPTR) : QBoxPlotSeries(parent) {QBoxPlotSeries_QBoxPlotSeries_QRegisterMetaType();};
	void Signal_BoxOutlineVisibilityChanged() { callbackQBoxPlotSeries_BoxOutlineVisibilityChanged(this); };
	void Signal_BoxWidthChanged() { callbackQBoxPlotSeries_BoxWidthChanged(this); };
	void Signal_BoxsetsAdded(QList<QBoxSet *> sets) { callbackQBoxPlotSeries_BoxsetsAdded(this, ({ QList<QBoxSet *>* tmpValue = new QList<QBoxSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BoxsetsRemoved(QList<QBoxSet *> sets) { callbackQBoxPlotSeries_BoxsetsRemoved(this, ({ QList<QBoxSet *>* tmpValue = new QList<QBoxSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BrushChanged() { callbackQBoxPlotSeries_BrushChanged(this); };
	void Signal_Clicked(QBoxSet * boxset) { callbackQBoxPlotSeries_Clicked(this, boxset); };
	void Signal_CountChanged() { callbackQBoxPlotSeries_CountChanged(this); };
	void Signal_DoubleClicked(QBoxSet * boxset) { callbackQBoxPlotSeries_DoubleClicked(this, boxset); };
	void Signal_Hovered(bool status, QBoxSet * boxset) { callbackQBoxPlotSeries_Hovered(this, status, boxset); };
	void Signal_PenChanged() { callbackQBoxPlotSeries_PenChanged(this); };
	void Signal_Pressed(QBoxSet * boxset) { callbackQBoxPlotSeries_Pressed(this, boxset); };
	void Signal_Released(QBoxSet * boxset) { callbackQBoxPlotSeries_Released(this, boxset); };
	 ~MyQBoxPlotSeries() { callbackQBoxPlotSeries_DestroyQBoxPlotSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQBoxPlotSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQBoxPlotSeries*)

int QBoxPlotSeries_QBoxPlotSeries_QRegisterMetaType(){qRegisterMetaType<QBoxPlotSeries*>(); return qRegisterMetaType<MyQBoxPlotSeries*>();}

void* QBoxPlotSeries_NewQBoxPlotSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBoxPlotSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQBoxPlotSeries(static_cast<QObject*>(parent));
	}
}

char QBoxPlotSeries_Append(void* ptr, void* set)
{
	return static_cast<QBoxPlotSeries*>(ptr)->append(static_cast<QBoxSet*>(set));
}

char QBoxPlotSeries_Append2(void* ptr, void* sets)
{
	return static_cast<QBoxPlotSeries*>(ptr)->append(({ QList<QBoxSet *>* tmpP = static_cast<QList<QBoxSet *>*>(sets); QList<QBoxSet *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

char QBoxPlotSeries_BoxOutlineVisible(void* ptr)
{
	return static_cast<QBoxPlotSeries*>(ptr)->boxOutlineVisible();
}

char QBoxPlotSeries_Insert(void* ptr, int index, void* set)
{
	return static_cast<QBoxPlotSeries*>(ptr)->insert(index, static_cast<QBoxSet*>(set));
}

char QBoxPlotSeries_Remove(void* ptr, void* set)
{
	return static_cast<QBoxPlotSeries*>(ptr)->remove(static_cast<QBoxSet*>(set));
}

char QBoxPlotSeries_Take(void* ptr, void* set)
{
	return static_cast<QBoxPlotSeries*>(ptr)->take(static_cast<QBoxSet*>(set));
}

double QBoxPlotSeries_BoxWidth(void* ptr)
{
	return static_cast<QBoxPlotSeries*>(ptr)->boxWidth();
}

void QBoxPlotSeries_ConnectBoxOutlineVisibilityChanged(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::boxOutlineVisibilityChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_BoxOutlineVisibilityChanged));
}

void QBoxPlotSeries_DisconnectBoxOutlineVisibilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::boxOutlineVisibilityChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_BoxOutlineVisibilityChanged));
}

void QBoxPlotSeries_BoxOutlineVisibilityChanged(void* ptr)
{
	static_cast<QBoxPlotSeries*>(ptr)->boxOutlineVisibilityChanged();
}

void QBoxPlotSeries_ConnectBoxWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::boxWidthChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_BoxWidthChanged));
}

void QBoxPlotSeries_DisconnectBoxWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::boxWidthChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_BoxWidthChanged));
}

void QBoxPlotSeries_BoxWidthChanged(void* ptr)
{
	static_cast<QBoxPlotSeries*>(ptr)->boxWidthChanged();
}

void QBoxPlotSeries_ConnectBoxsetsAdded(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QList<QBoxSet *>)>(&QBoxPlotSeries::boxsetsAdded), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QList<QBoxSet *>)>(&MyQBoxPlotSeries::Signal_BoxsetsAdded));
}

void QBoxPlotSeries_DisconnectBoxsetsAdded(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QList<QBoxSet *>)>(&QBoxPlotSeries::boxsetsAdded), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QList<QBoxSet *>)>(&MyQBoxPlotSeries::Signal_BoxsetsAdded));
}

void QBoxPlotSeries_BoxsetsAdded(void* ptr, void* sets)
{
	static_cast<QBoxPlotSeries*>(ptr)->boxsetsAdded(({ QList<QBoxSet *>* tmpP = static_cast<QList<QBoxSet *>*>(sets); QList<QBoxSet *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QBoxPlotSeries_ConnectBoxsetsRemoved(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QList<QBoxSet *>)>(&QBoxPlotSeries::boxsetsRemoved), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QList<QBoxSet *>)>(&MyQBoxPlotSeries::Signal_BoxsetsRemoved));
}

void QBoxPlotSeries_DisconnectBoxsetsRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QList<QBoxSet *>)>(&QBoxPlotSeries::boxsetsRemoved), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QList<QBoxSet *>)>(&MyQBoxPlotSeries::Signal_BoxsetsRemoved));
}

void QBoxPlotSeries_BoxsetsRemoved(void* ptr, void* sets)
{
	static_cast<QBoxPlotSeries*>(ptr)->boxsetsRemoved(({ QList<QBoxSet *>* tmpP = static_cast<QList<QBoxSet *>*>(sets); QList<QBoxSet *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QBoxPlotSeries_ConnectBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::brushChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_BrushChanged));
}

void QBoxPlotSeries_DisconnectBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::brushChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_BrushChanged));
}

void QBoxPlotSeries_BrushChanged(void* ptr)
{
	static_cast<QBoxPlotSeries*>(ptr)->brushChanged();
}

void QBoxPlotSeries_Clear(void* ptr)
{
	static_cast<QBoxPlotSeries*>(ptr)->clear();
}

void QBoxPlotSeries_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QBoxSet *)>(&QBoxPlotSeries::clicked), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QBoxSet *)>(&MyQBoxPlotSeries::Signal_Clicked));
}

void QBoxPlotSeries_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QBoxSet *)>(&QBoxPlotSeries::clicked), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QBoxSet *)>(&MyQBoxPlotSeries::Signal_Clicked));
}

void QBoxPlotSeries_Clicked(void* ptr, void* boxset)
{
	static_cast<QBoxPlotSeries*>(ptr)->clicked(static_cast<QBoxSet*>(boxset));
}

void QBoxPlotSeries_ConnectCountChanged(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::countChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_CountChanged));
}

void QBoxPlotSeries_DisconnectCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::countChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_CountChanged));
}

void QBoxPlotSeries_CountChanged(void* ptr)
{
	static_cast<QBoxPlotSeries*>(ptr)->countChanged();
}

void QBoxPlotSeries_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QBoxSet *)>(&QBoxPlotSeries::doubleClicked), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QBoxSet *)>(&MyQBoxPlotSeries::Signal_DoubleClicked));
}

void QBoxPlotSeries_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QBoxSet *)>(&QBoxPlotSeries::doubleClicked), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QBoxSet *)>(&MyQBoxPlotSeries::Signal_DoubleClicked));
}

void QBoxPlotSeries_DoubleClicked(void* ptr, void* boxset)
{
	static_cast<QBoxPlotSeries*>(ptr)->doubleClicked(static_cast<QBoxSet*>(boxset));
}

void QBoxPlotSeries_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(bool, QBoxSet *)>(&QBoxPlotSeries::hovered), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(bool, QBoxSet *)>(&MyQBoxPlotSeries::Signal_Hovered));
}

void QBoxPlotSeries_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(bool, QBoxSet *)>(&QBoxPlotSeries::hovered), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(bool, QBoxSet *)>(&MyQBoxPlotSeries::Signal_Hovered));
}

void QBoxPlotSeries_Hovered(void* ptr, char status, void* boxset)
{
	static_cast<QBoxPlotSeries*>(ptr)->hovered(status != 0, static_cast<QBoxSet*>(boxset));
}

void QBoxPlotSeries_ConnectPenChanged(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::penChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_PenChanged));
}

void QBoxPlotSeries_DisconnectPenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)()>(&QBoxPlotSeries::penChanged), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)()>(&MyQBoxPlotSeries::Signal_PenChanged));
}

void QBoxPlotSeries_PenChanged(void* ptr)
{
	static_cast<QBoxPlotSeries*>(ptr)->penChanged();
}

void QBoxPlotSeries_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QBoxSet *)>(&QBoxPlotSeries::pressed), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QBoxSet *)>(&MyQBoxPlotSeries::Signal_Pressed));
}

void QBoxPlotSeries_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QBoxSet *)>(&QBoxPlotSeries::pressed), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QBoxSet *)>(&MyQBoxPlotSeries::Signal_Pressed));
}

void QBoxPlotSeries_Pressed(void* ptr, void* boxset)
{
	static_cast<QBoxPlotSeries*>(ptr)->pressed(static_cast<QBoxSet*>(boxset));
}

void QBoxPlotSeries_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QBoxSet *)>(&QBoxPlotSeries::released), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QBoxSet *)>(&MyQBoxPlotSeries::Signal_Released));
}

void QBoxPlotSeries_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QBoxPlotSeries*>(ptr), static_cast<void (QBoxPlotSeries::*)(QBoxSet *)>(&QBoxPlotSeries::released), static_cast<MyQBoxPlotSeries*>(ptr), static_cast<void (MyQBoxPlotSeries::*)(QBoxSet *)>(&MyQBoxPlotSeries::Signal_Released));
}

void QBoxPlotSeries_Released(void* ptr, void* boxset)
{
	static_cast<QBoxPlotSeries*>(ptr)->released(static_cast<QBoxSet*>(boxset));
}

void QBoxPlotSeries_SetBoxOutlineVisible(void* ptr, char visible)
{
	static_cast<QBoxPlotSeries*>(ptr)->setBoxOutlineVisible(visible != 0);
}

void QBoxPlotSeries_SetBoxWidth(void* ptr, double width)
{
	static_cast<QBoxPlotSeries*>(ptr)->setBoxWidth(width);
}

void QBoxPlotSeries_SetBrush(void* ptr, void* brush)
{
	static_cast<QBoxPlotSeries*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QBoxPlotSeries_SetPen(void* ptr, void* pen)
{
	static_cast<QBoxPlotSeries*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QBoxPlotSeries_DestroyQBoxPlotSeries(void* ptr)
{
	static_cast<QBoxPlotSeries*>(ptr)->~QBoxPlotSeries();
}

void QBoxPlotSeries_DestroyQBoxPlotSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QBoxPlotSeries_Type(void* ptr)
{
	return static_cast<QBoxPlotSeries*>(ptr)->type();
}

long long QBoxPlotSeries_TypeDefault(void* ptr)
{
		return static_cast<QBoxPlotSeries*>(ptr)->QBoxPlotSeries::type();
}

void* QBoxPlotSeries_Brush(void* ptr)
{
	return new QBrush(static_cast<QBoxPlotSeries*>(ptr)->brush());
}

struct QtCharts_PackedList QBoxPlotSeries_BoxSets(void* ptr)
{
	return ({ QList<QBoxSet *>* tmpValue = new QList<QBoxSet *>(static_cast<QBoxPlotSeries*>(ptr)->boxSets()); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

void* QBoxPlotSeries_Pen(void* ptr)
{
	return new QPen(static_cast<QBoxPlotSeries*>(ptr)->pen());
}

int QBoxPlotSeries_Count(void* ptr)
{
	return static_cast<QBoxPlotSeries*>(ptr)->count();
}

void* QBoxPlotSeries___append_sets_atList2(void* ptr, int i)
{
	return ({QBoxSet * tmp = static_cast<QList<QBoxSet *>*>(ptr)->at(i); if (i == static_cast<QList<QBoxSet *>*>(ptr)->size()-1) { static_cast<QList<QBoxSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBoxPlotSeries___append_sets_setList2(void* ptr, void* i)
{
	static_cast<QList<QBoxSet *>*>(ptr)->append(static_cast<QBoxSet*>(i));
}

void* QBoxPlotSeries___append_sets_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QBoxSet *>();
}

void* QBoxPlotSeries___boxsetsAdded_sets_atList(void* ptr, int i)
{
	return ({QBoxSet * tmp = static_cast<QList<QBoxSet *>*>(ptr)->at(i); if (i == static_cast<QList<QBoxSet *>*>(ptr)->size()-1) { static_cast<QList<QBoxSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBoxPlotSeries___boxsetsAdded_sets_setList(void* ptr, void* i)
{
	static_cast<QList<QBoxSet *>*>(ptr)->append(static_cast<QBoxSet*>(i));
}

void* QBoxPlotSeries___boxsetsAdded_sets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QBoxSet *>();
}

void* QBoxPlotSeries___boxsetsRemoved_sets_atList(void* ptr, int i)
{
	return ({QBoxSet * tmp = static_cast<QList<QBoxSet *>*>(ptr)->at(i); if (i == static_cast<QList<QBoxSet *>*>(ptr)->size()-1) { static_cast<QList<QBoxSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBoxPlotSeries___boxsetsRemoved_sets_setList(void* ptr, void* i)
{
	static_cast<QList<QBoxSet *>*>(ptr)->append(static_cast<QBoxSet*>(i));
}

void* QBoxPlotSeries___boxsetsRemoved_sets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QBoxSet *>();
}

void* QBoxPlotSeries___boxSets_atList(void* ptr, int i)
{
	return ({QBoxSet * tmp = static_cast<QList<QBoxSet *>*>(ptr)->at(i); if (i == static_cast<QList<QBoxSet *>*>(ptr)->size()-1) { static_cast<QList<QBoxSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBoxPlotSeries___boxSets_setList(void* ptr, void* i)
{
	static_cast<QList<QBoxSet *>*>(ptr)->append(static_cast<QBoxSet*>(i));
}

void* QBoxPlotSeries___boxSets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QBoxSet *>();
}

class MyQBoxSet: public QBoxSet
{
public:
	MyQBoxSet(const QString label = QString(), QObject *parent = Q_NULLPTR) : QBoxSet(label, parent) {QBoxSet_QBoxSet_QRegisterMetaType();};
	MyQBoxSet(const qreal le, const qreal lq, const qreal m, const qreal uq, const qreal ue, const QString label = QString(), QObject *parent = Q_NULLPTR) : QBoxSet(le, lq, m, uq, ue, label, parent) {QBoxSet_QBoxSet_QRegisterMetaType();};
	void Signal_BrushChanged() { callbackQBoxSet_BrushChanged(this); };
	void Signal_Cleared() { callbackQBoxSet_Cleared(this); };
	void Signal_Clicked() { callbackQBoxSet_Clicked(this); };
	void Signal_DoubleClicked() { callbackQBoxSet_DoubleClicked(this); };
	void Signal_Hovered(bool status) { callbackQBoxSet_Hovered(this, status); };
	void Signal_PenChanged() { callbackQBoxSet_PenChanged(this); };
	void Signal_Pressed() { callbackQBoxSet_Pressed(this); };
	void Signal_Released() { callbackQBoxSet_Released(this); };
	void Signal_ValueChanged(int index) { callbackQBoxSet_ValueChanged(this, index); };
	void Signal_ValuesChanged() { callbackQBoxSet_ValuesChanged(this); };
	 ~MyQBoxSet() { callbackQBoxSet_DestroyQBoxSet(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQBoxSet_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQBoxSet_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQBoxSet_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQBoxSet_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQBoxSet_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQBoxSet_CustomEvent(this, event); };
	void deleteLater() { callbackQBoxSet_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQBoxSet_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQBoxSet_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQBoxSet_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQBoxSet_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQBoxSet*)

int QBoxSet_QBoxSet_QRegisterMetaType(){qRegisterMetaType<QBoxSet*>(); return qRegisterMetaType<MyQBoxSet*>();}

void* QBoxSet_NewQBoxSet(struct QtCharts_PackedString label, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQBoxSet(QString::fromUtf8(label.data, label.len), static_cast<QObject*>(parent));
	}
}

void* QBoxSet_NewQBoxSet2(double le, double lq, double m, double uq, double ue, struct QtCharts_PackedString label, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQBoxSet(le, lq, m, uq, ue, QString::fromUtf8(label.data, label.len), static_cast<QObject*>(parent));
	}
}

struct QtCharts_PackedString QBoxSet_QBoxSet_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t792b87 = QBoxSet::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t792b87.prepend("WHITESPACE").constData()+10), t792b87.size()-10 }; });
}

struct QtCharts_PackedString QBoxSet_QBoxSet_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t4739cc = QBoxSet::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t4739cc.prepend("WHITESPACE").constData()+10), t4739cc.size()-10 }; });
}

void QBoxSet_Append2(void* ptr, void* values)
{
	static_cast<QBoxSet*>(ptr)->append(*static_cast<QList<qreal>*>(values));
}

void QBoxSet_Append(void* ptr, double value)
{
	static_cast<QBoxSet*>(ptr)->append(value);
}

void QBoxSet_ConnectBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::brushChanged), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_BrushChanged));
}

void QBoxSet_DisconnectBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::brushChanged), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_BrushChanged));
}

void QBoxSet_BrushChanged(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->brushChanged();
}

void QBoxSet_Clear(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->clear();
}

void QBoxSet_ConnectCleared(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::cleared), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_Cleared));
}

void QBoxSet_DisconnectCleared(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::cleared), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_Cleared));
}

void QBoxSet_Cleared(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->cleared();
}

void QBoxSet_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::clicked), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_Clicked));
}

void QBoxSet_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::clicked), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_Clicked));
}

void QBoxSet_Clicked(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->clicked();
}

void QBoxSet_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::doubleClicked), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_DoubleClicked));
}

void QBoxSet_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::doubleClicked), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_DoubleClicked));
}

void QBoxSet_DoubleClicked(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->doubleClicked();
}

void QBoxSet_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)(bool)>(&QBoxSet::hovered), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)(bool)>(&MyQBoxSet::Signal_Hovered));
}

void QBoxSet_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)(bool)>(&QBoxSet::hovered), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)(bool)>(&MyQBoxSet::Signal_Hovered));
}

void QBoxSet_Hovered(void* ptr, char status)
{
	static_cast<QBoxSet*>(ptr)->hovered(status != 0);
}

void QBoxSet_ConnectPenChanged(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::penChanged), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_PenChanged));
}

void QBoxSet_DisconnectPenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::penChanged), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_PenChanged));
}

void QBoxSet_PenChanged(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->penChanged();
}

void QBoxSet_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::pressed), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_Pressed));
}

void QBoxSet_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::pressed), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_Pressed));
}

void QBoxSet_Pressed(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->pressed();
}

void QBoxSet_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::released), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_Released));
}

void QBoxSet_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::released), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_Released));
}

void QBoxSet_Released(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->released();
}

void QBoxSet_SetBrush(void* ptr, void* brush)
{
	static_cast<QBoxSet*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QBoxSet_SetLabel(void* ptr, struct QtCharts_PackedString label)
{
	static_cast<QBoxSet*>(ptr)->setLabel(QString::fromUtf8(label.data, label.len));
}

void QBoxSet_SetPen(void* ptr, void* pen)
{
	static_cast<QBoxSet*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QBoxSet_SetValue(void* ptr, int index, double value)
{
	static_cast<QBoxSet*>(ptr)->setValue(index, value);
}

void QBoxSet_ConnectValueChanged(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)(int)>(&QBoxSet::valueChanged), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)(int)>(&MyQBoxSet::Signal_ValueChanged));
}

void QBoxSet_DisconnectValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)(int)>(&QBoxSet::valueChanged), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)(int)>(&MyQBoxSet::Signal_ValueChanged));
}

void QBoxSet_ValueChanged(void* ptr, int index)
{
	static_cast<QBoxSet*>(ptr)->valueChanged(index);
}

void QBoxSet_ConnectValuesChanged(void* ptr)
{
	QObject::connect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::valuesChanged), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_ValuesChanged));
}

void QBoxSet_DisconnectValuesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBoxSet*>(ptr), static_cast<void (QBoxSet::*)()>(&QBoxSet::valuesChanged), static_cast<MyQBoxSet*>(ptr), static_cast<void (MyQBoxSet::*)()>(&MyQBoxSet::Signal_ValuesChanged));
}

void QBoxSet_ValuesChanged(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->valuesChanged();
}

void QBoxSet_DestroyQBoxSet(void* ptr)
{
	static_cast<QBoxSet*>(ptr)->~QBoxSet();
}

void QBoxSet_DestroyQBoxSetDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QBoxSet_Brush(void* ptr)
{
	return new QBrush(static_cast<QBoxSet*>(ptr)->brush());
}

void* QBoxSet_Pen(void* ptr)
{
	return new QPen(static_cast<QBoxSet*>(ptr)->pen());
}

struct QtCharts_PackedString QBoxSet_Label(void* ptr)
{
	return ({ QByteArray tfb4aff = static_cast<QBoxSet*>(ptr)->label().toUtf8(); QtCharts_PackedString { const_cast<char*>(tfb4aff.prepend("WHITESPACE").constData()+10), tfb4aff.size()-10 }; });
}

void* QBoxSet_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QBoxSet*>(ptr)->QBoxSet::metaObject());
}

int QBoxSet_Count(void* ptr)
{
	return static_cast<QBoxSet*>(ptr)->count();
}

double QBoxSet_At(void* ptr, int index)
{
	return static_cast<QBoxSet*>(ptr)->at(index);
}

double QBoxSet___append_values_atList2(void* ptr, int i)
{
	return ({qreal tmp = static_cast<QList<qreal>*>(ptr)->at(i); if (i == static_cast<QList<qreal>*>(ptr)->size()-1) { static_cast<QList<qreal>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBoxSet___append_values_setList2(void* ptr, double i)
{
	static_cast<QList<qreal>*>(ptr)->append(i);
}

void* QBoxSet___append_values_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qreal>();
}

void* QBoxSet___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QBoxSet___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QBoxSet___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QBoxSet___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBoxSet___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QBoxSet___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QBoxSet___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBoxSet___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QBoxSet___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QBoxSet___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBoxSet___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QBoxSet___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QBoxSet___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QBoxSet___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QBoxSet___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QBoxSet_EventDefault(void* ptr, void* e)
{
		return static_cast<QBoxSet*>(ptr)->QBoxSet::event(static_cast<QEvent*>(e));
}

char QBoxSet_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QBoxSet*>(ptr)->QBoxSet::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QBoxSet_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QBoxSet*>(ptr)->QBoxSet::childEvent(static_cast<QChildEvent*>(event));
}

void QBoxSet_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QBoxSet*>(ptr)->QBoxSet::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBoxSet_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QBoxSet*>(ptr)->QBoxSet::customEvent(static_cast<QEvent*>(event));
}

void QBoxSet_DeleteLaterDefault(void* ptr)
{
		static_cast<QBoxSet*>(ptr)->QBoxSet::deleteLater();
}

void QBoxSet_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QBoxSet*>(ptr)->QBoxSet::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QBoxSet_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QBoxSet*>(ptr)->QBoxSet::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQCandlestickLegendMarker: public QCandlestickLegendMarker
{
public:
	QCandlestickSeries * series() { return static_cast<QCandlestickSeries*>(callbackQCandlestickLegendMarker_Series(this)); };
	QLegendMarker::LegendMarkerType type() { return static_cast<QLegendMarker::LegendMarkerType>(callbackQCandlestickLegendMarker_Type(this)); };
	 ~MyQCandlestickLegendMarker() { callbackQCandlestickLegendMarker_DestroyQCandlestickLegendMarker(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLegendMarker_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BrushChanged() { callbackQLegendMarker_BrushChanged(this); };
	void Signal_Clicked() { callbackQLegendMarker_Clicked(this); };
	void Signal_FontChanged() { callbackQLegendMarker_FontChanged(this); };
	void Signal_Hovered(bool status) { callbackQLegendMarker_Hovered(this, status); };
	void Signal_LabelBrushChanged() { callbackQLegendMarker_LabelBrushChanged(this); };
	void Signal_LabelChanged() { callbackQLegendMarker_LabelChanged(this); };
	void Signal_PenChanged() { callbackQLegendMarker_PenChanged(this); };
	void Signal_ShapeChanged() { callbackQLegendMarker_ShapeChanged(this); };
	void Signal_VisibleChanged() { callbackQLegendMarker_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQLegendMarker_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLegendMarker_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLegendMarker_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLegendMarker_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLegendMarker_CustomEvent(this, event); };
	void deleteLater() { callbackQLegendMarker_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLegendMarker_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLegendMarker_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLegendMarker_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLegendMarker_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQCandlestickLegendMarker*)

int QCandlestickLegendMarker_QCandlestickLegendMarker_QRegisterMetaType(){qRegisterMetaType<QCandlestickLegendMarker*>(); return qRegisterMetaType<MyQCandlestickLegendMarker*>();}

void* QCandlestickLegendMarker_Series(void* ptr)
{
	return static_cast<QCandlestickLegendMarker*>(ptr)->series();
}

void* QCandlestickLegendMarker_SeriesDefault(void* ptr)
{
		return static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::series();
}

long long QCandlestickLegendMarker_Type(void* ptr)
{
	return static_cast<QCandlestickLegendMarker*>(ptr)->type();
}

long long QCandlestickLegendMarker_TypeDefault(void* ptr)
{
		return static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::type();
}

void QCandlestickLegendMarker_DestroyQCandlestickLegendMarker(void* ptr)
{
	static_cast<QCandlestickLegendMarker*>(ptr)->~QCandlestickLegendMarker();
}

void QCandlestickLegendMarker_DestroyQCandlestickLegendMarkerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQCandlestickModelMapper: public QCandlestickModelMapper
{
public:
	MyQCandlestickModelMapper(QObject *parent = Q_NULLPTR) : QCandlestickModelMapper(parent) {QCandlestickModelMapper_QCandlestickModelMapper_QRegisterMetaType();};
	void Signal_ModelReplaced() { callbackQCandlestickModelMapper_ModelReplaced(this); };
	void Signal_SeriesReplaced() { callbackQCandlestickModelMapper_SeriesReplaced(this); };
	Qt::Orientation orientation() const { return static_cast<Qt::Orientation>(callbackQCandlestickModelMapper_Orientation(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCandlestickModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQCandlestickModelMapper_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCandlestickModelMapper_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQCandlestickModelMapper_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCandlestickModelMapper_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCandlestickModelMapper_CustomEvent(this, event); };
	void deleteLater() { callbackQCandlestickModelMapper_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQCandlestickModelMapper_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCandlestickModelMapper_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQCandlestickModelMapper_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQCandlestickModelMapper_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQCandlestickModelMapper*)

int QCandlestickModelMapper_QCandlestickModelMapper_QRegisterMetaType(){qRegisterMetaType<QCandlestickModelMapper*>(); return qRegisterMetaType<MyQCandlestickModelMapper*>();}

void* QCandlestickModelMapper_NewQCandlestickModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQCandlestickModelMapper(static_cast<QObject*>(parent));
	}
}

struct QtCharts_PackedString QCandlestickModelMapper_QCandlestickModelMapper_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t4b5918 = QCandlestickModelMapper::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t4b5918.prepend("WHITESPACE").constData()+10), t4b5918.size()-10 }; });
}

struct QtCharts_PackedString QCandlestickModelMapper_QCandlestickModelMapper_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t19ba4e = QCandlestickModelMapper::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t19ba4e.prepend("WHITESPACE").constData()+10), t19ba4e.size()-10 }; });
}

void QCandlestickModelMapper_ConnectModelReplaced(void* ptr)
{
	QObject::connect(static_cast<QCandlestickModelMapper*>(ptr), static_cast<void (QCandlestickModelMapper::*)()>(&QCandlestickModelMapper::modelReplaced), static_cast<MyQCandlestickModelMapper*>(ptr), static_cast<void (MyQCandlestickModelMapper::*)()>(&MyQCandlestickModelMapper::Signal_ModelReplaced));
}

void QCandlestickModelMapper_DisconnectModelReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickModelMapper*>(ptr), static_cast<void (QCandlestickModelMapper::*)()>(&QCandlestickModelMapper::modelReplaced), static_cast<MyQCandlestickModelMapper*>(ptr), static_cast<void (MyQCandlestickModelMapper::*)()>(&MyQCandlestickModelMapper::Signal_ModelReplaced));
}

void QCandlestickModelMapper_ModelReplaced(void* ptr)
{
	static_cast<QCandlestickModelMapper*>(ptr)->modelReplaced();
}

void QCandlestickModelMapper_ConnectSeriesReplaced(void* ptr)
{
	QObject::connect(static_cast<QCandlestickModelMapper*>(ptr), static_cast<void (QCandlestickModelMapper::*)()>(&QCandlestickModelMapper::seriesReplaced), static_cast<MyQCandlestickModelMapper*>(ptr), static_cast<void (MyQCandlestickModelMapper::*)()>(&MyQCandlestickModelMapper::Signal_SeriesReplaced));
}

void QCandlestickModelMapper_DisconnectSeriesReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickModelMapper*>(ptr), static_cast<void (QCandlestickModelMapper::*)()>(&QCandlestickModelMapper::seriesReplaced), static_cast<MyQCandlestickModelMapper*>(ptr), static_cast<void (MyQCandlestickModelMapper::*)()>(&MyQCandlestickModelMapper::Signal_SeriesReplaced));
}

void QCandlestickModelMapper_SeriesReplaced(void* ptr)
{
	static_cast<QCandlestickModelMapper*>(ptr)->seriesReplaced();
}

void QCandlestickModelMapper_SetClose(void* ptr, int close)
{
	static_cast<QCandlestickModelMapper*>(ptr)->setClose(close);
}

void QCandlestickModelMapper_SetFirstSetSection(void* ptr, int firstSetSection)
{
	static_cast<QCandlestickModelMapper*>(ptr)->setFirstSetSection(firstSetSection);
}

void QCandlestickModelMapper_SetHigh(void* ptr, int high)
{
	static_cast<QCandlestickModelMapper*>(ptr)->setHigh(high);
}

void QCandlestickModelMapper_SetLastSetSection(void* ptr, int lastSetSection)
{
	static_cast<QCandlestickModelMapper*>(ptr)->setLastSetSection(lastSetSection);
}

void QCandlestickModelMapper_SetLow(void* ptr, int low)
{
	static_cast<QCandlestickModelMapper*>(ptr)->setLow(low);
}

void QCandlestickModelMapper_SetModel(void* ptr, void* model)
{
	static_cast<QCandlestickModelMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QCandlestickModelMapper_SetOpen(void* ptr, int open)
{
	static_cast<QCandlestickModelMapper*>(ptr)->setOpen(open);
}

void QCandlestickModelMapper_SetSeries(void* ptr, void* series)
{
	static_cast<QCandlestickModelMapper*>(ptr)->setSeries(static_cast<QCandlestickSeries*>(series));
}

void QCandlestickModelMapper_SetTimestamp(void* ptr, int timestamp)
{
	static_cast<QCandlestickModelMapper*>(ptr)->setTimestamp(timestamp);
}

void* QCandlestickModelMapper_Model(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->model();
}

void* QCandlestickModelMapper_Series(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->series();
}

long long QCandlestickModelMapper_Orientation(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->orientation();
}

void* QCandlestickModelMapper_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QVCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::metaObject());
	} else if (dynamic_cast<QHCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QCandlestickModelMapper*>(ptr)->QCandlestickModelMapper::metaObject());
	}
}

int QCandlestickModelMapper_Close(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->close();
}

int QCandlestickModelMapper_FirstSetSection(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->firstSetSection();
}

int QCandlestickModelMapper_High(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->high();
}

int QCandlestickModelMapper_LastSetSection(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->lastSetSection();
}

int QCandlestickModelMapper_Low(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->low();
}

int QCandlestickModelMapper_Open(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->open();
}

int QCandlestickModelMapper_Timestamp(void* ptr)
{
	return static_cast<QCandlestickModelMapper*>(ptr)->timestamp();
}

void* QCandlestickModelMapper___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QCandlestickModelMapper___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QCandlestickModelMapper___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QCandlestickModelMapper___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickModelMapper___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCandlestickModelMapper___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QCandlestickModelMapper___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickModelMapper___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCandlestickModelMapper___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QCandlestickModelMapper___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickModelMapper___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCandlestickModelMapper___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QCandlestickModelMapper___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickModelMapper___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCandlestickModelMapper___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QCandlestickModelMapper_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QVCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QHCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QCandlestickModelMapper*>(ptr)->QCandlestickModelMapper::event(static_cast<QEvent*>(e));
	}
}

char QCandlestickModelMapper_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QVCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		return static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QCandlestickModelMapper*>(ptr)->QCandlestickModelMapper::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QCandlestickModelMapper_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QVCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QHCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QCandlestickModelMapper*>(ptr)->QCandlestickModelMapper::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QCandlestickModelMapper_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QVCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QCandlestickModelMapper*>(ptr)->QCandlestickModelMapper::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QCandlestickModelMapper_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QVCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QCandlestickModelMapper*>(ptr)->QCandlestickModelMapper::customEvent(static_cast<QEvent*>(event));
	}
}

void QCandlestickModelMapper_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QVCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::deleteLater();
	} else if (dynamic_cast<QHCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::deleteLater();
	} else {
		static_cast<QCandlestickModelMapper*>(ptr)->QCandlestickModelMapper::deleteLater();
	}
}

void QCandlestickModelMapper_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QVCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QCandlestickModelMapper*>(ptr)->QCandlestickModelMapper::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QCandlestickModelMapper_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QVCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHCandlestickModelMapper*>(static_cast<QObject*>(ptr))) {
		static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QCandlestickModelMapper*>(ptr)->QCandlestickModelMapper::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

class MyQCandlestickSeries: public QCandlestickSeries
{
public:
	MyQCandlestickSeries(QObject *parent = Q_NULLPTR) : QCandlestickSeries(parent) {QCandlestickSeries_QCandlestickSeries_QRegisterMetaType();};
	void Signal_BodyOutlineVisibilityChanged() { callbackQCandlestickSeries_BodyOutlineVisibilityChanged(this); };
	void Signal_BodyWidthChanged() { callbackQCandlestickSeries_BodyWidthChanged(this); };
	void Signal_BrushChanged() { callbackQCandlestickSeries_BrushChanged(this); };
	void Signal_CandlestickSetsAdded(const QList<QCandlestickSet *> & sets) { callbackQCandlestickSeries_CandlestickSetsAdded(this, ({ QList<QCandlestickSet *>* tmpValue = const_cast<QList<QCandlestickSet *>*>(&sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_CandlestickSetsRemoved(const QList<QCandlestickSet *> & sets) { callbackQCandlestickSeries_CandlestickSetsRemoved(this, ({ QList<QCandlestickSet *>* tmpValue = const_cast<QList<QCandlestickSet *>*>(&sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_CapsVisibilityChanged() { callbackQCandlestickSeries_CapsVisibilityChanged(this); };
	void Signal_CapsWidthChanged() { callbackQCandlestickSeries_CapsWidthChanged(this); };
	void Signal_Clicked(QCandlestickSet * set) { callbackQCandlestickSeries_Clicked(this, set); };
	void Signal_CountChanged() { callbackQCandlestickSeries_CountChanged(this); };
	void Signal_DecreasingColorChanged() { callbackQCandlestickSeries_DecreasingColorChanged(this); };
	void Signal_DoubleClicked(QCandlestickSet * set) { callbackQCandlestickSeries_DoubleClicked(this, set); };
	void Signal_Hovered(bool status, QCandlestickSet * set) { callbackQCandlestickSeries_Hovered(this, status, set); };
	void Signal_IncreasingColorChanged() { callbackQCandlestickSeries_IncreasingColorChanged(this); };
	void Signal_MaximumColumnWidthChanged() { callbackQCandlestickSeries_MaximumColumnWidthChanged(this); };
	void Signal_MinimumColumnWidthChanged() { callbackQCandlestickSeries_MinimumColumnWidthChanged(this); };
	void Signal_PenChanged() { callbackQCandlestickSeries_PenChanged(this); };
	void Signal_Pressed(QCandlestickSet * set) { callbackQCandlestickSeries_Pressed(this, set); };
	void Signal_Released(QCandlestickSet * set) { callbackQCandlestickSeries_Released(this, set); };
	 ~MyQCandlestickSeries() { callbackQCandlestickSeries_DestroyQCandlestickSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQCandlestickSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQCandlestickSeries*)

int QCandlestickSeries_QCandlestickSeries_QRegisterMetaType(){qRegisterMetaType<QCandlestickSeries*>(); return qRegisterMetaType<MyQCandlestickSeries*>();}

void* QCandlestickSeries_NewQCandlestickSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQCandlestickSeries(static_cast<QObject*>(parent));
	}
}

char QCandlestickSeries_Append(void* ptr, void* set)
{
	return static_cast<QCandlestickSeries*>(ptr)->append(static_cast<QCandlestickSet*>(set));
}

char QCandlestickSeries_Append2(void* ptr, void* sets)
{
	return static_cast<QCandlestickSeries*>(ptr)->append(*static_cast<QList<QCandlestickSet *>*>(sets));
}

char QCandlestickSeries_Insert(void* ptr, int index, void* set)
{
	return static_cast<QCandlestickSeries*>(ptr)->insert(index, static_cast<QCandlestickSet*>(set));
}

char QCandlestickSeries_Remove(void* ptr, void* set)
{
	return static_cast<QCandlestickSeries*>(ptr)->remove(static_cast<QCandlestickSet*>(set));
}

char QCandlestickSeries_Remove2(void* ptr, void* sets)
{
	return static_cast<QCandlestickSeries*>(ptr)->remove(*static_cast<QList<QCandlestickSet *>*>(sets));
}

char QCandlestickSeries_Take(void* ptr, void* set)
{
	return static_cast<QCandlestickSeries*>(ptr)->take(static_cast<QCandlestickSet*>(set));
}

void QCandlestickSeries_ConnectBodyOutlineVisibilityChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::bodyOutlineVisibilityChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_BodyOutlineVisibilityChanged));
}

void QCandlestickSeries_DisconnectBodyOutlineVisibilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::bodyOutlineVisibilityChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_BodyOutlineVisibilityChanged));
}

void QCandlestickSeries_BodyOutlineVisibilityChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->bodyOutlineVisibilityChanged();
}

void QCandlestickSeries_ConnectBodyWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::bodyWidthChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_BodyWidthChanged));
}

void QCandlestickSeries_DisconnectBodyWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::bodyWidthChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_BodyWidthChanged));
}

void QCandlestickSeries_BodyWidthChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->bodyWidthChanged();
}

void QCandlestickSeries_ConnectBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::brushChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_BrushChanged));
}

void QCandlestickSeries_DisconnectBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::brushChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_BrushChanged));
}

void QCandlestickSeries_BrushChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->brushChanged();
}

void QCandlestickSeries_ConnectCandlestickSetsAdded(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(const QList<QCandlestickSet *> &)>(&QCandlestickSeries::candlestickSetsAdded), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(const QList<QCandlestickSet *> &)>(&MyQCandlestickSeries::Signal_CandlestickSetsAdded));
}

void QCandlestickSeries_DisconnectCandlestickSetsAdded(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(const QList<QCandlestickSet *> &)>(&QCandlestickSeries::candlestickSetsAdded), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(const QList<QCandlestickSet *> &)>(&MyQCandlestickSeries::Signal_CandlestickSetsAdded));
}

void QCandlestickSeries_CandlestickSetsAdded(void* ptr, void* sets)
{
	static_cast<QCandlestickSeries*>(ptr)->candlestickSetsAdded(*static_cast<QList<QCandlestickSet *>*>(sets));
}

void QCandlestickSeries_ConnectCandlestickSetsRemoved(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(const QList<QCandlestickSet *> &)>(&QCandlestickSeries::candlestickSetsRemoved), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(const QList<QCandlestickSet *> &)>(&MyQCandlestickSeries::Signal_CandlestickSetsRemoved));
}

void QCandlestickSeries_DisconnectCandlestickSetsRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(const QList<QCandlestickSet *> &)>(&QCandlestickSeries::candlestickSetsRemoved), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(const QList<QCandlestickSet *> &)>(&MyQCandlestickSeries::Signal_CandlestickSetsRemoved));
}

void QCandlestickSeries_CandlestickSetsRemoved(void* ptr, void* sets)
{
	static_cast<QCandlestickSeries*>(ptr)->candlestickSetsRemoved(*static_cast<QList<QCandlestickSet *>*>(sets));
}

void QCandlestickSeries_ConnectCapsVisibilityChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::capsVisibilityChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_CapsVisibilityChanged));
}

void QCandlestickSeries_DisconnectCapsVisibilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::capsVisibilityChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_CapsVisibilityChanged));
}

void QCandlestickSeries_CapsVisibilityChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->capsVisibilityChanged();
}

void QCandlestickSeries_ConnectCapsWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::capsWidthChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_CapsWidthChanged));
}

void QCandlestickSeries_DisconnectCapsWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::capsWidthChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_CapsWidthChanged));
}

void QCandlestickSeries_CapsWidthChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->capsWidthChanged();
}

void QCandlestickSeries_Clear(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->clear();
}

void QCandlestickSeries_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(QCandlestickSet *)>(&QCandlestickSeries::clicked), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(QCandlestickSet *)>(&MyQCandlestickSeries::Signal_Clicked));
}

void QCandlestickSeries_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(QCandlestickSet *)>(&QCandlestickSeries::clicked), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(QCandlestickSet *)>(&MyQCandlestickSeries::Signal_Clicked));
}

void QCandlestickSeries_Clicked(void* ptr, void* set)
{
	static_cast<QCandlestickSeries*>(ptr)->clicked(static_cast<QCandlestickSet*>(set));
}

void QCandlestickSeries_ConnectCountChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::countChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_CountChanged));
}

void QCandlestickSeries_DisconnectCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::countChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_CountChanged));
}

void QCandlestickSeries_CountChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->countChanged();
}

void QCandlestickSeries_ConnectDecreasingColorChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::decreasingColorChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_DecreasingColorChanged));
}

void QCandlestickSeries_DisconnectDecreasingColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::decreasingColorChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_DecreasingColorChanged));
}

void QCandlestickSeries_DecreasingColorChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->decreasingColorChanged();
}

void QCandlestickSeries_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(QCandlestickSet *)>(&QCandlestickSeries::doubleClicked), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(QCandlestickSet *)>(&MyQCandlestickSeries::Signal_DoubleClicked));
}

void QCandlestickSeries_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(QCandlestickSet *)>(&QCandlestickSeries::doubleClicked), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(QCandlestickSet *)>(&MyQCandlestickSeries::Signal_DoubleClicked));
}

void QCandlestickSeries_DoubleClicked(void* ptr, void* set)
{
	static_cast<QCandlestickSeries*>(ptr)->doubleClicked(static_cast<QCandlestickSet*>(set));
}

void QCandlestickSeries_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(bool, QCandlestickSet *)>(&QCandlestickSeries::hovered), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(bool, QCandlestickSet *)>(&MyQCandlestickSeries::Signal_Hovered));
}

void QCandlestickSeries_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(bool, QCandlestickSet *)>(&QCandlestickSeries::hovered), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(bool, QCandlestickSet *)>(&MyQCandlestickSeries::Signal_Hovered));
}

void QCandlestickSeries_Hovered(void* ptr, char status, void* set)
{
	static_cast<QCandlestickSeries*>(ptr)->hovered(status != 0, static_cast<QCandlestickSet*>(set));
}

void QCandlestickSeries_ConnectIncreasingColorChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::increasingColorChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_IncreasingColorChanged));
}

void QCandlestickSeries_DisconnectIncreasingColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::increasingColorChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_IncreasingColorChanged));
}

void QCandlestickSeries_IncreasingColorChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->increasingColorChanged();
}

void QCandlestickSeries_ConnectMaximumColumnWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::maximumColumnWidthChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_MaximumColumnWidthChanged));
}

void QCandlestickSeries_DisconnectMaximumColumnWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::maximumColumnWidthChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_MaximumColumnWidthChanged));
}

void QCandlestickSeries_MaximumColumnWidthChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->maximumColumnWidthChanged();
}

void QCandlestickSeries_ConnectMinimumColumnWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::minimumColumnWidthChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_MinimumColumnWidthChanged));
}

void QCandlestickSeries_DisconnectMinimumColumnWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::minimumColumnWidthChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_MinimumColumnWidthChanged));
}

void QCandlestickSeries_MinimumColumnWidthChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->minimumColumnWidthChanged();
}

void QCandlestickSeries_ConnectPenChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::penChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_PenChanged));
}

void QCandlestickSeries_DisconnectPenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)()>(&QCandlestickSeries::penChanged), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)()>(&MyQCandlestickSeries::Signal_PenChanged));
}

void QCandlestickSeries_PenChanged(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->penChanged();
}

void QCandlestickSeries_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(QCandlestickSet *)>(&QCandlestickSeries::pressed), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(QCandlestickSet *)>(&MyQCandlestickSeries::Signal_Pressed));
}

void QCandlestickSeries_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(QCandlestickSet *)>(&QCandlestickSeries::pressed), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(QCandlestickSet *)>(&MyQCandlestickSeries::Signal_Pressed));
}

void QCandlestickSeries_Pressed(void* ptr, void* set)
{
	static_cast<QCandlestickSeries*>(ptr)->pressed(static_cast<QCandlestickSet*>(set));
}

void QCandlestickSeries_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(QCandlestickSet *)>(&QCandlestickSeries::released), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(QCandlestickSet *)>(&MyQCandlestickSeries::Signal_Released));
}

void QCandlestickSeries_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSeries*>(ptr), static_cast<void (QCandlestickSeries::*)(QCandlestickSet *)>(&QCandlestickSeries::released), static_cast<MyQCandlestickSeries*>(ptr), static_cast<void (MyQCandlestickSeries::*)(QCandlestickSet *)>(&MyQCandlestickSeries::Signal_Released));
}

void QCandlestickSeries_Released(void* ptr, void* set)
{
	static_cast<QCandlestickSeries*>(ptr)->released(static_cast<QCandlestickSet*>(set));
}

void QCandlestickSeries_SetBodyOutlineVisible(void* ptr, char bodyOutlineVisible)
{
	static_cast<QCandlestickSeries*>(ptr)->setBodyOutlineVisible(bodyOutlineVisible != 0);
}

void QCandlestickSeries_SetBodyWidth(void* ptr, double bodyWidth)
{
	static_cast<QCandlestickSeries*>(ptr)->setBodyWidth(bodyWidth);
}

void QCandlestickSeries_SetBrush(void* ptr, void* brush)
{
	static_cast<QCandlestickSeries*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QCandlestickSeries_SetCapsVisible(void* ptr, char capsVisible)
{
	static_cast<QCandlestickSeries*>(ptr)->setCapsVisible(capsVisible != 0);
}

void QCandlestickSeries_SetCapsWidth(void* ptr, double capsWidth)
{
	static_cast<QCandlestickSeries*>(ptr)->setCapsWidth(capsWidth);
}

void QCandlestickSeries_SetDecreasingColor(void* ptr, void* decreasingColor)
{
	static_cast<QCandlestickSeries*>(ptr)->setDecreasingColor(*static_cast<QColor*>(decreasingColor));
}

void QCandlestickSeries_SetIncreasingColor(void* ptr, void* increasingColor)
{
	static_cast<QCandlestickSeries*>(ptr)->setIncreasingColor(*static_cast<QColor*>(increasingColor));
}

void QCandlestickSeries_SetMaximumColumnWidth(void* ptr, double maximumColumnWidth)
{
	static_cast<QCandlestickSeries*>(ptr)->setMaximumColumnWidth(maximumColumnWidth);
}

void QCandlestickSeries_SetMinimumColumnWidth(void* ptr, double minimumColumnWidth)
{
	static_cast<QCandlestickSeries*>(ptr)->setMinimumColumnWidth(minimumColumnWidth);
}

void QCandlestickSeries_SetPen(void* ptr, void* pen)
{
	static_cast<QCandlestickSeries*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QCandlestickSeries_DestroyQCandlestickSeries(void* ptr)
{
	static_cast<QCandlestickSeries*>(ptr)->~QCandlestickSeries();
}

void QCandlestickSeries_DestroyQCandlestickSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QCandlestickSeries_Type(void* ptr)
{
	return static_cast<QCandlestickSeries*>(ptr)->type();
}

long long QCandlestickSeries_TypeDefault(void* ptr)
{
		return static_cast<QCandlestickSeries*>(ptr)->QCandlestickSeries::type();
}

void* QCandlestickSeries_Brush(void* ptr)
{
	return new QBrush(static_cast<QCandlestickSeries*>(ptr)->brush());
}

void* QCandlestickSeries_DecreasingColor(void* ptr)
{
	return new QColor(static_cast<QCandlestickSeries*>(ptr)->decreasingColor());
}

void* QCandlestickSeries_IncreasingColor(void* ptr)
{
	return new QColor(static_cast<QCandlestickSeries*>(ptr)->increasingColor());
}

struct QtCharts_PackedList QCandlestickSeries_Sets(void* ptr)
{
	return ({ QList<QCandlestickSet *>* tmpValue = new QList<QCandlestickSet *>(static_cast<QCandlestickSeries*>(ptr)->sets()); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

void* QCandlestickSeries_Pen(void* ptr)
{
	return new QPen(static_cast<QCandlestickSeries*>(ptr)->pen());
}

char QCandlestickSeries_BodyOutlineVisible(void* ptr)
{
	return static_cast<QCandlestickSeries*>(ptr)->bodyOutlineVisible();
}

char QCandlestickSeries_CapsVisible(void* ptr)
{
	return static_cast<QCandlestickSeries*>(ptr)->capsVisible();
}

int QCandlestickSeries_Count(void* ptr)
{
	return static_cast<QCandlestickSeries*>(ptr)->count();
}

double QCandlestickSeries_BodyWidth(void* ptr)
{
	return static_cast<QCandlestickSeries*>(ptr)->bodyWidth();
}

double QCandlestickSeries_CapsWidth(void* ptr)
{
	return static_cast<QCandlestickSeries*>(ptr)->capsWidth();
}

double QCandlestickSeries_MaximumColumnWidth(void* ptr)
{
	return static_cast<QCandlestickSeries*>(ptr)->maximumColumnWidth();
}

double QCandlestickSeries_MinimumColumnWidth(void* ptr)
{
	return static_cast<QCandlestickSeries*>(ptr)->minimumColumnWidth();
}

void* QCandlestickSeries___append_sets_atList2(void* ptr, int i)
{
	return ({QCandlestickSet * tmp = static_cast<QList<QCandlestickSet *>*>(ptr)->at(i); if (i == static_cast<QList<QCandlestickSet *>*>(ptr)->size()-1) { static_cast<QList<QCandlestickSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickSeries___append_sets_setList2(void* ptr, void* i)
{
	static_cast<QList<QCandlestickSet *>*>(ptr)->append(static_cast<QCandlestickSet*>(i));
}

void* QCandlestickSeries___append_sets_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCandlestickSet *>();
}

void* QCandlestickSeries___remove_sets_atList2(void* ptr, int i)
{
	return ({QCandlestickSet * tmp = static_cast<QList<QCandlestickSet *>*>(ptr)->at(i); if (i == static_cast<QList<QCandlestickSet *>*>(ptr)->size()-1) { static_cast<QList<QCandlestickSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickSeries___remove_sets_setList2(void* ptr, void* i)
{
	static_cast<QList<QCandlestickSet *>*>(ptr)->append(static_cast<QCandlestickSet*>(i));
}

void* QCandlestickSeries___remove_sets_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCandlestickSet *>();
}

void* QCandlestickSeries___candlestickSetsAdded_sets_atList(void* ptr, int i)
{
	return ({QCandlestickSet * tmp = static_cast<QList<QCandlestickSet *>*>(ptr)->at(i); if (i == static_cast<QList<QCandlestickSet *>*>(ptr)->size()-1) { static_cast<QList<QCandlestickSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickSeries___candlestickSetsAdded_sets_setList(void* ptr, void* i)
{
	static_cast<QList<QCandlestickSet *>*>(ptr)->append(static_cast<QCandlestickSet*>(i));
}

void* QCandlestickSeries___candlestickSetsAdded_sets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCandlestickSet *>();
}

void* QCandlestickSeries___candlestickSetsRemoved_sets_atList(void* ptr, int i)
{
	return ({QCandlestickSet * tmp = static_cast<QList<QCandlestickSet *>*>(ptr)->at(i); if (i == static_cast<QList<QCandlestickSet *>*>(ptr)->size()-1) { static_cast<QList<QCandlestickSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickSeries___candlestickSetsRemoved_sets_setList(void* ptr, void* i)
{
	static_cast<QList<QCandlestickSet *>*>(ptr)->append(static_cast<QCandlestickSet*>(i));
}

void* QCandlestickSeries___candlestickSetsRemoved_sets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCandlestickSet *>();
}

void* QCandlestickSeries___sets_atList(void* ptr, int i)
{
	return ({QCandlestickSet * tmp = static_cast<QList<QCandlestickSet *>*>(ptr)->at(i); if (i == static_cast<QList<QCandlestickSet *>*>(ptr)->size()-1) { static_cast<QList<QCandlestickSet *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickSeries___sets_setList(void* ptr, void* i)
{
	static_cast<QList<QCandlestickSet *>*>(ptr)->append(static_cast<QCandlestickSet*>(i));
}

void* QCandlestickSeries___sets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCandlestickSet *>();
}

class MyQCandlestickSet: public QCandlestickSet
{
public:
	MyQCandlestickSet(qreal open, qreal high, qreal low, qreal close, qreal timestamp = 0.0, QObject *parent = Q_NULLPTR) : QCandlestickSet(open, high, low, close, timestamp, parent) {QCandlestickSet_QCandlestickSet_QRegisterMetaType();};
	MyQCandlestickSet(qreal timestamp = 0.0, QObject *parent = Q_NULLPTR) : QCandlestickSet(timestamp, parent) {QCandlestickSet_QCandlestickSet_QRegisterMetaType();};
	void Signal_BrushChanged() { callbackQCandlestickSet_BrushChanged(this); };
	void Signal_Clicked() { callbackQCandlestickSet_Clicked(this); };
	void Signal_CloseChanged() { callbackQCandlestickSet_CloseChanged(this); };
	void Signal_DoubleClicked() { callbackQCandlestickSet_DoubleClicked(this); };
	void Signal_HighChanged() { callbackQCandlestickSet_HighChanged(this); };
	void Signal_Hovered(bool status) { callbackQCandlestickSet_Hovered(this, status); };
	void Signal_LowChanged() { callbackQCandlestickSet_LowChanged(this); };
	void Signal_OpenChanged() { callbackQCandlestickSet_OpenChanged(this); };
	void Signal_PenChanged() { callbackQCandlestickSet_PenChanged(this); };
	void Signal_Pressed() { callbackQCandlestickSet_Pressed(this); };
	void Signal_Released() { callbackQCandlestickSet_Released(this); };
	void Signal_TimestampChanged() { callbackQCandlestickSet_TimestampChanged(this); };
	 ~MyQCandlestickSet() { callbackQCandlestickSet_DestroyQCandlestickSet(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCandlestickSet_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQCandlestickSet_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCandlestickSet_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQCandlestickSet_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCandlestickSet_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCandlestickSet_CustomEvent(this, event); };
	void deleteLater() { callbackQCandlestickSet_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQCandlestickSet_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCandlestickSet_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQCandlestickSet_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQCandlestickSet_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQCandlestickSet*)

int QCandlestickSet_QCandlestickSet_QRegisterMetaType(){qRegisterMetaType<QCandlestickSet*>(); return qRegisterMetaType<MyQCandlestickSet*>();}

void* QCandlestickSet_NewQCandlestickSet2(double open, double high, double low, double close, double timestamp, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QWindow*>(parent));
	} else {
		return new MyQCandlestickSet(open, high, low, close, timestamp, static_cast<QObject*>(parent));
	}
}

void* QCandlestickSet_NewQCandlestickSet(double timestamp, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCandlestickSet(timestamp, static_cast<QWindow*>(parent));
	} else {
		return new MyQCandlestickSet(timestamp, static_cast<QObject*>(parent));
	}
}

struct QtCharts_PackedString QCandlestickSet_QCandlestickSet_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t442373 = QCandlestickSet::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t442373.prepend("WHITESPACE").constData()+10), t442373.size()-10 }; });
}

struct QtCharts_PackedString QCandlestickSet_QCandlestickSet_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t808f5d = QCandlestickSet::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t808f5d.prepend("WHITESPACE").constData()+10), t808f5d.size()-10 }; });
}

void QCandlestickSet_ConnectBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::brushChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_BrushChanged));
}

void QCandlestickSet_DisconnectBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::brushChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_BrushChanged));
}

void QCandlestickSet_BrushChanged(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->brushChanged();
}

void QCandlestickSet_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::clicked), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_Clicked));
}

void QCandlestickSet_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::clicked), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_Clicked));
}

void QCandlestickSet_Clicked(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->clicked();
}

void QCandlestickSet_ConnectCloseChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::closeChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_CloseChanged));
}

void QCandlestickSet_DisconnectCloseChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::closeChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_CloseChanged));
}

void QCandlestickSet_CloseChanged(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->closeChanged();
}

void QCandlestickSet_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::doubleClicked), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_DoubleClicked));
}

void QCandlestickSet_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::doubleClicked), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_DoubleClicked));
}

void QCandlestickSet_DoubleClicked(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->doubleClicked();
}

void QCandlestickSet_ConnectHighChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::highChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_HighChanged));
}

void QCandlestickSet_DisconnectHighChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::highChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_HighChanged));
}

void QCandlestickSet_HighChanged(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->highChanged();
}

void QCandlestickSet_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)(bool)>(&QCandlestickSet::hovered), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)(bool)>(&MyQCandlestickSet::Signal_Hovered));
}

void QCandlestickSet_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)(bool)>(&QCandlestickSet::hovered), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)(bool)>(&MyQCandlestickSet::Signal_Hovered));
}

void QCandlestickSet_Hovered(void* ptr, char status)
{
	static_cast<QCandlestickSet*>(ptr)->hovered(status != 0);
}

void QCandlestickSet_ConnectLowChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::lowChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_LowChanged));
}

void QCandlestickSet_DisconnectLowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::lowChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_LowChanged));
}

void QCandlestickSet_LowChanged(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->lowChanged();
}

void QCandlestickSet_ConnectOpenChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::openChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_OpenChanged));
}

void QCandlestickSet_DisconnectOpenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::openChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_OpenChanged));
}

void QCandlestickSet_OpenChanged(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->openChanged();
}

void QCandlestickSet_ConnectPenChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::penChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_PenChanged));
}

void QCandlestickSet_DisconnectPenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::penChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_PenChanged));
}

void QCandlestickSet_PenChanged(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->penChanged();
}

void QCandlestickSet_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::pressed), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_Pressed));
}

void QCandlestickSet_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::pressed), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_Pressed));
}

void QCandlestickSet_Pressed(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->pressed();
}

void QCandlestickSet_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::released), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_Released));
}

void QCandlestickSet_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::released), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_Released));
}

void QCandlestickSet_Released(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->released();
}

void QCandlestickSet_SetBrush(void* ptr, void* brush)
{
	static_cast<QCandlestickSet*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QCandlestickSet_SetClose(void* ptr, double close)
{
	static_cast<QCandlestickSet*>(ptr)->setClose(close);
}

void QCandlestickSet_SetHigh(void* ptr, double high)
{
	static_cast<QCandlestickSet*>(ptr)->setHigh(high);
}

void QCandlestickSet_SetLow(void* ptr, double low)
{
	static_cast<QCandlestickSet*>(ptr)->setLow(low);
}

void QCandlestickSet_SetOpen(void* ptr, double open)
{
	static_cast<QCandlestickSet*>(ptr)->setOpen(open);
}

void QCandlestickSet_SetPen(void* ptr, void* pen)
{
	static_cast<QCandlestickSet*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QCandlestickSet_SetTimestamp(void* ptr, double timestamp)
{
	static_cast<QCandlestickSet*>(ptr)->setTimestamp(timestamp);
}

void QCandlestickSet_ConnectTimestampChanged(void* ptr)
{
	QObject::connect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::timestampChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_TimestampChanged));
}

void QCandlestickSet_DisconnectTimestampChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCandlestickSet*>(ptr), static_cast<void (QCandlestickSet::*)()>(&QCandlestickSet::timestampChanged), static_cast<MyQCandlestickSet*>(ptr), static_cast<void (MyQCandlestickSet::*)()>(&MyQCandlestickSet::Signal_TimestampChanged));
}

void QCandlestickSet_TimestampChanged(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->timestampChanged();
}

void QCandlestickSet_DestroyQCandlestickSet(void* ptr)
{
	static_cast<QCandlestickSet*>(ptr)->~QCandlestickSet();
}

void QCandlestickSet_DestroyQCandlestickSetDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QCandlestickSet_Brush(void* ptr)
{
	return new QBrush(static_cast<QCandlestickSet*>(ptr)->brush());
}

void* QCandlestickSet_Pen(void* ptr)
{
	return new QPen(static_cast<QCandlestickSet*>(ptr)->pen());
}

void* QCandlestickSet_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QCandlestickSet*>(ptr)->QCandlestickSet::metaObject());
}

double QCandlestickSet_Close(void* ptr)
{
	return static_cast<QCandlestickSet*>(ptr)->close();
}

double QCandlestickSet_High(void* ptr)
{
	return static_cast<QCandlestickSet*>(ptr)->high();
}

double QCandlestickSet_Low(void* ptr)
{
	return static_cast<QCandlestickSet*>(ptr)->low();
}

double QCandlestickSet_Open(void* ptr)
{
	return static_cast<QCandlestickSet*>(ptr)->open();
}

double QCandlestickSet_Timestamp(void* ptr)
{
	return static_cast<QCandlestickSet*>(ptr)->timestamp();
}

void* QCandlestickSet___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QCandlestickSet___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QCandlestickSet___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QCandlestickSet___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickSet___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCandlestickSet___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QCandlestickSet___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickSet___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCandlestickSet___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QCandlestickSet___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickSet___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCandlestickSet___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QCandlestickSet___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QCandlestickSet___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QCandlestickSet___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QCandlestickSet_EventDefault(void* ptr, void* e)
{
		return static_cast<QCandlestickSet*>(ptr)->QCandlestickSet::event(static_cast<QEvent*>(e));
}

char QCandlestickSet_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QCandlestickSet*>(ptr)->QCandlestickSet::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QCandlestickSet_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QCandlestickSet*>(ptr)->QCandlestickSet::childEvent(static_cast<QChildEvent*>(event));
}

void QCandlestickSet_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QCandlestickSet*>(ptr)->QCandlestickSet::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCandlestickSet_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QCandlestickSet*>(ptr)->QCandlestickSet::customEvent(static_cast<QEvent*>(event));
}

void QCandlestickSet_DeleteLaterDefault(void* ptr)
{
		static_cast<QCandlestickSet*>(ptr)->QCandlestickSet::deleteLater();
}

void QCandlestickSet_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QCandlestickSet*>(ptr)->QCandlestickSet::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCandlestickSet_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QCandlestickSet*>(ptr)->QCandlestickSet::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQCategoryAxis: public QCategoryAxis
{
public:
	MyQCategoryAxis(QObject *parent = Q_NULLPTR) : QCategoryAxis(parent) {QCategoryAxis_QCategoryAxis_QRegisterMetaType();};
	void Signal_CategoriesChanged() { callbackQCategoryAxis_CategoriesChanged(this); };
	void Signal_LabelsPositionChanged(QCategoryAxis::AxisLabelsPosition position) { callbackQCategoryAxis_LabelsPositionChanged(this, position); };
	 ~MyQCategoryAxis() { callbackQCategoryAxis_DestroyQCategoryAxis(this); };
	QAbstractAxis::AxisType type() const { return static_cast<QAbstractAxis::AxisType>(callbackQValueAxis_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractAxis_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void applyNiceNumbers() { callbackQValueAxis_ApplyNiceNumbers(this); };
	void Signal_LabelFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQValueAxis_LabelFormatChanged(this, formatPacked); };
	void Signal_MaxChanged(qreal max) { callbackQValueAxis_MaxChanged(this, max); };
	void Signal_MinChanged(qreal min) { callbackQValueAxis_MinChanged(this, min); };
	void Signal_MinorTickCountChanged(int minorTickCount) { callbackQValueAxis_MinorTickCountChanged(this, minorTickCount); };
	void Signal_RangeChanged(qreal min, qreal max) { callbackQValueAxis_RangeChanged(this, min, max); };
	void Signal_TickAnchorChanged(qreal anchor) { callbackQValueAxis_TickAnchorChanged(this, anchor); };
	void Signal_TickCountChanged(int tickCount) { callbackQValueAxis_TickCountChanged(this, tickCount); };
	void Signal_TickIntervalChanged(qreal interval) { callbackQValueAxis_TickIntervalChanged(this, interval); };
	void Signal_TickTypeChanged(QValueAxis::TickType ty) { callbackQValueAxis_TickTypeChanged(this, ty); };
	void Signal_ColorChanged(QColor color) { callbackQAbstractAxis_ColorChanged(this, new QColor(color)); };
	void Signal_GridLineColorChanged(const QColor & color) { callbackQAbstractAxis_GridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_GridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_GridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_GridVisibleChanged(bool visible) { callbackQAbstractAxis_GridVisibleChanged(this, visible); };
	void Signal_LabelsAngleChanged(int angle) { callbackQAbstractAxis_LabelsAngleChanged(this, angle); };
	void Signal_LabelsBrushChanged(const QBrush & brush) { callbackQAbstractAxis_LabelsBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_LabelsColorChanged(QColor color) { callbackQAbstractAxis_LabelsColorChanged(this, new QColor(color)); };
	void Signal_LabelsFontChanged(const QFont & font) { callbackQAbstractAxis_LabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_LabelsVisibleChanged(bool visible) { callbackQAbstractAxis_LabelsVisibleChanged(this, visible); };
	void Signal_LinePenChanged(const QPen & pen) { callbackQAbstractAxis_LinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_LineVisibleChanged(bool visible) { callbackQAbstractAxis_LineVisibleChanged(this, visible); };
	void Signal_MinorGridLineColorChanged(const QColor & color) { callbackQAbstractAxis_MinorGridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_MinorGridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_MinorGridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_MinorGridVisibleChanged(bool visible) { callbackQAbstractAxis_MinorGridVisibleChanged(this, visible); };
	void Signal_ReverseChanged(bool reverse) { callbackQAbstractAxis_ReverseChanged(this, reverse); };
	void Signal_ShadesBorderColorChanged(QColor color) { callbackQAbstractAxis_ShadesBorderColorChanged(this, new QColor(color)); };
	void Signal_ShadesBrushChanged(const QBrush & brush) { callbackQAbstractAxis_ShadesBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_ShadesColorChanged(QColor color) { callbackQAbstractAxis_ShadesColorChanged(this, new QColor(color)); };
	void Signal_ShadesPenChanged(const QPen & pen) { callbackQAbstractAxis_ShadesPenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_ShadesVisibleChanged(bool visible) { callbackQAbstractAxis_ShadesVisibleChanged(this, visible); };
	void Signal_TitleBrushChanged(const QBrush & brush) { callbackQAbstractAxis_TitleBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_TitleFontChanged(const QFont & font) { callbackQAbstractAxis_TitleFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_TitleTextChanged(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtCharts_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQAbstractAxis_TitleTextChanged(this, textPacked); };
	void Signal_TitleVisibleChanged(bool visible) { callbackQAbstractAxis_TitleVisibleChanged(this, visible); };
	void Signal_VisibleChanged(bool visible) { callbackQAbstractAxis_VisibleChanged(this, visible); };
	bool event(QEvent * e) { return callbackQAbstractAxis_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractAxis_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractAxis_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractAxis_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractAxis_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractAxis_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractAxis_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractAxis_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQCategoryAxis*)

int QCategoryAxis_QCategoryAxis_QRegisterMetaType(){qRegisterMetaType<QCategoryAxis*>(); return qRegisterMetaType<MyQCategoryAxis*>();}

void* QCategoryAxis_NewQCategoryAxis(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCategoryAxis(static_cast<QWindow*>(parent));
	} else {
		return new MyQCategoryAxis(static_cast<QObject*>(parent));
	}
}

struct QtCharts_PackedString QCategoryAxis_CategoriesLabels(void* ptr)
{
	return ({ QByteArray t6d5777 = static_cast<QCategoryAxis*>(ptr)->categoriesLabels().join("|").toUtf8(); QtCharts_PackedString { const_cast<char*>(t6d5777.prepend("WHITESPACE").constData()+10), t6d5777.size()-10 }; });
}

void QCategoryAxis_Append(void* ptr, struct QtCharts_PackedString categoryLabel, double categoryEndValue)
{
	static_cast<QCategoryAxis*>(ptr)->append(QString::fromUtf8(categoryLabel.data, categoryLabel.len), categoryEndValue);
}

void QCategoryAxis_ConnectCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QCategoryAxis*>(ptr), static_cast<void (QCategoryAxis::*)()>(&QCategoryAxis::categoriesChanged), static_cast<MyQCategoryAxis*>(ptr), static_cast<void (MyQCategoryAxis::*)()>(&MyQCategoryAxis::Signal_CategoriesChanged));
}

void QCategoryAxis_DisconnectCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCategoryAxis*>(ptr), static_cast<void (QCategoryAxis::*)()>(&QCategoryAxis::categoriesChanged), static_cast<MyQCategoryAxis*>(ptr), static_cast<void (MyQCategoryAxis::*)()>(&MyQCategoryAxis::Signal_CategoriesChanged));
}

void QCategoryAxis_CategoriesChanged(void* ptr)
{
	static_cast<QCategoryAxis*>(ptr)->categoriesChanged();
}

void QCategoryAxis_ConnectLabelsPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QCategoryAxis*>(ptr), static_cast<void (QCategoryAxis::*)(QCategoryAxis::AxisLabelsPosition)>(&QCategoryAxis::labelsPositionChanged), static_cast<MyQCategoryAxis*>(ptr), static_cast<void (MyQCategoryAxis::*)(QCategoryAxis::AxisLabelsPosition)>(&MyQCategoryAxis::Signal_LabelsPositionChanged));
}

void QCategoryAxis_DisconnectLabelsPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCategoryAxis*>(ptr), static_cast<void (QCategoryAxis::*)(QCategoryAxis::AxisLabelsPosition)>(&QCategoryAxis::labelsPositionChanged), static_cast<MyQCategoryAxis*>(ptr), static_cast<void (MyQCategoryAxis::*)(QCategoryAxis::AxisLabelsPosition)>(&MyQCategoryAxis::Signal_LabelsPositionChanged));
}

void QCategoryAxis_LabelsPositionChanged(void* ptr, long long position)
{
	static_cast<QCategoryAxis*>(ptr)->labelsPositionChanged(static_cast<QCategoryAxis::AxisLabelsPosition>(position));
}

void QCategoryAxis_Remove(void* ptr, struct QtCharts_PackedString categoryLabel)
{
	static_cast<QCategoryAxis*>(ptr)->remove(QString::fromUtf8(categoryLabel.data, categoryLabel.len));
}

void QCategoryAxis_ReplaceLabel(void* ptr, struct QtCharts_PackedString oldLabel, struct QtCharts_PackedString newLabel)
{
	static_cast<QCategoryAxis*>(ptr)->replaceLabel(QString::fromUtf8(oldLabel.data, oldLabel.len), QString::fromUtf8(newLabel.data, newLabel.len));
}

void QCategoryAxis_SetLabelsPosition(void* ptr, long long position)
{
	static_cast<QCategoryAxis*>(ptr)->setLabelsPosition(static_cast<QCategoryAxis::AxisLabelsPosition>(position));
}

void QCategoryAxis_SetStartValue(void* ptr, double min)
{
	static_cast<QCategoryAxis*>(ptr)->setStartValue(min);
}

void QCategoryAxis_DestroyQCategoryAxis(void* ptr)
{
	static_cast<QCategoryAxis*>(ptr)->~QCategoryAxis();
}

void QCategoryAxis_DestroyQCategoryAxisDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QCategoryAxis_LabelsPosition(void* ptr)
{
	return static_cast<QCategoryAxis*>(ptr)->labelsPosition();
}

int QCategoryAxis_Count(void* ptr)
{
	return static_cast<QCategoryAxis*>(ptr)->count();
}

double QCategoryAxis_EndValue(void* ptr, struct QtCharts_PackedString categoryLabel)
{
	return static_cast<QCategoryAxis*>(ptr)->endValue(QString::fromUtf8(categoryLabel.data, categoryLabel.len));
}

double QCategoryAxis_StartValue(void* ptr, struct QtCharts_PackedString categoryLabel)
{
	return static_cast<QCategoryAxis*>(ptr)->startValue(QString::fromUtf8(categoryLabel.data, categoryLabel.len));
}

class MyQChart: public QChart
{
public:
	MyQChart(QGraphicsItem *parent = Q_NULLPTR, Qt::WindowFlags wFlags = Qt::WindowFlags()) : QChart(parent, wFlags) {QChart_QChart_QRegisterMetaType();};
	void Signal_PlotAreaChanged(const QRectF & plotArea) { callbackQChart_PlotAreaChanged(this, const_cast<QRectF*>(&plotArea)); };
	 ~MyQChart() { callbackQChart_DestroyQChart(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQChart_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant & value) { return *static_cast<QVariant*>(callbackQChart_ItemChange(this, change, const_cast<QVariant*>(&value))); };
	bool close() { return callbackQChart_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQChart_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQChart_FocusNextPrevChild(this, next) != 0; };
	bool sceneEvent(QEvent * event) { return callbackQChart_SceneEvent(this, event) != 0; };
	bool windowFrameEvent(QEvent * event) { return callbackQChart_WindowFrameEvent(this, event) != 0; };
	void changeEvent(QEvent * event) { callbackQChart_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQChart_CloseEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQChart_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQChart_FocusOutEvent(this, event); };
	void Signal_GeometryChanged() { callbackQChart_GeometryChanged(this); };
	void grabKeyboardEvent(QEvent * event) { callbackQChart_GrabKeyboardEvent(this, event); };
	void grabMouseEvent(QEvent * event) { callbackQChart_GrabMouseEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQChart_HideEvent(this, event); };
	void hoverLeaveEvent(QGraphicsSceneHoverEvent * event) { callbackQChart_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QGraphicsSceneHoverEvent * event) { callbackQChart_HoverMoveEvent(this, event); };
	void moveEvent(QGraphicsSceneMoveEvent * event) { callbackQChart_MoveEvent(this, event); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQChart_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void paintWindowFrame(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQChart_PaintWindowFrame(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void polishEvent() { callbackQChart_PolishEvent(this); };
	void resizeEvent(QGraphicsSceneResizeEvent * event) { callbackQChart_ResizeEvent(this, event); };
	void setGeometry(const QRectF & rect) { callbackQChart_SetGeometry(this, const_cast<QRectF*>(&rect)); };
	void showEvent(QShowEvent * event) { callbackQChart_ShowEvent(this, event); };
	void ungrabKeyboardEvent(QEvent * event) { callbackQChart_UngrabKeyboardEvent(this, event); };
	void ungrabMouseEvent(QEvent * event) { callbackQChart_UngrabMouseEvent(this, event); };
	void updateGeometry() { callbackQChart_UpdateGeometry(this); };
	QPainterPath shape() const { return *static_cast<QPainterPath*>(callbackQChart_Shape(const_cast<void*>(static_cast<const void*>(this)))); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQChart_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQChart_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
	Qt::WindowFrameSection windowFrameSectionAt(const QPointF & pos) const { return static_cast<Qt::WindowFrameSection>(callbackQChart_WindowFrameSectionAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&pos))); };
	int type() const { return callbackQChart_Type(const_cast<void*>(static_cast<const void*>(this))); };
	void getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom) const { callbackQChart_GetContentsMargins(const_cast<void*>(static_cast<const void*>(this)), *left, *top, *right, *bottom); };
	void initStyleOption(QStyleOption * option) const { callbackQChart_InitStyleOption(const_cast<void*>(static_cast<const void*>(this)), option); };
	void Signal_EnabledChanged() { callbackQChart_EnabledChanged(this); };
	void Signal_OpacityChanged() { callbackQChart_OpacityChanged(this); };
	void Signal_ParentChanged() { callbackQChart_ParentChanged(this); };
	void Signal_RotationChanged() { callbackQChart_RotationChanged(this); };
	void Signal_ScaleChanged() { callbackQChart_ScaleChanged(this); };
	void updateMicroFocus() { callbackQChart_UpdateMicroFocus(this); };
	void Signal_VisibleChanged() { callbackQChart_VisibleChanged(this); };
	void Signal_XChanged() { callbackQChart_XChanged(this); };
	void Signal_YChanged() { callbackQChart_YChanged(this); };
	void Signal_ZChanged() { callbackQChart_ZChanged(this); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQChart_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQChart_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQChart_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQChart_CustomEvent(this, event); };
	void deleteLater() { callbackQChart_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQChart_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQChart_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQChart_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQChart_TimerEvent(this, event); };
	bool sceneEventFilter(QGraphicsItem * watched, QEvent * event) { return callbackQChart_SceneEventFilter(this, watched, event) != 0; };
	void advance(int phase) { callbackQChart_Advance(this, phase); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQChart_ContextMenuEvent(this, event); };
	void dragEnterEvent(QGraphicsSceneDragDropEvent * event) { callbackQChart_DragEnterEvent(this, event); };
	void dragLeaveEvent(QGraphicsSceneDragDropEvent * event) { callbackQChart_DragLeaveEvent(this, event); };
	void dragMoveEvent(QGraphicsSceneDragDropEvent * event) { callbackQChart_DragMoveEvent(this, event); };
	void dropEvent(QGraphicsSceneDragDropEvent * event) { callbackQChart_DropEvent(this, event); };
	void hoverEnterEvent(QGraphicsSceneHoverEvent * event) { callbackQChart_HoverEnterEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQChart_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQChart_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQChart_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQChart_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * event) { callbackQChart_MouseMoveEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQChart_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQChart_MouseReleaseEvent(this, event); };
	void wheelEvent(QGraphicsSceneWheelEvent * event) { callbackQChart_WheelEvent(this, event); };
	QPainterPath opaqueArea() const { return *static_cast<QPainterPath*>(callbackQChart_OpaqueArea(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQChart_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool collidesWithItem(const QGraphicsItem * other, Qt::ItemSelectionMode mode) const { return callbackQChart_CollidesWithItem(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(other), mode) != 0; };
	bool collidesWithPath(const QPainterPath & path, Qt::ItemSelectionMode mode) const { return callbackQChart_CollidesWithPath(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPainterPath*>(&path), mode) != 0; };
	bool contains(const QPointF & point) const { return callbackQChart_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	bool isObscuredBy(const QGraphicsItem * item) const { return callbackQChart_IsObscuredBy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(item)) != 0; };
};

Q_DECLARE_METATYPE(MyQChart*)

int QChart_QChart_QRegisterMetaType(){qRegisterMetaType<QChart*>(); return qRegisterMetaType<MyQChart*>();}

void* QChart_NewQChart(void* parent, long long wFlags)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQChart(static_cast<QGraphicsObject*>(parent), static_cast<Qt::WindowType>(wFlags));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQChart(static_cast<QGraphicsWidget*>(parent), static_cast<Qt::WindowType>(wFlags));
	} else {
		return new MyQChart(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
	}
}

void* QChart_MapToPosition(void* ptr, void* value, void* series)
{
	return ({ QPointF tmpValue = static_cast<QChart*>(ptr)->mapToPosition(*static_cast<QPointF*>(value), static_cast<QAbstractSeries*>(series)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void* QChart_MapToValue(void* ptr, void* position, void* series)
{
	return ({ QPointF tmpValue = static_cast<QChart*>(ptr)->mapToValue(*static_cast<QPointF*>(position), static_cast<QAbstractSeries*>(series)); new QPointF(tmpValue.x(), tmpValue.y()); });
}

struct QtCharts_PackedString QChart_QChart_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t6e2025 = QChart::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t6e2025.prepend("WHITESPACE").constData()+10), t6e2025.size()-10 }; });
}

struct QtCharts_PackedString QChart_QChart_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray taa5f7e = QChart::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(taa5f7e.prepend("WHITESPACE").constData()+10), taa5f7e.size()-10 }; });
}

char QChart_IsZoomed(void* ptr)
{
	return static_cast<QChart*>(ptr)->isZoomed();
}

void QChart_AddAxis(void* ptr, void* axis, long long alignment)
{
	static_cast<QChart*>(ptr)->addAxis(static_cast<QAbstractAxis*>(axis), static_cast<Qt::AlignmentFlag>(alignment));
}

void QChart_AddSeries(void* ptr, void* series)
{
	static_cast<QChart*>(ptr)->addSeries(static_cast<QAbstractSeries*>(series));
}

void QChart_CreateDefaultAxes(void* ptr)
{
	static_cast<QChart*>(ptr)->createDefaultAxes();
}

void QChart_ConnectPlotAreaChanged(void* ptr)
{
	QObject::connect(static_cast<QChart*>(ptr), static_cast<void (QChart::*)(const QRectF &)>(&QChart::plotAreaChanged), static_cast<MyQChart*>(ptr), static_cast<void (MyQChart::*)(const QRectF &)>(&MyQChart::Signal_PlotAreaChanged));
}

void QChart_DisconnectPlotAreaChanged(void* ptr)
{
	QObject::disconnect(static_cast<QChart*>(ptr), static_cast<void (QChart::*)(const QRectF &)>(&QChart::plotAreaChanged), static_cast<MyQChart*>(ptr), static_cast<void (MyQChart::*)(const QRectF &)>(&MyQChart::Signal_PlotAreaChanged));
}

void QChart_PlotAreaChanged(void* ptr, void* plotArea)
{
	static_cast<QChart*>(ptr)->plotAreaChanged(*static_cast<QRectF*>(plotArea));
}

void QChart_RemoveAllSeries(void* ptr)
{
	static_cast<QChart*>(ptr)->removeAllSeries();
}

void QChart_RemoveAxis(void* ptr, void* axis)
{
	static_cast<QChart*>(ptr)->removeAxis(static_cast<QAbstractAxis*>(axis));
}

void QChart_RemoveSeries(void* ptr, void* series)
{
	static_cast<QChart*>(ptr)->removeSeries(static_cast<QAbstractSeries*>(series));
}

void QChart_Scroll(void* ptr, double dx, double dy)
{
	static_cast<QChart*>(ptr)->scroll(dx, dy);
}

void QChart_SetAnimationDuration(void* ptr, int msecs)
{
	static_cast<QChart*>(ptr)->setAnimationDuration(msecs);
}

void QChart_SetAnimationEasingCurve(void* ptr, void* curve)
{
	static_cast<QChart*>(ptr)->setAnimationEasingCurve(*static_cast<QEasingCurve*>(curve));
}

void QChart_SetAnimationOptions(void* ptr, long long options)
{
	static_cast<QChart*>(ptr)->setAnimationOptions(static_cast<QChart::AnimationOption>(options));
}

void QChart_SetAxisX(void* ptr, void* axis, void* series)
{
	static_cast<QChart*>(ptr)->setAxisX(static_cast<QAbstractAxis*>(axis), static_cast<QAbstractSeries*>(series));
}

void QChart_SetAxisY(void* ptr, void* axis, void* series)
{
	static_cast<QChart*>(ptr)->setAxisY(static_cast<QAbstractAxis*>(axis), static_cast<QAbstractSeries*>(series));
}

void QChart_SetBackgroundBrush(void* ptr, void* brush)
{
	static_cast<QChart*>(ptr)->setBackgroundBrush(*static_cast<QBrush*>(brush));
}

void QChart_SetBackgroundPen(void* ptr, void* pen)
{
	static_cast<QChart*>(ptr)->setBackgroundPen(*static_cast<QPen*>(pen));
}

void QChart_SetBackgroundRoundness(void* ptr, double diameter)
{
	static_cast<QChart*>(ptr)->setBackgroundRoundness(diameter);
}

void QChart_SetBackgroundVisible(void* ptr, char visible)
{
	static_cast<QChart*>(ptr)->setBackgroundVisible(visible != 0);
}

void QChart_SetDropShadowEnabled(void* ptr, char enabled)
{
	static_cast<QChart*>(ptr)->setDropShadowEnabled(enabled != 0);
}

void QChart_SetLocale(void* ptr, void* locale)
{
	static_cast<QChart*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QChart_SetLocalizeNumbers(void* ptr, char localize)
{
	static_cast<QChart*>(ptr)->setLocalizeNumbers(localize != 0);
}

void QChart_SetMargins(void* ptr, void* margins)
{
	static_cast<QChart*>(ptr)->setMargins(*static_cast<QMargins*>(margins));
}

void QChart_SetPlotArea(void* ptr, void* rect)
{
	static_cast<QChart*>(ptr)->setPlotArea(*static_cast<QRectF*>(rect));
}

void QChart_SetPlotAreaBackgroundBrush(void* ptr, void* brush)
{
	static_cast<QChart*>(ptr)->setPlotAreaBackgroundBrush(*static_cast<QBrush*>(brush));
}

void QChart_SetPlotAreaBackgroundPen(void* ptr, void* pen)
{
	static_cast<QChart*>(ptr)->setPlotAreaBackgroundPen(*static_cast<QPen*>(pen));
}

void QChart_SetPlotAreaBackgroundVisible(void* ptr, char visible)
{
	static_cast<QChart*>(ptr)->setPlotAreaBackgroundVisible(visible != 0);
}

void QChart_SetTheme(void* ptr, long long theme)
{
	static_cast<QChart*>(ptr)->setTheme(static_cast<QChart::ChartTheme>(theme));
}

void QChart_SetTitle(void* ptr, struct QtCharts_PackedString title)
{
	static_cast<QChart*>(ptr)->setTitle(QString::fromUtf8(title.data, title.len));
}

void QChart_SetTitleBrush(void* ptr, void* brush)
{
	static_cast<QChart*>(ptr)->setTitleBrush(*static_cast<QBrush*>(brush));
}

void QChart_SetTitleFont(void* ptr, void* font)
{
	static_cast<QChart*>(ptr)->setTitleFont(*static_cast<QFont*>(font));
}

void QChart_Zoom(void* ptr, double factor)
{
	static_cast<QChart*>(ptr)->zoom(factor);
}

void QChart_ZoomIn(void* ptr)
{
	static_cast<QChart*>(ptr)->zoomIn();
}

void QChart_ZoomIn2(void* ptr, void* rect)
{
	static_cast<QChart*>(ptr)->zoomIn(*static_cast<QRectF*>(rect));
}

void QChart_ZoomOut(void* ptr)
{
	static_cast<QChart*>(ptr)->zoomOut();
}

void QChart_ZoomReset(void* ptr)
{
	static_cast<QChart*>(ptr)->zoomReset();
}

void QChart_DestroyQChart(void* ptr)
{
	static_cast<QChart*>(ptr)->~QChart();
}

void QChart_DestroyQChartDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QChart_AxisX(void* ptr, void* series)
{
	return static_cast<QChart*>(ptr)->axisX(static_cast<QAbstractSeries*>(series));
}

void* QChart_AxisY(void* ptr, void* series)
{
	return static_cast<QChart*>(ptr)->axisY(static_cast<QAbstractSeries*>(series));
}

void* QChart_BackgroundBrush(void* ptr)
{
	return new QBrush(static_cast<QChart*>(ptr)->backgroundBrush());
}

void* QChart_PlotAreaBackgroundBrush(void* ptr)
{
	return new QBrush(static_cast<QChart*>(ptr)->plotAreaBackgroundBrush());
}

void* QChart_TitleBrush(void* ptr)
{
	return new QBrush(static_cast<QChart*>(ptr)->titleBrush());
}

long long QChart_AnimationOptions(void* ptr)
{
	return static_cast<QChart*>(ptr)->animationOptions();
}

long long QChart_Theme(void* ptr)
{
	return static_cast<QChart*>(ptr)->theme();
}

long long QChart_ChartType(void* ptr)
{
	return static_cast<QChart*>(ptr)->chartType();
}

void* QChart_AnimationEasingCurve(void* ptr)
{
	return new QEasingCurve(static_cast<QChart*>(ptr)->animationEasingCurve());
}

void* QChart_TitleFont(void* ptr)
{
	return new QFont(static_cast<QChart*>(ptr)->titleFont());
}

void* QChart_Legend(void* ptr)
{
	return static_cast<QChart*>(ptr)->legend();
}

struct QtCharts_PackedList QChart_Axes(void* ptr, long long orientation, void* series)
{
	return ({ QList<QAbstractAxis *>* tmpValue = new QList<QAbstractAxis *>(static_cast<QChart*>(ptr)->axes(static_cast<Qt::Orientation>(orientation), static_cast<QAbstractSeries*>(series))); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtCharts_PackedList QChart_Series(void* ptr)
{
	return ({ QList<QAbstractSeries *>* tmpValue = new QList<QAbstractSeries *>(static_cast<QChart*>(ptr)->series()); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

void* QChart_Locale(void* ptr)
{
	return new QLocale(static_cast<QChart*>(ptr)->locale());
}

void* QChart_Margins(void* ptr)
{
	return ({ QMargins tmpValue = static_cast<QChart*>(ptr)->margins(); new QMargins(tmpValue.left(), tmpValue.top(), tmpValue.right(), tmpValue.bottom()); });
}

void* QChart_BackgroundPen(void* ptr)
{
	return new QPen(static_cast<QChart*>(ptr)->backgroundPen());
}

void* QChart_PlotAreaBackgroundPen(void* ptr)
{
	return new QPen(static_cast<QChart*>(ptr)->plotAreaBackgroundPen());
}

void* QChart_PlotArea(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QChart*>(ptr)->plotArea(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

struct QtCharts_PackedString QChart_Title(void* ptr)
{
	return ({ QByteArray tf2b090 = static_cast<QChart*>(ptr)->title().toUtf8(); QtCharts_PackedString { const_cast<char*>(tf2b090.prepend("WHITESPACE").constData()+10), tf2b090.size()-10 }; });
}

char QChart_IsBackgroundVisible(void* ptr)
{
	return static_cast<QChart*>(ptr)->isBackgroundVisible();
}

char QChart_IsDropShadowEnabled(void* ptr)
{
	return static_cast<QChart*>(ptr)->isDropShadowEnabled();
}

char QChart_IsPlotAreaBackgroundVisible(void* ptr)
{
	return static_cast<QChart*>(ptr)->isPlotAreaBackgroundVisible();
}

char QChart_LocalizeNumbers(void* ptr)
{
	return static_cast<QChart*>(ptr)->localizeNumbers();
}

void* QChart_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QPolarChart*>(ptr)->QPolarChart::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QChart*>(ptr)->QChart::metaObject());
	}
}

int QChart_AnimationDuration(void* ptr)
{
	return static_cast<QChart*>(ptr)->animationDuration();
}

double QChart_BackgroundRoundness(void* ptr)
{
	return static_cast<QChart*>(ptr)->backgroundRoundness();
}

void* QChart___axes_atList(void* ptr, int i)
{
	return ({QAbstractAxis * tmp = static_cast<QList<QAbstractAxis *>*>(ptr)->at(i); if (i == static_cast<QList<QAbstractAxis *>*>(ptr)->size()-1) { static_cast<QList<QAbstractAxis *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___axes_setList(void* ptr, void* i)
{
	static_cast<QList<QAbstractAxis *>*>(ptr)->append(static_cast<QAbstractAxis*>(i));
}

void* QChart___axes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAbstractAxis *>();
}

void* QChart___series_atList(void* ptr, int i)
{
	return ({QAbstractSeries * tmp = static_cast<QList<QAbstractSeries *>*>(ptr)->at(i); if (i == static_cast<QList<QAbstractSeries *>*>(ptr)->size()-1) { static_cast<QList<QAbstractSeries *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___series_setList(void* ptr, void* i)
{
	static_cast<QList<QAbstractSeries *>*>(ptr)->append(static_cast<QAbstractSeries*>(i));
}

void* QChart___series_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAbstractSeries *>();
}

void* QChart___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QChart___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QChart___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QChart___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QChart___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QChart___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QChart___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QChart___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QChart___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QChart___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QChart___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QChart___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QChart___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QChart___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QChart___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QChart___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QChart___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QChart___setTransformations_transformations_atList(void* ptr, int i)
{
	return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___setTransformations_transformations_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
}

void* QChart___setTransformations_transformations_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsTransform *>();
}

void* QChart___childItems_atList(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___childItems_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QChart___childItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QChart___collidingItems_atList(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___collidingItems_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QChart___collidingItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QChart___transformations_atList(void* ptr, int i)
{
	return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChart___transformations_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
}

void* QChart___transformations_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsTransform *>();
}

void* QChart_ItemChangeDefault(void* ptr, long long change, void* value)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QPolarChart*>(ptr)->QPolarChart::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
	} else {
		return new QVariant(static_cast<QChart*>(ptr)->QChart::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
	}
}

char QChart_CloseDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::close();
	} else {
		return static_cast<QChart*>(ptr)->QChart::close();
	}
}

char QChart_EventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::event(static_cast<QEvent*>(event));
	} else {
		return static_cast<QChart*>(ptr)->QChart::event(static_cast<QEvent*>(event));
	}
}

char QChart_FocusNextPrevChildDefault(void* ptr, char next)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::focusNextPrevChild(next != 0);
	} else {
		return static_cast<QChart*>(ptr)->QChart::focusNextPrevChild(next != 0);
	}
}

char QChart_SceneEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::sceneEvent(static_cast<QEvent*>(event));
	} else {
		return static_cast<QChart*>(ptr)->QChart::sceneEvent(static_cast<QEvent*>(event));
	}
}

char QChart_WindowFrameEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::windowFrameEvent(static_cast<QEvent*>(event));
	} else {
		return static_cast<QChart*>(ptr)->QChart::windowFrameEvent(static_cast<QEvent*>(event));
	}
}

void QChart_ChangeEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::changeEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::changeEvent(static_cast<QEvent*>(event));
	}
}

void QChart_CloseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::closeEvent(static_cast<QCloseEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::closeEvent(static_cast<QCloseEvent*>(event));
	}
}

void QChart_FocusInEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::focusInEvent(static_cast<QFocusEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::focusInEvent(static_cast<QFocusEvent*>(event));
	}
}

void QChart_FocusOutEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::focusOutEvent(static_cast<QFocusEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::focusOutEvent(static_cast<QFocusEvent*>(event));
	}
}

void QChart_GrabKeyboardEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::grabKeyboardEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::grabKeyboardEvent(static_cast<QEvent*>(event));
	}
}

void QChart_GrabMouseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::grabMouseEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::grabMouseEvent(static_cast<QEvent*>(event));
	}
}

void QChart_HideEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::hideEvent(static_cast<QHideEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::hideEvent(static_cast<QHideEvent*>(event));
	}
}

void QChart_HoverLeaveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
	}
}

void QChart_HoverMoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
	}
}

void QChart_MoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::moveEvent(static_cast<QGraphicsSceneMoveEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::moveEvent(static_cast<QGraphicsSceneMoveEvent*>(event));
	}
}

void QChart_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else {
		static_cast<QChart*>(ptr)->QChart::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	}
}

void QChart_PaintWindowFrameDefault(void* ptr, void* painter, void* option, void* widget)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::paintWindowFrame(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	} else {
		static_cast<QChart*>(ptr)->QChart::paintWindowFrame(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
	}
}

void QChart_PolishEventDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::polishEvent();
	} else {
		static_cast<QChart*>(ptr)->QChart::polishEvent();
	}
}

void QChart_ResizeEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::resizeEvent(static_cast<QGraphicsSceneResizeEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::resizeEvent(static_cast<QGraphicsSceneResizeEvent*>(event));
	}
}

void QChart_SetGeometryDefault(void* ptr, void* rect)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::setGeometry(*static_cast<QRectF*>(rect));
	} else {
		static_cast<QChart*>(ptr)->QChart::setGeometry(*static_cast<QRectF*>(rect));
	}
}

void QChart_ShowEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::showEvent(static_cast<QShowEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::showEvent(static_cast<QShowEvent*>(event));
	}
}

void QChart_UngrabKeyboardEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::ungrabKeyboardEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::ungrabKeyboardEvent(static_cast<QEvent*>(event));
	}
}

void QChart_UngrabMouseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::ungrabMouseEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::ungrabMouseEvent(static_cast<QEvent*>(event));
	}
}

void QChart_UpdateGeometryDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::updateGeometry();
	} else {
		static_cast<QChart*>(ptr)->QChart::updateGeometry();
	}
}

void* QChart_ShapeDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return new QPainterPath(static_cast<QPolarChart*>(ptr)->QPolarChart::shape());
	} else {
		return new QPainterPath(static_cast<QChart*>(ptr)->QChart::shape());
	}
}

void* QChart_BoundingRectDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return ({ QRectF tmpValue = static_cast<QPolarChart*>(ptr)->QPolarChart::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QRectF tmpValue = static_cast<QChart*>(ptr)->QChart::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
	}
}

void* QChart_SizeHintDefault(void* ptr, long long which, void* constraint)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return ({ QSizeF tmpValue = static_cast<QPolarChart*>(ptr)->QPolarChart::sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSizeF tmpValue = static_cast<QChart*>(ptr)->QChart::sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
	}
}

long long QChart_WindowFrameSectionAtDefault(void* ptr, void* pos)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::windowFrameSectionAt(*static_cast<QPointF*>(pos));
	} else {
		return static_cast<QChart*>(ptr)->QChart::windowFrameSectionAt(*static_cast<QPointF*>(pos));
	}
}

int QChart_TypeDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::type();
	} else {
		return static_cast<QChart*>(ptr)->QChart::type();
	}
}

void QChart_GetContentsMarginsDefault(void* ptr, double left, double top, double right, double bottom)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::getContentsMargins(&left, &top, &right, &bottom);
	} else {
		static_cast<QChart*>(ptr)->QChart::getContentsMargins(&left, &top, &right, &bottom);
	}
}

void QChart_InitStyleOptionDefault(void* ptr, void* option)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::initStyleOption(static_cast<QStyleOption*>(option));
	} else {
		static_cast<QChart*>(ptr)->QChart::initStyleOption(static_cast<QStyleOption*>(option));
	}
}

void QChart_UpdateMicroFocusDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::updateMicroFocus();
	} else {
		static_cast<QChart*>(ptr)->QChart::updateMicroFocus();
	}
}

char QChart_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QChart*>(ptr)->QChart::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QChart_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QChart_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QChart*>(ptr)->QChart::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QChart_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::customEvent(static_cast<QEvent*>(event));
	}
}

void QChart_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::deleteLater();
	} else {
		static_cast<QChart*>(ptr)->QChart::deleteLater();
	}
}

void QChart_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QChart*>(ptr)->QChart::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QChart_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

char QChart_SceneEventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QChart*>(ptr)->QChart::sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
	}
}

void QChart_AdvanceDefault(void* ptr, int phase)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::advance(phase);
	} else {
		static_cast<QChart*>(ptr)->QChart::advance(phase);
	}
}

void QChart_ContextMenuEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
	}
}

void QChart_DragEnterEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
	}
}

void QChart_DragLeaveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
	}
}

void QChart_DragMoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
	}
}

void QChart_DropEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
	}
}

void QChart_HoverEnterEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
	}
}

void QChart_InputMethodEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
	}
}

void QChart_KeyPressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::keyPressEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::keyPressEvent(static_cast<QKeyEvent*>(event));
	}
}

void QChart_KeyReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::keyReleaseEvent(static_cast<QKeyEvent*>(event));
	}
}

void QChart_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QChart_MouseMoveEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QChart_MousePressEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QChart_MouseReleaseEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
	}
}

void QChart_WheelEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		static_cast<QPolarChart*>(ptr)->QPolarChart::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
	} else {
		static_cast<QChart*>(ptr)->QChart::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
	}
}

void* QChart_OpaqueAreaDefault(void* ptr)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return new QPainterPath(static_cast<QPolarChart*>(ptr)->QPolarChart::opaqueArea());
	} else {
		return new QPainterPath(static_cast<QChart*>(ptr)->QChart::opaqueArea());
	}
}

void* QChart_InputMethodQueryDefault(void* ptr, long long query)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QPolarChart*>(ptr)->QPolarChart::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
	} else {
		return new QVariant(static_cast<QChart*>(ptr)->QChart::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
	}
}

char QChart_CollidesWithItemDefault(void* ptr, void* other, long long mode)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
	} else {
		return static_cast<QChart*>(ptr)->QChart::collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
	}
}

char QChart_CollidesWithPathDefault(void* ptr, void* path, long long mode)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
	} else {
		return static_cast<QChart*>(ptr)->QChart::collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
	}
}

char QChart_ContainsDefault(void* ptr, void* point)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::contains(*static_cast<QPointF*>(point));
	} else {
		return static_cast<QChart*>(ptr)->QChart::contains(*static_cast<QPointF*>(point));
	}
}

char QChart_IsObscuredByDefault(void* ptr, void* item)
{
	if (dynamic_cast<QPolarChart*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPolarChart*>(ptr)->QPolarChart::isObscuredBy(static_cast<QGraphicsItem*>(item));
	} else {
		return static_cast<QChart*>(ptr)->QChart::isObscuredBy(static_cast<QGraphicsItem*>(item));
	}
}

class MyQChartView: public QChartView
{
public:
	MyQChartView(QChart *chart, QWidget *parent = Q_NULLPTR) : QChartView(chart, parent) {QChartView_QChartView_QRegisterMetaType();};
	MyQChartView(QWidget *parent = Q_NULLPTR) : QChartView(parent) {QChartView_QChartView_QRegisterMetaType();};
	void mouseMoveEvent(QMouseEvent * event) { callbackQChartView_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQChartView_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQChartView_MouseReleaseEvent(this, event); };
	void resizeEvent(QResizeEvent * event) { callbackQChartView_ResizeEvent(this, event); };
	 ~MyQChartView() { callbackQChartView_DestroyQChartView(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQChartView_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * event) { return callbackQChartView_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQChartView_FocusNextPrevChild(this, next) != 0; };
	bool viewportEvent(QEvent * event) { return callbackQChartView_ViewportEvent(this, event) != 0; };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQChartView_ContextMenuEvent(this, event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQChartView_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQChartView_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQChartView_DragMoveEvent(this, event); };
	void drawBackground(QPainter * painter, const QRectF & rect) { callbackQChartView_DrawBackground(this, painter, const_cast<QRectF*>(&rect)); };
	void drawForeground(QPainter * painter, const QRectF & rect) { callbackQChartView_DrawForeground(this, painter, const_cast<QRectF*>(&rect)); };
	void dropEvent(QDropEvent * event) { callbackQChartView_DropEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQChartView_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQChartView_FocusOutEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQChartView_InputMethodEvent(this, event); };
	void invalidateScene(const QRectF & rect, QGraphicsScene::SceneLayers layers) { callbackQChartView_InvalidateScene(this, const_cast<QRectF*>(&rect), layers); };
	void keyPressEvent(QKeyEvent * event) { callbackQChartView_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQChartView_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQChartView_MouseDoubleClickEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQChartView_PaintEvent(this, event); };
	void Signal_RubberBandChanged(QRect rubberBandRect, QPointF fromScenePoint, QPointF toScenePoint) { callbackQChartView_RubberBandChanged(this, ({ QRect tmpValue = rubberBandRect; new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); }), ({ QPointF tmpValue = fromScenePoint; new QPointF(tmpValue.x(), tmpValue.y()); }), ({ QPointF tmpValue = toScenePoint; new QPointF(tmpValue.x(), tmpValue.y()); })); };
	void scrollContentsBy(int dx, int dy) { callbackQChartView_ScrollContentsBy(this, dx, dy); };
	void setupViewport(QWidget * widget) { callbackQChartView_SetupViewport(this, widget); };
	void showEvent(QShowEvent * event) { callbackQChartView_ShowEvent(this, event); };
	void updateScene(const QList<QRectF> & rects) { callbackQChartView_UpdateScene(this, ({ QList<QRectF>* tmpValue = const_cast<QList<QRectF>*>(&rects); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void updateSceneRect(const QRectF & rect) { callbackQChartView_UpdateSceneRect(this, const_cast<QRectF*>(&rect)); };
	void wheelEvent(QWheelEvent * event) { callbackQChartView_WheelEvent(this, event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQChartView_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQChartView_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQChartView_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize viewportSizeHint() const { return *static_cast<QSize*>(callbackQChartView_ViewportSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void changeEvent(QEvent * ev) { callbackQChartView_ChangeEvent(this, ev); };
	bool close() { return callbackQChartView_Close(this) != 0; };
	void actionEvent(QActionEvent * event) { callbackQChartView_ActionEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQChartView_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQChartView_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void enterEvent(QEvent * event) { callbackQChartView_EnterEvent(this, event); };
	void hide() { callbackQChartView_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQChartView_HideEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQChartView_LeaveEvent(this, event); };
	void lower() { callbackQChartView_Lower(this); };
	void moveEvent(QMoveEvent * event) { callbackQChartView_MoveEvent(this, event); };
	void raise() { callbackQChartView_Raise(this); };
	void repaint() { callbackQChartView_Repaint(this); };
	void setDisabled(bool disable) { callbackQChartView_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQChartView_SetEnabled(this, vbo); };
	void setFocus() { callbackQChartView_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQChartView_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtCharts_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQChartView_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQChartView_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQChartView_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtCharts_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQChartView_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQChartView_Show(this); };
	void showFullScreen() { callbackQChartView_ShowFullScreen(this); };
	void showMaximized() { callbackQChartView_ShowMaximized(this); };
	void showMinimized() { callbackQChartView_ShowMinimized(this); };
	void showNormal() { callbackQChartView_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQChartView_TabletEvent(this, event); };
	void update() { callbackQChartView_Update(this); };
	void updateMicroFocus() { callbackQChartView_UpdateMicroFocus(this); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQChartView_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtCharts_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQChartView_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQChartView_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQChartView_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQChartView_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQChartView_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackQChartView_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQChartView_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQChartView_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQChartView_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQChartView_CustomEvent(this, event); };
	void deleteLater() { callbackQChartView_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQChartView_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQChartView_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQChartView_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQChartView_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQChartView*)

int QChartView_QChartView_QRegisterMetaType(){qRegisterMetaType<QChartView*>(); return qRegisterMetaType<MyQChartView*>();}

void* QChartView_NewQChartView2(void* chart, void* parent)
{
		return new MyQChartView(static_cast<QChart*>(chart), static_cast<QWidget*>(parent));
}

void* QChartView_NewQChartView(void* parent)
{
		return new MyQChartView(static_cast<QWidget*>(parent));
}

struct QtCharts_PackedString QChartView_QChartView_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t31d26c = QChartView::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t31d26c.prepend("WHITESPACE").constData()+10), t31d26c.size()-10 }; });
}

struct QtCharts_PackedString QChartView_QChartView_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t573d3c = QChartView::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t573d3c.prepend("WHITESPACE").constData()+10), t573d3c.size()-10 }; });
}

void QChartView_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QChartView_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QChartView_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QChartView_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QChartView_SetChart(void* ptr, void* chart)
{
	static_cast<QChartView*>(ptr)->setChart(static_cast<QChart*>(chart));
}

void QChartView_SetRubberBand(void* ptr, long long rubberBand)
{
	static_cast<QChartView*>(ptr)->setRubberBand(static_cast<QChartView::RubberBand>(rubberBand));
}

void QChartView_DestroyQChartView(void* ptr)
{
	static_cast<QChartView*>(ptr)->~QChartView();
}

void QChartView_DestroyQChartViewDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QChartView_Chart(void* ptr)
{
	return static_cast<QChartView*>(ptr)->chart();
}

long long QChartView_RubberBand(void* ptr)
{
	return static_cast<QChartView*>(ptr)->rubberBand();
}

void* QChartView_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QChartView*>(ptr)->QChartView::metaObject());
}

void* QChartView___updateScene_rects_atList(void* ptr, int i)
{
	return ({ QRectF tmpValue = ({QRectF tmp = static_cast<QList<QRectF>*>(ptr)->at(i); if (i == static_cast<QList<QRectF>*>(ptr)->size()-1) { static_cast<QList<QRectF>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QChartView___updateScene_rects_setList(void* ptr, void* i)
{
	static_cast<QList<QRectF>*>(ptr)->append(*static_cast<QRectF*>(i));
}

void* QChartView___updateScene_rects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QRectF>();
}

void* QChartView___items_atList(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___items_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QChartView___items_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QChartView___items_atList7(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___items_setList7(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QChartView___items_newList7(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QChartView___items_atList2(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___items_setList2(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QChartView___items_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QChartView___items_atList6(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___items_setList6(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QChartView___items_newList6(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QChartView___items_atList4(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___items_setList4(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QChartView___items_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QChartView___items_atList3(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___items_setList3(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QChartView___items_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QChartView___items_atList5(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___items_setList5(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QChartView___items_newList5(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QChartView___scrollBarWidgets_atList(void* ptr, int i)
{
	return ({QWidget * tmp = static_cast<QList<QWidget *>*>(ptr)->at(i); if (i == static_cast<QList<QWidget *>*>(ptr)->size()-1) { static_cast<QList<QWidget *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___scrollBarWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QChartView___scrollBarWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>();
}

void* QChartView___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QChartView___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QChartView___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QChartView___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QChartView___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QChartView___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QChartView___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QChartView___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QChartView___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QChartView___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QChartView___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QChartView___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QChartView___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QChartView___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QChartView___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QChartView___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QChartView___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QChartView___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QChartView_EventDefault(void* ptr, void* event)
{
		return static_cast<QChartView*>(ptr)->QChartView::event(static_cast<QEvent*>(event));
}

char QChartView_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QChartView*>(ptr)->QChartView::focusNextPrevChild(next != 0);
}

char QChartView_ViewportEventDefault(void* ptr, void* event)
{
		return static_cast<QChartView*>(ptr)->QChartView::viewportEvent(static_cast<QEvent*>(event));
}

void QChartView_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QChartView_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QChartView_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QChartView_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QChartView_DrawBackgroundDefault(void* ptr, void* painter, void* rect)
{
		static_cast<QChartView*>(ptr)->QChartView::drawBackground(static_cast<QPainter*>(painter), *static_cast<QRectF*>(rect));
}

void QChartView_DrawForegroundDefault(void* ptr, void* painter, void* rect)
{
		static_cast<QChartView*>(ptr)->QChartView::drawForeground(static_cast<QPainter*>(painter), *static_cast<QRectF*>(rect));
}

void QChartView_DropEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::dropEvent(static_cast<QDropEvent*>(event));
}

void QChartView_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QChartView_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QChartView_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QChartView_InvalidateSceneDefault(void* ptr, void* rect, long long layers)
{
		static_cast<QChartView*>(ptr)->QChartView::invalidateScene(*static_cast<QRectF*>(rect), static_cast<QGraphicsScene::SceneLayer>(layers));
}

void QChartView_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QChartView_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QChartView_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QChartView_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::paintEvent(static_cast<QPaintEvent*>(event));
}

void QChartView_ScrollContentsByDefault(void* ptr, int dx, int dy)
{
		static_cast<QChartView*>(ptr)->QChartView::scrollContentsBy(dx, dy);
}

void QChartView_SetupViewportDefault(void* ptr, void* widget)
{
		static_cast<QChartView*>(ptr)->QChartView::setupViewport(static_cast<QWidget*>(widget));
}

void QChartView_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::showEvent(static_cast<QShowEvent*>(event));
}

void QChartView_UpdateSceneDefault(void* ptr, void* rects)
{
		static_cast<QChartView*>(ptr)->QChartView::updateScene(*static_cast<QList<QRectF>*>(rects));
}

void QChartView_UpdateSceneRectDefault(void* ptr, void* rect)
{
		static_cast<QChartView*>(ptr)->QChartView::updateSceneRect(*static_cast<QRectF*>(rect));
}

void QChartView_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QChartView_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QChartView*>(ptr)->QChartView::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QChartView_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QChartView*>(ptr)->QChartView::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QChartView_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QChartView*>(ptr)->QChartView::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QChartView_ViewportSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QChartView*>(ptr)->QChartView::viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QChartView_ChangeEventDefault(void* ptr, void* ev)
{
		static_cast<QChartView*>(ptr)->QChartView::changeEvent(static_cast<QEvent*>(ev));
}

char QChartView_CloseDefault(void* ptr)
{
		return static_cast<QChartView*>(ptr)->QChartView::close();
}

void QChartView_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::actionEvent(static_cast<QActionEvent*>(event));
}

void QChartView_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::closeEvent(static_cast<QCloseEvent*>(event));
}

void QChartView_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::enterEvent(static_cast<QEvent*>(event));
}

void QChartView_HideDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::hide();
}

void QChartView_HideEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::hideEvent(static_cast<QHideEvent*>(event));
}

void QChartView_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::leaveEvent(static_cast<QEvent*>(event));
}

void QChartView_LowerDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::lower();
}

void QChartView_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::moveEvent(static_cast<QMoveEvent*>(event));
}

void QChartView_RaiseDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::raise();
}

void QChartView_RepaintDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::repaint();
}

void QChartView_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QChartView*>(ptr)->QChartView::setDisabled(disable != 0);
}

void QChartView_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QChartView*>(ptr)->QChartView::setEnabled(vbo != 0);
}

void QChartView_SetFocus2Default(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::setFocus();
}

void QChartView_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QChartView*>(ptr)->QChartView::setHidden(hidden != 0);
}

void QChartView_SetStyleSheetDefault(void* ptr, struct QtCharts_PackedString styleSheet)
{
		static_cast<QChartView*>(ptr)->QChartView::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QChartView_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QChartView*>(ptr)->QChartView::setVisible(visible != 0);
}

void QChartView_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QChartView*>(ptr)->QChartView::setWindowModified(vbo != 0);
}

void QChartView_SetWindowTitleDefault(void* ptr, struct QtCharts_PackedString vqs)
{
		static_cast<QChartView*>(ptr)->QChartView::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QChartView_ShowDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::show();
}

void QChartView_ShowFullScreenDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::showFullScreen();
}

void QChartView_ShowMaximizedDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::showMaximized();
}

void QChartView_ShowMinimizedDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::showMinimized();
}

void QChartView_ShowNormalDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::showNormal();
}

void QChartView_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QChartView_UpdateDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::update();
}

void QChartView_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::updateMicroFocus();
}

void* QChartView_PaintEngineDefault(void* ptr)
{
		return static_cast<QChartView*>(ptr)->QChartView::paintEngine();
}

char QChartView_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QChartView*>(ptr)->QChartView::hasHeightForWidth();
}

int QChartView_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QChartView*>(ptr)->QChartView::heightForWidth(w);
}

int QChartView_MetricDefault(void* ptr, long long m)
{
		return static_cast<QChartView*>(ptr)->QChartView::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void QChartView_InitPainterDefault(void* ptr, void* painter)
{
		static_cast<QChartView*>(ptr)->QChartView::initPainter(static_cast<QPainter*>(painter));
}

char QChartView_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QChartView*>(ptr)->QChartView::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QChartView_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::childEvent(static_cast<QChildEvent*>(event));
}

void QChartView_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QChartView*>(ptr)->QChartView::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QChartView_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::customEvent(static_cast<QEvent*>(event));
}

void QChartView_DeleteLaterDefault(void* ptr)
{
		static_cast<QChartView*>(ptr)->QChartView::deleteLater();
}

void QChartView_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QChartView*>(ptr)->QChartView::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QChartView_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QChartView*>(ptr)->QChartView::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQDateTimeAxis: public QDateTimeAxis
{
public:
	MyQDateTimeAxis(QObject *parent = Q_NULLPTR) : QDateTimeAxis(parent) {QDateTimeAxis_QDateTimeAxis_QRegisterMetaType();};
	void Signal_FormatChanged(QString format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQDateTimeAxis_FormatChanged(this, formatPacked); };
	void Signal_MaxChanged(QDateTime max) { callbackQDateTimeAxis_MaxChanged(this, new QDateTime(max)); };
	void Signal_MinChanged(QDateTime min) { callbackQDateTimeAxis_MinChanged(this, new QDateTime(min)); };
	void Signal_RangeChanged(QDateTime min, QDateTime max) { callbackQDateTimeAxis_RangeChanged(this, new QDateTime(min), new QDateTime(max)); };
	void Signal_TickCountChanged(int tickCount) { callbackQDateTimeAxis_TickCountChanged(this, tickCount); };
	 ~MyQDateTimeAxis() { callbackQDateTimeAxis_DestroyQDateTimeAxis(this); };
	QAbstractAxis::AxisType type() const { return static_cast<QAbstractAxis::AxisType>(callbackQDateTimeAxis_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractAxis_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ColorChanged(QColor color) { callbackQAbstractAxis_ColorChanged(this, new QColor(color)); };
	void Signal_GridLineColorChanged(const QColor & color) { callbackQAbstractAxis_GridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_GridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_GridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_GridVisibleChanged(bool visible) { callbackQAbstractAxis_GridVisibleChanged(this, visible); };
	void Signal_LabelsAngleChanged(int angle) { callbackQAbstractAxis_LabelsAngleChanged(this, angle); };
	void Signal_LabelsBrushChanged(const QBrush & brush) { callbackQAbstractAxis_LabelsBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_LabelsColorChanged(QColor color) { callbackQAbstractAxis_LabelsColorChanged(this, new QColor(color)); };
	void Signal_LabelsFontChanged(const QFont & font) { callbackQAbstractAxis_LabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_LabelsVisibleChanged(bool visible) { callbackQAbstractAxis_LabelsVisibleChanged(this, visible); };
	void Signal_LinePenChanged(const QPen & pen) { callbackQAbstractAxis_LinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_LineVisibleChanged(bool visible) { callbackQAbstractAxis_LineVisibleChanged(this, visible); };
	void Signal_MinorGridLineColorChanged(const QColor & color) { callbackQAbstractAxis_MinorGridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_MinorGridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_MinorGridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_MinorGridVisibleChanged(bool visible) { callbackQAbstractAxis_MinorGridVisibleChanged(this, visible); };
	void Signal_ReverseChanged(bool reverse) { callbackQAbstractAxis_ReverseChanged(this, reverse); };
	void Signal_ShadesBorderColorChanged(QColor color) { callbackQAbstractAxis_ShadesBorderColorChanged(this, new QColor(color)); };
	void Signal_ShadesBrushChanged(const QBrush & brush) { callbackQAbstractAxis_ShadesBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_ShadesColorChanged(QColor color) { callbackQAbstractAxis_ShadesColorChanged(this, new QColor(color)); };
	void Signal_ShadesPenChanged(const QPen & pen) { callbackQAbstractAxis_ShadesPenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_ShadesVisibleChanged(bool visible) { callbackQAbstractAxis_ShadesVisibleChanged(this, visible); };
	void Signal_TitleBrushChanged(const QBrush & brush) { callbackQAbstractAxis_TitleBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_TitleFontChanged(const QFont & font) { callbackQAbstractAxis_TitleFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_TitleTextChanged(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtCharts_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQAbstractAxis_TitleTextChanged(this, textPacked); };
	void Signal_TitleVisibleChanged(bool visible) { callbackQAbstractAxis_TitleVisibleChanged(this, visible); };
	void Signal_VisibleChanged(bool visible) { callbackQAbstractAxis_VisibleChanged(this, visible); };
	bool event(QEvent * e) { return callbackQAbstractAxis_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractAxis_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractAxis_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractAxis_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractAxis_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractAxis_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractAxis_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractAxis_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQDateTimeAxis*)

int QDateTimeAxis_QDateTimeAxis_QRegisterMetaType(){qRegisterMetaType<QDateTimeAxis*>(); return qRegisterMetaType<MyQDateTimeAxis*>();}

void* QDateTimeAxis_NewQDateTimeAxis(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDateTimeAxis(static_cast<QWindow*>(parent));
	} else {
		return new MyQDateTimeAxis(static_cast<QObject*>(parent));
	}
}

void QDateTimeAxis_ConnectFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(QString)>(&QDateTimeAxis::formatChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(QString)>(&MyQDateTimeAxis::Signal_FormatChanged));
}

void QDateTimeAxis_DisconnectFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(QString)>(&QDateTimeAxis::formatChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(QString)>(&MyQDateTimeAxis::Signal_FormatChanged));
}

void QDateTimeAxis_FormatChanged(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QDateTimeAxis*>(ptr)->formatChanged(QString::fromUtf8(format.data, format.len));
}

void QDateTimeAxis_ConnectMaxChanged(void* ptr)
{
	QObject::connect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(QDateTime)>(&QDateTimeAxis::maxChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(QDateTime)>(&MyQDateTimeAxis::Signal_MaxChanged));
}

void QDateTimeAxis_DisconnectMaxChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(QDateTime)>(&QDateTimeAxis::maxChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(QDateTime)>(&MyQDateTimeAxis::Signal_MaxChanged));
}

void QDateTimeAxis_MaxChanged(void* ptr, void* max)
{
	static_cast<QDateTimeAxis*>(ptr)->maxChanged(*static_cast<QDateTime*>(max));
}

void QDateTimeAxis_ConnectMinChanged(void* ptr)
{
	QObject::connect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(QDateTime)>(&QDateTimeAxis::minChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(QDateTime)>(&MyQDateTimeAxis::Signal_MinChanged));
}

void QDateTimeAxis_DisconnectMinChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(QDateTime)>(&QDateTimeAxis::minChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(QDateTime)>(&MyQDateTimeAxis::Signal_MinChanged));
}

void QDateTimeAxis_MinChanged(void* ptr, void* min)
{
	static_cast<QDateTimeAxis*>(ptr)->minChanged(*static_cast<QDateTime*>(min));
}

void QDateTimeAxis_ConnectRangeChanged(void* ptr)
{
	QObject::connect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(QDateTime, QDateTime)>(&QDateTimeAxis::rangeChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(QDateTime, QDateTime)>(&MyQDateTimeAxis::Signal_RangeChanged));
}

void QDateTimeAxis_DisconnectRangeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(QDateTime, QDateTime)>(&QDateTimeAxis::rangeChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(QDateTime, QDateTime)>(&MyQDateTimeAxis::Signal_RangeChanged));
}

void QDateTimeAxis_RangeChanged(void* ptr, void* min, void* max)
{
	static_cast<QDateTimeAxis*>(ptr)->rangeChanged(*static_cast<QDateTime*>(min), *static_cast<QDateTime*>(max));
}

void QDateTimeAxis_SetFormat(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QDateTimeAxis*>(ptr)->setFormat(QString::fromUtf8(format.data, format.len));
}

void QDateTimeAxis_SetMax(void* ptr, void* max)
{
	static_cast<QDateTimeAxis*>(ptr)->setMax(*static_cast<QDateTime*>(max));
}

void QDateTimeAxis_SetMin(void* ptr, void* min)
{
	static_cast<QDateTimeAxis*>(ptr)->setMin(*static_cast<QDateTime*>(min));
}

void QDateTimeAxis_SetRange(void* ptr, void* min, void* max)
{
	static_cast<QDateTimeAxis*>(ptr)->setRange(*static_cast<QDateTime*>(min), *static_cast<QDateTime*>(max));
}

void QDateTimeAxis_SetTickCount(void* ptr, int count)
{
	static_cast<QDateTimeAxis*>(ptr)->setTickCount(count);
}

void QDateTimeAxis_ConnectTickCountChanged(void* ptr)
{
	QObject::connect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(int)>(&QDateTimeAxis::tickCountChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(int)>(&MyQDateTimeAxis::Signal_TickCountChanged));
}

void QDateTimeAxis_DisconnectTickCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDateTimeAxis*>(ptr), static_cast<void (QDateTimeAxis::*)(int)>(&QDateTimeAxis::tickCountChanged), static_cast<MyQDateTimeAxis*>(ptr), static_cast<void (MyQDateTimeAxis::*)(int)>(&MyQDateTimeAxis::Signal_TickCountChanged));
}

void QDateTimeAxis_TickCountChanged(void* ptr, int tickCount)
{
	static_cast<QDateTimeAxis*>(ptr)->tickCountChanged(tickCount);
}

void QDateTimeAxis_DestroyQDateTimeAxis(void* ptr)
{
	static_cast<QDateTimeAxis*>(ptr)->~QDateTimeAxis();
}

void QDateTimeAxis_DestroyQDateTimeAxisDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QDateTimeAxis_Type(void* ptr)
{
	return static_cast<QDateTimeAxis*>(ptr)->type();
}

long long QDateTimeAxis_TypeDefault(void* ptr)
{
		return static_cast<QDateTimeAxis*>(ptr)->QDateTimeAxis::type();
}

void* QDateTimeAxis_Max(void* ptr)
{
	return new QDateTime(static_cast<QDateTimeAxis*>(ptr)->max());
}

void* QDateTimeAxis_Min(void* ptr)
{
	return new QDateTime(static_cast<QDateTimeAxis*>(ptr)->min());
}

struct QtCharts_PackedString QDateTimeAxis_Format(void* ptr)
{
	return ({ QByteArray t34d3ef = static_cast<QDateTimeAxis*>(ptr)->format().toUtf8(); QtCharts_PackedString { const_cast<char*>(t34d3ef.prepend("WHITESPACE").constData()+10), t34d3ef.size()-10 }; });
}

int QDateTimeAxis_TickCount(void* ptr)
{
	return static_cast<QDateTimeAxis*>(ptr)->tickCount();
}

class MyQHBarModelMapper: public QHBarModelMapper
{
public:
	MyQHBarModelMapper(QObject *parent = Q_NULLPTR) : QHBarModelMapper(parent) {QHBarModelMapper_QHBarModelMapper_QRegisterMetaType();};
	void Signal_ColumnCountChanged() { callbackQHBarModelMapper_ColumnCountChanged(this); };
	void Signal_FirstBarSetRowChanged() { callbackQHBarModelMapper_FirstBarSetRowChanged(this); };
	void Signal_FirstColumnChanged() { callbackQHBarModelMapper_FirstColumnChanged(this); };
	void Signal_LastBarSetRowChanged() { callbackQHBarModelMapper_LastBarSetRowChanged(this); };
	void Signal_ModelReplaced() { callbackQHBarModelMapper_ModelReplaced(this); };
	void Signal_SeriesReplaced() { callbackQHBarModelMapper_SeriesReplaced(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHBarModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQHBarModelMapper_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHBarModelMapper_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHBarModelMapper_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHBarModelMapper_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHBarModelMapper_CustomEvent(this, event); };
	void deleteLater() { callbackQHBarModelMapper_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHBarModelMapper_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHBarModelMapper_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHBarModelMapper_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHBarModelMapper_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQHBarModelMapper*)

int QHBarModelMapper_QHBarModelMapper_QRegisterMetaType(){qRegisterMetaType<QHBarModelMapper*>(); return qRegisterMetaType<MyQHBarModelMapper*>();}

void* QHBarModelMapper_NewQHBarModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHBarModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQHBarModelMapper(static_cast<QObject*>(parent));
	}
}

struct QtCharts_PackedString QHBarModelMapper_QHBarModelMapper_Tr(char* s, char* c, int n)
{
	return ({ QByteArray tb32295 = QHBarModelMapper::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(tb32295.prepend("WHITESPACE").constData()+10), tb32295.size()-10 }; });
}

struct QtCharts_PackedString QHBarModelMapper_QHBarModelMapper_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tb5a44f = QHBarModelMapper::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(tb5a44f.prepend("WHITESPACE").constData()+10), tb5a44f.size()-10 }; });
}

void QHBarModelMapper_ConnectColumnCountChanged(void* ptr)
{
	QObject::connect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::columnCountChanged), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_ColumnCountChanged));
}

void QHBarModelMapper_DisconnectColumnCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::columnCountChanged), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_ColumnCountChanged));
}

void QHBarModelMapper_ColumnCountChanged(void* ptr)
{
	static_cast<QHBarModelMapper*>(ptr)->columnCountChanged();
}

void QHBarModelMapper_ConnectFirstBarSetRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::firstBarSetRowChanged), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_FirstBarSetRowChanged));
}

void QHBarModelMapper_DisconnectFirstBarSetRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::firstBarSetRowChanged), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_FirstBarSetRowChanged));
}

void QHBarModelMapper_FirstBarSetRowChanged(void* ptr)
{
	static_cast<QHBarModelMapper*>(ptr)->firstBarSetRowChanged();
}

void QHBarModelMapper_ConnectFirstColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::firstColumnChanged), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_FirstColumnChanged));
}

void QHBarModelMapper_DisconnectFirstColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::firstColumnChanged), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_FirstColumnChanged));
}

void QHBarModelMapper_FirstColumnChanged(void* ptr)
{
	static_cast<QHBarModelMapper*>(ptr)->firstColumnChanged();
}

void QHBarModelMapper_ConnectLastBarSetRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::lastBarSetRowChanged), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_LastBarSetRowChanged));
}

void QHBarModelMapper_DisconnectLastBarSetRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::lastBarSetRowChanged), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_LastBarSetRowChanged));
}

void QHBarModelMapper_LastBarSetRowChanged(void* ptr)
{
	static_cast<QHBarModelMapper*>(ptr)->lastBarSetRowChanged();
}

void QHBarModelMapper_ConnectModelReplaced(void* ptr)
{
	QObject::connect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::modelReplaced), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_ModelReplaced));
}

void QHBarModelMapper_DisconnectModelReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::modelReplaced), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_ModelReplaced));
}

void QHBarModelMapper_ModelReplaced(void* ptr)
{
	static_cast<QHBarModelMapper*>(ptr)->modelReplaced();
}

void QHBarModelMapper_ConnectSeriesReplaced(void* ptr)
{
	QObject::connect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::seriesReplaced), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_SeriesReplaced));
}

void QHBarModelMapper_DisconnectSeriesReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QHBarModelMapper*>(ptr), static_cast<void (QHBarModelMapper::*)()>(&QHBarModelMapper::seriesReplaced), static_cast<MyQHBarModelMapper*>(ptr), static_cast<void (MyQHBarModelMapper::*)()>(&MyQHBarModelMapper::Signal_SeriesReplaced));
}

void QHBarModelMapper_SeriesReplaced(void* ptr)
{
	static_cast<QHBarModelMapper*>(ptr)->seriesReplaced();
}

void QHBarModelMapper_SetColumnCount(void* ptr, int columnCount)
{
	static_cast<QHBarModelMapper*>(ptr)->setColumnCount(columnCount);
}

void QHBarModelMapper_SetFirstBarSetRow(void* ptr, int firstBarSetRow)
{
	static_cast<QHBarModelMapper*>(ptr)->setFirstBarSetRow(firstBarSetRow);
}

void QHBarModelMapper_SetFirstColumn(void* ptr, int firstColumn)
{
	static_cast<QHBarModelMapper*>(ptr)->setFirstColumn(firstColumn);
}

void QHBarModelMapper_SetLastBarSetRow(void* ptr, int lastBarSetRow)
{
	static_cast<QHBarModelMapper*>(ptr)->setLastBarSetRow(lastBarSetRow);
}

void QHBarModelMapper_SetModel(void* ptr, void* model)
{
	static_cast<QHBarModelMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QHBarModelMapper_SetSeries(void* ptr, void* series)
{
	static_cast<QHBarModelMapper*>(ptr)->setSeries(static_cast<QAbstractBarSeries*>(series));
}

void* QHBarModelMapper_Series(void* ptr)
{
	return static_cast<QHBarModelMapper*>(ptr)->series();
}

void* QHBarModelMapper_Model(void* ptr)
{
	return static_cast<QHBarModelMapper*>(ptr)->model();
}

void* QHBarModelMapper_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHBarModelMapper*>(ptr)->QHBarModelMapper::metaObject());
}

int QHBarModelMapper_ColumnCount(void* ptr)
{
	return static_cast<QHBarModelMapper*>(ptr)->columnCount();
}

int QHBarModelMapper_FirstBarSetRow(void* ptr)
{
	return static_cast<QHBarModelMapper*>(ptr)->firstBarSetRow();
}

int QHBarModelMapper_FirstColumn(void* ptr)
{
	return static_cast<QHBarModelMapper*>(ptr)->firstColumn();
}

int QHBarModelMapper_LastBarSetRow(void* ptr)
{
	return static_cast<QHBarModelMapper*>(ptr)->lastBarSetRow();
}

void* QHBarModelMapper___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QHBarModelMapper___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHBarModelMapper___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QHBarModelMapper___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHBarModelMapper___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHBarModelMapper___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QHBarModelMapper___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHBarModelMapper___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHBarModelMapper___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QHBarModelMapper___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHBarModelMapper___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHBarModelMapper___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QHBarModelMapper___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHBarModelMapper___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHBarModelMapper___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QHBarModelMapper_EventDefault(void* ptr, void* e)
{
		return static_cast<QHBarModelMapper*>(ptr)->QHBarModelMapper::event(static_cast<QEvent*>(e));
}

char QHBarModelMapper_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHBarModelMapper*>(ptr)->QHBarModelMapper::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHBarModelMapper_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHBarModelMapper*>(ptr)->QHBarModelMapper::childEvent(static_cast<QChildEvent*>(event));
}

void QHBarModelMapper_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHBarModelMapper*>(ptr)->QHBarModelMapper::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHBarModelMapper_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHBarModelMapper*>(ptr)->QHBarModelMapper::customEvent(static_cast<QEvent*>(event));
}

void QHBarModelMapper_DeleteLaterDefault(void* ptr)
{
		static_cast<QHBarModelMapper*>(ptr)->QHBarModelMapper::deleteLater();
}

void QHBarModelMapper_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHBarModelMapper*>(ptr)->QHBarModelMapper::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHBarModelMapper_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QHBarModelMapper*>(ptr)->QHBarModelMapper::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQHBoxPlotModelMapper: public QHBoxPlotModelMapper
{
public:
	MyQHBoxPlotModelMapper(QObject *parent = Q_NULLPTR) : QHBoxPlotModelMapper(parent) {QHBoxPlotModelMapper_QHBoxPlotModelMapper_QRegisterMetaType();};
	void Signal_ColumnCountChanged() { callbackQHBoxPlotModelMapper_ColumnCountChanged(this); };
	void Signal_FirstBoxSetRowChanged() { callbackQHBoxPlotModelMapper_FirstBoxSetRowChanged(this); };
	void Signal_FirstColumnChanged() { callbackQHBoxPlotModelMapper_FirstColumnChanged(this); };
	void Signal_LastBoxSetRowChanged() { callbackQHBoxPlotModelMapper_LastBoxSetRowChanged(this); };
	void Signal_ModelReplaced() { callbackQHBoxPlotModelMapper_ModelReplaced(this); };
	void Signal_SeriesReplaced() { callbackQHBoxPlotModelMapper_SeriesReplaced(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHBoxPlotModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQHBoxPlotModelMapper_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHBoxPlotModelMapper_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHBoxPlotModelMapper_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHBoxPlotModelMapper_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHBoxPlotModelMapper_CustomEvent(this, event); };
	void deleteLater() { callbackQHBoxPlotModelMapper_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHBoxPlotModelMapper_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHBoxPlotModelMapper_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHBoxPlotModelMapper_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHBoxPlotModelMapper_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQHBoxPlotModelMapper*)

int QHBoxPlotModelMapper_QHBoxPlotModelMapper_QRegisterMetaType(){qRegisterMetaType<QHBoxPlotModelMapper*>(); return qRegisterMetaType<MyQHBoxPlotModelMapper*>();}

void* QHBoxPlotModelMapper_NewQHBoxPlotModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHBoxPlotModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQHBoxPlotModelMapper(static_cast<QObject*>(parent));
	}
}

struct QtCharts_PackedString QHBoxPlotModelMapper_QHBoxPlotModelMapper_Tr(char* s, char* c, int n)
{
	return ({ QByteArray ta4366e = QHBoxPlotModelMapper::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(ta4366e.prepend("WHITESPACE").constData()+10), ta4366e.size()-10 }; });
}

struct QtCharts_PackedString QHBoxPlotModelMapper_QHBoxPlotModelMapper_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t684cf5 = QHBoxPlotModelMapper::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t684cf5.prepend("WHITESPACE").constData()+10), t684cf5.size()-10 }; });
}

void QHBoxPlotModelMapper_ConnectColumnCountChanged(void* ptr)
{
	QObject::connect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::columnCountChanged), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_ColumnCountChanged));
}

void QHBoxPlotModelMapper_DisconnectColumnCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::columnCountChanged), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_ColumnCountChanged));
}

void QHBoxPlotModelMapper_ColumnCountChanged(void* ptr)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->columnCountChanged();
}

void QHBoxPlotModelMapper_ConnectFirstBoxSetRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::firstBoxSetRowChanged), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_FirstBoxSetRowChanged));
}

void QHBoxPlotModelMapper_DisconnectFirstBoxSetRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::firstBoxSetRowChanged), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_FirstBoxSetRowChanged));
}

void QHBoxPlotModelMapper_FirstBoxSetRowChanged(void* ptr)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->firstBoxSetRowChanged();
}

void QHBoxPlotModelMapper_ConnectFirstColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::firstColumnChanged), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_FirstColumnChanged));
}

void QHBoxPlotModelMapper_DisconnectFirstColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::firstColumnChanged), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_FirstColumnChanged));
}

void QHBoxPlotModelMapper_FirstColumnChanged(void* ptr)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->firstColumnChanged();
}

void QHBoxPlotModelMapper_ConnectLastBoxSetRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::lastBoxSetRowChanged), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_LastBoxSetRowChanged));
}

void QHBoxPlotModelMapper_DisconnectLastBoxSetRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::lastBoxSetRowChanged), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_LastBoxSetRowChanged));
}

void QHBoxPlotModelMapper_LastBoxSetRowChanged(void* ptr)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->lastBoxSetRowChanged();
}

void QHBoxPlotModelMapper_ConnectModelReplaced(void* ptr)
{
	QObject::connect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::modelReplaced), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_ModelReplaced));
}

void QHBoxPlotModelMapper_DisconnectModelReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::modelReplaced), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_ModelReplaced));
}

void QHBoxPlotModelMapper_ModelReplaced(void* ptr)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->modelReplaced();
}

void QHBoxPlotModelMapper_ConnectSeriesReplaced(void* ptr)
{
	QObject::connect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::seriesReplaced), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_SeriesReplaced));
}

void QHBoxPlotModelMapper_DisconnectSeriesReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QHBoxPlotModelMapper*>(ptr), static_cast<void (QHBoxPlotModelMapper::*)()>(&QHBoxPlotModelMapper::seriesReplaced), static_cast<MyQHBoxPlotModelMapper*>(ptr), static_cast<void (MyQHBoxPlotModelMapper::*)()>(&MyQHBoxPlotModelMapper::Signal_SeriesReplaced));
}

void QHBoxPlotModelMapper_SeriesReplaced(void* ptr)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->seriesReplaced();
}

void QHBoxPlotModelMapper_SetColumnCount(void* ptr, int rowCount)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->setColumnCount(rowCount);
}

void QHBoxPlotModelMapper_SetFirstBoxSetRow(void* ptr, int firstBoxSetRow)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->setFirstBoxSetRow(firstBoxSetRow);
}

void QHBoxPlotModelMapper_SetFirstColumn(void* ptr, int firstColumn)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->setFirstColumn(firstColumn);
}

void QHBoxPlotModelMapper_SetLastBoxSetRow(void* ptr, int lastBoxSetRow)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->setLastBoxSetRow(lastBoxSetRow);
}

void QHBoxPlotModelMapper_SetModel(void* ptr, void* model)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QHBoxPlotModelMapper_SetSeries(void* ptr, void* series)
{
	static_cast<QHBoxPlotModelMapper*>(ptr)->setSeries(static_cast<QBoxPlotSeries*>(series));
}

void* QHBoxPlotModelMapper_Model(void* ptr)
{
	return static_cast<QHBoxPlotModelMapper*>(ptr)->model();
}

void* QHBoxPlotModelMapper_Series(void* ptr)
{
	return static_cast<QHBoxPlotModelMapper*>(ptr)->series();
}

void* QHBoxPlotModelMapper_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHBoxPlotModelMapper*>(ptr)->QHBoxPlotModelMapper::metaObject());
}

int QHBoxPlotModelMapper_ColumnCount(void* ptr)
{
	return static_cast<QHBoxPlotModelMapper*>(ptr)->columnCount();
}

int QHBoxPlotModelMapper_FirstBoxSetRow(void* ptr)
{
	return static_cast<QHBoxPlotModelMapper*>(ptr)->firstBoxSetRow();
}

int QHBoxPlotModelMapper_FirstColumn(void* ptr)
{
	return static_cast<QHBoxPlotModelMapper*>(ptr)->firstColumn();
}

int QHBoxPlotModelMapper_LastBoxSetRow(void* ptr)
{
	return static_cast<QHBoxPlotModelMapper*>(ptr)->lastBoxSetRow();
}

void* QHBoxPlotModelMapper___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QHBoxPlotModelMapper___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHBoxPlotModelMapper___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QHBoxPlotModelMapper___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHBoxPlotModelMapper___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHBoxPlotModelMapper___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QHBoxPlotModelMapper___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHBoxPlotModelMapper___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHBoxPlotModelMapper___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QHBoxPlotModelMapper___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHBoxPlotModelMapper___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHBoxPlotModelMapper___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QHBoxPlotModelMapper___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QHBoxPlotModelMapper___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHBoxPlotModelMapper___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QHBoxPlotModelMapper_EventDefault(void* ptr, void* e)
{
		return static_cast<QHBoxPlotModelMapper*>(ptr)->QHBoxPlotModelMapper::event(static_cast<QEvent*>(e));
}

char QHBoxPlotModelMapper_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHBoxPlotModelMapper*>(ptr)->QHBoxPlotModelMapper::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHBoxPlotModelMapper_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHBoxPlotModelMapper*>(ptr)->QHBoxPlotModelMapper::childEvent(static_cast<QChildEvent*>(event));
}

void QHBoxPlotModelMapper_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHBoxPlotModelMapper*>(ptr)->QHBoxPlotModelMapper::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHBoxPlotModelMapper_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHBoxPlotModelMapper*>(ptr)->QHBoxPlotModelMapper::customEvent(static_cast<QEvent*>(event));
}

void QHBoxPlotModelMapper_DeleteLaterDefault(void* ptr)
{
		static_cast<QHBoxPlotModelMapper*>(ptr)->QHBoxPlotModelMapper::deleteLater();
}

void QHBoxPlotModelMapper_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHBoxPlotModelMapper*>(ptr)->QHBoxPlotModelMapper::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHBoxPlotModelMapper_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QHBoxPlotModelMapper*>(ptr)->QHBoxPlotModelMapper::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQHCandlestickModelMapper: public QHCandlestickModelMapper
{
public:
	MyQHCandlestickModelMapper(QObject *parent = Q_NULLPTR) : QHCandlestickModelMapper(parent) {QHCandlestickModelMapper_QHCandlestickModelMapper_QRegisterMetaType();};
	void Signal_CloseColumnChanged() { callbackQHCandlestickModelMapper_CloseColumnChanged(this); };
	void Signal_FirstSetRowChanged() { callbackQHCandlestickModelMapper_FirstSetRowChanged(this); };
	void Signal_HighColumnChanged() { callbackQHCandlestickModelMapper_HighColumnChanged(this); };
	void Signal_LastSetRowChanged() { callbackQHCandlestickModelMapper_LastSetRowChanged(this); };
	void Signal_LowColumnChanged() { callbackQHCandlestickModelMapper_LowColumnChanged(this); };
	void Signal_OpenColumnChanged() { callbackQHCandlestickModelMapper_OpenColumnChanged(this); };
	void Signal_TimestampColumnChanged() { callbackQHCandlestickModelMapper_TimestampColumnChanged(this); };
	Qt::Orientation orientation() const { return static_cast<Qt::Orientation>(callbackQHCandlestickModelMapper_Orientation(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCandlestickModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ModelReplaced() { callbackQCandlestickModelMapper_ModelReplaced(this); };
	void Signal_SeriesReplaced() { callbackQCandlestickModelMapper_SeriesReplaced(this); };
	bool event(QEvent * e) { return callbackQCandlestickModelMapper_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCandlestickModelMapper_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQCandlestickModelMapper_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCandlestickModelMapper_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCandlestickModelMapper_CustomEvent(this, event); };
	void deleteLater() { callbackQCandlestickModelMapper_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQCandlestickModelMapper_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCandlestickModelMapper_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQCandlestickModelMapper_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQCandlestickModelMapper_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQHCandlestickModelMapper*)

int QHCandlestickModelMapper_QHCandlestickModelMapper_QRegisterMetaType(){qRegisterMetaType<QHCandlestickModelMapper*>(); return qRegisterMetaType<MyQHCandlestickModelMapper*>();}

void* QHCandlestickModelMapper_NewQHCandlestickModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHCandlestickModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQHCandlestickModelMapper(static_cast<QObject*>(parent));
	}
}

void QHCandlestickModelMapper_ConnectCloseColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::closeColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_CloseColumnChanged));
}

void QHCandlestickModelMapper_DisconnectCloseColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::closeColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_CloseColumnChanged));
}

void QHCandlestickModelMapper_CloseColumnChanged(void* ptr)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->closeColumnChanged();
}

void QHCandlestickModelMapper_ConnectFirstSetRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::firstSetRowChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_FirstSetRowChanged));
}

void QHCandlestickModelMapper_DisconnectFirstSetRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::firstSetRowChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_FirstSetRowChanged));
}

void QHCandlestickModelMapper_FirstSetRowChanged(void* ptr)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->firstSetRowChanged();
}

void QHCandlestickModelMapper_ConnectHighColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::highColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_HighColumnChanged));
}

void QHCandlestickModelMapper_DisconnectHighColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::highColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_HighColumnChanged));
}

void QHCandlestickModelMapper_HighColumnChanged(void* ptr)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->highColumnChanged();
}

void QHCandlestickModelMapper_ConnectLastSetRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::lastSetRowChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_LastSetRowChanged));
}

void QHCandlestickModelMapper_DisconnectLastSetRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::lastSetRowChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_LastSetRowChanged));
}

void QHCandlestickModelMapper_LastSetRowChanged(void* ptr)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->lastSetRowChanged();
}

void QHCandlestickModelMapper_ConnectLowColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::lowColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_LowColumnChanged));
}

void QHCandlestickModelMapper_DisconnectLowColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::lowColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_LowColumnChanged));
}

void QHCandlestickModelMapper_LowColumnChanged(void* ptr)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->lowColumnChanged();
}

void QHCandlestickModelMapper_ConnectOpenColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::openColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_OpenColumnChanged));
}

void QHCandlestickModelMapper_DisconnectOpenColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::openColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_OpenColumnChanged));
}

void QHCandlestickModelMapper_OpenColumnChanged(void* ptr)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->openColumnChanged();
}

void QHCandlestickModelMapper_SetCloseColumn(void* ptr, int closeColumn)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->setCloseColumn(closeColumn);
}

void QHCandlestickModelMapper_SetFirstSetRow(void* ptr, int firstSetRow)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->setFirstSetRow(firstSetRow);
}

void QHCandlestickModelMapper_SetHighColumn(void* ptr, int highColumn)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->setHighColumn(highColumn);
}

void QHCandlestickModelMapper_SetLastSetRow(void* ptr, int lastSetRow)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->setLastSetRow(lastSetRow);
}

void QHCandlestickModelMapper_SetLowColumn(void* ptr, int lowColumn)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->setLowColumn(lowColumn);
}

void QHCandlestickModelMapper_SetOpenColumn(void* ptr, int openColumn)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->setOpenColumn(openColumn);
}

void QHCandlestickModelMapper_SetTimestampColumn(void* ptr, int timestampColumn)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->setTimestampColumn(timestampColumn);
}

void QHCandlestickModelMapper_ConnectTimestampColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::timestampColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_TimestampColumnChanged));
}

void QHCandlestickModelMapper_DisconnectTimestampColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHCandlestickModelMapper*>(ptr), static_cast<void (QHCandlestickModelMapper::*)()>(&QHCandlestickModelMapper::timestampColumnChanged), static_cast<MyQHCandlestickModelMapper*>(ptr), static_cast<void (MyQHCandlestickModelMapper::*)()>(&MyQHCandlestickModelMapper::Signal_TimestampColumnChanged));
}

void QHCandlestickModelMapper_TimestampColumnChanged(void* ptr)
{
	static_cast<QHCandlestickModelMapper*>(ptr)->timestampColumnChanged();
}

long long QHCandlestickModelMapper_Orientation(void* ptr)
{
	return static_cast<QHCandlestickModelMapper*>(ptr)->orientation();
}

long long QHCandlestickModelMapper_OrientationDefault(void* ptr)
{
		return static_cast<QHCandlestickModelMapper*>(ptr)->QHCandlestickModelMapper::orientation();
}

int QHCandlestickModelMapper_CloseColumn(void* ptr)
{
	return static_cast<QHCandlestickModelMapper*>(ptr)->closeColumn();
}

int QHCandlestickModelMapper_FirstSetRow(void* ptr)
{
	return static_cast<QHCandlestickModelMapper*>(ptr)->firstSetRow();
}

int QHCandlestickModelMapper_HighColumn(void* ptr)
{
	return static_cast<QHCandlestickModelMapper*>(ptr)->highColumn();
}

int QHCandlestickModelMapper_LastSetRow(void* ptr)
{
	return static_cast<QHCandlestickModelMapper*>(ptr)->lastSetRow();
}

int QHCandlestickModelMapper_LowColumn(void* ptr)
{
	return static_cast<QHCandlestickModelMapper*>(ptr)->lowColumn();
}

int QHCandlestickModelMapper_OpenColumn(void* ptr)
{
	return static_cast<QHCandlestickModelMapper*>(ptr)->openColumn();
}

int QHCandlestickModelMapper_TimestampColumn(void* ptr)
{
	return static_cast<QHCandlestickModelMapper*>(ptr)->timestampColumn();
}

class MyQHPieModelMapper: public QHPieModelMapper
{
public:
	MyQHPieModelMapper(QObject *parent = Q_NULLPTR) : QHPieModelMapper(parent) {};
	void Signal_ColumnCountChanged() { callbackQHPieModelMapper_ColumnCountChanged(this); };
	void Signal_FirstColumnChanged() { callbackQHPieModelMapper_FirstColumnChanged(this); };
	void Signal_LabelsRowChanged() { callbackQHPieModelMapper_LabelsRowChanged(this); };
	void Signal_ModelReplaced() { callbackQHPieModelMapper_ModelReplaced(this); };
	void Signal_SeriesReplaced() { callbackQHPieModelMapper_SeriesReplaced(this); };
	void Signal_ValuesRowChanged() { callbackQHPieModelMapper_ValuesRowChanged(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHPieModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QHPieModelMapper_NewQHPieModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHPieModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQHPieModelMapper(static_cast<QObject*>(parent));
	}
}

struct QtCharts_PackedString QHPieModelMapper_QHPieModelMapper_Tr(char* s, char* c, int n)
{
	return ({ QByteArray td9016d = QHPieModelMapper::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(td9016d.prepend("WHITESPACE").constData()+10), td9016d.size()-10 }; });
}

struct QtCharts_PackedString QHPieModelMapper_QHPieModelMapper_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t15efa5 = QHPieModelMapper::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t15efa5.prepend("WHITESPACE").constData()+10), t15efa5.size()-10 }; });
}

void QHPieModelMapper_ConnectColumnCountChanged(void* ptr)
{
	QObject::connect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::columnCountChanged), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_ColumnCountChanged));
}

void QHPieModelMapper_DisconnectColumnCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::columnCountChanged), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_ColumnCountChanged));
}

void QHPieModelMapper_ColumnCountChanged(void* ptr)
{
	static_cast<QHPieModelMapper*>(ptr)->columnCountChanged();
}

void QHPieModelMapper_ConnectFirstColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::firstColumnChanged), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_FirstColumnChanged));
}

void QHPieModelMapper_DisconnectFirstColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::firstColumnChanged), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_FirstColumnChanged));
}

void QHPieModelMapper_FirstColumnChanged(void* ptr)
{
	static_cast<QHPieModelMapper*>(ptr)->firstColumnChanged();
}

void QHPieModelMapper_ConnectLabelsRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::labelsRowChanged), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_LabelsRowChanged));
}

void QHPieModelMapper_DisconnectLabelsRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::labelsRowChanged), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_LabelsRowChanged));
}

void QHPieModelMapper_LabelsRowChanged(void* ptr)
{
	static_cast<QHPieModelMapper*>(ptr)->labelsRowChanged();
}

void QHPieModelMapper_ConnectModelReplaced(void* ptr)
{
	QObject::connect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::modelReplaced), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_ModelReplaced));
}

void QHPieModelMapper_DisconnectModelReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::modelReplaced), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_ModelReplaced));
}

void QHPieModelMapper_ModelReplaced(void* ptr)
{
	static_cast<QHPieModelMapper*>(ptr)->modelReplaced();
}

void QHPieModelMapper_ConnectSeriesReplaced(void* ptr)
{
	QObject::connect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::seriesReplaced), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_SeriesReplaced));
}

void QHPieModelMapper_DisconnectSeriesReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::seriesReplaced), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_SeriesReplaced));
}

void QHPieModelMapper_SeriesReplaced(void* ptr)
{
	static_cast<QHPieModelMapper*>(ptr)->seriesReplaced();
}

void QHPieModelMapper_SetColumnCount(void* ptr, int columnCount)
{
	static_cast<QHPieModelMapper*>(ptr)->setColumnCount(columnCount);
}

void QHPieModelMapper_SetFirstColumn(void* ptr, int firstColumn)
{
	static_cast<QHPieModelMapper*>(ptr)->setFirstColumn(firstColumn);
}

void QHPieModelMapper_SetLabelsRow(void* ptr, int labelsRow)
{
	static_cast<QHPieModelMapper*>(ptr)->setLabelsRow(labelsRow);
}

void QHPieModelMapper_SetModel(void* ptr, void* model)
{
	static_cast<QHPieModelMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QHPieModelMapper_SetSeries(void* ptr, void* series)
{
	static_cast<QHPieModelMapper*>(ptr)->setSeries(static_cast<QPieSeries*>(series));
}

void QHPieModelMapper_SetValuesRow(void* ptr, int valuesRow)
{
	static_cast<QHPieModelMapper*>(ptr)->setValuesRow(valuesRow);
}

void QHPieModelMapper_ConnectValuesRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::valuesRowChanged), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_ValuesRowChanged));
}

void QHPieModelMapper_DisconnectValuesRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHPieModelMapper*>(ptr), static_cast<void (QHPieModelMapper::*)()>(&QHPieModelMapper::valuesRowChanged), static_cast<MyQHPieModelMapper*>(ptr), static_cast<void (MyQHPieModelMapper::*)()>(&MyQHPieModelMapper::Signal_ValuesRowChanged));
}

void QHPieModelMapper_ValuesRowChanged(void* ptr)
{
	static_cast<QHPieModelMapper*>(ptr)->valuesRowChanged();
}

void* QHPieModelMapper_Model(void* ptr)
{
	return static_cast<QHPieModelMapper*>(ptr)->model();
}

void* QHPieModelMapper_Series(void* ptr)
{
	return static_cast<QHPieModelMapper*>(ptr)->series();
}

void* QHPieModelMapper_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHPieModelMapper*>(ptr)->metaObject());
}

void* QHPieModelMapper_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHPieModelMapper*>(ptr)->QHPieModelMapper::metaObject());
}

int QHPieModelMapper_ColumnCount(void* ptr)
{
	return static_cast<QHPieModelMapper*>(ptr)->columnCount();
}

int QHPieModelMapper_FirstColumn(void* ptr)
{
	return static_cast<QHPieModelMapper*>(ptr)->firstColumn();
}

int QHPieModelMapper_LabelsRow(void* ptr)
{
	return static_cast<QHPieModelMapper*>(ptr)->labelsRow();
}

int QHPieModelMapper_ValuesRow(void* ptr)
{
	return static_cast<QHPieModelMapper*>(ptr)->valuesRow();
}

class MyQHXYModelMapper: public QHXYModelMapper
{
public:
	MyQHXYModelMapper(QObject *parent = Q_NULLPTR) : QHXYModelMapper(parent) {};
	void Signal_ColumnCountChanged() { callbackQHXYModelMapper_ColumnCountChanged(this); };
	void Signal_FirstColumnChanged() { callbackQHXYModelMapper_FirstColumnChanged(this); };
	void Signal_ModelReplaced() { callbackQHXYModelMapper_ModelReplaced(this); };
	void Signal_SeriesReplaced() { callbackQHXYModelMapper_SeriesReplaced(this); };
	void Signal_XRowChanged() { callbackQHXYModelMapper_XRowChanged(this); };
	void Signal_YRowChanged() { callbackQHXYModelMapper_YRowChanged(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHXYModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QHXYModelMapper_NewQHXYModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHXYModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQHXYModelMapper(static_cast<QObject*>(parent));
	}
}

struct QtCharts_PackedString QHXYModelMapper_QHXYModelMapper_Tr(char* s, char* c, int n)
{
	return ({ QByteArray td40d08 = QHXYModelMapper::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(td40d08.prepend("WHITESPACE").constData()+10), td40d08.size()-10 }; });
}

struct QtCharts_PackedString QHXYModelMapper_QHXYModelMapper_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t9ffaef = QHXYModelMapper::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t9ffaef.prepend("WHITESPACE").constData()+10), t9ffaef.size()-10 }; });
}

void QHXYModelMapper_ConnectColumnCountChanged(void* ptr)
{
	QObject::connect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::columnCountChanged), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_ColumnCountChanged));
}

void QHXYModelMapper_DisconnectColumnCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::columnCountChanged), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_ColumnCountChanged));
}

void QHXYModelMapper_ColumnCountChanged(void* ptr)
{
	static_cast<QHXYModelMapper*>(ptr)->columnCountChanged();
}

void QHXYModelMapper_ConnectFirstColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::firstColumnChanged), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_FirstColumnChanged));
}

void QHXYModelMapper_DisconnectFirstColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::firstColumnChanged), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_FirstColumnChanged));
}

void QHXYModelMapper_FirstColumnChanged(void* ptr)
{
	static_cast<QHXYModelMapper*>(ptr)->firstColumnChanged();
}

void QHXYModelMapper_ConnectModelReplaced(void* ptr)
{
	QObject::connect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::modelReplaced), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_ModelReplaced));
}

void QHXYModelMapper_DisconnectModelReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::modelReplaced), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_ModelReplaced));
}

void QHXYModelMapper_ModelReplaced(void* ptr)
{
	static_cast<QHXYModelMapper*>(ptr)->modelReplaced();
}

void QHXYModelMapper_ConnectSeriesReplaced(void* ptr)
{
	QObject::connect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::seriesReplaced), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_SeriesReplaced));
}

void QHXYModelMapper_DisconnectSeriesReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::seriesReplaced), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_SeriesReplaced));
}

void QHXYModelMapper_SeriesReplaced(void* ptr)
{
	static_cast<QHXYModelMapper*>(ptr)->seriesReplaced();
}

void QHXYModelMapper_SetColumnCount(void* ptr, int columnCount)
{
	static_cast<QHXYModelMapper*>(ptr)->setColumnCount(columnCount);
}

void QHXYModelMapper_SetFirstColumn(void* ptr, int firstColumn)
{
	static_cast<QHXYModelMapper*>(ptr)->setFirstColumn(firstColumn);
}

void QHXYModelMapper_SetModel(void* ptr, void* model)
{
	static_cast<QHXYModelMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QHXYModelMapper_SetSeries(void* ptr, void* series)
{
	static_cast<QHXYModelMapper*>(ptr)->setSeries(static_cast<QXYSeries*>(series));
}

void QHXYModelMapper_SetXRow(void* ptr, int xRow)
{
	static_cast<QHXYModelMapper*>(ptr)->setXRow(xRow);
}

void QHXYModelMapper_SetYRow(void* ptr, int yRow)
{
	static_cast<QHXYModelMapper*>(ptr)->setYRow(yRow);
}

void QHXYModelMapper_ConnectXRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::xRowChanged), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_XRowChanged));
}

void QHXYModelMapper_DisconnectXRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::xRowChanged), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_XRowChanged));
}

void QHXYModelMapper_XRowChanged(void* ptr)
{
	static_cast<QHXYModelMapper*>(ptr)->xRowChanged();
}

void QHXYModelMapper_ConnectYRowChanged(void* ptr)
{
	QObject::connect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::yRowChanged), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_YRowChanged));
}

void QHXYModelMapper_DisconnectYRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHXYModelMapper*>(ptr), static_cast<void (QHXYModelMapper::*)()>(&QHXYModelMapper::yRowChanged), static_cast<MyQHXYModelMapper*>(ptr), static_cast<void (MyQHXYModelMapper::*)()>(&MyQHXYModelMapper::Signal_YRowChanged));
}

void QHXYModelMapper_YRowChanged(void* ptr)
{
	static_cast<QHXYModelMapper*>(ptr)->yRowChanged();
}

void* QHXYModelMapper_Model(void* ptr)
{
	return static_cast<QHXYModelMapper*>(ptr)->model();
}

void* QHXYModelMapper_Series(void* ptr)
{
	return static_cast<QHXYModelMapper*>(ptr)->series();
}

void* QHXYModelMapper_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHXYModelMapper*>(ptr)->metaObject());
}

void* QHXYModelMapper_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHXYModelMapper*>(ptr)->QHXYModelMapper::metaObject());
}

int QHXYModelMapper_ColumnCount(void* ptr)
{
	return static_cast<QHXYModelMapper*>(ptr)->columnCount();
}

int QHXYModelMapper_FirstColumn(void* ptr)
{
	return static_cast<QHXYModelMapper*>(ptr)->firstColumn();
}

int QHXYModelMapper_XRow(void* ptr)
{
	return static_cast<QHXYModelMapper*>(ptr)->xRow();
}

int QHXYModelMapper_YRow(void* ptr)
{
	return static_cast<QHXYModelMapper*>(ptr)->yRow();
}

class MyQHorizontalBarSeries: public QHorizontalBarSeries
{
public:
	MyQHorizontalBarSeries(QObject *parent = Q_NULLPTR) : QHorizontalBarSeries(parent) {QHorizontalBarSeries_QHorizontalBarSeries_QRegisterMetaType();};
	 ~MyQHorizontalBarSeries() { callbackQHorizontalBarSeries_DestroyQHorizontalBarSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQHorizontalBarSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BarsetsAdded(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsAdded(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BarsetsRemoved(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsRemoved(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Clicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_Clicked(this, index, barset); };
	void Signal_CountChanged() { callbackQAbstractBarSeries_CountChanged(this); };
	void Signal_DoubleClicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_DoubleClicked(this, index, barset); };
	void Signal_Hovered(bool status, int index, QBarSet * barset) { callbackQAbstractBarSeries_Hovered(this, status, index, barset); };
	void Signal_LabelsAngleChanged(qreal angle) { callbackQAbstractBarSeries_LabelsAngleChanged(this, angle); };
	void Signal_LabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQAbstractBarSeries_LabelsFormatChanged(this, formatPacked); };
	void Signal_LabelsPositionChanged(QAbstractBarSeries::LabelsPosition position) { callbackQAbstractBarSeries_LabelsPositionChanged(this, position); };
	void Signal_LabelsPrecisionChanged(int precision) { callbackQAbstractBarSeries_LabelsPrecisionChanged(this, precision); };
	void Signal_LabelsVisibleChanged() { callbackQAbstractBarSeries_LabelsVisibleChanged(this); };
	void Signal_Pressed(int index, QBarSet * barset) { callbackQAbstractBarSeries_Pressed(this, index, barset); };
	void Signal_Released(int index, QBarSet * barset) { callbackQAbstractBarSeries_Released(this, index, barset); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQHorizontalBarSeries*)

int QHorizontalBarSeries_QHorizontalBarSeries_QRegisterMetaType(){qRegisterMetaType<QHorizontalBarSeries*>(); return qRegisterMetaType<MyQHorizontalBarSeries*>();}

void* QHorizontalBarSeries_NewQHorizontalBarSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalBarSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQHorizontalBarSeries(static_cast<QObject*>(parent));
	}
}

void QHorizontalBarSeries_DestroyQHorizontalBarSeries(void* ptr)
{
	static_cast<QHorizontalBarSeries*>(ptr)->~QHorizontalBarSeries();
}

void QHorizontalBarSeries_DestroyQHorizontalBarSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QHorizontalBarSeries_Type(void* ptr)
{
	return static_cast<QHorizontalBarSeries*>(ptr)->type();
}

long long QHorizontalBarSeries_TypeDefault(void* ptr)
{
		return static_cast<QHorizontalBarSeries*>(ptr)->QHorizontalBarSeries::type();
}

class MyQHorizontalPercentBarSeries: public QHorizontalPercentBarSeries
{
public:
	MyQHorizontalPercentBarSeries(QObject *parent = Q_NULLPTR) : QHorizontalPercentBarSeries(parent) {QHorizontalPercentBarSeries_QHorizontalPercentBarSeries_QRegisterMetaType();};
	 ~MyQHorizontalPercentBarSeries() { callbackQHorizontalPercentBarSeries_DestroyQHorizontalPercentBarSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQHorizontalPercentBarSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BarsetsAdded(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsAdded(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BarsetsRemoved(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsRemoved(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Clicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_Clicked(this, index, barset); };
	void Signal_CountChanged() { callbackQAbstractBarSeries_CountChanged(this); };
	void Signal_DoubleClicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_DoubleClicked(this, index, barset); };
	void Signal_Hovered(bool status, int index, QBarSet * barset) { callbackQAbstractBarSeries_Hovered(this, status, index, barset); };
	void Signal_LabelsAngleChanged(qreal angle) { callbackQAbstractBarSeries_LabelsAngleChanged(this, angle); };
	void Signal_LabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQAbstractBarSeries_LabelsFormatChanged(this, formatPacked); };
	void Signal_LabelsPositionChanged(QAbstractBarSeries::LabelsPosition position) { callbackQAbstractBarSeries_LabelsPositionChanged(this, position); };
	void Signal_LabelsPrecisionChanged(int precision) { callbackQAbstractBarSeries_LabelsPrecisionChanged(this, precision); };
	void Signal_LabelsVisibleChanged() { callbackQAbstractBarSeries_LabelsVisibleChanged(this); };
	void Signal_Pressed(int index, QBarSet * barset) { callbackQAbstractBarSeries_Pressed(this, index, barset); };
	void Signal_Released(int index, QBarSet * barset) { callbackQAbstractBarSeries_Released(this, index, barset); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQHorizontalPercentBarSeries*)

int QHorizontalPercentBarSeries_QHorizontalPercentBarSeries_QRegisterMetaType(){qRegisterMetaType<QHorizontalPercentBarSeries*>(); return qRegisterMetaType<MyQHorizontalPercentBarSeries*>();}

void* QHorizontalPercentBarSeries_NewQHorizontalPercentBarSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalPercentBarSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQHorizontalPercentBarSeries(static_cast<QObject*>(parent));
	}
}

void QHorizontalPercentBarSeries_DestroyQHorizontalPercentBarSeries(void* ptr)
{
	static_cast<QHorizontalPercentBarSeries*>(ptr)->~QHorizontalPercentBarSeries();
}

void QHorizontalPercentBarSeries_DestroyQHorizontalPercentBarSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QHorizontalPercentBarSeries_Type(void* ptr)
{
	return static_cast<QHorizontalPercentBarSeries*>(ptr)->type();
}

long long QHorizontalPercentBarSeries_TypeDefault(void* ptr)
{
		return static_cast<QHorizontalPercentBarSeries*>(ptr)->QHorizontalPercentBarSeries::type();
}

class MyQHorizontalStackedBarSeries: public QHorizontalStackedBarSeries
{
public:
	MyQHorizontalStackedBarSeries(QObject *parent = Q_NULLPTR) : QHorizontalStackedBarSeries(parent) {QHorizontalStackedBarSeries_QHorizontalStackedBarSeries_QRegisterMetaType();};
	 ~MyQHorizontalStackedBarSeries() { callbackQHorizontalStackedBarSeries_DestroyQHorizontalStackedBarSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQHorizontalStackedBarSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BarsetsAdded(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsAdded(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BarsetsRemoved(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsRemoved(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Clicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_Clicked(this, index, barset); };
	void Signal_CountChanged() { callbackQAbstractBarSeries_CountChanged(this); };
	void Signal_DoubleClicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_DoubleClicked(this, index, barset); };
	void Signal_Hovered(bool status, int index, QBarSet * barset) { callbackQAbstractBarSeries_Hovered(this, status, index, barset); };
	void Signal_LabelsAngleChanged(qreal angle) { callbackQAbstractBarSeries_LabelsAngleChanged(this, angle); };
	void Signal_LabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQAbstractBarSeries_LabelsFormatChanged(this, formatPacked); };
	void Signal_LabelsPositionChanged(QAbstractBarSeries::LabelsPosition position) { callbackQAbstractBarSeries_LabelsPositionChanged(this, position); };
	void Signal_LabelsPrecisionChanged(int precision) { callbackQAbstractBarSeries_LabelsPrecisionChanged(this, precision); };
	void Signal_LabelsVisibleChanged() { callbackQAbstractBarSeries_LabelsVisibleChanged(this); };
	void Signal_Pressed(int index, QBarSet * barset) { callbackQAbstractBarSeries_Pressed(this, index, barset); };
	void Signal_Released(int index, QBarSet * barset) { callbackQAbstractBarSeries_Released(this, index, barset); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQHorizontalStackedBarSeries*)

int QHorizontalStackedBarSeries_QHorizontalStackedBarSeries_QRegisterMetaType(){qRegisterMetaType<QHorizontalStackedBarSeries*>(); return qRegisterMetaType<MyQHorizontalStackedBarSeries*>();}

void* QHorizontalStackedBarSeries_NewQHorizontalStackedBarSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHorizontalStackedBarSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQHorizontalStackedBarSeries(static_cast<QObject*>(parent));
	}
}

void QHorizontalStackedBarSeries_DestroyQHorizontalStackedBarSeries(void* ptr)
{
	static_cast<QHorizontalStackedBarSeries*>(ptr)->~QHorizontalStackedBarSeries();
}

void QHorizontalStackedBarSeries_DestroyQHorizontalStackedBarSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QHorizontalStackedBarSeries_Type(void* ptr)
{
	return static_cast<QHorizontalStackedBarSeries*>(ptr)->type();
}

long long QHorizontalStackedBarSeries_TypeDefault(void* ptr)
{
		return static_cast<QHorizontalStackedBarSeries*>(ptr)->QHorizontalStackedBarSeries::type();
}

class MyQLegend: public QLegend
{
public:
	void Signal_BackgroundVisibleChanged(bool visible) { callbackQLegend_BackgroundVisibleChanged(this, visible); };
	void Signal_BorderColorChanged(QColor color) { callbackQLegend_BorderColorChanged(this, new QColor(color)); };
	void Signal_ColorChanged(QColor color) { callbackQLegend_ColorChanged(this, new QColor(color)); };
	void Signal_FontChanged(QFont font) { callbackQLegend_FontChanged(this, new QFont(font)); };
	void Signal_LabelColorChanged(QColor color) { callbackQLegend_LabelColorChanged(this, new QColor(color)); };
	void Signal_MarkerShapeChanged(QLegend::MarkerShape shape) { callbackQLegend_MarkerShapeChanged(this, shape); };
	void Signal_ReverseMarkersChanged(bool reverseMarkers) { callbackQLegend_ReverseMarkersChanged(this, reverseMarkers); };
	void Signal_ShowToolTipsChanged(bool showToolTips) { callbackQLegend_ShowToolTipsChanged(this, showToolTips); };
	 ~MyQLegend() { callbackQLegend_DestroyQLegend(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLegend_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant & value) { return *static_cast<QVariant*>(callbackQLegend_ItemChange(this, change, const_cast<QVariant*>(&value))); };
	bool close() { return callbackQLegend_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQLegend_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQLegend_FocusNextPrevChild(this, next) != 0; };
	bool sceneEvent(QEvent * event) { return callbackQLegend_SceneEvent(this, event) != 0; };
	bool windowFrameEvent(QEvent * event) { return callbackQLegend_WindowFrameEvent(this, event) != 0; };
	void changeEvent(QEvent * event) { callbackQLegend_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQLegend_CloseEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQLegend_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQLegend_FocusOutEvent(this, event); };
	void Signal_GeometryChanged() { callbackQLegend_GeometryChanged(this); };
	void grabKeyboardEvent(QEvent * event) { callbackQLegend_GrabKeyboardEvent(this, event); };
	void grabMouseEvent(QEvent * event) { callbackQLegend_GrabMouseEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQLegend_HideEvent(this, event); };
	void hoverLeaveEvent(QGraphicsSceneHoverEvent * event) { callbackQLegend_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QGraphicsSceneHoverEvent * event) { callbackQLegend_HoverMoveEvent(this, event); };
	void moveEvent(QGraphicsSceneMoveEvent * event) { callbackQLegend_MoveEvent(this, event); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQLegend_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void paintWindowFrame(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQLegend_PaintWindowFrame(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void polishEvent() { callbackQLegend_PolishEvent(this); };
	void resizeEvent(QGraphicsSceneResizeEvent * event) { callbackQLegend_ResizeEvent(this, event); };
	void setGeometry(const QRectF & rect) { callbackQLegend_SetGeometry(this, const_cast<QRectF*>(&rect)); };
	void showEvent(QShowEvent * event) { callbackQLegend_ShowEvent(this, event); };
	void ungrabKeyboardEvent(QEvent * event) { callbackQLegend_UngrabKeyboardEvent(this, event); };
	void ungrabMouseEvent(QEvent * event) { callbackQLegend_UngrabMouseEvent(this, event); };
	void updateGeometry() { callbackQLegend_UpdateGeometry(this); };
	QPainterPath shape() const { return *static_cast<QPainterPath*>(callbackQLegend_Shape(const_cast<void*>(static_cast<const void*>(this)))); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQLegend_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQLegend_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
	Qt::WindowFrameSection windowFrameSectionAt(const QPointF & pos) const { return static_cast<Qt::WindowFrameSection>(callbackQLegend_WindowFrameSectionAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&pos))); };
	int type() const { return callbackQLegend_Type(const_cast<void*>(static_cast<const void*>(this))); };
	void getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom) const { callbackQLegend_GetContentsMargins(const_cast<void*>(static_cast<const void*>(this)), *left, *top, *right, *bottom); };
	void initStyleOption(QStyleOption * option) const { callbackQLegend_InitStyleOption(const_cast<void*>(static_cast<const void*>(this)), option); };
	void Signal_EnabledChanged() { callbackQLegend_EnabledChanged(this); };
	void Signal_OpacityChanged() { callbackQLegend_OpacityChanged(this); };
	void Signal_ParentChanged() { callbackQLegend_ParentChanged(this); };
	void Signal_RotationChanged() { callbackQLegend_RotationChanged(this); };
	void Signal_ScaleChanged() { callbackQLegend_ScaleChanged(this); };
	void updateMicroFocus() { callbackQLegend_UpdateMicroFocus(this); };
	void Signal_VisibleChanged() { callbackQLegend_VisibleChanged(this); };
	void Signal_XChanged() { callbackQLegend_XChanged(this); };
	void Signal_YChanged() { callbackQLegend_YChanged(this); };
	void Signal_ZChanged() { callbackQLegend_ZChanged(this); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLegend_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLegend_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLegend_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLegend_CustomEvent(this, event); };
	void deleteLater() { callbackQLegend_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLegend_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLegend_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLegend_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLegend_TimerEvent(this, event); };
	bool sceneEventFilter(QGraphicsItem * watched, QEvent * event) { return callbackQLegend_SceneEventFilter(this, watched, event) != 0; };
	void advance(int phase) { callbackQLegend_Advance(this, phase); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQLegend_ContextMenuEvent(this, event); };
	void dragEnterEvent(QGraphicsSceneDragDropEvent * event) { callbackQLegend_DragEnterEvent(this, event); };
	void dragLeaveEvent(QGraphicsSceneDragDropEvent * event) { callbackQLegend_DragLeaveEvent(this, event); };
	void dragMoveEvent(QGraphicsSceneDragDropEvent * event) { callbackQLegend_DragMoveEvent(this, event); };
	void dropEvent(QGraphicsSceneDragDropEvent * event) { callbackQLegend_DropEvent(this, event); };
	void hoverEnterEvent(QGraphicsSceneHoverEvent * event) { callbackQLegend_HoverEnterEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQLegend_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQLegend_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQLegend_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQLegend_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * event) { callbackQLegend_MouseMoveEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQLegend_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQLegend_MouseReleaseEvent(this, event); };
	void wheelEvent(QGraphicsSceneWheelEvent * event) { callbackQLegend_WheelEvent(this, event); };
	QPainterPath opaqueArea() const { return *static_cast<QPainterPath*>(callbackQLegend_OpaqueArea(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQLegend_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool collidesWithItem(const QGraphicsItem * other, Qt::ItemSelectionMode mode) const { return callbackQLegend_CollidesWithItem(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(other), mode) != 0; };
	bool collidesWithPath(const QPainterPath & path, Qt::ItemSelectionMode mode) const { return callbackQLegend_CollidesWithPath(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPainterPath*>(&path), mode) != 0; };
	bool contains(const QPointF & point) const { return callbackQLegend_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	bool isObscuredBy(const QGraphicsItem * item) const { return callbackQLegend_IsObscuredBy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(item)) != 0; };
};

Q_DECLARE_METATYPE(MyQLegend*)

int QLegend_QLegend_QRegisterMetaType(){qRegisterMetaType<QLegend*>(); return qRegisterMetaType<MyQLegend*>();}

void* QLegend_BorderColor(void* ptr)
{
	return new QColor(static_cast<QLegend*>(ptr)->borderColor());
}

void* QLegend_Color(void* ptr)
{
	return new QColor(static_cast<QLegend*>(ptr)->color());
}

struct QtCharts_PackedString QLegend_QLegend_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t3c7212 = QLegend::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t3c7212.prepend("WHITESPACE").constData()+10), t3c7212.size()-10 }; });
}

struct QtCharts_PackedString QLegend_QLegend_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t208a86 = QLegend::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t208a86.prepend("WHITESPACE").constData()+10), t208a86.size()-10 }; });
}

char QLegend_IsAttachedToChart(void* ptr)
{
	return static_cast<QLegend*>(ptr)->isAttachedToChart();
}

char QLegend_ReverseMarkers(void* ptr)
{
	return static_cast<QLegend*>(ptr)->reverseMarkers();
}

void QLegend_AttachToChart(void* ptr)
{
	static_cast<QLegend*>(ptr)->attachToChart();
}

void QLegend_ConnectBackgroundVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(bool)>(&QLegend::backgroundVisibleChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(bool)>(&MyQLegend::Signal_BackgroundVisibleChanged));
}

void QLegend_DisconnectBackgroundVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(bool)>(&QLegend::backgroundVisibleChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(bool)>(&MyQLegend::Signal_BackgroundVisibleChanged));
}

void QLegend_BackgroundVisibleChanged(void* ptr, char visible)
{
	static_cast<QLegend*>(ptr)->backgroundVisibleChanged(visible != 0);
}

void QLegend_ConnectBorderColorChanged(void* ptr)
{
	QObject::connect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QColor)>(&QLegend::borderColorChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QColor)>(&MyQLegend::Signal_BorderColorChanged));
}

void QLegend_DisconnectBorderColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QColor)>(&QLegend::borderColorChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QColor)>(&MyQLegend::Signal_BorderColorChanged));
}

void QLegend_BorderColorChanged(void* ptr, void* color)
{
	static_cast<QLegend*>(ptr)->borderColorChanged(*static_cast<QColor*>(color));
}

void QLegend_ConnectColorChanged(void* ptr)
{
	QObject::connect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QColor)>(&QLegend::colorChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QColor)>(&MyQLegend::Signal_ColorChanged));
}

void QLegend_DisconnectColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QColor)>(&QLegend::colorChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QColor)>(&MyQLegend::Signal_ColorChanged));
}

void QLegend_ColorChanged(void* ptr, void* color)
{
	static_cast<QLegend*>(ptr)->colorChanged(*static_cast<QColor*>(color));
}

void QLegend_DetachFromChart(void* ptr)
{
	static_cast<QLegend*>(ptr)->detachFromChart();
}

void QLegend_ConnectFontChanged(void* ptr)
{
	QObject::connect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QFont)>(&QLegend::fontChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QFont)>(&MyQLegend::Signal_FontChanged));
}

void QLegend_DisconnectFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QFont)>(&QLegend::fontChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QFont)>(&MyQLegend::Signal_FontChanged));
}

void QLegend_FontChanged(void* ptr, void* font)
{
	static_cast<QLegend*>(ptr)->fontChanged(*static_cast<QFont*>(font));
}

void QLegend_ConnectLabelColorChanged(void* ptr)
{
	QObject::connect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QColor)>(&QLegend::labelColorChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QColor)>(&MyQLegend::Signal_LabelColorChanged));
}

void QLegend_DisconnectLabelColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QColor)>(&QLegend::labelColorChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QColor)>(&MyQLegend::Signal_LabelColorChanged));
}

void QLegend_LabelColorChanged(void* ptr, void* color)
{
	static_cast<QLegend*>(ptr)->labelColorChanged(*static_cast<QColor*>(color));
}

void QLegend_ConnectMarkerShapeChanged(void* ptr)
{
	QObject::connect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QLegend::MarkerShape)>(&QLegend::markerShapeChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QLegend::MarkerShape)>(&MyQLegend::Signal_MarkerShapeChanged));
}

void QLegend_DisconnectMarkerShapeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(QLegend::MarkerShape)>(&QLegend::markerShapeChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(QLegend::MarkerShape)>(&MyQLegend::Signal_MarkerShapeChanged));
}

void QLegend_MarkerShapeChanged(void* ptr, long long shape)
{
	static_cast<QLegend*>(ptr)->markerShapeChanged(static_cast<QLegend::MarkerShape>(shape));
}

void QLegend_ConnectReverseMarkersChanged(void* ptr)
{
	QObject::connect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(bool)>(&QLegend::reverseMarkersChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(bool)>(&MyQLegend::Signal_ReverseMarkersChanged));
}

void QLegend_DisconnectReverseMarkersChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(bool)>(&QLegend::reverseMarkersChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(bool)>(&MyQLegend::Signal_ReverseMarkersChanged));
}

void QLegend_ReverseMarkersChanged(void* ptr, char reverseMarkers)
{
	static_cast<QLegend*>(ptr)->reverseMarkersChanged(reverseMarkers != 0);
}

void QLegend_SetAlignment(void* ptr, long long alignment)
{
	static_cast<QLegend*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QLegend_SetBackgroundVisible(void* ptr, char visible)
{
	static_cast<QLegend*>(ptr)->setBackgroundVisible(visible != 0);
}

void QLegend_SetBorderColor(void* ptr, void* color)
{
	static_cast<QLegend*>(ptr)->setBorderColor(*static_cast<QColor*>(color));
}

void QLegend_SetBrush(void* ptr, void* brush)
{
	static_cast<QLegend*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QLegend_SetColor(void* ptr, void* color)
{
	static_cast<QLegend*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QLegend_SetLabelBrush(void* ptr, void* brush)
{
	static_cast<QLegend*>(ptr)->setLabelBrush(*static_cast<QBrush*>(brush));
}

void QLegend_SetLabelColor(void* ptr, void* color)
{
	static_cast<QLegend*>(ptr)->setLabelColor(*static_cast<QColor*>(color));
}

void QLegend_SetMarkerShape(void* ptr, long long shape)
{
	static_cast<QLegend*>(ptr)->setMarkerShape(static_cast<QLegend::MarkerShape>(shape));
}

void QLegend_SetPen(void* ptr, void* pen)
{
	static_cast<QLegend*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QLegend_SetReverseMarkers(void* ptr, char reverseMarkers)
{
	static_cast<QLegend*>(ptr)->setReverseMarkers(reverseMarkers != 0);
}

void QLegend_SetShowToolTips(void* ptr, char show)
{
	static_cast<QLegend*>(ptr)->setShowToolTips(show != 0);
}

void QLegend_ConnectShowToolTipsChanged(void* ptr)
{
	QObject::connect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(bool)>(&QLegend::showToolTipsChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(bool)>(&MyQLegend::Signal_ShowToolTipsChanged));
}

void QLegend_DisconnectShowToolTipsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegend*>(ptr), static_cast<void (QLegend::*)(bool)>(&QLegend::showToolTipsChanged), static_cast<MyQLegend*>(ptr), static_cast<void (MyQLegend::*)(bool)>(&MyQLegend::Signal_ShowToolTipsChanged));
}

void QLegend_ShowToolTipsChanged(void* ptr, char showToolTips)
{
	static_cast<QLegend*>(ptr)->showToolTipsChanged(showToolTips != 0);
}

void QLegend_DestroyQLegend(void* ptr)
{
	static_cast<QLegend*>(ptr)->~QLegend();
}

void QLegend_DestroyQLegendDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QLegend_Brush(void* ptr)
{
	return new QBrush(static_cast<QLegend*>(ptr)->brush());
}

void* QLegend_LabelBrush(void* ptr)
{
	return new QBrush(static_cast<QLegend*>(ptr)->labelBrush());
}

void* QLegend_LabelColor(void* ptr)
{
	return new QColor(static_cast<QLegend*>(ptr)->labelColor());
}

long long QLegend_MarkerShape(void* ptr)
{
	return static_cast<QLegend*>(ptr)->markerShape();
}

struct QtCharts_PackedList QLegend_Markers(void* ptr, void* series)
{
	return ({ QList<QLegendMarker *>* tmpValue = new QList<QLegendMarker *>(static_cast<QLegend*>(ptr)->markers(static_cast<QAbstractSeries*>(series))); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

void* QLegend_Pen(void* ptr)
{
	return new QPen(static_cast<QLegend*>(ptr)->pen());
}

long long QLegend_Alignment(void* ptr)
{
	return static_cast<QLegend*>(ptr)->alignment();
}

char QLegend_IsBackgroundVisible(void* ptr)
{
	return static_cast<QLegend*>(ptr)->isBackgroundVisible();
}

char QLegend_ShowToolTips(void* ptr)
{
	return static_cast<QLegend*>(ptr)->showToolTips();
}

void* QLegend_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QLegend*>(ptr)->QLegend::metaObject());
}

void* QLegend___markers_atList(void* ptr, int i)
{
	return ({QLegendMarker * tmp = static_cast<QList<QLegendMarker *>*>(ptr)->at(i); if (i == static_cast<QList<QLegendMarker *>*>(ptr)->size()-1) { static_cast<QList<QLegendMarker *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___markers_setList(void* ptr, void* i)
{
	static_cast<QList<QLegendMarker *>*>(ptr)->append(static_cast<QLegendMarker*>(i));
}

void* QLegend___markers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QLegendMarker *>();
}

void* QLegend___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QLegend___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QLegend___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QLegend___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QLegend___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QLegend___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* QLegend___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QLegend___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QLegend___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QLegend___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLegend___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLegend___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLegend___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLegend___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLegend___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLegend___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLegend___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QLegend___setTransformations_transformations_atList(void* ptr, int i)
{
	return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___setTransformations_transformations_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
}

void* QLegend___setTransformations_transformations_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsTransform *>();
}

void* QLegend___childItems_atList(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___childItems_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QLegend___childItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QLegend___collidingItems_atList(void* ptr, int i)
{
	return ({QGraphicsItem * tmp = static_cast<QList<QGraphicsItem *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsItem *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___collidingItems_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsItem *>*>(ptr)->append(static_cast<QGraphicsItem*>(i));
}

void* QLegend___collidingItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsItem *>();
}

void* QLegend___transformations_atList(void* ptr, int i)
{
	return ({QGraphicsTransform * tmp = static_cast<QList<QGraphicsTransform *>*>(ptr)->at(i); if (i == static_cast<QList<QGraphicsTransform *>*>(ptr)->size()-1) { static_cast<QList<QGraphicsTransform *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegend___transformations_setList(void* ptr, void* i)
{
	static_cast<QList<QGraphicsTransform *>*>(ptr)->append(static_cast<QGraphicsTransform*>(i));
}

void* QLegend___transformations_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGraphicsTransform *>();
}

void* QLegend_ItemChangeDefault(void* ptr, long long change, void* value)
{
		return new QVariant(static_cast<QLegend*>(ptr)->QLegend::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

char QLegend_CloseDefault(void* ptr)
{
		return static_cast<QLegend*>(ptr)->QLegend::close();
}

char QLegend_EventDefault(void* ptr, void* event)
{
		return static_cast<QLegend*>(ptr)->QLegend::event(static_cast<QEvent*>(event));
}

char QLegend_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QLegend*>(ptr)->QLegend::focusNextPrevChild(next != 0);
}

char QLegend_SceneEventDefault(void* ptr, void* event)
{
		return static_cast<QLegend*>(ptr)->QLegend::sceneEvent(static_cast<QEvent*>(event));
}

char QLegend_WindowFrameEventDefault(void* ptr, void* event)
{
		return static_cast<QLegend*>(ptr)->QLegend::windowFrameEvent(static_cast<QEvent*>(event));
}

void QLegend_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::changeEvent(static_cast<QEvent*>(event));
}

void QLegend_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::closeEvent(static_cast<QCloseEvent*>(event));
}

void QLegend_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QLegend_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QLegend_GrabKeyboardEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::grabKeyboardEvent(static_cast<QEvent*>(event));
}

void QLegend_GrabMouseEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::grabMouseEvent(static_cast<QEvent*>(event));
}

void QLegend_HideEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::hideEvent(static_cast<QHideEvent*>(event));
}

void QLegend_HoverLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QLegend_HoverMoveEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QLegend_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::moveEvent(static_cast<QGraphicsSceneMoveEvent*>(event));
}

void QLegend_PaintDefault(void* ptr, void* painter, void* option, void* widget)
{
		static_cast<QLegend*>(ptr)->QLegend::paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QLegend_PaintWindowFrameDefault(void* ptr, void* painter, void* option, void* widget)
{
		static_cast<QLegend*>(ptr)->QLegend::paintWindowFrame(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QLegend_PolishEventDefault(void* ptr)
{
		static_cast<QLegend*>(ptr)->QLegend::polishEvent();
}

void QLegend_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::resizeEvent(static_cast<QGraphicsSceneResizeEvent*>(event));
}

void QLegend_SetGeometryDefault(void* ptr, void* rect)
{
		static_cast<QLegend*>(ptr)->QLegend::setGeometry(*static_cast<QRectF*>(rect));
}

void QLegend_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::showEvent(static_cast<QShowEvent*>(event));
}

void QLegend_UngrabKeyboardEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::ungrabKeyboardEvent(static_cast<QEvent*>(event));
}

void QLegend_UngrabMouseEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::ungrabMouseEvent(static_cast<QEvent*>(event));
}

void QLegend_UpdateGeometryDefault(void* ptr)
{
		static_cast<QLegend*>(ptr)->QLegend::updateGeometry();
}

void* QLegend_ShapeDefault(void* ptr)
{
		return new QPainterPath(static_cast<QLegend*>(ptr)->QLegend::shape());
}

void* QLegend_BoundingRectDefault(void* ptr)
{
		return ({ QRectF tmpValue = static_cast<QLegend*>(ptr)->QLegend::boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QLegend_SizeHintDefault(void* ptr, long long which, void* constraint)
{
		return ({ QSizeF tmpValue = static_cast<QLegend*>(ptr)->QLegend::sizeHint(static_cast<Qt::SizeHint>(which), *static_cast<QSizeF*>(constraint)); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

long long QLegend_WindowFrameSectionAtDefault(void* ptr, void* pos)
{
		return static_cast<QLegend*>(ptr)->QLegend::windowFrameSectionAt(*static_cast<QPointF*>(pos));
}

int QLegend_TypeDefault(void* ptr)
{
		return static_cast<QLegend*>(ptr)->QLegend::type();
}

void QLegend_GetContentsMarginsDefault(void* ptr, double left, double top, double right, double bottom)
{
		static_cast<QLegend*>(ptr)->QLegend::getContentsMargins(&left, &top, &right, &bottom);
}

void QLegend_InitStyleOptionDefault(void* ptr, void* option)
{
		static_cast<QLegend*>(ptr)->QLegend::initStyleOption(static_cast<QStyleOption*>(option));
}

void QLegend_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QLegend*>(ptr)->QLegend::updateMicroFocus();
}

char QLegend_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QLegend*>(ptr)->QLegend::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QLegend_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::childEvent(static_cast<QChildEvent*>(event));
}

void QLegend_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QLegend*>(ptr)->QLegend::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLegend_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::customEvent(static_cast<QEvent*>(event));
}

void QLegend_DeleteLaterDefault(void* ptr)
{
		static_cast<QLegend*>(ptr)->QLegend::deleteLater();
}

void QLegend_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QLegend*>(ptr)->QLegend::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLegend_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::timerEvent(static_cast<QTimerEvent*>(event));
}

char QLegend_SceneEventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QLegend*>(ptr)->QLegend::sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

void QLegend_AdvanceDefault(void* ptr, int phase)
{
		static_cast<QLegend*>(ptr)->QLegend::advance(phase);
}

void QLegend_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QLegend_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QLegend_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QLegend_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QLegend_DropEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QLegend_HoverEnterEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QLegend_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QLegend_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QLegend_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QLegend_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QLegend_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QLegend_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QLegend_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QLegend_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QLegend*>(ptr)->QLegend::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

void* QLegend_OpaqueAreaDefault(void* ptr)
{
		return new QPainterPath(static_cast<QLegend*>(ptr)->QLegend::opaqueArea());
}

void* QLegend_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QLegend*>(ptr)->QLegend::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QLegend_CollidesWithItemDefault(void* ptr, void* other, long long mode)
{
		return static_cast<QLegend*>(ptr)->QLegend::collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

char QLegend_CollidesWithPathDefault(void* ptr, void* path, long long mode)
{
		return static_cast<QLegend*>(ptr)->QLegend::collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

char QLegend_ContainsDefault(void* ptr, void* point)
{
		return static_cast<QLegend*>(ptr)->QLegend::contains(*static_cast<QPointF*>(point));
}

char QLegend_IsObscuredByDefault(void* ptr, void* item)
{
		return static_cast<QLegend*>(ptr)->QLegend::isObscuredBy(static_cast<QGraphicsItem*>(item));
}

class MyQLegendMarker: public QLegendMarker
{
public:
	QAbstractSeries * series() { return static_cast<QAbstractSeries*>(callbackQLegendMarker_Series(this)); };
	QLegendMarker::LegendMarkerType type() { return static_cast<QLegendMarker::LegendMarkerType>(callbackQLegendMarker_Type(this)); };
	void Signal_BrushChanged() { callbackQLegendMarker_BrushChanged(this); };
	void Signal_Clicked() { callbackQLegendMarker_Clicked(this); };
	void Signal_FontChanged() { callbackQLegendMarker_FontChanged(this); };
	void Signal_Hovered(bool status) { callbackQLegendMarker_Hovered(this, status); };
	void Signal_LabelBrushChanged() { callbackQLegendMarker_LabelBrushChanged(this); };
	void Signal_LabelChanged() { callbackQLegendMarker_LabelChanged(this); };
	void Signal_PenChanged() { callbackQLegendMarker_PenChanged(this); };
	void Signal_ShapeChanged() { callbackQLegendMarker_ShapeChanged(this); };
	void Signal_VisibleChanged() { callbackQLegendMarker_VisibleChanged(this); };
	 ~MyQLegendMarker() { callbackQLegendMarker_DestroyQLegendMarker(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLegendMarker_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQLegendMarker_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLegendMarker_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLegendMarker_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLegendMarker_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLegendMarker_CustomEvent(this, event); };
	void deleteLater() { callbackQLegendMarker_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLegendMarker_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLegendMarker_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLegendMarker_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLegendMarker_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQLegendMarker*)

int QLegendMarker_QLegendMarker_QRegisterMetaType(){qRegisterMetaType<QLegendMarker*>(); return qRegisterMetaType<MyQLegendMarker*>();}

void* QLegendMarker_Series(void* ptr)
{
	return static_cast<QLegendMarker*>(ptr)->series();
}

long long QLegendMarker_Type(void* ptr)
{
	return static_cast<QLegendMarker*>(ptr)->type();
}

struct QtCharts_PackedString QLegendMarker_QLegendMarker_Tr(char* s, char* c, int n)
{
	return ({ QByteArray tb5ab52 = QLegendMarker::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(tb5ab52.prepend("WHITESPACE").constData()+10), tb5ab52.size()-10 }; });
}

struct QtCharts_PackedString QLegendMarker_QLegendMarker_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray te47f1a = QLegendMarker::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(te47f1a.prepend("WHITESPACE").constData()+10), te47f1a.size()-10 }; });
}

void QLegendMarker_ConnectBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::brushChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_BrushChanged));
}

void QLegendMarker_DisconnectBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::brushChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_BrushChanged));
}

void QLegendMarker_BrushChanged(void* ptr)
{
	static_cast<QLegendMarker*>(ptr)->brushChanged();
}

void QLegendMarker_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::clicked), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_Clicked));
}

void QLegendMarker_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::clicked), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_Clicked));
}

void QLegendMarker_Clicked(void* ptr)
{
	static_cast<QLegendMarker*>(ptr)->clicked();
}

void QLegendMarker_ConnectFontChanged(void* ptr)
{
	QObject::connect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::fontChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_FontChanged));
}

void QLegendMarker_DisconnectFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::fontChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_FontChanged));
}

void QLegendMarker_FontChanged(void* ptr)
{
	static_cast<QLegendMarker*>(ptr)->fontChanged();
}

void QLegendMarker_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)(bool)>(&QLegendMarker::hovered), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)(bool)>(&MyQLegendMarker::Signal_Hovered));
}

void QLegendMarker_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)(bool)>(&QLegendMarker::hovered), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)(bool)>(&MyQLegendMarker::Signal_Hovered));
}

void QLegendMarker_Hovered(void* ptr, char status)
{
	static_cast<QLegendMarker*>(ptr)->hovered(status != 0);
}

void QLegendMarker_ConnectLabelBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::labelBrushChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_LabelBrushChanged));
}

void QLegendMarker_DisconnectLabelBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::labelBrushChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_LabelBrushChanged));
}

void QLegendMarker_LabelBrushChanged(void* ptr)
{
	static_cast<QLegendMarker*>(ptr)->labelBrushChanged();
}

void QLegendMarker_ConnectLabelChanged(void* ptr)
{
	QObject::connect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::labelChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_LabelChanged));
}

void QLegendMarker_DisconnectLabelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::labelChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_LabelChanged));
}

void QLegendMarker_LabelChanged(void* ptr)
{
	static_cast<QLegendMarker*>(ptr)->labelChanged();
}

void QLegendMarker_ConnectPenChanged(void* ptr)
{
	QObject::connect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::penChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_PenChanged));
}

void QLegendMarker_DisconnectPenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::penChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_PenChanged));
}

void QLegendMarker_PenChanged(void* ptr)
{
	static_cast<QLegendMarker*>(ptr)->penChanged();
}

void QLegendMarker_SetBrush(void* ptr, void* brush)
{
	static_cast<QLegendMarker*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QLegendMarker_SetFont(void* ptr, void* font)
{
	static_cast<QLegendMarker*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QLegendMarker_SetLabel(void* ptr, struct QtCharts_PackedString label)
{
	static_cast<QLegendMarker*>(ptr)->setLabel(QString::fromUtf8(label.data, label.len));
}

void QLegendMarker_SetLabelBrush(void* ptr, void* brush)
{
	static_cast<QLegendMarker*>(ptr)->setLabelBrush(*static_cast<QBrush*>(brush));
}

void QLegendMarker_SetPen(void* ptr, void* pen)
{
	static_cast<QLegendMarker*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QLegendMarker_SetShape(void* ptr, long long shape)
{
	static_cast<QLegendMarker*>(ptr)->setShape(static_cast<QLegend::MarkerShape>(shape));
}

void QLegendMarker_SetVisible(void* ptr, char visible)
{
	static_cast<QLegendMarker*>(ptr)->setVisible(visible != 0);
}

void QLegendMarker_ConnectShapeChanged(void* ptr)
{
	QObject::connect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::shapeChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_ShapeChanged));
}

void QLegendMarker_DisconnectShapeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::shapeChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_ShapeChanged));
}

void QLegendMarker_ShapeChanged(void* ptr)
{
	static_cast<QLegendMarker*>(ptr)->shapeChanged();
}

void QLegendMarker_ConnectVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::visibleChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_VisibleChanged));
}

void QLegendMarker_DisconnectVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLegendMarker*>(ptr), static_cast<void (QLegendMarker::*)()>(&QLegendMarker::visibleChanged), static_cast<MyQLegendMarker*>(ptr), static_cast<void (MyQLegendMarker::*)()>(&MyQLegendMarker::Signal_VisibleChanged));
}

void QLegendMarker_VisibleChanged(void* ptr)
{
	static_cast<QLegendMarker*>(ptr)->visibleChanged();
}

void QLegendMarker_DestroyQLegendMarker(void* ptr)
{
	static_cast<QLegendMarker*>(ptr)->~QLegendMarker();
}

void QLegendMarker_DestroyQLegendMarkerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QLegendMarker_Brush(void* ptr)
{
	return new QBrush(static_cast<QLegendMarker*>(ptr)->brush());
}

void* QLegendMarker_LabelBrush(void* ptr)
{
	return new QBrush(static_cast<QLegendMarker*>(ptr)->labelBrush());
}

void* QLegendMarker_Font(void* ptr)
{
	return new QFont(static_cast<QLegendMarker*>(ptr)->font());
}

long long QLegendMarker_Shape(void* ptr)
{
	return static_cast<QLegendMarker*>(ptr)->shape();
}

void* QLegendMarker_Pen(void* ptr)
{
	return new QPen(static_cast<QLegendMarker*>(ptr)->pen());
}

struct QtCharts_PackedString QLegendMarker_Label(void* ptr)
{
	return ({ QByteArray t83e390 = static_cast<QLegendMarker*>(ptr)->label().toUtf8(); QtCharts_PackedString { const_cast<char*>(t83e390.prepend("WHITESPACE").constData()+10), t83e390.size()-10 }; });
}

char QLegendMarker_IsVisible(void* ptr)
{
	return static_cast<QLegendMarker*>(ptr)->isVisible();
}

void* QLegendMarker_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QXYLegendMarker*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::metaObject());
	} else if (dynamic_cast<QPieLegendMarker*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::metaObject());
	} else if (dynamic_cast<QCandlestickLegendMarker*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::metaObject());
	} else if (dynamic_cast<QBoxPlotLegendMarker*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::metaObject());
	} else if (dynamic_cast<QBarLegendMarker*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::metaObject());
	} else if (dynamic_cast<QAreaLegendMarker*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QLegendMarker*>(ptr)->QLegendMarker::metaObject());
	}
}

void* QLegendMarker___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QLegendMarker___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QLegendMarker___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QLegendMarker___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegendMarker___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLegendMarker___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLegendMarker___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegendMarker___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLegendMarker___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLegendMarker___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegendMarker___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLegendMarker___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QLegendMarker___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLegendMarker___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QLegendMarker___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QLegendMarker_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QXYLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QPieLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QCandlestickLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QBoxPlotLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QBarLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAreaLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QLegendMarker*>(ptr)->QLegendMarker::event(static_cast<QEvent*>(e));
	}
}

char QLegendMarker_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QXYLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPieLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCandlestickLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBoxPlotLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBarLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAreaLegendMarker*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QLegendMarker*>(ptr)->QLegendMarker::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QLegendMarker_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QXYLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QPieLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QCandlestickLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QBoxPlotLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QBarLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAreaLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QLegendMarker*>(ptr)->QLegendMarker::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QLegendMarker_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QXYLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPieLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCandlestickLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBoxPlotLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBarLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAreaLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QLegendMarker*>(ptr)->QLegendMarker::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QLegendMarker_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QXYLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPieLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCandlestickLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBoxPlotLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QBarLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAreaLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QLegendMarker*>(ptr)->QLegendMarker::customEvent(static_cast<QEvent*>(event));
	}
}

void QLegendMarker_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QXYLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::deleteLater();
	} else if (dynamic_cast<QPieLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::deleteLater();
	} else if (dynamic_cast<QCandlestickLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::deleteLater();
	} else if (dynamic_cast<QBoxPlotLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::deleteLater();
	} else if (dynamic_cast<QBarLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::deleteLater();
	} else if (dynamic_cast<QAreaLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::deleteLater();
	} else {
		static_cast<QLegendMarker*>(ptr)->QLegendMarker::deleteLater();
	}
}

void QLegendMarker_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QXYLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPieLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCandlestickLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBoxPlotLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QBarLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAreaLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QLegendMarker*>(ptr)->QLegendMarker::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QLegendMarker_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QXYLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QPieLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QCandlestickLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QCandlestickLegendMarker*>(ptr)->QCandlestickLegendMarker::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QBoxPlotLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBoxPlotLegendMarker*>(ptr)->QBoxPlotLegendMarker::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QBarLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QBarLegendMarker*>(ptr)->QBarLegendMarker::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAreaLegendMarker*>(static_cast<QObject*>(ptr))) {
		static_cast<QAreaLegendMarker*>(ptr)->QAreaLegendMarker::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QLegendMarker*>(ptr)->QLegendMarker::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

class MyQLineSeries: public QLineSeries
{
public:
	MyQLineSeries(QObject *parent = Q_NULLPTR) : QLineSeries(parent) {QLineSeries_QLineSeries_QRegisterMetaType();};
	 ~MyQLineSeries() { callbackQLineSeries_DestroyQLineSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQLineSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_Clicked(const QPointF & point) { callbackQXYSeries_Clicked(this, const_cast<QPointF*>(&point)); };
	void Signal_ColorChanged(QColor color) { callbackQXYSeries_ColorChanged(this, new QColor(color)); };
	void Signal_DoubleClicked(const QPointF & point) { callbackQXYSeries_DoubleClicked(this, const_cast<QPointF*>(&point)); };
	void Signal_Hovered(const QPointF & point, bool state) { callbackQXYSeries_Hovered(this, const_cast<QPointF*>(&point), state); };
	void Signal_PenChanged(const QPen & pen) { callbackQXYSeries_PenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_PointAdded(int index) { callbackQXYSeries_PointAdded(this, index); };
	void Signal_PointLabelsClippingChanged(bool clipping) { callbackQXYSeries_PointLabelsClippingChanged(this, clipping); };
	void Signal_PointLabelsColorChanged(const QColor & color) { callbackQXYSeries_PointLabelsColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_PointLabelsFontChanged(const QFont & font) { callbackQXYSeries_PointLabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_PointLabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQXYSeries_PointLabelsFormatChanged(this, formatPacked); };
	void Signal_PointLabelsVisibilityChanged(bool visible) { callbackQXYSeries_PointLabelsVisibilityChanged(this, visible); };
	void Signal_PointRemoved(int index) { callbackQXYSeries_PointRemoved(this, index); };
	void Signal_PointReplaced(int index) { callbackQXYSeries_PointReplaced(this, index); };
	void Signal_PointsRemoved(int index, int count) { callbackQXYSeries_PointsRemoved(this, index, count); };
	void Signal_PointsReplaced() { callbackQXYSeries_PointsReplaced(this); };
	void Signal_Pressed(const QPointF & point) { callbackQXYSeries_Pressed(this, const_cast<QPointF*>(&point)); };
	void Signal_Released(const QPointF & point) { callbackQXYSeries_Released(this, const_cast<QPointF*>(&point)); };
	void setBrush(const QBrush & brush) { callbackQXYSeries_SetBrush(this, const_cast<QBrush*>(&brush)); };
	void setColor(const QColor & color) { callbackQXYSeries_SetColor(this, const_cast<QColor*>(&color)); };
	void setPen(const QPen & pen) { callbackQXYSeries_SetPen(this, const_cast<QPen*>(&pen)); };
	QColor color() const { return *static_cast<QColor*>(callbackQXYSeries_Color(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQLineSeries*)

int QLineSeries_QLineSeries_QRegisterMetaType(){qRegisterMetaType<QLineSeries*>(); return qRegisterMetaType<MyQLineSeries*>();}

void* QLineSeries_NewQLineSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLineSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQLineSeries(static_cast<QObject*>(parent));
	}
}

void QLineSeries_DestroyQLineSeries(void* ptr)
{
	static_cast<QLineSeries*>(ptr)->~QLineSeries();
}

void QLineSeries_DestroyQLineSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QLineSeries_Type(void* ptr)
{
	return static_cast<QLineSeries*>(ptr)->type();
}

long long QLineSeries_TypeDefault(void* ptr)
{
	if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplineSeries*>(ptr)->QSplineSeries::type();
	} else {
		return static_cast<QLineSeries*>(ptr)->QLineSeries::type();
	}
}

class MyQLogValueAxis: public QLogValueAxis
{
public:
	MyQLogValueAxis(QObject *parent = Q_NULLPTR) : QLogValueAxis(parent) {QLogValueAxis_QLogValueAxis_QRegisterMetaType();};
	void Signal_BaseChanged(qreal base) { callbackQLogValueAxis_BaseChanged(this, base); };
	void Signal_LabelFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQLogValueAxis_LabelFormatChanged(this, formatPacked); };
	void Signal_MaxChanged(qreal max) { callbackQLogValueAxis_MaxChanged(this, max); };
	void Signal_MinChanged(qreal min) { callbackQLogValueAxis_MinChanged(this, min); };
	void Signal_MinorTickCountChanged(int minorTickCount) { callbackQLogValueAxis_MinorTickCountChanged(this, minorTickCount); };
	void Signal_RangeChanged(qreal min, qreal max) { callbackQLogValueAxis_RangeChanged(this, min, max); };
	void Signal_TickCountChanged(int tickCount) { callbackQLogValueAxis_TickCountChanged(this, tickCount); };
	 ~MyQLogValueAxis() { callbackQLogValueAxis_DestroyQLogValueAxis(this); };
	QAbstractAxis::AxisType type() const { return static_cast<QAbstractAxis::AxisType>(callbackQLogValueAxis_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractAxis_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ColorChanged(QColor color) { callbackQAbstractAxis_ColorChanged(this, new QColor(color)); };
	void Signal_GridLineColorChanged(const QColor & color) { callbackQAbstractAxis_GridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_GridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_GridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_GridVisibleChanged(bool visible) { callbackQAbstractAxis_GridVisibleChanged(this, visible); };
	void Signal_LabelsAngleChanged(int angle) { callbackQAbstractAxis_LabelsAngleChanged(this, angle); };
	void Signal_LabelsBrushChanged(const QBrush & brush) { callbackQAbstractAxis_LabelsBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_LabelsColorChanged(QColor color) { callbackQAbstractAxis_LabelsColorChanged(this, new QColor(color)); };
	void Signal_LabelsFontChanged(const QFont & font) { callbackQAbstractAxis_LabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_LabelsVisibleChanged(bool visible) { callbackQAbstractAxis_LabelsVisibleChanged(this, visible); };
	void Signal_LinePenChanged(const QPen & pen) { callbackQAbstractAxis_LinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_LineVisibleChanged(bool visible) { callbackQAbstractAxis_LineVisibleChanged(this, visible); };
	void Signal_MinorGridLineColorChanged(const QColor & color) { callbackQAbstractAxis_MinorGridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_MinorGridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_MinorGridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_MinorGridVisibleChanged(bool visible) { callbackQAbstractAxis_MinorGridVisibleChanged(this, visible); };
	void Signal_ReverseChanged(bool reverse) { callbackQAbstractAxis_ReverseChanged(this, reverse); };
	void Signal_ShadesBorderColorChanged(QColor color) { callbackQAbstractAxis_ShadesBorderColorChanged(this, new QColor(color)); };
	void Signal_ShadesBrushChanged(const QBrush & brush) { callbackQAbstractAxis_ShadesBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_ShadesColorChanged(QColor color) { callbackQAbstractAxis_ShadesColorChanged(this, new QColor(color)); };
	void Signal_ShadesPenChanged(const QPen & pen) { callbackQAbstractAxis_ShadesPenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_ShadesVisibleChanged(bool visible) { callbackQAbstractAxis_ShadesVisibleChanged(this, visible); };
	void Signal_TitleBrushChanged(const QBrush & brush) { callbackQAbstractAxis_TitleBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_TitleFontChanged(const QFont & font) { callbackQAbstractAxis_TitleFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_TitleTextChanged(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtCharts_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQAbstractAxis_TitleTextChanged(this, textPacked); };
	void Signal_TitleVisibleChanged(bool visible) { callbackQAbstractAxis_TitleVisibleChanged(this, visible); };
	void Signal_VisibleChanged(bool visible) { callbackQAbstractAxis_VisibleChanged(this, visible); };
	bool event(QEvent * e) { return callbackQAbstractAxis_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractAxis_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractAxis_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractAxis_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractAxis_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractAxis_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractAxis_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractAxis_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQLogValueAxis*)

int QLogValueAxis_QLogValueAxis_QRegisterMetaType(){qRegisterMetaType<QLogValueAxis*>(); return qRegisterMetaType<MyQLogValueAxis*>();}

void* QLogValueAxis_NewQLogValueAxis(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLogValueAxis(static_cast<QWindow*>(parent));
	} else {
		return new MyQLogValueAxis(static_cast<QObject*>(parent));
	}
}

void QLogValueAxis_ConnectBaseChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(qreal)>(&QLogValueAxis::baseChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(qreal)>(&MyQLogValueAxis::Signal_BaseChanged));
}

void QLogValueAxis_DisconnectBaseChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(qreal)>(&QLogValueAxis::baseChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(qreal)>(&MyQLogValueAxis::Signal_BaseChanged));
}

void QLogValueAxis_BaseChanged(void* ptr, double base)
{
	static_cast<QLogValueAxis*>(ptr)->baseChanged(base);
}

void QLogValueAxis_ConnectLabelFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(const QString &)>(&QLogValueAxis::labelFormatChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(const QString &)>(&MyQLogValueAxis::Signal_LabelFormatChanged));
}

void QLogValueAxis_DisconnectLabelFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(const QString &)>(&QLogValueAxis::labelFormatChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(const QString &)>(&MyQLogValueAxis::Signal_LabelFormatChanged));
}

void QLogValueAxis_LabelFormatChanged(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QLogValueAxis*>(ptr)->labelFormatChanged(QString::fromUtf8(format.data, format.len));
}

void QLogValueAxis_ConnectMaxChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(qreal)>(&QLogValueAxis::maxChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(qreal)>(&MyQLogValueAxis::Signal_MaxChanged));
}

void QLogValueAxis_DisconnectMaxChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(qreal)>(&QLogValueAxis::maxChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(qreal)>(&MyQLogValueAxis::Signal_MaxChanged));
}

void QLogValueAxis_MaxChanged(void* ptr, double max)
{
	static_cast<QLogValueAxis*>(ptr)->maxChanged(max);
}

void QLogValueAxis_ConnectMinChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(qreal)>(&QLogValueAxis::minChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(qreal)>(&MyQLogValueAxis::Signal_MinChanged));
}

void QLogValueAxis_DisconnectMinChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(qreal)>(&QLogValueAxis::minChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(qreal)>(&MyQLogValueAxis::Signal_MinChanged));
}

void QLogValueAxis_MinChanged(void* ptr, double min)
{
	static_cast<QLogValueAxis*>(ptr)->minChanged(min);
}

void QLogValueAxis_ConnectMinorTickCountChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(int)>(&QLogValueAxis::minorTickCountChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(int)>(&MyQLogValueAxis::Signal_MinorTickCountChanged));
}

void QLogValueAxis_DisconnectMinorTickCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(int)>(&QLogValueAxis::minorTickCountChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(int)>(&MyQLogValueAxis::Signal_MinorTickCountChanged));
}

void QLogValueAxis_MinorTickCountChanged(void* ptr, int minorTickCount)
{
	static_cast<QLogValueAxis*>(ptr)->minorTickCountChanged(minorTickCount);
}

void QLogValueAxis_ConnectRangeChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(qreal, qreal)>(&QLogValueAxis::rangeChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(qreal, qreal)>(&MyQLogValueAxis::Signal_RangeChanged));
}

void QLogValueAxis_DisconnectRangeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(qreal, qreal)>(&QLogValueAxis::rangeChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(qreal, qreal)>(&MyQLogValueAxis::Signal_RangeChanged));
}

void QLogValueAxis_RangeChanged(void* ptr, double min, double max)
{
	static_cast<QLogValueAxis*>(ptr)->rangeChanged(min, max);
}

void QLogValueAxis_SetBase(void* ptr, double base)
{
	static_cast<QLogValueAxis*>(ptr)->setBase(base);
}

void QLogValueAxis_SetLabelFormat(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QLogValueAxis*>(ptr)->setLabelFormat(QString::fromUtf8(format.data, format.len));
}

void QLogValueAxis_SetMax(void* ptr, double max)
{
	static_cast<QLogValueAxis*>(ptr)->setMax(max);
}

void QLogValueAxis_SetMin(void* ptr, double min)
{
	static_cast<QLogValueAxis*>(ptr)->setMin(min);
}

void QLogValueAxis_SetMinorTickCount(void* ptr, int minorTickCount)
{
	static_cast<QLogValueAxis*>(ptr)->setMinorTickCount(minorTickCount);
}

void QLogValueAxis_SetRange(void* ptr, double min, double max)
{
	static_cast<QLogValueAxis*>(ptr)->setRange(min, max);
}

void QLogValueAxis_ConnectTickCountChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(int)>(&QLogValueAxis::tickCountChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(int)>(&MyQLogValueAxis::Signal_TickCountChanged));
}

void QLogValueAxis_DisconnectTickCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValueAxis*>(ptr), static_cast<void (QLogValueAxis::*)(int)>(&QLogValueAxis::tickCountChanged), static_cast<MyQLogValueAxis*>(ptr), static_cast<void (MyQLogValueAxis::*)(int)>(&MyQLogValueAxis::Signal_TickCountChanged));
}

void QLogValueAxis_TickCountChanged(void* ptr, int tickCount)
{
	static_cast<QLogValueAxis*>(ptr)->tickCountChanged(tickCount);
}

void QLogValueAxis_DestroyQLogValueAxis(void* ptr)
{
	static_cast<QLogValueAxis*>(ptr)->~QLogValueAxis();
}

void QLogValueAxis_DestroyQLogValueAxisDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QLogValueAxis_Type(void* ptr)
{
	return static_cast<QLogValueAxis*>(ptr)->type();
}

long long QLogValueAxis_TypeDefault(void* ptr)
{
		return static_cast<QLogValueAxis*>(ptr)->QLogValueAxis::type();
}

struct QtCharts_PackedString QLogValueAxis_LabelFormat(void* ptr)
{
	return ({ QByteArray tb0d38b = static_cast<QLogValueAxis*>(ptr)->labelFormat().toUtf8(); QtCharts_PackedString { const_cast<char*>(tb0d38b.prepend("WHITESPACE").constData()+10), tb0d38b.size()-10 }; });
}

int QLogValueAxis_MinorTickCount(void* ptr)
{
	return static_cast<QLogValueAxis*>(ptr)->minorTickCount();
}

int QLogValueAxis_TickCount(void* ptr)
{
	return static_cast<QLogValueAxis*>(ptr)->tickCount();
}

double QLogValueAxis_Base(void* ptr)
{
	return static_cast<QLogValueAxis*>(ptr)->base();
}

double QLogValueAxis_Max(void* ptr)
{
	return static_cast<QLogValueAxis*>(ptr)->max();
}

double QLogValueAxis_Min(void* ptr)
{
	return static_cast<QLogValueAxis*>(ptr)->min();
}

class MyQPercentBarSeries: public QPercentBarSeries
{
public:
	MyQPercentBarSeries(QObject *parent = Q_NULLPTR) : QPercentBarSeries(parent) {QPercentBarSeries_QPercentBarSeries_QRegisterMetaType();};
	 ~MyQPercentBarSeries() { callbackQPercentBarSeries_DestroyQPercentBarSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQPercentBarSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BarsetsAdded(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsAdded(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BarsetsRemoved(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsRemoved(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Clicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_Clicked(this, index, barset); };
	void Signal_CountChanged() { callbackQAbstractBarSeries_CountChanged(this); };
	void Signal_DoubleClicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_DoubleClicked(this, index, barset); };
	void Signal_Hovered(bool status, int index, QBarSet * barset) { callbackQAbstractBarSeries_Hovered(this, status, index, barset); };
	void Signal_LabelsAngleChanged(qreal angle) { callbackQAbstractBarSeries_LabelsAngleChanged(this, angle); };
	void Signal_LabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQAbstractBarSeries_LabelsFormatChanged(this, formatPacked); };
	void Signal_LabelsPositionChanged(QAbstractBarSeries::LabelsPosition position) { callbackQAbstractBarSeries_LabelsPositionChanged(this, position); };
	void Signal_LabelsPrecisionChanged(int precision) { callbackQAbstractBarSeries_LabelsPrecisionChanged(this, precision); };
	void Signal_LabelsVisibleChanged() { callbackQAbstractBarSeries_LabelsVisibleChanged(this); };
	void Signal_Pressed(int index, QBarSet * barset) { callbackQAbstractBarSeries_Pressed(this, index, barset); };
	void Signal_Released(int index, QBarSet * barset) { callbackQAbstractBarSeries_Released(this, index, barset); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQPercentBarSeries*)

int QPercentBarSeries_QPercentBarSeries_QRegisterMetaType(){qRegisterMetaType<QPercentBarSeries*>(); return qRegisterMetaType<MyQPercentBarSeries*>();}

void* QPercentBarSeries_NewQPercentBarSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPercentBarSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQPercentBarSeries(static_cast<QObject*>(parent));
	}
}

void QPercentBarSeries_DestroyQPercentBarSeries(void* ptr)
{
	static_cast<QPercentBarSeries*>(ptr)->~QPercentBarSeries();
}

void QPercentBarSeries_DestroyQPercentBarSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QPercentBarSeries_Type(void* ptr)
{
	return static_cast<QPercentBarSeries*>(ptr)->type();
}

long long QPercentBarSeries_TypeDefault(void* ptr)
{
		return static_cast<QPercentBarSeries*>(ptr)->QPercentBarSeries::type();
}

class MyQPieLegendMarker: public QPieLegendMarker
{
public:
	QLegendMarker::LegendMarkerType type() { return static_cast<QLegendMarker::LegendMarkerType>(callbackQPieLegendMarker_Type(this)); };
	QPieSeries * series() { return static_cast<QPieSeries*>(callbackQPieLegendMarker_Series(this)); };
	 ~MyQPieLegendMarker() { callbackQPieLegendMarker_DestroyQPieLegendMarker(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLegendMarker_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BrushChanged() { callbackQLegendMarker_BrushChanged(this); };
	void Signal_Clicked() { callbackQLegendMarker_Clicked(this); };
	void Signal_FontChanged() { callbackQLegendMarker_FontChanged(this); };
	void Signal_Hovered(bool status) { callbackQLegendMarker_Hovered(this, status); };
	void Signal_LabelBrushChanged() { callbackQLegendMarker_LabelBrushChanged(this); };
	void Signal_LabelChanged() { callbackQLegendMarker_LabelChanged(this); };
	void Signal_PenChanged() { callbackQLegendMarker_PenChanged(this); };
	void Signal_ShapeChanged() { callbackQLegendMarker_ShapeChanged(this); };
	void Signal_VisibleChanged() { callbackQLegendMarker_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQLegendMarker_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLegendMarker_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLegendMarker_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLegendMarker_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLegendMarker_CustomEvent(this, event); };
	void deleteLater() { callbackQLegendMarker_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLegendMarker_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLegendMarker_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLegendMarker_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLegendMarker_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQPieLegendMarker*)

int QPieLegendMarker_QPieLegendMarker_QRegisterMetaType(){qRegisterMetaType<QPieLegendMarker*>(); return qRegisterMetaType<MyQPieLegendMarker*>();}

long long QPieLegendMarker_Type(void* ptr)
{
	return static_cast<QPieLegendMarker*>(ptr)->type();
}

long long QPieLegendMarker_TypeDefault(void* ptr)
{
		return static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::type();
}

void* QPieLegendMarker_Series(void* ptr)
{
	return static_cast<QPieLegendMarker*>(ptr)->series();
}

void* QPieLegendMarker_SeriesDefault(void* ptr)
{
		return static_cast<QPieLegendMarker*>(ptr)->QPieLegendMarker::series();
}

void* QPieLegendMarker_Slice(void* ptr)
{
	return static_cast<QPieLegendMarker*>(ptr)->slice();
}

void QPieLegendMarker_DestroyQPieLegendMarker(void* ptr)
{
	static_cast<QPieLegendMarker*>(ptr)->~QPieLegendMarker();
}

void QPieLegendMarker_DestroyQPieLegendMarkerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQPieSeries: public QPieSeries
{
public:
	MyQPieSeries(QObject *parent = Q_NULLPTR) : QPieSeries(parent) {QPieSeries_QPieSeries_QRegisterMetaType();};
	void Signal_Added(QList<QPieSlice *> slices) { callbackQPieSeries_Added(this, ({ QList<QPieSlice *>* tmpValue = new QList<QPieSlice *>(slices); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Clicked(QPieSlice * slice) { callbackQPieSeries_Clicked(this, slice); };
	void Signal_CountChanged() { callbackQPieSeries_CountChanged(this); };
	void Signal_DoubleClicked(QPieSlice * slice) { callbackQPieSeries_DoubleClicked(this, slice); };
	void Signal_Hovered(QPieSlice * slice, bool state) { callbackQPieSeries_Hovered(this, slice, state); };
	void Signal_Pressed(QPieSlice * slice) { callbackQPieSeries_Pressed(this, slice); };
	void Signal_Released(QPieSlice * slice) { callbackQPieSeries_Released(this, slice); };
	void Signal_Removed(QList<QPieSlice *> slices) { callbackQPieSeries_Removed(this, ({ QList<QPieSlice *>* tmpValue = new QList<QPieSlice *>(slices); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SumChanged() { callbackQPieSeries_SumChanged(this); };
	 ~MyQPieSeries() { callbackQPieSeries_DestroyQPieSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQPieSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQPieSeries*)

int QPieSeries_QPieSeries_QRegisterMetaType(){qRegisterMetaType<QPieSeries*>(); return qRegisterMetaType<MyQPieSeries*>();}

void* QPieSeries_NewQPieSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPieSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQPieSeries(static_cast<QObject*>(parent));
	}
}

void* QPieSeries_Append3(void* ptr, struct QtCharts_PackedString label, double value)
{
	return static_cast<QPieSeries*>(ptr)->append(QString::fromUtf8(label.data, label.len), value);
}

char QPieSeries_Append2(void* ptr, void* slices)
{
	return static_cast<QPieSeries*>(ptr)->append(({ QList<QPieSlice *>* tmpP = static_cast<QList<QPieSlice *>*>(slices); QList<QPieSlice *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

char QPieSeries_Append(void* ptr, void* slice)
{
	return static_cast<QPieSeries*>(ptr)->append(static_cast<QPieSlice*>(slice));
}

char QPieSeries_Insert(void* ptr, int index, void* slice)
{
	return static_cast<QPieSeries*>(ptr)->insert(index, static_cast<QPieSlice*>(slice));
}

char QPieSeries_Remove(void* ptr, void* slice)
{
	return static_cast<QPieSeries*>(ptr)->remove(static_cast<QPieSlice*>(slice));
}

char QPieSeries_Take(void* ptr, void* slice)
{
	return static_cast<QPieSeries*>(ptr)->take(static_cast<QPieSlice*>(slice));
}

void QPieSeries_ConnectAdded(void* ptr)
{
	QObject::connect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QList<QPieSlice *>)>(&QPieSeries::added), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QList<QPieSlice *>)>(&MyQPieSeries::Signal_Added));
}

void QPieSeries_DisconnectAdded(void* ptr)
{
	QObject::disconnect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QList<QPieSlice *>)>(&QPieSeries::added), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QList<QPieSlice *>)>(&MyQPieSeries::Signal_Added));
}

void QPieSeries_Added(void* ptr, void* slices)
{
	static_cast<QPieSeries*>(ptr)->added(({ QList<QPieSlice *>* tmpP = static_cast<QList<QPieSlice *>*>(slices); QList<QPieSlice *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QPieSeries_Clear(void* ptr)
{
	static_cast<QPieSeries*>(ptr)->clear();
}

void QPieSeries_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *)>(&QPieSeries::clicked), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *)>(&MyQPieSeries::Signal_Clicked));
}

void QPieSeries_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *)>(&QPieSeries::clicked), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *)>(&MyQPieSeries::Signal_Clicked));
}

void QPieSeries_Clicked(void* ptr, void* slice)
{
	static_cast<QPieSeries*>(ptr)->clicked(static_cast<QPieSlice*>(slice));
}

void QPieSeries_ConnectCountChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)()>(&QPieSeries::countChanged), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)()>(&MyQPieSeries::Signal_CountChanged));
}

void QPieSeries_DisconnectCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)()>(&QPieSeries::countChanged), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)()>(&MyQPieSeries::Signal_CountChanged));
}

void QPieSeries_CountChanged(void* ptr)
{
	static_cast<QPieSeries*>(ptr)->countChanged();
}

void QPieSeries_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *)>(&QPieSeries::doubleClicked), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *)>(&MyQPieSeries::Signal_DoubleClicked));
}

void QPieSeries_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *)>(&QPieSeries::doubleClicked), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *)>(&MyQPieSeries::Signal_DoubleClicked));
}

void QPieSeries_DoubleClicked(void* ptr, void* slice)
{
	static_cast<QPieSeries*>(ptr)->doubleClicked(static_cast<QPieSlice*>(slice));
}

void QPieSeries_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *, bool)>(&QPieSeries::hovered), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *, bool)>(&MyQPieSeries::Signal_Hovered));
}

void QPieSeries_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *, bool)>(&QPieSeries::hovered), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *, bool)>(&MyQPieSeries::Signal_Hovered));
}

void QPieSeries_Hovered(void* ptr, void* slice, char state)
{
	static_cast<QPieSeries*>(ptr)->hovered(static_cast<QPieSlice*>(slice), state != 0);
}

void QPieSeries_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *)>(&QPieSeries::pressed), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *)>(&MyQPieSeries::Signal_Pressed));
}

void QPieSeries_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *)>(&QPieSeries::pressed), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *)>(&MyQPieSeries::Signal_Pressed));
}

void QPieSeries_Pressed(void* ptr, void* slice)
{
	static_cast<QPieSeries*>(ptr)->pressed(static_cast<QPieSlice*>(slice));
}

void QPieSeries_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *)>(&QPieSeries::released), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *)>(&MyQPieSeries::Signal_Released));
}

void QPieSeries_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QPieSlice *)>(&QPieSeries::released), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QPieSlice *)>(&MyQPieSeries::Signal_Released));
}

void QPieSeries_Released(void* ptr, void* slice)
{
	static_cast<QPieSeries*>(ptr)->released(static_cast<QPieSlice*>(slice));
}

void QPieSeries_ConnectRemoved(void* ptr)
{
	QObject::connect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QList<QPieSlice *>)>(&QPieSeries::removed), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QList<QPieSlice *>)>(&MyQPieSeries::Signal_Removed));
}

void QPieSeries_DisconnectRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)(QList<QPieSlice *>)>(&QPieSeries::removed), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)(QList<QPieSlice *>)>(&MyQPieSeries::Signal_Removed));
}

void QPieSeries_Removed(void* ptr, void* slices)
{
	static_cast<QPieSeries*>(ptr)->removed(({ QList<QPieSlice *>* tmpP = static_cast<QList<QPieSlice *>*>(slices); QList<QPieSlice *> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QPieSeries_SetHoleSize(void* ptr, double holeSize)
{
	static_cast<QPieSeries*>(ptr)->setHoleSize(holeSize);
}

void QPieSeries_SetHorizontalPosition(void* ptr, double relativePosition)
{
	static_cast<QPieSeries*>(ptr)->setHorizontalPosition(relativePosition);
}

void QPieSeries_SetLabelsPosition(void* ptr, long long position)
{
	static_cast<QPieSeries*>(ptr)->setLabelsPosition(static_cast<QPieSlice::LabelPosition>(position));
}

void QPieSeries_SetLabelsVisible(void* ptr, char visible)
{
	static_cast<QPieSeries*>(ptr)->setLabelsVisible(visible != 0);
}

void QPieSeries_SetPieEndAngle(void* ptr, double angle)
{
	static_cast<QPieSeries*>(ptr)->setPieEndAngle(angle);
}

void QPieSeries_SetPieSize(void* ptr, double relativeSize)
{
	static_cast<QPieSeries*>(ptr)->setPieSize(relativeSize);
}

void QPieSeries_SetPieStartAngle(void* ptr, double startAngle)
{
	static_cast<QPieSeries*>(ptr)->setPieStartAngle(startAngle);
}

void QPieSeries_SetVerticalPosition(void* ptr, double relativePosition)
{
	static_cast<QPieSeries*>(ptr)->setVerticalPosition(relativePosition);
}

void QPieSeries_ConnectSumChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)()>(&QPieSeries::sumChanged), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)()>(&MyQPieSeries::Signal_SumChanged));
}

void QPieSeries_DisconnectSumChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSeries*>(ptr), static_cast<void (QPieSeries::*)()>(&QPieSeries::sumChanged), static_cast<MyQPieSeries*>(ptr), static_cast<void (MyQPieSeries::*)()>(&MyQPieSeries::Signal_SumChanged));
}

void QPieSeries_SumChanged(void* ptr)
{
	static_cast<QPieSeries*>(ptr)->sumChanged();
}

void QPieSeries_DestroyQPieSeries(void* ptr)
{
	static_cast<QPieSeries*>(ptr)->~QPieSeries();
}

void QPieSeries_DestroyQPieSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QPieSeries_Type(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->type();
}

long long QPieSeries_TypeDefault(void* ptr)
{
		return static_cast<QPieSeries*>(ptr)->QPieSeries::type();
}

struct QtCharts_PackedList QPieSeries_Slices(void* ptr)
{
	return ({ QList<QPieSlice *>* tmpValue = new QList<QPieSlice *>(static_cast<QPieSeries*>(ptr)->slices()); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

char QPieSeries_IsEmpty(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->isEmpty();
}

int QPieSeries_Count(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->count();
}

double QPieSeries_HoleSize(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->holeSize();
}

double QPieSeries_HorizontalPosition(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->horizontalPosition();
}

double QPieSeries_PieEndAngle(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->pieEndAngle();
}

double QPieSeries_PieSize(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->pieSize();
}

double QPieSeries_PieStartAngle(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->pieStartAngle();
}

double QPieSeries_Sum(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->sum();
}

double QPieSeries_VerticalPosition(void* ptr)
{
	return static_cast<QPieSeries*>(ptr)->verticalPosition();
}

void* QPieSeries___append_slices_atList2(void* ptr, int i)
{
	return ({QPieSlice * tmp = static_cast<QList<QPieSlice *>*>(ptr)->at(i); if (i == static_cast<QList<QPieSlice *>*>(ptr)->size()-1) { static_cast<QList<QPieSlice *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPieSeries___append_slices_setList2(void* ptr, void* i)
{
	static_cast<QList<QPieSlice *>*>(ptr)->append(static_cast<QPieSlice*>(i));
}

void* QPieSeries___append_slices_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPieSlice *>();
}

void* QPieSeries___added_slices_atList(void* ptr, int i)
{
	return ({QPieSlice * tmp = static_cast<QList<QPieSlice *>*>(ptr)->at(i); if (i == static_cast<QList<QPieSlice *>*>(ptr)->size()-1) { static_cast<QList<QPieSlice *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPieSeries___added_slices_setList(void* ptr, void* i)
{
	static_cast<QList<QPieSlice *>*>(ptr)->append(static_cast<QPieSlice*>(i));
}

void* QPieSeries___added_slices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPieSlice *>();
}

void* QPieSeries___removed_slices_atList(void* ptr, int i)
{
	return ({QPieSlice * tmp = static_cast<QList<QPieSlice *>*>(ptr)->at(i); if (i == static_cast<QList<QPieSlice *>*>(ptr)->size()-1) { static_cast<QList<QPieSlice *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPieSeries___removed_slices_setList(void* ptr, void* i)
{
	static_cast<QList<QPieSlice *>*>(ptr)->append(static_cast<QPieSlice*>(i));
}

void* QPieSeries___removed_slices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPieSlice *>();
}

void* QPieSeries___slices_atList(void* ptr, int i)
{
	return ({QPieSlice * tmp = static_cast<QList<QPieSlice *>*>(ptr)->at(i); if (i == static_cast<QList<QPieSlice *>*>(ptr)->size()-1) { static_cast<QList<QPieSlice *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPieSeries___slices_setList(void* ptr, void* i)
{
	static_cast<QList<QPieSlice *>*>(ptr)->append(static_cast<QPieSlice*>(i));
}

void* QPieSeries___slices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPieSlice *>();
}

class MyQPieSlice: public QPieSlice
{
public:
	MyQPieSlice(QObject *parent = Q_NULLPTR) : QPieSlice(parent) {QPieSlice_QPieSlice_QRegisterMetaType();};
	MyQPieSlice(QString label, qreal value, QObject *parent = Q_NULLPTR) : QPieSlice(label, value, parent) {QPieSlice_QPieSlice_QRegisterMetaType();};
	void Signal_AngleSpanChanged() { callbackQPieSlice_AngleSpanChanged(this); };
	void Signal_BorderColorChanged() { callbackQPieSlice_BorderColorChanged(this); };
	void Signal_BorderWidthChanged() { callbackQPieSlice_BorderWidthChanged(this); };
	void Signal_BrushChanged() { callbackQPieSlice_BrushChanged(this); };
	void Signal_Clicked() { callbackQPieSlice_Clicked(this); };
	void Signal_ColorChanged() { callbackQPieSlice_ColorChanged(this); };
	void Signal_DoubleClicked() { callbackQPieSlice_DoubleClicked(this); };
	void Signal_Hovered(bool state) { callbackQPieSlice_Hovered(this, state); };
	void Signal_LabelBrushChanged() { callbackQPieSlice_LabelBrushChanged(this); };
	void Signal_LabelChanged() { callbackQPieSlice_LabelChanged(this); };
	void Signal_LabelColorChanged() { callbackQPieSlice_LabelColorChanged(this); };
	void Signal_LabelFontChanged() { callbackQPieSlice_LabelFontChanged(this); };
	void Signal_LabelVisibleChanged() { callbackQPieSlice_LabelVisibleChanged(this); };
	void Signal_PenChanged() { callbackQPieSlice_PenChanged(this); };
	void Signal_PercentageChanged() { callbackQPieSlice_PercentageChanged(this); };
	void Signal_Pressed() { callbackQPieSlice_Pressed(this); };
	void Signal_Released() { callbackQPieSlice_Released(this); };
	void Signal_StartAngleChanged() { callbackQPieSlice_StartAngleChanged(this); };
	void Signal_ValueChanged() { callbackQPieSlice_ValueChanged(this); };
	 ~MyQPieSlice() { callbackQPieSlice_DestroyQPieSlice(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPieSlice_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQPieSlice_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQPieSlice_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQPieSlice_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPieSlice_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQPieSlice_CustomEvent(this, event); };
	void deleteLater() { callbackQPieSlice_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQPieSlice_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPieSlice_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQPieSlice_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQPieSlice_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQPieSlice*)

int QPieSlice_QPieSlice_QRegisterMetaType(){qRegisterMetaType<QPieSlice*>(); return qRegisterMetaType<MyQPieSlice*>();}

void* QPieSlice_BorderColor(void* ptr)
{
	return new QColor(static_cast<QPieSlice*>(ptr)->borderColor());
}

void* QPieSlice_Color(void* ptr)
{
	return new QColor(static_cast<QPieSlice*>(ptr)->color());
}

void* QPieSlice_LabelColor(void* ptr)
{
	return new QColor(static_cast<QPieSlice*>(ptr)->labelColor());
}

void* QPieSlice_NewQPieSlice(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(static_cast<QWindow*>(parent));
	} else {
		return new MyQPieSlice(static_cast<QObject*>(parent));
	}
}

void* QPieSlice_NewQPieSlice2(struct QtCharts_PackedString label, double value, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QWindow*>(parent));
	} else {
		return new MyQPieSlice(QString::fromUtf8(label.data, label.len), value, static_cast<QObject*>(parent));
	}
}

long long QPieSlice_LabelPosition(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->labelPosition();
}

struct QtCharts_PackedString QPieSlice_QPieSlice_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t9788e8 = QPieSlice::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t9788e8.prepend("WHITESPACE").constData()+10), t9788e8.size()-10 }; });
}

struct QtCharts_PackedString QPieSlice_QPieSlice_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray ta7f9e8 = QPieSlice::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(ta7f9e8.prepend("WHITESPACE").constData()+10), ta7f9e8.size()-10 }; });
}

int QPieSlice_BorderWidth(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->borderWidth();
}

void QPieSlice_ConnectAngleSpanChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::angleSpanChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_AngleSpanChanged));
}

void QPieSlice_DisconnectAngleSpanChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::angleSpanChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_AngleSpanChanged));
}

void QPieSlice_AngleSpanChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->angleSpanChanged();
}

void QPieSlice_ConnectBorderColorChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::borderColorChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_BorderColorChanged));
}

void QPieSlice_DisconnectBorderColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::borderColorChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_BorderColorChanged));
}

void QPieSlice_BorderColorChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->borderColorChanged();
}

void QPieSlice_ConnectBorderWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::borderWidthChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_BorderWidthChanged));
}

void QPieSlice_DisconnectBorderWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::borderWidthChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_BorderWidthChanged));
}

void QPieSlice_BorderWidthChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->borderWidthChanged();
}

void QPieSlice_ConnectBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::brushChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_BrushChanged));
}

void QPieSlice_DisconnectBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::brushChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_BrushChanged));
}

void QPieSlice_BrushChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->brushChanged();
}

void QPieSlice_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::clicked), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_Clicked));
}

void QPieSlice_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::clicked), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_Clicked));
}

void QPieSlice_Clicked(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->clicked();
}

void QPieSlice_ConnectColorChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::colorChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_ColorChanged));
}

void QPieSlice_DisconnectColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::colorChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_ColorChanged));
}

void QPieSlice_ColorChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->colorChanged();
}

void QPieSlice_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::doubleClicked), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_DoubleClicked));
}

void QPieSlice_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::doubleClicked), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_DoubleClicked));
}

void QPieSlice_DoubleClicked(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->doubleClicked();
}

void QPieSlice_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)(bool)>(&QPieSlice::hovered), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)(bool)>(&MyQPieSlice::Signal_Hovered));
}

void QPieSlice_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)(bool)>(&QPieSlice::hovered), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)(bool)>(&MyQPieSlice::Signal_Hovered));
}

void QPieSlice_Hovered(void* ptr, char state)
{
	static_cast<QPieSlice*>(ptr)->hovered(state != 0);
}

void QPieSlice_ConnectLabelBrushChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelBrushChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelBrushChanged));
}

void QPieSlice_DisconnectLabelBrushChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelBrushChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelBrushChanged));
}

void QPieSlice_LabelBrushChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->labelBrushChanged();
}

void QPieSlice_ConnectLabelChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelChanged));
}

void QPieSlice_DisconnectLabelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelChanged));
}

void QPieSlice_LabelChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->labelChanged();
}

void QPieSlice_ConnectLabelColorChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelColorChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelColorChanged));
}

void QPieSlice_DisconnectLabelColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelColorChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelColorChanged));
}

void QPieSlice_LabelColorChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->labelColorChanged();
}

void QPieSlice_ConnectLabelFontChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelFontChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelFontChanged));
}

void QPieSlice_DisconnectLabelFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelFontChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelFontChanged));
}

void QPieSlice_LabelFontChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->labelFontChanged();
}

void QPieSlice_ConnectLabelVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelVisibleChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelVisibleChanged));
}

void QPieSlice_DisconnectLabelVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::labelVisibleChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_LabelVisibleChanged));
}

void QPieSlice_LabelVisibleChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->labelVisibleChanged();
}

void QPieSlice_ConnectPenChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::penChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_PenChanged));
}

void QPieSlice_DisconnectPenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::penChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_PenChanged));
}

void QPieSlice_PenChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->penChanged();
}

void QPieSlice_ConnectPercentageChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::percentageChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_PercentageChanged));
}

void QPieSlice_DisconnectPercentageChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::percentageChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_PercentageChanged));
}

void QPieSlice_PercentageChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->percentageChanged();
}

void QPieSlice_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::pressed), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_Pressed));
}

void QPieSlice_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::pressed), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_Pressed));
}

void QPieSlice_Pressed(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->pressed();
}

void QPieSlice_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::released), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_Released));
}

void QPieSlice_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::released), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_Released));
}

void QPieSlice_Released(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->released();
}

void QPieSlice_SetBorderColor(void* ptr, void* color)
{
	static_cast<QPieSlice*>(ptr)->setBorderColor(*static_cast<QColor*>(color));
}

void QPieSlice_SetBorderWidth(void* ptr, int width)
{
	static_cast<QPieSlice*>(ptr)->setBorderWidth(width);
}

void QPieSlice_SetBrush(void* ptr, void* brush)
{
	static_cast<QPieSlice*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QPieSlice_SetColor(void* ptr, void* color)
{
	static_cast<QPieSlice*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QPieSlice_SetExplodeDistanceFactor(void* ptr, double factor)
{
	static_cast<QPieSlice*>(ptr)->setExplodeDistanceFactor(factor);
}

void QPieSlice_SetExploded(void* ptr, char exploded)
{
	static_cast<QPieSlice*>(ptr)->setExploded(exploded != 0);
}

void QPieSlice_SetLabel(void* ptr, struct QtCharts_PackedString label)
{
	static_cast<QPieSlice*>(ptr)->setLabel(QString::fromUtf8(label.data, label.len));
}

void QPieSlice_SetLabelArmLengthFactor(void* ptr, double factor)
{
	static_cast<QPieSlice*>(ptr)->setLabelArmLengthFactor(factor);
}

void QPieSlice_SetLabelBrush(void* ptr, void* brush)
{
	static_cast<QPieSlice*>(ptr)->setLabelBrush(*static_cast<QBrush*>(brush));
}

void QPieSlice_SetLabelColor(void* ptr, void* color)
{
	static_cast<QPieSlice*>(ptr)->setLabelColor(*static_cast<QColor*>(color));
}

void QPieSlice_SetLabelFont(void* ptr, void* font)
{
	static_cast<QPieSlice*>(ptr)->setLabelFont(*static_cast<QFont*>(font));
}

void QPieSlice_SetLabelPosition(void* ptr, long long position)
{
	static_cast<QPieSlice*>(ptr)->setLabelPosition(static_cast<QPieSlice::LabelPosition>(position));
}

void QPieSlice_SetLabelVisible(void* ptr, char visible)
{
	static_cast<QPieSlice*>(ptr)->setLabelVisible(visible != 0);
}

void QPieSlice_SetPen(void* ptr, void* pen)
{
	static_cast<QPieSlice*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QPieSlice_SetValue(void* ptr, double value)
{
	static_cast<QPieSlice*>(ptr)->setValue(value);
}

void QPieSlice_ConnectStartAngleChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::startAngleChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_StartAngleChanged));
}

void QPieSlice_DisconnectStartAngleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::startAngleChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_StartAngleChanged));
}

void QPieSlice_StartAngleChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->startAngleChanged();
}

void QPieSlice_ConnectValueChanged(void* ptr)
{
	QObject::connect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::valueChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_ValueChanged));
}

void QPieSlice_DisconnectValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QPieSlice*>(ptr), static_cast<void (QPieSlice::*)()>(&QPieSlice::valueChanged), static_cast<MyQPieSlice*>(ptr), static_cast<void (MyQPieSlice::*)()>(&MyQPieSlice::Signal_ValueChanged));
}

void QPieSlice_ValueChanged(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->valueChanged();
}

void QPieSlice_DestroyQPieSlice(void* ptr)
{
	static_cast<QPieSlice*>(ptr)->~QPieSlice();
}

void QPieSlice_DestroyQPieSliceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QPieSlice_Brush(void* ptr)
{
	return new QBrush(static_cast<QPieSlice*>(ptr)->brush());
}

void* QPieSlice_LabelBrush(void* ptr)
{
	return new QBrush(static_cast<QPieSlice*>(ptr)->labelBrush());
}

void* QPieSlice_LabelFont(void* ptr)
{
	return new QFont(static_cast<QPieSlice*>(ptr)->labelFont());
}

void* QPieSlice_Pen(void* ptr)
{
	return new QPen(static_cast<QPieSlice*>(ptr)->pen());
}

void* QPieSlice_Series(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->series();
}

struct QtCharts_PackedString QPieSlice_Label(void* ptr)
{
	return ({ QByteArray t7101a7 = static_cast<QPieSlice*>(ptr)->label().toUtf8(); QtCharts_PackedString { const_cast<char*>(t7101a7.prepend("WHITESPACE").constData()+10), t7101a7.size()-10 }; });
}

char QPieSlice_IsExploded(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->isExploded();
}

char QPieSlice_IsLabelVisible(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->isLabelVisible();
}

void* QPieSlice_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QPieSlice*>(ptr)->QPieSlice::metaObject());
}

double QPieSlice_AngleSpan(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->angleSpan();
}

double QPieSlice_ExplodeDistanceFactor(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->explodeDistanceFactor();
}

double QPieSlice_LabelArmLengthFactor(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->labelArmLengthFactor();
}

double QPieSlice_Percentage(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->percentage();
}

double QPieSlice_StartAngle(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->startAngle();
}

double QPieSlice_Value(void* ptr)
{
	return static_cast<QPieSlice*>(ptr)->value();
}

void* QPieSlice___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QPieSlice___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QPieSlice___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QPieSlice___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPieSlice___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPieSlice___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPieSlice___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPieSlice___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPieSlice___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPieSlice___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPieSlice___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPieSlice___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPieSlice___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPieSlice___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QPieSlice___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QPieSlice_EventDefault(void* ptr, void* e)
{
		return static_cast<QPieSlice*>(ptr)->QPieSlice::event(static_cast<QEvent*>(e));
}

char QPieSlice_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QPieSlice*>(ptr)->QPieSlice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QPieSlice_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QPieSlice*>(ptr)->QPieSlice::childEvent(static_cast<QChildEvent*>(event));
}

void QPieSlice_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QPieSlice*>(ptr)->QPieSlice::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPieSlice_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QPieSlice*>(ptr)->QPieSlice::customEvent(static_cast<QEvent*>(event));
}

void QPieSlice_DeleteLaterDefault(void* ptr)
{
		static_cast<QPieSlice*>(ptr)->QPieSlice::deleteLater();
}

void QPieSlice_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QPieSlice*>(ptr)->QPieSlice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPieSlice_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QPieSlice*>(ptr)->QPieSlice::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQPolarChart: public QPolarChart
{
public:
	MyQPolarChart(QGraphicsItem *parent = Q_NULLPTR, Qt::WindowFlags wFlags = Qt::WindowFlags()) : QPolarChart(parent, wFlags) {QPolarChart_QPolarChart_QRegisterMetaType();};
	 ~MyQPolarChart() { callbackQPolarChart_DestroyQPolarChart(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQChart_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_PlotAreaChanged(const QRectF & plotArea) { callbackQChart_PlotAreaChanged(this, const_cast<QRectF*>(&plotArea)); };
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant & value) { return *static_cast<QVariant*>(callbackQChart_ItemChange(this, change, const_cast<QVariant*>(&value))); };
	bool close() { return callbackQChart_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQChart_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQChart_FocusNextPrevChild(this, next) != 0; };
	bool sceneEvent(QEvent * event) { return callbackQChart_SceneEvent(this, event) != 0; };
	bool windowFrameEvent(QEvent * event) { return callbackQChart_WindowFrameEvent(this, event) != 0; };
	void changeEvent(QEvent * event) { callbackQChart_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQChart_CloseEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQChart_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQChart_FocusOutEvent(this, event); };
	void Signal_GeometryChanged() { callbackQChart_GeometryChanged(this); };
	void grabKeyboardEvent(QEvent * event) { callbackQChart_GrabKeyboardEvent(this, event); };
	void grabMouseEvent(QEvent * event) { callbackQChart_GrabMouseEvent(this, event); };
	void hideEvent(QHideEvent * event) { callbackQChart_HideEvent(this, event); };
	void hoverLeaveEvent(QGraphicsSceneHoverEvent * event) { callbackQChart_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QGraphicsSceneHoverEvent * event) { callbackQChart_HoverMoveEvent(this, event); };
	void moveEvent(QGraphicsSceneMoveEvent * event) { callbackQChart_MoveEvent(this, event); };
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQChart_Paint(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void paintWindowFrame(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { callbackQChart_PaintWindowFrame(this, painter, const_cast<QStyleOptionGraphicsItem*>(option), widget); };
	void polishEvent() { callbackQChart_PolishEvent(this); };
	void resizeEvent(QGraphicsSceneResizeEvent * event) { callbackQChart_ResizeEvent(this, event); };
	void setGeometry(const QRectF & rect) { callbackQChart_SetGeometry(this, const_cast<QRectF*>(&rect)); };
	void showEvent(QShowEvent * event) { callbackQChart_ShowEvent(this, event); };
	void ungrabKeyboardEvent(QEvent * event) { callbackQChart_UngrabKeyboardEvent(this, event); };
	void ungrabMouseEvent(QEvent * event) { callbackQChart_UngrabMouseEvent(this, event); };
	void updateGeometry() { callbackQChart_UpdateGeometry(this); };
	QPainterPath shape() const { return *static_cast<QPainterPath*>(callbackQChart_Shape(const_cast<void*>(static_cast<const void*>(this)))); };
	QRectF boundingRect() const { return *static_cast<QRectF*>(callbackQChart_BoundingRect(const_cast<void*>(static_cast<const void*>(this)))); };
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF & constraint) const { return *static_cast<QSizeF*>(callbackQChart_SizeHint(const_cast<void*>(static_cast<const void*>(this)), which, const_cast<QSizeF*>(&constraint))); };
	Qt::WindowFrameSection windowFrameSectionAt(const QPointF & pos) const { return static_cast<Qt::WindowFrameSection>(callbackQChart_WindowFrameSectionAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&pos))); };
	int type() const { return callbackQChart_Type(const_cast<void*>(static_cast<const void*>(this))); };
	void getContentsMargins(qreal * left, qreal * top, qreal * right, qreal * bottom) const { callbackQChart_GetContentsMargins(const_cast<void*>(static_cast<const void*>(this)), *left, *top, *right, *bottom); };
	void initStyleOption(QStyleOption * option) const { callbackQChart_InitStyleOption(const_cast<void*>(static_cast<const void*>(this)), option); };
	void Signal_EnabledChanged() { callbackQChart_EnabledChanged(this); };
	void Signal_OpacityChanged() { callbackQChart_OpacityChanged(this); };
	void Signal_ParentChanged() { callbackQChart_ParentChanged(this); };
	void Signal_RotationChanged() { callbackQChart_RotationChanged(this); };
	void Signal_ScaleChanged() { callbackQChart_ScaleChanged(this); };
	void updateMicroFocus() { callbackQChart_UpdateMicroFocus(this); };
	void Signal_VisibleChanged() { callbackQChart_VisibleChanged(this); };
	void Signal_XChanged() { callbackQChart_XChanged(this); };
	void Signal_YChanged() { callbackQChart_YChanged(this); };
	void Signal_ZChanged() { callbackQChart_ZChanged(this); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQChart_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQChart_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQChart_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQChart_CustomEvent(this, event); };
	void deleteLater() { callbackQChart_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQChart_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQChart_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQChart_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQChart_TimerEvent(this, event); };
	bool sceneEventFilter(QGraphicsItem * watched, QEvent * event) { return callbackQChart_SceneEventFilter(this, watched, event) != 0; };
	void advance(int phase) { callbackQChart_Advance(this, phase); };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQChart_ContextMenuEvent(this, event); };
	void dragEnterEvent(QGraphicsSceneDragDropEvent * event) { callbackQChart_DragEnterEvent(this, event); };
	void dragLeaveEvent(QGraphicsSceneDragDropEvent * event) { callbackQChart_DragLeaveEvent(this, event); };
	void dragMoveEvent(QGraphicsSceneDragDropEvent * event) { callbackQChart_DragMoveEvent(this, event); };
	void dropEvent(QGraphicsSceneDragDropEvent * event) { callbackQChart_DropEvent(this, event); };
	void hoverEnterEvent(QGraphicsSceneHoverEvent * event) { callbackQChart_HoverEnterEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQChart_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQChart_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQChart_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQChart_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * event) { callbackQChart_MouseMoveEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQChart_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQChart_MouseReleaseEvent(this, event); };
	void wheelEvent(QGraphicsSceneWheelEvent * event) { callbackQChart_WheelEvent(this, event); };
	QPainterPath opaqueArea() const { return *static_cast<QPainterPath*>(callbackQChart_OpaqueArea(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQChart_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool collidesWithItem(const QGraphicsItem * other, Qt::ItemSelectionMode mode) const { return callbackQChart_CollidesWithItem(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(other), mode) != 0; };
	bool collidesWithPath(const QPainterPath & path, Qt::ItemSelectionMode mode) const { return callbackQChart_CollidesWithPath(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPainterPath*>(&path), mode) != 0; };
	bool contains(const QPointF & point) const { return callbackQChart_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	bool isObscuredBy(const QGraphicsItem * item) const { return callbackQChart_IsObscuredBy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGraphicsItem*>(item)) != 0; };
};

Q_DECLARE_METATYPE(MyQPolarChart*)

int QPolarChart_QPolarChart_QRegisterMetaType(){qRegisterMetaType<QPolarChart*>(); return qRegisterMetaType<MyQPolarChart*>();}

void* QPolarChart_NewQPolarChart(void* parent, long long wFlags)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQPolarChart(static_cast<QGraphicsObject*>(parent), static_cast<Qt::WindowType>(wFlags));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPolarChart(static_cast<QGraphicsWidget*>(parent), static_cast<Qt::WindowType>(wFlags));
	} else {
		return new MyQPolarChart(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
	}
}

long long QPolarChart_QPolarChart_AxisPolarOrientation(void* axis)
{
	return QPolarChart::axisPolarOrientation(static_cast<QAbstractAxis*>(axis));
}

void QPolarChart_AddAxis(void* ptr, void* axis, long long polarOrientation)
{
	static_cast<QPolarChart*>(ptr)->addAxis(static_cast<QAbstractAxis*>(axis), static_cast<QPolarChart::PolarOrientation>(polarOrientation));
}

void QPolarChart_DestroyQPolarChart(void* ptr)
{
	static_cast<QPolarChart*>(ptr)->~QPolarChart();
}

void QPolarChart_DestroyQPolarChartDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtCharts_PackedList QPolarChart_Axes(void* ptr, long long polarOrientation, void* series)
{
	return ({ QList<QAbstractAxis *>* tmpValue = new QList<QAbstractAxis *>(static_cast<QPolarChart*>(ptr)->axes(static_cast<QPolarChart::PolarOrientation>(polarOrientation), static_cast<QAbstractSeries*>(series))); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

class MyQScatterSeries: public QScatterSeries
{
public:
	MyQScatterSeries(QObject *parent = Q_NULLPTR) : QScatterSeries(parent) {QScatterSeries_QScatterSeries_QRegisterMetaType();};
	void Signal_BorderColorChanged(QColor color) { callbackQScatterSeries_BorderColorChanged(this, new QColor(color)); };
	void Signal_ColorChanged(QColor color) { callbackQXYSeries_ColorChanged(this, new QColor(color)); };
	void Signal_MarkerShapeChanged(QScatterSeries::MarkerShape shape) { callbackQScatterSeries_MarkerShapeChanged(this, shape); };
	void Signal_MarkerSizeChanged(qreal size) { callbackQScatterSeries_MarkerSizeChanged(this, size); };
	void setBrush(const QBrush & brush) { callbackQXYSeries_SetBrush(this, const_cast<QBrush*>(&brush)); };
	void setColor(const QColor & color) { callbackQXYSeries_SetColor(this, const_cast<QColor*>(&color)); };
	void setPen(const QPen & pen) { callbackQXYSeries_SetPen(this, const_cast<QPen*>(&pen)); };
	 ~MyQScatterSeries() { callbackQScatterSeries_DestroyQScatterSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQScatterSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	QColor color() const { return *static_cast<QColor*>(callbackQXYSeries_Color(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_Clicked(const QPointF & point) { callbackQXYSeries_Clicked(this, const_cast<QPointF*>(&point)); };
	void Signal_DoubleClicked(const QPointF & point) { callbackQXYSeries_DoubleClicked(this, const_cast<QPointF*>(&point)); };
	void Signal_Hovered(const QPointF & point, bool state) { callbackQXYSeries_Hovered(this, const_cast<QPointF*>(&point), state); };
	void Signal_PenChanged(const QPen & pen) { callbackQXYSeries_PenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_PointAdded(int index) { callbackQXYSeries_PointAdded(this, index); };
	void Signal_PointLabelsClippingChanged(bool clipping) { callbackQXYSeries_PointLabelsClippingChanged(this, clipping); };
	void Signal_PointLabelsColorChanged(const QColor & color) { callbackQXYSeries_PointLabelsColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_PointLabelsFontChanged(const QFont & font) { callbackQXYSeries_PointLabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_PointLabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQXYSeries_PointLabelsFormatChanged(this, formatPacked); };
	void Signal_PointLabelsVisibilityChanged(bool visible) { callbackQXYSeries_PointLabelsVisibilityChanged(this, visible); };
	void Signal_PointRemoved(int index) { callbackQXYSeries_PointRemoved(this, index); };
	void Signal_PointReplaced(int index) { callbackQXYSeries_PointReplaced(this, index); };
	void Signal_PointsRemoved(int index, int count) { callbackQXYSeries_PointsRemoved(this, index, count); };
	void Signal_PointsReplaced() { callbackQXYSeries_PointsReplaced(this); };
	void Signal_Pressed(const QPointF & point) { callbackQXYSeries_Pressed(this, const_cast<QPointF*>(&point)); };
	void Signal_Released(const QPointF & point) { callbackQXYSeries_Released(this, const_cast<QPointF*>(&point)); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQScatterSeries*)

int QScatterSeries_QScatterSeries_QRegisterMetaType(){qRegisterMetaType<QScatterSeries*>(); return qRegisterMetaType<MyQScatterSeries*>();}

void* QScatterSeries_NewQScatterSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScatterSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQScatterSeries(static_cast<QObject*>(parent));
	}
}

void QScatterSeries_ConnectBorderColorChanged(void* ptr)
{
	QObject::connect(static_cast<QScatterSeries*>(ptr), static_cast<void (QScatterSeries::*)(QColor)>(&QScatterSeries::borderColorChanged), static_cast<MyQScatterSeries*>(ptr), static_cast<void (MyQScatterSeries::*)(QColor)>(&MyQScatterSeries::Signal_BorderColorChanged));
}

void QScatterSeries_DisconnectBorderColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScatterSeries*>(ptr), static_cast<void (QScatterSeries::*)(QColor)>(&QScatterSeries::borderColorChanged), static_cast<MyQScatterSeries*>(ptr), static_cast<void (MyQScatterSeries::*)(QColor)>(&MyQScatterSeries::Signal_BorderColorChanged));
}

void QScatterSeries_BorderColorChanged(void* ptr, void* color)
{
	static_cast<QScatterSeries*>(ptr)->borderColorChanged(*static_cast<QColor*>(color));
}

void QScatterSeries_ConnectMarkerShapeChanged(void* ptr)
{
	QObject::connect(static_cast<QScatterSeries*>(ptr), static_cast<void (QScatterSeries::*)(QScatterSeries::MarkerShape)>(&QScatterSeries::markerShapeChanged), static_cast<MyQScatterSeries*>(ptr), static_cast<void (MyQScatterSeries::*)(QScatterSeries::MarkerShape)>(&MyQScatterSeries::Signal_MarkerShapeChanged));
}

void QScatterSeries_DisconnectMarkerShapeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScatterSeries*>(ptr), static_cast<void (QScatterSeries::*)(QScatterSeries::MarkerShape)>(&QScatterSeries::markerShapeChanged), static_cast<MyQScatterSeries*>(ptr), static_cast<void (MyQScatterSeries::*)(QScatterSeries::MarkerShape)>(&MyQScatterSeries::Signal_MarkerShapeChanged));
}

void QScatterSeries_MarkerShapeChanged(void* ptr, long long shape)
{
	static_cast<QScatterSeries*>(ptr)->markerShapeChanged(static_cast<QScatterSeries::MarkerShape>(shape));
}

void QScatterSeries_ConnectMarkerSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QScatterSeries*>(ptr), static_cast<void (QScatterSeries::*)(qreal)>(&QScatterSeries::markerSizeChanged), static_cast<MyQScatterSeries*>(ptr), static_cast<void (MyQScatterSeries::*)(qreal)>(&MyQScatterSeries::Signal_MarkerSizeChanged));
}

void QScatterSeries_DisconnectMarkerSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScatterSeries*>(ptr), static_cast<void (QScatterSeries::*)(qreal)>(&QScatterSeries::markerSizeChanged), static_cast<MyQScatterSeries*>(ptr), static_cast<void (MyQScatterSeries::*)(qreal)>(&MyQScatterSeries::Signal_MarkerSizeChanged));
}

void QScatterSeries_MarkerSizeChanged(void* ptr, double size)
{
	static_cast<QScatterSeries*>(ptr)->markerSizeChanged(size);
}

void QScatterSeries_SetBorderColor(void* ptr, void* color)
{
	static_cast<QScatterSeries*>(ptr)->setBorderColor(*static_cast<QColor*>(color));
}

void QScatterSeries_SetMarkerShape(void* ptr, long long shape)
{
	static_cast<QScatterSeries*>(ptr)->setMarkerShape(static_cast<QScatterSeries::MarkerShape>(shape));
}

void QScatterSeries_SetMarkerSize(void* ptr, double size)
{
	static_cast<QScatterSeries*>(ptr)->setMarkerSize(size);
}

void QScatterSeries_DestroyQScatterSeries(void* ptr)
{
	static_cast<QScatterSeries*>(ptr)->~QScatterSeries();
}

void QScatterSeries_DestroyQScatterSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QScatterSeries_Type(void* ptr)
{
	return static_cast<QScatterSeries*>(ptr)->type();
}

long long QScatterSeries_TypeDefault(void* ptr)
{
		return static_cast<QScatterSeries*>(ptr)->QScatterSeries::type();
}

void* QScatterSeries_BorderColor(void* ptr)
{
	return new QColor(static_cast<QScatterSeries*>(ptr)->borderColor());
}

long long QScatterSeries_MarkerShape(void* ptr)
{
	return static_cast<QScatterSeries*>(ptr)->markerShape();
}

double QScatterSeries_MarkerSize(void* ptr)
{
	return static_cast<QScatterSeries*>(ptr)->markerSize();
}

class MyQSplineSeries: public QSplineSeries
{
public:
	MyQSplineSeries(QObject *parent = Q_NULLPTR) : QSplineSeries(parent) {QSplineSeries_QSplineSeries_QRegisterMetaType();};
	 ~MyQSplineSeries() { callbackQSplineSeries_DestroyQSplineSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQLineSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_Clicked(const QPointF & point) { callbackQXYSeries_Clicked(this, const_cast<QPointF*>(&point)); };
	void Signal_ColorChanged(QColor color) { callbackQXYSeries_ColorChanged(this, new QColor(color)); };
	void Signal_DoubleClicked(const QPointF & point) { callbackQXYSeries_DoubleClicked(this, const_cast<QPointF*>(&point)); };
	void Signal_Hovered(const QPointF & point, bool state) { callbackQXYSeries_Hovered(this, const_cast<QPointF*>(&point), state); };
	void Signal_PenChanged(const QPen & pen) { callbackQXYSeries_PenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_PointAdded(int index) { callbackQXYSeries_PointAdded(this, index); };
	void Signal_PointLabelsClippingChanged(bool clipping) { callbackQXYSeries_PointLabelsClippingChanged(this, clipping); };
	void Signal_PointLabelsColorChanged(const QColor & color) { callbackQXYSeries_PointLabelsColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_PointLabelsFontChanged(const QFont & font) { callbackQXYSeries_PointLabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_PointLabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQXYSeries_PointLabelsFormatChanged(this, formatPacked); };
	void Signal_PointLabelsVisibilityChanged(bool visible) { callbackQXYSeries_PointLabelsVisibilityChanged(this, visible); };
	void Signal_PointRemoved(int index) { callbackQXYSeries_PointRemoved(this, index); };
	void Signal_PointReplaced(int index) { callbackQXYSeries_PointReplaced(this, index); };
	void Signal_PointsRemoved(int index, int count) { callbackQXYSeries_PointsRemoved(this, index, count); };
	void Signal_PointsReplaced() { callbackQXYSeries_PointsReplaced(this); };
	void Signal_Pressed(const QPointF & point) { callbackQXYSeries_Pressed(this, const_cast<QPointF*>(&point)); };
	void Signal_Released(const QPointF & point) { callbackQXYSeries_Released(this, const_cast<QPointF*>(&point)); };
	void setBrush(const QBrush & brush) { callbackQXYSeries_SetBrush(this, const_cast<QBrush*>(&brush)); };
	void setColor(const QColor & color) { callbackQXYSeries_SetColor(this, const_cast<QColor*>(&color)); };
	void setPen(const QPen & pen) { callbackQXYSeries_SetPen(this, const_cast<QPen*>(&pen)); };
	QColor color() const { return *static_cast<QColor*>(callbackQXYSeries_Color(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQSplineSeries*)

int QSplineSeries_QSplineSeries_QRegisterMetaType(){qRegisterMetaType<QSplineSeries*>(); return qRegisterMetaType<MyQSplineSeries*>();}

void* QSplineSeries_NewQSplineSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSplineSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQSplineSeries(static_cast<QObject*>(parent));
	}
}

void QSplineSeries_DestroyQSplineSeries(void* ptr)
{
	static_cast<QSplineSeries*>(ptr)->~QSplineSeries();
}

void QSplineSeries_DestroyQSplineSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQStackedBarSeries: public QStackedBarSeries
{
public:
	MyQStackedBarSeries(QObject *parent = Q_NULLPTR) : QStackedBarSeries(parent) {QStackedBarSeries_QStackedBarSeries_QRegisterMetaType();};
	 ~MyQStackedBarSeries() { callbackQStackedBarSeries_DestroyQStackedBarSeries(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQStackedBarSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BarsetsAdded(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsAdded(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BarsetsRemoved(QList<QBarSet *> sets) { callbackQAbstractBarSeries_BarsetsRemoved(this, ({ QList<QBarSet *>* tmpValue = new QList<QBarSet *>(sets); QtCharts_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Clicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_Clicked(this, index, barset); };
	void Signal_CountChanged() { callbackQAbstractBarSeries_CountChanged(this); };
	void Signal_DoubleClicked(int index, QBarSet * barset) { callbackQAbstractBarSeries_DoubleClicked(this, index, barset); };
	void Signal_Hovered(bool status, int index, QBarSet * barset) { callbackQAbstractBarSeries_Hovered(this, status, index, barset); };
	void Signal_LabelsAngleChanged(qreal angle) { callbackQAbstractBarSeries_LabelsAngleChanged(this, angle); };
	void Signal_LabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQAbstractBarSeries_LabelsFormatChanged(this, formatPacked); };
	void Signal_LabelsPositionChanged(QAbstractBarSeries::LabelsPosition position) { callbackQAbstractBarSeries_LabelsPositionChanged(this, position); };
	void Signal_LabelsPrecisionChanged(int precision) { callbackQAbstractBarSeries_LabelsPrecisionChanged(this, precision); };
	void Signal_LabelsVisibleChanged() { callbackQAbstractBarSeries_LabelsVisibleChanged(this); };
	void Signal_Pressed(int index, QBarSet * barset) { callbackQAbstractBarSeries_Pressed(this, index, barset); };
	void Signal_Released(int index, QBarSet * barset) { callbackQAbstractBarSeries_Released(this, index, barset); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQStackedBarSeries*)

int QStackedBarSeries_QStackedBarSeries_QRegisterMetaType(){qRegisterMetaType<QStackedBarSeries*>(); return qRegisterMetaType<MyQStackedBarSeries*>();}

void* QStackedBarSeries_NewQStackedBarSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQStackedBarSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQStackedBarSeries(static_cast<QObject*>(parent));
	}
}

void QStackedBarSeries_DestroyQStackedBarSeries(void* ptr)
{
	static_cast<QStackedBarSeries*>(ptr)->~QStackedBarSeries();
}

void QStackedBarSeries_DestroyQStackedBarSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QStackedBarSeries_Type(void* ptr)
{
	return static_cast<QStackedBarSeries*>(ptr)->type();
}

long long QStackedBarSeries_TypeDefault(void* ptr)
{
		return static_cast<QStackedBarSeries*>(ptr)->QStackedBarSeries::type();
}

class MyQVBarModelMapper: public QVBarModelMapper
{
public:
	MyQVBarModelMapper(QObject *parent = Q_NULLPTR) : QVBarModelMapper(parent) {QVBarModelMapper_QVBarModelMapper_QRegisterMetaType();};
	void Signal_FirstBarSetColumnChanged() { callbackQVBarModelMapper_FirstBarSetColumnChanged(this); };
	void Signal_FirstRowChanged() { callbackQVBarModelMapper_FirstRowChanged(this); };
	void Signal_LastBarSetColumnChanged() { callbackQVBarModelMapper_LastBarSetColumnChanged(this); };
	void Signal_ModelReplaced() { callbackQVBarModelMapper_ModelReplaced(this); };
	void Signal_RowCountChanged() { callbackQVBarModelMapper_RowCountChanged(this); };
	void Signal_SeriesReplaced() { callbackQVBarModelMapper_SeriesReplaced(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVBarModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQVBarModelMapper_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVBarModelMapper_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQVBarModelMapper_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVBarModelMapper_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVBarModelMapper_CustomEvent(this, event); };
	void deleteLater() { callbackQVBarModelMapper_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQVBarModelMapper_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVBarModelMapper_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQVBarModelMapper_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQVBarModelMapper_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQVBarModelMapper*)

int QVBarModelMapper_QVBarModelMapper_QRegisterMetaType(){qRegisterMetaType<QVBarModelMapper*>(); return qRegisterMetaType<MyQVBarModelMapper*>();}

struct QtCharts_PackedString QVBarModelMapper_QVBarModelMapper_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t58ecd7 = QVBarModelMapper::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t58ecd7.prepend("WHITESPACE").constData()+10), t58ecd7.size()-10 }; });
}

struct QtCharts_PackedString QVBarModelMapper_QVBarModelMapper_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray te6614e = QVBarModelMapper::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(te6614e.prepend("WHITESPACE").constData()+10), te6614e.size()-10 }; });
}

void* QVBarModelMapper_NewQVBarModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVBarModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQVBarModelMapper(static_cast<QObject*>(parent));
	}
}

void QVBarModelMapper_ConnectFirstBarSetColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::firstBarSetColumnChanged), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_FirstBarSetColumnChanged));
}

void QVBarModelMapper_DisconnectFirstBarSetColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::firstBarSetColumnChanged), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_FirstBarSetColumnChanged));
}

void QVBarModelMapper_FirstBarSetColumnChanged(void* ptr)
{
	static_cast<QVBarModelMapper*>(ptr)->firstBarSetColumnChanged();
}

void QVBarModelMapper_ConnectFirstRowChanged(void* ptr)
{
	QObject::connect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::firstRowChanged), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_FirstRowChanged));
}

void QVBarModelMapper_DisconnectFirstRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::firstRowChanged), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_FirstRowChanged));
}

void QVBarModelMapper_FirstRowChanged(void* ptr)
{
	static_cast<QVBarModelMapper*>(ptr)->firstRowChanged();
}

void QVBarModelMapper_ConnectLastBarSetColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::lastBarSetColumnChanged), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_LastBarSetColumnChanged));
}

void QVBarModelMapper_DisconnectLastBarSetColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::lastBarSetColumnChanged), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_LastBarSetColumnChanged));
}

void QVBarModelMapper_LastBarSetColumnChanged(void* ptr)
{
	static_cast<QVBarModelMapper*>(ptr)->lastBarSetColumnChanged();
}

void QVBarModelMapper_ConnectModelReplaced(void* ptr)
{
	QObject::connect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::modelReplaced), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_ModelReplaced));
}

void QVBarModelMapper_DisconnectModelReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::modelReplaced), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_ModelReplaced));
}

void QVBarModelMapper_ModelReplaced(void* ptr)
{
	static_cast<QVBarModelMapper*>(ptr)->modelReplaced();
}

void QVBarModelMapper_ConnectRowCountChanged(void* ptr)
{
	QObject::connect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::rowCountChanged), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_RowCountChanged));
}

void QVBarModelMapper_DisconnectRowCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::rowCountChanged), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_RowCountChanged));
}

void QVBarModelMapper_RowCountChanged(void* ptr)
{
	static_cast<QVBarModelMapper*>(ptr)->rowCountChanged();
}

void QVBarModelMapper_ConnectSeriesReplaced(void* ptr)
{
	QObject::connect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::seriesReplaced), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_SeriesReplaced));
}

void QVBarModelMapper_DisconnectSeriesReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QVBarModelMapper*>(ptr), static_cast<void (QVBarModelMapper::*)()>(&QVBarModelMapper::seriesReplaced), static_cast<MyQVBarModelMapper*>(ptr), static_cast<void (MyQVBarModelMapper::*)()>(&MyQVBarModelMapper::Signal_SeriesReplaced));
}

void QVBarModelMapper_SeriesReplaced(void* ptr)
{
	static_cast<QVBarModelMapper*>(ptr)->seriesReplaced();
}

void QVBarModelMapper_SetFirstBarSetColumn(void* ptr, int firstBarSetColumn)
{
	static_cast<QVBarModelMapper*>(ptr)->setFirstBarSetColumn(firstBarSetColumn);
}

void QVBarModelMapper_SetFirstRow(void* ptr, int firstRow)
{
	static_cast<QVBarModelMapper*>(ptr)->setFirstRow(firstRow);
}

void QVBarModelMapper_SetLastBarSetColumn(void* ptr, int lastBarSetColumn)
{
	static_cast<QVBarModelMapper*>(ptr)->setLastBarSetColumn(lastBarSetColumn);
}

void QVBarModelMapper_SetModel(void* ptr, void* model)
{
	static_cast<QVBarModelMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QVBarModelMapper_SetRowCount(void* ptr, int rowCount)
{
	static_cast<QVBarModelMapper*>(ptr)->setRowCount(rowCount);
}

void QVBarModelMapper_SetSeries(void* ptr, void* series)
{
	static_cast<QVBarModelMapper*>(ptr)->setSeries(static_cast<QAbstractBarSeries*>(series));
}

void* QVBarModelMapper_Series(void* ptr)
{
	return static_cast<QVBarModelMapper*>(ptr)->series();
}

void* QVBarModelMapper_Model(void* ptr)
{
	return static_cast<QVBarModelMapper*>(ptr)->model();
}

void* QVBarModelMapper_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVBarModelMapper*>(ptr)->QVBarModelMapper::metaObject());
}

int QVBarModelMapper_FirstBarSetColumn(void* ptr)
{
	return static_cast<QVBarModelMapper*>(ptr)->firstBarSetColumn();
}

int QVBarModelMapper_FirstRow(void* ptr)
{
	return static_cast<QVBarModelMapper*>(ptr)->firstRow();
}

int QVBarModelMapper_LastBarSetColumn(void* ptr)
{
	return static_cast<QVBarModelMapper*>(ptr)->lastBarSetColumn();
}

int QVBarModelMapper_RowCount(void* ptr)
{
	return static_cast<QVBarModelMapper*>(ptr)->rowCount();
}

void* QVBarModelMapper___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVBarModelMapper___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QVBarModelMapper___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QVBarModelMapper___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVBarModelMapper___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBarModelMapper___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVBarModelMapper___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVBarModelMapper___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBarModelMapper___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVBarModelMapper___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVBarModelMapper___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBarModelMapper___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVBarModelMapper___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVBarModelMapper___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBarModelMapper___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QVBarModelMapper_EventDefault(void* ptr, void* e)
{
		return static_cast<QVBarModelMapper*>(ptr)->QVBarModelMapper::event(static_cast<QEvent*>(e));
}

char QVBarModelMapper_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QVBarModelMapper*>(ptr)->QVBarModelMapper::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QVBarModelMapper_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QVBarModelMapper*>(ptr)->QVBarModelMapper::childEvent(static_cast<QChildEvent*>(event));
}

void QVBarModelMapper_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVBarModelMapper*>(ptr)->QVBarModelMapper::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVBarModelMapper_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QVBarModelMapper*>(ptr)->QVBarModelMapper::customEvent(static_cast<QEvent*>(event));
}

void QVBarModelMapper_DeleteLaterDefault(void* ptr)
{
		static_cast<QVBarModelMapper*>(ptr)->QVBarModelMapper::deleteLater();
}

void QVBarModelMapper_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVBarModelMapper*>(ptr)->QVBarModelMapper::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVBarModelMapper_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QVBarModelMapper*>(ptr)->QVBarModelMapper::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQVBoxPlotModelMapper: public QVBoxPlotModelMapper
{
public:
	MyQVBoxPlotModelMapper(QObject *parent = Q_NULLPTR) : QVBoxPlotModelMapper(parent) {QVBoxPlotModelMapper_QVBoxPlotModelMapper_QRegisterMetaType();};
	void Signal_FirstBoxSetColumnChanged() { callbackQVBoxPlotModelMapper_FirstBoxSetColumnChanged(this); };
	void Signal_FirstRowChanged() { callbackQVBoxPlotModelMapper_FirstRowChanged(this); };
	void Signal_LastBoxSetColumnChanged() { callbackQVBoxPlotModelMapper_LastBoxSetColumnChanged(this); };
	void Signal_ModelReplaced() { callbackQVBoxPlotModelMapper_ModelReplaced(this); };
	void Signal_RowCountChanged() { callbackQVBoxPlotModelMapper_RowCountChanged(this); };
	void Signal_SeriesReplaced() { callbackQVBoxPlotModelMapper_SeriesReplaced(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVBoxPlotModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQVBoxPlotModelMapper_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVBoxPlotModelMapper_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQVBoxPlotModelMapper_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVBoxPlotModelMapper_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVBoxPlotModelMapper_CustomEvent(this, event); };
	void deleteLater() { callbackQVBoxPlotModelMapper_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQVBoxPlotModelMapper_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVBoxPlotModelMapper_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQVBoxPlotModelMapper_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQVBoxPlotModelMapper_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQVBoxPlotModelMapper*)

int QVBoxPlotModelMapper_QVBoxPlotModelMapper_QRegisterMetaType(){qRegisterMetaType<QVBoxPlotModelMapper*>(); return qRegisterMetaType<MyQVBoxPlotModelMapper*>();}

struct QtCharts_PackedString QVBoxPlotModelMapper_QVBoxPlotModelMapper_Tr(char* s, char* c, int n)
{
	return ({ QByteArray td7141a = QVBoxPlotModelMapper::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(td7141a.prepend("WHITESPACE").constData()+10), td7141a.size()-10 }; });
}

struct QtCharts_PackedString QVBoxPlotModelMapper_QVBoxPlotModelMapper_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t8ee883 = QVBoxPlotModelMapper::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t8ee883.prepend("WHITESPACE").constData()+10), t8ee883.size()-10 }; });
}

void* QVBoxPlotModelMapper_NewQVBoxPlotModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVBoxPlotModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQVBoxPlotModelMapper(static_cast<QObject*>(parent));
	}
}

void QVBoxPlotModelMapper_ConnectFirstBoxSetColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::firstBoxSetColumnChanged), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_FirstBoxSetColumnChanged));
}

void QVBoxPlotModelMapper_DisconnectFirstBoxSetColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::firstBoxSetColumnChanged), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_FirstBoxSetColumnChanged));
}

void QVBoxPlotModelMapper_FirstBoxSetColumnChanged(void* ptr)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->firstBoxSetColumnChanged();
}

void QVBoxPlotModelMapper_ConnectFirstRowChanged(void* ptr)
{
	QObject::connect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::firstRowChanged), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_FirstRowChanged));
}

void QVBoxPlotModelMapper_DisconnectFirstRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::firstRowChanged), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_FirstRowChanged));
}

void QVBoxPlotModelMapper_FirstRowChanged(void* ptr)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->firstRowChanged();
}

void QVBoxPlotModelMapper_ConnectLastBoxSetColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::lastBoxSetColumnChanged), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_LastBoxSetColumnChanged));
}

void QVBoxPlotModelMapper_DisconnectLastBoxSetColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::lastBoxSetColumnChanged), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_LastBoxSetColumnChanged));
}

void QVBoxPlotModelMapper_LastBoxSetColumnChanged(void* ptr)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->lastBoxSetColumnChanged();
}

void QVBoxPlotModelMapper_ConnectModelReplaced(void* ptr)
{
	QObject::connect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::modelReplaced), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_ModelReplaced));
}

void QVBoxPlotModelMapper_DisconnectModelReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::modelReplaced), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_ModelReplaced));
}

void QVBoxPlotModelMapper_ModelReplaced(void* ptr)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->modelReplaced();
}

void QVBoxPlotModelMapper_ConnectRowCountChanged(void* ptr)
{
	QObject::connect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::rowCountChanged), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_RowCountChanged));
}

void QVBoxPlotModelMapper_DisconnectRowCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::rowCountChanged), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_RowCountChanged));
}

void QVBoxPlotModelMapper_RowCountChanged(void* ptr)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->rowCountChanged();
}

void QVBoxPlotModelMapper_ConnectSeriesReplaced(void* ptr)
{
	QObject::connect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::seriesReplaced), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_SeriesReplaced));
}

void QVBoxPlotModelMapper_DisconnectSeriesReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QVBoxPlotModelMapper*>(ptr), static_cast<void (QVBoxPlotModelMapper::*)()>(&QVBoxPlotModelMapper::seriesReplaced), static_cast<MyQVBoxPlotModelMapper*>(ptr), static_cast<void (MyQVBoxPlotModelMapper::*)()>(&MyQVBoxPlotModelMapper::Signal_SeriesReplaced));
}

void QVBoxPlotModelMapper_SeriesReplaced(void* ptr)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->seriesReplaced();
}

void QVBoxPlotModelMapper_SetFirstBoxSetColumn(void* ptr, int firstBoxSetColumn)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->setFirstBoxSetColumn(firstBoxSetColumn);
}

void QVBoxPlotModelMapper_SetFirstRow(void* ptr, int firstRow)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->setFirstRow(firstRow);
}

void QVBoxPlotModelMapper_SetLastBoxSetColumn(void* ptr, int lastBoxSetColumn)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->setLastBoxSetColumn(lastBoxSetColumn);
}

void QVBoxPlotModelMapper_SetModel(void* ptr, void* model)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QVBoxPlotModelMapper_SetRowCount(void* ptr, int rowCount)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->setRowCount(rowCount);
}

void QVBoxPlotModelMapper_SetSeries(void* ptr, void* series)
{
	static_cast<QVBoxPlotModelMapper*>(ptr)->setSeries(static_cast<QBoxPlotSeries*>(series));
}

void* QVBoxPlotModelMapper_Model(void* ptr)
{
	return static_cast<QVBoxPlotModelMapper*>(ptr)->model();
}

void* QVBoxPlotModelMapper_Series(void* ptr)
{
	return static_cast<QVBoxPlotModelMapper*>(ptr)->series();
}

void* QVBoxPlotModelMapper_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVBoxPlotModelMapper*>(ptr)->QVBoxPlotModelMapper::metaObject());
}

int QVBoxPlotModelMapper_FirstBoxSetColumn(void* ptr)
{
	return static_cast<QVBoxPlotModelMapper*>(ptr)->firstBoxSetColumn();
}

int QVBoxPlotModelMapper_FirstRow(void* ptr)
{
	return static_cast<QVBoxPlotModelMapper*>(ptr)->firstRow();
}

int QVBoxPlotModelMapper_LastBoxSetColumn(void* ptr)
{
	return static_cast<QVBoxPlotModelMapper*>(ptr)->lastBoxSetColumn();
}

int QVBoxPlotModelMapper_RowCount(void* ptr)
{
	return static_cast<QVBoxPlotModelMapper*>(ptr)->rowCount();
}

void* QVBoxPlotModelMapper___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVBoxPlotModelMapper___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QVBoxPlotModelMapper___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QVBoxPlotModelMapper___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVBoxPlotModelMapper___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBoxPlotModelMapper___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVBoxPlotModelMapper___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVBoxPlotModelMapper___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBoxPlotModelMapper___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVBoxPlotModelMapper___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVBoxPlotModelMapper___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBoxPlotModelMapper___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QVBoxPlotModelMapper___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QVBoxPlotModelMapper___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBoxPlotModelMapper___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QVBoxPlotModelMapper_EventDefault(void* ptr, void* e)
{
		return static_cast<QVBoxPlotModelMapper*>(ptr)->QVBoxPlotModelMapper::event(static_cast<QEvent*>(e));
}

char QVBoxPlotModelMapper_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QVBoxPlotModelMapper*>(ptr)->QVBoxPlotModelMapper::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QVBoxPlotModelMapper_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QVBoxPlotModelMapper*>(ptr)->QVBoxPlotModelMapper::childEvent(static_cast<QChildEvent*>(event));
}

void QVBoxPlotModelMapper_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVBoxPlotModelMapper*>(ptr)->QVBoxPlotModelMapper::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVBoxPlotModelMapper_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QVBoxPlotModelMapper*>(ptr)->QVBoxPlotModelMapper::customEvent(static_cast<QEvent*>(event));
}

void QVBoxPlotModelMapper_DeleteLaterDefault(void* ptr)
{
		static_cast<QVBoxPlotModelMapper*>(ptr)->QVBoxPlotModelMapper::deleteLater();
}

void QVBoxPlotModelMapper_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QVBoxPlotModelMapper*>(ptr)->QVBoxPlotModelMapper::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVBoxPlotModelMapper_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QVBoxPlotModelMapper*>(ptr)->QVBoxPlotModelMapper::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQVCandlestickModelMapper: public QVCandlestickModelMapper
{
public:
	MyQVCandlestickModelMapper(QObject *parent = Q_NULLPTR) : QVCandlestickModelMapper(parent) {QVCandlestickModelMapper_QVCandlestickModelMapper_QRegisterMetaType();};
	void Signal_CloseRowChanged() { callbackQVCandlestickModelMapper_CloseRowChanged(this); };
	void Signal_FirstSetColumnChanged() { callbackQVCandlestickModelMapper_FirstSetColumnChanged(this); };
	void Signal_HighRowChanged() { callbackQVCandlestickModelMapper_HighRowChanged(this); };
	void Signal_LastSetColumnChanged() { callbackQVCandlestickModelMapper_LastSetColumnChanged(this); };
	void Signal_LowRowChanged() { callbackQVCandlestickModelMapper_LowRowChanged(this); };
	void Signal_OpenRowChanged() { callbackQVCandlestickModelMapper_OpenRowChanged(this); };
	void Signal_TimestampRowChanged() { callbackQVCandlestickModelMapper_TimestampRowChanged(this); };
	Qt::Orientation orientation() const { return static_cast<Qt::Orientation>(callbackQVCandlestickModelMapper_Orientation(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCandlestickModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ModelReplaced() { callbackQCandlestickModelMapper_ModelReplaced(this); };
	void Signal_SeriesReplaced() { callbackQCandlestickModelMapper_SeriesReplaced(this); };
	bool event(QEvent * e) { return callbackQCandlestickModelMapper_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCandlestickModelMapper_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQCandlestickModelMapper_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCandlestickModelMapper_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCandlestickModelMapper_CustomEvent(this, event); };
	void deleteLater() { callbackQCandlestickModelMapper_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQCandlestickModelMapper_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCandlestickModelMapper_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQCandlestickModelMapper_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQCandlestickModelMapper_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQVCandlestickModelMapper*)

int QVCandlestickModelMapper_QVCandlestickModelMapper_QRegisterMetaType(){qRegisterMetaType<QVCandlestickModelMapper*>(); return qRegisterMetaType<MyQVCandlestickModelMapper*>();}

void* QVCandlestickModelMapper_NewQVCandlestickModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVCandlestickModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQVCandlestickModelMapper(static_cast<QObject*>(parent));
	}
}

void QVCandlestickModelMapper_ConnectCloseRowChanged(void* ptr)
{
	QObject::connect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::closeRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_CloseRowChanged));
}

void QVCandlestickModelMapper_DisconnectCloseRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::closeRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_CloseRowChanged));
}

void QVCandlestickModelMapper_CloseRowChanged(void* ptr)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->closeRowChanged();
}

void QVCandlestickModelMapper_ConnectFirstSetColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::firstSetColumnChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_FirstSetColumnChanged));
}

void QVCandlestickModelMapper_DisconnectFirstSetColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::firstSetColumnChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_FirstSetColumnChanged));
}

void QVCandlestickModelMapper_FirstSetColumnChanged(void* ptr)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->firstSetColumnChanged();
}

void QVCandlestickModelMapper_ConnectHighRowChanged(void* ptr)
{
	QObject::connect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::highRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_HighRowChanged));
}

void QVCandlestickModelMapper_DisconnectHighRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::highRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_HighRowChanged));
}

void QVCandlestickModelMapper_HighRowChanged(void* ptr)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->highRowChanged();
}

void QVCandlestickModelMapper_ConnectLastSetColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::lastSetColumnChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_LastSetColumnChanged));
}

void QVCandlestickModelMapper_DisconnectLastSetColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::lastSetColumnChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_LastSetColumnChanged));
}

void QVCandlestickModelMapper_LastSetColumnChanged(void* ptr)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->lastSetColumnChanged();
}

void QVCandlestickModelMapper_ConnectLowRowChanged(void* ptr)
{
	QObject::connect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::lowRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_LowRowChanged));
}

void QVCandlestickModelMapper_DisconnectLowRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::lowRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_LowRowChanged));
}

void QVCandlestickModelMapper_LowRowChanged(void* ptr)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->lowRowChanged();
}

void QVCandlestickModelMapper_ConnectOpenRowChanged(void* ptr)
{
	QObject::connect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::openRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_OpenRowChanged));
}

void QVCandlestickModelMapper_DisconnectOpenRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::openRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_OpenRowChanged));
}

void QVCandlestickModelMapper_OpenRowChanged(void* ptr)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->openRowChanged();
}

void QVCandlestickModelMapper_SetCloseRow(void* ptr, int closeRow)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->setCloseRow(closeRow);
}

void QVCandlestickModelMapper_SetFirstSetColumn(void* ptr, int firstSetColumn)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->setFirstSetColumn(firstSetColumn);
}

void QVCandlestickModelMapper_SetHighRow(void* ptr, int highRow)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->setHighRow(highRow);
}

void QVCandlestickModelMapper_SetLastSetColumn(void* ptr, int lastSetColumn)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->setLastSetColumn(lastSetColumn);
}

void QVCandlestickModelMapper_SetLowRow(void* ptr, int lowRow)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->setLowRow(lowRow);
}

void QVCandlestickModelMapper_SetOpenRow(void* ptr, int openRow)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->setOpenRow(openRow);
}

void QVCandlestickModelMapper_SetTimestampRow(void* ptr, int timestampRow)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->setTimestampRow(timestampRow);
}

void QVCandlestickModelMapper_ConnectTimestampRowChanged(void* ptr)
{
	QObject::connect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::timestampRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_TimestampRowChanged));
}

void QVCandlestickModelMapper_DisconnectTimestampRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVCandlestickModelMapper*>(ptr), static_cast<void (QVCandlestickModelMapper::*)()>(&QVCandlestickModelMapper::timestampRowChanged), static_cast<MyQVCandlestickModelMapper*>(ptr), static_cast<void (MyQVCandlestickModelMapper::*)()>(&MyQVCandlestickModelMapper::Signal_TimestampRowChanged));
}

void QVCandlestickModelMapper_TimestampRowChanged(void* ptr)
{
	static_cast<QVCandlestickModelMapper*>(ptr)->timestampRowChanged();
}

long long QVCandlestickModelMapper_Orientation(void* ptr)
{
	return static_cast<QVCandlestickModelMapper*>(ptr)->orientation();
}

long long QVCandlestickModelMapper_OrientationDefault(void* ptr)
{
		return static_cast<QVCandlestickModelMapper*>(ptr)->QVCandlestickModelMapper::orientation();
}

int QVCandlestickModelMapper_CloseRow(void* ptr)
{
	return static_cast<QVCandlestickModelMapper*>(ptr)->closeRow();
}

int QVCandlestickModelMapper_FirstSetColumn(void* ptr)
{
	return static_cast<QVCandlestickModelMapper*>(ptr)->firstSetColumn();
}

int QVCandlestickModelMapper_HighRow(void* ptr)
{
	return static_cast<QVCandlestickModelMapper*>(ptr)->highRow();
}

int QVCandlestickModelMapper_LastSetColumn(void* ptr)
{
	return static_cast<QVCandlestickModelMapper*>(ptr)->lastSetColumn();
}

int QVCandlestickModelMapper_LowRow(void* ptr)
{
	return static_cast<QVCandlestickModelMapper*>(ptr)->lowRow();
}

int QVCandlestickModelMapper_OpenRow(void* ptr)
{
	return static_cast<QVCandlestickModelMapper*>(ptr)->openRow();
}

int QVCandlestickModelMapper_TimestampRow(void* ptr)
{
	return static_cast<QVCandlestickModelMapper*>(ptr)->timestampRow();
}

class MyQVPieModelMapper: public QVPieModelMapper
{
public:
	MyQVPieModelMapper(QObject *parent = Q_NULLPTR) : QVPieModelMapper(parent) {};
	void Signal_FirstRowChanged() { callbackQVPieModelMapper_FirstRowChanged(this); };
	void Signal_LabelsColumnChanged() { callbackQVPieModelMapper_LabelsColumnChanged(this); };
	void Signal_ModelReplaced() { callbackQVPieModelMapper_ModelReplaced(this); };
	void Signal_RowCountChanged() { callbackQVPieModelMapper_RowCountChanged(this); };
	void Signal_SeriesReplaced() { callbackQVPieModelMapper_SeriesReplaced(this); };
	void Signal_ValuesColumnChanged() { callbackQVPieModelMapper_ValuesColumnChanged(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVPieModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

struct QtCharts_PackedString QVPieModelMapper_QVPieModelMapper_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t94ee0a = QVPieModelMapper::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t94ee0a.prepend("WHITESPACE").constData()+10), t94ee0a.size()-10 }; });
}

struct QtCharts_PackedString QVPieModelMapper_QVPieModelMapper_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t8cf786 = QVPieModelMapper::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t8cf786.prepend("WHITESPACE").constData()+10), t8cf786.size()-10 }; });
}

void* QVPieModelMapper_NewQVPieModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVPieModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQVPieModelMapper(static_cast<QObject*>(parent));
	}
}

void QVPieModelMapper_ConnectFirstRowChanged(void* ptr)
{
	QObject::connect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::firstRowChanged), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_FirstRowChanged));
}

void QVPieModelMapper_DisconnectFirstRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::firstRowChanged), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_FirstRowChanged));
}

void QVPieModelMapper_FirstRowChanged(void* ptr)
{
	static_cast<QVPieModelMapper*>(ptr)->firstRowChanged();
}

void QVPieModelMapper_ConnectLabelsColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::labelsColumnChanged), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_LabelsColumnChanged));
}

void QVPieModelMapper_DisconnectLabelsColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::labelsColumnChanged), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_LabelsColumnChanged));
}

void QVPieModelMapper_LabelsColumnChanged(void* ptr)
{
	static_cast<QVPieModelMapper*>(ptr)->labelsColumnChanged();
}

void QVPieModelMapper_ConnectModelReplaced(void* ptr)
{
	QObject::connect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::modelReplaced), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_ModelReplaced));
}

void QVPieModelMapper_DisconnectModelReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::modelReplaced), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_ModelReplaced));
}

void QVPieModelMapper_ModelReplaced(void* ptr)
{
	static_cast<QVPieModelMapper*>(ptr)->modelReplaced();
}

void QVPieModelMapper_ConnectRowCountChanged(void* ptr)
{
	QObject::connect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::rowCountChanged), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_RowCountChanged));
}

void QVPieModelMapper_DisconnectRowCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::rowCountChanged), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_RowCountChanged));
}

void QVPieModelMapper_RowCountChanged(void* ptr)
{
	static_cast<QVPieModelMapper*>(ptr)->rowCountChanged();
}

void QVPieModelMapper_ConnectSeriesReplaced(void* ptr)
{
	QObject::connect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::seriesReplaced), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_SeriesReplaced));
}

void QVPieModelMapper_DisconnectSeriesReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::seriesReplaced), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_SeriesReplaced));
}

void QVPieModelMapper_SeriesReplaced(void* ptr)
{
	static_cast<QVPieModelMapper*>(ptr)->seriesReplaced();
}

void QVPieModelMapper_SetFirstRow(void* ptr, int firstRow)
{
	static_cast<QVPieModelMapper*>(ptr)->setFirstRow(firstRow);
}

void QVPieModelMapper_SetLabelsColumn(void* ptr, int labelsColumn)
{
	static_cast<QVPieModelMapper*>(ptr)->setLabelsColumn(labelsColumn);
}

void QVPieModelMapper_SetModel(void* ptr, void* model)
{
	static_cast<QVPieModelMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QVPieModelMapper_SetRowCount(void* ptr, int rowCount)
{
	static_cast<QVPieModelMapper*>(ptr)->setRowCount(rowCount);
}

void QVPieModelMapper_SetSeries(void* ptr, void* series)
{
	static_cast<QVPieModelMapper*>(ptr)->setSeries(static_cast<QPieSeries*>(series));
}

void QVPieModelMapper_SetValuesColumn(void* ptr, int valuesColumn)
{
	static_cast<QVPieModelMapper*>(ptr)->setValuesColumn(valuesColumn);
}

void QVPieModelMapper_ConnectValuesColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::valuesColumnChanged), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_ValuesColumnChanged));
}

void QVPieModelMapper_DisconnectValuesColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVPieModelMapper*>(ptr), static_cast<void (QVPieModelMapper::*)()>(&QVPieModelMapper::valuesColumnChanged), static_cast<MyQVPieModelMapper*>(ptr), static_cast<void (MyQVPieModelMapper::*)()>(&MyQVPieModelMapper::Signal_ValuesColumnChanged));
}

void QVPieModelMapper_ValuesColumnChanged(void* ptr)
{
	static_cast<QVPieModelMapper*>(ptr)->valuesColumnChanged();
}

void* QVPieModelMapper_Model(void* ptr)
{
	return static_cast<QVPieModelMapper*>(ptr)->model();
}

void* QVPieModelMapper_Series(void* ptr)
{
	return static_cast<QVPieModelMapper*>(ptr)->series();
}

void* QVPieModelMapper_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVPieModelMapper*>(ptr)->metaObject());
}

void* QVPieModelMapper_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVPieModelMapper*>(ptr)->QVPieModelMapper::metaObject());
}

int QVPieModelMapper_FirstRow(void* ptr)
{
	return static_cast<QVPieModelMapper*>(ptr)->firstRow();
}

int QVPieModelMapper_LabelsColumn(void* ptr)
{
	return static_cast<QVPieModelMapper*>(ptr)->labelsColumn();
}

int QVPieModelMapper_RowCount(void* ptr)
{
	return static_cast<QVPieModelMapper*>(ptr)->rowCount();
}

int QVPieModelMapper_ValuesColumn(void* ptr)
{
	return static_cast<QVPieModelMapper*>(ptr)->valuesColumn();
}

class MyQVXYModelMapper: public QVXYModelMapper
{
public:
	MyQVXYModelMapper(QObject *parent = Q_NULLPTR) : QVXYModelMapper(parent) {};
	void Signal_FirstRowChanged() { callbackQVXYModelMapper_FirstRowChanged(this); };
	void Signal_ModelReplaced() { callbackQVXYModelMapper_ModelReplaced(this); };
	void Signal_RowCountChanged() { callbackQVXYModelMapper_RowCountChanged(this); };
	void Signal_SeriesReplaced() { callbackQVXYModelMapper_SeriesReplaced(this); };
	void Signal_XColumnChanged() { callbackQVXYModelMapper_XColumnChanged(this); };
	void Signal_YColumnChanged() { callbackQVXYModelMapper_YColumnChanged(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVXYModelMapper_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

struct QtCharts_PackedString QVXYModelMapper_QVXYModelMapper_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t96ad4f = QVXYModelMapper::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(t96ad4f.prepend("WHITESPACE").constData()+10), t96ad4f.size()-10 }; });
}

struct QtCharts_PackedString QVXYModelMapper_QVXYModelMapper_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray td34d04 = QVXYModelMapper::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtCharts_PackedString { const_cast<char*>(td34d04.prepend("WHITESPACE").constData()+10), td34d04.size()-10 }; });
}

void* QVXYModelMapper_NewQVXYModelMapper(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQVXYModelMapper(static_cast<QWindow*>(parent));
	} else {
		return new MyQVXYModelMapper(static_cast<QObject*>(parent));
	}
}

void QVXYModelMapper_ConnectFirstRowChanged(void* ptr)
{
	QObject::connect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::firstRowChanged), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_FirstRowChanged));
}

void QVXYModelMapper_DisconnectFirstRowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::firstRowChanged), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_FirstRowChanged));
}

void QVXYModelMapper_FirstRowChanged(void* ptr)
{
	static_cast<QVXYModelMapper*>(ptr)->firstRowChanged();
}

void QVXYModelMapper_ConnectModelReplaced(void* ptr)
{
	QObject::connect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::modelReplaced), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_ModelReplaced));
}

void QVXYModelMapper_DisconnectModelReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::modelReplaced), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_ModelReplaced));
}

void QVXYModelMapper_ModelReplaced(void* ptr)
{
	static_cast<QVXYModelMapper*>(ptr)->modelReplaced();
}

void QVXYModelMapper_ConnectRowCountChanged(void* ptr)
{
	QObject::connect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::rowCountChanged), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_RowCountChanged));
}

void QVXYModelMapper_DisconnectRowCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::rowCountChanged), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_RowCountChanged));
}

void QVXYModelMapper_RowCountChanged(void* ptr)
{
	static_cast<QVXYModelMapper*>(ptr)->rowCountChanged();
}

void QVXYModelMapper_ConnectSeriesReplaced(void* ptr)
{
	QObject::connect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::seriesReplaced), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_SeriesReplaced));
}

void QVXYModelMapper_DisconnectSeriesReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::seriesReplaced), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_SeriesReplaced));
}

void QVXYModelMapper_SeriesReplaced(void* ptr)
{
	static_cast<QVXYModelMapper*>(ptr)->seriesReplaced();
}

void QVXYModelMapper_SetFirstRow(void* ptr, int firstRow)
{
	static_cast<QVXYModelMapper*>(ptr)->setFirstRow(firstRow);
}

void QVXYModelMapper_SetModel(void* ptr, void* model)
{
	static_cast<QVXYModelMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QVXYModelMapper_SetRowCount(void* ptr, int rowCount)
{
	static_cast<QVXYModelMapper*>(ptr)->setRowCount(rowCount);
}

void QVXYModelMapper_SetSeries(void* ptr, void* series)
{
	static_cast<QVXYModelMapper*>(ptr)->setSeries(static_cast<QXYSeries*>(series));
}

void QVXYModelMapper_SetXColumn(void* ptr, int xColumn)
{
	static_cast<QVXYModelMapper*>(ptr)->setXColumn(xColumn);
}

void QVXYModelMapper_SetYColumn(void* ptr, int yColumn)
{
	static_cast<QVXYModelMapper*>(ptr)->setYColumn(yColumn);
}

void QVXYModelMapper_ConnectXColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::xColumnChanged), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_XColumnChanged));
}

void QVXYModelMapper_DisconnectXColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::xColumnChanged), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_XColumnChanged));
}

void QVXYModelMapper_XColumnChanged(void* ptr)
{
	static_cast<QVXYModelMapper*>(ptr)->xColumnChanged();
}

void QVXYModelMapper_ConnectYColumnChanged(void* ptr)
{
	QObject::connect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::yColumnChanged), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_YColumnChanged));
}

void QVXYModelMapper_DisconnectYColumnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVXYModelMapper*>(ptr), static_cast<void (QVXYModelMapper::*)()>(&QVXYModelMapper::yColumnChanged), static_cast<MyQVXYModelMapper*>(ptr), static_cast<void (MyQVXYModelMapper::*)()>(&MyQVXYModelMapper::Signal_YColumnChanged));
}

void QVXYModelMapper_YColumnChanged(void* ptr)
{
	static_cast<QVXYModelMapper*>(ptr)->yColumnChanged();
}

void* QVXYModelMapper_Model(void* ptr)
{
	return static_cast<QVXYModelMapper*>(ptr)->model();
}

void* QVXYModelMapper_Series(void* ptr)
{
	return static_cast<QVXYModelMapper*>(ptr)->series();
}

void* QVXYModelMapper_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVXYModelMapper*>(ptr)->metaObject());
}

void* QVXYModelMapper_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QVXYModelMapper*>(ptr)->QVXYModelMapper::metaObject());
}

int QVXYModelMapper_FirstRow(void* ptr)
{
	return static_cast<QVXYModelMapper*>(ptr)->firstRow();
}

int QVXYModelMapper_RowCount(void* ptr)
{
	return static_cast<QVXYModelMapper*>(ptr)->rowCount();
}

int QVXYModelMapper_XColumn(void* ptr)
{
	return static_cast<QVXYModelMapper*>(ptr)->xColumn();
}

int QVXYModelMapper_YColumn(void* ptr)
{
	return static_cast<QVXYModelMapper*>(ptr)->yColumn();
}

class MyQValueAxis: public QValueAxis
{
public:
	MyQValueAxis(QObject *parent = Q_NULLPTR) : QValueAxis(parent) {QValueAxis_QValueAxis_QRegisterMetaType();};
	void applyNiceNumbers() { callbackQValueAxis_ApplyNiceNumbers(this); };
	void Signal_LabelFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQValueAxis_LabelFormatChanged(this, formatPacked); };
	void Signal_MaxChanged(qreal max) { callbackQValueAxis_MaxChanged(this, max); };
	void Signal_MinChanged(qreal min) { callbackQValueAxis_MinChanged(this, min); };
	void Signal_MinorTickCountChanged(int minorTickCount) { callbackQValueAxis_MinorTickCountChanged(this, minorTickCount); };
	void Signal_RangeChanged(qreal min, qreal max) { callbackQValueAxis_RangeChanged(this, min, max); };
	void Signal_TickAnchorChanged(qreal anchor) { callbackQValueAxis_TickAnchorChanged(this, anchor); };
	void Signal_TickCountChanged(int tickCount) { callbackQValueAxis_TickCountChanged(this, tickCount); };
	void Signal_TickIntervalChanged(qreal interval) { callbackQValueAxis_TickIntervalChanged(this, interval); };
	void Signal_TickTypeChanged(QValueAxis::TickType ty) { callbackQValueAxis_TickTypeChanged(this, ty); };
	 ~MyQValueAxis() { callbackQValueAxis_DestroyQValueAxis(this); };
	QAbstractAxis::AxisType type() const { return static_cast<QAbstractAxis::AxisType>(callbackQValueAxis_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractAxis_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ColorChanged(QColor color) { callbackQAbstractAxis_ColorChanged(this, new QColor(color)); };
	void Signal_GridLineColorChanged(const QColor & color) { callbackQAbstractAxis_GridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_GridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_GridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_GridVisibleChanged(bool visible) { callbackQAbstractAxis_GridVisibleChanged(this, visible); };
	void Signal_LabelsAngleChanged(int angle) { callbackQAbstractAxis_LabelsAngleChanged(this, angle); };
	void Signal_LabelsBrushChanged(const QBrush & brush) { callbackQAbstractAxis_LabelsBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_LabelsColorChanged(QColor color) { callbackQAbstractAxis_LabelsColorChanged(this, new QColor(color)); };
	void Signal_LabelsFontChanged(const QFont & font) { callbackQAbstractAxis_LabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_LabelsVisibleChanged(bool visible) { callbackQAbstractAxis_LabelsVisibleChanged(this, visible); };
	void Signal_LinePenChanged(const QPen & pen) { callbackQAbstractAxis_LinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_LineVisibleChanged(bool visible) { callbackQAbstractAxis_LineVisibleChanged(this, visible); };
	void Signal_MinorGridLineColorChanged(const QColor & color) { callbackQAbstractAxis_MinorGridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_MinorGridLinePenChanged(const QPen & pen) { callbackQAbstractAxis_MinorGridLinePenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_MinorGridVisibleChanged(bool visible) { callbackQAbstractAxis_MinorGridVisibleChanged(this, visible); };
	void Signal_ReverseChanged(bool reverse) { callbackQAbstractAxis_ReverseChanged(this, reverse); };
	void Signal_ShadesBorderColorChanged(QColor color) { callbackQAbstractAxis_ShadesBorderColorChanged(this, new QColor(color)); };
	void Signal_ShadesBrushChanged(const QBrush & brush) { callbackQAbstractAxis_ShadesBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_ShadesColorChanged(QColor color) { callbackQAbstractAxis_ShadesColorChanged(this, new QColor(color)); };
	void Signal_ShadesPenChanged(const QPen & pen) { callbackQAbstractAxis_ShadesPenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_ShadesVisibleChanged(bool visible) { callbackQAbstractAxis_ShadesVisibleChanged(this, visible); };
	void Signal_TitleBrushChanged(const QBrush & brush) { callbackQAbstractAxis_TitleBrushChanged(this, const_cast<QBrush*>(&brush)); };
	void Signal_TitleFontChanged(const QFont & font) { callbackQAbstractAxis_TitleFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_TitleTextChanged(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtCharts_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQAbstractAxis_TitleTextChanged(this, textPacked); };
	void Signal_TitleVisibleChanged(bool visible) { callbackQAbstractAxis_TitleVisibleChanged(this, visible); };
	void Signal_VisibleChanged(bool visible) { callbackQAbstractAxis_VisibleChanged(this, visible); };
	bool event(QEvent * e) { return callbackQAbstractAxis_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractAxis_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractAxis_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractAxis_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractAxis_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractAxis_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractAxis_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractAxis_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractAxis_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQValueAxis*)

int QValueAxis_QValueAxis_QRegisterMetaType(){qRegisterMetaType<QValueAxis*>(); return qRegisterMetaType<MyQValueAxis*>();}

void* QValueAxis_NewQValueAxis(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQValueAxis(static_cast<QWindow*>(parent));
	} else {
		return new MyQValueAxis(static_cast<QObject*>(parent));
	}
}

void QValueAxis_ApplyNiceNumbers(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QValueAxis*>(ptr), "applyNiceNumbers");
}

void QValueAxis_ApplyNiceNumbersDefault(void* ptr)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::applyNiceNumbers();
	} else {
		static_cast<QValueAxis*>(ptr)->QValueAxis::applyNiceNumbers();
	}
}

void QValueAxis_ConnectLabelFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(const QString &)>(&QValueAxis::labelFormatChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(const QString &)>(&MyQValueAxis::Signal_LabelFormatChanged));
}

void QValueAxis_DisconnectLabelFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(const QString &)>(&QValueAxis::labelFormatChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(const QString &)>(&MyQValueAxis::Signal_LabelFormatChanged));
}

void QValueAxis_LabelFormatChanged(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QValueAxis*>(ptr)->labelFormatChanged(QString::fromUtf8(format.data, format.len));
}

void QValueAxis_ConnectMaxChanged(void* ptr)
{
	QObject::connect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal)>(&QValueAxis::maxChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal)>(&MyQValueAxis::Signal_MaxChanged));
}

void QValueAxis_DisconnectMaxChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal)>(&QValueAxis::maxChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal)>(&MyQValueAxis::Signal_MaxChanged));
}

void QValueAxis_MaxChanged(void* ptr, double max)
{
	static_cast<QValueAxis*>(ptr)->maxChanged(max);
}

void QValueAxis_ConnectMinChanged(void* ptr)
{
	QObject::connect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal)>(&QValueAxis::minChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal)>(&MyQValueAxis::Signal_MinChanged));
}

void QValueAxis_DisconnectMinChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal)>(&QValueAxis::minChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal)>(&MyQValueAxis::Signal_MinChanged));
}

void QValueAxis_MinChanged(void* ptr, double min)
{
	static_cast<QValueAxis*>(ptr)->minChanged(min);
}

void QValueAxis_ConnectMinorTickCountChanged(void* ptr)
{
	QObject::connect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(int)>(&QValueAxis::minorTickCountChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(int)>(&MyQValueAxis::Signal_MinorTickCountChanged));
}

void QValueAxis_DisconnectMinorTickCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(int)>(&QValueAxis::minorTickCountChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(int)>(&MyQValueAxis::Signal_MinorTickCountChanged));
}

void QValueAxis_MinorTickCountChanged(void* ptr, int minorTickCount)
{
	static_cast<QValueAxis*>(ptr)->minorTickCountChanged(minorTickCount);
}

void QValueAxis_ConnectRangeChanged(void* ptr)
{
	QObject::connect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal, qreal)>(&QValueAxis::rangeChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal, qreal)>(&MyQValueAxis::Signal_RangeChanged));
}

void QValueAxis_DisconnectRangeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal, qreal)>(&QValueAxis::rangeChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal, qreal)>(&MyQValueAxis::Signal_RangeChanged));
}

void QValueAxis_RangeChanged(void* ptr, double min, double max)
{
	static_cast<QValueAxis*>(ptr)->rangeChanged(min, max);
}

void QValueAxis_SetLabelFormat(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QValueAxis*>(ptr)->setLabelFormat(QString::fromUtf8(format.data, format.len));
}

void QValueAxis_SetMax(void* ptr, double max)
{
	static_cast<QValueAxis*>(ptr)->setMax(max);
}

void QValueAxis_SetMin(void* ptr, double min)
{
	static_cast<QValueAxis*>(ptr)->setMin(min);
}

void QValueAxis_SetMinorTickCount(void* ptr, int count)
{
	static_cast<QValueAxis*>(ptr)->setMinorTickCount(count);
}

void QValueAxis_SetRange(void* ptr, double min, double max)
{
	static_cast<QValueAxis*>(ptr)->setRange(min, max);
}

void QValueAxis_SetTickAnchor(void* ptr, double anchor)
{
	static_cast<QValueAxis*>(ptr)->setTickAnchor(anchor);
}

void QValueAxis_SetTickCount(void* ptr, int count)
{
	static_cast<QValueAxis*>(ptr)->setTickCount(count);
}

void QValueAxis_SetTickInterval(void* ptr, double insterval)
{
	static_cast<QValueAxis*>(ptr)->setTickInterval(insterval);
}

void QValueAxis_SetTickType(void* ptr, long long ty)
{
	static_cast<QValueAxis*>(ptr)->setTickType(static_cast<QValueAxis::TickType>(ty));
}

void QValueAxis_ConnectTickAnchorChanged(void* ptr)
{
	QObject::connect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal)>(&QValueAxis::tickAnchorChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal)>(&MyQValueAxis::Signal_TickAnchorChanged));
}

void QValueAxis_DisconnectTickAnchorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal)>(&QValueAxis::tickAnchorChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal)>(&MyQValueAxis::Signal_TickAnchorChanged));
}

void QValueAxis_TickAnchorChanged(void* ptr, double anchor)
{
	static_cast<QValueAxis*>(ptr)->tickAnchorChanged(anchor);
}

void QValueAxis_ConnectTickCountChanged(void* ptr)
{
	QObject::connect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(int)>(&QValueAxis::tickCountChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(int)>(&MyQValueAxis::Signal_TickCountChanged));
}

void QValueAxis_DisconnectTickCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(int)>(&QValueAxis::tickCountChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(int)>(&MyQValueAxis::Signal_TickCountChanged));
}

void QValueAxis_TickCountChanged(void* ptr, int tickCount)
{
	static_cast<QValueAxis*>(ptr)->tickCountChanged(tickCount);
}

void QValueAxis_ConnectTickIntervalChanged(void* ptr)
{
	QObject::connect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal)>(&QValueAxis::tickIntervalChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal)>(&MyQValueAxis::Signal_TickIntervalChanged));
}

void QValueAxis_DisconnectTickIntervalChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(qreal)>(&QValueAxis::tickIntervalChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(qreal)>(&MyQValueAxis::Signal_TickIntervalChanged));
}

void QValueAxis_TickIntervalChanged(void* ptr, double interval)
{
	static_cast<QValueAxis*>(ptr)->tickIntervalChanged(interval);
}

void QValueAxis_ConnectTickTypeChanged(void* ptr)
{
	QObject::connect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(QValueAxis::TickType)>(&QValueAxis::tickTypeChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(QValueAxis::TickType)>(&MyQValueAxis::Signal_TickTypeChanged));
}

void QValueAxis_DisconnectTickTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValueAxis*>(ptr), static_cast<void (QValueAxis::*)(QValueAxis::TickType)>(&QValueAxis::tickTypeChanged), static_cast<MyQValueAxis*>(ptr), static_cast<void (MyQValueAxis::*)(QValueAxis::TickType)>(&MyQValueAxis::Signal_TickTypeChanged));
}

void QValueAxis_TickTypeChanged(void* ptr, long long ty)
{
	static_cast<QValueAxis*>(ptr)->tickTypeChanged(static_cast<QValueAxis::TickType>(ty));
}

void QValueAxis_DestroyQValueAxis(void* ptr)
{
	static_cast<QValueAxis*>(ptr)->~QValueAxis();
}

void QValueAxis_DestroyQValueAxisDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QValueAxis_Type(void* ptr)
{
	return static_cast<QValueAxis*>(ptr)->type();
}

long long QValueAxis_TypeDefault(void* ptr)
{
	if (dynamic_cast<QCategoryAxis*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCategoryAxis*>(ptr)->QCategoryAxis::type();
	} else {
		return static_cast<QValueAxis*>(ptr)->QValueAxis::type();
	}
}

struct QtCharts_PackedString QValueAxis_LabelFormat(void* ptr)
{
	return ({ QByteArray t607f3e = static_cast<QValueAxis*>(ptr)->labelFormat().toUtf8(); QtCharts_PackedString { const_cast<char*>(t607f3e.prepend("WHITESPACE").constData()+10), t607f3e.size()-10 }; });
}

long long QValueAxis_TickType(void* ptr)
{
	return static_cast<QValueAxis*>(ptr)->tickType();
}

int QValueAxis_MinorTickCount(void* ptr)
{
	return static_cast<QValueAxis*>(ptr)->minorTickCount();
}

int QValueAxis_TickCount(void* ptr)
{
	return static_cast<QValueAxis*>(ptr)->tickCount();
}

double QValueAxis_Max(void* ptr)
{
	return static_cast<QValueAxis*>(ptr)->max();
}

double QValueAxis_Min(void* ptr)
{
	return static_cast<QValueAxis*>(ptr)->min();
}

double QValueAxis_TickAnchor(void* ptr)
{
	return static_cast<QValueAxis*>(ptr)->tickAnchor();
}

double QValueAxis_TickInterval(void* ptr)
{
	return static_cast<QValueAxis*>(ptr)->tickInterval();
}

class MyQXYLegendMarker: public QXYLegendMarker
{
public:
	QLegendMarker::LegendMarkerType type() { return static_cast<QLegendMarker::LegendMarkerType>(callbackQXYLegendMarker_Type(this)); };
	QXYSeries * series() { return static_cast<QXYSeries*>(callbackQXYLegendMarker_Series(this)); };
	 ~MyQXYLegendMarker() { callbackQXYLegendMarker_DestroyQXYLegendMarker(this); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLegendMarker_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_BrushChanged() { callbackQLegendMarker_BrushChanged(this); };
	void Signal_Clicked() { callbackQLegendMarker_Clicked(this); };
	void Signal_FontChanged() { callbackQLegendMarker_FontChanged(this); };
	void Signal_Hovered(bool status) { callbackQLegendMarker_Hovered(this, status); };
	void Signal_LabelBrushChanged() { callbackQLegendMarker_LabelBrushChanged(this); };
	void Signal_LabelChanged() { callbackQLegendMarker_LabelChanged(this); };
	void Signal_PenChanged() { callbackQLegendMarker_PenChanged(this); };
	void Signal_ShapeChanged() { callbackQLegendMarker_ShapeChanged(this); };
	void Signal_VisibleChanged() { callbackQLegendMarker_VisibleChanged(this); };
	bool event(QEvent * e) { return callbackQLegendMarker_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLegendMarker_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQLegendMarker_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLegendMarker_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQLegendMarker_CustomEvent(this, event); };
	void deleteLater() { callbackQLegendMarker_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQLegendMarker_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLegendMarker_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQLegendMarker_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQLegendMarker_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQXYLegendMarker*)

int QXYLegendMarker_QXYLegendMarker_QRegisterMetaType(){qRegisterMetaType<QXYLegendMarker*>(); return qRegisterMetaType<MyQXYLegendMarker*>();}

long long QXYLegendMarker_Type(void* ptr)
{
	return static_cast<QXYLegendMarker*>(ptr)->type();
}

long long QXYLegendMarker_TypeDefault(void* ptr)
{
		return static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::type();
}

void* QXYLegendMarker_Series(void* ptr)
{
	return static_cast<QXYLegendMarker*>(ptr)->series();
}

void* QXYLegendMarker_SeriesDefault(void* ptr)
{
		return static_cast<QXYLegendMarker*>(ptr)->QXYLegendMarker::series();
}

void QXYLegendMarker_DestroyQXYLegendMarker(void* ptr)
{
	static_cast<QXYLegendMarker*>(ptr)->~QXYLegendMarker();
}

void QXYLegendMarker_DestroyQXYLegendMarkerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQXYSeries: public QXYSeries
{
public:
	void Signal_Clicked(const QPointF & point) { callbackQXYSeries_Clicked(this, const_cast<QPointF*>(&point)); };
	void Signal_ColorChanged(QColor color) { callbackQXYSeries_ColorChanged(this, new QColor(color)); };
	void Signal_DoubleClicked(const QPointF & point) { callbackQXYSeries_DoubleClicked(this, const_cast<QPointF*>(&point)); };
	void Signal_Hovered(const QPointF & point, bool state) { callbackQXYSeries_Hovered(this, const_cast<QPointF*>(&point), state); };
	void Signal_PenChanged(const QPen & pen) { callbackQXYSeries_PenChanged(this, const_cast<QPen*>(&pen)); };
	void Signal_PointAdded(int index) { callbackQXYSeries_PointAdded(this, index); };
	void Signal_PointLabelsClippingChanged(bool clipping) { callbackQXYSeries_PointLabelsClippingChanged(this, clipping); };
	void Signal_PointLabelsColorChanged(const QColor & color) { callbackQXYSeries_PointLabelsColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_PointLabelsFontChanged(const QFont & font) { callbackQXYSeries_PointLabelsFontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_PointLabelsFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtCharts_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQXYSeries_PointLabelsFormatChanged(this, formatPacked); };
	void Signal_PointLabelsVisibilityChanged(bool visible) { callbackQXYSeries_PointLabelsVisibilityChanged(this, visible); };
	void Signal_PointRemoved(int index) { callbackQXYSeries_PointRemoved(this, index); };
	void Signal_PointReplaced(int index) { callbackQXYSeries_PointReplaced(this, index); };
	void Signal_PointsRemoved(int index, int count) { callbackQXYSeries_PointsRemoved(this, index, count); };
	void Signal_PointsReplaced() { callbackQXYSeries_PointsReplaced(this); };
	void Signal_Pressed(const QPointF & point) { callbackQXYSeries_Pressed(this, const_cast<QPointF*>(&point)); };
	void Signal_Released(const QPointF & point) { callbackQXYSeries_Released(this, const_cast<QPointF*>(&point)); };
	void setBrush(const QBrush & brush) { callbackQXYSeries_SetBrush(this, const_cast<QBrush*>(&brush)); };
	void setColor(const QColor & color) { callbackQXYSeries_SetColor(this, const_cast<QColor*>(&color)); };
	void setPen(const QPen & pen) { callbackQXYSeries_SetPen(this, const_cast<QPen*>(&pen)); };
	 ~MyQXYSeries() { callbackQXYSeries_DestroyQXYSeries(this); };
	QColor color() const { return *static_cast<QColor*>(callbackQXYSeries_Color(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractSeries_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_NameChanged() { callbackQAbstractSeries_NameChanged(this); };
	void Signal_OpacityChanged() { callbackQAbstractSeries_OpacityChanged(this); };
	void Signal_UseOpenGLChanged() { callbackQAbstractSeries_UseOpenGLChanged(this); };
	void Signal_VisibleChanged() { callbackQAbstractSeries_VisibleChanged(this); };
	QAbstractSeries::SeriesType type() const { return static_cast<QAbstractSeries::SeriesType>(callbackQXYSeries_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQAbstractSeries_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractSeries_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQAbstractSeries_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractSeries_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractSeries_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractSeries_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractSeries_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCharts_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAbstractSeries_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractSeries_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQXYSeries*)

int QXYSeries_QXYSeries_QRegisterMetaType(){qRegisterMetaType<QXYSeries*>(); return qRegisterMetaType<MyQXYSeries*>();}

void QXYSeries_Append3(void* ptr, void* points)
{
	static_cast<QXYSeries*>(ptr)->append(*static_cast<QList<QPointF>*>(points));
}

void QXYSeries_Append2(void* ptr, void* point)
{
	static_cast<QXYSeries*>(ptr)->append(*static_cast<QPointF*>(point));
}

void QXYSeries_Append(void* ptr, double x, double y)
{
	static_cast<QXYSeries*>(ptr)->append(x, y);
}

void QXYSeries_Clear(void* ptr)
{
	static_cast<QXYSeries*>(ptr)->clear();
}

void QXYSeries_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &)>(&QXYSeries::clicked), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &)>(&MyQXYSeries::Signal_Clicked));
}

void QXYSeries_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &)>(&QXYSeries::clicked), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &)>(&MyQXYSeries::Signal_Clicked));
}

void QXYSeries_Clicked(void* ptr, void* point)
{
	static_cast<QXYSeries*>(ptr)->clicked(*static_cast<QPointF*>(point));
}

void QXYSeries_ConnectColorChanged(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(QColor)>(&QXYSeries::colorChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(QColor)>(&MyQXYSeries::Signal_ColorChanged));
}

void QXYSeries_DisconnectColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(QColor)>(&QXYSeries::colorChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(QColor)>(&MyQXYSeries::Signal_ColorChanged));
}

void QXYSeries_ColorChanged(void* ptr, void* color)
{
	static_cast<QXYSeries*>(ptr)->colorChanged(*static_cast<QColor*>(color));
}

void QXYSeries_ConnectDoubleClicked(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &)>(&QXYSeries::doubleClicked), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &)>(&MyQXYSeries::Signal_DoubleClicked));
}

void QXYSeries_DisconnectDoubleClicked(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &)>(&QXYSeries::doubleClicked), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &)>(&MyQXYSeries::Signal_DoubleClicked));
}

void QXYSeries_DoubleClicked(void* ptr, void* point)
{
	static_cast<QXYSeries*>(ptr)->doubleClicked(*static_cast<QPointF*>(point));
}

void QXYSeries_ConnectHovered(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &, bool)>(&QXYSeries::hovered), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &, bool)>(&MyQXYSeries::Signal_Hovered));
}

void QXYSeries_DisconnectHovered(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &, bool)>(&QXYSeries::hovered), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &, bool)>(&MyQXYSeries::Signal_Hovered));
}

void QXYSeries_Hovered(void* ptr, void* point, char state)
{
	static_cast<QXYSeries*>(ptr)->hovered(*static_cast<QPointF*>(point), state != 0);
}

void QXYSeries_Insert(void* ptr, int index, void* point)
{
	static_cast<QXYSeries*>(ptr)->insert(index, *static_cast<QPointF*>(point));
}

void QXYSeries_ConnectPenChanged(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPen &)>(&QXYSeries::penChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPen &)>(&MyQXYSeries::Signal_PenChanged));
}

void QXYSeries_DisconnectPenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPen &)>(&QXYSeries::penChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPen &)>(&MyQXYSeries::Signal_PenChanged));
}

void QXYSeries_PenChanged(void* ptr, void* pen)
{
	static_cast<QXYSeries*>(ptr)->penChanged(*static_cast<QPen*>(pen));
}

void QXYSeries_ConnectPointAdded(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(int)>(&QXYSeries::pointAdded), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(int)>(&MyQXYSeries::Signal_PointAdded));
}

void QXYSeries_DisconnectPointAdded(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(int)>(&QXYSeries::pointAdded), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(int)>(&MyQXYSeries::Signal_PointAdded));
}

void QXYSeries_PointAdded(void* ptr, int index)
{
	static_cast<QXYSeries*>(ptr)->pointAdded(index);
}

void QXYSeries_ConnectPointLabelsClippingChanged(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(bool)>(&QXYSeries::pointLabelsClippingChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(bool)>(&MyQXYSeries::Signal_PointLabelsClippingChanged));
}

void QXYSeries_DisconnectPointLabelsClippingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(bool)>(&QXYSeries::pointLabelsClippingChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(bool)>(&MyQXYSeries::Signal_PointLabelsClippingChanged));
}

void QXYSeries_PointLabelsClippingChanged(void* ptr, char clipping)
{
	static_cast<QXYSeries*>(ptr)->pointLabelsClippingChanged(clipping != 0);
}

void QXYSeries_ConnectPointLabelsColorChanged(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QColor &)>(&QXYSeries::pointLabelsColorChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QColor &)>(&MyQXYSeries::Signal_PointLabelsColorChanged));
}

void QXYSeries_DisconnectPointLabelsColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QColor &)>(&QXYSeries::pointLabelsColorChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QColor &)>(&MyQXYSeries::Signal_PointLabelsColorChanged));
}

void QXYSeries_PointLabelsColorChanged(void* ptr, void* color)
{
	static_cast<QXYSeries*>(ptr)->pointLabelsColorChanged(*static_cast<QColor*>(color));
}

void QXYSeries_ConnectPointLabelsFontChanged(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QFont &)>(&QXYSeries::pointLabelsFontChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QFont &)>(&MyQXYSeries::Signal_PointLabelsFontChanged));
}

void QXYSeries_DisconnectPointLabelsFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QFont &)>(&QXYSeries::pointLabelsFontChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QFont &)>(&MyQXYSeries::Signal_PointLabelsFontChanged));
}

void QXYSeries_PointLabelsFontChanged(void* ptr, void* font)
{
	static_cast<QXYSeries*>(ptr)->pointLabelsFontChanged(*static_cast<QFont*>(font));
}

void QXYSeries_ConnectPointLabelsFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QString &)>(&QXYSeries::pointLabelsFormatChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QString &)>(&MyQXYSeries::Signal_PointLabelsFormatChanged));
}

void QXYSeries_DisconnectPointLabelsFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QString &)>(&QXYSeries::pointLabelsFormatChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QString &)>(&MyQXYSeries::Signal_PointLabelsFormatChanged));
}

void QXYSeries_PointLabelsFormatChanged(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QXYSeries*>(ptr)->pointLabelsFormatChanged(QString::fromUtf8(format.data, format.len));
}

void QXYSeries_ConnectPointLabelsVisibilityChanged(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(bool)>(&QXYSeries::pointLabelsVisibilityChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(bool)>(&MyQXYSeries::Signal_PointLabelsVisibilityChanged));
}

void QXYSeries_DisconnectPointLabelsVisibilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(bool)>(&QXYSeries::pointLabelsVisibilityChanged), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(bool)>(&MyQXYSeries::Signal_PointLabelsVisibilityChanged));
}

void QXYSeries_PointLabelsVisibilityChanged(void* ptr, char visible)
{
	static_cast<QXYSeries*>(ptr)->pointLabelsVisibilityChanged(visible != 0);
}

void QXYSeries_ConnectPointRemoved(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(int)>(&QXYSeries::pointRemoved), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(int)>(&MyQXYSeries::Signal_PointRemoved));
}

void QXYSeries_DisconnectPointRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(int)>(&QXYSeries::pointRemoved), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(int)>(&MyQXYSeries::Signal_PointRemoved));
}

void QXYSeries_PointRemoved(void* ptr, int index)
{
	static_cast<QXYSeries*>(ptr)->pointRemoved(index);
}

void QXYSeries_ConnectPointReplaced(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(int)>(&QXYSeries::pointReplaced), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(int)>(&MyQXYSeries::Signal_PointReplaced));
}

void QXYSeries_DisconnectPointReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(int)>(&QXYSeries::pointReplaced), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(int)>(&MyQXYSeries::Signal_PointReplaced));
}

void QXYSeries_PointReplaced(void* ptr, int index)
{
	static_cast<QXYSeries*>(ptr)->pointReplaced(index);
}

void QXYSeries_ConnectPointsRemoved(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(int, int)>(&QXYSeries::pointsRemoved), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(int, int)>(&MyQXYSeries::Signal_PointsRemoved));
}

void QXYSeries_DisconnectPointsRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(int, int)>(&QXYSeries::pointsRemoved), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(int, int)>(&MyQXYSeries::Signal_PointsRemoved));
}

void QXYSeries_PointsRemoved(void* ptr, int index, int count)
{
	static_cast<QXYSeries*>(ptr)->pointsRemoved(index, count);
}

void QXYSeries_ConnectPointsReplaced(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)()>(&QXYSeries::pointsReplaced), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)()>(&MyQXYSeries::Signal_PointsReplaced));
}

void QXYSeries_DisconnectPointsReplaced(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)()>(&QXYSeries::pointsReplaced), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)()>(&MyQXYSeries::Signal_PointsReplaced));
}

void QXYSeries_PointsReplaced(void* ptr)
{
	static_cast<QXYSeries*>(ptr)->pointsReplaced();
}

void QXYSeries_ConnectPressed(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &)>(&QXYSeries::pressed), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &)>(&MyQXYSeries::Signal_Pressed));
}

void QXYSeries_DisconnectPressed(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &)>(&QXYSeries::pressed), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &)>(&MyQXYSeries::Signal_Pressed));
}

void QXYSeries_Pressed(void* ptr, void* point)
{
	static_cast<QXYSeries*>(ptr)->pressed(*static_cast<QPointF*>(point));
}

void QXYSeries_ConnectReleased(void* ptr)
{
	QObject::connect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &)>(&QXYSeries::released), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &)>(&MyQXYSeries::Signal_Released));
}

void QXYSeries_DisconnectReleased(void* ptr)
{
	QObject::disconnect(static_cast<QXYSeries*>(ptr), static_cast<void (QXYSeries::*)(const QPointF &)>(&QXYSeries::released), static_cast<MyQXYSeries*>(ptr), static_cast<void (MyQXYSeries::*)(const QPointF &)>(&MyQXYSeries::Signal_Released));
}

void QXYSeries_Released(void* ptr, void* point)
{
	static_cast<QXYSeries*>(ptr)->released(*static_cast<QPointF*>(point));
}

void QXYSeries_Remove2(void* ptr, void* point)
{
	static_cast<QXYSeries*>(ptr)->remove(*static_cast<QPointF*>(point));
}

void QXYSeries_Remove3(void* ptr, int index)
{
	static_cast<QXYSeries*>(ptr)->remove(index);
}

void QXYSeries_Remove(void* ptr, double x, double y)
{
	static_cast<QXYSeries*>(ptr)->remove(x, y);
}

void QXYSeries_RemovePoints(void* ptr, int index, int count)
{
	static_cast<QXYSeries*>(ptr)->removePoints(index, count);
}

void QXYSeries_Replace5(void* ptr, void* points)
{
	static_cast<QXYSeries*>(ptr)->replace(({ QList<QPointF>* tmpP = static_cast<QList<QPointF>*>(points); QList<QPointF> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QXYSeries_Replace6(void* ptr, void* points)
{
	static_cast<QXYSeries*>(ptr)->replace(({ QVector<QPointF>* tmpP = static_cast<QVector<QPointF>*>(points); QVector<QPointF> tmpV = *tmpP; tmpP->~QVector(); free(tmpP); tmpV; }));
}

void QXYSeries_Replace2(void* ptr, void* oldPoint, void* newPoint)
{
	static_cast<QXYSeries*>(ptr)->replace(*static_cast<QPointF*>(oldPoint), *static_cast<QPointF*>(newPoint));
}

void QXYSeries_Replace4(void* ptr, int index, void* newPoint)
{
	static_cast<QXYSeries*>(ptr)->replace(index, *static_cast<QPointF*>(newPoint));
}

void QXYSeries_Replace3(void* ptr, int index, double newX, double newY)
{
	static_cast<QXYSeries*>(ptr)->replace(index, newX, newY);
}

void QXYSeries_Replace(void* ptr, double oldX, double oldY, double newX, double newY)
{
	static_cast<QXYSeries*>(ptr)->replace(oldX, oldY, newX, newY);
}

void QXYSeries_SetBrush(void* ptr, void* brush)
{
	static_cast<QXYSeries*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QXYSeries_SetBrushDefault(void* ptr, void* brush)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QScatterSeries*>(ptr)->QScatterSeries::setBrush(*static_cast<QBrush*>(brush));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplineSeries*>(ptr)->QSplineSeries::setBrush(*static_cast<QBrush*>(brush));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineSeries*>(ptr)->QLineSeries::setBrush(*static_cast<QBrush*>(brush));
	} else {
		static_cast<QXYSeries*>(ptr)->QXYSeries::setBrush(*static_cast<QBrush*>(brush));
	}
}

void QXYSeries_SetColor(void* ptr, void* color)
{
	static_cast<QXYSeries*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QXYSeries_SetColorDefault(void* ptr, void* color)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QScatterSeries*>(ptr)->QScatterSeries::setColor(*static_cast<QColor*>(color));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplineSeries*>(ptr)->QSplineSeries::setColor(*static_cast<QColor*>(color));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineSeries*>(ptr)->QLineSeries::setColor(*static_cast<QColor*>(color));
	} else {
		static_cast<QXYSeries*>(ptr)->QXYSeries::setColor(*static_cast<QColor*>(color));
	}
}

void QXYSeries_SetPen(void* ptr, void* pen)
{
	static_cast<QXYSeries*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QXYSeries_SetPenDefault(void* ptr, void* pen)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QScatterSeries*>(ptr)->QScatterSeries::setPen(*static_cast<QPen*>(pen));
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QSplineSeries*>(ptr)->QSplineSeries::setPen(*static_cast<QPen*>(pen));
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		static_cast<QLineSeries*>(ptr)->QLineSeries::setPen(*static_cast<QPen*>(pen));
	} else {
		static_cast<QXYSeries*>(ptr)->QXYSeries::setPen(*static_cast<QPen*>(pen));
	}
}

void QXYSeries_SetPointLabelsClipping(void* ptr, char enabled)
{
	static_cast<QXYSeries*>(ptr)->setPointLabelsClipping(enabled != 0);
}

void QXYSeries_SetPointLabelsColor(void* ptr, void* color)
{
	static_cast<QXYSeries*>(ptr)->setPointLabelsColor(*static_cast<QColor*>(color));
}

void QXYSeries_SetPointLabelsFont(void* ptr, void* font)
{
	static_cast<QXYSeries*>(ptr)->setPointLabelsFont(*static_cast<QFont*>(font));
}

void QXYSeries_SetPointLabelsFormat(void* ptr, struct QtCharts_PackedString format)
{
	static_cast<QXYSeries*>(ptr)->setPointLabelsFormat(QString::fromUtf8(format.data, format.len));
}

void QXYSeries_SetPointLabelsVisible(void* ptr, char visible)
{
	static_cast<QXYSeries*>(ptr)->setPointLabelsVisible(visible != 0);
}

void QXYSeries_SetPointsVisible(void* ptr, char visible)
{
	static_cast<QXYSeries*>(ptr)->setPointsVisible(visible != 0);
}

void QXYSeries_DestroyQXYSeries(void* ptr)
{
	static_cast<QXYSeries*>(ptr)->~QXYSeries();
}

void QXYSeries_DestroyQXYSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QXYSeries_Brush(void* ptr)
{
	return new QBrush(static_cast<QXYSeries*>(ptr)->brush());
}

void* QXYSeries_Color(void* ptr)
{
	return new QColor(static_cast<QXYSeries*>(ptr)->color());
}

void* QXYSeries_ColorDefault(void* ptr)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		return new QColor(static_cast<QScatterSeries*>(ptr)->QScatterSeries::color());
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		return new QColor(static_cast<QSplineSeries*>(ptr)->QSplineSeries::color());
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		return new QColor(static_cast<QLineSeries*>(ptr)->QLineSeries::color());
	} else {
		return new QColor(static_cast<QXYSeries*>(ptr)->QXYSeries::color());
	}
}

void* QXYSeries_PointLabelsColor(void* ptr)
{
	return new QColor(static_cast<QXYSeries*>(ptr)->pointLabelsColor());
}

void* QXYSeries_PointLabelsFont(void* ptr)
{
	return new QFont(static_cast<QXYSeries*>(ptr)->pointLabelsFont());
}

struct QtCharts_PackedList QXYSeries_Points(void* ptr)
{
	return ({ QList<QPointF>* tmpValue = new QList<QPointF>(static_cast<QXYSeries*>(ptr)->points()); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

void* QXYSeries_Pen(void* ptr)
{
	return new QPen(static_cast<QXYSeries*>(ptr)->pen());
}

struct QtCharts_PackedString QXYSeries_PointLabelsFormat(void* ptr)
{
	return ({ QByteArray t00fe8f = static_cast<QXYSeries*>(ptr)->pointLabelsFormat().toUtf8(); QtCharts_PackedString { const_cast<char*>(t00fe8f.prepend("WHITESPACE").constData()+10), t00fe8f.size()-10 }; });
}

struct QtCharts_PackedList QXYSeries_PointsVector(void* ptr)
{
	return ({ QVector<QPointF>* tmpValue = new QVector<QPointF>(static_cast<QXYSeries*>(ptr)->pointsVector()); QtCharts_PackedList { tmpValue, tmpValue->size() }; });
}

char QXYSeries_PointLabelsClipping(void* ptr)
{
	return static_cast<QXYSeries*>(ptr)->pointLabelsClipping();
}

char QXYSeries_PointLabelsVisible(void* ptr)
{
	return static_cast<QXYSeries*>(ptr)->pointLabelsVisible();
}

char QXYSeries_PointsVisible(void* ptr)
{
	return static_cast<QXYSeries*>(ptr)->pointsVisible();
}

void* QXYSeries_At(void* ptr, int index)
{
	return const_cast<QPointF*>(&static_cast<QXYSeries*>(ptr)->at(index));
}

int QXYSeries_Count(void* ptr)
{
	return static_cast<QXYSeries*>(ptr)->count();
}

void* QXYSeries___append_points_atList3(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QList<QPointF>*>(ptr)->at(i); if (i == static_cast<QList<QPointF>*>(ptr)->size()-1) { static_cast<QList<QPointF>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QXYSeries___append_points_setList3(void* ptr, void* i)
{
	static_cast<QList<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QXYSeries___append_points_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPointF>();
}

void* QXYSeries___replace_points_atList5(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QList<QPointF>*>(ptr)->at(i); if (i == static_cast<QList<QPointF>*>(ptr)->size()-1) { static_cast<QList<QPointF>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QXYSeries___replace_points_setList5(void* ptr, void* i)
{
	static_cast<QList<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QXYSeries___replace_points_newList5(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPointF>();
}

void* QXYSeries___replace_points_atList6(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QXYSeries___replace_points_setList6(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QXYSeries___replace_points_newList6(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

void* QXYSeries___points_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QList<QPointF>*>(ptr)->at(i); if (i == static_cast<QList<QPointF>*>(ptr)->size()-1) { static_cast<QList<QPointF>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QXYSeries___points_setList(void* ptr, void* i)
{
	static_cast<QList<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QXYSeries___points_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPointF>();
}

void* QXYSeries___pointsVector_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QXYSeries___pointsVector_setList(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QXYSeries___pointsVector_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

long long QXYSeries_Type(void* ptr)
{
	return static_cast<QXYSeries*>(ptr)->type();
}

long long QXYSeries_TypeDefault(void* ptr)
{
	if (dynamic_cast<QScatterSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QScatterSeries*>(ptr)->QScatterSeries::type();
	} else if (dynamic_cast<QSplineSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSplineSeries*>(ptr)->QSplineSeries::type();
	} else if (dynamic_cast<QLineSeries*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLineSeries*>(ptr)->QLineSeries::type();
	} else {
	
	}
}

