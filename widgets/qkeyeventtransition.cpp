#include "qkeyeventtransition.h"
#include <QKeyEvent>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QState>
#include <QKeyEventTransition>
#include "_cgo_export.h"

class MyQKeyEventTransition: public QKeyEventTransition {
public:
};

void* QKeyEventTransition_NewQKeyEventTransition2(void* object, int ty, int key, void* sourceState){
	return new QKeyEventTransition(static_cast<QObject*>(object), static_cast<QEvent::Type>(ty), key, static_cast<QState*>(sourceState));
}

void* QKeyEventTransition_NewQKeyEventTransition(void* sourceState){
	return new QKeyEventTransition(static_cast<QState*>(sourceState));
}

int QKeyEventTransition_Key(void* ptr){
	return static_cast<QKeyEventTransition*>(ptr)->key();
}

int QKeyEventTransition_ModifierMask(void* ptr){
	return static_cast<QKeyEventTransition*>(ptr)->modifierMask();
}

void QKeyEventTransition_SetKey(void* ptr, int key){
	static_cast<QKeyEventTransition*>(ptr)->setKey(key);
}

void QKeyEventTransition_SetModifierMask(void* ptr, int modifierMask){
	static_cast<QKeyEventTransition*>(ptr)->setModifierMask(static_cast<Qt::KeyboardModifier>(modifierMask));
}

void QKeyEventTransition_DestroyQKeyEventTransition(void* ptr){
	static_cast<QKeyEventTransition*>(ptr)->~QKeyEventTransition();
}

