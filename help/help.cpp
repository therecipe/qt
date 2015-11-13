#include "qhelpsearchquerywidget.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpSearchQuery>
#include <QObject>
#include <QHelpSearchQueryWidget>
#include "_cgo_export.h"

class MyQHelpSearchQueryWidget: public QHelpSearchQueryWidget {
public:
void Signal_Search(){callbackQHelpSearchQueryWidgetSearch(this->objectName().toUtf8().data());};
};

int QHelpSearchQueryWidget_IsCompactMode(void* ptr){
	return static_cast<QHelpSearchQueryWidget*>(ptr)->isCompactMode();
}

void* QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(void* parent){
	return new QHelpSearchQueryWidget(static_cast<QWidget*>(parent));
}

void QHelpSearchQueryWidget_CollapseExtendedSearch(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->collapseExtendedSearch();
}

void QHelpSearchQueryWidget_ExpandExtendedSearch(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->expandExtendedSearch();
}

void QHelpSearchQueryWidget_ConnectSearch(void* ptr){
	QObject::connect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_DisconnectSearch(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->~QHelpSearchQueryWidget();
}

#include "qhelpenginecore.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpEngine>
#include <QObject>
#include <QByteArray>
#include <QString>
#include <QHelpEngineCore>
#include "_cgo_export.h"

class MyQHelpEngineCore: public QHelpEngineCore {
public:
void Signal_CurrentFilterChanged(const QString & newFilter){callbackQHelpEngineCoreCurrentFilterChanged(this->objectName().toUtf8().data(), newFilter.toUtf8().data());};
void Signal_ReadersAboutToBeInvalidated(){callbackQHelpEngineCoreReadersAboutToBeInvalidated(this->objectName().toUtf8().data());};
void Signal_SetupFinished(){callbackQHelpEngineCoreSetupFinished(this->objectName().toUtf8().data());};
void Signal_SetupStarted(){callbackQHelpEngineCoreSetupStarted(this->objectName().toUtf8().data());};
void Signal_Warning(const QString & msg){callbackQHelpEngineCoreWarning(this->objectName().toUtf8().data(), msg.toUtf8().data());};
};

int QHelpEngineCore_AutoSaveFilter(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->autoSaveFilter();
}

char* QHelpEngineCore_CollectionFile(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->collectionFile().toUtf8().data();
}

char* QHelpEngineCore_CurrentFilter(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->currentFilter().toUtf8().data();
}

void QHelpEngineCore_SetAutoSaveFilter(void* ptr, int save){
	static_cast<QHelpEngineCore*>(ptr)->setAutoSaveFilter(save != 0);
}

void QHelpEngineCore_SetCollectionFile(void* ptr, char* fileName){
	static_cast<QHelpEngineCore*>(ptr)->setCollectionFile(QString(fileName));
}

void QHelpEngineCore_SetCurrentFilter(void* ptr, char* filterName){
	static_cast<QHelpEngineCore*>(ptr)->setCurrentFilter(QString(filterName));
}

void* QHelpEngineCore_NewQHelpEngineCore(char* collectionFile, void* parent){
	return new QHelpEngineCore(QString(collectionFile), static_cast<QObject*>(parent));
}

int QHelpEngineCore_AddCustomFilter(void* ptr, char* filterName, char* attributes){
	return static_cast<QHelpEngineCore*>(ptr)->addCustomFilter(QString(filterName), QString(attributes).split("|", QString::SkipEmptyParts));
}

int QHelpEngineCore_CopyCollectionFile(void* ptr, char* fileName){
	return static_cast<QHelpEngineCore*>(ptr)->copyCollectionFile(QString(fileName));
}

void QHelpEngineCore_ConnectCurrentFilterChanged(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));;
}

void QHelpEngineCore_DisconnectCurrentFilterChanged(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));;
}

char* QHelpEngineCore_CustomFilters(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->customFilters().join("|").toUtf8().data();
}

