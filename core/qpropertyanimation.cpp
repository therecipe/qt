#include "qpropertyanimation.h"
#include <QModelIndex>
#include <QByteArray>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPropertyAnimation>
#include "_cgo_export.h"

class MyQPropertyAnimation: public QPropertyAnimation {
public:
};

void* QPropertyAnimation_PropertyName(void* ptr){
	return new QByteArray(static_cast<QPropertyAnimation*>(ptr)->propertyName());
}

void QPropertyAnimation_SetPropertyName(void* ptr, void* propertyName){
	static_cast<QPropertyAnimation*>(ptr)->setPropertyName(*static_cast<QByteArray*>(propertyName));
}

void QPropertyAnimation_SetTargetObject(void* ptr, void* target){
	static_cast<QPropertyAnimation*>(ptr)->setTargetObject(static_cast<QObject*>(target));
}

void* QPropertyAnimation_TargetObject(void* ptr){
	return static_cast<QPropertyAnimation*>(ptr)->targetObject();
}

void* QPropertyAnimation_NewQPropertyAnimation(void* parent){
	return new QPropertyAnimation(static_cast<QObject*>(parent));
}

void* QPropertyAnimation_NewQPropertyAnimation2(void* target, void* propertyName, void* parent){
	return new QPropertyAnimation(static_cast<QObject*>(target), *static_cast<QByteArray*>(propertyName), static_cast<QObject*>(parent));
}

void QPropertyAnimation_DestroyQPropertyAnimation(void* ptr){
	static_cast<QPropertyAnimation*>(ptr)->~QPropertyAnimation();
}

