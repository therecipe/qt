#include "qtimeline.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QEasingCurve>
#include <QTime>
#include <QTimeLine>
#include "_cgo_export.h"

class MyQTimeLine: public QTimeLine {
public:
void Signal_Finished(){callbackQTimeLineFinished(this->objectName().toUtf8().data());};
void Signal_FrameChanged(int frame){callbackQTimeLineFrameChanged(this->objectName().toUtf8().data(), frame);};
void Signal_StateChanged(QTimeLine::State newState){callbackQTimeLineStateChanged(this->objectName().toUtf8().data(), newState);};
};

int QTimeLine_CurrentTime(void* ptr){
	return static_cast<QTimeLine*>(ptr)->currentTime();
}

int QTimeLine_CurveShape(void* ptr){
	return static_cast<QTimeLine*>(ptr)->curveShape();
}

int QTimeLine_Direction(void* ptr){
	return static_cast<QTimeLine*>(ptr)->direction();
}

int QTimeLine_Duration(void* ptr){
	return static_cast<QTimeLine*>(ptr)->duration();
}

void* QTimeLine_EasingCurve(void* ptr){
	return new QEasingCurve(static_cast<QTimeLine*>(ptr)->easingCurve());
}

int QTimeLine_LoopCount(void* ptr){
	return static_cast<QTimeLine*>(ptr)->loopCount();
}

void QTimeLine_SetCurrentTime(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "setCurrentTime", Q_ARG(int, msec));
}

void QTimeLine_SetCurveShape(void* ptr, int shape){
	static_cast<QTimeLine*>(ptr)->setCurveShape(static_cast<QTimeLine::CurveShape>(shape));
}

void QTimeLine_SetDirection(void* ptr, int direction){
	static_cast<QTimeLine*>(ptr)->setDirection(static_cast<QTimeLine::Direction>(direction));
}

void QTimeLine_SetDuration(void* ptr, int duration){
	static_cast<QTimeLine*>(ptr)->setDuration(duration);
}

void QTimeLine_SetEasingCurve(void* ptr, void* curve){
	static_cast<QTimeLine*>(ptr)->setEasingCurve(*static_cast<QEasingCurve*>(curve));
}

void QTimeLine_SetLoopCount(void* ptr, int count){
	static_cast<QTimeLine*>(ptr)->setLoopCount(count);
}

void QTimeLine_SetUpdateInterval(void* ptr, int interval){
	static_cast<QTimeLine*>(ptr)->setUpdateInterval(interval);
}

int QTimeLine_UpdateInterval(void* ptr){
	return static_cast<QTimeLine*>(ptr)->updateInterval();
}

void* QTimeLine_NewQTimeLine(int duration, void* parent){
	return new QTimeLine(duration, static_cast<QObject*>(parent));
}

int QTimeLine_CurrentFrame(void* ptr){
	return static_cast<QTimeLine*>(ptr)->currentFrame();
}

double QTimeLine_CurrentValue(void* ptr){
	return static_cast<double>(static_cast<QTimeLine*>(ptr)->currentValue());
}

int QTimeLine_EndFrame(void* ptr){
	return static_cast<QTimeLine*>(ptr)->endFrame();
}

void QTimeLine_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::finished, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)()>(&MyQTimeLine::Signal_Finished));;
}

void QTimeLine_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::finished, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)()>(&MyQTimeLine::Signal_Finished));;
}

void QTimeLine_ConnectFrameChanged(void* ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::frameChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(int)>(&MyQTimeLine::Signal_FrameChanged));;
}

void QTimeLine_DisconnectFrameChanged(void* ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::frameChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(int)>(&MyQTimeLine::Signal_FrameChanged));;
}

int QTimeLine_FrameForTime(void* ptr, int msec){
	return static_cast<QTimeLine*>(ptr)->frameForTime(msec);
}

void QTimeLine_Resume(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "resume");
}

void QTimeLine_SetEndFrame(void* ptr, int frame){
	static_cast<QTimeLine*>(ptr)->setEndFrame(frame);
}

void QTimeLine_SetFrameRange(void* ptr, int startFrame, int endFrame){
	static_cast<QTimeLine*>(ptr)->setFrameRange(startFrame, endFrame);
}

void QTimeLine_SetPaused(void* ptr, int paused){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "setPaused", Q_ARG(bool, paused != 0));
}

void QTimeLine_SetStartFrame(void* ptr, int frame){
	static_cast<QTimeLine*>(ptr)->setStartFrame(frame);
}

void QTimeLine_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "start");
}

int QTimeLine_StartFrame(void* ptr){
	return static_cast<QTimeLine*>(ptr)->startFrame();
}

void QTimeLine_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QTimeLine*>(ptr), &QTimeLine::stateChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(QTimeLine::State)>(&MyQTimeLine::Signal_StateChanged));;
}

void QTimeLine_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QTimeLine*>(ptr), &QTimeLine::stateChanged, static_cast<MyQTimeLine*>(ptr), static_cast<void (MyQTimeLine::*)(QTimeLine::State)>(&MyQTimeLine::Signal_StateChanged));;
}

void QTimeLine_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "stop");
}

void QTimeLine_ToggleDirection(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimeLine*>(ptr), "toggleDirection");
}

double QTimeLine_ValueForTime(void* ptr, int msec){
	return static_cast<double>(static_cast<QTimeLine*>(ptr)->valueForTime(msec));
}

void QTimeLine_DestroyQTimeLine(void* ptr){
	static_cast<QTimeLine*>(ptr)->~QTimeLine();
}

