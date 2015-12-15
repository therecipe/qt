#include "multimedia.h"
#include "_cgo_export.h"

#include <QAbstractPlanarVideoBuffer>
#include <QAbstractVideoBuffer>
#include <QAbstractVideoFilter>
#include <QAbstractVideoSurface>
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
#include <QCameraViewfinderSettings>
#include <QCameraViewfinderSettingsControl>
#include <QCameraViewfinderSettingsControl2>
#include <QCameraZoomControl>
#include <QIODevice>
#include <QImage>
#include <QImageEncoderControl>
#include <QImageEncoderSettings>
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
#include <QMetaObject>
#include <QNetworkRequest>
#include <QObject>
#include <QPoint>
#include <QPointF>
#include <QRadioData>
#include <QRadioDataControl>
#include <QRadioTuner>
#include <QRadioTunerControl>
#include <QRect>
#include <QSize>
#include <QSound>
#include <QSoundEffect>
#include <QString>
#include <QUrl>
#include <QVariant>
#include <QVideoDeviceSelectorControl>
#include <QVideoEncoderSettings>
#include <QVideoEncoderSettingsControl>
#include <QVideoFrame>
#include <QVideoProbe>
#include <QVideoRendererControl>
#include <QVideoSurfaceFormat>
#include <QVideoWindowControl>

class MyQAbstractPlanarVideoBuffer: public QAbstractPlanarVideoBuffer {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

void QAbstractPlanarVideoBuffer_DestroyQAbstractPlanarVideoBuffer(void* ptr){
	static_cast<QAbstractPlanarVideoBuffer*>(ptr)->~QAbstractPlanarVideoBuffer();
}

char* QAbstractPlanarVideoBuffer_ObjectNameAbs(void* ptr){
	return static_cast<MyQAbstractPlanarVideoBuffer*>(ptr)->objectNameAbs().toUtf8().data();
}

void QAbstractPlanarVideoBuffer_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQAbstractPlanarVideoBuffer*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQAbstractVideoBuffer: public QAbstractVideoBuffer {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void release() { if (!callbackQAbstractVideoBufferRelease(this->objectNameAbs().toUtf8().data())) { QAbstractVideoBuffer::release(); }; };
protected:
};

void* QAbstractVideoBuffer_Handle(void* ptr){
	return new QVariant(static_cast<QAbstractVideoBuffer*>(ptr)->handle());
}

int QAbstractVideoBuffer_HandleType(void* ptr){
	return static_cast<QAbstractVideoBuffer*>(ptr)->handleType();
}

int QAbstractVideoBuffer_MapMode(void* ptr){
	return static_cast<QAbstractVideoBuffer*>(ptr)->mapMode();
}

void QAbstractVideoBuffer_Release(void* ptr){
	static_cast<QAbstractVideoBuffer*>(ptr)->release();
}

void QAbstractVideoBuffer_Unmap(void* ptr){
	static_cast<QAbstractVideoBuffer*>(ptr)->unmap();
}

void QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(void* ptr){
	static_cast<QAbstractVideoBuffer*>(ptr)->~QAbstractVideoBuffer();
}

char* QAbstractVideoBuffer_ObjectNameAbs(void* ptr){
	return static_cast<MyQAbstractVideoBuffer*>(ptr)->objectNameAbs().toUtf8().data();
}

void QAbstractVideoBuffer_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQAbstractVideoBuffer*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQAbstractVideoFilter: public QAbstractVideoFilter {
public:
	void Signal_ActiveChanged() { callbackQAbstractVideoFilterActiveChanged(this->objectName().toUtf8().data()); };
protected:
};

int QAbstractVideoFilter_IsActive(void* ptr){
	return static_cast<QAbstractVideoFilter*>(ptr)->isActive();
}

void QAbstractVideoFilter_SetActive(void* ptr, int v){
	static_cast<QAbstractVideoFilter*>(ptr)->setActive(v != 0);
}

void QAbstractVideoFilter_ConnectActiveChanged(void* ptr){
	QObject::connect(static_cast<QAbstractVideoFilter*>(ptr), static_cast<void (QAbstractVideoFilter::*)()>(&QAbstractVideoFilter::activeChanged), static_cast<MyQAbstractVideoFilter*>(ptr), static_cast<void (MyQAbstractVideoFilter::*)()>(&MyQAbstractVideoFilter::Signal_ActiveChanged));;
}

void QAbstractVideoFilter_DisconnectActiveChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractVideoFilter*>(ptr), static_cast<void (QAbstractVideoFilter::*)()>(&QAbstractVideoFilter::activeChanged), static_cast<MyQAbstractVideoFilter*>(ptr), static_cast<void (MyQAbstractVideoFilter::*)()>(&MyQAbstractVideoFilter::Signal_ActiveChanged));;
}

void* QAbstractVideoFilter_CreateFilterRunnable(void* ptr){
	return static_cast<QAbstractVideoFilter*>(ptr)->createFilterRunnable();
}

class MyQAbstractVideoSurface: public QAbstractVideoSurface {
public:
	void Signal_ActiveChanged(bool active) { callbackQAbstractVideoSurfaceActiveChanged(this->objectName().toUtf8().data(), active); };
	void stop() { if (!callbackQAbstractVideoSurfaceStop(this->objectName().toUtf8().data())) { QAbstractVideoSurface::stop(); }; };
	void Signal_SupportedFormatsChanged() { callbackQAbstractVideoSurfaceSupportedFormatsChanged(this->objectName().toUtf8().data()); };
protected:
};

void QAbstractVideoSurface_ConnectActiveChanged(void* ptr){
	QObject::connect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(bool)>(&QAbstractVideoSurface::activeChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(bool)>(&MyQAbstractVideoSurface::Signal_ActiveChanged));;
}

void QAbstractVideoSurface_DisconnectActiveChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(bool)>(&QAbstractVideoSurface::activeChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(bool)>(&MyQAbstractVideoSurface::Signal_ActiveChanged));;
}

int QAbstractVideoSurface_Error(void* ptr){
	return static_cast<QAbstractVideoSurface*>(ptr)->error();
}

int QAbstractVideoSurface_IsActive(void* ptr){
	return static_cast<QAbstractVideoSurface*>(ptr)->isActive();
}

int QAbstractVideoSurface_IsFormatSupported(void* ptr, void* format){
	return static_cast<QAbstractVideoSurface*>(ptr)->isFormatSupported(*static_cast<QVideoSurfaceFormat*>(format));
}

int QAbstractVideoSurface_Present(void* ptr, void* frame){
	return static_cast<QAbstractVideoSurface*>(ptr)->present(*static_cast<QVideoFrame*>(frame));
}

int QAbstractVideoSurface_Start(void* ptr, void* format){
	return static_cast<QAbstractVideoSurface*>(ptr)->start(*static_cast<QVideoSurfaceFormat*>(format));
}

void QAbstractVideoSurface_Stop(void* ptr){
	static_cast<QAbstractVideoSurface*>(ptr)->stop();
}

void QAbstractVideoSurface_ConnectSupportedFormatsChanged(void* ptr){
	QObject::connect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)()>(&QAbstractVideoSurface::supportedFormatsChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)()>(&MyQAbstractVideoSurface::Signal_SupportedFormatsChanged));;
}

void QAbstractVideoSurface_DisconnectSupportedFormatsChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)()>(&QAbstractVideoSurface::supportedFormatsChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)()>(&MyQAbstractVideoSurface::Signal_SupportedFormatsChanged));;
}

void QAbstractVideoSurface_DestroyQAbstractVideoSurface(void* ptr){
	static_cast<QAbstractVideoSurface*>(ptr)->~QAbstractVideoSurface();
}

void* QAudioBuffer_NewQAudioBuffer(){
	return new QAudioBuffer();
}

void* QAudioBuffer_NewQAudioBuffer3(void* other){
	return new QAudioBuffer(*static_cast<QAudioBuffer*>(other));
}

int QAudioBuffer_ByteCount(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->byteCount();
}

void* QAudioBuffer_ConstData(void* ptr){
	return const_cast<void*>(static_cast<QAudioBuffer*>(ptr)->constData());
}

void* QAudioBuffer_Data2(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->data();
}

void* QAudioBuffer_Data(void* ptr){
	return const_cast<void*>(static_cast<QAudioBuffer*>(ptr)->data());
}

int QAudioBuffer_FrameCount(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->frameCount();
}

int QAudioBuffer_IsValid(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->isValid();
}

int QAudioBuffer_SampleCount(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->sampleCount();
}

void QAudioBuffer_DestroyQAudioBuffer(void* ptr){
	static_cast<QAudioBuffer*>(ptr)->~QAudioBuffer();
}

class MyQAudioDecoder: public QAudioDecoder {
public:
	void Signal_BufferAvailableChanged(bool available) { callbackQAudioDecoderBufferAvailableChanged(this->objectName().toUtf8().data(), available); };
	void Signal_BufferReady() { callbackQAudioDecoderBufferReady(this->objectName().toUtf8().data()); };
	void Signal_Finished() { callbackQAudioDecoderFinished(this->objectName().toUtf8().data()); };
	void Signal_SourceChanged() { callbackQAudioDecoderSourceChanged(this->objectName().toUtf8().data()); };
	void Signal_StateChanged(QAudioDecoder::State state) { callbackQAudioDecoderStateChanged(this->objectName().toUtf8().data(), state); };
protected:
};

char* QAudioDecoder_ErrorString(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->errorString().toUtf8().data();
}

void* QAudioDecoder_NewQAudioDecoder(void* parent){
	return new QAudioDecoder(static_cast<QObject*>(parent));
}

int QAudioDecoder_BufferAvailable(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->bufferAvailable();
}

void QAudioDecoder_ConnectBufferAvailableChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(bool)>(&QAudioDecoder::bufferAvailableChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(bool)>(&MyQAudioDecoder::Signal_BufferAvailableChanged));;
}

void QAudioDecoder_DisconnectBufferAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(bool)>(&QAudioDecoder::bufferAvailableChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(bool)>(&MyQAudioDecoder::Signal_BufferAvailableChanged));;
}

void QAudioDecoder_ConnectBufferReady(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::bufferReady), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_BufferReady));;
}

void QAudioDecoder_DisconnectBufferReady(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::bufferReady), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_BufferReady));;
}

int QAudioDecoder_Error(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->error();
}

void QAudioDecoder_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::finished), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_Finished));;
}

void QAudioDecoder_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::finished), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_Finished));;
}

void QAudioDecoder_SetAudioFormat(void* ptr, void* format){
	static_cast<QAudioDecoder*>(ptr)->setAudioFormat(*static_cast<QAudioFormat*>(format));
}

void QAudioDecoder_SetSourceDevice(void* ptr, void* device){
	static_cast<QAudioDecoder*>(ptr)->setSourceDevice(static_cast<QIODevice*>(device));
}

void QAudioDecoder_SetSourceFilename(void* ptr, char* fileName){
	static_cast<QAudioDecoder*>(ptr)->setSourceFilename(QString(fileName));
}

void QAudioDecoder_ConnectSourceChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::sourceChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_SourceChanged));;
}

void QAudioDecoder_DisconnectSourceChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::sourceChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_SourceChanged));;
}

void* QAudioDecoder_SourceDevice(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->sourceDevice();
}

char* QAudioDecoder_SourceFilename(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->sourceFilename().toUtf8().data();
}

void QAudioDecoder_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAudioDecoder*>(ptr), "start");
}

void QAudioDecoder_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::State)>(&QAudioDecoder::stateChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::State)>(&MyQAudioDecoder::Signal_StateChanged));;
}

void QAudioDecoder_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::State)>(&QAudioDecoder::stateChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::State)>(&MyQAudioDecoder::Signal_StateChanged));;
}

void QAudioDecoder_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAudioDecoder*>(ptr), "stop");
}

void QAudioDecoder_DestroyQAudioDecoder(void* ptr){
	static_cast<QAudioDecoder*>(ptr)->~QAudioDecoder();
}

class MyQAudioDecoderControl: public QAudioDecoderControl {
public:
	void Signal_BufferAvailableChanged(bool available) { callbackQAudioDecoderControlBufferAvailableChanged(this->objectName().toUtf8().data(), available); };
	void Signal_BufferReady() { callbackQAudioDecoderControlBufferReady(this->objectName().toUtf8().data()); };
	void Signal_Error(int error, const QString & errorString) { callbackQAudioDecoderControlError(this->objectName().toUtf8().data(), error, errorString.toUtf8().data()); };
	void Signal_Finished() { callbackQAudioDecoderControlFinished(this->objectName().toUtf8().data()); };
	void Signal_SourceChanged() { callbackQAudioDecoderControlSourceChanged(this->objectName().toUtf8().data()); };
	void Signal_StateChanged(QAudioDecoder::State state) { callbackQAudioDecoderControlStateChanged(this->objectName().toUtf8().data(), state); };
protected:
};

int QAudioDecoderControl_BufferAvailable(void* ptr){
	return static_cast<QAudioDecoderControl*>(ptr)->bufferAvailable();
}

void QAudioDecoderControl_ConnectBufferAvailableChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(bool)>(&QAudioDecoderControl::bufferAvailableChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(bool)>(&MyQAudioDecoderControl::Signal_BufferAvailableChanged));;
}

void QAudioDecoderControl_DisconnectBufferAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(bool)>(&QAudioDecoderControl::bufferAvailableChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(bool)>(&MyQAudioDecoderControl::Signal_BufferAvailableChanged));;
}

void QAudioDecoderControl_ConnectBufferReady(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::bufferReady), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_BufferReady));;
}

void QAudioDecoderControl_DisconnectBufferReady(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::bufferReady), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_BufferReady));;
}

void QAudioDecoderControl_ConnectError(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(int, const QString &)>(&QAudioDecoderControl::error), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(int, const QString &)>(&MyQAudioDecoderControl::Signal_Error));;
}

void QAudioDecoderControl_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(int, const QString &)>(&QAudioDecoderControl::error), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(int, const QString &)>(&MyQAudioDecoderControl::Signal_Error));;
}

void QAudioDecoderControl_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::finished), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_Finished));;
}

void QAudioDecoderControl_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::finished), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_Finished));;
}

void QAudioDecoderControl_SetAudioFormat(void* ptr, void* format){
	static_cast<QAudioDecoderControl*>(ptr)->setAudioFormat(*static_cast<QAudioFormat*>(format));
}

void QAudioDecoderControl_SetSourceDevice(void* ptr, void* device){
	static_cast<QAudioDecoderControl*>(ptr)->setSourceDevice(static_cast<QIODevice*>(device));
}

void QAudioDecoderControl_SetSourceFilename(void* ptr, char* fileName){
	static_cast<QAudioDecoderControl*>(ptr)->setSourceFilename(QString(fileName));
}

void QAudioDecoderControl_ConnectSourceChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::sourceChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_SourceChanged));;
}

void QAudioDecoderControl_DisconnectSourceChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::sourceChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_SourceChanged));;
}

void* QAudioDecoderControl_SourceDevice(void* ptr){
	return static_cast<QAudioDecoderControl*>(ptr)->sourceDevice();
}

char* QAudioDecoderControl_SourceFilename(void* ptr){
	return static_cast<QAudioDecoderControl*>(ptr)->sourceFilename().toUtf8().data();
}

void QAudioDecoderControl_Start(void* ptr){
	static_cast<QAudioDecoderControl*>(ptr)->start();
}

void QAudioDecoderControl_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(QAudioDecoder::State)>(&QAudioDecoderControl::stateChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(QAudioDecoder::State)>(&MyQAudioDecoderControl::Signal_StateChanged));;
}

void QAudioDecoderControl_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(QAudioDecoder::State)>(&QAudioDecoderControl::stateChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(QAudioDecoder::State)>(&MyQAudioDecoderControl::Signal_StateChanged));;
}

void QAudioDecoderControl_Stop(void* ptr){
	static_cast<QAudioDecoderControl*>(ptr)->stop();
}

void QAudioDecoderControl_DestroyQAudioDecoderControl(void* ptr){
	static_cast<QAudioDecoderControl*>(ptr)->~QAudioDecoderControl();
}

void* QAudioDeviceInfo_NewQAudioDeviceInfo(){
	return new QAudioDeviceInfo();
}

void* QAudioDeviceInfo_NewQAudioDeviceInfo2(void* other){
	return new QAudioDeviceInfo(*static_cast<QAudioDeviceInfo*>(other));
}

char* QAudioDeviceInfo_DeviceName(void* ptr){
	return static_cast<QAudioDeviceInfo*>(ptr)->deviceName().toUtf8().data();
}

int QAudioDeviceInfo_IsFormatSupported(void* ptr, void* settings){
	return static_cast<QAudioDeviceInfo*>(ptr)->isFormatSupported(*static_cast<QAudioFormat*>(settings));
}

int QAudioDeviceInfo_IsNull(void* ptr){
	return static_cast<QAudioDeviceInfo*>(ptr)->isNull();
}

char* QAudioDeviceInfo_SupportedCodecs(void* ptr){
	return static_cast<QAudioDeviceInfo*>(ptr)->supportedCodecs().join(",,,").toUtf8().data();
}

void QAudioDeviceInfo_DestroyQAudioDeviceInfo(void* ptr){
	static_cast<QAudioDeviceInfo*>(ptr)->~QAudioDeviceInfo();
}

void* QAudioEncoderSettings_NewQAudioEncoderSettings(){
	return new QAudioEncoderSettings();
}

void* QAudioEncoderSettings_NewQAudioEncoderSettings2(void* other){
	return new QAudioEncoderSettings(*static_cast<QAudioEncoderSettings*>(other));
}

int QAudioEncoderSettings_BitRate(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->bitRate();
}

int QAudioEncoderSettings_ChannelCount(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->channelCount();
}

char* QAudioEncoderSettings_Codec(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->codec().toUtf8().data();
}

void* QAudioEncoderSettings_EncodingOption(void* ptr, char* option){
	return new QVariant(static_cast<QAudioEncoderSettings*>(ptr)->encodingOption(QString(option)));
}

int QAudioEncoderSettings_IsNull(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->isNull();
}

int QAudioEncoderSettings_SampleRate(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->sampleRate();
}

void QAudioEncoderSettings_SetBitRate(void* ptr, int rate){
	static_cast<QAudioEncoderSettings*>(ptr)->setBitRate(rate);
}

void QAudioEncoderSettings_SetChannelCount(void* ptr, int channels){
	static_cast<QAudioEncoderSettings*>(ptr)->setChannelCount(channels);
}

void QAudioEncoderSettings_SetCodec(void* ptr, char* codec){
	static_cast<QAudioEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QAudioEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value){
	static_cast<QAudioEncoderSettings*>(ptr)->setEncodingOption(QString(option), *static_cast<QVariant*>(value));
}

void QAudioEncoderSettings_SetSampleRate(void* ptr, int rate){
	static_cast<QAudioEncoderSettings*>(ptr)->setSampleRate(rate);
}

void QAudioEncoderSettings_DestroyQAudioEncoderSettings(void* ptr){
	static_cast<QAudioEncoderSettings*>(ptr)->~QAudioEncoderSettings();
}

class MyQAudioEncoderSettingsControl: public QAudioEncoderSettingsControl {
public:
protected:
};

char* QAudioEncoderSettingsControl_CodecDescription(void* ptr, char* codec){
	return static_cast<QAudioEncoderSettingsControl*>(ptr)->codecDescription(QString(codec)).toUtf8().data();
}

void QAudioEncoderSettingsControl_SetAudioSettings(void* ptr, void* settings){
	static_cast<QAudioEncoderSettingsControl*>(ptr)->setAudioSettings(*static_cast<QAudioEncoderSettings*>(settings));
}

char* QAudioEncoderSettingsControl_SupportedAudioCodecs(void* ptr){
	return static_cast<QAudioEncoderSettingsControl*>(ptr)->supportedAudioCodecs().join(",,,").toUtf8().data();
}

void QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(void* ptr){
	static_cast<QAudioEncoderSettingsControl*>(ptr)->~QAudioEncoderSettingsControl();
}

void* QAudioFormat_NewQAudioFormat(){
	return new QAudioFormat();
}

void* QAudioFormat_NewQAudioFormat2(void* other){
	return new QAudioFormat(*static_cast<QAudioFormat*>(other));
}

int QAudioFormat_ByteOrder(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->byteOrder();
}

int QAudioFormat_BytesPerFrame(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->bytesPerFrame();
}

int QAudioFormat_ChannelCount(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->channelCount();
}

char* QAudioFormat_Codec(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->codec().toUtf8().data();
}

int QAudioFormat_IsValid(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->isValid();
}

int QAudioFormat_SampleRate(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->sampleRate();
}

int QAudioFormat_SampleSize(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->sampleSize();
}

int QAudioFormat_SampleType(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->sampleType();
}

void QAudioFormat_SetByteOrder(void* ptr, int byteOrder){
	static_cast<QAudioFormat*>(ptr)->setByteOrder(static_cast<QAudioFormat::Endian>(byteOrder));
}

void QAudioFormat_SetChannelCount(void* ptr, int channels){
	static_cast<QAudioFormat*>(ptr)->setChannelCount(channels);
}

void QAudioFormat_SetCodec(void* ptr, char* codec){
	static_cast<QAudioFormat*>(ptr)->setCodec(QString(codec));
}

void QAudioFormat_SetSampleRate(void* ptr, int samplerate){
	static_cast<QAudioFormat*>(ptr)->setSampleRate(samplerate);
}

void QAudioFormat_SetSampleSize(void* ptr, int sampleSize){
	static_cast<QAudioFormat*>(ptr)->setSampleSize(sampleSize);
}

void QAudioFormat_SetSampleType(void* ptr, int sampleType){
	static_cast<QAudioFormat*>(ptr)->setSampleType(static_cast<QAudioFormat::SampleType>(sampleType));
}

void QAudioFormat_DestroyQAudioFormat(void* ptr){
	static_cast<QAudioFormat*>(ptr)->~QAudioFormat();
}

class MyQAudioInput: public QAudioInput {
public:
	void Signal_Notify() { callbackQAudioInputNotify(this->objectName().toUtf8().data()); };
protected:
};

void* QAudioInput_NewQAudioInput2(void* audioDevice, void* format, void* parent){
	return new QAudioInput(*static_cast<QAudioDeviceInfo*>(audioDevice), *static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

void* QAudioInput_NewQAudioInput(void* format, void* parent){
	return new QAudioInput(*static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

int QAudioInput_BufferSize(void* ptr){
	return static_cast<QAudioInput*>(ptr)->bufferSize();
}

int QAudioInput_BytesReady(void* ptr){
	return static_cast<QAudioInput*>(ptr)->bytesReady();
}

void QAudioInput_ConnectNotify(void* ptr){
	QObject::connect(static_cast<QAudioInput*>(ptr), static_cast<void (QAudioInput::*)()>(&QAudioInput::notify), static_cast<MyQAudioInput*>(ptr), static_cast<void (MyQAudioInput::*)()>(&MyQAudioInput::Signal_Notify));;
}

void QAudioInput_DisconnectNotify(void* ptr){
	QObject::disconnect(static_cast<QAudioInput*>(ptr), static_cast<void (QAudioInput::*)()>(&QAudioInput::notify), static_cast<MyQAudioInput*>(ptr), static_cast<void (MyQAudioInput::*)()>(&MyQAudioInput::Signal_Notify));;
}

int QAudioInput_NotifyInterval(void* ptr){
	return static_cast<QAudioInput*>(ptr)->notifyInterval();
}

int QAudioInput_PeriodSize(void* ptr){
	return static_cast<QAudioInput*>(ptr)->periodSize();
}

void QAudioInput_Reset(void* ptr){
	static_cast<QAudioInput*>(ptr)->reset();
}

void QAudioInput_Resume(void* ptr){
	static_cast<QAudioInput*>(ptr)->resume();
}

void QAudioInput_SetBufferSize(void* ptr, int value){
	static_cast<QAudioInput*>(ptr)->setBufferSize(value);
}

void QAudioInput_SetNotifyInterval(void* ptr, int ms){
	static_cast<QAudioInput*>(ptr)->setNotifyInterval(ms);
}

void QAudioInput_SetVolume(void* ptr, double volume){
	static_cast<QAudioInput*>(ptr)->setVolume(static_cast<qreal>(volume));
}

void* QAudioInput_Start2(void* ptr){
	return static_cast<QAudioInput*>(ptr)->start();
}

void QAudioInput_Start(void* ptr, void* device){
	static_cast<QAudioInput*>(ptr)->start(static_cast<QIODevice*>(device));
}

void QAudioInput_Stop(void* ptr){
	static_cast<QAudioInput*>(ptr)->stop();
}

void QAudioInput_Suspend(void* ptr){
	static_cast<QAudioInput*>(ptr)->suspend();
}

double QAudioInput_Volume(void* ptr){
	return static_cast<double>(static_cast<QAudioInput*>(ptr)->volume());
}

void QAudioInput_DestroyQAudioInput(void* ptr){
	static_cast<QAudioInput*>(ptr)->~QAudioInput();
}

class MyQAudioInputSelectorControl: public QAudioInputSelectorControl {
public:
	void Signal_ActiveInputChanged(const QString & name) { callbackQAudioInputSelectorControlActiveInputChanged(this->objectName().toUtf8().data(), name.toUtf8().data()); };
	void Signal_AvailableInputsChanged() { callbackQAudioInputSelectorControlAvailableInputsChanged(this->objectName().toUtf8().data()); };
protected:
};

char* QAudioInputSelectorControl_ActiveInput(void* ptr){
	return static_cast<QAudioInputSelectorControl*>(ptr)->activeInput().toUtf8().data();
}

void QAudioInputSelectorControl_ConnectActiveInputChanged(void* ptr){
	QObject::connect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)(const QString &)>(&QAudioInputSelectorControl::activeInputChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)(const QString &)>(&MyQAudioInputSelectorControl::Signal_ActiveInputChanged));;
}

void QAudioInputSelectorControl_DisconnectActiveInputChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)(const QString &)>(&QAudioInputSelectorControl::activeInputChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)(const QString &)>(&MyQAudioInputSelectorControl::Signal_ActiveInputChanged));;
}

