#include "qqmlpropertyvaluesource.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlProperty>
#include <QQmlPropertyValueSource>
#include "_cgo_export.h"

class MyQQmlPropertyValueSource: public QQmlPropertyValueSource {
public:
};

void QQmlPropertyValueSource_SetTarget(void* ptr, void* property){
	static_cast<QQmlPropertyValueSource*>(ptr)->setTarget(*static_cast<QQmlProperty*>(property));
}

void QQmlPropertyValueSource_DestroyQQmlPropertyValueSource(void* ptr){
	static_cast<QQmlPropertyValueSource*>(ptr)->~QQmlPropertyValueSource();
}

