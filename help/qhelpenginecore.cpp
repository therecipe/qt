#include "qhelpenginecore.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QHelpEngine>
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

int QHelpEngineCore_AutoSaveFilter(QtObjectPtr ptr){
	return static_cast<QHelpEngineCore*>(ptr)->autoSaveFilter();
}

char* QHelpEngineCore_CollectionFile(QtObjectPtr ptr){
	return static_cast<QHelpEngineCore*>(ptr)->collectionFile().toUtf8().data();
}

char* QHelpEngineCore_CurrentFilter(QtObjectPtr ptr){
	return static_cast<QHelpEngineCore*>(ptr)->currentFilter().toUtf8().data();
}

void QHelpEngineCore_SetAutoSaveFilter(QtObjectPtr ptr, int save){
	static_cast<QHelpEngineCore*>(ptr)->setAutoSaveFilter(save != 0);
}

void QHelpEngineCore_SetCollectionFile(QtObjectPtr ptr, char* fileName){
	static_cast<QHelpEngineCore*>(ptr)->setCollectionFile(QString(fileName));
}

void QHelpEngineCore_SetCurrentFilter(QtObjectPtr ptr, char* filterName){
	static_cast<QHelpEngineCore*>(ptr)->setCurrentFilter(QString(filterName));
}

QtObjectPtr QHelpEngineCore_NewQHelpEngineCore(char* collectionFile, QtObjectPtr parent){
	return new QHelpEngineCore(QString(collectionFile), static_cast<QObject*>(parent));
}

int QHelpEngineCore_AddCustomFilter(QtObjectPtr ptr, char* filterName, char* attributes){
	return static_cast<QHelpEngineCore*>(ptr)->addCustomFilter(QString(filterName), QString(attributes).split("|", QString::SkipEmptyParts));
}

int QHelpEngineCore_CopyCollectionFile(QtObjectPtr ptr, char* fileName){
	return static_cast<QHelpEngineCore*>(ptr)->copyCollectionFile(QString(fileName));
}

void QHelpEngineCore_ConnectCurrentFilterChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));;
}

void QHelpEngineCore_DisconnectCurrentFilterChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));;
}

char* QHelpEngineCore_CustomFilters(QtObjectPtr ptr){
	return static_cast<QHelpEngineCore*>(ptr)->customFilters().join("|").toUtf8().data();
}

char* QHelpEngineCore_CustomValue(QtObjectPtr ptr, char* key, char* defaultValue){
	return static_cast<QHelpEngineCore*>(ptr)->customValue(QString(key), QVariant(defaultValue)).toString().toUtf8().data();
}

char* QHelpEngineCore_DocumentationFileName(QtObjectPtr ptr, char* namespaceName){
	return static_cast<QHelpEngineCore*>(ptr)->documentationFileName(QString(namespaceName)).toUtf8().data();
}

char* QHelpEngineCore_Error(QtObjectPtr ptr){
	return static_cast<QHelpEngineCore*>(ptr)->error().toUtf8().data();
}

char* QHelpEngineCore_FilterAttributes(QtObjectPtr ptr){
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes().join("|").toUtf8().data();
}

char* QHelpEngineCore_FilterAttributes2(QtObjectPtr ptr, char* filterName){
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes(QString(filterName)).join("|").toUtf8().data();
}

char* QHelpEngineCore_FindFile(QtObjectPtr ptr, char* url){
	return static_cast<QHelpEngineCore*>(ptr)->findFile(QUrl(QString(url))).toString().toUtf8().data();
}

char* QHelpEngineCore_QHelpEngineCore_MetaData(char* documentationFileName, char* name){
	return QHelpEngineCore::metaData(QString(documentationFileName), QString(name)).toString().toUtf8().data();
}

char* QHelpEngineCore_QHelpEngineCore_NamespaceName(char* documentationFileName){
	return QHelpEngineCore::namespaceName(QString(documentationFileName)).toUtf8().data();
}

void QHelpEngineCore_ConnectReadersAboutToBeInvalidated(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));;
}

void QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));;
}

int QHelpEngineCore_RegisterDocumentation(QtObjectPtr ptr, char* documentationFileName){
	return static_cast<QHelpEngineCore*>(ptr)->registerDocumentation(QString(documentationFileName));
}

char* QHelpEngineCore_RegisteredDocumentations(QtObjectPtr ptr){
	return static_cast<QHelpEngineCore*>(ptr)->registeredDocumentations().join("|").toUtf8().data();
}

int QHelpEngineCore_RemoveCustomFilter(QtObjectPtr ptr, char* filterName){
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomFilter(QString(filterName));
}

int QHelpEngineCore_RemoveCustomValue(QtObjectPtr ptr, char* key){
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomValue(QString(key));
}

int QHelpEngineCore_SetCustomValue(QtObjectPtr ptr, char* key, char* value){
	return static_cast<QHelpEngineCore*>(ptr)->setCustomValue(QString(key), QVariant(value));
}

int QHelpEngineCore_SetupData(QtObjectPtr ptr){
	return static_cast<QHelpEngineCore*>(ptr)->setupData();
}

void QHelpEngineCore_ConnectSetupFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));;
}

void QHelpEngineCore_DisconnectSetupFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));;
}

void QHelpEngineCore_ConnectSetupStarted(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));;
}

void QHelpEngineCore_DisconnectSetupStarted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));;
}

int QHelpEngineCore_UnregisterDocumentation(QtObjectPtr ptr, char* namespaceName){
	return static_cast<QHelpEngineCore*>(ptr)->unregisterDocumentation(QString(namespaceName));
}

void QHelpEngineCore_ConnectWarning(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));;
}

void QHelpEngineCore_DisconnectWarning(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));;
}

void QHelpEngineCore_DestroyQHelpEngineCore(QtObjectPtr ptr){
	static_cast<QHelpEngineCore*>(ptr)->~QHelpEngineCore();
}