void QAudioInputSelectorControl_ConnectAvailableInputsChanged(void* ptr){
	QObject::connect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)()>(&QAudioInputSelectorControl::availableInputsChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)()>(&MyQAudioInputSelectorControl::Signal_AvailableInputsChanged));;
}

void QAudioInputSelectorControl_DisconnectAvailableInputsChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)()>(&QAudioInputSelectorControl::availableInputsChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)()>(&MyQAudioInputSelectorControl::Signal_AvailableInputsChanged));;
}

char* QAudioInputSelectorControl_DefaultInput(void* ptr){
	return static_cast<QAudioInputSelectorControl*>(ptr)->defaultInput().toUtf8().data();
}

char* QAudioInputSelectorControl_InputDescription(void* ptr, char* name){
	return static_cast<QAudioInputSelectorControl*>(ptr)->inputDescription(QString(name)).toUtf8().data();
}

void QAudioInputSelectorControl_SetActiveInput(void* ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QAudioInputSelectorControl*>(ptr), "setActiveInput", Q_ARG(QString, QString(name)));
}

void QAudioInputSelectorControl_DestroyQAudioInputSelectorControl(void* ptr){
	static_cast<QAudioInputSelectorControl*>(ptr)->~QAudioInputSelectorControl();
}

class MyQAudioOutput: public QAudioOutput {
public:
	void Signal_Notify() { callbackQAudioOutputNotify(this->objectName().toUtf8().data()); };
protected:
};

void* QAudioOutput_NewQAudioOutput2(void* audioDevice, void* format, void* parent){
	return new QAudioOutput(*static_cast<QAudioDeviceInfo*>(audioDevice), *static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

void* QAudioOutput_NewQAudioOutput(void* format, void* parent){
	return new QAudioOutput(*static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

int QAudioOutput_BufferSize(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->bufferSize();
}

int QAudioOutput_BytesFree(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->bytesFree();
}

char* QAudioOutput_Category(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->category().toUtf8().data();
}

void QAudioOutput_ConnectNotify(void* ptr){
	QObject::connect(static_cast<QAudioOutput*>(ptr), static_cast<void (QAudioOutput::*)()>(&QAudioOutput::notify), static_cast<MyQAudioOutput*>(ptr), static_cast<void (MyQAudioOutput::*)()>(&MyQAudioOutput::Signal_Notify));;
}

void QAudioOutput_DisconnectNotify(void* ptr){
	QObject::disconnect(static_cast<QAudioOutput*>(ptr), static_cast<void (QAudioOutput::*)()>(&QAudioOutput::notify), static_cast<MyQAudioOutput*>(ptr), static_cast<void (MyQAudioOutput::*)()>(&MyQAudioOutput::Signal_Notify));;
}

int QAudioOutput_NotifyInterval(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->notifyInterval();
}

int QAudioOutput_PeriodSize(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->periodSize();
}

void QAudioOutput_Reset(void* ptr){
	static_cast<QAudioOutput*>(ptr)->reset();
}

void QAudioOutput_Resume(void* ptr){
	static_cast<QAudioOutput*>(ptr)->resume();
}

void QAudioOutput_SetBufferSize(void* ptr, int value){
	static_cast<QAudioOutput*>(ptr)->setBufferSize(value);
}

void QAudioOutput_SetCategory(void* ptr, char* category){
	static_cast<QAudioOutput*>(ptr)->setCategory(QString(category));
}

void QAudioOutput_SetNotifyInterval(void* ptr, int ms){
	static_cast<QAudioOutput*>(ptr)->setNotifyInterval(ms);
}

void QAudioOutput_SetVolume(void* ptr, double volume){
	static_cast<QAudioOutput*>(ptr)->setVolume(static_cast<qreal>(volume));
}

void* QAudioOutput_Start2(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->start();
}

void QAudioOutput_Start(void* ptr, void* device){
	static_cast<QAudioOutput*>(ptr)->start(static_cast<QIODevice*>(device));
}

void QAudioOutput_Stop(void* ptr){
	static_cast<QAudioOutput*>(ptr)->stop();
}

void QAudioOutput_Suspend(void* ptr){
	static_cast<QAudioOutput*>(ptr)->suspend();
}

double QAudioOutput_Volume(void* ptr){
	return static_cast<double>(static_cast<QAudioOutput*>(ptr)->volume());
}

void QAudioOutput_DestroyQAudioOutput(void* ptr){
	static_cast<QAudioOutput*>(ptr)->~QAudioOutput();
}

class MyQAudioOutputSelectorControl: public QAudioOutputSelectorControl {
public:
	void Signal_ActiveOutputChanged(const QString & name) { callbackQAudioOutputSelectorControlActiveOutputChanged(this->objectName().toUtf8().data(), name.toUtf8().data()); };
	void Signal_AvailableOutputsChanged() { callbackQAudioOutputSelectorControlAvailableOutputsChanged(this->objectName().toUtf8().data()); };
protected:
};

char* QAudioOutputSelectorControl_ActiveOutput(void* ptr){
	return static_cast<QAudioOutputSelectorControl*>(ptr)->activeOutput().toUtf8().data();
}

void QAudioOutputSelectorControl_ConnectActiveOutputChanged(void* ptr){
	QObject::connect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)(const QString &)>(&QAudioOutputSelectorControl::activeOutputChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)(const QString &)>(&MyQAudioOutputSelectorControl::Signal_ActiveOutputChanged));;
}

void QAudioOutputSelectorControl_DisconnectActiveOutputChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)(const QString &)>(&QAudioOutputSelectorControl::activeOutputChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)(const QString &)>(&MyQAudioOutputSelectorControl::Signal_ActiveOutputChanged));;
}

void QAudioOutputSelectorControl_ConnectAvailableOutputsChanged(void* ptr){
	QObject::connect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)()>(&QAudioOutputSelectorControl::availableOutputsChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)()>(&MyQAudioOutputSelectorControl::Signal_AvailableOutputsChanged));;
}

void QAudioOutputSelectorControl_DisconnectAvailableOutputsChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)()>(&QAudioOutputSelectorControl::availableOutputsChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)()>(&MyQAudioOutputSelectorControl::Signal_AvailableOutputsChanged));;
}

char* QAudioOutputSelectorControl_DefaultOutput(void* ptr){
	return static_cast<QAudioOutputSelectorControl*>(ptr)->defaultOutput().toUtf8().data();
}

char* QAudioOutputSelectorControl_OutputDescription(void* ptr, char* name){
	return static_cast<QAudioOutputSelectorControl*>(ptr)->outputDescription(QString(name)).toUtf8().data();
}

void QAudioOutputSelectorControl_SetActiveOutput(void* ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QAudioOutputSelectorControl*>(ptr), "setActiveOutput", Q_ARG(QString, QString(name)));
}

void QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(void* ptr){
	static_cast<QAudioOutputSelectorControl*>(ptr)->~QAudioOutputSelectorControl();
}

class MyQAudioProbe: public QAudioProbe {
public:
	void Signal_Flush() { callbackQAudioProbeFlush(this->objectName().toUtf8().data()); };
protected:
};

void* QAudioProbe_NewQAudioProbe(void* parent){
	return new QAudioProbe(static_cast<QObject*>(parent));
}

void QAudioProbe_ConnectFlush(void* ptr){
	QObject::connect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)()>(&QAudioProbe::flush), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)()>(&MyQAudioProbe::Signal_Flush));;
}

void QAudioProbe_DisconnectFlush(void* ptr){
	QObject::disconnect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)()>(&QAudioProbe::flush), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)()>(&MyQAudioProbe::Signal_Flush));;
}

int QAudioProbe_IsActive(void* ptr){
	return static_cast<QAudioProbe*>(ptr)->isActive();
}

int QAudioProbe_SetSource(void* ptr, void* source){
	return static_cast<QAudioProbe*>(ptr)->setSource(static_cast<QMediaObject*>(source));
}

int QAudioProbe_SetSource2(void* ptr, void* mediaRecorder){
	return static_cast<QAudioProbe*>(ptr)->setSource(static_cast<QMediaRecorder*>(mediaRecorder));
}

void QAudioProbe_DestroyQAudioProbe(void* ptr){
	static_cast<QAudioProbe*>(ptr)->~QAudioProbe();
}

class MyQAudioRecorder: public QAudioRecorder {
public:
	void Signal_AudioInputChanged(const QString & name) { callbackQAudioRecorderAudioInputChanged(this->objectName().toUtf8().data(), name.toUtf8().data()); };
	void Signal_AvailableAudioInputsChanged() { callbackQAudioRecorderAvailableAudioInputsChanged(this->objectName().toUtf8().data()); };
protected:
};

void* QAudioRecorder_NewQAudioRecorder(void* parent){
	return new QAudioRecorder(static_cast<QObject*>(parent));
}

char* QAudioRecorder_AudioInput(void* ptr){
	return static_cast<QAudioRecorder*>(ptr)->audioInput().toUtf8().data();
}

void QAudioRecorder_ConnectAudioInputChanged(void* ptr){
	QObject::connect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)(const QString &)>(&QAudioRecorder::audioInputChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)(const QString &)>(&MyQAudioRecorder::Signal_AudioInputChanged));;
}

void QAudioRecorder_DisconnectAudioInputChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)(const QString &)>(&QAudioRecorder::audioInputChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)(const QString &)>(&MyQAudioRecorder::Signal_AudioInputChanged));;
}

char* QAudioRecorder_AudioInputDescription(void* ptr, char* name){
	return static_cast<QAudioRecorder*>(ptr)->audioInputDescription(QString(name)).toUtf8().data();
}

char* QAudioRecorder_AudioInputs(void* ptr){
	return static_cast<QAudioRecorder*>(ptr)->audioInputs().join(",,,").toUtf8().data();
}

void QAudioRecorder_ConnectAvailableAudioInputsChanged(void* ptr){
	QObject::connect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)()>(&QAudioRecorder::availableAudioInputsChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)()>(&MyQAudioRecorder::Signal_AvailableAudioInputsChanged));;
}

void QAudioRecorder_DisconnectAvailableAudioInputsChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)()>(&QAudioRecorder::availableAudioInputsChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)()>(&MyQAudioRecorder::Signal_AvailableAudioInputsChanged));;
}

char* QAudioRecorder_DefaultAudioInput(void* ptr){
	return static_cast<QAudioRecorder*>(ptr)->defaultAudioInput().toUtf8().data();
}

void QAudioRecorder_SetAudioInput(void* ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "setAudioInput", Q_ARG(QString, QString(name)));
}

void QAudioRecorder_DestroyQAudioRecorder(void* ptr){
	static_cast<QAudioRecorder*>(ptr)->~QAudioRecorder();
}

class MyQCamera: public QCamera {
public:
	MyQCamera(QCamera::Position position, QObject *parent) : QCamera(position, parent) {};
	MyQCamera(QObject *parent) : QCamera(parent) {};
	MyQCamera(const QByteArray &deviceName, QObject *parent) : QCamera(deviceName, parent) {};
	MyQCamera(const QCameraInfo &cameraInfo, QObject *parent) : QCamera(cameraInfo, parent) {};
	void Signal_CaptureModeChanged(QCamera::CaptureModes mode) { callbackQCameraCaptureModeChanged(this->objectName().toUtf8().data(), mode); };
	void Signal_LockFailed() { callbackQCameraLockFailed(this->objectName().toUtf8().data()); };
	void Signal_LockStatusChanged(QCamera::LockStatus status, QCamera::LockChangeReason reason) { callbackQCameraLockStatusChanged(this->objectName().toUtf8().data(), status, reason); };
	void Signal_Locked() { callbackQCameraLocked(this->objectName().toUtf8().data()); };
	void Signal_StateChanged(QCamera::State state) { callbackQCameraStateChanged(this->objectName().toUtf8().data(), state); };
	void Signal_StatusChanged(QCamera::Status status) { callbackQCameraStatusChanged(this->objectName().toUtf8().data(), status); };
protected:
};

int QCamera_CaptureMode(void* ptr){
	return static_cast<QCamera*>(ptr)->captureMode();
}

void QCamera_SearchAndLock2(void* ptr, int locks){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "searchAndLock", Q_ARG(QCamera::LockType, static_cast<QCamera::LockType>(locks)));
}

void QCamera_SetCaptureMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "setCaptureMode", Q_ARG(QCamera::CaptureMode, static_cast<QCamera::CaptureMode>(mode)));
}

int QCamera_Status(void* ptr){
	return static_cast<QCamera*>(ptr)->status();
}

void* QCamera_NewQCamera4(int position, void* parent){
	return new MyQCamera(static_cast<QCamera::Position>(position), static_cast<QObject*>(parent));
}

void* QCamera_NewQCamera(void* parent){
	return new MyQCamera(static_cast<QObject*>(parent));
}

void* QCamera_NewQCamera2(void* deviceName, void* parent){
	return new MyQCamera(*static_cast<QByteArray*>(deviceName), static_cast<QObject*>(parent));
}

void* QCamera_NewQCamera3(void* cameraInfo, void* parent){
	return new MyQCamera(*static_cast<QCameraInfo*>(cameraInfo), static_cast<QObject*>(parent));
}

void QCamera_ConnectCaptureModeChanged(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::CaptureModes)>(&QCamera::captureModeChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::CaptureModes)>(&MyQCamera::Signal_CaptureModeChanged));;
}

void QCamera_DisconnectCaptureModeChanged(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::CaptureModes)>(&QCamera::captureModeChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::CaptureModes)>(&MyQCamera::Signal_CaptureModeChanged));;
}

int QCamera_Error(void* ptr){
	return static_cast<QCamera*>(ptr)->error();
}

char* QCamera_ErrorString(void* ptr){
	return static_cast<QCamera*>(ptr)->errorString().toUtf8().data();
}

void* QCamera_Exposure(void* ptr){
	return static_cast<QCamera*>(ptr)->exposure();
}

void* QCamera_Focus(void* ptr){
	return static_cast<QCamera*>(ptr)->focus();
}

void* QCamera_ImageProcessing(void* ptr){
	return static_cast<QCamera*>(ptr)->imageProcessing();
}

int QCamera_IsCaptureModeSupported(void* ptr, int mode){
	return static_cast<QCamera*>(ptr)->isCaptureModeSupported(static_cast<QCamera::CaptureMode>(mode));
}

void QCamera_Load(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "load");
}

void QCamera_ConnectLockFailed(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::lockFailed), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_LockFailed));;
}

void QCamera_DisconnectLockFailed(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::lockFailed), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_LockFailed));;
}

int QCamera_LockStatus(void* ptr){
	return static_cast<QCamera*>(ptr)->lockStatus();
}

int QCamera_LockStatus2(void* ptr, int lockType){
	return static_cast<QCamera*>(ptr)->lockStatus(static_cast<QCamera::LockType>(lockType));
}

void QCamera_ConnectLockStatusChanged(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged));;
}

void QCamera_DisconnectLockStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&QCamera::lockStatusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCamera::Signal_LockStatusChanged));;
}

void QCamera_ConnectLocked(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::locked), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_Locked));;
}

void QCamera_DisconnectLocked(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)()>(&QCamera::locked), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)()>(&MyQCamera::Signal_Locked));;
}

int QCamera_RequestedLocks(void* ptr){
	return static_cast<QCamera*>(ptr)->requestedLocks();
}

void QCamera_SearchAndLock(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "searchAndLock");
}

void QCamera_SetViewfinder3(void* ptr, void* surface){
	static_cast<QCamera*>(ptr)->setViewfinder(static_cast<QAbstractVideoSurface*>(surface));
}

void QCamera_SetViewfinderSettings(void* ptr, void* settings){
	static_cast<QCamera*>(ptr)->setViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(settings));
}

void QCamera_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "start");
}

void QCamera_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::State)>(&QCamera::stateChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::State)>(&MyQCamera::Signal_StateChanged));;
}

void QCamera_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::State)>(&QCamera::stateChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::State)>(&MyQCamera::Signal_StateChanged));;
}

void QCamera_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Status)>(&QCamera::statusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Status)>(&MyQCamera::Signal_StatusChanged));;
}

void QCamera_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QCamera*>(ptr), static_cast<void (QCamera::*)(QCamera::Status)>(&QCamera::statusChanged), static_cast<MyQCamera*>(ptr), static_cast<void (MyQCamera::*)(QCamera::Status)>(&MyQCamera::Signal_StatusChanged));;
}

void QCamera_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "stop");
}

int QCamera_SupportedLocks(void* ptr){
	return static_cast<QCamera*>(ptr)->supportedLocks();
}

void QCamera_Unload(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unload");
}

void QCamera_Unlock(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unlock");
}

void QCamera_Unlock2(void* ptr, int locks){
	QMetaObject::invokeMethod(static_cast<QCamera*>(ptr), "unlock", Q_ARG(QCamera::LockType, static_cast<QCamera::LockType>(locks)));
}

void QCamera_DestroyQCamera(void* ptr){
	static_cast<QCamera*>(ptr)->~QCamera();
}

class MyQCameraCaptureBufferFormatControl: public QCameraCaptureBufferFormatControl {
public:
	void Signal_BufferFormatChanged(QVideoFrame::PixelFormat format) { callbackQCameraCaptureBufferFormatControlBufferFormatChanged(this->objectName().toUtf8().data(), format); };
protected:
};

int QCameraCaptureBufferFormatControl_BufferFormat(void* ptr){
	return static_cast<QCameraCaptureBufferFormatControl*>(ptr)->bufferFormat();
}

void QCameraCaptureBufferFormatControl_ConnectBufferFormatChanged(void* ptr){
	QObject::connect(static_cast<QCameraCaptureBufferFormatControl*>(ptr), static_cast<void (QCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&QCameraCaptureBufferFormatControl::bufferFormatChanged), static_cast<MyQCameraCaptureBufferFormatControl*>(ptr), static_cast<void (MyQCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&MyQCameraCaptureBufferFormatControl::Signal_BufferFormatChanged));;
}

void QCameraCaptureBufferFormatControl_DisconnectBufferFormatChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraCaptureBufferFormatControl*>(ptr), static_cast<void (QCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&QCameraCaptureBufferFormatControl::bufferFormatChanged), static_cast<MyQCameraCaptureBufferFormatControl*>(ptr), static_cast<void (MyQCameraCaptureBufferFormatControl::*)(QVideoFrame::PixelFormat)>(&MyQCameraCaptureBufferFormatControl::Signal_BufferFormatChanged));;
}

void QCameraCaptureBufferFormatControl_SetBufferFormat(void* ptr, int format){
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->setBufferFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraCaptureBufferFormatControl_DestroyQCameraCaptureBufferFormatControl(void* ptr){
	static_cast<QCameraCaptureBufferFormatControl*>(ptr)->~QCameraCaptureBufferFormatControl();
}

class MyQCameraCaptureDestinationControl: public QCameraCaptureDestinationControl {
public:
	void Signal_CaptureDestinationChanged(QCameraImageCapture::CaptureDestinations destination) { callbackQCameraCaptureDestinationControlCaptureDestinationChanged(this->objectName().toUtf8().data(), destination); };
protected:
};

int QCameraCaptureDestinationControl_CaptureDestination(void* ptr){
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->captureDestination();
}

void QCameraCaptureDestinationControl_ConnectCaptureDestinationChanged(void* ptr){
	QObject::connect(static_cast<QCameraCaptureDestinationControl*>(ptr), static_cast<void (QCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraCaptureDestinationControl::captureDestinationChanged), static_cast<MyQCameraCaptureDestinationControl*>(ptr), static_cast<void (MyQCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraCaptureDestinationControl::Signal_CaptureDestinationChanged));;
}

void QCameraCaptureDestinationControl_DisconnectCaptureDestinationChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraCaptureDestinationControl*>(ptr), static_cast<void (QCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraCaptureDestinationControl::captureDestinationChanged), static_cast<MyQCameraCaptureDestinationControl*>(ptr), static_cast<void (MyQCameraCaptureDestinationControl::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraCaptureDestinationControl::Signal_CaptureDestinationChanged));;
}

int QCameraCaptureDestinationControl_IsCaptureDestinationSupported(void* ptr, int destination){
	return static_cast<QCameraCaptureDestinationControl*>(ptr)->isCaptureDestinationSupported(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraCaptureDestinationControl_SetCaptureDestination(void* ptr, int destination){
	static_cast<QCameraCaptureDestinationControl*>(ptr)->setCaptureDestination(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraCaptureDestinationControl_DestroyQCameraCaptureDestinationControl(void* ptr){
	static_cast<QCameraCaptureDestinationControl*>(ptr)->~QCameraCaptureDestinationControl();
}

class MyQCameraControl: public QCameraControl {
public:
	void Signal_CaptureModeChanged(QCamera::CaptureModes mode) { callbackQCameraControlCaptureModeChanged(this->objectName().toUtf8().data(), mode); };
	void Signal_Error(int error, const QString & errorString) { callbackQCameraControlError(this->objectName().toUtf8().data(), error, errorString.toUtf8().data()); };
	void Signal_StateChanged(QCamera::State state) { callbackQCameraControlStateChanged(this->objectName().toUtf8().data(), state); };
	void Signal_StatusChanged(QCamera::Status status) { callbackQCameraControlStatusChanged(this->objectName().toUtf8().data(), status); };
protected:
};

int QCameraControl_CanChangeProperty(void* ptr, int changeType, int status){
	return static_cast<QCameraControl*>(ptr)->canChangeProperty(static_cast<QCameraControl::PropertyChangeType>(changeType), static_cast<QCamera::Status>(status));
}

int QCameraControl_CaptureMode(void* ptr){
	return static_cast<QCameraControl*>(ptr)->captureMode();
}

void QCameraControl_ConnectCaptureModeChanged(void* ptr){
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::CaptureModes)>(&QCameraControl::captureModeChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::CaptureModes)>(&MyQCameraControl::Signal_CaptureModeChanged));;
}

void QCameraControl_DisconnectCaptureModeChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::CaptureModes)>(&QCameraControl::captureModeChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::CaptureModes)>(&MyQCameraControl::Signal_CaptureModeChanged));;
}

void QCameraControl_ConnectError(void* ptr){
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(int, const QString &)>(&QCameraControl::error), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(int, const QString &)>(&MyQCameraControl::Signal_Error));;
}

void QCameraControl_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(int, const QString &)>(&QCameraControl::error), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(int, const QString &)>(&MyQCameraControl::Signal_Error));;
}

int QCameraControl_IsCaptureModeSupported(void* ptr, int mode){
	return static_cast<QCameraControl*>(ptr)->isCaptureModeSupported(static_cast<QCamera::CaptureMode>(mode));
}

void QCameraControl_SetCaptureMode(void* ptr, int mode){
	static_cast<QCameraControl*>(ptr)->setCaptureMode(static_cast<QCamera::CaptureMode>(mode));
}

void QCameraControl_SetState(void* ptr, int state){
	static_cast<QCameraControl*>(ptr)->setState(static_cast<QCamera::State>(state));
}

void QCameraControl_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::State)>(&QCameraControl::stateChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::State)>(&MyQCameraControl::Signal_StateChanged));;
}

void QCameraControl_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::State)>(&QCameraControl::stateChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::State)>(&MyQCameraControl::Signal_StateChanged));;
}

