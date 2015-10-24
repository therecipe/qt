#include "qscroller.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QScrollerProperties>
#include <QMetaObject>
#include <QObject>
#include <QScroller>
#include "_cgo_export.h"

class MyQScroller: public QScroller {
public:
void Signal_StateChanged(QScroller::State newState){callbackQScrollerStateChanged(this->objectName().toUtf8().data(), newState);};
};

void QScroller_SetScrollerProperties(QtObjectPtr ptr, QtObjectPtr prop){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "setScrollerProperties", Q_ARG(QScrollerProperties, *static_cast<QScrollerProperties*>(prop)));
}

int QScroller_QScroller_GrabGesture(QtObjectPtr target, int scrollGestureType){
	return QScroller::grabGesture(static_cast<QObject*>(target), static_cast<QScroller::ScrollerGestureType>(scrollGestureType));
}

int QScroller_QScroller_GrabbedGesture(QtObjectPtr target){
	return QScroller::grabbedGesture(static_cast<QObject*>(target));
}

int QScroller_QScroller_HasScroller(QtObjectPtr target){
	return QScroller::hasScroller(static_cast<QObject*>(target));
}

void QScroller_ResendPrepareEvent(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "resendPrepareEvent");
}

void QScroller_ScrollTo(QtObjectPtr ptr, QtObjectPtr pos){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "scrollTo", Q_ARG(QPointF, *static_cast<QPointF*>(pos)));
}

void QScroller_ScrollTo2(QtObjectPtr ptr, QtObjectPtr pos, int scrollTime){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "scrollTo", Q_ARG(QPointF, *static_cast<QPointF*>(pos)), Q_ARG(int, scrollTime));
}

QtObjectPtr QScroller_QScroller_Scroller(QtObjectPtr target){
	return QScroller::scroller(static_cast<QObject*>(target));
}

QtObjectPtr QScroller_QScroller_Scroller2(QtObjectPtr target){
	return const_cast<QScroller*>(QScroller::scroller(static_cast<QObject*>(target)));
}

void QScroller_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QScroller*>(ptr), static_cast<void (QScroller::*)(QScroller::State)>(&QScroller::stateChanged), static_cast<MyQScroller*>(ptr), static_cast<void (MyQScroller::*)(QScroller::State)>(&MyQScroller::Signal_StateChanged));;
}

void QScroller_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QScroller*>(ptr), static_cast<void (QScroller::*)(QScroller::State)>(&QScroller::stateChanged), static_cast<MyQScroller*>(ptr), static_cast<void (MyQScroller::*)(QScroller::State)>(&MyQScroller::Signal_StateChanged));;
}

void QScroller_Stop(QtObjectPtr ptr){
	static_cast<QScroller*>(ptr)->stop();
}

QtObjectPtr QScroller_Target(QtObjectPtr ptr){
	return static_cast<QScroller*>(ptr)->target();
}

void QScroller_QScroller_UngrabGesture(QtObjectPtr target){
	QScroller::ungrabGesture(static_cast<QObject*>(target));
}

