#include "qtimeline.h"
#include <QUrl>
#include <QModelIndex>
#include <QEasingCurve>
#include <QMetaObject>
#include <QTime>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QTimeLine>
#include "_cgo_export.h"

class MyQTimeLine: public QTimeLine {
public:
void Signal_Finished(){callbackQTimeLineFinished(this->objectName().toUtf8().data());};
void Signal_FrameChanged(int frame){callbackQTimeLineFrameChanged(this->objectName().toUtf8().data(), frame);};
void Signal_StateChanged(QTimeLine::State newState){callbackQTimeLineStateChanged(this->objectName().toUtf8().data(), newState);};
};

int QTimeLine_CurrentTime(QtObjectPtr ptr){
	return static_cast<QTimeLine*>(ptr)->currentTime();
}

int QTimeLine_CurveShape(QtObjectPtr ptr){
	return static_cast<QTimeLine*>(ptr)->curveShape();
}

int QTimeLine_Direction(QtObjectPtr ptr){
	return static_cast<QTimeLine*>(ptr)->direction();
}

int QTimeLine_Duration(QtObjectPtr ptr){
	return static_cast<QTimeLine*>(ptr)->duration();
}

int QTimeLine_LoopCount(QtObjectPtr ptr){
	return static_cast<QTimeLine*>(ptr)->loopCount();
}

void QTimeLine_SetCurrentTime(QtObjectPtr ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "setCurrentTime", Q_ARG(int, msec));
}

void QTimeLine_SetCurveShape(QtObjectPtr ptr, int shape){
	static_cast<QTimeLine*>(ptr)->setCurveShape(static_cast<QTimeLine::CurveShape>(shape));
}

void QTimeLine_SetDirection(QtObjectPtr ptr, int direction){
	static_cast<QTimeLine*>(ptr)->setDirection(static_cast<QTimeLine::Direction>(direction));
}

void QTimeLine_SetDuration(QtObjectPtr ptr, int duration){
	static_cast<QTimeLine*>(ptr)->setDuration(duration);
}

void QTimeLine_SetEasingCurve(QtObjectPtr ptr, QtObjectPtr curve){
	static_cast<QTimeLine*>(ptr)->setEasingCurve(*static_cast<QEasingCurve*>(curve));
}

void QTimeLine_SetLoopCount(QtObjectPtr ptr, int count){
	static_cast<QTimeLine*>(ptr)->setLoopCount(count);
}

void QTimeLine_SetUpdateInterval(QtObjectPtr ptr, int interval){
	static_cast<QTimeLine*>(ptr)->setUpdateInterval(interval);
}

int QTimeLine_UpdateInterval(QtObjectPtr ptr){
	return static_cast<QTimeLine*>(ptr)->updateInterval();
}

QtObjectPtr QTimeLine_NewQTimeLine(int duration, QtObjectPtr parent){
	return new QTimeLine(duration, static_cast<QObject*>(parent));
}

int QTimeLine_CurrentFrame(QtObjectPtr ptr){
	return static_cast<QTimeLine*>(ptr)->currentFrame();
}

int QTimeLine_EndFrame(QtObjectPtr ptr){
	return static_cast<QTimeLine*>(ptr)->endFrame();
}

void QTimeLine_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::finished, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)()>(&MyQTimeLine::Signal_Finished));;
}

void QTimeLine_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::finished, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)()>(&MyQTimeLine::Signal_Finished));;
}

void QTimeLine_ConnectFrameChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::frameChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(int)>(&MyQTimeLine::Signal_FrameChanged));;
}

void QTimeLine_DisconnectFrameChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::frameChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(int)>(&MyQTimeLine::Signal_FrameChanged));;
}

int QTimeLine_FrameForTime(QtObjectPtr ptr, int msec){
	return static_cast<QTimeLine*>(ptr)->frameForTime(msec);
}

void QTimeLine_Resume(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "resume");
}

void QTimeLine_SetEndFrame(QtObjectPtr ptr, int frame){
	static_cast<QTimeLine*>(ptr)->setEndFrame(frame);
}

void QTimeLine_SetFrameRange(QtObjectPtr ptr, int startFrame, int endFrame){
	static_cast<QTimeLine*>(ptr)->setFrameRange(startFrame, endFrame);
}

void QTimeLine_SetPaused(QtObjectPtr ptr, int paused){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "setPaused", Q_ARG(bool, paused != 0));
}

void QTimeLine_SetStartFrame(QtObjectPtr ptr, int frame){
	static_cast<QTimeLine*>(ptr)->setStartFrame(frame);
}

void QTimeLine_Start(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "start");
}

int QTimeLine_StartFrame(QtObjectPtr ptr){
	return static_cast<QTimeLine*>(ptr)->startFrame();
}

void QTimeLine_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::stateChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(QTimeLine::State)>(&MyQTimeLine::Signal_StateChanged));;
}

void QTimeLine_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::stateChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(QTimeLine::State)>(&MyQTimeLine::Signal_StateChanged));;
}

void QTimeLine_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "stop");
}

void QTimeLine_ToggleDirection(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "toggleDirection");
}

void QTimeLine_DestroyQTimeLine(QtObjectPtr ptr){
	static_cast<QTimeLine*>(ptr)->~QTimeLine();
}

