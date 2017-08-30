// +build !minimal

#define protected public
#define private public

#include "datavisualization.h"
#include "_cgo_export.h"

#include <Q3DBars>
#include <Q3DCamera>
#include <Q3DInputHandler>
#include <Q3DLight>
#include <Q3DObject>
#include <Q3DScatter>
#include <Q3DScene>
#include <Q3DSurface>
#include <Q3DTheme>
#include <QAbstract3DAxis>
#include <QAbstract3DGraph>
#include <QAbstract3DInputHandler>
#include <QAbstract3DSeries>
#include <QAbstractDataProxy>
#include <QAbstractItemModel>
#include <QBar3DSeries>
#include <QBarDataItem>
#include <QBarDataProxy>
#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QCategory3DAxis>
#include <QColor>
#include <QCustom3DItem>
#include <QCustom3DLabel>
#include <QCustom3DVolume>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFont>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHeightMapSurfaceDataProxy>
#include <QImage>
#include <QItemModelBarDataProxy>
#include <QItemModelScatterDataProxy>
#include <QItemModelSurfaceDataProxy>
#include <QLayout>
#include <QLine>
#include <QLinearGradient>
#include <QList>
#include <QLocale>
#include <QLogValue3DAxisFormatter>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMouseEvent>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPoint>
#include <QQuaternion>
#include <QQuickItem>
#include <QRadioData>
#include <QRect>
#include <QRegExp>
#include <QScatter3DSeries>
#include <QScatterDataItem>
#include <QScatterDataProxy>
#include <QSignalSpy>
#include <QSize>
#include <QSizeF>
#include <QString>
#include <QStringList>
#include <QSurface>
#include <QSurface3DSeries>
#include <QSurfaceDataItem>
#include <QSurfaceDataProxy>
#include <QSurfaceFormat>
#include <QTouch3DInputHandler>
#include <QTouchEvent>
#include <QValue3DAxis>
#include <QValue3DAxisFormatter>
#include <QVector>
#include <QVector3D>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>

typedef QtDataVisualization::Q3DBars Q3DBars;
typedef QtDataVisualization::Q3DCamera Q3DCamera;
typedef QtDataVisualization::Q3DInputHandler Q3DInputHandler;
typedef QtDataVisualization::Q3DLight Q3DLight;
typedef QtDataVisualization::Q3DObject Q3DObject;
typedef QtDataVisualization::Q3DScatter Q3DScatter;
typedef QtDataVisualization::Q3DScene Q3DScene;
typedef QtDataVisualization::Q3DSurface Q3DSurface;
typedef QtDataVisualization::Q3DTheme Q3DTheme;
typedef QtDataVisualization::QAbstract3DAxis QAbstract3DAxis;
typedef QtDataVisualization::QAbstract3DGraph QAbstract3DGraph;
typedef QtDataVisualization::QAbstract3DInputHandler QAbstract3DInputHandler;
typedef QtDataVisualization::QAbstract3DSeries QAbstract3DSeries;
typedef QtDataVisualization::QAbstractDataProxy QAbstractDataProxy;
typedef QtDataVisualization::QBar3DSeries QBar3DSeries;
typedef QtDataVisualization::QBarDataItem QBarDataItem;
typedef QtDataVisualization::QBarDataProxy QBarDataProxy;
typedef QtDataVisualization::QCategory3DAxis QCategory3DAxis;
typedef QtDataVisualization::QCustom3DItem QCustom3DItem;
typedef QtDataVisualization::QCustom3DLabel QCustom3DLabel;
typedef QtDataVisualization::QCustom3DVolume QCustom3DVolume;
typedef QtDataVisualization::QHeightMapSurfaceDataProxy QHeightMapSurfaceDataProxy;
typedef QtDataVisualization::QItemModelBarDataProxy QItemModelBarDataProxy;
typedef QtDataVisualization::QItemModelScatterDataProxy QItemModelScatterDataProxy;
typedef QtDataVisualization::QItemModelSurfaceDataProxy QItemModelSurfaceDataProxy;
typedef QtDataVisualization::QLogValue3DAxisFormatter QLogValue3DAxisFormatter;
typedef QtDataVisualization::QScatter3DSeries QScatter3DSeries;
typedef QtDataVisualization::QScatterDataItem QScatterDataItem;
typedef QtDataVisualization::QScatterDataProxy QScatterDataProxy;
typedef QtDataVisualization::QSurface3DSeries QSurface3DSeries;
typedef QtDataVisualization::QSurfaceDataItem QSurfaceDataItem;
typedef QtDataVisualization::QSurfaceDataProxy QSurfaceDataProxy;
typedef QtDataVisualization::QTouch3DInputHandler QTouch3DInputHandler;
typedef QtDataVisualization::QValue3DAxis QValue3DAxis;
typedef QtDataVisualization::QValue3DAxisFormatter QValue3DAxisFormatter;

class MyQ3DBars: public Q3DBars
{
public:
	MyQ3DBars(const QSurfaceFormat *format = Q_NULLPTR, QWindow *parent = Q_NULLPTR) : Q3DBars(format, parent) {};
	void Signal_BarSpacingChanged(const QSizeF & spacing) { callbackQ3DBars_BarSpacingChanged(this, const_cast<QSizeF*>(&spacing)); };
	void Signal_BarSpacingRelativeChanged(bool relative) { callbackQ3DBars_BarSpacingRelativeChanged(this, relative); };
	void Signal_BarThicknessChanged(float thicknessRatio) { callbackQ3DBars_BarThicknessChanged(this, thicknessRatio); };
	void Signal_ColumnAxisChanged(QCategory3DAxis * axis) { callbackQ3DBars_ColumnAxisChanged(this, axis); };
	void Signal_FloorLevelChanged(float level) { callbackQ3DBars_FloorLevelChanged(this, level); };
	void Signal_MultiSeriesUniformChanged(bool uniform) { callbackQ3DBars_MultiSeriesUniformChanged(this, uniform); };
	void Signal_PrimarySeriesChanged(QBar3DSeries * series) { callbackQ3DBars_PrimarySeriesChanged(this, series); };
	void Signal_RowAxisChanged(QCategory3DAxis * axis) { callbackQ3DBars_RowAxisChanged(this, axis); };
	void Signal_SelectedSeriesChanged(QBar3DSeries * series) { callbackQ3DBars_SelectedSeriesChanged(this, series); };
	void Signal_ValueAxisChanged(QValue3DAxis * axis) { callbackQ3DBars_ValueAxisChanged(this, axis); };
	 ~MyQ3DBars() { callbackQ3DBars_DestroyQ3DBars(this); };
};

void* Q3DBars_NewQ3DBars(void* format, void* parent)
{
	if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DBars(static_cast<QSurfaceFormat*>(format), static_cast<QPaintDeviceWindow*>(parent));
	} else {
		return new MyQ3DBars(static_cast<QSurfaceFormat*>(format), static_cast<QWindow*>(parent));
	}
}

void Q3DBars_AddAxis(void* ptr, void* axis)
{
	static_cast<Q3DBars*>(ptr)->addAxis(static_cast<QAbstract3DAxis*>(axis));
}

void Q3DBars_AddSeries(void* ptr, void* series)
{
	static_cast<Q3DBars*>(ptr)->addSeries(static_cast<QBar3DSeries*>(series));
}

void Q3DBars_ConnectBarSpacingChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(const QSizeF &)>(&Q3DBars::barSpacingChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(const QSizeF &)>(&MyQ3DBars::Signal_BarSpacingChanged));
}

void Q3DBars_DisconnectBarSpacingChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(const QSizeF &)>(&Q3DBars::barSpacingChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(const QSizeF &)>(&MyQ3DBars::Signal_BarSpacingChanged));
}

void Q3DBars_BarSpacingChanged(void* ptr, void* spacing)
{
	static_cast<Q3DBars*>(ptr)->barSpacingChanged(*static_cast<QSizeF*>(spacing));
}

void Q3DBars_ConnectBarSpacingRelativeChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(bool)>(&Q3DBars::barSpacingRelativeChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(bool)>(&MyQ3DBars::Signal_BarSpacingRelativeChanged));
}

void Q3DBars_DisconnectBarSpacingRelativeChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(bool)>(&Q3DBars::barSpacingRelativeChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(bool)>(&MyQ3DBars::Signal_BarSpacingRelativeChanged));
}

void Q3DBars_BarSpacingRelativeChanged(void* ptr, char relative)
{
	static_cast<Q3DBars*>(ptr)->barSpacingRelativeChanged(relative != 0);
}

void Q3DBars_ConnectBarThicknessChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(float)>(&Q3DBars::barThicknessChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(float)>(&MyQ3DBars::Signal_BarThicknessChanged));
}

void Q3DBars_DisconnectBarThicknessChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(float)>(&Q3DBars::barThicknessChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(float)>(&MyQ3DBars::Signal_BarThicknessChanged));
}

void Q3DBars_BarThicknessChanged(void* ptr, float thicknessRatio)
{
	static_cast<Q3DBars*>(ptr)->barThicknessChanged(thicknessRatio);
}

void Q3DBars_ConnectColumnAxisChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QCategory3DAxis *)>(&Q3DBars::columnAxisChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QCategory3DAxis *)>(&MyQ3DBars::Signal_ColumnAxisChanged));
}

void Q3DBars_DisconnectColumnAxisChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QCategory3DAxis *)>(&Q3DBars::columnAxisChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QCategory3DAxis *)>(&MyQ3DBars::Signal_ColumnAxisChanged));
}

void Q3DBars_ColumnAxisChanged(void* ptr, void* axis)
{
	static_cast<Q3DBars*>(ptr)->columnAxisChanged(static_cast<QCategory3DAxis*>(axis));
}

void Q3DBars_ConnectFloorLevelChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(float)>(&Q3DBars::floorLevelChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(float)>(&MyQ3DBars::Signal_FloorLevelChanged));
}

void Q3DBars_DisconnectFloorLevelChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(float)>(&Q3DBars::floorLevelChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(float)>(&MyQ3DBars::Signal_FloorLevelChanged));
}

void Q3DBars_FloorLevelChanged(void* ptr, float level)
{
	static_cast<Q3DBars*>(ptr)->floorLevelChanged(level);
}

void Q3DBars_InsertSeries(void* ptr, int index, void* series)
{
	static_cast<Q3DBars*>(ptr)->insertSeries(index, static_cast<QBar3DSeries*>(series));
}

void Q3DBars_ConnectMultiSeriesUniformChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(bool)>(&Q3DBars::multiSeriesUniformChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(bool)>(&MyQ3DBars::Signal_MultiSeriesUniformChanged));
}

void Q3DBars_DisconnectMultiSeriesUniformChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(bool)>(&Q3DBars::multiSeriesUniformChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(bool)>(&MyQ3DBars::Signal_MultiSeriesUniformChanged));
}

void Q3DBars_MultiSeriesUniformChanged(void* ptr, char uniform)
{
	static_cast<Q3DBars*>(ptr)->multiSeriesUniformChanged(uniform != 0);
}

void Q3DBars_ConnectPrimarySeriesChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QBar3DSeries *)>(&Q3DBars::primarySeriesChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QBar3DSeries *)>(&MyQ3DBars::Signal_PrimarySeriesChanged));
}

void Q3DBars_DisconnectPrimarySeriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QBar3DSeries *)>(&Q3DBars::primarySeriesChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QBar3DSeries *)>(&MyQ3DBars::Signal_PrimarySeriesChanged));
}

void Q3DBars_PrimarySeriesChanged(void* ptr, void* series)
{
	static_cast<Q3DBars*>(ptr)->primarySeriesChanged(static_cast<QBar3DSeries*>(series));
}

void Q3DBars_ReleaseAxis(void* ptr, void* axis)
{
	static_cast<Q3DBars*>(ptr)->releaseAxis(static_cast<QAbstract3DAxis*>(axis));
}

void Q3DBars_RemoveSeries(void* ptr, void* series)
{
	static_cast<Q3DBars*>(ptr)->removeSeries(static_cast<QBar3DSeries*>(series));
}

void Q3DBars_ConnectRowAxisChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QCategory3DAxis *)>(&Q3DBars::rowAxisChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QCategory3DAxis *)>(&MyQ3DBars::Signal_RowAxisChanged));
}

void Q3DBars_DisconnectRowAxisChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QCategory3DAxis *)>(&Q3DBars::rowAxisChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QCategory3DAxis *)>(&MyQ3DBars::Signal_RowAxisChanged));
}

void Q3DBars_RowAxisChanged(void* ptr, void* axis)
{
	static_cast<Q3DBars*>(ptr)->rowAxisChanged(static_cast<QCategory3DAxis*>(axis));
}

void Q3DBars_ConnectSelectedSeriesChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QBar3DSeries *)>(&Q3DBars::selectedSeriesChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QBar3DSeries *)>(&MyQ3DBars::Signal_SelectedSeriesChanged));
}

void Q3DBars_DisconnectSelectedSeriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QBar3DSeries *)>(&Q3DBars::selectedSeriesChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QBar3DSeries *)>(&MyQ3DBars::Signal_SelectedSeriesChanged));
}

void Q3DBars_SelectedSeriesChanged(void* ptr, void* series)
{
	static_cast<Q3DBars*>(ptr)->selectedSeriesChanged(static_cast<QBar3DSeries*>(series));
}

void Q3DBars_SetBarSpacing(void* ptr, void* spacing)
{
	static_cast<Q3DBars*>(ptr)->setBarSpacing(*static_cast<QSizeF*>(spacing));
}

void Q3DBars_SetBarSpacingRelative(void* ptr, char relative)
{
	static_cast<Q3DBars*>(ptr)->setBarSpacingRelative(relative != 0);
}

void Q3DBars_SetBarThickness(void* ptr, float thicknessRatio)
{
	static_cast<Q3DBars*>(ptr)->setBarThickness(thicknessRatio);
}

void Q3DBars_SetColumnAxis(void* ptr, void* axis)
{
	static_cast<Q3DBars*>(ptr)->setColumnAxis(static_cast<QCategory3DAxis*>(axis));
}

void Q3DBars_SetFloorLevel(void* ptr, float level)
{
	static_cast<Q3DBars*>(ptr)->setFloorLevel(level);
}

void Q3DBars_SetMultiSeriesUniform(void* ptr, char uniform)
{
	static_cast<Q3DBars*>(ptr)->setMultiSeriesUniform(uniform != 0);
}

void Q3DBars_SetPrimarySeries(void* ptr, void* series)
{
	static_cast<Q3DBars*>(ptr)->setPrimarySeries(static_cast<QBar3DSeries*>(series));
}

void Q3DBars_SetRowAxis(void* ptr, void* axis)
{
	static_cast<Q3DBars*>(ptr)->setRowAxis(static_cast<QCategory3DAxis*>(axis));
}

void Q3DBars_SetValueAxis(void* ptr, void* axis)
{
	static_cast<Q3DBars*>(ptr)->setValueAxis(static_cast<QValue3DAxis*>(axis));
}

void Q3DBars_ConnectValueAxisChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QValue3DAxis *)>(&Q3DBars::valueAxisChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QValue3DAxis *)>(&MyQ3DBars::Signal_ValueAxisChanged));
}

void Q3DBars_DisconnectValueAxisChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DBars*>(ptr), static_cast<void (Q3DBars::*)(QValue3DAxis *)>(&Q3DBars::valueAxisChanged), static_cast<MyQ3DBars*>(ptr), static_cast<void (MyQ3DBars::*)(QValue3DAxis *)>(&MyQ3DBars::Signal_ValueAxisChanged));
}

void Q3DBars_ValueAxisChanged(void* ptr, void* axis)
{
	static_cast<Q3DBars*>(ptr)->valueAxisChanged(static_cast<QValue3DAxis*>(axis));
}

void Q3DBars_DestroyQ3DBars(void* ptr)
{
	static_cast<Q3DBars*>(ptr)->~Q3DBars();
}

void Q3DBars_DestroyQ3DBarsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* Q3DBars_PrimarySeries(void* ptr)
{
	return static_cast<Q3DBars*>(ptr)->primarySeries();
}

void* Q3DBars_SelectedSeries(void* ptr)
{
	return static_cast<Q3DBars*>(ptr)->selectedSeries();
}

void* Q3DBars_ColumnAxis(void* ptr)
{
	return static_cast<Q3DBars*>(ptr)->columnAxis();
}

void* Q3DBars_RowAxis(void* ptr)
{
	return static_cast<Q3DBars*>(ptr)->rowAxis();
}

struct QtDataVisualization_PackedList Q3DBars_Axes(void* ptr)
{
	return ({ QList<QAbstract3DAxis *>* tmpValue = new QList<QAbstract3DAxis *>(static_cast<Q3DBars*>(ptr)->axes()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDataVisualization_PackedList Q3DBars_SeriesList(void* ptr)
{
	return ({ QList<QBar3DSeries *>* tmpValue = new QList<QBar3DSeries *>(static_cast<Q3DBars*>(ptr)->seriesList()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

void* Q3DBars_BarSpacing(void* ptr)
{
	return ({ QSizeF tmpValue = static_cast<Q3DBars*>(ptr)->barSpacing(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* Q3DBars_ValueAxis(void* ptr)
{
	return static_cast<Q3DBars*>(ptr)->valueAxis();
}

char Q3DBars_IsBarSpacingRelative(void* ptr)
{
	return static_cast<Q3DBars*>(ptr)->isBarSpacingRelative();
}

char Q3DBars_IsMultiSeriesUniform(void* ptr)
{
	return static_cast<Q3DBars*>(ptr)->isMultiSeriesUniform();
}

float Q3DBars_BarThickness(void* ptr)
{
	return static_cast<Q3DBars*>(ptr)->barThickness();
}

float Q3DBars_FloorLevel(void* ptr)
{
	return static_cast<Q3DBars*>(ptr)->floorLevel();
}

void* Q3DBars___axes_atList(void* ptr, int i)
{
	return const_cast<QAbstract3DAxis*>(static_cast<QList<QAbstract3DAxis *>*>(ptr)->at(i));
}

void Q3DBars___axes_setList(void* ptr, void* i)
{
	static_cast<QList<QAbstract3DAxis *>*>(ptr)->append(static_cast<QAbstract3DAxis*>(i));
}

void* Q3DBars___axes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAbstract3DAxis *>;
}

void* Q3DBars___seriesList_atList(void* ptr, int i)
{
	return const_cast<QBar3DSeries*>(static_cast<QList<QBar3DSeries *>*>(ptr)->at(i));
}

void Q3DBars___seriesList_setList(void* ptr, void* i)
{
	static_cast<QList<QBar3DSeries *>*>(ptr)->append(static_cast<QBar3DSeries*>(i));
}

void* Q3DBars___seriesList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QBar3DSeries *>;
}

class MyQ3DCamera: public Q3DCamera
{
public:
	MyQ3DCamera(QObject *parent = Q_NULLPTR) : Q3DCamera(parent) {};
	void Signal_CameraPresetChanged(Q3DCamera::CameraPreset preset) { callbackQ3DCamera_CameraPresetChanged(this, preset); };
	void copyValuesFrom(const Q3DObject & source) { callbackQ3DCamera_CopyValuesFrom(this, const_cast<Q3DObject*>(&source)); };
	void Signal_MaxZoomLevelChanged(float zoomLevel) { callbackQ3DCamera_MaxZoomLevelChanged(this, zoomLevel); };
	void Signal_MinZoomLevelChanged(float zoomLevel) { callbackQ3DCamera_MinZoomLevelChanged(this, zoomLevel); };
	void Signal_TargetChanged(const QVector3D & target) { callbackQ3DCamera_TargetChanged(this, const_cast<QVector3D*>(&target)); };
	void Signal_WrapXRotationChanged(bool isEnabled) { callbackQ3DCamera_WrapXRotationChanged(this, isEnabled); };
	void Signal_WrapYRotationChanged(bool isEnabled) { callbackQ3DCamera_WrapYRotationChanged(this, isEnabled); };
	void Signal_XRotationChanged(float rotation) { callbackQ3DCamera_XRotationChanged(this, rotation); };
	void Signal_YRotationChanged(float rotation) { callbackQ3DCamera_YRotationChanged(this, rotation); };
	void Signal_ZoomLevelChanged(float zoomLevel) { callbackQ3DCamera_ZoomLevelChanged(this, zoomLevel); };
	 ~MyQ3DCamera() { callbackQ3DCamera_DestroyQ3DCamera(this); };
};

void Q3DCamera_SetCameraPosition(void* ptr, float horizontal, float vertical, float zoom)
{
	static_cast<Q3DCamera*>(ptr)->setCameraPosition(horizontal, vertical, zoom);
}

void Q3DCamera_SetTarget(void* ptr, void* target)
{
	static_cast<Q3DCamera*>(ptr)->setTarget(*static_cast<QVector3D*>(target));
}

void* Q3DCamera_NewQ3DCamera(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DCamera(static_cast<QWindow*>(parent));
	} else {
		return new MyQ3DCamera(static_cast<QObject*>(parent));
	}
}

void Q3DCamera_ConnectCameraPresetChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(Q3DCamera::CameraPreset)>(&Q3DCamera::cameraPresetChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(Q3DCamera::CameraPreset)>(&MyQ3DCamera::Signal_CameraPresetChanged));
}

void Q3DCamera_DisconnectCameraPresetChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(Q3DCamera::CameraPreset)>(&Q3DCamera::cameraPresetChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(Q3DCamera::CameraPreset)>(&MyQ3DCamera::Signal_CameraPresetChanged));
}

void Q3DCamera_CameraPresetChanged(void* ptr, long long preset)
{
	static_cast<Q3DCamera*>(ptr)->cameraPresetChanged(static_cast<Q3DCamera::CameraPreset>(preset));
}

void Q3DCamera_CopyValuesFrom(void* ptr, void* source)
{
	static_cast<Q3DCamera*>(ptr)->copyValuesFrom(*static_cast<Q3DObject*>(source));
}

void Q3DCamera_CopyValuesFromDefault(void* ptr, void* source)
{
		static_cast<Q3DCamera*>(ptr)->Q3DCamera::copyValuesFrom(*static_cast<Q3DObject*>(source));
}

void Q3DCamera_ConnectMaxZoomLevelChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::maxZoomLevelChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_MaxZoomLevelChanged));
}

void Q3DCamera_DisconnectMaxZoomLevelChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::maxZoomLevelChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_MaxZoomLevelChanged));
}

void Q3DCamera_MaxZoomLevelChanged(void* ptr, float zoomLevel)
{
	static_cast<Q3DCamera*>(ptr)->maxZoomLevelChanged(zoomLevel);
}

void Q3DCamera_ConnectMinZoomLevelChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::minZoomLevelChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_MinZoomLevelChanged));
}

void Q3DCamera_DisconnectMinZoomLevelChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::minZoomLevelChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_MinZoomLevelChanged));
}

void Q3DCamera_MinZoomLevelChanged(void* ptr, float zoomLevel)
{
	static_cast<Q3DCamera*>(ptr)->minZoomLevelChanged(zoomLevel);
}

void Q3DCamera_SetCameraPreset(void* ptr, long long preset)
{
	static_cast<Q3DCamera*>(ptr)->setCameraPreset(static_cast<Q3DCamera::CameraPreset>(preset));
}

void Q3DCamera_SetMaxZoomLevel(void* ptr, float zoomLevel)
{
	static_cast<Q3DCamera*>(ptr)->setMaxZoomLevel(zoomLevel);
}

void Q3DCamera_SetMinZoomLevel(void* ptr, float zoomLevel)
{
	static_cast<Q3DCamera*>(ptr)->setMinZoomLevel(zoomLevel);
}

void Q3DCamera_SetWrapXRotation(void* ptr, char isEnabled)
{
	static_cast<Q3DCamera*>(ptr)->setWrapXRotation(isEnabled != 0);
}

void Q3DCamera_SetWrapYRotation(void* ptr, char isEnabled)
{
	static_cast<Q3DCamera*>(ptr)->setWrapYRotation(isEnabled != 0);
}

void Q3DCamera_SetXRotation(void* ptr, float rotation)
{
	static_cast<Q3DCamera*>(ptr)->setXRotation(rotation);
}

void Q3DCamera_SetYRotation(void* ptr, float rotation)
{
	static_cast<Q3DCamera*>(ptr)->setYRotation(rotation);
}

void Q3DCamera_SetZoomLevel(void* ptr, float zoomLevel)
{
	static_cast<Q3DCamera*>(ptr)->setZoomLevel(zoomLevel);
}

void Q3DCamera_ConnectTargetChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(const QVector3D &)>(&Q3DCamera::targetChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(const QVector3D &)>(&MyQ3DCamera::Signal_TargetChanged));
}

void Q3DCamera_DisconnectTargetChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(const QVector3D &)>(&Q3DCamera::targetChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(const QVector3D &)>(&MyQ3DCamera::Signal_TargetChanged));
}

void Q3DCamera_TargetChanged(void* ptr, void* target)
{
	static_cast<Q3DCamera*>(ptr)->targetChanged(*static_cast<QVector3D*>(target));
}

void Q3DCamera_ConnectWrapXRotationChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(bool)>(&Q3DCamera::wrapXRotationChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(bool)>(&MyQ3DCamera::Signal_WrapXRotationChanged));
}

void Q3DCamera_DisconnectWrapXRotationChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(bool)>(&Q3DCamera::wrapXRotationChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(bool)>(&MyQ3DCamera::Signal_WrapXRotationChanged));
}

void Q3DCamera_WrapXRotationChanged(void* ptr, char isEnabled)
{
	static_cast<Q3DCamera*>(ptr)->wrapXRotationChanged(isEnabled != 0);
}

void Q3DCamera_ConnectWrapYRotationChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(bool)>(&Q3DCamera::wrapYRotationChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(bool)>(&MyQ3DCamera::Signal_WrapYRotationChanged));
}

void Q3DCamera_DisconnectWrapYRotationChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(bool)>(&Q3DCamera::wrapYRotationChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(bool)>(&MyQ3DCamera::Signal_WrapYRotationChanged));
}

void Q3DCamera_WrapYRotationChanged(void* ptr, char isEnabled)
{
	static_cast<Q3DCamera*>(ptr)->wrapYRotationChanged(isEnabled != 0);
}

void Q3DCamera_ConnectXRotationChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::xRotationChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_XRotationChanged));
}

void Q3DCamera_DisconnectXRotationChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::xRotationChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_XRotationChanged));
}

void Q3DCamera_XRotationChanged(void* ptr, float rotation)
{
	static_cast<Q3DCamera*>(ptr)->xRotationChanged(rotation);
}

void Q3DCamera_ConnectYRotationChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::yRotationChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_YRotationChanged));
}

void Q3DCamera_DisconnectYRotationChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::yRotationChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_YRotationChanged));
}

void Q3DCamera_YRotationChanged(void* ptr, float rotation)
{
	static_cast<Q3DCamera*>(ptr)->yRotationChanged(rotation);
}

void Q3DCamera_ConnectZoomLevelChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::zoomLevelChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_ZoomLevelChanged));
}

void Q3DCamera_DisconnectZoomLevelChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DCamera*>(ptr), static_cast<void (Q3DCamera::*)(float)>(&Q3DCamera::zoomLevelChanged), static_cast<MyQ3DCamera*>(ptr), static_cast<void (MyQ3DCamera::*)(float)>(&MyQ3DCamera::Signal_ZoomLevelChanged));
}

void Q3DCamera_ZoomLevelChanged(void* ptr, float zoomLevel)
{
	static_cast<Q3DCamera*>(ptr)->zoomLevelChanged(zoomLevel);
}

void Q3DCamera_DestroyQ3DCamera(void* ptr)
{
	static_cast<Q3DCamera*>(ptr)->~Q3DCamera();
}

void Q3DCamera_DestroyQ3DCameraDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long Q3DCamera_CameraPreset(void* ptr)
{
	return static_cast<Q3DCamera*>(ptr)->cameraPreset();
}

void* Q3DCamera_Target(void* ptr)
{
	return new QVector3D(static_cast<Q3DCamera*>(ptr)->target());
}

char Q3DCamera_WrapXRotation(void* ptr)
{
	return static_cast<Q3DCamera*>(ptr)->wrapXRotation();
}

char Q3DCamera_WrapYRotation(void* ptr)
{
	return static_cast<Q3DCamera*>(ptr)->wrapYRotation();
}

float Q3DCamera_MaxZoomLevel(void* ptr)
{
	return static_cast<Q3DCamera*>(ptr)->maxZoomLevel();
}

float Q3DCamera_MinZoomLevel(void* ptr)
{
	return static_cast<Q3DCamera*>(ptr)->minZoomLevel();
}

float Q3DCamera_XRotation(void* ptr)
{
	return static_cast<Q3DCamera*>(ptr)->xRotation();
}

float Q3DCamera_YRotation(void* ptr)
{
	return static_cast<Q3DCamera*>(ptr)->yRotation();
}

float Q3DCamera_ZoomLevel(void* ptr)
{
	return static_cast<Q3DCamera*>(ptr)->zoomLevel();
}

class MyQ3DInputHandler: public Q3DInputHandler
{
public:
	MyQ3DInputHandler(QObject *parent = Q_NULLPTR) : Q3DInputHandler(parent) {};
	void mouseMoveEvent(QMouseEvent * event, const QPoint & mousePos) { callbackQ3DInputHandler_MouseMoveEvent(this, event, const_cast<QPoint*>(&mousePos)); };
	void mousePressEvent(QMouseEvent * event, const QPoint & mousePos) { callbackQ3DInputHandler_MousePressEvent(this, event, const_cast<QPoint*>(&mousePos)); };
	void mouseReleaseEvent(QMouseEvent * event, const QPoint & mousePos) { callbackQ3DInputHandler_MouseReleaseEvent(this, event, const_cast<QPoint*>(&mousePos)); };
	void Signal_RotationEnabledChanged(bool enable) { callbackQ3DInputHandler_RotationEnabledChanged(this, enable); };
	void Signal_SelectionEnabledChanged(bool enable) { callbackQ3DInputHandler_SelectionEnabledChanged(this, enable); };
	void wheelEvent(QWheelEvent * event) { callbackQ3DInputHandler_WheelEvent(this, event); };
	void Signal_ZoomAtTargetEnabledChanged(bool enable) { callbackQ3DInputHandler_ZoomAtTargetEnabledChanged(this, enable); };
	void Signal_ZoomEnabledChanged(bool enable) { callbackQ3DInputHandler_ZoomEnabledChanged(this, enable); };
	 ~MyQ3DInputHandler() { callbackQ3DInputHandler_DestroyQ3DInputHandler(this); };
};

void* Q3DInputHandler_NewQ3DInputHandler(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DInputHandler(static_cast<QWindow*>(parent));
	} else {
		return new MyQ3DInputHandler(static_cast<QObject*>(parent));
	}
}

void Q3DInputHandler_MouseMoveEvent(void* ptr, void* event, void* mousePos)
{
	static_cast<Q3DInputHandler*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void Q3DInputHandler_MouseMoveEventDefault(void* ptr, void* event, void* mousePos)
{
		static_cast<Q3DInputHandler*>(ptr)->Q3DInputHandler::mouseMoveEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void Q3DInputHandler_MousePressEvent(void* ptr, void* event, void* mousePos)
{
	static_cast<Q3DInputHandler*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void Q3DInputHandler_MousePressEventDefault(void* ptr, void* event, void* mousePos)
{
		static_cast<Q3DInputHandler*>(ptr)->Q3DInputHandler::mousePressEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void Q3DInputHandler_MouseReleaseEvent(void* ptr, void* event, void* mousePos)
{
	static_cast<Q3DInputHandler*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void Q3DInputHandler_MouseReleaseEventDefault(void* ptr, void* event, void* mousePos)
{
		static_cast<Q3DInputHandler*>(ptr)->Q3DInputHandler::mouseReleaseEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void Q3DInputHandler_ConnectRotationEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DInputHandler*>(ptr), static_cast<void (Q3DInputHandler::*)(bool)>(&Q3DInputHandler::rotationEnabledChanged), static_cast<MyQ3DInputHandler*>(ptr), static_cast<void (MyQ3DInputHandler::*)(bool)>(&MyQ3DInputHandler::Signal_RotationEnabledChanged));
}

void Q3DInputHandler_DisconnectRotationEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DInputHandler*>(ptr), static_cast<void (Q3DInputHandler::*)(bool)>(&Q3DInputHandler::rotationEnabledChanged), static_cast<MyQ3DInputHandler*>(ptr), static_cast<void (MyQ3DInputHandler::*)(bool)>(&MyQ3DInputHandler::Signal_RotationEnabledChanged));
}

void Q3DInputHandler_RotationEnabledChanged(void* ptr, char enable)
{
	static_cast<Q3DInputHandler*>(ptr)->rotationEnabledChanged(enable != 0);
}

void Q3DInputHandler_ConnectSelectionEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DInputHandler*>(ptr), static_cast<void (Q3DInputHandler::*)(bool)>(&Q3DInputHandler::selectionEnabledChanged), static_cast<MyQ3DInputHandler*>(ptr), static_cast<void (MyQ3DInputHandler::*)(bool)>(&MyQ3DInputHandler::Signal_SelectionEnabledChanged));
}

void Q3DInputHandler_DisconnectSelectionEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DInputHandler*>(ptr), static_cast<void (Q3DInputHandler::*)(bool)>(&Q3DInputHandler::selectionEnabledChanged), static_cast<MyQ3DInputHandler*>(ptr), static_cast<void (MyQ3DInputHandler::*)(bool)>(&MyQ3DInputHandler::Signal_SelectionEnabledChanged));
}

void Q3DInputHandler_SelectionEnabledChanged(void* ptr, char enable)
{
	static_cast<Q3DInputHandler*>(ptr)->selectionEnabledChanged(enable != 0);
}

void Q3DInputHandler_SetRotationEnabled(void* ptr, char enable)
{
	static_cast<Q3DInputHandler*>(ptr)->setRotationEnabled(enable != 0);
}

void Q3DInputHandler_SetSelectionEnabled(void* ptr, char enable)
{
	static_cast<Q3DInputHandler*>(ptr)->setSelectionEnabled(enable != 0);
}

void Q3DInputHandler_SetZoomAtTargetEnabled(void* ptr, char enable)
{
	static_cast<Q3DInputHandler*>(ptr)->setZoomAtTargetEnabled(enable != 0);
}

void Q3DInputHandler_SetZoomEnabled(void* ptr, char enable)
{
	static_cast<Q3DInputHandler*>(ptr)->setZoomEnabled(enable != 0);
}

void Q3DInputHandler_WheelEvent(void* ptr, void* event)
{
	static_cast<Q3DInputHandler*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void Q3DInputHandler_WheelEventDefault(void* ptr, void* event)
{
		static_cast<Q3DInputHandler*>(ptr)->Q3DInputHandler::wheelEvent(static_cast<QWheelEvent*>(event));
}

void Q3DInputHandler_ConnectZoomAtTargetEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DInputHandler*>(ptr), static_cast<void (Q3DInputHandler::*)(bool)>(&Q3DInputHandler::zoomAtTargetEnabledChanged), static_cast<MyQ3DInputHandler*>(ptr), static_cast<void (MyQ3DInputHandler::*)(bool)>(&MyQ3DInputHandler::Signal_ZoomAtTargetEnabledChanged));
}

void Q3DInputHandler_DisconnectZoomAtTargetEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DInputHandler*>(ptr), static_cast<void (Q3DInputHandler::*)(bool)>(&Q3DInputHandler::zoomAtTargetEnabledChanged), static_cast<MyQ3DInputHandler*>(ptr), static_cast<void (MyQ3DInputHandler::*)(bool)>(&MyQ3DInputHandler::Signal_ZoomAtTargetEnabledChanged));
}

void Q3DInputHandler_ZoomAtTargetEnabledChanged(void* ptr, char enable)
{
	static_cast<Q3DInputHandler*>(ptr)->zoomAtTargetEnabledChanged(enable != 0);
}

void Q3DInputHandler_ConnectZoomEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DInputHandler*>(ptr), static_cast<void (Q3DInputHandler::*)(bool)>(&Q3DInputHandler::zoomEnabledChanged), static_cast<MyQ3DInputHandler*>(ptr), static_cast<void (MyQ3DInputHandler::*)(bool)>(&MyQ3DInputHandler::Signal_ZoomEnabledChanged));
}

void Q3DInputHandler_DisconnectZoomEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DInputHandler*>(ptr), static_cast<void (Q3DInputHandler::*)(bool)>(&Q3DInputHandler::zoomEnabledChanged), static_cast<MyQ3DInputHandler*>(ptr), static_cast<void (MyQ3DInputHandler::*)(bool)>(&MyQ3DInputHandler::Signal_ZoomEnabledChanged));
}

void Q3DInputHandler_ZoomEnabledChanged(void* ptr, char enable)
{
	static_cast<Q3DInputHandler*>(ptr)->zoomEnabledChanged(enable != 0);
}

void Q3DInputHandler_DestroyQ3DInputHandler(void* ptr)
{
	static_cast<Q3DInputHandler*>(ptr)->~Q3DInputHandler();
}

void Q3DInputHandler_DestroyQ3DInputHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Q3DInputHandler_IsRotationEnabled(void* ptr)
{
	return static_cast<Q3DInputHandler*>(ptr)->isRotationEnabled();
}

char Q3DInputHandler_IsSelectionEnabled(void* ptr)
{
	return static_cast<Q3DInputHandler*>(ptr)->isSelectionEnabled();
}

char Q3DInputHandler_IsZoomAtTargetEnabled(void* ptr)
{
	return static_cast<Q3DInputHandler*>(ptr)->isZoomAtTargetEnabled();
}

char Q3DInputHandler_IsZoomEnabled(void* ptr)
{
	return static_cast<Q3DInputHandler*>(ptr)->isZoomEnabled();
}

class MyQ3DLight: public Q3DLight
{
public:
	MyQ3DLight(QObject *parent = Q_NULLPTR) : Q3DLight(parent) {};
	void Signal_AutoPositionChanged(bool autoPosition) { callbackQ3DLight_AutoPositionChanged(this, autoPosition); };
	 ~MyQ3DLight() { callbackQ3DLight_DestroyQ3DLight(this); };
};

void* Q3DLight_NewQ3DLight(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DLight(static_cast<QWindow*>(parent));
	} else {
		return new MyQ3DLight(static_cast<QObject*>(parent));
	}
}

char Q3DLight_IsAutoPosition(void* ptr)
{
	return static_cast<Q3DLight*>(ptr)->isAutoPosition();
}

void Q3DLight_ConnectAutoPositionChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DLight*>(ptr), static_cast<void (Q3DLight::*)(bool)>(&Q3DLight::autoPositionChanged), static_cast<MyQ3DLight*>(ptr), static_cast<void (MyQ3DLight::*)(bool)>(&MyQ3DLight::Signal_AutoPositionChanged));
}

void Q3DLight_DisconnectAutoPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DLight*>(ptr), static_cast<void (Q3DLight::*)(bool)>(&Q3DLight::autoPositionChanged), static_cast<MyQ3DLight*>(ptr), static_cast<void (MyQ3DLight::*)(bool)>(&MyQ3DLight::Signal_AutoPositionChanged));
}

void Q3DLight_AutoPositionChanged(void* ptr, char autoPosition)
{
	static_cast<Q3DLight*>(ptr)->autoPositionChanged(autoPosition != 0);
}

void Q3DLight_SetAutoPosition(void* ptr, char enabled)
{
	static_cast<Q3DLight*>(ptr)->setAutoPosition(enabled != 0);
}

void Q3DLight_DestroyQ3DLight(void* ptr)
{
	static_cast<Q3DLight*>(ptr)->~Q3DLight();
}

void Q3DLight_DestroyQ3DLightDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQ3DObject: public Q3DObject
{
public:
	MyQ3DObject(QObject *parent = Q_NULLPTR) : Q3DObject(parent) {};
	void copyValuesFrom(const Q3DObject & source) { callbackQ3DObject_CopyValuesFrom(this, const_cast<Q3DObject*>(&source)); };
	void Signal_PositionChanged(const QVector3D & position) { callbackQ3DObject_PositionChanged(this, const_cast<QVector3D*>(&position)); };
	 ~MyQ3DObject() { callbackQ3DObject_DestroyQ3DObject(this); };
};

void* Q3DObject_NewQ3DObject(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DObject(static_cast<QWindow*>(parent));
	} else {
		return new MyQ3DObject(static_cast<QObject*>(parent));
	}
}

void* Q3DObject_ParentScene(void* ptr)
{
	return static_cast<Q3DObject*>(ptr)->parentScene();
}

void Q3DObject_CopyValuesFrom(void* ptr, void* source)
{
	static_cast<Q3DObject*>(ptr)->copyValuesFrom(*static_cast<Q3DObject*>(source));
}

void Q3DObject_CopyValuesFromDefault(void* ptr, void* source)
{
		static_cast<Q3DObject*>(ptr)->Q3DObject::copyValuesFrom(*static_cast<Q3DObject*>(source));
}

void Q3DObject_ConnectPositionChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DObject*>(ptr), static_cast<void (Q3DObject::*)(const QVector3D &)>(&Q3DObject::positionChanged), static_cast<MyQ3DObject*>(ptr), static_cast<void (MyQ3DObject::*)(const QVector3D &)>(&MyQ3DObject::Signal_PositionChanged));
}

void Q3DObject_DisconnectPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DObject*>(ptr), static_cast<void (Q3DObject::*)(const QVector3D &)>(&Q3DObject::positionChanged), static_cast<MyQ3DObject*>(ptr), static_cast<void (MyQ3DObject::*)(const QVector3D &)>(&MyQ3DObject::Signal_PositionChanged));
}

void Q3DObject_PositionChanged(void* ptr, void* position)
{
	static_cast<Q3DObject*>(ptr)->positionChanged(*static_cast<QVector3D*>(position));
}

void Q3DObject_SetDirty(void* ptr, char dirty)
{
	static_cast<Q3DObject*>(ptr)->setDirty(dirty != 0);
}

void Q3DObject_SetPosition(void* ptr, void* position)
{
	static_cast<Q3DObject*>(ptr)->setPosition(*static_cast<QVector3D*>(position));
}

void Q3DObject_DestroyQ3DObject(void* ptr)
{
	static_cast<Q3DObject*>(ptr)->~Q3DObject();
}

void Q3DObject_DestroyQ3DObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* Q3DObject_Position(void* ptr)
{
	return new QVector3D(static_cast<Q3DObject*>(ptr)->position());
}

char Q3DObject_IsDirty(void* ptr)
{
	return static_cast<Q3DObject*>(ptr)->isDirty();
}

class MyQ3DScatter: public Q3DScatter
{
public:
	MyQ3DScatter(const QSurfaceFormat *format = Q_NULLPTR, QWindow *parent = Q_NULLPTR) : Q3DScatter(format, parent) {};
	void Signal_AxisXChanged(QValue3DAxis * axis) { callbackQ3DScatter_AxisXChanged(this, axis); };
	void Signal_AxisYChanged(QValue3DAxis * axis) { callbackQ3DScatter_AxisYChanged(this, axis); };
	void Signal_AxisZChanged(QValue3DAxis * axis) { callbackQ3DScatter_AxisZChanged(this, axis); };
	void Signal_SelectedSeriesChanged(QScatter3DSeries * series) { callbackQ3DScatter_SelectedSeriesChanged(this, series); };
	 ~MyQ3DScatter() { callbackQ3DScatter_DestroyQ3DScatter(this); };
};

void* Q3DScatter_NewQ3DScatter(void* format, void* parent)
{
	if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScatter(static_cast<QSurfaceFormat*>(format), static_cast<QPaintDeviceWindow*>(parent));
	} else {
		return new MyQ3DScatter(static_cast<QSurfaceFormat*>(format), static_cast<QWindow*>(parent));
	}
}

void Q3DScatter_AddAxis(void* ptr, void* axis)
{
	static_cast<Q3DScatter*>(ptr)->addAxis(static_cast<QValue3DAxis*>(axis));
}

void Q3DScatter_AddSeries(void* ptr, void* series)
{
	static_cast<Q3DScatter*>(ptr)->addSeries(static_cast<QScatter3DSeries*>(series));
}

void Q3DScatter_ConnectAxisXChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScatter*>(ptr), static_cast<void (Q3DScatter::*)(QValue3DAxis *)>(&Q3DScatter::axisXChanged), static_cast<MyQ3DScatter*>(ptr), static_cast<void (MyQ3DScatter::*)(QValue3DAxis *)>(&MyQ3DScatter::Signal_AxisXChanged));
}

void Q3DScatter_DisconnectAxisXChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScatter*>(ptr), static_cast<void (Q3DScatter::*)(QValue3DAxis *)>(&Q3DScatter::axisXChanged), static_cast<MyQ3DScatter*>(ptr), static_cast<void (MyQ3DScatter::*)(QValue3DAxis *)>(&MyQ3DScatter::Signal_AxisXChanged));
}

void Q3DScatter_AxisXChanged(void* ptr, void* axis)
{
	static_cast<Q3DScatter*>(ptr)->axisXChanged(static_cast<QValue3DAxis*>(axis));
}

void Q3DScatter_ConnectAxisYChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScatter*>(ptr), static_cast<void (Q3DScatter::*)(QValue3DAxis *)>(&Q3DScatter::axisYChanged), static_cast<MyQ3DScatter*>(ptr), static_cast<void (MyQ3DScatter::*)(QValue3DAxis *)>(&MyQ3DScatter::Signal_AxisYChanged));
}

void Q3DScatter_DisconnectAxisYChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScatter*>(ptr), static_cast<void (Q3DScatter::*)(QValue3DAxis *)>(&Q3DScatter::axisYChanged), static_cast<MyQ3DScatter*>(ptr), static_cast<void (MyQ3DScatter::*)(QValue3DAxis *)>(&MyQ3DScatter::Signal_AxisYChanged));
}

void Q3DScatter_AxisYChanged(void* ptr, void* axis)
{
	static_cast<Q3DScatter*>(ptr)->axisYChanged(static_cast<QValue3DAxis*>(axis));
}

void Q3DScatter_ConnectAxisZChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScatter*>(ptr), static_cast<void (Q3DScatter::*)(QValue3DAxis *)>(&Q3DScatter::axisZChanged), static_cast<MyQ3DScatter*>(ptr), static_cast<void (MyQ3DScatter::*)(QValue3DAxis *)>(&MyQ3DScatter::Signal_AxisZChanged));
}

void Q3DScatter_DisconnectAxisZChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScatter*>(ptr), static_cast<void (Q3DScatter::*)(QValue3DAxis *)>(&Q3DScatter::axisZChanged), static_cast<MyQ3DScatter*>(ptr), static_cast<void (MyQ3DScatter::*)(QValue3DAxis *)>(&MyQ3DScatter::Signal_AxisZChanged));
}

void Q3DScatter_AxisZChanged(void* ptr, void* axis)
{
	static_cast<Q3DScatter*>(ptr)->axisZChanged(static_cast<QValue3DAxis*>(axis));
}

void Q3DScatter_ReleaseAxis(void* ptr, void* axis)
{
	static_cast<Q3DScatter*>(ptr)->releaseAxis(static_cast<QValue3DAxis*>(axis));
}

void Q3DScatter_RemoveSeries(void* ptr, void* series)
{
	static_cast<Q3DScatter*>(ptr)->removeSeries(static_cast<QScatter3DSeries*>(series));
}

void Q3DScatter_ConnectSelectedSeriesChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScatter*>(ptr), static_cast<void (Q3DScatter::*)(QScatter3DSeries *)>(&Q3DScatter::selectedSeriesChanged), static_cast<MyQ3DScatter*>(ptr), static_cast<void (MyQ3DScatter::*)(QScatter3DSeries *)>(&MyQ3DScatter::Signal_SelectedSeriesChanged));
}

void Q3DScatter_DisconnectSelectedSeriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScatter*>(ptr), static_cast<void (Q3DScatter::*)(QScatter3DSeries *)>(&Q3DScatter::selectedSeriesChanged), static_cast<MyQ3DScatter*>(ptr), static_cast<void (MyQ3DScatter::*)(QScatter3DSeries *)>(&MyQ3DScatter::Signal_SelectedSeriesChanged));
}

void Q3DScatter_SelectedSeriesChanged(void* ptr, void* series)
{
	static_cast<Q3DScatter*>(ptr)->selectedSeriesChanged(static_cast<QScatter3DSeries*>(series));
}

void Q3DScatter_SetAxisX(void* ptr, void* axis)
{
	static_cast<Q3DScatter*>(ptr)->setAxisX(static_cast<QValue3DAxis*>(axis));
}

void Q3DScatter_SetAxisY(void* ptr, void* axis)
{
	static_cast<Q3DScatter*>(ptr)->setAxisY(static_cast<QValue3DAxis*>(axis));
}

void Q3DScatter_SetAxisZ(void* ptr, void* axis)
{
	static_cast<Q3DScatter*>(ptr)->setAxisZ(static_cast<QValue3DAxis*>(axis));
}

void Q3DScatter_DestroyQ3DScatter(void* ptr)
{
	static_cast<Q3DScatter*>(ptr)->~Q3DScatter();
}

void Q3DScatter_DestroyQ3DScatterDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDataVisualization_PackedList Q3DScatter_SeriesList(void* ptr)
{
	return ({ QList<QScatter3DSeries *>* tmpValue = new QList<QScatter3DSeries *>(static_cast<Q3DScatter*>(ptr)->seriesList()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDataVisualization_PackedList Q3DScatter_Axes(void* ptr)
{
	return ({ QList<QValue3DAxis *>* tmpValue = new QList<QValue3DAxis *>(static_cast<Q3DScatter*>(ptr)->axes()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

void* Q3DScatter_SelectedSeries(void* ptr)
{
	return static_cast<Q3DScatter*>(ptr)->selectedSeries();
}

void* Q3DScatter_AxisX(void* ptr)
{
	return static_cast<Q3DScatter*>(ptr)->axisX();
}

void* Q3DScatter_AxisY(void* ptr)
{
	return static_cast<Q3DScatter*>(ptr)->axisY();
}

void* Q3DScatter_AxisZ(void* ptr)
{
	return static_cast<Q3DScatter*>(ptr)->axisZ();
}

void* Q3DScatter___seriesList_atList(void* ptr, int i)
{
	return const_cast<QScatter3DSeries*>(static_cast<QList<QScatter3DSeries *>*>(ptr)->at(i));
}

void Q3DScatter___seriesList_setList(void* ptr, void* i)
{
	static_cast<QList<QScatter3DSeries *>*>(ptr)->append(static_cast<QScatter3DSeries*>(i));
}

void* Q3DScatter___seriesList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScatter3DSeries *>;
}

void* Q3DScatter___axes_atList(void* ptr, int i)
{
	return const_cast<QValue3DAxis*>(static_cast<QList<QValue3DAxis *>*>(ptr)->at(i));
}

void Q3DScatter___axes_setList(void* ptr, void* i)
{
	static_cast<QList<QValue3DAxis *>*>(ptr)->append(static_cast<QValue3DAxis*>(i));
}

void* Q3DScatter___axes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QValue3DAxis *>;
}

class MyQ3DScene: public Q3DScene
{
public:
	MyQ3DScene(QObject *parent = Q_NULLPTR) : Q3DScene(parent) {};
	void Signal_ActiveCameraChanged(Q3DCamera * camera) { callbackQ3DScene_ActiveCameraChanged(this, camera); };
	void Signal_ActiveLightChanged(Q3DLight * light) { callbackQ3DScene_ActiveLightChanged(this, light); };
	void Signal_DevicePixelRatioChanged(float pixelRatio) { callbackQ3DScene_DevicePixelRatioChanged(this, pixelRatio); };
	void Signal_GraphPositionQueryChanged(const QPoint & position) { callbackQ3DScene_GraphPositionQueryChanged(this, const_cast<QPoint*>(&position)); };
	void Signal_PrimarySubViewportChanged(const QRect & subViewport) { callbackQ3DScene_PrimarySubViewportChanged(this, const_cast<QRect*>(&subViewport)); };
	void Signal_SecondarySubViewportChanged(const QRect & subViewport) { callbackQ3DScene_SecondarySubViewportChanged(this, const_cast<QRect*>(&subViewport)); };
	void Signal_SecondarySubviewOnTopChanged(bool isSecondaryOnTop) { callbackQ3DScene_SecondarySubviewOnTopChanged(this, isSecondaryOnTop); };
	void Signal_SelectionQueryPositionChanged(const QPoint & position) { callbackQ3DScene_SelectionQueryPositionChanged(this, const_cast<QPoint*>(&position)); };
	void Signal_SlicingActiveChanged(bool isSlicingActive) { callbackQ3DScene_SlicingActiveChanged(this, isSlicingActive); };
	void Signal_ViewportChanged(const QRect & viewport) { callbackQ3DScene_ViewportChanged(this, const_cast<QRect*>(&viewport)); };
	 ~MyQ3DScene() { callbackQ3DScene_DestroyQ3DScene(this); };
};

void* Q3DScene_NewQ3DScene(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DScene(static_cast<QWindow*>(parent));
	} else {
		return new MyQ3DScene(static_cast<QObject*>(parent));
	}
}

void* Q3DScene_Q3DScene_InvalidSelectionPoint()
{
	return ({ QPoint tmpValue = Q3DScene::invalidSelectionPoint(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

char Q3DScene_IsPointInPrimarySubView(void* ptr, void* point)
{
	return static_cast<Q3DScene*>(ptr)->isPointInPrimarySubView(*static_cast<QPoint*>(point));
}

char Q3DScene_IsPointInSecondarySubView(void* ptr, void* point)
{
	return static_cast<Q3DScene*>(ptr)->isPointInSecondarySubView(*static_cast<QPoint*>(point));
}

void Q3DScene_ConnectActiveCameraChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(Q3DCamera *)>(&Q3DScene::activeCameraChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(Q3DCamera *)>(&MyQ3DScene::Signal_ActiveCameraChanged));
}

void Q3DScene_DisconnectActiveCameraChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(Q3DCamera *)>(&Q3DScene::activeCameraChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(Q3DCamera *)>(&MyQ3DScene::Signal_ActiveCameraChanged));
}

void Q3DScene_ActiveCameraChanged(void* ptr, void* camera)
{
	static_cast<Q3DScene*>(ptr)->activeCameraChanged(static_cast<Q3DCamera*>(camera));
}

void Q3DScene_ConnectActiveLightChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(Q3DLight *)>(&Q3DScene::activeLightChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(Q3DLight *)>(&MyQ3DScene::Signal_ActiveLightChanged));
}

void Q3DScene_DisconnectActiveLightChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(Q3DLight *)>(&Q3DScene::activeLightChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(Q3DLight *)>(&MyQ3DScene::Signal_ActiveLightChanged));
}

void Q3DScene_ActiveLightChanged(void* ptr, void* light)
{
	static_cast<Q3DScene*>(ptr)->activeLightChanged(static_cast<Q3DLight*>(light));
}

void Q3DScene_ConnectDevicePixelRatioChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(float)>(&Q3DScene::devicePixelRatioChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(float)>(&MyQ3DScene::Signal_DevicePixelRatioChanged));
}

void Q3DScene_DisconnectDevicePixelRatioChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(float)>(&Q3DScene::devicePixelRatioChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(float)>(&MyQ3DScene::Signal_DevicePixelRatioChanged));
}

void Q3DScene_DevicePixelRatioChanged(void* ptr, float pixelRatio)
{
	static_cast<Q3DScene*>(ptr)->devicePixelRatioChanged(pixelRatio);
}

void Q3DScene_ConnectGraphPositionQueryChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QPoint &)>(&Q3DScene::graphPositionQueryChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QPoint &)>(&MyQ3DScene::Signal_GraphPositionQueryChanged));
}

void Q3DScene_DisconnectGraphPositionQueryChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QPoint &)>(&Q3DScene::graphPositionQueryChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QPoint &)>(&MyQ3DScene::Signal_GraphPositionQueryChanged));
}

void Q3DScene_GraphPositionQueryChanged(void* ptr, void* position)
{
	static_cast<Q3DScene*>(ptr)->graphPositionQueryChanged(*static_cast<QPoint*>(position));
}

void Q3DScene_ConnectPrimarySubViewportChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QRect &)>(&Q3DScene::primarySubViewportChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QRect &)>(&MyQ3DScene::Signal_PrimarySubViewportChanged));
}

void Q3DScene_DisconnectPrimarySubViewportChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QRect &)>(&Q3DScene::primarySubViewportChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QRect &)>(&MyQ3DScene::Signal_PrimarySubViewportChanged));
}

void Q3DScene_PrimarySubViewportChanged(void* ptr, void* subViewport)
{
	static_cast<Q3DScene*>(ptr)->primarySubViewportChanged(*static_cast<QRect*>(subViewport));
}

void Q3DScene_ConnectSecondarySubViewportChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QRect &)>(&Q3DScene::secondarySubViewportChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QRect &)>(&MyQ3DScene::Signal_SecondarySubViewportChanged));
}

void Q3DScene_DisconnectSecondarySubViewportChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QRect &)>(&Q3DScene::secondarySubViewportChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QRect &)>(&MyQ3DScene::Signal_SecondarySubViewportChanged));
}

void Q3DScene_SecondarySubViewportChanged(void* ptr, void* subViewport)
{
	static_cast<Q3DScene*>(ptr)->secondarySubViewportChanged(*static_cast<QRect*>(subViewport));
}

void Q3DScene_ConnectSecondarySubviewOnTopChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(bool)>(&Q3DScene::secondarySubviewOnTopChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(bool)>(&MyQ3DScene::Signal_SecondarySubviewOnTopChanged));
}

void Q3DScene_DisconnectSecondarySubviewOnTopChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(bool)>(&Q3DScene::secondarySubviewOnTopChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(bool)>(&MyQ3DScene::Signal_SecondarySubviewOnTopChanged));
}

void Q3DScene_SecondarySubviewOnTopChanged(void* ptr, char isSecondaryOnTop)
{
	static_cast<Q3DScene*>(ptr)->secondarySubviewOnTopChanged(isSecondaryOnTop != 0);
}

void Q3DScene_ConnectSelectionQueryPositionChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QPoint &)>(&Q3DScene::selectionQueryPositionChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QPoint &)>(&MyQ3DScene::Signal_SelectionQueryPositionChanged));
}

void Q3DScene_DisconnectSelectionQueryPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QPoint &)>(&Q3DScene::selectionQueryPositionChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QPoint &)>(&MyQ3DScene::Signal_SelectionQueryPositionChanged));
}

void Q3DScene_SelectionQueryPositionChanged(void* ptr, void* position)
{
	static_cast<Q3DScene*>(ptr)->selectionQueryPositionChanged(*static_cast<QPoint*>(position));
}

void Q3DScene_SetActiveCamera(void* ptr, void* camera)
{
	static_cast<Q3DScene*>(ptr)->setActiveCamera(static_cast<Q3DCamera*>(camera));
}

void Q3DScene_SetActiveLight(void* ptr, void* light)
{
	static_cast<Q3DScene*>(ptr)->setActiveLight(static_cast<Q3DLight*>(light));
}

void Q3DScene_SetDevicePixelRatio(void* ptr, float pixelRatio)
{
	static_cast<Q3DScene*>(ptr)->setDevicePixelRatio(pixelRatio);
}

void Q3DScene_SetGraphPositionQuery(void* ptr, void* point)
{
	static_cast<Q3DScene*>(ptr)->setGraphPositionQuery(*static_cast<QPoint*>(point));
}

void Q3DScene_SetPrimarySubViewport(void* ptr, void* primarySubViewport)
{
	static_cast<Q3DScene*>(ptr)->setPrimarySubViewport(*static_cast<QRect*>(primarySubViewport));
}

void Q3DScene_SetSecondarySubViewport(void* ptr, void* secondarySubViewport)
{
	static_cast<Q3DScene*>(ptr)->setSecondarySubViewport(*static_cast<QRect*>(secondarySubViewport));
}

void Q3DScene_SetSecondarySubviewOnTop(void* ptr, char isSecondaryOnTop)
{
	static_cast<Q3DScene*>(ptr)->setSecondarySubviewOnTop(isSecondaryOnTop != 0);
}

void Q3DScene_SetSelectionQueryPosition(void* ptr, void* point)
{
	static_cast<Q3DScene*>(ptr)->setSelectionQueryPosition(*static_cast<QPoint*>(point));
}

void Q3DScene_SetSlicingActive(void* ptr, char isSlicing)
{
	static_cast<Q3DScene*>(ptr)->setSlicingActive(isSlicing != 0);
}

void Q3DScene_ConnectSlicingActiveChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(bool)>(&Q3DScene::slicingActiveChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(bool)>(&MyQ3DScene::Signal_SlicingActiveChanged));
}

void Q3DScene_DisconnectSlicingActiveChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(bool)>(&Q3DScene::slicingActiveChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(bool)>(&MyQ3DScene::Signal_SlicingActiveChanged));
}

void Q3DScene_SlicingActiveChanged(void* ptr, char isSlicingActive)
{
	static_cast<Q3DScene*>(ptr)->slicingActiveChanged(isSlicingActive != 0);
}

void Q3DScene_ConnectViewportChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QRect &)>(&Q3DScene::viewportChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QRect &)>(&MyQ3DScene::Signal_ViewportChanged));
}

void Q3DScene_DisconnectViewportChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DScene*>(ptr), static_cast<void (Q3DScene::*)(const QRect &)>(&Q3DScene::viewportChanged), static_cast<MyQ3DScene*>(ptr), static_cast<void (MyQ3DScene::*)(const QRect &)>(&MyQ3DScene::Signal_ViewportChanged));
}

void Q3DScene_ViewportChanged(void* ptr, void* viewport)
{
	static_cast<Q3DScene*>(ptr)->viewportChanged(*static_cast<QRect*>(viewport));
}

void Q3DScene_DestroyQ3DScene(void* ptr)
{
	static_cast<Q3DScene*>(ptr)->~Q3DScene();
}

void Q3DScene_DestroyQ3DSceneDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* Q3DScene_ActiveCamera(void* ptr)
{
	return static_cast<Q3DScene*>(ptr)->activeCamera();
}

void* Q3DScene_ActiveLight(void* ptr)
{
	return static_cast<Q3DScene*>(ptr)->activeLight();
}

void* Q3DScene_GraphPositionQuery(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<Q3DScene*>(ptr)->graphPositionQuery(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* Q3DScene_SelectionQueryPosition(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<Q3DScene*>(ptr)->selectionQueryPosition(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* Q3DScene_PrimarySubViewport(void* ptr)
{
	return ({ QRect tmpValue = static_cast<Q3DScene*>(ptr)->primarySubViewport(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* Q3DScene_SecondarySubViewport(void* ptr)
{
	return ({ QRect tmpValue = static_cast<Q3DScene*>(ptr)->secondarySubViewport(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* Q3DScene_Viewport(void* ptr)
{
	return ({ QRect tmpValue = static_cast<Q3DScene*>(ptr)->viewport(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

char Q3DScene_IsSecondarySubviewOnTop(void* ptr)
{
	return static_cast<Q3DScene*>(ptr)->isSecondarySubviewOnTop();
}

char Q3DScene_IsSlicingActive(void* ptr)
{
	return static_cast<Q3DScene*>(ptr)->isSlicingActive();
}

float Q3DScene_DevicePixelRatio(void* ptr)
{
	return static_cast<Q3DScene*>(ptr)->devicePixelRatio();
}

class MyQ3DSurface: public Q3DSurface
{
public:
	MyQ3DSurface(const QSurfaceFormat *format = Q_NULLPTR, QWindow *parent = Q_NULLPTR) : Q3DSurface(format, parent) {};
	void Signal_AxisXChanged(QValue3DAxis * axis) { callbackQ3DSurface_AxisXChanged(this, axis); };
	void Signal_AxisYChanged(QValue3DAxis * axis) { callbackQ3DSurface_AxisYChanged(this, axis); };
	void Signal_AxisZChanged(QValue3DAxis * axis) { callbackQ3DSurface_AxisZChanged(this, axis); };
	void Signal_FlipHorizontalGridChanged(bool flip) { callbackQ3DSurface_FlipHorizontalGridChanged(this, flip); };
	void Signal_SelectedSeriesChanged(QSurface3DSeries * series) { callbackQ3DSurface_SelectedSeriesChanged(this, series); };
	 ~MyQ3DSurface() { callbackQ3DSurface_DestroyQ3DSurface(this); };
};

void* Q3DSurface_NewQ3DSurface(void* format, void* parent)
{
	if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DSurface(static_cast<QSurfaceFormat*>(format), static_cast<QPaintDeviceWindow*>(parent));
	} else {
		return new MyQ3DSurface(static_cast<QSurfaceFormat*>(format), static_cast<QWindow*>(parent));
	}
}

void Q3DSurface_AddAxis(void* ptr, void* axis)
{
	static_cast<Q3DSurface*>(ptr)->addAxis(static_cast<QValue3DAxis*>(axis));
}

void Q3DSurface_AddSeries(void* ptr, void* series)
{
	static_cast<Q3DSurface*>(ptr)->addSeries(static_cast<QSurface3DSeries*>(series));
}

void Q3DSurface_ConnectAxisXChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(QValue3DAxis *)>(&Q3DSurface::axisXChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(QValue3DAxis *)>(&MyQ3DSurface::Signal_AxisXChanged));
}

void Q3DSurface_DisconnectAxisXChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(QValue3DAxis *)>(&Q3DSurface::axisXChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(QValue3DAxis *)>(&MyQ3DSurface::Signal_AxisXChanged));
}

void Q3DSurface_AxisXChanged(void* ptr, void* axis)
{
	static_cast<Q3DSurface*>(ptr)->axisXChanged(static_cast<QValue3DAxis*>(axis));
}

void Q3DSurface_ConnectAxisYChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(QValue3DAxis *)>(&Q3DSurface::axisYChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(QValue3DAxis *)>(&MyQ3DSurface::Signal_AxisYChanged));
}

void Q3DSurface_DisconnectAxisYChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(QValue3DAxis *)>(&Q3DSurface::axisYChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(QValue3DAxis *)>(&MyQ3DSurface::Signal_AxisYChanged));
}

void Q3DSurface_AxisYChanged(void* ptr, void* axis)
{
	static_cast<Q3DSurface*>(ptr)->axisYChanged(static_cast<QValue3DAxis*>(axis));
}

void Q3DSurface_ConnectAxisZChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(QValue3DAxis *)>(&Q3DSurface::axisZChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(QValue3DAxis *)>(&MyQ3DSurface::Signal_AxisZChanged));
}

void Q3DSurface_DisconnectAxisZChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(QValue3DAxis *)>(&Q3DSurface::axisZChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(QValue3DAxis *)>(&MyQ3DSurface::Signal_AxisZChanged));
}

void Q3DSurface_AxisZChanged(void* ptr, void* axis)
{
	static_cast<Q3DSurface*>(ptr)->axisZChanged(static_cast<QValue3DAxis*>(axis));
}

void Q3DSurface_ConnectFlipHorizontalGridChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(bool)>(&Q3DSurface::flipHorizontalGridChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(bool)>(&MyQ3DSurface::Signal_FlipHorizontalGridChanged));
}

void Q3DSurface_DisconnectFlipHorizontalGridChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(bool)>(&Q3DSurface::flipHorizontalGridChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(bool)>(&MyQ3DSurface::Signal_FlipHorizontalGridChanged));
}

void Q3DSurface_FlipHorizontalGridChanged(void* ptr, char flip)
{
	static_cast<Q3DSurface*>(ptr)->flipHorizontalGridChanged(flip != 0);
}

void Q3DSurface_ReleaseAxis(void* ptr, void* axis)
{
	static_cast<Q3DSurface*>(ptr)->releaseAxis(static_cast<QValue3DAxis*>(axis));
}

void Q3DSurface_RemoveSeries(void* ptr, void* series)
{
	static_cast<Q3DSurface*>(ptr)->removeSeries(static_cast<QSurface3DSeries*>(series));
}

void Q3DSurface_ConnectSelectedSeriesChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(QSurface3DSeries *)>(&Q3DSurface::selectedSeriesChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(QSurface3DSeries *)>(&MyQ3DSurface::Signal_SelectedSeriesChanged));
}

void Q3DSurface_DisconnectSelectedSeriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DSurface*>(ptr), static_cast<void (Q3DSurface::*)(QSurface3DSeries *)>(&Q3DSurface::selectedSeriesChanged), static_cast<MyQ3DSurface*>(ptr), static_cast<void (MyQ3DSurface::*)(QSurface3DSeries *)>(&MyQ3DSurface::Signal_SelectedSeriesChanged));
}

void Q3DSurface_SelectedSeriesChanged(void* ptr, void* series)
{
	static_cast<Q3DSurface*>(ptr)->selectedSeriesChanged(static_cast<QSurface3DSeries*>(series));
}

void Q3DSurface_SetAxisX(void* ptr, void* axis)
{
	static_cast<Q3DSurface*>(ptr)->setAxisX(static_cast<QValue3DAxis*>(axis));
}

void Q3DSurface_SetAxisY(void* ptr, void* axis)
{
	static_cast<Q3DSurface*>(ptr)->setAxisY(static_cast<QValue3DAxis*>(axis));
}

void Q3DSurface_SetAxisZ(void* ptr, void* axis)
{
	static_cast<Q3DSurface*>(ptr)->setAxisZ(static_cast<QValue3DAxis*>(axis));
}

void Q3DSurface_SetFlipHorizontalGrid(void* ptr, char flip)
{
	static_cast<Q3DSurface*>(ptr)->setFlipHorizontalGrid(flip != 0);
}

void Q3DSurface_DestroyQ3DSurface(void* ptr)
{
	static_cast<Q3DSurface*>(ptr)->~Q3DSurface();
}

void Q3DSurface_DestroyQ3DSurfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDataVisualization_PackedList Q3DSurface_SeriesList(void* ptr)
{
	return ({ QList<QSurface3DSeries *>* tmpValue = new QList<QSurface3DSeries *>(static_cast<Q3DSurface*>(ptr)->seriesList()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDataVisualization_PackedList Q3DSurface_Axes(void* ptr)
{
	return ({ QList<QValue3DAxis *>* tmpValue = new QList<QValue3DAxis *>(static_cast<Q3DSurface*>(ptr)->axes()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

void* Q3DSurface_SelectedSeries(void* ptr)
{
	return static_cast<Q3DSurface*>(ptr)->selectedSeries();
}

void* Q3DSurface_AxisX(void* ptr)
{
	return static_cast<Q3DSurface*>(ptr)->axisX();
}

void* Q3DSurface_AxisY(void* ptr)
{
	return static_cast<Q3DSurface*>(ptr)->axisY();
}

void* Q3DSurface_AxisZ(void* ptr)
{
	return static_cast<Q3DSurface*>(ptr)->axisZ();
}

char Q3DSurface_FlipHorizontalGrid(void* ptr)
{
	return static_cast<Q3DSurface*>(ptr)->flipHorizontalGrid();
}

void* Q3DSurface___seriesList_atList(void* ptr, int i)
{
	return const_cast<QSurface3DSeries*>(static_cast<QList<QSurface3DSeries *>*>(ptr)->at(i));
}

void Q3DSurface___seriesList_setList(void* ptr, void* i)
{
	static_cast<QList<QSurface3DSeries *>*>(ptr)->append(static_cast<QSurface3DSeries*>(i));
}

void* Q3DSurface___seriesList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSurface3DSeries *>;
}

void* Q3DSurface___axes_atList(void* ptr, int i)
{
	return const_cast<QValue3DAxis*>(static_cast<QList<QValue3DAxis *>*>(ptr)->at(i));
}

void Q3DSurface___axes_setList(void* ptr, void* i)
{
	static_cast<QList<QValue3DAxis *>*>(ptr)->append(static_cast<QValue3DAxis*>(i));
}

void* Q3DSurface___axes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QValue3DAxis *>;
}

class MyQ3DTheme: public Q3DTheme
{
public:
	MyQ3DTheme(QObject *parent = Q_NULLPTR) : Q3DTheme(parent) {};
	MyQ3DTheme(Theme themeType, QObject *parent = Q_NULLPTR) : Q3DTheme(themeType, parent) {};
	void Signal_AmbientLightStrengthChanged(float strength) { callbackQ3DTheme_AmbientLightStrengthChanged(this, strength); };
	void Signal_BackgroundColorChanged(const QColor & color) { callbackQ3DTheme_BackgroundColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_BackgroundEnabledChanged(bool enabled) { callbackQ3DTheme_BackgroundEnabledChanged(this, enabled); };
	void Signal_BaseColorsChanged(const QList<QColor> & colors) { callbackQ3DTheme_BaseColorsChanged(this, ({ QList<QColor>* tmpValue = const_cast<QList<QColor>*>(&colors); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_BaseGradientsChanged(const QList<QLinearGradient> & gradients) { callbackQ3DTheme_BaseGradientsChanged(this, ({ QList<QLinearGradient>* tmpValue = const_cast<QList<QLinearGradient>*>(&gradients); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_ColorStyleChanged(Q3DTheme::ColorStyle style) { callbackQ3DTheme_ColorStyleChanged(this, style); };
	void Signal_FontChanged(const QFont & font) { callbackQ3DTheme_FontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_GridEnabledChanged(bool enabled) { callbackQ3DTheme_GridEnabledChanged(this, enabled); };
	void Signal_GridLineColorChanged(const QColor & color) { callbackQ3DTheme_GridLineColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_HighlightLightStrengthChanged(float strength) { callbackQ3DTheme_HighlightLightStrengthChanged(this, strength); };
	void Signal_LabelBackgroundColorChanged(const QColor & color) { callbackQ3DTheme_LabelBackgroundColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_LabelBackgroundEnabledChanged(bool enabled) { callbackQ3DTheme_LabelBackgroundEnabledChanged(this, enabled); };
	void Signal_LabelBorderEnabledChanged(bool enabled) { callbackQ3DTheme_LabelBorderEnabledChanged(this, enabled); };
	void Signal_LabelTextColorChanged(const QColor & color) { callbackQ3DTheme_LabelTextColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_LightColorChanged(const QColor & color) { callbackQ3DTheme_LightColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_LightStrengthChanged(float strength) { callbackQ3DTheme_LightStrengthChanged(this, strength); };
	void Signal_MultiHighlightColorChanged(const QColor & color) { callbackQ3DTheme_MultiHighlightColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_MultiHighlightGradientChanged(const QLinearGradient & gradient) { callbackQ3DTheme_MultiHighlightGradientChanged(this, const_cast<QLinearGradient*>(&gradient)); };
	void Signal_SingleHighlightColorChanged(const QColor & color) { callbackQ3DTheme_SingleHighlightColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_SingleHighlightGradientChanged(const QLinearGradient & gradient) { callbackQ3DTheme_SingleHighlightGradientChanged(this, const_cast<QLinearGradient*>(&gradient)); };
	void Signal_TypeChanged(Q3DTheme::Theme themeType) { callbackQ3DTheme_TypeChanged(this, themeType); };
	void Signal_WindowColorChanged(const QColor & color) { callbackQ3DTheme_WindowColorChanged(this, const_cast<QColor*>(&color)); };
	 ~MyQ3DTheme() { callbackQ3DTheme_DestroyQ3DTheme(this); };
};

void* Q3DTheme_NewQ3DTheme(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<QWindow*>(parent));
	} else {
		return new MyQ3DTheme(static_cast<QObject*>(parent));
	}
}

void* Q3DTheme_NewQ3DTheme2(long long themeType, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QWindow*>(parent));
	} else {
		return new MyQ3DTheme(static_cast<Q3DTheme::Theme>(themeType), static_cast<QObject*>(parent));
	}
}

void Q3DTheme_ConnectAmbientLightStrengthChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(float)>(&Q3DTheme::ambientLightStrengthChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(float)>(&MyQ3DTheme::Signal_AmbientLightStrengthChanged));
}

void Q3DTheme_DisconnectAmbientLightStrengthChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(float)>(&Q3DTheme::ambientLightStrengthChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(float)>(&MyQ3DTheme::Signal_AmbientLightStrengthChanged));
}

void Q3DTheme_AmbientLightStrengthChanged(void* ptr, float strength)
{
	static_cast<Q3DTheme*>(ptr)->ambientLightStrengthChanged(strength);
}

void Q3DTheme_ConnectBackgroundColorChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::backgroundColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_BackgroundColorChanged));
}

void Q3DTheme_DisconnectBackgroundColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::backgroundColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_BackgroundColorChanged));
}

void Q3DTheme_BackgroundColorChanged(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->backgroundColorChanged(*static_cast<QColor*>(color));
}

void Q3DTheme_ConnectBackgroundEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(bool)>(&Q3DTheme::backgroundEnabledChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(bool)>(&MyQ3DTheme::Signal_BackgroundEnabledChanged));
}

void Q3DTheme_DisconnectBackgroundEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(bool)>(&Q3DTheme::backgroundEnabledChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(bool)>(&MyQ3DTheme::Signal_BackgroundEnabledChanged));
}

void Q3DTheme_BackgroundEnabledChanged(void* ptr, char enabled)
{
	static_cast<Q3DTheme*>(ptr)->backgroundEnabledChanged(enabled != 0);
}

void Q3DTheme_ConnectBaseColorsChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QList<QColor> &)>(&Q3DTheme::baseColorsChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QList<QColor> &)>(&MyQ3DTheme::Signal_BaseColorsChanged));
}

void Q3DTheme_DisconnectBaseColorsChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QList<QColor> &)>(&Q3DTheme::baseColorsChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QList<QColor> &)>(&MyQ3DTheme::Signal_BaseColorsChanged));
}

void Q3DTheme_BaseColorsChanged(void* ptr, void* colors)
{
	static_cast<Q3DTheme*>(ptr)->baseColorsChanged(*static_cast<QList<QColor>*>(colors));
}

void Q3DTheme_ConnectBaseGradientsChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QList<QLinearGradient> &)>(&Q3DTheme::baseGradientsChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QList<QLinearGradient> &)>(&MyQ3DTheme::Signal_BaseGradientsChanged));
}

void Q3DTheme_DisconnectBaseGradientsChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QList<QLinearGradient> &)>(&Q3DTheme::baseGradientsChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QList<QLinearGradient> &)>(&MyQ3DTheme::Signal_BaseGradientsChanged));
}

void Q3DTheme_BaseGradientsChanged(void* ptr, void* gradients)
{
	static_cast<Q3DTheme*>(ptr)->baseGradientsChanged(*static_cast<QList<QLinearGradient>*>(gradients));
}

void Q3DTheme_ConnectColorStyleChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(Q3DTheme::ColorStyle)>(&Q3DTheme::colorStyleChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(Q3DTheme::ColorStyle)>(&MyQ3DTheme::Signal_ColorStyleChanged));
}

void Q3DTheme_DisconnectColorStyleChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(Q3DTheme::ColorStyle)>(&Q3DTheme::colorStyleChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(Q3DTheme::ColorStyle)>(&MyQ3DTheme::Signal_ColorStyleChanged));
}

void Q3DTheme_ColorStyleChanged(void* ptr, long long style)
{
	static_cast<Q3DTheme*>(ptr)->colorStyleChanged(static_cast<Q3DTheme::ColorStyle>(style));
}

void Q3DTheme_ConnectFontChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QFont &)>(&Q3DTheme::fontChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QFont &)>(&MyQ3DTheme::Signal_FontChanged));
}

void Q3DTheme_DisconnectFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QFont &)>(&Q3DTheme::fontChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QFont &)>(&MyQ3DTheme::Signal_FontChanged));
}

void Q3DTheme_FontChanged(void* ptr, void* font)
{
	static_cast<Q3DTheme*>(ptr)->fontChanged(*static_cast<QFont*>(font));
}

void Q3DTheme_ConnectGridEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(bool)>(&Q3DTheme::gridEnabledChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(bool)>(&MyQ3DTheme::Signal_GridEnabledChanged));
}

void Q3DTheme_DisconnectGridEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(bool)>(&Q3DTheme::gridEnabledChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(bool)>(&MyQ3DTheme::Signal_GridEnabledChanged));
}

void Q3DTheme_GridEnabledChanged(void* ptr, char enabled)
{
	static_cast<Q3DTheme*>(ptr)->gridEnabledChanged(enabled != 0);
}

void Q3DTheme_ConnectGridLineColorChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::gridLineColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_GridLineColorChanged));
}

void Q3DTheme_DisconnectGridLineColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::gridLineColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_GridLineColorChanged));
}

void Q3DTheme_GridLineColorChanged(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->gridLineColorChanged(*static_cast<QColor*>(color));
}

void Q3DTheme_ConnectHighlightLightStrengthChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(float)>(&Q3DTheme::highlightLightStrengthChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(float)>(&MyQ3DTheme::Signal_HighlightLightStrengthChanged));
}

void Q3DTheme_DisconnectHighlightLightStrengthChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(float)>(&Q3DTheme::highlightLightStrengthChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(float)>(&MyQ3DTheme::Signal_HighlightLightStrengthChanged));
}

void Q3DTheme_HighlightLightStrengthChanged(void* ptr, float strength)
{
	static_cast<Q3DTheme*>(ptr)->highlightLightStrengthChanged(strength);
}

void Q3DTheme_ConnectLabelBackgroundColorChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::labelBackgroundColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_LabelBackgroundColorChanged));
}

void Q3DTheme_DisconnectLabelBackgroundColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::labelBackgroundColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_LabelBackgroundColorChanged));
}

void Q3DTheme_LabelBackgroundColorChanged(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->labelBackgroundColorChanged(*static_cast<QColor*>(color));
}

void Q3DTheme_ConnectLabelBackgroundEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(bool)>(&Q3DTheme::labelBackgroundEnabledChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(bool)>(&MyQ3DTheme::Signal_LabelBackgroundEnabledChanged));
}

void Q3DTheme_DisconnectLabelBackgroundEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(bool)>(&Q3DTheme::labelBackgroundEnabledChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(bool)>(&MyQ3DTheme::Signal_LabelBackgroundEnabledChanged));
}

void Q3DTheme_LabelBackgroundEnabledChanged(void* ptr, char enabled)
{
	static_cast<Q3DTheme*>(ptr)->labelBackgroundEnabledChanged(enabled != 0);
}

void Q3DTheme_ConnectLabelBorderEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(bool)>(&Q3DTheme::labelBorderEnabledChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(bool)>(&MyQ3DTheme::Signal_LabelBorderEnabledChanged));
}

void Q3DTheme_DisconnectLabelBorderEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(bool)>(&Q3DTheme::labelBorderEnabledChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(bool)>(&MyQ3DTheme::Signal_LabelBorderEnabledChanged));
}

void Q3DTheme_LabelBorderEnabledChanged(void* ptr, char enabled)
{
	static_cast<Q3DTheme*>(ptr)->labelBorderEnabledChanged(enabled != 0);
}

void Q3DTheme_ConnectLabelTextColorChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::labelTextColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_LabelTextColorChanged));
}

void Q3DTheme_DisconnectLabelTextColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::labelTextColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_LabelTextColorChanged));
}

void Q3DTheme_LabelTextColorChanged(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->labelTextColorChanged(*static_cast<QColor*>(color));
}

void Q3DTheme_ConnectLightColorChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::lightColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_LightColorChanged));
}

void Q3DTheme_DisconnectLightColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::lightColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_LightColorChanged));
}

void Q3DTheme_LightColorChanged(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->lightColorChanged(*static_cast<QColor*>(color));
}

void Q3DTheme_ConnectLightStrengthChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(float)>(&Q3DTheme::lightStrengthChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(float)>(&MyQ3DTheme::Signal_LightStrengthChanged));
}

void Q3DTheme_DisconnectLightStrengthChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(float)>(&Q3DTheme::lightStrengthChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(float)>(&MyQ3DTheme::Signal_LightStrengthChanged));
}

void Q3DTheme_LightStrengthChanged(void* ptr, float strength)
{
	static_cast<Q3DTheme*>(ptr)->lightStrengthChanged(strength);
}

void Q3DTheme_ConnectMultiHighlightColorChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::multiHighlightColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_MultiHighlightColorChanged));
}

void Q3DTheme_DisconnectMultiHighlightColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::multiHighlightColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_MultiHighlightColorChanged));
}

void Q3DTheme_MultiHighlightColorChanged(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->multiHighlightColorChanged(*static_cast<QColor*>(color));
}

void Q3DTheme_ConnectMultiHighlightGradientChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QLinearGradient &)>(&Q3DTheme::multiHighlightGradientChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QLinearGradient &)>(&MyQ3DTheme::Signal_MultiHighlightGradientChanged));
}

void Q3DTheme_DisconnectMultiHighlightGradientChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QLinearGradient &)>(&Q3DTheme::multiHighlightGradientChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QLinearGradient &)>(&MyQ3DTheme::Signal_MultiHighlightGradientChanged));
}

void Q3DTheme_MultiHighlightGradientChanged(void* ptr, void* gradient)
{
	static_cast<Q3DTheme*>(ptr)->multiHighlightGradientChanged(*static_cast<QLinearGradient*>(gradient));
}

void Q3DTheme_SetAmbientLightStrength(void* ptr, float strength)
{
	static_cast<Q3DTheme*>(ptr)->setAmbientLightStrength(strength);
}

void Q3DTheme_SetBackgroundColor(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void Q3DTheme_SetBackgroundEnabled(void* ptr, char enabled)
{
	static_cast<Q3DTheme*>(ptr)->setBackgroundEnabled(enabled != 0);
}

void Q3DTheme_SetBaseColors(void* ptr, void* colors)
{
	static_cast<Q3DTheme*>(ptr)->setBaseColors(*static_cast<QList<QColor>*>(colors));
}

void Q3DTheme_SetBaseGradients(void* ptr, void* gradients)
{
	static_cast<Q3DTheme*>(ptr)->setBaseGradients(*static_cast<QList<QLinearGradient>*>(gradients));
}

void Q3DTheme_SetColorStyle(void* ptr, long long style)
{
	static_cast<Q3DTheme*>(ptr)->setColorStyle(static_cast<Q3DTheme::ColorStyle>(style));
}

void Q3DTheme_SetFont(void* ptr, void* font)
{
	static_cast<Q3DTheme*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void Q3DTheme_SetGridEnabled(void* ptr, char enabled)
{
	static_cast<Q3DTheme*>(ptr)->setGridEnabled(enabled != 0);
}

void Q3DTheme_SetGridLineColor(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->setGridLineColor(*static_cast<QColor*>(color));
}

void Q3DTheme_SetHighlightLightStrength(void* ptr, float strength)
{
	static_cast<Q3DTheme*>(ptr)->setHighlightLightStrength(strength);
}

void Q3DTheme_SetLabelBackgroundColor(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->setLabelBackgroundColor(*static_cast<QColor*>(color));
}

void Q3DTheme_SetLabelBackgroundEnabled(void* ptr, char enabled)
{
	static_cast<Q3DTheme*>(ptr)->setLabelBackgroundEnabled(enabled != 0);
}

void Q3DTheme_SetLabelBorderEnabled(void* ptr, char enabled)
{
	static_cast<Q3DTheme*>(ptr)->setLabelBorderEnabled(enabled != 0);
}

void Q3DTheme_SetLabelTextColor(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->setLabelTextColor(*static_cast<QColor*>(color));
}

void Q3DTheme_SetLightColor(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->setLightColor(*static_cast<QColor*>(color));
}

void Q3DTheme_SetLightStrength(void* ptr, float strength)
{
	static_cast<Q3DTheme*>(ptr)->setLightStrength(strength);
}

void Q3DTheme_SetMultiHighlightColor(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->setMultiHighlightColor(*static_cast<QColor*>(color));
}

void Q3DTheme_SetMultiHighlightGradient(void* ptr, void* gradient)
{
	static_cast<Q3DTheme*>(ptr)->setMultiHighlightGradient(*static_cast<QLinearGradient*>(gradient));
}

void Q3DTheme_SetSingleHighlightColor(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->setSingleHighlightColor(*static_cast<QColor*>(color));
}

void Q3DTheme_SetSingleHighlightGradient(void* ptr, void* gradient)
{
	static_cast<Q3DTheme*>(ptr)->setSingleHighlightGradient(*static_cast<QLinearGradient*>(gradient));
}

void Q3DTheme_SetType(void* ptr, long long themeType)
{
	static_cast<Q3DTheme*>(ptr)->setType(static_cast<Q3DTheme::Theme>(themeType));
}

void Q3DTheme_SetWindowColor(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->setWindowColor(*static_cast<QColor*>(color));
}

void Q3DTheme_ConnectSingleHighlightColorChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::singleHighlightColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_SingleHighlightColorChanged));
}

void Q3DTheme_DisconnectSingleHighlightColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::singleHighlightColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_SingleHighlightColorChanged));
}

void Q3DTheme_SingleHighlightColorChanged(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->singleHighlightColorChanged(*static_cast<QColor*>(color));
}

void Q3DTheme_ConnectSingleHighlightGradientChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QLinearGradient &)>(&Q3DTheme::singleHighlightGradientChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QLinearGradient &)>(&MyQ3DTheme::Signal_SingleHighlightGradientChanged));
}

void Q3DTheme_DisconnectSingleHighlightGradientChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QLinearGradient &)>(&Q3DTheme::singleHighlightGradientChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QLinearGradient &)>(&MyQ3DTheme::Signal_SingleHighlightGradientChanged));
}

void Q3DTheme_SingleHighlightGradientChanged(void* ptr, void* gradient)
{
	static_cast<Q3DTheme*>(ptr)->singleHighlightGradientChanged(*static_cast<QLinearGradient*>(gradient));
}

void Q3DTheme_ConnectTypeChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(Q3DTheme::Theme)>(&Q3DTheme::typeChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(Q3DTheme::Theme)>(&MyQ3DTheme::Signal_TypeChanged));
}

void Q3DTheme_DisconnectTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(Q3DTheme::Theme)>(&Q3DTheme::typeChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(Q3DTheme::Theme)>(&MyQ3DTheme::Signal_TypeChanged));
}

void Q3DTheme_TypeChanged(void* ptr, long long themeType)
{
	static_cast<Q3DTheme*>(ptr)->typeChanged(static_cast<Q3DTheme::Theme>(themeType));
}

void Q3DTheme_ConnectWindowColorChanged(void* ptr)
{
	QObject::connect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::windowColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_WindowColorChanged));
}

void Q3DTheme_DisconnectWindowColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<Q3DTheme*>(ptr), static_cast<void (Q3DTheme::*)(const QColor &)>(&Q3DTheme::windowColorChanged), static_cast<MyQ3DTheme*>(ptr), static_cast<void (MyQ3DTheme::*)(const QColor &)>(&MyQ3DTheme::Signal_WindowColorChanged));
}

void Q3DTheme_WindowColorChanged(void* ptr, void* color)
{
	static_cast<Q3DTheme*>(ptr)->windowColorChanged(*static_cast<QColor*>(color));
}

void Q3DTheme_DestroyQ3DTheme(void* ptr)
{
	static_cast<Q3DTheme*>(ptr)->~Q3DTheme();
}

void Q3DTheme_DestroyQ3DThemeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long Q3DTheme_ColorStyle(void* ptr)
{
	return static_cast<Q3DTheme*>(ptr)->colorStyle();
}

void* Q3DTheme_BackgroundColor(void* ptr)
{
	return new QColor(static_cast<Q3DTheme*>(ptr)->backgroundColor());
}

void* Q3DTheme_GridLineColor(void* ptr)
{
	return new QColor(static_cast<Q3DTheme*>(ptr)->gridLineColor());
}

void* Q3DTheme_LabelBackgroundColor(void* ptr)
{
	return new QColor(static_cast<Q3DTheme*>(ptr)->labelBackgroundColor());
}

void* Q3DTheme_LabelTextColor(void* ptr)
{
	return new QColor(static_cast<Q3DTheme*>(ptr)->labelTextColor());
}

void* Q3DTheme_LightColor(void* ptr)
{
	return new QColor(static_cast<Q3DTheme*>(ptr)->lightColor());
}

void* Q3DTheme_MultiHighlightColor(void* ptr)
{
	return new QColor(static_cast<Q3DTheme*>(ptr)->multiHighlightColor());
}

void* Q3DTheme_SingleHighlightColor(void* ptr)
{
	return new QColor(static_cast<Q3DTheme*>(ptr)->singleHighlightColor());
}

void* Q3DTheme_WindowColor(void* ptr)
{
	return new QColor(static_cast<Q3DTheme*>(ptr)->windowColor());
}

void* Q3DTheme_Font(void* ptr)
{
	return new QFont(static_cast<Q3DTheme*>(ptr)->font());
}

void* Q3DTheme_MultiHighlightGradient(void* ptr)
{
	return new QLinearGradient(static_cast<Q3DTheme*>(ptr)->multiHighlightGradient());
}

void* Q3DTheme_SingleHighlightGradient(void* ptr)
{
	return new QLinearGradient(static_cast<Q3DTheme*>(ptr)->singleHighlightGradient());
}

struct QtDataVisualization_PackedList Q3DTheme_BaseColors(void* ptr)
{
	return ({ QList<QColor>* tmpValue = new QList<QColor>(static_cast<Q3DTheme*>(ptr)->baseColors()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDataVisualization_PackedList Q3DTheme_BaseGradients(void* ptr)
{
	return ({ QList<QLinearGradient>* tmpValue = new QList<QLinearGradient>(static_cast<Q3DTheme*>(ptr)->baseGradients()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

long long Q3DTheme_Type(void* ptr)
{
	return static_cast<Q3DTheme*>(ptr)->type();
}

char Q3DTheme_IsBackgroundEnabled(void* ptr)
{
	return static_cast<Q3DTheme*>(ptr)->isBackgroundEnabled();
}

char Q3DTheme_IsGridEnabled(void* ptr)
{
	return static_cast<Q3DTheme*>(ptr)->isGridEnabled();
}

char Q3DTheme_IsLabelBackgroundEnabled(void* ptr)
{
	return static_cast<Q3DTheme*>(ptr)->isLabelBackgroundEnabled();
}

char Q3DTheme_IsLabelBorderEnabled(void* ptr)
{
	return static_cast<Q3DTheme*>(ptr)->isLabelBorderEnabled();
}

float Q3DTheme_AmbientLightStrength(void* ptr)
{
	return static_cast<Q3DTheme*>(ptr)->ambientLightStrength();
}

float Q3DTheme_HighlightLightStrength(void* ptr)
{
	return static_cast<Q3DTheme*>(ptr)->highlightLightStrength();
}

float Q3DTheme_LightStrength(void* ptr)
{
	return static_cast<Q3DTheme*>(ptr)->lightStrength();
}

void* Q3DTheme___baseColorsChanged_colors_atList(void* ptr, int i)
{
	return new QColor(static_cast<QList<QColor>*>(ptr)->at(i));
}

void Q3DTheme___baseColorsChanged_colors_setList(void* ptr, void* i)
{
	static_cast<QList<QColor>*>(ptr)->append(*static_cast<QColor*>(i));
}

void* Q3DTheme___baseColorsChanged_colors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QColor>;
}

void* Q3DTheme___baseGradientsChanged_gradients_atList(void* ptr, int i)
{
	return new QLinearGradient(static_cast<QList<QLinearGradient>*>(ptr)->at(i));
}

void Q3DTheme___baseGradientsChanged_gradients_setList(void* ptr, void* i)
{
	static_cast<QList<QLinearGradient>*>(ptr)->append(*static_cast<QLinearGradient*>(i));
}

void* Q3DTheme___baseGradientsChanged_gradients_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QLinearGradient>;
}

void* Q3DTheme___setBaseColors_colors_atList(void* ptr, int i)
{
	return new QColor(static_cast<QList<QColor>*>(ptr)->at(i));
}

void Q3DTheme___setBaseColors_colors_setList(void* ptr, void* i)
{
	static_cast<QList<QColor>*>(ptr)->append(*static_cast<QColor*>(i));
}

void* Q3DTheme___setBaseColors_colors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QColor>;
}

void* Q3DTheme___setBaseGradients_gradients_atList(void* ptr, int i)
{
	return new QLinearGradient(static_cast<QList<QLinearGradient>*>(ptr)->at(i));
}

void Q3DTheme___setBaseGradients_gradients_setList(void* ptr, void* i)
{
	static_cast<QList<QLinearGradient>*>(ptr)->append(*static_cast<QLinearGradient*>(i));
}

void* Q3DTheme___setBaseGradients_gradients_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QLinearGradient>;
}

void* Q3DTheme___baseColors_atList(void* ptr, int i)
{
	return new QColor(static_cast<QList<QColor>*>(ptr)->at(i));
}

void Q3DTheme___baseColors_setList(void* ptr, void* i)
{
	static_cast<QList<QColor>*>(ptr)->append(*static_cast<QColor*>(i));
}

void* Q3DTheme___baseColors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QColor>;
}

void* Q3DTheme___baseGradients_atList(void* ptr, int i)
{
	return new QLinearGradient(static_cast<QList<QLinearGradient>*>(ptr)->at(i));
}

void Q3DTheme___baseGradients_setList(void* ptr, void* i)
{
	static_cast<QList<QLinearGradient>*>(ptr)->append(*static_cast<QLinearGradient*>(i));
}

void* Q3DTheme___baseGradients_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QLinearGradient>;
}

class MyQAbstract3DAxis: public QAbstract3DAxis
{
public:
	void Signal_RangeChanged(float min, float max) { callbackQAbstract3DAxis_RangeChanged(this, min, max); };
	void Signal_AutoAdjustRangeChanged(bool autoAdjust) { callbackQAbstract3DAxis_AutoAdjustRangeChanged(this, autoAdjust); };
	void Signal_LabelAutoRotationChanged(float angle) { callbackQAbstract3DAxis_LabelAutoRotationChanged(this, angle); };
	void Signal_LabelsChanged() { callbackQAbstract3DAxis_LabelsChanged(this); };
	void Signal_MaxChanged(float value) { callbackQAbstract3DAxis_MaxChanged(this, value); };
	void Signal_MinChanged(float value) { callbackQAbstract3DAxis_MinChanged(this, value); };
	void Signal_OrientationChanged(QAbstract3DAxis::AxisOrientation orientation) { callbackQAbstract3DAxis_OrientationChanged(this, orientation); };
	void Signal_TitleChanged(const QString & newTitle) { QByteArray tee1e04 = newTitle.toUtf8(); QtDataVisualization_PackedString newTitlePacked = { const_cast<char*>(tee1e04.prepend("WHITESPACE").constData()+10), tee1e04.size()-10 };callbackQAbstract3DAxis_TitleChanged(this, newTitlePacked); };
	void Signal_TitleFixedChanged(bool fixed) { callbackQAbstract3DAxis_TitleFixedChanged(this, fixed); };
	void Signal_TitleVisibilityChanged(bool visible) { callbackQAbstract3DAxis_TitleVisibilityChanged(this, visible); };
	 ~MyQAbstract3DAxis() { callbackQAbstract3DAxis_DestroyQAbstract3DAxis(this); };
};

void QAbstract3DAxis_ConnectRangeChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(float, float)>(&QAbstract3DAxis::rangeChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(float, float)>(&MyQAbstract3DAxis::Signal_RangeChanged));
}

void QAbstract3DAxis_DisconnectRangeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(float, float)>(&QAbstract3DAxis::rangeChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(float, float)>(&MyQAbstract3DAxis::Signal_RangeChanged));
}

void QAbstract3DAxis_RangeChanged(void* ptr, float min, float max)
{
	static_cast<QAbstract3DAxis*>(ptr)->rangeChanged(min, max);
}

void QAbstract3DAxis_SetTitleFixed(void* ptr, char fixed)
{
	static_cast<QAbstract3DAxis*>(ptr)->setTitleFixed(fixed != 0);
}

void QAbstract3DAxis_ConnectAutoAdjustRangeChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(bool)>(&QAbstract3DAxis::autoAdjustRangeChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(bool)>(&MyQAbstract3DAxis::Signal_AutoAdjustRangeChanged));
}

void QAbstract3DAxis_DisconnectAutoAdjustRangeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(bool)>(&QAbstract3DAxis::autoAdjustRangeChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(bool)>(&MyQAbstract3DAxis::Signal_AutoAdjustRangeChanged));
}

void QAbstract3DAxis_AutoAdjustRangeChanged(void* ptr, char autoAdjust)
{
	static_cast<QAbstract3DAxis*>(ptr)->autoAdjustRangeChanged(autoAdjust != 0);
}

void QAbstract3DAxis_ConnectLabelAutoRotationChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(float)>(&QAbstract3DAxis::labelAutoRotationChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(float)>(&MyQAbstract3DAxis::Signal_LabelAutoRotationChanged));
}

void QAbstract3DAxis_DisconnectLabelAutoRotationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(float)>(&QAbstract3DAxis::labelAutoRotationChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(float)>(&MyQAbstract3DAxis::Signal_LabelAutoRotationChanged));
}

void QAbstract3DAxis_LabelAutoRotationChanged(void* ptr, float angle)
{
	static_cast<QAbstract3DAxis*>(ptr)->labelAutoRotationChanged(angle);
}

void QAbstract3DAxis_ConnectLabelsChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)()>(&QAbstract3DAxis::labelsChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)()>(&MyQAbstract3DAxis::Signal_LabelsChanged));
}

void QAbstract3DAxis_DisconnectLabelsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)()>(&QAbstract3DAxis::labelsChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)()>(&MyQAbstract3DAxis::Signal_LabelsChanged));
}

void QAbstract3DAxis_LabelsChanged(void* ptr)
{
	static_cast<QAbstract3DAxis*>(ptr)->labelsChanged();
}

void QAbstract3DAxis_ConnectMaxChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(float)>(&QAbstract3DAxis::maxChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(float)>(&MyQAbstract3DAxis::Signal_MaxChanged));
}

void QAbstract3DAxis_DisconnectMaxChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(float)>(&QAbstract3DAxis::maxChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(float)>(&MyQAbstract3DAxis::Signal_MaxChanged));
}

void QAbstract3DAxis_MaxChanged(void* ptr, float value)
{
	static_cast<QAbstract3DAxis*>(ptr)->maxChanged(value);
}

void QAbstract3DAxis_ConnectMinChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(float)>(&QAbstract3DAxis::minChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(float)>(&MyQAbstract3DAxis::Signal_MinChanged));
}

void QAbstract3DAxis_DisconnectMinChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(float)>(&QAbstract3DAxis::minChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(float)>(&MyQAbstract3DAxis::Signal_MinChanged));
}

void QAbstract3DAxis_MinChanged(void* ptr, float value)
{
	static_cast<QAbstract3DAxis*>(ptr)->minChanged(value);
}

void QAbstract3DAxis_ConnectOrientationChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(QAbstract3DAxis::AxisOrientation)>(&QAbstract3DAxis::orientationChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(QAbstract3DAxis::AxisOrientation)>(&MyQAbstract3DAxis::Signal_OrientationChanged));
}

void QAbstract3DAxis_DisconnectOrientationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(QAbstract3DAxis::AxisOrientation)>(&QAbstract3DAxis::orientationChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(QAbstract3DAxis::AxisOrientation)>(&MyQAbstract3DAxis::Signal_OrientationChanged));
}

void QAbstract3DAxis_OrientationChanged(void* ptr, long long orientation)
{
	static_cast<QAbstract3DAxis*>(ptr)->orientationChanged(static_cast<QAbstract3DAxis::AxisOrientation>(orientation));
}

void QAbstract3DAxis_SetAutoAdjustRange(void* ptr, char autoAdjust)
{
	static_cast<QAbstract3DAxis*>(ptr)->setAutoAdjustRange(autoAdjust != 0);
}

void QAbstract3DAxis_SetLabelAutoRotation(void* ptr, float angle)
{
	static_cast<QAbstract3DAxis*>(ptr)->setLabelAutoRotation(angle);
}

void QAbstract3DAxis_SetLabels(void* ptr, struct QtDataVisualization_PackedString labels)
{
	static_cast<QAbstract3DAxis*>(ptr)->setLabels(QString::fromUtf8(labels.data, labels.len).split("|", QString::SkipEmptyParts));
}

void QAbstract3DAxis_SetMax(void* ptr, float max)
{
	static_cast<QAbstract3DAxis*>(ptr)->setMax(max);
}

void QAbstract3DAxis_SetMin(void* ptr, float min)
{
	static_cast<QAbstract3DAxis*>(ptr)->setMin(min);
}

void QAbstract3DAxis_SetRange(void* ptr, float min, float max)
{
	static_cast<QAbstract3DAxis*>(ptr)->setRange(min, max);
}

void QAbstract3DAxis_SetTitle(void* ptr, struct QtDataVisualization_PackedString title)
{
	static_cast<QAbstract3DAxis*>(ptr)->setTitle(QString::fromUtf8(title.data, title.len));
}

void QAbstract3DAxis_SetTitleVisible(void* ptr, char visible)
{
	static_cast<QAbstract3DAxis*>(ptr)->setTitleVisible(visible != 0);
}

void QAbstract3DAxis_ConnectTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(const QString &)>(&QAbstract3DAxis::titleChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(const QString &)>(&MyQAbstract3DAxis::Signal_TitleChanged));
}

void QAbstract3DAxis_DisconnectTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(const QString &)>(&QAbstract3DAxis::titleChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(const QString &)>(&MyQAbstract3DAxis::Signal_TitleChanged));
}

void QAbstract3DAxis_TitleChanged(void* ptr, struct QtDataVisualization_PackedString newTitle)
{
	static_cast<QAbstract3DAxis*>(ptr)->titleChanged(QString::fromUtf8(newTitle.data, newTitle.len));
}

void QAbstract3DAxis_ConnectTitleFixedChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(bool)>(&QAbstract3DAxis::titleFixedChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(bool)>(&MyQAbstract3DAxis::Signal_TitleFixedChanged));
}

void QAbstract3DAxis_DisconnectTitleFixedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(bool)>(&QAbstract3DAxis::titleFixedChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(bool)>(&MyQAbstract3DAxis::Signal_TitleFixedChanged));
}

void QAbstract3DAxis_TitleFixedChanged(void* ptr, char fixed)
{
	static_cast<QAbstract3DAxis*>(ptr)->titleFixedChanged(fixed != 0);
}

void QAbstract3DAxis_ConnectTitleVisibilityChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(bool)>(&QAbstract3DAxis::titleVisibilityChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(bool)>(&MyQAbstract3DAxis::Signal_TitleVisibilityChanged));
}

void QAbstract3DAxis_DisconnectTitleVisibilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DAxis*>(ptr), static_cast<void (QAbstract3DAxis::*)(bool)>(&QAbstract3DAxis::titleVisibilityChanged), static_cast<MyQAbstract3DAxis*>(ptr), static_cast<void (MyQAbstract3DAxis::*)(bool)>(&MyQAbstract3DAxis::Signal_TitleVisibilityChanged));
}

void QAbstract3DAxis_TitleVisibilityChanged(void* ptr, char visible)
{
	static_cast<QAbstract3DAxis*>(ptr)->titleVisibilityChanged(visible != 0);
}

void QAbstract3DAxis_DestroyQAbstract3DAxis(void* ptr)
{
	static_cast<QAbstract3DAxis*>(ptr)->~QAbstract3DAxis();
}

void QAbstract3DAxis_DestroyQAbstract3DAxisDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QAbstract3DAxis_Orientation(void* ptr)
{
	return static_cast<QAbstract3DAxis*>(ptr)->orientation();
}

long long QAbstract3DAxis_Type(void* ptr)
{
	return static_cast<QAbstract3DAxis*>(ptr)->type();
}

struct QtDataVisualization_PackedString QAbstract3DAxis_Title(void* ptr)
{
	return ({ QByteArray tcaf968 = static_cast<QAbstract3DAxis*>(ptr)->title().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tcaf968.prepend("WHITESPACE").constData()+10), tcaf968.size()-10 }; });
}

struct QtDataVisualization_PackedString QAbstract3DAxis_Labels(void* ptr)
{
	return ({ QByteArray t7dab2f = static_cast<QAbstract3DAxis*>(ptr)->labels().join("|").toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t7dab2f.prepend("WHITESPACE").constData()+10), t7dab2f.size()-10 }; });
}

char QAbstract3DAxis_IsAutoAdjustRange(void* ptr)
{
	return static_cast<QAbstract3DAxis*>(ptr)->isAutoAdjustRange();
}

char QAbstract3DAxis_IsTitleFixed(void* ptr)
{
	return static_cast<QAbstract3DAxis*>(ptr)->isTitleFixed();
}

char QAbstract3DAxis_IsTitleVisible(void* ptr)
{
	return static_cast<QAbstract3DAxis*>(ptr)->isTitleVisible();
}

float QAbstract3DAxis_LabelAutoRotation(void* ptr)
{
	return static_cast<QAbstract3DAxis*>(ptr)->labelAutoRotation();
}

float QAbstract3DAxis_Max(void* ptr)
{
	return static_cast<QAbstract3DAxis*>(ptr)->max();
}

float QAbstract3DAxis_Min(void* ptr)
{
	return static_cast<QAbstract3DAxis*>(ptr)->min();
}

class MyQAbstract3DGraph: public QAbstract3DGraph
{
public:
	void Signal_ActiveInputHandlerChanged(QAbstract3DInputHandler * inputHandler) { callbackQAbstract3DGraph_ActiveInputHandlerChanged(this, inputHandler); };
	void Signal_ActiveThemeChanged(Q3DTheme * theme) { callbackQAbstract3DGraph_ActiveThemeChanged(this, theme); };
	void Signal_AspectRatioChanged(qreal ratio) { callbackQAbstract3DGraph_AspectRatioChanged(this, ratio); };
	void Signal_CurrentFpsChanged(qreal fps) { callbackQAbstract3DGraph_CurrentFpsChanged(this, fps); };
	void Signal_HorizontalAspectRatioChanged(qreal ratio) { callbackQAbstract3DGraph_HorizontalAspectRatioChanged(this, ratio); };
	void Signal_LocaleChanged(const QLocale & locale) { callbackQAbstract3DGraph_LocaleChanged(this, const_cast<QLocale*>(&locale)); };
	void Signal_MarginChanged(qreal margin) { callbackQAbstract3DGraph_MarginChanged(this, margin); };
	void Signal_MeasureFpsChanged(bool enabled) { callbackQAbstract3DGraph_MeasureFpsChanged(this, enabled); };
	void Signal_OptimizationHintsChanged(QAbstract3DGraph::OptimizationHints hints) { callbackQAbstract3DGraph_OptimizationHintsChanged(this, hints); };
	void Signal_OrthoProjectionChanged(bool enabled) { callbackQAbstract3DGraph_OrthoProjectionChanged(this, enabled); };
	void Signal_PolarChanged(bool enabled) { callbackQAbstract3DGraph_PolarChanged(this, enabled); };
	void Signal_QueriedGraphPositionChanged(const QVector3D & data) { callbackQAbstract3DGraph_QueriedGraphPositionChanged(this, const_cast<QVector3D*>(&data)); };
	void Signal_RadialLabelOffsetChanged(float offset) { callbackQAbstract3DGraph_RadialLabelOffsetChanged(this, offset); };
	void Signal_ReflectionChanged(bool enabled) { callbackQAbstract3DGraph_ReflectionChanged(this, enabled); };
	void Signal_ReflectivityChanged(qreal reflectivity) { callbackQAbstract3DGraph_ReflectivityChanged(this, reflectivity); };
	void Signal_SelectedElementChanged(QAbstract3DGraph::ElementType ty) { callbackQAbstract3DGraph_SelectedElementChanged(this, ty); };
	void Signal_SelectionModeChanged(QAbstract3DGraph::SelectionFlags mode) { callbackQAbstract3DGraph_SelectionModeChanged(this, mode); };
	void Signal_ShadowQualityChanged(QAbstract3DGraph::ShadowQuality quality) { callbackQAbstract3DGraph_ShadowQualityChanged(this, quality); };
	 ~MyQAbstract3DGraph() { callbackQAbstract3DGraph_DestroyQAbstract3DGraph(this); };
	bool shadowsSupported() const { return callbackQAbstract3DGraph_ShadowsSupported(const_cast<void*>(static_cast<const void*>(this))) != 0; };
};

void* QAbstract3DGraph_RenderToImage(void* ptr, int msaaSamples, void* imageSize)
{
	return new QImage(static_cast<QAbstract3DGraph*>(ptr)->renderToImage(msaaSamples, *static_cast<QSize*>(imageSize)));
}

int QAbstract3DGraph_AddCustomItem(void* ptr, void* item)
{
	return static_cast<QAbstract3DGraph*>(ptr)->addCustomItem(static_cast<QCustom3DItem*>(item));
}

void QAbstract3DGraph_ConnectActiveInputHandlerChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DInputHandler *)>(&QAbstract3DGraph::activeInputHandlerChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DInputHandler *)>(&MyQAbstract3DGraph::Signal_ActiveInputHandlerChanged));
}

void QAbstract3DGraph_DisconnectActiveInputHandlerChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DInputHandler *)>(&QAbstract3DGraph::activeInputHandlerChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DInputHandler *)>(&MyQAbstract3DGraph::Signal_ActiveInputHandlerChanged));
}

void QAbstract3DGraph_ActiveInputHandlerChanged(void* ptr, void* inputHandler)
{
	static_cast<QAbstract3DGraph*>(ptr)->activeInputHandlerChanged(static_cast<QAbstract3DInputHandler*>(inputHandler));
}

void QAbstract3DGraph_ConnectActiveThemeChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(Q3DTheme *)>(&QAbstract3DGraph::activeThemeChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(Q3DTheme *)>(&MyQAbstract3DGraph::Signal_ActiveThemeChanged));
}

void QAbstract3DGraph_DisconnectActiveThemeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(Q3DTheme *)>(&QAbstract3DGraph::activeThemeChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(Q3DTheme *)>(&MyQAbstract3DGraph::Signal_ActiveThemeChanged));
}

void QAbstract3DGraph_ActiveThemeChanged(void* ptr, void* theme)
{
	static_cast<QAbstract3DGraph*>(ptr)->activeThemeChanged(static_cast<Q3DTheme*>(theme));
}

void QAbstract3DGraph_AddInputHandler(void* ptr, void* inputHandler)
{
	static_cast<QAbstract3DGraph*>(ptr)->addInputHandler(static_cast<QAbstract3DInputHandler*>(inputHandler));
}

void QAbstract3DGraph_AddTheme(void* ptr, void* theme)
{
	static_cast<QAbstract3DGraph*>(ptr)->addTheme(static_cast<Q3DTheme*>(theme));
}

void QAbstract3DGraph_ConnectAspectRatioChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::aspectRatioChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_AspectRatioChanged));
}

void QAbstract3DGraph_DisconnectAspectRatioChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::aspectRatioChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_AspectRatioChanged));
}

void QAbstract3DGraph_AspectRatioChanged(void* ptr, double ratio)
{
	static_cast<QAbstract3DGraph*>(ptr)->aspectRatioChanged(ratio);
}

void QAbstract3DGraph_ClearSelection(void* ptr)
{
	static_cast<QAbstract3DGraph*>(ptr)->clearSelection();
}

void QAbstract3DGraph_ConnectCurrentFpsChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::currentFpsChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_CurrentFpsChanged));
}

void QAbstract3DGraph_DisconnectCurrentFpsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::currentFpsChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_CurrentFpsChanged));
}

void QAbstract3DGraph_CurrentFpsChanged(void* ptr, double fps)
{
	static_cast<QAbstract3DGraph*>(ptr)->currentFpsChanged(fps);
}

void QAbstract3DGraph_ConnectHorizontalAspectRatioChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::horizontalAspectRatioChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_HorizontalAspectRatioChanged));
}

void QAbstract3DGraph_DisconnectHorizontalAspectRatioChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::horizontalAspectRatioChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_HorizontalAspectRatioChanged));
}

void QAbstract3DGraph_HorizontalAspectRatioChanged(void* ptr, double ratio)
{
	static_cast<QAbstract3DGraph*>(ptr)->horizontalAspectRatioChanged(ratio);
}

void QAbstract3DGraph_ConnectLocaleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(const QLocale &)>(&QAbstract3DGraph::localeChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(const QLocale &)>(&MyQAbstract3DGraph::Signal_LocaleChanged));
}

void QAbstract3DGraph_DisconnectLocaleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(const QLocale &)>(&QAbstract3DGraph::localeChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(const QLocale &)>(&MyQAbstract3DGraph::Signal_LocaleChanged));
}

void QAbstract3DGraph_LocaleChanged(void* ptr, void* locale)
{
	static_cast<QAbstract3DGraph*>(ptr)->localeChanged(*static_cast<QLocale*>(locale));
}

void QAbstract3DGraph_ConnectMarginChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::marginChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_MarginChanged));
}

void QAbstract3DGraph_DisconnectMarginChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::marginChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_MarginChanged));
}

void QAbstract3DGraph_MarginChanged(void* ptr, double margin)
{
	static_cast<QAbstract3DGraph*>(ptr)->marginChanged(margin);
}

void QAbstract3DGraph_ConnectMeasureFpsChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(bool)>(&QAbstract3DGraph::measureFpsChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(bool)>(&MyQAbstract3DGraph::Signal_MeasureFpsChanged));
}

void QAbstract3DGraph_DisconnectMeasureFpsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(bool)>(&QAbstract3DGraph::measureFpsChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(bool)>(&MyQAbstract3DGraph::Signal_MeasureFpsChanged));
}

void QAbstract3DGraph_MeasureFpsChanged(void* ptr, char enabled)
{
	static_cast<QAbstract3DGraph*>(ptr)->measureFpsChanged(enabled != 0);
}

void QAbstract3DGraph_ConnectOptimizationHintsChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DGraph::OptimizationHints)>(&QAbstract3DGraph::optimizationHintsChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DGraph::OptimizationHints)>(&MyQAbstract3DGraph::Signal_OptimizationHintsChanged));
}

void QAbstract3DGraph_DisconnectOptimizationHintsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DGraph::OptimizationHints)>(&QAbstract3DGraph::optimizationHintsChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DGraph::OptimizationHints)>(&MyQAbstract3DGraph::Signal_OptimizationHintsChanged));
}

void QAbstract3DGraph_OptimizationHintsChanged(void* ptr, long long hints)
{
	static_cast<QAbstract3DGraph*>(ptr)->optimizationHintsChanged(static_cast<QAbstract3DGraph::OptimizationHint>(hints));
}

void QAbstract3DGraph_ConnectOrthoProjectionChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(bool)>(&QAbstract3DGraph::orthoProjectionChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(bool)>(&MyQAbstract3DGraph::Signal_OrthoProjectionChanged));
}

void QAbstract3DGraph_DisconnectOrthoProjectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(bool)>(&QAbstract3DGraph::orthoProjectionChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(bool)>(&MyQAbstract3DGraph::Signal_OrthoProjectionChanged));
}

void QAbstract3DGraph_OrthoProjectionChanged(void* ptr, char enabled)
{
	static_cast<QAbstract3DGraph*>(ptr)->orthoProjectionChanged(enabled != 0);
}

void QAbstract3DGraph_ConnectPolarChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(bool)>(&QAbstract3DGraph::polarChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(bool)>(&MyQAbstract3DGraph::Signal_PolarChanged));
}

void QAbstract3DGraph_DisconnectPolarChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(bool)>(&QAbstract3DGraph::polarChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(bool)>(&MyQAbstract3DGraph::Signal_PolarChanged));
}

void QAbstract3DGraph_PolarChanged(void* ptr, char enabled)
{
	static_cast<QAbstract3DGraph*>(ptr)->polarChanged(enabled != 0);
}

void QAbstract3DGraph_ConnectQueriedGraphPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(const QVector3D &)>(&QAbstract3DGraph::queriedGraphPositionChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(const QVector3D &)>(&MyQAbstract3DGraph::Signal_QueriedGraphPositionChanged));
}

void QAbstract3DGraph_DisconnectQueriedGraphPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(const QVector3D &)>(&QAbstract3DGraph::queriedGraphPositionChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(const QVector3D &)>(&MyQAbstract3DGraph::Signal_QueriedGraphPositionChanged));
}

void QAbstract3DGraph_QueriedGraphPositionChanged(void* ptr, void* data)
{
	static_cast<QAbstract3DGraph*>(ptr)->queriedGraphPositionChanged(*static_cast<QVector3D*>(data));
}

void QAbstract3DGraph_ConnectRadialLabelOffsetChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(float)>(&QAbstract3DGraph::radialLabelOffsetChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(float)>(&MyQAbstract3DGraph::Signal_RadialLabelOffsetChanged));
}

void QAbstract3DGraph_DisconnectRadialLabelOffsetChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(float)>(&QAbstract3DGraph::radialLabelOffsetChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(float)>(&MyQAbstract3DGraph::Signal_RadialLabelOffsetChanged));
}

void QAbstract3DGraph_RadialLabelOffsetChanged(void* ptr, float offset)
{
	static_cast<QAbstract3DGraph*>(ptr)->radialLabelOffsetChanged(offset);
}

void QAbstract3DGraph_ConnectReflectionChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(bool)>(&QAbstract3DGraph::reflectionChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(bool)>(&MyQAbstract3DGraph::Signal_ReflectionChanged));
}

void QAbstract3DGraph_DisconnectReflectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(bool)>(&QAbstract3DGraph::reflectionChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(bool)>(&MyQAbstract3DGraph::Signal_ReflectionChanged));
}

void QAbstract3DGraph_ReflectionChanged(void* ptr, char enabled)
{
	static_cast<QAbstract3DGraph*>(ptr)->reflectionChanged(enabled != 0);
}

void QAbstract3DGraph_ConnectReflectivityChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::reflectivityChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_ReflectivityChanged));
}

void QAbstract3DGraph_DisconnectReflectivityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(qreal)>(&QAbstract3DGraph::reflectivityChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(qreal)>(&MyQAbstract3DGraph::Signal_ReflectivityChanged));
}

void QAbstract3DGraph_ReflectivityChanged(void* ptr, double reflectivity)
{
	static_cast<QAbstract3DGraph*>(ptr)->reflectivityChanged(reflectivity);
}

void QAbstract3DGraph_ReleaseCustomItem(void* ptr, void* item)
{
	static_cast<QAbstract3DGraph*>(ptr)->releaseCustomItem(static_cast<QCustom3DItem*>(item));
}

void QAbstract3DGraph_ReleaseInputHandler(void* ptr, void* inputHandler)
{
	static_cast<QAbstract3DGraph*>(ptr)->releaseInputHandler(static_cast<QAbstract3DInputHandler*>(inputHandler));
}

void QAbstract3DGraph_ReleaseTheme(void* ptr, void* theme)
{
	static_cast<QAbstract3DGraph*>(ptr)->releaseTheme(static_cast<Q3DTheme*>(theme));
}

void QAbstract3DGraph_RemoveCustomItem(void* ptr, void* item)
{
	static_cast<QAbstract3DGraph*>(ptr)->removeCustomItem(static_cast<QCustom3DItem*>(item));
}

void QAbstract3DGraph_RemoveCustomItemAt(void* ptr, void* position)
{
	static_cast<QAbstract3DGraph*>(ptr)->removeCustomItemAt(*static_cast<QVector3D*>(position));
}

void QAbstract3DGraph_RemoveCustomItems(void* ptr)
{
	static_cast<QAbstract3DGraph*>(ptr)->removeCustomItems();
}

void QAbstract3DGraph_ConnectSelectedElementChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DGraph::ElementType)>(&QAbstract3DGraph::selectedElementChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DGraph::ElementType)>(&MyQAbstract3DGraph::Signal_SelectedElementChanged));
}

void QAbstract3DGraph_DisconnectSelectedElementChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DGraph::ElementType)>(&QAbstract3DGraph::selectedElementChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DGraph::ElementType)>(&MyQAbstract3DGraph::Signal_SelectedElementChanged));
}

void QAbstract3DGraph_SelectedElementChanged(void* ptr, long long ty)
{
	static_cast<QAbstract3DGraph*>(ptr)->selectedElementChanged(static_cast<QAbstract3DGraph::ElementType>(ty));
}

void QAbstract3DGraph_ConnectSelectionModeChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DGraph::SelectionFlags)>(&QAbstract3DGraph::selectionModeChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DGraph::SelectionFlags)>(&MyQAbstract3DGraph::Signal_SelectionModeChanged));
}

void QAbstract3DGraph_DisconnectSelectionModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DGraph::SelectionFlags)>(&QAbstract3DGraph::selectionModeChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DGraph::SelectionFlags)>(&MyQAbstract3DGraph::Signal_SelectionModeChanged));
}

void QAbstract3DGraph_SelectionModeChanged(void* ptr, long long mode)
{
	static_cast<QAbstract3DGraph*>(ptr)->selectionModeChanged(static_cast<QAbstract3DGraph::SelectionFlag>(mode));
}

void QAbstract3DGraph_SetActiveInputHandler(void* ptr, void* inputHandler)
{
	static_cast<QAbstract3DGraph*>(ptr)->setActiveInputHandler(static_cast<QAbstract3DInputHandler*>(inputHandler));
}

void QAbstract3DGraph_SetActiveTheme(void* ptr, void* theme)
{
	static_cast<QAbstract3DGraph*>(ptr)->setActiveTheme(static_cast<Q3DTheme*>(theme));
}

void QAbstract3DGraph_SetAspectRatio(void* ptr, double ratio)
{
	static_cast<QAbstract3DGraph*>(ptr)->setAspectRatio(ratio);
}

void QAbstract3DGraph_SetHorizontalAspectRatio(void* ptr, double ratio)
{
	static_cast<QAbstract3DGraph*>(ptr)->setHorizontalAspectRatio(ratio);
}

void QAbstract3DGraph_SetLocale(void* ptr, void* locale)
{
	static_cast<QAbstract3DGraph*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QAbstract3DGraph_SetMargin(void* ptr, double margin)
{
	static_cast<QAbstract3DGraph*>(ptr)->setMargin(margin);
}

void QAbstract3DGraph_SetMeasureFps(void* ptr, char enable)
{
	static_cast<QAbstract3DGraph*>(ptr)->setMeasureFps(enable != 0);
}

void QAbstract3DGraph_SetOptimizationHints(void* ptr, long long hints)
{
	static_cast<QAbstract3DGraph*>(ptr)->setOptimizationHints(static_cast<QAbstract3DGraph::OptimizationHint>(hints));
}

void QAbstract3DGraph_SetOrthoProjection(void* ptr, char enable)
{
	static_cast<QAbstract3DGraph*>(ptr)->setOrthoProjection(enable != 0);
}

void QAbstract3DGraph_SetPolar(void* ptr, char enable)
{
	static_cast<QAbstract3DGraph*>(ptr)->setPolar(enable != 0);
}

void QAbstract3DGraph_SetRadialLabelOffset(void* ptr, float offset)
{
	static_cast<QAbstract3DGraph*>(ptr)->setRadialLabelOffset(offset);
}

void QAbstract3DGraph_SetReflection(void* ptr, char enable)
{
	static_cast<QAbstract3DGraph*>(ptr)->setReflection(enable != 0);
}

void QAbstract3DGraph_SetReflectivity(void* ptr, double reflectivity)
{
	static_cast<QAbstract3DGraph*>(ptr)->setReflectivity(reflectivity);
}

void QAbstract3DGraph_SetSelectionMode(void* ptr, long long mode)
{
	static_cast<QAbstract3DGraph*>(ptr)->setSelectionMode(static_cast<QAbstract3DGraph::SelectionFlag>(mode));
}

void QAbstract3DGraph_SetShadowQuality(void* ptr, long long quality)
{
	static_cast<QAbstract3DGraph*>(ptr)->setShadowQuality(static_cast<QAbstract3DGraph::ShadowQuality>(quality));
}

void QAbstract3DGraph_ConnectShadowQualityChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DGraph::ShadowQuality)>(&QAbstract3DGraph::shadowQualityChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DGraph::ShadowQuality)>(&MyQAbstract3DGraph::Signal_ShadowQualityChanged));
}

void QAbstract3DGraph_DisconnectShadowQualityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DGraph*>(ptr), static_cast<void (QAbstract3DGraph::*)(QAbstract3DGraph::ShadowQuality)>(&QAbstract3DGraph::shadowQualityChanged), static_cast<MyQAbstract3DGraph*>(ptr), static_cast<void (MyQAbstract3DGraph::*)(QAbstract3DGraph::ShadowQuality)>(&MyQAbstract3DGraph::Signal_ShadowQualityChanged));
}

void QAbstract3DGraph_ShadowQualityChanged(void* ptr, long long quality)
{
	static_cast<QAbstract3DGraph*>(ptr)->shadowQualityChanged(static_cast<QAbstract3DGraph::ShadowQuality>(quality));
}

void QAbstract3DGraph_DestroyQAbstract3DGraph(void* ptr)
{
	static_cast<QAbstract3DGraph*>(ptr)->~QAbstract3DGraph();
}

void QAbstract3DGraph_DestroyQAbstract3DGraphDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QAbstract3DGraph_SelectedElement(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->selectedElement();
}

long long QAbstract3DGraph_OptimizationHints(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->optimizationHints();
}

void* QAbstract3DGraph_Scene(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->scene();
}

void* QAbstract3DGraph_ActiveTheme(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->activeTheme();
}

void* QAbstract3DGraph_SelectedAxis(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->selectedAxis();
}

void* QAbstract3DGraph_ActiveInputHandler(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->activeInputHandler();
}

void* QAbstract3DGraph_SelectedCustomItem(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->selectedCustomItem();
}

struct QtDataVisualization_PackedList QAbstract3DGraph_Themes(void* ptr)
{
	return ({ QList<Q3DTheme *>* tmpValue = new QList<Q3DTheme *>(static_cast<QAbstract3DGraph*>(ptr)->themes()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDataVisualization_PackedList QAbstract3DGraph_InputHandlers(void* ptr)
{
	return ({ QList<QAbstract3DInputHandler *>* tmpValue = new QList<QAbstract3DInputHandler *>(static_cast<QAbstract3DGraph*>(ptr)->inputHandlers()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDataVisualization_PackedList QAbstract3DGraph_CustomItems(void* ptr)
{
	return ({ QList<QCustom3DItem *>* tmpValue = new QList<QCustom3DItem *>(static_cast<QAbstract3DGraph*>(ptr)->customItems()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

void* QAbstract3DGraph_Locale(void* ptr)
{
	return new QLocale(static_cast<QAbstract3DGraph*>(ptr)->locale());
}

void* QAbstract3DGraph_QueriedGraphPosition(void* ptr)
{
	return new QVector3D(static_cast<QAbstract3DGraph*>(ptr)->queriedGraphPosition());
}

long long QAbstract3DGraph_SelectionMode(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->selectionMode();
}

long long QAbstract3DGraph_ShadowQuality(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->shadowQuality();
}

char QAbstract3DGraph_HasContext(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->hasContext();
}

char QAbstract3DGraph_IsOrthoProjection(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->isOrthoProjection();
}

char QAbstract3DGraph_IsPolar(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->isPolar();
}

char QAbstract3DGraph_IsReflection(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->isReflection();
}

char QAbstract3DGraph_MeasureFps(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->measureFps();
}

char QAbstract3DGraph_ShadowsSupported(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->shadowsSupported();
}

char QAbstract3DGraph_ShadowsSupportedDefault(void* ptr)
{
		return static_cast<QAbstract3DGraph*>(ptr)->QAbstract3DGraph::shadowsSupported();
}

float QAbstract3DGraph_RadialLabelOffset(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->radialLabelOffset();
}

int QAbstract3DGraph_SelectedCustomItemIndex(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->selectedCustomItemIndex();
}

int QAbstract3DGraph_SelectedLabelIndex(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->selectedLabelIndex();
}

double QAbstract3DGraph_AspectRatio(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->aspectRatio();
}

double QAbstract3DGraph_CurrentFps(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->currentFps();
}

double QAbstract3DGraph_HorizontalAspectRatio(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->horizontalAspectRatio();
}

double QAbstract3DGraph_Margin(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->margin();
}

double QAbstract3DGraph_Reflectivity(void* ptr)
{
	return static_cast<QAbstract3DGraph*>(ptr)->reflectivity();
}

void* QAbstract3DGraph___themes_atList(void* ptr, int i)
{
	return const_cast<Q3DTheme*>(static_cast<QList<Q3DTheme *>*>(ptr)->at(i));
}

void QAbstract3DGraph___themes_setList(void* ptr, void* i)
{
	static_cast<QList<Q3DTheme *>*>(ptr)->append(static_cast<Q3DTheme*>(i));
}

void* QAbstract3DGraph___themes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Q3DTheme *>;
}

void* QAbstract3DGraph___inputHandlers_atList(void* ptr, int i)
{
	return const_cast<QAbstract3DInputHandler*>(static_cast<QList<QAbstract3DInputHandler *>*>(ptr)->at(i));
}

void QAbstract3DGraph___inputHandlers_setList(void* ptr, void* i)
{
	static_cast<QList<QAbstract3DInputHandler *>*>(ptr)->append(static_cast<QAbstract3DInputHandler*>(i));
}

void* QAbstract3DGraph___inputHandlers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAbstract3DInputHandler *>;
}

void* QAbstract3DGraph___customItems_atList(void* ptr, int i)
{
	return const_cast<QCustom3DItem*>(static_cast<QList<QCustom3DItem *>*>(ptr)->at(i));
}

void QAbstract3DGraph___customItems_setList(void* ptr, void* i)
{
	static_cast<QList<QCustom3DItem *>*>(ptr)->append(static_cast<QCustom3DItem*>(i));
}

void* QAbstract3DGraph___customItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QCustom3DItem *>;
}

class MyQAbstract3DInputHandler: public QAbstract3DInputHandler
{
public:
	MyQAbstract3DInputHandler(QObject *parent = Q_NULLPTR) : QAbstract3DInputHandler(parent) {};
	void Signal_InputViewChanged(QAbstract3DInputHandler::InputView view) { callbackQAbstract3DInputHandler_InputViewChanged(this, view); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQAbstract3DInputHandler_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event, const QPoint & mousePos) { callbackQAbstract3DInputHandler_MouseMoveEvent(this, event, const_cast<QPoint*>(&mousePos)); };
	void mousePressEvent(QMouseEvent * event, const QPoint & mousePos) { callbackQAbstract3DInputHandler_MousePressEvent(this, event, const_cast<QPoint*>(&mousePos)); };
	void mouseReleaseEvent(QMouseEvent * event, const QPoint & mousePos) { callbackQAbstract3DInputHandler_MouseReleaseEvent(this, event, const_cast<QPoint*>(&mousePos)); };
	void Signal_PositionChanged(const QPoint & position) { callbackQAbstract3DInputHandler_PositionChanged(this, const_cast<QPoint*>(&position)); };
	void Signal_SceneChanged(Q3DScene * scene) { callbackQAbstract3DInputHandler_SceneChanged(this, scene); };
	void touchEvent(QTouchEvent * event) { callbackQAbstract3DInputHandler_TouchEvent(this, event); };
	void wheelEvent(QWheelEvent * event) { callbackQAbstract3DInputHandler_WheelEvent(this, event); };
	 ~MyQAbstract3DInputHandler() { callbackQAbstract3DInputHandler_DestroyQAbstract3DInputHandler(this); };
};

void QAbstract3DInputHandler_SetInputPosition(void* ptr, void* position)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->setInputPosition(*static_cast<QPoint*>(position));
}

void QAbstract3DInputHandler_SetInputView(void* ptr, long long inputView)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->setInputView(static_cast<QAbstract3DInputHandler::InputView>(inputView));
}

long long QAbstract3DInputHandler_InputView(void* ptr)
{
	return static_cast<QAbstract3DInputHandler*>(ptr)->inputView();
}

void* QAbstract3DInputHandler_NewQAbstract3DInputHandler(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstract3DInputHandler(static_cast<QWindow*>(parent));
	} else {
		return new MyQAbstract3DInputHandler(static_cast<QObject*>(parent));
	}
}

void QAbstract3DInputHandler_ConnectInputViewChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DInputHandler*>(ptr), static_cast<void (QAbstract3DInputHandler::*)(QAbstract3DInputHandler::InputView)>(&QAbstract3DInputHandler::inputViewChanged), static_cast<MyQAbstract3DInputHandler*>(ptr), static_cast<void (MyQAbstract3DInputHandler::*)(QAbstract3DInputHandler::InputView)>(&MyQAbstract3DInputHandler::Signal_InputViewChanged));
}

void QAbstract3DInputHandler_DisconnectInputViewChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DInputHandler*>(ptr), static_cast<void (QAbstract3DInputHandler::*)(QAbstract3DInputHandler::InputView)>(&QAbstract3DInputHandler::inputViewChanged), static_cast<MyQAbstract3DInputHandler*>(ptr), static_cast<void (MyQAbstract3DInputHandler::*)(QAbstract3DInputHandler::InputView)>(&MyQAbstract3DInputHandler::Signal_InputViewChanged));
}

void QAbstract3DInputHandler_InputViewChanged(void* ptr, long long view)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->inputViewChanged(static_cast<QAbstract3DInputHandler::InputView>(view));
}

void QAbstract3DInputHandler_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QAbstract3DInputHandler_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QAbstract3DInputHandler*>(ptr)->QAbstract3DInputHandler::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QAbstract3DInputHandler_MouseMoveEvent(void* ptr, void* event, void* mousePos)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void QAbstract3DInputHandler_MouseMoveEventDefault(void* ptr, void* event, void* mousePos)
{
		static_cast<QAbstract3DInputHandler*>(ptr)->QAbstract3DInputHandler::mouseMoveEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void QAbstract3DInputHandler_MousePressEvent(void* ptr, void* event, void* mousePos)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void QAbstract3DInputHandler_MousePressEventDefault(void* ptr, void* event, void* mousePos)
{
		static_cast<QAbstract3DInputHandler*>(ptr)->QAbstract3DInputHandler::mousePressEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void QAbstract3DInputHandler_MouseReleaseEvent(void* ptr, void* event, void* mousePos)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void QAbstract3DInputHandler_MouseReleaseEventDefault(void* ptr, void* event, void* mousePos)
{
		static_cast<QAbstract3DInputHandler*>(ptr)->QAbstract3DInputHandler::mouseReleaseEvent(static_cast<QMouseEvent*>(event), *static_cast<QPoint*>(mousePos));
}

void QAbstract3DInputHandler_ConnectPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DInputHandler*>(ptr), static_cast<void (QAbstract3DInputHandler::*)(const QPoint &)>(&QAbstract3DInputHandler::positionChanged), static_cast<MyQAbstract3DInputHandler*>(ptr), static_cast<void (MyQAbstract3DInputHandler::*)(const QPoint &)>(&MyQAbstract3DInputHandler::Signal_PositionChanged));
}

void QAbstract3DInputHandler_DisconnectPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DInputHandler*>(ptr), static_cast<void (QAbstract3DInputHandler::*)(const QPoint &)>(&QAbstract3DInputHandler::positionChanged), static_cast<MyQAbstract3DInputHandler*>(ptr), static_cast<void (MyQAbstract3DInputHandler::*)(const QPoint &)>(&MyQAbstract3DInputHandler::Signal_PositionChanged));
}

void QAbstract3DInputHandler_PositionChanged(void* ptr, void* position)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->positionChanged(*static_cast<QPoint*>(position));
}

void QAbstract3DInputHandler_ConnectSceneChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DInputHandler*>(ptr), static_cast<void (QAbstract3DInputHandler::*)(Q3DScene *)>(&QAbstract3DInputHandler::sceneChanged), static_cast<MyQAbstract3DInputHandler*>(ptr), static_cast<void (MyQAbstract3DInputHandler::*)(Q3DScene *)>(&MyQAbstract3DInputHandler::Signal_SceneChanged));
}

void QAbstract3DInputHandler_DisconnectSceneChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DInputHandler*>(ptr), static_cast<void (QAbstract3DInputHandler::*)(Q3DScene *)>(&QAbstract3DInputHandler::sceneChanged), static_cast<MyQAbstract3DInputHandler*>(ptr), static_cast<void (MyQAbstract3DInputHandler::*)(Q3DScene *)>(&MyQAbstract3DInputHandler::Signal_SceneChanged));
}

void QAbstract3DInputHandler_SceneChanged(void* ptr, void* scene)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->sceneChanged(static_cast<Q3DScene*>(scene));
}

void QAbstract3DInputHandler_SetPrevDistance(void* ptr, int distance)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->setPrevDistance(distance);
}

void QAbstract3DInputHandler_SetPreviousInputPos(void* ptr, void* position)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->setPreviousInputPos(*static_cast<QPoint*>(position));
}

void QAbstract3DInputHandler_SetScene(void* ptr, void* scene)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->setScene(static_cast<Q3DScene*>(scene));
}

void QAbstract3DInputHandler_TouchEvent(void* ptr, void* event)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->touchEvent(static_cast<QTouchEvent*>(event));
}

void QAbstract3DInputHandler_TouchEventDefault(void* ptr, void* event)
{
		static_cast<QAbstract3DInputHandler*>(ptr)->QAbstract3DInputHandler::touchEvent(static_cast<QTouchEvent*>(event));
}

void QAbstract3DInputHandler_WheelEvent(void* ptr, void* event)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QAbstract3DInputHandler_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QAbstract3DInputHandler*>(ptr)->QAbstract3DInputHandler::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QAbstract3DInputHandler_DestroyQAbstract3DInputHandler(void* ptr)
{
	static_cast<QAbstract3DInputHandler*>(ptr)->~QAbstract3DInputHandler();
}

void QAbstract3DInputHandler_DestroyQAbstract3DInputHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstract3DInputHandler_Scene(void* ptr)
{
	return static_cast<QAbstract3DInputHandler*>(ptr)->scene();
}

void* QAbstract3DInputHandler_InputPosition(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QAbstract3DInputHandler*>(ptr)->inputPosition(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QAbstract3DInputHandler_PreviousInputPos(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QAbstract3DInputHandler*>(ptr)->previousInputPos(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

int QAbstract3DInputHandler_PrevDistance(void* ptr)
{
	return static_cast<QAbstract3DInputHandler*>(ptr)->prevDistance();
}

class MyQAbstract3DSeries: public QAbstract3DSeries
{
public:
	void Signal_MultiHighlightGradientChanged(const QLinearGradient & gradient) { callbackQAbstract3DSeries_MultiHighlightGradientChanged(this, const_cast<QLinearGradient*>(&gradient)); };
	void Signal_BaseColorChanged(const QColor & color) { callbackQAbstract3DSeries_BaseColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_BaseGradientChanged(const QLinearGradient & gradient) { callbackQAbstract3DSeries_BaseGradientChanged(this, const_cast<QLinearGradient*>(&gradient)); };
	void Signal_ColorStyleChanged(Q3DTheme::ColorStyle style) { callbackQAbstract3DSeries_ColorStyleChanged(this, style); };
	void Signal_ItemLabelChanged(const QString & label) { QByteArray t64c653 = label.toUtf8(); QtDataVisualization_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };callbackQAbstract3DSeries_ItemLabelChanged(this, labelPacked); };
	void Signal_ItemLabelFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtDataVisualization_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQAbstract3DSeries_ItemLabelFormatChanged(this, formatPacked); };
	void Signal_ItemLabelVisibilityChanged(bool visible) { callbackQAbstract3DSeries_ItemLabelVisibilityChanged(this, visible); };
	void Signal_MeshChanged(QAbstract3DSeries::Mesh mesh) { callbackQAbstract3DSeries_MeshChanged(this, mesh); };
	void Signal_MeshRotationChanged(const QQuaternion & rotation) { callbackQAbstract3DSeries_MeshRotationChanged(this, const_cast<QQuaternion*>(&rotation)); };
	void Signal_MeshSmoothChanged(bool enabled) { callbackQAbstract3DSeries_MeshSmoothChanged(this, enabled); };
	void Signal_MultiHighlightColorChanged(const QColor & color) { callbackQAbstract3DSeries_MultiHighlightColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_NameChanged(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtDataVisualization_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQAbstract3DSeries_NameChanged(this, namePacked); };
	void Signal_SingleHighlightColorChanged(const QColor & color) { callbackQAbstract3DSeries_SingleHighlightColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_SingleHighlightGradientChanged(const QLinearGradient & gradient) { callbackQAbstract3DSeries_SingleHighlightGradientChanged(this, const_cast<QLinearGradient*>(&gradient)); };
	void Signal_UserDefinedMeshChanged(const QString & fileName) { QByteArray td83e09 = fileName.toUtf8(); QtDataVisualization_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQAbstract3DSeries_UserDefinedMeshChanged(this, fileNamePacked); };
	void Signal_VisibilityChanged(bool visible) { callbackQAbstract3DSeries_VisibilityChanged(this, visible); };
	 ~MyQAbstract3DSeries() { callbackQAbstract3DSeries_DestroyQAbstract3DSeries(this); };
};

void QAbstract3DSeries_ConnectMultiHighlightGradientChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QLinearGradient &)>(&QAbstract3DSeries::multiHighlightGradientChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QLinearGradient &)>(&MyQAbstract3DSeries::Signal_MultiHighlightGradientChanged));
}

void QAbstract3DSeries_DisconnectMultiHighlightGradientChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QLinearGradient &)>(&QAbstract3DSeries::multiHighlightGradientChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QLinearGradient &)>(&MyQAbstract3DSeries::Signal_MultiHighlightGradientChanged));
}

void QAbstract3DSeries_MultiHighlightGradientChanged(void* ptr, void* gradient)
{
	static_cast<QAbstract3DSeries*>(ptr)->multiHighlightGradientChanged(*static_cast<QLinearGradient*>(gradient));
}

void QAbstract3DSeries_SetMeshRotation(void* ptr, void* rotation)
{
	static_cast<QAbstract3DSeries*>(ptr)->setMeshRotation(*static_cast<QQuaternion*>(rotation));
}

void QAbstract3DSeries_ConnectBaseColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QColor &)>(&QAbstract3DSeries::baseColorChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QColor &)>(&MyQAbstract3DSeries::Signal_BaseColorChanged));
}

void QAbstract3DSeries_DisconnectBaseColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QColor &)>(&QAbstract3DSeries::baseColorChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QColor &)>(&MyQAbstract3DSeries::Signal_BaseColorChanged));
}

void QAbstract3DSeries_BaseColorChanged(void* ptr, void* color)
{
	static_cast<QAbstract3DSeries*>(ptr)->baseColorChanged(*static_cast<QColor*>(color));
}

void QAbstract3DSeries_ConnectBaseGradientChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QLinearGradient &)>(&QAbstract3DSeries::baseGradientChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QLinearGradient &)>(&MyQAbstract3DSeries::Signal_BaseGradientChanged));
}

void QAbstract3DSeries_DisconnectBaseGradientChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QLinearGradient &)>(&QAbstract3DSeries::baseGradientChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QLinearGradient &)>(&MyQAbstract3DSeries::Signal_BaseGradientChanged));
}

void QAbstract3DSeries_BaseGradientChanged(void* ptr, void* gradient)
{
	static_cast<QAbstract3DSeries*>(ptr)->baseGradientChanged(*static_cast<QLinearGradient*>(gradient));
}

void QAbstract3DSeries_ConnectColorStyleChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(Q3DTheme::ColorStyle)>(&QAbstract3DSeries::colorStyleChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(Q3DTheme::ColorStyle)>(&MyQAbstract3DSeries::Signal_ColorStyleChanged));
}

void QAbstract3DSeries_DisconnectColorStyleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(Q3DTheme::ColorStyle)>(&QAbstract3DSeries::colorStyleChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(Q3DTheme::ColorStyle)>(&MyQAbstract3DSeries::Signal_ColorStyleChanged));
}

void QAbstract3DSeries_ColorStyleChanged(void* ptr, long long style)
{
	static_cast<QAbstract3DSeries*>(ptr)->colorStyleChanged(static_cast<Q3DTheme::ColorStyle>(style));
}

void QAbstract3DSeries_ConnectItemLabelChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QString &)>(&QAbstract3DSeries::itemLabelChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QString &)>(&MyQAbstract3DSeries::Signal_ItemLabelChanged));
}

void QAbstract3DSeries_DisconnectItemLabelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QString &)>(&QAbstract3DSeries::itemLabelChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QString &)>(&MyQAbstract3DSeries::Signal_ItemLabelChanged));
}

void QAbstract3DSeries_ItemLabelChanged(void* ptr, struct QtDataVisualization_PackedString label)
{
	static_cast<QAbstract3DSeries*>(ptr)->itemLabelChanged(QString::fromUtf8(label.data, label.len));
}

void QAbstract3DSeries_ConnectItemLabelFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QString &)>(&QAbstract3DSeries::itemLabelFormatChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QString &)>(&MyQAbstract3DSeries::Signal_ItemLabelFormatChanged));
}

void QAbstract3DSeries_DisconnectItemLabelFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QString &)>(&QAbstract3DSeries::itemLabelFormatChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QString &)>(&MyQAbstract3DSeries::Signal_ItemLabelFormatChanged));
}

void QAbstract3DSeries_ItemLabelFormatChanged(void* ptr, struct QtDataVisualization_PackedString format)
{
	static_cast<QAbstract3DSeries*>(ptr)->itemLabelFormatChanged(QString::fromUtf8(format.data, format.len));
}

void QAbstract3DSeries_ConnectItemLabelVisibilityChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(bool)>(&QAbstract3DSeries::itemLabelVisibilityChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(bool)>(&MyQAbstract3DSeries::Signal_ItemLabelVisibilityChanged));
}

void QAbstract3DSeries_DisconnectItemLabelVisibilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(bool)>(&QAbstract3DSeries::itemLabelVisibilityChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(bool)>(&MyQAbstract3DSeries::Signal_ItemLabelVisibilityChanged));
}

void QAbstract3DSeries_ItemLabelVisibilityChanged(void* ptr, char visible)
{
	static_cast<QAbstract3DSeries*>(ptr)->itemLabelVisibilityChanged(visible != 0);
}

void QAbstract3DSeries_ConnectMeshChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(QAbstract3DSeries::Mesh)>(&QAbstract3DSeries::meshChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(QAbstract3DSeries::Mesh)>(&MyQAbstract3DSeries::Signal_MeshChanged));
}

void QAbstract3DSeries_DisconnectMeshChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(QAbstract3DSeries::Mesh)>(&QAbstract3DSeries::meshChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(QAbstract3DSeries::Mesh)>(&MyQAbstract3DSeries::Signal_MeshChanged));
}

void QAbstract3DSeries_MeshChanged(void* ptr, long long mesh)
{
	static_cast<QAbstract3DSeries*>(ptr)->meshChanged(static_cast<QAbstract3DSeries::Mesh>(mesh));
}

void QAbstract3DSeries_ConnectMeshRotationChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QQuaternion &)>(&QAbstract3DSeries::meshRotationChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QQuaternion &)>(&MyQAbstract3DSeries::Signal_MeshRotationChanged));
}

void QAbstract3DSeries_DisconnectMeshRotationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QQuaternion &)>(&QAbstract3DSeries::meshRotationChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QQuaternion &)>(&MyQAbstract3DSeries::Signal_MeshRotationChanged));
}

void QAbstract3DSeries_MeshRotationChanged(void* ptr, void* rotation)
{
	static_cast<QAbstract3DSeries*>(ptr)->meshRotationChanged(*static_cast<QQuaternion*>(rotation));
}

void QAbstract3DSeries_ConnectMeshSmoothChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(bool)>(&QAbstract3DSeries::meshSmoothChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(bool)>(&MyQAbstract3DSeries::Signal_MeshSmoothChanged));
}

void QAbstract3DSeries_DisconnectMeshSmoothChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(bool)>(&QAbstract3DSeries::meshSmoothChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(bool)>(&MyQAbstract3DSeries::Signal_MeshSmoothChanged));
}

void QAbstract3DSeries_MeshSmoothChanged(void* ptr, char enabled)
{
	static_cast<QAbstract3DSeries*>(ptr)->meshSmoothChanged(enabled != 0);
}

void QAbstract3DSeries_ConnectMultiHighlightColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QColor &)>(&QAbstract3DSeries::multiHighlightColorChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QColor &)>(&MyQAbstract3DSeries::Signal_MultiHighlightColorChanged));
}

void QAbstract3DSeries_DisconnectMultiHighlightColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QColor &)>(&QAbstract3DSeries::multiHighlightColorChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QColor &)>(&MyQAbstract3DSeries::Signal_MultiHighlightColorChanged));
}

void QAbstract3DSeries_MultiHighlightColorChanged(void* ptr, void* color)
{
	static_cast<QAbstract3DSeries*>(ptr)->multiHighlightColorChanged(*static_cast<QColor*>(color));
}

void QAbstract3DSeries_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QString &)>(&QAbstract3DSeries::nameChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QString &)>(&MyQAbstract3DSeries::Signal_NameChanged));
}

void QAbstract3DSeries_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QString &)>(&QAbstract3DSeries::nameChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QString &)>(&MyQAbstract3DSeries::Signal_NameChanged));
}

void QAbstract3DSeries_NameChanged(void* ptr, struct QtDataVisualization_PackedString name)
{
	static_cast<QAbstract3DSeries*>(ptr)->nameChanged(QString::fromUtf8(name.data, name.len));
}

void QAbstract3DSeries_SetBaseColor(void* ptr, void* color)
{
	static_cast<QAbstract3DSeries*>(ptr)->setBaseColor(*static_cast<QColor*>(color));
}

void QAbstract3DSeries_SetBaseGradient(void* ptr, void* gradient)
{
	static_cast<QAbstract3DSeries*>(ptr)->setBaseGradient(*static_cast<QLinearGradient*>(gradient));
}

void QAbstract3DSeries_SetColorStyle(void* ptr, long long style)
{
	static_cast<QAbstract3DSeries*>(ptr)->setColorStyle(static_cast<Q3DTheme::ColorStyle>(style));
}

void QAbstract3DSeries_SetItemLabelFormat(void* ptr, struct QtDataVisualization_PackedString format)
{
	static_cast<QAbstract3DSeries*>(ptr)->setItemLabelFormat(QString::fromUtf8(format.data, format.len));
}

void QAbstract3DSeries_SetItemLabelVisible(void* ptr, char visible)
{
	static_cast<QAbstract3DSeries*>(ptr)->setItemLabelVisible(visible != 0);
}

void QAbstract3DSeries_SetMesh(void* ptr, long long mesh)
{
	static_cast<QAbstract3DSeries*>(ptr)->setMesh(static_cast<QAbstract3DSeries::Mesh>(mesh));
}

void QAbstract3DSeries_SetMeshAxisAndAngle(void* ptr, void* axis, float angle)
{
	static_cast<QAbstract3DSeries*>(ptr)->setMeshAxisAndAngle(*static_cast<QVector3D*>(axis), angle);
}

void QAbstract3DSeries_SetMeshSmooth(void* ptr, char enable)
{
	static_cast<QAbstract3DSeries*>(ptr)->setMeshSmooth(enable != 0);
}

void QAbstract3DSeries_SetMultiHighlightColor(void* ptr, void* color)
{
	static_cast<QAbstract3DSeries*>(ptr)->setMultiHighlightColor(*static_cast<QColor*>(color));
}

void QAbstract3DSeries_SetMultiHighlightGradient(void* ptr, void* gradient)
{
	static_cast<QAbstract3DSeries*>(ptr)->setMultiHighlightGradient(*static_cast<QLinearGradient*>(gradient));
}

void QAbstract3DSeries_SetName(void* ptr, struct QtDataVisualization_PackedString name)
{
	static_cast<QAbstract3DSeries*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QAbstract3DSeries_SetSingleHighlightColor(void* ptr, void* color)
{
	static_cast<QAbstract3DSeries*>(ptr)->setSingleHighlightColor(*static_cast<QColor*>(color));
}

void QAbstract3DSeries_SetSingleHighlightGradient(void* ptr, void* gradient)
{
	static_cast<QAbstract3DSeries*>(ptr)->setSingleHighlightGradient(*static_cast<QLinearGradient*>(gradient));
}

void QAbstract3DSeries_SetUserDefinedMesh(void* ptr, struct QtDataVisualization_PackedString fileName)
{
	static_cast<QAbstract3DSeries*>(ptr)->setUserDefinedMesh(QString::fromUtf8(fileName.data, fileName.len));
}

void QAbstract3DSeries_SetVisible(void* ptr, char visible)
{
	static_cast<QAbstract3DSeries*>(ptr)->setVisible(visible != 0);
}

void QAbstract3DSeries_ConnectSingleHighlightColorChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QColor &)>(&QAbstract3DSeries::singleHighlightColorChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QColor &)>(&MyQAbstract3DSeries::Signal_SingleHighlightColorChanged));
}

void QAbstract3DSeries_DisconnectSingleHighlightColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QColor &)>(&QAbstract3DSeries::singleHighlightColorChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QColor &)>(&MyQAbstract3DSeries::Signal_SingleHighlightColorChanged));
}

void QAbstract3DSeries_SingleHighlightColorChanged(void* ptr, void* color)
{
	static_cast<QAbstract3DSeries*>(ptr)->singleHighlightColorChanged(*static_cast<QColor*>(color));
}

void QAbstract3DSeries_ConnectSingleHighlightGradientChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QLinearGradient &)>(&QAbstract3DSeries::singleHighlightGradientChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QLinearGradient &)>(&MyQAbstract3DSeries::Signal_SingleHighlightGradientChanged));
}

void QAbstract3DSeries_DisconnectSingleHighlightGradientChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QLinearGradient &)>(&QAbstract3DSeries::singleHighlightGradientChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QLinearGradient &)>(&MyQAbstract3DSeries::Signal_SingleHighlightGradientChanged));
}

void QAbstract3DSeries_SingleHighlightGradientChanged(void* ptr, void* gradient)
{
	static_cast<QAbstract3DSeries*>(ptr)->singleHighlightGradientChanged(*static_cast<QLinearGradient*>(gradient));
}

void QAbstract3DSeries_ConnectUserDefinedMeshChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QString &)>(&QAbstract3DSeries::userDefinedMeshChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QString &)>(&MyQAbstract3DSeries::Signal_UserDefinedMeshChanged));
}

void QAbstract3DSeries_DisconnectUserDefinedMeshChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(const QString &)>(&QAbstract3DSeries::userDefinedMeshChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(const QString &)>(&MyQAbstract3DSeries::Signal_UserDefinedMeshChanged));
}

void QAbstract3DSeries_UserDefinedMeshChanged(void* ptr, struct QtDataVisualization_PackedString fileName)
{
	static_cast<QAbstract3DSeries*>(ptr)->userDefinedMeshChanged(QString::fromUtf8(fileName.data, fileName.len));
}

void QAbstract3DSeries_ConnectVisibilityChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(bool)>(&QAbstract3DSeries::visibilityChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(bool)>(&MyQAbstract3DSeries::Signal_VisibilityChanged));
}

void QAbstract3DSeries_DisconnectVisibilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstract3DSeries*>(ptr), static_cast<void (QAbstract3DSeries::*)(bool)>(&QAbstract3DSeries::visibilityChanged), static_cast<MyQAbstract3DSeries*>(ptr), static_cast<void (MyQAbstract3DSeries::*)(bool)>(&MyQAbstract3DSeries::Signal_VisibilityChanged));
}

void QAbstract3DSeries_VisibilityChanged(void* ptr, char visible)
{
	static_cast<QAbstract3DSeries*>(ptr)->visibilityChanged(visible != 0);
}

void QAbstract3DSeries_DestroyQAbstract3DSeries(void* ptr)
{
	static_cast<QAbstract3DSeries*>(ptr)->~QAbstract3DSeries();
}

void QAbstract3DSeries_DestroyQAbstract3DSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QAbstract3DSeries_Mesh(void* ptr)
{
	return static_cast<QAbstract3DSeries*>(ptr)->mesh();
}

long long QAbstract3DSeries_ColorStyle(void* ptr)
{
	return static_cast<QAbstract3DSeries*>(ptr)->colorStyle();
}

void* QAbstract3DSeries_BaseColor(void* ptr)
{
	return new QColor(static_cast<QAbstract3DSeries*>(ptr)->baseColor());
}

void* QAbstract3DSeries_MultiHighlightColor(void* ptr)
{
	return new QColor(static_cast<QAbstract3DSeries*>(ptr)->multiHighlightColor());
}

void* QAbstract3DSeries_SingleHighlightColor(void* ptr)
{
	return new QColor(static_cast<QAbstract3DSeries*>(ptr)->singleHighlightColor());
}

void* QAbstract3DSeries_BaseGradient(void* ptr)
{
	return new QLinearGradient(static_cast<QAbstract3DSeries*>(ptr)->baseGradient());
}

void* QAbstract3DSeries_MultiHighlightGradient(void* ptr)
{
	return new QLinearGradient(static_cast<QAbstract3DSeries*>(ptr)->multiHighlightGradient());
}

void* QAbstract3DSeries_SingleHighlightGradient(void* ptr)
{
	return new QLinearGradient(static_cast<QAbstract3DSeries*>(ptr)->singleHighlightGradient());
}

void* QAbstract3DSeries_MeshRotation(void* ptr)
{
	return new QQuaternion(static_cast<QAbstract3DSeries*>(ptr)->meshRotation());
}

struct QtDataVisualization_PackedString QAbstract3DSeries_ItemLabel(void* ptr)
{
	return ({ QByteArray tdad3c7 = static_cast<QAbstract3DSeries*>(ptr)->itemLabel().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tdad3c7.prepend("WHITESPACE").constData()+10), tdad3c7.size()-10 }; });
}

struct QtDataVisualization_PackedString QAbstract3DSeries_ItemLabelFormat(void* ptr)
{
	return ({ QByteArray tf0cb25 = static_cast<QAbstract3DSeries*>(ptr)->itemLabelFormat().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tf0cb25.prepend("WHITESPACE").constData()+10), tf0cb25.size()-10 }; });
}

struct QtDataVisualization_PackedString QAbstract3DSeries_Name(void* ptr)
{
	return ({ QByteArray t7f4311 = static_cast<QAbstract3DSeries*>(ptr)->name().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t7f4311.prepend("WHITESPACE").constData()+10), t7f4311.size()-10 }; });
}

struct QtDataVisualization_PackedString QAbstract3DSeries_UserDefinedMesh(void* ptr)
{
	return ({ QByteArray t448627 = static_cast<QAbstract3DSeries*>(ptr)->userDefinedMesh().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t448627.prepend("WHITESPACE").constData()+10), t448627.size()-10 }; });
}

long long QAbstract3DSeries_Type(void* ptr)
{
	return static_cast<QAbstract3DSeries*>(ptr)->type();
}

char QAbstract3DSeries_IsItemLabelVisible(void* ptr)
{
	return static_cast<QAbstract3DSeries*>(ptr)->isItemLabelVisible();
}

char QAbstract3DSeries_IsMeshSmooth(void* ptr)
{
	return static_cast<QAbstract3DSeries*>(ptr)->isMeshSmooth();
}

char QAbstract3DSeries_IsVisible(void* ptr)
{
	return static_cast<QAbstract3DSeries*>(ptr)->isVisible();
}

class MyQAbstractDataProxy: public QAbstractDataProxy
{
public:
	 ~MyQAbstractDataProxy() { callbackQAbstractDataProxy_DestroyQAbstractDataProxy(this); };
};

long long QAbstractDataProxy_Type(void* ptr)
{
	return static_cast<QAbstractDataProxy*>(ptr)->type();
}

void QAbstractDataProxy_DestroyQAbstractDataProxy(void* ptr)
{
	static_cast<QAbstractDataProxy*>(ptr)->~QAbstractDataProxy();
}

void QAbstractDataProxy_DestroyQAbstractDataProxyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQBar3DSeries: public QBar3DSeries
{
public:
	MyQBar3DSeries(QBarDataProxy *dataProxy, QObject *parent = Q_NULLPTR) : QBar3DSeries(dataProxy, parent) {};
	MyQBar3DSeries(QObject *parent = Q_NULLPTR) : QBar3DSeries(parent) {};
	void Signal_DataProxyChanged(QBarDataProxy * proxy) { callbackQBar3DSeries_DataProxyChanged(this, proxy); };
	void Signal_MeshAngleChanged(float angle) { callbackQBar3DSeries_MeshAngleChanged(this, angle); };
	void Signal_SelectedBarChanged(const QPoint & position) { callbackQBar3DSeries_SelectedBarChanged(this, const_cast<QPoint*>(&position)); };
	 ~MyQBar3DSeries() { callbackQBar3DSeries_DestroyQBar3DSeries(this); };
};

void* QBar3DSeries_NewQBar3DSeries2(void* dataProxy, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QWindow*>(parent));
	} else {
		return new MyQBar3DSeries(static_cast<QBarDataProxy*>(dataProxy), static_cast<QObject*>(parent));
	}
}

void* QBar3DSeries_NewQBar3DSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBar3DSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQBar3DSeries(static_cast<QObject*>(parent));
	}
}

void* QBar3DSeries_QBar3DSeries_InvalidSelectionPosition()
{
	return ({ QPoint tmpValue = QBar3DSeries::invalidSelectionPosition(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QBar3DSeries_ConnectDataProxyChanged(void* ptr)
{
	QObject::connect(static_cast<QBar3DSeries*>(ptr), static_cast<void (QBar3DSeries::*)(QBarDataProxy *)>(&QBar3DSeries::dataProxyChanged), static_cast<MyQBar3DSeries*>(ptr), static_cast<void (MyQBar3DSeries::*)(QBarDataProxy *)>(&MyQBar3DSeries::Signal_DataProxyChanged));
}

void QBar3DSeries_DisconnectDataProxyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBar3DSeries*>(ptr), static_cast<void (QBar3DSeries::*)(QBarDataProxy *)>(&QBar3DSeries::dataProxyChanged), static_cast<MyQBar3DSeries*>(ptr), static_cast<void (MyQBar3DSeries::*)(QBarDataProxy *)>(&MyQBar3DSeries::Signal_DataProxyChanged));
}

void QBar3DSeries_DataProxyChanged(void* ptr, void* proxy)
{
	static_cast<QBar3DSeries*>(ptr)->dataProxyChanged(static_cast<QBarDataProxy*>(proxy));
}

void QBar3DSeries_ConnectMeshAngleChanged(void* ptr)
{
	QObject::connect(static_cast<QBar3DSeries*>(ptr), static_cast<void (QBar3DSeries::*)(float)>(&QBar3DSeries::meshAngleChanged), static_cast<MyQBar3DSeries*>(ptr), static_cast<void (MyQBar3DSeries::*)(float)>(&MyQBar3DSeries::Signal_MeshAngleChanged));
}

void QBar3DSeries_DisconnectMeshAngleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBar3DSeries*>(ptr), static_cast<void (QBar3DSeries::*)(float)>(&QBar3DSeries::meshAngleChanged), static_cast<MyQBar3DSeries*>(ptr), static_cast<void (MyQBar3DSeries::*)(float)>(&MyQBar3DSeries::Signal_MeshAngleChanged));
}

void QBar3DSeries_MeshAngleChanged(void* ptr, float angle)
{
	static_cast<QBar3DSeries*>(ptr)->meshAngleChanged(angle);
}

void QBar3DSeries_ConnectSelectedBarChanged(void* ptr)
{
	QObject::connect(static_cast<QBar3DSeries*>(ptr), static_cast<void (QBar3DSeries::*)(const QPoint &)>(&QBar3DSeries::selectedBarChanged), static_cast<MyQBar3DSeries*>(ptr), static_cast<void (MyQBar3DSeries::*)(const QPoint &)>(&MyQBar3DSeries::Signal_SelectedBarChanged));
}

void QBar3DSeries_DisconnectSelectedBarChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBar3DSeries*>(ptr), static_cast<void (QBar3DSeries::*)(const QPoint &)>(&QBar3DSeries::selectedBarChanged), static_cast<MyQBar3DSeries*>(ptr), static_cast<void (MyQBar3DSeries::*)(const QPoint &)>(&MyQBar3DSeries::Signal_SelectedBarChanged));
}

void QBar3DSeries_SelectedBarChanged(void* ptr, void* position)
{
	static_cast<QBar3DSeries*>(ptr)->selectedBarChanged(*static_cast<QPoint*>(position));
}

void QBar3DSeries_SetDataProxy(void* ptr, void* proxy)
{
	static_cast<QBar3DSeries*>(ptr)->setDataProxy(static_cast<QBarDataProxy*>(proxy));
}

void QBar3DSeries_SetMeshAngle(void* ptr, float angle)
{
	static_cast<QBar3DSeries*>(ptr)->setMeshAngle(angle);
}

void QBar3DSeries_SetSelectedBar(void* ptr, void* position)
{
	static_cast<QBar3DSeries*>(ptr)->setSelectedBar(*static_cast<QPoint*>(position));
}

void QBar3DSeries_DestroyQBar3DSeries(void* ptr)
{
	static_cast<QBar3DSeries*>(ptr)->~QBar3DSeries();
}

void QBar3DSeries_DestroyQBar3DSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QBar3DSeries_DataProxy(void* ptr)
{
	return static_cast<QBar3DSeries*>(ptr)->dataProxy();
}

void* QBar3DSeries_SelectedBar(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QBar3DSeries*>(ptr)->selectedBar(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

float QBar3DSeries_MeshAngle(void* ptr)
{
	return static_cast<QBar3DSeries*>(ptr)->meshAngle();
}

void* QBarDataItem_NewQBarDataItem()
{
	return new QBarDataItem();
}

void* QBarDataItem_NewQBarDataItem4(void* other)
{
	return new QBarDataItem(*static_cast<QBarDataItem*>(other));
}

void* QBarDataItem_NewQBarDataItem2(float value)
{
	return new QBarDataItem(value);
}

void* QBarDataItem_NewQBarDataItem3(float value, float angle)
{
	return new QBarDataItem(value, angle);
}

void QBarDataItem_SetRotation(void* ptr, float angle)
{
	static_cast<QBarDataItem*>(ptr)->setRotation(angle);
}

void QBarDataItem_SetValue(void* ptr, float val)
{
	static_cast<QBarDataItem*>(ptr)->setValue(val);
}

void QBarDataItem_DestroyQBarDataItem(void* ptr)
{
	static_cast<QBarDataItem*>(ptr)->~QBarDataItem();
}

float QBarDataItem_Rotation(void* ptr)
{
	return static_cast<QBarDataItem*>(ptr)->rotation();
}

float QBarDataItem_Value(void* ptr)
{
	return static_cast<QBarDataItem*>(ptr)->value();
}

class MyQBarDataProxy: public QBarDataProxy
{
public:
	MyQBarDataProxy(QObject *parent = Q_NULLPTR) : QBarDataProxy(parent) {};
	void Signal_ArrayReset() { callbackQBarDataProxy_ArrayReset(this); };
	void Signal_ColumnLabelsChanged() { callbackQBarDataProxy_ColumnLabelsChanged(this); };
	void Signal_ItemChanged(int rowIndex, int columnIndex) { callbackQBarDataProxy_ItemChanged(this, rowIndex, columnIndex); };
	void Signal_RowCountChanged(int count) { callbackQBarDataProxy_RowCountChanged(this, count); };
	void Signal_RowLabelsChanged() { callbackQBarDataProxy_RowLabelsChanged(this); };
	void Signal_RowsAdded(int startIndex, int count) { callbackQBarDataProxy_RowsAdded(this, startIndex, count); };
	void Signal_RowsChanged(int startIndex, int count) { callbackQBarDataProxy_RowsChanged(this, startIndex, count); };
	void Signal_RowsInserted(int startIndex, int count) { callbackQBarDataProxy_RowsInserted(this, startIndex, count); };
	void Signal_RowsRemoved(int startIndex, int count) { callbackQBarDataProxy_RowsRemoved(this, startIndex, count); };
	void Signal_SeriesChanged(QBar3DSeries * series) { callbackQBarDataProxy_SeriesChanged(this, series); };
	 ~MyQBarDataProxy() { callbackQBarDataProxy_DestroyQBarDataProxy(this); };
};

void* QBarDataProxy_NewQBarDataProxy(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQBarDataProxy(static_cast<QWindow*>(parent));
	} else {
		return new MyQBarDataProxy(static_cast<QObject*>(parent));
	}
}

void QBarDataProxy_ConnectArrayReset(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)()>(&QBarDataProxy::arrayReset), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)()>(&MyQBarDataProxy::Signal_ArrayReset));
}

void QBarDataProxy_DisconnectArrayReset(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)()>(&QBarDataProxy::arrayReset), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)()>(&MyQBarDataProxy::Signal_ArrayReset));
}

void QBarDataProxy_ArrayReset(void* ptr)
{
	static_cast<QBarDataProxy*>(ptr)->arrayReset();
}

void QBarDataProxy_ConnectColumnLabelsChanged(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)()>(&QBarDataProxy::columnLabelsChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)()>(&MyQBarDataProxy::Signal_ColumnLabelsChanged));
}

void QBarDataProxy_DisconnectColumnLabelsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)()>(&QBarDataProxy::columnLabelsChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)()>(&MyQBarDataProxy::Signal_ColumnLabelsChanged));
}

void QBarDataProxy_ColumnLabelsChanged(void* ptr)
{
	static_cast<QBarDataProxy*>(ptr)->columnLabelsChanged();
}

void QBarDataProxy_ConnectItemChanged(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::itemChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_ItemChanged));
}

void QBarDataProxy_DisconnectItemChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::itemChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_ItemChanged));
}

void QBarDataProxy_ItemChanged(void* ptr, int rowIndex, int columnIndex)
{
	static_cast<QBarDataProxy*>(ptr)->itemChanged(rowIndex, columnIndex);
}

void QBarDataProxy_RemoveRows(void* ptr, int rowIndex, int removeCount, char removeLabels)
{
	static_cast<QBarDataProxy*>(ptr)->removeRows(rowIndex, removeCount, removeLabels != 0);
}

void QBarDataProxy_ResetArray(void* ptr)
{
	static_cast<QBarDataProxy*>(ptr)->resetArray();
}

void QBarDataProxy_ConnectRowCountChanged(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int)>(&QBarDataProxy::rowCountChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int)>(&MyQBarDataProxy::Signal_RowCountChanged));
}

void QBarDataProxy_DisconnectRowCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int)>(&QBarDataProxy::rowCountChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int)>(&MyQBarDataProxy::Signal_RowCountChanged));
}

void QBarDataProxy_RowCountChanged(void* ptr, int count)
{
	static_cast<QBarDataProxy*>(ptr)->rowCountChanged(count);
}

void QBarDataProxy_ConnectRowLabelsChanged(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)()>(&QBarDataProxy::rowLabelsChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)()>(&MyQBarDataProxy::Signal_RowLabelsChanged));
}

void QBarDataProxy_DisconnectRowLabelsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)()>(&QBarDataProxy::rowLabelsChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)()>(&MyQBarDataProxy::Signal_RowLabelsChanged));
}

void QBarDataProxy_RowLabelsChanged(void* ptr)
{
	static_cast<QBarDataProxy*>(ptr)->rowLabelsChanged();
}

void QBarDataProxy_ConnectRowsAdded(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::rowsAdded), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_RowsAdded));
}

void QBarDataProxy_DisconnectRowsAdded(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::rowsAdded), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_RowsAdded));
}

void QBarDataProxy_RowsAdded(void* ptr, int startIndex, int count)
{
	static_cast<QBarDataProxy*>(ptr)->rowsAdded(startIndex, count);
}

void QBarDataProxy_ConnectRowsChanged(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::rowsChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_RowsChanged));
}

void QBarDataProxy_DisconnectRowsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::rowsChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_RowsChanged));
}

void QBarDataProxy_RowsChanged(void* ptr, int startIndex, int count)
{
	static_cast<QBarDataProxy*>(ptr)->rowsChanged(startIndex, count);
}

void QBarDataProxy_ConnectRowsInserted(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::rowsInserted), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_RowsInserted));
}

void QBarDataProxy_DisconnectRowsInserted(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::rowsInserted), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_RowsInserted));
}

void QBarDataProxy_RowsInserted(void* ptr, int startIndex, int count)
{
	static_cast<QBarDataProxy*>(ptr)->rowsInserted(startIndex, count);
}

void QBarDataProxy_ConnectRowsRemoved(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::rowsRemoved), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_RowsRemoved));
}

void QBarDataProxy_DisconnectRowsRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(int, int)>(&QBarDataProxy::rowsRemoved), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(int, int)>(&MyQBarDataProxy::Signal_RowsRemoved));
}

void QBarDataProxy_RowsRemoved(void* ptr, int startIndex, int count)
{
	static_cast<QBarDataProxy*>(ptr)->rowsRemoved(startIndex, count);
}

void QBarDataProxy_ConnectSeriesChanged(void* ptr)
{
	QObject::connect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(QBar3DSeries *)>(&QBarDataProxy::seriesChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(QBar3DSeries *)>(&MyQBarDataProxy::Signal_SeriesChanged));
}

void QBarDataProxy_DisconnectSeriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QBarDataProxy*>(ptr), static_cast<void (QBarDataProxy::*)(QBar3DSeries *)>(&QBarDataProxy::seriesChanged), static_cast<MyQBarDataProxy*>(ptr), static_cast<void (MyQBarDataProxy::*)(QBar3DSeries *)>(&MyQBarDataProxy::Signal_SeriesChanged));
}

void QBarDataProxy_SeriesChanged(void* ptr, void* series)
{
	static_cast<QBarDataProxy*>(ptr)->seriesChanged(static_cast<QBar3DSeries*>(series));
}

void QBarDataProxy_SetColumnLabels(void* ptr, struct QtDataVisualization_PackedString labels)
{
	static_cast<QBarDataProxy*>(ptr)->setColumnLabels(QString::fromUtf8(labels.data, labels.len).split("|", QString::SkipEmptyParts));
}

void QBarDataProxy_SetItem2(void* ptr, void* position, void* item)
{
	static_cast<QBarDataProxy*>(ptr)->setItem(*static_cast<QPoint*>(position), *static_cast<QBarDataItem*>(item));
}

void QBarDataProxy_SetItem(void* ptr, int rowIndex, int columnIndex, void* item)
{
	static_cast<QBarDataProxy*>(ptr)->setItem(rowIndex, columnIndex, *static_cast<QBarDataItem*>(item));
}

void QBarDataProxy_SetRowLabels(void* ptr, struct QtDataVisualization_PackedString labels)
{
	static_cast<QBarDataProxy*>(ptr)->setRowLabels(QString::fromUtf8(labels.data, labels.len).split("|", QString::SkipEmptyParts));
}

void QBarDataProxy_DestroyQBarDataProxy(void* ptr)
{
	static_cast<QBarDataProxy*>(ptr)->~QBarDataProxy();
}

void QBarDataProxy_DestroyQBarDataProxyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QBarDataProxy_Series(void* ptr)
{
	return static_cast<QBarDataProxy*>(ptr)->series();
}

struct QtDataVisualization_PackedString QBarDataProxy_ColumnLabels(void* ptr)
{
	return ({ QByteArray t7f1685 = static_cast<QBarDataProxy*>(ptr)->columnLabels().join("|").toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t7f1685.prepend("WHITESPACE").constData()+10), t7f1685.size()-10 }; });
}

struct QtDataVisualization_PackedString QBarDataProxy_RowLabels(void* ptr)
{
	return ({ QByteArray t0565e1 = static_cast<QBarDataProxy*>(ptr)->rowLabels().join("|").toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t0565e1.prepend("WHITESPACE").constData()+10), t0565e1.size()-10 }; });
}

void* QBarDataProxy_ItemAt2(void* ptr, void* position)
{
	return const_cast<QBarDataItem*>(static_cast<QBarDataProxy*>(ptr)->itemAt(*static_cast<QPoint*>(position)));
}

void* QBarDataProxy_ItemAt(void* ptr, int rowIndex, int columnIndex)
{
	return const_cast<QBarDataItem*>(static_cast<QBarDataProxy*>(ptr)->itemAt(rowIndex, columnIndex));
}

int QBarDataProxy_RowCount(void* ptr)
{
	return static_cast<QBarDataProxy*>(ptr)->rowCount();
}

class MyQCategory3DAxis: public QCategory3DAxis
{
public:
	MyQCategory3DAxis(QObject *parent = Q_NULLPTR) : QCategory3DAxis(parent) {};
	void Signal_LabelsChanged() { callbackQCategory3DAxis_LabelsChanged(this); };
	 ~MyQCategory3DAxis() { callbackQCategory3DAxis_DestroyQCategory3DAxis(this); };
};

void* QCategory3DAxis_NewQCategory3DAxis(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCategory3DAxis(static_cast<QWindow*>(parent));
	} else {
		return new MyQCategory3DAxis(static_cast<QObject*>(parent));
	}
}

void QCategory3DAxis_ConnectLabelsChanged(void* ptr)
{
	QObject::connect(static_cast<QCategory3DAxis*>(ptr), static_cast<void (QCategory3DAxis::*)()>(&QCategory3DAxis::labelsChanged), static_cast<MyQCategory3DAxis*>(ptr), static_cast<void (MyQCategory3DAxis::*)()>(&MyQCategory3DAxis::Signal_LabelsChanged));
}

void QCategory3DAxis_DisconnectLabelsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCategory3DAxis*>(ptr), static_cast<void (QCategory3DAxis::*)()>(&QCategory3DAxis::labelsChanged), static_cast<MyQCategory3DAxis*>(ptr), static_cast<void (MyQCategory3DAxis::*)()>(&MyQCategory3DAxis::Signal_LabelsChanged));
}

void QCategory3DAxis_LabelsChanged(void* ptr)
{
	static_cast<QCategory3DAxis*>(ptr)->labelsChanged();
}

void QCategory3DAxis_SetLabels(void* ptr, struct QtDataVisualization_PackedString labels)
{
	static_cast<QCategory3DAxis*>(ptr)->setLabels(QString::fromUtf8(labels.data, labels.len).split("|", QString::SkipEmptyParts));
}

void QCategory3DAxis_DestroyQCategory3DAxis(void* ptr)
{
	static_cast<QCategory3DAxis*>(ptr)->~QCategory3DAxis();
}

void QCategory3DAxis_DestroyQCategory3DAxisDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDataVisualization_PackedString QCategory3DAxis_Labels(void* ptr)
{
	return ({ QByteArray t2a027f = static_cast<QCategory3DAxis*>(ptr)->labels().join("|").toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t2a027f.prepend("WHITESPACE").constData()+10), t2a027f.size()-10 }; });
}

class MyQCustom3DItem: public QCustom3DItem
{
public:
	MyQCustom3DItem(QObject *parent = Q_NULLPTR) : QCustom3DItem(parent) {};
	MyQCustom3DItem(const QString &meshFile, const QVector3D &position, const QVector3D &scaling, const QQuaternion &rotation, const QImage &texture, QObject *parent = Q_NULLPTR) : QCustom3DItem(meshFile, position, scaling, rotation, texture, parent) {};
	void Signal_MeshFileChanged(const QString & meshFile) { QByteArray tdc0aec = meshFile.toUtf8(); QtDataVisualization_PackedString meshFilePacked = { const_cast<char*>(tdc0aec.prepend("WHITESPACE").constData()+10), tdc0aec.size()-10 };callbackQCustom3DItem_MeshFileChanged(this, meshFilePacked); };
	void Signal_PositionAbsoluteChanged(bool positionAbsolute) { callbackQCustom3DItem_PositionAbsoluteChanged(this, positionAbsolute); };
	void Signal_PositionChanged(const QVector3D & position) { callbackQCustom3DItem_PositionChanged(this, const_cast<QVector3D*>(&position)); };
	void Signal_RotationChanged(const QQuaternion & rotation) { callbackQCustom3DItem_RotationChanged(this, const_cast<QQuaternion*>(&rotation)); };
	void Signal_ScalingAbsoluteChanged(bool scalingAbsolute) { callbackQCustom3DItem_ScalingAbsoluteChanged(this, scalingAbsolute); };
	void Signal_ScalingChanged(const QVector3D & scaling) { callbackQCustom3DItem_ScalingChanged(this, const_cast<QVector3D*>(&scaling)); };
	void Signal_ShadowCastingChanged(bool shadowCasting) { callbackQCustom3DItem_ShadowCastingChanged(this, shadowCasting); };
	void Signal_TextureFileChanged(const QString & textureFile) { QByteArray t19f53d = textureFile.toUtf8(); QtDataVisualization_PackedString textureFilePacked = { const_cast<char*>(t19f53d.prepend("WHITESPACE").constData()+10), t19f53d.size()-10 };callbackQCustom3DItem_TextureFileChanged(this, textureFilePacked); };
	void Signal_VisibleChanged(bool visible) { callbackQCustom3DItem_VisibleChanged(this, visible); };
	 ~MyQCustom3DItem() { callbackQCustom3DItem_DestroyQCustom3DItem(this); };
};

void* QCustom3DItem_NewQCustom3DItem(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(static_cast<QWindow*>(parent));
	} else {
		return new MyQCustom3DItem(static_cast<QObject*>(parent));
	}
}

void* QCustom3DItem_NewQCustom3DItem2(struct QtDataVisualization_PackedString meshFile, void* position, void* scaling, void* rotation, void* texture, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QWindow*>(parent));
	} else {
		return new MyQCustom3DItem(QString::fromUtf8(meshFile.data, meshFile.len), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), *static_cast<QImage*>(texture), static_cast<QObject*>(parent));
	}
}

void* QCustom3DItem_Rotation(void* ptr)
{
	return new QQuaternion(static_cast<QCustom3DItem*>(ptr)->rotation());
}

void QCustom3DItem_ConnectMeshFileChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QString &)>(&QCustom3DItem::meshFileChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QString &)>(&MyQCustom3DItem::Signal_MeshFileChanged));
}

void QCustom3DItem_DisconnectMeshFileChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QString &)>(&QCustom3DItem::meshFileChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QString &)>(&MyQCustom3DItem::Signal_MeshFileChanged));
}

void QCustom3DItem_MeshFileChanged(void* ptr, struct QtDataVisualization_PackedString meshFile)
{
	static_cast<QCustom3DItem*>(ptr)->meshFileChanged(QString::fromUtf8(meshFile.data, meshFile.len));
}

void QCustom3DItem_ConnectPositionAbsoluteChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(bool)>(&QCustom3DItem::positionAbsoluteChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(bool)>(&MyQCustom3DItem::Signal_PositionAbsoluteChanged));
}

void QCustom3DItem_DisconnectPositionAbsoluteChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(bool)>(&QCustom3DItem::positionAbsoluteChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(bool)>(&MyQCustom3DItem::Signal_PositionAbsoluteChanged));
}

void QCustom3DItem_PositionAbsoluteChanged(void* ptr, char positionAbsolute)
{
	static_cast<QCustom3DItem*>(ptr)->positionAbsoluteChanged(positionAbsolute != 0);
}

void QCustom3DItem_ConnectPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QVector3D &)>(&QCustom3DItem::positionChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QVector3D &)>(&MyQCustom3DItem::Signal_PositionChanged));
}

void QCustom3DItem_DisconnectPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QVector3D &)>(&QCustom3DItem::positionChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QVector3D &)>(&MyQCustom3DItem::Signal_PositionChanged));
}

void QCustom3DItem_PositionChanged(void* ptr, void* position)
{
	static_cast<QCustom3DItem*>(ptr)->positionChanged(*static_cast<QVector3D*>(position));
}

void QCustom3DItem_ConnectRotationChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QQuaternion &)>(&QCustom3DItem::rotationChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QQuaternion &)>(&MyQCustom3DItem::Signal_RotationChanged));
}

void QCustom3DItem_DisconnectRotationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QQuaternion &)>(&QCustom3DItem::rotationChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QQuaternion &)>(&MyQCustom3DItem::Signal_RotationChanged));
}

void QCustom3DItem_RotationChanged(void* ptr, void* rotation)
{
	static_cast<QCustom3DItem*>(ptr)->rotationChanged(*static_cast<QQuaternion*>(rotation));
}

void QCustom3DItem_ConnectScalingAbsoluteChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(bool)>(&QCustom3DItem::scalingAbsoluteChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(bool)>(&MyQCustom3DItem::Signal_ScalingAbsoluteChanged));
}

void QCustom3DItem_DisconnectScalingAbsoluteChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(bool)>(&QCustom3DItem::scalingAbsoluteChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(bool)>(&MyQCustom3DItem::Signal_ScalingAbsoluteChanged));
}

void QCustom3DItem_ScalingAbsoluteChanged(void* ptr, char scalingAbsolute)
{
	static_cast<QCustom3DItem*>(ptr)->scalingAbsoluteChanged(scalingAbsolute != 0);
}

void QCustom3DItem_ConnectScalingChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QVector3D &)>(&QCustom3DItem::scalingChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QVector3D &)>(&MyQCustom3DItem::Signal_ScalingChanged));
}

void QCustom3DItem_DisconnectScalingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QVector3D &)>(&QCustom3DItem::scalingChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QVector3D &)>(&MyQCustom3DItem::Signal_ScalingChanged));
}

void QCustom3DItem_ScalingChanged(void* ptr, void* scaling)
{
	static_cast<QCustom3DItem*>(ptr)->scalingChanged(*static_cast<QVector3D*>(scaling));
}

void QCustom3DItem_SetMeshFile(void* ptr, struct QtDataVisualization_PackedString meshFile)
{
	static_cast<QCustom3DItem*>(ptr)->setMeshFile(QString::fromUtf8(meshFile.data, meshFile.len));
}

void QCustom3DItem_SetPosition(void* ptr, void* position)
{
	static_cast<QCustom3DItem*>(ptr)->setPosition(*static_cast<QVector3D*>(position));
}

void QCustom3DItem_SetPositionAbsolute(void* ptr, char positionAbsolute)
{
	static_cast<QCustom3DItem*>(ptr)->setPositionAbsolute(positionAbsolute != 0);
}

void QCustom3DItem_SetRotation(void* ptr, void* rotation)
{
	static_cast<QCustom3DItem*>(ptr)->setRotation(*static_cast<QQuaternion*>(rotation));
}

void QCustom3DItem_SetRotationAxisAndAngle(void* ptr, void* axis, float angle)
{
	static_cast<QCustom3DItem*>(ptr)->setRotationAxisAndAngle(*static_cast<QVector3D*>(axis), angle);
}

void QCustom3DItem_SetScaling(void* ptr, void* scaling)
{
	static_cast<QCustom3DItem*>(ptr)->setScaling(*static_cast<QVector3D*>(scaling));
}

void QCustom3DItem_SetScalingAbsolute(void* ptr, char scalingAbsolute)
{
	static_cast<QCustom3DItem*>(ptr)->setScalingAbsolute(scalingAbsolute != 0);
}

void QCustom3DItem_SetShadowCasting(void* ptr, char enabled)
{
	static_cast<QCustom3DItem*>(ptr)->setShadowCasting(enabled != 0);
}

void QCustom3DItem_SetTextureFile(void* ptr, struct QtDataVisualization_PackedString textureFile)
{
	static_cast<QCustom3DItem*>(ptr)->setTextureFile(QString::fromUtf8(textureFile.data, textureFile.len));
}

void QCustom3DItem_SetTextureImage(void* ptr, void* textureImage)
{
	static_cast<QCustom3DItem*>(ptr)->setTextureImage(*static_cast<QImage*>(textureImage));
}

void QCustom3DItem_SetVisible(void* ptr, char visible)
{
	static_cast<QCustom3DItem*>(ptr)->setVisible(visible != 0);
}

void QCustom3DItem_ConnectShadowCastingChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(bool)>(&QCustom3DItem::shadowCastingChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(bool)>(&MyQCustom3DItem::Signal_ShadowCastingChanged));
}

void QCustom3DItem_DisconnectShadowCastingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(bool)>(&QCustom3DItem::shadowCastingChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(bool)>(&MyQCustom3DItem::Signal_ShadowCastingChanged));
}

void QCustom3DItem_ShadowCastingChanged(void* ptr, char shadowCasting)
{
	static_cast<QCustom3DItem*>(ptr)->shadowCastingChanged(shadowCasting != 0);
}

void QCustom3DItem_ConnectTextureFileChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QString &)>(&QCustom3DItem::textureFileChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QString &)>(&MyQCustom3DItem::Signal_TextureFileChanged));
}

void QCustom3DItem_DisconnectTextureFileChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(const QString &)>(&QCustom3DItem::textureFileChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(const QString &)>(&MyQCustom3DItem::Signal_TextureFileChanged));
}

void QCustom3DItem_TextureFileChanged(void* ptr, struct QtDataVisualization_PackedString textureFile)
{
	static_cast<QCustom3DItem*>(ptr)->textureFileChanged(QString::fromUtf8(textureFile.data, textureFile.len));
}

void QCustom3DItem_ConnectVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(bool)>(&QCustom3DItem::visibleChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(bool)>(&MyQCustom3DItem::Signal_VisibleChanged));
}

void QCustom3DItem_DisconnectVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DItem*>(ptr), static_cast<void (QCustom3DItem::*)(bool)>(&QCustom3DItem::visibleChanged), static_cast<MyQCustom3DItem*>(ptr), static_cast<void (MyQCustom3DItem::*)(bool)>(&MyQCustom3DItem::Signal_VisibleChanged));
}

void QCustom3DItem_VisibleChanged(void* ptr, char visible)
{
	static_cast<QCustom3DItem*>(ptr)->visibleChanged(visible != 0);
}

void QCustom3DItem_DestroyQCustom3DItem(void* ptr)
{
	static_cast<QCustom3DItem*>(ptr)->~QCustom3DItem();
}

void QCustom3DItem_DestroyQCustom3DItemDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDataVisualization_PackedString QCustom3DItem_MeshFile(void* ptr)
{
	return ({ QByteArray t3766ee = static_cast<QCustom3DItem*>(ptr)->meshFile().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t3766ee.prepend("WHITESPACE").constData()+10), t3766ee.size()-10 }; });
}

struct QtDataVisualization_PackedString QCustom3DItem_TextureFile(void* ptr)
{
	return ({ QByteArray t0c954e = static_cast<QCustom3DItem*>(ptr)->textureFile().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t0c954e.prepend("WHITESPACE").constData()+10), t0c954e.size()-10 }; });
}

void* QCustom3DItem_Position(void* ptr)
{
	return new QVector3D(static_cast<QCustom3DItem*>(ptr)->position());
}

void* QCustom3DItem_Scaling(void* ptr)
{
	return new QVector3D(static_cast<QCustom3DItem*>(ptr)->scaling());
}

char QCustom3DItem_IsPositionAbsolute(void* ptr)
{
	return static_cast<QCustom3DItem*>(ptr)->isPositionAbsolute();
}

char QCustom3DItem_IsScalingAbsolute(void* ptr)
{
	return static_cast<QCustom3DItem*>(ptr)->isScalingAbsolute();
}

char QCustom3DItem_IsShadowCasting(void* ptr)
{
	return static_cast<QCustom3DItem*>(ptr)->isShadowCasting();
}

char QCustom3DItem_IsVisible(void* ptr)
{
	return static_cast<QCustom3DItem*>(ptr)->isVisible();
}

class MyQCustom3DLabel: public QCustom3DLabel
{
public:
	MyQCustom3DLabel(QObject *parent = Q_NULLPTR) : QCustom3DLabel(parent) {};
	MyQCustom3DLabel(const QString &text, const QFont &font, const QVector3D &position, const QVector3D &scaling, const QQuaternion &rotation, QObject *parent = Q_NULLPTR) : QCustom3DLabel(text, font, position, scaling, rotation, parent) {};
	void Signal_BackgroundColorChanged(const QColor & color) { callbackQCustom3DLabel_BackgroundColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_BackgroundEnabledChanged(bool enabled) { callbackQCustom3DLabel_BackgroundEnabledChanged(this, enabled); };
	void Signal_BorderEnabledChanged(bool enabled) { callbackQCustom3DLabel_BorderEnabledChanged(this, enabled); };
	void Signal_FacingCameraChanged(bool enabled) { callbackQCustom3DLabel_FacingCameraChanged(this, enabled); };
	void Signal_FontChanged(const QFont & font) { callbackQCustom3DLabel_FontChanged(this, const_cast<QFont*>(&font)); };
	void Signal_TextChanged(const QString & text) { QByteArray t372ea0 = text.toUtf8(); QtDataVisualization_PackedString textPacked = { const_cast<char*>(t372ea0.prepend("WHITESPACE").constData()+10), t372ea0.size()-10 };callbackQCustom3DLabel_TextChanged(this, textPacked); };
	void Signal_TextColorChanged(const QColor & color) { callbackQCustom3DLabel_TextColorChanged(this, const_cast<QColor*>(&color)); };
	 ~MyQCustom3DLabel() { callbackQCustom3DLabel_DestroyQCustom3DLabel(this); };
};

void* QCustom3DLabel_NewQCustom3DLabel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(static_cast<QWindow*>(parent));
	} else {
		return new MyQCustom3DLabel(static_cast<QObject*>(parent));
	}
}

void* QCustom3DLabel_NewQCustom3DLabel2(struct QtDataVisualization_PackedString text, void* font, void* position, void* scaling, void* rotation, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QWindow*>(parent));
	} else {
		return new MyQCustom3DLabel(QString::fromUtf8(text.data, text.len), *static_cast<QFont*>(font), *static_cast<QVector3D*>(position), *static_cast<QVector3D*>(scaling), *static_cast<QQuaternion*>(rotation), static_cast<QObject*>(parent));
	}
}

void QCustom3DLabel_ConnectBackgroundColorChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(const QColor &)>(&QCustom3DLabel::backgroundColorChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(const QColor &)>(&MyQCustom3DLabel::Signal_BackgroundColorChanged));
}

void QCustom3DLabel_DisconnectBackgroundColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(const QColor &)>(&QCustom3DLabel::backgroundColorChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(const QColor &)>(&MyQCustom3DLabel::Signal_BackgroundColorChanged));
}

void QCustom3DLabel_BackgroundColorChanged(void* ptr, void* color)
{
	static_cast<QCustom3DLabel*>(ptr)->backgroundColorChanged(*static_cast<QColor*>(color));
}

void QCustom3DLabel_ConnectBackgroundEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(bool)>(&QCustom3DLabel::backgroundEnabledChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(bool)>(&MyQCustom3DLabel::Signal_BackgroundEnabledChanged));
}

void QCustom3DLabel_DisconnectBackgroundEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(bool)>(&QCustom3DLabel::backgroundEnabledChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(bool)>(&MyQCustom3DLabel::Signal_BackgroundEnabledChanged));
}

void QCustom3DLabel_BackgroundEnabledChanged(void* ptr, char enabled)
{
	static_cast<QCustom3DLabel*>(ptr)->backgroundEnabledChanged(enabled != 0);
}

void QCustom3DLabel_ConnectBorderEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(bool)>(&QCustom3DLabel::borderEnabledChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(bool)>(&MyQCustom3DLabel::Signal_BorderEnabledChanged));
}

void QCustom3DLabel_DisconnectBorderEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(bool)>(&QCustom3DLabel::borderEnabledChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(bool)>(&MyQCustom3DLabel::Signal_BorderEnabledChanged));
}

void QCustom3DLabel_BorderEnabledChanged(void* ptr, char enabled)
{
	static_cast<QCustom3DLabel*>(ptr)->borderEnabledChanged(enabled != 0);
}

void QCustom3DLabel_ConnectFacingCameraChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(bool)>(&QCustom3DLabel::facingCameraChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(bool)>(&MyQCustom3DLabel::Signal_FacingCameraChanged));
}

void QCustom3DLabel_DisconnectFacingCameraChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(bool)>(&QCustom3DLabel::facingCameraChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(bool)>(&MyQCustom3DLabel::Signal_FacingCameraChanged));
}

void QCustom3DLabel_FacingCameraChanged(void* ptr, char enabled)
{
	static_cast<QCustom3DLabel*>(ptr)->facingCameraChanged(enabled != 0);
}

void QCustom3DLabel_ConnectFontChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(const QFont &)>(&QCustom3DLabel::fontChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(const QFont &)>(&MyQCustom3DLabel::Signal_FontChanged));
}

void QCustom3DLabel_DisconnectFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(const QFont &)>(&QCustom3DLabel::fontChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(const QFont &)>(&MyQCustom3DLabel::Signal_FontChanged));
}

void QCustom3DLabel_FontChanged(void* ptr, void* font)
{
	static_cast<QCustom3DLabel*>(ptr)->fontChanged(*static_cast<QFont*>(font));
}

void QCustom3DLabel_SetBackgroundColor(void* ptr, void* color)
{
	static_cast<QCustom3DLabel*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void QCustom3DLabel_SetBackgroundEnabled(void* ptr, char enabled)
{
	static_cast<QCustom3DLabel*>(ptr)->setBackgroundEnabled(enabled != 0);
}

void QCustom3DLabel_SetBorderEnabled(void* ptr, char enabled)
{
	static_cast<QCustom3DLabel*>(ptr)->setBorderEnabled(enabled != 0);
}

void QCustom3DLabel_SetFacingCamera(void* ptr, char enabled)
{
	static_cast<QCustom3DLabel*>(ptr)->setFacingCamera(enabled != 0);
}

void QCustom3DLabel_SetFont(void* ptr, void* font)
{
	static_cast<QCustom3DLabel*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QCustom3DLabel_SetText(void* ptr, struct QtDataVisualization_PackedString text)
{
	static_cast<QCustom3DLabel*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

void QCustom3DLabel_SetTextColor(void* ptr, void* color)
{
	static_cast<QCustom3DLabel*>(ptr)->setTextColor(*static_cast<QColor*>(color));
}

void QCustom3DLabel_ConnectTextChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(const QString &)>(&QCustom3DLabel::textChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(const QString &)>(&MyQCustom3DLabel::Signal_TextChanged));
}

void QCustom3DLabel_DisconnectTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(const QString &)>(&QCustom3DLabel::textChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(const QString &)>(&MyQCustom3DLabel::Signal_TextChanged));
}

void QCustom3DLabel_TextChanged(void* ptr, struct QtDataVisualization_PackedString text)
{
	static_cast<QCustom3DLabel*>(ptr)->textChanged(QString::fromUtf8(text.data, text.len));
}

void QCustom3DLabel_ConnectTextColorChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(const QColor &)>(&QCustom3DLabel::textColorChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(const QColor &)>(&MyQCustom3DLabel::Signal_TextColorChanged));
}

void QCustom3DLabel_DisconnectTextColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DLabel*>(ptr), static_cast<void (QCustom3DLabel::*)(const QColor &)>(&QCustom3DLabel::textColorChanged), static_cast<MyQCustom3DLabel*>(ptr), static_cast<void (MyQCustom3DLabel::*)(const QColor &)>(&MyQCustom3DLabel::Signal_TextColorChanged));
}

void QCustom3DLabel_TextColorChanged(void* ptr, void* color)
{
	static_cast<QCustom3DLabel*>(ptr)->textColorChanged(*static_cast<QColor*>(color));
}

void QCustom3DLabel_DestroyQCustom3DLabel(void* ptr)
{
	static_cast<QCustom3DLabel*>(ptr)->~QCustom3DLabel();
}

void QCustom3DLabel_DestroyQCustom3DLabelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QCustom3DLabel_BackgroundColor(void* ptr)
{
	return new QColor(static_cast<QCustom3DLabel*>(ptr)->backgroundColor());
}

void* QCustom3DLabel_TextColor(void* ptr)
{
	return new QColor(static_cast<QCustom3DLabel*>(ptr)->textColor());
}

void* QCustom3DLabel_Font(void* ptr)
{
	return new QFont(static_cast<QCustom3DLabel*>(ptr)->font());
}

struct QtDataVisualization_PackedString QCustom3DLabel_Text(void* ptr)
{
	return ({ QByteArray ta11a00 = static_cast<QCustom3DLabel*>(ptr)->text().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(ta11a00.prepend("WHITESPACE").constData()+10), ta11a00.size()-10 }; });
}

char QCustom3DLabel_IsBackgroundEnabled(void* ptr)
{
	return static_cast<QCustom3DLabel*>(ptr)->isBackgroundEnabled();
}

char QCustom3DLabel_IsBorderEnabled(void* ptr)
{
	return static_cast<QCustom3DLabel*>(ptr)->isBorderEnabled();
}

char QCustom3DLabel_IsFacingCamera(void* ptr)
{
	return static_cast<QCustom3DLabel*>(ptr)->isFacingCamera();
}

class MyQCustom3DVolume: public QCustom3DVolume
{
public:
	MyQCustom3DVolume(QObject *parent = Q_NULLPTR) : QCustom3DVolume(parent) {};
	void Signal_AlphaMultiplierChanged(float mult) { callbackQCustom3DVolume_AlphaMultiplierChanged(this, mult); };
	void Signal_ColorTableChanged() { callbackQCustom3DVolume_ColorTableChanged(this); };
	void Signal_DrawSliceFramesChanged(bool enabled) { callbackQCustom3DVolume_DrawSliceFramesChanged(this, enabled); };
	void Signal_DrawSlicesChanged(bool enabled) { callbackQCustom3DVolume_DrawSlicesChanged(this, enabled); };
	void Signal_PreserveOpacityChanged(bool enabled) { callbackQCustom3DVolume_PreserveOpacityChanged(this, enabled); };
	void Signal_SliceFrameColorChanged(const QColor & color) { callbackQCustom3DVolume_SliceFrameColorChanged(this, const_cast<QColor*>(&color)); };
	void Signal_SliceFrameGapsChanged(const QVector3D & values) { callbackQCustom3DVolume_SliceFrameGapsChanged(this, const_cast<QVector3D*>(&values)); };
	void Signal_SliceFrameThicknessesChanged(const QVector3D & values) { callbackQCustom3DVolume_SliceFrameThicknessesChanged(this, const_cast<QVector3D*>(&values)); };
	void Signal_SliceFrameWidthsChanged(const QVector3D & values) { callbackQCustom3DVolume_SliceFrameWidthsChanged(this, const_cast<QVector3D*>(&values)); };
	void Signal_SliceIndexXChanged(int value) { callbackQCustom3DVolume_SliceIndexXChanged(this, value); };
	void Signal_SliceIndexYChanged(int value) { callbackQCustom3DVolume_SliceIndexYChanged(this, value); };
	void Signal_SliceIndexZChanged(int value) { callbackQCustom3DVolume_SliceIndexZChanged(this, value); };
	void Signal_TextureDepthChanged(int value) { callbackQCustom3DVolume_TextureDepthChanged(this, value); };
	void Signal_TextureFormatChanged(QImage::Format format) { callbackQCustom3DVolume_TextureFormatChanged(this, format); };
	void Signal_TextureHeightChanged(int value) { callbackQCustom3DVolume_TextureHeightChanged(this, value); };
	void Signal_TextureWidthChanged(int value) { callbackQCustom3DVolume_TextureWidthChanged(this, value); };
	void Signal_UseHighDefShaderChanged(bool enabled) { callbackQCustom3DVolume_UseHighDefShaderChanged(this, enabled); };
	 ~MyQCustom3DVolume() { callbackQCustom3DVolume_DestroyQCustom3DVolume(this); };
};

void* QCustom3DVolume_NewQCustom3DVolume(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCustom3DVolume(static_cast<QWindow*>(parent));
	} else {
		return new MyQCustom3DVolume(static_cast<QObject*>(parent));
	}
}

void* QCustom3DVolume_RenderSlice(void* ptr, long long axis, int index)
{
	return new QImage(static_cast<QCustom3DVolume*>(ptr)->renderSlice(static_cast<Qt::Axis>(axis), index));
}

void QCustom3DVolume_ConnectAlphaMultiplierChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(float)>(&QCustom3DVolume::alphaMultiplierChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(float)>(&MyQCustom3DVolume::Signal_AlphaMultiplierChanged));
}

void QCustom3DVolume_DisconnectAlphaMultiplierChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(float)>(&QCustom3DVolume::alphaMultiplierChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(float)>(&MyQCustom3DVolume::Signal_AlphaMultiplierChanged));
}

void QCustom3DVolume_AlphaMultiplierChanged(void* ptr, float mult)
{
	static_cast<QCustom3DVolume*>(ptr)->alphaMultiplierChanged(mult);
}

void QCustom3DVolume_ConnectColorTableChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)()>(&QCustom3DVolume::colorTableChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)()>(&MyQCustom3DVolume::Signal_ColorTableChanged));
}

void QCustom3DVolume_DisconnectColorTableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)()>(&QCustom3DVolume::colorTableChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)()>(&MyQCustom3DVolume::Signal_ColorTableChanged));
}

void QCustom3DVolume_ColorTableChanged(void* ptr)
{
	static_cast<QCustom3DVolume*>(ptr)->colorTableChanged();
}

void QCustom3DVolume_ConnectDrawSliceFramesChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(bool)>(&QCustom3DVolume::drawSliceFramesChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(bool)>(&MyQCustom3DVolume::Signal_DrawSliceFramesChanged));
}

void QCustom3DVolume_DisconnectDrawSliceFramesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(bool)>(&QCustom3DVolume::drawSliceFramesChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(bool)>(&MyQCustom3DVolume::Signal_DrawSliceFramesChanged));
}

void QCustom3DVolume_DrawSliceFramesChanged(void* ptr, char enabled)
{
	static_cast<QCustom3DVolume*>(ptr)->drawSliceFramesChanged(enabled != 0);
}

void QCustom3DVolume_ConnectDrawSlicesChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(bool)>(&QCustom3DVolume::drawSlicesChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(bool)>(&MyQCustom3DVolume::Signal_DrawSlicesChanged));
}

void QCustom3DVolume_DisconnectDrawSlicesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(bool)>(&QCustom3DVolume::drawSlicesChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(bool)>(&MyQCustom3DVolume::Signal_DrawSlicesChanged));
}

void QCustom3DVolume_DrawSlicesChanged(void* ptr, char enabled)
{
	static_cast<QCustom3DVolume*>(ptr)->drawSlicesChanged(enabled != 0);
}

void QCustom3DVolume_ConnectPreserveOpacityChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(bool)>(&QCustom3DVolume::preserveOpacityChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(bool)>(&MyQCustom3DVolume::Signal_PreserveOpacityChanged));
}

void QCustom3DVolume_DisconnectPreserveOpacityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(bool)>(&QCustom3DVolume::preserveOpacityChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(bool)>(&MyQCustom3DVolume::Signal_PreserveOpacityChanged));
}

void QCustom3DVolume_PreserveOpacityChanged(void* ptr, char enabled)
{
	static_cast<QCustom3DVolume*>(ptr)->preserveOpacityChanged(enabled != 0);
}

void QCustom3DVolume_SetAlphaMultiplier(void* ptr, float mult)
{
	static_cast<QCustom3DVolume*>(ptr)->setAlphaMultiplier(mult);
}

void QCustom3DVolume_SetDrawSliceFrames(void* ptr, char enable)
{
	static_cast<QCustom3DVolume*>(ptr)->setDrawSliceFrames(enable != 0);
}

void QCustom3DVolume_SetDrawSlices(void* ptr, char enable)
{
	static_cast<QCustom3DVolume*>(ptr)->setDrawSlices(enable != 0);
}

void QCustom3DVolume_SetPreserveOpacity(void* ptr, char enable)
{
	static_cast<QCustom3DVolume*>(ptr)->setPreserveOpacity(enable != 0);
}

void QCustom3DVolume_SetSliceFrameColor(void* ptr, void* color)
{
	static_cast<QCustom3DVolume*>(ptr)->setSliceFrameColor(*static_cast<QColor*>(color));
}

void QCustom3DVolume_SetSliceFrameGaps(void* ptr, void* values)
{
	static_cast<QCustom3DVolume*>(ptr)->setSliceFrameGaps(*static_cast<QVector3D*>(values));
}

void QCustom3DVolume_SetSliceFrameThicknesses(void* ptr, void* values)
{
	static_cast<QCustom3DVolume*>(ptr)->setSliceFrameThicknesses(*static_cast<QVector3D*>(values));
}

void QCustom3DVolume_SetSliceFrameWidths(void* ptr, void* values)
{
	static_cast<QCustom3DVolume*>(ptr)->setSliceFrameWidths(*static_cast<QVector3D*>(values));
}

void QCustom3DVolume_SetSliceIndexX(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->setSliceIndexX(value);
}

void QCustom3DVolume_SetSliceIndexY(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->setSliceIndexY(value);
}

void QCustom3DVolume_SetSliceIndexZ(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->setSliceIndexZ(value);
}

void QCustom3DVolume_SetSliceIndices(void* ptr, int x, int y, int z)
{
	static_cast<QCustom3DVolume*>(ptr)->setSliceIndices(x, y, z);
}

void QCustom3DVolume_SetSubTextureData2(void* ptr, long long axis, int index, void* image)
{
	static_cast<QCustom3DVolume*>(ptr)->setSubTextureData(static_cast<Qt::Axis>(axis), index, *static_cast<QImage*>(image));
}

void QCustom3DVolume_SetSubTextureData(void* ptr, long long axis, int index, char* data)
{
	static_cast<QCustom3DVolume*>(ptr)->setSubTextureData(static_cast<Qt::Axis>(axis), index, const_cast<const uchar*>(static_cast<uchar*>(static_cast<void*>(data))));
}

void QCustom3DVolume_SetTextureDepth(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->setTextureDepth(value);
}

void QCustom3DVolume_SetTextureDimensions(void* ptr, int width, int height, int depth)
{
	static_cast<QCustom3DVolume*>(ptr)->setTextureDimensions(width, height, depth);
}

void QCustom3DVolume_SetTextureFormat(void* ptr, long long format)
{
	static_cast<QCustom3DVolume*>(ptr)->setTextureFormat(static_cast<QImage::Format>(format));
}

void QCustom3DVolume_SetTextureHeight(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->setTextureHeight(value);
}

void QCustom3DVolume_SetTextureWidth(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->setTextureWidth(value);
}

void QCustom3DVolume_SetUseHighDefShader(void* ptr, char enable)
{
	static_cast<QCustom3DVolume*>(ptr)->setUseHighDefShader(enable != 0);
}

void QCustom3DVolume_ConnectSliceFrameColorChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(const QColor &)>(&QCustom3DVolume::sliceFrameColorChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(const QColor &)>(&MyQCustom3DVolume::Signal_SliceFrameColorChanged));
}

void QCustom3DVolume_DisconnectSliceFrameColorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(const QColor &)>(&QCustom3DVolume::sliceFrameColorChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(const QColor &)>(&MyQCustom3DVolume::Signal_SliceFrameColorChanged));
}

void QCustom3DVolume_SliceFrameColorChanged(void* ptr, void* color)
{
	static_cast<QCustom3DVolume*>(ptr)->sliceFrameColorChanged(*static_cast<QColor*>(color));
}

void QCustom3DVolume_ConnectSliceFrameGapsChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(const QVector3D &)>(&QCustom3DVolume::sliceFrameGapsChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(const QVector3D &)>(&MyQCustom3DVolume::Signal_SliceFrameGapsChanged));
}

void QCustom3DVolume_DisconnectSliceFrameGapsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(const QVector3D &)>(&QCustom3DVolume::sliceFrameGapsChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(const QVector3D &)>(&MyQCustom3DVolume::Signal_SliceFrameGapsChanged));
}

void QCustom3DVolume_SliceFrameGapsChanged(void* ptr, void* values)
{
	static_cast<QCustom3DVolume*>(ptr)->sliceFrameGapsChanged(*static_cast<QVector3D*>(values));
}

void QCustom3DVolume_ConnectSliceFrameThicknessesChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(const QVector3D &)>(&QCustom3DVolume::sliceFrameThicknessesChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(const QVector3D &)>(&MyQCustom3DVolume::Signal_SliceFrameThicknessesChanged));
}

void QCustom3DVolume_DisconnectSliceFrameThicknessesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(const QVector3D &)>(&QCustom3DVolume::sliceFrameThicknessesChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(const QVector3D &)>(&MyQCustom3DVolume::Signal_SliceFrameThicknessesChanged));
}

void QCustom3DVolume_SliceFrameThicknessesChanged(void* ptr, void* values)
{
	static_cast<QCustom3DVolume*>(ptr)->sliceFrameThicknessesChanged(*static_cast<QVector3D*>(values));
}

void QCustom3DVolume_ConnectSliceFrameWidthsChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(const QVector3D &)>(&QCustom3DVolume::sliceFrameWidthsChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(const QVector3D &)>(&MyQCustom3DVolume::Signal_SliceFrameWidthsChanged));
}

void QCustom3DVolume_DisconnectSliceFrameWidthsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(const QVector3D &)>(&QCustom3DVolume::sliceFrameWidthsChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(const QVector3D &)>(&MyQCustom3DVolume::Signal_SliceFrameWidthsChanged));
}

void QCustom3DVolume_SliceFrameWidthsChanged(void* ptr, void* values)
{
	static_cast<QCustom3DVolume*>(ptr)->sliceFrameWidthsChanged(*static_cast<QVector3D*>(values));
}

void QCustom3DVolume_ConnectSliceIndexXChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::sliceIndexXChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_SliceIndexXChanged));
}

void QCustom3DVolume_DisconnectSliceIndexXChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::sliceIndexXChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_SliceIndexXChanged));
}

void QCustom3DVolume_SliceIndexXChanged(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->sliceIndexXChanged(value);
}

void QCustom3DVolume_ConnectSliceIndexYChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::sliceIndexYChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_SliceIndexYChanged));
}

void QCustom3DVolume_DisconnectSliceIndexYChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::sliceIndexYChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_SliceIndexYChanged));
}

void QCustom3DVolume_SliceIndexYChanged(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->sliceIndexYChanged(value);
}

void QCustom3DVolume_ConnectSliceIndexZChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::sliceIndexZChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_SliceIndexZChanged));
}

void QCustom3DVolume_DisconnectSliceIndexZChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::sliceIndexZChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_SliceIndexZChanged));
}

void QCustom3DVolume_SliceIndexZChanged(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->sliceIndexZChanged(value);
}

void QCustom3DVolume_ConnectTextureDepthChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::textureDepthChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_TextureDepthChanged));
}

void QCustom3DVolume_DisconnectTextureDepthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::textureDepthChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_TextureDepthChanged));
}

void QCustom3DVolume_TextureDepthChanged(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->textureDepthChanged(value);
}

void QCustom3DVolume_ConnectTextureFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(QImage::Format)>(&QCustom3DVolume::textureFormatChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(QImage::Format)>(&MyQCustom3DVolume::Signal_TextureFormatChanged));
}

void QCustom3DVolume_DisconnectTextureFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(QImage::Format)>(&QCustom3DVolume::textureFormatChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(QImage::Format)>(&MyQCustom3DVolume::Signal_TextureFormatChanged));
}

void QCustom3DVolume_TextureFormatChanged(void* ptr, long long format)
{
	static_cast<QCustom3DVolume*>(ptr)->textureFormatChanged(static_cast<QImage::Format>(format));
}

void QCustom3DVolume_ConnectTextureHeightChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::textureHeightChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_TextureHeightChanged));
}

void QCustom3DVolume_DisconnectTextureHeightChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::textureHeightChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_TextureHeightChanged));
}

void QCustom3DVolume_TextureHeightChanged(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->textureHeightChanged(value);
}

void QCustom3DVolume_ConnectTextureWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::textureWidthChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_TextureWidthChanged));
}

void QCustom3DVolume_DisconnectTextureWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(int)>(&QCustom3DVolume::textureWidthChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(int)>(&MyQCustom3DVolume::Signal_TextureWidthChanged));
}

void QCustom3DVolume_TextureWidthChanged(void* ptr, int value)
{
	static_cast<QCustom3DVolume*>(ptr)->textureWidthChanged(value);
}

void QCustom3DVolume_ConnectUseHighDefShaderChanged(void* ptr)
{
	QObject::connect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(bool)>(&QCustom3DVolume::useHighDefShaderChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(bool)>(&MyQCustom3DVolume::Signal_UseHighDefShaderChanged));
}

void QCustom3DVolume_DisconnectUseHighDefShaderChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCustom3DVolume*>(ptr), static_cast<void (QCustom3DVolume::*)(bool)>(&QCustom3DVolume::useHighDefShaderChanged), static_cast<MyQCustom3DVolume*>(ptr), static_cast<void (MyQCustom3DVolume::*)(bool)>(&MyQCustom3DVolume::Signal_UseHighDefShaderChanged));
}

void QCustom3DVolume_UseHighDefShaderChanged(void* ptr, char enabled)
{
	static_cast<QCustom3DVolume*>(ptr)->useHighDefShaderChanged(enabled != 0);
}

void QCustom3DVolume_DestroyQCustom3DVolume(void* ptr)
{
	static_cast<QCustom3DVolume*>(ptr)->~QCustom3DVolume();
}

void QCustom3DVolume_DestroyQCustom3DVolumeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QCustom3DVolume_SliceFrameColor(void* ptr)
{
	return new QColor(static_cast<QCustom3DVolume*>(ptr)->sliceFrameColor());
}

long long QCustom3DVolume_TextureFormat(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->textureFormat();
}

void* QCustom3DVolume_SliceFrameGaps(void* ptr)
{
	return new QVector3D(static_cast<QCustom3DVolume*>(ptr)->sliceFrameGaps());
}

void* QCustom3DVolume_SliceFrameThicknesses(void* ptr)
{
	return new QVector3D(static_cast<QCustom3DVolume*>(ptr)->sliceFrameThicknesses());
}

void* QCustom3DVolume_SliceFrameWidths(void* ptr)
{
	return new QVector3D(static_cast<QCustom3DVolume*>(ptr)->sliceFrameWidths());
}

char QCustom3DVolume_DrawSliceFrames(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->drawSliceFrames();
}

char QCustom3DVolume_DrawSlices(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->drawSlices();
}

char QCustom3DVolume_PreserveOpacity(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->preserveOpacity();
}

char QCustom3DVolume_UseHighDefShader(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->useHighDefShader();
}

float QCustom3DVolume_AlphaMultiplier(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->alphaMultiplier();
}

int QCustom3DVolume_SliceIndexX(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->sliceIndexX();
}

int QCustom3DVolume_SliceIndexY(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->sliceIndexY();
}

int QCustom3DVolume_SliceIndexZ(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->sliceIndexZ();
}

int QCustom3DVolume_TextureDepth(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->textureDepth();
}

int QCustom3DVolume_TextureHeight(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->textureHeight();
}

int QCustom3DVolume_TextureWidth(void* ptr)
{
	return static_cast<QCustom3DVolume*>(ptr)->textureWidth();
}

void* QCustom3DVolume___QCustom3DVolume_colorTable_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRgb>;
}

void* QCustom3DVolume___setColorTable_colors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRgb>;
}

void* QCustom3DVolume___colorTable_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRgb>;
}

class MyQHeightMapSurfaceDataProxy: public QHeightMapSurfaceDataProxy
{
public:
	MyQHeightMapSurfaceDataProxy(QObject *parent = Q_NULLPTR) : QHeightMapSurfaceDataProxy(parent) {};
	MyQHeightMapSurfaceDataProxy(const QImage &image, QObject *parent = Q_NULLPTR) : QHeightMapSurfaceDataProxy(image, parent) {};
	MyQHeightMapSurfaceDataProxy(const QString &filename, QObject *parent = Q_NULLPTR) : QHeightMapSurfaceDataProxy(filename, parent) {};
	void Signal_HeightMapChanged(const QImage & image) { callbackQHeightMapSurfaceDataProxy_HeightMapChanged(this, const_cast<QImage*>(&image)); };
	void Signal_HeightMapFileChanged(const QString & filename) { QByteArray t08deae = filename.toUtf8(); QtDataVisualization_PackedString filenamePacked = { const_cast<char*>(t08deae.prepend("WHITESPACE").constData()+10), t08deae.size()-10 };callbackQHeightMapSurfaceDataProxy_HeightMapFileChanged(this, filenamePacked); };
	void Signal_MaxXValueChanged(float value) { callbackQHeightMapSurfaceDataProxy_MaxXValueChanged(this, value); };
	void Signal_MaxZValueChanged(float value) { callbackQHeightMapSurfaceDataProxy_MaxZValueChanged(this, value); };
	void Signal_MinXValueChanged(float value) { callbackQHeightMapSurfaceDataProxy_MinXValueChanged(this, value); };
	void Signal_MinZValueChanged(float value) { callbackQHeightMapSurfaceDataProxy_MinZValueChanged(this, value); };
	 ~MyQHeightMapSurfaceDataProxy() { callbackQHeightMapSurfaceDataProxy_DestroyQHeightMapSurfaceDataProxy(this); };
};

void* QHeightMapSurfaceDataProxy_NewQHeightMapSurfaceDataProxy(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QWindow*>(parent));
	} else {
		return new MyQHeightMapSurfaceDataProxy(static_cast<QObject*>(parent));
	}
}

void* QHeightMapSurfaceDataProxy_NewQHeightMapSurfaceDataProxy2(void* image, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QWindow*>(parent));
	} else {
		return new MyQHeightMapSurfaceDataProxy(*static_cast<QImage*>(image), static_cast<QObject*>(parent));
	}
}

void* QHeightMapSurfaceDataProxy_NewQHeightMapSurfaceDataProxy3(struct QtDataVisualization_PackedString filename, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQHeightMapSurfaceDataProxy(QString::fromUtf8(filename.data, filename.len), static_cast<QObject*>(parent));
	}
}

void QHeightMapSurfaceDataProxy_ConnectHeightMapChanged(void* ptr)
{
	QObject::connect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(const QImage &)>(&QHeightMapSurfaceDataProxy::heightMapChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(const QImage &)>(&MyQHeightMapSurfaceDataProxy::Signal_HeightMapChanged));
}

void QHeightMapSurfaceDataProxy_DisconnectHeightMapChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(const QImage &)>(&QHeightMapSurfaceDataProxy::heightMapChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(const QImage &)>(&MyQHeightMapSurfaceDataProxy::Signal_HeightMapChanged));
}

void QHeightMapSurfaceDataProxy_HeightMapChanged(void* ptr, void* image)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->heightMapChanged(*static_cast<QImage*>(image));
}

void QHeightMapSurfaceDataProxy_ConnectHeightMapFileChanged(void* ptr)
{
	QObject::connect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(const QString &)>(&QHeightMapSurfaceDataProxy::heightMapFileChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(const QString &)>(&MyQHeightMapSurfaceDataProxy::Signal_HeightMapFileChanged));
}

void QHeightMapSurfaceDataProxy_DisconnectHeightMapFileChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(const QString &)>(&QHeightMapSurfaceDataProxy::heightMapFileChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(const QString &)>(&MyQHeightMapSurfaceDataProxy::Signal_HeightMapFileChanged));
}

void QHeightMapSurfaceDataProxy_HeightMapFileChanged(void* ptr, struct QtDataVisualization_PackedString filename)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->heightMapFileChanged(QString::fromUtf8(filename.data, filename.len));
}

void QHeightMapSurfaceDataProxy_ConnectMaxXValueChanged(void* ptr)
{
	QObject::connect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(float)>(&QHeightMapSurfaceDataProxy::maxXValueChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(float)>(&MyQHeightMapSurfaceDataProxy::Signal_MaxXValueChanged));
}

void QHeightMapSurfaceDataProxy_DisconnectMaxXValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(float)>(&QHeightMapSurfaceDataProxy::maxXValueChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(float)>(&MyQHeightMapSurfaceDataProxy::Signal_MaxXValueChanged));
}

void QHeightMapSurfaceDataProxy_MaxXValueChanged(void* ptr, float value)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->maxXValueChanged(value);
}

void QHeightMapSurfaceDataProxy_ConnectMaxZValueChanged(void* ptr)
{
	QObject::connect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(float)>(&QHeightMapSurfaceDataProxy::maxZValueChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(float)>(&MyQHeightMapSurfaceDataProxy::Signal_MaxZValueChanged));
}

void QHeightMapSurfaceDataProxy_DisconnectMaxZValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(float)>(&QHeightMapSurfaceDataProxy::maxZValueChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(float)>(&MyQHeightMapSurfaceDataProxy::Signal_MaxZValueChanged));
}

void QHeightMapSurfaceDataProxy_MaxZValueChanged(void* ptr, float value)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->maxZValueChanged(value);
}

void QHeightMapSurfaceDataProxy_ConnectMinXValueChanged(void* ptr)
{
	QObject::connect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(float)>(&QHeightMapSurfaceDataProxy::minXValueChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(float)>(&MyQHeightMapSurfaceDataProxy::Signal_MinXValueChanged));
}

void QHeightMapSurfaceDataProxy_DisconnectMinXValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(float)>(&QHeightMapSurfaceDataProxy::minXValueChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(float)>(&MyQHeightMapSurfaceDataProxy::Signal_MinXValueChanged));
}

void QHeightMapSurfaceDataProxy_MinXValueChanged(void* ptr, float value)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->minXValueChanged(value);
}

void QHeightMapSurfaceDataProxy_ConnectMinZValueChanged(void* ptr)
{
	QObject::connect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(float)>(&QHeightMapSurfaceDataProxy::minZValueChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(float)>(&MyQHeightMapSurfaceDataProxy::Signal_MinZValueChanged));
}

void QHeightMapSurfaceDataProxy_DisconnectMinZValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHeightMapSurfaceDataProxy*>(ptr), static_cast<void (QHeightMapSurfaceDataProxy::*)(float)>(&QHeightMapSurfaceDataProxy::minZValueChanged), static_cast<MyQHeightMapSurfaceDataProxy*>(ptr), static_cast<void (MyQHeightMapSurfaceDataProxy::*)(float)>(&MyQHeightMapSurfaceDataProxy::Signal_MinZValueChanged));
}

void QHeightMapSurfaceDataProxy_MinZValueChanged(void* ptr, float value)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->minZValueChanged(value);
}

void QHeightMapSurfaceDataProxy_SetHeightMap(void* ptr, void* image)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->setHeightMap(*static_cast<QImage*>(image));
}

void QHeightMapSurfaceDataProxy_SetHeightMapFile(void* ptr, struct QtDataVisualization_PackedString filename)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->setHeightMapFile(QString::fromUtf8(filename.data, filename.len));
}

void QHeightMapSurfaceDataProxy_SetMaxXValue(void* ptr, float max)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->setMaxXValue(max);
}

void QHeightMapSurfaceDataProxy_SetMaxZValue(void* ptr, float max)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->setMaxZValue(max);
}

void QHeightMapSurfaceDataProxy_SetMinXValue(void* ptr, float min)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->setMinXValue(min);
}

void QHeightMapSurfaceDataProxy_SetMinZValue(void* ptr, float min)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->setMinZValue(min);
}

void QHeightMapSurfaceDataProxy_SetValueRanges(void* ptr, float minX, float maxX, float minZ, float maxZ)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->setValueRanges(minX, maxX, minZ, maxZ);
}

void QHeightMapSurfaceDataProxy_DestroyQHeightMapSurfaceDataProxy(void* ptr)
{
	static_cast<QHeightMapSurfaceDataProxy*>(ptr)->~QHeightMapSurfaceDataProxy();
}

void QHeightMapSurfaceDataProxy_DestroyQHeightMapSurfaceDataProxyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QHeightMapSurfaceDataProxy_HeightMap(void* ptr)
{
	return new QImage(static_cast<QHeightMapSurfaceDataProxy*>(ptr)->heightMap());
}

struct QtDataVisualization_PackedString QHeightMapSurfaceDataProxy_HeightMapFile(void* ptr)
{
	return ({ QByteArray t26dd34 = static_cast<QHeightMapSurfaceDataProxy*>(ptr)->heightMapFile().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t26dd34.prepend("WHITESPACE").constData()+10), t26dd34.size()-10 }; });
}

float QHeightMapSurfaceDataProxy_MaxXValue(void* ptr)
{
	return static_cast<QHeightMapSurfaceDataProxy*>(ptr)->maxXValue();
}

float QHeightMapSurfaceDataProxy_MaxZValue(void* ptr)
{
	return static_cast<QHeightMapSurfaceDataProxy*>(ptr)->maxZValue();
}

float QHeightMapSurfaceDataProxy_MinXValue(void* ptr)
{
	return static_cast<QHeightMapSurfaceDataProxy*>(ptr)->minXValue();
}

float QHeightMapSurfaceDataProxy_MinZValue(void* ptr)
{
	return static_cast<QHeightMapSurfaceDataProxy*>(ptr)->minZValue();
}

class MyQItemModelBarDataProxy: public QItemModelBarDataProxy
{
public:
	MyQItemModelBarDataProxy(QAbstractItemModel *itemModel, QObject *parent = Q_NULLPTR) : QItemModelBarDataProxy(itemModel, parent) {};
	MyQItemModelBarDataProxy(QAbstractItemModel *itemModel, const QString &rowRole, const QString &columnRole, const QString &valueRole, QObject *parent = Q_NULLPTR) : QItemModelBarDataProxy(itemModel, rowRole, columnRole, valueRole, parent) {};
	MyQItemModelBarDataProxy(QAbstractItemModel *itemModel, const QString &rowRole, const QString &columnRole, const QString &valueRole, const QString &rotationRole, QObject *parent = Q_NULLPTR) : QItemModelBarDataProxy(itemModel, rowRole, columnRole, valueRole, rotationRole, parent) {};
	MyQItemModelBarDataProxy(QAbstractItemModel *itemModel, const QString &rowRole, const QString &columnRole, const QString &valueRole, const QString &rotationRole, const QStringList &rowCategories, const QStringList &columnCategories, QObject *parent = Q_NULLPTR) : QItemModelBarDataProxy(itemModel, rowRole, columnRole, valueRole, rotationRole, rowCategories, columnCategories, parent) {};
	MyQItemModelBarDataProxy(QAbstractItemModel *itemModel, const QString &rowRole, const QString &columnRole, const QString &valueRole, const QStringList &rowCategories, const QStringList &columnCategories, QObject *parent = Q_NULLPTR) : QItemModelBarDataProxy(itemModel, rowRole, columnRole, valueRole, rowCategories, columnCategories, parent) {};
	MyQItemModelBarDataProxy(QAbstractItemModel *itemModel, const QString &valueRole, QObject *parent = Q_NULLPTR) : QItemModelBarDataProxy(itemModel, valueRole, parent) {};
	MyQItemModelBarDataProxy(QObject *parent = Q_NULLPTR) : QItemModelBarDataProxy(parent) {};
	void Signal_AutoColumnCategoriesChanged(bool enable) { callbackQItemModelBarDataProxy_AutoColumnCategoriesChanged(this, enable); };
	void Signal_AutoRowCategoriesChanged(bool enable) { callbackQItemModelBarDataProxy_AutoRowCategoriesChanged(this, enable); };
	void Signal_ColumnCategoriesChanged() { callbackQItemModelBarDataProxy_ColumnCategoriesChanged(this); };
	void Signal_ColumnRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelBarDataProxy_ColumnRoleChanged(this, rolePacked); };
	void Signal_ColumnRolePatternChanged(const QRegExp & pattern) { callbackQItemModelBarDataProxy_ColumnRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_ColumnRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelBarDataProxy_ColumnRoleReplaceChanged(this, replacePacked); };
	void Signal_ItemModelChanged(const QAbstractItemModel * itemModel) { callbackQItemModelBarDataProxy_ItemModelChanged(this, const_cast<QAbstractItemModel*>(itemModel)); };
	void Signal_MultiMatchBehaviorChanged(QItemModelBarDataProxy::MultiMatchBehavior behavior) { callbackQItemModelBarDataProxy_MultiMatchBehaviorChanged(this, behavior); };
	void Signal_RotationRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelBarDataProxy_RotationRoleChanged(this, rolePacked); };
	void Signal_RotationRolePatternChanged(const QRegExp & pattern) { callbackQItemModelBarDataProxy_RotationRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_RotationRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelBarDataProxy_RotationRoleReplaceChanged(this, replacePacked); };
	void Signal_RowCategoriesChanged() { callbackQItemModelBarDataProxy_RowCategoriesChanged(this); };
	void Signal_RowRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelBarDataProxy_RowRoleChanged(this, rolePacked); };
	void Signal_RowRolePatternChanged(const QRegExp & pattern) { callbackQItemModelBarDataProxy_RowRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_RowRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelBarDataProxy_RowRoleReplaceChanged(this, replacePacked); };
	void Signal_UseModelCategoriesChanged(bool enable) { callbackQItemModelBarDataProxy_UseModelCategoriesChanged(this, enable); };
	void Signal_ValueRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelBarDataProxy_ValueRoleChanged(this, rolePacked); };
	void Signal_ValueRolePatternChanged(const QRegExp & pattern) { callbackQItemModelBarDataProxy_ValueRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_ValueRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelBarDataProxy_ValueRoleReplaceChanged(this, replacePacked); };
	 ~MyQItemModelBarDataProxy() { callbackQItemModelBarDataProxy_DestroyQItemModelBarDataProxy(this); };
};

void* QItemModelBarDataProxy_NewQItemModelBarDataProxy2(void* itemModel, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QObject*>(parent));
	}
}

void* QItemModelBarDataProxy_NewQItemModelBarDataProxy4(void* itemModel, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString valueRole, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QObject*>(parent));
	}
}

void* QItemModelBarDataProxy_NewQItemModelBarDataProxy5(void* itemModel, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString valueRole, struct QtDataVisualization_PackedString rotationRole, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QObject*>(parent));
	}
}

void* QItemModelBarDataProxy_NewQItemModelBarDataProxy7(void* itemModel, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString valueRole, struct QtDataVisualization_PackedString rotationRole, struct QtDataVisualization_PackedString rowCategories, struct QtDataVisualization_PackedString columnCategories, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
	}
}

void* QItemModelBarDataProxy_NewQItemModelBarDataProxy6(void* itemModel, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString valueRole, struct QtDataVisualization_PackedString rowCategories, struct QtDataVisualization_PackedString columnCategories, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
	}
}

void* QItemModelBarDataProxy_NewQItemModelBarDataProxy3(void* itemModel, struct QtDataVisualization_PackedString valueRole, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelBarDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(valueRole.data, valueRole.len), static_cast<QObject*>(parent));
	}
}

void* QItemModelBarDataProxy_NewQItemModelBarDataProxy(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelBarDataProxy(static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelBarDataProxy(static_cast<QObject*>(parent));
	}
}

int QItemModelBarDataProxy_ColumnCategoryIndex(void* ptr, struct QtDataVisualization_PackedString category)
{
	return static_cast<QItemModelBarDataProxy*>(ptr)->columnCategoryIndex(QString::fromUtf8(category.data, category.len));
}

int QItemModelBarDataProxy_RowCategoryIndex(void* ptr, struct QtDataVisualization_PackedString category)
{
	return static_cast<QItemModelBarDataProxy*>(ptr)->rowCategoryIndex(QString::fromUtf8(category.data, category.len));
}

void QItemModelBarDataProxy_ConnectAutoColumnCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(bool)>(&QItemModelBarDataProxy::autoColumnCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(bool)>(&MyQItemModelBarDataProxy::Signal_AutoColumnCategoriesChanged));
}

void QItemModelBarDataProxy_DisconnectAutoColumnCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(bool)>(&QItemModelBarDataProxy::autoColumnCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(bool)>(&MyQItemModelBarDataProxy::Signal_AutoColumnCategoriesChanged));
}

void QItemModelBarDataProxy_AutoColumnCategoriesChanged(void* ptr, char enable)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->autoColumnCategoriesChanged(enable != 0);
}

void QItemModelBarDataProxy_ConnectAutoRowCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(bool)>(&QItemModelBarDataProxy::autoRowCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(bool)>(&MyQItemModelBarDataProxy::Signal_AutoRowCategoriesChanged));
}

void QItemModelBarDataProxy_DisconnectAutoRowCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(bool)>(&QItemModelBarDataProxy::autoRowCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(bool)>(&MyQItemModelBarDataProxy::Signal_AutoRowCategoriesChanged));
}

void QItemModelBarDataProxy_AutoRowCategoriesChanged(void* ptr, char enable)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->autoRowCategoriesChanged(enable != 0);
}

void QItemModelBarDataProxy_ConnectColumnCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)()>(&QItemModelBarDataProxy::columnCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)()>(&MyQItemModelBarDataProxy::Signal_ColumnCategoriesChanged));
}

void QItemModelBarDataProxy_DisconnectColumnCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)()>(&QItemModelBarDataProxy::columnCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)()>(&MyQItemModelBarDataProxy::Signal_ColumnCategoriesChanged));
}

void QItemModelBarDataProxy_ColumnCategoriesChanged(void* ptr)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->columnCategoriesChanged();
}

void QItemModelBarDataProxy_ConnectColumnRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::columnRoleChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_ColumnRoleChanged));
}

void QItemModelBarDataProxy_DisconnectColumnRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::columnRoleChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_ColumnRoleChanged));
}

void QItemModelBarDataProxy_ColumnRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->columnRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelBarDataProxy_ConnectColumnRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QRegExp &)>(&QItemModelBarDataProxy::columnRolePatternChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QRegExp &)>(&MyQItemModelBarDataProxy::Signal_ColumnRolePatternChanged));
}

void QItemModelBarDataProxy_DisconnectColumnRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QRegExp &)>(&QItemModelBarDataProxy::columnRolePatternChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QRegExp &)>(&MyQItemModelBarDataProxy::Signal_ColumnRolePatternChanged));
}

void QItemModelBarDataProxy_ColumnRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->columnRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelBarDataProxy_ConnectColumnRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::columnRoleReplaceChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_ColumnRoleReplaceChanged));
}

void QItemModelBarDataProxy_DisconnectColumnRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::columnRoleReplaceChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_ColumnRoleReplaceChanged));
}

void QItemModelBarDataProxy_ColumnRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->columnRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelBarDataProxy_ConnectItemModelChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QAbstractItemModel *)>(&QItemModelBarDataProxy::itemModelChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QAbstractItemModel *)>(&MyQItemModelBarDataProxy::Signal_ItemModelChanged));
}

void QItemModelBarDataProxy_DisconnectItemModelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QAbstractItemModel *)>(&QItemModelBarDataProxy::itemModelChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QAbstractItemModel *)>(&MyQItemModelBarDataProxy::Signal_ItemModelChanged));
}

void QItemModelBarDataProxy_ItemModelChanged(void* ptr, void* itemModel)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->itemModelChanged(static_cast<QAbstractItemModel*>(itemModel));
}

void QItemModelBarDataProxy_ConnectMultiMatchBehaviorChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(QItemModelBarDataProxy::MultiMatchBehavior)>(&QItemModelBarDataProxy::multiMatchBehaviorChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(QItemModelBarDataProxy::MultiMatchBehavior)>(&MyQItemModelBarDataProxy::Signal_MultiMatchBehaviorChanged));
}

void QItemModelBarDataProxy_DisconnectMultiMatchBehaviorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(QItemModelBarDataProxy::MultiMatchBehavior)>(&QItemModelBarDataProxy::multiMatchBehaviorChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(QItemModelBarDataProxy::MultiMatchBehavior)>(&MyQItemModelBarDataProxy::Signal_MultiMatchBehaviorChanged));
}

void QItemModelBarDataProxy_MultiMatchBehaviorChanged(void* ptr, long long behavior)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->multiMatchBehaviorChanged(static_cast<QItemModelBarDataProxy::MultiMatchBehavior>(behavior));
}

void QItemModelBarDataProxy_Remap(void* ptr, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString valueRole, struct QtDataVisualization_PackedString rotationRole, struct QtDataVisualization_PackedString rowCategories, struct QtDataVisualization_PackedString columnCategories)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->remap(QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(valueRole.data, valueRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts));
}

void QItemModelBarDataProxy_ConnectRotationRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::rotationRoleChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_RotationRoleChanged));
}

void QItemModelBarDataProxy_DisconnectRotationRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::rotationRoleChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_RotationRoleChanged));
}

void QItemModelBarDataProxy_RotationRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->rotationRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelBarDataProxy_ConnectRotationRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QRegExp &)>(&QItemModelBarDataProxy::rotationRolePatternChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QRegExp &)>(&MyQItemModelBarDataProxy::Signal_RotationRolePatternChanged));
}

void QItemModelBarDataProxy_DisconnectRotationRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QRegExp &)>(&QItemModelBarDataProxy::rotationRolePatternChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QRegExp &)>(&MyQItemModelBarDataProxy::Signal_RotationRolePatternChanged));
}

void QItemModelBarDataProxy_RotationRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->rotationRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelBarDataProxy_ConnectRotationRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::rotationRoleReplaceChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_RotationRoleReplaceChanged));
}

void QItemModelBarDataProxy_DisconnectRotationRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::rotationRoleReplaceChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_RotationRoleReplaceChanged));
}

void QItemModelBarDataProxy_RotationRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->rotationRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelBarDataProxy_ConnectRowCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)()>(&QItemModelBarDataProxy::rowCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)()>(&MyQItemModelBarDataProxy::Signal_RowCategoriesChanged));
}

void QItemModelBarDataProxy_DisconnectRowCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)()>(&QItemModelBarDataProxy::rowCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)()>(&MyQItemModelBarDataProxy::Signal_RowCategoriesChanged));
}

void QItemModelBarDataProxy_RowCategoriesChanged(void* ptr)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->rowCategoriesChanged();
}

void QItemModelBarDataProxy_ConnectRowRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::rowRoleChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_RowRoleChanged));
}

void QItemModelBarDataProxy_DisconnectRowRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::rowRoleChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_RowRoleChanged));
}

void QItemModelBarDataProxy_RowRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->rowRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelBarDataProxy_ConnectRowRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QRegExp &)>(&QItemModelBarDataProxy::rowRolePatternChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QRegExp &)>(&MyQItemModelBarDataProxy::Signal_RowRolePatternChanged));
}

void QItemModelBarDataProxy_DisconnectRowRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QRegExp &)>(&QItemModelBarDataProxy::rowRolePatternChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QRegExp &)>(&MyQItemModelBarDataProxy::Signal_RowRolePatternChanged));
}

void QItemModelBarDataProxy_RowRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->rowRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelBarDataProxy_ConnectRowRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::rowRoleReplaceChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_RowRoleReplaceChanged));
}

void QItemModelBarDataProxy_DisconnectRowRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::rowRoleReplaceChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_RowRoleReplaceChanged));
}

void QItemModelBarDataProxy_RowRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->rowRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelBarDataProxy_SetAutoColumnCategories(void* ptr, char enable)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setAutoColumnCategories(enable != 0);
}

void QItemModelBarDataProxy_SetAutoRowCategories(void* ptr, char enable)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setAutoRowCategories(enable != 0);
}

void QItemModelBarDataProxy_SetColumnCategories(void* ptr, struct QtDataVisualization_PackedString categories)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setColumnCategories(QString::fromUtf8(categories.data, categories.len).split("|", QString::SkipEmptyParts));
}

void QItemModelBarDataProxy_SetColumnRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setColumnRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelBarDataProxy_SetColumnRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setColumnRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelBarDataProxy_SetColumnRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setColumnRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelBarDataProxy_SetItemModel(void* ptr, void* itemModel)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setItemModel(static_cast<QAbstractItemModel*>(itemModel));
}

void QItemModelBarDataProxy_SetMultiMatchBehavior(void* ptr, long long behavior)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setMultiMatchBehavior(static_cast<QItemModelBarDataProxy::MultiMatchBehavior>(behavior));
}

void QItemModelBarDataProxy_SetRotationRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setRotationRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelBarDataProxy_SetRotationRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setRotationRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelBarDataProxy_SetRotationRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setRotationRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelBarDataProxy_SetRowCategories(void* ptr, struct QtDataVisualization_PackedString categories)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setRowCategories(QString::fromUtf8(categories.data, categories.len).split("|", QString::SkipEmptyParts));
}

void QItemModelBarDataProxy_SetRowRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setRowRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelBarDataProxy_SetRowRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setRowRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelBarDataProxy_SetRowRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setRowRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelBarDataProxy_SetUseModelCategories(void* ptr, char enable)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setUseModelCategories(enable != 0);
}

void QItemModelBarDataProxy_SetValueRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setValueRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelBarDataProxy_SetValueRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setValueRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelBarDataProxy_SetValueRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->setValueRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelBarDataProxy_ConnectUseModelCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(bool)>(&QItemModelBarDataProxy::useModelCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(bool)>(&MyQItemModelBarDataProxy::Signal_UseModelCategoriesChanged));
}

void QItemModelBarDataProxy_DisconnectUseModelCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(bool)>(&QItemModelBarDataProxy::useModelCategoriesChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(bool)>(&MyQItemModelBarDataProxy::Signal_UseModelCategoriesChanged));
}

void QItemModelBarDataProxy_UseModelCategoriesChanged(void* ptr, char enable)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->useModelCategoriesChanged(enable != 0);
}

void QItemModelBarDataProxy_ConnectValueRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::valueRoleChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_ValueRoleChanged));
}

void QItemModelBarDataProxy_DisconnectValueRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::valueRoleChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_ValueRoleChanged));
}

void QItemModelBarDataProxy_ValueRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->valueRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelBarDataProxy_ConnectValueRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QRegExp &)>(&QItemModelBarDataProxy::valueRolePatternChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QRegExp &)>(&MyQItemModelBarDataProxy::Signal_ValueRolePatternChanged));
}

void QItemModelBarDataProxy_DisconnectValueRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QRegExp &)>(&QItemModelBarDataProxy::valueRolePatternChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QRegExp &)>(&MyQItemModelBarDataProxy::Signal_ValueRolePatternChanged));
}

void QItemModelBarDataProxy_ValueRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->valueRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelBarDataProxy_ConnectValueRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::valueRoleReplaceChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_ValueRoleReplaceChanged));
}

void QItemModelBarDataProxy_DisconnectValueRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelBarDataProxy*>(ptr), static_cast<void (QItemModelBarDataProxy::*)(const QString &)>(&QItemModelBarDataProxy::valueRoleReplaceChanged), static_cast<MyQItemModelBarDataProxy*>(ptr), static_cast<void (MyQItemModelBarDataProxy::*)(const QString &)>(&MyQItemModelBarDataProxy::Signal_ValueRoleReplaceChanged));
}

void QItemModelBarDataProxy_ValueRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->valueRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelBarDataProxy_DestroyQItemModelBarDataProxy(void* ptr)
{
	static_cast<QItemModelBarDataProxy*>(ptr)->~QItemModelBarDataProxy();
}

void QItemModelBarDataProxy_DestroyQItemModelBarDataProxyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QItemModelBarDataProxy_MultiMatchBehavior(void* ptr)
{
	return static_cast<QItemModelBarDataProxy*>(ptr)->multiMatchBehavior();
}

void* QItemModelBarDataProxy_ItemModel(void* ptr)
{
	return static_cast<QItemModelBarDataProxy*>(ptr)->itemModel();
}

void* QItemModelBarDataProxy_ColumnRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelBarDataProxy*>(ptr)->columnRolePattern());
}

void* QItemModelBarDataProxy_RotationRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelBarDataProxy*>(ptr)->rotationRolePattern());
}

void* QItemModelBarDataProxy_RowRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelBarDataProxy*>(ptr)->rowRolePattern());
}

void* QItemModelBarDataProxy_ValueRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelBarDataProxy*>(ptr)->valueRolePattern());
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_ColumnRole(void* ptr)
{
	return ({ QByteArray t1d4975 = static_cast<QItemModelBarDataProxy*>(ptr)->columnRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t1d4975.prepend("WHITESPACE").constData()+10), t1d4975.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_ColumnRoleReplace(void* ptr)
{
	return ({ QByteArray tf1dce7 = static_cast<QItemModelBarDataProxy*>(ptr)->columnRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tf1dce7.prepend("WHITESPACE").constData()+10), tf1dce7.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_RotationRole(void* ptr)
{
	return ({ QByteArray ta1c8cc = static_cast<QItemModelBarDataProxy*>(ptr)->rotationRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(ta1c8cc.prepend("WHITESPACE").constData()+10), ta1c8cc.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_RotationRoleReplace(void* ptr)
{
	return ({ QByteArray tf65c4d = static_cast<QItemModelBarDataProxy*>(ptr)->rotationRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tf65c4d.prepend("WHITESPACE").constData()+10), tf65c4d.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_RowRole(void* ptr)
{
	return ({ QByteArray t609fcb = static_cast<QItemModelBarDataProxy*>(ptr)->rowRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t609fcb.prepend("WHITESPACE").constData()+10), t609fcb.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_RowRoleReplace(void* ptr)
{
	return ({ QByteArray td8a1ca = static_cast<QItemModelBarDataProxy*>(ptr)->rowRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(td8a1ca.prepend("WHITESPACE").constData()+10), td8a1ca.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_ValueRole(void* ptr)
{
	return ({ QByteArray t03456f = static_cast<QItemModelBarDataProxy*>(ptr)->valueRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t03456f.prepend("WHITESPACE").constData()+10), t03456f.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_ValueRoleReplace(void* ptr)
{
	return ({ QByteArray t1ef472 = static_cast<QItemModelBarDataProxy*>(ptr)->valueRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t1ef472.prepend("WHITESPACE").constData()+10), t1ef472.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_ColumnCategories(void* ptr)
{
	return ({ QByteArray t86b84b = static_cast<QItemModelBarDataProxy*>(ptr)->columnCategories().join("|").toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t86b84b.prepend("WHITESPACE").constData()+10), t86b84b.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelBarDataProxy_RowCategories(void* ptr)
{
	return ({ QByteArray t50f4e1 = static_cast<QItemModelBarDataProxy*>(ptr)->rowCategories().join("|").toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t50f4e1.prepend("WHITESPACE").constData()+10), t50f4e1.size()-10 }; });
}

char QItemModelBarDataProxy_AutoColumnCategories(void* ptr)
{
	return static_cast<QItemModelBarDataProxy*>(ptr)->autoColumnCategories();
}

char QItemModelBarDataProxy_AutoRowCategories(void* ptr)
{
	return static_cast<QItemModelBarDataProxy*>(ptr)->autoRowCategories();
}

char QItemModelBarDataProxy_UseModelCategories(void* ptr)
{
	return static_cast<QItemModelBarDataProxy*>(ptr)->useModelCategories();
}

class MyQItemModelScatterDataProxy: public QItemModelScatterDataProxy
{
public:
	MyQItemModelScatterDataProxy(QAbstractItemModel *itemModel, QObject *parent = Q_NULLPTR) : QItemModelScatterDataProxy(itemModel, parent) {};
	MyQItemModelScatterDataProxy(QAbstractItemModel *itemModel, const QString &xPosRole, const QString &yPosRole, const QString &zPosRole, QObject *parent = Q_NULLPTR) : QItemModelScatterDataProxy(itemModel, xPosRole, yPosRole, zPosRole, parent) {};
	MyQItemModelScatterDataProxy(QAbstractItemModel *itemModel, const QString &xPosRole, const QString &yPosRole, const QString &zPosRole, const QString &rotationRole, QObject *parent = Q_NULLPTR) : QItemModelScatterDataProxy(itemModel, xPosRole, yPosRole, zPosRole, rotationRole, parent) {};
	MyQItemModelScatterDataProxy(QObject *parent = Q_NULLPTR) : QItemModelScatterDataProxy(parent) {};
	void Signal_ItemModelChanged(const QAbstractItemModel * itemModel) { callbackQItemModelScatterDataProxy_ItemModelChanged(this, const_cast<QAbstractItemModel*>(itemModel)); };
	void Signal_RotationRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelScatterDataProxy_RotationRoleChanged(this, rolePacked); };
	void Signal_RotationRolePatternChanged(const QRegExp & pattern) { callbackQItemModelScatterDataProxy_RotationRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_RotationRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelScatterDataProxy_RotationRoleReplaceChanged(this, replacePacked); };
	void Signal_XPosRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelScatterDataProxy_XPosRoleChanged(this, rolePacked); };
	void Signal_XPosRolePatternChanged(const QRegExp & pattern) { callbackQItemModelScatterDataProxy_XPosRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_XPosRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelScatterDataProxy_XPosRoleReplaceChanged(this, replacePacked); };
	void Signal_YPosRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelScatterDataProxy_YPosRoleChanged(this, rolePacked); };
	void Signal_YPosRolePatternChanged(const QRegExp & pattern) { callbackQItemModelScatterDataProxy_YPosRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_YPosRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelScatterDataProxy_YPosRoleReplaceChanged(this, replacePacked); };
	void Signal_ZPosRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelScatterDataProxy_ZPosRoleChanged(this, rolePacked); };
	void Signal_ZPosRolePatternChanged(const QRegExp & pattern) { callbackQItemModelScatterDataProxy_ZPosRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_ZPosRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelScatterDataProxy_ZPosRoleReplaceChanged(this, replacePacked); };
	 ~MyQItemModelScatterDataProxy() { callbackQItemModelScatterDataProxy_DestroyQItemModelScatterDataProxy(this); };
};

void* QItemModelScatterDataProxy_NewQItemModelScatterDataProxy2(void* itemModel, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QObject*>(parent));
	}
}

void* QItemModelScatterDataProxy_NewQItemModelScatterDataProxy3(void* itemModel, struct QtDataVisualization_PackedString xPosRole, struct QtDataVisualization_PackedString yPosRole, struct QtDataVisualization_PackedString zPosRole, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QObject*>(parent));
	}
}

void* QItemModelScatterDataProxy_NewQItemModelScatterDataProxy4(void* itemModel, struct QtDataVisualization_PackedString xPosRole, struct QtDataVisualization_PackedString yPosRole, struct QtDataVisualization_PackedString zPosRole, struct QtDataVisualization_PackedString rotationRole, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelScatterDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len), static_cast<QObject*>(parent));
	}
}

void* QItemModelScatterDataProxy_NewQItemModelScatterDataProxy(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelScatterDataProxy(static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelScatterDataProxy(static_cast<QObject*>(parent));
	}
}

void QItemModelScatterDataProxy_ConnectItemModelChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QAbstractItemModel *)>(&QItemModelScatterDataProxy::itemModelChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QAbstractItemModel *)>(&MyQItemModelScatterDataProxy::Signal_ItemModelChanged));
}

void QItemModelScatterDataProxy_DisconnectItemModelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QAbstractItemModel *)>(&QItemModelScatterDataProxy::itemModelChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QAbstractItemModel *)>(&MyQItemModelScatterDataProxy::Signal_ItemModelChanged));
}

void QItemModelScatterDataProxy_ItemModelChanged(void* ptr, void* itemModel)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->itemModelChanged(static_cast<QAbstractItemModel*>(itemModel));
}

void QItemModelScatterDataProxy_Remap(void* ptr, struct QtDataVisualization_PackedString xPosRole, struct QtDataVisualization_PackedString yPosRole, struct QtDataVisualization_PackedString zPosRole, struct QtDataVisualization_PackedString rotationRole)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->remap(QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rotationRole.data, rotationRole.len));
}

void QItemModelScatterDataProxy_ConnectRotationRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::rotationRoleChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_RotationRoleChanged));
}

void QItemModelScatterDataProxy_DisconnectRotationRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::rotationRoleChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_RotationRoleChanged));
}

void QItemModelScatterDataProxy_RotationRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->rotationRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelScatterDataProxy_ConnectRotationRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QRegExp &)>(&QItemModelScatterDataProxy::rotationRolePatternChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QRegExp &)>(&MyQItemModelScatterDataProxy::Signal_RotationRolePatternChanged));
}

void QItemModelScatterDataProxy_DisconnectRotationRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QRegExp &)>(&QItemModelScatterDataProxy::rotationRolePatternChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QRegExp &)>(&MyQItemModelScatterDataProxy::Signal_RotationRolePatternChanged));
}

void QItemModelScatterDataProxy_RotationRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->rotationRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelScatterDataProxy_ConnectRotationRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::rotationRoleReplaceChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_RotationRoleReplaceChanged));
}

void QItemModelScatterDataProxy_DisconnectRotationRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::rotationRoleReplaceChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_RotationRoleReplaceChanged));
}

void QItemModelScatterDataProxy_RotationRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->rotationRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelScatterDataProxy_SetItemModel(void* ptr, void* itemModel)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setItemModel(static_cast<QAbstractItemModel*>(itemModel));
}

void QItemModelScatterDataProxy_SetRotationRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setRotationRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelScatterDataProxy_SetRotationRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setRotationRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelScatterDataProxy_SetRotationRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setRotationRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelScatterDataProxy_SetXPosRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setXPosRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelScatterDataProxy_SetXPosRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setXPosRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelScatterDataProxy_SetXPosRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setXPosRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelScatterDataProxy_SetYPosRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setYPosRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelScatterDataProxy_SetYPosRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setYPosRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelScatterDataProxy_SetYPosRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setYPosRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelScatterDataProxy_SetZPosRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setZPosRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelScatterDataProxy_SetZPosRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setZPosRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelScatterDataProxy_SetZPosRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->setZPosRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelScatterDataProxy_ConnectXPosRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::xPosRoleChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_XPosRoleChanged));
}

void QItemModelScatterDataProxy_DisconnectXPosRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::xPosRoleChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_XPosRoleChanged));
}

void QItemModelScatterDataProxy_XPosRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->xPosRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelScatterDataProxy_ConnectXPosRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QRegExp &)>(&QItemModelScatterDataProxy::xPosRolePatternChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QRegExp &)>(&MyQItemModelScatterDataProxy::Signal_XPosRolePatternChanged));
}

void QItemModelScatterDataProxy_DisconnectXPosRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QRegExp &)>(&QItemModelScatterDataProxy::xPosRolePatternChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QRegExp &)>(&MyQItemModelScatterDataProxy::Signal_XPosRolePatternChanged));
}

void QItemModelScatterDataProxy_XPosRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->xPosRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelScatterDataProxy_ConnectXPosRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::xPosRoleReplaceChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_XPosRoleReplaceChanged));
}

void QItemModelScatterDataProxy_DisconnectXPosRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::xPosRoleReplaceChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_XPosRoleReplaceChanged));
}

void QItemModelScatterDataProxy_XPosRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->xPosRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelScatterDataProxy_ConnectYPosRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::yPosRoleChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_YPosRoleChanged));
}

void QItemModelScatterDataProxy_DisconnectYPosRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::yPosRoleChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_YPosRoleChanged));
}

void QItemModelScatterDataProxy_YPosRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->yPosRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelScatterDataProxy_ConnectYPosRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QRegExp &)>(&QItemModelScatterDataProxy::yPosRolePatternChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QRegExp &)>(&MyQItemModelScatterDataProxy::Signal_YPosRolePatternChanged));
}

void QItemModelScatterDataProxy_DisconnectYPosRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QRegExp &)>(&QItemModelScatterDataProxy::yPosRolePatternChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QRegExp &)>(&MyQItemModelScatterDataProxy::Signal_YPosRolePatternChanged));
}

void QItemModelScatterDataProxy_YPosRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->yPosRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelScatterDataProxy_ConnectYPosRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::yPosRoleReplaceChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_YPosRoleReplaceChanged));
}

void QItemModelScatterDataProxy_DisconnectYPosRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::yPosRoleReplaceChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_YPosRoleReplaceChanged));
}

void QItemModelScatterDataProxy_YPosRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->yPosRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelScatterDataProxy_ConnectZPosRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::zPosRoleChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_ZPosRoleChanged));
}

void QItemModelScatterDataProxy_DisconnectZPosRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::zPosRoleChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_ZPosRoleChanged));
}

void QItemModelScatterDataProxy_ZPosRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->zPosRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelScatterDataProxy_ConnectZPosRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QRegExp &)>(&QItemModelScatterDataProxy::zPosRolePatternChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QRegExp &)>(&MyQItemModelScatterDataProxy::Signal_ZPosRolePatternChanged));
}

void QItemModelScatterDataProxy_DisconnectZPosRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QRegExp &)>(&QItemModelScatterDataProxy::zPosRolePatternChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QRegExp &)>(&MyQItemModelScatterDataProxy::Signal_ZPosRolePatternChanged));
}

void QItemModelScatterDataProxy_ZPosRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->zPosRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelScatterDataProxy_ConnectZPosRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::zPosRoleReplaceChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_ZPosRoleReplaceChanged));
}

void QItemModelScatterDataProxy_DisconnectZPosRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelScatterDataProxy*>(ptr), static_cast<void (QItemModelScatterDataProxy::*)(const QString &)>(&QItemModelScatterDataProxy::zPosRoleReplaceChanged), static_cast<MyQItemModelScatterDataProxy*>(ptr), static_cast<void (MyQItemModelScatterDataProxy::*)(const QString &)>(&MyQItemModelScatterDataProxy::Signal_ZPosRoleReplaceChanged));
}

void QItemModelScatterDataProxy_ZPosRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->zPosRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelScatterDataProxy_DestroyQItemModelScatterDataProxy(void* ptr)
{
	static_cast<QItemModelScatterDataProxy*>(ptr)->~QItemModelScatterDataProxy();
}

void QItemModelScatterDataProxy_DestroyQItemModelScatterDataProxyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QItemModelScatterDataProxy_ItemModel(void* ptr)
{
	return static_cast<QItemModelScatterDataProxy*>(ptr)->itemModel();
}

void* QItemModelScatterDataProxy_RotationRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelScatterDataProxy*>(ptr)->rotationRolePattern());
}

void* QItemModelScatterDataProxy_XPosRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelScatterDataProxy*>(ptr)->xPosRolePattern());
}

void* QItemModelScatterDataProxy_YPosRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelScatterDataProxy*>(ptr)->yPosRolePattern());
}

void* QItemModelScatterDataProxy_ZPosRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelScatterDataProxy*>(ptr)->zPosRolePattern());
}

struct QtDataVisualization_PackedString QItemModelScatterDataProxy_RotationRole(void* ptr)
{
	return ({ QByteArray tb50293 = static_cast<QItemModelScatterDataProxy*>(ptr)->rotationRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tb50293.prepend("WHITESPACE").constData()+10), tb50293.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelScatterDataProxy_RotationRoleReplace(void* ptr)
{
	return ({ QByteArray t94bb1d = static_cast<QItemModelScatterDataProxy*>(ptr)->rotationRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t94bb1d.prepend("WHITESPACE").constData()+10), t94bb1d.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelScatterDataProxy_XPosRole(void* ptr)
{
	return ({ QByteArray t4f9501 = static_cast<QItemModelScatterDataProxy*>(ptr)->xPosRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t4f9501.prepend("WHITESPACE").constData()+10), t4f9501.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelScatterDataProxy_XPosRoleReplace(void* ptr)
{
	return ({ QByteArray tb11de2 = static_cast<QItemModelScatterDataProxy*>(ptr)->xPosRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tb11de2.prepend("WHITESPACE").constData()+10), tb11de2.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelScatterDataProxy_YPosRole(void* ptr)
{
	return ({ QByteArray td91cd2 = static_cast<QItemModelScatterDataProxy*>(ptr)->yPosRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(td91cd2.prepend("WHITESPACE").constData()+10), td91cd2.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelScatterDataProxy_YPosRoleReplace(void* ptr)
{
	return ({ QByteArray t4307ad = static_cast<QItemModelScatterDataProxy*>(ptr)->yPosRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t4307ad.prepend("WHITESPACE").constData()+10), t4307ad.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelScatterDataProxy_ZPosRole(void* ptr)
{
	return ({ QByteArray t831570 = static_cast<QItemModelScatterDataProxy*>(ptr)->zPosRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t831570.prepend("WHITESPACE").constData()+10), t831570.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelScatterDataProxy_ZPosRoleReplace(void* ptr)
{
	return ({ QByteArray t4c6acb = static_cast<QItemModelScatterDataProxy*>(ptr)->zPosRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t4c6acb.prepend("WHITESPACE").constData()+10), t4c6acb.size()-10 }; });
}

class MyQItemModelSurfaceDataProxy: public QItemModelSurfaceDataProxy
{
public:
	MyQItemModelSurfaceDataProxy(QAbstractItemModel *itemModel, QObject *parent = Q_NULLPTR) : QItemModelSurfaceDataProxy(itemModel, parent) {};
	MyQItemModelSurfaceDataProxy(QAbstractItemModel *itemModel, const QString &rowRole, const QString &columnRole, const QString &xPosRole, const QString &yPosRole, const QString &zPosRole, QObject *parent = Q_NULLPTR) : QItemModelSurfaceDataProxy(itemModel, rowRole, columnRole, xPosRole, yPosRole, zPosRole, parent) {};
	MyQItemModelSurfaceDataProxy(QAbstractItemModel *itemModel, const QString &rowRole, const QString &columnRole, const QString &xPosRole, const QString &yPosRole, const QString &zPosRole, const QStringList &rowCategories, const QStringList &columnCategories, QObject *parent = Q_NULLPTR) : QItemModelSurfaceDataProxy(itemModel, rowRole, columnRole, xPosRole, yPosRole, zPosRole, rowCategories, columnCategories, parent) {};
	MyQItemModelSurfaceDataProxy(QAbstractItemModel *itemModel, const QString &rowRole, const QString &columnRole, const QString &yPosRole, QObject *parent = Q_NULLPTR) : QItemModelSurfaceDataProxy(itemModel, rowRole, columnRole, yPosRole, parent) {};
	MyQItemModelSurfaceDataProxy(QAbstractItemModel *itemModel, const QString &rowRole, const QString &columnRole, const QString &yPosRole, const QStringList &rowCategories, const QStringList &columnCategories, QObject *parent = Q_NULLPTR) : QItemModelSurfaceDataProxy(itemModel, rowRole, columnRole, yPosRole, rowCategories, columnCategories, parent) {};
	MyQItemModelSurfaceDataProxy(QAbstractItemModel *itemModel, const QString &yPosRole, QObject *parent = Q_NULLPTR) : QItemModelSurfaceDataProxy(itemModel, yPosRole, parent) {};
	MyQItemModelSurfaceDataProxy(QObject *parent = Q_NULLPTR) : QItemModelSurfaceDataProxy(parent) {};
	void Signal_AutoColumnCategoriesChanged(bool enable) { callbackQItemModelSurfaceDataProxy_AutoColumnCategoriesChanged(this, enable); };
	void Signal_AutoRowCategoriesChanged(bool enable) { callbackQItemModelSurfaceDataProxy_AutoRowCategoriesChanged(this, enable); };
	void Signal_ColumnCategoriesChanged() { callbackQItemModelSurfaceDataProxy_ColumnCategoriesChanged(this); };
	void Signal_ColumnRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelSurfaceDataProxy_ColumnRoleChanged(this, rolePacked); };
	void Signal_ColumnRolePatternChanged(const QRegExp & pattern) { callbackQItemModelSurfaceDataProxy_ColumnRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_ColumnRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelSurfaceDataProxy_ColumnRoleReplaceChanged(this, replacePacked); };
	void Signal_ItemModelChanged(const QAbstractItemModel * itemModel) { callbackQItemModelSurfaceDataProxy_ItemModelChanged(this, const_cast<QAbstractItemModel*>(itemModel)); };
	void Signal_MultiMatchBehaviorChanged(QItemModelSurfaceDataProxy::MultiMatchBehavior behavior) { callbackQItemModelSurfaceDataProxy_MultiMatchBehaviorChanged(this, behavior); };
	void Signal_RowCategoriesChanged() { callbackQItemModelSurfaceDataProxy_RowCategoriesChanged(this); };
	void Signal_RowRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelSurfaceDataProxy_RowRoleChanged(this, rolePacked); };
	void Signal_RowRolePatternChanged(const QRegExp & pattern) { callbackQItemModelSurfaceDataProxy_RowRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_RowRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelSurfaceDataProxy_RowRoleReplaceChanged(this, replacePacked); };
	void Signal_UseModelCategoriesChanged(bool enable) { callbackQItemModelSurfaceDataProxy_UseModelCategoriesChanged(this, enable); };
	void Signal_XPosRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelSurfaceDataProxy_XPosRoleChanged(this, rolePacked); };
	void Signal_XPosRolePatternChanged(const QRegExp & pattern) { callbackQItemModelSurfaceDataProxy_XPosRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_XPosRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelSurfaceDataProxy_XPosRoleReplaceChanged(this, replacePacked); };
	void Signal_YPosRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelSurfaceDataProxy_YPosRoleChanged(this, rolePacked); };
	void Signal_YPosRolePatternChanged(const QRegExp & pattern) { callbackQItemModelSurfaceDataProxy_YPosRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_YPosRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelSurfaceDataProxy_YPosRoleReplaceChanged(this, replacePacked); };
	void Signal_ZPosRoleChanged(const QString & role) { QByteArray t8dca46 = role.toUtf8(); QtDataVisualization_PackedString rolePacked = { const_cast<char*>(t8dca46.prepend("WHITESPACE").constData()+10), t8dca46.size()-10 };callbackQItemModelSurfaceDataProxy_ZPosRoleChanged(this, rolePacked); };
	void Signal_ZPosRolePatternChanged(const QRegExp & pattern) { callbackQItemModelSurfaceDataProxy_ZPosRolePatternChanged(this, const_cast<QRegExp*>(&pattern)); };
	void Signal_ZPosRoleReplaceChanged(const QString & replace) { QByteArray t3cacc7 = replace.toUtf8(); QtDataVisualization_PackedString replacePacked = { const_cast<char*>(t3cacc7.prepend("WHITESPACE").constData()+10), t3cacc7.size()-10 };callbackQItemModelSurfaceDataProxy_ZPosRoleReplaceChanged(this, replacePacked); };
	 ~MyQItemModelSurfaceDataProxy() { callbackQItemModelSurfaceDataProxy_DestroyQItemModelSurfaceDataProxy(this); };
};

void* QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy2(void* itemModel, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), static_cast<QObject*>(parent));
	}
}

void* QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy5(void* itemModel, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString xPosRole, struct QtDataVisualization_PackedString yPosRole, struct QtDataVisualization_PackedString zPosRole, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), static_cast<QObject*>(parent));
	}
}

void* QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy7(void* itemModel, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString xPosRole, struct QtDataVisualization_PackedString yPosRole, struct QtDataVisualization_PackedString zPosRole, struct QtDataVisualization_PackedString rowCategories, struct QtDataVisualization_PackedString columnCategories, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
	}
}

void* QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy4(void* itemModel, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString yPosRole, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QObject*>(parent));
	}
}

void* QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy6(void* itemModel, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString yPosRole, struct QtDataVisualization_PackedString rowCategories, struct QtDataVisualization_PackedString columnCategories, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
	}
}

void* QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy3(void* itemModel, struct QtDataVisualization_PackedString yPosRole, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelSurfaceDataProxy(static_cast<QAbstractItemModel*>(itemModel), QString::fromUtf8(yPosRole.data, yPosRole.len), static_cast<QObject*>(parent));
	}
}

void* QItemModelSurfaceDataProxy_NewQItemModelSurfaceDataProxy(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQItemModelSurfaceDataProxy(static_cast<QWindow*>(parent));
	} else {
		return new MyQItemModelSurfaceDataProxy(static_cast<QObject*>(parent));
	}
}

int QItemModelSurfaceDataProxy_ColumnCategoryIndex(void* ptr, struct QtDataVisualization_PackedString category)
{
	return static_cast<QItemModelSurfaceDataProxy*>(ptr)->columnCategoryIndex(QString::fromUtf8(category.data, category.len));
}

int QItemModelSurfaceDataProxy_RowCategoryIndex(void* ptr, struct QtDataVisualization_PackedString category)
{
	return static_cast<QItemModelSurfaceDataProxy*>(ptr)->rowCategoryIndex(QString::fromUtf8(category.data, category.len));
}

void QItemModelSurfaceDataProxy_ConnectAutoColumnCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(bool)>(&QItemModelSurfaceDataProxy::autoColumnCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(bool)>(&MyQItemModelSurfaceDataProxy::Signal_AutoColumnCategoriesChanged));
}

void QItemModelSurfaceDataProxy_DisconnectAutoColumnCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(bool)>(&QItemModelSurfaceDataProxy::autoColumnCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(bool)>(&MyQItemModelSurfaceDataProxy::Signal_AutoColumnCategoriesChanged));
}

void QItemModelSurfaceDataProxy_AutoColumnCategoriesChanged(void* ptr, char enable)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->autoColumnCategoriesChanged(enable != 0);
}

void QItemModelSurfaceDataProxy_ConnectAutoRowCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(bool)>(&QItemModelSurfaceDataProxy::autoRowCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(bool)>(&MyQItemModelSurfaceDataProxy::Signal_AutoRowCategoriesChanged));
}

void QItemModelSurfaceDataProxy_DisconnectAutoRowCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(bool)>(&QItemModelSurfaceDataProxy::autoRowCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(bool)>(&MyQItemModelSurfaceDataProxy::Signal_AutoRowCategoriesChanged));
}

void QItemModelSurfaceDataProxy_AutoRowCategoriesChanged(void* ptr, char enable)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->autoRowCategoriesChanged(enable != 0);
}

void QItemModelSurfaceDataProxy_ConnectColumnCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)()>(&QItemModelSurfaceDataProxy::columnCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)()>(&MyQItemModelSurfaceDataProxy::Signal_ColumnCategoriesChanged));
}

void QItemModelSurfaceDataProxy_DisconnectColumnCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)()>(&QItemModelSurfaceDataProxy::columnCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)()>(&MyQItemModelSurfaceDataProxy::Signal_ColumnCategoriesChanged));
}

void QItemModelSurfaceDataProxy_ColumnCategoriesChanged(void* ptr)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->columnCategoriesChanged();
}

void QItemModelSurfaceDataProxy_ConnectColumnRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::columnRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_ColumnRoleChanged));
}

void QItemModelSurfaceDataProxy_DisconnectColumnRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::columnRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_ColumnRoleChanged));
}

void QItemModelSurfaceDataProxy_ColumnRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->columnRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_ConnectColumnRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::columnRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_ColumnRolePatternChanged));
}

void QItemModelSurfaceDataProxy_DisconnectColumnRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::columnRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_ColumnRolePatternChanged));
}

void QItemModelSurfaceDataProxy_ColumnRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->columnRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_ConnectColumnRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::columnRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_ColumnRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_DisconnectColumnRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::columnRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_ColumnRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_ColumnRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->columnRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_ConnectItemModelChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QAbstractItemModel *)>(&QItemModelSurfaceDataProxy::itemModelChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QAbstractItemModel *)>(&MyQItemModelSurfaceDataProxy::Signal_ItemModelChanged));
}

void QItemModelSurfaceDataProxy_DisconnectItemModelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QAbstractItemModel *)>(&QItemModelSurfaceDataProxy::itemModelChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QAbstractItemModel *)>(&MyQItemModelSurfaceDataProxy::Signal_ItemModelChanged));
}

void QItemModelSurfaceDataProxy_ItemModelChanged(void* ptr, void* itemModel)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->itemModelChanged(static_cast<QAbstractItemModel*>(itemModel));
}

void QItemModelSurfaceDataProxy_ConnectMultiMatchBehaviorChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(QItemModelSurfaceDataProxy::MultiMatchBehavior)>(&QItemModelSurfaceDataProxy::multiMatchBehaviorChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(QItemModelSurfaceDataProxy::MultiMatchBehavior)>(&MyQItemModelSurfaceDataProxy::Signal_MultiMatchBehaviorChanged));
}

void QItemModelSurfaceDataProxy_DisconnectMultiMatchBehaviorChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(QItemModelSurfaceDataProxy::MultiMatchBehavior)>(&QItemModelSurfaceDataProxy::multiMatchBehaviorChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(QItemModelSurfaceDataProxy::MultiMatchBehavior)>(&MyQItemModelSurfaceDataProxy::Signal_MultiMatchBehaviorChanged));
}

void QItemModelSurfaceDataProxy_MultiMatchBehaviorChanged(void* ptr, long long behavior)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->multiMatchBehaviorChanged(static_cast<QItemModelSurfaceDataProxy::MultiMatchBehavior>(behavior));
}

void QItemModelSurfaceDataProxy_Remap(void* ptr, struct QtDataVisualization_PackedString rowRole, struct QtDataVisualization_PackedString columnRole, struct QtDataVisualization_PackedString xPosRole, struct QtDataVisualization_PackedString yPosRole, struct QtDataVisualization_PackedString zPosRole, struct QtDataVisualization_PackedString rowCategories, struct QtDataVisualization_PackedString columnCategories)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->remap(QString::fromUtf8(rowRole.data, rowRole.len), QString::fromUtf8(columnRole.data, columnRole.len), QString::fromUtf8(xPosRole.data, xPosRole.len), QString::fromUtf8(yPosRole.data, yPosRole.len), QString::fromUtf8(zPosRole.data, zPosRole.len), QString::fromUtf8(rowCategories.data, rowCategories.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(columnCategories.data, columnCategories.len).split("|", QString::SkipEmptyParts));
}

void QItemModelSurfaceDataProxy_ConnectRowCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)()>(&QItemModelSurfaceDataProxy::rowCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)()>(&MyQItemModelSurfaceDataProxy::Signal_RowCategoriesChanged));
}

void QItemModelSurfaceDataProxy_DisconnectRowCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)()>(&QItemModelSurfaceDataProxy::rowCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)()>(&MyQItemModelSurfaceDataProxy::Signal_RowCategoriesChanged));
}

void QItemModelSurfaceDataProxy_RowCategoriesChanged(void* ptr)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->rowCategoriesChanged();
}

void QItemModelSurfaceDataProxy_ConnectRowRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::rowRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_RowRoleChanged));
}

void QItemModelSurfaceDataProxy_DisconnectRowRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::rowRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_RowRoleChanged));
}

void QItemModelSurfaceDataProxy_RowRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->rowRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_ConnectRowRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::rowRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_RowRolePatternChanged));
}

void QItemModelSurfaceDataProxy_DisconnectRowRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::rowRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_RowRolePatternChanged));
}

void QItemModelSurfaceDataProxy_RowRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->rowRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_ConnectRowRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::rowRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_RowRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_DisconnectRowRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::rowRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_RowRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_RowRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->rowRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_SetAutoColumnCategories(void* ptr, char enable)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setAutoColumnCategories(enable != 0);
}

void QItemModelSurfaceDataProxy_SetAutoRowCategories(void* ptr, char enable)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setAutoRowCategories(enable != 0);
}

void QItemModelSurfaceDataProxy_SetColumnCategories(void* ptr, struct QtDataVisualization_PackedString categories)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setColumnCategories(QString::fromUtf8(categories.data, categories.len).split("|", QString::SkipEmptyParts));
}

void QItemModelSurfaceDataProxy_SetColumnRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setColumnRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_SetColumnRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setColumnRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_SetColumnRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setColumnRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_SetItemModel(void* ptr, void* itemModel)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setItemModel(static_cast<QAbstractItemModel*>(itemModel));
}

void QItemModelSurfaceDataProxy_SetMultiMatchBehavior(void* ptr, long long behavior)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setMultiMatchBehavior(static_cast<QItemModelSurfaceDataProxy::MultiMatchBehavior>(behavior));
}

void QItemModelSurfaceDataProxy_SetRowCategories(void* ptr, struct QtDataVisualization_PackedString categories)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setRowCategories(QString::fromUtf8(categories.data, categories.len).split("|", QString::SkipEmptyParts));
}

void QItemModelSurfaceDataProxy_SetRowRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setRowRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_SetRowRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setRowRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_SetRowRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setRowRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_SetUseModelCategories(void* ptr, char enable)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setUseModelCategories(enable != 0);
}

void QItemModelSurfaceDataProxy_SetXPosRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setXPosRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_SetXPosRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setXPosRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_SetXPosRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setXPosRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_SetYPosRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setYPosRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_SetYPosRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setYPosRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_SetYPosRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setYPosRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_SetZPosRole(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setZPosRole(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_SetZPosRolePattern(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setZPosRolePattern(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_SetZPosRoleReplace(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->setZPosRoleReplace(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_ConnectUseModelCategoriesChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(bool)>(&QItemModelSurfaceDataProxy::useModelCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(bool)>(&MyQItemModelSurfaceDataProxy::Signal_UseModelCategoriesChanged));
}

void QItemModelSurfaceDataProxy_DisconnectUseModelCategoriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(bool)>(&QItemModelSurfaceDataProxy::useModelCategoriesChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(bool)>(&MyQItemModelSurfaceDataProxy::Signal_UseModelCategoriesChanged));
}

void QItemModelSurfaceDataProxy_UseModelCategoriesChanged(void* ptr, char enable)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->useModelCategoriesChanged(enable != 0);
}

void QItemModelSurfaceDataProxy_ConnectXPosRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::xPosRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_XPosRoleChanged));
}

void QItemModelSurfaceDataProxy_DisconnectXPosRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::xPosRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_XPosRoleChanged));
}

void QItemModelSurfaceDataProxy_XPosRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->xPosRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_ConnectXPosRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::xPosRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_XPosRolePatternChanged));
}

void QItemModelSurfaceDataProxy_DisconnectXPosRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::xPosRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_XPosRolePatternChanged));
}

void QItemModelSurfaceDataProxy_XPosRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->xPosRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_ConnectXPosRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::xPosRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_XPosRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_DisconnectXPosRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::xPosRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_XPosRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_XPosRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->xPosRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_ConnectYPosRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::yPosRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_YPosRoleChanged));
}

void QItemModelSurfaceDataProxy_DisconnectYPosRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::yPosRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_YPosRoleChanged));
}

void QItemModelSurfaceDataProxy_YPosRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->yPosRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_ConnectYPosRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::yPosRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_YPosRolePatternChanged));
}

void QItemModelSurfaceDataProxy_DisconnectYPosRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::yPosRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_YPosRolePatternChanged));
}

void QItemModelSurfaceDataProxy_YPosRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->yPosRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_ConnectYPosRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::yPosRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_YPosRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_DisconnectYPosRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::yPosRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_YPosRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_YPosRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->yPosRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_ConnectZPosRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::zPosRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_ZPosRoleChanged));
}

void QItemModelSurfaceDataProxy_DisconnectZPosRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::zPosRoleChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_ZPosRoleChanged));
}

void QItemModelSurfaceDataProxy_ZPosRoleChanged(void* ptr, struct QtDataVisualization_PackedString role)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->zPosRoleChanged(QString::fromUtf8(role.data, role.len));
}

void QItemModelSurfaceDataProxy_ConnectZPosRolePatternChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::zPosRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_ZPosRolePatternChanged));
}

void QItemModelSurfaceDataProxy_DisconnectZPosRolePatternChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QRegExp &)>(&QItemModelSurfaceDataProxy::zPosRolePatternChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QRegExp &)>(&MyQItemModelSurfaceDataProxy::Signal_ZPosRolePatternChanged));
}

void QItemModelSurfaceDataProxy_ZPosRolePatternChanged(void* ptr, void* pattern)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->zPosRolePatternChanged(*static_cast<QRegExp*>(pattern));
}

void QItemModelSurfaceDataProxy_ConnectZPosRoleReplaceChanged(void* ptr)
{
	QObject::connect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::zPosRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_ZPosRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_DisconnectZPosRoleReplaceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QItemModelSurfaceDataProxy*>(ptr), static_cast<void (QItemModelSurfaceDataProxy::*)(const QString &)>(&QItemModelSurfaceDataProxy::zPosRoleReplaceChanged), static_cast<MyQItemModelSurfaceDataProxy*>(ptr), static_cast<void (MyQItemModelSurfaceDataProxy::*)(const QString &)>(&MyQItemModelSurfaceDataProxy::Signal_ZPosRoleReplaceChanged));
}

void QItemModelSurfaceDataProxy_ZPosRoleReplaceChanged(void* ptr, struct QtDataVisualization_PackedString replace)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->zPosRoleReplaceChanged(QString::fromUtf8(replace.data, replace.len));
}

void QItemModelSurfaceDataProxy_DestroyQItemModelSurfaceDataProxy(void* ptr)
{
	static_cast<QItemModelSurfaceDataProxy*>(ptr)->~QItemModelSurfaceDataProxy();
}

void QItemModelSurfaceDataProxy_DestroyQItemModelSurfaceDataProxyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QItemModelSurfaceDataProxy_MultiMatchBehavior(void* ptr)
{
	return static_cast<QItemModelSurfaceDataProxy*>(ptr)->multiMatchBehavior();
}

void* QItemModelSurfaceDataProxy_ItemModel(void* ptr)
{
	return static_cast<QItemModelSurfaceDataProxy*>(ptr)->itemModel();
}

void* QItemModelSurfaceDataProxy_ColumnRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelSurfaceDataProxy*>(ptr)->columnRolePattern());
}

void* QItemModelSurfaceDataProxy_RowRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelSurfaceDataProxy*>(ptr)->rowRolePattern());
}

void* QItemModelSurfaceDataProxy_XPosRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelSurfaceDataProxy*>(ptr)->xPosRolePattern());
}

void* QItemModelSurfaceDataProxy_YPosRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelSurfaceDataProxy*>(ptr)->yPosRolePattern());
}

void* QItemModelSurfaceDataProxy_ZPosRolePattern(void* ptr)
{
	return new QRegExp(static_cast<QItemModelSurfaceDataProxy*>(ptr)->zPosRolePattern());
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_ColumnRole(void* ptr)
{
	return ({ QByteArray tb79848 = static_cast<QItemModelSurfaceDataProxy*>(ptr)->columnRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tb79848.prepend("WHITESPACE").constData()+10), tb79848.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_ColumnRoleReplace(void* ptr)
{
	return ({ QByteArray t935d1d = static_cast<QItemModelSurfaceDataProxy*>(ptr)->columnRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t935d1d.prepend("WHITESPACE").constData()+10), t935d1d.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_RowRole(void* ptr)
{
	return ({ QByteArray t5df08d = static_cast<QItemModelSurfaceDataProxy*>(ptr)->rowRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t5df08d.prepend("WHITESPACE").constData()+10), t5df08d.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_RowRoleReplace(void* ptr)
{
	return ({ QByteArray t2f4502 = static_cast<QItemModelSurfaceDataProxy*>(ptr)->rowRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t2f4502.prepend("WHITESPACE").constData()+10), t2f4502.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_XPosRole(void* ptr)
{
	return ({ QByteArray t72f466 = static_cast<QItemModelSurfaceDataProxy*>(ptr)->xPosRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t72f466.prepend("WHITESPACE").constData()+10), t72f466.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_XPosRoleReplace(void* ptr)
{
	return ({ QByteArray t9a9018 = static_cast<QItemModelSurfaceDataProxy*>(ptr)->xPosRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t9a9018.prepend("WHITESPACE").constData()+10), t9a9018.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_YPosRole(void* ptr)
{
	return ({ QByteArray ta04979 = static_cast<QItemModelSurfaceDataProxy*>(ptr)->yPosRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(ta04979.prepend("WHITESPACE").constData()+10), ta04979.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_YPosRoleReplace(void* ptr)
{
	return ({ QByteArray te4bace = static_cast<QItemModelSurfaceDataProxy*>(ptr)->yPosRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(te4bace.prepend("WHITESPACE").constData()+10), te4bace.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_ZPosRole(void* ptr)
{
	return ({ QByteArray tba9475 = static_cast<QItemModelSurfaceDataProxy*>(ptr)->zPosRole().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tba9475.prepend("WHITESPACE").constData()+10), tba9475.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_ZPosRoleReplace(void* ptr)
{
	return ({ QByteArray t5d7f56 = static_cast<QItemModelSurfaceDataProxy*>(ptr)->zPosRoleReplace().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t5d7f56.prepend("WHITESPACE").constData()+10), t5d7f56.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_ColumnCategories(void* ptr)
{
	return ({ QByteArray tb94ba4 = static_cast<QItemModelSurfaceDataProxy*>(ptr)->columnCategories().join("|").toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tb94ba4.prepend("WHITESPACE").constData()+10), tb94ba4.size()-10 }; });
}

struct QtDataVisualization_PackedString QItemModelSurfaceDataProxy_RowCategories(void* ptr)
{
	return ({ QByteArray t02a480 = static_cast<QItemModelSurfaceDataProxy*>(ptr)->rowCategories().join("|").toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t02a480.prepend("WHITESPACE").constData()+10), t02a480.size()-10 }; });
}

char QItemModelSurfaceDataProxy_AutoColumnCategories(void* ptr)
{
	return static_cast<QItemModelSurfaceDataProxy*>(ptr)->autoColumnCategories();
}

char QItemModelSurfaceDataProxy_AutoRowCategories(void* ptr)
{
	return static_cast<QItemModelSurfaceDataProxy*>(ptr)->autoRowCategories();
}

char QItemModelSurfaceDataProxy_UseModelCategories(void* ptr)
{
	return static_cast<QItemModelSurfaceDataProxy*>(ptr)->useModelCategories();
}

class MyQLogValue3DAxisFormatter: public QLogValue3DAxisFormatter
{
public:
	MyQLogValue3DAxisFormatter(QObject *parent = Q_NULLPTR) : QLogValue3DAxisFormatter(parent) {};
	void Signal_AutoSubGridChanged(bool enabled) { callbackQLogValue3DAxisFormatter_AutoSubGridChanged(this, enabled); };
	void Signal_BaseChanged(qreal base) { callbackQLogValue3DAxisFormatter_BaseChanged(this, base); };
	void Signal_ShowEdgeLabelsChanged(bool enabled) { callbackQLogValue3DAxisFormatter_ShowEdgeLabelsChanged(this, enabled); };
	 ~MyQLogValue3DAxisFormatter() { callbackQLogValue3DAxisFormatter_DestroyQLogValue3DAxisFormatter(this); };
};

void* QLogValue3DAxisFormatter_NewQLogValue3DAxisFormatter(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLogValue3DAxisFormatter(static_cast<QWindow*>(parent));
	} else {
		return new MyQLogValue3DAxisFormatter(static_cast<QObject*>(parent));
	}
}

void QLogValue3DAxisFormatter_ConnectAutoSubGridChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValue3DAxisFormatter*>(ptr), static_cast<void (QLogValue3DAxisFormatter::*)(bool)>(&QLogValue3DAxisFormatter::autoSubGridChanged), static_cast<MyQLogValue3DAxisFormatter*>(ptr), static_cast<void (MyQLogValue3DAxisFormatter::*)(bool)>(&MyQLogValue3DAxisFormatter::Signal_AutoSubGridChanged));
}

void QLogValue3DAxisFormatter_DisconnectAutoSubGridChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValue3DAxisFormatter*>(ptr), static_cast<void (QLogValue3DAxisFormatter::*)(bool)>(&QLogValue3DAxisFormatter::autoSubGridChanged), static_cast<MyQLogValue3DAxisFormatter*>(ptr), static_cast<void (MyQLogValue3DAxisFormatter::*)(bool)>(&MyQLogValue3DAxisFormatter::Signal_AutoSubGridChanged));
}

void QLogValue3DAxisFormatter_AutoSubGridChanged(void* ptr, char enabled)
{
	static_cast<QLogValue3DAxisFormatter*>(ptr)->autoSubGridChanged(enabled != 0);
}

void QLogValue3DAxisFormatter_ConnectBaseChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValue3DAxisFormatter*>(ptr), static_cast<void (QLogValue3DAxisFormatter::*)(qreal)>(&QLogValue3DAxisFormatter::baseChanged), static_cast<MyQLogValue3DAxisFormatter*>(ptr), static_cast<void (MyQLogValue3DAxisFormatter::*)(qreal)>(&MyQLogValue3DAxisFormatter::Signal_BaseChanged));
}

void QLogValue3DAxisFormatter_DisconnectBaseChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValue3DAxisFormatter*>(ptr), static_cast<void (QLogValue3DAxisFormatter::*)(qreal)>(&QLogValue3DAxisFormatter::baseChanged), static_cast<MyQLogValue3DAxisFormatter*>(ptr), static_cast<void (MyQLogValue3DAxisFormatter::*)(qreal)>(&MyQLogValue3DAxisFormatter::Signal_BaseChanged));
}

void QLogValue3DAxisFormatter_BaseChanged(void* ptr, double base)
{
	static_cast<QLogValue3DAxisFormatter*>(ptr)->baseChanged(base);
}

void QLogValue3DAxisFormatter_SetAutoSubGrid(void* ptr, char enabled)
{
	static_cast<QLogValue3DAxisFormatter*>(ptr)->setAutoSubGrid(enabled != 0);
}

void QLogValue3DAxisFormatter_SetBase(void* ptr, double base)
{
	static_cast<QLogValue3DAxisFormatter*>(ptr)->setBase(base);
}

void QLogValue3DAxisFormatter_SetShowEdgeLabels(void* ptr, char enabled)
{
	static_cast<QLogValue3DAxisFormatter*>(ptr)->setShowEdgeLabels(enabled != 0);
}

void QLogValue3DAxisFormatter_ConnectShowEdgeLabelsChanged(void* ptr)
{
	QObject::connect(static_cast<QLogValue3DAxisFormatter*>(ptr), static_cast<void (QLogValue3DAxisFormatter::*)(bool)>(&QLogValue3DAxisFormatter::showEdgeLabelsChanged), static_cast<MyQLogValue3DAxisFormatter*>(ptr), static_cast<void (MyQLogValue3DAxisFormatter::*)(bool)>(&MyQLogValue3DAxisFormatter::Signal_ShowEdgeLabelsChanged));
}

void QLogValue3DAxisFormatter_DisconnectShowEdgeLabelsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLogValue3DAxisFormatter*>(ptr), static_cast<void (QLogValue3DAxisFormatter::*)(bool)>(&QLogValue3DAxisFormatter::showEdgeLabelsChanged), static_cast<MyQLogValue3DAxisFormatter*>(ptr), static_cast<void (MyQLogValue3DAxisFormatter::*)(bool)>(&MyQLogValue3DAxisFormatter::Signal_ShowEdgeLabelsChanged));
}

void QLogValue3DAxisFormatter_ShowEdgeLabelsChanged(void* ptr, char enabled)
{
	static_cast<QLogValue3DAxisFormatter*>(ptr)->showEdgeLabelsChanged(enabled != 0);
}

void QLogValue3DAxisFormatter_DestroyQLogValue3DAxisFormatter(void* ptr)
{
	static_cast<QLogValue3DAxisFormatter*>(ptr)->~QLogValue3DAxisFormatter();
}

void QLogValue3DAxisFormatter_DestroyQLogValue3DAxisFormatterDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char QLogValue3DAxisFormatter_AutoSubGrid(void* ptr)
{
	return static_cast<QLogValue3DAxisFormatter*>(ptr)->autoSubGrid();
}

char QLogValue3DAxisFormatter_ShowEdgeLabels(void* ptr)
{
	return static_cast<QLogValue3DAxisFormatter*>(ptr)->showEdgeLabels();
}

double QLogValue3DAxisFormatter_Base(void* ptr)
{
	return static_cast<QLogValue3DAxisFormatter*>(ptr)->base();
}

class MyQScatter3DSeries: public QScatter3DSeries
{
public:
	MyQScatter3DSeries(QObject *parent = Q_NULLPTR) : QScatter3DSeries(parent) {};
	MyQScatter3DSeries(QScatterDataProxy *dataProxy, QObject *parent = Q_NULLPTR) : QScatter3DSeries(dataProxy, parent) {};
	void Signal_DataProxyChanged(QScatterDataProxy * proxy) { callbackQScatter3DSeries_DataProxyChanged(this, proxy); };
	void Signal_ItemSizeChanged(float size) { callbackQScatter3DSeries_ItemSizeChanged(this, size); };
	void Signal_SelectedItemChanged(int index) { callbackQScatter3DSeries_SelectedItemChanged(this, index); };
	 ~MyQScatter3DSeries() { callbackQScatter3DSeries_DestroyQScatter3DSeries(this); };
};

void* QScatter3DSeries_NewQScatter3DSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQScatter3DSeries(static_cast<QObject*>(parent));
	}
}

void* QScatter3DSeries_NewQScatter3DSeries2(void* dataProxy, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QWindow*>(parent));
	} else {
		return new MyQScatter3DSeries(static_cast<QScatterDataProxy*>(dataProxy), static_cast<QObject*>(parent));
	}
}

int QScatter3DSeries_QScatter3DSeries_InvalidSelectionIndex()
{
	return QScatter3DSeries::invalidSelectionIndex();
}

void QScatter3DSeries_ConnectDataProxyChanged(void* ptr)
{
	QObject::connect(static_cast<QScatter3DSeries*>(ptr), static_cast<void (QScatter3DSeries::*)(QScatterDataProxy *)>(&QScatter3DSeries::dataProxyChanged), static_cast<MyQScatter3DSeries*>(ptr), static_cast<void (MyQScatter3DSeries::*)(QScatterDataProxy *)>(&MyQScatter3DSeries::Signal_DataProxyChanged));
}

void QScatter3DSeries_DisconnectDataProxyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScatter3DSeries*>(ptr), static_cast<void (QScatter3DSeries::*)(QScatterDataProxy *)>(&QScatter3DSeries::dataProxyChanged), static_cast<MyQScatter3DSeries*>(ptr), static_cast<void (MyQScatter3DSeries::*)(QScatterDataProxy *)>(&MyQScatter3DSeries::Signal_DataProxyChanged));
}

void QScatter3DSeries_DataProxyChanged(void* ptr, void* proxy)
{
	static_cast<QScatter3DSeries*>(ptr)->dataProxyChanged(static_cast<QScatterDataProxy*>(proxy));
}

void QScatter3DSeries_ConnectItemSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QScatter3DSeries*>(ptr), static_cast<void (QScatter3DSeries::*)(float)>(&QScatter3DSeries::itemSizeChanged), static_cast<MyQScatter3DSeries*>(ptr), static_cast<void (MyQScatter3DSeries::*)(float)>(&MyQScatter3DSeries::Signal_ItemSizeChanged));
}

void QScatter3DSeries_DisconnectItemSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScatter3DSeries*>(ptr), static_cast<void (QScatter3DSeries::*)(float)>(&QScatter3DSeries::itemSizeChanged), static_cast<MyQScatter3DSeries*>(ptr), static_cast<void (MyQScatter3DSeries::*)(float)>(&MyQScatter3DSeries::Signal_ItemSizeChanged));
}

void QScatter3DSeries_ItemSizeChanged(void* ptr, float size)
{
	static_cast<QScatter3DSeries*>(ptr)->itemSizeChanged(size);
}

void QScatter3DSeries_ConnectSelectedItemChanged(void* ptr)
{
	QObject::connect(static_cast<QScatter3DSeries*>(ptr), static_cast<void (QScatter3DSeries::*)(int)>(&QScatter3DSeries::selectedItemChanged), static_cast<MyQScatter3DSeries*>(ptr), static_cast<void (MyQScatter3DSeries::*)(int)>(&MyQScatter3DSeries::Signal_SelectedItemChanged));
}

void QScatter3DSeries_DisconnectSelectedItemChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScatter3DSeries*>(ptr), static_cast<void (QScatter3DSeries::*)(int)>(&QScatter3DSeries::selectedItemChanged), static_cast<MyQScatter3DSeries*>(ptr), static_cast<void (MyQScatter3DSeries::*)(int)>(&MyQScatter3DSeries::Signal_SelectedItemChanged));
}

void QScatter3DSeries_SelectedItemChanged(void* ptr, int index)
{
	static_cast<QScatter3DSeries*>(ptr)->selectedItemChanged(index);
}

void QScatter3DSeries_SetDataProxy(void* ptr, void* proxy)
{
	static_cast<QScatter3DSeries*>(ptr)->setDataProxy(static_cast<QScatterDataProxy*>(proxy));
}

void QScatter3DSeries_SetItemSize(void* ptr, float size)
{
	static_cast<QScatter3DSeries*>(ptr)->setItemSize(size);
}

void QScatter3DSeries_SetSelectedItem(void* ptr, int index)
{
	static_cast<QScatter3DSeries*>(ptr)->setSelectedItem(index);
}

void QScatter3DSeries_DestroyQScatter3DSeries(void* ptr)
{
	static_cast<QScatter3DSeries*>(ptr)->~QScatter3DSeries();
}

void QScatter3DSeries_DestroyQScatter3DSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QScatter3DSeries_DataProxy(void* ptr)
{
	return static_cast<QScatter3DSeries*>(ptr)->dataProxy();
}

float QScatter3DSeries_ItemSize(void* ptr)
{
	return static_cast<QScatter3DSeries*>(ptr)->itemSize();
}

int QScatter3DSeries_SelectedItem(void* ptr)
{
	return static_cast<QScatter3DSeries*>(ptr)->selectedItem();
}

void* QScatterDataItem_NewQScatterDataItem()
{
	return new QScatterDataItem();
}

void* QScatterDataItem_NewQScatterDataItem4(void* other)
{
	return new QScatterDataItem(*static_cast<QScatterDataItem*>(other));
}

void* QScatterDataItem_NewQScatterDataItem2(void* position)
{
	return new QScatterDataItem(*static_cast<QVector3D*>(position));
}

void* QScatterDataItem_NewQScatterDataItem3(void* position, void* rotation)
{
	return new QScatterDataItem(*static_cast<QVector3D*>(position), *static_cast<QQuaternion*>(rotation));
}

void QScatterDataItem_SetPosition(void* ptr, void* pos)
{
	static_cast<QScatterDataItem*>(ptr)->setPosition(*static_cast<QVector3D*>(pos));
}

void QScatterDataItem_SetRotation(void* ptr, void* rot)
{
	static_cast<QScatterDataItem*>(ptr)->setRotation(*static_cast<QQuaternion*>(rot));
}

void QScatterDataItem_SetX(void* ptr, float value)
{
	static_cast<QScatterDataItem*>(ptr)->setX(value);
}

void QScatterDataItem_SetY(void* ptr, float value)
{
	static_cast<QScatterDataItem*>(ptr)->setY(value);
}

void QScatterDataItem_SetZ(void* ptr, float value)
{
	static_cast<QScatterDataItem*>(ptr)->setZ(value);
}

void QScatterDataItem_DestroyQScatterDataItem(void* ptr)
{
	static_cast<QScatterDataItem*>(ptr)->~QScatterDataItem();
}

void* QScatterDataItem_Rotation(void* ptr)
{
	return new QQuaternion(static_cast<QScatterDataItem*>(ptr)->rotation());
}

void* QScatterDataItem_Position(void* ptr)
{
	return new QVector3D(static_cast<QScatterDataItem*>(ptr)->position());
}

float QScatterDataItem_X(void* ptr)
{
	return static_cast<QScatterDataItem*>(ptr)->x();
}

float QScatterDataItem_Y(void* ptr)
{
	return static_cast<QScatterDataItem*>(ptr)->y();
}

float QScatterDataItem_Z(void* ptr)
{
	return static_cast<QScatterDataItem*>(ptr)->z();
}

class MyQScatterDataProxy: public QScatterDataProxy
{
public:
	MyQScatterDataProxy(QObject *parent = Q_NULLPTR) : QScatterDataProxy(parent) {};
	void Signal_ArrayReset() { callbackQScatterDataProxy_ArrayReset(this); };
	void Signal_ItemCountChanged(int count) { callbackQScatterDataProxy_ItemCountChanged(this, count); };
	void Signal_ItemsAdded(int startIndex, int count) { callbackQScatterDataProxy_ItemsAdded(this, startIndex, count); };
	void Signal_ItemsChanged(int startIndex, int count) { callbackQScatterDataProxy_ItemsChanged(this, startIndex, count); };
	void Signal_ItemsInserted(int startIndex, int count) { callbackQScatterDataProxy_ItemsInserted(this, startIndex, count); };
	void Signal_ItemsRemoved(int startIndex, int count) { callbackQScatterDataProxy_ItemsRemoved(this, startIndex, count); };
	void Signal_SeriesChanged(QScatter3DSeries * series) { callbackQScatterDataProxy_SeriesChanged(this, series); };
	 ~MyQScatterDataProxy() { callbackQScatterDataProxy_DestroyQScatterDataProxy(this); };
};

void* QScatterDataProxy_NewQScatterDataProxy(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQScatterDataProxy(static_cast<QWindow*>(parent));
	} else {
		return new MyQScatterDataProxy(static_cast<QObject*>(parent));
	}
}

int QScatterDataProxy_AddItem(void* ptr, void* item)
{
	return static_cast<QScatterDataProxy*>(ptr)->addItem(*static_cast<QScatterDataItem*>(item));
}

void QScatterDataProxy_ConnectArrayReset(void* ptr)
{
	QObject::connect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)()>(&QScatterDataProxy::arrayReset), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)()>(&MyQScatterDataProxy::Signal_ArrayReset));
}

void QScatterDataProxy_DisconnectArrayReset(void* ptr)
{
	QObject::disconnect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)()>(&QScatterDataProxy::arrayReset), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)()>(&MyQScatterDataProxy::Signal_ArrayReset));
}

void QScatterDataProxy_ArrayReset(void* ptr)
{
	static_cast<QScatterDataProxy*>(ptr)->arrayReset();
}

void QScatterDataProxy_InsertItem(void* ptr, int index, void* item)
{
	static_cast<QScatterDataProxy*>(ptr)->insertItem(index, *static_cast<QScatterDataItem*>(item));
}

void QScatterDataProxy_ConnectItemCountChanged(void* ptr)
{
	QObject::connect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int)>(&QScatterDataProxy::itemCountChanged), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int)>(&MyQScatterDataProxy::Signal_ItemCountChanged));
}

void QScatterDataProxy_DisconnectItemCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int)>(&QScatterDataProxy::itemCountChanged), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int)>(&MyQScatterDataProxy::Signal_ItemCountChanged));
}

void QScatterDataProxy_ItemCountChanged(void* ptr, int count)
{
	static_cast<QScatterDataProxy*>(ptr)->itemCountChanged(count);
}

void QScatterDataProxy_ConnectItemsAdded(void* ptr)
{
	QObject::connect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int, int)>(&QScatterDataProxy::itemsAdded), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int, int)>(&MyQScatterDataProxy::Signal_ItemsAdded));
}

void QScatterDataProxy_DisconnectItemsAdded(void* ptr)
{
	QObject::disconnect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int, int)>(&QScatterDataProxy::itemsAdded), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int, int)>(&MyQScatterDataProxy::Signal_ItemsAdded));
}

void QScatterDataProxy_ItemsAdded(void* ptr, int startIndex, int count)
{
	static_cast<QScatterDataProxy*>(ptr)->itemsAdded(startIndex, count);
}

void QScatterDataProxy_ConnectItemsChanged(void* ptr)
{
	QObject::connect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int, int)>(&QScatterDataProxy::itemsChanged), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int, int)>(&MyQScatterDataProxy::Signal_ItemsChanged));
}

void QScatterDataProxy_DisconnectItemsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int, int)>(&QScatterDataProxy::itemsChanged), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int, int)>(&MyQScatterDataProxy::Signal_ItemsChanged));
}

void QScatterDataProxy_ItemsChanged(void* ptr, int startIndex, int count)
{
	static_cast<QScatterDataProxy*>(ptr)->itemsChanged(startIndex, count);
}

void QScatterDataProxy_ConnectItemsInserted(void* ptr)
{
	QObject::connect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int, int)>(&QScatterDataProxy::itemsInserted), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int, int)>(&MyQScatterDataProxy::Signal_ItemsInserted));
}

void QScatterDataProxy_DisconnectItemsInserted(void* ptr)
{
	QObject::disconnect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int, int)>(&QScatterDataProxy::itemsInserted), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int, int)>(&MyQScatterDataProxy::Signal_ItemsInserted));
}

void QScatterDataProxy_ItemsInserted(void* ptr, int startIndex, int count)
{
	static_cast<QScatterDataProxy*>(ptr)->itemsInserted(startIndex, count);
}

void QScatterDataProxy_ConnectItemsRemoved(void* ptr)
{
	QObject::connect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int, int)>(&QScatterDataProxy::itemsRemoved), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int, int)>(&MyQScatterDataProxy::Signal_ItemsRemoved));
}

void QScatterDataProxy_DisconnectItemsRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(int, int)>(&QScatterDataProxy::itemsRemoved), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(int, int)>(&MyQScatterDataProxy::Signal_ItemsRemoved));
}

void QScatterDataProxy_ItemsRemoved(void* ptr, int startIndex, int count)
{
	static_cast<QScatterDataProxy*>(ptr)->itemsRemoved(startIndex, count);
}

void QScatterDataProxy_RemoveItems(void* ptr, int index, int removeCount)
{
	static_cast<QScatterDataProxy*>(ptr)->removeItems(index, removeCount);
}

void QScatterDataProxy_ConnectSeriesChanged(void* ptr)
{
	QObject::connect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(QScatter3DSeries *)>(&QScatterDataProxy::seriesChanged), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(QScatter3DSeries *)>(&MyQScatterDataProxy::Signal_SeriesChanged));
}

void QScatterDataProxy_DisconnectSeriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScatterDataProxy*>(ptr), static_cast<void (QScatterDataProxy::*)(QScatter3DSeries *)>(&QScatterDataProxy::seriesChanged), static_cast<MyQScatterDataProxy*>(ptr), static_cast<void (MyQScatterDataProxy::*)(QScatter3DSeries *)>(&MyQScatterDataProxy::Signal_SeriesChanged));
}

void QScatterDataProxy_SeriesChanged(void* ptr, void* series)
{
	static_cast<QScatterDataProxy*>(ptr)->seriesChanged(static_cast<QScatter3DSeries*>(series));
}

void QScatterDataProxy_SetItem(void* ptr, int index, void* item)
{
	static_cast<QScatterDataProxy*>(ptr)->setItem(index, *static_cast<QScatterDataItem*>(item));
}

void QScatterDataProxy_DestroyQScatterDataProxy(void* ptr)
{
	static_cast<QScatterDataProxy*>(ptr)->~QScatterDataProxy();
}

void QScatterDataProxy_DestroyQScatterDataProxyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QScatterDataProxy_Series(void* ptr)
{
	return static_cast<QScatterDataProxy*>(ptr)->series();
}

void* QScatterDataProxy_ItemAt(void* ptr, int index)
{
	return const_cast<QScatterDataItem*>(static_cast<QScatterDataProxy*>(ptr)->itemAt(index));
}

int QScatterDataProxy_ItemCount(void* ptr)
{
	return static_cast<QScatterDataProxy*>(ptr)->itemCount();
}

class MyQSurface3DSeries: public QSurface3DSeries
{
public:
	MyQSurface3DSeries(QObject *parent = Q_NULLPTR) : QSurface3DSeries(parent) {};
	MyQSurface3DSeries(QSurfaceDataProxy *dataProxy, QObject *parent = Q_NULLPTR) : QSurface3DSeries(dataProxy, parent) {};
	void Signal_SelectedPointChanged(const QPoint & position) { callbackQSurface3DSeries_SelectedPointChanged(this, const_cast<QPoint*>(&position)); };
	void Signal_DataProxyChanged(QSurfaceDataProxy * proxy) { callbackQSurface3DSeries_DataProxyChanged(this, proxy); };
	void Signal_DrawModeChanged(QSurface3DSeries::DrawFlags mode) { callbackQSurface3DSeries_DrawModeChanged(this, mode); };
	void Signal_FlatShadingEnabledChanged(bool enable) { callbackQSurface3DSeries_FlatShadingEnabledChanged(this, enable); };
	void Signal_FlatShadingSupportedChanged(bool enable) { callbackQSurface3DSeries_FlatShadingSupportedChanged(this, enable); };
	void Signal_TextureChanged(const QImage & image) { callbackQSurface3DSeries_TextureChanged(this, const_cast<QImage*>(&image)); };
	void Signal_TextureFileChanged(const QString & filename) { QByteArray t08deae = filename.toUtf8(); QtDataVisualization_PackedString filenamePacked = { const_cast<char*>(t08deae.prepend("WHITESPACE").constData()+10), t08deae.size()-10 };callbackQSurface3DSeries_TextureFileChanged(this, filenamePacked); };
	 ~MyQSurface3DSeries() { callbackQSurface3DSeries_DestroyQSurface3DSeries(this); };
};

void QSurface3DSeries_ConnectSelectedPointChanged(void* ptr)
{
	QObject::connect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(const QPoint &)>(&QSurface3DSeries::selectedPointChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(const QPoint &)>(&MyQSurface3DSeries::Signal_SelectedPointChanged));
}

void QSurface3DSeries_DisconnectSelectedPointChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(const QPoint &)>(&QSurface3DSeries::selectedPointChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(const QPoint &)>(&MyQSurface3DSeries::Signal_SelectedPointChanged));
}

void QSurface3DSeries_SelectedPointChanged(void* ptr, void* position)
{
	static_cast<QSurface3DSeries*>(ptr)->selectedPointChanged(*static_cast<QPoint*>(position));
}

void QSurface3DSeries_SetDataProxy(void* ptr, void* proxy)
{
	static_cast<QSurface3DSeries*>(ptr)->setDataProxy(static_cast<QSurfaceDataProxy*>(proxy));
}

void* QSurface3DSeries_QSurface3DSeries_InvalidSelectionPosition()
{
	return ({ QPoint tmpValue = QSurface3DSeries::invalidSelectionPosition(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QSurface3DSeries_NewQSurface3DSeries(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QWindow*>(parent));
	} else {
		return new MyQSurface3DSeries(static_cast<QObject*>(parent));
	}
}

void* QSurface3DSeries_NewQSurface3DSeries2(void* dataProxy, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QWindow*>(parent));
	} else {
		return new MyQSurface3DSeries(static_cast<QSurfaceDataProxy*>(dataProxy), static_cast<QObject*>(parent));
	}
}

void QSurface3DSeries_ConnectDataProxyChanged(void* ptr)
{
	QObject::connect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(QSurfaceDataProxy *)>(&QSurface3DSeries::dataProxyChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(QSurfaceDataProxy *)>(&MyQSurface3DSeries::Signal_DataProxyChanged));
}

void QSurface3DSeries_DisconnectDataProxyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(QSurfaceDataProxy *)>(&QSurface3DSeries::dataProxyChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(QSurfaceDataProxy *)>(&MyQSurface3DSeries::Signal_DataProxyChanged));
}

void QSurface3DSeries_DataProxyChanged(void* ptr, void* proxy)
{
	static_cast<QSurface3DSeries*>(ptr)->dataProxyChanged(static_cast<QSurfaceDataProxy*>(proxy));
}

void QSurface3DSeries_ConnectDrawModeChanged(void* ptr)
{
	QObject::connect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(QSurface3DSeries::DrawFlags)>(&QSurface3DSeries::drawModeChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(QSurface3DSeries::DrawFlags)>(&MyQSurface3DSeries::Signal_DrawModeChanged));
}

void QSurface3DSeries_DisconnectDrawModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(QSurface3DSeries::DrawFlags)>(&QSurface3DSeries::drawModeChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(QSurface3DSeries::DrawFlags)>(&MyQSurface3DSeries::Signal_DrawModeChanged));
}

void QSurface3DSeries_DrawModeChanged(void* ptr, long long mode)
{
	static_cast<QSurface3DSeries*>(ptr)->drawModeChanged(static_cast<QSurface3DSeries::DrawFlag>(mode));
}

void QSurface3DSeries_ConnectFlatShadingEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(bool)>(&QSurface3DSeries::flatShadingEnabledChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(bool)>(&MyQSurface3DSeries::Signal_FlatShadingEnabledChanged));
}

void QSurface3DSeries_DisconnectFlatShadingEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(bool)>(&QSurface3DSeries::flatShadingEnabledChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(bool)>(&MyQSurface3DSeries::Signal_FlatShadingEnabledChanged));
}

void QSurface3DSeries_FlatShadingEnabledChanged(void* ptr, char enable)
{
	static_cast<QSurface3DSeries*>(ptr)->flatShadingEnabledChanged(enable != 0);
}

void QSurface3DSeries_ConnectFlatShadingSupportedChanged(void* ptr)
{
	QObject::connect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(bool)>(&QSurface3DSeries::flatShadingSupportedChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(bool)>(&MyQSurface3DSeries::Signal_FlatShadingSupportedChanged));
}

void QSurface3DSeries_DisconnectFlatShadingSupportedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(bool)>(&QSurface3DSeries::flatShadingSupportedChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(bool)>(&MyQSurface3DSeries::Signal_FlatShadingSupportedChanged));
}

void QSurface3DSeries_FlatShadingSupportedChanged(void* ptr, char enable)
{
	static_cast<QSurface3DSeries*>(ptr)->flatShadingSupportedChanged(enable != 0);
}

void QSurface3DSeries_SetDrawMode(void* ptr, long long mode)
{
	static_cast<QSurface3DSeries*>(ptr)->setDrawMode(static_cast<QSurface3DSeries::DrawFlag>(mode));
}

void QSurface3DSeries_SetFlatShadingEnabled(void* ptr, char enabled)
{
	static_cast<QSurface3DSeries*>(ptr)->setFlatShadingEnabled(enabled != 0);
}

void QSurface3DSeries_SetSelectedPoint(void* ptr, void* position)
{
	static_cast<QSurface3DSeries*>(ptr)->setSelectedPoint(*static_cast<QPoint*>(position));
}

void QSurface3DSeries_SetTexture(void* ptr, void* texture)
{
	static_cast<QSurface3DSeries*>(ptr)->setTexture(*static_cast<QImage*>(texture));
}

void QSurface3DSeries_SetTextureFile(void* ptr, struct QtDataVisualization_PackedString filename)
{
	static_cast<QSurface3DSeries*>(ptr)->setTextureFile(QString::fromUtf8(filename.data, filename.len));
}

void QSurface3DSeries_ConnectTextureChanged(void* ptr)
{
	QObject::connect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(const QImage &)>(&QSurface3DSeries::textureChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(const QImage &)>(&MyQSurface3DSeries::Signal_TextureChanged));
}

void QSurface3DSeries_DisconnectTextureChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(const QImage &)>(&QSurface3DSeries::textureChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(const QImage &)>(&MyQSurface3DSeries::Signal_TextureChanged));
}

void QSurface3DSeries_TextureChanged(void* ptr, void* image)
{
	static_cast<QSurface3DSeries*>(ptr)->textureChanged(*static_cast<QImage*>(image));
}

void QSurface3DSeries_ConnectTextureFileChanged(void* ptr)
{
	QObject::connect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(const QString &)>(&QSurface3DSeries::textureFileChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(const QString &)>(&MyQSurface3DSeries::Signal_TextureFileChanged));
}

void QSurface3DSeries_DisconnectTextureFileChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurface3DSeries*>(ptr), static_cast<void (QSurface3DSeries::*)(const QString &)>(&QSurface3DSeries::textureFileChanged), static_cast<MyQSurface3DSeries*>(ptr), static_cast<void (MyQSurface3DSeries::*)(const QString &)>(&MyQSurface3DSeries::Signal_TextureFileChanged));
}

void QSurface3DSeries_TextureFileChanged(void* ptr, struct QtDataVisualization_PackedString filename)
{
	static_cast<QSurface3DSeries*>(ptr)->textureFileChanged(QString::fromUtf8(filename.data, filename.len));
}

void QSurface3DSeries_DestroyQSurface3DSeries(void* ptr)
{
	static_cast<QSurface3DSeries*>(ptr)->~QSurface3DSeries();
}

void QSurface3DSeries_DestroyQSurface3DSeriesDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSurface3DSeries_Texture(void* ptr)
{
	return new QImage(static_cast<QSurface3DSeries*>(ptr)->texture());
}

void* QSurface3DSeries_SelectedPoint(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QSurface3DSeries*>(ptr)->selectedPoint(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

struct QtDataVisualization_PackedString QSurface3DSeries_TextureFile(void* ptr)
{
	return ({ QByteArray tf3d489 = static_cast<QSurface3DSeries*>(ptr)->textureFile().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(tf3d489.prepend("WHITESPACE").constData()+10), tf3d489.size()-10 }; });
}

long long QSurface3DSeries_DrawMode(void* ptr)
{
	return static_cast<QSurface3DSeries*>(ptr)->drawMode();
}

void* QSurface3DSeries_DataProxy(void* ptr)
{
	return static_cast<QSurface3DSeries*>(ptr)->dataProxy();
}

char QSurface3DSeries_IsFlatShadingEnabled(void* ptr)
{
	return static_cast<QSurface3DSeries*>(ptr)->isFlatShadingEnabled();
}

char QSurface3DSeries_IsFlatShadingSupported(void* ptr)
{
	return static_cast<QSurface3DSeries*>(ptr)->isFlatShadingSupported();
}

void* QSurfaceDataItem_NewQSurfaceDataItem()
{
	return new QSurfaceDataItem();
}

void* QSurfaceDataItem_NewQSurfaceDataItem3(void* other)
{
	return new QSurfaceDataItem(*static_cast<QSurfaceDataItem*>(other));
}

void* QSurfaceDataItem_NewQSurfaceDataItem2(void* position)
{
	return new QSurfaceDataItem(*static_cast<QVector3D*>(position));
}

void QSurfaceDataItem_SetPosition(void* ptr, void* pos)
{
	static_cast<QSurfaceDataItem*>(ptr)->setPosition(*static_cast<QVector3D*>(pos));
}

void QSurfaceDataItem_SetX(void* ptr, float value)
{
	static_cast<QSurfaceDataItem*>(ptr)->setX(value);
}

void QSurfaceDataItem_SetY(void* ptr, float value)
{
	static_cast<QSurfaceDataItem*>(ptr)->setY(value);
}

void QSurfaceDataItem_SetZ(void* ptr, float value)
{
	static_cast<QSurfaceDataItem*>(ptr)->setZ(value);
}

void QSurfaceDataItem_DestroyQSurfaceDataItem(void* ptr)
{
	static_cast<QSurfaceDataItem*>(ptr)->~QSurfaceDataItem();
}

void* QSurfaceDataItem_Position(void* ptr)
{
	return new QVector3D(static_cast<QSurfaceDataItem*>(ptr)->position());
}

float QSurfaceDataItem_X(void* ptr)
{
	return static_cast<QSurfaceDataItem*>(ptr)->x();
}

float QSurfaceDataItem_Y(void* ptr)
{
	return static_cast<QSurfaceDataItem*>(ptr)->y();
}

float QSurfaceDataItem_Z(void* ptr)
{
	return static_cast<QSurfaceDataItem*>(ptr)->z();
}

class MyQSurfaceDataProxy: public QSurfaceDataProxy
{
public:
	MyQSurfaceDataProxy(QObject *parent = Q_NULLPTR) : QSurfaceDataProxy(parent) {};
	void Signal_ArrayReset() { callbackQSurfaceDataProxy_ArrayReset(this); };
	void Signal_ColumnCountChanged(int count) { callbackQSurfaceDataProxy_ColumnCountChanged(this, count); };
	void Signal_ItemChanged(int rowIndex, int columnIndex) { callbackQSurfaceDataProxy_ItemChanged(this, rowIndex, columnIndex); };
	void Signal_RowCountChanged(int count) { callbackQSurfaceDataProxy_RowCountChanged(this, count); };
	void Signal_RowsAdded(int startIndex, int count) { callbackQSurfaceDataProxy_RowsAdded(this, startIndex, count); };
	void Signal_RowsChanged(int startIndex, int count) { callbackQSurfaceDataProxy_RowsChanged(this, startIndex, count); };
	void Signal_RowsInserted(int startIndex, int count) { callbackQSurfaceDataProxy_RowsInserted(this, startIndex, count); };
	void Signal_RowsRemoved(int startIndex, int count) { callbackQSurfaceDataProxy_RowsRemoved(this, startIndex, count); };
	void Signal_SeriesChanged(QSurface3DSeries * series) { callbackQSurfaceDataProxy_SeriesChanged(this, series); };
	 ~MyQSurfaceDataProxy() { callbackQSurfaceDataProxy_DestroyQSurfaceDataProxy(this); };
};

void* QSurfaceDataProxy_NewQSurfaceDataProxy(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSurfaceDataProxy(static_cast<QWindow*>(parent));
	} else {
		return new MyQSurfaceDataProxy(static_cast<QObject*>(parent));
	}
}

void QSurfaceDataProxy_ConnectArrayReset(void* ptr)
{
	QObject::connect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)()>(&QSurfaceDataProxy::arrayReset), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)()>(&MyQSurfaceDataProxy::Signal_ArrayReset));
}

void QSurfaceDataProxy_DisconnectArrayReset(void* ptr)
{
	QObject::disconnect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)()>(&QSurfaceDataProxy::arrayReset), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)()>(&MyQSurfaceDataProxy::Signal_ArrayReset));
}

void QSurfaceDataProxy_ArrayReset(void* ptr)
{
	static_cast<QSurfaceDataProxy*>(ptr)->arrayReset();
}

void QSurfaceDataProxy_ConnectColumnCountChanged(void* ptr)
{
	QObject::connect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int)>(&QSurfaceDataProxy::columnCountChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int)>(&MyQSurfaceDataProxy::Signal_ColumnCountChanged));
}

void QSurfaceDataProxy_DisconnectColumnCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int)>(&QSurfaceDataProxy::columnCountChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int)>(&MyQSurfaceDataProxy::Signal_ColumnCountChanged));
}

void QSurfaceDataProxy_ColumnCountChanged(void* ptr, int count)
{
	static_cast<QSurfaceDataProxy*>(ptr)->columnCountChanged(count);
}

void QSurfaceDataProxy_ConnectItemChanged(void* ptr)
{
	QObject::connect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::itemChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_ItemChanged));
}

void QSurfaceDataProxy_DisconnectItemChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::itemChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_ItemChanged));
}

void QSurfaceDataProxy_ItemChanged(void* ptr, int rowIndex, int columnIndex)
{
	static_cast<QSurfaceDataProxy*>(ptr)->itemChanged(rowIndex, columnIndex);
}

void QSurfaceDataProxy_RemoveRows(void* ptr, int rowIndex, int removeCount)
{
	static_cast<QSurfaceDataProxy*>(ptr)->removeRows(rowIndex, removeCount);
}

void QSurfaceDataProxy_ConnectRowCountChanged(void* ptr)
{
	QObject::connect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int)>(&QSurfaceDataProxy::rowCountChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int)>(&MyQSurfaceDataProxy::Signal_RowCountChanged));
}

void QSurfaceDataProxy_DisconnectRowCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int)>(&QSurfaceDataProxy::rowCountChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int)>(&MyQSurfaceDataProxy::Signal_RowCountChanged));
}

void QSurfaceDataProxy_RowCountChanged(void* ptr, int count)
{
	static_cast<QSurfaceDataProxy*>(ptr)->rowCountChanged(count);
}

void QSurfaceDataProxy_ConnectRowsAdded(void* ptr)
{
	QObject::connect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::rowsAdded), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_RowsAdded));
}

void QSurfaceDataProxy_DisconnectRowsAdded(void* ptr)
{
	QObject::disconnect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::rowsAdded), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_RowsAdded));
}

void QSurfaceDataProxy_RowsAdded(void* ptr, int startIndex, int count)
{
	static_cast<QSurfaceDataProxy*>(ptr)->rowsAdded(startIndex, count);
}

void QSurfaceDataProxy_ConnectRowsChanged(void* ptr)
{
	QObject::connect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::rowsChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_RowsChanged));
}

void QSurfaceDataProxy_DisconnectRowsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::rowsChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_RowsChanged));
}

void QSurfaceDataProxy_RowsChanged(void* ptr, int startIndex, int count)
{
	static_cast<QSurfaceDataProxy*>(ptr)->rowsChanged(startIndex, count);
}

void QSurfaceDataProxy_ConnectRowsInserted(void* ptr)
{
	QObject::connect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::rowsInserted), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_RowsInserted));
}

void QSurfaceDataProxy_DisconnectRowsInserted(void* ptr)
{
	QObject::disconnect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::rowsInserted), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_RowsInserted));
}

void QSurfaceDataProxy_RowsInserted(void* ptr, int startIndex, int count)
{
	static_cast<QSurfaceDataProxy*>(ptr)->rowsInserted(startIndex, count);
}

void QSurfaceDataProxy_ConnectRowsRemoved(void* ptr)
{
	QObject::connect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::rowsRemoved), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_RowsRemoved));
}

void QSurfaceDataProxy_DisconnectRowsRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(int, int)>(&QSurfaceDataProxy::rowsRemoved), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(int, int)>(&MyQSurfaceDataProxy::Signal_RowsRemoved));
}

void QSurfaceDataProxy_RowsRemoved(void* ptr, int startIndex, int count)
{
	static_cast<QSurfaceDataProxy*>(ptr)->rowsRemoved(startIndex, count);
}

void QSurfaceDataProxy_ConnectSeriesChanged(void* ptr)
{
	QObject::connect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(QSurface3DSeries *)>(&QSurfaceDataProxy::seriesChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(QSurface3DSeries *)>(&MyQSurfaceDataProxy::Signal_SeriesChanged));
}

void QSurfaceDataProxy_DisconnectSeriesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSurfaceDataProxy*>(ptr), static_cast<void (QSurfaceDataProxy::*)(QSurface3DSeries *)>(&QSurfaceDataProxy::seriesChanged), static_cast<MyQSurfaceDataProxy*>(ptr), static_cast<void (MyQSurfaceDataProxy::*)(QSurface3DSeries *)>(&MyQSurfaceDataProxy::Signal_SeriesChanged));
}

void QSurfaceDataProxy_SeriesChanged(void* ptr, void* series)
{
	static_cast<QSurfaceDataProxy*>(ptr)->seriesChanged(static_cast<QSurface3DSeries*>(series));
}

void QSurfaceDataProxy_SetItem2(void* ptr, void* position, void* item)
{
	static_cast<QSurfaceDataProxy*>(ptr)->setItem(*static_cast<QPoint*>(position), *static_cast<QSurfaceDataItem*>(item));
}

void QSurfaceDataProxy_SetItem(void* ptr, int rowIndex, int columnIndex, void* item)
{
	static_cast<QSurfaceDataProxy*>(ptr)->setItem(rowIndex, columnIndex, *static_cast<QSurfaceDataItem*>(item));
}

void QSurfaceDataProxy_DestroyQSurfaceDataProxy(void* ptr)
{
	static_cast<QSurfaceDataProxy*>(ptr)->~QSurfaceDataProxy();
}

void QSurfaceDataProxy_DestroyQSurfaceDataProxyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSurfaceDataProxy_Series(void* ptr)
{
	return static_cast<QSurfaceDataProxy*>(ptr)->series();
}

void* QSurfaceDataProxy_ItemAt2(void* ptr, void* position)
{
	return const_cast<QSurfaceDataItem*>(static_cast<QSurfaceDataProxy*>(ptr)->itemAt(*static_cast<QPoint*>(position)));
}

void* QSurfaceDataProxy_ItemAt(void* ptr, int rowIndex, int columnIndex)
{
	return const_cast<QSurfaceDataItem*>(static_cast<QSurfaceDataProxy*>(ptr)->itemAt(rowIndex, columnIndex));
}

int QSurfaceDataProxy_ColumnCount(void* ptr)
{
	return static_cast<QSurfaceDataProxy*>(ptr)->columnCount();
}

int QSurfaceDataProxy_RowCount(void* ptr)
{
	return static_cast<QSurfaceDataProxy*>(ptr)->rowCount();
}

class MyQTouch3DInputHandler: public QTouch3DInputHandler
{
public:
	MyQTouch3DInputHandler(QObject *parent = Q_NULLPTR) : QTouch3DInputHandler(parent) {};
	void touchEvent(QTouchEvent * event) { callbackQTouch3DInputHandler_TouchEvent(this, event); };
	 ~MyQTouch3DInputHandler() { callbackQTouch3DInputHandler_DestroyQTouch3DInputHandler(this); };
};

void* QTouch3DInputHandler_NewQTouch3DInputHandler(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTouch3DInputHandler(static_cast<QWindow*>(parent));
	} else {
		return new MyQTouch3DInputHandler(static_cast<QObject*>(parent));
	}
}

void QTouch3DInputHandler_TouchEvent(void* ptr, void* event)
{
	static_cast<QTouch3DInputHandler*>(ptr)->touchEvent(static_cast<QTouchEvent*>(event));
}

void QTouch3DInputHandler_TouchEventDefault(void* ptr, void* event)
{
		static_cast<QTouch3DInputHandler*>(ptr)->QTouch3DInputHandler::touchEvent(static_cast<QTouchEvent*>(event));
}

void QTouch3DInputHandler_DestroyQTouch3DInputHandler(void* ptr)
{
	static_cast<QTouch3DInputHandler*>(ptr)->~QTouch3DInputHandler();
}

void QTouch3DInputHandler_DestroyQTouch3DInputHandlerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQValue3DAxis: public QValue3DAxis
{
public:
	MyQValue3DAxis(QObject *parent = Q_NULLPTR) : QValue3DAxis(parent) {};
	void Signal_FormatterChanged(QValue3DAxisFormatter * formatter) { callbackQValue3DAxis_FormatterChanged(this, formatter); };
	void Signal_LabelFormatChanged(const QString & format) { QByteArray t785987 = format.toUtf8(); QtDataVisualization_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQValue3DAxis_LabelFormatChanged(this, formatPacked); };
	void Signal_ReversedChanged(bool enable) { callbackQValue3DAxis_ReversedChanged(this, enable); };
	void Signal_SegmentCountChanged(int count) { callbackQValue3DAxis_SegmentCountChanged(this, count); };
	void Signal_SubSegmentCountChanged(int count) { callbackQValue3DAxis_SubSegmentCountChanged(this, count); };
	 ~MyQValue3DAxis() { callbackQValue3DAxis_DestroyQValue3DAxis(this); };
};

void* QValue3DAxis_NewQValue3DAxis(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxis(static_cast<QWindow*>(parent));
	} else {
		return new MyQValue3DAxis(static_cast<QObject*>(parent));
	}
}

void QValue3DAxis_ConnectFormatterChanged(void* ptr)
{
	QObject::connect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(QValue3DAxisFormatter *)>(&QValue3DAxis::formatterChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(QValue3DAxisFormatter *)>(&MyQValue3DAxis::Signal_FormatterChanged));
}

void QValue3DAxis_DisconnectFormatterChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(QValue3DAxisFormatter *)>(&QValue3DAxis::formatterChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(QValue3DAxisFormatter *)>(&MyQValue3DAxis::Signal_FormatterChanged));
}

void QValue3DAxis_FormatterChanged(void* ptr, void* formatter)
{
	static_cast<QValue3DAxis*>(ptr)->formatterChanged(static_cast<QValue3DAxisFormatter*>(formatter));
}

void QValue3DAxis_ConnectLabelFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(const QString &)>(&QValue3DAxis::labelFormatChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(const QString &)>(&MyQValue3DAxis::Signal_LabelFormatChanged));
}

void QValue3DAxis_DisconnectLabelFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(const QString &)>(&QValue3DAxis::labelFormatChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(const QString &)>(&MyQValue3DAxis::Signal_LabelFormatChanged));
}

void QValue3DAxis_LabelFormatChanged(void* ptr, struct QtDataVisualization_PackedString format)
{
	static_cast<QValue3DAxis*>(ptr)->labelFormatChanged(QString::fromUtf8(format.data, format.len));
}

void QValue3DAxis_ConnectReversedChanged(void* ptr)
{
	QObject::connect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(bool)>(&QValue3DAxis::reversedChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(bool)>(&MyQValue3DAxis::Signal_ReversedChanged));
}

void QValue3DAxis_DisconnectReversedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(bool)>(&QValue3DAxis::reversedChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(bool)>(&MyQValue3DAxis::Signal_ReversedChanged));
}

void QValue3DAxis_ReversedChanged(void* ptr, char enable)
{
	static_cast<QValue3DAxis*>(ptr)->reversedChanged(enable != 0);
}

void QValue3DAxis_ConnectSegmentCountChanged(void* ptr)
{
	QObject::connect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(int)>(&QValue3DAxis::segmentCountChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(int)>(&MyQValue3DAxis::Signal_SegmentCountChanged));
}

void QValue3DAxis_DisconnectSegmentCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(int)>(&QValue3DAxis::segmentCountChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(int)>(&MyQValue3DAxis::Signal_SegmentCountChanged));
}

void QValue3DAxis_SegmentCountChanged(void* ptr, int count)
{
	static_cast<QValue3DAxis*>(ptr)->segmentCountChanged(count);
}

void QValue3DAxis_SetFormatter(void* ptr, void* formatter)
{
	static_cast<QValue3DAxis*>(ptr)->setFormatter(static_cast<QValue3DAxisFormatter*>(formatter));
}

void QValue3DAxis_SetLabelFormat(void* ptr, struct QtDataVisualization_PackedString format)
{
	static_cast<QValue3DAxis*>(ptr)->setLabelFormat(QString::fromUtf8(format.data, format.len));
}

void QValue3DAxis_SetReversed(void* ptr, char enable)
{
	static_cast<QValue3DAxis*>(ptr)->setReversed(enable != 0);
}

void QValue3DAxis_SetSegmentCount(void* ptr, int count)
{
	static_cast<QValue3DAxis*>(ptr)->setSegmentCount(count);
}

void QValue3DAxis_SetSubSegmentCount(void* ptr, int count)
{
	static_cast<QValue3DAxis*>(ptr)->setSubSegmentCount(count);
}

void QValue3DAxis_ConnectSubSegmentCountChanged(void* ptr)
{
	QObject::connect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(int)>(&QValue3DAxis::subSegmentCountChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(int)>(&MyQValue3DAxis::Signal_SubSegmentCountChanged));
}

void QValue3DAxis_DisconnectSubSegmentCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValue3DAxis*>(ptr), static_cast<void (QValue3DAxis::*)(int)>(&QValue3DAxis::subSegmentCountChanged), static_cast<MyQValue3DAxis*>(ptr), static_cast<void (MyQValue3DAxis::*)(int)>(&MyQValue3DAxis::Signal_SubSegmentCountChanged));
}

void QValue3DAxis_SubSegmentCountChanged(void* ptr, int count)
{
	static_cast<QValue3DAxis*>(ptr)->subSegmentCountChanged(count);
}

void QValue3DAxis_DestroyQValue3DAxis(void* ptr)
{
	static_cast<QValue3DAxis*>(ptr)->~QValue3DAxis();
}

void QValue3DAxis_DestroyQValue3DAxisDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtDataVisualization_PackedString QValue3DAxis_LabelFormat(void* ptr)
{
	return ({ QByteArray teadd17 = static_cast<QValue3DAxis*>(ptr)->labelFormat().toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(teadd17.prepend("WHITESPACE").constData()+10), teadd17.size()-10 }; });
}

void* QValue3DAxis_Formatter(void* ptr)
{
	return static_cast<QValue3DAxis*>(ptr)->formatter();
}

char QValue3DAxis_Reversed(void* ptr)
{
	return static_cast<QValue3DAxis*>(ptr)->reversed();
}

int QValue3DAxis_SegmentCount(void* ptr)
{
	return static_cast<QValue3DAxis*>(ptr)->segmentCount();
}

int QValue3DAxis_SubSegmentCount(void* ptr)
{
	return static_cast<QValue3DAxis*>(ptr)->subSegmentCount();
}

class MyQValue3DAxisFormatter: public QValue3DAxisFormatter
{
public:
	MyQValue3DAxisFormatter(QObject *parent = Q_NULLPTR) : QValue3DAxisFormatter(parent) {};
	void recalculate() { callbackQValue3DAxisFormatter_Recalculate(this); };
	 ~MyQValue3DAxisFormatter() { callbackQValue3DAxisFormatter_DestroyQValue3DAxisFormatter(this); };
	QString stringForValue(qreal value, const QString & format) const { QByteArray t785987 = format.toUtf8(); QtDataVisualization_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };return ({ QtDataVisualization_PackedString tempVal = callbackQValue3DAxisFormatter_StringForValue(const_cast<void*>(static_cast<const void*>(this)), value, formatPacked); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QValue3DAxisFormatter * createNewInstance() const { return static_cast<QValue3DAxisFormatter*>(callbackQValue3DAxisFormatter_CreateNewInstance(const_cast<void*>(static_cast<const void*>(this)))); };
	float positionAt(float value) const { return callbackQValue3DAxisFormatter_PositionAt(const_cast<void*>(static_cast<const void*>(this)), value); };
	float valueAt(float position) const { return callbackQValue3DAxisFormatter_ValueAt(const_cast<void*>(static_cast<const void*>(this)), position); };
	void populateCopy(QValue3DAxisFormatter & copy) const { callbackQValue3DAxisFormatter_PopulateCopy(const_cast<void*>(static_cast<const void*>(this)), static_cast<QValue3DAxisFormatter*>(&copy)); };
};

void* QValue3DAxisFormatter_NewQValue3DAxisFormatter(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQValue3DAxisFormatter(static_cast<QWindow*>(parent));
	} else {
		return new MyQValue3DAxisFormatter(static_cast<QObject*>(parent));
	}
}

void QValue3DAxisFormatter_MarkDirty(void* ptr, char labelsChange)
{
	static_cast<QValue3DAxisFormatter*>(ptr)->markDirty(labelsChange != 0);
}

void QValue3DAxisFormatter_Recalculate(void* ptr)
{
	static_cast<QValue3DAxisFormatter*>(ptr)->recalculate();
}

void QValue3DAxisFormatter_RecalculateDefault(void* ptr)
{
		static_cast<QValue3DAxisFormatter*>(ptr)->QValue3DAxisFormatter::recalculate();
}

void QValue3DAxisFormatter_SetAllowNegatives(void* ptr, char allow)
{
	static_cast<QValue3DAxisFormatter*>(ptr)->setAllowNegatives(allow != 0);
}

void QValue3DAxisFormatter_SetAllowZero(void* ptr, char allow)
{
	static_cast<QValue3DAxisFormatter*>(ptr)->setAllowZero(allow != 0);
}

void QValue3DAxisFormatter_SetLocale(void* ptr, void* locale)
{
	static_cast<QValue3DAxisFormatter*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QValue3DAxisFormatter_DestroyQValue3DAxisFormatter(void* ptr)
{
	static_cast<QValue3DAxisFormatter*>(ptr)->~QValue3DAxisFormatter();
}

void QValue3DAxisFormatter_DestroyQValue3DAxisFormatterDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QValue3DAxisFormatter_Locale(void* ptr)
{
	return new QLocale(static_cast<QValue3DAxisFormatter*>(ptr)->locale());
}

struct QtDataVisualization_PackedString QValue3DAxisFormatter_StringForValue(void* ptr, double value, struct QtDataVisualization_PackedString format)
{
	return ({ QByteArray t1b2a5f = static_cast<QValue3DAxisFormatter*>(ptr)->stringForValue(value, QString::fromUtf8(format.data, format.len)).toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(t1b2a5f.prepend("WHITESPACE").constData()+10), t1b2a5f.size()-10 }; });
}

struct QtDataVisualization_PackedString QValue3DAxisFormatter_StringForValueDefault(void* ptr, double value, struct QtDataVisualization_PackedString format)
{
		return ({ QByteArray te8d0b4 = static_cast<QValue3DAxisFormatter*>(ptr)->QValue3DAxisFormatter::stringForValue(value, QString::fromUtf8(format.data, format.len)).toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(te8d0b4.prepend("WHITESPACE").constData()+10), te8d0b4.size()-10 }; });
}

struct QtDataVisualization_PackedString QValue3DAxisFormatter_LabelStrings(void* ptr)
{
	return ({ QByteArray td5e8ca = static_cast<QValue3DAxisFormatter*>(ptr)->labelStrings().join("|").toUtf8(); QtDataVisualization_PackedString { const_cast<char*>(td5e8ca.prepend("WHITESPACE").constData()+10), td5e8ca.size()-10 }; });
}

void* QValue3DAxisFormatter_Axis(void* ptr)
{
	return static_cast<QValue3DAxisFormatter*>(ptr)->axis();
}

void* QValue3DAxisFormatter_CreateNewInstance(void* ptr)
{
	return static_cast<QValue3DAxisFormatter*>(ptr)->createNewInstance();
}

void* QValue3DAxisFormatter_CreateNewInstanceDefault(void* ptr)
{
		return static_cast<QValue3DAxisFormatter*>(ptr)->QValue3DAxisFormatter::createNewInstance();
}

struct QtDataVisualization_PackedList QValue3DAxisFormatter_GridPositions(void* ptr)
{
	return ({ QVector<float>* tmpValue = new QVector<float>(static_cast<QValue3DAxisFormatter*>(ptr)->gridPositions()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDataVisualization_PackedList QValue3DAxisFormatter_LabelPositions(void* ptr)
{
	return ({ QVector<float>* tmpValue = new QVector<float>(static_cast<QValue3DAxisFormatter*>(ptr)->labelPositions()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDataVisualization_PackedList QValue3DAxisFormatter_SubGridPositions(void* ptr)
{
	return ({ QVector<float>* tmpValue = new QVector<float>(static_cast<QValue3DAxisFormatter*>(ptr)->subGridPositions()); QtDataVisualization_PackedList { tmpValue, tmpValue->size() }; });
}

char QValue3DAxisFormatter_AllowNegatives(void* ptr)
{
	return static_cast<QValue3DAxisFormatter*>(ptr)->allowNegatives();
}

char QValue3DAxisFormatter_AllowZero(void* ptr)
{
	return static_cast<QValue3DAxisFormatter*>(ptr)->allowZero();
}

float QValue3DAxisFormatter_PositionAt(void* ptr, float value)
{
	return static_cast<QValue3DAxisFormatter*>(ptr)->positionAt(value);
}

float QValue3DAxisFormatter_PositionAtDefault(void* ptr, float value)
{
		return static_cast<QValue3DAxisFormatter*>(ptr)->QValue3DAxisFormatter::positionAt(value);
}

float QValue3DAxisFormatter_ValueAt(void* ptr, float position)
{
	return static_cast<QValue3DAxisFormatter*>(ptr)->valueAt(position);
}

float QValue3DAxisFormatter_ValueAtDefault(void* ptr, float position)
{
		return static_cast<QValue3DAxisFormatter*>(ptr)->QValue3DAxisFormatter::valueAt(position);
}

void QValue3DAxisFormatter_PopulateCopy(void* ptr, void* copy)
{
	static_cast<QValue3DAxisFormatter*>(ptr)->populateCopy(*static_cast<QValue3DAxisFormatter*>(copy));
}

void QValue3DAxisFormatter_PopulateCopyDefault(void* ptr, void* copy)
{
		static_cast<QValue3DAxisFormatter*>(ptr)->QValue3DAxisFormatter::populateCopy(*static_cast<QValue3DAxisFormatter*>(copy));
}

float QValue3DAxisFormatter___gridPositions_atList(void* ptr, int i)
{
	return static_cast<QVector<float>*>(ptr)->at(i);
}

void QValue3DAxisFormatter___gridPositions_setList(void* ptr, float i)
{
	static_cast<QVector<float>*>(ptr)->append(i);
}

void* QValue3DAxisFormatter___gridPositions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<float>;
}

float QValue3DAxisFormatter___labelPositions_atList(void* ptr, int i)
{
	return static_cast<QVector<float>*>(ptr)->at(i);
}

void QValue3DAxisFormatter___labelPositions_setList(void* ptr, float i)
{
	static_cast<QVector<float>*>(ptr)->append(i);
}

void* QValue3DAxisFormatter___labelPositions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<float>;
}

float QValue3DAxisFormatter___subGridPositions_atList(void* ptr, int i)
{
	return static_cast<QVector<float>*>(ptr)->at(i);
}

void QValue3DAxisFormatter___subGridPositions_setList(void* ptr, float i)
{
	static_cast<QVector<float>*>(ptr)->append(i);
}

void* QValue3DAxisFormatter___subGridPositions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<float>;
}

