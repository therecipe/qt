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

void* QSurfaceFormat_NewQSurfaceFormat(){
	return new QSurfaceFormat();
}

void* QSurfaceFormat_NewQSurfaceFormat2(int options){
	return new QSurfaceFormat(static_cast<QSurfaceFormat::FormatOption>(options));
}

void* QSurfaceFormat_NewQSurfaceFormat3(void* other){
	return new QSurfaceFormat(*static_cast<QSurfaceFormat*>(other));
}

int QSurfaceFormat_AlphaBufferSize(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->alphaBufferSize();
}

int QSurfaceFormat_BlueBufferSize(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->blueBufferSize();
}

int QSurfaceFormat_DepthBufferSize(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->depthBufferSize();
}

int QSurfaceFormat_GreenBufferSize(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->greenBufferSize();
}

int QSurfaceFormat_HasAlpha(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->hasAlpha();
}

int QSurfaceFormat_MajorVersion(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->majorVersion();
}

int QSurfaceFormat_MinorVersion(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->minorVersion();
}

int QSurfaceFormat_Options(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->options();
}

int QSurfaceFormat_Profile(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->profile();
}

int QSurfaceFormat_RedBufferSize(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->redBufferSize();
}

int QSurfaceFormat_RenderableType(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->renderableType();
}

int QSurfaceFormat_Samples(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->samples();
}

void QSurfaceFormat_SetAlphaBufferSize(void* ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setAlphaBufferSize(size);
}

void QSurfaceFormat_SetBlueBufferSize(void* ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setBlueBufferSize(size);
}

void QSurfaceFormat_QSurfaceFormat_SetDefaultFormat(void* format){
	QSurfaceFormat::setDefaultFormat(*static_cast<QSurfaceFormat*>(format));
}

void QSurfaceFormat_SetDepthBufferSize(void* ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setDepthBufferSize(size);
}

void QSurfaceFormat_SetGreenBufferSize(void* ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setGreenBufferSize(size);
}

void QSurfaceFormat_SetMajorVersion(void* ptr, int major){
	static_cast<QSurfaceFormat*>(ptr)->setMajorVersion(major);
}

void QSurfaceFormat_SetMinorVersion(void* ptr, int minor){
	static_cast<QSurfaceFormat*>(ptr)->setMinorVersion(minor);
}

void QSurfaceFormat_SetOption(void* ptr, int option, int on){
	static_cast<QSurfaceFormat*>(ptr)->setOption(static_cast<QSurfaceFormat::FormatOption>(option), on != 0);
}

void QSurfaceFormat_SetOptions(void* ptr, int options){
	static_cast<QSurfaceFormat*>(ptr)->setOptions(static_cast<QSurfaceFormat::FormatOption>(options));
}

void QSurfaceFormat_SetProfile(void* ptr, int profile){
	static_cast<QSurfaceFormat*>(ptr)->setProfile(static_cast<QSurfaceFormat::OpenGLContextProfile>(profile));
}

void QSurfaceFormat_SetRedBufferSize(void* ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setRedBufferSize(size);
}

void QSurfaceFormat_SetRenderableType(void* ptr, int ty){
	static_cast<QSurfaceFormat*>(ptr)->setRenderableType(static_cast<QSurfaceFormat::RenderableType>(ty));
}

void QSurfaceFormat_SetSamples(void* ptr, int numSamples){
	static_cast<QSurfaceFormat*>(ptr)->setSamples(numSamples);
}

void QSurfaceFormat_SetStencilBufferSize(void* ptr, int size){
	static_cast<QSurfaceFormat*>(ptr)->setStencilBufferSize(size);
}

void QSurfaceFormat_SetStereo(void* ptr, int enable){
	static_cast<QSurfaceFormat*>(ptr)->setStereo(enable != 0);
}

void QSurfaceFormat_SetSwapBehavior(void* ptr, int behavior){
	static_cast<QSurfaceFormat*>(ptr)->setSwapBehavior(static_cast<QSurfaceFormat::SwapBehavior>(behavior));
}

void QSurfaceFormat_SetSwapInterval(void* ptr, int interval){
	static_cast<QSurfaceFormat*>(ptr)->setSwapInterval(interval);
}

void QSurfaceFormat_SetVersion(void* ptr, int major, int minor){
	static_cast<QSurfaceFormat*>(ptr)->setVersion(major, minor);
}

int QSurfaceFormat_StencilBufferSize(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->stencilBufferSize();
}

int QSurfaceFormat_Stereo(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->stereo();
}

int QSurfaceFormat_SwapBehavior(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->swapBehavior();
}

int QSurfaceFormat_SwapInterval(void* ptr){
	return static_cast<QSurfaceFormat*>(ptr)->swapInterval();
}

int QSurfaceFormat_TestOption(void* ptr, int option){
	return static_cast<QSurfaceFormat*>(ptr)->testOption(static_cast<QSurfaceFormat::FormatOption>(option));
}

void QSurfaceFormat_DestroyQSurfaceFormat(void* ptr){
	static_cast<QSurfaceFormat*>(ptr)->~QSurfaceFormat();
}

