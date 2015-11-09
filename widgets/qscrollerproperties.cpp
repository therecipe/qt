#include "qscrollerproperties.h"
#include <QUrl>
#include <QModelIndex>
#include <QScroller>
#include <QString>
#include <QVariant>
#include <QScrollerProperties>
#include "_cgo_export.h"

class MyQScrollerProperties: public QScrollerProperties {
public:
};

void* QScrollerProperties_NewQScrollerProperties(){
	return new QScrollerProperties();
}

void* QScrollerProperties_NewQScrollerProperties2(void* sp){
	return new QScrollerProperties(*static_cast<QScrollerProperties*>(sp));
}

void* QScrollerProperties_ScrollMetric(void* ptr, int metric){
	return new QVariant(static_cast<QScrollerProperties*>(ptr)->scrollMetric(static_cast<QScrollerProperties::ScrollMetric>(metric)));
}

void QScrollerProperties_QScrollerProperties_SetDefaultScrollerProperties(void* sp){
	QScrollerProperties::setDefaultScrollerProperties(*static_cast<QScrollerProperties*>(sp));
}

void QScrollerProperties_SetScrollMetric(void* ptr, int metric, void* value){
	static_cast<QScrollerProperties*>(ptr)->setScrollMetric(static_cast<QScrollerProperties::ScrollMetric>(metric), *static_cast<QVariant*>(value));
}

void QScrollerProperties_QScrollerProperties_UnsetDefaultScrollerProperties(){
	QScrollerProperties::unsetDefaultScrollerProperties();
}

void QScrollerProperties_DestroyQScrollerProperties(void* ptr){
	static_cast<QScrollerProperties*>(ptr)->~QScrollerProperties();
}

