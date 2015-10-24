#include "qpropertyanimation.h"
#include <QObject>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPropertyAnimation>
#include "_cgo_export.h"

class MyQPropertyAnimation: public QPropertyAnimation {
public:
};

void QPropertyAnimation_SetPropertyName(QtObjectPtr ptr, QtObjectPtr propertyName){
	static_cast<QPropertyAnimation*>(ptr)->setPropertyName(*static_cast<QByteArray*>(propertyName));
}

void QPropertyAnimation_SetTargetObject(QtObjectPtr ptr, QtObjectPtr target){
	static_cast<QPropertyAnimation*>(ptr)->setTargetObject(static_cast<QObject*>(target));
}

QtObjectPtr QPropertyAnimation_TargetObject(QtObjectPtr ptr){
	return static_cast<QPropertyAnimation*>(ptr)->targetObject();
}

QtObjectPtr QPropertyAnimation_NewQPropertyAnimation(QtObjectPtr parent){
	return new QPropertyAnimation(static_cast<QObject*>(parent));
}

QtObjectPtr QPropertyAnimation_NewQPropertyAnimation2(QtObjectPtr target, QtObjectPtr propertyName, QtObjectPtr parent){
	return new QPropertyAnimation(static_cast<QObject*>(target), *static_cast<QByteArray*>(propertyName), static_cast<QObject*>(parent));
}

void QPropertyAnimation_DestroyQPropertyAnimation(QtObjectPtr ptr){
	static_cast<QPropertyAnimation*>(ptr)->~QPropertyAnimation();
}

