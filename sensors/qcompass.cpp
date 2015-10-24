#include "qcompass.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCompass>
#include "_cgo_export.h"

class MyQCompass: public QCompass {
public:
};

QtObjectPtr QCompass_Reading(QtObjectPtr ptr){
	return static_cast<QCompass*>(ptr)->reading();
}

QtObjectPtr QCompass_NewQCompass(QtObjectPtr parent){
	return new QCompass(static_cast<QObject*>(parent));
}

void QCompass_DestroyQCompass(QtObjectPtr ptr){
	static_cast<QCompass*>(ptr)->~QCompass();
}

