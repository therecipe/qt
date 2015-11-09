#include "qhelpindexmodel.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
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

