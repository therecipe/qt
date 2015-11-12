#include "qmovie.h"
#include <QImageReader>
#include <QSize>
#include <QColor>
#include <QString>
#include <QUrl>
#include <QMetaObject>
#include <QByteArray>
#include <QObject>
#include <QVariant>
#include <QModelIndex>
#include <QImage>
#include <QIODevice>
#include <QMovie>
#include "_cgo_export.h"

class MyQMovie: public QMovie {
public:
void Signal_Error(QImageReader::ImageReaderError error){callbackQMovieError(this->objectName().toUtf8().data(), error);};
void Signal_Finished(){callbackQMovieFinished(this->objectName().toUtf8().data());};
void Signal_FrameChanged(int frameNumber){callbackQMovieFrameChanged(this->objectName().toUtf8().data(), frameNumber);};
void Signal_Started(){callbackQMovieStarted(this->objectName().toUtf8().data());};
void Signal_StateChanged(QMovie::MovieState state){callbackQMovieStateChanged(this->objectName().toUtf8().data(), state);};
};

int QMovie_CacheMode(void* ptr){
	return static_cast<QMovie*>(ptr)->cacheMode();
}

void QMovie_SetCacheMode(void* ptr, int mode){
	static_cast<QMovie*>(ptr)->setCacheMode(static_cast<QMovie::CacheMode>(mode));
}

void QMovie_SetSpeed(void* ptr, int percentSpeed){
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "setSpeed", Q_ARG(int, percentSpeed));
}

int QMovie_Speed(void* ptr){
	return static_cast<QMovie*>(ptr)->speed();
}

void* QMovie_NewQMovie2(void* device, void* format, void* parent){
	return new QMovie(static_cast<QIODevice*>(device), *static_cast<QByteArray*>(format), static_cast<QObject*>(parent));
}

void* QMovie_NewQMovie(void* parent){
	return new QMovie(static_cast<QObject*>(parent));
}

void* QMovie_NewQMovie3(char* fileName, void* format, void* parent){
	return new QMovie(QString(fileName), *static_cast<QByteArray*>(format), static_cast<QObject*>(parent));
}

void* QMovie_BackgroundColor(void* ptr){
	return new QColor(static_cast<QMovie*>(ptr)->backgroundColor());
}

int QMovie_CurrentFrameNumber(void* ptr){
	return static_cast<QMovie*>(ptr)->currentFrameNumber();
}

void* QMovie_Device(void* ptr){
	return static_cast<QMovie*>(ptr)->device();
}

void QMovie_ConnectError(void* ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QImageReader::ImageReaderError)>(&QMovie::error), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QImageReader::ImageReaderError)>(&MyQMovie::Signal_Error));;
}

void QMovie_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QImageReader::ImageReaderError)>(&QMovie::error), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QImageReader::ImageReaderError)>(&MyQMovie::Signal_Error));;
}

char* QMovie_FileName(void* ptr){
	return static_cast<QMovie*>(ptr)->fileName().toUtf8().data();
}

void QMovie_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::finished), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Finished));;
}

void QMovie_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::finished), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Finished));;
}

void* QMovie_Format(void* ptr){
	return new QByteArray(static_cast<QMovie*>(ptr)->format());
}

void QMovie_ConnectFrameChanged(void* ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(int)>(&QMovie::frameChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(int)>(&MyQMovie::Signal_FrameChanged));;
}

void QMovie_DisconnectFrameChanged(void* ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(int)>(&QMovie::frameChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(int)>(&MyQMovie::Signal_FrameChanged));;
}

int QMovie_FrameCount(void* ptr){
	return static_cast<QMovie*>(ptr)->frameCount();
}

int QMovie_IsValid(void* ptr){
	return static_cast<QMovie*>(ptr)->isValid();
}

int QMovie_JumpToFrame(void* ptr, int frameNumber){
	return static_cast<QMovie*>(ptr)->jumpToFrame(frameNumber);
}

int QMovie_JumpToNextFrame(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "jumpToNextFrame");
}

int QMovie_LoopCount(void* ptr){
	return static_cast<QMovie*>(ptr)->loopCount();
}

int QMovie_NextFrameDelay(void* ptr){
	return static_cast<QMovie*>(ptr)->nextFrameDelay();
}

void QMovie_SetBackgroundColor(void* ptr, void* color){
	static_cast<QMovie*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void QMovie_SetDevice(void* ptr, void* device){
	static_cast<QMovie*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QMovie_SetFileName(void* ptr, char* fileName){
	static_cast<QMovie*>(ptr)->setFileName(QString(fileName));
}

void QMovie_SetFormat(void* ptr, void* format){
	static_cast<QMovie*>(ptr)->setFormat(*static_cast<QByteArray*>(format));
}

void QMovie_SetPaused(void* ptr, int paused){
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "setPaused", Q_ARG(bool, paused != 0));
}

void QMovie_SetScaledSize(void* ptr, void* size){
	static_cast<QMovie*>(ptr)->setScaledSize(*static_cast<QSize*>(size));
}

void QMovie_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "start");
}

void QMovie_ConnectStarted(void* ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::started), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Started));;
}

void QMovie_DisconnectStarted(void* ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::started), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Started));;
}

void QMovie_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QMovie::MovieState)>(&QMovie::stateChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QMovie::MovieState)>(&MyQMovie::Signal_StateChanged));;
}

void QMovie_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QMovie::MovieState)>(&QMovie::stateChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QMovie::MovieState)>(&MyQMovie::Signal_StateChanged));;
}

void QMovie_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "stop");
}

void QMovie_DestroyQMovie(void* ptr){
	static_cast<QMovie*>(ptr)->~QMovie();
}