int QCameraControl_Status(void* ptr){
	return static_cast<QCameraControl*>(ptr)->status();
}

void QCameraControl_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::Status)>(&QCameraControl::statusChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::Status)>(&MyQCameraControl::Signal_StatusChanged));;
}

void QCameraControl_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraControl*>(ptr), static_cast<void (QCameraControl::*)(QCamera::Status)>(&QCameraControl::statusChanged), static_cast<MyQCameraControl*>(ptr), static_cast<void (MyQCameraControl::*)(QCamera::Status)>(&MyQCameraControl::Signal_StatusChanged));;
}

void QCameraControl_DestroyQCameraControl(void* ptr){
	static_cast<QCameraControl*>(ptr)->~QCameraControl();
}

class MyQCameraExposure: public QCameraExposure {
public:
	void Signal_ApertureRangeChanged() { callbackQCameraExposureApertureRangeChanged(this->objectName().toUtf8().data()); };
	void Signal_FlashReady(bool ready) { callbackQCameraExposureFlashReady(this->objectName().toUtf8().data(), ready); };
	void Signal_IsoSensitivityChanged(int value) { callbackQCameraExposureIsoSensitivityChanged(this->objectName().toUtf8().data(), value); };
	void Signal_ShutterSpeedRangeChanged() { callbackQCameraExposureShutterSpeedRangeChanged(this->objectName().toUtf8().data()); };
protected:
};

double QCameraExposure_Aperture(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->aperture());
}

double QCameraExposure_ExposureCompensation(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->exposureCompensation());
}

int QCameraExposure_ExposureMode(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->exposureMode();
}

int QCameraExposure_FlashMode(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->flashMode();
}

int QCameraExposure_IsoSensitivity(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->isoSensitivity();
}

int QCameraExposure_MeteringMode(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->meteringMode();
}

void QCameraExposure_SetAutoAperture(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoAperture");
}

void QCameraExposure_SetAutoIsoSensitivity(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoIsoSensitivity");
}

void QCameraExposure_SetExposureCompensation(void* ptr, double ev){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setExposureCompensation", Q_ARG(qreal, static_cast<qreal>(ev)));
}

void QCameraExposure_SetExposureMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setExposureMode", Q_ARG(QCameraExposure::ExposureMode, static_cast<QCameraExposure::ExposureMode>(mode)));
}

void QCameraExposure_SetFlashMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setFlashMode", Q_ARG(QCameraExposure::FlashMode, static_cast<QCameraExposure::FlashMode>(mode)));
}

void QCameraExposure_SetManualAperture(void* ptr, double aperture){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualAperture", Q_ARG(qreal, static_cast<qreal>(aperture)));
}

void QCameraExposure_SetManualIsoSensitivity(void* ptr, int iso){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualIsoSensitivity", Q_ARG(int, iso));
}

void QCameraExposure_SetMeteringMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setMeteringMode", Q_ARG(QCameraExposure::MeteringMode, static_cast<QCameraExposure::MeteringMode>(mode)));
}

void QCameraExposure_SetSpotMeteringPoint(void* ptr, void* point){
	static_cast<QCameraExposure*>(ptr)->setSpotMeteringPoint(*static_cast<QPointF*>(point));
}

void QCameraExposure_ConnectApertureRangeChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::apertureRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ApertureRangeChanged));;
}

void QCameraExposure_DisconnectApertureRangeChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::apertureRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ApertureRangeChanged));;
}

void QCameraExposure_ConnectFlashReady(void* ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(bool)>(&QCameraExposure::flashReady), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(bool)>(&MyQCameraExposure::Signal_FlashReady));;
}

void QCameraExposure_DisconnectFlashReady(void* ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(bool)>(&QCameraExposure::flashReady), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(bool)>(&MyQCameraExposure::Signal_FlashReady));;
}

int QCameraExposure_IsAvailable(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->isAvailable();
}

int QCameraExposure_IsExposureModeSupported(void* ptr, int mode){
	return static_cast<QCameraExposure*>(ptr)->isExposureModeSupported(static_cast<QCameraExposure::ExposureMode>(mode));
}

int QCameraExposure_IsFlashModeSupported(void* ptr, int mode){
	return static_cast<QCameraExposure*>(ptr)->isFlashModeSupported(static_cast<QCameraExposure::FlashMode>(mode));
}

int QCameraExposure_IsFlashReady(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->isFlashReady();
}

int QCameraExposure_IsMeteringModeSupported(void* ptr, int mode){
	return static_cast<QCameraExposure*>(ptr)->isMeteringModeSupported(static_cast<QCameraExposure::MeteringMode>(mode));
}

void QCameraExposure_ConnectIsoSensitivityChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(int)>(&QCameraExposure::isoSensitivityChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(int)>(&MyQCameraExposure::Signal_IsoSensitivityChanged));;
}

void QCameraExposure_DisconnectIsoSensitivityChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(int)>(&QCameraExposure::isoSensitivityChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(int)>(&MyQCameraExposure::Signal_IsoSensitivityChanged));;
}

double QCameraExposure_RequestedAperture(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->requestedAperture());
}

int QCameraExposure_RequestedIsoSensitivity(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->requestedIsoSensitivity();
}

double QCameraExposure_RequestedShutterSpeed(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->requestedShutterSpeed());
}

void QCameraExposure_SetAutoShutterSpeed(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoShutterSpeed");
}

void QCameraExposure_SetManualShutterSpeed(void* ptr, double seconds){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualShutterSpeed", Q_ARG(qreal, static_cast<qreal>(seconds)));
}

double QCameraExposure_ShutterSpeed(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->shutterSpeed());
}

void QCameraExposure_ConnectShutterSpeedRangeChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::shutterSpeedRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ShutterSpeedRangeChanged));;
}

void QCameraExposure_DisconnectShutterSpeedRangeChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::shutterSpeedRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ShutterSpeedRangeChanged));;
}

class MyQCameraExposureControl: public QCameraExposureControl {
public:
	void Signal_ActualValueChanged(int parameter) { callbackQCameraExposureControlActualValueChanged(this->objectName().toUtf8().data(), parameter); };
	void Signal_ParameterRangeChanged(int parameter) { callbackQCameraExposureControlParameterRangeChanged(this->objectName().toUtf8().data(), parameter); };
	void Signal_RequestedValueChanged(int parameter) { callbackQCameraExposureControlRequestedValueChanged(this->objectName().toUtf8().data(), parameter); };
protected:
};

void* QCameraExposureControl_ActualValue(void* ptr, int parameter){
	return new QVariant(static_cast<QCameraExposureControl*>(ptr)->actualValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter)));
}

void QCameraExposureControl_ConnectActualValueChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::actualValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ActualValueChanged));;
}

void QCameraExposureControl_DisconnectActualValueChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::actualValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ActualValueChanged));;
}

int QCameraExposureControl_IsParameterSupported(void* ptr, int parameter){
	return static_cast<QCameraExposureControl*>(ptr)->isParameterSupported(static_cast<QCameraExposureControl::ExposureParameter>(parameter));
}

void QCameraExposureControl_ConnectParameterRangeChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::parameterRangeChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ParameterRangeChanged));;
}

void QCameraExposureControl_DisconnectParameterRangeChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::parameterRangeChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ParameterRangeChanged));;
}

void* QCameraExposureControl_RequestedValue(void* ptr, int parameter){
	return new QVariant(static_cast<QCameraExposureControl*>(ptr)->requestedValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter)));
}

void QCameraExposureControl_ConnectRequestedValueChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::requestedValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_RequestedValueChanged));;
}

void QCameraExposureControl_DisconnectRequestedValueChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::requestedValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_RequestedValueChanged));;
}

int QCameraExposureControl_SetValue(void* ptr, int parameter, void* value){
	return static_cast<QCameraExposureControl*>(ptr)->setValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter), *static_cast<QVariant*>(value));
}

void QCameraExposureControl_DestroyQCameraExposureControl(void* ptr){
	static_cast<QCameraExposureControl*>(ptr)->~QCameraExposureControl();
}

int QCameraFeedbackControl_IsEventFeedbackEnabled(void* ptr, int event){
	return static_cast<QCameraFeedbackControl*>(ptr)->isEventFeedbackEnabled(static_cast<QCameraFeedbackControl::EventType>(event));
}

int QCameraFeedbackControl_IsEventFeedbackLocked(void* ptr, int event){
	return static_cast<QCameraFeedbackControl*>(ptr)->isEventFeedbackLocked(static_cast<QCameraFeedbackControl::EventType>(event));
}

void QCameraFeedbackControl_ResetEventFeedback(void* ptr, int event){
	static_cast<QCameraFeedbackControl*>(ptr)->resetEventFeedback(static_cast<QCameraFeedbackControl::EventType>(event));
}

int QCameraFeedbackControl_SetEventFeedbackEnabled(void* ptr, int event, int enabled){
	return static_cast<QCameraFeedbackControl*>(ptr)->setEventFeedbackEnabled(static_cast<QCameraFeedbackControl::EventType>(event), enabled != 0);
}

int QCameraFeedbackControl_SetEventFeedbackSound(void* ptr, int event, char* filePath){
	return static_cast<QCameraFeedbackControl*>(ptr)->setEventFeedbackSound(static_cast<QCameraFeedbackControl::EventType>(event), QString(filePath));
}

void QCameraFeedbackControl_DestroyQCameraFeedbackControl(void* ptr){
	static_cast<QCameraFeedbackControl*>(ptr)->~QCameraFeedbackControl();
}

class MyQCameraFlashControl: public QCameraFlashControl {
public:
	void Signal_FlashReady(bool ready) { callbackQCameraFlashControlFlashReady(this->objectName().toUtf8().data(), ready); };
protected:
};

int QCameraFlashControl_FlashMode(void* ptr){
	return static_cast<QCameraFlashControl*>(ptr)->flashMode();
}

void QCameraFlashControl_ConnectFlashReady(void* ptr){
	QObject::connect(static_cast<QCameraFlashControl*>(ptr), static_cast<void (QCameraFlashControl::*)(bool)>(&QCameraFlashControl::flashReady), static_cast<MyQCameraFlashControl*>(ptr), static_cast<void (MyQCameraFlashControl::*)(bool)>(&MyQCameraFlashControl::Signal_FlashReady));;
}

void QCameraFlashControl_DisconnectFlashReady(void* ptr){
	QObject::disconnect(static_cast<QCameraFlashControl*>(ptr), static_cast<void (QCameraFlashControl::*)(bool)>(&QCameraFlashControl::flashReady), static_cast<MyQCameraFlashControl*>(ptr), static_cast<void (MyQCameraFlashControl::*)(bool)>(&MyQCameraFlashControl::Signal_FlashReady));;
}

int QCameraFlashControl_IsFlashModeSupported(void* ptr, int mode){
	return static_cast<QCameraFlashControl*>(ptr)->isFlashModeSupported(static_cast<QCameraExposure::FlashMode>(mode));
}

int QCameraFlashControl_IsFlashReady(void* ptr){
	return static_cast<QCameraFlashControl*>(ptr)->isFlashReady();
}

void QCameraFlashControl_SetFlashMode(void* ptr, int mode){
	static_cast<QCameraFlashControl*>(ptr)->setFlashMode(static_cast<QCameraExposure::FlashMode>(mode));
}

void QCameraFlashControl_DestroyQCameraFlashControl(void* ptr){
	static_cast<QCameraFlashControl*>(ptr)->~QCameraFlashControl();
}

class MyQCameraFocus: public QCameraFocus {
public:
	void Signal_FocusZonesChanged() { callbackQCameraFocusFocusZonesChanged(this->objectName().toUtf8().data()); };
protected:
};

double QCameraFocus_DigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraFocus*>(ptr)->digitalZoom());
}

int QCameraFocus_FocusMode(void* ptr){
	return static_cast<QCameraFocus*>(ptr)->focusMode();
}

int QCameraFocus_FocusPointMode(void* ptr){
	return static_cast<QCameraFocus*>(ptr)->focusPointMode();
}

double QCameraFocus_OpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraFocus*>(ptr)->opticalZoom());
}

void QCameraFocus_SetCustomFocusPoint(void* ptr, void* point){
	static_cast<QCameraFocus*>(ptr)->setCustomFocusPoint(*static_cast<QPointF*>(point));
}

void QCameraFocus_SetFocusMode(void* ptr, int mode){
	static_cast<QCameraFocus*>(ptr)->setFocusMode(static_cast<QCameraFocus::FocusMode>(mode));
}

void QCameraFocus_SetFocusPointMode(void* ptr, int mode){
	static_cast<QCameraFocus*>(ptr)->setFocusPointMode(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocus_ConnectFocusZonesChanged(void* ptr){
	QObject::connect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)()>(&QCameraFocus::focusZonesChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)()>(&MyQCameraFocus::Signal_FocusZonesChanged));;
}

void QCameraFocus_DisconnectFocusZonesChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)()>(&QCameraFocus::focusZonesChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)()>(&MyQCameraFocus::Signal_FocusZonesChanged));;
}

int QCameraFocus_IsAvailable(void* ptr){
	return static_cast<QCameraFocus*>(ptr)->isAvailable();
}

int QCameraFocus_IsFocusModeSupported(void* ptr, int mode){
	return static_cast<QCameraFocus*>(ptr)->isFocusModeSupported(static_cast<QCameraFocus::FocusMode>(mode));
}

int QCameraFocus_IsFocusPointModeSupported(void* ptr, int mode){
	return static_cast<QCameraFocus*>(ptr)->isFocusPointModeSupported(static_cast<QCameraFocus::FocusPointMode>(mode));
}

double QCameraFocus_MaximumDigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraFocus*>(ptr)->maximumDigitalZoom());
}

double QCameraFocus_MaximumOpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraFocus*>(ptr)->maximumOpticalZoom());
}

void QCameraFocus_ZoomTo(void* ptr, double optical, double digital){
	static_cast<QCameraFocus*>(ptr)->zoomTo(static_cast<qreal>(optical), static_cast<qreal>(digital));
}

class MyQCameraFocusControl: public QCameraFocusControl {
public:
	void Signal_FocusModeChanged(QCameraFocus::FocusModes mode) { callbackQCameraFocusControlFocusModeChanged(this->objectName().toUtf8().data(), mode); };
	void Signal_FocusPointModeChanged(QCameraFocus::FocusPointMode mode) { callbackQCameraFocusControlFocusPointModeChanged(this->objectName().toUtf8().data(), mode); };
	void Signal_FocusZonesChanged() { callbackQCameraFocusControlFocusZonesChanged(this->objectName().toUtf8().data()); };
protected:
};

int QCameraFocusControl_FocusMode(void* ptr){
	return static_cast<QCameraFocusControl*>(ptr)->focusMode();
}

void QCameraFocusControl_ConnectFocusModeChanged(void* ptr){
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusModes)>(&QCameraFocusControl::focusModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusModes)>(&MyQCameraFocusControl::Signal_FocusModeChanged));;
}

void QCameraFocusControl_DisconnectFocusModeChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusModes)>(&QCameraFocusControl::focusModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusModes)>(&MyQCameraFocusControl::Signal_FocusModeChanged));;
}

int QCameraFocusControl_FocusPointMode(void* ptr){
	return static_cast<QCameraFocusControl*>(ptr)->focusPointMode();
}

void QCameraFocusControl_ConnectFocusPointModeChanged(void* ptr){
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&QCameraFocusControl::focusPointModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&MyQCameraFocusControl::Signal_FocusPointModeChanged));;
}

void QCameraFocusControl_DisconnectFocusPointModeChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&QCameraFocusControl::focusPointModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&MyQCameraFocusControl::Signal_FocusPointModeChanged));;
}

void QCameraFocusControl_ConnectFocusZonesChanged(void* ptr){
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)()>(&QCameraFocusControl::focusZonesChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)()>(&MyQCameraFocusControl::Signal_FocusZonesChanged));;
}

void QCameraFocusControl_DisconnectFocusZonesChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)()>(&QCameraFocusControl::focusZonesChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)()>(&MyQCameraFocusControl::Signal_FocusZonesChanged));;
}

int QCameraFocusControl_IsFocusModeSupported(void* ptr, int mode){
	return static_cast<QCameraFocusControl*>(ptr)->isFocusModeSupported(static_cast<QCameraFocus::FocusMode>(mode));
}

int QCameraFocusControl_IsFocusPointModeSupported(void* ptr, int mode){
	return static_cast<QCameraFocusControl*>(ptr)->isFocusPointModeSupported(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocusControl_SetCustomFocusPoint(void* ptr, void* point){
	static_cast<QCameraFocusControl*>(ptr)->setCustomFocusPoint(*static_cast<QPointF*>(point));
}

void QCameraFocusControl_SetFocusMode(void* ptr, int mode){
	static_cast<QCameraFocusControl*>(ptr)->setFocusMode(static_cast<QCameraFocus::FocusMode>(mode));
}

void QCameraFocusControl_SetFocusPointMode(void* ptr, int mode){
	static_cast<QCameraFocusControl*>(ptr)->setFocusPointMode(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocusControl_DestroyQCameraFocusControl(void* ptr){
	static_cast<QCameraFocusControl*>(ptr)->~QCameraFocusControl();
}

void* QCameraFocusZone_NewQCameraFocusZone(void* other){
	return new QCameraFocusZone(*static_cast<QCameraFocusZone*>(other));
}

int QCameraFocusZone_IsValid(void* ptr){
	return static_cast<QCameraFocusZone*>(ptr)->isValid();
}

int QCameraFocusZone_Status(void* ptr){
	return static_cast<QCameraFocusZone*>(ptr)->status();
}

void QCameraFocusZone_DestroyQCameraFocusZone(void* ptr){
	static_cast<QCameraFocusZone*>(ptr)->~QCameraFocusZone();
}

class MyQCameraImageCapture: public QCameraImageCapture {
public:
	MyQCameraImageCapture(QMediaObject *mediaObject, QObject *parent) : QCameraImageCapture(mediaObject, parent) {};
	void Signal_BufferFormatChanged(QVideoFrame::PixelFormat format) { callbackQCameraImageCaptureBufferFormatChanged(this->objectName().toUtf8().data(), format); };
	void Signal_CaptureDestinationChanged(QCameraImageCapture::CaptureDestinations destination) { callbackQCameraImageCaptureCaptureDestinationChanged(this->objectName().toUtf8().data(), destination); };
	void Signal_ImageExposed(int id) { callbackQCameraImageCaptureImageExposed(this->objectName().toUtf8().data(), id); };
	void Signal_ImageMetadataAvailable(int id, const QString & key, const QVariant & value) { callbackQCameraImageCaptureImageMetadataAvailable(this->objectName().toUtf8().data(), id, key.toUtf8().data(), new QVariant(value)); };
	void Signal_ImageSaved(int id, const QString & fileName) { callbackQCameraImageCaptureImageSaved(this->objectName().toUtf8().data(), id, fileName.toUtf8().data()); };
	void Signal_ReadyForCaptureChanged(bool ready) { callbackQCameraImageCaptureReadyForCaptureChanged(this->objectName().toUtf8().data(), ready); };
protected:
};

int QCameraImageCapture_IsReadyForCapture(void* ptr){
	return static_cast<QCameraImageCapture*>(ptr)->isReadyForCapture();
}

void* QCameraImageCapture_NewQCameraImageCapture(void* mediaObject, void* parent){
	return new MyQCameraImageCapture(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

int QCameraImageCapture_BufferFormat(void* ptr){
	return static_cast<QCameraImageCapture*>(ptr)->bufferFormat();
}

void QCameraImageCapture_ConnectBufferFormatChanged(void* ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&QCameraImageCapture::bufferFormatChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&MyQCameraImageCapture::Signal_BufferFormatChanged));;
}

void QCameraImageCapture_DisconnectBufferFormatChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&QCameraImageCapture::bufferFormatChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QVideoFrame::PixelFormat)>(&MyQCameraImageCapture::Signal_BufferFormatChanged));;
}

void QCameraImageCapture_CancelCapture(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCameraImageCapture*>(ptr), "cancelCapture");
}

int QCameraImageCapture_Capture(void* ptr, char* file){
	return QMetaObject::invokeMethod(static_cast<QCameraImageCapture*>(ptr), "capture", Q_ARG(QString, QString(file)));
}

int QCameraImageCapture_CaptureDestination(void* ptr){
	return static_cast<QCameraImageCapture*>(ptr)->captureDestination();
}

void QCameraImageCapture_ConnectCaptureDestinationChanged(void* ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraImageCapture::captureDestinationChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraImageCapture::Signal_CaptureDestinationChanged));;
}

void QCameraImageCapture_DisconnectCaptureDestinationChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&QCameraImageCapture::captureDestinationChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(QCameraImageCapture::CaptureDestinations)>(&MyQCameraImageCapture::Signal_CaptureDestinationChanged));;
}

int QCameraImageCapture_Error(void* ptr){
	return static_cast<QCameraImageCapture*>(ptr)->error();
}

char* QCameraImageCapture_ErrorString(void* ptr){
	return static_cast<QCameraImageCapture*>(ptr)->errorString().toUtf8().data();
}

char* QCameraImageCapture_ImageCodecDescription(void* ptr, char* codec){
	return static_cast<QCameraImageCapture*>(ptr)->imageCodecDescription(QString(codec)).toUtf8().data();
}

void QCameraImageCapture_ConnectImageExposed(void* ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int)>(&QCameraImageCapture::imageExposed), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int)>(&MyQCameraImageCapture::Signal_ImageExposed));;
}

void QCameraImageCapture_DisconnectImageExposed(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int)>(&QCameraImageCapture::imageExposed), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int)>(&MyQCameraImageCapture::Signal_ImageExposed));;
}

void QCameraImageCapture_ConnectImageMetadataAvailable(void* ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &, const QVariant &)>(&QCameraImageCapture::imageMetadataAvailable), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCapture::Signal_ImageMetadataAvailable));;
}

void QCameraImageCapture_DisconnectImageMetadataAvailable(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &, const QVariant &)>(&QCameraImageCapture::imageMetadataAvailable), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCapture::Signal_ImageMetadataAvailable));;
}

void QCameraImageCapture_ConnectImageSaved(void* ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &)>(&QCameraImageCapture::imageSaved), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &)>(&MyQCameraImageCapture::Signal_ImageSaved));;
}

void QCameraImageCapture_DisconnectImageSaved(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(int, const QString &)>(&QCameraImageCapture::imageSaved), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(int, const QString &)>(&MyQCameraImageCapture::Signal_ImageSaved));;
}

int QCameraImageCapture_IsAvailable(void* ptr){
	return static_cast<QCameraImageCapture*>(ptr)->isAvailable();
}

int QCameraImageCapture_IsCaptureDestinationSupported(void* ptr, int destination){
	return static_cast<QCameraImageCapture*>(ptr)->isCaptureDestinationSupported(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void* QCameraImageCapture_MediaObject(void* ptr){
	return static_cast<QCameraImageCapture*>(ptr)->mediaObject();
}

void QCameraImageCapture_ConnectReadyForCaptureChanged(void* ptr){
	QObject::connect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(bool)>(&QCameraImageCapture::readyForCaptureChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(bool)>(&MyQCameraImageCapture::Signal_ReadyForCaptureChanged));;
}

void QCameraImageCapture_DisconnectReadyForCaptureChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCapture*>(ptr), static_cast<void (QCameraImageCapture::*)(bool)>(&QCameraImageCapture::readyForCaptureChanged), static_cast<MyQCameraImageCapture*>(ptr), static_cast<void (MyQCameraImageCapture::*)(bool)>(&MyQCameraImageCapture::Signal_ReadyForCaptureChanged));;
}

