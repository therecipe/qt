#include "qhelpcontentmodel.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
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

