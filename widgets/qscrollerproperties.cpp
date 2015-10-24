#include "qscrollerproperties.h"
#include <QModelIndex>
#include <QScroller>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QScrollerProperties>
#include "_cgo_export.h"

class MyQScrollerProperties: public QScrollerProperties {
public:
};

QtObjectPtr QScrollerProperties_NewQScrollerProperties(){
	return new QScrollerProperties();
}

QtObjectPtr QScrollerProperties_NewQScrollerProperties2(QtObjectPtr sp){
	return new QScrollerProperties(*static_cast<QScrollerProperties*>(sp));
}

char* QScrollerProperties_ScrollMetric(QtObjectPtr ptr, int metric){
	return static_cast<QScrollerProperties*>(ptr)->scrollMetric(static_cast<QScrollerProperties::ScrollMetric>(metric)).toString().toUtf8().data();
}

void QScrollerProperties_QScrollerProperties_SetDefaultScrollerProperties(QtObjectPtr sp){
	QScrollerProperties::setDefaultScrollerProperties(*static_cast<QScrollerProperties*>(sp));
}

void QScrollerProperties_SetScrollMetric(QtObjectPtr ptr, int metric, char* value){
	static_cast<QScrollerProperties*>(ptr)->setScrollMetric(static_cast<QScrollerProperties::ScrollMetric>(metric), QVariant(value));
}

void QScrollerProperties_QScrollerProperties_UnsetDefaultScrollerProperties(){
	QScrollerProperties::unsetDefaultScrollerProperties();
}

void QScrollerProperties_DestroyQScrollerProperties(QtObjectPtr ptr){
	static_cast<QScrollerProperties*>(ptr)->~QScrollerProperties();
}