void QCameraImageCapture_SetBufferFormat(void* ptr, int format){
	static_cast<QCameraImageCapture*>(ptr)->setBufferFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraImageCapture_SetCaptureDestination(void* ptr, int destination){
	static_cast<QCameraImageCapture*>(ptr)->setCaptureDestination(static_cast<QCameraImageCapture::CaptureDestination>(destination));
}

void QCameraImageCapture_SetEncodingSettings(void* ptr, void* settings){
	static_cast<QCameraImageCapture*>(ptr)->setEncodingSettings(*static_cast<QImageEncoderSettings*>(settings));
}

char* QCameraImageCapture_SupportedImageCodecs(void* ptr){
	return static_cast<QCameraImageCapture*>(ptr)->supportedImageCodecs().join(",,,").toUtf8().data();
}

void QCameraImageCapture_DestroyQCameraImageCapture(void* ptr){
	static_cast<QCameraImageCapture*>(ptr)->~QCameraImageCapture();
}

class MyQCameraImageCaptureControl: public QCameraImageCaptureControl {
public:
	void Signal_Error(int id, int error, const QString & errorString) { callbackQCameraImageCaptureControlError(this->objectName().toUtf8().data(), id, error, errorString.toUtf8().data()); };
	void Signal_ImageExposed(int requestId) { callbackQCameraImageCaptureControlImageExposed(this->objectName().toUtf8().data(), requestId); };
	void Signal_ImageMetadataAvailable(int id, const QString & key, const QVariant & value) { callbackQCameraImageCaptureControlImageMetadataAvailable(this->objectName().toUtf8().data(), id, key.toUtf8().data(), new QVariant(value)); };
	void Signal_ImageSaved(int requestId, const QString & fileName) { callbackQCameraImageCaptureControlImageSaved(this->objectName().toUtf8().data(), requestId, fileName.toUtf8().data()); };
	void Signal_ReadyForCaptureChanged(bool ready) { callbackQCameraImageCaptureControlReadyForCaptureChanged(this->objectName().toUtf8().data(), ready); };
protected:
};

void QCameraImageCaptureControl_CancelCapture(void* ptr){
	static_cast<QCameraImageCaptureControl*>(ptr)->cancelCapture();
}

int QCameraImageCaptureControl_Capture(void* ptr, char* fileName){
	return static_cast<QCameraImageCaptureControl*>(ptr)->capture(QString(fileName));
}

int QCameraImageCaptureControl_DriveMode(void* ptr){
	return static_cast<QCameraImageCaptureControl*>(ptr)->driveMode();
}

void QCameraImageCaptureControl_ConnectError(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, int, const QString &)>(&QCameraImageCaptureControl::error), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, int, const QString &)>(&MyQCameraImageCaptureControl::Signal_Error));;
}

void QCameraImageCaptureControl_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, int, const QString &)>(&QCameraImageCaptureControl::error), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, int, const QString &)>(&MyQCameraImageCaptureControl::Signal_Error));;
}

void QCameraImageCaptureControl_ConnectImageExposed(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int)>(&QCameraImageCaptureControl::imageExposed), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int)>(&MyQCameraImageCaptureControl::Signal_ImageExposed));;
}

void QCameraImageCaptureControl_DisconnectImageExposed(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int)>(&QCameraImageCaptureControl::imageExposed), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int)>(&MyQCameraImageCaptureControl::Signal_ImageExposed));;
}

void QCameraImageCaptureControl_ConnectImageMetadataAvailable(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&QCameraImageCaptureControl::imageMetadataAvailable), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCaptureControl::Signal_ImageMetadataAvailable));;
}

void QCameraImageCaptureControl_DisconnectImageMetadataAvailable(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&QCameraImageCaptureControl::imageMetadataAvailable), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &, const QVariant &)>(&MyQCameraImageCaptureControl::Signal_ImageMetadataAvailable));;
}

void QCameraImageCaptureControl_ConnectImageSaved(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &)>(&QCameraImageCaptureControl::imageSaved), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &)>(&MyQCameraImageCaptureControl::Signal_ImageSaved));;
}

void QCameraImageCaptureControl_DisconnectImageSaved(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(int, const QString &)>(&QCameraImageCaptureControl::imageSaved), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(int, const QString &)>(&MyQCameraImageCaptureControl::Signal_ImageSaved));;
}

int QCameraImageCaptureControl_IsReadyForCapture(void* ptr){
	return static_cast<QCameraImageCaptureControl*>(ptr)->isReadyForCapture();
}

void QCameraImageCaptureControl_ConnectReadyForCaptureChanged(void* ptr){
	QObject::connect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(bool)>(&QCameraImageCaptureControl::readyForCaptureChanged), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(bool)>(&MyQCameraImageCaptureControl::Signal_ReadyForCaptureChanged));;
}

void QCameraImageCaptureControl_DisconnectReadyForCaptureChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraImageCaptureControl*>(ptr), static_cast<void (QCameraImageCaptureControl::*)(bool)>(&QCameraImageCaptureControl::readyForCaptureChanged), static_cast<MyQCameraImageCaptureControl*>(ptr), static_cast<void (MyQCameraImageCaptureControl::*)(bool)>(&MyQCameraImageCaptureControl::Signal_ReadyForCaptureChanged));;
}

void QCameraImageCaptureControl_SetDriveMode(void* ptr, int mode){
	static_cast<QCameraImageCaptureControl*>(ptr)->setDriveMode(static_cast<QCameraImageCapture::DriveMode>(mode));
}

void QCameraImageCaptureControl_DestroyQCameraImageCaptureControl(void* ptr){
	static_cast<QCameraImageCaptureControl*>(ptr)->~QCameraImageCaptureControl();
}

int QCameraImageProcessing_ColorFilter(void* ptr){
	return static_cast<QCameraImageProcessing*>(ptr)->colorFilter();
}

double QCameraImageProcessing_Contrast(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->contrast());
}

double QCameraImageProcessing_DenoisingLevel(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->denoisingLevel());
}

int QCameraImageProcessing_IsAvailable(void* ptr){
	return static_cast<QCameraImageProcessing*>(ptr)->isAvailable();
}

int QCameraImageProcessing_IsColorFilterSupported(void* ptr, int filter){
	return static_cast<QCameraImageProcessing*>(ptr)->isColorFilterSupported(static_cast<QCameraImageProcessing::ColorFilter>(filter));
}

int QCameraImageProcessing_IsWhiteBalanceModeSupported(void* ptr, int mode){
	return static_cast<QCameraImageProcessing*>(ptr)->isWhiteBalanceModeSupported(static_cast<QCameraImageProcessing::WhiteBalanceMode>(mode));
}

double QCameraImageProcessing_ManualWhiteBalance(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->manualWhiteBalance());
}

double QCameraImageProcessing_Saturation(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->saturation());
}

void QCameraImageProcessing_SetColorFilter(void* ptr, int filter){
	static_cast<QCameraImageProcessing*>(ptr)->setColorFilter(static_cast<QCameraImageProcessing::ColorFilter>(filter));
}

void QCameraImageProcessing_SetContrast(void* ptr, double value){
	static_cast<QCameraImageProcessing*>(ptr)->setContrast(static_cast<qreal>(value));
}

void QCameraImageProcessing_SetDenoisingLevel(void* ptr, double level){
	static_cast<QCameraImageProcessing*>(ptr)->setDenoisingLevel(static_cast<qreal>(level));
}

void QCameraImageProcessing_SetManualWhiteBalance(void* ptr, double colorTemperature){
	static_cast<QCameraImageProcessing*>(ptr)->setManualWhiteBalance(static_cast<qreal>(colorTemperature));
}

void QCameraImageProcessing_SetSaturation(void* ptr, double value){
	static_cast<QCameraImageProcessing*>(ptr)->setSaturation(static_cast<qreal>(value));
}

void QCameraImageProcessing_SetSharpeningLevel(void* ptr, double level){
	static_cast<QCameraImageProcessing*>(ptr)->setSharpeningLevel(static_cast<qreal>(level));
}

void QCameraImageProcessing_SetWhiteBalanceMode(void* ptr, int mode){
	static_cast<QCameraImageProcessing*>(ptr)->setWhiteBalanceMode(static_cast<QCameraImageProcessing::WhiteBalanceMode>(mode));
}

double QCameraImageProcessing_SharpeningLevel(void* ptr){
	return static_cast<double>(static_cast<QCameraImageProcessing*>(ptr)->sharpeningLevel());
}

int QCameraImageProcessing_WhiteBalanceMode(void* ptr){
	return static_cast<QCameraImageProcessing*>(ptr)->whiteBalanceMode();
}

int QCameraImageProcessingControl_IsParameterSupported(void* ptr, int parameter){
	return static_cast<QCameraImageProcessingControl*>(ptr)->isParameterSupported(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter));
}

int QCameraImageProcessingControl_IsParameterValueSupported(void* ptr, int parameter, void* value){
	return static_cast<QCameraImageProcessingControl*>(ptr)->isParameterValueSupported(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter), *static_cast<QVariant*>(value));
}

void* QCameraImageProcessingControl_Parameter(void* ptr, int parameter){
	return new QVariant(static_cast<QCameraImageProcessingControl*>(ptr)->parameter(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter)));
}

void QCameraImageProcessingControl_SetParameter(void* ptr, int parameter, void* value){
	static_cast<QCameraImageProcessingControl*>(ptr)->setParameter(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter), *static_cast<QVariant*>(value));
}

void QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(void* ptr){
	static_cast<QCameraImageProcessingControl*>(ptr)->~QCameraImageProcessingControl();
}

void* QCameraInfo_NewQCameraInfo(void* name){
	return new QCameraInfo(*static_cast<QByteArray*>(name));
}

void* QCameraInfo_NewQCameraInfo2(void* camera){
	return new QCameraInfo(*static_cast<QCamera*>(camera));
}

void* QCameraInfo_NewQCameraInfo3(void* other){
	return new QCameraInfo(*static_cast<QCameraInfo*>(other));
}

char* QCameraInfo_Description(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->description().toUtf8().data();
}

char* QCameraInfo_DeviceName(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->deviceName().toUtf8().data();
}

int QCameraInfo_IsNull(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->isNull();
}

int QCameraInfo_Orientation(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->orientation();
}

int QCameraInfo_Position(void* ptr){
	return static_cast<QCameraInfo*>(ptr)->position();
}

void QCameraInfo_DestroyQCameraInfo(void* ptr){
	static_cast<QCameraInfo*>(ptr)->~QCameraInfo();
}

class MyQCameraInfoControl: public QCameraInfoControl {
public:
protected:
};

int QCameraInfoControl_CameraOrientation(void* ptr, char* deviceName){
	return static_cast<QCameraInfoControl*>(ptr)->cameraOrientation(QString(deviceName));
}

int QCameraInfoControl_CameraPosition(void* ptr, char* deviceName){
	return static_cast<QCameraInfoControl*>(ptr)->cameraPosition(QString(deviceName));
}

void QCameraInfoControl_DestroyQCameraInfoControl(void* ptr){
	static_cast<QCameraInfoControl*>(ptr)->~QCameraInfoControl();
}

class MyQCameraLocksControl: public QCameraLocksControl {
public:
	void Signal_LockStatusChanged(QCamera::LockType lock, QCamera::LockStatus status, QCamera::LockChangeReason reason) { callbackQCameraLocksControlLockStatusChanged(this->objectName().toUtf8().data(), lock, status, reason); };
protected:
};

int QCameraLocksControl_LockStatus(void* ptr, int lock){
	return static_cast<QCameraLocksControl*>(ptr)->lockStatus(static_cast<QCamera::LockType>(lock));
}

void QCameraLocksControl_ConnectLockStatusChanged(void* ptr){
	QObject::connect(static_cast<QCameraLocksControl*>(ptr), static_cast<void (QCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&QCameraLocksControl::lockStatusChanged), static_cast<MyQCameraLocksControl*>(ptr), static_cast<void (MyQCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCameraLocksControl::Signal_LockStatusChanged));;
}

void QCameraLocksControl_DisconnectLockStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraLocksControl*>(ptr), static_cast<void (QCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&QCameraLocksControl::lockStatusChanged), static_cast<MyQCameraLocksControl*>(ptr), static_cast<void (MyQCameraLocksControl::*)(QCamera::LockType, QCamera::LockStatus, QCamera::LockChangeReason)>(&MyQCameraLocksControl::Signal_LockStatusChanged));;
}

void QCameraLocksControl_SearchAndLock(void* ptr, int locks){
	static_cast<QCameraLocksControl*>(ptr)->searchAndLock(static_cast<QCamera::LockType>(locks));
}

int QCameraLocksControl_SupportedLocks(void* ptr){
	return static_cast<QCameraLocksControl*>(ptr)->supportedLocks();
}

void QCameraLocksControl_Unlock(void* ptr, int locks){
	static_cast<QCameraLocksControl*>(ptr)->unlock(static_cast<QCamera::LockType>(locks));
}

void QCameraLocksControl_DestroyQCameraLocksControl(void* ptr){
	static_cast<QCameraLocksControl*>(ptr)->~QCameraLocksControl();
}

void* QCameraViewfinderSettings_NewQCameraViewfinderSettings(){
	return new QCameraViewfinderSettings();
}

void* QCameraViewfinderSettings_NewQCameraViewfinderSettings2(void* other){
	return new QCameraViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(other));
}

int QCameraViewfinderSettings_IsNull(void* ptr){
	return static_cast<QCameraViewfinderSettings*>(ptr)->isNull();
}

double QCameraViewfinderSettings_MaximumFrameRate(void* ptr){
	return static_cast<double>(static_cast<QCameraViewfinderSettings*>(ptr)->maximumFrameRate());
}

double QCameraViewfinderSettings_MinimumFrameRate(void* ptr){
	return static_cast<double>(static_cast<QCameraViewfinderSettings*>(ptr)->minimumFrameRate());
}

int QCameraViewfinderSettings_PixelFormat(void* ptr){
	return static_cast<QCameraViewfinderSettings*>(ptr)->pixelFormat();
}

void QCameraViewfinderSettings_SetMaximumFrameRate(void* ptr, double rate){
	static_cast<QCameraViewfinderSettings*>(ptr)->setMaximumFrameRate(static_cast<qreal>(rate));
}

void QCameraViewfinderSettings_SetMinimumFrameRate(void* ptr, double rate){
	static_cast<QCameraViewfinderSettings*>(ptr)->setMinimumFrameRate(static_cast<qreal>(rate));
}

void QCameraViewfinderSettings_SetPixelAspectRatio(void* ptr, void* ratio){
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelAspectRatio(*static_cast<QSize*>(ratio));
}

void QCameraViewfinderSettings_SetPixelAspectRatio2(void* ptr, int horizontal, int vertical){
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelAspectRatio(horizontal, vertical);
}

void QCameraViewfinderSettings_SetPixelFormat(void* ptr, int format){
	static_cast<QCameraViewfinderSettings*>(ptr)->setPixelFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

void QCameraViewfinderSettings_SetResolution(void* ptr, void* resolution){
	static_cast<QCameraViewfinderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QCameraViewfinderSettings_SetResolution2(void* ptr, int width, int height){
	static_cast<QCameraViewfinderSettings*>(ptr)->setResolution(width, height);
}

void QCameraViewfinderSettings_Swap(void* ptr, void* other){
	static_cast<QCameraViewfinderSettings*>(ptr)->swap(*static_cast<QCameraViewfinderSettings*>(other));
}

void QCameraViewfinderSettings_DestroyQCameraViewfinderSettings(void* ptr){
	static_cast<QCameraViewfinderSettings*>(ptr)->~QCameraViewfinderSettings();
}

int QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(void* ptr, int parameter){
	return static_cast<QCameraViewfinderSettingsControl*>(ptr)->isViewfinderParameterSupported(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter));
}

void QCameraViewfinderSettingsControl_SetViewfinderParameter(void* ptr, int parameter, void* value){
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->setViewfinderParameter(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter), *static_cast<QVariant*>(value));
}

void* QCameraViewfinderSettingsControl_ViewfinderParameter(void* ptr, int parameter){
	return new QVariant(static_cast<QCameraViewfinderSettingsControl*>(ptr)->viewfinderParameter(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter)));
}

void QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(void* ptr){
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->~QCameraViewfinderSettingsControl();
}

class MyQCameraViewfinderSettingsControl2: public QCameraViewfinderSettingsControl2 {
public:
protected:
};

void QCameraViewfinderSettingsControl2_SetViewfinderSettings(void* ptr, void* settings){
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->setViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(settings));
}

void QCameraViewfinderSettingsControl2_DestroyQCameraViewfinderSettingsControl2(void* ptr){
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->~QCameraViewfinderSettingsControl2();
}

class MyQCameraZoomControl: public QCameraZoomControl {
public:
protected:
};

double QCameraZoomControl_CurrentDigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->currentDigitalZoom());
}

double QCameraZoomControl_CurrentOpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->currentOpticalZoom());
}

double QCameraZoomControl_MaximumDigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->maximumDigitalZoom());
}

double QCameraZoomControl_MaximumOpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->maximumOpticalZoom());
}

double QCameraZoomControl_RequestedDigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->requestedDigitalZoom());
}

double QCameraZoomControl_RequestedOpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->requestedOpticalZoom());
}

void QCameraZoomControl_ZoomTo(void* ptr, double optical, double digital){
	static_cast<QCameraZoomControl*>(ptr)->zoomTo(static_cast<qreal>(optical), static_cast<qreal>(digital));
}

void QCameraZoomControl_DestroyQCameraZoomControl(void* ptr){
	static_cast<QCameraZoomControl*>(ptr)->~QCameraZoomControl();
}

class MyQImageEncoderControl: public QImageEncoderControl {
public:
protected:
};

char* QImageEncoderControl_ImageCodecDescription(void* ptr, char* codec){
	return static_cast<QImageEncoderControl*>(ptr)->imageCodecDescription(QString(codec)).toUtf8().data();
}

void QImageEncoderControl_SetImageSettings(void* ptr, void* settings){
	static_cast<QImageEncoderControl*>(ptr)->setImageSettings(*static_cast<QImageEncoderSettings*>(settings));
}

char* QImageEncoderControl_SupportedImageCodecs(void* ptr){
	return static_cast<QImageEncoderControl*>(ptr)->supportedImageCodecs().join(",,,").toUtf8().data();
}

void QImageEncoderControl_DestroyQImageEncoderControl(void* ptr){
	static_cast<QImageEncoderControl*>(ptr)->~QImageEncoderControl();
}

void* QImageEncoderSettings_NewQImageEncoderSettings(){
	return new QImageEncoderSettings();
}

void* QImageEncoderSettings_NewQImageEncoderSettings2(void* other){
	return new QImageEncoderSettings(*static_cast<QImageEncoderSettings*>(other));
}

char* QImageEncoderSettings_Codec(void* ptr){
	return static_cast<QImageEncoderSettings*>(ptr)->codec().toUtf8().data();
}

void* QImageEncoderSettings_EncodingOption(void* ptr, char* option){
	return new QVariant(static_cast<QImageEncoderSettings*>(ptr)->encodingOption(QString(option)));
}

int QImageEncoderSettings_IsNull(void* ptr){
	return static_cast<QImageEncoderSettings*>(ptr)->isNull();
}

void QImageEncoderSettings_SetCodec(void* ptr, char* codec){
	static_cast<QImageEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QImageEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value){
	static_cast<QImageEncoderSettings*>(ptr)->setEncodingOption(QString(option), *static_cast<QVariant*>(value));
}

void QImageEncoderSettings_SetResolution(void* ptr, void* resolution){
	static_cast<QImageEncoderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QImageEncoderSettings_SetResolution2(void* ptr, int width, int height){
	static_cast<QImageEncoderSettings*>(ptr)->setResolution(width, height);
}

void QImageEncoderSettings_DestroyQImageEncoderSettings(void* ptr){
	static_cast<QImageEncoderSettings*>(ptr)->~QImageEncoderSettings();
}

class MyQMediaAudioProbeControl: public QMediaAudioProbeControl {
public:
	void Signal_Flush() { callbackQMediaAudioProbeControlFlush(this->objectName().toUtf8().data()); };
protected:
};

void QMediaAudioProbeControl_ConnectFlush(void* ptr){
	QObject::connect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)()>(&QMediaAudioProbeControl::flush), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)()>(&MyQMediaAudioProbeControl::Signal_Flush));;
}

void QMediaAudioProbeControl_DisconnectFlush(void* ptr){
	QObject::disconnect(static_cast<QMediaAudioProbeControl*>(ptr), static_cast<void (QMediaAudioProbeControl::*)()>(&QMediaAudioProbeControl::flush), static_cast<MyQMediaAudioProbeControl*>(ptr), static_cast<void (MyQMediaAudioProbeControl::*)()>(&MyQMediaAudioProbeControl::Signal_Flush));;
}

void QMediaAudioProbeControl_DestroyQMediaAudioProbeControl(void* ptr){
	static_cast<QMediaAudioProbeControl*>(ptr)->~QMediaAudioProbeControl();
}

class MyQMediaAvailabilityControl: public QMediaAvailabilityControl {
public:
protected:
};

void QMediaAvailabilityControl_DestroyQMediaAvailabilityControl(void* ptr){
	static_cast<QMediaAvailabilityControl*>(ptr)->~QMediaAvailabilityControl();
}

class MyQMediaBindableInterface: public QMediaBindableInterface {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

void* QMediaBindableInterface_MediaObject(void* ptr){
	return static_cast<QMediaBindableInterface*>(ptr)->mediaObject();
}

void QMediaBindableInterface_DestroyQMediaBindableInterface(void* ptr){
	static_cast<QMediaBindableInterface*>(ptr)->~QMediaBindableInterface();
}

char* QMediaBindableInterface_ObjectNameAbs(void* ptr){
	return static_cast<MyQMediaBindableInterface*>(ptr)->objectNameAbs().toUtf8().data();
}

void QMediaBindableInterface_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQMediaBindableInterface*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQMediaContainerControl: public QMediaContainerControl {
public:
protected:
};

char* QMediaContainerControl_ContainerDescription(void* ptr, char* format){
	return static_cast<QMediaContainerControl*>(ptr)->containerDescription(QString(format)).toUtf8().data();
}

char* QMediaContainerControl_ContainerFormat(void* ptr){
	return static_cast<QMediaContainerControl*>(ptr)->containerFormat().toUtf8().data();
}

void QMediaContainerControl_SetContainerFormat(void* ptr, char* format){
	static_cast<QMediaContainerControl*>(ptr)->setContainerFormat(QString(format));
}

char* QMediaContainerControl_SupportedContainers(void* ptr){
	return static_cast<QMediaContainerControl*>(ptr)->supportedContainers().join(",,,").toUtf8().data();
}

void QMediaContainerControl_DestroyQMediaContainerControl(void* ptr){
	static_cast<QMediaContainerControl*>(ptr)->~QMediaContainerControl();
}

void* QMediaContent_NewQMediaContent(){
	return new QMediaContent();
}

void* QMediaContent_NewQMediaContent7(void* playlist, void* contentUrl, int takeOwnership){
	return new QMediaContent(static_cast<QMediaPlaylist*>(playlist), *static_cast<QUrl*>(contentUrl), takeOwnership != 0);
}

void* QMediaContent_NewQMediaContent6(void* other){
	return new QMediaContent(*static_cast<QMediaContent*>(other));
}

void* QMediaContent_NewQMediaContent4(void* resource){
	return new QMediaContent(*static_cast<QMediaResource*>(resource));
}

void* QMediaContent_NewQMediaContent3(void* request){
	return new QMediaContent(*static_cast<QNetworkRequest*>(request));
}

void* QMediaContent_NewQMediaContent2(void* url){
	return new QMediaContent(*static_cast<QUrl*>(url));
}

int QMediaContent_IsNull(void* ptr){
	return static_cast<QMediaContent*>(ptr)->isNull();
}

void* QMediaContent_Playlist(void* ptr){
	return static_cast<QMediaContent*>(ptr)->playlist();
}

void QMediaContent_DestroyQMediaContent(void* ptr){
	static_cast<QMediaContent*>(ptr)->~QMediaContent();
}

void QMediaControl_DestroyQMediaControl(void* ptr){
	static_cast<QMediaControl*>(ptr)->~QMediaControl();
}

class MyQMediaGaplessPlaybackControl: public QMediaGaplessPlaybackControl {
public:
	void Signal_AdvancedToNextMedia() { callbackQMediaGaplessPlaybackControlAdvancedToNextMedia(this->objectName().toUtf8().data()); };
protected:
};

void QMediaGaplessPlaybackControl_ConnectAdvancedToNextMedia(void* ptr){
	QObject::connect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)()>(&QMediaGaplessPlaybackControl::advancedToNextMedia), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)()>(&MyQMediaGaplessPlaybackControl::Signal_AdvancedToNextMedia));;
}

void QMediaGaplessPlaybackControl_DisconnectAdvancedToNextMedia(void* ptr){
	QObject::disconnect(static_cast<QMediaGaplessPlaybackControl*>(ptr), static_cast<void (QMediaGaplessPlaybackControl::*)()>(&QMediaGaplessPlaybackControl::advancedToNextMedia), static_cast<MyQMediaGaplessPlaybackControl*>(ptr), static_cast<void (MyQMediaGaplessPlaybackControl::*)()>(&MyQMediaGaplessPlaybackControl::Signal_AdvancedToNextMedia));;
}

double QMediaGaplessPlaybackControl_CrossfadeTime(void* ptr){
	return static_cast<double>(static_cast<QMediaGaplessPlaybackControl*>(ptr)->crossfadeTime());
}

int QMediaGaplessPlaybackControl_IsCrossfadeSupported(void* ptr){
	return static_cast<QMediaGaplessPlaybackControl*>(ptr)->isCrossfadeSupported();
}

