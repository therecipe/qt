#include "qhelpengine.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QHelpEngine>
#include "_cgo_export.h"

class MyQHelpEngine: public QHelpEngine {
public:
};

QtObjectPtr QHelpEngine_NewQHelpEngine(char* collectionFile, QtObjectPtr parent){
	return new QHelpEngine(QString(collectionFile), static_cast<QObject*>(parent));
}

QtObjectPtr QHelpEngine_ContentModel(QtObjectPtr ptr){
	return static_cast<QHelpEngine*>(ptr)->contentModel();
}

QtObjectPtr QHelpEngine_ContentWidget(QtObjectPtr ptr){
	return static_cast<QHelpEngine*>(ptr)->contentWidget();
}

QtObjectPtr QHelpEngine_IndexModel(QtObjectPtr ptr){
	return static_cast<QHelpEngine*>(ptr)->indexModel();
}

QtObjectPtr QHelpEngine_IndexWidget(QtObjectPtr ptr){
	return static_cast<QHelpEngine*>(ptr)->indexWidget();
}

QtObjectPtr QHelpEngine_SearchEngine(QtObjectPtr ptr){
	return static_cast<QHelpEngine*>(ptr)->searchEngine();
}

void QHelpEngine_DestroyQHelpEngine(QtObjectPtr ptr){
	static_cast<QHelpEngine*>(ptr)->~QHelpEngine();
}