void* QHelpEngineCore_CustomValue(void* ptr, char* key, void* defaultValue){
	return new QVariant(static_cast<QHelpEngineCore*>(ptr)->customValue(QString(key), *static_cast<QVariant*>(defaultValue)));
}

char* QHelpEngineCore_DocumentationFileName(void* ptr, char* namespaceName){
	return static_cast<QHelpEngineCore*>(ptr)->documentationFileName(QString(namespaceName)).toUtf8().data();
}

char* QHelpEngineCore_Error(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->error().toUtf8().data();
}

void* QHelpEngineCore_FileData(void* ptr, void* url){
	return new QByteArray(static_cast<QHelpEngineCore*>(ptr)->fileData(*static_cast<QUrl*>(url)));
}

char* QHelpEngineCore_FilterAttributes(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes().join("|").toUtf8().data();
}

char* QHelpEngineCore_FilterAttributes2(void* ptr, char* filterName){
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes(QString(filterName)).join("|").toUtf8().data();
}

void* QHelpEngineCore_QHelpEngineCore_MetaData(char* documentationFileName, char* name){
	return new QVariant(QHelpEngineCore::metaData(QString(documentationFileName), QString(name)));
}

char* QHelpEngineCore_QHelpEngineCore_NamespaceName(char* documentationFileName){
	return QHelpEngineCore::namespaceName(QString(documentationFileName)).toUtf8().data();
}

void QHelpEngineCore_ConnectReadersAboutToBeInvalidated(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));;
}

void QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));;
}

int QHelpEngineCore_RegisterDocumentation(void* ptr, char* documentationFileName){
	return static_cast<QHelpEngineCore*>(ptr)->registerDocumentation(QString(documentationFileName));
}

char* QHelpEngineCore_RegisteredDocumentations(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->registeredDocumentations().join("|").toUtf8().data();
}

int QHelpEngineCore_RemoveCustomFilter(void* ptr, char* filterName){
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomFilter(QString(filterName));
}

int QHelpEngineCore_RemoveCustomValue(void* ptr, char* key){
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomValue(QString(key));
}

int QHelpEngineCore_SetCustomValue(void* ptr, char* key, void* value){
	return static_cast<QHelpEngineCore*>(ptr)->setCustomValue(QString(key), *static_cast<QVariant*>(value));
}

int QHelpEngineCore_SetupData(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->setupData();
}

void QHelpEngineCore_ConnectSetupFinished(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));;
}

void QHelpEngineCore_DisconnectSetupFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));;
}

void QHelpEngineCore_ConnectSetupStarted(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));;
}

void QHelpEngineCore_DisconnectSetupStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));;
}

int QHelpEngineCore_UnregisterDocumentation(void* ptr, char* namespaceName){
	return static_cast<QHelpEngineCore*>(ptr)->unregisterDocumentation(QString(namespaceName));
}

void QHelpEngineCore_ConnectWarning(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));;
}

void QHelpEngineCore_DisconnectWarning(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));;
}

void QHelpEngineCore_DestroyQHelpEngineCore(void* ptr){
	static_cast<QHelpEngineCore*>(ptr)->~QHelpEngineCore();
}

#include "qhelpcontentmodel.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QHelpContentModel>
#include "_cgo_export.h"

class MyQHelpContentModel: public QHelpContentModel {
public:
void Signal_ContentsCreated(){callbackQHelpContentModelContentsCreated(this->objectName().toUtf8().data());};
void Signal_ContentsCreationStarted(){callbackQHelpContentModelContentsCreationStarted(this->objectName().toUtf8().data());};
};

int QHelpContentModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QHelpContentModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QHelpContentModel_ContentItemAt(void* ptr, void* index){
	return static_cast<QHelpContentModel*>(ptr)->contentItemAt(*static_cast<QModelIndex*>(index));
}

