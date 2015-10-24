#include "qmouseeventtransition.h"
#include <QUrl>
#include <QMouseEvent>
#include <QPainter>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QEvent>
#include <QState>
#include <QPainterPath>
#include <QMouseEventTransition>
#include "_cgo_export.h"

class MyQMouseEventTransition: public QMouseEventTransition {
public:
};

QtObjectPtr QMouseEventTransition_NewQMouseEventTransition2(QtObjectPtr object, int ty, int button, QtObjectPtr sourceState){
	return new QMouseEventTransition(static_cast<QObject*>(object), static_cast<QEvent::Type>(ty), static_cast<Qt::MouseButton>(button), static_cast<QState*>(sourceState));
}

QtObjectPtr QMouseEventTransition_NewQMouseEventTransition(QtObjectPtr sourceState){
	return new QMouseEventTransition(static_cast<QState*>(sourceState));
}

int QMouseEventTransition_Button(QtObjectPtr ptr){
	return static_cast<QMouseEventTransition*>(ptr)->button();
}

int QMouseEventTransition_ModifierMask(QtObjectPtr ptr){
	return static_cast<QMouseEventTransition*>(ptr)->modifierMask();
}

void QMouseEventTransition_SetButton(QtObjectPtr ptr, int button){
	static_cast<QMouseEventTransition*>(ptr)->setButton(static_cast<Qt::MouseButton>(button));
}

void QMouseEventTransition_SetHitTestPath(QtObjectPtr ptr, QtObjectPtr path){
	static_cast<QMouseEventTransition*>(ptr)->setHitTestPath(*static_cast<QPainterPath*>(path));
}

void QMouseEventTransition_SetModifierMask(QtObjectPtr ptr, int modifierMask){
	static_cast<QMouseEventTransition*>(ptr)->setModifierMask(static_cast<Qt::KeyboardModifier>(modifierMask));
}

void QMouseEventTransition_DestroyQMouseEventTransition(QtObjectPtr ptr){
	static_cast<QMouseEventTransition*>(ptr)->~QMouseEventTransition();
}

