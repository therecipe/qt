#include "qabstractanimation.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractAnimation>
#include "_cgo_export.h"

class MyQAbstractAnimation: public QAbstractAnimation {
public:
void Signal_CurrentLoopChanged(int currentLoop){callbackQAbstractAnimationCurrentLoopChanged(this->objectName().toUtf8().data(), currentLoop);};
void Signal_DirectionChanged(QAbstractAnimation::Direction newDirection){callbackQAbstractAnimationDirectionChanged(this->objectName().toUtf8().data(), newDirection);};
void Signal_Finished(){callbackQAbstractAnimationFinished(this->objectName().toUtf8().data());};
void Signal_StateChanged(QAbstractAnimation::State newState, QAbstractAnimation::State oldState){callbackQAbstractAnimationStateChanged(this->objectName().toUtf8().data(), newState, oldState);};
};

int QAbstractAnimation_CurrentLoop(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->currentLoop();
}

int QAbstractAnimation_CurrentTime(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->currentTime();
}

int QAbstractAnimation_Direction(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->direction();
}

int QAbstractAnimation_LoopCount(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->loopCount();
}

void QAbstractAnimation_SetCurrentTime(void* ptr, int msecs){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "setCurrentTime", Q_ARG(int, msecs));
}

void QAbstractAnimation_SetDirection(void* ptr, int direction){
	static_cast<QAbstractAnimation*>(ptr)->setDirection(static_cast<QAbstractAnimation::Direction>(direction));
}

void QAbstractAnimation_SetLoopCount(void* ptr, int loopCount){
	static_cast<QAbstractAnimation*>(ptr)->setLoopCount(loopCount);
}

void QAbstractAnimation_ConnectCurrentLoopChanged(void* ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(int)>(&QAbstractAnimation::currentLoopChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(int)>(&MyQAbstractAnimation::Signal_CurrentLoopChanged));;
}

void QAbstractAnimation_DisconnectCurrentLoopChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(int)>(&QAbstractAnimation::currentLoopChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(int)>(&MyQAbstractAnimation::Signal_CurrentLoopChanged));;
}

int QAbstractAnimation_CurrentLoopTime(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->currentLoopTime();
}

void QAbstractAnimation_ConnectDirectionChanged(void* ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::Direction)>(&QAbstractAnimation::directionChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::Direction)>(&MyQAbstractAnimation::Signal_DirectionChanged));;
}

void QAbstractAnimation_DisconnectDirectionChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::Direction)>(&QAbstractAnimation::directionChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::Direction)>(&MyQAbstractAnimation::Signal_DirectionChanged));;
}

int QAbstractAnimation_Duration(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->duration();
}

void QAbstractAnimation_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)()>(&QAbstractAnimation::finished), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)()>(&MyQAbstractAnimation::Signal_Finished));;
}

void QAbstractAnimation_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)()>(&QAbstractAnimation::finished), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)()>(&MyQAbstractAnimation::Signal_Finished));;
}

void* QAbstractAnimation_Group(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->group();
}

void QAbstractAnimation_Pause(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "pause");
}

void QAbstractAnimation_Resume(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "resume");
}

void QAbstractAnimation_SetPaused(void* ptr, int paused){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "setPaused", Q_ARG(bool, paused != 0));
}

void QAbstractAnimation_Start(void* ptr, int policy){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "start", Q_ARG(QAbstractAnimation::DeletionPolicy, static_cast<QAbstractAnimation::DeletionPolicy>(policy)));
}

void QAbstractAnimation_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&QAbstractAnimation::stateChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&MyQAbstractAnimation::Signal_StateChanged));;
}

void QAbstractAnimation_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&QAbstractAnimation::stateChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&MyQAbstractAnimation::Signal_StateChanged));;
}

void QAbstractAnimation_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "stop");
}

int QAbstractAnimation_TotalDuration(void* ptr){
	return static_cast<QAbstractAnimation*>(ptr)->totalDuration();
}

void QAbstractAnimation_DestroyQAbstractAnimation(void* ptr){
	static_cast<QAbstractAnimation*>(ptr)->~QAbstractAnimation();
}

