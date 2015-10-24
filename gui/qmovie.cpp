#include "qmovie.h"
#include <QString>
#include <QVariant>
#include <QImageReader>
#include <QColor>
#include <QSize>
#include <QObject>
#include <QByteArray>
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QMetaObject>
#include <QImage>
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

int QMovie_CacheMode(QtObjectPtr ptr){
	return static_cast<QMovie*>(ptr)->cacheMode();
}

void QMovie_SetCacheMode(QtObjectPtr ptr, int mode){
	static_cast<QMovie*>(ptr)->setCacheMode(static_cast<QMovie::CacheMode>(mode));
}

void QMovie_SetSpeed(QtObjectPtr ptr, int percentSpeed){
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "setSpeed", Q_ARG(int, percentSpeed));
}

int QMovie_Speed(QtObjectPtr ptr){
	return static_cast<QMovie*>(ptr)->speed();
}

QtObjectPtr QMovie_NewQMovie2(QtObjectPtr device, QtObjectPtr format, QtObjectPtr parent){
	return new QMovie(static_cast<QIODevice*>(device), *static_cast<QByteArray*>(format), static_cast<QObject*>(parent));
}

QtObjectPtr QMovie_NewQMovie(QtObjectPtr parent){
	return new QMovie(static_cast<QObject*>(parent));
}

QtObjectPtr QMovie_NewQMovie3(char* fileName, QtObjectPtr format, QtObjectPtr parent){
	return new QMovie(QString(fileName), *static_cast<QByteArray*>(format), static_cast<QObject*>(parent));
}

int QMovie_CurrentFrameNumber(QtObjectPtr ptr){
	return static_cast<QMovie*>(ptr)->currentFrameNumber();
}

QtObjectPtr QMovie_Device(QtObjectPtr ptr){
	return static_cast<QMovie*>(ptr)->device();
}

void QMovie_ConnectError(QtObjectPtr ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QImageReader::ImageReaderError)>(&QMovie::error), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QImageReader::ImageReaderError)>(&MyQMovie::Signal_Error));;
}

void QMovie_DisconnectError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QImageReader::ImageReaderError)>(&QMovie::error), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QImageReader::ImageReaderError)>(&MyQMovie::Signal_Error));;
}

char* QMovie_FileName(QtObjectPtr ptr){
	return static_cast<QMovie*>(ptr)->fileName().toUtf8().data();
}

void QMovie_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::finished), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Finished));;
}

void QMovie_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::finished), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Finished));;
}

void QMovie_ConnectFrameChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(int)>(&QMovie::frameChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(int)>(&MyQMovie::Signal_FrameChanged));;
}

void QMovie_DisconnectFrameChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(int)>(&QMovie::frameChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(int)>(&MyQMovie::Signal_FrameChanged));;
}

int QMovie_FrameCount(QtObjectPtr ptr){
	return static_cast<QMovie*>(ptr)->frameCount();
}

int QMovie_IsValid(QtObjectPtr ptr){
	return static_cast<QMovie*>(ptr)->isValid();
}

int QMovie_JumpToFrame(QtObjectPtr ptr, int frameNumber){
	return static_cast<QMovie*>(ptr)->jumpToFrame(frameNumber);
}

int QMovie_JumpToNextFrame(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "jumpToNextFrame");
}

int QMovie_LoopCount(QtObjectPtr ptr){
	return static_cast<QMovie*>(ptr)->loopCount();
}

int QMovie_NextFrameDelay(QtObjectPtr ptr){
	return static_cast<QMovie*>(ptr)->nextFrameDelay();
}

void QMovie_SetBackgroundColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QMovie*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void QMovie_SetDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QMovie*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QMovie_SetFileName(QtObjectPtr ptr, char* fileName){
	static_cast<QMovie*>(ptr)->setFileName(QString(fileName));
}

void QMovie_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QMovie*>(ptr)->setFormat(*static_cast<QByteArray*>(format));
}

void QMovie_SetPaused(QtObjectPtr ptr, int paused){
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "setPaused", Q_ARG(bool, paused != 0));
}

void QMovie_SetScaledSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QMovie*>(ptr)->setScaledSize(*static_cast<QSize*>(size));
}

void QMovie_Start(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "start");
}

void QMovie_ConnectStarted(QtObjectPtr ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::started), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Started));;
}

void QMovie_DisconnectStarted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::started), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Started));;
}

void QMovie_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QMovie::MovieState)>(&QMovie::stateChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QMovie::MovieState)>(&MyQMovie::Signal_StateChanged));;
}

void QMovie_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QMovie::MovieState)>(&QMovie::stateChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QMovie::MovieState)>(&MyQMovie::Signal_StateChanged));;
}

void QMovie_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "stop");
}

void QMovie_DestroyQMovie(QtObjectPtr ptr){
	static_cast<QMovie*>(ptr)->~QMovie();
}