void QHelpContentModel_ConnectContentsCreated(void* ptr){
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));;
}

void QHelpContentModel_DisconnectContentsCreated(void* ptr){
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));;
}

void QHelpContentModel_ConnectContentsCreationStarted(void* ptr){
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));;
}

void QHelpContentModel_DisconnectContentsCreationStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));;
}

void QHelpContentModel_CreateContents(void* ptr, char* customFilterName){
	static_cast<QHelpContentModel*>(ptr)->createContents(QString(customFilterName));
}

void* QHelpContentModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QHelpContentModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QHelpContentModel_Index(void* ptr, int row, int column, void* parent){
	return static_cast<QHelpContentModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QHelpContentModel_IsCreatingContents(void* ptr){
	return static_cast<QHelpContentModel*>(ptr)->isCreatingContents();
}

void* QHelpContentModel_Parent(void* ptr, void* index){
	return static_cast<QHelpContentModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QHelpContentModel_RowCount(void* ptr, void* parent){
	return static_cast<QHelpContentModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QHelpContentModel_DestroyQHelpContentModel(void* ptr){
	static_cast<QHelpContentModel*>(ptr)->~QHelpContentModel();
}

#include "qhelpcontentitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpContentItem>
#include "_cgo_export.h"

class MyQHelpContentItem: public QHelpContentItem {
public:
};

void* QHelpContentItem_Child(void* ptr, int row){
	return static_cast<QHelpContentItem*>(ptr)->child(row);
}

int QHelpContentItem_ChildCount(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->childCount();
}

int QHelpContentItem_ChildPosition(void* ptr, void* child){
	return static_cast<QHelpContentItem*>(ptr)->childPosition(static_cast<QHelpContentItem*>(child));
}

void* QHelpContentItem_Parent(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->parent();
}

int QHelpContentItem_Row(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->row();
}

char* QHelpContentItem_Title(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->title().toUtf8().data();
}

void QHelpContentItem_DestroyQHelpContentItem(void* ptr){
	static_cast<QHelpContentItem*>(ptr)->~QHelpContentItem();
}

#include "qhelpindexwidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QHelpIndexWidget>
#include "_cgo_export.h"

class MyQHelpIndexWidget: public QHelpIndexWidget {
public:
};

void QHelpIndexWidget_ActivateCurrentItem(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "activateCurrentItem");
}

void QHelpIndexWidget_FilterIndices(void* ptr, char* filter, char* wildcard){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "filterIndices", Q_ARG(QString, QString(filter)), Q_ARG(QString, QString(wildcard)));
}

#include "qhelpsearchquery.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QHelpSearchQuery>
#include "_cgo_export.h"

class MyQHelpSearchQuery: public QHelpSearchQuery {
public:
};

void* QHelpSearchQuery_NewQHelpSearchQuery(){
	return new QHelpSearchQuery();
}

void* QHelpSearchQuery_NewQHelpSearchQuery2(int field, char* wordList){
	return new QHelpSearchQuery(static_cast<QHelpSearchQuery::FieldName>(field), QString(wordList).split("|", QString::SkipEmptyParts));
}

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

void* QHelpEngine_NewQHelpEngine(char* collectionFile, void* parent){
	return new QHelpEngine(QString(collectionFile), static_cast<QObject*>(parent));
}

void* QHelpEngine_ContentModel(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->contentModel();
}

void* QHelpEngine_ContentWidget(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->contentWidget();
}

void* QHelpEngine_IndexModel(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->indexModel();
}

void* QHelpEngine_IndexWidget(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->indexWidget();
}

void* QHelpEngine_SearchEngine(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->searchEngine();
}

void QHelpEngine_DestroyQHelpEngine(void* ptr){
	static_cast<QHelpEngine*>(ptr)->~QHelpEngine();
}

#include "qhelpsearchengine.h"
#include <QHelpEngineCore>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpEngine>
#include <QObject>
#include <QHelpSearchEngine>
#include "_cgo_export.h"

class MyQHelpSearchEngine: public QHelpSearchEngine {
public:
void Signal_IndexingFinished(){callbackQHelpSearchEngineIndexingFinished(this->objectName().toUtf8().data());};
void Signal_IndexingStarted(){callbackQHelpSearchEngineIndexingStarted(this->objectName().toUtf8().data());};
void Signal_SearchingFinished(int hits){callbackQHelpSearchEngineSearchingFinished(this->objectName().toUtf8().data(), hits);};
void Signal_SearchingStarted(){callbackQHelpSearchEngineSearchingStarted(this->objectName().toUtf8().data());};
};

void* QHelpSearchEngine_NewQHelpSearchEngine(void* helpEngine, void* parent){
	return new QHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QObject*>(parent));
}

