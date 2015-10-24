#include "qabstractanimation.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QAbstractAnimation>
#include "_cgo_export.h"

class MyQAbstractAnimation: public QAbstractAnimation {
public:
void Signal_CurrentLoopChanged(int currentLoop){callbackQAbstractAnimationCurrentLoopChanged(this->objectName().toUtf8().data(), currentLoop);};
void Signal_DirectionChanged(QAbstractAnimation::Direction newDirection){callbackQAbstractAnimationDirectionChanged(this->objectName().toUtf8().data(), newDirection);};
void Signal_Finished(){callbackQAbstractAnimationFinished(this->objectName().toUtf8().data());};
void Signal_StateChanged(QAbstractAnimation::State newState, QAbstractAnimation::State oldState){callbackQAbstractAnimationStateChanged(this->objectName().toUtf8().data(), newState, oldState);};
};

int QAbstractAnimation_CurrentLoop(QtObjectPtr ptr){
	return static_cast<QAbstractAnimation*>(ptr)->currentLoop();
}

int QAbstractAnimation_CurrentTime(QtObjectPtr ptr){
	return static_cast<QAbstractAnimation*>(ptr)->currentTime();
}

int QAbstractAnimation_Direction(QtObjectPtr ptr){
	return static_cast<QAbstractAnimation*>(ptr)->direction();
}

int QAbstractAnimation_LoopCount(QtObjectPtr ptr){
	return static_cast<QAbstractAnimation*>(ptr)->loopCount();
}

void QAbstractAnimation_SetCurrentTime(QtObjectPtr ptr, int msecs){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "setCurrentTime", Q_ARG(int, msecs));
}

void QAbstractAnimation_SetDirection(QtObjectPtr ptr, int direction){
	static_cast<QAbstractAnimation*>(ptr)->setDirection(static_cast<QAbstractAnimation::Direction>(direction));
}

void QAbstractAnimation_SetLoopCount(QtObjectPtr ptr, int loopCount){
	static_cast<QAbstractAnimation*>(ptr)->setLoopCount(loopCount);
}

void QAbstractAnimation_ConnectCurrentLoopChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(int)>(&QAbstractAnimation::currentLoopChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(int)>(&MyQAbstractAnimation::Signal_CurrentLoopChanged));;
}

void QAbstractAnimation_DisconnectCurrentLoopChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(int)>(&QAbstractAnimation::currentLoopChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(int)>(&MyQAbstractAnimation::Signal_CurrentLoopChanged));;
}

int QAbstractAnimation_CurrentLoopTime(QtObjectPtr ptr){
	return static_cast<QAbstractAnimation*>(ptr)->currentLoopTime();
}

void QAbstractAnimation_ConnectDirectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::Direction)>(&QAbstractAnimation::directionChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::Direction)>(&MyQAbstractAnimation::Signal_DirectionChanged));;
}

void QAbstractAnimation_DisconnectDirectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::Direction)>(&QAbstractAnimation::directionChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::Direction)>(&MyQAbstractAnimation::Signal_DirectionChanged));;
}

int QAbstractAnimation_Duration(QtObjectPtr ptr){
	return static_cast<QAbstractAnimation*>(ptr)->duration();
}

void QAbstractAnimation_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)()>(&QAbstractAnimation::finished), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)()>(&MyQAbstractAnimation::Signal_Finished));;
}

void QAbstractAnimation_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)()>(&QAbstractAnimation::finished), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)()>(&MyQAbstractAnimation::Signal_Finished));;
}

QtObjectPtr QAbstractAnimation_Group(QtObjectPtr ptr){
	return static_cast<QAbstractAnimation*>(ptr)->group();
}

void QAbstractAnimation_Pause(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "pause");
}

void QAbstractAnimation_Resume(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "resume");
}

void QAbstractAnimation_SetPaused(QtObjectPtr ptr, int paused){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "setPaused", Q_ARG(bool, paused != 0));
}

void QAbstractAnimation_Start(QtObjectPtr ptr, int policy){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "start", Q_ARG(QAbstractAnimation::DeletionPolicy, static_cast<QAbstractAnimation::DeletionPolicy>(policy)));
}

void QAbstractAnimation_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&QAbstractAnimation::stateChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&MyQAbstractAnimation::Signal_StateChanged));;
}

void QAbstractAnimation_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractAnimation*>(ptr), static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&QAbstractAnimation::stateChanged), static_cast<MyQAbstractAnimation*>(ptr), static_cast<void (MyQAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&MyQAbstractAnimation::Signal_StateChanged));;
}

void QAbstractAnimation_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractAnimation*>(ptr), "stop");
}

int QAbstractAnimation_TotalDuration(QtObjectPtr ptr){
	return static_cast<QAbstractAnimation*>(ptr)->totalDuration();
}

void QAbstractAnimation_DestroyQAbstractAnimation(QtObjectPtr ptr){
	static_cast<QAbstractAnimation*>(ptr)->~QAbstractAnimation();
}

