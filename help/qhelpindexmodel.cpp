#include "qhelpindexmodel.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QHelpIndexModel>
#include "_cgo_export.h"

class MyQHelpIndexModel: public QHelpIndexModel {
public:
void Signal_IndexCreated(){callbackQHelpIndexModelIndexCreated(this->objectName().toUtf8().data());};
void Signal_IndexCreationStarted(){callbackQHelpIndexModelIndexCreationStarted(this->objectName().toUtf8().data());};
};

void QHelpIndexModel_CreateIndex(QtObjectPtr ptr, char* customFilterName){
	static_cast<QHelpIndexModel*>(ptr)->createIndex(QString(customFilterName));
}

QtObjectPtr QHelpIndexModel_Filter(QtObjectPtr ptr, char* filter, char* wildcard){
	return static_cast<QHelpIndexModel*>(ptr)->filter(QString(filter), QString(wildcard)).internalPointer();
}

void QHelpIndexModel_ConnectIndexCreated(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));;
}

void QHelpIndexModel_DisconnectIndexCreated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));;
}

void QHelpIndexModel_ConnectIndexCreationStarted(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));;
}

void QHelpIndexModel_DisconnectIndexCreationStarted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));;
}

int QHelpIndexModel_IsCreatingIndex(QtObjectPtr ptr){
	return static_cast<QHelpIndexModel*>(ptr)->isCreatingIndex();
}

