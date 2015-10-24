#include "qgyroscope.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QGyroscope>
#include "_cgo_export.h"

class MyQGyroscope: public QGyroscope {
public:
};

QtObjectPtr QGyroscope_Reading(QtObjectPtr ptr){
	return static_cast<QGyroscope*>(ptr)->reading();
}

QtObjectPtr QGyroscope_NewQGyroscope(QtObjectPtr parent){
	return new QGyroscope(static_cast<QObject*>(parent));
}

void QGyroscope_DestroyQGyroscope(QtObjectPtr ptr){
	static_cast<QGyroscope*>(ptr)->~QGyroscope();
}

