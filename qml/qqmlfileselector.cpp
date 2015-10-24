#include "qqmlfileselector.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFileSelector>
#include <QQmlEngine>
#include <QObject>
#include <QFile>
#include <QString>
#include <QQmlFileSelector>
#include "_cgo_export.h"

class MyQQmlFileSelector: public QQmlFileSelector {
public:
};

QtObjectPtr QQmlFileSelector_NewQQmlFileSelector(QtObjectPtr engine, QtObjectPtr parent){
	return new QQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

QtObjectPtr QQmlFileSelector_QQmlFileSelector_Get(QtObjectPtr engine){
	return QQmlFileSelector::get(static_cast<QQmlEngine*>(engine));
}

void QQmlFileSelector_SetExtraSelectors(QtObjectPtr ptr, char* strin){
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(QString(strin).split("|", QString::SkipEmptyParts));
}

void QQmlFileSelector_SetExtraSelectors2(QtObjectPtr ptr, char* strin){
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(QString(strin).split("|", QString::SkipEmptyParts));
}

void QQmlFileSelector_SetSelector(QtObjectPtr ptr, QtObjectPtr selector){
	static_cast<QQmlFileSelector*>(ptr)->setSelector(static_cast<QFileSelector*>(selector));
}

void QQmlFileSelector_DestroyQQmlFileSelector(QtObjectPtr ptr){
	static_cast<QQmlFileSelector*>(ptr)->~QQmlFileSelector();
}

