#include "qsurfaceformat.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSurface>
#include <QSurfaceFormat>
#include "_cgo_export.h"

class MyQSurfaceFormat: public QSurfaceFormat {
public:
};

QtObjectPtr QSurfaceFormat_NewQSurfaceFormat(){
	return new QSurfaceFormat();
}

QtObjectPtr QSurfaceFormat_NewQSurfaceFormat2(int options){
	return new QSurfaceFormat(static_cast<QSurfaceFormat::FormatOption>(options));
}

QtObjectPtr QSurfaceFormat_NewQSurfaceFormat3(QtObjectPtr other){
	return new QSurfaceFormat(*static_cast<QSurfaceFormat*>(other));
}

int QSurfaceFormat_AlphaBufferSize(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->alphaBufferSize();
}

int QSurfaceFormat_BlueBufferSize(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->blueBufferSize();
}

int QSurfaceFormat_DepthBufferSize(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->depthBufferSize();
}

int QSurfaceFormat_GreenBufferSize(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->greenBufferSize();
}

int QSurfaceFormat_HasAlpha(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->hasAlpha();
}

int QSurfaceFormat_MajorVersion(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->majorVersion();
}

int QSurfaceFormat_MinorVersion(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->minorVersion();
}

int QSurfaceFormat_Options(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->options();
}

int QSurfaceFormat_Profile(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->profile();
}

int QSurfaceFormat_RedBufferSize(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->redBufferSize();
}

int QSurfaceFormat_RenderableType(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->renderableType();
}

int QSurfaceFormat_Samples(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->samples();
}

void QSurfaceFormat_SetAlphaBufferSize(QtObjectPtr ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setAlphaBufferSize(size);
}

void QSurfaceFormat_SetBlueBufferSize(QtObjectPtr ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setBlueBufferSize(size);
}

void QSurfaceFormat_QSurfaceFormat_SetDefaultFormat(QtObjectPtr format){
	QSurfaceFormat::setDefaultFormat(*static_cast<QSurfaceFormat*>(format));
}

void QSurfaceFormat_SetDepthBufferSize(QtObjectPtr ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setDepthBufferSize(size);
}

void QSurfaceFormat_SetGreenBufferSize(QtObjectPtr ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setGreenBufferSize(size);
}

void QSurfaceFormat_SetMajorVersion(QtObjectPtr ptr, int major){
	static_cast<QSurfaceFormat*>(ptr)->setMajorVersion(major);
}

void QSurfaceFormat_SetMinorVersion(QtObjectPtr ptr, int minor){
	static_cast<QSurfaceFormat*>(ptr)->setMinorVersion(minor);
}

void QSurfaceFormat_SetOption(QtObjectPtr ptr, int option, int on){
	static_cast<QSurfaceFormat*>(ptr)->setOption(static_cast<QSurfaceFormat::FormatOption>(option), on != 0);
}

void QSurfaceFormat_SetOptions(QtObjectPtr ptr, int options){
	static_cast<QSurfaceFormat*>(ptr)->setOptions(static_cast<QSurfaceFormat::FormatOption>(options));
}

void QSurfaceFormat_SetProfile(QtObjectPtr ptr, int profile){
	static_cast<QSurfaceFormat*>(ptr)->setProfile(static_cast<QSurfaceFormat::OpenGLContextProfile>(profile));
}

void QSurfaceFormat_SetRedBufferSize(QtObjectPtr ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setRedBufferSize(size);
}

void QSurfaceFormat_SetRenderableType(QtObjectPtr ptr, int ty){
	static_cast<QSurfaceFormat*>(ptr)->setRenderableType(static_cast<QSurfaceFormat::RenderableType>(ty));
}

void QSurfaceFormat_SetSamples(QtObjectPtr ptr, int numSamples){
	static_cast<QSurfaceFormat*>(ptr)->setSamples(numSamples);
}

void QSurfaceFormat_SetStencilBufferSize(QtObjectPtr ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setStencilBufferSize(size);
}

void QSurfaceFormat_SetStereo(QtObjectPtr ptr, int enable){
	static_cast<QSurfaceFormat*>(ptr)->setStereo(enable != 0);
}

void QSurfaceFormat_SetSwapBehavior(QtObjectPtr ptr, int behavior){
	static_cast<QSurfaceFormat*>(ptr)->setSwapBehavior(static_cast<QSurfaceFormat::SwapBehavior>(behavior));
}

void QSurfaceFormat_SetSwapInterval(QtObjectPtr ptr, int interval){
	static_cast<QSurfaceFormat*>(ptr)->setSwapInterval(interval);
}

void QSurfaceFormat_SetVersion(QtObjectPtr ptr, int major, int minor){
	static_cast<QSurfaceFormat*>(ptr)->setVersion(major, minor);
}

int QSurfaceFormat_StencilBufferSize(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->stencilBufferSize();
}

int QSurfaceFormat_Stereo(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->stereo();
}

int QSurfaceFormat_SwapBehavior(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->swapBehavior();
}

int QSurfaceFormat_SwapInterval(QtObjectPtr ptr){
	return static_cast<QSurfaceFormat*>(ptr)->swapInterval();
}

int QSurfaceFormat_TestOption(QtObjectPtr ptr, int option){
	return static_cast<QSurfaceFormat*>(ptr)->testOption(static_cast<QSurfaceFormat::FormatOption>(option));
}

void QSurfaceFormat_DestroyQSurfaceFormat(QtObjectPtr ptr){
	static_cast<QSurfaceFormat*>(ptr)->~QSurfaceFormat();
}

