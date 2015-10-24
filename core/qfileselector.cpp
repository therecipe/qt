#include "qfileselector.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QFile>
#include <QFileSelector>
#include "_cgo_export.h"

class MyQFileSelector: public QFileSelector {
public:
};

QtObjectPtr QFileSelector_NewQFileSelector(QtObjectPtr parent){
	return new QFileSelector(static_cast<QObject*>(parent));
}

char* QFileSelector_AllSelectors(QtObjectPtr ptr){
	return static_cast<QFileSelector*>(ptr)->allSelectors().join("|").toUtf8().data();
}

char* QFileSelector_ExtraSelectors(QtObjectPtr ptr){
	return static_cast<QFileSelector*>(ptr)->extraSelectors().join("|").toUtf8().data();
}

char* QFileSelector_Select(QtObjectPtr ptr, char* filePath){
	return static_cast<QFileSelector*>(ptr)->select(QString(filePath)).toUtf8().data();
}

char* QFileSelector_Select2(QtObjectPtr ptr, char* filePath){
	return static_cast<QFileSelector*>(ptr)->select(QUrl(QString(filePath))).toString().toUtf8().data();
}

void QFileSelector_SetExtraSelectors(QtObjectPtr ptr, char* list){
	static_cast<QFileSelector*>(ptr)->setExtraSelectors(QString(list).split("|", QString::SkipEmptyParts));
}

void QFileSelector_DestroyQFileSelector(QtObjectPtr ptr){
	static_cast<QFileSelector*>(ptr)->~QFileSelector();
}

