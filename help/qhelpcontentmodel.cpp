#include "qhelpcontentmodel.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpContentModel>
#include "_cgo_export.h"

class MyQHelpContentModel: public QHelpContentModel {
public:
void Signal_ContentsCreated(){callbackQHelpContentModelContentsCreated(this->objectName().toUtf8().data());};
void Signal_ContentsCreationStarted(){callbackQHelpContentModelContentsCreationStarted(this->objectName().toUtf8().data());};
};

int QHelpContentModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QHelpContentModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

QtObjectPtr QHelpContentModel_ContentItemAt(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QHelpContentModel*>(ptr)->contentItemAt(*static_cast<QModelIndex*>(index));
}

void QHelpContentModel_ConnectContentsCreated(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));;
}

void QHelpContentModel_DisconnectContentsCreated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));;
}

void QHelpContentModel_ConnectContentsCreationStarted(QtObjectPtr ptr){
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));;
}

void QHelpContentModel_DisconnectContentsCreationStarted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));;
}

void QHelpContentModel_CreateContents(QtObjectPtr ptr, char* customFilterName){
	static_cast<QHelpContentModel*>(ptr)->createContents(QString(customFilterName));
}

char* QHelpContentModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role){
	return static_cast<QHelpContentModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role).toString().toUtf8().data();
}

QtObjectPtr QHelpContentModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent){
	return static_cast<QHelpContentModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QHelpContentModel_IsCreatingContents(QtObjectPtr ptr){
	return static_cast<QHelpContentModel*>(ptr)->isCreatingContents();
}

QtObjectPtr QHelpContentModel_Parent(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QHelpContentModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QHelpContentModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QHelpContentModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QHelpContentModel_DestroyQHelpContentModel(QtObjectPtr ptr){
	static_cast<QHelpContentModel*>(ptr)->~QHelpContentModel();
}

