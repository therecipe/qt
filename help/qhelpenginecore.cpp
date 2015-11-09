#include "qhelpenginecore.h"
#include <QByteArray>
#include <QHelpEngine>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
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