void QMediaGaplessPlaybackControl_SetCrossfadeTime(void* ptr, double crossfadeTime){
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->setCrossfadeTime(static_cast<qreal>(crossfadeTime));
}

void QMediaGaplessPlaybackControl_SetNextMedia(void* ptr, void* media){
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->setNextMedia(*static_cast<QMediaContent*>(media));
}

void QMediaGaplessPlaybackControl_DestroyQMediaGaplessPlaybackControl(void* ptr){
	static_cast<QMediaGaplessPlaybackControl*>(ptr)->~QMediaGaplessPlaybackControl();
}

class MyQMediaNetworkAccessControl: public QMediaNetworkAccessControl {
public:
protected:
};

void QMediaNetworkAccessControl_DestroyQMediaNetworkAccessControl(void* ptr){
	static_cast<QMediaNetworkAccessControl*>(ptr)->~QMediaNetworkAccessControl();
}

class MyQMediaObject: public QMediaObject {
public:
	void Signal_AvailabilityChanged(bool available) { callbackQMediaObjectAvailabilityChanged(this->objectName().toUtf8().data(), available); };
	void Signal_MetaDataAvailableChanged(bool available) { callbackQMediaObjectMetaDataAvailableChanged(this->objectName().toUtf8().data(), available); };
	void Signal_MetaDataChanged() { callbackQMediaObjectMetaDataChanged(this->objectName().toUtf8().data()); };
	void Signal_NotifyIntervalChanged(int milliseconds) { callbackQMediaObjectNotifyIntervalChanged(this->objectName().toUtf8().data(), milliseconds); };
	void unbind(QObject * object) { if (!callbackQMediaObjectUnbind(this->objectName().toUtf8().data(), object)) { QMediaObject::unbind(object); }; };
protected:
};

int QMediaObject_NotifyInterval(void* ptr){
	return static_cast<QMediaObject*>(ptr)->notifyInterval();
}

void QMediaObject_SetNotifyInterval(void* ptr, int milliSeconds){
	static_cast<QMediaObject*>(ptr)->setNotifyInterval(milliSeconds);
}

void QMediaObject_ConnectAvailabilityChanged(void* ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_AvailabilityChanged));;
}

void QMediaObject_DisconnectAvailabilityChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_AvailabilityChanged));;
}

char* QMediaObject_AvailableMetaData(void* ptr){
	return static_cast<QMediaObject*>(ptr)->availableMetaData().join(",,,").toUtf8().data();
}

int QMediaObject_Bind(void* ptr, void* object){
	return static_cast<QMediaObject*>(ptr)->bind(static_cast<QObject*>(object));
}

int QMediaObject_IsAvailable(void* ptr){
	return static_cast<QMediaObject*>(ptr)->isAvailable();
}

int QMediaObject_IsMetaDataAvailable(void* ptr){
	return static_cast<QMediaObject*>(ptr)->isMetaDataAvailable();
}

void* QMediaObject_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QMediaObject*>(ptr)->metaData(QString(key)));
}

void QMediaObject_ConnectMetaDataAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::metaDataAvailableChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_MetaDataAvailableChanged));;
}

void QMediaObject_DisconnectMetaDataAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::metaDataAvailableChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_MetaDataAvailableChanged));;
}

void QMediaObject_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)()>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)()>(&MyQMediaObject::Signal_MetaDataChanged));;
}

void QMediaObject_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)()>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)()>(&MyQMediaObject::Signal_MetaDataChanged));;
}

void QMediaObject_ConnectNotifyIntervalChanged(void* ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(int)>(&QMediaObject::notifyIntervalChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(int)>(&MyQMediaObject::Signal_NotifyIntervalChanged));;
}

void QMediaObject_DisconnectNotifyIntervalChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(int)>(&QMediaObject::notifyIntervalChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(int)>(&MyQMediaObject::Signal_NotifyIntervalChanged));;
}

void* QMediaObject_Service(void* ptr){
	return static_cast<QMediaObject*>(ptr)->service();
}

void QMediaObject_Unbind(void* ptr, void* object){
	static_cast<QMediaObject*>(ptr)->unbind(static_cast<QObject*>(object));
}

void QMediaObject_DestroyQMediaObject(void* ptr){
	static_cast<QMediaObject*>(ptr)->~QMediaObject();
}

class MyQMediaPlayer: public QMediaPlayer {
public:
	MyQMediaPlayer(QObject *parent, Flags flags) : QMediaPlayer(parent, flags) {};
	void Signal_AudioAvailableChanged(bool available) { callbackQMediaPlayerAudioAvailableChanged(this->objectName().toUtf8().data(), available); };
	void Signal_BufferStatusChanged(int percentFilled) { callbackQMediaPlayerBufferStatusChanged(this->objectName().toUtf8().data(), percentFilled); };
	void Signal_MediaStatusChanged(QMediaPlayer::MediaStatus status) { callbackQMediaPlayerMediaStatusChanged(this->objectName().toUtf8().data(), status); };
	void Signal_MutedChanged(bool muted) { callbackQMediaPlayerMutedChanged(this->objectName().toUtf8().data(), muted); };
	void Signal_SeekableChanged(bool seekable) { callbackQMediaPlayerSeekableChanged(this->objectName().toUtf8().data(), seekable); };
	void Signal_StateChanged(QMediaPlayer::State state) { callbackQMediaPlayerStateChanged(this->objectName().toUtf8().data(), state); };
	void Signal_VideoAvailableChanged(bool videoAvailable) { callbackQMediaPlayerVideoAvailableChanged(this->objectName().toUtf8().data(), videoAvailable); };
	void Signal_VolumeChanged(int volume) { callbackQMediaPlayerVolumeChanged(this->objectName().toUtf8().data(), volume); };
protected:
};

int QMediaPlayer_BufferStatus(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->bufferStatus();
}

char* QMediaPlayer_ErrorString(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->errorString().toUtf8().data();
}

int QMediaPlayer_IsAudioAvailable(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->isAudioAvailable();
}

int QMediaPlayer_IsMuted(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->isMuted();
}

int QMediaPlayer_IsSeekable(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->isSeekable();
}

int QMediaPlayer_IsVideoAvailable(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->isVideoAvailable();
}

int QMediaPlayer_MediaStatus(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->mediaStatus();
}

double QMediaPlayer_PlaybackRate(void* ptr){
	return static_cast<double>(static_cast<QMediaPlayer*>(ptr)->playbackRate());
}

void* QMediaPlayer_Playlist(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->playlist();
}

void QMediaPlayer_SetMuted(void* ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

void QMediaPlayer_SetPlaybackRate(void* ptr, double rate){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setPlaybackRate", Q_ARG(qreal, static_cast<qreal>(rate)));
}

void QMediaPlayer_SetPlaylist(void* ptr, void* playlist){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setPlaylist", Q_ARG(QMediaPlaylist*, static_cast<QMediaPlaylist*>(playlist)));
}

void QMediaPlayer_SetVolume(void* ptr, int volume){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setVolume", Q_ARG(int, volume));
}

int QMediaPlayer_Volume(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->volume();
}

void* QMediaPlayer_NewQMediaPlayer(void* parent, int flags){
	return new MyQMediaPlayer(static_cast<QObject*>(parent), static_cast<QMediaPlayer::Flag>(flags));
}

void QMediaPlayer_ConnectAudioAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::audioAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_AudioAvailableChanged));;
}

void QMediaPlayer_DisconnectAudioAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::audioAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_AudioAvailableChanged));;
}

void QMediaPlayer_ConnectBufferStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::bufferStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_BufferStatusChanged));;
}

void QMediaPlayer_DisconnectBufferStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::bufferStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_BufferStatusChanged));;
}

int QMediaPlayer_Error(void* ptr){
	return static_cast<QMediaPlayer*>(ptr)->error();
}

void QMediaPlayer_ConnectMediaStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayer::mediaStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayer::Signal_MediaStatusChanged));;
}

void QMediaPlayer_DisconnectMediaStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayer::mediaStatusChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayer::Signal_MediaStatusChanged));;
}

void* QMediaPlayer_MediaStream(void* ptr){
	return const_cast<QIODevice*>(static_cast<QMediaPlayer*>(ptr)->mediaStream());
}

void QMediaPlayer_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::mutedChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_MutedChanged));;
}

void QMediaPlayer_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::mutedChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_MutedChanged));;
}

void QMediaPlayer_Pause(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "pause");
}

void QMediaPlayer_Play(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "play");
}

void QMediaPlayer_ConnectSeekableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::seekableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_SeekableChanged));;
}

void QMediaPlayer_DisconnectSeekableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::seekableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_SeekableChanged));;
}

void QMediaPlayer_SetMedia(void* ptr, void* media, void* stream){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "setMedia", Q_ARG(QMediaContent, *static_cast<QMediaContent*>(media)), Q_ARG(QIODevice*, static_cast<QIODevice*>(stream)));
}

void QMediaPlayer_SetVideoOutput3(void* ptr, void* surface){
	static_cast<QMediaPlayer*>(ptr)->setVideoOutput(static_cast<QAbstractVideoSurface*>(surface));
}

void QMediaPlayer_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::State)>(&QMediaPlayer::stateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::State)>(&MyQMediaPlayer::Signal_StateChanged));;
}

void QMediaPlayer_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(QMediaPlayer::State)>(&QMediaPlayer::stateChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(QMediaPlayer::State)>(&MyQMediaPlayer::Signal_StateChanged));;
}

void QMediaPlayer_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlayer*>(ptr), "stop");
}

void QMediaPlayer_ConnectVideoAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::videoAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_VideoAvailableChanged));;
}

void QMediaPlayer_DisconnectVideoAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(bool)>(&QMediaPlayer::videoAvailableChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(bool)>(&MyQMediaPlayer::Signal_VideoAvailableChanged));;
}

void QMediaPlayer_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::volumeChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_VolumeChanged));;
}

void QMediaPlayer_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayer*>(ptr), static_cast<void (QMediaPlayer::*)(int)>(&QMediaPlayer::volumeChanged), static_cast<MyQMediaPlayer*>(ptr), static_cast<void (MyQMediaPlayer::*)(int)>(&MyQMediaPlayer::Signal_VolumeChanged));;
}

void QMediaPlayer_DestroyQMediaPlayer(void* ptr){
	static_cast<QMediaPlayer*>(ptr)->~QMediaPlayer();
}

class MyQMediaPlayerControl: public QMediaPlayerControl {
public:
	void Signal_AudioAvailableChanged(bool audio) { callbackQMediaPlayerControlAudioAvailableChanged(this->objectName().toUtf8().data(), audio); };
	void Signal_BufferStatusChanged(int progress) { callbackQMediaPlayerControlBufferStatusChanged(this->objectName().toUtf8().data(), progress); };
	void Signal_Error(int error, const QString & errorString) { callbackQMediaPlayerControlError(this->objectName().toUtf8().data(), error, errorString.toUtf8().data()); };
	void Signal_MediaStatusChanged(QMediaPlayer::MediaStatus status) { callbackQMediaPlayerControlMediaStatusChanged(this->objectName().toUtf8().data(), status); };
	void Signal_MutedChanged(bool mute) { callbackQMediaPlayerControlMutedChanged(this->objectName().toUtf8().data(), mute); };
	void Signal_SeekableChanged(bool seekable) { callbackQMediaPlayerControlSeekableChanged(this->objectName().toUtf8().data(), seekable); };
	void Signal_StateChanged(QMediaPlayer::State state) { callbackQMediaPlayerControlStateChanged(this->objectName().toUtf8().data(), state); };
	void Signal_VideoAvailableChanged(bool video) { callbackQMediaPlayerControlVideoAvailableChanged(this->objectName().toUtf8().data(), video); };
	void Signal_VolumeChanged(int volume) { callbackQMediaPlayerControlVolumeChanged(this->objectName().toUtf8().data(), volume); };
protected:
};

void QMediaPlayerControl_ConnectAudioAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::audioAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_AudioAvailableChanged));;
}

void QMediaPlayerControl_DisconnectAudioAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::audioAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_AudioAvailableChanged));;
}

int QMediaPlayerControl_BufferStatus(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->bufferStatus();
}

void QMediaPlayerControl_ConnectBufferStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::bufferStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_BufferStatusChanged));;
}

void QMediaPlayerControl_DisconnectBufferStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::bufferStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_BufferStatusChanged));;
}

void QMediaPlayerControl_ConnectError(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int, const QString &)>(&QMediaPlayerControl::error), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int, const QString &)>(&MyQMediaPlayerControl::Signal_Error));;
}

void QMediaPlayerControl_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int, const QString &)>(&QMediaPlayerControl::error), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int, const QString &)>(&MyQMediaPlayerControl::Signal_Error));;
}

int QMediaPlayerControl_IsAudioAvailable(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->isAudioAvailable();
}

int QMediaPlayerControl_IsMuted(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->isMuted();
}

int QMediaPlayerControl_IsSeekable(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->isSeekable();
}

int QMediaPlayerControl_IsVideoAvailable(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->isVideoAvailable();
}

int QMediaPlayerControl_MediaStatus(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->mediaStatus();
}

void QMediaPlayerControl_ConnectMediaStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayerControl::mediaStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayerControl::Signal_MediaStatusChanged));;
}

void QMediaPlayerControl_DisconnectMediaStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&QMediaPlayerControl::mediaStatusChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::MediaStatus)>(&MyQMediaPlayerControl::Signal_MediaStatusChanged));;
}

void* QMediaPlayerControl_MediaStream(void* ptr){
	return const_cast<QIODevice*>(static_cast<QMediaPlayerControl*>(ptr)->mediaStream());
}

void QMediaPlayerControl_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::mutedChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_MutedChanged));;
}

void QMediaPlayerControl_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::mutedChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_MutedChanged));;
}

void QMediaPlayerControl_Pause(void* ptr){
	static_cast<QMediaPlayerControl*>(ptr)->pause();
}

void QMediaPlayerControl_Play(void* ptr){
	static_cast<QMediaPlayerControl*>(ptr)->play();
}

double QMediaPlayerControl_PlaybackRate(void* ptr){
	return static_cast<double>(static_cast<QMediaPlayerControl*>(ptr)->playbackRate());
}

void QMediaPlayerControl_ConnectSeekableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::seekableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_SeekableChanged));;
}

void QMediaPlayerControl_DisconnectSeekableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::seekableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_SeekableChanged));;
}

void QMediaPlayerControl_SetMedia(void* ptr, void* media, void* stream){
	static_cast<QMediaPlayerControl*>(ptr)->setMedia(*static_cast<QMediaContent*>(media), static_cast<QIODevice*>(stream));
}

void QMediaPlayerControl_SetMuted(void* ptr, int mute){
	static_cast<QMediaPlayerControl*>(ptr)->setMuted(mute != 0);
}

void QMediaPlayerControl_SetPlaybackRate(void* ptr, double rate){
	static_cast<QMediaPlayerControl*>(ptr)->setPlaybackRate(static_cast<qreal>(rate));
}

void QMediaPlayerControl_SetVolume(void* ptr, int volume){
	static_cast<QMediaPlayerControl*>(ptr)->setVolume(volume);
}

void QMediaPlayerControl_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::State)>(&QMediaPlayerControl::stateChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::State)>(&MyQMediaPlayerControl::Signal_StateChanged));;
}

void QMediaPlayerControl_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(QMediaPlayer::State)>(&QMediaPlayerControl::stateChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(QMediaPlayer::State)>(&MyQMediaPlayerControl::Signal_StateChanged));;
}

void QMediaPlayerControl_Stop(void* ptr){
	static_cast<QMediaPlayerControl*>(ptr)->stop();
}

void QMediaPlayerControl_ConnectVideoAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::videoAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_VideoAvailableChanged));;
}

void QMediaPlayerControl_DisconnectVideoAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(bool)>(&QMediaPlayerControl::videoAvailableChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(bool)>(&MyQMediaPlayerControl::Signal_VideoAvailableChanged));;
}

int QMediaPlayerControl_Volume(void* ptr){
	return static_cast<QMediaPlayerControl*>(ptr)->volume();
}

void QMediaPlayerControl_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::volumeChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_VolumeChanged));;
}

void QMediaPlayerControl_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlayerControl*>(ptr), static_cast<void (QMediaPlayerControl::*)(int)>(&QMediaPlayerControl::volumeChanged), static_cast<MyQMediaPlayerControl*>(ptr), static_cast<void (MyQMediaPlayerControl::*)(int)>(&MyQMediaPlayerControl::Signal_VolumeChanged));;
}

void QMediaPlayerControl_DestroyQMediaPlayerControl(void* ptr){
	static_cast<QMediaPlayerControl*>(ptr)->~QMediaPlayerControl();
}

class MyQMediaPlaylist: public QMediaPlaylist {
public:
	MyQMediaPlaylist(QObject *parent) : QMediaPlaylist(parent) {};
	void Signal_CurrentIndexChanged(int position) { callbackQMediaPlaylistCurrentIndexChanged(this->objectName().toUtf8().data(), position); };
	void Signal_LoadFailed() { callbackQMediaPlaylistLoadFailed(this->objectName().toUtf8().data()); };
	void Signal_Loaded() { callbackQMediaPlaylistLoaded(this->objectName().toUtf8().data()); };
	void Signal_MediaAboutToBeInserted(int start, int end) { callbackQMediaPlaylistMediaAboutToBeInserted(this->objectName().toUtf8().data(), start, end); };
	void Signal_MediaAboutToBeRemoved(int start, int end) { callbackQMediaPlaylistMediaAboutToBeRemoved(this->objectName().toUtf8().data(), start, end); };
	void Signal_MediaChanged(int start, int end) { callbackQMediaPlaylistMediaChanged(this->objectName().toUtf8().data(), start, end); };
	void Signal_MediaInserted(int start, int end) { callbackQMediaPlaylistMediaInserted(this->objectName().toUtf8().data(), start, end); };
	void Signal_MediaRemoved(int start, int end) { callbackQMediaPlaylistMediaRemoved(this->objectName().toUtf8().data(), start, end); };
	void Signal_PlaybackModeChanged(QMediaPlaylist::PlaybackMode mode) { callbackQMediaPlaylistPlaybackModeChanged(this->objectName().toUtf8().data(), mode); };
protected:
};

int QMediaPlaylist_PlaybackMode(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->playbackMode();
}

void QMediaPlaylist_SetPlaybackMode(void* ptr, int mode){
	static_cast<QMediaPlaylist*>(ptr)->setPlaybackMode(static_cast<QMediaPlaylist::PlaybackMode>(mode));
}

void* QMediaPlaylist_NewQMediaPlaylist(void* parent){
	return new MyQMediaPlaylist(static_cast<QObject*>(parent));
}

int QMediaPlaylist_AddMedia(void* ptr, void* content){
	return static_cast<QMediaPlaylist*>(ptr)->addMedia(*static_cast<QMediaContent*>(content));
}

int QMediaPlaylist_Clear(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->clear();
}

int QMediaPlaylist_CurrentIndex(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->currentIndex();
}

void QMediaPlaylist_ConnectCurrentIndexChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int)>(&QMediaPlaylist::currentIndexChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int)>(&MyQMediaPlaylist::Signal_CurrentIndexChanged));;
}

void QMediaPlaylist_DisconnectCurrentIndexChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int)>(&QMediaPlaylist::currentIndexChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int)>(&MyQMediaPlaylist::Signal_CurrentIndexChanged));;
}

int QMediaPlaylist_Error(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->error();
}

char* QMediaPlaylist_ErrorString(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->errorString().toUtf8().data();
}

int QMediaPlaylist_InsertMedia(void* ptr, int pos, void* content){
	return static_cast<QMediaPlaylist*>(ptr)->insertMedia(pos, *static_cast<QMediaContent*>(content));
}

int QMediaPlaylist_IsEmpty(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->isEmpty();
}

int QMediaPlaylist_IsReadOnly(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->isReadOnly();
}

void QMediaPlaylist_Load3(void* ptr, void* device, char* format){
	static_cast<QMediaPlaylist*>(ptr)->load(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

void QMediaPlaylist_Load(void* ptr, void* request, char* format){
	static_cast<QMediaPlaylist*>(ptr)->load(*static_cast<QNetworkRequest*>(request), const_cast<const char*>(format));
}

void QMediaPlaylist_Load2(void* ptr, void* location, char* format){
	static_cast<QMediaPlaylist*>(ptr)->load(*static_cast<QUrl*>(location), const_cast<const char*>(format));
}

void QMediaPlaylist_ConnectLoadFailed(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loadFailed), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_LoadFailed));;
}

void QMediaPlaylist_DisconnectLoadFailed(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loadFailed), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_LoadFailed));;
}

void QMediaPlaylist_ConnectLoaded(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loaded), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_Loaded));;
}

void QMediaPlaylist_DisconnectLoaded(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)()>(&QMediaPlaylist::loaded), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)()>(&MyQMediaPlaylist::Signal_Loaded));;
}

void QMediaPlaylist_ConnectMediaAboutToBeInserted(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeInserted));;
}

void QMediaPlaylist_DisconnectMediaAboutToBeInserted(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeInserted));;
}

void QMediaPlaylist_ConnectMediaAboutToBeRemoved(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeRemoved));;
}

void QMediaPlaylist_DisconnectMediaAboutToBeRemoved(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaAboutToBeRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaAboutToBeRemoved));;
}

void QMediaPlaylist_ConnectMediaChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaChanged));;
}

void QMediaPlaylist_DisconnectMediaChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaChanged));;
}

int QMediaPlaylist_MediaCount(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->mediaCount();
}

void QMediaPlaylist_ConnectMediaInserted(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaInserted));;
}

void QMediaPlaylist_DisconnectMediaInserted(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaInserted), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaInserted));;
}

void* QMediaPlaylist_MediaObject(void* ptr){
	return static_cast<QMediaPlaylist*>(ptr)->mediaObject();
}

void QMediaPlaylist_ConnectMediaRemoved(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaRemoved));;
}

void QMediaPlaylist_DisconnectMediaRemoved(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(int, int)>(&QMediaPlaylist::mediaRemoved), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(int, int)>(&MyQMediaPlaylist::Signal_MediaRemoved));;
}

void QMediaPlaylist_Next(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "next");
}

int QMediaPlaylist_NextIndex(void* ptr, int steps){
	return static_cast<QMediaPlaylist*>(ptr)->nextIndex(steps);
}

void QMediaPlaylist_ConnectPlaybackModeChanged(void* ptr){
	QObject::connect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&QMediaPlaylist::playbackModeChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&MyQMediaPlaylist::Signal_PlaybackModeChanged));;
}

void QMediaPlaylist_DisconnectPlaybackModeChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaPlaylist*>(ptr), static_cast<void (QMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&QMediaPlaylist::playbackModeChanged), static_cast<MyQMediaPlaylist*>(ptr), static_cast<void (MyQMediaPlaylist::*)(QMediaPlaylist::PlaybackMode)>(&MyQMediaPlaylist::Signal_PlaybackModeChanged));;
}

void QMediaPlaylist_Previous(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "previous");
}

int QMediaPlaylist_PreviousIndex(void* ptr, int steps){
	return static_cast<QMediaPlaylist*>(ptr)->previousIndex(steps);
}

int QMediaPlaylist_RemoveMedia(void* ptr, int pos){
	return static_cast<QMediaPlaylist*>(ptr)->removeMedia(pos);
}

int QMediaPlaylist_RemoveMedia2(void* ptr, int start, int end){
	return static_cast<QMediaPlaylist*>(ptr)->removeMedia(start, end);
}

int QMediaPlaylist_Save2(void* ptr, void* device, char* format){
	return static_cast<QMediaPlaylist*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

int QMediaPlaylist_Save(void* ptr, void* location, char* format){
	return static_cast<QMediaPlaylist*>(ptr)->save(*static_cast<QUrl*>(location), const_cast<const char*>(format));
}

void QMediaPlaylist_SetCurrentIndex(void* ptr, int playlistPosition){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "setCurrentIndex", Q_ARG(int, playlistPosition));
}

void QMediaPlaylist_Shuffle(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaPlaylist*>(ptr), "shuffle");
}

void QMediaPlaylist_DestroyQMediaPlaylist(void* ptr){
	static_cast<QMediaPlaylist*>(ptr)->~QMediaPlaylist();
}

