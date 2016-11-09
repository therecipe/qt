// +build !minimal

#define protected public
#define private public

#include "multimedia.h"
#include "_cgo_export.h"

#include <QAbstractPlanarVideoBuffer>
#include <QAbstractVideoBuffer>
#include <QAbstractVideoFilter>
#include <QAbstractVideoSurface>
#include <QAction>
#include <QActionEvent>
#include <QAudio>
#include <QAudioBuffer>
#include <QAudioDecoder>
#include <QAudioDecoderControl>
#include <QAudioDeviceInfo>
#include <QAudioEncoderSettings>
#include <QAudioEncoderSettingsControl>
#include <QAudioFormat>
#include <QAudioInput>
#include <QAudioInputSelectorControl>
#include <QAudioOutput>
#include <QAudioOutputSelectorControl>
#include <QAudioProbe>
#include <QAudioRecorder>
#include <QAudioRoleControl>
#include <QByteArray>
#include <QCamera>
#include <QCameraCaptureBufferFormatControl>
#include <QCameraCaptureDestinationControl>
#include <QCameraControl>
#include <QCameraExposure>
#include <QCameraExposureControl>
#include <QCameraFeedbackControl>
#include <QCameraFlashControl>
#include <QCameraFocus>
#include <QCameraFocusControl>
#include <QCameraFocusZone>
#include <QCameraImageCapture>
#include <QCameraImageCaptureControl>
#include <QCameraImageProcessing>
#include <QCameraImageProcessingControl>
#include <QCameraInfo>
#include <QCameraInfoControl>
#include <QCameraLocksControl>
#include <QCameraViewfinder>
#include <QCameraViewfinderSettings>
#include <QCameraViewfinderSettingsControl>
#include <QCameraViewfinderSettingsControl2>
#include <QCameraZoomControl>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QGraphicsItem>
#include <QGraphicsScene>
#include <QGraphicsSceneContextMenuEvent>
#include <QGraphicsSceneDragDropEvent>
#include <QGraphicsSceneHoverEvent>
#include <QGraphicsSceneMouseEvent>
#include <QGraphicsSceneWheelEvent>
#include <QGraphicsVideoItem>
#include <QHideEvent>
#include <QIODevice>
#include <QImage>
#include <QImageEncoderControl>
#include <QImageEncoderSettings>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMediaAudioProbeControl>
#include <QMediaAvailabilityControl>
#include <QMediaBindableInterface>
#include <QMediaContainerControl>
#include <QMediaContent>
#include <QMediaControl>
#include <QMediaGaplessPlaybackControl>
#include <QMediaNetworkAccessControl>
#include <QMediaObject>
#include <QMediaPlayer>
#include <QMediaPlayerControl>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMediaRecorderControl>
#include <QMediaResource>
#include <QMediaService>
#include <QMediaServiceCameraInfoInterface>
#include <QMediaServiceDefaultDeviceInterface>
#include <QMediaServiceFeaturesInterface>
#include <QMediaServiceProviderPlugin>
#include <QMediaServiceSupportedDevicesInterface>
#include <QMediaServiceSupportedFormatsInterface>
#include <QMediaStreamsControl>
#include <QMediaTimeInterval>
#include <QMediaTimeRange>
#include <QMediaVideoProbeControl>
#include <QMetaDataReaderControl>
#include <QMetaDataWriterControl>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QMultimedia>
#include <QNetworkConfiguration>
#include <QNetworkRequest>
#include <QObject>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPoint>
#include <QPointF>
#include <QRadioData>
#include <QRadioDataControl>
#include <QRadioTuner>
#include <QRadioTunerControl>
#include <QRect>
#include <QRectF>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QSound>
#include <QSoundEffect>
#include <QString>
#include <QStringList>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionGraphicsItem>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QVideoDeviceSelectorControl>
#include <QVideoEncoderSettings>
#include <QVideoEncoderSettingsControl>
#include <QVideoFilterRunnable>
#include <QVideoFrame>
#include <QVideoProbe>
#include <QVideoRendererControl>
#include <QVideoSurfaceFormat>
#include <QVideoWidget>
#include <QVideoWidgetControl>
#include <QVideoWindowControl>
#include <QWheelEvent>
#include <QWidget>

class MyQAbstractPlanarVideoBuffer: public QAbstractPlanarVideoBuffer
{
public:
	 ~MyQAbstractPlanarVideoBuffer() { callbackQAbstractPlanarVideoBuffer_DestroyQAbstractPlanarVideoBuffer(this); };
	QVariant handle() const { return *static_cast<QVariant*>(callbackQAbstractPlanarVideoBuffer_Handle(const_cast<MyQAbstractPlanarVideoBuffer*>(this))); };
	MapMode mapMode() const { return static_cast<QAbstractVideoBuffer::MapMode>(callbackQAbstractPlanarVideoBuffer_MapMode(const_cast<MyQAbstractPlanarVideoBuffer*>(this))); };
	void release() { callbackQAbstractPlanarVideoBuffer_Release(this); };
	void unmap() { callbackQAbstractPlanarVideoBuffer_Unmap(this); };
};

void QAbstractPlanarVideoBuffer_DestroyQAbstractPlanarVideoBuffer(void* ptr)
{
	static_cast<QAbstractPlanarVideoBuffer*>(ptr)->~QAbstractPlanarVideoBuffer();
}

void QAbstractPlanarVideoBuffer_DestroyQAbstractPlanarVideoBufferDefault(void* ptr)
{

}

void* QAbstractPlanarVideoBuffer_Handle(void* ptr)
{
	return new QVariant(static_cast<QAbstractPlanarVideoBuffer*>(ptr)->handle());
}

void* QAbstractPlanarVideoBuffer_HandleDefault(void* ptr)
{
	return new QVariant(static_cast<QAbstractPlanarVideoBuffer*>(ptr)->QAbstractPlanarVideoBuffer::handle());
}

long long QAbstractPlanarVideoBuffer_MapMode(void* ptr)
{
	return static_cast<QAbstractPlanarVideoBuffer*>(ptr)->mapMode();
}



void QAbstractPlanarVideoBuffer_Release(void* ptr)
{
	static_cast<QAbstractPlanarVideoBuffer*>(ptr)->release();
}

void QAbstractPlanarVideoBuffer_ReleaseDefault(void* ptr)
{
	static_cast<QAbstractPlanarVideoBuffer*>(ptr)->QAbstractPlanarVideoBuffer::release();
}

void QAbstractPlanarVideoBuffer_Unmap(void* ptr)
{
	static_cast<QAbstractPlanarVideoBuffer*>(ptr)->unmap();
}



class MyQAbstractVideoBuffer: public QAbstractVideoBuffer
{
public:
	MyQAbstractVideoBuffer(HandleType type) : QAbstractVideoBuffer(type) {};
	QVariant handle() const { return *static_cast<QVariant*>(callbackQAbstractVideoBuffer_Handle(const_cast<MyQAbstractVideoBuffer*>(this))); };
	uchar * map(QAbstractVideoBuffer::MapMode mode, int * numBytes, int * bytesPerLine) { return static_cast<uchar*>(static_cast<void*>(callbackQAbstractVideoBuffer_Map(this, mode, *numBytes, *bytesPerLine))); };
	MapMode mapMode() const { return static_cast<QAbstractVideoBuffer::MapMode>(callbackQAbstractVideoBuffer_MapMode(const_cast<MyQAbstractVideoBuffer*>(this))); };
	void release() { callbackQAbstractVideoBuffer_Release(this); };
	void unmap() { callbackQAbstractVideoBuffer_Unmap(this); };
	 ~MyQAbstractVideoBuffer() { callbackQAbstractVideoBuffer_DestroyQAbstractVideoBuffer(this); };
};

void* QAbstractVideoBuffer_NewQAbstractVideoBuffer(long long ty)
{
	return new MyQAbstractVideoBuffer(static_cast<QAbstractVideoBuffer::HandleType>(ty));
}

void* QAbstractVideoBuffer_Handle(void* ptr)
{
	return new QVariant(static_cast<QAbstractVideoBuffer*>(ptr)->handle());
}

void* QAbstractVideoBuffer_HandleDefault(void* ptr)
{
	return new QVariant(static_cast<QAbstractVideoBuffer*>(ptr)->QAbstractVideoBuffer::handle());
}

long long QAbstractVideoBuffer_HandleType(void* ptr)
{
	return static_cast<QAbstractVideoBuffer*>(ptr)->handleType();
}

struct QtMultimedia_PackedString QAbstractVideoBuffer_Map(void* ptr, long long mode, int numBytes, int bytesPerLine)
{
	return ({ char* t5cec49 = static_cast<char*>(static_cast<void*>(static_cast<QAbstractVideoBuffer*>(ptr)->map(static_cast<QAbstractVideoBuffer::MapMode>(mode), &numBytes, &bytesPerLine))); QtMultimedia_PackedString { t5cec49, -1 }; });
}

long long QAbstractVideoBuffer_MapMode(void* ptr)
{
	return static_cast<QAbstractVideoBuffer*>(ptr)->mapMode();
}

void QAbstractVideoBuffer_Release(void* ptr)
{
	static_cast<QAbstractVideoBuffer*>(ptr)->release();
}

void QAbstractVideoBuffer_ReleaseDefault(void* ptr)
{
	static_cast<QAbstractVideoBuffer*>(ptr)->QAbstractVideoBuffer::release();
}

void QAbstractVideoBuffer_Unmap(void* ptr)
{
	static_cast<QAbstractVideoBuffer*>(ptr)->unmap();
}

void QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(void* ptr)
{
	static_cast<QAbstractVideoBuffer*>(ptr)->~QAbstractVideoBuffer();
}

void QAbstractVideoBuffer_DestroyQAbstractVideoBufferDefault(void* ptr)
{

}

long long QAbstractVideoBuffer_M_type(void* ptr)
{
	return static_cast<QAbstractVideoBuffer*>(ptr)->m_type;
}

void QAbstractVideoBuffer_SetM_type(void* ptr, long long vha)
{
	static_cast<QAbstractVideoBuffer*>(ptr)->m_type = static_cast<QAbstractVideoBuffer::HandleType>(vha);
}

class MyQAbstractVideoFilter: public QAbstractVideoFilter
{
public:
	MyQAbstractVideoFilter(QObject *parent) : QAbstractVideoFilter(parent) {};
	void Signal_ActiveChanged() { callbackQAbstractVideoFilter_ActiveChanged(this); };
	QVideoFilterRunnable * createFilterRunnable() { return static_cast<QVideoFilterRunnable*>(callbackQAbstractVideoFilter_CreateFilterRunnable(this)); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractVideoFilter_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAbstractVideoFilter_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractVideoFilter_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractVideoFilter_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractVideoFilter_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractVideoFilter_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAbstractVideoFilter_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractVideoFilter_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractVideoFilter_MetaObject(const_cast<MyQAbstractVideoFilter*>(this))); };
};

char QAbstractVideoFilter_IsActive(void* ptr)
{
	return static_cast<QAbstractVideoFilter*>(ptr)->isActive();
}

void QAbstractVideoFilter_SetActive(void* ptr, char v)
{
	static_cast<QAbstractVideoFilter*>(ptr)->setActive(v != 0);
}

void* QAbstractVideoFilter_NewQAbstractVideoFilter(void* parent)
{
	return new MyQAbstractVideoFilter(static_cast<QObject*>(parent));
}

void QAbstractVideoFilter_ConnectActiveChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractVideoFilter*>(ptr), static_cast<void (QAbstractVideoFilter::*)()>(&QAbstractVideoFilter::activeChanged), static_cast<MyQAbstractVideoFilter*>(ptr), static_cast<void (MyQAbstractVideoFilter::*)()>(&MyQAbstractVideoFilter::Signal_ActiveChanged));
}

void QAbstractVideoFilter_DisconnectActiveChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractVideoFilter*>(ptr), static_cast<void (QAbstractVideoFilter::*)()>(&QAbstractVideoFilter::activeChanged), static_cast<MyQAbstractVideoFilter*>(ptr), static_cast<void (MyQAbstractVideoFilter::*)()>(&MyQAbstractVideoFilter::Signal_ActiveChanged));
}

void QAbstractVideoFilter_ActiveChanged(void* ptr)
{
	static_cast<QAbstractVideoFilter*>(ptr)->activeChanged();
}

void* QAbstractVideoFilter_CreateFilterRunnable(void* ptr)
{
	return static_cast<QAbstractVideoFilter*>(ptr)->createFilterRunnable();
}

void QAbstractVideoFilter_TimerEvent(void* ptr, void* event)
{
	static_cast<QAbstractVideoFilter*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractVideoFilter_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractVideoFilter*>(ptr)->QAbstractVideoFilter::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractVideoFilter_ChildEvent(void* ptr, void* event)
{
	static_cast<QAbstractVideoFilter*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractVideoFilter_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractVideoFilter*>(ptr)->QAbstractVideoFilter::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractVideoFilter_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractVideoFilter*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractVideoFilter_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractVideoFilter*>(ptr)->QAbstractVideoFilter::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractVideoFilter_CustomEvent(void* ptr, void* event)
{
	static_cast<QAbstractVideoFilter*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractVideoFilter_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractVideoFilter*>(ptr)->QAbstractVideoFilter::customEvent(static_cast<QEvent*>(event));
}

void QAbstractVideoFilter_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractVideoFilter*>(ptr), "deleteLater");
}

void QAbstractVideoFilter_DeleteLaterDefault(void* ptr)
{
	static_cast<QAbstractVideoFilter*>(ptr)->QAbstractVideoFilter::deleteLater();
}

void QAbstractVideoFilter_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractVideoFilter*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractVideoFilter_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractVideoFilter*>(ptr)->QAbstractVideoFilter::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAbstractVideoFilter_Event(void* ptr, void* e)
{
	return static_cast<QAbstractVideoFilter*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAbstractVideoFilter_EventDefault(void* ptr, void* e)
{
	return static_cast<QAbstractVideoFilter*>(ptr)->QAbstractVideoFilter::event(static_cast<QEvent*>(e));
}

char QAbstractVideoFilter_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractVideoFilter*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAbstractVideoFilter_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractVideoFilter*>(ptr)->QAbstractVideoFilter::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAbstractVideoFilter_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractVideoFilter*>(ptr)->metaObject());
}

void* QAbstractVideoFilter_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractVideoFilter*>(ptr)->QAbstractVideoFilter::metaObject());
}

class MyQAbstractVideoSurface: public QAbstractVideoSurface
{
public:
	void Signal_ActiveChanged(bool active) { callbackQAbstractVideoSurface_ActiveChanged(this, active); };
	bool isFormatSupported(const QVideoSurfaceFormat & format) const { return callbackQAbstractVideoSurface_IsFormatSupported(const_cast<MyQAbstractVideoSurface*>(this), const_cast<QVideoSurfaceFormat*>(&format)) != 0; };
	void Signal_NativeResolutionChanged(const QSize & resolution) { callbackQAbstractVideoSurface_NativeResolutionChanged(this, const_cast<QSize*>(&resolution)); };
	QVideoSurfaceFormat nearestFormat(const QVideoSurfaceFormat & format) const { return *static_cast<QVideoSurfaceFormat*>(callbackQAbstractVideoSurface_NearestFormat(const_cast<MyQAbstractVideoSurface*>(this), const_cast<QVideoSurfaceFormat*>(&format))); };
	bool present(const QVideoFrame & frame) { return callbackQAbstractVideoSurface_Present(this, const_cast<QVideoFrame*>(&frame)) != 0; };
	bool start(const QVideoSurfaceFormat & format) { return callbackQAbstractVideoSurface_Start(this, const_cast<QVideoSurfaceFormat*>(&format)) != 0; };
	void stop() { callbackQAbstractVideoSurface_Stop(this); };
	void Signal_SupportedFormatsChanged() { callbackQAbstractVideoSurface_SupportedFormatsChanged(this); };
	void Signal_SurfaceFormatChanged(const QVideoSurfaceFormat & format) { callbackQAbstractVideoSurface_SurfaceFormatChanged(this, const_cast<QVideoSurfaceFormat*>(&format)); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractVideoSurface_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAbstractVideoSurface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractVideoSurface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractVideoSurface_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractVideoSurface_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractVideoSurface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAbstractVideoSurface_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractVideoSurface_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractVideoSurface_MetaObject(const_cast<MyQAbstractVideoSurface*>(this))); };
};

void* QAbstractVideoSurface_NativeResolution(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QAbstractVideoSurface*>(ptr)->nativeResolution(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QAbstractVideoSurface_ConnectActiveChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(bool)>(&QAbstractVideoSurface::activeChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(bool)>(&MyQAbstractVideoSurface::Signal_ActiveChanged));
}

void QAbstractVideoSurface_DisconnectActiveChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(bool)>(&QAbstractVideoSurface::activeChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(bool)>(&MyQAbstractVideoSurface::Signal_ActiveChanged));
}

void QAbstractVideoSurface_ActiveChanged(void* ptr, char active)
{
	static_cast<QAbstractVideoSurface*>(ptr)->activeChanged(active != 0);
}

long long QAbstractVideoSurface_Error(void* ptr)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->error();
}

char QAbstractVideoSurface_IsActive(void* ptr)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->isActive();
}

char QAbstractVideoSurface_IsFormatSupported(void* ptr, void* format)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->isFormatSupported(*static_cast<QVideoSurfaceFormat*>(format));
}

char QAbstractVideoSurface_IsFormatSupportedDefault(void* ptr, void* format)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::isFormatSupported(*static_cast<QVideoSurfaceFormat*>(format));
}

void QAbstractVideoSurface_ConnectNativeResolutionChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(const QSize &)>(&QAbstractVideoSurface::nativeResolutionChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(const QSize &)>(&MyQAbstractVideoSurface::Signal_NativeResolutionChanged));
}

void QAbstractVideoSurface_DisconnectNativeResolutionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(const QSize &)>(&QAbstractVideoSurface::nativeResolutionChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(const QSize &)>(&MyQAbstractVideoSurface::Signal_NativeResolutionChanged));
}

void QAbstractVideoSurface_NativeResolutionChanged(void* ptr, void* resolution)
{
	static_cast<QAbstractVideoSurface*>(ptr)->nativeResolutionChanged(*static_cast<QSize*>(resolution));
}

void* QAbstractVideoSurface_NearestFormat(void* ptr, void* format)
{
	return new QVideoSurfaceFormat(static_cast<QAbstractVideoSurface*>(ptr)->nearestFormat(*static_cast<QVideoSurfaceFormat*>(format)));
}

void* QAbstractVideoSurface_NearestFormatDefault(void* ptr, void* format)
{
	return new QVideoSurfaceFormat(static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::nearestFormat(*static_cast<QVideoSurfaceFormat*>(format)));
}

char QAbstractVideoSurface_Present(void* ptr, void* frame)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->present(*static_cast<QVideoFrame*>(frame));
}

void QAbstractVideoSurface_SetError(void* ptr, long long error)
{
	static_cast<QAbstractVideoSurface*>(ptr)->setError(static_cast<QAbstractVideoSurface::Error>(error));
}

void QAbstractVideoSurface_SetNativeResolution(void* ptr, void* resolution)
{
	static_cast<QAbstractVideoSurface*>(ptr)->setNativeResolution(*static_cast<QSize*>(resolution));
}

char QAbstractVideoSurface_Start(void* ptr, void* format)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->start(*static_cast<QVideoSurfaceFormat*>(format));
}

char QAbstractVideoSurface_StartDefault(void* ptr, void* format)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::start(*static_cast<QVideoSurfaceFormat*>(format));
}

void QAbstractVideoSurface_Stop(void* ptr)
{
	static_cast<QAbstractVideoSurface*>(ptr)->stop();
}

void QAbstractVideoSurface_StopDefault(void* ptr)
{
	static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::stop();
}

void QAbstractVideoSurface_ConnectSupportedFormatsChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)()>(&QAbstractVideoSurface::supportedFormatsChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)()>(&MyQAbstractVideoSurface::Signal_SupportedFormatsChanged));
}

void QAbstractVideoSurface_DisconnectSupportedFormatsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)()>(&QAbstractVideoSurface::supportedFormatsChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)()>(&MyQAbstractVideoSurface::Signal_SupportedFormatsChanged));
}

void QAbstractVideoSurface_SupportedFormatsChanged(void* ptr)
{
	static_cast<QAbstractVideoSurface*>(ptr)->supportedFormatsChanged();
}

void* QAbstractVideoSurface_SurfaceFormat(void* ptr)
{
	return new QVideoSurfaceFormat(static_cast<QAbstractVideoSurface*>(ptr)->surfaceFormat());
}

void QAbstractVideoSurface_ConnectSurfaceFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(const QVideoSurfaceFormat &)>(&QAbstractVideoSurface::surfaceFormatChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(const QVideoSurfaceFormat &)>(&MyQAbstractVideoSurface::Signal_SurfaceFormatChanged));
}

void QAbstractVideoSurface_DisconnectSurfaceFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(const QVideoSurfaceFormat &)>(&QAbstractVideoSurface::surfaceFormatChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(const QVideoSurfaceFormat &)>(&MyQAbstractVideoSurface::Signal_SurfaceFormatChanged));
}

void QAbstractVideoSurface_SurfaceFormatChanged(void* ptr, void* format)
{
	static_cast<QAbstractVideoSurface*>(ptr)->surfaceFormatChanged(*static_cast<QVideoSurfaceFormat*>(format));
}

void QAbstractVideoSurface_DestroyQAbstractVideoSurface(void* ptr)
{
	static_cast<QAbstractVideoSurface*>(ptr)->~QAbstractVideoSurface();
}

void QAbstractVideoSurface_TimerEvent(void* ptr, void* event)
{
	static_cast<QAbstractVideoSurface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractVideoSurface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractVideoSurface_ChildEvent(void* ptr, void* event)
{
	static_cast<QAbstractVideoSurface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractVideoSurface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractVideoSurface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractVideoSurface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractVideoSurface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractVideoSurface_CustomEvent(void* ptr, void* event)
{
	static_cast<QAbstractVideoSurface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractVideoSurface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::customEvent(static_cast<QEvent*>(event));
}

void QAbstractVideoSurface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractVideoSurface*>(ptr), "deleteLater");
}

void QAbstractVideoSurface_DeleteLaterDefault(void* ptr)
{
	static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::deleteLater();
}

void QAbstractVideoSurface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractVideoSurface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractVideoSurface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAbstractVideoSurface_Event(void* ptr, void* e)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAbstractVideoSurface_EventDefault(void* ptr, void* e)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::event(static_cast<QEvent*>(e));
}

char QAbstractVideoSurface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAbstractVideoSurface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAbstractVideoSurface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractVideoSurface*>(ptr)->metaObject());
}

void* QAbstractVideoSurface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractVideoSurface*>(ptr)->QAbstractVideoSurface::metaObject());
}

void* QAudioBuffer_NewQAudioBuffer()
{
	return new QAudioBuffer();
}

void* QAudioBuffer_NewQAudioBuffer3(void* other)
{
	return new QAudioBuffer(*static_cast<QAudioBuffer*>(other));
}

void* QAudioBuffer_NewQAudioBuffer4(void* data, void* format, long long startTime)
{
	return new QAudioBuffer(*static_cast<QByteArray*>(data), *static_cast<QAudioFormat*>(format), startTime);
}

void* QAudioBuffer_NewQAudioBuffer5(int numFrames, void* format, long long startTime)
{
	return new QAudioBuffer(numFrames, *static_cast<QAudioFormat*>(format), startTime);
}

int QAudioBuffer_ByteCount(void* ptr)
{
	return static_cast<QAudioBuffer*>(ptr)->byteCount();
}

void* QAudioBuffer_ConstData2(void* ptr)
{
	return const_cast<void*>(static_cast<QAudioBuffer*>(ptr)->constData());
}

void* QAudioBuffer_ConstData(void* ptr)
{
	return const_cast<void*>(static_cast<QAudioBuffer*>(ptr)->constData());
}

void* QAudioBuffer_Data4(void* ptr)
{
	return static_cast<QAudioBuffer*>(ptr)->data();
}

void* QAudioBuffer_Data2(void* ptr)
{
	return static_cast<QAudioBuffer*>(ptr)->data();
}

void* QAudioBuffer_Data3(void* ptr)
{
	return const_cast<void*>(static_cast<QAudioBuffer*>(ptr)->data());
}

void* QAudioBuffer_Data(void* ptr)
{
	return const_cast<void*>(static_cast<QAudioBuffer*>(ptr)->data());
}

long long QAudioBuffer_Duration(void* ptr)
{
	return static_cast<QAudioBuffer*>(ptr)->duration();
}

void* QAudioBuffer_Format(void* ptr)
{
	return new QAudioFormat(static_cast<QAudioBuffer*>(ptr)->format());
}

int QAudioBuffer_FrameCount(void* ptr)
{
	return static_cast<QAudioBuffer*>(ptr)->frameCount();
}

char QAudioBuffer_IsValid(void* ptr)
{
	return static_cast<QAudioBuffer*>(ptr)->isValid();
}

int QAudioBuffer_SampleCount(void* ptr)
{
	return static_cast<QAudioBuffer*>(ptr)->sampleCount();
}

long long QAudioBuffer_StartTime(void* ptr)
{
	return static_cast<QAudioBuffer*>(ptr)->startTime();
}

void QAudioBuffer_DestroyQAudioBuffer(void* ptr)
{
	static_cast<QAudioBuffer*>(ptr)->~QAudioBuffer();
}

class MyQAudioDecoder: public QAudioDecoder
{
public:
	MyQAudioDecoder(QObject *parent) : QAudioDecoder(parent) {};
	void Signal_BufferAvailableChanged(bool available) { callbackQAudioDecoder_BufferAvailableChanged(this, available); };
	void Signal_BufferReady() { callbackQAudioDecoder_BufferReady(this); };
	void Signal_DurationChanged(qint64 duration) { callbackQAudioDecoder_DurationChanged(this, duration); };
	void Signal_Error2(QAudioDecoder::Error error) { callbackQAudioDecoder_Error2(this, error); };
	void Signal_Finished() { callbackQAudioDecoder_Finished(this); };
	void Signal_FormatChanged(const QAudioFormat & format) { callbackQAudioDecoder_FormatChanged(this, const_cast<QAudioFormat*>(&format)); };
	void Signal_PositionChanged(qint64 position) { callbackQAudioDecoder_PositionChanged(this, position); };
	void Signal_SourceChanged() { callbackQAudioDecoder_SourceChanged(this); };
	void start() { callbackQAudioDecoder_Start(this); };
	void Signal_StateChanged(QAudioDecoder::State state) { callbackQAudioDecoder_StateChanged(this, state); };
	void stop() { callbackQAudioDecoder_Stop(this); };
	QMultimedia::AvailabilityStatus availability() const { return static_cast<QMultimedia::AvailabilityStatus>(callbackQAudioDecoder_Availability(const_cast<MyQAudioDecoder*>(this))); };
	bool bind(QObject * object) { return callbackQAudioDecoder_Bind(this, object) != 0; };
	bool isAvailable() const { return callbackQAudioDecoder_IsAvailable(const_cast<MyQAudioDecoder*>(this)) != 0; };
	QMediaService * service() const { return static_cast<QMediaService*>(callbackQAudioDecoder_Service(const_cast<MyQAudioDecoder*>(this))); };
	void unbind(QObject * object) { callbackQAudioDecoder_Unbind(this, object); };
	void timerEvent(QTimerEvent * event) { callbackQAudioDecoder_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioDecoder_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioDecoder_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioDecoder_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioDecoder_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioDecoder_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioDecoder_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioDecoder_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioDecoder_MetaObject(const_cast<MyQAudioDecoder*>(this))); };
};

struct QtMultimedia_PackedString QAudioDecoder_ErrorString(void* ptr)
{
	return ({ QByteArray t574770 = static_cast<QAudioDecoder*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t574770.prepend("WHITESPACE").constData()+10), t574770.size()-10 }; });
}

long long QAudioDecoder_State(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->state();
}

void* QAudioDecoder_NewQAudioDecoder(void* parent)
{
	return new MyQAudioDecoder(static_cast<QObject*>(parent));
}

void* QAudioDecoder_AudioFormat(void* ptr)
{
	return new QAudioFormat(static_cast<QAudioDecoder*>(ptr)->audioFormat());
}

char QAudioDecoder_BufferAvailable(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->bufferAvailable();
}

void QAudioDecoder_ConnectBufferAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(bool)>(&QAudioDecoder::bufferAvailableChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(bool)>(&MyQAudioDecoder::Signal_BufferAvailableChanged));
}

void QAudioDecoder_DisconnectBufferAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(bool)>(&QAudioDecoder::bufferAvailableChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(bool)>(&MyQAudioDecoder::Signal_BufferAvailableChanged));
}

void QAudioDecoder_BufferAvailableChanged(void* ptr, char available)
{
	static_cast<QAudioDecoder*>(ptr)->bufferAvailableChanged(available != 0);
}

void QAudioDecoder_ConnectBufferReady(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::bufferReady), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_BufferReady));
}

void QAudioDecoder_DisconnectBufferReady(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::bufferReady), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_BufferReady));
}

void QAudioDecoder_BufferReady(void* ptr)
{
	static_cast<QAudioDecoder*>(ptr)->bufferReady();
}

long long QAudioDecoder_Duration(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->duration();
}

void QAudioDecoder_ConnectDurationChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(qint64)>(&QAudioDecoder::durationChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(qint64)>(&MyQAudioDecoder::Signal_DurationChanged));
}

void QAudioDecoder_DisconnectDurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(qint64)>(&QAudioDecoder::durationChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(qint64)>(&MyQAudioDecoder::Signal_DurationChanged));
}

void QAudioDecoder_DurationChanged(void* ptr, long long duration)
{
	static_cast<QAudioDecoder*>(ptr)->durationChanged(duration);
}

void QAudioDecoder_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::Error)>(&QAudioDecoder::error), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::Error)>(&MyQAudioDecoder::Signal_Error2));
}

void QAudioDecoder_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::Error)>(&QAudioDecoder::error), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::Error)>(&MyQAudioDecoder::Signal_Error2));
}

void QAudioDecoder_Error2(void* ptr, long long error)
{
	static_cast<QAudioDecoder*>(ptr)->error(static_cast<QAudioDecoder::Error>(error));
}

long long QAudioDecoder_Error(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->error();
}

void QAudioDecoder_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::finished), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_Finished));
}

void QAudioDecoder_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::finished), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_Finished));
}

void QAudioDecoder_Finished(void* ptr)
{
	static_cast<QAudioDecoder*>(ptr)->finished();
}

void QAudioDecoder_ConnectFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(const QAudioFormat &)>(&QAudioDecoder::formatChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(const QAudioFormat &)>(&MyQAudioDecoder::Signal_FormatChanged));
}

void QAudioDecoder_DisconnectFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(const QAudioFormat &)>(&QAudioDecoder::formatChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(const QAudioFormat &)>(&MyQAudioDecoder::Signal_FormatChanged));
}

void QAudioDecoder_FormatChanged(void* ptr, void* format)
{
	static_cast<QAudioDecoder*>(ptr)->formatChanged(*static_cast<QAudioFormat*>(format));
}

long long QAudioDecoder_QAudioDecoder_HasSupport(char* mimeType, char* codecs)
{
	return QAudioDecoder::hasSupport(QString(mimeType), QString(codecs).split("|", QString::SkipEmptyParts));
}

long long QAudioDecoder_Position(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->position();
}

void QAudioDecoder_ConnectPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(qint64)>(&QAudioDecoder::positionChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(qint64)>(&MyQAudioDecoder::Signal_PositionChanged));
}

void QAudioDecoder_DisconnectPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(qint64)>(&QAudioDecoder::positionChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(qint64)>(&MyQAudioDecoder::Signal_PositionChanged));
}

void QAudioDecoder_PositionChanged(void* ptr, long long position)
{
	static_cast<QAudioDecoder*>(ptr)->positionChanged(position);
}

void* QAudioDecoder_Read(void* ptr)
{
	return new QAudioBuffer(static_cast<QAudioDecoder*>(ptr)->read());
}

void QAudioDecoder_SetAudioFormat(void* ptr, void* format)
{
	static_cast<QAudioDecoder*>(ptr)->setAudioFormat(*static_cast<QAudioFormat*>(format));
}

void QAudioDecoder_SetSourceDevice(void* ptr, void* device)
{
	static_cast<QAudioDecoder*>(ptr)->setSourceDevice(static_cast<QIODevice*>(device));
}

void QAudioDecoder_SetSourceFilename(void* ptr, char* fileName)
{
	static_cast<QAudioDecoder*>(ptr)->setSourceFilename(QString(fileName));
}

void QAudioDecoder_ConnectSourceChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::sourceChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_SourceChanged));
}

void QAudioDecoder_DisconnectSourceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::sourceChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_SourceChanged));
}

void QAudioDecoder_SourceChanged(void* ptr)
{
	static_cast<QAudioDecoder*>(ptr)->sourceChanged();
}

void* QAudioDecoder_SourceDevice(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->sourceDevice();
}

struct QtMultimedia_PackedString QAudioDecoder_SourceFilename(void* ptr)
{
	return ({ QByteArray t0d51d3 = static_cast<QAudioDecoder*>(ptr)->sourceFilename().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t0d51d3.prepend("WHITESPACE").constData()+10), t0d51d3.size()-10 }; });
}

void QAudioDecoder_Start(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioDecoder*>(ptr), "start");
}

void QAudioDecoder_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::State)>(&QAudioDecoder::stateChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::State)>(&MyQAudioDecoder::Signal_StateChanged));
}

void QAudioDecoder_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::State)>(&QAudioDecoder::stateChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::State)>(&MyQAudioDecoder::Signal_StateChanged));
}

void QAudioDecoder_StateChanged(void* ptr, long long state)
{
	static_cast<QAudioDecoder*>(ptr)->stateChanged(static_cast<QAudioDecoder::State>(state));
}

void QAudioDecoder_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioDecoder*>(ptr), "stop");
}

void QAudioDecoder_DestroyQAudioDecoder(void* ptr)
{
	static_cast<QAudioDecoder*>(ptr)->~QAudioDecoder();
}

long long QAudioDecoder_Availability(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->availability();
}

long long QAudioDecoder_AvailabilityDefault(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::availability();
}

char QAudioDecoder_Bind(void* ptr, void* object)
{
	return static_cast<QAudioDecoder*>(ptr)->bind(static_cast<QObject*>(object));
}

char QAudioDecoder_BindDefault(void* ptr, void* object)
{
	return static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::bind(static_cast<QObject*>(object));
}

char QAudioDecoder_IsAvailable(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->isAvailable();
}

char QAudioDecoder_IsAvailableDefault(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::isAvailable();
}

void* QAudioDecoder_Service(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->service();
}

void* QAudioDecoder_ServiceDefault(void* ptr)
{
	return static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::service();
}

void QAudioDecoder_Unbind(void* ptr, void* object)
{
	static_cast<QAudioDecoder*>(ptr)->unbind(static_cast<QObject*>(object));
}

void QAudioDecoder_UnbindDefault(void* ptr, void* object)
{
	static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::unbind(static_cast<QObject*>(object));
}

void QAudioDecoder_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioDecoder*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioDecoder_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioDecoder_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioDecoder*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioDecoder_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioDecoder_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioDecoder*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioDecoder_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioDecoder_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioDecoder*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioDecoder_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::customEvent(static_cast<QEvent*>(event));
}

void QAudioDecoder_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioDecoder*>(ptr), "deleteLater");
}

void QAudioDecoder_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::deleteLater();
}

void QAudioDecoder_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioDecoder*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioDecoder_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioDecoder_Event(void* ptr, void* e)
{
	return static_cast<QAudioDecoder*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioDecoder_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::event(static_cast<QEvent*>(e));
}

char QAudioDecoder_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioDecoder*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioDecoder_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioDecoder_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioDecoder*>(ptr)->metaObject());
}

void* QAudioDecoder_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioDecoder*>(ptr)->QAudioDecoder::metaObject());
}

class MyQAudioDecoderControl: public QAudioDecoderControl
{
public:
	MyQAudioDecoderControl(QObject *parent) : QAudioDecoderControl(parent) {};
	QAudioFormat audioFormat() const { return *static_cast<QAudioFormat*>(callbackQAudioDecoderControl_AudioFormat(const_cast<MyQAudioDecoderControl*>(this))); };
	bool bufferAvailable() const { return callbackQAudioDecoderControl_BufferAvailable(const_cast<MyQAudioDecoderControl*>(this)) != 0; };
	void Signal_BufferAvailableChanged(bool available) { callbackQAudioDecoderControl_BufferAvailableChanged(this, available); };
	void Signal_BufferReady() { callbackQAudioDecoderControl_BufferReady(this); };
	qint64 duration() const { return callbackQAudioDecoderControl_Duration(const_cast<MyQAudioDecoderControl*>(this)); };
	void Signal_DurationChanged(qint64 duration) { callbackQAudioDecoderControl_DurationChanged(this, duration); };
	void Signal_Error(int error, const QString & errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtMultimedia_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQAudioDecoderControl_Error(this, error, errorStringPacked); };
	void Signal_Finished() { callbackQAudioDecoderControl_Finished(this); };
	void Signal_FormatChanged(const QAudioFormat & format) { callbackQAudioDecoderControl_FormatChanged(this, const_cast<QAudioFormat*>(&format)); };
	qint64 position() const { return callbackQAudioDecoderControl_Position(const_cast<MyQAudioDecoderControl*>(this)); };
	void Signal_PositionChanged(qint64 position) { callbackQAudioDecoderControl_PositionChanged(this, position); };
	QAudioBuffer read() { return *static_cast<QAudioBuffer*>(callbackQAudioDecoderControl_Read(this)); };
	void setAudioFormat(const QAudioFormat & format) { callbackQAudioDecoderControl_SetAudioFormat(this, const_cast<QAudioFormat*>(&format)); };
	void setSourceDevice(QIODevice * device) { callbackQAudioDecoderControl_SetSourceDevice(this, device); };
	void setSourceFilename(const QString & fileName) { QByteArray td83e09 = fileName.toUtf8(); QtMultimedia_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQAudioDecoderControl_SetSourceFilename(this, fileNamePacked); };
	void Signal_SourceChanged() { callbackQAudioDecoderControl_SourceChanged(this); };
	QIODevice * sourceDevice() const { return static_cast<QIODevice*>(callbackQAudioDecoderControl_SourceDevice(const_cast<MyQAudioDecoderControl*>(this))); };
	QString sourceFilename() const { return QString(callbackQAudioDecoderControl_SourceFilename(const_cast<MyQAudioDecoderControl*>(this))); };
	void start() { callbackQAudioDecoderControl_Start(this); };
	QAudioDecoder::State state() const { return static_cast<QAudioDecoder::State>(callbackQAudioDecoderControl_State(const_cast<MyQAudioDecoderControl*>(this))); };
	void Signal_StateChanged(QAudioDecoder::State state) { callbackQAudioDecoderControl_StateChanged(this, state); };
	void stop() { callbackQAudioDecoderControl_Stop(this); };
	void timerEvent(QTimerEvent * event) { callbackQAudioDecoderControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioDecoderControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioDecoderControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioDecoderControl_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioDecoderControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioDecoderControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioDecoderControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioDecoderControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioDecoderControl_MetaObject(const_cast<MyQAudioDecoderControl*>(this))); };
};

void* QAudioDecoderControl_NewQAudioDecoderControl(void* parent)
{
	return new MyQAudioDecoderControl(static_cast<QObject*>(parent));
}

void* QAudioDecoderControl_AudioFormat(void* ptr)
{
	return new QAudioFormat(static_cast<QAudioDecoderControl*>(ptr)->audioFormat());
}

char QAudioDecoderControl_BufferAvailable(void* ptr)
{
	return static_cast<QAudioDecoderControl*>(ptr)->bufferAvailable();
}

void QAudioDecoderControl_ConnectBufferAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(bool)>(&QAudioDecoderControl::bufferAvailableChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(bool)>(&MyQAudioDecoderControl::Signal_BufferAvailableChanged));
}

void QAudioDecoderControl_DisconnectBufferAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(bool)>(&QAudioDecoderControl::bufferAvailableChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(bool)>(&MyQAudioDecoderControl::Signal_BufferAvailableChanged));
}

void QAudioDecoderControl_BufferAvailableChanged(void* ptr, char available)
{
	static_cast<QAudioDecoderControl*>(ptr)->bufferAvailableChanged(available != 0);
}

void QAudioDecoderControl_ConnectBufferReady(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::bufferReady), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_BufferReady));
}

void QAudioDecoderControl_DisconnectBufferReady(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::bufferReady), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_BufferReady));
}

void QAudioDecoderControl_BufferReady(void* ptr)
{
	static_cast<QAudioDecoderControl*>(ptr)->bufferReady();
}

long long QAudioDecoderControl_Duration(void* ptr)
{
	return static_cast<QAudioDecoderControl*>(ptr)->duration();
}

void QAudioDecoderControl_ConnectDurationChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(qint64)>(&QAudioDecoderControl::durationChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(qint64)>(&MyQAudioDecoderControl::Signal_DurationChanged));
}

void QAudioDecoderControl_DisconnectDurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(qint64)>(&QAudioDecoderControl::durationChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(qint64)>(&MyQAudioDecoderControl::Signal_DurationChanged));
}

void QAudioDecoderControl_DurationChanged(void* ptr, long long duration)
{
	static_cast<QAudioDecoderControl*>(ptr)->durationChanged(duration);
}

void QAudioDecoderControl_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(int, const QString &)>(&QAudioDecoderControl::error), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(int, const QString &)>(&MyQAudioDecoderControl::Signal_Error));
}

void QAudioDecoderControl_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(int, const QString &)>(&QAudioDecoderControl::error), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(int, const QString &)>(&MyQAudioDecoderControl::Signal_Error));
}

void QAudioDecoderControl_Error(void* ptr, int error, char* errorString)
{
	static_cast<QAudioDecoderControl*>(ptr)->error(error, QString(errorString));
}

void QAudioDecoderControl_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::finished), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_Finished));
}

void QAudioDecoderControl_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::finished), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_Finished));
}

void QAudioDecoderControl_Finished(void* ptr)
{
	static_cast<QAudioDecoderControl*>(ptr)->finished();
}

void QAudioDecoderControl_ConnectFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(const QAudioFormat &)>(&QAudioDecoderControl::formatChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(const QAudioFormat &)>(&MyQAudioDecoderControl::Signal_FormatChanged));
}

void QAudioDecoderControl_DisconnectFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(const QAudioFormat &)>(&QAudioDecoderControl::formatChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(const QAudioFormat &)>(&MyQAudioDecoderControl::Signal_FormatChanged));
}

void QAudioDecoderControl_FormatChanged(void* ptr, void* format)
{
	static_cast<QAudioDecoderControl*>(ptr)->formatChanged(*static_cast<QAudioFormat*>(format));
}

long long QAudioDecoderControl_Position(void* ptr)
{
	return static_cast<QAudioDecoderControl*>(ptr)->position();
}

void QAudioDecoderControl_ConnectPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(qint64)>(&QAudioDecoderControl::positionChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(qint64)>(&MyQAudioDecoderControl::Signal_PositionChanged));
}

void QAudioDecoderControl_DisconnectPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(qint64)>(&QAudioDecoderControl::positionChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(qint64)>(&MyQAudioDecoderControl::Signal_PositionChanged));
}

void QAudioDecoderControl_PositionChanged(void* ptr, long long position)
{
	static_cast<QAudioDecoderControl*>(ptr)->positionChanged(position);
}

void* QAudioDecoderControl_Read(void* ptr)
{
	return new QAudioBuffer(static_cast<QAudioDecoderControl*>(ptr)->read());
}

void QAudioDecoderControl_SetAudioFormat(void* ptr, void* format)
{
	static_cast<QAudioDecoderControl*>(ptr)->setAudioFormat(*static_cast<QAudioFormat*>(format));
}

void QAudioDecoderControl_SetSourceDevice(void* ptr, void* device)
{
	static_cast<QAudioDecoderControl*>(ptr)->setSourceDevice(static_cast<QIODevice*>(device));
}

void QAudioDecoderControl_SetSourceFilename(void* ptr, char* fileName)
{
	static_cast<QAudioDecoderControl*>(ptr)->setSourceFilename(QString(fileName));
}

void QAudioDecoderControl_ConnectSourceChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::sourceChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_SourceChanged));
}

void QAudioDecoderControl_DisconnectSourceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::sourceChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_SourceChanged));
}

void QAudioDecoderControl_SourceChanged(void* ptr)
{
	static_cast<QAudioDecoderControl*>(ptr)->sourceChanged();
}

void* QAudioDecoderControl_SourceDevice(void* ptr)
{
	return static_cast<QAudioDecoderControl*>(ptr)->sourceDevice();
}

struct QtMultimedia_PackedString QAudioDecoderControl_SourceFilename(void* ptr)
{
	return ({ QByteArray tcf9fb0 = static_cast<QAudioDecoderControl*>(ptr)->sourceFilename().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tcf9fb0.prepend("WHITESPACE").constData()+10), tcf9fb0.size()-10 }; });
}

void QAudioDecoderControl_Start(void* ptr)
{
	static_cast<QAudioDecoderControl*>(ptr)->start();
}

long long QAudioDecoderControl_State(void* ptr)
{
	return static_cast<QAudioDecoderControl*>(ptr)->state();
}

void QAudioDecoderControl_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(QAudioDecoder::State)>(&QAudioDecoderControl::stateChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(QAudioDecoder::State)>(&MyQAudioDecoderControl::Signal_StateChanged));
}

void QAudioDecoderControl_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(QAudioDecoder::State)>(&QAudioDecoderControl::stateChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(QAudioDecoder::State)>(&MyQAudioDecoderControl::Signal_StateChanged));
}

void QAudioDecoderControl_StateChanged(void* ptr, long long state)
{
	static_cast<QAudioDecoderControl*>(ptr)->stateChanged(static_cast<QAudioDecoder::State>(state));
}

void QAudioDecoderControl_Stop(void* ptr)
{
	static_cast<QAudioDecoderControl*>(ptr)->stop();
}

void QAudioDecoderControl_DestroyQAudioDecoderControl(void* ptr)
{
	static_cast<QAudioDecoderControl*>(ptr)->~QAudioDecoderControl();
}

void QAudioDecoderControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioDecoderControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioDecoderControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioDecoderControl*>(ptr)->QAudioDecoderControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioDecoderControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioDecoderControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioDecoderControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioDecoderControl*>(ptr)->QAudioDecoderControl::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioDecoderControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioDecoderControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioDecoderControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioDecoderControl*>(ptr)->QAudioDecoderControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioDecoderControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioDecoderControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioDecoderControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioDecoderControl*>(ptr)->QAudioDecoderControl::customEvent(static_cast<QEvent*>(event));
}

void QAudioDecoderControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioDecoderControl*>(ptr), "deleteLater");
}

void QAudioDecoderControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioDecoderControl*>(ptr)->QAudioDecoderControl::deleteLater();
}

void QAudioDecoderControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioDecoderControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioDecoderControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioDecoderControl*>(ptr)->QAudioDecoderControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioDecoderControl_Event(void* ptr, void* e)
{
	return static_cast<QAudioDecoderControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioDecoderControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioDecoderControl*>(ptr)->QAudioDecoderControl::event(static_cast<QEvent*>(e));
}

char QAudioDecoderControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioDecoderControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioDecoderControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioDecoderControl*>(ptr)->QAudioDecoderControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioDecoderControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioDecoderControl*>(ptr)->metaObject());
}

void* QAudioDecoderControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioDecoderControl*>(ptr)->QAudioDecoderControl::metaObject());
}

void* QAudioDeviceInfo_NewQAudioDeviceInfo()
{
	return new QAudioDeviceInfo();
}

void* QAudioDeviceInfo_NewQAudioDeviceInfo2(void* other)
{
	return new QAudioDeviceInfo(*static_cast<QAudioDeviceInfo*>(other));
}

void* QAudioDeviceInfo_QAudioDeviceInfo_DefaultInputDevice()
{
	return new QAudioDeviceInfo(QAudioDeviceInfo::defaultInputDevice());
}

void* QAudioDeviceInfo_QAudioDeviceInfo_DefaultOutputDevice()
{
	return new QAudioDeviceInfo(QAudioDeviceInfo::defaultOutputDevice());
}

struct QtMultimedia_PackedString QAudioDeviceInfo_DeviceName(void* ptr)
{
	return ({ QByteArray t7b618f = static_cast<QAudioDeviceInfo*>(ptr)->deviceName().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t7b618f.prepend("WHITESPACE").constData()+10), t7b618f.size()-10 }; });
}

char QAudioDeviceInfo_IsFormatSupported(void* ptr, void* settings)
{
	return static_cast<QAudioDeviceInfo*>(ptr)->isFormatSupported(*static_cast<QAudioFormat*>(settings));
}

char QAudioDeviceInfo_IsNull(void* ptr)
{
	return static_cast<QAudioDeviceInfo*>(ptr)->isNull();
}

void* QAudioDeviceInfo_NearestFormat(void* ptr, void* settings)
{
	return new QAudioFormat(static_cast<QAudioDeviceInfo*>(ptr)->nearestFormat(*static_cast<QAudioFormat*>(settings)));
}

void* QAudioDeviceInfo_PreferredFormat(void* ptr)
{
	return new QAudioFormat(static_cast<QAudioDeviceInfo*>(ptr)->preferredFormat());
}

struct QtMultimedia_PackedString QAudioDeviceInfo_SupportedCodecs(void* ptr)
{
	return ({ QByteArray t91bb7c = static_cast<QAudioDeviceInfo*>(ptr)->supportedCodecs().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t91bb7c.prepend("WHITESPACE").constData()+10), t91bb7c.size()-10 }; });
}

void QAudioDeviceInfo_DestroyQAudioDeviceInfo(void* ptr)
{
	static_cast<QAudioDeviceInfo*>(ptr)->~QAudioDeviceInfo();
}

void* QAudioEncoderSettings_NewQAudioEncoderSettings()
{
	return new QAudioEncoderSettings();
}

void* QAudioEncoderSettings_NewQAudioEncoderSettings2(void* other)
{
	return new QAudioEncoderSettings(*static_cast<QAudioEncoderSettings*>(other));
}

int QAudioEncoderSettings_BitRate(void* ptr)
{
	return static_cast<QAudioEncoderSettings*>(ptr)->bitRate();
}

int QAudioEncoderSettings_ChannelCount(void* ptr)
{
	return static_cast<QAudioEncoderSettings*>(ptr)->channelCount();
}

struct QtMultimedia_PackedString QAudioEncoderSettings_Codec(void* ptr)
{
	return ({ QByteArray t4add78 = static_cast<QAudioEncoderSettings*>(ptr)->codec().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t4add78.prepend("WHITESPACE").constData()+10), t4add78.size()-10 }; });
}

long long QAudioEncoderSettings_EncodingMode(void* ptr)
{
	return static_cast<QAudioEncoderSettings*>(ptr)->encodingMode();
}

void* QAudioEncoderSettings_EncodingOption(void* ptr, char* option)
{
	return new QVariant(static_cast<QAudioEncoderSettings*>(ptr)->encodingOption(QString(option)));
}

char QAudioEncoderSettings_IsNull(void* ptr)
{
	return static_cast<QAudioEncoderSettings*>(ptr)->isNull();
}

long long QAudioEncoderSettings_Quality(void* ptr)
{
	return static_cast<QAudioEncoderSettings*>(ptr)->quality();
}

int QAudioEncoderSettings_SampleRate(void* ptr)
{
	return static_cast<QAudioEncoderSettings*>(ptr)->sampleRate();
}

void QAudioEncoderSettings_SetBitRate(void* ptr, int rate)
{
	static_cast<QAudioEncoderSettings*>(ptr)->setBitRate(rate);
}

void QAudioEncoderSettings_SetChannelCount(void* ptr, int channels)
{
	static_cast<QAudioEncoderSettings*>(ptr)->setChannelCount(channels);
}

void QAudioEncoderSettings_SetCodec(void* ptr, char* codec)
{
	static_cast<QAudioEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QAudioEncoderSettings_SetEncodingMode(void* ptr, long long mode)
{
	static_cast<QAudioEncoderSettings*>(ptr)->setEncodingMode(static_cast<QMultimedia::EncodingMode>(mode));
}

void QAudioEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value)
{
	static_cast<QAudioEncoderSettings*>(ptr)->setEncodingOption(QString(option), *static_cast<QVariant*>(value));
}

void QAudioEncoderSettings_SetQuality(void* ptr, long long quality)
{
	static_cast<QAudioEncoderSettings*>(ptr)->setQuality(static_cast<QMultimedia::EncodingQuality>(quality));
}

void QAudioEncoderSettings_SetSampleRate(void* ptr, int rate)
{
	static_cast<QAudioEncoderSettings*>(ptr)->setSampleRate(rate);
}

void QAudioEncoderSettings_DestroyQAudioEncoderSettings(void* ptr)
{
	static_cast<QAudioEncoderSettings*>(ptr)->~QAudioEncoderSettings();
}

class MyQAudioEncoderSettingsControl: public QAudioEncoderSettingsControl
{
public:
	QAudioEncoderSettings audioSettings() const { return *static_cast<QAudioEncoderSettings*>(callbackQAudioEncoderSettingsControl_AudioSettings(const_cast<MyQAudioEncoderSettingsControl*>(this))); };
	QString codecDescription(const QString & codec) const { QByteArray td061f6 = codec.toUtf8(); QtMultimedia_PackedString codecPacked = { const_cast<char*>(td061f6.prepend("WHITESPACE").constData()+10), td061f6.size()-10 };return QString(callbackQAudioEncoderSettingsControl_CodecDescription(const_cast<MyQAudioEncoderSettingsControl*>(this), codecPacked)); };
	void setAudioSettings(const QAudioEncoderSettings & settings) { callbackQAudioEncoderSettingsControl_SetAudioSettings(this, const_cast<QAudioEncoderSettings*>(&settings)); };
	QStringList supportedAudioCodecs() const { return QString(callbackQAudioEncoderSettingsControl_SupportedAudioCodecs(const_cast<MyQAudioEncoderSettingsControl*>(this))).split("|", QString::SkipEmptyParts); };
	 ~MyQAudioEncoderSettingsControl() { callbackQAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQAudioEncoderSettingsControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioEncoderSettingsControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioEncoderSettingsControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioEncoderSettingsControl_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioEncoderSettingsControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioEncoderSettingsControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioEncoderSettingsControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioEncoderSettingsControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioEncoderSettingsControl_MetaObject(const_cast<MyQAudioEncoderSettingsControl*>(this))); };
};

void* QAudioEncoderSettingsControl_AudioSettings(void* ptr)
{
	return new QAudioEncoderSettings(static_cast<QAudioEncoderSettingsControl*>(ptr)->audioSettings());
}

struct QtMultimedia_PackedString QAudioEncoderSettingsControl_CodecDescription(void* ptr, char* codec)
{
	return ({ QByteArray ta4f9aa = static_cast<QAudioEncoderSettingsControl*>(ptr)->codecDescription(QString(codec)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(ta4f9aa.prepend("WHITESPACE").constData()+10), ta4f9aa.size()-10 }; });
}

void QAudioEncoderSettingsControl_SetAudioSettings(void* ptr, void* settings)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->setAudioSettings(*static_cast<QAudioEncoderSettings*>(settings));
}

struct QtMultimedia_PackedString QAudioEncoderSettingsControl_SupportedAudioCodecs(void* ptr)
{
	return ({ QByteArray t640ebe = static_cast<QAudioEncoderSettingsControl*>(ptr)->supportedAudioCodecs().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t640ebe.prepend("WHITESPACE").constData()+10), t640ebe.size()-10 }; });
}

void QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(void* ptr)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->~QAudioEncoderSettingsControl();
}

void QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControlDefault(void* ptr)
{

}

void QAudioEncoderSettingsControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioEncoderSettingsControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->QAudioEncoderSettingsControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioEncoderSettingsControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioEncoderSettingsControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->QAudioEncoderSettingsControl::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioEncoderSettingsControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioEncoderSettingsControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->QAudioEncoderSettingsControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioEncoderSettingsControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioEncoderSettingsControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->QAudioEncoderSettingsControl::customEvent(static_cast<QEvent*>(event));
}

void QAudioEncoderSettingsControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioEncoderSettingsControl*>(ptr), "deleteLater");
}

void QAudioEncoderSettingsControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->QAudioEncoderSettingsControl::deleteLater();
}

void QAudioEncoderSettingsControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioEncoderSettingsControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioEncoderSettingsControl*>(ptr)->QAudioEncoderSettingsControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioEncoderSettingsControl_Event(void* ptr, void* e)
{
	return static_cast<QAudioEncoderSettingsControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioEncoderSettingsControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioEncoderSettingsControl*>(ptr)->QAudioEncoderSettingsControl::event(static_cast<QEvent*>(e));
}

char QAudioEncoderSettingsControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioEncoderSettingsControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioEncoderSettingsControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioEncoderSettingsControl*>(ptr)->QAudioEncoderSettingsControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioEncoderSettingsControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioEncoderSettingsControl*>(ptr)->metaObject());
}

void* QAudioEncoderSettingsControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioEncoderSettingsControl*>(ptr)->QAudioEncoderSettingsControl::metaObject());
}

void* QAudioFormat_NewQAudioFormat()
{
	return new QAudioFormat();
}

void* QAudioFormat_NewQAudioFormat2(void* other)
{
	return new QAudioFormat(*static_cast<QAudioFormat*>(other));
}

long long QAudioFormat_ByteOrder(void* ptr)
{
	return static_cast<QAudioFormat*>(ptr)->byteOrder();
}

int QAudioFormat_BytesForDuration(void* ptr, long long duration)
{
	return static_cast<QAudioFormat*>(ptr)->bytesForDuration(duration);
}

int QAudioFormat_BytesForFrames(void* ptr, int frameCount)
{
	return static_cast<QAudioFormat*>(ptr)->bytesForFrames(frameCount);
}

int QAudioFormat_BytesPerFrame(void* ptr)
{
	return static_cast<QAudioFormat*>(ptr)->bytesPerFrame();
}

int QAudioFormat_ChannelCount(void* ptr)
{
	return static_cast<QAudioFormat*>(ptr)->channelCount();
}

struct QtMultimedia_PackedString QAudioFormat_Codec(void* ptr)
{
	return ({ QByteArray t6e9963 = static_cast<QAudioFormat*>(ptr)->codec().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t6e9963.prepend("WHITESPACE").constData()+10), t6e9963.size()-10 }; });
}

long long QAudioFormat_DurationForBytes(void* ptr, int bytes)
{
	return static_cast<QAudioFormat*>(ptr)->durationForBytes(bytes);
}

long long QAudioFormat_DurationForFrames(void* ptr, int frameCount)
{
	return static_cast<QAudioFormat*>(ptr)->durationForFrames(frameCount);
}

int QAudioFormat_FramesForBytes(void* ptr, int byteCount)
{
	return static_cast<QAudioFormat*>(ptr)->framesForBytes(byteCount);
}

int QAudioFormat_FramesForDuration(void* ptr, long long duration)
{
	return static_cast<QAudioFormat*>(ptr)->framesForDuration(duration);
}

char QAudioFormat_IsValid(void* ptr)
{
	return static_cast<QAudioFormat*>(ptr)->isValid();
}

int QAudioFormat_SampleRate(void* ptr)
{
	return static_cast<QAudioFormat*>(ptr)->sampleRate();
}

int QAudioFormat_SampleSize(void* ptr)
{
	return static_cast<QAudioFormat*>(ptr)->sampleSize();
}

long long QAudioFormat_SampleType(void* ptr)
{
	return static_cast<QAudioFormat*>(ptr)->sampleType();
}

void QAudioFormat_SetByteOrder(void* ptr, long long byteOrder)
{
	static_cast<QAudioFormat*>(ptr)->setByteOrder(static_cast<QAudioFormat::Endian>(byteOrder));
}

void QAudioFormat_SetChannelCount(void* ptr, int channels)
{
	static_cast<QAudioFormat*>(ptr)->setChannelCount(channels);
}

void QAudioFormat_SetCodec(void* ptr, char* codec)
{
	static_cast<QAudioFormat*>(ptr)->setCodec(QString(codec));
}

void QAudioFormat_SetSampleRate(void* ptr, int samplerate)
{
	static_cast<QAudioFormat*>(ptr)->setSampleRate(samplerate);
}

void QAudioFormat_SetSampleSize(void* ptr, int sampleSize)
{
	static_cast<QAudioFormat*>(ptr)->setSampleSize(sampleSize);
}

void QAudioFormat_SetSampleType(void* ptr, long long sampleType)
{
	static_cast<QAudioFormat*>(ptr)->setSampleType(static_cast<QAudioFormat::SampleType>(sampleType));
}

void QAudioFormat_DestroyQAudioFormat(void* ptr)
{
	static_cast<QAudioFormat*>(ptr)->~QAudioFormat();
}

class MyQAudioInput: public QAudioInput
{
public:
	MyQAudioInput(const QAudioDeviceInfo &audioDevice, const QAudioFormat &format, QObject *parent) : QAudioInput(audioDevice, format, parent) {};
	MyQAudioInput(const QAudioFormat &format, QObject *parent) : QAudioInput(format, parent) {};
	void Signal_StateChanged(QAudio::State state) { callbackQAudioInput_StateChanged(this, state); };
	void timerEvent(QTimerEvent * event) { callbackQAudioInput_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioInput_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioInput_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioInput_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioInput_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioInput_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioInput_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioInput_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioInput_MetaObject(const_cast<MyQAudioInput*>(this))); };
};

void* QAudioInput_NewQAudioInput2(void* audioDevice, void* format, void* parent)
{
	return new MyQAudioInput(*static_cast<QAudioDeviceInfo*>(audioDevice), *static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

void* QAudioInput_NewQAudioInput(void* format, void* parent)
{
	return new MyQAudioInput(*static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

int QAudioInput_BufferSize(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->bufferSize();
}

int QAudioInput_BytesReady(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->bytesReady();
}

long long QAudioInput_ElapsedUSecs(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->elapsedUSecs();
}

long long QAudioInput_Error(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->error();
}

void* QAudioInput_Format(void* ptr)
{
	return new QAudioFormat(static_cast<QAudioInput*>(ptr)->format());
}

int QAudioInput_NotifyInterval(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->notifyInterval();
}

int QAudioInput_PeriodSize(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->periodSize();
}

long long QAudioInput_ProcessedUSecs(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->processedUSecs();
}

void QAudioInput_Reset(void* ptr)
{
	static_cast<QAudioInput*>(ptr)->reset();
}

void QAudioInput_Resume(void* ptr)
{
	static_cast<QAudioInput*>(ptr)->resume();
}

void QAudioInput_SetBufferSize(void* ptr, int value)
{
	static_cast<QAudioInput*>(ptr)->setBufferSize(value);
}

void QAudioInput_SetNotifyInterval(void* ptr, int ms)
{
	static_cast<QAudioInput*>(ptr)->setNotifyInterval(ms);
}

void QAudioInput_SetVolume(void* ptr, double volume)
{
	static_cast<QAudioInput*>(ptr)->setVolume(volume);
}

void* QAudioInput_Start2(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->start();
}

void QAudioInput_Start(void* ptr, void* device)
{
	static_cast<QAudioInput*>(ptr)->start(static_cast<QIODevice*>(device));
}

long long QAudioInput_State(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->state();
}

void QAudioInput_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioInput*>(ptr), static_cast<void (QAudioInput::*)(QAudio::State)>(&QAudioInput::stateChanged), static_cast<MyQAudioInput*>(ptr), static_cast<void (MyQAudioInput::*)(QAudio::State)>(&MyQAudioInput::Signal_StateChanged));
}

void QAudioInput_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioInput*>(ptr), static_cast<void (QAudioInput::*)(QAudio::State)>(&QAudioInput::stateChanged), static_cast<MyQAudioInput*>(ptr), static_cast<void (MyQAudioInput::*)(QAudio::State)>(&MyQAudioInput::Signal_StateChanged));
}

void QAudioInput_StateChanged(void* ptr, long long state)
{
	static_cast<QAudioInput*>(ptr)->stateChanged(static_cast<QAudio::State>(state));
}

void QAudioInput_Stop(void* ptr)
{
	static_cast<QAudioInput*>(ptr)->stop();
}

void QAudioInput_Suspend(void* ptr)
{
	static_cast<QAudioInput*>(ptr)->suspend();
}

double QAudioInput_Volume(void* ptr)
{
	return static_cast<QAudioInput*>(ptr)->volume();
}

void QAudioInput_DestroyQAudioInput(void* ptr)
{
	static_cast<QAudioInput*>(ptr)->~QAudioInput();
}

void QAudioInput_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioInput*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioInput_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioInput*>(ptr)->QAudioInput::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioInput_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioInput*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioInput_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioInput*>(ptr)->QAudioInput::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioInput_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioInput*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioInput_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioInput*>(ptr)->QAudioInput::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioInput_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioInput*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioInput_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioInput*>(ptr)->QAudioInput::customEvent(static_cast<QEvent*>(event));
}

void QAudioInput_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioInput*>(ptr), "deleteLater");
}

void QAudioInput_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioInput*>(ptr)->QAudioInput::deleteLater();
}

void QAudioInput_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioInput*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioInput_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioInput*>(ptr)->QAudioInput::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioInput_Event(void* ptr, void* e)
{
	return static_cast<QAudioInput*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioInput_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioInput*>(ptr)->QAudioInput::event(static_cast<QEvent*>(e));
}

char QAudioInput_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioInput*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioInput_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioInput*>(ptr)->QAudioInput::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioInput_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioInput*>(ptr)->metaObject());
}

void* QAudioInput_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioInput*>(ptr)->QAudioInput::metaObject());
}

class MyQAudioInputSelectorControl: public QAudioInputSelectorControl
{
public:
	QString activeInput() const { return QString(callbackQAudioInputSelectorControl_ActiveInput(const_cast<MyQAudioInputSelectorControl*>(this))); };
	void Signal_ActiveInputChanged(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtMultimedia_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQAudioInputSelectorControl_ActiveInputChanged(this, namePacked); };
	void Signal_AvailableInputsChanged() { callbackQAudioInputSelectorControl_AvailableInputsChanged(this); };
	QString defaultInput() const { return QString(callbackQAudioInputSelectorControl_DefaultInput(const_cast<MyQAudioInputSelectorControl*>(this))); };
	QString inputDescription(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtMultimedia_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return QString(callbackQAudioInputSelectorControl_InputDescription(const_cast<MyQAudioInputSelectorControl*>(this), namePacked)); };
	void setActiveInput(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtMultimedia_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQAudioInputSelectorControl_SetActiveInput(this, namePacked); };
	 ~MyQAudioInputSelectorControl() { callbackQAudioInputSelectorControl_DestroyQAudioInputSelectorControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQAudioInputSelectorControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioInputSelectorControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioInputSelectorControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioInputSelectorControl_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioInputSelectorControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioInputSelectorControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioInputSelectorControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioInputSelectorControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioInputSelectorControl_MetaObject(const_cast<MyQAudioInputSelectorControl*>(this))); };
};

struct QtMultimedia_PackedString QAudioInputSelectorControl_ActiveInput(void* ptr)
{
	return ({ QByteArray t730771 = static_cast<QAudioInputSelectorControl*>(ptr)->activeInput().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t730771.prepend("WHITESPACE").constData()+10), t730771.size()-10 }; });
}

void QAudioInputSelectorControl_ConnectActiveInputChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)(const QString &)>(&QAudioInputSelectorControl::activeInputChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)(const QString &)>(&MyQAudioInputSelectorControl::Signal_ActiveInputChanged));
}

void QAudioInputSelectorControl_DisconnectActiveInputChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)(const QString &)>(&QAudioInputSelectorControl::activeInputChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)(const QString &)>(&MyQAudioInputSelectorControl::Signal_ActiveInputChanged));
}

void QAudioInputSelectorControl_ActiveInputChanged(void* ptr, char* name)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->activeInputChanged(QString(name));
}

void QAudioInputSelectorControl_ConnectAvailableInputsChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)()>(&QAudioInputSelectorControl::availableInputsChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)()>(&MyQAudioInputSelectorControl::Signal_AvailableInputsChanged));
}

void QAudioInputSelectorControl_DisconnectAvailableInputsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)()>(&QAudioInputSelectorControl::availableInputsChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)()>(&MyQAudioInputSelectorControl::Signal_AvailableInputsChanged));
}

void QAudioInputSelectorControl_AvailableInputsChanged(void* ptr)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->availableInputsChanged();
}

struct QtMultimedia_PackedString QAudioInputSelectorControl_DefaultInput(void* ptr)
{
	return ({ QByteArray tdf7663 = static_cast<QAudioInputSelectorControl*>(ptr)->defaultInput().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tdf7663.prepend("WHITESPACE").constData()+10), tdf7663.size()-10 }; });
}

struct QtMultimedia_PackedString QAudioInputSelectorControl_InputDescription(void* ptr, char* name)
{
	return ({ QByteArray tf9d830 = static_cast<QAudioInputSelectorControl*>(ptr)->inputDescription(QString(name)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tf9d830.prepend("WHITESPACE").constData()+10), tf9d830.size()-10 }; });
}

void QAudioInputSelectorControl_SetActiveInput(void* ptr, char* name)
{
	QMetaObject::invokeMethod(static_cast<QAudioInputSelectorControl*>(ptr), "setActiveInput", Q_ARG(QString, QString(name)));
}

void QAudioInputSelectorControl_DestroyQAudioInputSelectorControl(void* ptr)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->~QAudioInputSelectorControl();
}

void QAudioInputSelectorControl_DestroyQAudioInputSelectorControlDefault(void* ptr)
{

}

void QAudioInputSelectorControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioInputSelectorControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->QAudioInputSelectorControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioInputSelectorControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioInputSelectorControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->QAudioInputSelectorControl::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioInputSelectorControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioInputSelectorControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->QAudioInputSelectorControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioInputSelectorControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioInputSelectorControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->QAudioInputSelectorControl::customEvent(static_cast<QEvent*>(event));
}

void QAudioInputSelectorControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioInputSelectorControl*>(ptr), "deleteLater");
}

void QAudioInputSelectorControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->QAudioInputSelectorControl::deleteLater();
}

void QAudioInputSelectorControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioInputSelectorControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioInputSelectorControl*>(ptr)->QAudioInputSelectorControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioInputSelectorControl_Event(void* ptr, void* e)
{
	return static_cast<QAudioInputSelectorControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioInputSelectorControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioInputSelectorControl*>(ptr)->QAudioInputSelectorControl::event(static_cast<QEvent*>(e));
}

char QAudioInputSelectorControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioInputSelectorControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioInputSelectorControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioInputSelectorControl*>(ptr)->QAudioInputSelectorControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioInputSelectorControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioInputSelectorControl*>(ptr)->metaObject());
}

void* QAudioInputSelectorControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioInputSelectorControl*>(ptr)->QAudioInputSelectorControl::metaObject());
}

class MyQAudioOutput: public QAudioOutput
{
public:
	MyQAudioOutput(const QAudioDeviceInfo &audioDevice, const QAudioFormat &format, QObject *parent) : QAudioOutput(audioDevice, format, parent) {};
	MyQAudioOutput(const QAudioFormat &format, QObject *parent) : QAudioOutput(format, parent) {};
	void Signal_StateChanged(QAudio::State state) { callbackQAudioOutput_StateChanged(this, state); };
	void timerEvent(QTimerEvent * event) { callbackQAudioOutput_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioOutput_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioOutput_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioOutput_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioOutput_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioOutput_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioOutput_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioOutput_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioOutput_MetaObject(const_cast<MyQAudioOutput*>(this))); };
};

void* QAudioOutput_NewQAudioOutput2(void* audioDevice, void* format, void* parent)
{
	return new MyQAudioOutput(*static_cast<QAudioDeviceInfo*>(audioDevice), *static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

void* QAudioOutput_NewQAudioOutput(void* format, void* parent)
{
	return new MyQAudioOutput(*static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

int QAudioOutput_BufferSize(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->bufferSize();
}

int QAudioOutput_BytesFree(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->bytesFree();
}

struct QtMultimedia_PackedString QAudioOutput_Category(void* ptr)
{
	return ({ QByteArray t5fbcba = static_cast<QAudioOutput*>(ptr)->category().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t5fbcba.prepend("WHITESPACE").constData()+10), t5fbcba.size()-10 }; });
}

long long QAudioOutput_ElapsedUSecs(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->elapsedUSecs();
}

long long QAudioOutput_Error(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->error();
}

void* QAudioOutput_Format(void* ptr)
{
	return new QAudioFormat(static_cast<QAudioOutput*>(ptr)->format());
}

int QAudioOutput_NotifyInterval(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->notifyInterval();
}

int QAudioOutput_PeriodSize(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->periodSize();
}

long long QAudioOutput_ProcessedUSecs(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->processedUSecs();
}

void QAudioOutput_Reset(void* ptr)
{
	static_cast<QAudioOutput*>(ptr)->reset();
}

void QAudioOutput_Resume(void* ptr)
{
	static_cast<QAudioOutput*>(ptr)->resume();
}

void QAudioOutput_SetBufferSize(void* ptr, int value)
{
	static_cast<QAudioOutput*>(ptr)->setBufferSize(value);
}

void QAudioOutput_SetCategory(void* ptr, char* category)
{
	static_cast<QAudioOutput*>(ptr)->setCategory(QString(category));
}

void QAudioOutput_SetNotifyInterval(void* ptr, int ms)
{
	static_cast<QAudioOutput*>(ptr)->setNotifyInterval(ms);
}

void QAudioOutput_SetVolume(void* ptr, double volume)
{
	static_cast<QAudioOutput*>(ptr)->setVolume(volume);
}

void* QAudioOutput_Start2(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->start();
}

void QAudioOutput_Start(void* ptr, void* device)
{
	static_cast<QAudioOutput*>(ptr)->start(static_cast<QIODevice*>(device));
}

long long QAudioOutput_State(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->state();
}

void QAudioOutput_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioOutput*>(ptr), static_cast<void (QAudioOutput::*)(QAudio::State)>(&QAudioOutput::stateChanged), static_cast<MyQAudioOutput*>(ptr), static_cast<void (MyQAudioOutput::*)(QAudio::State)>(&MyQAudioOutput::Signal_StateChanged));
}

void QAudioOutput_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioOutput*>(ptr), static_cast<void (QAudioOutput::*)(QAudio::State)>(&QAudioOutput::stateChanged), static_cast<MyQAudioOutput*>(ptr), static_cast<void (MyQAudioOutput::*)(QAudio::State)>(&MyQAudioOutput::Signal_StateChanged));
}

void QAudioOutput_StateChanged(void* ptr, long long state)
{
	static_cast<QAudioOutput*>(ptr)->stateChanged(static_cast<QAudio::State>(state));
}

void QAudioOutput_Stop(void* ptr)
{
	static_cast<QAudioOutput*>(ptr)->stop();
}

void QAudioOutput_Suspend(void* ptr)
{
	static_cast<QAudioOutput*>(ptr)->suspend();
}

double QAudioOutput_Volume(void* ptr)
{
	return static_cast<QAudioOutput*>(ptr)->volume();
}

void QAudioOutput_DestroyQAudioOutput(void* ptr)
{
	static_cast<QAudioOutput*>(ptr)->~QAudioOutput();
}

void QAudioOutput_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioOutput*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioOutput_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioOutput*>(ptr)->QAudioOutput::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioOutput_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioOutput*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioOutput_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioOutput*>(ptr)->QAudioOutput::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioOutput_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioOutput*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioOutput_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioOutput*>(ptr)->QAudioOutput::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioOutput_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioOutput*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioOutput_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioOutput*>(ptr)->QAudioOutput::customEvent(static_cast<QEvent*>(event));
}

void QAudioOutput_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioOutput*>(ptr), "deleteLater");
}

void QAudioOutput_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioOutput*>(ptr)->QAudioOutput::deleteLater();
}

void QAudioOutput_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioOutput*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioOutput_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioOutput*>(ptr)->QAudioOutput::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioOutput_Event(void* ptr, void* e)
{
	return static_cast<QAudioOutput*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioOutput_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioOutput*>(ptr)->QAudioOutput::event(static_cast<QEvent*>(e));
}

char QAudioOutput_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioOutput*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioOutput_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioOutput*>(ptr)->QAudioOutput::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioOutput_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioOutput*>(ptr)->metaObject());
}

void* QAudioOutput_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioOutput*>(ptr)->QAudioOutput::metaObject());
}

class MyQAudioOutputSelectorControl: public QAudioOutputSelectorControl
{
public:
	QString activeOutput() const { return QString(callbackQAudioOutputSelectorControl_ActiveOutput(const_cast<MyQAudioOutputSelectorControl*>(this))); };
	void Signal_ActiveOutputChanged(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtMultimedia_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQAudioOutputSelectorControl_ActiveOutputChanged(this, namePacked); };
	void Signal_AvailableOutputsChanged() { callbackQAudioOutputSelectorControl_AvailableOutputsChanged(this); };
	QString defaultOutput() const { return QString(callbackQAudioOutputSelectorControl_DefaultOutput(const_cast<MyQAudioOutputSelectorControl*>(this))); };
	QString outputDescription(const QString & name) const { QByteArray t6ae999 = name.toUtf8(); QtMultimedia_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };return QString(callbackQAudioOutputSelectorControl_OutputDescription(const_cast<MyQAudioOutputSelectorControl*>(this), namePacked)); };
	void setActiveOutput(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtMultimedia_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQAudioOutputSelectorControl_SetActiveOutput(this, namePacked); };
	 ~MyQAudioOutputSelectorControl() { callbackQAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQAudioOutputSelectorControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioOutputSelectorControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioOutputSelectorControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioOutputSelectorControl_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioOutputSelectorControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioOutputSelectorControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioOutputSelectorControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioOutputSelectorControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioOutputSelectorControl_MetaObject(const_cast<MyQAudioOutputSelectorControl*>(this))); };
};

struct QtMultimedia_PackedString QAudioOutputSelectorControl_ActiveOutput(void* ptr)
{
	return ({ QByteArray t4c7caf = static_cast<QAudioOutputSelectorControl*>(ptr)->activeOutput().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t4c7caf.prepend("WHITESPACE").constData()+10), t4c7caf.size()-10 }; });
}

void QAudioOutputSelectorControl_ConnectActiveOutputChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)(const QString &)>(&QAudioOutputSelectorControl::activeOutputChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)(const QString &)>(&MyQAudioOutputSelectorControl::Signal_ActiveOutputChanged));
}

void QAudioOutputSelectorControl_DisconnectActiveOutputChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)(const QString &)>(&QAudioOutputSelectorControl::activeOutputChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)(const QString &)>(&MyQAudioOutputSelectorControl::Signal_ActiveOutputChanged));
}

void QAudioOutputSelectorControl_ActiveOutputChanged(void* ptr, char* name)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->activeOutputChanged(QString(name));
}

void QAudioOutputSelectorControl_ConnectAvailableOutputsChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)()>(&QAudioOutputSelectorControl::availableOutputsChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)()>(&MyQAudioOutputSelectorControl::Signal_AvailableOutputsChanged));
}

void QAudioOutputSelectorControl_DisconnectAvailableOutputsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)()>(&QAudioOutputSelectorControl::availableOutputsChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)()>(&MyQAudioOutputSelectorControl::Signal_AvailableOutputsChanged));
}

void QAudioOutputSelectorControl_AvailableOutputsChanged(void* ptr)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->availableOutputsChanged();
}

struct QtMultimedia_PackedString QAudioOutputSelectorControl_DefaultOutput(void* ptr)
{
	return ({ QByteArray t3b7fff = static_cast<QAudioOutputSelectorControl*>(ptr)->defaultOutput().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t3b7fff.prepend("WHITESPACE").constData()+10), t3b7fff.size()-10 }; });
}

struct QtMultimedia_PackedString QAudioOutputSelectorControl_OutputDescription(void* ptr, char* name)
{
	return ({ QByteArray t1cf920 = static_cast<QAudioOutputSelectorControl*>(ptr)->outputDescription(QString(name)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t1cf920.prepend("WHITESPACE").constData()+10), t1cf920.size()-10 }; });
}

void QAudioOutputSelectorControl_SetActiveOutput(void* ptr, char* name)
{
	QMetaObject::invokeMethod(static_cast<QAudioOutputSelectorControl*>(ptr), "setActiveOutput", Q_ARG(QString, QString(name)));
}

void QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(void* ptr)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->~QAudioOutputSelectorControl();
}

void QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControlDefault(void* ptr)
{

}

void QAudioOutputSelectorControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioOutputSelectorControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->QAudioOutputSelectorControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioOutputSelectorControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioOutputSelectorControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->QAudioOutputSelectorControl::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioOutputSelectorControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioOutputSelectorControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->QAudioOutputSelectorControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioOutputSelectorControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioOutputSelectorControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->QAudioOutputSelectorControl::customEvent(static_cast<QEvent*>(event));
}

void QAudioOutputSelectorControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioOutputSelectorControl*>(ptr), "deleteLater");
}

void QAudioOutputSelectorControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->QAudioOutputSelectorControl::deleteLater();
}

void QAudioOutputSelectorControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioOutputSelectorControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioOutputSelectorControl*>(ptr)->QAudioOutputSelectorControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioOutputSelectorControl_Event(void* ptr, void* e)
{
	return static_cast<QAudioOutputSelectorControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioOutputSelectorControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioOutputSelectorControl*>(ptr)->QAudioOutputSelectorControl::event(static_cast<QEvent*>(e));
}

char QAudioOutputSelectorControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioOutputSelectorControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioOutputSelectorControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioOutputSelectorControl*>(ptr)->QAudioOutputSelectorControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioOutputSelectorControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioOutputSelectorControl*>(ptr)->metaObject());
}

void* QAudioOutputSelectorControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioOutputSelectorControl*>(ptr)->QAudioOutputSelectorControl::metaObject());
}

class MyQAudioProbe: public QAudioProbe
{
public:
	MyQAudioProbe(QObject *parent) : QAudioProbe(parent) {};
	void Signal_AudioBufferProbed(const QAudioBuffer & buffer) { callbackQAudioProbe_AudioBufferProbed(this, const_cast<QAudioBuffer*>(&buffer)); };
	void Signal_Flush() { callbackQAudioProbe_Flush(this); };
	void timerEvent(QTimerEvent * event) { callbackQAudioProbe_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioProbe_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioProbe_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioProbe_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioProbe_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioProbe_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioProbe_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioProbe_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioProbe_MetaObject(const_cast<MyQAudioProbe*>(this))); };
};

void* QAudioProbe_NewQAudioProbe(void* parent)
{
	return new MyQAudioProbe(static_cast<QObject*>(parent));
}

void QAudioProbe_ConnectAudioBufferProbed(void* ptr)
{
	QObject::connect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)(const QAudioBuffer &)>(&QAudioProbe::audioBufferProbed), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)(const QAudioBuffer &)>(&MyQAudioProbe::Signal_AudioBufferProbed));
}

void QAudioProbe_DisconnectAudioBufferProbed(void* ptr)
{
	QObject::disconnect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)(const QAudioBuffer &)>(&QAudioProbe::audioBufferProbed), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)(const QAudioBuffer &)>(&MyQAudioProbe::Signal_AudioBufferProbed));
}

void QAudioProbe_AudioBufferProbed(void* ptr, void* buffer)
{
	static_cast<QAudioProbe*>(ptr)->audioBufferProbed(*static_cast<QAudioBuffer*>(buffer));
}

void QAudioProbe_ConnectFlush(void* ptr)
{
	QObject::connect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)()>(&QAudioProbe::flush), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)()>(&MyQAudioProbe::Signal_Flush));
}

void QAudioProbe_DisconnectFlush(void* ptr)
{
	QObject::disconnect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)()>(&QAudioProbe::flush), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)()>(&MyQAudioProbe::Signal_Flush));
}

void QAudioProbe_Flush(void* ptr)
{
	static_cast<QAudioProbe*>(ptr)->flush();
}

char QAudioProbe_IsActive(void* ptr)
{
	return static_cast<QAudioProbe*>(ptr)->isActive();
}

char QAudioProbe_SetSource(void* ptr, void* source)
{
	return static_cast<QAudioProbe*>(ptr)->setSource(static_cast<QMediaObject*>(source));
}

char QAudioProbe_SetSource2(void* ptr, void* mediaRecorder)
{
	return static_cast<QAudioProbe*>(ptr)->setSource(static_cast<QMediaRecorder*>(mediaRecorder));
}

void QAudioProbe_DestroyQAudioProbe(void* ptr)
{
	static_cast<QAudioProbe*>(ptr)->~QAudioProbe();
}

void QAudioProbe_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioProbe*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioProbe_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioProbe*>(ptr)->QAudioProbe::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioProbe_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioProbe*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioProbe_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioProbe*>(ptr)->QAudioProbe::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioProbe_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioProbe*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioProbe_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioProbe*>(ptr)->QAudioProbe::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioProbe_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioProbe*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioProbe_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioProbe*>(ptr)->QAudioProbe::customEvent(static_cast<QEvent*>(event));
}

void QAudioProbe_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioProbe*>(ptr), "deleteLater");
}

void QAudioProbe_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioProbe*>(ptr)->QAudioProbe::deleteLater();
}

void QAudioProbe_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioProbe*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioProbe_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioProbe*>(ptr)->QAudioProbe::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioProbe_Event(void* ptr, void* e)
{
	return static_cast<QAudioProbe*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioProbe_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioProbe*>(ptr)->QAudioProbe::event(static_cast<QEvent*>(e));
}

char QAudioProbe_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioProbe*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioProbe_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioProbe*>(ptr)->QAudioProbe::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioProbe_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioProbe*>(ptr)->metaObject());
}

void* QAudioProbe_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioProbe*>(ptr)->QAudioProbe::metaObject());
}

class MyQAudioRecorder: public QAudioRecorder
{
public:
	MyQAudioRecorder(QObject *parent) : QAudioRecorder(parent) {};
	void Signal_AudioInputChanged(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtMultimedia_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQAudioRecorder_AudioInputChanged(this, namePacked); };
	void Signal_AvailableAudioInputsChanged() { callbackQAudioRecorder_AvailableAudioInputsChanged(this); };
	void setAudioInput(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtMultimedia_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQAudioRecorder_SetAudioInput(this, namePacked); };
	void setMuted(bool muted) { callbackQAudioRecorder_SetMuted(this, muted); };
	void setVolume(qreal volume) { callbackQAudioRecorder_SetVolume(this, volume); };
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQAudioRecorder_MediaObject(const_cast<MyQAudioRecorder*>(this))); };
	void pause() { callbackQAudioRecorder_Pause(this); };
	void record() { callbackQAudioRecorder_Record(this); };
	void stop() { callbackQAudioRecorder_Stop(this); };
	void timerEvent(QTimerEvent * event) { callbackQAudioRecorder_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioRecorder_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioRecorder_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioRecorder_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioRecorder_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioRecorder_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioRecorder_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioRecorder_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioRecorder_MetaObject(const_cast<MyQAudioRecorder*>(this))); };
	bool setMediaObject(QMediaObject * object) { return callbackQAudioRecorder_SetMediaObject(this, object) != 0; };
};

void* QAudioRecorder_NewQAudioRecorder(void* parent)
{
	return new MyQAudioRecorder(static_cast<QObject*>(parent));
}

struct QtMultimedia_PackedString QAudioRecorder_AudioInput(void* ptr)
{
	return ({ QByteArray t7b4459 = static_cast<QAudioRecorder*>(ptr)->audioInput().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t7b4459.prepend("WHITESPACE").constData()+10), t7b4459.size()-10 }; });
}

void QAudioRecorder_ConnectAudioInputChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)(const QString &)>(&QAudioRecorder::audioInputChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)(const QString &)>(&MyQAudioRecorder::Signal_AudioInputChanged));
}

void QAudioRecorder_DisconnectAudioInputChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)(const QString &)>(&QAudioRecorder::audioInputChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)(const QString &)>(&MyQAudioRecorder::Signal_AudioInputChanged));
}

void QAudioRecorder_AudioInputChanged(void* ptr, char* name)
{
	static_cast<QAudioRecorder*>(ptr)->audioInputChanged(QString(name));
}

struct QtMultimedia_PackedString QAudioRecorder_AudioInputDescription(void* ptr, char* name)
{
	return ({ QByteArray t492593 = static_cast<QAudioRecorder*>(ptr)->audioInputDescription(QString(name)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t492593.prepend("WHITESPACE").constData()+10), t492593.size()-10 }; });
}

struct QtMultimedia_PackedString QAudioRecorder_AudioInputs(void* ptr)
{
	return ({ QByteArray tf48c37 = static_cast<QAudioRecorder*>(ptr)->audioInputs().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tf48c37.prepend("WHITESPACE").constData()+10), tf48c37.size()-10 }; });
}

void QAudioRecorder_ConnectAvailableAudioInputsChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)()>(&QAudioRecorder::availableAudioInputsChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)()>(&MyQAudioRecorder::Signal_AvailableAudioInputsChanged));
}

void QAudioRecorder_DisconnectAvailableAudioInputsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)()>(&QAudioRecorder::availableAudioInputsChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)()>(&MyQAudioRecorder::Signal_AvailableAudioInputsChanged));
}

void QAudioRecorder_AvailableAudioInputsChanged(void* ptr)
{
	static_cast<QAudioRecorder*>(ptr)->availableAudioInputsChanged();
}

struct QtMultimedia_PackedString QAudioRecorder_DefaultAudioInput(void* ptr)
{
	return ({ QByteArray tf2898e = static_cast<QAudioRecorder*>(ptr)->defaultAudioInput().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tf2898e.prepend("WHITESPACE").constData()+10), tf2898e.size()-10 }; });
}

void QAudioRecorder_SetAudioInput(void* ptr, char* name)
{
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "setAudioInput", Q_ARG(QString, QString(name)));
}

void QAudioRecorder_DestroyQAudioRecorder(void* ptr)
{
	static_cast<QAudioRecorder*>(ptr)->~QAudioRecorder();
}

void QAudioRecorder_SetMuted(void* ptr, char muted)
{
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

void QAudioRecorder_SetMutedDefault(void* ptr, char muted)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::setMuted(muted != 0);
}

void QAudioRecorder_SetVolume(void* ptr, double volume)
{
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "setVolume", Q_ARG(qreal, volume));
}

void QAudioRecorder_SetVolumeDefault(void* ptr, double volume)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::setVolume(volume);
}

void* QAudioRecorder_MediaObject(void* ptr)
{
	return static_cast<QAudioRecorder*>(ptr)->mediaObject();
}

void* QAudioRecorder_MediaObjectDefault(void* ptr)
{
	return static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::mediaObject();
}

void QAudioRecorder_Pause(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "pause");
}

void QAudioRecorder_PauseDefault(void* ptr)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::pause();
}

void QAudioRecorder_Record(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "record");
}

void QAudioRecorder_RecordDefault(void* ptr)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::record();
}

void QAudioRecorder_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "stop");
}

void QAudioRecorder_StopDefault(void* ptr)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::stop();
}

void QAudioRecorder_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioRecorder*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioRecorder_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioRecorder_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioRecorder*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioRecorder_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioRecorder_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioRecorder*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioRecorder_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioRecorder_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioRecorder*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioRecorder_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::customEvent(static_cast<QEvent*>(event));
}

void QAudioRecorder_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "deleteLater");
}

void QAudioRecorder_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::deleteLater();
}

void QAudioRecorder_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioRecorder*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioRecorder_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioRecorder_Event(void* ptr, void* e)
{
	return static_cast<QAudioRecorder*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioRecorder_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::event(static_cast<QEvent*>(e));
}

char QAudioRecorder_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioRecorder*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioRecorder_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioRecorder_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioRecorder*>(ptr)->metaObject());
}

void* QAudioRecorder_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::metaObject());
}

char QAudioRecorder_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QAudioRecorder*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

char QAudioRecorder_SetMediaObjectDefault(void* ptr, void* object)
{
	return static_cast<QAudioRecorder*>(ptr)->QAudioRecorder::setMediaObject(static_cast<QMediaObject*>(object));
}

class MyQAudioRoleControl: public QAudioRoleControl
{
public:
	QAudio::Role audioRole() const { return static_cast<QAudio::Role>(callbackQAudioRoleControl_AudioRole(const_cast<MyQAudioRoleControl*>(this))); };
	void Signal_AudioRoleChanged(QAudio::Role role) { callbackQAudioRoleControl_AudioRoleChanged(this, role); };
	void setAudioRole(QAudio::Role role) { callbackQAudioRoleControl_SetAudioRole(this, role); };
	 ~MyQAudioRoleControl() { callbackQAudioRoleControl_DestroyQAudioRoleControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQAudioRoleControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQAudioRoleControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAudioRoleControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAudioRoleControl_CustomEvent(this, event); };
	void deleteLater() { callbackQAudioRoleControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAudioRoleControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAudioRoleControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAudioRoleControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAudioRoleControl_MetaObject(const_cast<MyQAudioRoleControl*>(this))); };
};

long long QAudioRoleControl_AudioRole(void* ptr)
{
	return static_cast<QAudioRoleControl*>(ptr)->audioRole();
}

void QAudioRoleControl_ConnectAudioRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QAudioRoleControl*>(ptr), static_cast<void (QAudioRoleControl::*)(QAudio::Role)>(&QAudioRoleControl::audioRoleChanged), static_cast<MyQAudioRoleControl*>(ptr), static_cast<void (MyQAudioRoleControl::*)(QAudio::Role)>(&MyQAudioRoleControl::Signal_AudioRoleChanged));
}

void QAudioRoleControl_DisconnectAudioRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAudioRoleControl*>(ptr), static_cast<void (QAudioRoleControl::*)(QAudio::Role)>(&QAudioRoleControl::audioRoleChanged), static_cast<MyQAudioRoleControl*>(ptr), static_cast<void (MyQAudioRoleControl::*)(QAudio::Role)>(&MyQAudioRoleControl::Signal_AudioRoleChanged));
}

void QAudioRoleControl_AudioRoleChanged(void* ptr, long long role)
{
	static_cast<QAudioRoleControl*>(ptr)->audioRoleChanged(static_cast<QAudio::Role>(role));
}

void QAudioRoleControl_SetAudioRole(void* ptr, long long role)
{
	static_cast<QAudioRoleControl*>(ptr)->setAudioRole(static_cast<QAudio::Role>(role));
}

void QAudioRoleControl_DestroyQAudioRoleControl(void* ptr)
{
	static_cast<QAudioRoleControl*>(ptr)->~QAudioRoleControl();
}

void QAudioRoleControl_DestroyQAudioRoleControlDefault(void* ptr)
{

}

void QAudioRoleControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QAudioRoleControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioRoleControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAudioRoleControl*>(ptr)->QAudioRoleControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAudioRoleControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QAudioRoleControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAudioRoleControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAudioRoleControl*>(ptr)->QAudioRoleControl::childEvent(static_cast<QChildEvent*>(event));
}

void QAudioRoleControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioRoleControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioRoleControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioRoleControl*>(ptr)->QAudioRoleControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioRoleControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QAudioRoleControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAudioRoleControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAudioRoleControl*>(ptr)->QAudioRoleControl::customEvent(static_cast<QEvent*>(event));
}

void QAudioRoleControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAudioRoleControl*>(ptr), "deleteLater");
}

void QAudioRoleControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QAudioRoleControl*>(ptr)->QAudioRoleControl::deleteLater();
}

void QAudioRoleControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAudioRoleControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAudioRoleControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAudioRoleControl*>(ptr)->QAudioRoleControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAudioRoleControl_Event(void* ptr, void* e)
{
	return static_cast<QAudioRoleControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QAudioRoleControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QAudioRoleControl*>(ptr)->QAudioRoleControl::event(static_cast<QEvent*>(e));
}

char QAudioRoleControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioRoleControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QAudioRoleControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAudioRoleControl*>(ptr)->QAudioRoleControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAudioRoleControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioRoleControl*>(ptr)->metaObject());
}

void* QAudioRoleControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAudioRoleControl*>(ptr)->QAudioRoleControl::metaObject());
}

class MyQCamera: public QCamera
{
public:
	MyQCamera(QCamera::Position position, QObject *parent) : QCamera(position, parent) {};
	MyQCamera(QObject *parent) : QCamera(parent) {};
	MyQCamera(const QByteArray &deviceName, QObject *parent) : QCamera(deviceName, parent) {};
	MyQCamera(const QCameraInfo &cameraInfo, QObject *parent) : QCamera(cameraInfo, parent) {};
	void searchAndLock(QCamera::LockTypes locks) { callbackQCamera_SearchAndLock2(this, locks); };
	void setCaptureMode(QCamera::CaptureModes mode) { callbackQCamera_SetCaptureMode(this, mode); };
	QMultimedia::AvailabilityStatus availability() const { return static_cast<QMultimedia::AvailabilityStatus>(callbackQCamera_Availability(const_cast<MyQCamera*>(this))); };
	void Signal_CaptureModeChanged(QCamera::CaptureModes mode) { callbackQCamera_CaptureModeChanged(this, mode); };
	void Signal_Error2(QCamera::Error value) { callbackQCamera_Error2(this, value); };
	void load() { callbackQCamera_Load(this); };
	void Signal_LockFailed() { callbackQCamera_LockFailed(this); };
	void Signal_LockStatusChanged(QCamera::LockStatus status, QCamera::LockChangeReason reason) { callbackQCamera_LockStatusChanged(this, status, reason); };
	void Signal_LockStatusChanged2(QCamera::LockType lock, QCamera::LockStatus status, QCamera::LockChangeReason reason) { callbackQCamera_LockStatusChanged2(this, lock, status, reason); };
	void Signal_Locked() { callbackQCamera_Locked(this); };
	void searchAndLock() { callbackQCamera_SearchAndLock(this); };
	void start() { callbackQCamera_Start(this); };
	void Signal_StateChanged(QCamera::State state) { callbackQCamera_StateChanged(this, state); };
	void Signal_StatusChanged(QCamera::Status status) { callbackQCamera_StatusChanged(this, status); };
	void stop() { callbackQCamera_Stop(this); };
	void unload() { callbackQCamera_Unload(this); };
	void unlock() { callbackQCamera_Unlock(this); };
	void unlock(QCamera::LockTypes locks) { callbackQCamera_Unlock2(this, locks); };
	bool bind(QObject * object) { return callbackQCamera_Bind(this, object) != 0; };
	bool isAvailable() const { return callbackQCamera_IsAvailable(const_cast<MyQCamera*>(this)) != 0; };
	QMediaService * service() const { return static_cast<QMediaService*>(callbackQCamera_Service(const_cast<MyQCamera*>(this))); };
	void unbind(QObject * object) { callbackQCamera_Unbind(this, object); };
	void timerEvent(QTimerEvent * event) { callbackQCamera_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCamera_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCamera_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCamera_CustomEvent(this, event); };
	void deleteLater() { callbackQCamera_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCamera_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCamera_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCamera_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCamera_MetaObject(const_cast<MyQCamera*>(this))); };
};

long long QCamera_CaptureMode(void* ptr)
{
	return static_cast<QCamera*>(ptr)->captureMode();
}

void QCamera_SearchAndLock2(void* ptr, long long locks)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "searchAndLock", Q_ARG(QCamera::LockType, static_cast<QCamera::LockType>(locks)));
}

void QCamera_SetCaptureMode(void* ptr, long long mode)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "setCaptureMode", Q_ARG(QCamera::CaptureMode, static_cast<QCamera::CaptureMode>(mode)));
}

long long QCamera_State(void* ptr)
{
	return static_cast<QCamera*>(ptr)->state();
}

long long QCamera_Status(void* ptr)
{
	return static_cast<QCamera*>(ptr)->status();
}

void* QCamera_NewQCamera4(long long position, void* parent)
{
	return new MyQCamera(static_cast<QCamera::Position>(position), static_cast<QObject*>(parent));
}

void* QCamera_NewQCamera(void* parent)
{
	return new MyQCamera(static_cast<QObject*>(parent));
}

void* QCamera_NewQCamera2(void* deviceName, void* parent)
{
	return new MyQCamera(*static_cast<QByteArray*>(deviceName), static_cast<QObject*>(parent));
}

void* QCamera_NewQCamera3(void* cameraInfo, void* parent)
{
	return new MyQCamera(*static_cast<QCameraInfo*>(cameraInfo), static_cast<QObject*>(parent));
}

long long QCamera_Availability(void* ptr)
{
	return static_cast<QCamera*>(ptr)->availability();
}

long long QCamera_AvailabilityDefault(void* ptr)
{
	return static_cast<QCamera*>(ptr)->QCamera::availability();
}

void QCamera_ConnectCaptureModeChanged(void* ptr)
{
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::CaptureModes)>(&QCamera::captureModeChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::CaptureModes)>(&MyQCamera::Signal_CaptureModeChanged));
}

void QCamera_DisconnectCaptureModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::CaptureModes)>(&QCamera::captureModeChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::CaptureModes)>(&MyQCamera::Signal_CaptureModeChanged));
}

void QCamera_CaptureModeChanged(void* ptr, long long mode)
{
	static_cast<QCamera*>(ptr)->captureModeChanged(static_cast<QCamera::CaptureMode>(mode));
}

void QCamera_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Error)>(&QCamera::error), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Error)>(&MyQCamera::Signal_Error2));
}

void QCamera_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Error)>(&QCamera::error), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Error)>(&MyQCamera::Signal_Error2));
}

void QCamera_Error2(void* ptr, long long value)
{
	static_cast<QCamera*>(ptr)->error(static_cast<QCamera::Error>(value));
}

long long QCamera_Error(void* ptr)
{
	return static_cast<QCamera*>(ptr)->error();
}

struct QtMultimedia_PackedString QCamera_ErrorString(void* ptr)
{
	return ({ QByteArray t14c12c = static_cast<QCamera*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t14c12c.prepend("WHITESPACE").constData()+10), t14c12c.size()-10 }; });
}

void* QCamera_Exposure(void* ptr)
{
	return static_cast<QCamera*>(ptr)->exposure();
}

void* QCamera_Focus(void* ptr)
{
	return static_cast<QCamera*>(ptr)->focus();
}

void* QCamera_ImageProcessing(void* ptr)
{
	return static_cast<QCamera*>(ptr)->imageProcessing();
}

char QCamera_IsCaptureModeSupported(void* ptr, long long mode)
{
	return static_cast<QCamera*>(ptr)->isCaptureModeSupported(static_cast<QCamera::CaptureMode>(mode));
}

void QCamera_Load(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "load");
}

void QCamera_ConnectLockFailed(void* ptr)
{
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::lockFailed), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_LockFailed));
}

void QCamera_DisconnectLockFailed(void* ptr)
{
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::lockFailed), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_LockFailed));
}

void QCamera_LockFailed(void* ptr)
{
	static_cast<QCamera*>(ptr)->lockFailed();
}

long long QCamera_LockStatus(void* ptr)
{
	return static_cast<QCamera*>(ptr)->lockStatus();
}

long long QCamera_LockStatus2(void* ptr, long long lockType)
{
	return static_cast<QCamera*>(ptr)->lockStatus(static_cast<QCamera::LockType>(lockType));
}

void QCamera_ConnectLockStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged));
}

void QCamera_DisconnectLockStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged));
}

void QCamera_LockStatusChanged(void* ptr, long long status, long long reason)
{
	static_cast<QCamera*>(ptr)->lockStatusChanged(static_cast<QCamera::LockStatus>(status), static_cast<QCamera::LockChangeReason>(reason));
}

void QCamera_ConnectLockStatusChanged2(void* ptr)
{
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged2));
}

void QCamera_DisconnectLockStatusChanged2(void* ptr)
{
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged2));
}

void QCamera_LockStatusChanged2(void* ptr, long long lock, long long status, long long reason)
{
	static_cast<QCamera*>(ptr)->lockStatusChanged(static_cast<QCamera::LockType>(lock), static_cast<QCamera::LockStatus>(status), static_cast<QCamera::LockChangeReason>(reason));
}

void QCamera_ConnectLocked(void* ptr)
{
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::locked), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_Locked));
}

void QCamera_DisconnectLocked(void* ptr)
{
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::locked), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_Locked));
}

void QCamera_Locked(void* ptr)
{
	static_cast<QCamera*>(ptr)->locked();
}

long long QCamera_RequestedLocks(void* ptr)
{
	return static_cast<QCamera*>(ptr)->requestedLocks();
}

void QCamera_SearchAndLock(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "searchAndLock");
}

void QCamera_SetViewfinder3(void* ptr, void* surface)
{
	static_cast<QCamera*>(ptr)->setViewfinder(static_cast<QAbstractVideoSurface*>(surface));
}

void QCamera_SetViewfinder2(void* ptr, void* viewfinder)
{
	static_cast<QCamera*>(ptr)->setViewfinder(static_cast<QGraphicsVideoItem*>(viewfinder));
}

void QCamera_SetViewfinder(void* ptr, void* viewfinder)
{
	static_cast<QCamera*>(ptr)->setViewfinder(static_cast<QVideoWidget*>(viewfinder));
}

void QCamera_SetViewfinderSettings(void* ptr, void* settings)
{
	static_cast<QCamera*>(ptr)->setViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(settings));
}

void QCamera_Start(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "start");
}

void QCamera_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::State)>(&QCamera::stateChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::State)>(&MyQCamera::Signal_StateChanged));
}

void QCamera_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::State)>(&QCamera::stateChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::State)>(&MyQCamera::Signal_StateChanged));
}

void QCamera_StateChanged(void* ptr, long long state)
{
	static_cast<QCamera*>(ptr)->stateChanged(static_cast<QCamera::State>(state));
}

void QCamera_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Status)>(&QCamera::statusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Status)>(&MyQCamera::Signal_StatusChanged));
}

void QCamera_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Status)>(&QCamera::statusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Status)>(&MyQCamera::Signal_StatusChanged));
}

void QCamera_StatusChanged(void* ptr, long long status)
{
	static_cast<QCamera*>(ptr)->statusChanged(static_cast<QCamera::Status>(status));
}

void QCamera_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "stop");
}

long long QCamera_SupportedLocks(void* ptr)
{
	return static_cast<QCamera*>(ptr)->supportedLocks();
}

void QCamera_Unload(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unload");
}

void QCamera_Unlock(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unlock");
}

void QCamera_Unlock2(void* ptr, long long locks)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unlock", Q_ARG(QCamera::LockType, static_cast<QCamera::LockType>(locks)));
}

void* QCamera_ViewfinderSettings(void* ptr)
{
	return new QCameraViewfinderSettings(static_cast<QCamera*>(ptr)->viewfinderSettings());
}

void QCamera_DestroyQCamera(void* ptr)
{
	static_cast<QCamera*>(ptr)->~QCamera();
}

char QCamera_Bind(void* ptr, void* object)
{
	return static_cast<QCamera*>(ptr)->bind(static_cast<QObject*>(object));
}

char QCamera_BindDefault(void* ptr, void* object)
{
	return static_cast<QCamera*>(ptr)->QCamera::bind(static_cast<QObject*>(object));
}

char QCamera_IsAvailable(void* ptr)
{
	return static_cast<QCamera*>(ptr)->isAvailable();
}

char QCamera_IsAvailableDefault(void* ptr)
{
	return static_cast<QCamera*>(ptr)->QCamera::isAvailable();
}

void* QCamera_Service(void* ptr)
{
	return static_cast<QCamera*>(ptr)->service();
}

void* QCamera_ServiceDefault(void* ptr)
{
	return static_cast<QCamera*>(ptr)->QCamera::service();
}

void QCamera_Unbind(void* ptr, void* object)
{
	static_cast<QCamera*>(ptr)->unbind(static_cast<QObject*>(object));
}

void QCamera_UnbindDefault(void* ptr, void* object)
{
	static_cast<QCamera*>(ptr)->QCamera::unbind(static_cast<QObject*>(object));
}

void QCamera_TimerEvent(void* ptr, void* event)
{
	static_cast<QCamera*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCamera_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCamera*>(ptr)->QCamera::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCamera_ChildEvent(void* ptr, void* event)
{
	static_cast<QCamera*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCamera_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCamera*>(ptr)->QCamera::childEvent(static_cast<QChildEvent*>(event));
}

void QCamera_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCamera*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCamera_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCamera*>(ptr)->QCamera::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCamera_CustomEvent(void* ptr, void* event)
{
	static_cast<QCamera*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCamera_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCamera*>(ptr)->QCamera::customEvent(static_cast<QEvent*>(event));
}

void QCamera_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "deleteLater");
}

void QCamera_DeleteLaterDefault(void* ptr)
{
	static_cast<QCamera*>(ptr)->QCamera::deleteLater();
}

void QCamera_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCamera*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCamera_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCamera*>(ptr)->QCamera::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCamera_Event(void* ptr, void* e)
{
	return static_cast<QCamera*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCamera_EventDefault(void* ptr, void* e)
{
	return static_cast<QCamera*>(ptr)->QCamera::event(static_cast<QEvent*>(e));
}

char QCamera_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCamera*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCamera_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCamera*>(ptr)->QCamera::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCamera_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCamera*>(ptr)->metaObject());
}

void* QCamera_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCamera*>(ptr)->QCamera::metaObject());
}

class MyQCameraCaptureBufferFormatControl: public QCameraCaptureBufferFormatControl
{
public:
	QVideoFrame::PixelFormat bufferFormat() const { return static_cast<QVideoFrame::PixelFormat>(callbackQCameraCaptureBufferFormatControl_BufferFormat(const_cast<MyQCameraCaptureBufferFormatControl*>(this))); };
	void Signal_BufferFormatChanged(QVideoFrame::PixelFormat format) { callbackQCameraCaptureBufferFormatControl_BufferFormatChanged(this, format); };
	void setBufferFormat(QVideoFrame::PixelFormat format) { callbackQCameraCaptureBufferFormatControl_SetBufferFormat(this, format); };
	void timerEvent(QTimerEvent * event) { callbackQCameraCaptureBufferFormatControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraCaptureBufferFormatControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraCaptureBufferFormatControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraCaptureBufferFormatControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraCaptureBufferFormatControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraCaptureBufferFormatControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraCaptureBufferFormatControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraCaptureBufferFormatControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraCaptureBufferFormatControl_MetaObject(const_cast<MyQCameraCaptureBufferFormatControl*>(this))); };
};

long long QCameraCaptureBufferFormatControl_BufferFormat(void* ptr)
{
	return static_cast<QCameraCaptureBufferFormatControl*>(ptr)->bufferFormat();
}

void QCameraCaptureBufferFormatControl_ConnectBufferFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraCaptureBufferFormatControl*>(ptr), static_cast<void (QCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&QCameraCaptureBufferFormatControl::bufferFormatChanged), static_cast<MyQCameraCaptureBufferFormatControl*>(ptr), static_cast<void (MyQCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&MyQCameraCaptureBufferFormatControl::Signal_BufferFormatChanged));
}

void QCameraCaptureBufferFormatControl_DisconnectBufferFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraCaptureBufferFormatControl*>(ptr), static_cast<void (QCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&QCameraCaptureBufferFormatControl::bufferFormatChanged), static_cast<MyQCameraCaptureBufferFormatControl*>(ptr), static_cast<void (MyQCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&MyQCameraCaptureBufferFormatControl::Signal_BufferFormatChanged));
}

void QCameraCaptureBufferFormatControl_BufferFormatChanged(void* ptr, long long format)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->bufferFormatChanged(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraCaptureBufferFormatControl_SetBufferFormat(void* ptr, long long format)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->setBufferFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraCaptureBufferFormatControl_DestroyQCameraCaptureBufferFormatControl(void* ptr)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->~QCameraCaptureBufferFormatControl();
}

void QCameraCaptureBufferFormatControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraCaptureBufferFormatControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->QCameraCaptureBufferFormatControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraCaptureBufferFormatControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraCaptureBufferFormatControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->QCameraCaptureBufferFormatControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraCaptureBufferFormatControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraCaptureBufferFormatControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->QCameraCaptureBufferFormatControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraCaptureBufferFormatControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraCaptureBufferFormatControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->QCameraCaptureBufferFormatControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraCaptureBufferFormatControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraCaptureBufferFormatControl*>(ptr), "deleteLater");
}

void QCameraCaptureBufferFormatControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->QCameraCaptureBufferFormatControl::deleteLater();
}

void QCameraCaptureBufferFormatControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraCaptureBufferFormatControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->QCameraCaptureBufferFormatControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraCaptureBufferFormatControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraCaptureBufferFormatControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraCaptureBufferFormatControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraCaptureBufferFormatControl*>(ptr)->QCameraCaptureBufferFormatControl::event(static_cast<QEvent*>(e));
}

char QCameraCaptureBufferFormatControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraCaptureBufferFormatControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraCaptureBufferFormatControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraCaptureBufferFormatControl*>(ptr)->QCameraCaptureBufferFormatControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraCaptureBufferFormatControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraCaptureBufferFormatControl*>(ptr)->metaObject());
}

void* QCameraCaptureBufferFormatControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraCaptureBufferFormatControl*>(ptr)->QCameraCaptureBufferFormatControl::metaObject());
}

class MyQCameraCaptureDestinationControl: public QCameraCaptureDestinationControl
{
public:
	MyQCameraCaptureDestinationControl(QObject *parent) : QCameraCaptureDestinationControl(parent) {};
	QCameraImageCapture::CaptureDestinations captureDestination() const { return static_cast<QCameraImageCapture::CaptureDestination>(callbackQCameraCaptureDestinationControl_CaptureDestination(const_cast<MyQCameraCaptureDestinationControl*>(this))); };
	void Signal_CaptureDestinationChanged(QCameraImageCapture::CaptureDestinations destination) { callbackQCameraCaptureDestinationControl_CaptureDestinationChanged(this, destination); };
	bool isCaptureDestinationSupported(QCameraImageCapture::CaptureDestinations destination) const { return callbackQCameraCaptureDestinationControl_IsCaptureDestinationSupported(const_cast<MyQCameraCaptureDestinationControl*>(this), destination) != 0; };
	void setCaptureDestination(QCameraImageCapture::CaptureDestinations destination) { callbackQCameraCaptureDestinationControl_SetCaptureDestination(this, destination); };
	void timerEvent(QTimerEvent * event) { callbackQCameraCaptureDestinationControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraCaptureDestinationControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraCaptureDestinationControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraCaptureDestinationControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraCaptureDestinationControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraCaptureDestinationControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraCaptureDestinationControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraCaptureDestinationControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraCaptureDestinationControl_MetaObject(const_cast<MyQCameraCaptureDestinationControl*>(this))); };
};

void* QCameraCaptureDestinationControl_NewQCameraCaptureDestinationControl(void* parent)
{
	return new MyQCameraCaptureDestinationControl(static_cast<QObject*>(parent));
}

long long QCameraCaptureDestinationControl_CaptureDestination(void* ptr)
{
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->captureDestination();
}

void QCameraCaptureDestinationControl_ConnectCaptureDestinationChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraCaptureDestinationControl*>(ptr), static_cast<void (QCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraCaptureDestinationControl::captureDestinationChanged), static_cast<MyQCameraCaptureDestinationControl*>(ptr), static_cast<void (MyQCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraCaptureDestinationControl::Signal_CaptureDestinationChanged));
}

void QCameraCaptureDestinationControl_DisconnectCaptureDestinationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraCaptureDestinationControl*>(ptr), static_cast<void (QCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraCaptureDestinationControl::captureDestinationChanged), static_cast<MyQCameraCaptureDestinationControl*>(ptr), static_cast<void (MyQCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraCaptureDestinationControl::Signal_CaptureDestinationChanged));
}

void QCameraCaptureDestinationControl_CaptureDestinationChanged(void* ptr, long long destination)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->captureDestinationChanged(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

char QCameraCaptureDestinationControl_IsCaptureDestinationSupported(void* ptr, long long destination)
{
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->isCaptureDestinationSupported(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraCaptureDestinationControl_SetCaptureDestination(void* ptr, long long destination)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->setCaptureDestination(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraCaptureDestinationControl_DestroyQCameraCaptureDestinationControl(void* ptr)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->~QCameraCaptureDestinationControl();
}

void QCameraCaptureDestinationControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraCaptureDestinationControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->QCameraCaptureDestinationControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraCaptureDestinationControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraCaptureDestinationControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->QCameraCaptureDestinationControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraCaptureDestinationControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraCaptureDestinationControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->QCameraCaptureDestinationControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraCaptureDestinationControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraCaptureDestinationControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->QCameraCaptureDestinationControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraCaptureDestinationControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraCaptureDestinationControl*>(ptr), "deleteLater");
}

void QCameraCaptureDestinationControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->QCameraCaptureDestinationControl::deleteLater();
}

void QCameraCaptureDestinationControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraCaptureDestinationControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraCaptureDestinationControl*>(ptr)->QCameraCaptureDestinationControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraCaptureDestinationControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraCaptureDestinationControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->QCameraCaptureDestinationControl::event(static_cast<QEvent*>(e));
}

char QCameraCaptureDestinationControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraCaptureDestinationControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->QCameraCaptureDestinationControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraCaptureDestinationControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraCaptureDestinationControl*>(ptr)->metaObject());
}

void* QCameraCaptureDestinationControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraCaptureDestinationControl*>(ptr)->QCameraCaptureDestinationControl::metaObject());
}

class MyQCameraControl: public QCameraControl
{
public:
	MyQCameraControl(QObject *parent) : QCameraControl(parent) {};
	bool canChangeProperty(QCameraControl::PropertyChangeType changeType, QCamera::Status status) const { return callbackQCameraControl_CanChangeProperty(const_cast<MyQCameraControl*>(this), changeType, status) != 0; };
	QCamera::CaptureModes captureMode() const { return static_cast<QCamera::CaptureMode>(callbackQCameraControl_CaptureMode(const_cast<MyQCameraControl*>(this))); };
	void Signal_CaptureModeChanged(QCamera::CaptureModes mode) { callbackQCameraControl_CaptureModeChanged(this, mode); };
	void Signal_Error(int error, const QString & errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtMultimedia_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQCameraControl_Error(this, error, errorStringPacked); };
	bool isCaptureModeSupported(QCamera::CaptureModes mode) const { return callbackQCameraControl_IsCaptureModeSupported(const_cast<MyQCameraControl*>(this), mode) != 0; };
	void setCaptureMode(QCamera::CaptureModes mode) { callbackQCameraControl_SetCaptureMode(this, mode); };
	void setState(QCamera::State state) { callbackQCameraControl_SetState(this, state); };
	QCamera::State state() const { return static_cast<QCamera::State>(callbackQCameraControl_State(const_cast<MyQCameraControl*>(this))); };
	void Signal_StateChanged(QCamera::State state) { callbackQCameraControl_StateChanged(this, state); };
	QCamera::Status status() const { return static_cast<QCamera::Status>(callbackQCameraControl_Status(const_cast<MyQCameraControl*>(this))); };
	void Signal_StatusChanged(QCamera::Status status) { callbackQCameraControl_StatusChanged(this, status); };
	void timerEvent(QTimerEvent * event) { callbackQCameraControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraControl_MetaObject(const_cast<MyQCameraControl*>(this))); };
};

void* QCameraControl_NewQCameraControl(void* parent)
{
	return new MyQCameraControl(static_cast<QObject*>(parent));
}

char QCameraControl_CanChangeProperty(void* ptr, long long changeType, long long status)
{
	return static_cast<QCameraControl*>(ptr)->canChangeProperty(static_cast<QCameraControl::PropertyChangeType>(changeType), static_cast<QCamera::Status>(status));
}

long long QCameraControl_CaptureMode(void* ptr)
{
	return static_cast<QCameraControl*>(ptr)->captureMode();
}

void QCameraControl_ConnectCaptureModeChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::CaptureModes)>(&QCameraControl::captureModeChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::CaptureModes)>(&MyQCameraControl::Signal_CaptureModeChanged));
}

void QCameraControl_DisconnectCaptureModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::CaptureModes)>(&QCameraControl::captureModeChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::CaptureModes)>(&MyQCameraControl::Signal_CaptureModeChanged));
}

void QCameraControl_CaptureModeChanged(void* ptr, long long mode)
{
	static_cast<QCameraControl*>(ptr)->captureModeChanged(static_cast<QCamera::CaptureMode>(mode));
}

void QCameraControl_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(int, const QString &)>(&QCameraControl::error), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(int, const QString &)>(&MyQCameraControl::Signal_Error));
}

void QCameraControl_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(int, const QString &)>(&QCameraControl::error), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(int, const QString &)>(&MyQCameraControl::Signal_Error));
}

void QCameraControl_Error(void* ptr, int error, char* errorString)
{
	static_cast<QCameraControl*>(ptr)->error(error, QString(errorString));
}

char QCameraControl_IsCaptureModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraControl*>(ptr)->isCaptureModeSupported(static_cast<QCamera::CaptureMode>(mode));
}

void QCameraControl_SetCaptureMode(void* ptr, long long mode)
{
	static_cast<QCameraControl*>(ptr)->setCaptureMode(static_cast<QCamera::CaptureMode>(mode));
}

void QCameraControl_SetState(void* ptr, long long state)
{
	static_cast<QCameraControl*>(ptr)->setState(static_cast<QCamera::State>(state));
}

long long QCameraControl_State(void* ptr)
{
	return static_cast<QCameraControl*>(ptr)->state();
}

void QCameraControl_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::State)>(&QCameraControl::stateChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::State)>(&MyQCameraControl::Signal_StateChanged));
}

void QCameraControl_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::State)>(&QCameraControl::stateChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::State)>(&MyQCameraControl::Signal_StateChanged));
}

void QCameraControl_StateChanged(void* ptr, long long state)
{
	static_cast<QCameraControl*>(ptr)->stateChanged(static_cast<QCamera::State>(state));
}

long long QCameraControl_Status(void* ptr)
{
	return static_cast<QCameraControl*>(ptr)->status();
}

void QCameraControl_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::Status)>(&QCameraControl::statusChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::Status)>(&MyQCameraControl::Signal_StatusChanged));
}

void QCameraControl_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::Status)>(&QCameraControl::statusChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::Status)>(&MyQCameraControl::Signal_StatusChanged));
}

void QCameraControl_StatusChanged(void* ptr, long long status)
{
	static_cast<QCameraControl*>(ptr)->statusChanged(static_cast<QCamera::Status>(status));
}

void QCameraControl_DestroyQCameraControl(void* ptr)
{
	static_cast<QCameraControl*>(ptr)->~QCameraControl();
}

void QCameraControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraControl*>(ptr)->QCameraControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraControl*>(ptr)->QCameraControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraControl*>(ptr)->QCameraControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraControl*>(ptr)->QCameraControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraControl*>(ptr), "deleteLater");
}

void QCameraControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraControl*>(ptr)->QCameraControl::deleteLater();
}

void QCameraControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraControl*>(ptr)->QCameraControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraControl*>(ptr)->QCameraControl::event(static_cast<QEvent*>(e));
}

char QCameraControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraControl*>(ptr)->QCameraControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraControl*>(ptr)->metaObject());
}

void* QCameraControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraControl*>(ptr)->QCameraControl::metaObject());
}

class MyQCameraExposure: public QCameraExposure
{
public:
	void setAutoAperture() { callbackQCameraExposure_SetAutoAperture(this); };
	void setAutoIsoSensitivity() { callbackQCameraExposure_SetAutoIsoSensitivity(this); };
	void setExposureCompensation(qreal ev) { callbackQCameraExposure_SetExposureCompensation(this, ev); };
	void setExposureMode(QCameraExposure::ExposureMode mode) { callbackQCameraExposure_SetExposureMode(this, mode); };
	void setFlashMode(QCameraExposure::FlashModes mode) { callbackQCameraExposure_SetFlashMode(this, mode); };
	void setManualAperture(qreal aperture) { callbackQCameraExposure_SetManualAperture(this, aperture); };
	void setManualIsoSensitivity(int iso) { callbackQCameraExposure_SetManualIsoSensitivity(this, iso); };
	void setMeteringMode(QCameraExposure::MeteringMode mode) { callbackQCameraExposure_SetMeteringMode(this, mode); };
	void Signal_ApertureChanged(qreal value) { callbackQCameraExposure_ApertureChanged(this, value); };
	void Signal_ApertureRangeChanged() { callbackQCameraExposure_ApertureRangeChanged(this); };
	void Signal_ExposureCompensationChanged(qreal value) { callbackQCameraExposure_ExposureCompensationChanged(this, value); };
	void Signal_FlashReady(bool ready) { callbackQCameraExposure_FlashReady(this, ready); };
	void Signal_IsoSensitivityChanged(int value) { callbackQCameraExposure_IsoSensitivityChanged(this, value); };
	void setAutoShutterSpeed() { callbackQCameraExposure_SetAutoShutterSpeed(this); };
	void setManualShutterSpeed(qreal seconds) { callbackQCameraExposure_SetManualShutterSpeed(this, seconds); };
	void Signal_ShutterSpeedChanged(qreal speed) { callbackQCameraExposure_ShutterSpeedChanged(this, speed); };
	void Signal_ShutterSpeedRangeChanged() { callbackQCameraExposure_ShutterSpeedRangeChanged(this); };
	void timerEvent(QTimerEvent * event) { callbackQCameraExposure_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraExposure_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraExposure_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraExposure_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraExposure_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraExposure_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraExposure_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraExposure_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraExposure_MetaObject(const_cast<MyQCameraExposure*>(this))); };
};

double QCameraExposure_Aperture(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->aperture();
}

double QCameraExposure_ExposureCompensation(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->exposureCompensation();
}

long long QCameraExposure_ExposureMode(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->exposureMode();
}

long long QCameraExposure_FlashMode(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->flashMode();
}

int QCameraExposure_IsoSensitivity(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->isoSensitivity();
}

long long QCameraExposure_MeteringMode(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->meteringMode();
}

void QCameraExposure_SetAutoAperture(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoAperture");
}

void QCameraExposure_SetAutoIsoSensitivity(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoIsoSensitivity");
}

void QCameraExposure_SetExposureCompensation(void* ptr, double ev)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setExposureCompensation", Q_ARG(qreal, ev));
}

void QCameraExposure_SetExposureMode(void* ptr, long long mode)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setExposureMode", Q_ARG(QCameraExposure::ExposureMode, static_cast<QCameraExposure::ExposureMode>(mode)));
}

void QCameraExposure_SetFlashMode(void* ptr, long long mode)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setFlashMode", Q_ARG(QCameraExposure::FlashMode, static_cast<QCameraExposure::FlashMode>(mode)));
}

void QCameraExposure_SetManualAperture(void* ptr, double aperture)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualAperture", Q_ARG(qreal, aperture));
}

void QCameraExposure_SetManualIsoSensitivity(void* ptr, int iso)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualIsoSensitivity", Q_ARG(int, iso));
}

void QCameraExposure_SetMeteringMode(void* ptr, long long mode)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setMeteringMode", Q_ARG(QCameraExposure::MeteringMode, static_cast<QCameraExposure::MeteringMode>(mode)));
}

void QCameraExposure_SetSpotMeteringPoint(void* ptr, void* point)
{
	static_cast<QCameraExposure*>(ptr)->setSpotMeteringPoint(*static_cast<QPointF*>(point));
}

void* QCameraExposure_SpotMeteringPoint(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QCameraExposure*>(ptr)->spotMeteringPoint(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QCameraExposure_ConnectApertureChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(qreal)>(&QCameraExposure::apertureChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(qreal)>(&MyQCameraExposure::Signal_ApertureChanged));
}

void QCameraExposure_DisconnectApertureChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(qreal)>(&QCameraExposure::apertureChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(qreal)>(&MyQCameraExposure::Signal_ApertureChanged));
}

void QCameraExposure_ApertureChanged(void* ptr, double value)
{
	static_cast<QCameraExposure*>(ptr)->apertureChanged(value);
}

void QCameraExposure_ConnectApertureRangeChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::apertureRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ApertureRangeChanged));
}

void QCameraExposure_DisconnectApertureRangeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::apertureRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ApertureRangeChanged));
}

void QCameraExposure_ApertureRangeChanged(void* ptr)
{
	static_cast<QCameraExposure*>(ptr)->apertureRangeChanged();
}

void QCameraExposure_ConnectExposureCompensationChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(qreal)>(&QCameraExposure::exposureCompensationChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(qreal)>(&MyQCameraExposure::Signal_ExposureCompensationChanged));
}

void QCameraExposure_DisconnectExposureCompensationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(qreal)>(&QCameraExposure::exposureCompensationChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(qreal)>(&MyQCameraExposure::Signal_ExposureCompensationChanged));
}

void QCameraExposure_ExposureCompensationChanged(void* ptr, double value)
{
	static_cast<QCameraExposure*>(ptr)->exposureCompensationChanged(value);
}

void QCameraExposure_ConnectFlashReady(void* ptr)
{
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(bool)>(&QCameraExposure::flashReady), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(bool)>(&MyQCameraExposure::Signal_FlashReady));
}

void QCameraExposure_DisconnectFlashReady(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(bool)>(&QCameraExposure::flashReady), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(bool)>(&MyQCameraExposure::Signal_FlashReady));
}

void QCameraExposure_FlashReady(void* ptr, char ready)
{
	static_cast<QCameraExposure*>(ptr)->flashReady(ready != 0);
}

char QCameraExposure_IsAvailable(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->isAvailable();
}

char QCameraExposure_IsExposureModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraExposure*>(ptr)->isExposureModeSupported(static_cast<QCameraExposure::ExposureMode>(mode));
}

char QCameraExposure_IsFlashModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraExposure*>(ptr)->isFlashModeSupported(static_cast<QCameraExposure::FlashMode>(mode));
}

char QCameraExposure_IsFlashReady(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->isFlashReady();
}

char QCameraExposure_IsMeteringModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraExposure*>(ptr)->isMeteringModeSupported(static_cast<QCameraExposure::MeteringMode>(mode));
}

void QCameraExposure_ConnectIsoSensitivityChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(int)>(&QCameraExposure::isoSensitivityChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(int)>(&MyQCameraExposure::Signal_IsoSensitivityChanged));
}

void QCameraExposure_DisconnectIsoSensitivityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(int)>(&QCameraExposure::isoSensitivityChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(int)>(&MyQCameraExposure::Signal_IsoSensitivityChanged));
}

void QCameraExposure_IsoSensitivityChanged(void* ptr, int value)
{
	static_cast<QCameraExposure*>(ptr)->isoSensitivityChanged(value);
}

double QCameraExposure_RequestedAperture(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->requestedAperture();
}

int QCameraExposure_RequestedIsoSensitivity(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->requestedIsoSensitivity();
}

double QCameraExposure_RequestedShutterSpeed(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->requestedShutterSpeed();
}

void QCameraExposure_SetAutoShutterSpeed(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoShutterSpeed");
}

void QCameraExposure_SetManualShutterSpeed(void* ptr, double seconds)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualShutterSpeed", Q_ARG(qreal, seconds));
}

double QCameraExposure_ShutterSpeed(void* ptr)
{
	return static_cast<QCameraExposure*>(ptr)->shutterSpeed();
}

void QCameraExposure_ConnectShutterSpeedChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(qreal)>(&QCameraExposure::shutterSpeedChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(qreal)>(&MyQCameraExposure::Signal_ShutterSpeedChanged));
}

void QCameraExposure_DisconnectShutterSpeedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(qreal)>(&QCameraExposure::shutterSpeedChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(qreal)>(&MyQCameraExposure::Signal_ShutterSpeedChanged));
}

void QCameraExposure_ShutterSpeedChanged(void* ptr, double speed)
{
	static_cast<QCameraExposure*>(ptr)->shutterSpeedChanged(speed);
}

void QCameraExposure_ConnectShutterSpeedRangeChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::shutterSpeedRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ShutterSpeedRangeChanged));
}

void QCameraExposure_DisconnectShutterSpeedRangeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::shutterSpeedRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ShutterSpeedRangeChanged));
}

void QCameraExposure_ShutterSpeedRangeChanged(void* ptr)
{
	static_cast<QCameraExposure*>(ptr)->shutterSpeedRangeChanged();
}

void QCameraExposure_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraExposure*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraExposure_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraExposure*>(ptr)->QCameraExposure::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraExposure_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraExposure*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraExposure_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraExposure*>(ptr)->QCameraExposure::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraExposure_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraExposure*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraExposure_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraExposure*>(ptr)->QCameraExposure::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraExposure_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraExposure*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraExposure_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraExposure*>(ptr)->QCameraExposure::customEvent(static_cast<QEvent*>(event));
}

void QCameraExposure_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "deleteLater");
}

void QCameraExposure_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraExposure*>(ptr)->QCameraExposure::deleteLater();
}

void QCameraExposure_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraExposure*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraExposure_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraExposure*>(ptr)->QCameraExposure::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraExposure_Event(void* ptr, void* e)
{
	return static_cast<QCameraExposure*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraExposure_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraExposure*>(ptr)->QCameraExposure::event(static_cast<QEvent*>(e));
}

char QCameraExposure_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraExposure*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraExposure_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraExposure*>(ptr)->QCameraExposure::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraExposure_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraExposure*>(ptr)->metaObject());
}

void* QCameraExposure_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraExposure*>(ptr)->QCameraExposure::metaObject());
}

class MyQCameraExposureControl: public QCameraExposureControl
{
public:
	QVariant actualValue(QCameraExposureControl::ExposureParameter parameter) const { return *static_cast<QVariant*>(callbackQCameraExposureControl_ActualValue(const_cast<MyQCameraExposureControl*>(this), parameter)); };
	void Signal_ActualValueChanged(int parameter) { callbackQCameraExposureControl_ActualValueChanged(this, parameter); };
	bool isParameterSupported(QCameraExposureControl::ExposureParameter parameter) const { return callbackQCameraExposureControl_IsParameterSupported(const_cast<MyQCameraExposureControl*>(this), parameter) != 0; };
	void Signal_ParameterRangeChanged(int parameter) { callbackQCameraExposureControl_ParameterRangeChanged(this, parameter); };
	QVariant requestedValue(QCameraExposureControl::ExposureParameter parameter) const { return *static_cast<QVariant*>(callbackQCameraExposureControl_RequestedValue(const_cast<MyQCameraExposureControl*>(this), parameter)); };
	void Signal_RequestedValueChanged(int parameter) { callbackQCameraExposureControl_RequestedValueChanged(this, parameter); };
	bool setValue(QCameraExposureControl::ExposureParameter parameter, const QVariant & value) { return callbackQCameraExposureControl_SetValue(this, parameter, const_cast<QVariant*>(&value)) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQCameraExposureControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraExposureControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraExposureControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraExposureControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraExposureControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraExposureControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraExposureControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraExposureControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraExposureControl_MetaObject(const_cast<MyQCameraExposureControl*>(this))); };
};

void* QCameraExposureControl_ActualValue(void* ptr, long long parameter)
{
	return new QVariant(static_cast<QCameraExposureControl*>(ptr)->actualValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter)));
}

void QCameraExposureControl_ConnectActualValueChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::actualValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ActualValueChanged));
}

void QCameraExposureControl_DisconnectActualValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::actualValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ActualValueChanged));
}

void QCameraExposureControl_ActualValueChanged(void* ptr, int parameter)
{
	static_cast<QCameraExposureControl*>(ptr)->actualValueChanged(parameter);
}

char QCameraExposureControl_IsParameterSupported(void* ptr, long long parameter)
{
	return static_cast<QCameraExposureControl*>(ptr)->isParameterSupported(static_cast<QCameraExposureControl::ExposureParameter>(parameter));
}

void QCameraExposureControl_ConnectParameterRangeChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::parameterRangeChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ParameterRangeChanged));
}

void QCameraExposureControl_DisconnectParameterRangeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::parameterRangeChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ParameterRangeChanged));
}

void QCameraExposureControl_ParameterRangeChanged(void* ptr, int parameter)
{
	static_cast<QCameraExposureControl*>(ptr)->parameterRangeChanged(parameter);
}

void* QCameraExposureControl_RequestedValue(void* ptr, long long parameter)
{
	return new QVariant(static_cast<QCameraExposureControl*>(ptr)->requestedValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter)));
}

void QCameraExposureControl_ConnectRequestedValueChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::requestedValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_RequestedValueChanged));
}

void QCameraExposureControl_DisconnectRequestedValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::requestedValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_RequestedValueChanged));
}

void QCameraExposureControl_RequestedValueChanged(void* ptr, int parameter)
{
	static_cast<QCameraExposureControl*>(ptr)->requestedValueChanged(parameter);
}

char QCameraExposureControl_SetValue(void* ptr, long long parameter, void* value)
{
	return static_cast<QCameraExposureControl*>(ptr)->setValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter), *static_cast<QVariant*>(value));
}

void QCameraExposureControl_DestroyQCameraExposureControl(void* ptr)
{
	static_cast<QCameraExposureControl*>(ptr)->~QCameraExposureControl();
}

void QCameraExposureControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraExposureControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraExposureControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraExposureControl*>(ptr)->QCameraExposureControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraExposureControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraExposureControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraExposureControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraExposureControl*>(ptr)->QCameraExposureControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraExposureControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraExposureControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraExposureControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraExposureControl*>(ptr)->QCameraExposureControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraExposureControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraExposureControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraExposureControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraExposureControl*>(ptr)->QCameraExposureControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraExposureControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraExposureControl*>(ptr), "deleteLater");
}

void QCameraExposureControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraExposureControl*>(ptr)->QCameraExposureControl::deleteLater();
}

void QCameraExposureControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraExposureControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraExposureControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraExposureControl*>(ptr)->QCameraExposureControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraExposureControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraExposureControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraExposureControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraExposureControl*>(ptr)->QCameraExposureControl::event(static_cast<QEvent*>(e));
}

char QCameraExposureControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraExposureControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraExposureControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraExposureControl*>(ptr)->QCameraExposureControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraExposureControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraExposureControl*>(ptr)->metaObject());
}

void* QCameraExposureControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraExposureControl*>(ptr)->QCameraExposureControl::metaObject());
}

class MyQCameraFeedbackControl: public QCameraFeedbackControl
{
public:
	MyQCameraFeedbackControl(QObject *parent) : QCameraFeedbackControl(parent) {};
	bool isEventFeedbackEnabled(QCameraFeedbackControl::EventType event) const { return callbackQCameraFeedbackControl_IsEventFeedbackEnabled(const_cast<MyQCameraFeedbackControl*>(this), event) != 0; };
	bool isEventFeedbackLocked(QCameraFeedbackControl::EventType event) const { return callbackQCameraFeedbackControl_IsEventFeedbackLocked(const_cast<MyQCameraFeedbackControl*>(this), event) != 0; };
	void resetEventFeedback(QCameraFeedbackControl::EventType event) { callbackQCameraFeedbackControl_ResetEventFeedback(this, event); };
	bool setEventFeedbackEnabled(QCameraFeedbackControl::EventType event, bool enabled) { return callbackQCameraFeedbackControl_SetEventFeedbackEnabled(this, event, enabled) != 0; };
	bool setEventFeedbackSound(QCameraFeedbackControl::EventType event, const QString & filePath) { QByteArray t7df503 = filePath.toUtf8(); QtMultimedia_PackedString filePathPacked = { const_cast<char*>(t7df503.prepend("WHITESPACE").constData()+10), t7df503.size()-10 };return callbackQCameraFeedbackControl_SetEventFeedbackSound(this, event, filePathPacked) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQCameraFeedbackControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraFeedbackControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraFeedbackControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraFeedbackControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraFeedbackControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraFeedbackControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraFeedbackControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraFeedbackControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraFeedbackControl_MetaObject(const_cast<MyQCameraFeedbackControl*>(this))); };
};

void* QCameraFeedbackControl_NewQCameraFeedbackControl(void* parent)
{
	return new MyQCameraFeedbackControl(static_cast<QObject*>(parent));
}

char QCameraFeedbackControl_IsEventFeedbackEnabled(void* ptr, long long event)
{
	return static_cast<QCameraFeedbackControl*>(ptr)->isEventFeedbackEnabled(static_cast<QCameraFeedbackControl::EventType>(event));
}

char QCameraFeedbackControl_IsEventFeedbackLocked(void* ptr, long long event)
{
	return static_cast<QCameraFeedbackControl*>(ptr)->isEventFeedbackLocked(static_cast<QCameraFeedbackControl::EventType>(event));
}

void QCameraFeedbackControl_ResetEventFeedback(void* ptr, long long event)
{
	static_cast<QCameraFeedbackControl*>(ptr)->resetEventFeedback(static_cast<QCameraFeedbackControl::EventType>(event));
}

char QCameraFeedbackControl_SetEventFeedbackEnabled(void* ptr, long long event, char enabled)
{
	return static_cast<QCameraFeedbackControl*>(ptr)->setEventFeedbackEnabled(static_cast<QCameraFeedbackControl::EventType>(event), enabled != 0);
}

char QCameraFeedbackControl_SetEventFeedbackSound(void* ptr, long long event, char* filePath)
{
	return static_cast<QCameraFeedbackControl*>(ptr)->setEventFeedbackSound(static_cast<QCameraFeedbackControl::EventType>(event), QString(filePath));
}

void QCameraFeedbackControl_DestroyQCameraFeedbackControl(void* ptr)
{
	static_cast<QCameraFeedbackControl*>(ptr)->~QCameraFeedbackControl();
}

void QCameraFeedbackControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraFeedbackControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraFeedbackControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFeedbackControl*>(ptr)->QCameraFeedbackControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraFeedbackControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraFeedbackControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraFeedbackControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFeedbackControl*>(ptr)->QCameraFeedbackControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraFeedbackControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraFeedbackControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFeedbackControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraFeedbackControl*>(ptr)->QCameraFeedbackControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFeedbackControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraFeedbackControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraFeedbackControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFeedbackControl*>(ptr)->QCameraFeedbackControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraFeedbackControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraFeedbackControl*>(ptr), "deleteLater");
}

void QCameraFeedbackControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraFeedbackControl*>(ptr)->QCameraFeedbackControl::deleteLater();
}

void QCameraFeedbackControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraFeedbackControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFeedbackControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraFeedbackControl*>(ptr)->QCameraFeedbackControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraFeedbackControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraFeedbackControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraFeedbackControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraFeedbackControl*>(ptr)->QCameraFeedbackControl::event(static_cast<QEvent*>(e));
}

char QCameraFeedbackControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraFeedbackControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraFeedbackControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraFeedbackControl*>(ptr)->QCameraFeedbackControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraFeedbackControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraFeedbackControl*>(ptr)->metaObject());
}

void* QCameraFeedbackControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraFeedbackControl*>(ptr)->QCameraFeedbackControl::metaObject());
}

class MyQCameraFlashControl: public QCameraFlashControl
{
public:
	MyQCameraFlashControl(QObject *parent) : QCameraFlashControl(parent) {};
	QCameraExposure::FlashModes flashMode() const { return static_cast<QCameraExposure::FlashMode>(callbackQCameraFlashControl_FlashMode(const_cast<MyQCameraFlashControl*>(this))); };
	void Signal_FlashReady(bool ready) { callbackQCameraFlashControl_FlashReady(this, ready); };
	bool isFlashModeSupported(QCameraExposure::FlashModes mode) const { return callbackQCameraFlashControl_IsFlashModeSupported(const_cast<MyQCameraFlashControl*>(this), mode) != 0; };
	bool isFlashReady() const { return callbackQCameraFlashControl_IsFlashReady(const_cast<MyQCameraFlashControl*>(this)) != 0; };
	void setFlashMode(QCameraExposure::FlashModes mode) { callbackQCameraFlashControl_SetFlashMode(this, mode); };
	void timerEvent(QTimerEvent * event) { callbackQCameraFlashControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraFlashControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraFlashControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraFlashControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraFlashControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraFlashControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraFlashControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraFlashControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraFlashControl_MetaObject(const_cast<MyQCameraFlashControl*>(this))); };
};

void* QCameraFlashControl_NewQCameraFlashControl(void* parent)
{
	return new MyQCameraFlashControl(static_cast<QObject*>(parent));
}

long long QCameraFlashControl_FlashMode(void* ptr)
{
	return static_cast<QCameraFlashControl*>(ptr)->flashMode();
}

void QCameraFlashControl_ConnectFlashReady(void* ptr)
{
	QObject::connect(static_cast<QCameraFlashControl*>(ptr), static_cast<void (QCameraFlashControl::*)(bool)>(&QCameraFlashControl::flashReady), static_cast<MyQCameraFlashControl*>(ptr), static_cast<void (MyQCameraFlashControl::*)(bool)>(&MyQCameraFlashControl::Signal_FlashReady));
}

void QCameraFlashControl_DisconnectFlashReady(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFlashControl*>(ptr), static_cast<void (QCameraFlashControl::*)(bool)>(&QCameraFlashControl::flashReady), static_cast<MyQCameraFlashControl*>(ptr), static_cast<void (MyQCameraFlashControl::*)(bool)>(&MyQCameraFlashControl::Signal_FlashReady));
}

void QCameraFlashControl_FlashReady(void* ptr, char ready)
{
	static_cast<QCameraFlashControl*>(ptr)->flashReady(ready != 0);
}

char QCameraFlashControl_IsFlashModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraFlashControl*>(ptr)->isFlashModeSupported(static_cast<QCameraExposure::FlashMode>(mode));
}

char QCameraFlashControl_IsFlashReady(void* ptr)
{
	return static_cast<QCameraFlashControl*>(ptr)->isFlashReady();
}

void QCameraFlashControl_SetFlashMode(void* ptr, long long mode)
{
	static_cast<QCameraFlashControl*>(ptr)->setFlashMode(static_cast<QCameraExposure::FlashMode>(mode));
}

void QCameraFlashControl_DestroyQCameraFlashControl(void* ptr)
{
	static_cast<QCameraFlashControl*>(ptr)->~QCameraFlashControl();
}

void QCameraFlashControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraFlashControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraFlashControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFlashControl*>(ptr)->QCameraFlashControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraFlashControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraFlashControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraFlashControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFlashControl*>(ptr)->QCameraFlashControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraFlashControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraFlashControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFlashControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraFlashControl*>(ptr)->QCameraFlashControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFlashControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraFlashControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraFlashControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFlashControl*>(ptr)->QCameraFlashControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraFlashControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraFlashControl*>(ptr), "deleteLater");
}

void QCameraFlashControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraFlashControl*>(ptr)->QCameraFlashControl::deleteLater();
}

void QCameraFlashControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraFlashControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFlashControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraFlashControl*>(ptr)->QCameraFlashControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraFlashControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraFlashControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraFlashControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraFlashControl*>(ptr)->QCameraFlashControl::event(static_cast<QEvent*>(e));
}

char QCameraFlashControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraFlashControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraFlashControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraFlashControl*>(ptr)->QCameraFlashControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraFlashControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraFlashControl*>(ptr)->metaObject());
}

void* QCameraFlashControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraFlashControl*>(ptr)->QCameraFlashControl::metaObject());
}

class MyQCameraFocus: public QCameraFocus
{
public:
	void Signal_DigitalZoomChanged(qreal value) { callbackQCameraFocus_DigitalZoomChanged(this, value); };
	void Signal_FocusZonesChanged() { callbackQCameraFocus_FocusZonesChanged(this); };
	void Signal_MaximumDigitalZoomChanged(qreal zoom) { callbackQCameraFocus_MaximumDigitalZoomChanged(this, zoom); };
	void Signal_MaximumOpticalZoomChanged(qreal zoom) { callbackQCameraFocus_MaximumOpticalZoomChanged(this, zoom); };
	void Signal_OpticalZoomChanged(qreal value) { callbackQCameraFocus_OpticalZoomChanged(this, value); };
	void timerEvent(QTimerEvent * event) { callbackQCameraFocus_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraFocus_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraFocus_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraFocus_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraFocus_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraFocus_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraFocus_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraFocus_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraFocus_MetaObject(const_cast<MyQCameraFocus*>(this))); };
};

void* QCameraFocus_CustomFocusPoint(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QCameraFocus*>(ptr)->customFocusPoint(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

double QCameraFocus_DigitalZoom(void* ptr)
{
	return static_cast<QCameraFocus*>(ptr)->digitalZoom();
}

long long QCameraFocus_FocusMode(void* ptr)
{
	return static_cast<QCameraFocus*>(ptr)->focusMode();
}

long long QCameraFocus_FocusPointMode(void* ptr)
{
	return static_cast<QCameraFocus*>(ptr)->focusPointMode();
}

double QCameraFocus_OpticalZoom(void* ptr)
{
	return static_cast<QCameraFocus*>(ptr)->opticalZoom();
}

void QCameraFocus_SetCustomFocusPoint(void* ptr, void* point)
{
	static_cast<QCameraFocus*>(ptr)->setCustomFocusPoint(*static_cast<QPointF*>(point));
}

void QCameraFocus_SetFocusMode(void* ptr, long long mode)
{
	static_cast<QCameraFocus*>(ptr)->setFocusMode(static_cast<QCameraFocus::FocusMode>(mode));
}

void QCameraFocus_SetFocusPointMode(void* ptr, long long mode)
{
	static_cast<QCameraFocus*>(ptr)->setFocusPointMode(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocus_ConnectDigitalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)(qreal)>(&QCameraFocus::digitalZoomChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)(qreal)>(&MyQCameraFocus::Signal_DigitalZoomChanged));
}

void QCameraFocus_DisconnectDigitalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)(qreal)>(&QCameraFocus::digitalZoomChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)(qreal)>(&MyQCameraFocus::Signal_DigitalZoomChanged));
}

void QCameraFocus_DigitalZoomChanged(void* ptr, double value)
{
	static_cast<QCameraFocus*>(ptr)->digitalZoomChanged(value);
}

void QCameraFocus_ConnectFocusZonesChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)()>(&QCameraFocus::focusZonesChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)()>(&MyQCameraFocus::Signal_FocusZonesChanged));
}

void QCameraFocus_DisconnectFocusZonesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)()>(&QCameraFocus::focusZonesChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)()>(&MyQCameraFocus::Signal_FocusZonesChanged));
}

void QCameraFocus_FocusZonesChanged(void* ptr)
{
	static_cast<QCameraFocus*>(ptr)->focusZonesChanged();
}

char QCameraFocus_IsAvailable(void* ptr)
{
	return static_cast<QCameraFocus*>(ptr)->isAvailable();
}

char QCameraFocus_IsFocusModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraFocus*>(ptr)->isFocusModeSupported(static_cast<QCameraFocus::FocusMode>(mode));
}

char QCameraFocus_IsFocusPointModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraFocus*>(ptr)->isFocusPointModeSupported(static_cast<QCameraFocus::FocusPointMode>(mode));
}

double QCameraFocus_MaximumDigitalZoom(void* ptr)
{
	return static_cast<QCameraFocus*>(ptr)->maximumDigitalZoom();
}

void QCameraFocus_ConnectMaximumDigitalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)(qreal)>(&QCameraFocus::maximumDigitalZoomChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)(qreal)>(&MyQCameraFocus::Signal_MaximumDigitalZoomChanged));
}

void QCameraFocus_DisconnectMaximumDigitalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)(qreal)>(&QCameraFocus::maximumDigitalZoomChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)(qreal)>(&MyQCameraFocus::Signal_MaximumDigitalZoomChanged));
}

void QCameraFocus_MaximumDigitalZoomChanged(void* ptr, double zoom)
{
	static_cast<QCameraFocus*>(ptr)->maximumDigitalZoomChanged(zoom);
}

double QCameraFocus_MaximumOpticalZoom(void* ptr)
{
	return static_cast<QCameraFocus*>(ptr)->maximumOpticalZoom();
}

void QCameraFocus_ConnectMaximumOpticalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)(qreal)>(&QCameraFocus::maximumOpticalZoomChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)(qreal)>(&MyQCameraFocus::Signal_MaximumOpticalZoomChanged));
}

void QCameraFocus_DisconnectMaximumOpticalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)(qreal)>(&QCameraFocus::maximumOpticalZoomChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)(qreal)>(&MyQCameraFocus::Signal_MaximumOpticalZoomChanged));
}

void QCameraFocus_MaximumOpticalZoomChanged(void* ptr, double zoom)
{
	static_cast<QCameraFocus*>(ptr)->maximumOpticalZoomChanged(zoom);
}

void QCameraFocus_ConnectOpticalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)(qreal)>(&QCameraFocus::opticalZoomChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)(qreal)>(&MyQCameraFocus::Signal_OpticalZoomChanged));
}

void QCameraFocus_DisconnectOpticalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)(qreal)>(&QCameraFocus::opticalZoomChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)(qreal)>(&MyQCameraFocus::Signal_OpticalZoomChanged));
}

void QCameraFocus_OpticalZoomChanged(void* ptr, double value)
{
	static_cast<QCameraFocus*>(ptr)->opticalZoomChanged(value);
}

void QCameraFocus_ZoomTo(void* ptr, double optical, double digital)
{
	static_cast<QCameraFocus*>(ptr)->zoomTo(optical, digital);
}

void QCameraFocus_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraFocus*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraFocus_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFocus*>(ptr)->QCameraFocus::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraFocus_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraFocus*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraFocus_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFocus*>(ptr)->QCameraFocus::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraFocus_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraFocus*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFocus_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraFocus*>(ptr)->QCameraFocus::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFocus_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraFocus*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraFocus_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFocus*>(ptr)->QCameraFocus::customEvent(static_cast<QEvent*>(event));
}

void QCameraFocus_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraFocus*>(ptr), "deleteLater");
}

void QCameraFocus_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraFocus*>(ptr)->QCameraFocus::deleteLater();
}

void QCameraFocus_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraFocus*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFocus_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraFocus*>(ptr)->QCameraFocus::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraFocus_Event(void* ptr, void* e)
{
	return static_cast<QCameraFocus*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraFocus_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraFocus*>(ptr)->QCameraFocus::event(static_cast<QEvent*>(e));
}

char QCameraFocus_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraFocus*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraFocus_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraFocus*>(ptr)->QCameraFocus::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraFocus_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraFocus*>(ptr)->metaObject());
}

void* QCameraFocus_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraFocus*>(ptr)->QCameraFocus::metaObject());
}

class MyQCameraFocusControl: public QCameraFocusControl
{
public:
	QPointF customFocusPoint() const { return *static_cast<QPointF*>(callbackQCameraFocusControl_CustomFocusPoint(const_cast<MyQCameraFocusControl*>(this))); };
	void Signal_CustomFocusPointChanged(const QPointF & point) { callbackQCameraFocusControl_CustomFocusPointChanged(this, const_cast<QPointF*>(&point)); };
	QCameraFocus::FocusModes focusMode() const { return static_cast<QCameraFocus::FocusMode>(callbackQCameraFocusControl_FocusMode(const_cast<MyQCameraFocusControl*>(this))); };
	void Signal_FocusModeChanged(QCameraFocus::FocusModes mode) { callbackQCameraFocusControl_FocusModeChanged(this, mode); };
	QCameraFocus::FocusPointMode focusPointMode() const { return static_cast<QCameraFocus::FocusPointMode>(callbackQCameraFocusControl_FocusPointMode(const_cast<MyQCameraFocusControl*>(this))); };
	void Signal_FocusPointModeChanged(QCameraFocus::FocusPointMode mode) { callbackQCameraFocusControl_FocusPointModeChanged(this, mode); };
	void Signal_FocusZonesChanged() { callbackQCameraFocusControl_FocusZonesChanged(this); };
	bool isFocusModeSupported(QCameraFocus::FocusModes mode) const { return callbackQCameraFocusControl_IsFocusModeSupported(const_cast<MyQCameraFocusControl*>(this), mode) != 0; };
	bool isFocusPointModeSupported(QCameraFocus::FocusPointMode mode) const { return callbackQCameraFocusControl_IsFocusPointModeSupported(const_cast<MyQCameraFocusControl*>(this), mode) != 0; };
	void setCustomFocusPoint(const QPointF & point) { callbackQCameraFocusControl_SetCustomFocusPoint(this, const_cast<QPointF*>(&point)); };
	void setFocusMode(QCameraFocus::FocusModes mode) { callbackQCameraFocusControl_SetFocusMode(this, mode); };
	void setFocusPointMode(QCameraFocus::FocusPointMode mode) { callbackQCameraFocusControl_SetFocusPointMode(this, mode); };
	void timerEvent(QTimerEvent * event) { callbackQCameraFocusControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraFocusControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraFocusControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraFocusControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraFocusControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraFocusControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraFocusControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraFocusControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraFocusControl_MetaObject(const_cast<MyQCameraFocusControl*>(this))); };
};

void* QCameraFocusControl_CustomFocusPoint(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QCameraFocusControl*>(ptr)->customFocusPoint(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QCameraFocusControl_ConnectCustomFocusPointChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(const QPointF &)>(&QCameraFocusControl::customFocusPointChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(const QPointF &)>(&MyQCameraFocusControl::Signal_CustomFocusPointChanged));
}

void QCameraFocusControl_DisconnectCustomFocusPointChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(const QPointF &)>(&QCameraFocusControl::customFocusPointChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(const QPointF &)>(&MyQCameraFocusControl::Signal_CustomFocusPointChanged));
}

void QCameraFocusControl_CustomFocusPointChanged(void* ptr, void* point)
{
	static_cast<QCameraFocusControl*>(ptr)->customFocusPointChanged(*static_cast<QPointF*>(point));
}

long long QCameraFocusControl_FocusMode(void* ptr)
{
	return static_cast<QCameraFocusControl*>(ptr)->focusMode();
}

void QCameraFocusControl_ConnectFocusModeChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusModes)>(&QCameraFocusControl::focusModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusModes)>(&MyQCameraFocusControl::Signal_FocusModeChanged));
}

void QCameraFocusControl_DisconnectFocusModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusModes)>(&QCameraFocusControl::focusModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusModes)>(&MyQCameraFocusControl::Signal_FocusModeChanged));
}

void QCameraFocusControl_FocusModeChanged(void* ptr, long long mode)
{
	static_cast<QCameraFocusControl*>(ptr)->focusModeChanged(static_cast<QCameraFocus::FocusMode>(mode));
}

long long QCameraFocusControl_FocusPointMode(void* ptr)
{
	return static_cast<QCameraFocusControl*>(ptr)->focusPointMode();
}

void QCameraFocusControl_ConnectFocusPointModeChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&QCameraFocusControl::focusPointModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&MyQCameraFocusControl::Signal_FocusPointModeChanged));
}

void QCameraFocusControl_DisconnectFocusPointModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&QCameraFocusControl::focusPointModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&MyQCameraFocusControl::Signal_FocusPointModeChanged));
}

void QCameraFocusControl_FocusPointModeChanged(void* ptr, long long mode)
{
	static_cast<QCameraFocusControl*>(ptr)->focusPointModeChanged(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocusControl_ConnectFocusZonesChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)()>(&QCameraFocusControl::focusZonesChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)()>(&MyQCameraFocusControl::Signal_FocusZonesChanged));
}

void QCameraFocusControl_DisconnectFocusZonesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)()>(&QCameraFocusControl::focusZonesChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)()>(&MyQCameraFocusControl::Signal_FocusZonesChanged));
}

void QCameraFocusControl_FocusZonesChanged(void* ptr)
{
	static_cast<QCameraFocusControl*>(ptr)->focusZonesChanged();
}

char QCameraFocusControl_IsFocusModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraFocusControl*>(ptr)->isFocusModeSupported(static_cast<QCameraFocus::FocusMode>(mode));
}

char QCameraFocusControl_IsFocusPointModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraFocusControl*>(ptr)->isFocusPointModeSupported(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocusControl_SetCustomFocusPoint(void* ptr, void* point)
{
	static_cast<QCameraFocusControl*>(ptr)->setCustomFocusPoint(*static_cast<QPointF*>(point));
}

void QCameraFocusControl_SetFocusMode(void* ptr, long long mode)
{
	static_cast<QCameraFocusControl*>(ptr)->setFocusMode(static_cast<QCameraFocus::FocusMode>(mode));
}

void QCameraFocusControl_SetFocusPointMode(void* ptr, long long mode)
{
	static_cast<QCameraFocusControl*>(ptr)->setFocusPointMode(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocusControl_DestroyQCameraFocusControl(void* ptr)
{
	static_cast<QCameraFocusControl*>(ptr)->~QCameraFocusControl();
}

void QCameraFocusControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraFocusControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraFocusControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFocusControl*>(ptr)->QCameraFocusControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraFocusControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraFocusControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraFocusControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFocusControl*>(ptr)->QCameraFocusControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraFocusControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraFocusControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFocusControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraFocusControl*>(ptr)->QCameraFocusControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFocusControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraFocusControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraFocusControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraFocusControl*>(ptr)->QCameraFocusControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraFocusControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraFocusControl*>(ptr), "deleteLater");
}

void QCameraFocusControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraFocusControl*>(ptr)->QCameraFocusControl::deleteLater();
}

void QCameraFocusControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraFocusControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraFocusControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraFocusControl*>(ptr)->QCameraFocusControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraFocusControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraFocusControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraFocusControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraFocusControl*>(ptr)->QCameraFocusControl::event(static_cast<QEvent*>(e));
}

char QCameraFocusControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraFocusControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraFocusControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraFocusControl*>(ptr)->QCameraFocusControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraFocusControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraFocusControl*>(ptr)->metaObject());
}

void* QCameraFocusControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraFocusControl*>(ptr)->QCameraFocusControl::metaObject());
}

void* QCameraFocusZone_NewQCameraFocusZone(void* other)
{
	return new QCameraFocusZone(*static_cast<QCameraFocusZone*>(other));
}

void* QCameraFocusZone_Area(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QCameraFocusZone*>(ptr)->area(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

char QCameraFocusZone_IsValid(void* ptr)
{
	return static_cast<QCameraFocusZone*>(ptr)->isValid();
}

long long QCameraFocusZone_Status(void* ptr)
{
	return static_cast<QCameraFocusZone*>(ptr)->status();
}

void QCameraFocusZone_DestroyQCameraFocusZone(void* ptr)
{
	static_cast<QCameraFocusZone*>(ptr)->~QCameraFocusZone();
}

class MyQCameraImageCapture: public QCameraImageCapture
{
public:
	MyQCameraImageCapture(QMediaObject *mediaObject, QObject *parent) : QCameraImageCapture(mediaObject, parent) {};
	void Signal_BufferFormatChanged(QVideoFrame::PixelFormat format) { callbackQCameraImageCapture_BufferFormatChanged(this, format); };
	void cancelCapture() { callbackQCameraImageCapture_CancelCapture(this); };
	int capture(const QString & file) { QByteArray t971c41 = file.toUtf8(); QtMultimedia_PackedString filePacked = { const_cast<char*>(t971c41.prepend("WHITESPACE").constData()+10), t971c41.size()-10 };return callbackQCameraImageCapture_Capture(this, filePacked); };
	void Signal_CaptureDestinationChanged(QCameraImageCapture::CaptureDestinations destination) { callbackQCameraImageCapture_CaptureDestinationChanged(this, destination); };
	void Signal_Error2(int id, QCameraImageCapture::Error error, const QString & errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtMultimedia_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQCameraImageCapture_Error2(this, id, error, errorStringPacked); };
	void Signal_ImageAvailable(int id, const QVideoFrame & buffer) { callbackQCameraImageCapture_ImageAvailable(this, id, const_cast<QVideoFrame*>(&buffer)); };
	void Signal_ImageCaptured(int id, const QImage & preview) { callbackQCameraImageCapture_ImageCaptured(this, id, const_cast<QImage*>(&preview)); };
	void Signal_ImageExposed(int id) { callbackQCameraImageCapture_ImageExposed(this, id); };
	void Signal_ImageMetadataAvailable(int id, const QString & key, const QVariant & value) { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };callbackQCameraImageCapture_ImageMetadataAvailable(this, id, keyPacked, const_cast<QVariant*>(&value)); };
	void Signal_ImageSaved(int id, const QString & fileName) { QByteArray td83e09 = fileName.toUtf8(); QtMultimedia_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQCameraImageCapture_ImageSaved(this, id, fileNamePacked); };
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQCameraImageCapture_MediaObject(const_cast<MyQCameraImageCapture*>(this))); };
	void Signal_ReadyForCaptureChanged(bool ready) { callbackQCameraImageCapture_ReadyForCaptureChanged(this, ready); };
	bool setMediaObject(QMediaObject * mediaObject) { return callbackQCameraImageCapture_SetMediaObject(this, mediaObject) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQCameraImageCapture_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraImageCapture_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraImageCapture_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraImageCapture_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraImageCapture_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraImageCapture_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraImageCapture_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraImageCapture_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraImageCapture_MetaObject(const_cast<MyQCameraImageCapture*>(this))); };
};

char QCameraImageCapture_IsReadyForCapture(void* ptr)
{
	return static_cast<QCameraImageCapture*>(ptr)->isReadyForCapture();
}

void* QCameraImageCapture_NewQCameraImageCapture(void* mediaObject, void* parent)
{
	return new MyQCameraImageCapture(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

long long QCameraImageCapture_Availability(void* ptr)
{
	return static_cast<QCameraImageCapture*>(ptr)->availability();
}

long long QCameraImageCapture_BufferFormat(void* ptr)
{
	return static_cast<QCameraImageCapture*>(ptr)->bufferFormat();
}

void QCameraImageCapture_ConnectBufferFormatChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&QCameraImageCapture::bufferFormatChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&MyQCameraImageCapture::Signal_BufferFormatChanged));
}

void QCameraImageCapture_DisconnectBufferFormatChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&QCameraImageCapture::bufferFormatChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&MyQCameraImageCapture::Signal_BufferFormatChanged));
}

void QCameraImageCapture_BufferFormatChanged(void* ptr, long long format)
{
	static_cast<QCameraImageCapture*>(ptr)->bufferFormatChanged(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraImageCapture_CancelCapture(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraImageCapture*>(ptr), "cancelCapture");
}

int QCameraImageCapture_Capture(void* ptr, char* file)
{
	int returnArg;
	QMetaObject::invokeMethod(static_cast<QCameraImageCapture*>(ptr), "capture", Q_RETURN_ARG(int, returnArg), Q_ARG(QString, QString(file)));
	return returnArg;
}

long long QCameraImageCapture_CaptureDestination(void* ptr)
{
	return static_cast<QCameraImageCapture*>(ptr)->captureDestination();
}

void QCameraImageCapture_ConnectCaptureDestinationChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraImageCapture::captureDestinationChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraImageCapture::Signal_CaptureDestinationChanged));
}

void QCameraImageCapture_DisconnectCaptureDestinationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraImageCapture::captureDestinationChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraImageCapture::Signal_CaptureDestinationChanged));
}

void QCameraImageCapture_CaptureDestinationChanged(void* ptr, long long destination)
{
	static_cast<QCameraImageCapture*>(ptr)->captureDestinationChanged(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void* QCameraImageCapture_EncodingSettings(void* ptr)
{
	return new QImageEncoderSettings(static_cast<QCameraImageCapture*>(ptr)->encodingSettings());
}

void QCameraImageCapture_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, QCameraImageCapture::Error, const QString &)>(&QCameraImageCapture::error), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, QCameraImageCapture::Error, const QString &)>(&MyQCameraImageCapture::Signal_Error2));
}

void QCameraImageCapture_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, QCameraImageCapture::Error, const QString &)>(&QCameraImageCapture::error), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, QCameraImageCapture::Error, const QString &)>(&MyQCameraImageCapture::Signal_Error2));
}

void QCameraImageCapture_Error2(void* ptr, int id, long long error, char* errorString)
{
	static_cast<QCameraImageCapture*>(ptr)->error(id, static_cast<QCameraImageCapture::Error>(error), QString(errorString));
}

long long QCameraImageCapture_Error(void* ptr)
{
	return static_cast<QCameraImageCapture*>(ptr)->error();
}

struct QtMultimedia_PackedString QCameraImageCapture_ErrorString(void* ptr)
{
	return ({ QByteArray t5acece = static_cast<QCameraImageCapture*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t5acece.prepend("WHITESPACE").constData()+10), t5acece.size()-10 }; });
}

void QCameraImageCapture_ConnectImageAvailable(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QVideoFrame &)>(&QCameraImageCapture::imageAvailable), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QVideoFrame &)>(&MyQCameraImageCapture::Signal_ImageAvailable));
}

void QCameraImageCapture_DisconnectImageAvailable(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QVideoFrame &)>(&QCameraImageCapture::imageAvailable), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QVideoFrame &)>(&MyQCameraImageCapture::Signal_ImageAvailable));
}

void QCameraImageCapture_ImageAvailable(void* ptr, int id, void* buffer)
{
	static_cast<QCameraImageCapture*>(ptr)->imageAvailable(id, *static_cast<QVideoFrame*>(buffer));
}

void QCameraImageCapture_ConnectImageCaptured(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QImage &)>(&QCameraImageCapture::imageCaptured), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QImage &)>(&MyQCameraImageCapture::Signal_ImageCaptured));
}

void QCameraImageCapture_DisconnectImageCaptured(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QImage &)>(&QCameraImageCapture::imageCaptured), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QImage &)>(&MyQCameraImageCapture::Signal_ImageCaptured));
}

void QCameraImageCapture_ImageCaptured(void* ptr, int id, void* preview)
{
	static_cast<QCameraImageCapture*>(ptr)->imageCaptured(id, *static_cast<QImage*>(preview));
}

struct QtMultimedia_PackedString QCameraImageCapture_ImageCodecDescription(void* ptr, char* codec)
{
	return ({ QByteArray tf414d7 = static_cast<QCameraImageCapture*>(ptr)->imageCodecDescription(QString(codec)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tf414d7.prepend("WHITESPACE").constData()+10), tf414d7.size()-10 }; });
}

void QCameraImageCapture_ConnectImageExposed(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int)>(&QCameraImageCapture::imageExposed), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int)>(&MyQCameraImageCapture::Signal_ImageExposed));
}

void QCameraImageCapture_DisconnectImageExposed(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int)>(&QCameraImageCapture::imageExposed), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int)>(&MyQCameraImageCapture::Signal_ImageExposed));
}

void QCameraImageCapture_ImageExposed(void* ptr, int id)
{
	static_cast<QCameraImageCapture*>(ptr)->imageExposed(id);
}

void QCameraImageCapture_ConnectImageMetadataAvailable(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &, const QVariant &)>(&QCameraImageCapture::imageMetadataAvailable), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCapture::Signal_ImageMetadataAvailable));
}

void QCameraImageCapture_DisconnectImageMetadataAvailable(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &, const QVariant &)>(&QCameraImageCapture::imageMetadataAvailable), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCapture::Signal_ImageMetadataAvailable));
}

void QCameraImageCapture_ImageMetadataAvailable(void* ptr, int id, char* key, void* value)
{
	static_cast<QCameraImageCapture*>(ptr)->imageMetadataAvailable(id, QString(key), *static_cast<QVariant*>(value));
}

void QCameraImageCapture_ConnectImageSaved(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &)>(&QCameraImageCapture::imageSaved), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &)>(&MyQCameraImageCapture::Signal_ImageSaved));
}

void QCameraImageCapture_DisconnectImageSaved(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &)>(&QCameraImageCapture::imageSaved), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &)>(&MyQCameraImageCapture::Signal_ImageSaved));
}

void QCameraImageCapture_ImageSaved(void* ptr, int id, char* fileName)
{
	static_cast<QCameraImageCapture*>(ptr)->imageSaved(id, QString(fileName));
}

char QCameraImageCapture_IsAvailable(void* ptr)
{
	return static_cast<QCameraImageCapture*>(ptr)->isAvailable();
}

char QCameraImageCapture_IsCaptureDestinationSupported(void* ptr, long long destination)
{
	return static_cast<QCameraImageCapture*>(ptr)->isCaptureDestinationSupported(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void* QCameraImageCapture_MediaObject(void* ptr)
{
	return static_cast<QCameraImageCapture*>(ptr)->mediaObject();
}

void* QCameraImageCapture_MediaObjectDefault(void* ptr)
{
	return static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::mediaObject();
}

void QCameraImageCapture_ConnectReadyForCaptureChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(bool)>(&QCameraImageCapture::readyForCaptureChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(bool)>(&MyQCameraImageCapture::Signal_ReadyForCaptureChanged));
}

void QCameraImageCapture_DisconnectReadyForCaptureChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(bool)>(&QCameraImageCapture::readyForCaptureChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(bool)>(&MyQCameraImageCapture::Signal_ReadyForCaptureChanged));
}

void QCameraImageCapture_ReadyForCaptureChanged(void* ptr, char ready)
{
	static_cast<QCameraImageCapture*>(ptr)->readyForCaptureChanged(ready != 0);
}

void QCameraImageCapture_SetBufferFormat(void* ptr, long long format)
{
	static_cast<QCameraImageCapture*>(ptr)->setBufferFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraImageCapture_SetCaptureDestination(void* ptr, long long destination)
{
	static_cast<QCameraImageCapture*>(ptr)->setCaptureDestination(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraImageCapture_SetEncodingSettings(void* ptr, void* settings)
{
	static_cast<QCameraImageCapture*>(ptr)->setEncodingSettings(*static_cast<QImageEncoderSettings*>(settings));
}

char QCameraImageCapture_SetMediaObject(void* ptr, void* mediaObject)
{
	return static_cast<QCameraImageCapture*>(ptr)->setMediaObject(static_cast<QMediaObject*>(mediaObject));
}

char QCameraImageCapture_SetMediaObjectDefault(void* ptr, void* mediaObject)
{
	return static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::setMediaObject(static_cast<QMediaObject*>(mediaObject));
}

struct QtMultimedia_PackedString QCameraImageCapture_SupportedImageCodecs(void* ptr)
{
	return ({ QByteArray t52ceb4 = static_cast<QCameraImageCapture*>(ptr)->supportedImageCodecs().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t52ceb4.prepend("WHITESPACE").constData()+10), t52ceb4.size()-10 }; });
}

void QCameraImageCapture_DestroyQCameraImageCapture(void* ptr)
{
	static_cast<QCameraImageCapture*>(ptr)->~QCameraImageCapture();
}

void QCameraImageCapture_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraImageCapture*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraImageCapture_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraImageCapture_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraImageCapture*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraImageCapture_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraImageCapture_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraImageCapture*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageCapture_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageCapture_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraImageCapture*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraImageCapture_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::customEvent(static_cast<QEvent*>(event));
}

void QCameraImageCapture_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraImageCapture*>(ptr), "deleteLater");
}

void QCameraImageCapture_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::deleteLater();
}

void QCameraImageCapture_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraImageCapture*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageCapture_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraImageCapture_Event(void* ptr, void* e)
{
	return static_cast<QCameraImageCapture*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraImageCapture_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::event(static_cast<QEvent*>(e));
}

char QCameraImageCapture_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraImageCapture*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraImageCapture_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraImageCapture_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraImageCapture*>(ptr)->metaObject());
}

void* QCameraImageCapture_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraImageCapture*>(ptr)->QCameraImageCapture::metaObject());
}

class MyQCameraImageCaptureControl: public QCameraImageCaptureControl
{
public:
	MyQCameraImageCaptureControl(QObject *parent) : QCameraImageCaptureControl(parent) {};
	void cancelCapture() { callbackQCameraImageCaptureControl_CancelCapture(this); };
	int capture(const QString & fileName) { QByteArray td83e09 = fileName.toUtf8(); QtMultimedia_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };return callbackQCameraImageCaptureControl_Capture(this, fileNamePacked); };
	QCameraImageCapture::DriveMode driveMode() const { return static_cast<QCameraImageCapture::DriveMode>(callbackQCameraImageCaptureControl_DriveMode(const_cast<MyQCameraImageCaptureControl*>(this))); };
	void Signal_Error(int id, int error, const QString & errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtMultimedia_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQCameraImageCaptureControl_Error(this, id, error, errorStringPacked); };
	void Signal_ImageAvailable(int requestId, const QVideoFrame & buffer) { callbackQCameraImageCaptureControl_ImageAvailable(this, requestId, const_cast<QVideoFrame*>(&buffer)); };
	void Signal_ImageCaptured(int requestId, const QImage & preview) { callbackQCameraImageCaptureControl_ImageCaptured(this, requestId, const_cast<QImage*>(&preview)); };
	void Signal_ImageExposed(int requestId) { callbackQCameraImageCaptureControl_ImageExposed(this, requestId); };
	void Signal_ImageMetadataAvailable(int id, const QString & key, const QVariant & value) { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };callbackQCameraImageCaptureControl_ImageMetadataAvailable(this, id, keyPacked, const_cast<QVariant*>(&value)); };
	void Signal_ImageSaved(int requestId, const QString & fileName) { QByteArray td83e09 = fileName.toUtf8(); QtMultimedia_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQCameraImageCaptureControl_ImageSaved(this, requestId, fileNamePacked); };
	bool isReadyForCapture() const { return callbackQCameraImageCaptureControl_IsReadyForCapture(const_cast<MyQCameraImageCaptureControl*>(this)) != 0; };
	void Signal_ReadyForCaptureChanged(bool ready) { callbackQCameraImageCaptureControl_ReadyForCaptureChanged(this, ready); };
	void setDriveMode(QCameraImageCapture::DriveMode mode) { callbackQCameraImageCaptureControl_SetDriveMode(this, mode); };
	void timerEvent(QTimerEvent * event) { callbackQCameraImageCaptureControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraImageCaptureControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraImageCaptureControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraImageCaptureControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraImageCaptureControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraImageCaptureControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraImageCaptureControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraImageCaptureControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraImageCaptureControl_MetaObject(const_cast<MyQCameraImageCaptureControl*>(this))); };
};

void* QCameraImageCaptureControl_NewQCameraImageCaptureControl(void* parent)
{
	return new MyQCameraImageCaptureControl(static_cast<QObject*>(parent));
}

void QCameraImageCaptureControl_CancelCapture(void* ptr)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->cancelCapture();
}

int QCameraImageCaptureControl_Capture(void* ptr, char* fileName)
{
	return static_cast<QCameraImageCaptureControl*>(ptr)->capture(QString(fileName));
}

long long QCameraImageCaptureControl_DriveMode(void* ptr)
{
	return static_cast<QCameraImageCaptureControl*>(ptr)->driveMode();
}

void QCameraImageCaptureControl_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, int, const QString &)>(&QCameraImageCaptureControl::error), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, int, const QString &)>(&MyQCameraImageCaptureControl::Signal_Error));
}

void QCameraImageCaptureControl_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, int, const QString &)>(&QCameraImageCaptureControl::error), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, int, const QString &)>(&MyQCameraImageCaptureControl::Signal_Error));
}

void QCameraImageCaptureControl_Error(void* ptr, int id, int error, char* errorString)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->error(id, error, QString(errorString));
}

void QCameraImageCaptureControl_ConnectImageAvailable(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QVideoFrame &)>(&QCameraImageCaptureControl::imageAvailable), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QVideoFrame &)>(&MyQCameraImageCaptureControl::Signal_ImageAvailable));
}

void QCameraImageCaptureControl_DisconnectImageAvailable(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QVideoFrame &)>(&QCameraImageCaptureControl::imageAvailable), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QVideoFrame &)>(&MyQCameraImageCaptureControl::Signal_ImageAvailable));
}

void QCameraImageCaptureControl_ImageAvailable(void* ptr, int requestId, void* buffer)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->imageAvailable(requestId, *static_cast<QVideoFrame*>(buffer));
}

void QCameraImageCaptureControl_ConnectImageCaptured(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QImage &)>(&QCameraImageCaptureControl::imageCaptured), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QImage &)>(&MyQCameraImageCaptureControl::Signal_ImageCaptured));
}

void QCameraImageCaptureControl_DisconnectImageCaptured(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QImage &)>(&QCameraImageCaptureControl::imageCaptured), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QImage &)>(&MyQCameraImageCaptureControl::Signal_ImageCaptured));
}

void QCameraImageCaptureControl_ImageCaptured(void* ptr, int requestId, void* preview)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->imageCaptured(requestId, *static_cast<QImage*>(preview));
}

void QCameraImageCaptureControl_ConnectImageExposed(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int)>(&QCameraImageCaptureControl::imageExposed), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int)>(&MyQCameraImageCaptureControl::Signal_ImageExposed));
}

void QCameraImageCaptureControl_DisconnectImageExposed(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int)>(&QCameraImageCaptureControl::imageExposed), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int)>(&MyQCameraImageCaptureControl::Signal_ImageExposed));
}

void QCameraImageCaptureControl_ImageExposed(void* ptr, int requestId)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->imageExposed(requestId);
}

void QCameraImageCaptureControl_ConnectImageMetadataAvailable(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&QCameraImageCaptureControl::imageMetadataAvailable), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCaptureControl::Signal_ImageMetadataAvailable));
}

void QCameraImageCaptureControl_DisconnectImageMetadataAvailable(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&QCameraImageCaptureControl::imageMetadataAvailable), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCaptureControl::Signal_ImageMetadataAvailable));
}

void QCameraImageCaptureControl_ImageMetadataAvailable(void* ptr, int id, char* key, void* value)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->imageMetadataAvailable(id, QString(key), *static_cast<QVariant*>(value));
}

void QCameraImageCaptureControl_ConnectImageSaved(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &)>(&QCameraImageCaptureControl::imageSaved), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &)>(&MyQCameraImageCaptureControl::Signal_ImageSaved));
}

void QCameraImageCaptureControl_DisconnectImageSaved(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &)>(&QCameraImageCaptureControl::imageSaved), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &)>(&MyQCameraImageCaptureControl::Signal_ImageSaved));
}

void QCameraImageCaptureControl_ImageSaved(void* ptr, int requestId, char* fileName)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->imageSaved(requestId, QString(fileName));
}

char QCameraImageCaptureControl_IsReadyForCapture(void* ptr)
{
	return static_cast<QCameraImageCaptureControl*>(ptr)->isReadyForCapture();
}

void QCameraImageCaptureControl_ConnectReadyForCaptureChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(bool)>(&QCameraImageCaptureControl::readyForCaptureChanged), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(bool)>(&MyQCameraImageCaptureControl::Signal_ReadyForCaptureChanged));
}

void QCameraImageCaptureControl_DisconnectReadyForCaptureChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(bool)>(&QCameraImageCaptureControl::readyForCaptureChanged), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(bool)>(&MyQCameraImageCaptureControl::Signal_ReadyForCaptureChanged));
}

void QCameraImageCaptureControl_ReadyForCaptureChanged(void* ptr, char ready)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->readyForCaptureChanged(ready != 0);
}

void QCameraImageCaptureControl_SetDriveMode(void* ptr, long long mode)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->setDriveMode(static_cast<QCameraImageCapture::DriveMode>(mode));
}

void QCameraImageCaptureControl_DestroyQCameraImageCaptureControl(void* ptr)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->~QCameraImageCaptureControl();
}

void QCameraImageCaptureControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraImageCaptureControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->QCameraImageCaptureControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraImageCaptureControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraImageCaptureControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->QCameraImageCaptureControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraImageCaptureControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageCaptureControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->QCameraImageCaptureControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageCaptureControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraImageCaptureControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->QCameraImageCaptureControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraImageCaptureControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraImageCaptureControl*>(ptr), "deleteLater");
}

void QCameraImageCaptureControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->QCameraImageCaptureControl::deleteLater();
}

void QCameraImageCaptureControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageCaptureControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraImageCaptureControl*>(ptr)->QCameraImageCaptureControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraImageCaptureControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraImageCaptureControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraImageCaptureControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraImageCaptureControl*>(ptr)->QCameraImageCaptureControl::event(static_cast<QEvent*>(e));
}

char QCameraImageCaptureControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraImageCaptureControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraImageCaptureControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraImageCaptureControl*>(ptr)->QCameraImageCaptureControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraImageCaptureControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraImageCaptureControl*>(ptr)->metaObject());
}

void* QCameraImageCaptureControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraImageCaptureControl*>(ptr)->QCameraImageCaptureControl::metaObject());
}

double QCameraImageProcessing_Brightness(void* ptr)
{
	return static_cast<QCameraImageProcessing*>(ptr)->brightness();
}

long long QCameraImageProcessing_ColorFilter(void* ptr)
{
	return static_cast<QCameraImageProcessing*>(ptr)->colorFilter();
}

double QCameraImageProcessing_Contrast(void* ptr)
{
	return static_cast<QCameraImageProcessing*>(ptr)->contrast();
}

double QCameraImageProcessing_DenoisingLevel(void* ptr)
{
	return static_cast<QCameraImageProcessing*>(ptr)->denoisingLevel();
}

char QCameraImageProcessing_IsAvailable(void* ptr)
{
	return static_cast<QCameraImageProcessing*>(ptr)->isAvailable();
}

char QCameraImageProcessing_IsColorFilterSupported(void* ptr, long long filter)
{
	return static_cast<QCameraImageProcessing*>(ptr)->isColorFilterSupported(static_cast<QCameraImageProcessing::ColorFilter>(filter));
}

char QCameraImageProcessing_IsWhiteBalanceModeSupported(void* ptr, long long mode)
{
	return static_cast<QCameraImageProcessing*>(ptr)->isWhiteBalanceModeSupported(static_cast<QCameraImageProcessing::WhiteBalanceMode>(mode));
}

double QCameraImageProcessing_ManualWhiteBalance(void* ptr)
{
	return static_cast<QCameraImageProcessing*>(ptr)->manualWhiteBalance();
}

double QCameraImageProcessing_Saturation(void* ptr)
{
	return static_cast<QCameraImageProcessing*>(ptr)->saturation();
}

void QCameraImageProcessing_SetBrightness(void* ptr, double value)
{
	static_cast<QCameraImageProcessing*>(ptr)->setBrightness(value);
}

void QCameraImageProcessing_SetColorFilter(void* ptr, long long filter)
{
	static_cast<QCameraImageProcessing*>(ptr)->setColorFilter(static_cast<QCameraImageProcessing::ColorFilter>(filter));
}

void QCameraImageProcessing_SetContrast(void* ptr, double value)
{
	static_cast<QCameraImageProcessing*>(ptr)->setContrast(value);
}

void QCameraImageProcessing_SetDenoisingLevel(void* ptr, double level)
{
	static_cast<QCameraImageProcessing*>(ptr)->setDenoisingLevel(level);
}

void QCameraImageProcessing_SetManualWhiteBalance(void* ptr, double colorTemperature)
{
	static_cast<QCameraImageProcessing*>(ptr)->setManualWhiteBalance(colorTemperature);
}

void QCameraImageProcessing_SetSaturation(void* ptr, double value)
{
	static_cast<QCameraImageProcessing*>(ptr)->setSaturation(value);
}

void QCameraImageProcessing_SetSharpeningLevel(void* ptr, double level)
{
	static_cast<QCameraImageProcessing*>(ptr)->setSharpeningLevel(level);
}

void QCameraImageProcessing_SetWhiteBalanceMode(void* ptr, long long mode)
{
	static_cast<QCameraImageProcessing*>(ptr)->setWhiteBalanceMode(static_cast<QCameraImageProcessing::WhiteBalanceMode>(mode));
}

double QCameraImageProcessing_SharpeningLevel(void* ptr)
{
	return static_cast<QCameraImageProcessing*>(ptr)->sharpeningLevel();
}

long long QCameraImageProcessing_WhiteBalanceMode(void* ptr)
{
	return static_cast<QCameraImageProcessing*>(ptr)->whiteBalanceMode();
}

void QCameraImageProcessing_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraImageProcessing*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraImageProcessing_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageProcessing*>(ptr)->QCameraImageProcessing::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraImageProcessing_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraImageProcessing*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraImageProcessing_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageProcessing*>(ptr)->QCameraImageProcessing::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraImageProcessing_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraImageProcessing*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageProcessing_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraImageProcessing*>(ptr)->QCameraImageProcessing::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageProcessing_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraImageProcessing*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraImageProcessing_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageProcessing*>(ptr)->QCameraImageProcessing::customEvent(static_cast<QEvent*>(event));
}

void QCameraImageProcessing_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraImageProcessing*>(ptr), "deleteLater");
}

void QCameraImageProcessing_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraImageProcessing*>(ptr)->QCameraImageProcessing::deleteLater();
}

void QCameraImageProcessing_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraImageProcessing*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageProcessing_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraImageProcessing*>(ptr)->QCameraImageProcessing::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraImageProcessing_Event(void* ptr, void* e)
{
	return static_cast<QCameraImageProcessing*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraImageProcessing_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraImageProcessing*>(ptr)->QCameraImageProcessing::event(static_cast<QEvent*>(e));
}

char QCameraImageProcessing_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraImageProcessing*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraImageProcessing_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraImageProcessing*>(ptr)->QCameraImageProcessing::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraImageProcessing_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraImageProcessing*>(ptr)->metaObject());
}

void* QCameraImageProcessing_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraImageProcessing*>(ptr)->QCameraImageProcessing::metaObject());
}

class MyQCameraImageProcessingControl: public QCameraImageProcessingControl
{
public:
	MyQCameraImageProcessingControl(QObject *parent) : QCameraImageProcessingControl(parent) {};
	bool isParameterSupported(QCameraImageProcessingControl::ProcessingParameter parameter) const { return callbackQCameraImageProcessingControl_IsParameterSupported(const_cast<MyQCameraImageProcessingControl*>(this), parameter) != 0; };
	bool isParameterValueSupported(QCameraImageProcessingControl::ProcessingParameter parameter, const QVariant & value) const { return callbackQCameraImageProcessingControl_IsParameterValueSupported(const_cast<MyQCameraImageProcessingControl*>(this), parameter, const_cast<QVariant*>(&value)) != 0; };
	QVariant parameter(QCameraImageProcessingControl::ProcessingParameter parameter) const { return *static_cast<QVariant*>(callbackQCameraImageProcessingControl_Parameter(const_cast<MyQCameraImageProcessingControl*>(this), parameter)); };
	void setParameter(QCameraImageProcessingControl::ProcessingParameter parameter, const QVariant & value) { callbackQCameraImageProcessingControl_SetParameter(this, parameter, const_cast<QVariant*>(&value)); };
	void timerEvent(QTimerEvent * event) { callbackQCameraImageProcessingControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraImageProcessingControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraImageProcessingControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraImageProcessingControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraImageProcessingControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraImageProcessingControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraImageProcessingControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraImageProcessingControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraImageProcessingControl_MetaObject(const_cast<MyQCameraImageProcessingControl*>(this))); };
};

void* QCameraImageProcessingControl_NewQCameraImageProcessingControl(void* parent)
{
	return new MyQCameraImageProcessingControl(static_cast<QObject*>(parent));
}

char QCameraImageProcessingControl_IsParameterSupported(void* ptr, long long parameter)
{
	return static_cast<QCameraImageProcessingControl*>(ptr)->isParameterSupported(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter));
}

char QCameraImageProcessingControl_IsParameterValueSupported(void* ptr, long long parameter, void* value)
{
	return static_cast<QCameraImageProcessingControl*>(ptr)->isParameterValueSupported(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter), *static_cast<QVariant*>(value));
}

void* QCameraImageProcessingControl_Parameter(void* ptr, long long parameter)
{
	return new QVariant(static_cast<QCameraImageProcessingControl*>(ptr)->parameter(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter)));
}

void QCameraImageProcessingControl_SetParameter(void* ptr, long long parameter, void* value)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->setParameter(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter), *static_cast<QVariant*>(value));
}

void QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(void* ptr)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->~QCameraImageProcessingControl();
}

void QCameraImageProcessingControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraImageProcessingControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->QCameraImageProcessingControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraImageProcessingControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraImageProcessingControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->QCameraImageProcessingControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraImageProcessingControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageProcessingControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->QCameraImageProcessingControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageProcessingControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraImageProcessingControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->QCameraImageProcessingControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraImageProcessingControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraImageProcessingControl*>(ptr), "deleteLater");
}

void QCameraImageProcessingControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->QCameraImageProcessingControl::deleteLater();
}

void QCameraImageProcessingControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraImageProcessingControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraImageProcessingControl*>(ptr)->QCameraImageProcessingControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraImageProcessingControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraImageProcessingControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraImageProcessingControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraImageProcessingControl*>(ptr)->QCameraImageProcessingControl::event(static_cast<QEvent*>(e));
}

char QCameraImageProcessingControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraImageProcessingControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraImageProcessingControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraImageProcessingControl*>(ptr)->QCameraImageProcessingControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraImageProcessingControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraImageProcessingControl*>(ptr)->metaObject());
}

void* QCameraImageProcessingControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraImageProcessingControl*>(ptr)->QCameraImageProcessingControl::metaObject());
}

void* QCameraInfo_NewQCameraInfo(void* name)
{
	return new QCameraInfo(*static_cast<QByteArray*>(name));
}

void* QCameraInfo_NewQCameraInfo2(void* camera)
{
	return new QCameraInfo(*static_cast<QCamera*>(camera));
}

void* QCameraInfo_NewQCameraInfo3(void* other)
{
	return new QCameraInfo(*static_cast<QCameraInfo*>(other));
}

void* QCameraInfo_QCameraInfo_DefaultCamera()
{
	return new QCameraInfo(QCameraInfo::defaultCamera());
}

struct QtMultimedia_PackedString QCameraInfo_Description(void* ptr)
{
	return ({ QByteArray t58a341 = static_cast<QCameraInfo*>(ptr)->description().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t58a341.prepend("WHITESPACE").constData()+10), t58a341.size()-10 }; });
}

struct QtMultimedia_PackedString QCameraInfo_DeviceName(void* ptr)
{
	return ({ QByteArray tb7aa82 = static_cast<QCameraInfo*>(ptr)->deviceName().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tb7aa82.prepend("WHITESPACE").constData()+10), tb7aa82.size()-10 }; });
}

char QCameraInfo_IsNull(void* ptr)
{
	return static_cast<QCameraInfo*>(ptr)->isNull();
}

int QCameraInfo_Orientation(void* ptr)
{
	return static_cast<QCameraInfo*>(ptr)->orientation();
}

long long QCameraInfo_Position(void* ptr)
{
	return static_cast<QCameraInfo*>(ptr)->position();
}

void QCameraInfo_DestroyQCameraInfo(void* ptr)
{
	static_cast<QCameraInfo*>(ptr)->~QCameraInfo();
}

class MyQCameraInfoControl: public QCameraInfoControl
{
public:
	MyQCameraInfoControl(QObject *parent) : QCameraInfoControl(parent) {};
	int cameraOrientation(const QString & deviceName) const { QByteArray t0f5a7b = deviceName.toUtf8(); QtMultimedia_PackedString deviceNamePacked = { const_cast<char*>(t0f5a7b.prepend("WHITESPACE").constData()+10), t0f5a7b.size()-10 };return callbackQCameraInfoControl_CameraOrientation(const_cast<MyQCameraInfoControl*>(this), deviceNamePacked); };
	QCamera::Position cameraPosition(const QString & deviceName) const { QByteArray t0f5a7b = deviceName.toUtf8(); QtMultimedia_PackedString deviceNamePacked = { const_cast<char*>(t0f5a7b.prepend("WHITESPACE").constData()+10), t0f5a7b.size()-10 };return static_cast<QCamera::Position>(callbackQCameraInfoControl_CameraPosition(const_cast<MyQCameraInfoControl*>(this), deviceNamePacked)); };
	 ~MyQCameraInfoControl() { callbackQCameraInfoControl_DestroyQCameraInfoControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQCameraInfoControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraInfoControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraInfoControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraInfoControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraInfoControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraInfoControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraInfoControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraInfoControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraInfoControl_MetaObject(const_cast<MyQCameraInfoControl*>(this))); };
};

void* QCameraInfoControl_NewQCameraInfoControl(void* parent)
{
	return new MyQCameraInfoControl(static_cast<QObject*>(parent));
}

int QCameraInfoControl_CameraOrientation(void* ptr, char* deviceName)
{
	return static_cast<QCameraInfoControl*>(ptr)->cameraOrientation(QString(deviceName));
}

long long QCameraInfoControl_CameraPosition(void* ptr, char* deviceName)
{
	return static_cast<QCameraInfoControl*>(ptr)->cameraPosition(QString(deviceName));
}

void QCameraInfoControl_DestroyQCameraInfoControl(void* ptr)
{
	static_cast<QCameraInfoControl*>(ptr)->~QCameraInfoControl();
}

void QCameraInfoControl_DestroyQCameraInfoControlDefault(void* ptr)
{

}

void QCameraInfoControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraInfoControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraInfoControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraInfoControl*>(ptr)->QCameraInfoControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraInfoControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraInfoControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraInfoControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraInfoControl*>(ptr)->QCameraInfoControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraInfoControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraInfoControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraInfoControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraInfoControl*>(ptr)->QCameraInfoControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraInfoControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraInfoControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraInfoControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraInfoControl*>(ptr)->QCameraInfoControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraInfoControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraInfoControl*>(ptr), "deleteLater");
}

void QCameraInfoControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraInfoControl*>(ptr)->QCameraInfoControl::deleteLater();
}

void QCameraInfoControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraInfoControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraInfoControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraInfoControl*>(ptr)->QCameraInfoControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraInfoControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraInfoControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraInfoControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraInfoControl*>(ptr)->QCameraInfoControl::event(static_cast<QEvent*>(e));
}

char QCameraInfoControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraInfoControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraInfoControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraInfoControl*>(ptr)->QCameraInfoControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraInfoControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraInfoControl*>(ptr)->metaObject());
}

void* QCameraInfoControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraInfoControl*>(ptr)->QCameraInfoControl::metaObject());
}

class MyQCameraLocksControl: public QCameraLocksControl
{
public:
	MyQCameraLocksControl(QObject *parent) : QCameraLocksControl(parent) {};
	QCamera::LockStatus lockStatus(QCamera::LockType lock) const { return static_cast<QCamera::LockStatus>(callbackQCameraLocksControl_LockStatus(const_cast<MyQCameraLocksControl*>(this), lock)); };
	void Signal_LockStatusChanged(QCamera::LockType lock, QCamera::LockStatus status, QCamera::LockChangeReason reason) { callbackQCameraLocksControl_LockStatusChanged(this, lock, status, reason); };
	void searchAndLock(QCamera::LockTypes locks) { callbackQCameraLocksControl_SearchAndLock(this, locks); };
	QCamera::LockTypes supportedLocks() const { return static_cast<QCamera::LockType>(callbackQCameraLocksControl_SupportedLocks(const_cast<MyQCameraLocksControl*>(this))); };
	void unlock(QCamera::LockTypes locks) { callbackQCameraLocksControl_Unlock(this, locks); };
	void timerEvent(QTimerEvent * event) { callbackQCameraLocksControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraLocksControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraLocksControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraLocksControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraLocksControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraLocksControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraLocksControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraLocksControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraLocksControl_MetaObject(const_cast<MyQCameraLocksControl*>(this))); };
};

void* QCameraLocksControl_NewQCameraLocksControl(void* parent)
{
	return new MyQCameraLocksControl(static_cast<QObject*>(parent));
}

long long QCameraLocksControl_LockStatus(void* ptr, long long lock)
{
	return static_cast<QCameraLocksControl*>(ptr)->lockStatus(static_cast<QCamera::LockType>(lock));
}

void QCameraLocksControl_ConnectLockStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraLocksControl*>(ptr), static_cast<void (QCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&QCameraLocksControl::lockStatusChanged), static_cast<MyQCameraLocksControl*>(ptr), static_cast<void (MyQCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCameraLocksControl::Signal_LockStatusChanged));
}

void QCameraLocksControl_DisconnectLockStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraLocksControl*>(ptr), static_cast<void (QCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&QCameraLocksControl::lockStatusChanged), static_cast<MyQCameraLocksControl*>(ptr), static_cast<void (MyQCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCameraLocksControl::Signal_LockStatusChanged));
}

void QCameraLocksControl_LockStatusChanged(void* ptr, long long lock, long long status, long long reason)
{
	static_cast<QCameraLocksControl*>(ptr)->lockStatusChanged(static_cast<QCamera::LockType>(lock), static_cast<QCamera::LockStatus>(status), static_cast<QCamera::LockChangeReason>(reason));
}

void QCameraLocksControl_SearchAndLock(void* ptr, long long locks)
{
	static_cast<QCameraLocksControl*>(ptr)->searchAndLock(static_cast<QCamera::LockType>(locks));
}

long long QCameraLocksControl_SupportedLocks(void* ptr)
{
	return static_cast<QCameraLocksControl*>(ptr)->supportedLocks();
}

void QCameraLocksControl_Unlock(void* ptr, long long locks)
{
	static_cast<QCameraLocksControl*>(ptr)->unlock(static_cast<QCamera::LockType>(locks));
}

void QCameraLocksControl_DestroyQCameraLocksControl(void* ptr)
{
	static_cast<QCameraLocksControl*>(ptr)->~QCameraLocksControl();
}

void QCameraLocksControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraLocksControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraLocksControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraLocksControl*>(ptr)->QCameraLocksControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraLocksControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraLocksControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraLocksControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraLocksControl*>(ptr)->QCameraLocksControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraLocksControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraLocksControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraLocksControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraLocksControl*>(ptr)->QCameraLocksControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraLocksControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraLocksControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraLocksControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraLocksControl*>(ptr)->QCameraLocksControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraLocksControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraLocksControl*>(ptr), "deleteLater");
}

void QCameraLocksControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraLocksControl*>(ptr)->QCameraLocksControl::deleteLater();
}

void QCameraLocksControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraLocksControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraLocksControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraLocksControl*>(ptr)->QCameraLocksControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraLocksControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraLocksControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraLocksControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraLocksControl*>(ptr)->QCameraLocksControl::event(static_cast<QEvent*>(e));
}

char QCameraLocksControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraLocksControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraLocksControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraLocksControl*>(ptr)->QCameraLocksControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraLocksControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraLocksControl*>(ptr)->metaObject());
}

void* QCameraLocksControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraLocksControl*>(ptr)->QCameraLocksControl::metaObject());
}

class MyQCameraViewfinder: public QCameraViewfinder
{
public:
	MyQCameraViewfinder(QWidget *parent) : QCameraViewfinder(parent) {};
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQCameraViewfinder_MediaObject(const_cast<MyQCameraViewfinder*>(this))); };
	bool setMediaObject(QMediaObject * object) { return callbackQCameraViewfinder_SetMediaObject(this, object) != 0; };
	void setAspectRatioMode(Qt::AspectRatioMode mode) { callbackQCameraViewfinder_SetAspectRatioMode(this, mode); };
	void setBrightness(int brightness) { callbackQCameraViewfinder_SetBrightness(this, brightness); };
	void setContrast(int contrast) { callbackQCameraViewfinder_SetContrast(this, contrast); };
	void setFullScreen(bool fullScreen) { callbackQCameraViewfinder_SetFullScreen(this, fullScreen); };
	void setHue(int hue) { callbackQCameraViewfinder_SetHue(this, hue); };
	void setSaturation(int saturation) { callbackQCameraViewfinder_SetSaturation(this, saturation); };
	void actionEvent(QActionEvent * event) { callbackQCameraViewfinder_ActionEvent(this, event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQCameraViewfinder_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQCameraViewfinder_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQCameraViewfinder_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQCameraViewfinder_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQCameraViewfinder_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQCameraViewfinder_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQCameraViewfinder_FocusOutEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQCameraViewfinder_LeaveEvent(this, event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQCameraViewfinder_MinimumSizeHint(const_cast<MyQCameraViewfinder*>(this))); };
	void setEnabled(bool vbo) { callbackQCameraViewfinder_SetEnabled(this, vbo); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtMultimedia_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQCameraViewfinder_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQCameraViewfinder_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQCameraViewfinder_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtMultimedia_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQCameraViewfinder_SetWindowTitle(this, vqsPacked); };
	void changeEvent(QEvent * event) { callbackQCameraViewfinder_ChangeEvent(this, event); };
	bool close() { return callbackQCameraViewfinder_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQCameraViewfinder_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQCameraViewfinder_ContextMenuEvent(this, event); };
	bool focusNextPrevChild(bool next) { return callbackQCameraViewfinder_FocusNextPrevChild(this, next) != 0; };
	bool hasHeightForWidth() const { return callbackQCameraViewfinder_HasHeightForWidth(const_cast<MyQCameraViewfinder*>(this)) != 0; };
	int heightForWidth(int w) const { return callbackQCameraViewfinder_HeightForWidth(const_cast<MyQCameraViewfinder*>(this), w); };
	void hide() { callbackQCameraViewfinder_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQCameraViewfinder_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQCameraViewfinder_InputMethodQuery(const_cast<MyQCameraViewfinder*>(this), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQCameraViewfinder_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQCameraViewfinder_KeyReleaseEvent(this, event); };
	void lower() { callbackQCameraViewfinder_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQCameraViewfinder_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQCameraViewfinder_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQCameraViewfinder_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQCameraViewfinder_MouseReleaseEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQCameraViewfinder_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void raise() { callbackQCameraViewfinder_Raise(this); };
	void repaint() { callbackQCameraViewfinder_Repaint(this); };
	void setDisabled(bool disable) { callbackQCameraViewfinder_SetDisabled(this, disable); };
	void setFocus() { callbackQCameraViewfinder_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQCameraViewfinder_SetHidden(this, hidden); };
	void show() { callbackQCameraViewfinder_Show(this); };
	void showFullScreen() { callbackQCameraViewfinder_ShowFullScreen(this); };
	void showMaximized() { callbackQCameraViewfinder_ShowMaximized(this); };
	void showMinimized() { callbackQCameraViewfinder_ShowMinimized(this); };
	void showNormal() { callbackQCameraViewfinder_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQCameraViewfinder_TabletEvent(this, event); };
	void update() { callbackQCameraViewfinder_Update(this); };
	void updateMicroFocus() { callbackQCameraViewfinder_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQCameraViewfinder_WheelEvent(this, event); };
	void timerEvent(QTimerEvent * event) { callbackQCameraViewfinder_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraViewfinder_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraViewfinder_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraViewfinder_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraViewfinder_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraViewfinder_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraViewfinder_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraViewfinder_MetaObject(const_cast<MyQCameraViewfinder*>(this))); };
};

void* QCameraViewfinder_NewQCameraViewfinder(void* parent)
{
	return new MyQCameraViewfinder(static_cast<QWidget*>(parent));
}

void* QCameraViewfinder_MediaObject(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->mediaObject();
}

void* QCameraViewfinder_MediaObjectDefault(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mediaObject();
}

char QCameraViewfinder_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QCameraViewfinder*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

char QCameraViewfinder_SetMediaObjectDefault(void* ptr, void* object)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setMediaObject(static_cast<QMediaObject*>(object));
}

void QCameraViewfinder_DestroyQCameraViewfinder(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->~QCameraViewfinder();
}

void QCameraViewfinder_SetAspectRatioMode(void* ptr, long long mode)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setAspectRatioMode", Q_ARG(Qt::AspectRatioMode, static_cast<Qt::AspectRatioMode>(mode)));
}

void QCameraViewfinder_SetAspectRatioModeDefault(void* ptr, long long mode)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QCameraViewfinder_SetBrightness(void* ptr, int brightness)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setBrightness", Q_ARG(int, brightness));
}

void QCameraViewfinder_SetBrightnessDefault(void* ptr, int brightness)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setBrightness(brightness);
}

void QCameraViewfinder_SetContrast(void* ptr, int contrast)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setContrast", Q_ARG(int, contrast));
}

void QCameraViewfinder_SetContrastDefault(void* ptr, int contrast)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setContrast(contrast);
}

void QCameraViewfinder_SetFullScreen(void* ptr, char fullScreen)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setFullScreen", Q_ARG(bool, fullScreen != 0));
}

void QCameraViewfinder_SetFullScreenDefault(void* ptr, char fullScreen)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setFullScreen(fullScreen != 0);
}

void QCameraViewfinder_SetHue(void* ptr, int hue)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setHue", Q_ARG(int, hue));
}

void QCameraViewfinder_SetHueDefault(void* ptr, int hue)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setHue(hue);
}

void QCameraViewfinder_SetSaturation(void* ptr, int saturation)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setSaturation", Q_ARG(int, saturation));
}

void QCameraViewfinder_SetSaturationDefault(void* ptr, int saturation)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setSaturation(saturation);
}

void QCameraViewfinder_ActionEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QCameraViewfinder_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::actionEvent(static_cast<QActionEvent*>(event));
}

void QCameraViewfinder_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QCameraViewfinder_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QCameraViewfinder_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QCameraViewfinder_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QCameraViewfinder_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QCameraViewfinder_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QCameraViewfinder_DropEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QCameraViewfinder_DropEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::dropEvent(static_cast<QDropEvent*>(event));
}

void QCameraViewfinder_EnterEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::enterEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_FocusInEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QCameraViewfinder_LeaveEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::leaveEvent(static_cast<QEvent*>(event));
}

void* QCameraViewfinder_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QCameraViewfinder*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QCameraViewfinder_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QCameraViewfinder_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QCameraViewfinder_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setEnabled(vbo != 0);
}

void QCameraViewfinder_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QCameraViewfinder_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setStyleSheet(QString(styleSheet));
}

void QCameraViewfinder_SetVisible(void* ptr, char visible)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QCameraViewfinder_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setVisible(visible != 0);
}

void QCameraViewfinder_SetWindowModified(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QCameraViewfinder_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setWindowModified(vbo != 0);
}

void QCameraViewfinder_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QCameraViewfinder_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setWindowTitle(QString(vqs));
}

void QCameraViewfinder_ChangeEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::changeEvent(static_cast<QEvent*>(event));
}

char QCameraViewfinder_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QCameraViewfinder_CloseDefault(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::close();
}

void QCameraViewfinder_CloseEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QCameraViewfinder_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::closeEvent(static_cast<QCloseEvent*>(event));
}

void QCameraViewfinder_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QCameraViewfinder_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

char QCameraViewfinder_FocusNextPrevChild(void* ptr, char next)
{
	return static_cast<QCameraViewfinder*>(ptr)->focusNextPrevChild(next != 0);
}

char QCameraViewfinder_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::focusNextPrevChild(next != 0);
}

char QCameraViewfinder_HasHeightForWidth(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->hasHeightForWidth();
}

char QCameraViewfinder_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::hasHeightForWidth();
}

int QCameraViewfinder_HeightForWidth(void* ptr, int w)
{
	return static_cast<QCameraViewfinder*>(ptr)->heightForWidth(w);
}

int QCameraViewfinder_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::heightForWidth(w);
}

void QCameraViewfinder_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "hide");
}

void QCameraViewfinder_HideDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::hide();
}

void QCameraViewfinder_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QCameraViewfinder_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QCameraViewfinder_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QCameraViewfinder*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QCameraViewfinder_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QCameraViewfinder_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QCameraViewfinder_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "lower");
}

void QCameraViewfinder_LowerDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::lower();
}

void QCameraViewfinder_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MousePressEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QCameraViewfinder_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

char QCameraViewfinder_NativeEvent(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QCameraViewfinder*>(ptr)->nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

char QCameraViewfinder_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QCameraViewfinder_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "raise");
}

void QCameraViewfinder_RaiseDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::raise();
}

void QCameraViewfinder_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "repaint");
}

void QCameraViewfinder_RepaintDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::repaint();
}

void QCameraViewfinder_SetDisabled(void* ptr, char disable)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QCameraViewfinder_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setDisabled(disable != 0);
}

void QCameraViewfinder_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setFocus");
}

void QCameraViewfinder_SetFocus2Default(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setFocus();
}

void QCameraViewfinder_SetHidden(void* ptr, char hidden)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QCameraViewfinder_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::setHidden(hidden != 0);
}

void QCameraViewfinder_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "show");
}

void QCameraViewfinder_ShowDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::show();
}

void QCameraViewfinder_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "showFullScreen");
}

void QCameraViewfinder_ShowFullScreenDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::showFullScreen();
}

void QCameraViewfinder_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "showMaximized");
}

void QCameraViewfinder_ShowMaximizedDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::showMaximized();
}

void QCameraViewfinder_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "showMinimized");
}

void QCameraViewfinder_ShowMinimizedDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::showMinimized();
}

void QCameraViewfinder_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "showNormal");
}

void QCameraViewfinder_ShowNormalDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::showNormal();
}

void QCameraViewfinder_TabletEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QCameraViewfinder_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QCameraViewfinder_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "update");
}

void QCameraViewfinder_UpdateDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::update();
}

void QCameraViewfinder_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "updateMicroFocus");
}

void QCameraViewfinder_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::updateMicroFocus();
}

void QCameraViewfinder_WheelEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QCameraViewfinder_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QCameraViewfinder_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinder_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinder_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinder_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinder_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraViewfinder*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinder_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinder_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::customEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinder_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinder*>(ptr), "deleteLater");
}

void QCameraViewfinder_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::deleteLater();
}

void QCameraViewfinder_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraViewfinder*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinder_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraViewfinder_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraViewfinder*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraViewfinder_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraViewfinder_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraViewfinder*>(ptr)->metaObject());
}

void* QCameraViewfinder_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraViewfinder*>(ptr)->QCameraViewfinder::metaObject());
}

void* QCameraViewfinderSettings_NewQCameraViewfinderSettings()
{
	return new QCameraViewfinderSettings();
}

void* QCameraViewfinderSettings_NewQCameraViewfinderSettings2(void* other)
{
	return new QCameraViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(other));
}

char QCameraViewfinderSettings_IsNull(void* ptr)
{
	return static_cast<QCameraViewfinderSettings*>(ptr)->isNull();
}

double QCameraViewfinderSettings_MaximumFrameRate(void* ptr)
{
	return static_cast<QCameraViewfinderSettings*>(ptr)->maximumFrameRate();
}

double QCameraViewfinderSettings_MinimumFrameRate(void* ptr)
{
	return static_cast<QCameraViewfinderSettings*>(ptr)->minimumFrameRate();
}

void* QCameraViewfinderSettings_PixelAspectRatio(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QCameraViewfinderSettings*>(ptr)->pixelAspectRatio(); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long QCameraViewfinderSettings_PixelFormat(void* ptr)
{
	return static_cast<QCameraViewfinderSettings*>(ptr)->pixelFormat();
}

void* QCameraViewfinderSettings_Resolution(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QCameraViewfinderSettings*>(ptr)->resolution(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QCameraViewfinderSettings_SetMaximumFrameRate(void* ptr, double rate)
{
	static_cast<QCameraViewfinderSettings*>(ptr)->setMaximumFrameRate(rate);
}

void QCameraViewfinderSettings_SetMinimumFrameRate(void* ptr, double rate)
{
	static_cast<QCameraViewfinderSettings*>(ptr)->setMinimumFrameRate(rate);
}

void QCameraViewfinderSettings_SetPixelAspectRatio(void* ptr, void* ratio)
{
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelAspectRatio(*static_cast<QSize*>(ratio));
}

void QCameraViewfinderSettings_SetPixelAspectRatio2(void* ptr, int horizontal, int vertical)
{
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelAspectRatio(horizontal, vertical);
}

void QCameraViewfinderSettings_SetPixelFormat(void* ptr, long long format)
{
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraViewfinderSettings_SetResolution(void* ptr, void* resolution)
{
	static_cast<QCameraViewfinderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QCameraViewfinderSettings_SetResolution2(void* ptr, int width, int height)
{
	static_cast<QCameraViewfinderSettings*>(ptr)->setResolution(width, height);
}

void QCameraViewfinderSettings_Swap(void* ptr, void* other)
{
	static_cast<QCameraViewfinderSettings*>(ptr)->swap(*static_cast<QCameraViewfinderSettings*>(other));
}

void QCameraViewfinderSettings_DestroyQCameraViewfinderSettings(void* ptr)
{
	static_cast<QCameraViewfinderSettings*>(ptr)->~QCameraViewfinderSettings();
}

class MyQCameraViewfinderSettingsControl: public QCameraViewfinderSettingsControl
{
public:
	MyQCameraViewfinderSettingsControl(QObject *parent) : QCameraViewfinderSettingsControl(parent) {};
	bool isViewfinderParameterSupported(QCameraViewfinderSettingsControl::ViewfinderParameter parameter) const { return callbackQCameraViewfinderSettingsControl_IsViewfinderParameterSupported(const_cast<MyQCameraViewfinderSettingsControl*>(this), parameter) != 0; };
	void setViewfinderParameter(QCameraViewfinderSettingsControl::ViewfinderParameter parameter, const QVariant & value) { callbackQCameraViewfinderSettingsControl_SetViewfinderParameter(this, parameter, const_cast<QVariant*>(&value)); };
	QVariant viewfinderParameter(QCameraViewfinderSettingsControl::ViewfinderParameter parameter) const { return *static_cast<QVariant*>(callbackQCameraViewfinderSettingsControl_ViewfinderParameter(const_cast<MyQCameraViewfinderSettingsControl*>(this), parameter)); };
	void timerEvent(QTimerEvent * event) { callbackQCameraViewfinderSettingsControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraViewfinderSettingsControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraViewfinderSettingsControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraViewfinderSettingsControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraViewfinderSettingsControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraViewfinderSettingsControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraViewfinderSettingsControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraViewfinderSettingsControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraViewfinderSettingsControl_MetaObject(const_cast<MyQCameraViewfinderSettingsControl*>(this))); };
};

void* QCameraViewfinderSettingsControl_NewQCameraViewfinderSettingsControl(void* parent)
{
	return new MyQCameraViewfinderSettingsControl(static_cast<QObject*>(parent));
}

char QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(void* ptr, long long parameter)
{
	return static_cast<QCameraViewfinderSettingsControl*>(ptr)->isViewfinderParameterSupported(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter));
}

void QCameraViewfinderSettingsControl_SetViewfinderParameter(void* ptr, long long parameter, void* value)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->setViewfinderParameter(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter), *static_cast<QVariant*>(value));
}

void* QCameraViewfinderSettingsControl_ViewfinderParameter(void* ptr, long long parameter)
{
	return new QVariant(static_cast<QCameraViewfinderSettingsControl*>(ptr)->viewfinderParameter(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter)));
}

void QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(void* ptr)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->~QCameraViewfinderSettingsControl();
}

void QCameraViewfinderSettingsControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinderSettingsControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->QCameraViewfinderSettingsControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinderSettingsControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinderSettingsControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->QCameraViewfinderSettingsControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinderSettingsControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinderSettingsControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->QCameraViewfinderSettingsControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinderSettingsControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinderSettingsControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->QCameraViewfinderSettingsControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinderSettingsControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinderSettingsControl*>(ptr), "deleteLater");
}

void QCameraViewfinderSettingsControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->QCameraViewfinderSettingsControl::deleteLater();
}

void QCameraViewfinderSettingsControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinderSettingsControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->QCameraViewfinderSettingsControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraViewfinderSettingsControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraViewfinderSettingsControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraViewfinderSettingsControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraViewfinderSettingsControl*>(ptr)->QCameraViewfinderSettingsControl::event(static_cast<QEvent*>(e));
}

char QCameraViewfinderSettingsControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraViewfinderSettingsControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraViewfinderSettingsControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraViewfinderSettingsControl*>(ptr)->QCameraViewfinderSettingsControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraViewfinderSettingsControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraViewfinderSettingsControl*>(ptr)->metaObject());
}

void* QCameraViewfinderSettingsControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraViewfinderSettingsControl*>(ptr)->QCameraViewfinderSettingsControl::metaObject());
}

class MyQCameraViewfinderSettingsControl2: public QCameraViewfinderSettingsControl2
{
public:
	void setViewfinderSettings(const QCameraViewfinderSettings & settings) { callbackQCameraViewfinderSettingsControl2_SetViewfinderSettings(this, const_cast<QCameraViewfinderSettings*>(&settings)); };
	QCameraViewfinderSettings viewfinderSettings() const { return *static_cast<QCameraViewfinderSettings*>(callbackQCameraViewfinderSettingsControl2_ViewfinderSettings(const_cast<MyQCameraViewfinderSettingsControl2*>(this))); };
	 ~MyQCameraViewfinderSettingsControl2() { callbackQCameraViewfinderSettingsControl2_DestroyQCameraViewfinderSettingsControl2(this); };
	void timerEvent(QTimerEvent * event) { callbackQCameraViewfinderSettingsControl2_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraViewfinderSettingsControl2_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraViewfinderSettingsControl2_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraViewfinderSettingsControl2_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraViewfinderSettingsControl2_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraViewfinderSettingsControl2_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraViewfinderSettingsControl2_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraViewfinderSettingsControl2_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraViewfinderSettingsControl2_MetaObject(const_cast<MyQCameraViewfinderSettingsControl2*>(this))); };
};

void QCameraViewfinderSettingsControl2_SetViewfinderSettings(void* ptr, void* settings)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->setViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(settings));
}

void* QCameraViewfinderSettingsControl2_ViewfinderSettings(void* ptr)
{
	return new QCameraViewfinderSettings(static_cast<QCameraViewfinderSettingsControl2*>(ptr)->viewfinderSettings());
}

void QCameraViewfinderSettingsControl2_DestroyQCameraViewfinderSettingsControl2(void* ptr)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->~QCameraViewfinderSettingsControl2();
}

void QCameraViewfinderSettingsControl2_DestroyQCameraViewfinderSettingsControl2Default(void* ptr)
{

}

void QCameraViewfinderSettingsControl2_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinderSettingsControl2_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->QCameraViewfinderSettingsControl2::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraViewfinderSettingsControl2_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinderSettingsControl2_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->QCameraViewfinderSettingsControl2::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraViewfinderSettingsControl2_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinderSettingsControl2_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->QCameraViewfinderSettingsControl2::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinderSettingsControl2_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinderSettingsControl2_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->QCameraViewfinderSettingsControl2::customEvent(static_cast<QEvent*>(event));
}

void QCameraViewfinderSettingsControl2_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraViewfinderSettingsControl2*>(ptr), "deleteLater");
}

void QCameraViewfinderSettingsControl2_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->QCameraViewfinderSettingsControl2::deleteLater();
}

void QCameraViewfinderSettingsControl2_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraViewfinderSettingsControl2_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->QCameraViewfinderSettingsControl2::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraViewfinderSettingsControl2_Event(void* ptr, void* e)
{
	return static_cast<QCameraViewfinderSettingsControl2*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraViewfinderSettingsControl2_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraViewfinderSettingsControl2*>(ptr)->QCameraViewfinderSettingsControl2::event(static_cast<QEvent*>(e));
}

char QCameraViewfinderSettingsControl2_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraViewfinderSettingsControl2*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraViewfinderSettingsControl2_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraViewfinderSettingsControl2*>(ptr)->QCameraViewfinderSettingsControl2::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraViewfinderSettingsControl2_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraViewfinderSettingsControl2*>(ptr)->metaObject());
}

void* QCameraViewfinderSettingsControl2_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraViewfinderSettingsControl2*>(ptr)->QCameraViewfinderSettingsControl2::metaObject());
}

class MyQCameraZoomControl: public QCameraZoomControl
{
public:
	MyQCameraZoomControl(QObject *parent) : QCameraZoomControl(parent) {};
	qreal currentDigitalZoom() const { return callbackQCameraZoomControl_CurrentDigitalZoom(const_cast<MyQCameraZoomControl*>(this)); };
	void Signal_CurrentDigitalZoomChanged(qreal zoom) { callbackQCameraZoomControl_CurrentDigitalZoomChanged(this, zoom); };
	qreal currentOpticalZoom() const { return callbackQCameraZoomControl_CurrentOpticalZoom(const_cast<MyQCameraZoomControl*>(this)); };
	void Signal_CurrentOpticalZoomChanged(qreal zoom) { callbackQCameraZoomControl_CurrentOpticalZoomChanged(this, zoom); };
	qreal maximumDigitalZoom() const { return callbackQCameraZoomControl_MaximumDigitalZoom(const_cast<MyQCameraZoomControl*>(this)); };
	void Signal_MaximumDigitalZoomChanged(qreal zoom) { callbackQCameraZoomControl_MaximumDigitalZoomChanged(this, zoom); };
	qreal maximumOpticalZoom() const { return callbackQCameraZoomControl_MaximumOpticalZoom(const_cast<MyQCameraZoomControl*>(this)); };
	void Signal_MaximumOpticalZoomChanged(qreal zoom) { callbackQCameraZoomControl_MaximumOpticalZoomChanged(this, zoom); };
	qreal requestedDigitalZoom() const { return callbackQCameraZoomControl_RequestedDigitalZoom(const_cast<MyQCameraZoomControl*>(this)); };
	void Signal_RequestedDigitalZoomChanged(qreal zoom) { callbackQCameraZoomControl_RequestedDigitalZoomChanged(this, zoom); };
	qreal requestedOpticalZoom() const { return callbackQCameraZoomControl_RequestedOpticalZoom(const_cast<MyQCameraZoomControl*>(this)); };
	void Signal_RequestedOpticalZoomChanged(qreal zoom) { callbackQCameraZoomControl_RequestedOpticalZoomChanged(this, zoom); };
	void zoomTo(qreal optical, qreal digital) { callbackQCameraZoomControl_ZoomTo(this, optical, digital); };
	void timerEvent(QTimerEvent * event) { callbackQCameraZoomControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQCameraZoomControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCameraZoomControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQCameraZoomControl_CustomEvent(this, event); };
	void deleteLater() { callbackQCameraZoomControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCameraZoomControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQCameraZoomControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCameraZoomControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCameraZoomControl_MetaObject(const_cast<MyQCameraZoomControl*>(this))); };
};

void* QCameraZoomControl_NewQCameraZoomControl(void* parent)
{
	return new MyQCameraZoomControl(static_cast<QObject*>(parent));
}

double QCameraZoomControl_CurrentDigitalZoom(void* ptr)
{
	return static_cast<QCameraZoomControl*>(ptr)->currentDigitalZoom();
}

void QCameraZoomControl_ConnectCurrentDigitalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::currentDigitalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_CurrentDigitalZoomChanged));
}

void QCameraZoomControl_DisconnectCurrentDigitalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::currentDigitalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_CurrentDigitalZoomChanged));
}

void QCameraZoomControl_CurrentDigitalZoomChanged(void* ptr, double zoom)
{
	static_cast<QCameraZoomControl*>(ptr)->currentDigitalZoomChanged(zoom);
}

double QCameraZoomControl_CurrentOpticalZoom(void* ptr)
{
	return static_cast<QCameraZoomControl*>(ptr)->currentOpticalZoom();
}

void QCameraZoomControl_ConnectCurrentOpticalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::currentOpticalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_CurrentOpticalZoomChanged));
}

void QCameraZoomControl_DisconnectCurrentOpticalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::currentOpticalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_CurrentOpticalZoomChanged));
}

void QCameraZoomControl_CurrentOpticalZoomChanged(void* ptr, double zoom)
{
	static_cast<QCameraZoomControl*>(ptr)->currentOpticalZoomChanged(zoom);
}

double QCameraZoomControl_MaximumDigitalZoom(void* ptr)
{
	return static_cast<QCameraZoomControl*>(ptr)->maximumDigitalZoom();
}

void QCameraZoomControl_ConnectMaximumDigitalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::maximumDigitalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_MaximumDigitalZoomChanged));
}

void QCameraZoomControl_DisconnectMaximumDigitalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::maximumDigitalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_MaximumDigitalZoomChanged));
}

void QCameraZoomControl_MaximumDigitalZoomChanged(void* ptr, double zoom)
{
	static_cast<QCameraZoomControl*>(ptr)->maximumDigitalZoomChanged(zoom);
}

double QCameraZoomControl_MaximumOpticalZoom(void* ptr)
{
	return static_cast<QCameraZoomControl*>(ptr)->maximumOpticalZoom();
}

void QCameraZoomControl_ConnectMaximumOpticalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::maximumOpticalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_MaximumOpticalZoomChanged));
}

void QCameraZoomControl_DisconnectMaximumOpticalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::maximumOpticalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_MaximumOpticalZoomChanged));
}

void QCameraZoomControl_MaximumOpticalZoomChanged(void* ptr, double zoom)
{
	static_cast<QCameraZoomControl*>(ptr)->maximumOpticalZoomChanged(zoom);
}

double QCameraZoomControl_RequestedDigitalZoom(void* ptr)
{
	return static_cast<QCameraZoomControl*>(ptr)->requestedDigitalZoom();
}

void QCameraZoomControl_ConnectRequestedDigitalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::requestedDigitalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_RequestedDigitalZoomChanged));
}

void QCameraZoomControl_DisconnectRequestedDigitalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::requestedDigitalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_RequestedDigitalZoomChanged));
}

void QCameraZoomControl_RequestedDigitalZoomChanged(void* ptr, double zoom)
{
	static_cast<QCameraZoomControl*>(ptr)->requestedDigitalZoomChanged(zoom);
}

double QCameraZoomControl_RequestedOpticalZoom(void* ptr)
{
	return static_cast<QCameraZoomControl*>(ptr)->requestedOpticalZoom();
}

void QCameraZoomControl_ConnectRequestedOpticalZoomChanged(void* ptr)
{
	QObject::connect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::requestedOpticalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_RequestedOpticalZoomChanged));
}

void QCameraZoomControl_DisconnectRequestedOpticalZoomChanged(void* ptr)
{
	QObject::disconnect(static_cast<QCameraZoomControl*>(ptr), static_cast<void (QCameraZoomControl::*)(qreal)>(&QCameraZoomControl::requestedOpticalZoomChanged), static_cast<MyQCameraZoomControl*>(ptr), static_cast<void (MyQCameraZoomControl::*)(qreal)>(&MyQCameraZoomControl::Signal_RequestedOpticalZoomChanged));
}

void QCameraZoomControl_RequestedOpticalZoomChanged(void* ptr, double zoom)
{
	static_cast<QCameraZoomControl*>(ptr)->requestedOpticalZoomChanged(zoom);
}

void QCameraZoomControl_ZoomTo(void* ptr, double optical, double digital)
{
	static_cast<QCameraZoomControl*>(ptr)->zoomTo(optical, digital);
}

void QCameraZoomControl_DestroyQCameraZoomControl(void* ptr)
{
	static_cast<QCameraZoomControl*>(ptr)->~QCameraZoomControl();
}

void QCameraZoomControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QCameraZoomControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraZoomControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCameraZoomControl*>(ptr)->QCameraZoomControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCameraZoomControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QCameraZoomControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCameraZoomControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCameraZoomControl*>(ptr)->QCameraZoomControl::childEvent(static_cast<QChildEvent*>(event));
}

void QCameraZoomControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraZoomControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraZoomControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraZoomControl*>(ptr)->QCameraZoomControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraZoomControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QCameraZoomControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCameraZoomControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCameraZoomControl*>(ptr)->QCameraZoomControl::customEvent(static_cast<QEvent*>(event));
}

void QCameraZoomControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCameraZoomControl*>(ptr), "deleteLater");
}

void QCameraZoomControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QCameraZoomControl*>(ptr)->QCameraZoomControl::deleteLater();
}

void QCameraZoomControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCameraZoomControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCameraZoomControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCameraZoomControl*>(ptr)->QCameraZoomControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QCameraZoomControl_Event(void* ptr, void* e)
{
	return static_cast<QCameraZoomControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QCameraZoomControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QCameraZoomControl*>(ptr)->QCameraZoomControl::event(static_cast<QEvent*>(e));
}

char QCameraZoomControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraZoomControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QCameraZoomControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCameraZoomControl*>(ptr)->QCameraZoomControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCameraZoomControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraZoomControl*>(ptr)->metaObject());
}

void* QCameraZoomControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCameraZoomControl*>(ptr)->QCameraZoomControl::metaObject());
}

class MyQGraphicsVideoItem: public QGraphicsVideoItem
{
public:
	MyQGraphicsVideoItem(QGraphicsItem *parent) : QGraphicsVideoItem(parent) {};
	void Signal_NativeSizeChanged(const QSizeF & size) { callbackQGraphicsVideoItem_NativeSizeChanged(this, const_cast<QSizeF*>(&size)); };
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQGraphicsVideoItem_MediaObject(const_cast<MyQGraphicsVideoItem*>(this))); };
	void updateMicroFocus() { callbackQGraphicsVideoItem_UpdateMicroFocus(this); };
	void timerEvent(QTimerEvent * event) { callbackQGraphicsVideoItem_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQGraphicsVideoItem_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGraphicsVideoItem_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGraphicsVideoItem_CustomEvent(this, event); };
	void deleteLater() { callbackQGraphicsVideoItem_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGraphicsVideoItem_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGraphicsVideoItem_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGraphicsVideoItem_MetaObject(const_cast<MyQGraphicsVideoItem*>(this))); };
	void advance(int phase) { callbackQGraphicsVideoItem_Advance(this, phase); };
	bool collidesWithItem(const QGraphicsItem * other, Qt::ItemSelectionMode mode) const { return callbackQGraphicsVideoItem_CollidesWithItem(const_cast<MyQGraphicsVideoItem*>(this), const_cast<QGraphicsItem*>(other), mode) != 0; };
	bool collidesWithPath(const QPainterPath & path, Qt::ItemSelectionMode mode) const { return callbackQGraphicsVideoItem_CollidesWithPath(const_cast<MyQGraphicsVideoItem*>(this), const_cast<QPainterPath*>(&path), mode) != 0; };
	bool contains(const QPointF & point) const { return callbackQGraphicsVideoItem_Contains(const_cast<MyQGraphicsVideoItem*>(this), const_cast<QPointF*>(&point)) != 0; };
	void contextMenuEvent(QGraphicsSceneContextMenuEvent * event) { callbackQGraphicsVideoItem_ContextMenuEvent(this, event); };
	void dragEnterEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsVideoItem_DragEnterEvent(this, event); };
	void dragLeaveEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsVideoItem_DragLeaveEvent(this, event); };
	void dragMoveEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsVideoItem_DragMoveEvent(this, event); };
	void dropEvent(QGraphicsSceneDragDropEvent * event) { callbackQGraphicsVideoItem_DropEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQGraphicsVideoItem_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQGraphicsVideoItem_FocusOutEvent(this, event); };
	void hoverEnterEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsVideoItem_HoverEnterEvent(this, event); };
	void hoverLeaveEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsVideoItem_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QGraphicsSceneHoverEvent * event) { callbackQGraphicsVideoItem_HoverMoveEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQGraphicsVideoItem_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQGraphicsVideoItem_InputMethodQuery(const_cast<MyQGraphicsVideoItem*>(this), query)); };
	bool isObscuredBy(const QGraphicsItem * item) const { return callbackQGraphicsVideoItem_IsObscuredBy(const_cast<MyQGraphicsVideoItem*>(this), const_cast<QGraphicsItem*>(item)) != 0; };
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant & value) { return *static_cast<QVariant*>(callbackQGraphicsVideoItem_ItemChange(this, change, const_cast<QVariant*>(&value))); };
	void keyPressEvent(QKeyEvent * event) { callbackQGraphicsVideoItem_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQGraphicsVideoItem_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsVideoItem_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsVideoItem_MouseMoveEvent(this, event); };
	void mousePressEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsVideoItem_MousePressEvent(this, event); };
	void mouseReleaseEvent(QGraphicsSceneMouseEvent * event) { callbackQGraphicsVideoItem_MouseReleaseEvent(this, event); };
	QPainterPath opaqueArea() const { return *static_cast<QPainterPath*>(callbackQGraphicsVideoItem_OpaqueArea(const_cast<MyQGraphicsVideoItem*>(this))); };
	bool sceneEvent(QEvent * event) { return callbackQGraphicsVideoItem_SceneEvent(this, event) != 0; };
	bool sceneEventFilter(QGraphicsItem * watched, QEvent * event) { return callbackQGraphicsVideoItem_SceneEventFilter(this, watched, event) != 0; };
	QPainterPath shape() const { return *static_cast<QPainterPath*>(callbackQGraphicsVideoItem_Shape(const_cast<MyQGraphicsVideoItem*>(this))); };
	int type() const { return callbackQGraphicsVideoItem_Type(const_cast<MyQGraphicsVideoItem*>(this)); };
	void wheelEvent(QGraphicsSceneWheelEvent * event) { callbackQGraphicsVideoItem_WheelEvent(this, event); };
	bool setMediaObject(QMediaObject * object) { return callbackQGraphicsVideoItem_SetMediaObject(this, object) != 0; };
};

void QGraphicsVideoItem_ConnectNativeSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QGraphicsVideoItem*>(ptr), static_cast<void (QGraphicsVideoItem::*)(const QSizeF &)>(&QGraphicsVideoItem::nativeSizeChanged), static_cast<MyQGraphicsVideoItem*>(ptr), static_cast<void (MyQGraphicsVideoItem::*)(const QSizeF &)>(&MyQGraphicsVideoItem::Signal_NativeSizeChanged));
}

void QGraphicsVideoItem_DisconnectNativeSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGraphicsVideoItem*>(ptr), static_cast<void (QGraphicsVideoItem::*)(const QSizeF &)>(&QGraphicsVideoItem::nativeSizeChanged), static_cast<MyQGraphicsVideoItem*>(ptr), static_cast<void (MyQGraphicsVideoItem::*)(const QSizeF &)>(&MyQGraphicsVideoItem::Signal_NativeSizeChanged));
}

void QGraphicsVideoItem_NativeSizeChanged(void* ptr, void* size)
{
	static_cast<QGraphicsVideoItem*>(ptr)->nativeSizeChanged(*static_cast<QSizeF*>(size));
}

void* QGraphicsVideoItem_NewQGraphicsVideoItem(void* parent)
{
	return new MyQGraphicsVideoItem(static_cast<QGraphicsItem*>(parent));
}

long long QGraphicsVideoItem_AspectRatioMode(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->aspectRatioMode();
}

void* QGraphicsVideoItem_BoundingRect(void* ptr)
{
	return ({ QRectF tmpValue = static_cast<QGraphicsVideoItem*>(ptr)->boundingRect(); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsVideoItem_MediaObject(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->mediaObject();
}

void* QGraphicsVideoItem_MediaObjectDefault(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mediaObject();
}

void* QGraphicsVideoItem_NativeSize(void* ptr)
{
	return ({ QSizeF tmpValue = static_cast<QGraphicsVideoItem*>(ptr)->nativeSize(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void* QGraphicsVideoItem_Offset(void* ptr)
{
	return ({ QPointF tmpValue = static_cast<QGraphicsVideoItem*>(ptr)->offset(); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QGraphicsVideoItem_Paint(void* ptr, void* painter, void* option, void* widget)
{
	static_cast<QGraphicsVideoItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsVideoItem_SetAspectRatioMode(void* ptr, long long mode)
{
	static_cast<QGraphicsVideoItem*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QGraphicsVideoItem_SetOffset(void* ptr, void* offset)
{
	static_cast<QGraphicsVideoItem*>(ptr)->setOffset(*static_cast<QPointF*>(offset));
}

void QGraphicsVideoItem_SetSize(void* ptr, void* size)
{
	static_cast<QGraphicsVideoItem*>(ptr)->setSize(*static_cast<QSizeF*>(size));
}

void* QGraphicsVideoItem_Size(void* ptr)
{
	return ({ QSizeF tmpValue = static_cast<QGraphicsVideoItem*>(ptr)->size(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

void QGraphicsVideoItem_DestroyQGraphicsVideoItem(void* ptr)
{
	static_cast<QGraphicsVideoItem*>(ptr)->~QGraphicsVideoItem();
}

void QGraphicsVideoItem_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsVideoItem*>(ptr), "updateMicroFocus");
}

void QGraphicsVideoItem_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::updateMicroFocus();
}

void QGraphicsVideoItem_TimerEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsVideoItem_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGraphicsVideoItem_ChildEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsVideoItem_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::childEvent(static_cast<QChildEvent*>(event));
}

void QGraphicsVideoItem_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGraphicsVideoItem*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsVideoItem_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsVideoItem_CustomEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGraphicsVideoItem_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::customEvent(static_cast<QEvent*>(event));
}

void QGraphicsVideoItem_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGraphicsVideoItem*>(ptr), "deleteLater");
}

void QGraphicsVideoItem_DeleteLaterDefault(void* ptr)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::deleteLater();
}

void QGraphicsVideoItem_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGraphicsVideoItem*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGraphicsVideoItem_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGraphicsVideoItem_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGraphicsVideoItem_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGraphicsVideoItem_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGraphicsVideoItem*>(ptr)->metaObject());
}

void* QGraphicsVideoItem_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::metaObject());
}

void QGraphicsVideoItem_Advance(void* ptr, int phase)
{
	static_cast<QGraphicsVideoItem*>(ptr)->advance(phase);
}

void QGraphicsVideoItem_AdvanceDefault(void* ptr, int phase)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::advance(phase);
}

char QGraphicsVideoItem_CollidesWithItem(void* ptr, void* other, long long mode)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsVideoItem_CollidesWithItemDefault(void* ptr, void* other, long long mode)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::collidesWithItem(static_cast<QGraphicsItem*>(other), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsVideoItem_CollidesWithPath(void* ptr, void* path, long long mode)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsVideoItem_CollidesWithPathDefault(void* ptr, void* path, long long mode)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::collidesWithPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode));
}

char QGraphicsVideoItem_Contains(void* ptr, void* point)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

char QGraphicsVideoItem_ContainsDefault(void* ptr, void* point)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::contains(*static_cast<QPointF*>(point));
}

void QGraphicsVideoItem_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QGraphicsVideoItem_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::contextMenuEvent(static_cast<QGraphicsSceneContextMenuEvent*>(event));
}

void QGraphicsVideoItem_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::dragEnterEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::dragLeaveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::dragMoveEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DropEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_DropEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::dropEvent(static_cast<QGraphicsSceneDragDropEvent*>(event));
}

void QGraphicsVideoItem_FocusInEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsVideoItem_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsVideoItem_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsVideoItem_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QGraphicsVideoItem_HoverEnterEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverEnterEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::hoverEnterEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverLeaveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::hoverLeaveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverMoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_HoverMoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::hoverMoveEvent(static_cast<QGraphicsSceneHoverEvent*>(event));
}

void QGraphicsVideoItem_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QGraphicsVideoItem_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QGraphicsVideoItem_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QGraphicsVideoItem*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QGraphicsVideoItem_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QGraphicsVideoItem_IsObscuredBy(void* ptr, void* item)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

char QGraphicsVideoItem_IsObscuredByDefault(void* ptr, void* item)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void* QGraphicsVideoItem_ItemChange(void* ptr, long long change, void* value)
{
	return new QVariant(static_cast<QGraphicsVideoItem*>(ptr)->itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

void* QGraphicsVideoItem_ItemChangeDefault(void* ptr, long long change, void* value)
{
	return new QVariant(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::itemChange(static_cast<QGraphicsItem::GraphicsItemChange>(change), *static_cast<QVariant*>(value)));
}

void QGraphicsVideoItem_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsVideoItem_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsVideoItem_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsVideoItem_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QGraphicsVideoItem_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mouseDoubleClickEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mouseMoveEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MousePressEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mousePressEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void QGraphicsVideoItem_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::mouseReleaseEvent(static_cast<QGraphicsSceneMouseEvent*>(event));
}

void* QGraphicsVideoItem_OpaqueArea(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsVideoItem*>(ptr)->opaqueArea());
}

void* QGraphicsVideoItem_OpaqueAreaDefault(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::opaqueArea());
}

char QGraphicsVideoItem_SceneEvent(void* ptr, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->sceneEvent(static_cast<QEvent*>(event));
}

char QGraphicsVideoItem_SceneEventDefault(void* ptr, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::sceneEvent(static_cast<QEvent*>(event));
}

char QGraphicsVideoItem_SceneEventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

char QGraphicsVideoItem_SceneEventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::sceneEventFilter(static_cast<QGraphicsItem*>(watched), static_cast<QEvent*>(event));
}

void* QGraphicsVideoItem_Shape(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsVideoItem*>(ptr)->shape());
}

void* QGraphicsVideoItem_ShapeDefault(void* ptr)
{
	return new QPainterPath(static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::shape());
}

int QGraphicsVideoItem_Type(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->type();
}

int QGraphicsVideoItem_TypeDefault(void* ptr)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::type();
}

void QGraphicsVideoItem_WheelEvent(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

void QGraphicsVideoItem_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::wheelEvent(static_cast<QGraphicsSceneWheelEvent*>(event));
}

char QGraphicsVideoItem_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

char QGraphicsVideoItem_SetMediaObjectDefault(void* ptr, void* object)
{
	return static_cast<QGraphicsVideoItem*>(ptr)->QGraphicsVideoItem::setMediaObject(static_cast<QMediaObject*>(object));
}

class MyQImageEncoderControl: public QImageEncoderControl
{
public:
	QString imageCodecDescription(const QString & codec) const { QByteArray td061f6 = codec.toUtf8(); QtMultimedia_PackedString codecPacked = { const_cast<char*>(td061f6.prepend("WHITESPACE").constData()+10), td061f6.size()-10 };return QString(callbackQImageEncoderControl_ImageCodecDescription(const_cast<MyQImageEncoderControl*>(this), codecPacked)); };
	QImageEncoderSettings imageSettings() const { return *static_cast<QImageEncoderSettings*>(callbackQImageEncoderControl_ImageSettings(const_cast<MyQImageEncoderControl*>(this))); };
	void setImageSettings(const QImageEncoderSettings & settings) { callbackQImageEncoderControl_SetImageSettings(this, const_cast<QImageEncoderSettings*>(&settings)); };
	QStringList supportedImageCodecs() const { return QString(callbackQImageEncoderControl_SupportedImageCodecs(const_cast<MyQImageEncoderControl*>(this))).split("|", QString::SkipEmptyParts); };
	 ~MyQImageEncoderControl() { callbackQImageEncoderControl_DestroyQImageEncoderControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQImageEncoderControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQImageEncoderControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQImageEncoderControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQImageEncoderControl_CustomEvent(this, event); };
	void deleteLater() { callbackQImageEncoderControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQImageEncoderControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQImageEncoderControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQImageEncoderControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQImageEncoderControl_MetaObject(const_cast<MyQImageEncoderControl*>(this))); };
};

struct QtMultimedia_PackedString QImageEncoderControl_ImageCodecDescription(void* ptr, char* codec)
{
	return ({ QByteArray t7716ec = static_cast<QImageEncoderControl*>(ptr)->imageCodecDescription(QString(codec)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t7716ec.prepend("WHITESPACE").constData()+10), t7716ec.size()-10 }; });
}

void* QImageEncoderControl_ImageSettings(void* ptr)
{
	return new QImageEncoderSettings(static_cast<QImageEncoderControl*>(ptr)->imageSettings());
}

void QImageEncoderControl_SetImageSettings(void* ptr, void* settings)
{
	static_cast<QImageEncoderControl*>(ptr)->setImageSettings(*static_cast<QImageEncoderSettings*>(settings));
}

struct QtMultimedia_PackedString QImageEncoderControl_SupportedImageCodecs(void* ptr)
{
	return ({ QByteArray t0f36dd = static_cast<QImageEncoderControl*>(ptr)->supportedImageCodecs().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t0f36dd.prepend("WHITESPACE").constData()+10), t0f36dd.size()-10 }; });
}

void QImageEncoderControl_DestroyQImageEncoderControl(void* ptr)
{
	static_cast<QImageEncoderControl*>(ptr)->~QImageEncoderControl();
}

void QImageEncoderControl_DestroyQImageEncoderControlDefault(void* ptr)
{

}

void QImageEncoderControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QImageEncoderControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QImageEncoderControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QImageEncoderControl*>(ptr)->QImageEncoderControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QImageEncoderControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QImageEncoderControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QImageEncoderControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QImageEncoderControl*>(ptr)->QImageEncoderControl::childEvent(static_cast<QChildEvent*>(event));
}

void QImageEncoderControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QImageEncoderControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QImageEncoderControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QImageEncoderControl*>(ptr)->QImageEncoderControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QImageEncoderControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QImageEncoderControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QImageEncoderControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QImageEncoderControl*>(ptr)->QImageEncoderControl::customEvent(static_cast<QEvent*>(event));
}

void QImageEncoderControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QImageEncoderControl*>(ptr), "deleteLater");
}

void QImageEncoderControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QImageEncoderControl*>(ptr)->QImageEncoderControl::deleteLater();
}

void QImageEncoderControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QImageEncoderControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QImageEncoderControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QImageEncoderControl*>(ptr)->QImageEncoderControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QImageEncoderControl_Event(void* ptr, void* e)
{
	return static_cast<QImageEncoderControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QImageEncoderControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QImageEncoderControl*>(ptr)->QImageEncoderControl::event(static_cast<QEvent*>(e));
}

char QImageEncoderControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QImageEncoderControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QImageEncoderControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QImageEncoderControl*>(ptr)->QImageEncoderControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QImageEncoderControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QImageEncoderControl*>(ptr)->metaObject());
}

void* QImageEncoderControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QImageEncoderControl*>(ptr)->QImageEncoderControl::metaObject());
}

void* QImageEncoderSettings_NewQImageEncoderSettings()
{
	return new QImageEncoderSettings();
}

void* QImageEncoderSettings_NewQImageEncoderSettings2(void* other)
{
	return new QImageEncoderSettings(*static_cast<QImageEncoderSettings*>(other));
}

struct QtMultimedia_PackedString QImageEncoderSettings_Codec(void* ptr)
{
	return ({ QByteArray t008bcd = static_cast<QImageEncoderSettings*>(ptr)->codec().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t008bcd.prepend("WHITESPACE").constData()+10), t008bcd.size()-10 }; });
}

void* QImageEncoderSettings_EncodingOption(void* ptr, char* option)
{
	return new QVariant(static_cast<QImageEncoderSettings*>(ptr)->encodingOption(QString(option)));
}

char QImageEncoderSettings_IsNull(void* ptr)
{
	return static_cast<QImageEncoderSettings*>(ptr)->isNull();
}

long long QImageEncoderSettings_Quality(void* ptr)
{
	return static_cast<QImageEncoderSettings*>(ptr)->quality();
}

void* QImageEncoderSettings_Resolution(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QImageEncoderSettings*>(ptr)->resolution(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QImageEncoderSettings_SetCodec(void* ptr, char* codec)
{
	static_cast<QImageEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QImageEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value)
{
	static_cast<QImageEncoderSettings*>(ptr)->setEncodingOption(QString(option), *static_cast<QVariant*>(value));
}

void QImageEncoderSettings_SetQuality(void* ptr, long long quality)
{
	static_cast<QImageEncoderSettings*>(ptr)->setQuality(static_cast<QMultimedia::EncodingQuality>(quality));
}

void QImageEncoderSettings_SetResolution(void* ptr, void* resolution)
{
	static_cast<QImageEncoderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QImageEncoderSettings_SetResolution2(void* ptr, int width, int height)
{
	static_cast<QImageEncoderSettings*>(ptr)->setResolution(width, height);
}

void QImageEncoderSettings_DestroyQImageEncoderSettings(void* ptr)
{
	static_cast<QImageEncoderSettings*>(ptr)->~QImageEncoderSettings();
}

class MyQMediaAudioProbeControl: public QMediaAudioProbeControl
{
public:
	MyQMediaAudioProbeControl(QObject *parent) : QMediaAudioProbeControl(parent) {};
	void Signal_AudioBufferProbed(const QAudioBuffer & buffer) { callbackQMediaAudioProbeControl_AudioBufferProbed(this, const_cast<QAudioBuffer*>(&buffer)); };
	void Signal_Flush() { callbackQMediaAudioProbeControl_Flush(this); };
	 ~MyQMediaAudioProbeControl() { callbackQMediaAudioProbeControl_DestroyQMediaAudioProbeControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQMediaAudioProbeControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaAudioProbeControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaAudioProbeControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaAudioProbeControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaAudioProbeControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaAudioProbeControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaAudioProbeControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaAudioProbeControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaAudioProbeControl_MetaObject(const_cast<MyQMediaAudioProbeControl*>(this))); };
};

void* QMediaAudioProbeControl_NewQMediaAudioProbeControl(void* parent)
{
	return new MyQMediaAudioProbeControl(static_cast<QObject*>(parent));
}

void QMediaAudioProbeControl_ConnectAudioBufferProbed(void* ptr)
{
	QObject::connect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)(const QAudioBuffer &)>(&QMediaAudioProbeControl::audioBufferProbed), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)(const QAudioBuffer &)>(&MyQMediaAudioProbeControl::Signal_AudioBufferProbed));
}

void QMediaAudioProbeControl_DisconnectAudioBufferProbed(void* ptr)
{
	QObject::disconnect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)(const QAudioBuffer &)>(&QMediaAudioProbeControl::audioBufferProbed), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)(const QAudioBuffer &)>(&MyQMediaAudioProbeControl::Signal_AudioBufferProbed));
}

void QMediaAudioProbeControl_AudioBufferProbed(void* ptr, void* buffer)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->audioBufferProbed(*static_cast<QAudioBuffer*>(buffer));
}

void QMediaAudioProbeControl_ConnectFlush(void* ptr)
{
	QObject::connect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)()>(&QMediaAudioProbeControl::flush), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)()>(&MyQMediaAudioProbeControl::Signal_Flush));
}

void QMediaAudioProbeControl_DisconnectFlush(void* ptr)
{
	QObject::disconnect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)()>(&QMediaAudioProbeControl::flush), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)()>(&MyQMediaAudioProbeControl::Signal_Flush));
}

void QMediaAudioProbeControl_Flush(void* ptr)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->flush();
}

void QMediaAudioProbeControl_DestroyQMediaAudioProbeControl(void* ptr)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->~QMediaAudioProbeControl();
}

void QMediaAudioProbeControl_DestroyQMediaAudioProbeControlDefault(void* ptr)
{

}

void QMediaAudioProbeControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaAudioProbeControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->QMediaAudioProbeControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaAudioProbeControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaAudioProbeControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->QMediaAudioProbeControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaAudioProbeControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaAudioProbeControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->QMediaAudioProbeControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaAudioProbeControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaAudioProbeControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->QMediaAudioProbeControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaAudioProbeControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaAudioProbeControl*>(ptr), "deleteLater");
}

void QMediaAudioProbeControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->QMediaAudioProbeControl::deleteLater();
}

void QMediaAudioProbeControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaAudioProbeControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaAudioProbeControl*>(ptr)->QMediaAudioProbeControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaAudioProbeControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaAudioProbeControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaAudioProbeControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaAudioProbeControl*>(ptr)->QMediaAudioProbeControl::event(static_cast<QEvent*>(e));
}

char QMediaAudioProbeControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaAudioProbeControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaAudioProbeControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaAudioProbeControl*>(ptr)->QMediaAudioProbeControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaAudioProbeControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaAudioProbeControl*>(ptr)->metaObject());
}

void* QMediaAudioProbeControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaAudioProbeControl*>(ptr)->QMediaAudioProbeControl::metaObject());
}

class MyQMediaAvailabilityControl: public QMediaAvailabilityControl
{
public:
	MyQMediaAvailabilityControl(QObject *parent) : QMediaAvailabilityControl(parent) {};
	QMultimedia::AvailabilityStatus availability() const { return static_cast<QMultimedia::AvailabilityStatus>(callbackQMediaAvailabilityControl_Availability(const_cast<MyQMediaAvailabilityControl*>(this))); };
	void Signal_AvailabilityChanged(QMultimedia::AvailabilityStatus availability) { callbackQMediaAvailabilityControl_AvailabilityChanged(this, availability); };
	void timerEvent(QTimerEvent * event) { callbackQMediaAvailabilityControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaAvailabilityControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaAvailabilityControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaAvailabilityControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaAvailabilityControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaAvailabilityControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaAvailabilityControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaAvailabilityControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaAvailabilityControl_MetaObject(const_cast<MyQMediaAvailabilityControl*>(this))); };
};

void* QMediaAvailabilityControl_NewQMediaAvailabilityControl(void* parent)
{
	return new MyQMediaAvailabilityControl(static_cast<QObject*>(parent));
}

long long QMediaAvailabilityControl_Availability(void* ptr)
{
	return static_cast<QMediaAvailabilityControl*>(ptr)->availability();
}

void QMediaAvailabilityControl_ConnectAvailabilityChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaAvailabilityControl*>(ptr), static_cast<void (QMediaAvailabilityControl::*)(QMultimedia::AvailabilityStatus)>(&QMediaAvailabilityControl::availabilityChanged), static_cast<MyQMediaAvailabilityControl*>(ptr), static_cast<void (MyQMediaAvailabilityControl::*)(QMultimedia::AvailabilityStatus)>(&MyQMediaAvailabilityControl::Signal_AvailabilityChanged));
}

void QMediaAvailabilityControl_DisconnectAvailabilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaAvailabilityControl*>(ptr), static_cast<void (QMediaAvailabilityControl::*)(QMultimedia::AvailabilityStatus)>(&QMediaAvailabilityControl::availabilityChanged), static_cast<MyQMediaAvailabilityControl*>(ptr), static_cast<void (MyQMediaAvailabilityControl::*)(QMultimedia::AvailabilityStatus)>(&MyQMediaAvailabilityControl::Signal_AvailabilityChanged));
}

void QMediaAvailabilityControl_AvailabilityChanged(void* ptr, long long availability)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->availabilityChanged(static_cast<QMultimedia::AvailabilityStatus>(availability));
}

void QMediaAvailabilityControl_DestroyQMediaAvailabilityControl(void* ptr)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->~QMediaAvailabilityControl();
}

void QMediaAvailabilityControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaAvailabilityControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->QMediaAvailabilityControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaAvailabilityControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaAvailabilityControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->QMediaAvailabilityControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaAvailabilityControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaAvailabilityControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->QMediaAvailabilityControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaAvailabilityControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaAvailabilityControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->QMediaAvailabilityControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaAvailabilityControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaAvailabilityControl*>(ptr), "deleteLater");
}

void QMediaAvailabilityControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->QMediaAvailabilityControl::deleteLater();
}

void QMediaAvailabilityControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaAvailabilityControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaAvailabilityControl*>(ptr)->QMediaAvailabilityControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaAvailabilityControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaAvailabilityControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaAvailabilityControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaAvailabilityControl*>(ptr)->QMediaAvailabilityControl::event(static_cast<QEvent*>(e));
}

char QMediaAvailabilityControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaAvailabilityControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaAvailabilityControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaAvailabilityControl*>(ptr)->QMediaAvailabilityControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaAvailabilityControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaAvailabilityControl*>(ptr)->metaObject());
}

void* QMediaAvailabilityControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaAvailabilityControl*>(ptr)->QMediaAvailabilityControl::metaObject());
}

class MyQMediaBindableInterface: public QMediaBindableInterface
{
public:
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQMediaBindableInterface_MediaObject(const_cast<MyQMediaBindableInterface*>(this))); };
	bool setMediaObject(QMediaObject * object) { return callbackQMediaBindableInterface_SetMediaObject(this, object) != 0; };
	 ~MyQMediaBindableInterface() { callbackQMediaBindableInterface_DestroyQMediaBindableInterface(this); };
};

void* QMediaBindableInterface_MediaObject(void* ptr)
{
	return static_cast<QMediaBindableInterface*>(ptr)->mediaObject();
}

char QMediaBindableInterface_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QMediaBindableInterface*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

void QMediaBindableInterface_DestroyQMediaBindableInterface(void* ptr)
{
	static_cast<QMediaBindableInterface*>(ptr)->~QMediaBindableInterface();
}

void QMediaBindableInterface_DestroyQMediaBindableInterfaceDefault(void* ptr)
{

}

class MyQMediaContainerControl: public QMediaContainerControl
{
public:
	MyQMediaContainerControl(QObject *parent) : QMediaContainerControl(parent) {};
	QString containerDescription(const QString & format) const { QByteArray t785987 = format.toUtf8(); QtMultimedia_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };return QString(callbackQMediaContainerControl_ContainerDescription(const_cast<MyQMediaContainerControl*>(this), formatPacked)); };
	QString containerFormat() const { return QString(callbackQMediaContainerControl_ContainerFormat(const_cast<MyQMediaContainerControl*>(this))); };
	void setContainerFormat(const QString & format) { QByteArray t785987 = format.toUtf8(); QtMultimedia_PackedString formatPacked = { const_cast<char*>(t785987.prepend("WHITESPACE").constData()+10), t785987.size()-10 };callbackQMediaContainerControl_SetContainerFormat(this, formatPacked); };
	QStringList supportedContainers() const { return QString(callbackQMediaContainerControl_SupportedContainers(const_cast<MyQMediaContainerControl*>(this))).split("|", QString::SkipEmptyParts); };
	 ~MyQMediaContainerControl() { callbackQMediaContainerControl_DestroyQMediaContainerControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQMediaContainerControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaContainerControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaContainerControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaContainerControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaContainerControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaContainerControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaContainerControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaContainerControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaContainerControl_MetaObject(const_cast<MyQMediaContainerControl*>(this))); };
};

void* QMediaContainerControl_NewQMediaContainerControl(void* parent)
{
	return new MyQMediaContainerControl(static_cast<QObject*>(parent));
}

struct QtMultimedia_PackedString QMediaContainerControl_ContainerDescription(void* ptr, char* format)
{
	return ({ QByteArray t92eea4 = static_cast<QMediaContainerControl*>(ptr)->containerDescription(QString(format)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t92eea4.prepend("WHITESPACE").constData()+10), t92eea4.size()-10 }; });
}

struct QtMultimedia_PackedString QMediaContainerControl_ContainerFormat(void* ptr)
{
	return ({ QByteArray t46098d = static_cast<QMediaContainerControl*>(ptr)->containerFormat().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t46098d.prepend("WHITESPACE").constData()+10), t46098d.size()-10 }; });
}

void QMediaContainerControl_SetContainerFormat(void* ptr, char* format)
{
	static_cast<QMediaContainerControl*>(ptr)->setContainerFormat(QString(format));
}

struct QtMultimedia_PackedString QMediaContainerControl_SupportedContainers(void* ptr)
{
	return ({ QByteArray t36aec1 = static_cast<QMediaContainerControl*>(ptr)->supportedContainers().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t36aec1.prepend("WHITESPACE").constData()+10), t36aec1.size()-10 }; });
}

void QMediaContainerControl_DestroyQMediaContainerControl(void* ptr)
{
	static_cast<QMediaContainerControl*>(ptr)->~QMediaContainerControl();
}

void QMediaContainerControl_DestroyQMediaContainerControlDefault(void* ptr)
{

}

void QMediaContainerControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaContainerControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaContainerControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaContainerControl*>(ptr)->QMediaContainerControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaContainerControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaContainerControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaContainerControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaContainerControl*>(ptr)->QMediaContainerControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaContainerControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaContainerControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaContainerControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaContainerControl*>(ptr)->QMediaContainerControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaContainerControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaContainerControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaContainerControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaContainerControl*>(ptr)->QMediaContainerControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaContainerControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaContainerControl*>(ptr), "deleteLater");
}

void QMediaContainerControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaContainerControl*>(ptr)->QMediaContainerControl::deleteLater();
}

void QMediaContainerControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaContainerControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaContainerControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaContainerControl*>(ptr)->QMediaContainerControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaContainerControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaContainerControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaContainerControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaContainerControl*>(ptr)->QMediaContainerControl::event(static_cast<QEvent*>(e));
}

char QMediaContainerControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaContainerControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaContainerControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaContainerControl*>(ptr)->QMediaContainerControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaContainerControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaContainerControl*>(ptr)->metaObject());
}

void* QMediaContainerControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaContainerControl*>(ptr)->QMediaContainerControl::metaObject());
}

void* QMediaContent_NewQMediaContent()
{
	return new QMediaContent();
}

void* QMediaContent_NewQMediaContent7(void* playlist, void* contentUrl, char takeOwnership)
{
	return new QMediaContent(static_cast<QMediaPlaylist*>(playlist), *static_cast<QUrl*>(contentUrl), takeOwnership != 0);
}

void* QMediaContent_NewQMediaContent6(void* other)
{
	return new QMediaContent(*static_cast<QMediaContent*>(other));
}

void* QMediaContent_NewQMediaContent4(void* resource)
{
	return new QMediaContent(*static_cast<QMediaResource*>(resource));
}

void* QMediaContent_NewQMediaContent3(void* request)
{
	return new QMediaContent(*static_cast<QNetworkRequest*>(request));
}

void* QMediaContent_NewQMediaContent2(void* url)
{
	return new QMediaContent(*static_cast<QUrl*>(url));
}

void* QMediaContent_CanonicalRequest(void* ptr)
{
	return new QNetworkRequest(static_cast<QMediaContent*>(ptr)->canonicalRequest());
}

void* QMediaContent_CanonicalResource(void* ptr)
{
	return new QMediaResource(static_cast<QMediaContent*>(ptr)->canonicalResource());
}

void* QMediaContent_CanonicalUrl(void* ptr)
{
	return new QUrl(static_cast<QMediaContent*>(ptr)->canonicalUrl());
}

char QMediaContent_IsNull(void* ptr)
{
	return static_cast<QMediaContent*>(ptr)->isNull();
}

void* QMediaContent_Playlist(void* ptr)
{
	return static_cast<QMediaContent*>(ptr)->playlist();
}

void QMediaContent_DestroyQMediaContent(void* ptr)
{
	static_cast<QMediaContent*>(ptr)->~QMediaContent();
}

void* QMediaControl_NewQMediaControl(void* parent)
{
	return new QMediaControl(static_cast<QObject*>(parent));
}

void QMediaControl_DestroyQMediaControl(void* ptr)
{
	static_cast<QMediaControl*>(ptr)->~QMediaControl();
}

void QMediaControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaControl*>(ptr)->QMediaControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaControl*>(ptr)->QMediaControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaControl*>(ptr)->QMediaControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaControl*>(ptr)->QMediaControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaControl*>(ptr), "deleteLater");
}

void QMediaControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaControl*>(ptr)->QMediaControl::deleteLater();
}

void QMediaControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaControl*>(ptr)->QMediaControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaControl*>(ptr)->QMediaControl::event(static_cast<QEvent*>(e));
}

char QMediaControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaControl*>(ptr)->QMediaControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaControl*>(ptr)->metaObject());
}

void* QMediaControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaControl*>(ptr)->QMediaControl::metaObject());
}

class MyQMediaGaplessPlaybackControl: public QMediaGaplessPlaybackControl
{
public:
	MyQMediaGaplessPlaybackControl(QObject *parent) : QMediaGaplessPlaybackControl(parent) {};
	void Signal_AdvancedToNextMedia() { callbackQMediaGaplessPlaybackControl_AdvancedToNextMedia(this); };
	qreal crossfadeTime() const { return callbackQMediaGaplessPlaybackControl_CrossfadeTime(const_cast<MyQMediaGaplessPlaybackControl*>(this)); };
	void Signal_CrossfadeTimeChanged(qreal crossfadeTime) { callbackQMediaGaplessPlaybackControl_CrossfadeTimeChanged(this, crossfadeTime); };
	bool isCrossfadeSupported() const { return callbackQMediaGaplessPlaybackControl_IsCrossfadeSupported(const_cast<MyQMediaGaplessPlaybackControl*>(this)) != 0; };
	QMediaContent nextMedia() const { return *static_cast<QMediaContent*>(callbackQMediaGaplessPlaybackControl_NextMedia(const_cast<MyQMediaGaplessPlaybackControl*>(this))); };
	void Signal_NextMediaChanged(const QMediaContent & media) { callbackQMediaGaplessPlaybackControl_NextMediaChanged(this, const_cast<QMediaContent*>(&media)); };
	void setCrossfadeTime(qreal crossfadeTime) { callbackQMediaGaplessPlaybackControl_SetCrossfadeTime(this, crossfadeTime); };
	void setNextMedia(const QMediaContent & media) { callbackQMediaGaplessPlaybackControl_SetNextMedia(this, const_cast<QMediaContent*>(&media)); };
	 ~MyQMediaGaplessPlaybackControl() { callbackQMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQMediaGaplessPlaybackControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaGaplessPlaybackControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaGaplessPlaybackControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaGaplessPlaybackControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaGaplessPlaybackControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaGaplessPlaybackControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaGaplessPlaybackControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaGaplessPlaybackControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaGaplessPlaybackControl_MetaObject(const_cast<MyQMediaGaplessPlaybackControl*>(this))); };
};

void* QMediaGaplessPlaybackControl_NewQMediaGaplessPlaybackControl(void* parent)
{
	return new MyQMediaGaplessPlaybackControl(static_cast<QObject*>(parent));
}

void QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(void* ptr)
{
	QObject::connect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)()>(&QMediaGaplessPlaybackControl::advancedToNextMedia), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)()>(&MyQMediaGaplessPlaybackControl::Signal_AdvancedToNextMedia));
}

void QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(void* ptr)
{
	QObject::disconnect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)()>(&QMediaGaplessPlaybackControl::advancedToNextMedia), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)()>(&MyQMediaGaplessPlaybackControl::Signal_AdvancedToNextMedia));
}

void QMediaGaplessPlaybackControl_AdvancedToNextMedia(void* ptr)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->advancedToNextMedia();
}

double QMediaGaplessPlaybackControl_CrossfadeTime(void* ptr)
{
	return static_cast<QMediaGaplessPlaybackControl*>(ptr)->crossfadeTime();
}

void QMediaGaplessPlaybackControl_ConnectCrossfadeTimeChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)(qreal)>(&QMediaGaplessPlaybackControl::crossfadeTimeChanged), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)(qreal)>(&MyQMediaGaplessPlaybackControl::Signal_CrossfadeTimeChanged));
}

void QMediaGaplessPlaybackControl_DisconnectCrossfadeTimeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)(qreal)>(&QMediaGaplessPlaybackControl::crossfadeTimeChanged), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)(qreal)>(&MyQMediaGaplessPlaybackControl::Signal_CrossfadeTimeChanged));
}

void QMediaGaplessPlaybackControl_CrossfadeTimeChanged(void* ptr, double crossfadeTime)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->crossfadeTimeChanged(crossfadeTime);
}

char QMediaGaplessPlaybackControl_IsCrossfadeSupported(void* ptr)
{
	return static_cast<QMediaGaplessPlaybackControl*>(ptr)->isCrossfadeSupported();
}

void* QMediaGaplessPlaybackControl_NextMedia(void* ptr)
{
	return new QMediaContent(static_cast<QMediaGaplessPlaybackControl*>(ptr)->nextMedia());
}

void QMediaGaplessPlaybackControl_ConnectNextMediaChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)(const QMediaContent &)>(&QMediaGaplessPlaybackControl::nextMediaChanged), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)(const QMediaContent &)>(&MyQMediaGaplessPlaybackControl::Signal_NextMediaChanged));
}

void QMediaGaplessPlaybackControl_DisconnectNextMediaChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)(const QMediaContent &)>(&QMediaGaplessPlaybackControl::nextMediaChanged), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)(const QMediaContent &)>(&MyQMediaGaplessPlaybackControl::Signal_NextMediaChanged));
}

void QMediaGaplessPlaybackControl_NextMediaChanged(void* ptr, void* media)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->nextMediaChanged(*static_cast<QMediaContent*>(media));
}

void QMediaGaplessPlaybackControl_SetCrossfadeTime(void* ptr, double crossfadeTime)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->setCrossfadeTime(crossfadeTime);
}

void QMediaGaplessPlaybackControl_SetNextMedia(void* ptr, void* media)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->setNextMedia(*static_cast<QMediaContent*>(media));
}

void QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(void* ptr)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->~QMediaGaplessPlaybackControl();
}

void QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControlDefault(void* ptr)
{

}

void QMediaGaplessPlaybackControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaGaplessPlaybackControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->QMediaGaplessPlaybackControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaGaplessPlaybackControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaGaplessPlaybackControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->QMediaGaplessPlaybackControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaGaplessPlaybackControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaGaplessPlaybackControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->QMediaGaplessPlaybackControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaGaplessPlaybackControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaGaplessPlaybackControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->QMediaGaplessPlaybackControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaGaplessPlaybackControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaGaplessPlaybackControl*>(ptr), "deleteLater");
}

void QMediaGaplessPlaybackControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->QMediaGaplessPlaybackControl::deleteLater();
}

void QMediaGaplessPlaybackControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaGaplessPlaybackControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->QMediaGaplessPlaybackControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaGaplessPlaybackControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaGaplessPlaybackControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaGaplessPlaybackControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaGaplessPlaybackControl*>(ptr)->QMediaGaplessPlaybackControl::event(static_cast<QEvent*>(e));
}

char QMediaGaplessPlaybackControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaGaplessPlaybackControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaGaplessPlaybackControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaGaplessPlaybackControl*>(ptr)->QMediaGaplessPlaybackControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaGaplessPlaybackControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaGaplessPlaybackControl*>(ptr)->metaObject());
}

void* QMediaGaplessPlaybackControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaGaplessPlaybackControl*>(ptr)->QMediaGaplessPlaybackControl::metaObject());
}

class MyQMediaNetworkAccessControl: public QMediaNetworkAccessControl
{
public:
	void Signal_ConfigurationChanged(const QNetworkConfiguration & configuration) { callbackQMediaNetworkAccessControl_ConfigurationChanged(this, const_cast<QNetworkConfiguration*>(&configuration)); };
	QNetworkConfiguration currentConfiguration() const { return *static_cast<QNetworkConfiguration*>(callbackQMediaNetworkAccessControl_CurrentConfiguration(const_cast<MyQMediaNetworkAccessControl*>(this))); };
	 ~MyQMediaNetworkAccessControl() { callbackQMediaNetworkAccessControl_DestroyQMediaNetworkAccessControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQMediaNetworkAccessControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaNetworkAccessControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaNetworkAccessControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaNetworkAccessControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaNetworkAccessControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaNetworkAccessControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaNetworkAccessControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaNetworkAccessControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaNetworkAccessControl_MetaObject(const_cast<MyQMediaNetworkAccessControl*>(this))); };
};

void QMediaNetworkAccessControl_ConnectConfigurationChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaNetworkAccessControl*>(ptr), static_cast<void (QMediaNetworkAccessControl::*)(const QNetworkConfiguration &)>(&QMediaNetworkAccessControl::configurationChanged), static_cast<MyQMediaNetworkAccessControl*>(ptr), static_cast<void (MyQMediaNetworkAccessControl::*)(const QNetworkConfiguration &)>(&MyQMediaNetworkAccessControl::Signal_ConfigurationChanged));
}

void QMediaNetworkAccessControl_DisconnectConfigurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaNetworkAccessControl*>(ptr), static_cast<void (QMediaNetworkAccessControl::*)(const QNetworkConfiguration &)>(&QMediaNetworkAccessControl::configurationChanged), static_cast<MyQMediaNetworkAccessControl*>(ptr), static_cast<void (MyQMediaNetworkAccessControl::*)(const QNetworkConfiguration &)>(&MyQMediaNetworkAccessControl::Signal_ConfigurationChanged));
}

void QMediaNetworkAccessControl_ConfigurationChanged(void* ptr, void* configuration)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->configurationChanged(*static_cast<QNetworkConfiguration*>(configuration));
}

void* QMediaNetworkAccessControl_CurrentConfiguration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QMediaNetworkAccessControl*>(ptr)->currentConfiguration());
}

void QMediaNetworkAccessControl_DestroyQMediaNetworkAccessControl(void* ptr)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->~QMediaNetworkAccessControl();
}

void QMediaNetworkAccessControl_DestroyQMediaNetworkAccessControlDefault(void* ptr)
{

}

void QMediaNetworkAccessControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaNetworkAccessControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->QMediaNetworkAccessControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaNetworkAccessControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaNetworkAccessControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->QMediaNetworkAccessControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaNetworkAccessControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaNetworkAccessControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->QMediaNetworkAccessControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaNetworkAccessControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaNetworkAccessControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->QMediaNetworkAccessControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaNetworkAccessControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaNetworkAccessControl*>(ptr), "deleteLater");
}

void QMediaNetworkAccessControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->QMediaNetworkAccessControl::deleteLater();
}

void QMediaNetworkAccessControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaNetworkAccessControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaNetworkAccessControl*>(ptr)->QMediaNetworkAccessControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaNetworkAccessControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaNetworkAccessControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaNetworkAccessControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaNetworkAccessControl*>(ptr)->QMediaNetworkAccessControl::event(static_cast<QEvent*>(e));
}

char QMediaNetworkAccessControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaNetworkAccessControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaNetworkAccessControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaNetworkAccessControl*>(ptr)->QMediaNetworkAccessControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaNetworkAccessControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaNetworkAccessControl*>(ptr)->metaObject());
}

void* QMediaNetworkAccessControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaNetworkAccessControl*>(ptr)->QMediaNetworkAccessControl::metaObject());
}

class MyQMediaObject: public QMediaObject
{
public:
	MyQMediaObject(QObject *parent, QMediaService *service) : QMediaObject(parent, service) {};
	QMultimedia::AvailabilityStatus availability() const { return static_cast<QMultimedia::AvailabilityStatus>(callbackQMediaObject_Availability(const_cast<MyQMediaObject*>(this))); };
	void Signal_AvailabilityChanged2(QMultimedia::AvailabilityStatus availability) { callbackQMediaObject_AvailabilityChanged2(this, availability); };
	void Signal_AvailabilityChanged(bool available) { callbackQMediaObject_AvailabilityChanged(this, available); };
	bool bind(QObject * object) { return callbackQMediaObject_Bind(this, object) != 0; };
	bool isAvailable() const { return callbackQMediaObject_IsAvailable(const_cast<MyQMediaObject*>(this)) != 0; };
	void Signal_MetaDataAvailableChanged(bool available) { callbackQMediaObject_MetaDataAvailableChanged(this, available); };
	void Signal_MetaDataChanged() { callbackQMediaObject_MetaDataChanged(this); };
	void Signal_MetaDataChanged2(const QString & key, const QVariant & value) { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };callbackQMediaObject_MetaDataChanged2(this, keyPacked, const_cast<QVariant*>(&value)); };
	void Signal_NotifyIntervalChanged(int milliseconds) { callbackQMediaObject_NotifyIntervalChanged(this, milliseconds); };
	QMediaService * service() const { return static_cast<QMediaService*>(callbackQMediaObject_Service(const_cast<MyQMediaObject*>(this))); };
	void unbind(QObject * object) { callbackQMediaObject_Unbind(this, object); };
	void timerEvent(QTimerEvent * event) { callbackQMediaObject_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaObject_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaObject_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaObject_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaObject_MetaObject(const_cast<MyQMediaObject*>(this))); };
};

int QMediaObject_NotifyInterval(void* ptr)
{
	return static_cast<QMediaObject*>(ptr)->notifyInterval();
}

void QMediaObject_SetNotifyInterval(void* ptr, int milliSeconds)
{
	static_cast<QMediaObject*>(ptr)->setNotifyInterval(milliSeconds);
}

void* QMediaObject_NewQMediaObject(void* parent, void* service)
{
	return new MyQMediaObject(static_cast<QObject*>(parent), static_cast<QMediaService*>(service));
}

void QMediaObject_AddPropertyWatch(void* ptr, void* name)
{
	static_cast<QMediaObject*>(ptr)->addPropertyWatch(*static_cast<QByteArray*>(name));
}

long long QMediaObject_Availability(void* ptr)
{
	return static_cast<QMediaObject*>(ptr)->availability();
}

long long QMediaObject_AvailabilityDefault(void* ptr)
{
	return static_cast<QMediaObject*>(ptr)->QMediaObject::availability();
}

void QMediaObject_ConnectAvailabilityChanged2(void* ptr)
{
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(QMultimedia::AvailabilityStatus)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(QMultimedia::AvailabilityStatus)>(&MyQMediaObject::Signal_AvailabilityChanged2));
}

void QMediaObject_DisconnectAvailabilityChanged2(void* ptr)
{
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(QMultimedia::AvailabilityStatus)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(QMultimedia::AvailabilityStatus)>(&MyQMediaObject::Signal_AvailabilityChanged2));
}

void QMediaObject_AvailabilityChanged2(void* ptr, long long availability)
{
	static_cast<QMediaObject*>(ptr)->availabilityChanged(static_cast<QMultimedia::AvailabilityStatus>(availability));
}

void QMediaObject_ConnectAvailabilityChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_AvailabilityChanged));
}

void QMediaObject_DisconnectAvailabilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_AvailabilityChanged));
}

void QMediaObject_AvailabilityChanged(void* ptr, char available)
{
	static_cast<QMediaObject*>(ptr)->availabilityChanged(available != 0);
}

struct QtMultimedia_PackedString QMediaObject_AvailableMetaData(void* ptr)
{
	return ({ QByteArray t26e55a = static_cast<QMediaObject*>(ptr)->availableMetaData().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t26e55a.prepend("WHITESPACE").constData()+10), t26e55a.size()-10 }; });
}

char QMediaObject_Bind(void* ptr, void* object)
{
	return static_cast<QMediaObject*>(ptr)->bind(static_cast<QObject*>(object));
}

char QMediaObject_BindDefault(void* ptr, void* object)
{
	return static_cast<QMediaObject*>(ptr)->QMediaObject::bind(static_cast<QObject*>(object));
}

char QMediaObject_IsAvailable(void* ptr)
{
	return static_cast<QMediaObject*>(ptr)->isAvailable();
}

char QMediaObject_IsAvailableDefault(void* ptr)
{
	return static_cast<QMediaObject*>(ptr)->QMediaObject::isAvailable();
}

char QMediaObject_IsMetaDataAvailable(void* ptr)
{
	return static_cast<QMediaObject*>(ptr)->isMetaDataAvailable();
}

void* QMediaObject_MetaData(void* ptr, char* key)
{
	return new QVariant(static_cast<QMediaObject*>(ptr)->metaData(QString(key)));
}

void QMediaObject_ConnectMetaDataAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::metaDataAvailableChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_MetaDataAvailableChanged));
}

void QMediaObject_DisconnectMetaDataAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::metaDataAvailableChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_MetaDataAvailableChanged));
}

void QMediaObject_MetaDataAvailableChanged(void* ptr, char available)
{
	static_cast<QMediaObject*>(ptr)->metaDataAvailableChanged(available != 0);
}

void QMediaObject_ConnectMetaDataChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)()>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)()>(&MyQMediaObject::Signal_MetaDataChanged));
}

void QMediaObject_DisconnectMetaDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)()>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)()>(&MyQMediaObject::Signal_MetaDataChanged));
}

void QMediaObject_MetaDataChanged(void* ptr)
{
	static_cast<QMediaObject*>(ptr)->metaDataChanged();
}

void QMediaObject_ConnectMetaDataChanged2(void* ptr)
{
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(const QString &, const QVariant &)>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(const QString &, const QVariant &)>(&MyQMediaObject::Signal_MetaDataChanged2));
}

void QMediaObject_DisconnectMetaDataChanged2(void* ptr)
{
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(const QString &, const QVariant &)>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(const QString &, const QVariant &)>(&MyQMediaObject::Signal_MetaDataChanged2));
}

void QMediaObject_MetaDataChanged2(void* ptr, char* key, void* value)
{
	static_cast<QMediaObject*>(ptr)->metaDataChanged(QString(key), *static_cast<QVariant*>(value));
}

void QMediaObject_ConnectNotifyIntervalChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(int)>(&QMediaObject::notifyIntervalChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(int)>(&MyQMediaObject::Signal_NotifyIntervalChanged));
}

void QMediaObject_DisconnectNotifyIntervalChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(int)>(&QMediaObject::notifyIntervalChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(int)>(&MyQMediaObject::Signal_NotifyIntervalChanged));
}

void QMediaObject_NotifyIntervalChanged(void* ptr, int milliseconds)
{
	static_cast<QMediaObject*>(ptr)->notifyIntervalChanged(milliseconds);
}

void QMediaObject_RemovePropertyWatch(void* ptr, void* name)
{
	static_cast<QMediaObject*>(ptr)->removePropertyWatch(*static_cast<QByteArray*>(name));
}

void* QMediaObject_Service(void* ptr)
{
	return static_cast<QMediaObject*>(ptr)->service();
}

void* QMediaObject_ServiceDefault(void* ptr)
{
	return static_cast<QMediaObject*>(ptr)->QMediaObject::service();
}

void QMediaObject_Unbind(void* ptr, void* object)
{
	static_cast<QMediaObject*>(ptr)->unbind(static_cast<QObject*>(object));
}

void QMediaObject_UnbindDefault(void* ptr, void* object)
{
	static_cast<QMediaObject*>(ptr)->QMediaObject::unbind(static_cast<QObject*>(object));
}

void QMediaObject_DestroyQMediaObject(void* ptr)
{
	static_cast<QMediaObject*>(ptr)->~QMediaObject();
}

void QMediaObject_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaObject_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaObject*>(ptr)->QMediaObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaObject_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaObject_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaObject*>(ptr)->QMediaObject::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaObject_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaObject*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaObject_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaObject*>(ptr)->QMediaObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaObject_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaObject_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaObject*>(ptr)->QMediaObject::customEvent(static_cast<QEvent*>(event));
}

void QMediaObject_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaObject*>(ptr), "deleteLater");
}

void QMediaObject_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaObject*>(ptr)->QMediaObject::deleteLater();
}

void QMediaObject_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaObject*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaObject_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaObject*>(ptr)->QMediaObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaObject_Event(void* ptr, void* e)
{
	return static_cast<QMediaObject*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaObject_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaObject*>(ptr)->QMediaObject::event(static_cast<QEvent*>(e));
}

char QMediaObject_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaObject_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaObject*>(ptr)->QMediaObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaObject_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaObject*>(ptr)->metaObject());
}

void* QMediaObject_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaObject*>(ptr)->QMediaObject::metaObject());
}

class MyQMediaPlayer: public QMediaPlayer
{
public:
	MyQMediaPlayer(QObject *parent, Flags flags) : QMediaPlayer(parent, flags) {};
	void setMuted(bool muted) { callbackQMediaPlayer_SetMuted(this, muted); };
	void setPlaybackRate(qreal rate) { callbackQMediaPlayer_SetPlaybackRate(this, rate); };
	void setPlaylist(QMediaPlaylist * playlist) { callbackQMediaPlayer_SetPlaylist(this, playlist); };
	void setPosition(qint64 position) { callbackQMediaPlayer_SetPosition(this, position); };
	void setVolume(int volume) { callbackQMediaPlayer_SetVolume(this, volume); };
	void Signal_AudioAvailableChanged(bool available) { callbackQMediaPlayer_AudioAvailableChanged(this, available); };
	void Signal_AudioRoleChanged(QAudio::Role role) { callbackQMediaPlayer_AudioRoleChanged(this, role); };
	QMultimedia::AvailabilityStatus availability() const { return static_cast<QMultimedia::AvailabilityStatus>(callbackQMediaPlayer_Availability(const_cast<MyQMediaPlayer*>(this))); };
	void Signal_BufferStatusChanged(int percentFilled) { callbackQMediaPlayer_BufferStatusChanged(this, percentFilled); };
	void Signal_CurrentMediaChanged(const QMediaContent & media) { callbackQMediaPlayer_CurrentMediaChanged(this, const_cast<QMediaContent*>(&media)); };
	void Signal_DurationChanged(qint64 duration) { callbackQMediaPlayer_DurationChanged(this, duration); };
	void Signal_Error2(QMediaPlayer::Error error) { callbackQMediaPlayer_Error2(this, error); };
	void Signal_MediaChanged(const QMediaContent & media) { callbackQMediaPlayer_MediaChanged(this, const_cast<QMediaContent*>(&media)); };
	void Signal_MediaStatusChanged(QMediaPlayer::MediaStatus status) { callbackQMediaPlayer_MediaStatusChanged(this, status); };
	void Signal_MutedChanged(bool muted) { callbackQMediaPlayer_MutedChanged(this, muted); };
	void Signal_NetworkConfigurationChanged(const QNetworkConfiguration & configuration) { callbackQMediaPlayer_NetworkConfigurationChanged(this, const_cast<QNetworkConfiguration*>(&configuration)); };
	void pause() { callbackQMediaPlayer_Pause(this); };
	void play() { callbackQMediaPlayer_Play(this); };
	void Signal_PlaybackRateChanged(qreal rate) { callbackQMediaPlayer_PlaybackRateChanged(this, rate); };
	void Signal_PositionChanged(qint64 position) { callbackQMediaPlayer_PositionChanged(this, position); };
	void Signal_SeekableChanged(bool seekable) { callbackQMediaPlayer_SeekableChanged(this, seekable); };
	void setMedia(const QMediaContent & media, QIODevice * stream) { callbackQMediaPlayer_SetMedia(this, const_cast<QMediaContent*>(&media), stream); };
	void Signal_StateChanged(QMediaPlayer::State state) { callbackQMediaPlayer_StateChanged(this, state); };
	void stop() { callbackQMediaPlayer_Stop(this); };
	void Signal_VideoAvailableChanged(bool videoAvailable) { callbackQMediaPlayer_VideoAvailableChanged(this, videoAvailable); };
	void Signal_VolumeChanged(int volume) { callbackQMediaPlayer_VolumeChanged(this, volume); };
	bool bind(QObject * object) { return callbackQMediaPlayer_Bind(this, object) != 0; };
	bool isAvailable() const { return callbackQMediaPlayer_IsAvailable(const_cast<MyQMediaPlayer*>(this)) != 0; };
	QMediaService * service() const { return static_cast<QMediaService*>(callbackQMediaPlayer_Service(const_cast<MyQMediaPlayer*>(this))); };
	void unbind(QObject * object) { callbackQMediaPlayer_Unbind(this, object); };
	void timerEvent(QTimerEvent * event) { callbackQMediaPlayer_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaPlayer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaPlayer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaPlayer_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaPlayer_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaPlayer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaPlayer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaPlayer_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaPlayer_MetaObject(const_cast<MyQMediaPlayer*>(this))); };
};

long long QMediaPlayer_AudioRole(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->audioRole();
}

int QMediaPlayer_BufferStatus(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->bufferStatus();
}

void* QMediaPlayer_CurrentMedia(void* ptr)
{
	return new QMediaContent(static_cast<QMediaPlayer*>(ptr)->currentMedia());
}

long long QMediaPlayer_Duration(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->duration();
}

struct QtMultimedia_PackedString QMediaPlayer_ErrorString(void* ptr)
{
	return ({ QByteArray t55328b = static_cast<QMediaPlayer*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t55328b.prepend("WHITESPACE").constData()+10), t55328b.size()-10 }; });
}

char QMediaPlayer_IsAudioAvailable(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->isAudioAvailable();
}

char QMediaPlayer_IsMuted(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->isMuted();
}

char QMediaPlayer_IsSeekable(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->isSeekable();
}

char QMediaPlayer_IsVideoAvailable(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->isVideoAvailable();
}

void* QMediaPlayer_Media(void* ptr)
{
	return new QMediaContent(static_cast<QMediaPlayer*>(ptr)->media());
}

long long QMediaPlayer_MediaStatus(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->mediaStatus();
}

double QMediaPlayer_PlaybackRate(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->playbackRate();
}

void* QMediaPlayer_Playlist(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->playlist();
}

long long QMediaPlayer_Position(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->position();
}

void QMediaPlayer_SetAudioRole(void* ptr, long long audioRole)
{
	static_cast<QMediaPlayer*>(ptr)->setAudioRole(static_cast<QAudio::Role>(audioRole));
}

void QMediaPlayer_SetMuted(void* ptr, char muted)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

void QMediaPlayer_SetPlaybackRate(void* ptr, double rate)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setPlaybackRate", Q_ARG(qreal, rate));
}

void QMediaPlayer_SetPlaylist(void* ptr, void* playlist)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setPlaylist", Q_ARG(QMediaPlaylist*, static_cast<QMediaPlaylist*>(playlist)));
}

void QMediaPlayer_SetPosition(void* ptr, long long position)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setPosition", Q_ARG(qint64, position));
}

void QMediaPlayer_SetVideoOutput2(void* ptr, void* output)
{
	static_cast<QMediaPlayer*>(ptr)->setVideoOutput(static_cast<QGraphicsVideoItem*>(output));
}

void QMediaPlayer_SetVideoOutput(void* ptr, void* output)
{
	static_cast<QMediaPlayer*>(ptr)->setVideoOutput(static_cast<QVideoWidget*>(output));
}

void QMediaPlayer_SetVolume(void* ptr, int volume)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setVolume", Q_ARG(int, volume));
}

long long QMediaPlayer_State(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->state();
}

int QMediaPlayer_Volume(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->volume();
}

void* QMediaPlayer_NewQMediaPlayer(void* parent, long long flags)
{
	return new MyQMediaPlayer(static_cast<QObject*>(parent), static_cast<QMediaPlayer::Flag>(flags));
}

void QMediaPlayer_ConnectAudioAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::audioAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_AudioAvailableChanged));
}

void QMediaPlayer_DisconnectAudioAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::audioAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_AudioAvailableChanged));
}

void QMediaPlayer_AudioAvailableChanged(void* ptr, char available)
{
	static_cast<QMediaPlayer*>(ptr)->audioAvailableChanged(available != 0);
}

void QMediaPlayer_ConnectAudioRoleChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QAudio::Role)>(&QMediaPlayer::audioRoleChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QAudio::Role)>(&MyQMediaPlayer::Signal_AudioRoleChanged));
}

void QMediaPlayer_DisconnectAudioRoleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QAudio::Role)>(&QMediaPlayer::audioRoleChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QAudio::Role)>(&MyQMediaPlayer::Signal_AudioRoleChanged));
}

void QMediaPlayer_AudioRoleChanged(void* ptr, long long role)
{
	static_cast<QMediaPlayer*>(ptr)->audioRoleChanged(static_cast<QAudio::Role>(role));
}

long long QMediaPlayer_Availability(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->availability();
}

long long QMediaPlayer_AvailabilityDefault(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::availability();
}

void QMediaPlayer_ConnectBufferStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::bufferStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_BufferStatusChanged));
}

void QMediaPlayer_DisconnectBufferStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::bufferStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_BufferStatusChanged));
}

void QMediaPlayer_BufferStatusChanged(void* ptr, int percentFilled)
{
	static_cast<QMediaPlayer*>(ptr)->bufferStatusChanged(percentFilled);
}

void QMediaPlayer_ConnectCurrentMediaChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(const QMediaContent &)>(&QMediaPlayer::currentMediaChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(const QMediaContent &)>(&MyQMediaPlayer::Signal_CurrentMediaChanged));
}

void QMediaPlayer_DisconnectCurrentMediaChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(const QMediaContent &)>(&QMediaPlayer::currentMediaChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(const QMediaContent &)>(&MyQMediaPlayer::Signal_CurrentMediaChanged));
}

void QMediaPlayer_CurrentMediaChanged(void* ptr, void* media)
{
	static_cast<QMediaPlayer*>(ptr)->currentMediaChanged(*static_cast<QMediaContent*>(media));
}

void* QMediaPlayer_CurrentNetworkConfiguration(void* ptr)
{
	return new QNetworkConfiguration(static_cast<QMediaPlayer*>(ptr)->currentNetworkConfiguration());
}

void QMediaPlayer_ConnectDurationChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(qint64)>(&QMediaPlayer::durationChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(qint64)>(&MyQMediaPlayer::Signal_DurationChanged));
}

void QMediaPlayer_DisconnectDurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(qint64)>(&QMediaPlayer::durationChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(qint64)>(&MyQMediaPlayer::Signal_DurationChanged));
}

void QMediaPlayer_DurationChanged(void* ptr, long long duration)
{
	static_cast<QMediaPlayer*>(ptr)->durationChanged(duration);
}

void QMediaPlayer_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::Error)>(&QMediaPlayer::error), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::Error)>(&MyQMediaPlayer::Signal_Error2));
}

void QMediaPlayer_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::Error)>(&QMediaPlayer::error), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::Error)>(&MyQMediaPlayer::Signal_Error2));
}

void QMediaPlayer_Error2(void* ptr, long long error)
{
	static_cast<QMediaPlayer*>(ptr)->error(static_cast<QMediaPlayer::Error>(error));
}

long long QMediaPlayer_Error(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->error();
}

long long QMediaPlayer_QMediaPlayer_HasSupport(char* mimeType, char* codecs, long long flags)
{
	return QMediaPlayer::hasSupport(QString(mimeType), QString(codecs).split("|", QString::SkipEmptyParts), static_cast<QMediaPlayer::Flag>(flags));
}

void QMediaPlayer_ConnectMediaChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(const QMediaContent &)>(&QMediaPlayer::mediaChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(const QMediaContent &)>(&MyQMediaPlayer::Signal_MediaChanged));
}

void QMediaPlayer_DisconnectMediaChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(const QMediaContent &)>(&QMediaPlayer::mediaChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(const QMediaContent &)>(&MyQMediaPlayer::Signal_MediaChanged));
}

void QMediaPlayer_MediaChanged(void* ptr, void* media)
{
	static_cast<QMediaPlayer*>(ptr)->mediaChanged(*static_cast<QMediaContent*>(media));
}

void QMediaPlayer_ConnectMediaStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayer::mediaStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayer::Signal_MediaStatusChanged));
}

void QMediaPlayer_DisconnectMediaStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayer::mediaStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayer::Signal_MediaStatusChanged));
}

void QMediaPlayer_MediaStatusChanged(void* ptr, long long status)
{
	static_cast<QMediaPlayer*>(ptr)->mediaStatusChanged(static_cast<QMediaPlayer::MediaStatus>(status));
}

void* QMediaPlayer_MediaStream(void* ptr)
{
	return const_cast<QIODevice*>(static_cast<QMediaPlayer*>(ptr)->mediaStream());
}

void QMediaPlayer_ConnectMutedChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::mutedChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_MutedChanged));
}

void QMediaPlayer_DisconnectMutedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::mutedChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_MutedChanged));
}

void QMediaPlayer_MutedChanged(void* ptr, char muted)
{
	static_cast<QMediaPlayer*>(ptr)->mutedChanged(muted != 0);
}

void QMediaPlayer_ConnectNetworkConfigurationChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(const QNetworkConfiguration &)>(&QMediaPlayer::networkConfigurationChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(const QNetworkConfiguration &)>(&MyQMediaPlayer::Signal_NetworkConfigurationChanged));
}

void QMediaPlayer_DisconnectNetworkConfigurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(const QNetworkConfiguration &)>(&QMediaPlayer::networkConfigurationChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(const QNetworkConfiguration &)>(&MyQMediaPlayer::Signal_NetworkConfigurationChanged));
}

void QMediaPlayer_NetworkConfigurationChanged(void* ptr, void* configuration)
{
	static_cast<QMediaPlayer*>(ptr)->networkConfigurationChanged(*static_cast<QNetworkConfiguration*>(configuration));
}

void QMediaPlayer_Pause(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "pause");
}

void QMediaPlayer_Play(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "play");
}

void QMediaPlayer_ConnectPlaybackRateChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(qreal)>(&QMediaPlayer::playbackRateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(qreal)>(&MyQMediaPlayer::Signal_PlaybackRateChanged));
}

void QMediaPlayer_DisconnectPlaybackRateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(qreal)>(&QMediaPlayer::playbackRateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(qreal)>(&MyQMediaPlayer::Signal_PlaybackRateChanged));
}

void QMediaPlayer_PlaybackRateChanged(void* ptr, double rate)
{
	static_cast<QMediaPlayer*>(ptr)->playbackRateChanged(rate);
}

void QMediaPlayer_ConnectPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(qint64)>(&QMediaPlayer::positionChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(qint64)>(&MyQMediaPlayer::Signal_PositionChanged));
}

void QMediaPlayer_DisconnectPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(qint64)>(&QMediaPlayer::positionChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(qint64)>(&MyQMediaPlayer::Signal_PositionChanged));
}

void QMediaPlayer_PositionChanged(void* ptr, long long position)
{
	static_cast<QMediaPlayer*>(ptr)->positionChanged(position);
}

void QMediaPlayer_ConnectSeekableChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::seekableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_SeekableChanged));
}

void QMediaPlayer_DisconnectSeekableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::seekableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_SeekableChanged));
}

void QMediaPlayer_SeekableChanged(void* ptr, char seekable)
{
	static_cast<QMediaPlayer*>(ptr)->seekableChanged(seekable != 0);
}

void QMediaPlayer_SetMedia(void* ptr, void* media, void* stream)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setMedia", Q_ARG(QMediaContent, *static_cast<QMediaContent*>(media)), Q_ARG(QIODevice*, static_cast<QIODevice*>(stream)));
}

void QMediaPlayer_SetVideoOutput3(void* ptr, void* surface)
{
	static_cast<QMediaPlayer*>(ptr)->setVideoOutput(static_cast<QAbstractVideoSurface*>(surface));
}

void QMediaPlayer_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::State)>(&QMediaPlayer::stateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::State)>(&MyQMediaPlayer::Signal_StateChanged));
}

void QMediaPlayer_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::State)>(&QMediaPlayer::stateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::State)>(&MyQMediaPlayer::Signal_StateChanged));
}

void QMediaPlayer_StateChanged(void* ptr, long long state)
{
	static_cast<QMediaPlayer*>(ptr)->stateChanged(static_cast<QMediaPlayer::State>(state));
}

void QMediaPlayer_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "stop");
}

void QMediaPlayer_ConnectVideoAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::videoAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_VideoAvailableChanged));
}

void QMediaPlayer_DisconnectVideoAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::videoAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_VideoAvailableChanged));
}

void QMediaPlayer_VideoAvailableChanged(void* ptr, char videoAvailable)
{
	static_cast<QMediaPlayer*>(ptr)->videoAvailableChanged(videoAvailable != 0);
}

void QMediaPlayer_ConnectVolumeChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::volumeChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_VolumeChanged));
}

void QMediaPlayer_DisconnectVolumeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::volumeChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_VolumeChanged));
}

void QMediaPlayer_VolumeChanged(void* ptr, int volume)
{
	static_cast<QMediaPlayer*>(ptr)->volumeChanged(volume);
}

void QMediaPlayer_DestroyQMediaPlayer(void* ptr)
{
	static_cast<QMediaPlayer*>(ptr)->~QMediaPlayer();
}

char QMediaPlayer_Bind(void* ptr, void* object)
{
	return static_cast<QMediaPlayer*>(ptr)->bind(static_cast<QObject*>(object));
}

char QMediaPlayer_BindDefault(void* ptr, void* object)
{
	return static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::bind(static_cast<QObject*>(object));
}

char QMediaPlayer_IsAvailable(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->isAvailable();
}

char QMediaPlayer_IsAvailableDefault(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::isAvailable();
}

void* QMediaPlayer_Service(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->service();
}

void* QMediaPlayer_ServiceDefault(void* ptr)
{
	return static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::service();
}

void QMediaPlayer_Unbind(void* ptr, void* object)
{
	static_cast<QMediaPlayer*>(ptr)->unbind(static_cast<QObject*>(object));
}

void QMediaPlayer_UnbindDefault(void* ptr, void* object)
{
	static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::unbind(static_cast<QObject*>(object));
}

void QMediaPlayer_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaPlayer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaPlayer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaPlayer_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaPlayer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaPlayer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaPlayer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaPlayer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaPlayer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaPlayer_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaPlayer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaPlayer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::customEvent(static_cast<QEvent*>(event));
}

void QMediaPlayer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "deleteLater");
}

void QMediaPlayer_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::deleteLater();
}

void QMediaPlayer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaPlayer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaPlayer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaPlayer_Event(void* ptr, void* e)
{
	return static_cast<QMediaPlayer*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaPlayer_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::event(static_cast<QEvent*>(e));
}

char QMediaPlayer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaPlayer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaPlayer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaPlayer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaPlayer*>(ptr)->metaObject());
}

void* QMediaPlayer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaPlayer*>(ptr)->QMediaPlayer::metaObject());
}

class MyQMediaPlayerControl: public QMediaPlayerControl
{
public:
	MyQMediaPlayerControl(QObject *parent) : QMediaPlayerControl(parent) {};
	void Signal_AudioAvailableChanged(bool audio) { callbackQMediaPlayerControl_AudioAvailableChanged(this, audio); };
	QMediaTimeRange availablePlaybackRanges() const { return *static_cast<QMediaTimeRange*>(callbackQMediaPlayerControl_AvailablePlaybackRanges(const_cast<MyQMediaPlayerControl*>(this))); };
	void Signal_AvailablePlaybackRangesChanged(const QMediaTimeRange & ranges) { callbackQMediaPlayerControl_AvailablePlaybackRangesChanged(this, const_cast<QMediaTimeRange*>(&ranges)); };
	int bufferStatus() const { return callbackQMediaPlayerControl_BufferStatus(const_cast<MyQMediaPlayerControl*>(this)); };
	void Signal_BufferStatusChanged(int progress) { callbackQMediaPlayerControl_BufferStatusChanged(this, progress); };
	qint64 duration() const { return callbackQMediaPlayerControl_Duration(const_cast<MyQMediaPlayerControl*>(this)); };
	void Signal_DurationChanged(qint64 duration) { callbackQMediaPlayerControl_DurationChanged(this, duration); };
	void Signal_Error(int error, const QString & errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtMultimedia_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQMediaPlayerControl_Error(this, error, errorStringPacked); };
	bool isAudioAvailable() const { return callbackQMediaPlayerControl_IsAudioAvailable(const_cast<MyQMediaPlayerControl*>(this)) != 0; };
	bool isMuted() const { return callbackQMediaPlayerControl_IsMuted(const_cast<MyQMediaPlayerControl*>(this)) != 0; };
	bool isSeekable() const { return callbackQMediaPlayerControl_IsSeekable(const_cast<MyQMediaPlayerControl*>(this)) != 0; };
	bool isVideoAvailable() const { return callbackQMediaPlayerControl_IsVideoAvailable(const_cast<MyQMediaPlayerControl*>(this)) != 0; };
	QMediaContent media() const { return *static_cast<QMediaContent*>(callbackQMediaPlayerControl_Media(const_cast<MyQMediaPlayerControl*>(this))); };
	void Signal_MediaChanged(const QMediaContent & content) { callbackQMediaPlayerControl_MediaChanged(this, const_cast<QMediaContent*>(&content)); };
	QMediaPlayer::MediaStatus mediaStatus() const { return static_cast<QMediaPlayer::MediaStatus>(callbackQMediaPlayerControl_MediaStatus(const_cast<MyQMediaPlayerControl*>(this))); };
	void Signal_MediaStatusChanged(QMediaPlayer::MediaStatus status) { callbackQMediaPlayerControl_MediaStatusChanged(this, status); };
	const QIODevice * mediaStream() const { return static_cast<QIODevice*>(callbackQMediaPlayerControl_MediaStream(const_cast<MyQMediaPlayerControl*>(this))); };
	void Signal_MutedChanged(bool mute) { callbackQMediaPlayerControl_MutedChanged(this, mute); };
	void pause() { callbackQMediaPlayerControl_Pause(this); };
	void play() { callbackQMediaPlayerControl_Play(this); };
	qreal playbackRate() const { return callbackQMediaPlayerControl_PlaybackRate(const_cast<MyQMediaPlayerControl*>(this)); };
	void Signal_PlaybackRateChanged(qreal rate) { callbackQMediaPlayerControl_PlaybackRateChanged(this, rate); };
	qint64 position() const { return callbackQMediaPlayerControl_Position(const_cast<MyQMediaPlayerControl*>(this)); };
	void Signal_PositionChanged(qint64 position) { callbackQMediaPlayerControl_PositionChanged(this, position); };
	void Signal_SeekableChanged(bool seekable) { callbackQMediaPlayerControl_SeekableChanged(this, seekable); };
	void setMedia(const QMediaContent & media, QIODevice * stream) { callbackQMediaPlayerControl_SetMedia(this, const_cast<QMediaContent*>(&media), stream); };
	void setMuted(bool mute) { callbackQMediaPlayerControl_SetMuted(this, mute); };
	void setPlaybackRate(qreal rate) { callbackQMediaPlayerControl_SetPlaybackRate(this, rate); };
	void setPosition(qint64 position) { callbackQMediaPlayerControl_SetPosition(this, position); };
	void setVolume(int volume) { callbackQMediaPlayerControl_SetVolume(this, volume); };
	QMediaPlayer::State state() const { return static_cast<QMediaPlayer::State>(callbackQMediaPlayerControl_State(const_cast<MyQMediaPlayerControl*>(this))); };
	void Signal_StateChanged(QMediaPlayer::State state) { callbackQMediaPlayerControl_StateChanged(this, state); };
	void stop() { callbackQMediaPlayerControl_Stop(this); };
	void Signal_VideoAvailableChanged(bool video) { callbackQMediaPlayerControl_VideoAvailableChanged(this, video); };
	int volume() const { return callbackQMediaPlayerControl_Volume(const_cast<MyQMediaPlayerControl*>(this)); };
	void Signal_VolumeChanged(int volume) { callbackQMediaPlayerControl_VolumeChanged(this, volume); };
	void timerEvent(QTimerEvent * event) { callbackQMediaPlayerControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaPlayerControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaPlayerControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaPlayerControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaPlayerControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaPlayerControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaPlayerControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaPlayerControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaPlayerControl_MetaObject(const_cast<MyQMediaPlayerControl*>(this))); };
};

void* QMediaPlayerControl_NewQMediaPlayerControl(void* parent)
{
	return new MyQMediaPlayerControl(static_cast<QObject*>(parent));
}

void QMediaPlayerControl_ConnectAudioAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::audioAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_AudioAvailableChanged));
}

void QMediaPlayerControl_DisconnectAudioAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::audioAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_AudioAvailableChanged));
}

void QMediaPlayerControl_AudioAvailableChanged(void* ptr, char audio)
{
	static_cast<QMediaPlayerControl*>(ptr)->audioAvailableChanged(audio != 0);
}

void* QMediaPlayerControl_AvailablePlaybackRanges(void* ptr)
{
	return new QMediaTimeRange(static_cast<QMediaPlayerControl*>(ptr)->availablePlaybackRanges());
}

void QMediaPlayerControl_ConnectAvailablePlaybackRangesChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(const QMediaTimeRange &)>(&QMediaPlayerControl::availablePlaybackRangesChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(const QMediaTimeRange &)>(&MyQMediaPlayerControl::Signal_AvailablePlaybackRangesChanged));
}

void QMediaPlayerControl_DisconnectAvailablePlaybackRangesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(const QMediaTimeRange &)>(&QMediaPlayerControl::availablePlaybackRangesChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(const QMediaTimeRange &)>(&MyQMediaPlayerControl::Signal_AvailablePlaybackRangesChanged));
}

void QMediaPlayerControl_AvailablePlaybackRangesChanged(void* ptr, void* ranges)
{
	static_cast<QMediaPlayerControl*>(ptr)->availablePlaybackRangesChanged(*static_cast<QMediaTimeRange*>(ranges));
}

int QMediaPlayerControl_BufferStatus(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->bufferStatus();
}

void QMediaPlayerControl_ConnectBufferStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::bufferStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_BufferStatusChanged));
}

void QMediaPlayerControl_DisconnectBufferStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::bufferStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_BufferStatusChanged));
}

void QMediaPlayerControl_BufferStatusChanged(void* ptr, int progress)
{
	static_cast<QMediaPlayerControl*>(ptr)->bufferStatusChanged(progress);
}

long long QMediaPlayerControl_Duration(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->duration();
}

void QMediaPlayerControl_ConnectDurationChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(qint64)>(&QMediaPlayerControl::durationChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(qint64)>(&MyQMediaPlayerControl::Signal_DurationChanged));
}

void QMediaPlayerControl_DisconnectDurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(qint64)>(&QMediaPlayerControl::durationChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(qint64)>(&MyQMediaPlayerControl::Signal_DurationChanged));
}

void QMediaPlayerControl_DurationChanged(void* ptr, long long duration)
{
	static_cast<QMediaPlayerControl*>(ptr)->durationChanged(duration);
}

void QMediaPlayerControl_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int, const QString &)>(&QMediaPlayerControl::error), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int, const QString &)>(&MyQMediaPlayerControl::Signal_Error));
}

void QMediaPlayerControl_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int, const QString &)>(&QMediaPlayerControl::error), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int, const QString &)>(&MyQMediaPlayerControl::Signal_Error));
}

void QMediaPlayerControl_Error(void* ptr, int error, char* errorString)
{
	static_cast<QMediaPlayerControl*>(ptr)->error(error, QString(errorString));
}

char QMediaPlayerControl_IsAudioAvailable(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->isAudioAvailable();
}

char QMediaPlayerControl_IsMuted(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->isMuted();
}

char QMediaPlayerControl_IsSeekable(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->isSeekable();
}

char QMediaPlayerControl_IsVideoAvailable(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->isVideoAvailable();
}

void* QMediaPlayerControl_Media(void* ptr)
{
	return new QMediaContent(static_cast<QMediaPlayerControl*>(ptr)->media());
}

void QMediaPlayerControl_ConnectMediaChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(const QMediaContent &)>(&QMediaPlayerControl::mediaChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(const QMediaContent &)>(&MyQMediaPlayerControl::Signal_MediaChanged));
}

void QMediaPlayerControl_DisconnectMediaChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(const QMediaContent &)>(&QMediaPlayerControl::mediaChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(const QMediaContent &)>(&MyQMediaPlayerControl::Signal_MediaChanged));
}

void QMediaPlayerControl_MediaChanged(void* ptr, void* content)
{
	static_cast<QMediaPlayerControl*>(ptr)->mediaChanged(*static_cast<QMediaContent*>(content));
}

long long QMediaPlayerControl_MediaStatus(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->mediaStatus();
}

void QMediaPlayerControl_ConnectMediaStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayerControl::mediaStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayerControl::Signal_MediaStatusChanged));
}

void QMediaPlayerControl_DisconnectMediaStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayerControl::mediaStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayerControl::Signal_MediaStatusChanged));
}

void QMediaPlayerControl_MediaStatusChanged(void* ptr, long long status)
{
	static_cast<QMediaPlayerControl*>(ptr)->mediaStatusChanged(static_cast<QMediaPlayer::MediaStatus>(status));
}

void* QMediaPlayerControl_MediaStream(void* ptr)
{
	return const_cast<QIODevice*>(static_cast<QMediaPlayerControl*>(ptr)->mediaStream());
}

void QMediaPlayerControl_ConnectMutedChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::mutedChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_MutedChanged));
}

void QMediaPlayerControl_DisconnectMutedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::mutedChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_MutedChanged));
}

void QMediaPlayerControl_MutedChanged(void* ptr, char mute)
{
	static_cast<QMediaPlayerControl*>(ptr)->mutedChanged(mute != 0);
}

void QMediaPlayerControl_Pause(void* ptr)
{
	static_cast<QMediaPlayerControl*>(ptr)->pause();
}

void QMediaPlayerControl_Play(void* ptr)
{
	static_cast<QMediaPlayerControl*>(ptr)->play();
}

double QMediaPlayerControl_PlaybackRate(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->playbackRate();
}

void QMediaPlayerControl_ConnectPlaybackRateChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(qreal)>(&QMediaPlayerControl::playbackRateChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(qreal)>(&MyQMediaPlayerControl::Signal_PlaybackRateChanged));
}

void QMediaPlayerControl_DisconnectPlaybackRateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(qreal)>(&QMediaPlayerControl::playbackRateChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(qreal)>(&MyQMediaPlayerControl::Signal_PlaybackRateChanged));
}

void QMediaPlayerControl_PlaybackRateChanged(void* ptr, double rate)
{
	static_cast<QMediaPlayerControl*>(ptr)->playbackRateChanged(rate);
}

long long QMediaPlayerControl_Position(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->position();
}

void QMediaPlayerControl_ConnectPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(qint64)>(&QMediaPlayerControl::positionChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(qint64)>(&MyQMediaPlayerControl::Signal_PositionChanged));
}

void QMediaPlayerControl_DisconnectPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(qint64)>(&QMediaPlayerControl::positionChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(qint64)>(&MyQMediaPlayerControl::Signal_PositionChanged));
}

void QMediaPlayerControl_PositionChanged(void* ptr, long long position)
{
	static_cast<QMediaPlayerControl*>(ptr)->positionChanged(position);
}

void QMediaPlayerControl_ConnectSeekableChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::seekableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_SeekableChanged));
}

void QMediaPlayerControl_DisconnectSeekableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::seekableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_SeekableChanged));
}

void QMediaPlayerControl_SeekableChanged(void* ptr, char seekable)
{
	static_cast<QMediaPlayerControl*>(ptr)->seekableChanged(seekable != 0);
}

void QMediaPlayerControl_SetMedia(void* ptr, void* media, void* stream)
{
	static_cast<QMediaPlayerControl*>(ptr)->setMedia(*static_cast<QMediaContent*>(media), static_cast<QIODevice*>(stream));
}

void QMediaPlayerControl_SetMuted(void* ptr, char mute)
{
	static_cast<QMediaPlayerControl*>(ptr)->setMuted(mute != 0);
}

void QMediaPlayerControl_SetPlaybackRate(void* ptr, double rate)
{
	static_cast<QMediaPlayerControl*>(ptr)->setPlaybackRate(rate);
}

void QMediaPlayerControl_SetPosition(void* ptr, long long position)
{
	static_cast<QMediaPlayerControl*>(ptr)->setPosition(position);
}

void QMediaPlayerControl_SetVolume(void* ptr, int volume)
{
	static_cast<QMediaPlayerControl*>(ptr)->setVolume(volume);
}

long long QMediaPlayerControl_State(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->state();
}

void QMediaPlayerControl_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::State)>(&QMediaPlayerControl::stateChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::State)>(&MyQMediaPlayerControl::Signal_StateChanged));
}

void QMediaPlayerControl_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::State)>(&QMediaPlayerControl::stateChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::State)>(&MyQMediaPlayerControl::Signal_StateChanged));
}

void QMediaPlayerControl_StateChanged(void* ptr, long long state)
{
	static_cast<QMediaPlayerControl*>(ptr)->stateChanged(static_cast<QMediaPlayer::State>(state));
}

void QMediaPlayerControl_Stop(void* ptr)
{
	static_cast<QMediaPlayerControl*>(ptr)->stop();
}

void QMediaPlayerControl_ConnectVideoAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::videoAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_VideoAvailableChanged));
}

void QMediaPlayerControl_DisconnectVideoAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::videoAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_VideoAvailableChanged));
}

void QMediaPlayerControl_VideoAvailableChanged(void* ptr, char video)
{
	static_cast<QMediaPlayerControl*>(ptr)->videoAvailableChanged(video != 0);
}

int QMediaPlayerControl_Volume(void* ptr)
{
	return static_cast<QMediaPlayerControl*>(ptr)->volume();
}

void QMediaPlayerControl_ConnectVolumeChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::volumeChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_VolumeChanged));
}

void QMediaPlayerControl_DisconnectVolumeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::volumeChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_VolumeChanged));
}

void QMediaPlayerControl_VolumeChanged(void* ptr, int volume)
{
	static_cast<QMediaPlayerControl*>(ptr)->volumeChanged(volume);
}

void QMediaPlayerControl_DestroyQMediaPlayerControl(void* ptr)
{
	static_cast<QMediaPlayerControl*>(ptr)->~QMediaPlayerControl();
}

void QMediaPlayerControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaPlayerControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaPlayerControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaPlayerControl*>(ptr)->QMediaPlayerControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaPlayerControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaPlayerControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaPlayerControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaPlayerControl*>(ptr)->QMediaPlayerControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaPlayerControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaPlayerControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaPlayerControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaPlayerControl*>(ptr)->QMediaPlayerControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaPlayerControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaPlayerControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaPlayerControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaPlayerControl*>(ptr)->QMediaPlayerControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaPlayerControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlayerControl*>(ptr), "deleteLater");
}

void QMediaPlayerControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaPlayerControl*>(ptr)->QMediaPlayerControl::deleteLater();
}

void QMediaPlayerControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaPlayerControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaPlayerControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaPlayerControl*>(ptr)->QMediaPlayerControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaPlayerControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaPlayerControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaPlayerControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaPlayerControl*>(ptr)->QMediaPlayerControl::event(static_cast<QEvent*>(e));
}

char QMediaPlayerControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaPlayerControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaPlayerControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaPlayerControl*>(ptr)->QMediaPlayerControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaPlayerControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaPlayerControl*>(ptr)->metaObject());
}

void* QMediaPlayerControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaPlayerControl*>(ptr)->QMediaPlayerControl::metaObject());
}

class MyQMediaPlaylist: public QMediaPlaylist
{
public:
	MyQMediaPlaylist(QObject *parent) : QMediaPlaylist(parent) {};
	void Signal_CurrentIndexChanged(int position) { callbackQMediaPlaylist_CurrentIndexChanged(this, position); };
	void Signal_CurrentMediaChanged(const QMediaContent & content) { callbackQMediaPlaylist_CurrentMediaChanged(this, const_cast<QMediaContent*>(&content)); };
	void Signal_LoadFailed() { callbackQMediaPlaylist_LoadFailed(this); };
	void Signal_Loaded() { callbackQMediaPlaylist_Loaded(this); };
	void Signal_MediaAboutToBeInserted(int start, int end) { callbackQMediaPlaylist_MediaAboutToBeInserted(this, start, end); };
	void Signal_MediaAboutToBeRemoved(int start, int end) { callbackQMediaPlaylist_MediaAboutToBeRemoved(this, start, end); };
	void Signal_MediaChanged(int start, int end) { callbackQMediaPlaylist_MediaChanged(this, start, end); };
	void Signal_MediaInserted(int start, int end) { callbackQMediaPlaylist_MediaInserted(this, start, end); };
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQMediaPlaylist_MediaObject(const_cast<MyQMediaPlaylist*>(this))); };
	void Signal_MediaRemoved(int start, int end) { callbackQMediaPlaylist_MediaRemoved(this, start, end); };
	void next() { callbackQMediaPlaylist_Next(this); };
	void Signal_PlaybackModeChanged(QMediaPlaylist::PlaybackMode mode) { callbackQMediaPlaylist_PlaybackModeChanged(this, mode); };
	void previous() { callbackQMediaPlaylist_Previous(this); };
	void setCurrentIndex(int playlistPosition) { callbackQMediaPlaylist_SetCurrentIndex(this, playlistPosition); };
	void shuffle() { callbackQMediaPlaylist_Shuffle(this); };
	 ~MyQMediaPlaylist() { callbackQMediaPlaylist_DestroyQMediaPlaylist(this); };
	void timerEvent(QTimerEvent * event) { callbackQMediaPlaylist_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaPlaylist_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaPlaylist_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaPlaylist_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaPlaylist_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaPlaylist_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaPlaylist_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaPlaylist_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaPlaylist_MetaObject(const_cast<MyQMediaPlaylist*>(this))); };
	bool setMediaObject(QMediaObject * object) { return callbackQMediaPlaylist_SetMediaObject(this, object) != 0; };
};

long long QMediaPlaylist_PlaybackMode(void* ptr)
{
	return static_cast<QMediaPlaylist*>(ptr)->playbackMode();
}

void QMediaPlaylist_SetPlaybackMode(void* ptr, long long mode)
{
	static_cast<QMediaPlaylist*>(ptr)->setPlaybackMode(static_cast<QMediaPlaylist::PlaybackMode>(mode));
}

void* QMediaPlaylist_NewQMediaPlaylist(void* parent)
{
	return new MyQMediaPlaylist(static_cast<QObject*>(parent));
}

char QMediaPlaylist_AddMedia(void* ptr, void* content)
{
	return static_cast<QMediaPlaylist*>(ptr)->addMedia(*static_cast<QMediaContent*>(content));
}

char QMediaPlaylist_Clear(void* ptr)
{
	return static_cast<QMediaPlaylist*>(ptr)->clear();
}

int QMediaPlaylist_CurrentIndex(void* ptr)
{
	return static_cast<QMediaPlaylist*>(ptr)->currentIndex();
}

void QMediaPlaylist_ConnectCurrentIndexChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int)>(&QMediaPlaylist::currentIndexChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int)>(&MyQMediaPlaylist::Signal_CurrentIndexChanged));
}

void QMediaPlaylist_DisconnectCurrentIndexChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int)>(&QMediaPlaylist::currentIndexChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int)>(&MyQMediaPlaylist::Signal_CurrentIndexChanged));
}

void QMediaPlaylist_CurrentIndexChanged(void* ptr, int position)
{
	static_cast<QMediaPlaylist*>(ptr)->currentIndexChanged(position);
}

void* QMediaPlaylist_CurrentMedia(void* ptr)
{
	return new QMediaContent(static_cast<QMediaPlaylist*>(ptr)->currentMedia());
}

void QMediaPlaylist_ConnectCurrentMediaChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(const QMediaContent &)>(&QMediaPlaylist::currentMediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(const QMediaContent &)>(&MyQMediaPlaylist::Signal_CurrentMediaChanged));
}

void QMediaPlaylist_DisconnectCurrentMediaChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(const QMediaContent &)>(&QMediaPlaylist::currentMediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(const QMediaContent &)>(&MyQMediaPlaylist::Signal_CurrentMediaChanged));
}

void QMediaPlaylist_CurrentMediaChanged(void* ptr, void* content)
{
	static_cast<QMediaPlaylist*>(ptr)->currentMediaChanged(*static_cast<QMediaContent*>(content));
}

long long QMediaPlaylist_Error(void* ptr)
{
	return static_cast<QMediaPlaylist*>(ptr)->error();
}

struct QtMultimedia_PackedString QMediaPlaylist_ErrorString(void* ptr)
{
	return ({ QByteArray t91ff89 = static_cast<QMediaPlaylist*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t91ff89.prepend("WHITESPACE").constData()+10), t91ff89.size()-10 }; });
}

char QMediaPlaylist_InsertMedia(void* ptr, int pos, void* content)
{
	return static_cast<QMediaPlaylist*>(ptr)->insertMedia(pos, *static_cast<QMediaContent*>(content));
}

char QMediaPlaylist_IsEmpty(void* ptr)
{
	return static_cast<QMediaPlaylist*>(ptr)->isEmpty();
}

char QMediaPlaylist_IsReadOnly(void* ptr)
{
	return static_cast<QMediaPlaylist*>(ptr)->isReadOnly();
}

void QMediaPlaylist_Load3(void* ptr, void* device, char* format)
{
	static_cast<QMediaPlaylist*>(ptr)->load(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

void QMediaPlaylist_Load(void* ptr, void* request, char* format)
{
	static_cast<QMediaPlaylist*>(ptr)->load(*static_cast<QNetworkRequest*>(request), const_cast<const char*>(format));
}

void QMediaPlaylist_Load2(void* ptr, void* location, char* format)
{
	static_cast<QMediaPlaylist*>(ptr)->load(*static_cast<QUrl*>(location), const_cast<const char*>(format));
}

void QMediaPlaylist_ConnectLoadFailed(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loadFailed), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_LoadFailed));
}

void QMediaPlaylist_DisconnectLoadFailed(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loadFailed), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_LoadFailed));
}

void QMediaPlaylist_LoadFailed(void* ptr)
{
	static_cast<QMediaPlaylist*>(ptr)->loadFailed();
}

void QMediaPlaylist_ConnectLoaded(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loaded), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_Loaded));
}

void QMediaPlaylist_DisconnectLoaded(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loaded), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_Loaded));
}

void QMediaPlaylist_Loaded(void* ptr)
{
	static_cast<QMediaPlaylist*>(ptr)->loaded();
}

void* QMediaPlaylist_Media(void* ptr, int index)
{
	return new QMediaContent(static_cast<QMediaPlaylist*>(ptr)->media(index));
}

void QMediaPlaylist_ConnectMediaAboutToBeInserted(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeInserted));
}

void QMediaPlaylist_DisconnectMediaAboutToBeInserted(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeInserted));
}

void QMediaPlaylist_MediaAboutToBeInserted(void* ptr, int start, int end)
{
	static_cast<QMediaPlaylist*>(ptr)->mediaAboutToBeInserted(start, end);
}

void QMediaPlaylist_ConnectMediaAboutToBeRemoved(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeRemoved));
}

void QMediaPlaylist_DisconnectMediaAboutToBeRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeRemoved));
}

void QMediaPlaylist_MediaAboutToBeRemoved(void* ptr, int start, int end)
{
	static_cast<QMediaPlaylist*>(ptr)->mediaAboutToBeRemoved(start, end);
}

void QMediaPlaylist_ConnectMediaChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaChanged));
}

void QMediaPlaylist_DisconnectMediaChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaChanged));
}

void QMediaPlaylist_MediaChanged(void* ptr, int start, int end)
{
	static_cast<QMediaPlaylist*>(ptr)->mediaChanged(start, end);
}

int QMediaPlaylist_MediaCount(void* ptr)
{
	return static_cast<QMediaPlaylist*>(ptr)->mediaCount();
}

void QMediaPlaylist_ConnectMediaInserted(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaInserted));
}

void QMediaPlaylist_DisconnectMediaInserted(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaInserted));
}

void QMediaPlaylist_MediaInserted(void* ptr, int start, int end)
{
	static_cast<QMediaPlaylist*>(ptr)->mediaInserted(start, end);
}

void* QMediaPlaylist_MediaObject(void* ptr)
{
	return static_cast<QMediaPlaylist*>(ptr)->mediaObject();
}

void* QMediaPlaylist_MediaObjectDefault(void* ptr)
{
	return static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::mediaObject();
}

void QMediaPlaylist_ConnectMediaRemoved(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaRemoved));
}

void QMediaPlaylist_DisconnectMediaRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaRemoved));
}

void QMediaPlaylist_MediaRemoved(void* ptr, int start, int end)
{
	static_cast<QMediaPlaylist*>(ptr)->mediaRemoved(start, end);
}

char QMediaPlaylist_MoveMedia(void* ptr, int from, int to)
{
	return static_cast<QMediaPlaylist*>(ptr)->moveMedia(from, to);
}

void QMediaPlaylist_Next(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "next");
}

int QMediaPlaylist_NextIndex(void* ptr, int steps)
{
	return static_cast<QMediaPlaylist*>(ptr)->nextIndex(steps);
}

void QMediaPlaylist_ConnectPlaybackModeChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&QMediaPlaylist::playbackModeChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&MyQMediaPlaylist::Signal_PlaybackModeChanged));
}

void QMediaPlaylist_DisconnectPlaybackModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&QMediaPlaylist::playbackModeChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&MyQMediaPlaylist::Signal_PlaybackModeChanged));
}

void QMediaPlaylist_PlaybackModeChanged(void* ptr, long long mode)
{
	static_cast<QMediaPlaylist*>(ptr)->playbackModeChanged(static_cast<QMediaPlaylist::PlaybackMode>(mode));
}

void QMediaPlaylist_Previous(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "previous");
}

int QMediaPlaylist_PreviousIndex(void* ptr, int steps)
{
	return static_cast<QMediaPlaylist*>(ptr)->previousIndex(steps);
}

char QMediaPlaylist_RemoveMedia(void* ptr, int pos)
{
	return static_cast<QMediaPlaylist*>(ptr)->removeMedia(pos);
}

char QMediaPlaylist_RemoveMedia2(void* ptr, int start, int end)
{
	return static_cast<QMediaPlaylist*>(ptr)->removeMedia(start, end);
}

char QMediaPlaylist_Save2(void* ptr, void* device, char* format)
{
	return static_cast<QMediaPlaylist*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

char QMediaPlaylist_Save(void* ptr, void* location, char* format)
{
	return static_cast<QMediaPlaylist*>(ptr)->save(*static_cast<QUrl*>(location), const_cast<const char*>(format));
}

void QMediaPlaylist_SetCurrentIndex(void* ptr, int playlistPosition)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "setCurrentIndex", Q_ARG(int, playlistPosition));
}

void QMediaPlaylist_Shuffle(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "shuffle");
}

void QMediaPlaylist_DestroyQMediaPlaylist(void* ptr)
{
	static_cast<QMediaPlaylist*>(ptr)->~QMediaPlaylist();
}

void QMediaPlaylist_DestroyQMediaPlaylistDefault(void* ptr)
{

}

void QMediaPlaylist_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaPlaylist*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaPlaylist_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaPlaylist_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaPlaylist*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaPlaylist_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaPlaylist_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaPlaylist*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaPlaylist_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaPlaylist_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaPlaylist*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaPlaylist_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::customEvent(static_cast<QEvent*>(event));
}

void QMediaPlaylist_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "deleteLater");
}

void QMediaPlaylist_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::deleteLater();
}

void QMediaPlaylist_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaPlaylist*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaPlaylist_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaPlaylist_Event(void* ptr, void* e)
{
	return static_cast<QMediaPlaylist*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaPlaylist_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::event(static_cast<QEvent*>(e));
}

char QMediaPlaylist_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaPlaylist*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaPlaylist_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaPlaylist_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaPlaylist*>(ptr)->metaObject());
}

void* QMediaPlaylist_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::metaObject());
}

char QMediaPlaylist_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QMediaPlaylist*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

char QMediaPlaylist_SetMediaObjectDefault(void* ptr, void* object)
{
	return static_cast<QMediaPlaylist*>(ptr)->QMediaPlaylist::setMediaObject(static_cast<QMediaObject*>(object));
}

class MyQMediaRecorder: public QMediaRecorder
{
public:
	MyQMediaRecorder(QMediaObject *mediaObject, QObject *parent) : QMediaRecorder(mediaObject, parent) {};
	void setMuted(bool muted) { callbackQMediaRecorder_SetMuted(this, muted); };
	void setVolume(qreal volume) { callbackQMediaRecorder_SetVolume(this, volume); };
	void Signal_ActualLocationChanged(const QUrl & location) { callbackQMediaRecorder_ActualLocationChanged(this, const_cast<QUrl*>(&location)); };
	void Signal_AvailabilityChanged2(QMultimedia::AvailabilityStatus availability) { callbackQMediaRecorder_AvailabilityChanged2(this, availability); };
	void Signal_AvailabilityChanged(bool available) { callbackQMediaRecorder_AvailabilityChanged(this, available); };
	void Signal_DurationChanged(qint64 duration) { callbackQMediaRecorder_DurationChanged(this, duration); };
	void Signal_Error2(QMediaRecorder::Error error) { callbackQMediaRecorder_Error2(this, error); };
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQMediaRecorder_MediaObject(const_cast<MyQMediaRecorder*>(this))); };
	void Signal_MetaDataAvailableChanged(bool available) { callbackQMediaRecorder_MetaDataAvailableChanged(this, available); };
	void Signal_MetaDataChanged() { callbackQMediaRecorder_MetaDataChanged(this); };
	void Signal_MetaDataChanged2(const QString & key, const QVariant & value) { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };callbackQMediaRecorder_MetaDataChanged2(this, keyPacked, const_cast<QVariant*>(&value)); };
	void Signal_MetaDataWritableChanged(bool writable) { callbackQMediaRecorder_MetaDataWritableChanged(this, writable); };
	void Signal_MutedChanged(bool muted) { callbackQMediaRecorder_MutedChanged(this, muted); };
	void pause() { callbackQMediaRecorder_Pause(this); };
	void record() { callbackQMediaRecorder_Record(this); };
	void Signal_StateChanged(QMediaRecorder::State state) { callbackQMediaRecorder_StateChanged(this, state); };
	void Signal_StatusChanged(QMediaRecorder::Status status) { callbackQMediaRecorder_StatusChanged(this, status); };
	void stop() { callbackQMediaRecorder_Stop(this); };
	void Signal_VolumeChanged(qreal volume) { callbackQMediaRecorder_VolumeChanged(this, volume); };
	void timerEvent(QTimerEvent * event) { callbackQMediaRecorder_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaRecorder_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaRecorder_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaRecorder_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaRecorder_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaRecorder_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaRecorder_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaRecorder_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaRecorder_MetaObject(const_cast<MyQMediaRecorder*>(this))); };
	bool setMediaObject(QMediaObject * object) { return callbackQMediaRecorder_SetMediaObject(this, object) != 0; };
};

void* QMediaRecorder_ActualLocation(void* ptr)
{
	return new QUrl(static_cast<QMediaRecorder*>(ptr)->actualLocation());
}

long long QMediaRecorder_Duration(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->duration();
}

char QMediaRecorder_IsMetaDataAvailable(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->isMetaDataAvailable();
}

char QMediaRecorder_IsMetaDataWritable(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->isMetaDataWritable();
}

char QMediaRecorder_IsMuted(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->isMuted();
}

void* QMediaRecorder_OutputLocation(void* ptr)
{
	return new QUrl(static_cast<QMediaRecorder*>(ptr)->outputLocation());
}

void QMediaRecorder_SetMuted(void* ptr, char muted)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

char QMediaRecorder_SetOutputLocation(void* ptr, void* location)
{
	return static_cast<QMediaRecorder*>(ptr)->setOutputLocation(*static_cast<QUrl*>(location));
}

void QMediaRecorder_SetVolume(void* ptr, double volume)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "setVolume", Q_ARG(qreal, volume));
}

double QMediaRecorder_Volume(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->volume();
}

void* QMediaRecorder_NewQMediaRecorder(void* mediaObject, void* parent)
{
	return new MyQMediaRecorder(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

void QMediaRecorder_ConnectActualLocationChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(const QUrl &)>(&QMediaRecorder::actualLocationChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(const QUrl &)>(&MyQMediaRecorder::Signal_ActualLocationChanged));
}

void QMediaRecorder_DisconnectActualLocationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(const QUrl &)>(&QMediaRecorder::actualLocationChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(const QUrl &)>(&MyQMediaRecorder::Signal_ActualLocationChanged));
}

void QMediaRecorder_ActualLocationChanged(void* ptr, void* location)
{
	static_cast<QMediaRecorder*>(ptr)->actualLocationChanged(*static_cast<QUrl*>(location));
}

struct QtMultimedia_PackedString QMediaRecorder_AudioCodecDescription(void* ptr, char* codec)
{
	return ({ QByteArray tbd8228 = static_cast<QMediaRecorder*>(ptr)->audioCodecDescription(QString(codec)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tbd8228.prepend("WHITESPACE").constData()+10), tbd8228.size()-10 }; });
}

void* QMediaRecorder_AudioSettings(void* ptr)
{
	return new QAudioEncoderSettings(static_cast<QMediaRecorder*>(ptr)->audioSettings());
}

long long QMediaRecorder_Availability(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->availability();
}

void QMediaRecorder_ConnectAvailabilityChanged2(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMultimedia::AvailabilityStatus)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMultimedia::AvailabilityStatus)>(&MyQMediaRecorder::Signal_AvailabilityChanged2));
}

void QMediaRecorder_DisconnectAvailabilityChanged2(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMultimedia::AvailabilityStatus)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMultimedia::AvailabilityStatus)>(&MyQMediaRecorder::Signal_AvailabilityChanged2));
}

void QMediaRecorder_AvailabilityChanged2(void* ptr, long long availability)
{
	static_cast<QMediaRecorder*>(ptr)->availabilityChanged(static_cast<QMultimedia::AvailabilityStatus>(availability));
}

void QMediaRecorder_ConnectAvailabilityChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_AvailabilityChanged));
}

void QMediaRecorder_DisconnectAvailabilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_AvailabilityChanged));
}

void QMediaRecorder_AvailabilityChanged(void* ptr, char available)
{
	static_cast<QMediaRecorder*>(ptr)->availabilityChanged(available != 0);
}

struct QtMultimedia_PackedString QMediaRecorder_AvailableMetaData(void* ptr)
{
	return ({ QByteArray t148b06 = static_cast<QMediaRecorder*>(ptr)->availableMetaData().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t148b06.prepend("WHITESPACE").constData()+10), t148b06.size()-10 }; });
}

struct QtMultimedia_PackedString QMediaRecorder_ContainerDescription(void* ptr, char* format)
{
	return ({ QByteArray tf7fe4c = static_cast<QMediaRecorder*>(ptr)->containerDescription(QString(format)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tf7fe4c.prepend("WHITESPACE").constData()+10), tf7fe4c.size()-10 }; });
}

struct QtMultimedia_PackedString QMediaRecorder_ContainerFormat(void* ptr)
{
	return ({ QByteArray tce675e = static_cast<QMediaRecorder*>(ptr)->containerFormat().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tce675e.prepend("WHITESPACE").constData()+10), tce675e.size()-10 }; });
}

void QMediaRecorder_ConnectDurationChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(qint64)>(&QMediaRecorder::durationChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(qint64)>(&MyQMediaRecorder::Signal_DurationChanged));
}

void QMediaRecorder_DisconnectDurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(qint64)>(&QMediaRecorder::durationChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(qint64)>(&MyQMediaRecorder::Signal_DurationChanged));
}

void QMediaRecorder_DurationChanged(void* ptr, long long duration)
{
	static_cast<QMediaRecorder*>(ptr)->durationChanged(duration);
}

void QMediaRecorder_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Error)>(&QMediaRecorder::error), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Error)>(&MyQMediaRecorder::Signal_Error2));
}

void QMediaRecorder_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Error)>(&QMediaRecorder::error), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Error)>(&MyQMediaRecorder::Signal_Error2));
}

void QMediaRecorder_Error2(void* ptr, long long error)
{
	static_cast<QMediaRecorder*>(ptr)->error(static_cast<QMediaRecorder::Error>(error));
}

long long QMediaRecorder_Error(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->error();
}

struct QtMultimedia_PackedString QMediaRecorder_ErrorString(void* ptr)
{
	return ({ QByteArray t8975f0 = static_cast<QMediaRecorder*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t8975f0.prepend("WHITESPACE").constData()+10), t8975f0.size()-10 }; });
}

char QMediaRecorder_IsAvailable(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->isAvailable();
}

void* QMediaRecorder_MediaObject(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->mediaObject();
}

void* QMediaRecorder_MediaObjectDefault(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::mediaObject();
}

void* QMediaRecorder_MetaData(void* ptr, char* key)
{
	return new QVariant(static_cast<QMediaRecorder*>(ptr)->metaData(QString(key)));
}

void QMediaRecorder_ConnectMetaDataAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataAvailableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataAvailableChanged));
}

void QMediaRecorder_DisconnectMetaDataAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataAvailableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataAvailableChanged));
}

void QMediaRecorder_MetaDataAvailableChanged(void* ptr, char available)
{
	static_cast<QMediaRecorder*>(ptr)->metaDataAvailableChanged(available != 0);
}

void QMediaRecorder_ConnectMetaDataChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)()>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)()>(&MyQMediaRecorder::Signal_MetaDataChanged));
}

void QMediaRecorder_DisconnectMetaDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)()>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)()>(&MyQMediaRecorder::Signal_MetaDataChanged));
}

void QMediaRecorder_MetaDataChanged(void* ptr)
{
	static_cast<QMediaRecorder*>(ptr)->metaDataChanged();
}

void QMediaRecorder_ConnectMetaDataChanged2(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(const QString &, const QVariant &)>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(const QString &, const QVariant &)>(&MyQMediaRecorder::Signal_MetaDataChanged2));
}

void QMediaRecorder_DisconnectMetaDataChanged2(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(const QString &, const QVariant &)>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(const QString &, const QVariant &)>(&MyQMediaRecorder::Signal_MetaDataChanged2));
}

void QMediaRecorder_MetaDataChanged2(void* ptr, char* key, void* value)
{
	static_cast<QMediaRecorder*>(ptr)->metaDataChanged(QString(key), *static_cast<QVariant*>(value));
}

void QMediaRecorder_ConnectMetaDataWritableChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataWritableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataWritableChanged));
}

void QMediaRecorder_DisconnectMetaDataWritableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataWritableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataWritableChanged));
}

void QMediaRecorder_MetaDataWritableChanged(void* ptr, char writable)
{
	static_cast<QMediaRecorder*>(ptr)->metaDataWritableChanged(writable != 0);
}

void QMediaRecorder_ConnectMutedChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::mutedChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MutedChanged));
}

void QMediaRecorder_DisconnectMutedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::mutedChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MutedChanged));
}

void QMediaRecorder_MutedChanged(void* ptr, char muted)
{
	static_cast<QMediaRecorder*>(ptr)->mutedChanged(muted != 0);
}

void QMediaRecorder_Pause(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "pause");
}

void QMediaRecorder_Record(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "record");
}

void QMediaRecorder_SetAudioSettings(void* ptr, void* settings)
{
	static_cast<QMediaRecorder*>(ptr)->setAudioSettings(*static_cast<QAudioEncoderSettings*>(settings));
}

void QMediaRecorder_SetContainerFormat(void* ptr, char* container)
{
	static_cast<QMediaRecorder*>(ptr)->setContainerFormat(QString(container));
}

void QMediaRecorder_SetEncodingSettings(void* ptr, void* audio, void* video, char* container)
{
	static_cast<QMediaRecorder*>(ptr)->setEncodingSettings(*static_cast<QAudioEncoderSettings*>(audio), *static_cast<QVideoEncoderSettings*>(video), QString(container));
}

void QMediaRecorder_SetMetaData(void* ptr, char* key, void* value)
{
	static_cast<QMediaRecorder*>(ptr)->setMetaData(QString(key), *static_cast<QVariant*>(value));
}

void QMediaRecorder_SetVideoSettings(void* ptr, void* settings)
{
	static_cast<QMediaRecorder*>(ptr)->setVideoSettings(*static_cast<QVideoEncoderSettings*>(settings));
}

long long QMediaRecorder_State(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->state();
}

void QMediaRecorder_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::State)>(&QMediaRecorder::stateChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::State)>(&MyQMediaRecorder::Signal_StateChanged));
}

void QMediaRecorder_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::State)>(&QMediaRecorder::stateChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::State)>(&MyQMediaRecorder::Signal_StateChanged));
}

void QMediaRecorder_StateChanged(void* ptr, long long state)
{
	static_cast<QMediaRecorder*>(ptr)->stateChanged(static_cast<QMediaRecorder::State>(state));
}

long long QMediaRecorder_Status(void* ptr)
{
	return static_cast<QMediaRecorder*>(ptr)->status();
}

void QMediaRecorder_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Status)>(&QMediaRecorder::statusChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Status)>(&MyQMediaRecorder::Signal_StatusChanged));
}

void QMediaRecorder_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Status)>(&QMediaRecorder::statusChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Status)>(&MyQMediaRecorder::Signal_StatusChanged));
}

void QMediaRecorder_StatusChanged(void* ptr, long long status)
{
	static_cast<QMediaRecorder*>(ptr)->statusChanged(static_cast<QMediaRecorder::Status>(status));
}

void QMediaRecorder_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "stop");
}

struct QtMultimedia_PackedString QMediaRecorder_SupportedAudioCodecs(void* ptr)
{
	return ({ QByteArray t489f71 = static_cast<QMediaRecorder*>(ptr)->supportedAudioCodecs().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t489f71.prepend("WHITESPACE").constData()+10), t489f71.size()-10 }; });
}

struct QtMultimedia_PackedString QMediaRecorder_SupportedContainers(void* ptr)
{
	return ({ QByteArray t948956 = static_cast<QMediaRecorder*>(ptr)->supportedContainers().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t948956.prepend("WHITESPACE").constData()+10), t948956.size()-10 }; });
}

struct QtMultimedia_PackedString QMediaRecorder_SupportedVideoCodecs(void* ptr)
{
	return ({ QByteArray tb262ef = static_cast<QMediaRecorder*>(ptr)->supportedVideoCodecs().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tb262ef.prepend("WHITESPACE").constData()+10), tb262ef.size()-10 }; });
}

struct QtMultimedia_PackedString QMediaRecorder_VideoCodecDescription(void* ptr, char* codec)
{
	return ({ QByteArray ta0c17f = static_cast<QMediaRecorder*>(ptr)->videoCodecDescription(QString(codec)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(ta0c17f.prepend("WHITESPACE").constData()+10), ta0c17f.size()-10 }; });
}

void* QMediaRecorder_VideoSettings(void* ptr)
{
	return new QVideoEncoderSettings(static_cast<QMediaRecorder*>(ptr)->videoSettings());
}

void QMediaRecorder_ConnectVolumeChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(qreal)>(&QMediaRecorder::volumeChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(qreal)>(&MyQMediaRecorder::Signal_VolumeChanged));
}

void QMediaRecorder_DisconnectVolumeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(qreal)>(&QMediaRecorder::volumeChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(qreal)>(&MyQMediaRecorder::Signal_VolumeChanged));
}

void QMediaRecorder_VolumeChanged(void* ptr, double volume)
{
	static_cast<QMediaRecorder*>(ptr)->volumeChanged(volume);
}

void QMediaRecorder_DestroyQMediaRecorder(void* ptr)
{
	static_cast<QMediaRecorder*>(ptr)->~QMediaRecorder();
}

void QMediaRecorder_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaRecorder*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaRecorder_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaRecorder_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaRecorder*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaRecorder_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaRecorder_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaRecorder*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaRecorder_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaRecorder_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaRecorder*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaRecorder_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::customEvent(static_cast<QEvent*>(event));
}

void QMediaRecorder_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "deleteLater");
}

void QMediaRecorder_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::deleteLater();
}

void QMediaRecorder_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaRecorder*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaRecorder_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaRecorder_Event(void* ptr, void* e)
{
	return static_cast<QMediaRecorder*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaRecorder_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::event(static_cast<QEvent*>(e));
}

char QMediaRecorder_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaRecorder*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaRecorder_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaRecorder_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaRecorder*>(ptr)->metaObject());
}

void* QMediaRecorder_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::metaObject());
}

char QMediaRecorder_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QMediaRecorder*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

char QMediaRecorder_SetMediaObjectDefault(void* ptr, void* object)
{
	return static_cast<QMediaRecorder*>(ptr)->QMediaRecorder::setMediaObject(static_cast<QMediaObject*>(object));
}

class MyQMediaRecorderControl: public QMediaRecorderControl
{
public:
	MyQMediaRecorderControl(QObject *parent) : QMediaRecorderControl(parent) {};
	void Signal_ActualLocationChanged(const QUrl & location) { callbackQMediaRecorderControl_ActualLocationChanged(this, const_cast<QUrl*>(&location)); };
	void applySettings() { callbackQMediaRecorderControl_ApplySettings(this); };
	qint64 duration() const { return callbackQMediaRecorderControl_Duration(const_cast<MyQMediaRecorderControl*>(this)); };
	void Signal_DurationChanged(qint64 duration) { callbackQMediaRecorderControl_DurationChanged(this, duration); };
	void Signal_Error(int error, const QString & errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtMultimedia_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQMediaRecorderControl_Error(this, error, errorStringPacked); };
	bool isMuted() const { return callbackQMediaRecorderControl_IsMuted(const_cast<MyQMediaRecorderControl*>(this)) != 0; };
	void Signal_MutedChanged(bool muted) { callbackQMediaRecorderControl_MutedChanged(this, muted); };
	QUrl outputLocation() const { return *static_cast<QUrl*>(callbackQMediaRecorderControl_OutputLocation(const_cast<MyQMediaRecorderControl*>(this))); };
	void setMuted(bool muted) { callbackQMediaRecorderControl_SetMuted(this, muted); };
	bool setOutputLocation(const QUrl & location) { return callbackQMediaRecorderControl_SetOutputLocation(this, const_cast<QUrl*>(&location)) != 0; };
	void setState(QMediaRecorder::State state) { callbackQMediaRecorderControl_SetState(this, state); };
	void setVolume(qreal gain) { callbackQMediaRecorderControl_SetVolume(this, gain); };
	QMediaRecorder::State state() const { return static_cast<QMediaRecorder::State>(callbackQMediaRecorderControl_State(const_cast<MyQMediaRecorderControl*>(this))); };
	void Signal_StateChanged(QMediaRecorder::State state) { callbackQMediaRecorderControl_StateChanged(this, state); };
	QMediaRecorder::Status status() const { return static_cast<QMediaRecorder::Status>(callbackQMediaRecorderControl_Status(const_cast<MyQMediaRecorderControl*>(this))); };
	void Signal_StatusChanged(QMediaRecorder::Status status) { callbackQMediaRecorderControl_StatusChanged(this, status); };
	qreal volume() const { return callbackQMediaRecorderControl_Volume(const_cast<MyQMediaRecorderControl*>(this)); };
	void Signal_VolumeChanged(qreal gain) { callbackQMediaRecorderControl_VolumeChanged(this, gain); };
	 ~MyQMediaRecorderControl() { callbackQMediaRecorderControl_DestroyQMediaRecorderControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQMediaRecorderControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaRecorderControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaRecorderControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaRecorderControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaRecorderControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaRecorderControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaRecorderControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaRecorderControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaRecorderControl_MetaObject(const_cast<MyQMediaRecorderControl*>(this))); };
};

void* QMediaRecorderControl_NewQMediaRecorderControl(void* parent)
{
	return new MyQMediaRecorderControl(static_cast<QObject*>(parent));
}

void QMediaRecorderControl_ConnectActualLocationChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(const QUrl &)>(&QMediaRecorderControl::actualLocationChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(const QUrl &)>(&MyQMediaRecorderControl::Signal_ActualLocationChanged));
}

void QMediaRecorderControl_DisconnectActualLocationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(const QUrl &)>(&QMediaRecorderControl::actualLocationChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(const QUrl &)>(&MyQMediaRecorderControl::Signal_ActualLocationChanged));
}

void QMediaRecorderControl_ActualLocationChanged(void* ptr, void* location)
{
	static_cast<QMediaRecorderControl*>(ptr)->actualLocationChanged(*static_cast<QUrl*>(location));
}

void QMediaRecorderControl_ApplySettings(void* ptr)
{
	static_cast<QMediaRecorderControl*>(ptr)->applySettings();
}

long long QMediaRecorderControl_Duration(void* ptr)
{
	return static_cast<QMediaRecorderControl*>(ptr)->duration();
}

void QMediaRecorderControl_ConnectDurationChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(qint64)>(&QMediaRecorderControl::durationChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(qint64)>(&MyQMediaRecorderControl::Signal_DurationChanged));
}

void QMediaRecorderControl_DisconnectDurationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(qint64)>(&QMediaRecorderControl::durationChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(qint64)>(&MyQMediaRecorderControl::Signal_DurationChanged));
}

void QMediaRecorderControl_DurationChanged(void* ptr, long long duration)
{
	static_cast<QMediaRecorderControl*>(ptr)->durationChanged(duration);
}

void QMediaRecorderControl_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(int, const QString &)>(&QMediaRecorderControl::error), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(int, const QString &)>(&MyQMediaRecorderControl::Signal_Error));
}

void QMediaRecorderControl_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(int, const QString &)>(&QMediaRecorderControl::error), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(int, const QString &)>(&MyQMediaRecorderControl::Signal_Error));
}

void QMediaRecorderControl_Error(void* ptr, int error, char* errorString)
{
	static_cast<QMediaRecorderControl*>(ptr)->error(error, QString(errorString));
}

char QMediaRecorderControl_IsMuted(void* ptr)
{
	return static_cast<QMediaRecorderControl*>(ptr)->isMuted();
}

void QMediaRecorderControl_ConnectMutedChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(bool)>(&QMediaRecorderControl::mutedChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(bool)>(&MyQMediaRecorderControl::Signal_MutedChanged));
}

void QMediaRecorderControl_DisconnectMutedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(bool)>(&QMediaRecorderControl::mutedChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(bool)>(&MyQMediaRecorderControl::Signal_MutedChanged));
}

void QMediaRecorderControl_MutedChanged(void* ptr, char muted)
{
	static_cast<QMediaRecorderControl*>(ptr)->mutedChanged(muted != 0);
}

void* QMediaRecorderControl_OutputLocation(void* ptr)
{
	return new QUrl(static_cast<QMediaRecorderControl*>(ptr)->outputLocation());
}

void QMediaRecorderControl_SetMuted(void* ptr, char muted)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

char QMediaRecorderControl_SetOutputLocation(void* ptr, void* location)
{
	return static_cast<QMediaRecorderControl*>(ptr)->setOutputLocation(*static_cast<QUrl*>(location));
}

void QMediaRecorderControl_SetState(void* ptr, long long state)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setState", Q_ARG(QMediaRecorder::State, static_cast<QMediaRecorder::State>(state)));
}

void QMediaRecorderControl_SetVolume(void* ptr, double gain)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setVolume", Q_ARG(qreal, gain));
}

long long QMediaRecorderControl_State(void* ptr)
{
	return static_cast<QMediaRecorderControl*>(ptr)->state();
}

void QMediaRecorderControl_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::State)>(&QMediaRecorderControl::stateChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::State)>(&MyQMediaRecorderControl::Signal_StateChanged));
}

void QMediaRecorderControl_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::State)>(&QMediaRecorderControl::stateChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::State)>(&MyQMediaRecorderControl::Signal_StateChanged));
}

void QMediaRecorderControl_StateChanged(void* ptr, long long state)
{
	static_cast<QMediaRecorderControl*>(ptr)->stateChanged(static_cast<QMediaRecorder::State>(state));
}

long long QMediaRecorderControl_Status(void* ptr)
{
	return static_cast<QMediaRecorderControl*>(ptr)->status();
}

void QMediaRecorderControl_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::Status)>(&QMediaRecorderControl::statusChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::Status)>(&MyQMediaRecorderControl::Signal_StatusChanged));
}

void QMediaRecorderControl_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::Status)>(&QMediaRecorderControl::statusChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::Status)>(&MyQMediaRecorderControl::Signal_StatusChanged));
}

void QMediaRecorderControl_StatusChanged(void* ptr, long long status)
{
	static_cast<QMediaRecorderControl*>(ptr)->statusChanged(static_cast<QMediaRecorder::Status>(status));
}

double QMediaRecorderControl_Volume(void* ptr)
{
	return static_cast<QMediaRecorderControl*>(ptr)->volume();
}

void QMediaRecorderControl_ConnectVolumeChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(qreal)>(&QMediaRecorderControl::volumeChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(qreal)>(&MyQMediaRecorderControl::Signal_VolumeChanged));
}

void QMediaRecorderControl_DisconnectVolumeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(qreal)>(&QMediaRecorderControl::volumeChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(qreal)>(&MyQMediaRecorderControl::Signal_VolumeChanged));
}

void QMediaRecorderControl_VolumeChanged(void* ptr, double gain)
{
	static_cast<QMediaRecorderControl*>(ptr)->volumeChanged(gain);
}

void QMediaRecorderControl_DestroyQMediaRecorderControl(void* ptr)
{
	static_cast<QMediaRecorderControl*>(ptr)->~QMediaRecorderControl();
}

void QMediaRecorderControl_DestroyQMediaRecorderControlDefault(void* ptr)
{

}

void QMediaRecorderControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaRecorderControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaRecorderControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaRecorderControl*>(ptr)->QMediaRecorderControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaRecorderControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaRecorderControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaRecorderControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaRecorderControl*>(ptr)->QMediaRecorderControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaRecorderControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaRecorderControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaRecorderControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaRecorderControl*>(ptr)->QMediaRecorderControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaRecorderControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaRecorderControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaRecorderControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaRecorderControl*>(ptr)->QMediaRecorderControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaRecorderControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "deleteLater");
}

void QMediaRecorderControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaRecorderControl*>(ptr)->QMediaRecorderControl::deleteLater();
}

void QMediaRecorderControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaRecorderControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaRecorderControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaRecorderControl*>(ptr)->QMediaRecorderControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaRecorderControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaRecorderControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaRecorderControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaRecorderControl*>(ptr)->QMediaRecorderControl::event(static_cast<QEvent*>(e));
}

char QMediaRecorderControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaRecorderControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaRecorderControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaRecorderControl*>(ptr)->QMediaRecorderControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaRecorderControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaRecorderControl*>(ptr)->metaObject());
}

void* QMediaRecorderControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaRecorderControl*>(ptr)->QMediaRecorderControl::metaObject());
}

void* QMediaResource_NewQMediaResource()
{
	return new QMediaResource();
}

void* QMediaResource_NewQMediaResource4(void* other)
{
	return new QMediaResource(*static_cast<QMediaResource*>(other));
}

void* QMediaResource_NewQMediaResource3(void* request, char* mimeType)
{
	return new QMediaResource(*static_cast<QNetworkRequest*>(request), QString(mimeType));
}

void* QMediaResource_NewQMediaResource2(void* url, char* mimeType)
{
	return new QMediaResource(*static_cast<QUrl*>(url), QString(mimeType));
}

int QMediaResource_AudioBitRate(void* ptr)
{
	return static_cast<QMediaResource*>(ptr)->audioBitRate();
}

struct QtMultimedia_PackedString QMediaResource_AudioCodec(void* ptr)
{
	return ({ QByteArray t3ad21c = static_cast<QMediaResource*>(ptr)->audioCodec().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t3ad21c.prepend("WHITESPACE").constData()+10), t3ad21c.size()-10 }; });
}

int QMediaResource_ChannelCount(void* ptr)
{
	return static_cast<QMediaResource*>(ptr)->channelCount();
}

long long QMediaResource_DataSize(void* ptr)
{
	return static_cast<QMediaResource*>(ptr)->dataSize();
}

char QMediaResource_IsNull(void* ptr)
{
	return static_cast<QMediaResource*>(ptr)->isNull();
}

struct QtMultimedia_PackedString QMediaResource_Language(void* ptr)
{
	return ({ QByteArray t69ac0c = static_cast<QMediaResource*>(ptr)->language().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t69ac0c.prepend("WHITESPACE").constData()+10), t69ac0c.size()-10 }; });
}

struct QtMultimedia_PackedString QMediaResource_MimeType(void* ptr)
{
	return ({ QByteArray ta44245 = static_cast<QMediaResource*>(ptr)->mimeType().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(ta44245.prepend("WHITESPACE").constData()+10), ta44245.size()-10 }; });
}

void* QMediaResource_Request(void* ptr)
{
	return new QNetworkRequest(static_cast<QMediaResource*>(ptr)->request());
}

void* QMediaResource_Resolution(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QMediaResource*>(ptr)->resolution(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QMediaResource_SampleRate(void* ptr)
{
	return static_cast<QMediaResource*>(ptr)->sampleRate();
}

void QMediaResource_SetAudioBitRate(void* ptr, int rate)
{
	static_cast<QMediaResource*>(ptr)->setAudioBitRate(rate);
}

void QMediaResource_SetAudioCodec(void* ptr, char* codec)
{
	static_cast<QMediaResource*>(ptr)->setAudioCodec(QString(codec));
}

void QMediaResource_SetChannelCount(void* ptr, int channels)
{
	static_cast<QMediaResource*>(ptr)->setChannelCount(channels);
}

void QMediaResource_SetDataSize(void* ptr, long long size)
{
	static_cast<QMediaResource*>(ptr)->setDataSize(size);
}

void QMediaResource_SetLanguage(void* ptr, char* language)
{
	static_cast<QMediaResource*>(ptr)->setLanguage(QString(language));
}

void QMediaResource_SetResolution(void* ptr, void* resolution)
{
	static_cast<QMediaResource*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QMediaResource_SetResolution2(void* ptr, int width, int height)
{
	static_cast<QMediaResource*>(ptr)->setResolution(width, height);
}

void QMediaResource_SetSampleRate(void* ptr, int sampleRate)
{
	static_cast<QMediaResource*>(ptr)->setSampleRate(sampleRate);
}

void QMediaResource_SetVideoBitRate(void* ptr, int rate)
{
	static_cast<QMediaResource*>(ptr)->setVideoBitRate(rate);
}

void QMediaResource_SetVideoCodec(void* ptr, char* codec)
{
	static_cast<QMediaResource*>(ptr)->setVideoCodec(QString(codec));
}

void* QMediaResource_Url(void* ptr)
{
	return new QUrl(static_cast<QMediaResource*>(ptr)->url());
}

int QMediaResource_VideoBitRate(void* ptr)
{
	return static_cast<QMediaResource*>(ptr)->videoBitRate();
}

struct QtMultimedia_PackedString QMediaResource_VideoCodec(void* ptr)
{
	return ({ QByteArray tfe661d = static_cast<QMediaResource*>(ptr)->videoCodec().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tfe661d.prepend("WHITESPACE").constData()+10), tfe661d.size()-10 }; });
}

void QMediaResource_DestroyQMediaResource(void* ptr)
{
	static_cast<QMediaResource*>(ptr)->~QMediaResource();
}

class MyQMediaService: public QMediaService
{
public:
	MyQMediaService(QObject *parent) : QMediaService(parent) {};
	void releaseControl(QMediaControl * control) { callbackQMediaService_ReleaseControl(this, control); };
	QMediaControl * requestControl(const char * interfa) { QtMultimedia_PackedString interfaPacked = { const_cast<char*>(interfa), -1 };return static_cast<QMediaControl*>(callbackQMediaService_RequestControl(this, interfaPacked)); };
	void timerEvent(QTimerEvent * event) { callbackQMediaService_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaService_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaService_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaService_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaService_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaService_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaService_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaService_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaService_MetaObject(const_cast<MyQMediaService*>(this))); };
};

void* QMediaService_NewQMediaService(void* parent)
{
	return new MyQMediaService(static_cast<QObject*>(parent));
}

void QMediaService_ReleaseControl(void* ptr, void* control)
{
	static_cast<QMediaService*>(ptr)->releaseControl(static_cast<QMediaControl*>(control));
}

void* QMediaService_RequestControl(void* ptr, char* interfa)
{
	return static_cast<QMediaService*>(ptr)->requestControl(const_cast<const char*>(interfa));
}

void* QMediaService_RequestControl2(void* ptr)
{
	return static_cast<QMediaService*>(ptr)->requestControl<QMediaService*>();
}

void QMediaService_DestroyQMediaService(void* ptr)
{
	static_cast<QMediaService*>(ptr)->~QMediaService();
}

void QMediaService_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaService*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaService_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaService*>(ptr)->QMediaService::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaService_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaService*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaService_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaService*>(ptr)->QMediaService::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaService_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaService*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaService_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaService*>(ptr)->QMediaService::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaService_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaService*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaService_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaService*>(ptr)->QMediaService::customEvent(static_cast<QEvent*>(event));
}

void QMediaService_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaService*>(ptr), "deleteLater");
}

void QMediaService_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaService*>(ptr)->QMediaService::deleteLater();
}

void QMediaService_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaService*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaService_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaService*>(ptr)->QMediaService::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaService_Event(void* ptr, void* e)
{
	return static_cast<QMediaService*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaService_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaService*>(ptr)->QMediaService::event(static_cast<QEvent*>(e));
}

char QMediaService_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaService*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaService_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaService*>(ptr)->QMediaService::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaService_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaService*>(ptr)->metaObject());
}

void* QMediaService_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaService*>(ptr)->QMediaService::metaObject());
}

class MyQMediaServiceCameraInfoInterface: public QMediaServiceCameraInfoInterface
{
public:
	int cameraOrientation(const QByteArray & device) const { return callbackQMediaServiceCameraInfoInterface_CameraOrientation(const_cast<MyQMediaServiceCameraInfoInterface*>(this), const_cast<QByteArray*>(&device)); };
	QCamera::Position cameraPosition(const QByteArray & device) const { return static_cast<QCamera::Position>(callbackQMediaServiceCameraInfoInterface_CameraPosition(const_cast<MyQMediaServiceCameraInfoInterface*>(this), const_cast<QByteArray*>(&device))); };
	 ~MyQMediaServiceCameraInfoInterface() { callbackQMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterface(this); };
};

int QMediaServiceCameraInfoInterface_CameraOrientation(void* ptr, void* device)
{
	return static_cast<QMediaServiceCameraInfoInterface*>(ptr)->cameraOrientation(*static_cast<QByteArray*>(device));
}

long long QMediaServiceCameraInfoInterface_CameraPosition(void* ptr, void* device)
{
	return static_cast<QMediaServiceCameraInfoInterface*>(ptr)->cameraPosition(*static_cast<QByteArray*>(device));
}

void QMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterface(void* ptr)
{
	static_cast<QMediaServiceCameraInfoInterface*>(ptr)->~QMediaServiceCameraInfoInterface();
}

void QMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterfaceDefault(void* ptr)
{

}

class MyQMediaServiceDefaultDeviceInterface: public QMediaServiceDefaultDeviceInterface
{
public:
	QByteArray defaultDevice(const QByteArray & service) const { return *static_cast<QByteArray*>(callbackQMediaServiceDefaultDeviceInterface_DefaultDevice(const_cast<MyQMediaServiceDefaultDeviceInterface*>(this), const_cast<QByteArray*>(&service))); };
	 ~MyQMediaServiceDefaultDeviceInterface() { callbackQMediaServiceDefaultDeviceInterface_DestroyQMediaServiceDefaultDeviceInterface(this); };
};

void* QMediaServiceDefaultDeviceInterface_DefaultDevice(void* ptr, void* service)
{
	return new QByteArray(static_cast<QMediaServiceDefaultDeviceInterface*>(ptr)->defaultDevice(*static_cast<QByteArray*>(service)));
}

void QMediaServiceDefaultDeviceInterface_DestroyQMediaServiceDefaultDeviceInterface(void* ptr)
{
	static_cast<QMediaServiceDefaultDeviceInterface*>(ptr)->~QMediaServiceDefaultDeviceInterface();
}

void QMediaServiceDefaultDeviceInterface_DestroyQMediaServiceDefaultDeviceInterfaceDefault(void* ptr)
{

}

class MyQMediaServiceFeaturesInterface: public QMediaServiceFeaturesInterface
{
public:
	 ~MyQMediaServiceFeaturesInterface() { callbackQMediaServiceFeaturesInterface_DestroyQMediaServiceFeaturesInterface(this); };
};

void QMediaServiceFeaturesInterface_DestroyQMediaServiceFeaturesInterface(void* ptr)
{
	static_cast<QMediaServiceFeaturesInterface*>(ptr)->~QMediaServiceFeaturesInterface();
}

void QMediaServiceFeaturesInterface_DestroyQMediaServiceFeaturesInterfaceDefault(void* ptr)
{

}

class MyQMediaServiceProviderPlugin: public QMediaServiceProviderPlugin
{
public:
	QMediaService * create(const QString & key) { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };return static_cast<QMediaService*>(callbackQMediaServiceProviderPlugin_Create(this, keyPacked)); };
	void release(QMediaService * service) { callbackQMediaServiceProviderPlugin_Release(this, service); };
	void timerEvent(QTimerEvent * event) { callbackQMediaServiceProviderPlugin_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaServiceProviderPlugin_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaServiceProviderPlugin_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaServiceProviderPlugin_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaServiceProviderPlugin_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaServiceProviderPlugin_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaServiceProviderPlugin_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaServiceProviderPlugin_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaServiceProviderPlugin_MetaObject(const_cast<MyQMediaServiceProviderPlugin*>(this))); };
};

void* QMediaServiceProviderPlugin_Create(void* ptr, char* key)
{
	return static_cast<QMediaServiceProviderPlugin*>(ptr)->create(QString(key));
}

void QMediaServiceProviderPlugin_Release(void* ptr, void* service)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->release(static_cast<QMediaService*>(service));
}

void QMediaServiceProviderPlugin_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaServiceProviderPlugin_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->QMediaServiceProviderPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaServiceProviderPlugin_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaServiceProviderPlugin_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->QMediaServiceProviderPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaServiceProviderPlugin_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaServiceProviderPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->QMediaServiceProviderPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaServiceProviderPlugin_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaServiceProviderPlugin_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->QMediaServiceProviderPlugin::customEvent(static_cast<QEvent*>(event));
}

void QMediaServiceProviderPlugin_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaServiceProviderPlugin*>(ptr), "deleteLater");
}

void QMediaServiceProviderPlugin_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->QMediaServiceProviderPlugin::deleteLater();
}

void QMediaServiceProviderPlugin_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaServiceProviderPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaServiceProviderPlugin*>(ptr)->QMediaServiceProviderPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaServiceProviderPlugin_Event(void* ptr, void* e)
{
	return static_cast<QMediaServiceProviderPlugin*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaServiceProviderPlugin_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaServiceProviderPlugin*>(ptr)->QMediaServiceProviderPlugin::event(static_cast<QEvent*>(e));
}

char QMediaServiceProviderPlugin_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaServiceProviderPlugin*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaServiceProviderPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaServiceProviderPlugin*>(ptr)->QMediaServiceProviderPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaServiceProviderPlugin_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaServiceProviderPlugin*>(ptr)->metaObject());
}

void* QMediaServiceProviderPlugin_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaServiceProviderPlugin*>(ptr)->QMediaServiceProviderPlugin::metaObject());
}

class MyQMediaServiceSupportedDevicesInterface: public QMediaServiceSupportedDevicesInterface
{
public:
	QString deviceDescription(const QByteArray & service, const QByteArray & device) { return QString(callbackQMediaServiceSupportedDevicesInterface_DeviceDescription(this, const_cast<QByteArray*>(&service), const_cast<QByteArray*>(&device))); };
	 ~MyQMediaServiceSupportedDevicesInterface() { callbackQMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterface(this); };
};

struct QtMultimedia_PackedString QMediaServiceSupportedDevicesInterface_DeviceDescription(void* ptr, void* service, void* device)
{
	return ({ QByteArray tc3bd23 = static_cast<QMediaServiceSupportedDevicesInterface*>(ptr)->deviceDescription(*static_cast<QByteArray*>(service), *static_cast<QByteArray*>(device)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tc3bd23.prepend("WHITESPACE").constData()+10), tc3bd23.size()-10 }; });
}

void QMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterface(void* ptr)
{
	static_cast<QMediaServiceSupportedDevicesInterface*>(ptr)->~QMediaServiceSupportedDevicesInterface();
}

void QMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterfaceDefault(void* ptr)
{

}

class MyQMediaServiceSupportedFormatsInterface: public QMediaServiceSupportedFormatsInterface
{
public:
	QMultimedia::SupportEstimate hasSupport(const QString & mimeType, const QStringList & codecs) const { QByteArray t3313b8 = mimeType.toUtf8(); QtMultimedia_PackedString mimeTypePacked = { const_cast<char*>(t3313b8.prepend("WHITESPACE").constData()+10), t3313b8.size()-10 };QByteArray t7222c0 = codecs.join("|").toUtf8(); QtMultimedia_PackedString codecsPacked = { const_cast<char*>(t7222c0.prepend("WHITESPACE").constData()+10), t7222c0.size()-10 };return static_cast<QMultimedia::SupportEstimate>(callbackQMediaServiceSupportedFormatsInterface_HasSupport(const_cast<MyQMediaServiceSupportedFormatsInterface*>(this), mimeTypePacked, codecsPacked)); };
	QStringList supportedMimeTypes() const { return QString(callbackQMediaServiceSupportedFormatsInterface_SupportedMimeTypes(const_cast<MyQMediaServiceSupportedFormatsInterface*>(this))).split("|", QString::SkipEmptyParts); };
	 ~MyQMediaServiceSupportedFormatsInterface() { callbackQMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterface(this); };
};

long long QMediaServiceSupportedFormatsInterface_HasSupport(void* ptr, char* mimeType, char* codecs)
{
	return static_cast<QMediaServiceSupportedFormatsInterface*>(ptr)->hasSupport(QString(mimeType), QString(codecs).split("|", QString::SkipEmptyParts));
}

struct QtMultimedia_PackedString QMediaServiceSupportedFormatsInterface_SupportedMimeTypes(void* ptr)
{
	return ({ QByteArray te6e5ca = static_cast<QMediaServiceSupportedFormatsInterface*>(ptr)->supportedMimeTypes().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(te6e5ca.prepend("WHITESPACE").constData()+10), te6e5ca.size()-10 }; });
}

void QMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterface(void* ptr)
{
	static_cast<QMediaServiceSupportedFormatsInterface*>(ptr)->~QMediaServiceSupportedFormatsInterface();
}

void QMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterfaceDefault(void* ptr)
{

}

class MyQMediaStreamsControl: public QMediaStreamsControl
{
public:
	MyQMediaStreamsControl(QObject *parent) : QMediaStreamsControl(parent) {};
	void Signal_ActiveStreamsChanged() { callbackQMediaStreamsControl_ActiveStreamsChanged(this); };
	bool isActive(int stream) { return callbackQMediaStreamsControl_IsActive(this, stream) != 0; };
	QVariant metaData(int stream, const QString & key) { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };return *static_cast<QVariant*>(callbackQMediaStreamsControl_MetaData(this, stream, keyPacked)); };
	void setActive(int stream, bool state) { callbackQMediaStreamsControl_SetActive(this, stream, state); };
	int streamCount() { return callbackQMediaStreamsControl_StreamCount(this); };
	StreamType streamType(int stream) { return static_cast<QMediaStreamsControl::StreamType>(callbackQMediaStreamsControl_StreamType(this, stream)); };
	void Signal_StreamsChanged() { callbackQMediaStreamsControl_StreamsChanged(this); };
	 ~MyQMediaStreamsControl() { callbackQMediaStreamsControl_DestroyQMediaStreamsControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQMediaStreamsControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaStreamsControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaStreamsControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaStreamsControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaStreamsControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaStreamsControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaStreamsControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaStreamsControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaStreamsControl_MetaObject(const_cast<MyQMediaStreamsControl*>(this))); };
};

void* QMediaStreamsControl_NewQMediaStreamsControl(void* parent)
{
	return new MyQMediaStreamsControl(static_cast<QObject*>(parent));
}

void QMediaStreamsControl_ConnectActiveStreamsChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::activeStreamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_ActiveStreamsChanged));
}

void QMediaStreamsControl_DisconnectActiveStreamsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::activeStreamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_ActiveStreamsChanged));
}

void QMediaStreamsControl_ActiveStreamsChanged(void* ptr)
{
	static_cast<QMediaStreamsControl*>(ptr)->activeStreamsChanged();
}

char QMediaStreamsControl_IsActive(void* ptr, int stream)
{
	return static_cast<QMediaStreamsControl*>(ptr)->isActive(stream);
}

void* QMediaStreamsControl_MetaData(void* ptr, int stream, char* key)
{
	return new QVariant(static_cast<QMediaStreamsControl*>(ptr)->metaData(stream, QString(key)));
}

void QMediaStreamsControl_SetActive(void* ptr, int stream, char state)
{
	static_cast<QMediaStreamsControl*>(ptr)->setActive(stream, state != 0);
}

int QMediaStreamsControl_StreamCount(void* ptr)
{
	return static_cast<QMediaStreamsControl*>(ptr)->streamCount();
}

long long QMediaStreamsControl_StreamType(void* ptr, int stream)
{
	return static_cast<QMediaStreamsControl*>(ptr)->streamType(stream);
}

void QMediaStreamsControl_ConnectStreamsChanged(void* ptr)
{
	QObject::connect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::streamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_StreamsChanged));
}

void QMediaStreamsControl_DisconnectStreamsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::streamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_StreamsChanged));
}

void QMediaStreamsControl_StreamsChanged(void* ptr)
{
	static_cast<QMediaStreamsControl*>(ptr)->streamsChanged();
}

void QMediaStreamsControl_DestroyQMediaStreamsControl(void* ptr)
{
	static_cast<QMediaStreamsControl*>(ptr)->~QMediaStreamsControl();
}

void QMediaStreamsControl_DestroyQMediaStreamsControlDefault(void* ptr)
{

}

void QMediaStreamsControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaStreamsControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaStreamsControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaStreamsControl*>(ptr)->QMediaStreamsControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaStreamsControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaStreamsControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaStreamsControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaStreamsControl*>(ptr)->QMediaStreamsControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaStreamsControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaStreamsControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaStreamsControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaStreamsControl*>(ptr)->QMediaStreamsControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaStreamsControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaStreamsControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaStreamsControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaStreamsControl*>(ptr)->QMediaStreamsControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaStreamsControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaStreamsControl*>(ptr), "deleteLater");
}

void QMediaStreamsControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaStreamsControl*>(ptr)->QMediaStreamsControl::deleteLater();
}

void QMediaStreamsControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaStreamsControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaStreamsControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaStreamsControl*>(ptr)->QMediaStreamsControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaStreamsControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaStreamsControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaStreamsControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaStreamsControl*>(ptr)->QMediaStreamsControl::event(static_cast<QEvent*>(e));
}

char QMediaStreamsControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaStreamsControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaStreamsControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaStreamsControl*>(ptr)->QMediaStreamsControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaStreamsControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaStreamsControl*>(ptr)->metaObject());
}

void* QMediaStreamsControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaStreamsControl*>(ptr)->QMediaStreamsControl::metaObject());
}

void* QMediaTimeInterval_NewQMediaTimeInterval()
{
	return new QMediaTimeInterval();
}

void* QMediaTimeInterval_NewQMediaTimeInterval3(void* other)
{
	return new QMediaTimeInterval(*static_cast<QMediaTimeInterval*>(other));
}

void* QMediaTimeInterval_NewQMediaTimeInterval2(long long start, long long end)
{
	return new QMediaTimeInterval(start, end);
}

char QMediaTimeInterval_Contains(void* ptr, long long time)
{
	return static_cast<QMediaTimeInterval*>(ptr)->contains(time);
}

long long QMediaTimeInterval_End(void* ptr)
{
	return static_cast<QMediaTimeInterval*>(ptr)->end();
}

char QMediaTimeInterval_IsNormal(void* ptr)
{
	return static_cast<QMediaTimeInterval*>(ptr)->isNormal();
}

void* QMediaTimeInterval_Normalized(void* ptr)
{
	return new QMediaTimeInterval(static_cast<QMediaTimeInterval*>(ptr)->normalized());
}

long long QMediaTimeInterval_Start(void* ptr)
{
	return static_cast<QMediaTimeInterval*>(ptr)->start();
}

void* QMediaTimeInterval_Translated(void* ptr, long long offset)
{
	return new QMediaTimeInterval(static_cast<QMediaTimeInterval*>(ptr)->translated(offset));
}

void* QMediaTimeRange_NewQMediaTimeRange()
{
	return new QMediaTimeRange();
}

void* QMediaTimeRange_NewQMediaTimeRange3(void* interval)
{
	return new QMediaTimeRange(*static_cast<QMediaTimeInterval*>(interval));
}

void* QMediaTimeRange_NewQMediaTimeRange4(void* ran)
{
	return new QMediaTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void* QMediaTimeRange_NewQMediaTimeRange2(long long start, long long end)
{
	return new QMediaTimeRange(start, end);
}

void QMediaTimeRange_AddInterval(void* ptr, void* interval)
{
	static_cast<QMediaTimeRange*>(ptr)->addInterval(*static_cast<QMediaTimeInterval*>(interval));
}

void QMediaTimeRange_AddInterval2(void* ptr, long long start, long long end)
{
	static_cast<QMediaTimeRange*>(ptr)->addInterval(start, end);
}

void QMediaTimeRange_AddTimeRange(void* ptr, void* ran)
{
	static_cast<QMediaTimeRange*>(ptr)->addTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_Clear(void* ptr)
{
	static_cast<QMediaTimeRange*>(ptr)->clear();
}

char QMediaTimeRange_Contains(void* ptr, long long time)
{
	return static_cast<QMediaTimeRange*>(ptr)->contains(time);
}

long long QMediaTimeRange_EarliestTime(void* ptr)
{
	return static_cast<QMediaTimeRange*>(ptr)->earliestTime();
}

char QMediaTimeRange_IsContinuous(void* ptr)
{
	return static_cast<QMediaTimeRange*>(ptr)->isContinuous();
}

char QMediaTimeRange_IsEmpty(void* ptr)
{
	return static_cast<QMediaTimeRange*>(ptr)->isEmpty();
}

long long QMediaTimeRange_LatestTime(void* ptr)
{
	return static_cast<QMediaTimeRange*>(ptr)->latestTime();
}

void QMediaTimeRange_RemoveInterval(void* ptr, void* interval)
{
	static_cast<QMediaTimeRange*>(ptr)->removeInterval(*static_cast<QMediaTimeInterval*>(interval));
}

void QMediaTimeRange_RemoveInterval2(void* ptr, long long start, long long end)
{
	static_cast<QMediaTimeRange*>(ptr)->removeInterval(start, end);
}

void QMediaTimeRange_RemoveTimeRange(void* ptr, void* ran)
{
	static_cast<QMediaTimeRange*>(ptr)->removeTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_DestroyQMediaTimeRange(void* ptr)
{
	static_cast<QMediaTimeRange*>(ptr)->~QMediaTimeRange();
}

class MyQMediaVideoProbeControl: public QMediaVideoProbeControl
{
public:
	MyQMediaVideoProbeControl(QObject *parent) : QMediaVideoProbeControl(parent) {};
	void Signal_Flush() { callbackQMediaVideoProbeControl_Flush(this); };
	void Signal_VideoFrameProbed(const QVideoFrame & frame) { callbackQMediaVideoProbeControl_VideoFrameProbed(this, const_cast<QVideoFrame*>(&frame)); };
	 ~MyQMediaVideoProbeControl() { callbackQMediaVideoProbeControl_DestroyQMediaVideoProbeControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQMediaVideoProbeControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMediaVideoProbeControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMediaVideoProbeControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMediaVideoProbeControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMediaVideoProbeControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMediaVideoProbeControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMediaVideoProbeControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMediaVideoProbeControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMediaVideoProbeControl_MetaObject(const_cast<MyQMediaVideoProbeControl*>(this))); };
};

void* QMediaVideoProbeControl_NewQMediaVideoProbeControl(void* parent)
{
	return new MyQMediaVideoProbeControl(static_cast<QObject*>(parent));
}

void QMediaVideoProbeControl_ConnectFlush(void* ptr)
{
	QObject::connect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)()>(&QMediaVideoProbeControl::flush), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)()>(&MyQMediaVideoProbeControl::Signal_Flush));
}

void QMediaVideoProbeControl_DisconnectFlush(void* ptr)
{
	QObject::disconnect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)()>(&QMediaVideoProbeControl::flush), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)()>(&MyQMediaVideoProbeControl::Signal_Flush));
}

void QMediaVideoProbeControl_Flush(void* ptr)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->flush();
}

void QMediaVideoProbeControl_ConnectVideoFrameProbed(void* ptr)
{
	QObject::connect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)(const QVideoFrame &)>(&QMediaVideoProbeControl::videoFrameProbed), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)(const QVideoFrame &)>(&MyQMediaVideoProbeControl::Signal_VideoFrameProbed));
}

void QMediaVideoProbeControl_DisconnectVideoFrameProbed(void* ptr)
{
	QObject::disconnect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)(const QVideoFrame &)>(&QMediaVideoProbeControl::videoFrameProbed), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)(const QVideoFrame &)>(&MyQMediaVideoProbeControl::Signal_VideoFrameProbed));
}

void QMediaVideoProbeControl_VideoFrameProbed(void* ptr, void* frame)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->videoFrameProbed(*static_cast<QVideoFrame*>(frame));
}

void QMediaVideoProbeControl_DestroyQMediaVideoProbeControl(void* ptr)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->~QMediaVideoProbeControl();
}

void QMediaVideoProbeControl_DestroyQMediaVideoProbeControlDefault(void* ptr)
{

}

void QMediaVideoProbeControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaVideoProbeControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->QMediaVideoProbeControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMediaVideoProbeControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMediaVideoProbeControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->QMediaVideoProbeControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMediaVideoProbeControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaVideoProbeControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->QMediaVideoProbeControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaVideoProbeControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMediaVideoProbeControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->QMediaVideoProbeControl::customEvent(static_cast<QEvent*>(event));
}

void QMediaVideoProbeControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMediaVideoProbeControl*>(ptr), "deleteLater");
}

void QMediaVideoProbeControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->QMediaVideoProbeControl::deleteLater();
}

void QMediaVideoProbeControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMediaVideoProbeControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMediaVideoProbeControl*>(ptr)->QMediaVideoProbeControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMediaVideoProbeControl_Event(void* ptr, void* e)
{
	return static_cast<QMediaVideoProbeControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMediaVideoProbeControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMediaVideoProbeControl*>(ptr)->QMediaVideoProbeControl::event(static_cast<QEvent*>(e));
}

char QMediaVideoProbeControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaVideoProbeControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMediaVideoProbeControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMediaVideoProbeControl*>(ptr)->QMediaVideoProbeControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMediaVideoProbeControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaVideoProbeControl*>(ptr)->metaObject());
}

void* QMediaVideoProbeControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMediaVideoProbeControl*>(ptr)->QMediaVideoProbeControl::metaObject());
}

class MyQMetaDataReaderControl: public QMetaDataReaderControl
{
public:
	MyQMetaDataReaderControl(QObject *parent) : QMetaDataReaderControl(parent) {};
	QStringList availableMetaData() const { return QString(callbackQMetaDataReaderControl_AvailableMetaData(const_cast<MyQMetaDataReaderControl*>(this))).split("|", QString::SkipEmptyParts); };
	bool isMetaDataAvailable() const { return callbackQMetaDataReaderControl_IsMetaDataAvailable(const_cast<MyQMetaDataReaderControl*>(this)) != 0; };
	QVariant metaData(const QString & key) const { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };return *static_cast<QVariant*>(callbackQMetaDataReaderControl_MetaData(const_cast<MyQMetaDataReaderControl*>(this), keyPacked)); };
	void Signal_MetaDataAvailableChanged(bool available) { callbackQMetaDataReaderControl_MetaDataAvailableChanged(this, available); };
	void Signal_MetaDataChanged() { callbackQMetaDataReaderControl_MetaDataChanged(this); };
	void Signal_MetaDataChanged2(const QString & key, const QVariant & value) { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };callbackQMetaDataReaderControl_MetaDataChanged2(this, keyPacked, const_cast<QVariant*>(&value)); };
	void timerEvent(QTimerEvent * event) { callbackQMetaDataReaderControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMetaDataReaderControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMetaDataReaderControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMetaDataReaderControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMetaDataReaderControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMetaDataReaderControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMetaDataReaderControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMetaDataReaderControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMetaDataReaderControl_MetaObject(const_cast<MyQMetaDataReaderControl*>(this))); };
};

void* QMetaDataReaderControl_NewQMetaDataReaderControl(void* parent)
{
	return new MyQMetaDataReaderControl(static_cast<QObject*>(parent));
}

struct QtMultimedia_PackedString QMetaDataReaderControl_AvailableMetaData(void* ptr)
{
	return ({ QByteArray tff69ae = static_cast<QMetaDataReaderControl*>(ptr)->availableMetaData().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tff69ae.prepend("WHITESPACE").constData()+10), tff69ae.size()-10 }; });
}

char QMetaDataReaderControl_IsMetaDataAvailable(void* ptr)
{
	return static_cast<QMetaDataReaderControl*>(ptr)->isMetaDataAvailable();
}

void* QMetaDataReaderControl_MetaData(void* ptr, char* key)
{
	return new QVariant(static_cast<QMetaDataReaderControl*>(ptr)->metaData(QString(key)));
}

void QMetaDataReaderControl_ConnectMetaDataAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(bool)>(&QMetaDataReaderControl::metaDataAvailableChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(bool)>(&MyQMetaDataReaderControl::Signal_MetaDataAvailableChanged));
}

void QMetaDataReaderControl_DisconnectMetaDataAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(bool)>(&QMetaDataReaderControl::metaDataAvailableChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(bool)>(&MyQMetaDataReaderControl::Signal_MetaDataAvailableChanged));
}

void QMetaDataReaderControl_MetaDataAvailableChanged(void* ptr, char available)
{
	static_cast<QMetaDataReaderControl*>(ptr)->metaDataAvailableChanged(available != 0);
}

void QMetaDataReaderControl_ConnectMetaDataChanged(void* ptr)
{
	QObject::connect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)()>(&MyQMetaDataReaderControl::Signal_MetaDataChanged));
}

void QMetaDataReaderControl_DisconnectMetaDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)()>(&MyQMetaDataReaderControl::Signal_MetaDataChanged));
}

void QMetaDataReaderControl_MetaDataChanged(void* ptr)
{
	static_cast<QMetaDataReaderControl*>(ptr)->metaDataChanged();
}

void QMetaDataReaderControl_ConnectMetaDataChanged2(void* ptr)
{
	QObject::connect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(const QString &, const QVariant &)>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(const QString &, const QVariant &)>(&MyQMetaDataReaderControl::Signal_MetaDataChanged2));
}

void QMetaDataReaderControl_DisconnectMetaDataChanged2(void* ptr)
{
	QObject::disconnect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(const QString &, const QVariant &)>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(const QString &, const QVariant &)>(&MyQMetaDataReaderControl::Signal_MetaDataChanged2));
}

void QMetaDataReaderControl_MetaDataChanged2(void* ptr, char* key, void* value)
{
	static_cast<QMetaDataReaderControl*>(ptr)->metaDataChanged(QString(key), *static_cast<QVariant*>(value));
}

void QMetaDataReaderControl_DestroyQMetaDataReaderControl(void* ptr)
{
	static_cast<QMetaDataReaderControl*>(ptr)->~QMetaDataReaderControl();
}

void QMetaDataReaderControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMetaDataReaderControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMetaDataReaderControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMetaDataReaderControl*>(ptr)->QMetaDataReaderControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMetaDataReaderControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMetaDataReaderControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMetaDataReaderControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMetaDataReaderControl*>(ptr)->QMetaDataReaderControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMetaDataReaderControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMetaDataReaderControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMetaDataReaderControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMetaDataReaderControl*>(ptr)->QMetaDataReaderControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMetaDataReaderControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMetaDataReaderControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMetaDataReaderControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMetaDataReaderControl*>(ptr)->QMetaDataReaderControl::customEvent(static_cast<QEvent*>(event));
}

void QMetaDataReaderControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMetaDataReaderControl*>(ptr), "deleteLater");
}

void QMetaDataReaderControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMetaDataReaderControl*>(ptr)->QMetaDataReaderControl::deleteLater();
}

void QMetaDataReaderControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMetaDataReaderControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMetaDataReaderControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMetaDataReaderControl*>(ptr)->QMetaDataReaderControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMetaDataReaderControl_Event(void* ptr, void* e)
{
	return static_cast<QMetaDataReaderControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMetaDataReaderControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMetaDataReaderControl*>(ptr)->QMetaDataReaderControl::event(static_cast<QEvent*>(e));
}

char QMetaDataReaderControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMetaDataReaderControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMetaDataReaderControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMetaDataReaderControl*>(ptr)->QMetaDataReaderControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMetaDataReaderControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMetaDataReaderControl*>(ptr)->metaObject());
}

void* QMetaDataReaderControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMetaDataReaderControl*>(ptr)->QMetaDataReaderControl::metaObject());
}

class MyQMetaDataWriterControl: public QMetaDataWriterControl
{
public:
	MyQMetaDataWriterControl(QObject *parent) : QMetaDataWriterControl(parent) {};
	QStringList availableMetaData() const { return QString(callbackQMetaDataWriterControl_AvailableMetaData(const_cast<MyQMetaDataWriterControl*>(this))).split("|", QString::SkipEmptyParts); };
	bool isMetaDataAvailable() const { return callbackQMetaDataWriterControl_IsMetaDataAvailable(const_cast<MyQMetaDataWriterControl*>(this)) != 0; };
	bool isWritable() const { return callbackQMetaDataWriterControl_IsWritable(const_cast<MyQMetaDataWriterControl*>(this)) != 0; };
	QVariant metaData(const QString & key) const { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };return *static_cast<QVariant*>(callbackQMetaDataWriterControl_MetaData(const_cast<MyQMetaDataWriterControl*>(this), keyPacked)); };
	void Signal_MetaDataAvailableChanged(bool available) { callbackQMetaDataWriterControl_MetaDataAvailableChanged(this, available); };
	void Signal_MetaDataChanged() { callbackQMetaDataWriterControl_MetaDataChanged(this); };
	void Signal_MetaDataChanged2(const QString & key, const QVariant & value) { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };callbackQMetaDataWriterControl_MetaDataChanged2(this, keyPacked, const_cast<QVariant*>(&value)); };
	void setMetaData(const QString & key, const QVariant & value) { QByteArray ta62f22 = key.toUtf8(); QtMultimedia_PackedString keyPacked = { const_cast<char*>(ta62f22.prepend("WHITESPACE").constData()+10), ta62f22.size()-10 };callbackQMetaDataWriterControl_SetMetaData(this, keyPacked, const_cast<QVariant*>(&value)); };
	void Signal_WritableChanged(bool writable) { callbackQMetaDataWriterControl_WritableChanged(this, writable); };
	void timerEvent(QTimerEvent * event) { callbackQMetaDataWriterControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQMetaDataWriterControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMetaDataWriterControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQMetaDataWriterControl_CustomEvent(this, event); };
	void deleteLater() { callbackQMetaDataWriterControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMetaDataWriterControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQMetaDataWriterControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMetaDataWriterControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMetaDataWriterControl_MetaObject(const_cast<MyQMetaDataWriterControl*>(this))); };
};

void* QMetaDataWriterControl_NewQMetaDataWriterControl(void* parent)
{
	return new MyQMetaDataWriterControl(static_cast<QObject*>(parent));
}

struct QtMultimedia_PackedString QMetaDataWriterControl_AvailableMetaData(void* ptr)
{
	return ({ QByteArray te3166d = static_cast<QMetaDataWriterControl*>(ptr)->availableMetaData().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(te3166d.prepend("WHITESPACE").constData()+10), te3166d.size()-10 }; });
}

char QMetaDataWriterControl_IsMetaDataAvailable(void* ptr)
{
	return static_cast<QMetaDataWriterControl*>(ptr)->isMetaDataAvailable();
}

char QMetaDataWriterControl_IsWritable(void* ptr)
{
	return static_cast<QMetaDataWriterControl*>(ptr)->isWritable();
}

void* QMetaDataWriterControl_MetaData(void* ptr, char* key)
{
	return new QVariant(static_cast<QMetaDataWriterControl*>(ptr)->metaData(QString(key)));
}

void QMetaDataWriterControl_ConnectMetaDataAvailableChanged(void* ptr)
{
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::metaDataAvailableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_MetaDataAvailableChanged));
}

void QMetaDataWriterControl_DisconnectMetaDataAvailableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::metaDataAvailableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_MetaDataAvailableChanged));
}

void QMetaDataWriterControl_MetaDataAvailableChanged(void* ptr, char available)
{
	static_cast<QMetaDataWriterControl*>(ptr)->metaDataAvailableChanged(available != 0);
}

void QMetaDataWriterControl_ConnectMetaDataChanged(void* ptr)
{
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)()>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)()>(&MyQMetaDataWriterControl::Signal_MetaDataChanged));
}

void QMetaDataWriterControl_DisconnectMetaDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)()>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)()>(&MyQMetaDataWriterControl::Signal_MetaDataChanged));
}

void QMetaDataWriterControl_MetaDataChanged(void* ptr)
{
	static_cast<QMetaDataWriterControl*>(ptr)->metaDataChanged();
}

void QMetaDataWriterControl_ConnectMetaDataChanged2(void* ptr)
{
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(const QString &, const QVariant &)>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(const QString &, const QVariant &)>(&MyQMetaDataWriterControl::Signal_MetaDataChanged2));
}

void QMetaDataWriterControl_DisconnectMetaDataChanged2(void* ptr)
{
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(const QString &, const QVariant &)>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(const QString &, const QVariant &)>(&MyQMetaDataWriterControl::Signal_MetaDataChanged2));
}

void QMetaDataWriterControl_MetaDataChanged2(void* ptr, char* key, void* value)
{
	static_cast<QMetaDataWriterControl*>(ptr)->metaDataChanged(QString(key), *static_cast<QVariant*>(value));
}

void QMetaDataWriterControl_SetMetaData(void* ptr, char* key, void* value)
{
	static_cast<QMetaDataWriterControl*>(ptr)->setMetaData(QString(key), *static_cast<QVariant*>(value));
}

void QMetaDataWriterControl_ConnectWritableChanged(void* ptr)
{
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::writableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_WritableChanged));
}

void QMetaDataWriterControl_DisconnectWritableChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::writableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_WritableChanged));
}

void QMetaDataWriterControl_WritableChanged(void* ptr, char writable)
{
	static_cast<QMetaDataWriterControl*>(ptr)->writableChanged(writable != 0);
}

void QMetaDataWriterControl_DestroyQMetaDataWriterControl(void* ptr)
{
	static_cast<QMetaDataWriterControl*>(ptr)->~QMetaDataWriterControl();
}

void QMetaDataWriterControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QMetaDataWriterControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMetaDataWriterControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMetaDataWriterControl*>(ptr)->QMetaDataWriterControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMetaDataWriterControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QMetaDataWriterControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMetaDataWriterControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMetaDataWriterControl*>(ptr)->QMetaDataWriterControl::childEvent(static_cast<QChildEvent*>(event));
}

void QMetaDataWriterControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMetaDataWriterControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMetaDataWriterControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMetaDataWriterControl*>(ptr)->QMetaDataWriterControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMetaDataWriterControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QMetaDataWriterControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMetaDataWriterControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMetaDataWriterControl*>(ptr)->QMetaDataWriterControl::customEvent(static_cast<QEvent*>(event));
}

void QMetaDataWriterControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMetaDataWriterControl*>(ptr), "deleteLater");
}

void QMetaDataWriterControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QMetaDataWriterControl*>(ptr)->QMetaDataWriterControl::deleteLater();
}

void QMetaDataWriterControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMetaDataWriterControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMetaDataWriterControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMetaDataWriterControl*>(ptr)->QMetaDataWriterControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QMetaDataWriterControl_Event(void* ptr, void* e)
{
	return static_cast<QMetaDataWriterControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QMetaDataWriterControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QMetaDataWriterControl*>(ptr)->QMetaDataWriterControl::event(static_cast<QEvent*>(e));
}

char QMetaDataWriterControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMetaDataWriterControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QMetaDataWriterControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMetaDataWriterControl*>(ptr)->QMetaDataWriterControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMetaDataWriterControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMetaDataWriterControl*>(ptr)->metaObject());
}

void* QMetaDataWriterControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMetaDataWriterControl*>(ptr)->QMetaDataWriterControl::metaObject());
}

class MyQRadioData: public QRadioData
{
public:
	MyQRadioData(QMediaObject *mediaObject, QObject *parent) : QRadioData(mediaObject, parent) {};
	void setAlternativeFrequenciesEnabled(bool enabled) { callbackQRadioData_SetAlternativeFrequenciesEnabled(this, enabled); };
	void Signal_AlternativeFrequenciesEnabledChanged(bool enabled) { callbackQRadioData_AlternativeFrequenciesEnabledChanged(this, enabled); };
	void Signal_Error2(QRadioData::Error error) { callbackQRadioData_Error2(this, error); };
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQRadioData_MediaObject(const_cast<MyQRadioData*>(this))); };
	void Signal_ProgramTypeChanged(QRadioData::ProgramType programType) { callbackQRadioData_ProgramTypeChanged(this, programType); };
	void Signal_ProgramTypeNameChanged(QString programTypeName) { QByteArray t2400f1 = programTypeName.toUtf8(); QtMultimedia_PackedString programTypeNamePacked = { const_cast<char*>(t2400f1.prepend("WHITESPACE").constData()+10), t2400f1.size()-10 };callbackQRadioData_ProgramTypeNameChanged(this, programTypeNamePacked); };
	void Signal_RadioTextChanged(QString radioText) { QByteArray t7ba492 = radioText.toUtf8(); QtMultimedia_PackedString radioTextPacked = { const_cast<char*>(t7ba492.prepend("WHITESPACE").constData()+10), t7ba492.size()-10 };callbackQRadioData_RadioTextChanged(this, radioTextPacked); };
	bool setMediaObject(QMediaObject * mediaObject) { return callbackQRadioData_SetMediaObject(this, mediaObject) != 0; };
	void Signal_StationIdChanged(QString stationId) { QByteArray t6518b0 = stationId.toUtf8(); QtMultimedia_PackedString stationIdPacked = { const_cast<char*>(t6518b0.prepend("WHITESPACE").constData()+10), t6518b0.size()-10 };callbackQRadioData_StationIdChanged(this, stationIdPacked); };
	void Signal_StationNameChanged(QString stationName) { QByteArray t32b816 = stationName.toUtf8(); QtMultimedia_PackedString stationNamePacked = { const_cast<char*>(t32b816.prepend("WHITESPACE").constData()+10), t32b816.size()-10 };callbackQRadioData_StationNameChanged(this, stationNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRadioData_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQRadioData_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRadioData_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRadioData_CustomEvent(this, event); };
	void deleteLater() { callbackQRadioData_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRadioData_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRadioData_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRadioData_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRadioData_MetaObject(const_cast<MyQRadioData*>(this))); };
};

char QRadioData_IsAlternativeFrequenciesEnabled(void* ptr)
{
	return static_cast<QRadioData*>(ptr)->isAlternativeFrequenciesEnabled();
}

long long QRadioData_ProgramType(void* ptr)
{
	return static_cast<QRadioData*>(ptr)->programType();
}

struct QtMultimedia_PackedString QRadioData_ProgramTypeName(void* ptr)
{
	return ({ QByteArray tf5ec6b = static_cast<QRadioData*>(ptr)->programTypeName().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tf5ec6b.prepend("WHITESPACE").constData()+10), tf5ec6b.size()-10 }; });
}

struct QtMultimedia_PackedString QRadioData_RadioText(void* ptr)
{
	return ({ QByteArray t809fb8 = static_cast<QRadioData*>(ptr)->radioText().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t809fb8.prepend("WHITESPACE").constData()+10), t809fb8.size()-10 }; });
}

void QRadioData_SetAlternativeFrequenciesEnabled(void* ptr, char enabled)
{
	QMetaObject::invokeMethod(static_cast<QRadioData*>(ptr), "setAlternativeFrequenciesEnabled", Q_ARG(bool, enabled != 0));
}

struct QtMultimedia_PackedString QRadioData_StationId(void* ptr)
{
	return ({ QByteArray t645596 = static_cast<QRadioData*>(ptr)->stationId().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t645596.prepend("WHITESPACE").constData()+10), t645596.size()-10 }; });
}

struct QtMultimedia_PackedString QRadioData_StationName(void* ptr)
{
	return ({ QByteArray t79f056 = static_cast<QRadioData*>(ptr)->stationName().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t79f056.prepend("WHITESPACE").constData()+10), t79f056.size()-10 }; });
}

void* QRadioData_NewQRadioData(void* mediaObject, void* parent)
{
	return new MyQRadioData(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

void QRadioData_ConnectAlternativeFrequenciesEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(bool)>(&QRadioData::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(bool)>(&MyQRadioData::Signal_AlternativeFrequenciesEnabledChanged));
}

void QRadioData_DisconnectAlternativeFrequenciesEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(bool)>(&QRadioData::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(bool)>(&MyQRadioData::Signal_AlternativeFrequenciesEnabledChanged));
}

void QRadioData_AlternativeFrequenciesEnabledChanged(void* ptr, char enabled)
{
	static_cast<QRadioData*>(ptr)->alternativeFrequenciesEnabledChanged(enabled != 0);
}

long long QRadioData_Availability(void* ptr)
{
	return static_cast<QRadioData*>(ptr)->availability();
}

void QRadioData_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QRadioData::Error)>(&QRadioData::error), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QRadioData::Error)>(&MyQRadioData::Signal_Error2));
}

void QRadioData_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QRadioData::Error)>(&QRadioData::error), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QRadioData::Error)>(&MyQRadioData::Signal_Error2));
}

void QRadioData_Error2(void* ptr, long long error)
{
	static_cast<QRadioData*>(ptr)->error(static_cast<QRadioData::Error>(error));
}

long long QRadioData_Error(void* ptr)
{
	return static_cast<QRadioData*>(ptr)->error();
}

struct QtMultimedia_PackedString QRadioData_ErrorString(void* ptr)
{
	return ({ QByteArray t763fe2 = static_cast<QRadioData*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t763fe2.prepend("WHITESPACE").constData()+10), t763fe2.size()-10 }; });
}

void* QRadioData_MediaObject(void* ptr)
{
	return static_cast<QRadioData*>(ptr)->mediaObject();
}

void* QRadioData_MediaObjectDefault(void* ptr)
{
	return static_cast<QRadioData*>(ptr)->QRadioData::mediaObject();
}

void QRadioData_ConnectProgramTypeChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QRadioData::ProgramType)>(&QRadioData::programTypeChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QRadioData::ProgramType)>(&MyQRadioData::Signal_ProgramTypeChanged));
}

void QRadioData_DisconnectProgramTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QRadioData::ProgramType)>(&QRadioData::programTypeChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QRadioData::ProgramType)>(&MyQRadioData::Signal_ProgramTypeChanged));
}

void QRadioData_ProgramTypeChanged(void* ptr, long long programType)
{
	static_cast<QRadioData*>(ptr)->programTypeChanged(static_cast<QRadioData::ProgramType>(programType));
}

void QRadioData_ConnectProgramTypeNameChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::programTypeNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_ProgramTypeNameChanged));
}

void QRadioData_DisconnectProgramTypeNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::programTypeNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_ProgramTypeNameChanged));
}

void QRadioData_ProgramTypeNameChanged(void* ptr, char* programTypeName)
{
	static_cast<QRadioData*>(ptr)->programTypeNameChanged(QString(programTypeName));
}

void QRadioData_ConnectRadioTextChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::radioTextChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_RadioTextChanged));
}

void QRadioData_DisconnectRadioTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::radioTextChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_RadioTextChanged));
}

void QRadioData_RadioTextChanged(void* ptr, char* radioText)
{
	static_cast<QRadioData*>(ptr)->radioTextChanged(QString(radioText));
}

char QRadioData_SetMediaObject(void* ptr, void* mediaObject)
{
	return static_cast<QRadioData*>(ptr)->setMediaObject(static_cast<QMediaObject*>(mediaObject));
}

char QRadioData_SetMediaObjectDefault(void* ptr, void* mediaObject)
{
	return static_cast<QRadioData*>(ptr)->QRadioData::setMediaObject(static_cast<QMediaObject*>(mediaObject));
}

void QRadioData_ConnectStationIdChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationIdChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationIdChanged));
}

void QRadioData_DisconnectStationIdChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationIdChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationIdChanged));
}

void QRadioData_StationIdChanged(void* ptr, char* stationId)
{
	static_cast<QRadioData*>(ptr)->stationIdChanged(QString(stationId));
}

void QRadioData_ConnectStationNameChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationNameChanged));
}

void QRadioData_DisconnectStationNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationNameChanged));
}

void QRadioData_StationNameChanged(void* ptr, char* stationName)
{
	static_cast<QRadioData*>(ptr)->stationNameChanged(QString(stationName));
}

void QRadioData_DestroyQRadioData(void* ptr)
{
	static_cast<QRadioData*>(ptr)->~QRadioData();
}

void QRadioData_TimerEvent(void* ptr, void* event)
{
	static_cast<QRadioData*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRadioData_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRadioData*>(ptr)->QRadioData::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRadioData_ChildEvent(void* ptr, void* event)
{
	static_cast<QRadioData*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRadioData_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QRadioData*>(ptr)->QRadioData::childEvent(static_cast<QChildEvent*>(event));
}

void QRadioData_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QRadioData*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioData_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRadioData*>(ptr)->QRadioData::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioData_CustomEvent(void* ptr, void* event)
{
	static_cast<QRadioData*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRadioData_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRadioData*>(ptr)->QRadioData::customEvent(static_cast<QEvent*>(event));
}

void QRadioData_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRadioData*>(ptr), "deleteLater");
}

void QRadioData_DeleteLaterDefault(void* ptr)
{
	static_cast<QRadioData*>(ptr)->QRadioData::deleteLater();
}

void QRadioData_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QRadioData*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioData_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRadioData*>(ptr)->QRadioData::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QRadioData_Event(void* ptr, void* e)
{
	return static_cast<QRadioData*>(ptr)->event(static_cast<QEvent*>(e));
}

char QRadioData_EventDefault(void* ptr, void* e)
{
	return static_cast<QRadioData*>(ptr)->QRadioData::event(static_cast<QEvent*>(e));
}

char QRadioData_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QRadioData*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QRadioData_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRadioData*>(ptr)->QRadioData::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QRadioData_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRadioData*>(ptr)->metaObject());
}

void* QRadioData_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRadioData*>(ptr)->QRadioData::metaObject());
}

class MyQRadioDataControl: public QRadioDataControl
{
public:
	MyQRadioDataControl(QObject *parent) : QRadioDataControl(parent) {};
	void Signal_AlternativeFrequenciesEnabledChanged(bool enabled) { callbackQRadioDataControl_AlternativeFrequenciesEnabledChanged(this, enabled); };
	void Signal_Error2(QRadioData::Error error) { callbackQRadioDataControl_Error2(this, error); };
	QRadioData::Error error() const { return static_cast<QRadioData::Error>(callbackQRadioDataControl_Error(const_cast<MyQRadioDataControl*>(this))); };
	QString errorString() const { return QString(callbackQRadioDataControl_ErrorString(const_cast<MyQRadioDataControl*>(this))); };
	bool isAlternativeFrequenciesEnabled() const { return callbackQRadioDataControl_IsAlternativeFrequenciesEnabled(const_cast<MyQRadioDataControl*>(this)) != 0; };
	QRadioData::ProgramType programType() const { return static_cast<QRadioData::ProgramType>(callbackQRadioDataControl_ProgramType(const_cast<MyQRadioDataControl*>(this))); };
	void Signal_ProgramTypeChanged(QRadioData::ProgramType programType) { callbackQRadioDataControl_ProgramTypeChanged(this, programType); };
	QString programTypeName() const { return QString(callbackQRadioDataControl_ProgramTypeName(const_cast<MyQRadioDataControl*>(this))); };
	void Signal_ProgramTypeNameChanged(QString programTypeName) { QByteArray t2400f1 = programTypeName.toUtf8(); QtMultimedia_PackedString programTypeNamePacked = { const_cast<char*>(t2400f1.prepend("WHITESPACE").constData()+10), t2400f1.size()-10 };callbackQRadioDataControl_ProgramTypeNameChanged(this, programTypeNamePacked); };
	QString radioText() const { return QString(callbackQRadioDataControl_RadioText(const_cast<MyQRadioDataControl*>(this))); };
	void Signal_RadioTextChanged(QString radioText) { QByteArray t7ba492 = radioText.toUtf8(); QtMultimedia_PackedString radioTextPacked = { const_cast<char*>(t7ba492.prepend("WHITESPACE").constData()+10), t7ba492.size()-10 };callbackQRadioDataControl_RadioTextChanged(this, radioTextPacked); };
	void setAlternativeFrequenciesEnabled(bool enabled) { callbackQRadioDataControl_SetAlternativeFrequenciesEnabled(this, enabled); };
	QString stationId() const { return QString(callbackQRadioDataControl_StationId(const_cast<MyQRadioDataControl*>(this))); };
	void Signal_StationIdChanged(QString stationId) { QByteArray t6518b0 = stationId.toUtf8(); QtMultimedia_PackedString stationIdPacked = { const_cast<char*>(t6518b0.prepend("WHITESPACE").constData()+10), t6518b0.size()-10 };callbackQRadioDataControl_StationIdChanged(this, stationIdPacked); };
	QString stationName() const { return QString(callbackQRadioDataControl_StationName(const_cast<MyQRadioDataControl*>(this))); };
	void Signal_StationNameChanged(QString stationName) { QByteArray t32b816 = stationName.toUtf8(); QtMultimedia_PackedString stationNamePacked = { const_cast<char*>(t32b816.prepend("WHITESPACE").constData()+10), t32b816.size()-10 };callbackQRadioDataControl_StationNameChanged(this, stationNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRadioDataControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQRadioDataControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRadioDataControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRadioDataControl_CustomEvent(this, event); };
	void deleteLater() { callbackQRadioDataControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRadioDataControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRadioDataControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRadioDataControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRadioDataControl_MetaObject(const_cast<MyQRadioDataControl*>(this))); };
};

void* QRadioDataControl_NewQRadioDataControl(void* parent)
{
	return new MyQRadioDataControl(static_cast<QObject*>(parent));
}

void QRadioDataControl_ConnectAlternativeFrequenciesEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(bool)>(&QRadioDataControl::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(bool)>(&MyQRadioDataControl::Signal_AlternativeFrequenciesEnabledChanged));
}

void QRadioDataControl_DisconnectAlternativeFrequenciesEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(bool)>(&QRadioDataControl::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(bool)>(&MyQRadioDataControl::Signal_AlternativeFrequenciesEnabledChanged));
}

void QRadioDataControl_AlternativeFrequenciesEnabledChanged(void* ptr, char enabled)
{
	static_cast<QRadioDataControl*>(ptr)->alternativeFrequenciesEnabledChanged(enabled != 0);
}

void QRadioDataControl_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QRadioData::Error)>(&QRadioDataControl::error), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QRadioData::Error)>(&MyQRadioDataControl::Signal_Error2));
}

void QRadioDataControl_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QRadioData::Error)>(&QRadioDataControl::error), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QRadioData::Error)>(&MyQRadioDataControl::Signal_Error2));
}

void QRadioDataControl_Error2(void* ptr, long long error)
{
	static_cast<QRadioDataControl*>(ptr)->error(static_cast<QRadioData::Error>(error));
}

long long QRadioDataControl_Error(void* ptr)
{
	return static_cast<QRadioDataControl*>(ptr)->error();
}

struct QtMultimedia_PackedString QRadioDataControl_ErrorString(void* ptr)
{
	return ({ QByteArray tcfd03f = static_cast<QRadioDataControl*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tcfd03f.prepend("WHITESPACE").constData()+10), tcfd03f.size()-10 }; });
}

char QRadioDataControl_IsAlternativeFrequenciesEnabled(void* ptr)
{
	return static_cast<QRadioDataControl*>(ptr)->isAlternativeFrequenciesEnabled();
}

long long QRadioDataControl_ProgramType(void* ptr)
{
	return static_cast<QRadioDataControl*>(ptr)->programType();
}

void QRadioDataControl_ConnectProgramTypeChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QRadioData::ProgramType)>(&QRadioDataControl::programTypeChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QRadioData::ProgramType)>(&MyQRadioDataControl::Signal_ProgramTypeChanged));
}

void QRadioDataControl_DisconnectProgramTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QRadioData::ProgramType)>(&QRadioDataControl::programTypeChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QRadioData::ProgramType)>(&MyQRadioDataControl::Signal_ProgramTypeChanged));
}

void QRadioDataControl_ProgramTypeChanged(void* ptr, long long programType)
{
	static_cast<QRadioDataControl*>(ptr)->programTypeChanged(static_cast<QRadioData::ProgramType>(programType));
}

struct QtMultimedia_PackedString QRadioDataControl_ProgramTypeName(void* ptr)
{
	return ({ QByteArray tf82b2c = static_cast<QRadioDataControl*>(ptr)->programTypeName().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tf82b2c.prepend("WHITESPACE").constData()+10), tf82b2c.size()-10 }; });
}

void QRadioDataControl_ConnectProgramTypeNameChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::programTypeNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_ProgramTypeNameChanged));
}

void QRadioDataControl_DisconnectProgramTypeNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::programTypeNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_ProgramTypeNameChanged));
}

void QRadioDataControl_ProgramTypeNameChanged(void* ptr, char* programTypeName)
{
	static_cast<QRadioDataControl*>(ptr)->programTypeNameChanged(QString(programTypeName));
}

struct QtMultimedia_PackedString QRadioDataControl_RadioText(void* ptr)
{
	return ({ QByteArray t97ccf4 = static_cast<QRadioDataControl*>(ptr)->radioText().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t97ccf4.prepend("WHITESPACE").constData()+10), t97ccf4.size()-10 }; });
}

void QRadioDataControl_ConnectRadioTextChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::radioTextChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_RadioTextChanged));
}

void QRadioDataControl_DisconnectRadioTextChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::radioTextChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_RadioTextChanged));
}

void QRadioDataControl_RadioTextChanged(void* ptr, char* radioText)
{
	static_cast<QRadioDataControl*>(ptr)->radioTextChanged(QString(radioText));
}

void QRadioDataControl_SetAlternativeFrequenciesEnabled(void* ptr, char enabled)
{
	static_cast<QRadioDataControl*>(ptr)->setAlternativeFrequenciesEnabled(enabled != 0);
}

struct QtMultimedia_PackedString QRadioDataControl_StationId(void* ptr)
{
	return ({ QByteArray tbf44fc = static_cast<QRadioDataControl*>(ptr)->stationId().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tbf44fc.prepend("WHITESPACE").constData()+10), tbf44fc.size()-10 }; });
}

void QRadioDataControl_ConnectStationIdChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationIdChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationIdChanged));
}

void QRadioDataControl_DisconnectStationIdChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationIdChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationIdChanged));
}

void QRadioDataControl_StationIdChanged(void* ptr, char* stationId)
{
	static_cast<QRadioDataControl*>(ptr)->stationIdChanged(QString(stationId));
}

struct QtMultimedia_PackedString QRadioDataControl_StationName(void* ptr)
{
	return ({ QByteArray tb06572 = static_cast<QRadioDataControl*>(ptr)->stationName().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tb06572.prepend("WHITESPACE").constData()+10), tb06572.size()-10 }; });
}

void QRadioDataControl_ConnectStationNameChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationNameChanged));
}

void QRadioDataControl_DisconnectStationNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationNameChanged));
}

void QRadioDataControl_StationNameChanged(void* ptr, char* stationName)
{
	static_cast<QRadioDataControl*>(ptr)->stationNameChanged(QString(stationName));
}

void QRadioDataControl_DestroyQRadioDataControl(void* ptr)
{
	static_cast<QRadioDataControl*>(ptr)->~QRadioDataControl();
}

void QRadioDataControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QRadioDataControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRadioDataControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRadioDataControl*>(ptr)->QRadioDataControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRadioDataControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QRadioDataControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRadioDataControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QRadioDataControl*>(ptr)->QRadioDataControl::childEvent(static_cast<QChildEvent*>(event));
}

void QRadioDataControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QRadioDataControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioDataControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRadioDataControl*>(ptr)->QRadioDataControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioDataControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QRadioDataControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRadioDataControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRadioDataControl*>(ptr)->QRadioDataControl::customEvent(static_cast<QEvent*>(event));
}

void QRadioDataControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRadioDataControl*>(ptr), "deleteLater");
}

void QRadioDataControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QRadioDataControl*>(ptr)->QRadioDataControl::deleteLater();
}

void QRadioDataControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QRadioDataControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioDataControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRadioDataControl*>(ptr)->QRadioDataControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QRadioDataControl_Event(void* ptr, void* e)
{
	return static_cast<QRadioDataControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QRadioDataControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QRadioDataControl*>(ptr)->QRadioDataControl::event(static_cast<QEvent*>(e));
}

char QRadioDataControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QRadioDataControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QRadioDataControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRadioDataControl*>(ptr)->QRadioDataControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QRadioDataControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRadioDataControl*>(ptr)->metaObject());
}

void* QRadioDataControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRadioDataControl*>(ptr)->QRadioDataControl::metaObject());
}

class MyQRadioTuner: public QRadioTuner
{
public:
	MyQRadioTuner(QObject *parent) : QRadioTuner(parent) {};
	void setMuted(bool muted) { callbackQRadioTuner_SetMuted(this, muted); };
	void setVolume(int volume) { callbackQRadioTuner_SetVolume(this, volume); };
	void Signal_AntennaConnectedChanged(bool connectionStatus) { callbackQRadioTuner_AntennaConnectedChanged(this, connectionStatus); };
	QMultimedia::AvailabilityStatus availability() const { return static_cast<QMultimedia::AvailabilityStatus>(callbackQRadioTuner_Availability(const_cast<MyQRadioTuner*>(this))); };
	void Signal_BandChanged(QRadioTuner::Band band) { callbackQRadioTuner_BandChanged(this, band); };
	void cancelSearch() { callbackQRadioTuner_CancelSearch(this); };
	void Signal_Error2(QRadioTuner::Error error) { callbackQRadioTuner_Error2(this, error); };
	void Signal_FrequencyChanged(int frequency) { callbackQRadioTuner_FrequencyChanged(this, frequency); };
	void Signal_MutedChanged(bool muted) { callbackQRadioTuner_MutedChanged(this, muted); };
	void searchAllStations(QRadioTuner::SearchMode searchMode) { callbackQRadioTuner_SearchAllStations(this, searchMode); };
	void searchBackward() { callbackQRadioTuner_SearchBackward(this); };
	void searchForward() { callbackQRadioTuner_SearchForward(this); };
	void Signal_SearchingChanged(bool searching) { callbackQRadioTuner_SearchingChanged(this, searching); };
	void setBand(QRadioTuner::Band band) { callbackQRadioTuner_SetBand(this, band); };
	void setFrequency(int frequency) { callbackQRadioTuner_SetFrequency(this, frequency); };
	void Signal_SignalStrengthChanged(int strength) { callbackQRadioTuner_SignalStrengthChanged(this, strength); };
	void start() { callbackQRadioTuner_Start(this); };
	void Signal_StateChanged(QRadioTuner::State state) { callbackQRadioTuner_StateChanged(this, state); };
	void Signal_StationFound(int frequency, QString stationId) { QByteArray t6518b0 = stationId.toUtf8(); QtMultimedia_PackedString stationIdPacked = { const_cast<char*>(t6518b0.prepend("WHITESPACE").constData()+10), t6518b0.size()-10 };callbackQRadioTuner_StationFound(this, frequency, stationIdPacked); };
	void Signal_StereoStatusChanged(bool stereo) { callbackQRadioTuner_StereoStatusChanged(this, stereo); };
	void stop() { callbackQRadioTuner_Stop(this); };
	void Signal_VolumeChanged(int volume) { callbackQRadioTuner_VolumeChanged(this, volume); };
	bool bind(QObject * object) { return callbackQRadioTuner_Bind(this, object) != 0; };
	bool isAvailable() const { return callbackQRadioTuner_IsAvailable(const_cast<MyQRadioTuner*>(this)) != 0; };
	QMediaService * service() const { return static_cast<QMediaService*>(callbackQRadioTuner_Service(const_cast<MyQRadioTuner*>(this))); };
	void unbind(QObject * object) { callbackQRadioTuner_Unbind(this, object); };
	void timerEvent(QTimerEvent * event) { callbackQRadioTuner_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQRadioTuner_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRadioTuner_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRadioTuner_CustomEvent(this, event); };
	void deleteLater() { callbackQRadioTuner_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRadioTuner_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRadioTuner_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRadioTuner_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRadioTuner_MetaObject(const_cast<MyQRadioTuner*>(this))); };
};

long long QRadioTuner_Band(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->band();
}

int QRadioTuner_Frequency(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->frequency();
}

char QRadioTuner_IsAntennaConnected(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->isAntennaConnected();
}

char QRadioTuner_IsMuted(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->isMuted();
}

char QRadioTuner_IsSearching(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->isSearching();
}

char QRadioTuner_IsStereo(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->isStereo();
}

void* QRadioTuner_RadioData(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->radioData();
}

void QRadioTuner_SetMuted(void* ptr, char muted)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

void QRadioTuner_SetStereoMode(void* ptr, long long mode)
{
	static_cast<QRadioTuner*>(ptr)->setStereoMode(static_cast<QRadioTuner::StereoMode>(mode));
}

void QRadioTuner_SetVolume(void* ptr, int volume)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setVolume", Q_ARG(int, volume));
}

int QRadioTuner_SignalStrength(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->signalStrength();
}

long long QRadioTuner_State(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->state();
}

long long QRadioTuner_StereoMode(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->stereoMode();
}

int QRadioTuner_Volume(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->volume();
}

void* QRadioTuner_NewQRadioTuner(void* parent)
{
	return new MyQRadioTuner(static_cast<QObject*>(parent));
}

void QRadioTuner_ConnectAntennaConnectedChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::antennaConnectedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_AntennaConnectedChanged));
}

void QRadioTuner_DisconnectAntennaConnectedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::antennaConnectedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_AntennaConnectedChanged));
}

void QRadioTuner_AntennaConnectedChanged(void* ptr, char connectionStatus)
{
	static_cast<QRadioTuner*>(ptr)->antennaConnectedChanged(connectionStatus != 0);
}

long long QRadioTuner_Availability(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->availability();
}

long long QRadioTuner_AvailabilityDefault(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->QRadioTuner::availability();
}

void QRadioTuner_ConnectBandChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Band)>(&QRadioTuner::bandChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Band)>(&MyQRadioTuner::Signal_BandChanged));
}

void QRadioTuner_DisconnectBandChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Band)>(&QRadioTuner::bandChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Band)>(&MyQRadioTuner::Signal_BandChanged));
}

void QRadioTuner_BandChanged(void* ptr, long long band)
{
	static_cast<QRadioTuner*>(ptr)->bandChanged(static_cast<QRadioTuner::Band>(band));
}

void QRadioTuner_CancelSearch(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "cancelSearch");
}

void QRadioTuner_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Error)>(&QRadioTuner::error), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Error)>(&MyQRadioTuner::Signal_Error2));
}

void QRadioTuner_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Error)>(&QRadioTuner::error), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Error)>(&MyQRadioTuner::Signal_Error2));
}

void QRadioTuner_Error2(void* ptr, long long error)
{
	static_cast<QRadioTuner*>(ptr)->error(static_cast<QRadioTuner::Error>(error));
}

long long QRadioTuner_Error(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->error();
}

struct QtMultimedia_PackedString QRadioTuner_ErrorString(void* ptr)
{
	return ({ QByteArray t26ab64 = static_cast<QRadioTuner*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t26ab64.prepend("WHITESPACE").constData()+10), t26ab64.size()-10 }; });
}

void QRadioTuner_ConnectFrequencyChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::frequencyChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_FrequencyChanged));
}

void QRadioTuner_DisconnectFrequencyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::frequencyChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_FrequencyChanged));
}

void QRadioTuner_FrequencyChanged(void* ptr, int frequency)
{
	static_cast<QRadioTuner*>(ptr)->frequencyChanged(frequency);
}

int QRadioTuner_FrequencyStep(void* ptr, long long band)
{
	return static_cast<QRadioTuner*>(ptr)->frequencyStep(static_cast<QRadioTuner::Band>(band));
}

char QRadioTuner_IsBandSupported(void* ptr, long long band)
{
	return static_cast<QRadioTuner*>(ptr)->isBandSupported(static_cast<QRadioTuner::Band>(band));
}

void QRadioTuner_ConnectMutedChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::mutedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_MutedChanged));
}

void QRadioTuner_DisconnectMutedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::mutedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_MutedChanged));
}

void QRadioTuner_MutedChanged(void* ptr, char muted)
{
	static_cast<QRadioTuner*>(ptr)->mutedChanged(muted != 0);
}

void QRadioTuner_SearchAllStations(void* ptr, long long searchMode)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchAllStations", Q_ARG(QRadioTuner::SearchMode, static_cast<QRadioTuner::SearchMode>(searchMode)));
}

void QRadioTuner_SearchBackward(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchBackward");
}

void QRadioTuner_SearchForward(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchForward");
}

void QRadioTuner_ConnectSearchingChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::searchingChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_SearchingChanged));
}

void QRadioTuner_DisconnectSearchingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::searchingChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_SearchingChanged));
}

void QRadioTuner_SearchingChanged(void* ptr, char searching)
{
	static_cast<QRadioTuner*>(ptr)->searchingChanged(searching != 0);
}

void QRadioTuner_SetBand(void* ptr, long long band)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setBand", Q_ARG(QRadioTuner::Band, static_cast<QRadioTuner::Band>(band)));
}

void QRadioTuner_SetFrequency(void* ptr, int frequency)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setFrequency", Q_ARG(int, frequency));
}

void QRadioTuner_ConnectSignalStrengthChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::signalStrengthChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_SignalStrengthChanged));
}

void QRadioTuner_DisconnectSignalStrengthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::signalStrengthChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_SignalStrengthChanged));
}

void QRadioTuner_SignalStrengthChanged(void* ptr, int strength)
{
	static_cast<QRadioTuner*>(ptr)->signalStrengthChanged(strength);
}

void QRadioTuner_Start(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "start");
}

void QRadioTuner_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::State)>(&QRadioTuner::stateChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::State)>(&MyQRadioTuner::Signal_StateChanged));
}

void QRadioTuner_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::State)>(&QRadioTuner::stateChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::State)>(&MyQRadioTuner::Signal_StateChanged));
}

void QRadioTuner_StateChanged(void* ptr, long long state)
{
	static_cast<QRadioTuner*>(ptr)->stateChanged(static_cast<QRadioTuner::State>(state));
}

void QRadioTuner_ConnectStationFound(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int, QString)>(&QRadioTuner::stationFound), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int, QString)>(&MyQRadioTuner::Signal_StationFound));
}

void QRadioTuner_DisconnectStationFound(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int, QString)>(&QRadioTuner::stationFound), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int, QString)>(&MyQRadioTuner::Signal_StationFound));
}

void QRadioTuner_StationFound(void* ptr, int frequency, char* stationId)
{
	static_cast<QRadioTuner*>(ptr)->stationFound(frequency, QString(stationId));
}

void QRadioTuner_ConnectStereoStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::stereoStatusChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_StereoStatusChanged));
}

void QRadioTuner_DisconnectStereoStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::stereoStatusChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_StereoStatusChanged));
}

void QRadioTuner_StereoStatusChanged(void* ptr, char stereo)
{
	static_cast<QRadioTuner*>(ptr)->stereoStatusChanged(stereo != 0);
}

void QRadioTuner_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "stop");
}

void QRadioTuner_ConnectVolumeChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::volumeChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_VolumeChanged));
}

void QRadioTuner_DisconnectVolumeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::volumeChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_VolumeChanged));
}

void QRadioTuner_VolumeChanged(void* ptr, int volume)
{
	static_cast<QRadioTuner*>(ptr)->volumeChanged(volume);
}

void QRadioTuner_DestroyQRadioTuner(void* ptr)
{
	static_cast<QRadioTuner*>(ptr)->~QRadioTuner();
}

char QRadioTuner_Bind(void* ptr, void* object)
{
	return static_cast<QRadioTuner*>(ptr)->bind(static_cast<QObject*>(object));
}

char QRadioTuner_BindDefault(void* ptr, void* object)
{
	return static_cast<QRadioTuner*>(ptr)->QRadioTuner::bind(static_cast<QObject*>(object));
}

char QRadioTuner_IsAvailable(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->isAvailable();
}

char QRadioTuner_IsAvailableDefault(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->QRadioTuner::isAvailable();
}

void* QRadioTuner_Service(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->service();
}

void* QRadioTuner_ServiceDefault(void* ptr)
{
	return static_cast<QRadioTuner*>(ptr)->QRadioTuner::service();
}

void QRadioTuner_Unbind(void* ptr, void* object)
{
	static_cast<QRadioTuner*>(ptr)->unbind(static_cast<QObject*>(object));
}

void QRadioTuner_UnbindDefault(void* ptr, void* object)
{
	static_cast<QRadioTuner*>(ptr)->QRadioTuner::unbind(static_cast<QObject*>(object));
}

void QRadioTuner_TimerEvent(void* ptr, void* event)
{
	static_cast<QRadioTuner*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRadioTuner_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRadioTuner*>(ptr)->QRadioTuner::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRadioTuner_ChildEvent(void* ptr, void* event)
{
	static_cast<QRadioTuner*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRadioTuner_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QRadioTuner*>(ptr)->QRadioTuner::childEvent(static_cast<QChildEvent*>(event));
}

void QRadioTuner_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QRadioTuner*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioTuner_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRadioTuner*>(ptr)->QRadioTuner::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioTuner_CustomEvent(void* ptr, void* event)
{
	static_cast<QRadioTuner*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRadioTuner_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRadioTuner*>(ptr)->QRadioTuner::customEvent(static_cast<QEvent*>(event));
}

void QRadioTuner_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "deleteLater");
}

void QRadioTuner_DeleteLaterDefault(void* ptr)
{
	static_cast<QRadioTuner*>(ptr)->QRadioTuner::deleteLater();
}

void QRadioTuner_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QRadioTuner*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioTuner_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRadioTuner*>(ptr)->QRadioTuner::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QRadioTuner_Event(void* ptr, void* e)
{
	return static_cast<QRadioTuner*>(ptr)->event(static_cast<QEvent*>(e));
}

char QRadioTuner_EventDefault(void* ptr, void* e)
{
	return static_cast<QRadioTuner*>(ptr)->QRadioTuner::event(static_cast<QEvent*>(e));
}

char QRadioTuner_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QRadioTuner*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QRadioTuner_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRadioTuner*>(ptr)->QRadioTuner::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QRadioTuner_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRadioTuner*>(ptr)->metaObject());
}

void* QRadioTuner_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRadioTuner*>(ptr)->QRadioTuner::metaObject());
}

class MyQRadioTunerControl: public QRadioTunerControl
{
public:
	void Signal_AntennaConnectedChanged(bool connectionStatus) { callbackQRadioTunerControl_AntennaConnectedChanged(this, connectionStatus); };
	QRadioTuner::Band band() const { return static_cast<QRadioTuner::Band>(callbackQRadioTunerControl_Band(const_cast<MyQRadioTunerControl*>(this))); };
	void Signal_BandChanged(QRadioTuner::Band band) { callbackQRadioTunerControl_BandChanged(this, band); };
	void cancelSearch() { callbackQRadioTunerControl_CancelSearch(this); };
	void Signal_Error2(QRadioTuner::Error error) { callbackQRadioTunerControl_Error2(this, error); };
	QRadioTuner::Error error() const { return static_cast<QRadioTuner::Error>(callbackQRadioTunerControl_Error(const_cast<MyQRadioTunerControl*>(this))); };
	QString errorString() const { return QString(callbackQRadioTunerControl_ErrorString(const_cast<MyQRadioTunerControl*>(this))); };
	int frequency() const { return callbackQRadioTunerControl_Frequency(const_cast<MyQRadioTunerControl*>(this)); };
	void Signal_FrequencyChanged(int frequency) { callbackQRadioTunerControl_FrequencyChanged(this, frequency); };
	int frequencyStep(QRadioTuner::Band band) const { return callbackQRadioTunerControl_FrequencyStep(const_cast<MyQRadioTunerControl*>(this), band); };
	bool isAntennaConnected() const { return callbackQRadioTunerControl_IsAntennaConnected(const_cast<MyQRadioTunerControl*>(this)) != 0; };
	bool isBandSupported(QRadioTuner::Band band) const { return callbackQRadioTunerControl_IsBandSupported(const_cast<MyQRadioTunerControl*>(this), band) != 0; };
	bool isMuted() const { return callbackQRadioTunerControl_IsMuted(const_cast<MyQRadioTunerControl*>(this)) != 0; };
	bool isSearching() const { return callbackQRadioTunerControl_IsSearching(const_cast<MyQRadioTunerControl*>(this)) != 0; };
	bool isStereo() const { return callbackQRadioTunerControl_IsStereo(const_cast<MyQRadioTunerControl*>(this)) != 0; };
	void Signal_MutedChanged(bool muted) { callbackQRadioTunerControl_MutedChanged(this, muted); };
	void searchAllStations(QRadioTuner::SearchMode searchMode) { callbackQRadioTunerControl_SearchAllStations(this, searchMode); };
	void searchBackward() { callbackQRadioTunerControl_SearchBackward(this); };
	void searchForward() { callbackQRadioTunerControl_SearchForward(this); };
	void Signal_SearchingChanged(bool searching) { callbackQRadioTunerControl_SearchingChanged(this, searching); };
	void setBand(QRadioTuner::Band band) { callbackQRadioTunerControl_SetBand(this, band); };
	void setFrequency(int frequency) { callbackQRadioTunerControl_SetFrequency(this, frequency); };
	void setMuted(bool muted) { callbackQRadioTunerControl_SetMuted(this, muted); };
	void setStereoMode(QRadioTuner::StereoMode mode) { callbackQRadioTunerControl_SetStereoMode(this, mode); };
	void setVolume(int volume) { callbackQRadioTunerControl_SetVolume(this, volume); };
	int signalStrength() const { return callbackQRadioTunerControl_SignalStrength(const_cast<MyQRadioTunerControl*>(this)); };
	void Signal_SignalStrengthChanged(int strength) { callbackQRadioTunerControl_SignalStrengthChanged(this, strength); };
	void start() { callbackQRadioTunerControl_Start(this); };
	QRadioTuner::State state() const { return static_cast<QRadioTuner::State>(callbackQRadioTunerControl_State(const_cast<MyQRadioTunerControl*>(this))); };
	void Signal_StateChanged(QRadioTuner::State state) { callbackQRadioTunerControl_StateChanged(this, state); };
	void Signal_StationFound(int frequency, QString stationId) { QByteArray t6518b0 = stationId.toUtf8(); QtMultimedia_PackedString stationIdPacked = { const_cast<char*>(t6518b0.prepend("WHITESPACE").constData()+10), t6518b0.size()-10 };callbackQRadioTunerControl_StationFound(this, frequency, stationIdPacked); };
	QRadioTuner::StereoMode stereoMode() const { return static_cast<QRadioTuner::StereoMode>(callbackQRadioTunerControl_StereoMode(const_cast<MyQRadioTunerControl*>(this))); };
	void Signal_StereoStatusChanged(bool stereo) { callbackQRadioTunerControl_StereoStatusChanged(this, stereo); };
	void stop() { callbackQRadioTunerControl_Stop(this); };
	int volume() const { return callbackQRadioTunerControl_Volume(const_cast<MyQRadioTunerControl*>(this)); };
	void Signal_VolumeChanged(int volume) { callbackQRadioTunerControl_VolumeChanged(this, volume); };
	void timerEvent(QTimerEvent * event) { callbackQRadioTunerControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQRadioTunerControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRadioTunerControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRadioTunerControl_CustomEvent(this, event); };
	void deleteLater() { callbackQRadioTunerControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRadioTunerControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRadioTunerControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRadioTunerControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRadioTunerControl_MetaObject(const_cast<MyQRadioTunerControl*>(this))); };
};

void QRadioTunerControl_ConnectAntennaConnectedChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::antennaConnectedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_AntennaConnectedChanged));
}

void QRadioTunerControl_DisconnectAntennaConnectedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::antennaConnectedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_AntennaConnectedChanged));
}

void QRadioTunerControl_AntennaConnectedChanged(void* ptr, char connectionStatus)
{
	static_cast<QRadioTunerControl*>(ptr)->antennaConnectedChanged(connectionStatus != 0);
}

long long QRadioTunerControl_Band(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->band();
}

void QRadioTunerControl_ConnectBandChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::Band)>(&QRadioTunerControl::bandChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::Band)>(&MyQRadioTunerControl::Signal_BandChanged));
}

void QRadioTunerControl_DisconnectBandChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::Band)>(&QRadioTunerControl::bandChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::Band)>(&MyQRadioTunerControl::Signal_BandChanged));
}

void QRadioTunerControl_BandChanged(void* ptr, long long band)
{
	static_cast<QRadioTunerControl*>(ptr)->bandChanged(static_cast<QRadioTuner::Band>(band));
}

void QRadioTunerControl_CancelSearch(void* ptr)
{
	static_cast<QRadioTunerControl*>(ptr)->cancelSearch();
}

void QRadioTunerControl_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::Error)>(&QRadioTunerControl::error), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::Error)>(&MyQRadioTunerControl::Signal_Error2));
}

void QRadioTunerControl_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::Error)>(&QRadioTunerControl::error), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::Error)>(&MyQRadioTunerControl::Signal_Error2));
}

void QRadioTunerControl_Error2(void* ptr, long long error)
{
	static_cast<QRadioTunerControl*>(ptr)->error(static_cast<QRadioTuner::Error>(error));
}

long long QRadioTunerControl_Error(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->error();
}

struct QtMultimedia_PackedString QRadioTunerControl_ErrorString(void* ptr)
{
	return ({ QByteArray t29e878 = static_cast<QRadioTunerControl*>(ptr)->errorString().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t29e878.prepend("WHITESPACE").constData()+10), t29e878.size()-10 }; });
}

int QRadioTunerControl_Frequency(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->frequency();
}

void QRadioTunerControl_ConnectFrequencyChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::frequencyChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_FrequencyChanged));
}

void QRadioTunerControl_DisconnectFrequencyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::frequencyChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_FrequencyChanged));
}

void QRadioTunerControl_FrequencyChanged(void* ptr, int frequency)
{
	static_cast<QRadioTunerControl*>(ptr)->frequencyChanged(frequency);
}

int QRadioTunerControl_FrequencyStep(void* ptr, long long band)
{
	return static_cast<QRadioTunerControl*>(ptr)->frequencyStep(static_cast<QRadioTuner::Band>(band));
}

char QRadioTunerControl_IsAntennaConnected(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->isAntennaConnected();
}

char QRadioTunerControl_IsAntennaConnectedDefault(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::isAntennaConnected();
}

char QRadioTunerControl_IsBandSupported(void* ptr, long long band)
{
	return static_cast<QRadioTunerControl*>(ptr)->isBandSupported(static_cast<QRadioTuner::Band>(band));
}

char QRadioTunerControl_IsMuted(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->isMuted();
}

char QRadioTunerControl_IsSearching(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->isSearching();
}

char QRadioTunerControl_IsStereo(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->isStereo();
}

void QRadioTunerControl_ConnectMutedChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::mutedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_MutedChanged));
}

void QRadioTunerControl_DisconnectMutedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::mutedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_MutedChanged));
}

void QRadioTunerControl_MutedChanged(void* ptr, char muted)
{
	static_cast<QRadioTunerControl*>(ptr)->mutedChanged(muted != 0);
}

void QRadioTunerControl_SearchAllStations(void* ptr, long long searchMode)
{
	static_cast<QRadioTunerControl*>(ptr)->searchAllStations(static_cast<QRadioTuner::SearchMode>(searchMode));
}

void QRadioTunerControl_SearchBackward(void* ptr)
{
	static_cast<QRadioTunerControl*>(ptr)->searchBackward();
}

void QRadioTunerControl_SearchForward(void* ptr)
{
	static_cast<QRadioTunerControl*>(ptr)->searchForward();
}

void QRadioTunerControl_ConnectSearchingChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::searchingChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_SearchingChanged));
}

void QRadioTunerControl_DisconnectSearchingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::searchingChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_SearchingChanged));
}

void QRadioTunerControl_SearchingChanged(void* ptr, char searching)
{
	static_cast<QRadioTunerControl*>(ptr)->searchingChanged(searching != 0);
}

void QRadioTunerControl_SetBand(void* ptr, long long band)
{
	static_cast<QRadioTunerControl*>(ptr)->setBand(static_cast<QRadioTuner::Band>(band));
}

void QRadioTunerControl_SetFrequency(void* ptr, int frequency)
{
	static_cast<QRadioTunerControl*>(ptr)->setFrequency(frequency);
}

void QRadioTunerControl_SetMuted(void* ptr, char muted)
{
	static_cast<QRadioTunerControl*>(ptr)->setMuted(muted != 0);
}

void QRadioTunerControl_SetStereoMode(void* ptr, long long mode)
{
	static_cast<QRadioTunerControl*>(ptr)->setStereoMode(static_cast<QRadioTuner::StereoMode>(mode));
}

void QRadioTunerControl_SetVolume(void* ptr, int volume)
{
	static_cast<QRadioTunerControl*>(ptr)->setVolume(volume);
}

int QRadioTunerControl_SignalStrength(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->signalStrength();
}

void QRadioTunerControl_ConnectSignalStrengthChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::signalStrengthChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_SignalStrengthChanged));
}

void QRadioTunerControl_DisconnectSignalStrengthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::signalStrengthChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_SignalStrengthChanged));
}

void QRadioTunerControl_SignalStrengthChanged(void* ptr, int strength)
{
	static_cast<QRadioTunerControl*>(ptr)->signalStrengthChanged(strength);
}

void QRadioTunerControl_Start(void* ptr)
{
	static_cast<QRadioTunerControl*>(ptr)->start();
}

long long QRadioTunerControl_State(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->state();
}

void QRadioTunerControl_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::State)>(&QRadioTunerControl::stateChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::State)>(&MyQRadioTunerControl::Signal_StateChanged));
}

void QRadioTunerControl_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::State)>(&QRadioTunerControl::stateChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::State)>(&MyQRadioTunerControl::Signal_StateChanged));
}

void QRadioTunerControl_StateChanged(void* ptr, long long state)
{
	static_cast<QRadioTunerControl*>(ptr)->stateChanged(static_cast<QRadioTuner::State>(state));
}

void QRadioTunerControl_ConnectStationFound(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int, QString)>(&QRadioTunerControl::stationFound), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int, QString)>(&MyQRadioTunerControl::Signal_StationFound));
}

void QRadioTunerControl_DisconnectStationFound(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int, QString)>(&QRadioTunerControl::stationFound), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int, QString)>(&MyQRadioTunerControl::Signal_StationFound));
}

void QRadioTunerControl_StationFound(void* ptr, int frequency, char* stationId)
{
	static_cast<QRadioTunerControl*>(ptr)->stationFound(frequency, QString(stationId));
}

long long QRadioTunerControl_StereoMode(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->stereoMode();
}

void QRadioTunerControl_ConnectStereoStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::stereoStatusChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_StereoStatusChanged));
}

void QRadioTunerControl_DisconnectStereoStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::stereoStatusChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_StereoStatusChanged));
}

void QRadioTunerControl_StereoStatusChanged(void* ptr, char stereo)
{
	static_cast<QRadioTunerControl*>(ptr)->stereoStatusChanged(stereo != 0);
}

void QRadioTunerControl_Stop(void* ptr)
{
	static_cast<QRadioTunerControl*>(ptr)->stop();
}

int QRadioTunerControl_Volume(void* ptr)
{
	return static_cast<QRadioTunerControl*>(ptr)->volume();
}

void QRadioTunerControl_ConnectVolumeChanged(void* ptr)
{
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::volumeChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_VolumeChanged));
}

void QRadioTunerControl_DisconnectVolumeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::volumeChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_VolumeChanged));
}

void QRadioTunerControl_VolumeChanged(void* ptr, int volume)
{
	static_cast<QRadioTunerControl*>(ptr)->volumeChanged(volume);
}

void QRadioTunerControl_DestroyQRadioTunerControl(void* ptr)
{
	static_cast<QRadioTunerControl*>(ptr)->~QRadioTunerControl();
}

void QRadioTunerControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QRadioTunerControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRadioTunerControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRadioTunerControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QRadioTunerControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRadioTunerControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::childEvent(static_cast<QChildEvent*>(event));
}

void QRadioTunerControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QRadioTunerControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioTunerControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioTunerControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QRadioTunerControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRadioTunerControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::customEvent(static_cast<QEvent*>(event));
}

void QRadioTunerControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRadioTunerControl*>(ptr), "deleteLater");
}

void QRadioTunerControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::deleteLater();
}

void QRadioTunerControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QRadioTunerControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRadioTunerControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QRadioTunerControl_Event(void* ptr, void* e)
{
	return static_cast<QRadioTunerControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QRadioTunerControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::event(static_cast<QEvent*>(e));
}

char QRadioTunerControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QRadioTunerControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QRadioTunerControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QRadioTunerControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRadioTunerControl*>(ptr)->metaObject());
}

void* QRadioTunerControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRadioTunerControl*>(ptr)->QRadioTunerControl::metaObject());
}

class MyQSound: public QSound
{
public:
	MyQSound(const QString &filename, QObject *parent) : QSound(filename, parent) {};
	void play() { callbackQSound_Play2(this); };
	void stop() { callbackQSound_Stop(this); };
	void timerEvent(QTimerEvent * event) { callbackQSound_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSound_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSound_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSound_CustomEvent(this, event); };
	void deleteLater() { callbackQSound_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSound_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSound_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSound_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSound_MetaObject(const_cast<MyQSound*>(this))); };
};

void QSound_SetLoops(void* ptr, int number)
{
	static_cast<QSound*>(ptr)->setLoops(number);
}

void* QSound_NewQSound(char* filename, void* parent)
{
	return new MyQSound(QString(filename), static_cast<QObject*>(parent));
}

struct QtMultimedia_PackedString QSound_FileName(void* ptr)
{
	return ({ QByteArray t54b0ee = static_cast<QSound*>(ptr)->fileName().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t54b0ee.prepend("WHITESPACE").constData()+10), t54b0ee.size()-10 }; });
}

char QSound_IsFinished(void* ptr)
{
	return static_cast<QSound*>(ptr)->isFinished();
}

int QSound_Loops(void* ptr)
{
	return static_cast<QSound*>(ptr)->loops();
}

int QSound_LoopsRemaining(void* ptr)
{
	return static_cast<QSound*>(ptr)->loopsRemaining();
}

void QSound_Play2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSound*>(ptr), "play");
}

void QSound_QSound_Play(char* filename)
{
	QSound::play(QString(filename));
}

void QSound_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSound*>(ptr), "stop");
}

void QSound_DestroyQSound(void* ptr)
{
	static_cast<QSound*>(ptr)->~QSound();
}

void QSound_TimerEvent(void* ptr, void* event)
{
	static_cast<QSound*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSound_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSound*>(ptr)->QSound::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSound_ChildEvent(void* ptr, void* event)
{
	static_cast<QSound*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSound_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSound*>(ptr)->QSound::childEvent(static_cast<QChildEvent*>(event));
}

void QSound_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSound*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSound_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSound*>(ptr)->QSound::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSound_CustomEvent(void* ptr, void* event)
{
	static_cast<QSound*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSound_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSound*>(ptr)->QSound::customEvent(static_cast<QEvent*>(event));
}

void QSound_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSound*>(ptr), "deleteLater");
}

void QSound_DeleteLaterDefault(void* ptr)
{
	static_cast<QSound*>(ptr)->QSound::deleteLater();
}

void QSound_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSound*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSound_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSound*>(ptr)->QSound::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSound_Event(void* ptr, void* e)
{
	return static_cast<QSound*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSound_EventDefault(void* ptr, void* e)
{
	return static_cast<QSound*>(ptr)->QSound::event(static_cast<QEvent*>(e));
}

char QSound_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSound*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSound_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSound*>(ptr)->QSound::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSound_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSound*>(ptr)->metaObject());
}

void* QSound_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSound*>(ptr)->QSound::metaObject());
}

class MyQSoundEffect: public QSoundEffect
{
public:
	MyQSoundEffect(QObject *parent) : QSoundEffect(parent) {};
	void play() { callbackQSoundEffect_Play(this); };
	void stop() { callbackQSoundEffect_Stop(this); };
	void Signal_CategoryChanged() { callbackQSoundEffect_CategoryChanged(this); };
	void Signal_LoadedChanged() { callbackQSoundEffect_LoadedChanged(this); };
	void Signal_LoopCountChanged() { callbackQSoundEffect_LoopCountChanged(this); };
	void Signal_LoopsRemainingChanged() { callbackQSoundEffect_LoopsRemainingChanged(this); };
	void Signal_MutedChanged() { callbackQSoundEffect_MutedChanged(this); };
	void Signal_PlayingChanged() { callbackQSoundEffect_PlayingChanged(this); };
	void Signal_SourceChanged() { callbackQSoundEffect_SourceChanged(this); };
	void Signal_StatusChanged() { callbackQSoundEffect_StatusChanged(this); };
	void Signal_VolumeChanged() { callbackQSoundEffect_VolumeChanged(this); };
	void timerEvent(QTimerEvent * event) { callbackQSoundEffect_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQSoundEffect_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSoundEffect_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSoundEffect_CustomEvent(this, event); };
	void deleteLater() { callbackQSoundEffect_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSoundEffect_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQSoundEffect_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSoundEffect_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSoundEffect_MetaObject(const_cast<MyQSoundEffect*>(this))); };
};

char QSoundEffect_IsLoaded(void* ptr)
{
	return static_cast<QSoundEffect*>(ptr)->isLoaded();
}

int QSoundEffect_LoopsRemaining(void* ptr)
{
	return static_cast<QSoundEffect*>(ptr)->loopsRemaining();
}

void QSoundEffect_Play(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSoundEffect*>(ptr), "play");
}

void QSoundEffect_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSoundEffect*>(ptr), "stop");
}

struct QtMultimedia_PackedString QSoundEffect_QSoundEffect_SupportedMimeTypes()
{
	return ({ QByteArray t9ca848 = QSoundEffect::supportedMimeTypes().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t9ca848.prepend("WHITESPACE").constData()+10), t9ca848.size()-10 }; });
}

void* QSoundEffect_NewQSoundEffect(void* parent)
{
	return new MyQSoundEffect(static_cast<QObject*>(parent));
}

struct QtMultimedia_PackedString QSoundEffect_Category(void* ptr)
{
	return ({ QByteArray t036c43 = static_cast<QSoundEffect*>(ptr)->category().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t036c43.prepend("WHITESPACE").constData()+10), t036c43.size()-10 }; });
}

void QSoundEffect_ConnectCategoryChanged(void* ptr)
{
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::categoryChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_CategoryChanged));
}

void QSoundEffect_DisconnectCategoryChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::categoryChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_CategoryChanged));
}

void QSoundEffect_CategoryChanged(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->categoryChanged();
}

char QSoundEffect_IsMuted(void* ptr)
{
	return static_cast<QSoundEffect*>(ptr)->isMuted();
}

char QSoundEffect_IsPlaying(void* ptr)
{
	return static_cast<QSoundEffect*>(ptr)->isPlaying();
}

void QSoundEffect_ConnectLoadedChanged(void* ptr)
{
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loadedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoadedChanged));
}

void QSoundEffect_DisconnectLoadedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loadedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoadedChanged));
}

void QSoundEffect_LoadedChanged(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->loadedChanged();
}

int QSoundEffect_LoopCount(void* ptr)
{
	return static_cast<QSoundEffect*>(ptr)->loopCount();
}

void QSoundEffect_ConnectLoopCountChanged(void* ptr)
{
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopCountChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopCountChanged));
}

void QSoundEffect_DisconnectLoopCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopCountChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopCountChanged));
}

void QSoundEffect_LoopCountChanged(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->loopCountChanged();
}

void QSoundEffect_ConnectLoopsRemainingChanged(void* ptr)
{
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopsRemainingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopsRemainingChanged));
}

void QSoundEffect_DisconnectLoopsRemainingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopsRemainingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopsRemainingChanged));
}

void QSoundEffect_LoopsRemainingChanged(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->loopsRemainingChanged();
}

void QSoundEffect_ConnectMutedChanged(void* ptr)
{
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::mutedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_MutedChanged));
}

void QSoundEffect_DisconnectMutedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::mutedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_MutedChanged));
}

void QSoundEffect_MutedChanged(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->mutedChanged();
}

void QSoundEffect_ConnectPlayingChanged(void* ptr)
{
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::playingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_PlayingChanged));
}

void QSoundEffect_DisconnectPlayingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::playingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_PlayingChanged));
}

void QSoundEffect_PlayingChanged(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->playingChanged();
}

void QSoundEffect_SetCategory(void* ptr, char* category)
{
	static_cast<QSoundEffect*>(ptr)->setCategory(QString(category));
}

void QSoundEffect_SetLoopCount(void* ptr, int loopCount)
{
	static_cast<QSoundEffect*>(ptr)->setLoopCount(loopCount);
}

void QSoundEffect_SetMuted(void* ptr, char muted)
{
	static_cast<QSoundEffect*>(ptr)->setMuted(muted != 0);
}

void QSoundEffect_SetSource(void* ptr, void* url)
{
	static_cast<QSoundEffect*>(ptr)->setSource(*static_cast<QUrl*>(url));
}

void QSoundEffect_SetVolume(void* ptr, double volume)
{
	static_cast<QSoundEffect*>(ptr)->setVolume(volume);
}

void* QSoundEffect_Source(void* ptr)
{
	return new QUrl(static_cast<QSoundEffect*>(ptr)->source());
}

void QSoundEffect_ConnectSourceChanged(void* ptr)
{
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::sourceChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_SourceChanged));
}

void QSoundEffect_DisconnectSourceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::sourceChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_SourceChanged));
}

void QSoundEffect_SourceChanged(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->sourceChanged();
}

long long QSoundEffect_Status(void* ptr)
{
	return static_cast<QSoundEffect*>(ptr)->status();
}

void QSoundEffect_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::statusChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_StatusChanged));
}

void QSoundEffect_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::statusChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_StatusChanged));
}

void QSoundEffect_StatusChanged(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->statusChanged();
}

double QSoundEffect_Volume(void* ptr)
{
	return static_cast<QSoundEffect*>(ptr)->volume();
}

void QSoundEffect_ConnectVolumeChanged(void* ptr)
{
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::volumeChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_VolumeChanged));
}

void QSoundEffect_DisconnectVolumeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::volumeChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_VolumeChanged));
}

void QSoundEffect_VolumeChanged(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->volumeChanged();
}

void QSoundEffect_DestroyQSoundEffect(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->~QSoundEffect();
}

void QSoundEffect_TimerEvent(void* ptr, void* event)
{
	static_cast<QSoundEffect*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSoundEffect_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSoundEffect*>(ptr)->QSoundEffect::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSoundEffect_ChildEvent(void* ptr, void* event)
{
	static_cast<QSoundEffect*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSoundEffect_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSoundEffect*>(ptr)->QSoundEffect::childEvent(static_cast<QChildEvent*>(event));
}

void QSoundEffect_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSoundEffect*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSoundEffect_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSoundEffect*>(ptr)->QSoundEffect::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSoundEffect_CustomEvent(void* ptr, void* event)
{
	static_cast<QSoundEffect*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSoundEffect_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSoundEffect*>(ptr)->QSoundEffect::customEvent(static_cast<QEvent*>(event));
}

void QSoundEffect_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSoundEffect*>(ptr), "deleteLater");
}

void QSoundEffect_DeleteLaterDefault(void* ptr)
{
	static_cast<QSoundEffect*>(ptr)->QSoundEffect::deleteLater();
}

void QSoundEffect_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSoundEffect*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSoundEffect_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSoundEffect*>(ptr)->QSoundEffect::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSoundEffect_Event(void* ptr, void* e)
{
	return static_cast<QSoundEffect*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSoundEffect_EventDefault(void* ptr, void* e)
{
	return static_cast<QSoundEffect*>(ptr)->QSoundEffect::event(static_cast<QEvent*>(e));
}

char QSoundEffect_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSoundEffect*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QSoundEffect_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSoundEffect*>(ptr)->QSoundEffect::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSoundEffect_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSoundEffect*>(ptr)->metaObject());
}

void* QSoundEffect_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSoundEffect*>(ptr)->QSoundEffect::metaObject());
}

class MyQVideoDeviceSelectorControl: public QVideoDeviceSelectorControl
{
public:
	MyQVideoDeviceSelectorControl(QObject *parent) : QVideoDeviceSelectorControl(parent) {};
	int defaultDevice() const { return callbackQVideoDeviceSelectorControl_DefaultDevice(const_cast<MyQVideoDeviceSelectorControl*>(this)); };
	int deviceCount() const { return callbackQVideoDeviceSelectorControl_DeviceCount(const_cast<MyQVideoDeviceSelectorControl*>(this)); };
	QString deviceDescription(int index) const { return QString(callbackQVideoDeviceSelectorControl_DeviceDescription(const_cast<MyQVideoDeviceSelectorControl*>(this), index)); };
	QString deviceName(int index) const { return QString(callbackQVideoDeviceSelectorControl_DeviceName(const_cast<MyQVideoDeviceSelectorControl*>(this), index)); };
	void Signal_DevicesChanged() { callbackQVideoDeviceSelectorControl_DevicesChanged(this); };
	int selectedDevice() const { return callbackQVideoDeviceSelectorControl_SelectedDevice(const_cast<MyQVideoDeviceSelectorControl*>(this)); };
	void Signal_SelectedDeviceChanged2(const QString & name) { QByteArray t6ae999 = name.toUtf8(); QtMultimedia_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQVideoDeviceSelectorControl_SelectedDeviceChanged2(this, namePacked); };
	void Signal_SelectedDeviceChanged(int index) { callbackQVideoDeviceSelectorControl_SelectedDeviceChanged(this, index); };
	void setSelectedDevice(int index) { callbackQVideoDeviceSelectorControl_SetSelectedDevice(this, index); };
	 ~MyQVideoDeviceSelectorControl() { callbackQVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQVideoDeviceSelectorControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQVideoDeviceSelectorControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVideoDeviceSelectorControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVideoDeviceSelectorControl_CustomEvent(this, event); };
	void deleteLater() { callbackQVideoDeviceSelectorControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVideoDeviceSelectorControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVideoDeviceSelectorControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVideoDeviceSelectorControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVideoDeviceSelectorControl_MetaObject(const_cast<MyQVideoDeviceSelectorControl*>(this))); };
};

void* QVideoDeviceSelectorControl_NewQVideoDeviceSelectorControl(void* parent)
{
	return new MyQVideoDeviceSelectorControl(static_cast<QObject*>(parent));
}

int QVideoDeviceSelectorControl_DefaultDevice(void* ptr)
{
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->defaultDevice();
}

int QVideoDeviceSelectorControl_DeviceCount(void* ptr)
{
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->deviceCount();
}

struct QtMultimedia_PackedString QVideoDeviceSelectorControl_DeviceDescription(void* ptr, int index)
{
	return ({ QByteArray t62636e = static_cast<QVideoDeviceSelectorControl*>(ptr)->deviceDescription(index).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t62636e.prepend("WHITESPACE").constData()+10), t62636e.size()-10 }; });
}

struct QtMultimedia_PackedString QVideoDeviceSelectorControl_DeviceName(void* ptr, int index)
{
	return ({ QByteArray td80c77 = static_cast<QVideoDeviceSelectorControl*>(ptr)->deviceName(index).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(td80c77.prepend("WHITESPACE").constData()+10), td80c77.size()-10 }; });
}

void QVideoDeviceSelectorControl_ConnectDevicesChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)()>(&QVideoDeviceSelectorControl::devicesChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)()>(&MyQVideoDeviceSelectorControl::Signal_DevicesChanged));
}

void QVideoDeviceSelectorControl_DisconnectDevicesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)()>(&QVideoDeviceSelectorControl::devicesChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)()>(&MyQVideoDeviceSelectorControl::Signal_DevicesChanged));
}

void QVideoDeviceSelectorControl_DevicesChanged(void* ptr)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->devicesChanged();
}

int QVideoDeviceSelectorControl_SelectedDevice(void* ptr)
{
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->selectedDevice();
}

void QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged2(void* ptr)
{
	QObject::connect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)(const QString &)>(&QVideoDeviceSelectorControl::selectedDeviceChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)(const QString &)>(&MyQVideoDeviceSelectorControl::Signal_SelectedDeviceChanged2));
}

void QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged2(void* ptr)
{
	QObject::disconnect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)(const QString &)>(&QVideoDeviceSelectorControl::selectedDeviceChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)(const QString &)>(&MyQVideoDeviceSelectorControl::Signal_SelectedDeviceChanged2));
}

void QVideoDeviceSelectorControl_SelectedDeviceChanged2(void* ptr, char* name)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->selectedDeviceChanged(QString(name));
}

void QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)(int)>(&QVideoDeviceSelectorControl::selectedDeviceChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)(int)>(&MyQVideoDeviceSelectorControl::Signal_SelectedDeviceChanged));
}

void QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)(int)>(&QVideoDeviceSelectorControl::selectedDeviceChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)(int)>(&MyQVideoDeviceSelectorControl::Signal_SelectedDeviceChanged));
}

void QVideoDeviceSelectorControl_SelectedDeviceChanged(void* ptr, int index)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->selectedDeviceChanged(index);
}

void QVideoDeviceSelectorControl_SetSelectedDevice(void* ptr, int index)
{
	QMetaObject::invokeMethod(static_cast<QVideoDeviceSelectorControl*>(ptr), "setSelectedDevice", Q_ARG(int, index));
}

void QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(void* ptr)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->~QVideoDeviceSelectorControl();
}

void QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControlDefault(void* ptr)
{

}

void QVideoDeviceSelectorControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoDeviceSelectorControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->QVideoDeviceSelectorControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoDeviceSelectorControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoDeviceSelectorControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->QVideoDeviceSelectorControl::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoDeviceSelectorControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoDeviceSelectorControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->QVideoDeviceSelectorControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoDeviceSelectorControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoDeviceSelectorControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->QVideoDeviceSelectorControl::customEvent(static_cast<QEvent*>(event));
}

void QVideoDeviceSelectorControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoDeviceSelectorControl*>(ptr), "deleteLater");
}

void QVideoDeviceSelectorControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->QVideoDeviceSelectorControl::deleteLater();
}

void QVideoDeviceSelectorControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoDeviceSelectorControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoDeviceSelectorControl*>(ptr)->QVideoDeviceSelectorControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVideoDeviceSelectorControl_Event(void* ptr, void* e)
{
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QVideoDeviceSelectorControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->QVideoDeviceSelectorControl::event(static_cast<QEvent*>(e));
}

char QVideoDeviceSelectorControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QVideoDeviceSelectorControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->QVideoDeviceSelectorControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVideoDeviceSelectorControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoDeviceSelectorControl*>(ptr)->metaObject());
}

void* QVideoDeviceSelectorControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoDeviceSelectorControl*>(ptr)->QVideoDeviceSelectorControl::metaObject());
}

void QVideoEncoderSettings_SetFrameRate(void* ptr, double rate)
{
	static_cast<QVideoEncoderSettings*>(ptr)->setFrameRate(rate);
}

void* QVideoEncoderSettings_NewQVideoEncoderSettings()
{
	return new QVideoEncoderSettings();
}

void* QVideoEncoderSettings_NewQVideoEncoderSettings2(void* other)
{
	return new QVideoEncoderSettings(*static_cast<QVideoEncoderSettings*>(other));
}

int QVideoEncoderSettings_BitRate(void* ptr)
{
	return static_cast<QVideoEncoderSettings*>(ptr)->bitRate();
}

struct QtMultimedia_PackedString QVideoEncoderSettings_Codec(void* ptr)
{
	return ({ QByteArray t09ac64 = static_cast<QVideoEncoderSettings*>(ptr)->codec().toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t09ac64.prepend("WHITESPACE").constData()+10), t09ac64.size()-10 }; });
}

long long QVideoEncoderSettings_EncodingMode(void* ptr)
{
	return static_cast<QVideoEncoderSettings*>(ptr)->encodingMode();
}

void* QVideoEncoderSettings_EncodingOption(void* ptr, char* option)
{
	return new QVariant(static_cast<QVideoEncoderSettings*>(ptr)->encodingOption(QString(option)));
}

double QVideoEncoderSettings_FrameRate(void* ptr)
{
	return static_cast<QVideoEncoderSettings*>(ptr)->frameRate();
}

char QVideoEncoderSettings_IsNull(void* ptr)
{
	return static_cast<QVideoEncoderSettings*>(ptr)->isNull();
}

long long QVideoEncoderSettings_Quality(void* ptr)
{
	return static_cast<QVideoEncoderSettings*>(ptr)->quality();
}

void* QVideoEncoderSettings_Resolution(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVideoEncoderSettings*>(ptr)->resolution(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QVideoEncoderSettings_SetBitRate(void* ptr, int value)
{
	static_cast<QVideoEncoderSettings*>(ptr)->setBitRate(value);
}

void QVideoEncoderSettings_SetCodec(void* ptr, char* codec)
{
	static_cast<QVideoEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QVideoEncoderSettings_SetEncodingMode(void* ptr, long long mode)
{
	static_cast<QVideoEncoderSettings*>(ptr)->setEncodingMode(static_cast<QMultimedia::EncodingMode>(mode));
}

void QVideoEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value)
{
	static_cast<QVideoEncoderSettings*>(ptr)->setEncodingOption(QString(option), *static_cast<QVariant*>(value));
}

void QVideoEncoderSettings_SetQuality(void* ptr, long long quality)
{
	static_cast<QVideoEncoderSettings*>(ptr)->setQuality(static_cast<QMultimedia::EncodingQuality>(quality));
}

void QVideoEncoderSettings_SetResolution(void* ptr, void* resolution)
{
	static_cast<QVideoEncoderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QVideoEncoderSettings_SetResolution2(void* ptr, int width, int height)
{
	static_cast<QVideoEncoderSettings*>(ptr)->setResolution(width, height);
}

void QVideoEncoderSettings_DestroyQVideoEncoderSettings(void* ptr)
{
	static_cast<QVideoEncoderSettings*>(ptr)->~QVideoEncoderSettings();
}

class MyQVideoEncoderSettingsControl: public QVideoEncoderSettingsControl
{
public:
	void setVideoSettings(const QVideoEncoderSettings & settings) { callbackQVideoEncoderSettingsControl_SetVideoSettings(this, const_cast<QVideoEncoderSettings*>(&settings)); };
	QStringList supportedVideoCodecs() const { return QString(callbackQVideoEncoderSettingsControl_SupportedVideoCodecs(const_cast<MyQVideoEncoderSettingsControl*>(this))).split("|", QString::SkipEmptyParts); };
	QString videoCodecDescription(const QString & codec) const { QByteArray td061f6 = codec.toUtf8(); QtMultimedia_PackedString codecPacked = { const_cast<char*>(td061f6.prepend("WHITESPACE").constData()+10), td061f6.size()-10 };return QString(callbackQVideoEncoderSettingsControl_VideoCodecDescription(const_cast<MyQVideoEncoderSettingsControl*>(this), codecPacked)); };
	QVideoEncoderSettings videoSettings() const { return *static_cast<QVideoEncoderSettings*>(callbackQVideoEncoderSettingsControl_VideoSettings(const_cast<MyQVideoEncoderSettingsControl*>(this))); };
	 ~MyQVideoEncoderSettingsControl() { callbackQVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQVideoEncoderSettingsControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQVideoEncoderSettingsControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVideoEncoderSettingsControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVideoEncoderSettingsControl_CustomEvent(this, event); };
	void deleteLater() { callbackQVideoEncoderSettingsControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVideoEncoderSettingsControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVideoEncoderSettingsControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVideoEncoderSettingsControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVideoEncoderSettingsControl_MetaObject(const_cast<MyQVideoEncoderSettingsControl*>(this))); };
};

void QVideoEncoderSettingsControl_SetVideoSettings(void* ptr, void* settings)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->setVideoSettings(*static_cast<QVideoEncoderSettings*>(settings));
}

struct QtMultimedia_PackedString QVideoEncoderSettingsControl_SupportedVideoCodecs(void* ptr)
{
	return ({ QByteArray tf4a8d7 = static_cast<QVideoEncoderSettingsControl*>(ptr)->supportedVideoCodecs().join("|").toUtf8(); QtMultimedia_PackedString { const_cast<char*>(tf4a8d7.prepend("WHITESPACE").constData()+10), tf4a8d7.size()-10 }; });
}

struct QtMultimedia_PackedString QVideoEncoderSettingsControl_VideoCodecDescription(void* ptr, char* codec)
{
	return ({ QByteArray t5ec3fe = static_cast<QVideoEncoderSettingsControl*>(ptr)->videoCodecDescription(QString(codec)).toUtf8(); QtMultimedia_PackedString { const_cast<char*>(t5ec3fe.prepend("WHITESPACE").constData()+10), t5ec3fe.size()-10 }; });
}

void* QVideoEncoderSettingsControl_VideoSettings(void* ptr)
{
	return new QVideoEncoderSettings(static_cast<QVideoEncoderSettingsControl*>(ptr)->videoSettings());
}

void QVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControl(void* ptr)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->~QVideoEncoderSettingsControl();
}

void QVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControlDefault(void* ptr)
{

}

void QVideoEncoderSettingsControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoEncoderSettingsControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->QVideoEncoderSettingsControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoEncoderSettingsControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoEncoderSettingsControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->QVideoEncoderSettingsControl::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoEncoderSettingsControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoEncoderSettingsControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->QVideoEncoderSettingsControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoEncoderSettingsControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoEncoderSettingsControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->QVideoEncoderSettingsControl::customEvent(static_cast<QEvent*>(event));
}

void QVideoEncoderSettingsControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoEncoderSettingsControl*>(ptr), "deleteLater");
}

void QVideoEncoderSettingsControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->QVideoEncoderSettingsControl::deleteLater();
}

void QVideoEncoderSettingsControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoEncoderSettingsControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoEncoderSettingsControl*>(ptr)->QVideoEncoderSettingsControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVideoEncoderSettingsControl_Event(void* ptr, void* e)
{
	return static_cast<QVideoEncoderSettingsControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QVideoEncoderSettingsControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QVideoEncoderSettingsControl*>(ptr)->QVideoEncoderSettingsControl::event(static_cast<QEvent*>(e));
}

char QVideoEncoderSettingsControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoEncoderSettingsControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QVideoEncoderSettingsControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoEncoderSettingsControl*>(ptr)->QVideoEncoderSettingsControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVideoEncoderSettingsControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoEncoderSettingsControl*>(ptr)->metaObject());
}

void* QVideoEncoderSettingsControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoEncoderSettingsControl*>(ptr)->QVideoEncoderSettingsControl::metaObject());
}

class MyQVideoFilterRunnable: public QVideoFilterRunnable
{
public:
	QVideoFrame run(QVideoFrame * input, const QVideoSurfaceFormat & surfaceFormat, QVideoFilterRunnable::RunFlags flags) { return *static_cast<QVideoFrame*>(callbackQVideoFilterRunnable_Run(this, input, const_cast<QVideoSurfaceFormat*>(&surfaceFormat), flags)); };
};

void* QVideoFilterRunnable_Run(void* ptr, void* input, void* surfaceFormat, long long flags)
{
	return new QVideoFrame(static_cast<QVideoFilterRunnable*>(ptr)->run(static_cast<QVideoFrame*>(input), *static_cast<QVideoSurfaceFormat*>(surfaceFormat), static_cast<QVideoFilterRunnable::RunFlag>(flags)));
}

void* QVideoFrame_NewQVideoFrame()
{
	return new QVideoFrame();
}

void* QVideoFrame_NewQVideoFrame2(void* buffer, void* size, long long format)
{
	return new QVideoFrame(static_cast<QAbstractVideoBuffer*>(buffer), *static_cast<QSize*>(size), static_cast<QVideoFrame::PixelFormat>(format));
}

void* QVideoFrame_NewQVideoFrame4(void* image)
{
	return new QVideoFrame(*static_cast<QImage*>(image));
}

void* QVideoFrame_NewQVideoFrame5(void* other)
{
	return new QVideoFrame(*static_cast<QVideoFrame*>(other));
}

void* QVideoFrame_NewQVideoFrame3(int bytes, void* size, int bytesPerLine, long long format)
{
	return new QVideoFrame(bytes, *static_cast<QSize*>(size), bytesPerLine, static_cast<QVideoFrame::PixelFormat>(format));
}

struct QtMultimedia_PackedString QVideoFrame_Bits(void* ptr)
{
	return ({ char* t262dd0 = static_cast<char*>(static_cast<void*>(static_cast<QVideoFrame*>(ptr)->bits())); QtMultimedia_PackedString { t262dd0, -1 }; });
}

struct QtMultimedia_PackedString QVideoFrame_Bits2(void* ptr, int plane)
{
	return ({ char* t8420e9 = static_cast<char*>(static_cast<void*>(static_cast<QVideoFrame*>(ptr)->bits(plane))); QtMultimedia_PackedString { t8420e9, -1 }; });
}

struct QtMultimedia_PackedString QVideoFrame_Bits3(void* ptr)
{
	return ({ char* t262dd0 = static_cast<char*>(static_cast<void*>(const_cast<uchar*>(static_cast<QVideoFrame*>(ptr)->bits()))); QtMultimedia_PackedString { t262dd0, -1 }; });
}

struct QtMultimedia_PackedString QVideoFrame_Bits4(void* ptr, int plane)
{
	return ({ char* t8420e9 = static_cast<char*>(static_cast<void*>(const_cast<uchar*>(static_cast<QVideoFrame*>(ptr)->bits(plane)))); QtMultimedia_PackedString { t8420e9, -1 }; });
}

int QVideoFrame_BytesPerLine(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->bytesPerLine();
}

int QVideoFrame_BytesPerLine2(void* ptr, int plane)
{
	return static_cast<QVideoFrame*>(ptr)->bytesPerLine(plane);
}

long long QVideoFrame_EndTime(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->endTime();
}

long long QVideoFrame_FieldType(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->fieldType();
}

void* QVideoFrame_Handle(void* ptr)
{
	return new QVariant(static_cast<QVideoFrame*>(ptr)->handle());
}

long long QVideoFrame_HandleType(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->handleType();
}

int QVideoFrame_Height(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->height();
}

long long QVideoFrame_QVideoFrame_ImageFormatFromPixelFormat(long long format)
{
	return QVideoFrame::imageFormatFromPixelFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

char QVideoFrame_IsMapped(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->isMapped();
}

char QVideoFrame_IsReadable(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->isReadable();
}

char QVideoFrame_IsValid(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->isValid();
}

char QVideoFrame_IsWritable(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->isWritable();
}

char QVideoFrame_Map(void* ptr, long long mode)
{
	return static_cast<QVideoFrame*>(ptr)->map(static_cast<QAbstractVideoBuffer::MapMode>(mode));
}

long long QVideoFrame_MapMode(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->mapMode();
}

int QVideoFrame_MappedBytes(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->mappedBytes();
}

void* QVideoFrame_MetaData(void* ptr, char* key)
{
	return new QVariant(static_cast<QVideoFrame*>(ptr)->metaData(QString(key)));
}

long long QVideoFrame_PixelFormat(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->pixelFormat();
}

long long QVideoFrame_QVideoFrame_PixelFormatFromImageFormat(long long format)
{
	return QVideoFrame::pixelFormatFromImageFormat(static_cast<QImage::Format>(format));
}

int QVideoFrame_PlaneCount(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->planeCount();
}

void QVideoFrame_SetEndTime(void* ptr, long long time)
{
	static_cast<QVideoFrame*>(ptr)->setEndTime(time);
}

void QVideoFrame_SetFieldType(void* ptr, long long field)
{
	static_cast<QVideoFrame*>(ptr)->setFieldType(static_cast<QVideoFrame::FieldType>(field));
}

void QVideoFrame_SetMetaData(void* ptr, char* key, void* value)
{
	static_cast<QVideoFrame*>(ptr)->setMetaData(QString(key), *static_cast<QVariant*>(value));
}

void QVideoFrame_SetStartTime(void* ptr, long long time)
{
	static_cast<QVideoFrame*>(ptr)->setStartTime(time);
}

void* QVideoFrame_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVideoFrame*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long QVideoFrame_StartTime(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->startTime();
}

void QVideoFrame_Unmap(void* ptr)
{
	static_cast<QVideoFrame*>(ptr)->unmap();
}

int QVideoFrame_Width(void* ptr)
{
	return static_cast<QVideoFrame*>(ptr)->width();
}

void QVideoFrame_DestroyQVideoFrame(void* ptr)
{
	static_cast<QVideoFrame*>(ptr)->~QVideoFrame();
}

class MyQVideoProbe: public QVideoProbe
{
public:
	MyQVideoProbe(QObject *parent) : QVideoProbe(parent) {};
	void Signal_Flush() { callbackQVideoProbe_Flush(this); };
	void Signal_VideoFrameProbed(const QVideoFrame & frame) { callbackQVideoProbe_VideoFrameProbed(this, const_cast<QVideoFrame*>(&frame)); };
	void timerEvent(QTimerEvent * event) { callbackQVideoProbe_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQVideoProbe_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVideoProbe_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVideoProbe_CustomEvent(this, event); };
	void deleteLater() { callbackQVideoProbe_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVideoProbe_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVideoProbe_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVideoProbe_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVideoProbe_MetaObject(const_cast<MyQVideoProbe*>(this))); };
};

void* QVideoProbe_NewQVideoProbe(void* parent)
{
	return new MyQVideoProbe(static_cast<QObject*>(parent));
}

void QVideoProbe_ConnectFlush(void* ptr)
{
	QObject::connect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)()>(&QVideoProbe::flush), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)()>(&MyQVideoProbe::Signal_Flush));
}

void QVideoProbe_DisconnectFlush(void* ptr)
{
	QObject::disconnect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)()>(&QVideoProbe::flush), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)()>(&MyQVideoProbe::Signal_Flush));
}

void QVideoProbe_Flush(void* ptr)
{
	static_cast<QVideoProbe*>(ptr)->flush();
}

char QVideoProbe_IsActive(void* ptr)
{
	return static_cast<QVideoProbe*>(ptr)->isActive();
}

char QVideoProbe_SetSource(void* ptr, void* source)
{
	return static_cast<QVideoProbe*>(ptr)->setSource(static_cast<QMediaObject*>(source));
}

char QVideoProbe_SetSource2(void* ptr, void* mediaRecorder)
{
	return static_cast<QVideoProbe*>(ptr)->setSource(static_cast<QMediaRecorder*>(mediaRecorder));
}

void QVideoProbe_ConnectVideoFrameProbed(void* ptr)
{
	QObject::connect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)(const QVideoFrame &)>(&QVideoProbe::videoFrameProbed), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)(const QVideoFrame &)>(&MyQVideoProbe::Signal_VideoFrameProbed));
}

void QVideoProbe_DisconnectVideoFrameProbed(void* ptr)
{
	QObject::disconnect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)(const QVideoFrame &)>(&QVideoProbe::videoFrameProbed), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)(const QVideoFrame &)>(&MyQVideoProbe::Signal_VideoFrameProbed));
}

void QVideoProbe_VideoFrameProbed(void* ptr, void* frame)
{
	static_cast<QVideoProbe*>(ptr)->videoFrameProbed(*static_cast<QVideoFrame*>(frame));
}

void QVideoProbe_DestroyQVideoProbe(void* ptr)
{
	static_cast<QVideoProbe*>(ptr)->~QVideoProbe();
}

void QVideoProbe_TimerEvent(void* ptr, void* event)
{
	static_cast<QVideoProbe*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoProbe_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVideoProbe*>(ptr)->QVideoProbe::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoProbe_ChildEvent(void* ptr, void* event)
{
	static_cast<QVideoProbe*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoProbe_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QVideoProbe*>(ptr)->QVideoProbe::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoProbe_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoProbe*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoProbe_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoProbe*>(ptr)->QVideoProbe::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoProbe_CustomEvent(void* ptr, void* event)
{
	static_cast<QVideoProbe*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoProbe_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVideoProbe*>(ptr)->QVideoProbe::customEvent(static_cast<QEvent*>(event));
}

void QVideoProbe_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoProbe*>(ptr), "deleteLater");
}

void QVideoProbe_DeleteLaterDefault(void* ptr)
{
	static_cast<QVideoProbe*>(ptr)->QVideoProbe::deleteLater();
}

void QVideoProbe_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoProbe*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoProbe_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoProbe*>(ptr)->QVideoProbe::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVideoProbe_Event(void* ptr, void* e)
{
	return static_cast<QVideoProbe*>(ptr)->event(static_cast<QEvent*>(e));
}

char QVideoProbe_EventDefault(void* ptr, void* e)
{
	return static_cast<QVideoProbe*>(ptr)->QVideoProbe::event(static_cast<QEvent*>(e));
}

char QVideoProbe_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoProbe*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QVideoProbe_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoProbe*>(ptr)->QVideoProbe::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVideoProbe_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoProbe*>(ptr)->metaObject());
}

void* QVideoProbe_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoProbe*>(ptr)->QVideoProbe::metaObject());
}

class MyQVideoRendererControl: public QVideoRendererControl
{
public:
	MyQVideoRendererControl(QObject *parent) : QVideoRendererControl(parent) {};
	void setSurface(QAbstractVideoSurface * surface) { callbackQVideoRendererControl_SetSurface(this, surface); };
	QAbstractVideoSurface * surface() const { return static_cast<QAbstractVideoSurface*>(callbackQVideoRendererControl_Surface(const_cast<MyQVideoRendererControl*>(this))); };
	void timerEvent(QTimerEvent * event) { callbackQVideoRendererControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQVideoRendererControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVideoRendererControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVideoRendererControl_CustomEvent(this, event); };
	void deleteLater() { callbackQVideoRendererControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVideoRendererControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVideoRendererControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVideoRendererControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVideoRendererControl_MetaObject(const_cast<MyQVideoRendererControl*>(this))); };
};

void* QVideoRendererControl_NewQVideoRendererControl(void* parent)
{
	return new MyQVideoRendererControl(static_cast<QObject*>(parent));
}

void QVideoRendererControl_SetSurface(void* ptr, void* surface)
{
	static_cast<QVideoRendererControl*>(ptr)->setSurface(static_cast<QAbstractVideoSurface*>(surface));
}

void* QVideoRendererControl_Surface(void* ptr)
{
	return static_cast<QVideoRendererControl*>(ptr)->surface();
}

void QVideoRendererControl_DestroyQVideoRendererControl(void* ptr)
{
	static_cast<QVideoRendererControl*>(ptr)->~QVideoRendererControl();
}

void QVideoRendererControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QVideoRendererControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoRendererControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVideoRendererControl*>(ptr)->QVideoRendererControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoRendererControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QVideoRendererControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoRendererControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QVideoRendererControl*>(ptr)->QVideoRendererControl::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoRendererControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoRendererControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoRendererControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoRendererControl*>(ptr)->QVideoRendererControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoRendererControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QVideoRendererControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoRendererControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVideoRendererControl*>(ptr)->QVideoRendererControl::customEvent(static_cast<QEvent*>(event));
}

void QVideoRendererControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoRendererControl*>(ptr), "deleteLater");
}

void QVideoRendererControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QVideoRendererControl*>(ptr)->QVideoRendererControl::deleteLater();
}

void QVideoRendererControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoRendererControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoRendererControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoRendererControl*>(ptr)->QVideoRendererControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVideoRendererControl_Event(void* ptr, void* e)
{
	return static_cast<QVideoRendererControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QVideoRendererControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QVideoRendererControl*>(ptr)->QVideoRendererControl::event(static_cast<QEvent*>(e));
}

char QVideoRendererControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoRendererControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QVideoRendererControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoRendererControl*>(ptr)->QVideoRendererControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVideoRendererControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoRendererControl*>(ptr)->metaObject());
}

void* QVideoRendererControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoRendererControl*>(ptr)->QVideoRendererControl::metaObject());
}

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat()
{
	return new QVideoSurfaceFormat();
}

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat2(void* size, long long format, long long ty)
{
	return new QVideoSurfaceFormat(*static_cast<QSize*>(size), static_cast<QVideoFrame::PixelFormat>(format), static_cast<QAbstractVideoBuffer::HandleType>(ty));
}

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat3(void* other)
{
	return new QVideoSurfaceFormat(*static_cast<QVideoSurfaceFormat*>(other));
}

int QVideoSurfaceFormat_FrameHeight(void* ptr)
{
	return static_cast<QVideoSurfaceFormat*>(ptr)->frameHeight();
}

double QVideoSurfaceFormat_FrameRate(void* ptr)
{
	return static_cast<QVideoSurfaceFormat*>(ptr)->frameRate();
}

void* QVideoSurfaceFormat_FrameSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVideoSurfaceFormat*>(ptr)->frameSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QVideoSurfaceFormat_FrameWidth(void* ptr)
{
	return static_cast<QVideoSurfaceFormat*>(ptr)->frameWidth();
}

long long QVideoSurfaceFormat_HandleType(void* ptr)
{
	return static_cast<QVideoSurfaceFormat*>(ptr)->handleType();
}

char QVideoSurfaceFormat_IsValid(void* ptr)
{
	return static_cast<QVideoSurfaceFormat*>(ptr)->isValid();
}

void* QVideoSurfaceFormat_PixelAspectRatio(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVideoSurfaceFormat*>(ptr)->pixelAspectRatio(); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long QVideoSurfaceFormat_PixelFormat(void* ptr)
{
	return static_cast<QVideoSurfaceFormat*>(ptr)->pixelFormat();
}

void* QVideoSurfaceFormat_Property(void* ptr, char* name)
{
	return new QVariant(static_cast<QVideoSurfaceFormat*>(ptr)->property(const_cast<const char*>(name)));
}

long long QVideoSurfaceFormat_ScanLineDirection(void* ptr)
{
	return static_cast<QVideoSurfaceFormat*>(ptr)->scanLineDirection();
}

void QVideoSurfaceFormat_SetFrameRate(void* ptr, double rate)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameRate(rate);
}

void QVideoSurfaceFormat_SetFrameSize(void* ptr, void* size)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameSize(*static_cast<QSize*>(size));
}

void QVideoSurfaceFormat_SetFrameSize2(void* ptr, int width, int height)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameSize(width, height);
}

void QVideoSurfaceFormat_SetPixelAspectRatio(void* ptr, void* ratio)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->setPixelAspectRatio(*static_cast<QSize*>(ratio));
}

void QVideoSurfaceFormat_SetPixelAspectRatio2(void* ptr, int horizontal, int vertical)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->setPixelAspectRatio(horizontal, vertical);
}

void QVideoSurfaceFormat_SetProperty(void* ptr, char* name, void* value)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->setProperty(const_cast<const char*>(name), *static_cast<QVariant*>(value));
}

void QVideoSurfaceFormat_SetScanLineDirection(void* ptr, long long direction)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->setScanLineDirection(static_cast<QVideoSurfaceFormat::Direction>(direction));
}

void QVideoSurfaceFormat_SetViewport(void* ptr, void* viewport)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->setViewport(*static_cast<QRect*>(viewport));
}

void QVideoSurfaceFormat_SetYCbCrColorSpace(void* ptr, long long space)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->setYCbCrColorSpace(static_cast<QVideoSurfaceFormat::YCbCrColorSpace>(space));
}

void* QVideoSurfaceFormat_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVideoSurfaceFormat*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QVideoSurfaceFormat_Viewport(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QVideoSurfaceFormat*>(ptr)->viewport(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

long long QVideoSurfaceFormat_YCbCrColorSpace(void* ptr)
{
	return static_cast<QVideoSurfaceFormat*>(ptr)->yCbCrColorSpace();
}

void QVideoSurfaceFormat_DestroyQVideoSurfaceFormat(void* ptr)
{
	static_cast<QVideoSurfaceFormat*>(ptr)->~QVideoSurfaceFormat();
}

class MyQVideoWidget: public QVideoWidget
{
public:
	MyQVideoWidget(QWidget *parent) : QVideoWidget(parent) {};
	QMediaObject * mediaObject() const { return static_cast<QMediaObject*>(callbackQVideoWidget_MediaObject(const_cast<MyQVideoWidget*>(this))); };
	void setAspectRatioMode(Qt::AspectRatioMode mode) { callbackQVideoWidget_SetAspectRatioMode(this, mode); };
	void setBrightness(int brightness) { callbackQVideoWidget_SetBrightness(this, brightness); };
	void setContrast(int contrast) { callbackQVideoWidget_SetContrast(this, contrast); };
	void setFullScreen(bool fullScreen) { callbackQVideoWidget_SetFullScreen(this, fullScreen); };
	void setHue(int hue) { callbackQVideoWidget_SetHue(this, hue); };
	void setSaturation(int saturation) { callbackQVideoWidget_SetSaturation(this, saturation); };
	void Signal_BrightnessChanged(int brightness) { callbackQVideoWidget_BrightnessChanged(this, brightness); };
	void Signal_ContrastChanged(int contrast) { callbackQVideoWidget_ContrastChanged(this, contrast); };
	void Signal_FullScreenChanged(bool fullScreen) { callbackQVideoWidget_FullScreenChanged(this, fullScreen); };
	void Signal_HueChanged(int hue) { callbackQVideoWidget_HueChanged(this, hue); };
	void Signal_SaturationChanged(int saturation) { callbackQVideoWidget_SaturationChanged(this, saturation); };
	void actionEvent(QActionEvent * event) { callbackQVideoWidget_ActionEvent(this, event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQVideoWidget_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQVideoWidget_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQVideoWidget_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQVideoWidget_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQVideoWidget_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQVideoWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQVideoWidget_FocusOutEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQVideoWidget_LeaveEvent(this, event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQVideoWidget_MinimumSizeHint(const_cast<MyQVideoWidget*>(this))); };
	void setEnabled(bool vbo) { callbackQVideoWidget_SetEnabled(this, vbo); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtMultimedia_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQVideoWidget_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQVideoWidget_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQVideoWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtMultimedia_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQVideoWidget_SetWindowTitle(this, vqsPacked); };
	void changeEvent(QEvent * event) { callbackQVideoWidget_ChangeEvent(this, event); };
	bool close() { return callbackQVideoWidget_Close(this) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQVideoWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQVideoWidget_ContextMenuEvent(this, event); };
	bool focusNextPrevChild(bool next) { return callbackQVideoWidget_FocusNextPrevChild(this, next) != 0; };
	bool hasHeightForWidth() const { return callbackQVideoWidget_HasHeightForWidth(const_cast<MyQVideoWidget*>(this)) != 0; };
	int heightForWidth(int w) const { return callbackQVideoWidget_HeightForWidth(const_cast<MyQVideoWidget*>(this), w); };
	void hide() { callbackQVideoWidget_Hide(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQVideoWidget_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQVideoWidget_InputMethodQuery(const_cast<MyQVideoWidget*>(this), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQVideoWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQVideoWidget_KeyReleaseEvent(this, event); };
	void lower() { callbackQVideoWidget_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQVideoWidget_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQVideoWidget_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQVideoWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQVideoWidget_MouseReleaseEvent(this, event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQVideoWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void raise() { callbackQVideoWidget_Raise(this); };
	void repaint() { callbackQVideoWidget_Repaint(this); };
	void setDisabled(bool disable) { callbackQVideoWidget_SetDisabled(this, disable); };
	void setFocus() { callbackQVideoWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQVideoWidget_SetHidden(this, hidden); };
	void show() { callbackQVideoWidget_Show(this); };
	void showFullScreen() { callbackQVideoWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQVideoWidget_ShowMaximized(this); };
	void showMinimized() { callbackQVideoWidget_ShowMinimized(this); };
	void showNormal() { callbackQVideoWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQVideoWidget_TabletEvent(this, event); };
	void update() { callbackQVideoWidget_Update(this); };
	void updateMicroFocus() { callbackQVideoWidget_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQVideoWidget_WheelEvent(this, event); };
	void timerEvent(QTimerEvent * event) { callbackQVideoWidget_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQVideoWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVideoWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVideoWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQVideoWidget_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVideoWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVideoWidget_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVideoWidget_MetaObject(const_cast<MyQVideoWidget*>(this))); };
	bool setMediaObject(QMediaObject * object) { return callbackQVideoWidget_SetMediaObject(this, object) != 0; };
};

long long QVideoWidget_AspectRatioMode(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->aspectRatioMode();
}

int QVideoWidget_Brightness(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->brightness();
}

int QVideoWidget_Contrast(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->contrast();
}

int QVideoWidget_Hue(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->hue();
}

void* QVideoWidget_MediaObject(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->mediaObject();
}

void* QVideoWidget_MediaObjectDefault(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::mediaObject();
}

int QVideoWidget_Saturation(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->saturation();
}

void QVideoWidget_SetAspectRatioMode(void* ptr, long long mode)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setAspectRatioMode", Q_ARG(Qt::AspectRatioMode, static_cast<Qt::AspectRatioMode>(mode)));
}

void QVideoWidget_SetBrightness(void* ptr, int brightness)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setBrightness", Q_ARG(int, brightness));
}

void QVideoWidget_SetContrast(void* ptr, int contrast)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setContrast", Q_ARG(int, contrast));
}

void QVideoWidget_SetFullScreen(void* ptr, char fullScreen)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setFullScreen", Q_ARG(bool, fullScreen != 0));
}

void QVideoWidget_SetHue(void* ptr, int hue)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setHue", Q_ARG(int, hue));
}

void QVideoWidget_SetSaturation(void* ptr, int saturation)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setSaturation", Q_ARG(int, saturation));
}

void* QVideoWidget_NewQVideoWidget(void* parent)
{
	return new MyQVideoWidget(static_cast<QWidget*>(parent));
}

void QVideoWidget_ConnectBrightnessChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::brightnessChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_BrightnessChanged));
}

void QVideoWidget_DisconnectBrightnessChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::brightnessChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_BrightnessChanged));
}

void QVideoWidget_BrightnessChanged(void* ptr, int brightness)
{
	static_cast<QVideoWidget*>(ptr)->brightnessChanged(brightness);
}

void QVideoWidget_ConnectContrastChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::contrastChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_ContrastChanged));
}

void QVideoWidget_DisconnectContrastChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::contrastChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_ContrastChanged));
}

void QVideoWidget_ContrastChanged(void* ptr, int contrast)
{
	static_cast<QVideoWidget*>(ptr)->contrastChanged(contrast);
}

char QVideoWidget_Event(void* ptr, void* event)
{
	return static_cast<QVideoWidget*>(ptr)->event(static_cast<QEvent*>(event));
}

void QVideoWidget_ConnectFullScreenChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(bool)>(&QVideoWidget::fullScreenChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(bool)>(&MyQVideoWidget::Signal_FullScreenChanged));
}

void QVideoWidget_DisconnectFullScreenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(bool)>(&QVideoWidget::fullScreenChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(bool)>(&MyQVideoWidget::Signal_FullScreenChanged));
}

void QVideoWidget_FullScreenChanged(void* ptr, char fullScreen)
{
	static_cast<QVideoWidget*>(ptr)->fullScreenChanged(fullScreen != 0);
}

void QVideoWidget_HideEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QVideoWidget_ConnectHueChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::hueChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_HueChanged));
}

void QVideoWidget_DisconnectHueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::hueChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_HueChanged));
}

void QVideoWidget_HueChanged(void* ptr, int hue)
{
	static_cast<QVideoWidget*>(ptr)->hueChanged(hue);
}

char QVideoWidget_IsFullScreen(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->isFullScreen();
}

void QVideoWidget_MoveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QVideoWidget_PaintEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QVideoWidget_ResizeEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QVideoWidget_ConnectSaturationChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::saturationChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_SaturationChanged));
}

void QVideoWidget_DisconnectSaturationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::saturationChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_SaturationChanged));
}

void QVideoWidget_SaturationChanged(void* ptr, int saturation)
{
	static_cast<QVideoWidget*>(ptr)->saturationChanged(saturation);
}

void QVideoWidget_ShowEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void* QVideoWidget_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVideoWidget*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QVideoWidget_DestroyQVideoWidget(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->~QVideoWidget();
}

void QVideoWidget_ActionEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QVideoWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QVideoWidget_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QVideoWidget_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QVideoWidget_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QVideoWidget_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QVideoWidget_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QVideoWidget_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QVideoWidget_DropEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QVideoWidget_DropEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QVideoWidget_EnterEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::enterEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_FocusInEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QVideoWidget_LeaveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::leaveEvent(static_cast<QEvent*>(event));
}

void* QVideoWidget_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVideoWidget*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QVideoWidget_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVideoWidget*>(ptr)->QVideoWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QVideoWidget_SetEnabled(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QVideoWidget_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::setEnabled(vbo != 0);
}

void QVideoWidget_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QVideoWidget_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::setStyleSheet(QString(styleSheet));
}

void QVideoWidget_SetVisible(void* ptr, char visible)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QVideoWidget_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::setVisible(visible != 0);
}

void QVideoWidget_SetWindowModified(void* ptr, char vbo)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QVideoWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::setWindowModified(vbo != 0);
}

void QVideoWidget_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QVideoWidget_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::setWindowTitle(QString(vqs));
}

void QVideoWidget_ChangeEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::changeEvent(static_cast<QEvent*>(event));
}

char QVideoWidget_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QVideoWidget_CloseDefault(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::close();
}

void QVideoWidget_CloseEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QVideoWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QVideoWidget_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QVideoWidget_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

char QVideoWidget_FocusNextPrevChild(void* ptr, char next)
{
	return static_cast<QVideoWidget*>(ptr)->focusNextPrevChild(next != 0);
}

char QVideoWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::focusNextPrevChild(next != 0);
}

char QVideoWidget_HasHeightForWidth(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->hasHeightForWidth();
}

char QVideoWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::hasHeightForWidth();
}

int QVideoWidget_HeightForWidth(void* ptr, int w)
{
	return static_cast<QVideoWidget*>(ptr)->heightForWidth(w);
}

int QVideoWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::heightForWidth(w);
}

void QVideoWidget_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "hide");
}

void QVideoWidget_HideDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::hide();
}

void QVideoWidget_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QVideoWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QVideoWidget_InputMethodQuery(void* ptr, long long query)
{
	return new QVariant(static_cast<QVideoWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QVideoWidget_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<QVideoWidget*>(ptr)->QVideoWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QVideoWidget_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QVideoWidget_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "lower");
}

void QVideoWidget_LowerDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::lower();
}

void QVideoWidget_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MousePressEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QVideoWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

char QVideoWidget_NativeEvent(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QVideoWidget*>(ptr)->nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

char QVideoWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QVideoWidget_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "raise");
}

void QVideoWidget_RaiseDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::raise();
}

void QVideoWidget_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "repaint");
}

void QVideoWidget_RepaintDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::repaint();
}

void QVideoWidget_SetDisabled(void* ptr, char disable)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QVideoWidget_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::setDisabled(disable != 0);
}

void QVideoWidget_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setFocus");
}

void QVideoWidget_SetFocus2Default(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::setFocus();
}

void QVideoWidget_SetHidden(void* ptr, char hidden)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QVideoWidget_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::setHidden(hidden != 0);
}

void QVideoWidget_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "show");
}

void QVideoWidget_ShowDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::show();
}

void QVideoWidget_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "showFullScreen");
}

void QVideoWidget_ShowFullScreenDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::showFullScreen();
}

void QVideoWidget_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "showMaximized");
}

void QVideoWidget_ShowMaximizedDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::showMaximized();
}

void QVideoWidget_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "showMinimized");
}

void QVideoWidget_ShowMinimizedDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::showMinimized();
}

void QVideoWidget_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "showNormal");
}

void QVideoWidget_ShowNormalDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::showNormal();
}

void QVideoWidget_TabletEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QVideoWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QVideoWidget_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "update");
}

void QVideoWidget_UpdateDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::update();
}

void QVideoWidget_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "updateMicroFocus");
}

void QVideoWidget_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::updateMicroFocus();
}

void QVideoWidget_WheelEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QVideoWidget_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QVideoWidget_TimerEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidget_ChildEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidget_CustomEvent(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::customEvent(static_cast<QEvent*>(event));
}

void QVideoWidget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "deleteLater");
}

void QVideoWidget_DeleteLaterDefault(void* ptr)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::deleteLater();
}

void QVideoWidget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWidget*>(ptr)->QVideoWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVideoWidget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QVideoWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVideoWidget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWidget*>(ptr)->metaObject());
}

void* QVideoWidget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWidget*>(ptr)->QVideoWidget::metaObject());
}

char QVideoWidget_SetMediaObject(void* ptr, void* object)
{
	return static_cast<QVideoWidget*>(ptr)->setMediaObject(static_cast<QMediaObject*>(object));
}

char QVideoWidget_SetMediaObjectDefault(void* ptr, void* object)
{
	return static_cast<QVideoWidget*>(ptr)->QVideoWidget::setMediaObject(static_cast<QMediaObject*>(object));
}

class MyQVideoWidgetControl: public QVideoWidgetControl
{
public:
	MyQVideoWidgetControl(QObject *parent) : QVideoWidgetControl(parent) {};
	Qt::AspectRatioMode aspectRatioMode() const { return static_cast<Qt::AspectRatioMode>(callbackQVideoWidgetControl_AspectRatioMode(const_cast<MyQVideoWidgetControl*>(this))); };
	int brightness() const { return callbackQVideoWidgetControl_Brightness(const_cast<MyQVideoWidgetControl*>(this)); };
	void Signal_BrightnessChanged(int brightness) { callbackQVideoWidgetControl_BrightnessChanged(this, brightness); };
	int contrast() const { return callbackQVideoWidgetControl_Contrast(const_cast<MyQVideoWidgetControl*>(this)); };
	void Signal_ContrastChanged(int contrast) { callbackQVideoWidgetControl_ContrastChanged(this, contrast); };
	void Signal_FullScreenChanged(bool fullScreen) { callbackQVideoWidgetControl_FullScreenChanged(this, fullScreen); };
	int hue() const { return callbackQVideoWidgetControl_Hue(const_cast<MyQVideoWidgetControl*>(this)); };
	void Signal_HueChanged(int hue) { callbackQVideoWidgetControl_HueChanged(this, hue); };
	bool isFullScreen() const { return callbackQVideoWidgetControl_IsFullScreen(const_cast<MyQVideoWidgetControl*>(this)) != 0; };
	int saturation() const { return callbackQVideoWidgetControl_Saturation(const_cast<MyQVideoWidgetControl*>(this)); };
	void Signal_SaturationChanged(int saturation) { callbackQVideoWidgetControl_SaturationChanged(this, saturation); };
	void setAspectRatioMode(Qt::AspectRatioMode mode) { callbackQVideoWidgetControl_SetAspectRatioMode(this, mode); };
	void setBrightness(int brightness) { callbackQVideoWidgetControl_SetBrightness(this, brightness); };
	void setContrast(int contrast) { callbackQVideoWidgetControl_SetContrast(this, contrast); };
	void setFullScreen(bool fullScreen) { callbackQVideoWidgetControl_SetFullScreen(this, fullScreen); };
	void setHue(int hue) { callbackQVideoWidgetControl_SetHue(this, hue); };
	void setSaturation(int saturation) { callbackQVideoWidgetControl_SetSaturation(this, saturation); };
	QWidget * videoWidget() { return static_cast<QWidget*>(callbackQVideoWidgetControl_VideoWidget(this)); };
	 ~MyQVideoWidgetControl() { callbackQVideoWidgetControl_DestroyQVideoWidgetControl(this); };
	void timerEvent(QTimerEvent * event) { callbackQVideoWidgetControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQVideoWidgetControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVideoWidgetControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVideoWidgetControl_CustomEvent(this, event); };
	void deleteLater() { callbackQVideoWidgetControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVideoWidgetControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVideoWidgetControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVideoWidgetControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVideoWidgetControl_MetaObject(const_cast<MyQVideoWidgetControl*>(this))); };
};

void* QVideoWidgetControl_NewQVideoWidgetControl(void* parent)
{
	return new MyQVideoWidgetControl(static_cast<QObject*>(parent));
}

long long QVideoWidgetControl_AspectRatioMode(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->aspectRatioMode();
}

int QVideoWidgetControl_Brightness(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->brightness();
}

void QVideoWidgetControl_ConnectBrightnessChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));
}

void QVideoWidgetControl_DisconnectBrightnessChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));
}

void QVideoWidgetControl_BrightnessChanged(void* ptr, int brightness)
{
	static_cast<QVideoWidgetControl*>(ptr)->brightnessChanged(brightness);
}

int QVideoWidgetControl_Contrast(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->contrast();
}

void QVideoWidgetControl_ConnectContrastChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));
}

void QVideoWidgetControl_DisconnectContrastChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));
}

void QVideoWidgetControl_ContrastChanged(void* ptr, int contrast)
{
	static_cast<QVideoWidgetControl*>(ptr)->contrastChanged(contrast);
}

void QVideoWidgetControl_ConnectFullScreenChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));
}

void QVideoWidgetControl_DisconnectFullScreenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));
}

void QVideoWidgetControl_FullScreenChanged(void* ptr, char fullScreen)
{
	static_cast<QVideoWidgetControl*>(ptr)->fullScreenChanged(fullScreen != 0);
}

int QVideoWidgetControl_Hue(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->hue();
}

void QVideoWidgetControl_ConnectHueChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));
}

void QVideoWidgetControl_DisconnectHueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));
}

void QVideoWidgetControl_HueChanged(void* ptr, int hue)
{
	static_cast<QVideoWidgetControl*>(ptr)->hueChanged(hue);
}

char QVideoWidgetControl_IsFullScreen(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->isFullScreen();
}

int QVideoWidgetControl_Saturation(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->saturation();
}

void QVideoWidgetControl_ConnectSaturationChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));
}

void QVideoWidgetControl_DisconnectSaturationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));
}

void QVideoWidgetControl_SaturationChanged(void* ptr, int saturation)
{
	static_cast<QVideoWidgetControl*>(ptr)->saturationChanged(saturation);
}

void QVideoWidgetControl_SetAspectRatioMode(void* ptr, long long mode)
{
	static_cast<QVideoWidgetControl*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QVideoWidgetControl_SetBrightness(void* ptr, int brightness)
{
	static_cast<QVideoWidgetControl*>(ptr)->setBrightness(brightness);
}

void QVideoWidgetControl_SetContrast(void* ptr, int contrast)
{
	static_cast<QVideoWidgetControl*>(ptr)->setContrast(contrast);
}

void QVideoWidgetControl_SetFullScreen(void* ptr, char fullScreen)
{
	static_cast<QVideoWidgetControl*>(ptr)->setFullScreen(fullScreen != 0);
}

void QVideoWidgetControl_SetHue(void* ptr, int hue)
{
	static_cast<QVideoWidgetControl*>(ptr)->setHue(hue);
}

void QVideoWidgetControl_SetSaturation(void* ptr, int saturation)
{
	static_cast<QVideoWidgetControl*>(ptr)->setSaturation(saturation);
}

void* QVideoWidgetControl_VideoWidget(void* ptr)
{
	return static_cast<QVideoWidgetControl*>(ptr)->videoWidget();
}

void QVideoWidgetControl_DestroyQVideoWidgetControl(void* ptr)
{
	static_cast<QVideoWidgetControl*>(ptr)->~QVideoWidgetControl();
}

void QVideoWidgetControl_DestroyQVideoWidgetControlDefault(void* ptr)
{

}

void QVideoWidgetControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidgetControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWidgetControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidgetControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWidgetControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWidgetControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidgetControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidgetControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoWidgetControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::customEvent(static_cast<QEvent*>(event));
}

void QVideoWidgetControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWidgetControl*>(ptr), "deleteLater");
}

void QVideoWidgetControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::deleteLater();
}

void QVideoWidgetControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWidgetControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWidgetControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVideoWidgetControl_Event(void* ptr, void* e)
{
	return static_cast<QVideoWidgetControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QVideoWidgetControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::event(static_cast<QEvent*>(e));
}

char QVideoWidgetControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWidgetControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QVideoWidgetControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVideoWidgetControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWidgetControl*>(ptr)->metaObject());
}

void* QVideoWidgetControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWidgetControl*>(ptr)->QVideoWidgetControl::metaObject());
}

class MyQVideoWindowControl: public QVideoWindowControl
{
public:
	MyQVideoWindowControl(QObject *parent) : QVideoWindowControl(parent) {};
	Qt::AspectRatioMode aspectRatioMode() const { return static_cast<Qt::AspectRatioMode>(callbackQVideoWindowControl_AspectRatioMode(const_cast<MyQVideoWindowControl*>(this))); };
	int brightness() const { return callbackQVideoWindowControl_Brightness(const_cast<MyQVideoWindowControl*>(this)); };
	void Signal_BrightnessChanged(int brightness) { callbackQVideoWindowControl_BrightnessChanged(this, brightness); };
	int contrast() const { return callbackQVideoWindowControl_Contrast(const_cast<MyQVideoWindowControl*>(this)); };
	void Signal_ContrastChanged(int contrast) { callbackQVideoWindowControl_ContrastChanged(this, contrast); };
	QRect displayRect() const { return *static_cast<QRect*>(callbackQVideoWindowControl_DisplayRect(const_cast<MyQVideoWindowControl*>(this))); };
	void Signal_FullScreenChanged(bool fullScreen) { callbackQVideoWindowControl_FullScreenChanged(this, fullScreen); };
	int hue() const { return callbackQVideoWindowControl_Hue(const_cast<MyQVideoWindowControl*>(this)); };
	void Signal_HueChanged(int hue) { callbackQVideoWindowControl_HueChanged(this, hue); };
	bool isFullScreen() const { return callbackQVideoWindowControl_IsFullScreen(const_cast<MyQVideoWindowControl*>(this)) != 0; };
	QSize nativeSize() const { return *static_cast<QSize*>(callbackQVideoWindowControl_NativeSize(const_cast<MyQVideoWindowControl*>(this))); };
	void Signal_NativeSizeChanged() { callbackQVideoWindowControl_NativeSizeChanged(this); };
	void repaint() { callbackQVideoWindowControl_Repaint(this); };
	int saturation() const { return callbackQVideoWindowControl_Saturation(const_cast<MyQVideoWindowControl*>(this)); };
	void Signal_SaturationChanged(int saturation) { callbackQVideoWindowControl_SaturationChanged(this, saturation); };
	void setAspectRatioMode(Qt::AspectRatioMode mode) { callbackQVideoWindowControl_SetAspectRatioMode(this, mode); };
	void setBrightness(int brightness) { callbackQVideoWindowControl_SetBrightness(this, brightness); };
	void setContrast(int contrast) { callbackQVideoWindowControl_SetContrast(this, contrast); };
	void setDisplayRect(const QRect & rect) { callbackQVideoWindowControl_SetDisplayRect(this, const_cast<QRect*>(&rect)); };
	void setFullScreen(bool fullScreen) { callbackQVideoWindowControl_SetFullScreen(this, fullScreen); };
	void setHue(int hue) { callbackQVideoWindowControl_SetHue(this, hue); };
	void setSaturation(int saturation) { callbackQVideoWindowControl_SetSaturation(this, saturation); };
	void setWinId(WId id) { callbackQVideoWindowControl_SetWinId(this, id); };
	WId winId() const { return callbackQVideoWindowControl_WinId(const_cast<MyQVideoWindowControl*>(this)); };
	void timerEvent(QTimerEvent * event) { callbackQVideoWindowControl_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQVideoWindowControl_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQVideoWindowControl_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVideoWindowControl_CustomEvent(this, event); };
	void deleteLater() { callbackQVideoWindowControl_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVideoWindowControl_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQVideoWindowControl_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVideoWindowControl_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQVideoWindowControl_MetaObject(const_cast<MyQVideoWindowControl*>(this))); };
};

void* QVideoWindowControl_NewQVideoWindowControl(void* parent)
{
	return new MyQVideoWindowControl(static_cast<QObject*>(parent));
}

long long QVideoWindowControl_AspectRatioMode(void* ptr)
{
	return static_cast<QVideoWindowControl*>(ptr)->aspectRatioMode();
}

int QVideoWindowControl_Brightness(void* ptr)
{
	return static_cast<QVideoWindowControl*>(ptr)->brightness();
}

void QVideoWindowControl_ConnectBrightnessChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::brightnessChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_BrightnessChanged));
}

void QVideoWindowControl_DisconnectBrightnessChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::brightnessChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_BrightnessChanged));
}

void QVideoWindowControl_BrightnessChanged(void* ptr, int brightness)
{
	static_cast<QVideoWindowControl*>(ptr)->brightnessChanged(brightness);
}

int QVideoWindowControl_Contrast(void* ptr)
{
	return static_cast<QVideoWindowControl*>(ptr)->contrast();
}

void QVideoWindowControl_ConnectContrastChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::contrastChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_ContrastChanged));
}

void QVideoWindowControl_DisconnectContrastChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::contrastChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_ContrastChanged));
}

void QVideoWindowControl_ContrastChanged(void* ptr, int contrast)
{
	static_cast<QVideoWindowControl*>(ptr)->contrastChanged(contrast);
}

void* QVideoWindowControl_DisplayRect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QVideoWindowControl*>(ptr)->displayRect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QVideoWindowControl_ConnectFullScreenChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(bool)>(&QVideoWindowControl::fullScreenChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(bool)>(&MyQVideoWindowControl::Signal_FullScreenChanged));
}

void QVideoWindowControl_DisconnectFullScreenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(bool)>(&QVideoWindowControl::fullScreenChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(bool)>(&MyQVideoWindowControl::Signal_FullScreenChanged));
}

void QVideoWindowControl_FullScreenChanged(void* ptr, char fullScreen)
{
	static_cast<QVideoWindowControl*>(ptr)->fullScreenChanged(fullScreen != 0);
}

int QVideoWindowControl_Hue(void* ptr)
{
	return static_cast<QVideoWindowControl*>(ptr)->hue();
}

void QVideoWindowControl_ConnectHueChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::hueChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_HueChanged));
}

void QVideoWindowControl_DisconnectHueChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::hueChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_HueChanged));
}

void QVideoWindowControl_HueChanged(void* ptr, int hue)
{
	static_cast<QVideoWindowControl*>(ptr)->hueChanged(hue);
}

char QVideoWindowControl_IsFullScreen(void* ptr)
{
	return static_cast<QVideoWindowControl*>(ptr)->isFullScreen();
}

void* QVideoWindowControl_NativeSize(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVideoWindowControl*>(ptr)->nativeSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QVideoWindowControl_ConnectNativeSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)()>(&QVideoWindowControl::nativeSizeChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)()>(&MyQVideoWindowControl::Signal_NativeSizeChanged));
}

void QVideoWindowControl_DisconnectNativeSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)()>(&QVideoWindowControl::nativeSizeChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)()>(&MyQVideoWindowControl::Signal_NativeSizeChanged));
}

void QVideoWindowControl_NativeSizeChanged(void* ptr)
{
	static_cast<QVideoWindowControl*>(ptr)->nativeSizeChanged();
}

void QVideoWindowControl_Repaint(void* ptr)
{
	static_cast<QVideoWindowControl*>(ptr)->repaint();
}

int QVideoWindowControl_Saturation(void* ptr)
{
	return static_cast<QVideoWindowControl*>(ptr)->saturation();
}

void QVideoWindowControl_ConnectSaturationChanged(void* ptr)
{
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::saturationChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_SaturationChanged));
}

void QVideoWindowControl_DisconnectSaturationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::saturationChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_SaturationChanged));
}

void QVideoWindowControl_SaturationChanged(void* ptr, int saturation)
{
	static_cast<QVideoWindowControl*>(ptr)->saturationChanged(saturation);
}

void QVideoWindowControl_SetAspectRatioMode(void* ptr, long long mode)
{
	static_cast<QVideoWindowControl*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QVideoWindowControl_SetBrightness(void* ptr, int brightness)
{
	static_cast<QVideoWindowControl*>(ptr)->setBrightness(brightness);
}

void QVideoWindowControl_SetContrast(void* ptr, int contrast)
{
	static_cast<QVideoWindowControl*>(ptr)->setContrast(contrast);
}

void QVideoWindowControl_SetDisplayRect(void* ptr, void* rect)
{
	static_cast<QVideoWindowControl*>(ptr)->setDisplayRect(*static_cast<QRect*>(rect));
}

void QVideoWindowControl_SetFullScreen(void* ptr, char fullScreen)
{
	static_cast<QVideoWindowControl*>(ptr)->setFullScreen(fullScreen != 0);
}

void QVideoWindowControl_SetHue(void* ptr, int hue)
{
	static_cast<QVideoWindowControl*>(ptr)->setHue(hue);
}

void QVideoWindowControl_SetSaturation(void* ptr, int saturation)
{
	static_cast<QVideoWindowControl*>(ptr)->setSaturation(saturation);
}

void QVideoWindowControl_SetWinId(void* ptr, uintptr_t id)
{
	static_cast<QVideoWindowControl*>(ptr)->setWinId(id);
}

uintptr_t QVideoWindowControl_WinId(void* ptr)
{
	return static_cast<QVideoWindowControl*>(ptr)->winId();
}

void QVideoWindowControl_DestroyQVideoWindowControl(void* ptr)
{
	static_cast<QVideoWindowControl*>(ptr)->~QVideoWindowControl();
}

void QVideoWindowControl_TimerEvent(void* ptr, void* event)
{
	static_cast<QVideoWindowControl*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWindowControl_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWindowControl*>(ptr)->QVideoWindowControl::timerEvent(static_cast<QTimerEvent*>(event));
}

void QVideoWindowControl_ChildEvent(void* ptr, void* event)
{
	static_cast<QVideoWindowControl*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWindowControl_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWindowControl*>(ptr)->QVideoWindowControl::childEvent(static_cast<QChildEvent*>(event));
}

void QVideoWindowControl_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWindowControl*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWindowControl_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWindowControl*>(ptr)->QVideoWindowControl::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWindowControl_CustomEvent(void* ptr, void* event)
{
	static_cast<QVideoWindowControl*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QVideoWindowControl_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVideoWindowControl*>(ptr)->QVideoWindowControl::customEvent(static_cast<QEvent*>(event));
}

void QVideoWindowControl_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QVideoWindowControl*>(ptr), "deleteLater");
}

void QVideoWindowControl_DeleteLaterDefault(void* ptr)
{
	static_cast<QVideoWindowControl*>(ptr)->QVideoWindowControl::deleteLater();
}

void QVideoWindowControl_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QVideoWindowControl*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVideoWindowControl_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVideoWindowControl*>(ptr)->QVideoWindowControl::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QVideoWindowControl_Event(void* ptr, void* e)
{
	return static_cast<QVideoWindowControl*>(ptr)->event(static_cast<QEvent*>(e));
}

char QVideoWindowControl_EventDefault(void* ptr, void* e)
{
	return static_cast<QVideoWindowControl*>(ptr)->QVideoWindowControl::event(static_cast<QEvent*>(e));
}

char QVideoWindowControl_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWindowControl*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QVideoWindowControl_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVideoWindowControl*>(ptr)->QVideoWindowControl::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QVideoWindowControl_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWindowControl*>(ptr)->metaObject());
}

void* QVideoWindowControl_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QVideoWindowControl*>(ptr)->QVideoWindowControl::metaObject());
}