void QHelpSearchEngine_CancelIndexing(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelIndexing");
}

void QHelpSearchEngine_CancelSearching(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelSearching");
}

int QHelpSearchEngine_HitCount(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->hitCount();
}

void QHelpSearchEngine_ConnectIndexingFinished(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_DisconnectIndexingFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_ConnectIndexingStarted(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

void QHelpSearchEngine_DisconnectIndexingStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

void* QHelpSearchEngine_QueryWidget(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->queryWidget();
}

void QHelpSearchEngine_ReindexDocumentation(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "reindexDocumentation");
}

void* QHelpSearchEngine_ResultWidget(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->resultWidget();
}

void QHelpSearchEngine_ConnectSearchingFinished(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_DisconnectSearchingFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_ConnectSearchingStarted(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_DisconnectSearchingStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_DestroyQHelpSearchEngine(void* ptr){
	static_cast<QHelpSearchEngine*>(ptr)->~QHelpSearchEngine();
}

#include "qhelpsearchresultwidget.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QHelpSearchResultWidget>
#include "_cgo_export.h"

class MyQHelpSearchResultWidget: public QHelpSearchResultWidget {
public:
};

void QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(void* ptr){
	static_cast<QHelpSearchResultWidget*>(ptr)->~QHelpSearchResultWidget();
}

#include "qhelpcontentwidget.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QHelpContentWidget>
#include "_cgo_export.h"

class MyQHelpContentWidget: public QHelpContentWidget {
public:
};

void* QHelpContentWidget_IndexOf(void* ptr, void* link){
	return static_cast<QHelpContentWidget*>(ptr)->indexOf(*static_cast<QUrl*>(link)).internalPointer();
}

#include "qhelpindexmodel.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QHelpIndexModel>
#include "_cgo_export.h"

class MyQHelpIndexModel: public QHelpIndexModel {
public:
void Signal_IndexCreated(){callbackQHelpIndexModelIndexCreated(this->objectName().toUtf8().data());};
void Signal_IndexCreationStarted(){callbackQHelpIndexModelIndexCreationStarted(this->objectName().toUtf8().data());};
};

void QHelpIndexModel_CreateIndex(void* ptr, char* customFilterName){
	static_cast<QHelpIndexModel*>(ptr)->createIndex(QString(customFilterName));
}

void* QHelpIndexModel_Filter(void* ptr, char* filter, char* wildcard){
	return static_cast<QHelpIndexModel*>(ptr)->filter(QString(filter), QString(wildcard)).internalPointer();
}

void QHelpIndexModel_ConnectIndexCreated(void* ptr){
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));;
}

void QHelpIndexModel_DisconnectIndexCreated(void* ptr){
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));;
}

void QHelpIndexModel_ConnectIndexCreationStarted(void* ptr){
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));;
}

void QHelpIndexModel_DisconnectIndexCreationStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));;
}

int QHelpIndexModel_IsCreatingIndex(void* ptr){
	return static_cast<QHelpIndexModel*>(ptr)->isCreatingIndex();
}

