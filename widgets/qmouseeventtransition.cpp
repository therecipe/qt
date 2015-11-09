#include "qmouseeventtransition.h"
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QEvent>
#include <QVariant>
#include <QState>
#include <QPainterPath>
#include <QPainter>
#include <QMouseEvent>
#include <QMouseEventTransition>
#include "_cgo_export.h"

class MyQMouseEventTransition: public QMouseEventTransition {
public:
};

void* QMouseEventTransition_NewQMouseEventTransition2(void* object, int ty, int button, void* sourceState){
	return new QMouseEventTransition(static_cast<QObject*>(object), static_cast<QEvent::Type>(ty), static_cast<Qt::MouseButton>(button), static_cast<QState*>(sourceState));
}

void* QMouseEventTransition_NewQMouseEventTransition(void* sourceState){
	return new QMouseEventTransition(static_cast<QState*>(sourceState));
}

int QMouseEventTransition_Button(void* ptr){
	return static_cast<QMouseEventTransition*>(ptr)->button();
}

int QMouseEventTransition_ModifierMask(void* ptr){
	return static_cast<QMouseEventTransition*>(ptr)->modifierMask();
}

void QMouseEventTransition_SetButton(void* ptr, int button){
	static_cast<QMouseEventTransition*>(ptr)->setButton(static_cast<Qt::MouseButton>(button));
}

void QMouseEventTransition_SetHitTestPath(void* ptr, void* path){
	static_cast<QMouseEventTransition*>(ptr)->setHitTestPath(*static_cast<QPainterPath*>(path));
}

void QMouseEventTransition_SetModifierMask(void* ptr, int modifierMask){
	static_cast<QMouseEventTransition*>(ptr)->setModifierMask(static_cast<Qt::KeyboardModifier>(modifierMask));
}

void QMouseEventTransition_DestroyQMouseEventTransition(void* ptr){
	static_cast<QMouseEventTransition*>(ptr)->~QMouseEventTransition();
}