class MyQMediaRecorder: public QMediaRecorder {
public:
	MyQMediaRecorder(QMediaObject *mediaObject, QObject *parent) : QMediaRecorder(mediaObject, parent) {};
	void Signal_AvailabilityChanged(bool available) { callbackQMediaRecorderAvailabilityChanged(this->objectName().toUtf8().data(), available); };
	void Signal_MetaDataAvailableChanged(bool available) { callbackQMediaRecorderMetaDataAvailableChanged(this->objectName().toUtf8().data(), available); };
	void Signal_MetaDataChanged() { callbackQMediaRecorderMetaDataChanged(this->objectName().toUtf8().data()); };
	void Signal_MetaDataWritableChanged(bool writable) { callbackQMediaRecorderMetaDataWritableChanged(this->objectName().toUtf8().data(), writable); };
	void Signal_MutedChanged(bool muted) { callbackQMediaRecorderMutedChanged(this->objectName().toUtf8().data(), muted); };
	void Signal_StateChanged(QMediaRecorder::State state) { callbackQMediaRecorderStateChanged(this->objectName().toUtf8().data(), state); };
	void Signal_StatusChanged(QMediaRecorder::Status status) { callbackQMediaRecorderStatusChanged(this->objectName().toUtf8().data(), status); };
protected:
};

int QMediaRecorder_IsMetaDataAvailable(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->isMetaDataAvailable();
}

int QMediaRecorder_IsMetaDataWritable(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->isMetaDataWritable();
}

int QMediaRecorder_IsMuted(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->isMuted();
}

void QMediaRecorder_SetMuted(void* ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

int QMediaRecorder_SetOutputLocation(void* ptr, void* location){
	return static_cast<QMediaRecorder*>(ptr)->setOutputLocation(*static_cast<QUrl*>(location));
}

void QMediaRecorder_SetVolume(void* ptr, double volume){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "setVolume", Q_ARG(qreal, static_cast<qreal>(volume)));
}

double QMediaRecorder_Volume(void* ptr){
	return static_cast<double>(static_cast<QMediaRecorder*>(ptr)->volume());
}

void* QMediaRecorder_NewQMediaRecorder(void* mediaObject, void* parent){
	return new MyQMediaRecorder(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

char* QMediaRecorder_AudioCodecDescription(void* ptr, char* codec){
	return static_cast<QMediaRecorder*>(ptr)->audioCodecDescription(QString(codec)).toUtf8().data();
}

void QMediaRecorder_ConnectAvailabilityChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_AvailabilityChanged));;
}

void QMediaRecorder_DisconnectAvailabilityChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_AvailabilityChanged));;
}

char* QMediaRecorder_AvailableMetaData(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->availableMetaData().join(",,,").toUtf8().data();
}

char* QMediaRecorder_ContainerDescription(void* ptr, char* format){
	return static_cast<QMediaRecorder*>(ptr)->containerDescription(QString(format)).toUtf8().data();
}

char* QMediaRecorder_ContainerFormat(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->containerFormat().toUtf8().data();
}

int QMediaRecorder_Error(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->error();
}

char* QMediaRecorder_ErrorString(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->errorString().toUtf8().data();
}

int QMediaRecorder_IsAvailable(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->isAvailable();
}

void* QMediaRecorder_MediaObject(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->mediaObject();
}

void* QMediaRecorder_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QMediaRecorder*>(ptr)->metaData(QString(key)));
}

void QMediaRecorder_ConnectMetaDataAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataAvailableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataAvailableChanged));;
}

void QMediaRecorder_DisconnectMetaDataAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataAvailableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataAvailableChanged));;
}

void QMediaRecorder_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)()>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)()>(&MyQMediaRecorder::Signal_MetaDataChanged));;
}

void QMediaRecorder_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)()>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)()>(&MyQMediaRecorder::Signal_MetaDataChanged));;
}

void QMediaRecorder_ConnectMetaDataWritableChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataWritableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataWritableChanged));;
}

void QMediaRecorder_DisconnectMetaDataWritableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataWritableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataWritableChanged));;
}

void QMediaRecorder_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::mutedChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MutedChanged));;
}

void QMediaRecorder_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::mutedChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MutedChanged));;
}

void QMediaRecorder_Pause(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "pause");
}

void QMediaRecorder_Record(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "record");
}

void QMediaRecorder_SetAudioSettings(void* ptr, void* settings){
	static_cast<QMediaRecorder*>(ptr)->setAudioSettings(*static_cast<QAudioEncoderSettings*>(settings));
}

void QMediaRecorder_SetContainerFormat(void* ptr, char* container){
	static_cast<QMediaRecorder*>(ptr)->setContainerFormat(QString(container));
}

void QMediaRecorder_SetEncodingSettings(void* ptr, void* audio, void* video, char* container){
	static_cast<QMediaRecorder*>(ptr)->setEncodingSettings(*static_cast<QAudioEncoderSettings*>(audio), *static_cast<QVideoEncoderSettings*>(video), QString(container));
}

void QMediaRecorder_SetMetaData(void* ptr, char* key, void* value){
	static_cast<QMediaRecorder*>(ptr)->setMetaData(QString(key), *static_cast<QVariant*>(value));
}

void QMediaRecorder_SetVideoSettings(void* ptr, void* settings){
	static_cast<QMediaRecorder*>(ptr)->setVideoSettings(*static_cast<QVideoEncoderSettings*>(settings));
}

void QMediaRecorder_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::State)>(&QMediaRecorder::stateChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::State)>(&MyQMediaRecorder::Signal_StateChanged));;
}

void QMediaRecorder_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::State)>(&QMediaRecorder::stateChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::State)>(&MyQMediaRecorder::Signal_StateChanged));;
}

int QMediaRecorder_Status(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->status();
}

void QMediaRecorder_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Status)>(&QMediaRecorder::statusChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Status)>(&MyQMediaRecorder::Signal_StatusChanged));;
}

void QMediaRecorder_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Status)>(&QMediaRecorder::statusChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Status)>(&MyQMediaRecorder::Signal_StatusChanged));;
}

void QMediaRecorder_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "stop");
}

char* QMediaRecorder_SupportedAudioCodecs(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->supportedAudioCodecs().join(",,,").toUtf8().data();
}

char* QMediaRecorder_SupportedContainers(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->supportedContainers().join(",,,").toUtf8().data();
}

char* QMediaRecorder_SupportedVideoCodecs(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->supportedVideoCodecs().join(",,,").toUtf8().data();
}

char* QMediaRecorder_VideoCodecDescription(void* ptr, char* codec){
	return static_cast<QMediaRecorder*>(ptr)->videoCodecDescription(QString(codec)).toUtf8().data();
}

void QMediaRecorder_DestroyQMediaRecorder(void* ptr){
	static_cast<QMediaRecorder*>(ptr)->~QMediaRecorder();
}

class MyQMediaRecorderControl: public QMediaRecorderControl {
public:
	void Signal_Error(int error, const QString & errorString) { callbackQMediaRecorderControlError(this->objectName().toUtf8().data(), error, errorString.toUtf8().data()); };
	void Signal_MutedChanged(bool muted) { callbackQMediaRecorderControlMutedChanged(this->objectName().toUtf8().data(), muted); };
	void Signal_StateChanged(QMediaRecorder::State state) { callbackQMediaRecorderControlStateChanged(this->objectName().toUtf8().data(), state); };
	void Signal_StatusChanged(QMediaRecorder::Status status) { callbackQMediaRecorderControlStatusChanged(this->objectName().toUtf8().data(), status); };
protected:
};

void QMediaRecorderControl_ApplySettings(void* ptr){
	static_cast<QMediaRecorderControl*>(ptr)->applySettings();
}

void QMediaRecorderControl_ConnectError(void* ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(int, const QString &)>(&QMediaRecorderControl::error), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(int, const QString &)>(&MyQMediaRecorderControl::Signal_Error));;
}

void QMediaRecorderControl_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(int, const QString &)>(&QMediaRecorderControl::error), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(int, const QString &)>(&MyQMediaRecorderControl::Signal_Error));;
}

int QMediaRecorderControl_IsMuted(void* ptr){
	return static_cast<QMediaRecorderControl*>(ptr)->isMuted();
}

void QMediaRecorderControl_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(bool)>(&QMediaRecorderControl::mutedChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(bool)>(&MyQMediaRecorderControl::Signal_MutedChanged));;
}

void QMediaRecorderControl_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(bool)>(&QMediaRecorderControl::mutedChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(bool)>(&MyQMediaRecorderControl::Signal_MutedChanged));;
}

void QMediaRecorderControl_SetMuted(void* ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

int QMediaRecorderControl_SetOutputLocation(void* ptr, void* location){
	return static_cast<QMediaRecorderControl*>(ptr)->setOutputLocation(*static_cast<QUrl*>(location));
}

void QMediaRecorderControl_SetState(void* ptr, int state){
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setState", Q_ARG(QMediaRecorder::State, static_cast<QMediaRecorder::State>(state)));
}

void QMediaRecorderControl_SetVolume(void* ptr, double gain){
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setVolume", Q_ARG(qreal, static_cast<qreal>(gain)));
}

void QMediaRecorderControl_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::State)>(&QMediaRecorderControl::stateChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::State)>(&MyQMediaRecorderControl::Signal_StateChanged));;
}

void QMediaRecorderControl_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::State)>(&QMediaRecorderControl::stateChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::State)>(&MyQMediaRecorderControl::Signal_StateChanged));;
}

int QMediaRecorderControl_Status(void* ptr){
	return static_cast<QMediaRecorderControl*>(ptr)->status();
}

void QMediaRecorderControl_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::Status)>(&QMediaRecorderControl::statusChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::Status)>(&MyQMediaRecorderControl::Signal_StatusChanged));;
}

void QMediaRecorderControl_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::Status)>(&QMediaRecorderControl::statusChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::Status)>(&MyQMediaRecorderControl::Signal_StatusChanged));;
}

double QMediaRecorderControl_Volume(void* ptr){
	return static_cast<double>(static_cast<QMediaRecorderControl*>(ptr)->volume());
}

void QMediaRecorderControl_DestroyQMediaRecorderControl(void* ptr){
	static_cast<QMediaRecorderControl*>(ptr)->~QMediaRecorderControl();
}

void* QMediaResource_NewQMediaResource(){
	return new QMediaResource();
}

void* QMediaResource_NewQMediaResource4(void* other){
	return new QMediaResource(*static_cast<QMediaResource*>(other));
}

void* QMediaResource_NewQMediaResource3(void* request, char* mimeType){
	return new QMediaResource(*static_cast<QNetworkRequest*>(request), QString(mimeType));
}

void* QMediaResource_NewQMediaResource2(void* url, char* mimeType){
	return new QMediaResource(*static_cast<QUrl*>(url), QString(mimeType));
}

int QMediaResource_AudioBitRate(void* ptr){
	return static_cast<QMediaResource*>(ptr)->audioBitRate();
}

char* QMediaResource_AudioCodec(void* ptr){
	return static_cast<QMediaResource*>(ptr)->audioCodec().toUtf8().data();
}

int QMediaResource_ChannelCount(void* ptr){
	return static_cast<QMediaResource*>(ptr)->channelCount();
}

int QMediaResource_IsNull(void* ptr){
	return static_cast<QMediaResource*>(ptr)->isNull();
}

char* QMediaResource_Language(void* ptr){
	return static_cast<QMediaResource*>(ptr)->language().toUtf8().data();
}

char* QMediaResource_MimeType(void* ptr){
	return static_cast<QMediaResource*>(ptr)->mimeType().toUtf8().data();
}

int QMediaResource_SampleRate(void* ptr){
	return static_cast<QMediaResource*>(ptr)->sampleRate();
}

void QMediaResource_SetAudioBitRate(void* ptr, int rate){
	static_cast<QMediaResource*>(ptr)->setAudioBitRate(rate);
}

void QMediaResource_SetAudioCodec(void* ptr, char* codec){
	static_cast<QMediaResource*>(ptr)->setAudioCodec(QString(codec));
}

void QMediaResource_SetChannelCount(void* ptr, int channels){
	static_cast<QMediaResource*>(ptr)->setChannelCount(channels);
}

void QMediaResource_SetLanguage(void* ptr, char* language){
	static_cast<QMediaResource*>(ptr)->setLanguage(QString(language));
}

void QMediaResource_SetResolution(void* ptr, void* resolution){
	static_cast<QMediaResource*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QMediaResource_SetResolution2(void* ptr, int width, int height){
	static_cast<QMediaResource*>(ptr)->setResolution(width, height);
}

void QMediaResource_SetSampleRate(void* ptr, int sampleRate){
	static_cast<QMediaResource*>(ptr)->setSampleRate(sampleRate);
}

void QMediaResource_SetVideoBitRate(void* ptr, int rate){
	static_cast<QMediaResource*>(ptr)->setVideoBitRate(rate);
}

void QMediaResource_SetVideoCodec(void* ptr, char* codec){
	static_cast<QMediaResource*>(ptr)->setVideoCodec(QString(codec));
}

int QMediaResource_VideoBitRate(void* ptr){
	return static_cast<QMediaResource*>(ptr)->videoBitRate();
}

char* QMediaResource_VideoCodec(void* ptr){
	return static_cast<QMediaResource*>(ptr)->videoCodec().toUtf8().data();
}

void QMediaResource_DestroyQMediaResource(void* ptr){
	static_cast<QMediaResource*>(ptr)->~QMediaResource();
}

void QMediaService_ReleaseControl(void* ptr, void* control){
	static_cast<QMediaService*>(ptr)->releaseControl(static_cast<QMediaControl*>(control));
}

void* QMediaService_RequestControl(void* ptr, char* interfa){
	return static_cast<QMediaService*>(ptr)->requestControl(const_cast<const char*>(interfa));
}

void* QMediaService_RequestControl2(void* ptr){
	return static_cast<QMediaService*>(ptr)->requestControl<QMediaControl*>();
}

void QMediaService_DestroyQMediaService(void* ptr){
	static_cast<QMediaService*>(ptr)->~QMediaService();
}

class MyQMediaServiceCameraInfoInterface: public QMediaServiceCameraInfoInterface {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

int QMediaServiceCameraInfoInterface_CameraOrientation(void* ptr, void* device){
	return static_cast<QMediaServiceCameraInfoInterface*>(ptr)->cameraOrientation(*static_cast<QByteArray*>(device));
}

int QMediaServiceCameraInfoInterface_CameraPosition(void* ptr, void* device){
	return static_cast<QMediaServiceCameraInfoInterface*>(ptr)->cameraPosition(*static_cast<QByteArray*>(device));
}

void QMediaServiceCameraInfoInterface_DestroyQMediaServiceCameraInfoInterface(void* ptr){
	static_cast<QMediaServiceCameraInfoInterface*>(ptr)->~QMediaServiceCameraInfoInterface();
}

char* QMediaServiceCameraInfoInterface_ObjectNameAbs(void* ptr){
	return static_cast<MyQMediaServiceCameraInfoInterface*>(ptr)->objectNameAbs().toUtf8().data();
}

void QMediaServiceCameraInfoInterface_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQMediaServiceCameraInfoInterface*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQMediaServiceDefaultDeviceInterface: public QMediaServiceDefaultDeviceInterface {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

void* QMediaServiceDefaultDeviceInterface_DefaultDevice(void* ptr, void* service){
	return new QByteArray(static_cast<QMediaServiceDefaultDeviceInterface*>(ptr)->defaultDevice(*static_cast<QByteArray*>(service)));
}

void QMediaServiceDefaultDeviceInterface_DestroyQMediaServiceDefaultDeviceInterface(void* ptr){
	static_cast<QMediaServiceDefaultDeviceInterface*>(ptr)->~QMediaServiceDefaultDeviceInterface();
}

char* QMediaServiceDefaultDeviceInterface_ObjectNameAbs(void* ptr){
	return static_cast<MyQMediaServiceDefaultDeviceInterface*>(ptr)->objectNameAbs().toUtf8().data();
}

void QMediaServiceDefaultDeviceInterface_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQMediaServiceDefaultDeviceInterface*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQMediaServiceFeaturesInterface: public QMediaServiceFeaturesInterface {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

void QMediaServiceFeaturesInterface_DestroyQMediaServiceFeaturesInterface(void* ptr){
	static_cast<QMediaServiceFeaturesInterface*>(ptr)->~QMediaServiceFeaturesInterface();
}

char* QMediaServiceFeaturesInterface_ObjectNameAbs(void* ptr){
	return static_cast<MyQMediaServiceFeaturesInterface*>(ptr)->objectNameAbs().toUtf8().data();
}

void QMediaServiceFeaturesInterface_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQMediaServiceFeaturesInterface*>(ptr)->setObjectNameAbs(QString(name));
}

void* QMediaServiceProviderPlugin_Create(void* ptr, char* key){
	return static_cast<QMediaServiceProviderPlugin*>(ptr)->create(QString(key));
}

void QMediaServiceProviderPlugin_Release(void* ptr, void* service){
	static_cast<QMediaServiceProviderPlugin*>(ptr)->release(static_cast<QMediaService*>(service));
}

class MyQMediaServiceSupportedDevicesInterface: public QMediaServiceSupportedDevicesInterface {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

char* QMediaServiceSupportedDevicesInterface_DeviceDescription(void* ptr, void* service, void* device){
	return static_cast<QMediaServiceSupportedDevicesInterface*>(ptr)->deviceDescription(*static_cast<QByteArray*>(service), *static_cast<QByteArray*>(device)).toUtf8().data();
}

void QMediaServiceSupportedDevicesInterface_DestroyQMediaServiceSupportedDevicesInterface(void* ptr){
	static_cast<QMediaServiceSupportedDevicesInterface*>(ptr)->~QMediaServiceSupportedDevicesInterface();
}

char* QMediaServiceSupportedDevicesInterface_ObjectNameAbs(void* ptr){
	return static_cast<MyQMediaServiceSupportedDevicesInterface*>(ptr)->objectNameAbs().toUtf8().data();
}

void QMediaServiceSupportedDevicesInterface_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQMediaServiceSupportedDevicesInterface*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQMediaServiceSupportedFormatsInterface: public QMediaServiceSupportedFormatsInterface {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

char* QMediaServiceSupportedFormatsInterface_SupportedMimeTypes(void* ptr){
	return static_cast<QMediaServiceSupportedFormatsInterface*>(ptr)->supportedMimeTypes().join(",,,").toUtf8().data();
}

void QMediaServiceSupportedFormatsInterface_DestroyQMediaServiceSupportedFormatsInterface(void* ptr){
	static_cast<QMediaServiceSupportedFormatsInterface*>(ptr)->~QMediaServiceSupportedFormatsInterface();
}

char* QMediaServiceSupportedFormatsInterface_ObjectNameAbs(void* ptr){
	return static_cast<MyQMediaServiceSupportedFormatsInterface*>(ptr)->objectNameAbs().toUtf8().data();
}

void QMediaServiceSupportedFormatsInterface_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQMediaServiceSupportedFormatsInterface*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQMediaStreamsControl: public QMediaStreamsControl {
public:
	void Signal_ActiveStreamsChanged() { callbackQMediaStreamsControlActiveStreamsChanged(this->objectName().toUtf8().data()); };
	void Signal_StreamsChanged() { callbackQMediaStreamsControlStreamsChanged(this->objectName().toUtf8().data()); };
protected:
};

void QMediaStreamsControl_ConnectActiveStreamsChanged(void* ptr){
	QObject::connect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::activeStreamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_ActiveStreamsChanged));;
}

void QMediaStreamsControl_DisconnectActiveStreamsChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::activeStreamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_ActiveStreamsChanged));;
}

int QMediaStreamsControl_IsActive(void* ptr, int stream){
	return static_cast<QMediaStreamsControl*>(ptr)->isActive(stream);
}

void* QMediaStreamsControl_MetaData(void* ptr, int stream, char* key){
	return new QVariant(static_cast<QMediaStreamsControl*>(ptr)->metaData(stream, QString(key)));
}

void QMediaStreamsControl_SetActive(void* ptr, int stream, int state){
	static_cast<QMediaStreamsControl*>(ptr)->setActive(stream, state != 0);
}

int QMediaStreamsControl_StreamCount(void* ptr){
	return static_cast<QMediaStreamsControl*>(ptr)->streamCount();
}

int QMediaStreamsControl_StreamType(void* ptr, int stream){
	return static_cast<QMediaStreamsControl*>(ptr)->streamType(stream);
}

void QMediaStreamsControl_ConnectStreamsChanged(void* ptr){
	QObject::connect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::streamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_StreamsChanged));;
}

void QMediaStreamsControl_DisconnectStreamsChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::streamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_StreamsChanged));;
}

void QMediaStreamsControl_DestroyQMediaStreamsControl(void* ptr){
	static_cast<QMediaStreamsControl*>(ptr)->~QMediaStreamsControl();
}

void* QMediaTimeInterval_NewQMediaTimeInterval(){
	return new QMediaTimeInterval();
}

void* QMediaTimeInterval_NewQMediaTimeInterval3(void* other){
	return new QMediaTimeInterval(*static_cast<QMediaTimeInterval*>(other));
}

int QMediaTimeInterval_IsNormal(void* ptr){
	return static_cast<QMediaTimeInterval*>(ptr)->isNormal();
}

void* QMediaTimeRange_NewQMediaTimeRange(){
	return new QMediaTimeRange();
}

void* QMediaTimeRange_NewQMediaTimeRange3(void* interval){
	return new QMediaTimeRange(*static_cast<QMediaTimeInterval*>(interval));
}

void* QMediaTimeRange_NewQMediaTimeRange4(void* ran){
	return new QMediaTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_AddInterval(void* ptr, void* interval){
	static_cast<QMediaTimeRange*>(ptr)->addInterval(*static_cast<QMediaTimeInterval*>(interval));
}

void QMediaTimeRange_AddTimeRange(void* ptr, void* ran){
	static_cast<QMediaTimeRange*>(ptr)->addTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_Clear(void* ptr){
	static_cast<QMediaTimeRange*>(ptr)->clear();
}

int QMediaTimeRange_IsContinuous(void* ptr){
	return static_cast<QMediaTimeRange*>(ptr)->isContinuous();
}

int QMediaTimeRange_IsEmpty(void* ptr){
	return static_cast<QMediaTimeRange*>(ptr)->isEmpty();
}

void QMediaTimeRange_RemoveInterval(void* ptr, void* interval){
	static_cast<QMediaTimeRange*>(ptr)->removeInterval(*static_cast<QMediaTimeInterval*>(interval));
}

void QMediaTimeRange_RemoveTimeRange(void* ptr, void* ran){
	static_cast<QMediaTimeRange*>(ptr)->removeTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_DestroyQMediaTimeRange(void* ptr){
	static_cast<QMediaTimeRange*>(ptr)->~QMediaTimeRange();
}

class MyQMediaVideoProbeControl: public QMediaVideoProbeControl {
public:
	void Signal_Flush() { callbackQMediaVideoProbeControlFlush(this->objectName().toUtf8().data()); };
protected:
};

void QMediaVideoProbeControl_ConnectFlush(void* ptr){
	QObject::connect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)()>(&QMediaVideoProbeControl::flush), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)()>(&MyQMediaVideoProbeControl::Signal_Flush));;
}

void QMediaVideoProbeControl_DisconnectFlush(void* ptr){
	QObject::disconnect(static_cast<QMediaVideoProbeControl*>(ptr), static_cast<void (QMediaVideoProbeControl::*)()>(&QMediaVideoProbeControl::flush), static_cast<MyQMediaVideoProbeControl*>(ptr), static_cast<void (MyQMediaVideoProbeControl::*)()>(&MyQMediaVideoProbeControl::Signal_Flush));;
}

void QMediaVideoProbeControl_DestroyQMediaVideoProbeControl(void* ptr){
	static_cast<QMediaVideoProbeControl*>(ptr)->~QMediaVideoProbeControl();
}

class MyQMetaDataReaderControl: public QMetaDataReaderControl {
public:
	void Signal_MetaDataAvailableChanged(bool available) { callbackQMetaDataReaderControlMetaDataAvailableChanged(this->objectName().toUtf8().data(), available); };
	void Signal_MetaDataChanged() { callbackQMetaDataReaderControlMetaDataChanged(this->objectName().toUtf8().data()); };
protected:
};

char* QMetaDataReaderControl_AvailableMetaData(void* ptr){
	return static_cast<QMetaDataReaderControl*>(ptr)->availableMetaData().join(",,,").toUtf8().data();
}

int QMetaDataReaderControl_IsMetaDataAvailable(void* ptr){
	return static_cast<QMetaDataReaderControl*>(ptr)->isMetaDataAvailable();
}

void* QMetaDataReaderControl_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QMetaDataReaderControl*>(ptr)->metaData(QString(key)));
}

