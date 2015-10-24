#include "qhelpsearchengine.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpEngine>
#include <QHelpEngineCore>
#include <QHelpSearchEngine>
#include "_cgo_export.h"

class MyQHelpSearchEngine: public QHelpSearchEngine {
public:
void Signal_IndexingFinished(){callbackQHelpSearchEngineIndexingFinished(this->objectName().toUtf8().data());};
void Signal_IndexingStarted(){callbackQHelpSearchEngineIndexingStarted(this->objectName().toUtf8().data());};
void Signal_SearchingFinished(int hits){callbackQHelpSearchEngineSearchingFinished(this->objectName().toUtf8().data(), hits);};
void Signal_SearchingStarted(){callbackQHelpSearchEngineSearchingStarted(this->objectName().toUtf8().data());};
};

QtObjectPtr QHelpSearchEngine_NewQHelpSearchEngine(QtObjectPtr helpEngine, QtObjectPtr parent){
	return new QHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QObject*>(parent));
}

void QHelpSearchEngine_CancelIndexing(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelIndexing");
}

void QHelpSearchEngine_CancelSearching(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelSearching");
}

int QHelpSearchEngine_HitCount(QtObjectPtr ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->hitCount();
}

void QHelpSearchEngine_ConnectIndexingFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_DisconnectIndexingFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_ConnectIndexingStarted(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

void QHelpSearchEngine_DisconnectIndexingStarted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

QtObjectPtr QHelpSearchEngine_QueryWidget(QtObjectPtr ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->queryWidget();
}

void QHelpSearchEngine_ReindexDocumentation(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "reindexDocumentation");
}

QtObjectPtr QHelpSearchEngine_ResultWidget(QtObjectPtr ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->resultWidget();
}

void QHelpSearchEngine_ConnectSearchingFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_DisconnectSearchingFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_ConnectSearchingStarted(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_DisconnectSearchingStarted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_DestroyQHelpSearchEngine(QtObjectPtr ptr){
	static_cast<QHelpSearchEngine*>(ptr)->~QHelpSearchEngine();
}

