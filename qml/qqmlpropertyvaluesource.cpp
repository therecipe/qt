#include "qqmlpropertyvaluesource.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlProperty>
#include <QString>
#include <QQmlPropertyValueSource>
#include "_cgo_export.h"

class MyQQmlPropertyValueSource: public QQmlPropertyValueSource {
public:
};

void QQmlPropertyValueSource_SetTarget(QtObjectPtr ptr, QtObjectPtr property){
	static_cast<QQmlPropertyValueSource*>(ptr)->setTarget(*static_cast<QQmlProperty*>(property));
}

void QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(QtObjectPtr ptr){
	static_cast<QQmlPropertyValueSource*>(ptr)->~QQmlPropertyValueSource();
}