void QMetaDataReaderControl_ConnectMetaDataAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(bool)>(&QMetaDataReaderControl::metaDataAvailableChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(bool)>(&MyQMetaDataReaderControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataReaderControl_DisconnectMetaDataAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(bool)>(&QMetaDataReaderControl::metaDataAvailableChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(bool)>(&MyQMetaDataReaderControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataReaderControl_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)()>(&MyQMetaDataReaderControl::Signal_MetaDataChanged));;
}

void QMetaDataReaderControl_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)()>(&MyQMetaDataReaderControl::Signal_MetaDataChanged));;
}

void QMetaDataReaderControl_DestroyQMetaDataReaderControl(void* ptr){
	static_cast<QMetaDataReaderControl*>(ptr)->~QMetaDataReaderControl();
}

class MyQMetaDataWriterControl: public QMetaDataWriterControl {
public:
	void Signal_MetaDataAvailableChanged(bool available) { callbackQMetaDataWriterControlMetaDataAvailableChanged(this->objectName().toUtf8().data(), available); };
	void Signal_MetaDataChanged() { callbackQMetaDataWriterControlMetaDataChanged(this->objectName().toUtf8().data()); };
	void Signal_WritableChanged(bool writable) { callbackQMetaDataWriterControlWritableChanged(this->objectName().toUtf8().data(), writable); };
protected:
};

char* QMetaDataWriterControl_AvailableMetaData(void* ptr){
	return static_cast<QMetaDataWriterControl*>(ptr)->availableMetaData().join(",,,").toUtf8().data();
}

int QMetaDataWriterControl_IsMetaDataAvailable(void* ptr){
	return static_cast<QMetaDataWriterControl*>(ptr)->isMetaDataAvailable();
}

int QMetaDataWriterControl_IsWritable(void* ptr){
	return static_cast<QMetaDataWriterControl*>(ptr)->isWritable();
}

void* QMetaDataWriterControl_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QMetaDataWriterControl*>(ptr)->metaData(QString(key)));
}

void QMetaDataWriterControl_ConnectMetaDataAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::metaDataAvailableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataWriterControl_DisconnectMetaDataAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::metaDataAvailableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataWriterControl_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)()>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)()>(&MyQMetaDataWriterControl::Signal_MetaDataChanged));;
}

void QMetaDataWriterControl_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)()>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)()>(&MyQMetaDataWriterControl::Signal_MetaDataChanged));;
}

void QMetaDataWriterControl_SetMetaData(void* ptr, char* key, void* value){
	static_cast<QMetaDataWriterControl*>(ptr)->setMetaData(QString(key), *static_cast<QVariant*>(value));
}

void QMetaDataWriterControl_ConnectWritableChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::writableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_WritableChanged));;
}

void QMetaDataWriterControl_DisconnectWritableChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::writableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_WritableChanged));;
}

void QMetaDataWriterControl_DestroyQMetaDataWriterControl(void* ptr){
	static_cast<QMetaDataWriterControl*>(ptr)->~QMetaDataWriterControl();
}

class MyQRadioData: public QRadioData {
public:
	MyQRadioData(QMediaObject *mediaObject, QObject *parent) : QRadioData(mediaObject, parent) {};
	void Signal_AlternativeFrequenciesEnabledChanged(bool enabled) { callbackQRadioDataAlternativeFrequenciesEnabledChanged(this->objectName().toUtf8().data(), enabled); };
	void Signal_ProgramTypeChanged(QRadioData::ProgramType programType) { callbackQRadioDataProgramTypeChanged(this->objectName().toUtf8().data(), programType); };
	void Signal_ProgramTypeNameChanged(QString programTypeName) { callbackQRadioDataProgramTypeNameChanged(this->objectName().toUtf8().data(), programTypeName.toUtf8().data()); };
	void Signal_RadioTextChanged(QString radioText) { callbackQRadioDataRadioTextChanged(this->objectName().toUtf8().data(), radioText.toUtf8().data()); };
	void Signal_StationIdChanged(QString stationId) { callbackQRadioDataStationIdChanged(this->objectName().toUtf8().data(), stationId.toUtf8().data()); };
	void Signal_StationNameChanged(QString stationName) { callbackQRadioDataStationNameChanged(this->objectName().toUtf8().data(), stationName.toUtf8().data()); };
protected:
};

int QRadioData_IsAlternativeFrequenciesEnabled(void* ptr){
	return static_cast<QRadioData*>(ptr)->isAlternativeFrequenciesEnabled();
}

int QRadioData_ProgramType(void* ptr){
	return static_cast<QRadioData*>(ptr)->programType();
}

char* QRadioData_ProgramTypeName(void* ptr){
	return static_cast<QRadioData*>(ptr)->programTypeName().toUtf8().data();
}

char* QRadioData_RadioText(void* ptr){
	return static_cast<QRadioData*>(ptr)->radioText().toUtf8().data();
}

void QRadioData_SetAlternativeFrequenciesEnabled(void* ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QRadioData*>(ptr), "setAlternativeFrequenciesEnabled", Q_ARG(bool, enabled != 0));
}

char* QRadioData_StationId(void* ptr){
	return static_cast<QRadioData*>(ptr)->stationId().toUtf8().data();
}

char* QRadioData_StationName(void* ptr){
	return static_cast<QRadioData*>(ptr)->stationName().toUtf8().data();
}

void* QRadioData_NewQRadioData(void* mediaObject, void* parent){
	return new MyQRadioData(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

void QRadioData_ConnectAlternativeFrequenciesEnabledChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(bool)>(&QRadioData::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(bool)>(&MyQRadioData::Signal_AlternativeFrequenciesEnabledChanged));;
}

void QRadioData_DisconnectAlternativeFrequenciesEnabledChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(bool)>(&QRadioData::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(bool)>(&MyQRadioData::Signal_AlternativeFrequenciesEnabledChanged));;
}

int QRadioData_Error(void* ptr){
	return static_cast<QRadioData*>(ptr)->error();
}

char* QRadioData_ErrorString(void* ptr){
	return static_cast<QRadioData*>(ptr)->errorString().toUtf8().data();
}

void* QRadioData_MediaObject(void* ptr){
	return static_cast<QRadioData*>(ptr)->mediaObject();
}

void QRadioData_ConnectProgramTypeChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QRadioData::ProgramType)>(&QRadioData::programTypeChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QRadioData::ProgramType)>(&MyQRadioData::Signal_ProgramTypeChanged));;
}

void QRadioData_DisconnectProgramTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QRadioData::ProgramType)>(&QRadioData::programTypeChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QRadioData::ProgramType)>(&MyQRadioData::Signal_ProgramTypeChanged));;
}

void QRadioData_ConnectProgramTypeNameChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::programTypeNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_ProgramTypeNameChanged));;
}

void QRadioData_DisconnectProgramTypeNameChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::programTypeNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_ProgramTypeNameChanged));;
}

void QRadioData_ConnectRadioTextChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::radioTextChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_RadioTextChanged));;
}

void QRadioData_DisconnectRadioTextChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::radioTextChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_RadioTextChanged));;
}

void QRadioData_ConnectStationIdChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationIdChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationIdChanged));;
}

void QRadioData_DisconnectStationIdChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationIdChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationIdChanged));;
}

void QRadioData_ConnectStationNameChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationNameChanged));;
}

void QRadioData_DisconnectStationNameChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationNameChanged));;
}

void QRadioData_DestroyQRadioData(void* ptr){
	static_cast<QRadioData*>(ptr)->~QRadioData();
}

class MyQRadioDataControl: public QRadioDataControl {
public:
	void Signal_AlternativeFrequenciesEnabledChanged(bool enabled) { callbackQRadioDataControlAlternativeFrequenciesEnabledChanged(this->objectName().toUtf8().data(), enabled); };
	void Signal_ProgramTypeChanged(QRadioData::ProgramType programType) { callbackQRadioDataControlProgramTypeChanged(this->objectName().toUtf8().data(), programType); };
	void Signal_ProgramTypeNameChanged(QString programTypeName) { callbackQRadioDataControlProgramTypeNameChanged(this->objectName().toUtf8().data(), programTypeName.toUtf8().data()); };
	void Signal_RadioTextChanged(QString radioText) { callbackQRadioDataControlRadioTextChanged(this->objectName().toUtf8().data(), radioText.toUtf8().data()); };
	void Signal_StationIdChanged(QString stationId) { callbackQRadioDataControlStationIdChanged(this->objectName().toUtf8().data(), stationId.toUtf8().data()); };
	void Signal_StationNameChanged(QString stationName) { callbackQRadioDataControlStationNameChanged(this->objectName().toUtf8().data(), stationName.toUtf8().data()); };
protected:
};

void QRadioDataControl_ConnectAlternativeFrequenciesEnabledChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(bool)>(&QRadioDataControl::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(bool)>(&MyQRadioDataControl::Signal_AlternativeFrequenciesEnabledChanged));;
}

void QRadioDataControl_DisconnectAlternativeFrequenciesEnabledChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(bool)>(&QRadioDataControl::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(bool)>(&MyQRadioDataControl::Signal_AlternativeFrequenciesEnabledChanged));;
}

int QRadioDataControl_Error(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->error();
}

char* QRadioDataControl_ErrorString(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->errorString().toUtf8().data();
}

int QRadioDataControl_IsAlternativeFrequenciesEnabled(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->isAlternativeFrequenciesEnabled();
}

int QRadioDataControl_ProgramType(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->programType();
}

void QRadioDataControl_ConnectProgramTypeChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QRadioData::ProgramType)>(&QRadioDataControl::programTypeChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QRadioData::ProgramType)>(&MyQRadioDataControl::Signal_ProgramTypeChanged));;
}

void QRadioDataControl_DisconnectProgramTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QRadioData::ProgramType)>(&QRadioDataControl::programTypeChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QRadioData::ProgramType)>(&MyQRadioDataControl::Signal_ProgramTypeChanged));;
}

char* QRadioDataControl_ProgramTypeName(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->programTypeName().toUtf8().data();
}

void QRadioDataControl_ConnectProgramTypeNameChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::programTypeNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_ProgramTypeNameChanged));;
}

void QRadioDataControl_DisconnectProgramTypeNameChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::programTypeNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_ProgramTypeNameChanged));;
}

char* QRadioDataControl_RadioText(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->radioText().toUtf8().data();
}

void QRadioDataControl_ConnectRadioTextChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::radioTextChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_RadioTextChanged));;
}

void QRadioDataControl_DisconnectRadioTextChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::radioTextChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_RadioTextChanged));;
}

void QRadioDataControl_SetAlternativeFrequenciesEnabled(void* ptr, int enabled){
	static_cast<QRadioDataControl*>(ptr)->setAlternativeFrequenciesEnabled(enabled != 0);
}

char* QRadioDataControl_StationId(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->stationId().toUtf8().data();
}

void QRadioDataControl_ConnectStationIdChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationIdChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationIdChanged));;
}

void QRadioDataControl_DisconnectStationIdChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationIdChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationIdChanged));;
}

char* QRadioDataControl_StationName(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->stationName().toUtf8().data();
}

void QRadioDataControl_ConnectStationNameChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationNameChanged));;
}

void QRadioDataControl_DisconnectStationNameChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationNameChanged));;
}

void QRadioDataControl_DestroyQRadioDataControl(void* ptr){
	static_cast<QRadioDataControl*>(ptr)->~QRadioDataControl();
}

class MyQRadioTuner: public QRadioTuner {
public:
	MyQRadioTuner(QObject *parent) : QRadioTuner(parent) {};
	void Signal_AntennaConnectedChanged(bool connectionStatus) { callbackQRadioTunerAntennaConnectedChanged(this->objectName().toUtf8().data(), connectionStatus); };
	void Signal_BandChanged(QRadioTuner::Band band) { callbackQRadioTunerBandChanged(this->objectName().toUtf8().data(), band); };
	void Signal_FrequencyChanged(int frequency) { callbackQRadioTunerFrequencyChanged(this->objectName().toUtf8().data(), frequency); };
	void Signal_MutedChanged(bool muted) { callbackQRadioTunerMutedChanged(this->objectName().toUtf8().data(), muted); };
	void Signal_SearchingChanged(bool searching) { callbackQRadioTunerSearchingChanged(this->objectName().toUtf8().data(), searching); };
	void Signal_SignalStrengthChanged(int strength) { callbackQRadioTunerSignalStrengthChanged(this->objectName().toUtf8().data(), strength); };
	void Signal_StateChanged(QRadioTuner::State state) { callbackQRadioTunerStateChanged(this->objectName().toUtf8().data(), state); };
	void Signal_StationFound(int frequency, QString stationId) { callbackQRadioTunerStationFound(this->objectName().toUtf8().data(), frequency, stationId.toUtf8().data()); };
	void Signal_StereoStatusChanged(bool stereo) { callbackQRadioTunerStereoStatusChanged(this->objectName().toUtf8().data(), stereo); };
	void Signal_VolumeChanged(int volume) { callbackQRadioTunerVolumeChanged(this->objectName().toUtf8().data(), volume); };
protected:
};

int QRadioTuner_Band(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->band();
}

int QRadioTuner_Frequency(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->frequency();
}

int QRadioTuner_IsAntennaConnected(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->isAntennaConnected();
}

int QRadioTuner_IsMuted(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->isMuted();
}

int QRadioTuner_IsSearching(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->isSearching();
}

int QRadioTuner_IsStereo(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->isStereo();
}

void* QRadioTuner_RadioData(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->radioData();
}

void QRadioTuner_SetMuted(void* ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

void QRadioTuner_SetStereoMode(void* ptr, int mode){
	static_cast<QRadioTuner*>(ptr)->setStereoMode(static_cast<QRadioTuner::StereoMode>(mode));
}

void QRadioTuner_SetVolume(void* ptr, int volume){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setVolume", Q_ARG(int, volume));
}

int QRadioTuner_SignalStrength(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->signalStrength();
}

int QRadioTuner_StereoMode(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->stereoMode();
}

int QRadioTuner_Volume(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->volume();
}

void* QRadioTuner_NewQRadioTuner(void* parent){
	return new MyQRadioTuner(static_cast<QObject*>(parent));
}

void QRadioTuner_ConnectAntennaConnectedChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::antennaConnectedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_AntennaConnectedChanged));;
}

void QRadioTuner_DisconnectAntennaConnectedChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::antennaConnectedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_AntennaConnectedChanged));;
}

void QRadioTuner_ConnectBandChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Band)>(&QRadioTuner::bandChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Band)>(&MyQRadioTuner::Signal_BandChanged));;
}

void QRadioTuner_DisconnectBandChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Band)>(&QRadioTuner::bandChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Band)>(&MyQRadioTuner::Signal_BandChanged));;
}

void QRadioTuner_CancelSearch(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "cancelSearch");
}

int QRadioTuner_Error(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->error();
}

char* QRadioTuner_ErrorString(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->errorString().toUtf8().data();
}

void QRadioTuner_ConnectFrequencyChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::frequencyChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_FrequencyChanged));;
}

void QRadioTuner_DisconnectFrequencyChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::frequencyChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_FrequencyChanged));;
}

int QRadioTuner_FrequencyStep(void* ptr, int band){
	return static_cast<QRadioTuner*>(ptr)->frequencyStep(static_cast<QRadioTuner::Band>(band));
}

int QRadioTuner_IsBandSupported(void* ptr, int band){
	return static_cast<QRadioTuner*>(ptr)->isBandSupported(static_cast<QRadioTuner::Band>(band));
}

void QRadioTuner_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::mutedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_MutedChanged));;
}

void QRadioTuner_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::mutedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_MutedChanged));;
}

void QRadioTuner_SearchAllStations(void* ptr, int searchMode){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchAllStations", Q_ARG(QRadioTuner::SearchMode, static_cast<QRadioTuner::SearchMode>(searchMode)));
}

void QRadioTuner_SearchBackward(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchBackward");
}

void QRadioTuner_SearchForward(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchForward");
}

void QRadioTuner_ConnectSearchingChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::searchingChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_SearchingChanged));;
}

void QRadioTuner_DisconnectSearchingChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::searchingChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_SearchingChanged));;
}

void QRadioTuner_SetBand(void* ptr, int band){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setBand", Q_ARG(QRadioTuner::Band, static_cast<QRadioTuner::Band>(band)));
}

void QRadioTuner_SetFrequency(void* ptr, int frequency){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setFrequency", Q_ARG(int, frequency));
}

void QRadioTuner_ConnectSignalStrengthChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::signalStrengthChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_SignalStrengthChanged));;
}

void QRadioTuner_DisconnectSignalStrengthChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::signalStrengthChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_SignalStrengthChanged));;
}

void QRadioTuner_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "start");
}

void QRadioTuner_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::State)>(&QRadioTuner::stateChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::State)>(&MyQRadioTuner::Signal_StateChanged));;
}

void QRadioTuner_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::State)>(&QRadioTuner::stateChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::State)>(&MyQRadioTuner::Signal_StateChanged));;
}

void QRadioTuner_ConnectStationFound(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int, QString)>(&QRadioTuner::stationFound), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int, QString)>(&MyQRadioTuner::Signal_StationFound));;
}

void QRadioTuner_DisconnectStationFound(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int, QString)>(&QRadioTuner::stationFound), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int, QString)>(&MyQRadioTuner::Signal_StationFound));;
}

void QRadioTuner_ConnectStereoStatusChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::stereoStatusChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_StereoStatusChanged));;
}

void QRadioTuner_DisconnectStereoStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::stereoStatusChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_StereoStatusChanged));;
}

void QRadioTuner_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "stop");
}

void QRadioTuner_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::volumeChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_VolumeChanged));;
}

void QRadioTuner_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::volumeChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_VolumeChanged));;
}

void QRadioTuner_DestroyQRadioTuner(void* ptr){
	static_cast<QRadioTuner*>(ptr)->~QRadioTuner();
}

class MyQRadioTunerControl: public QRadioTunerControl {
public:
	void Signal_AntennaConnectedChanged(bool connectionStatus) { callbackQRadioTunerControlAntennaConnectedChanged(this->objectName().toUtf8().data(), connectionStatus); };
	void Signal_BandChanged(QRadioTuner::Band band) { callbackQRadioTunerControlBandChanged(this->objectName().toUtf8().data(), band); };
	void Signal_FrequencyChanged(int frequency) { callbackQRadioTunerControlFrequencyChanged(this->objectName().toUtf8().data(), frequency); };
	void Signal_MutedChanged(bool muted) { callbackQRadioTunerControlMutedChanged(this->objectName().toUtf8().data(), muted); };
	void Signal_SearchingChanged(bool searching) { callbackQRadioTunerControlSearchingChanged(this->objectName().toUtf8().data(), searching); };
	void Signal_SignalStrengthChanged(int strength) { callbackQRadioTunerControlSignalStrengthChanged(this->objectName().toUtf8().data(), strength); };
	void Signal_StateChanged(QRadioTuner::State state) { callbackQRadioTunerControlStateChanged(this->objectName().toUtf8().data(), state); };
	void Signal_StationFound(int frequency, QString stationId) { callbackQRadioTunerControlStationFound(this->objectName().toUtf8().data(), frequency, stationId.toUtf8().data()); };
	void Signal_StereoStatusChanged(bool stereo) { callbackQRadioTunerControlStereoStatusChanged(this->objectName().toUtf8().data(), stereo); };
	void Signal_VolumeChanged(int volume) { callbackQRadioTunerControlVolumeChanged(this->objectName().toUtf8().data(), volume); };
protected:
};

void QRadioTunerControl_ConnectAntennaConnectedChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::antennaConnectedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_AntennaConnectedChanged));;
}

void QRadioTunerControl_DisconnectAntennaConnectedChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::antennaConnectedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_AntennaConnectedChanged));;
}

int QRadioTunerControl_Band(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->band();
}

void QRadioTunerControl_ConnectBandChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::Band)>(&QRadioTunerControl::bandChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::Band)>(&MyQRadioTunerControl::Signal_BandChanged));;
}

void QRadioTunerControl_DisconnectBandChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::Band)>(&QRadioTunerControl::bandChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::Band)>(&MyQRadioTunerControl::Signal_BandChanged));;
}

void QRadioTunerControl_CancelSearch(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->cancelSearch();
}

int QRadioTunerControl_Error(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->error();
}

char* QRadioTunerControl_ErrorString(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->errorString().toUtf8().data();
}

int QRadioTunerControl_Frequency(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->frequency();
}

void QRadioTunerControl_ConnectFrequencyChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::frequencyChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_FrequencyChanged));;
}

void QRadioTunerControl_DisconnectFrequencyChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::frequencyChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_FrequencyChanged));;
}

int QRadioTunerControl_FrequencyStep(void* ptr, int band){
	return static_cast<QRadioTunerControl*>(ptr)->frequencyStep(static_cast<QRadioTuner::Band>(band));
}

int QRadioTunerControl_IsAntennaConnected(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->isAntennaConnected();
}

int QRadioTunerControl_IsBandSupported(void* ptr, int band){
	return static_cast<QRadioTunerControl*>(ptr)->isBandSupported(static_cast<QRadioTuner::Band>(band));
}

int QRadioTunerControl_IsMuted(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->isMuted();
}

int QRadioTunerControl_IsSearching(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->isSearching();
}

int QRadioTunerControl_IsStereo(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->isStereo();
}

void QRadioTunerControl_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::mutedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_MutedChanged));;
}

void QRadioTunerControl_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::mutedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_MutedChanged));;
}

void QRadioTunerControl_SearchAllStations(void* ptr, int searchMode){
	static_cast<QRadioTunerControl*>(ptr)->searchAllStations(static_cast<QRadioTuner::SearchMode>(searchMode));
}

void QRadioTunerControl_SearchBackward(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->searchBackward();
}

void QRadioTunerControl_SearchForward(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->searchForward();
}

void QRadioTunerControl_ConnectSearchingChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::searchingChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_SearchingChanged));;
}

void QRadioTunerControl_DisconnectSearchingChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::searchingChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_SearchingChanged));;
}

void QRadioTunerControl_SetBand(void* ptr, int band){
	static_cast<QRadioTunerControl*>(ptr)->setBand(static_cast<QRadioTuner::Band>(band));
}

void QRadioTunerControl_SetFrequency(void* ptr, int frequency){
	static_cast<QRadioTunerControl*>(ptr)->setFrequency(frequency);
}

void QRadioTunerControl_SetMuted(void* ptr, int muted){
	static_cast<QRadioTunerControl*>(ptr)->setMuted(muted != 0);
}

void QRadioTunerControl_SetStereoMode(void* ptr, int mode){
	static_cast<QRadioTunerControl*>(ptr)->setStereoMode(static_cast<QRadioTuner::StereoMode>(mode));
}

void QRadioTunerControl_SetVolume(void* ptr, int volume){
	static_cast<QRadioTunerControl*>(ptr)->setVolume(volume);
}

int QRadioTunerControl_SignalStrength(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->signalStrength();
}

void QRadioTunerControl_ConnectSignalStrengthChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::signalStrengthChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_SignalStrengthChanged));;
}

void QRadioTunerControl_DisconnectSignalStrengthChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::signalStrengthChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_SignalStrengthChanged));;
}

void QRadioTunerControl_Start(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->start();
}

void QRadioTunerControl_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::State)>(&QRadioTunerControl::stateChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::State)>(&MyQRadioTunerControl::Signal_StateChanged));;
}

void QRadioTunerControl_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::State)>(&QRadioTunerControl::stateChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::State)>(&MyQRadioTunerControl::Signal_StateChanged));;
}

void QRadioTunerControl_ConnectStationFound(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int, QString)>(&QRadioTunerControl::stationFound), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int, QString)>(&MyQRadioTunerControl::Signal_StationFound));;
}

void QRadioTunerControl_DisconnectStationFound(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int, QString)>(&QRadioTunerControl::stationFound), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int, QString)>(&MyQRadioTunerControl::Signal_StationFound));;
}

int QRadioTunerControl_StereoMode(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->stereoMode();
}

void QRadioTunerControl_ConnectStereoStatusChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::stereoStatusChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_StereoStatusChanged));;
}

void QRadioTunerControl_DisconnectStereoStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::stereoStatusChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_StereoStatusChanged));;
}

void QRadioTunerControl_Stop(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->stop();
}

int QRadioTunerControl_Volume(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->volume();
}

void QRadioTunerControl_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::volumeChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_VolumeChanged));;
}

void QRadioTunerControl_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::volumeChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_VolumeChanged));;
}

void QRadioTunerControl_DestroyQRadioTunerControl(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->~QRadioTunerControl();
}

void QSound_SetLoops(void* ptr, int number){
	static_cast<QSound*>(ptr)->setLoops(number);
}

