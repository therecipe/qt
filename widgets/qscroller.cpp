#include "qscroller.h"
#include <QVariant>
#include <QUrl>
#include <QPoint>
#include <QObject>
#include <QRect>
#include <QMetaObject>
#include <QScrollerProperties>
#include <QString>
#include <QRectF>
#include <QPointF>
#include <QModelIndex>
#include <QScroller>
#include "_cgo_export.h"

class MyQScroller: public QScroller {
public:
void Signal_StateChanged(QScroller::State newState){callbackQScrollerStateChanged(this->objectName().toUtf8().data(), newState);};
};

void QScroller_SetScrollerProperties(void* ptr, void* prop){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "setScrollerProperties", Q_ARG(QScrollerProperties, *static_cast<QScrollerProperties*>(prop)));
}

void QScroller_EnsureVisible(void* ptr, void* rect, double xmargin, double ymargin){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "ensureVisible", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(qreal, static_cast<qreal>(xmargin)), Q_ARG(qreal, static_cast<qreal>(ymargin)));
}

void QScroller_EnsureVisible2(void* ptr, void* rect, double xmargin, double ymargin, int scrollTime){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "ensureVisible", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(qreal, static_cast<qreal>(xmargin)), Q_ARG(qreal, static_cast<qreal>(ymargin)), Q_ARG(int, scrollTime));
}

int QScroller_QScroller_GrabGesture(void* target, int scrollGestureType){
	return QScroller::grabGesture(static_cast<QObject*>(target), static_cast<QScroller::ScrollerGestureType>(scrollGestureType));
}

int QScroller_QScroller_GrabbedGesture(void* target){
	return QScroller::grabbedGesture(static_cast<QObject*>(target));
}

int QScroller_QScroller_HasScroller(void* target){
	return QScroller::hasScroller(static_cast<QObject*>(target));
}

void QScroller_ResendPrepareEvent(void* ptr){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "resendPrepareEvent");
}

void QScroller_ScrollTo(void* ptr, void* pos){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "scrollTo", Q_ARG(QPointF, *static_cast<QPointF*>(pos)));
}

void QScroller_ScrollTo2(void* ptr, void* pos, int scrollTime){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "scrollTo", Q_ARG(QPointF, *static_cast<QPointF*>(pos)), Q_ARG(int, scrollTime));
}

void* QScroller_QScroller_Scroller(void* target){
	return QScroller::scroller(static_cast<QObject*>(target));
}

void* QScroller_QScroller_Scroller2(void* target){
	return const_cast<QScroller*>(QScroller::scroller(static_cast<QObject*>(target)));
}

void QScroller_SetSnapPositionsX2(void* ptr, double first, double interval){
	static_cast<QScroller*>(ptr)->setSnapPositionsX(static_cast<qreal>(first), static_cast<qreal>(interval));
}

void QScroller_SetSnapPositionsY2(void* ptr, double first, double interval){
	static_cast<QScroller*>(ptr)->setSnapPositionsY(static_cast<qreal>(first), static_cast<qreal>(interval));
}

void QScroller_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QScroller*>(ptr), static_cast<void (QScroller::*)(QScroller::State)>(&QScroller::stateChanged), static_cast<MyQScroller*>(ptr), static_cast<void (MyQScroller::*)(QScroller::State)>(&MyQScroller::Signal_StateChanged));;
}

void QScroller_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QScroller*>(ptr), static_cast<void (QScroller::*)(QScroller::State)>(&QScroller::stateChanged), static_cast<MyQScroller*>(ptr), static_cast<void (MyQScroller::*)(QScroller::State)>(&MyQScroller::Signal_StateChanged));;
}

void QScroller_Stop(void* ptr){
	static_cast<QScroller*>(ptr)->stop();
}

void* QScroller_Target(void* ptr){
	return static_cast<QScroller*>(ptr)->target();
}

void QScroller_QScroller_UngrabGesture(void* target){
	QScroller::ungrabGesture(static_cast<QObject*>(target));
}

