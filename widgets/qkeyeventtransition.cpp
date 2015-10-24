#include "qkeyeventtransition.h"
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QState>
#include <QObject>
#include <QKeyEvent>
#include <QString>
#include <QVariant>
#include <QKeyEventTransition>
#include "_cgo_export.h"

class MyQKeyEventTransition: public QKeyEventTransition {
public:
};

QtObjectPtr QKeyEventTransition_NewQKeyEventTransition2(QtObjectPtr object, int ty, int key, QtObjectPtr sourceState){
	return new QKeyEventTransition(static_cast<QObject*>(object), static_cast<QEvent::Type>(ty), key, static_cast<QState*>(sourceState));
}

QtObjectPtr QKeyEventTransition_NewQKeyEventTransition(QtObjectPtr sourceState){
	return new QKeyEventTransition(static_cast<QState*>(sourceState));
}

int QKeyEventTransition_Key(QtObjectPtr ptr){
	return static_cast<QKeyEventTransition*>(ptr)->key();
}

int QKeyEventTransition_ModifierMask(QtObjectPtr ptr){
	return static_cast<QKeyEventTransition*>(ptr)->modifierMask();
}

void QKeyEventTransition_SetKey(QtObjectPtr ptr, int key){
	static_cast<QKeyEventTransition*>(ptr)->setKey(key);
}

void QKeyEventTransition_SetModifierMask(QtObjectPtr ptr, int modifierMask){
	static_cast<QKeyEventTransition*>(ptr)->setModifierMask(static_cast<Qt::KeyboardModifier>(modifierMask));
}

void QKeyEventTransition_DestroyQKeyEventTransition(QtObjectPtr ptr){
	static_cast<QKeyEventTransition*>(ptr)->~QKeyEventTransition();
}