void* QSound_NewQSound(char* filename, void* parent){
	return new QSound(QString(filename), static_cast<QObject*>(parent));
}

char* QSound_FileName(void* ptr){
	return static_cast<QSound*>(ptr)->fileName().toUtf8().data();
}

int QSound_IsFinished(void* ptr){
	return static_cast<QSound*>(ptr)->isFinished();
}

int QSound_Loops(void* ptr){
	return static_cast<QSound*>(ptr)->loops();
}

int QSound_LoopsRemaining(void* ptr){
	return static_cast<QSound*>(ptr)->loopsRemaining();
}

void QSound_Play2(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSound*>(ptr), "play");
}

void QSound_QSound_Play(char* filename){
	QSound::play(QString(filename));
}

void QSound_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSound*>(ptr), "stop");
}

void QSound_DestroyQSound(void* ptr){
	static_cast<QSound*>(ptr)->~QSound();
}

class MyQSoundEffect: public QSoundEffect {
public:
	void Signal_CategoryChanged() { callbackQSoundEffectCategoryChanged(this->objectName().toUtf8().data()); };
	void Signal_LoadedChanged() { callbackQSoundEffectLoadedChanged(this->objectName().toUtf8().data()); };
	void Signal_LoopCountChanged() { callbackQSoundEffectLoopCountChanged(this->objectName().toUtf8().data()); };
	void Signal_LoopsRemainingChanged() { callbackQSoundEffectLoopsRemainingChanged(this->objectName().toUtf8().data()); };
	void Signal_MutedChanged() { callbackQSoundEffectMutedChanged(this->objectName().toUtf8().data()); };
	void Signal_PlayingChanged() { callbackQSoundEffectPlayingChanged(this->objectName().toUtf8().data()); };
	void Signal_SourceChanged() { callbackQSoundEffectSourceChanged(this->objectName().toUtf8().data()); };
	void Signal_StatusChanged() { callbackQSoundEffectStatusChanged(this->objectName().toUtf8().data()); };
	void Signal_VolumeChanged() { callbackQSoundEffectVolumeChanged(this->objectName().toUtf8().data()); };
protected:
};

int QSoundEffect_IsLoaded(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->isLoaded();
}

int QSoundEffect_LoopsRemaining(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->loopsRemaining();
}

void QSoundEffect_Play(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSoundEffect*>(ptr), "play");
}

void QSoundEffect_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSoundEffect*>(ptr), "stop");
}

char* QSoundEffect_QSoundEffect_SupportedMimeTypes(){
	return QSoundEffect::supportedMimeTypes().join(",,,").toUtf8().data();
}

void* QSoundEffect_NewQSoundEffect(void* parent){
	return new QSoundEffect(static_cast<QObject*>(parent));
}

char* QSoundEffect_Category(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->category().toUtf8().data();
}

void QSoundEffect_ConnectCategoryChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::categoryChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_CategoryChanged));;
}

void QSoundEffect_DisconnectCategoryChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::categoryChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_CategoryChanged));;
}

int QSoundEffect_IsMuted(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->isMuted();
}

int QSoundEffect_IsPlaying(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->isPlaying();
}

void QSoundEffect_ConnectLoadedChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loadedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoadedChanged));;
}

void QSoundEffect_DisconnectLoadedChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loadedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoadedChanged));;
}

int QSoundEffect_LoopCount(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->loopCount();
}

void QSoundEffect_ConnectLoopCountChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopCountChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopCountChanged));;
}

void QSoundEffect_DisconnectLoopCountChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopCountChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopCountChanged));;
}

void QSoundEffect_ConnectLoopsRemainingChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopsRemainingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopsRemainingChanged));;
}

void QSoundEffect_DisconnectLoopsRemainingChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopsRemainingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_LoopsRemainingChanged));;
}

void QSoundEffect_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::mutedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_MutedChanged));;
}

void QSoundEffect_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::mutedChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_MutedChanged));;
}

void QSoundEffect_ConnectPlayingChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::playingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_PlayingChanged));;
}

void QSoundEffect_DisconnectPlayingChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::playingChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_PlayingChanged));;
}

void QSoundEffect_SetCategory(void* ptr, char* category){
	static_cast<QSoundEffect*>(ptr)->setCategory(QString(category));
}

void QSoundEffect_SetLoopCount(void* ptr, int loopCount){
	static_cast<QSoundEffect*>(ptr)->setLoopCount(loopCount);
}

void QSoundEffect_SetMuted(void* ptr, int muted){
	static_cast<QSoundEffect*>(ptr)->setMuted(muted != 0);
}

void QSoundEffect_SetSource(void* ptr, void* url){
	static_cast<QSoundEffect*>(ptr)->setSource(*static_cast<QUrl*>(url));
}

void QSoundEffect_SetVolume(void* ptr, double volume){
	static_cast<QSoundEffect*>(ptr)->setVolume(static_cast<qreal>(volume));
}

void QSoundEffect_ConnectSourceChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::sourceChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_SourceChanged));;
}

void QSoundEffect_DisconnectSourceChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::sourceChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_SourceChanged));;
}

int QSoundEffect_Status(void* ptr){
	return static_cast<QSoundEffect*>(ptr)->status();
}

void QSoundEffect_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::statusChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_StatusChanged));;
}

void QSoundEffect_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::statusChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_StatusChanged));;
}

double QSoundEffect_Volume(void* ptr){
	return static_cast<double>(static_cast<QSoundEffect*>(ptr)->volume());
}

void QSoundEffect_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::volumeChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_VolumeChanged));;
}

void QSoundEffect_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QSoundEffect*>(ptr), static_cast<void (QSoundEffect::*)()>(&QSoundEffect::volumeChanged), static_cast<MyQSoundEffect*>(ptr), static_cast<void (MyQSoundEffect::*)()>(&MyQSoundEffect::Signal_VolumeChanged));;
}

void QSoundEffect_DestroyQSoundEffect(void* ptr){
	static_cast<QSoundEffect*>(ptr)->~QSoundEffect();
}

class MyQVideoDeviceSelectorControl: public QVideoDeviceSelectorControl {
public:
	void Signal_DevicesChanged() { callbackQVideoDeviceSelectorControlDevicesChanged(this->objectName().toUtf8().data()); };
	void Signal_SelectedDeviceChanged(int index) { callbackQVideoDeviceSelectorControlSelectedDeviceChanged(this->objectName().toUtf8().data(), index); };
protected:
};

int QVideoDeviceSelectorControl_DefaultDevice(void* ptr){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->defaultDevice();
}

int QVideoDeviceSelectorControl_DeviceCount(void* ptr){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->deviceCount();
}

char* QVideoDeviceSelectorControl_DeviceDescription(void* ptr, int index){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->deviceDescription(index).toUtf8().data();
}

char* QVideoDeviceSelectorControl_DeviceName(void* ptr, int index){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->deviceName(index).toUtf8().data();
}

void QVideoDeviceSelectorControl_ConnectDevicesChanged(void* ptr){
	QObject::connect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)()>(&QVideoDeviceSelectorControl::devicesChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)()>(&MyQVideoDeviceSelectorControl::Signal_DevicesChanged));;
}

void QVideoDeviceSelectorControl_DisconnectDevicesChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)()>(&QVideoDeviceSelectorControl::devicesChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)()>(&MyQVideoDeviceSelectorControl::Signal_DevicesChanged));;
}

int QVideoDeviceSelectorControl_SelectedDevice(void* ptr){
	return static_cast<QVideoDeviceSelectorControl*>(ptr)->selectedDevice();
}

void QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged(void* ptr){
	QObject::connect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)(int)>(&QVideoDeviceSelectorControl::selectedDeviceChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)(int)>(&MyQVideoDeviceSelectorControl::Signal_SelectedDeviceChanged));;
}

void QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoDeviceSelectorControl*>(ptr), static_cast<void (QVideoDeviceSelectorControl::*)(int)>(&QVideoDeviceSelectorControl::selectedDeviceChanged), static_cast<MyQVideoDeviceSelectorControl*>(ptr), static_cast<void (MyQVideoDeviceSelectorControl::*)(int)>(&MyQVideoDeviceSelectorControl::Signal_SelectedDeviceChanged));;
}

void QVideoDeviceSelectorControl_SetSelectedDevice(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QVideoDeviceSelectorControl*>(ptr), "setSelectedDevice", Q_ARG(int, index));
}

void QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(void* ptr){
	static_cast<QVideoDeviceSelectorControl*>(ptr)->~QVideoDeviceSelectorControl();
}

void QVideoEncoderSettings_SetFrameRate(void* ptr, double rate){
	static_cast<QVideoEncoderSettings*>(ptr)->setFrameRate(static_cast<qreal>(rate));
}

void* QVideoEncoderSettings_NewQVideoEncoderSettings(){
	return new QVideoEncoderSettings();
}

void* QVideoEncoderSettings_NewQVideoEncoderSettings2(void* other){
	return new QVideoEncoderSettings(*static_cast<QVideoEncoderSettings*>(other));
}

int QVideoEncoderSettings_BitRate(void* ptr){
	return static_cast<QVideoEncoderSettings*>(ptr)->bitRate();
}

char* QVideoEncoderSettings_Codec(void* ptr){
	return static_cast<QVideoEncoderSettings*>(ptr)->codec().toUtf8().data();
}

void* QVideoEncoderSettings_EncodingOption(void* ptr, char* option){
	return new QVariant(static_cast<QVideoEncoderSettings*>(ptr)->encodingOption(QString(option)));
}

double QVideoEncoderSettings_FrameRate(void* ptr){
	return static_cast<double>(static_cast<QVideoEncoderSettings*>(ptr)->frameRate());
}

int QVideoEncoderSettings_IsNull(void* ptr){
	return static_cast<QVideoEncoderSettings*>(ptr)->isNull();
}

void QVideoEncoderSettings_SetBitRate(void* ptr, int value){
	static_cast<QVideoEncoderSettings*>(ptr)->setBitRate(value);
}

void QVideoEncoderSettings_SetCodec(void* ptr, char* codec){
	static_cast<QVideoEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QVideoEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value){
	static_cast<QVideoEncoderSettings*>(ptr)->setEncodingOption(QString(option), *static_cast<QVariant*>(value));
}

void QVideoEncoderSettings_SetResolution(void* ptr, void* resolution){
	static_cast<QVideoEncoderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QVideoEncoderSettings_SetResolution2(void* ptr, int width, int height){
	static_cast<QVideoEncoderSettings*>(ptr)->setResolution(width, height);
}

void QVideoEncoderSettings_DestroyQVideoEncoderSettings(void* ptr){
	static_cast<QVideoEncoderSettings*>(ptr)->~QVideoEncoderSettings();
}

class MyQVideoEncoderSettingsControl: public QVideoEncoderSettingsControl {
public:
protected:
};

void QVideoEncoderSettingsControl_SetVideoSettings(void* ptr, void* settings){
	static_cast<QVideoEncoderSettingsControl*>(ptr)->setVideoSettings(*static_cast<QVideoEncoderSettings*>(settings));
}

char* QVideoEncoderSettingsControl_SupportedVideoCodecs(void* ptr){
	return static_cast<QVideoEncoderSettingsControl*>(ptr)->supportedVideoCodecs().join(",,,").toUtf8().data();
}

char* QVideoEncoderSettingsControl_VideoCodecDescription(void* ptr, char* codec){
	return static_cast<QVideoEncoderSettingsControl*>(ptr)->videoCodecDescription(QString(codec)).toUtf8().data();
}

void QVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControl(void* ptr){
	static_cast<QVideoEncoderSettingsControl*>(ptr)->~QVideoEncoderSettingsControl();
}

void* QVideoFrame_NewQVideoFrame(){
	return new QVideoFrame();
}

void* QVideoFrame_NewQVideoFrame2(void* buffer, void* size, int format){
	return new QVideoFrame(static_cast<QAbstractVideoBuffer*>(buffer), *static_cast<QSize*>(size), static_cast<QVideoFrame::PixelFormat>(format));
}

void* QVideoFrame_NewQVideoFrame4(void* image){
	return new QVideoFrame(*static_cast<QImage*>(image));
}

void* QVideoFrame_NewQVideoFrame5(void* other){
	return new QVideoFrame(*static_cast<QVideoFrame*>(other));
}

void* QVideoFrame_NewQVideoFrame3(int bytes, void* size, int bytesPerLine, int format){
	return new QVideoFrame(bytes, *static_cast<QSize*>(size), bytesPerLine, static_cast<QVideoFrame::PixelFormat>(format));
}

int QVideoFrame_BytesPerLine(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->bytesPerLine();
}

int QVideoFrame_BytesPerLine2(void* ptr, int plane){
	return static_cast<QVideoFrame*>(ptr)->bytesPerLine(plane);
}

int QVideoFrame_FieldType(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->fieldType();
}

void* QVideoFrame_Handle(void* ptr){
	return new QVariant(static_cast<QVideoFrame*>(ptr)->handle());
}

int QVideoFrame_HandleType(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->handleType();
}

int QVideoFrame_Height(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->height();
}

int QVideoFrame_QVideoFrame_ImageFormatFromPixelFormat(int format){
	return QVideoFrame::imageFormatFromPixelFormat(static_cast<QVideoFrame::PixelFormat>(format));
}

int QVideoFrame_IsMapped(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->isMapped();
}

int QVideoFrame_IsReadable(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->isReadable();
}

int QVideoFrame_IsValid(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->isValid();
}

int QVideoFrame_IsWritable(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->isWritable();
}

int QVideoFrame_Map(void* ptr, int mode){
	return static_cast<QVideoFrame*>(ptr)->map(static_cast<QAbstractVideoBuffer::MapMode>(mode));
}

int QVideoFrame_MapMode(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->mapMode();
}

int QVideoFrame_MappedBytes(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->mappedBytes();
}

void* QVideoFrame_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QVideoFrame*>(ptr)->metaData(QString(key)));
}

int QVideoFrame_PixelFormat(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->pixelFormat();
}

int QVideoFrame_QVideoFrame_PixelFormatFromImageFormat(int format){
	return QVideoFrame::pixelFormatFromImageFormat(static_cast<QImage::Format>(format));
}

int QVideoFrame_PlaneCount(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->planeCount();
}

void QVideoFrame_SetFieldType(void* ptr, int field){
	static_cast<QVideoFrame*>(ptr)->setFieldType(static_cast<QVideoFrame::FieldType>(field));
}

void QVideoFrame_SetMetaData(void* ptr, char* key, void* value){
	static_cast<QVideoFrame*>(ptr)->setMetaData(QString(key), *static_cast<QVariant*>(value));
}

void QVideoFrame_Unmap(void* ptr){
	static_cast<QVideoFrame*>(ptr)->unmap();
}

int QVideoFrame_Width(void* ptr){
	return static_cast<QVideoFrame*>(ptr)->width();
}

void QVideoFrame_DestroyQVideoFrame(void* ptr){
	static_cast<QVideoFrame*>(ptr)->~QVideoFrame();
}

class MyQVideoProbe: public QVideoProbe {
public:
	void Signal_Flush() { callbackQVideoProbeFlush(this->objectName().toUtf8().data()); };
protected:
};

void* QVideoProbe_NewQVideoProbe(void* parent){
	return new QVideoProbe(static_cast<QObject*>(parent));
}

void QVideoProbe_ConnectFlush(void* ptr){
	QObject::connect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)()>(&QVideoProbe::flush), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)()>(&MyQVideoProbe::Signal_Flush));;
}

void QVideoProbe_DisconnectFlush(void* ptr){
	QObject::disconnect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)()>(&QVideoProbe::flush), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)()>(&MyQVideoProbe::Signal_Flush));;
}

int QVideoProbe_IsActive(void* ptr){
	return static_cast<QVideoProbe*>(ptr)->isActive();
}

int QVideoProbe_SetSource(void* ptr, void* source){
	return static_cast<QVideoProbe*>(ptr)->setSource(static_cast<QMediaObject*>(source));
}

int QVideoProbe_SetSource2(void* ptr, void* mediaRecorder){
	return static_cast<QVideoProbe*>(ptr)->setSource(static_cast<QMediaRecorder*>(mediaRecorder));
}

void QVideoProbe_DestroyQVideoProbe(void* ptr){
	static_cast<QVideoProbe*>(ptr)->~QVideoProbe();
}

void QVideoRendererControl_SetSurface(void* ptr, void* surface){
	static_cast<QVideoRendererControl*>(ptr)->setSurface(static_cast<QAbstractVideoSurface*>(surface));
}

void* QVideoRendererControl_Surface(void* ptr){
	return static_cast<QVideoRendererControl*>(ptr)->surface();
}

void QVideoRendererControl_DestroyQVideoRendererControl(void* ptr){
	static_cast<QVideoRendererControl*>(ptr)->~QVideoRendererControl();
}

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat(){
	return new QVideoSurfaceFormat();
}

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat2(void* size, int format, int ty){
	return new QVideoSurfaceFormat(*static_cast<QSize*>(size), static_cast<QVideoFrame::PixelFormat>(format), static_cast<QAbstractVideoBuffer::HandleType>(ty));
}

void* QVideoSurfaceFormat_NewQVideoSurfaceFormat3(void* other){
	return new QVideoSurfaceFormat(*static_cast<QVideoSurfaceFormat*>(other));
}

int QVideoSurfaceFormat_FrameHeight(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->frameHeight();
}

double QVideoSurfaceFormat_FrameRate(void* ptr){
	return static_cast<double>(static_cast<QVideoSurfaceFormat*>(ptr)->frameRate());
}

int QVideoSurfaceFormat_FrameWidth(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->frameWidth();
}

int QVideoSurfaceFormat_HandleType(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->handleType();
}

int QVideoSurfaceFormat_IsValid(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->isValid();
}

int QVideoSurfaceFormat_PixelFormat(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->pixelFormat();
}

void* QVideoSurfaceFormat_Property(void* ptr, char* name){
	return new QVariant(static_cast<QVideoSurfaceFormat*>(ptr)->property(const_cast<const char*>(name)));
}

int QVideoSurfaceFormat_ScanLineDirection(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->scanLineDirection();
}

void QVideoSurfaceFormat_SetFrameRate(void* ptr, double rate){
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameRate(static_cast<qreal>(rate));
}

void QVideoSurfaceFormat_SetFrameSize(void* ptr, void* size){
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameSize(*static_cast<QSize*>(size));
}

void QVideoSurfaceFormat_SetFrameSize2(void* ptr, int width, int height){
	static_cast<QVideoSurfaceFormat*>(ptr)->setFrameSize(width, height);
}

void QVideoSurfaceFormat_SetPixelAspectRatio(void* ptr, void* ratio){
	static_cast<QVideoSurfaceFormat*>(ptr)->setPixelAspectRatio(*static_cast<QSize*>(ratio));
}

void QVideoSurfaceFormat_SetPixelAspectRatio2(void* ptr, int horizontal, int vertical){
	static_cast<QVideoSurfaceFormat*>(ptr)->setPixelAspectRatio(horizontal, vertical);
}

void QVideoSurfaceFormat_SetProperty(void* ptr, char* name, void* value){
	static_cast<QVideoSurfaceFormat*>(ptr)->setProperty(const_cast<const char*>(name), *static_cast<QVariant*>(value));
}

void QVideoSurfaceFormat_SetScanLineDirection(void* ptr, int direction){
	static_cast<QVideoSurfaceFormat*>(ptr)->setScanLineDirection(static_cast<QVideoSurfaceFormat::Direction>(direction));
}

void QVideoSurfaceFormat_SetViewport(void* ptr, void* viewport){
	static_cast<QVideoSurfaceFormat*>(ptr)->setViewport(*static_cast<QRect*>(viewport));
}

void QVideoSurfaceFormat_SetYCbCrColorSpace(void* ptr, int space){
	static_cast<QVideoSurfaceFormat*>(ptr)->setYCbCrColorSpace(static_cast<QVideoSurfaceFormat::YCbCrColorSpace>(space));
}

int QVideoSurfaceFormat_YCbCrColorSpace(void* ptr){
	return static_cast<QVideoSurfaceFormat*>(ptr)->yCbCrColorSpace();
}

void QVideoSurfaceFormat_DestroyQVideoSurfaceFormat(void* ptr){
	static_cast<QVideoSurfaceFormat*>(ptr)->~QVideoSurfaceFormat();
}

class MyQVideoWindowControl: public QVideoWindowControl {
public:
	void Signal_BrightnessChanged(int brightness) { callbackQVideoWindowControlBrightnessChanged(this->objectName().toUtf8().data(), brightness); };
	void Signal_ContrastChanged(int contrast) { callbackQVideoWindowControlContrastChanged(this->objectName().toUtf8().data(), contrast); };
	void Signal_FullScreenChanged(bool fullScreen) { callbackQVideoWindowControlFullScreenChanged(this->objectName().toUtf8().data(), fullScreen); };
	void Signal_HueChanged(int hue) { callbackQVideoWindowControlHueChanged(this->objectName().toUtf8().data(), hue); };
	void Signal_NativeSizeChanged() { callbackQVideoWindowControlNativeSizeChanged(this->objectName().toUtf8().data()); };
	void Signal_SaturationChanged(int saturation) { callbackQVideoWindowControlSaturationChanged(this->objectName().toUtf8().data(), saturation); };
protected:
};

int QVideoWindowControl_AspectRatioMode(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->aspectRatioMode();
}

int QVideoWindowControl_Brightness(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->brightness();
}

void QVideoWindowControl_ConnectBrightnessChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::brightnessChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_BrightnessChanged));;
}

void QVideoWindowControl_DisconnectBrightnessChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::brightnessChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_BrightnessChanged));;
}

int QVideoWindowControl_Contrast(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->contrast();
}

void QVideoWindowControl_ConnectContrastChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::contrastChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_ContrastChanged));;
}

void QVideoWindowControl_DisconnectContrastChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::contrastChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_ContrastChanged));;
}

void QVideoWindowControl_ConnectFullScreenChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(bool)>(&QVideoWindowControl::fullScreenChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(bool)>(&MyQVideoWindowControl::Signal_FullScreenChanged));;
}

void QVideoWindowControl_DisconnectFullScreenChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(bool)>(&QVideoWindowControl::fullScreenChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(bool)>(&MyQVideoWindowControl::Signal_FullScreenChanged));;
}

int QVideoWindowControl_Hue(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->hue();
}

void QVideoWindowControl_ConnectHueChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::hueChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_HueChanged));;
}

void QVideoWindowControl_DisconnectHueChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::hueChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_HueChanged));;
}

int QVideoWindowControl_IsFullScreen(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->isFullScreen();
}

void QVideoWindowControl_ConnectNativeSizeChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)()>(&QVideoWindowControl::nativeSizeChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)()>(&MyQVideoWindowControl::Signal_NativeSizeChanged));;
}

void QVideoWindowControl_DisconnectNativeSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)()>(&QVideoWindowControl::nativeSizeChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)()>(&MyQVideoWindowControl::Signal_NativeSizeChanged));;
}

void QVideoWindowControl_Repaint(void* ptr){
	static_cast<QVideoWindowControl*>(ptr)->repaint();
}

int QVideoWindowControl_Saturation(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->saturation();
}

void QVideoWindowControl_ConnectSaturationChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::saturationChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_SaturationChanged));;
}

void QVideoWindowControl_DisconnectSaturationChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::saturationChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_SaturationChanged));;
}

void QVideoWindowControl_SetAspectRatioMode(void* ptr, int mode){
	static_cast<QVideoWindowControl*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QVideoWindowControl_SetBrightness(void* ptr, int brightness){
	static_cast<QVideoWindowControl*>(ptr)->setBrightness(brightness);
}

void QVideoWindowControl_SetContrast(void* ptr, int contrast){
	static_cast<QVideoWindowControl*>(ptr)->setContrast(contrast);
}

void QVideoWindowControl_SetDisplayRect(void* ptr, void* rect){
	static_cast<QVideoWindowControl*>(ptr)->setDisplayRect(*static_cast<QRect*>(rect));
}

void QVideoWindowControl_SetFullScreen(void* ptr, int fullScreen){
	static_cast<QVideoWindowControl*>(ptr)->setFullScreen(fullScreen != 0);
}

void QVideoWindowControl_SetHue(void* ptr, int hue){
	static_cast<QVideoWindowControl*>(ptr)->setHue(hue);
}

void QVideoWindowControl_SetSaturation(void* ptr, int saturation){
	static_cast<QVideoWindowControl*>(ptr)->setSaturation(saturation);
}

void QVideoWindowControl_DestroyQVideoWindowControl(void* ptr){
	static_cast<QVideoWindowControl*>(ptr)->~QVideoWindowControl();
}

