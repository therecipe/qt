#include "qstringlistmodel.h"
#include <QList>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QList>
#include <QStringList>
#include <QStringListModel>
#include "_cgo_export.h"

class MyQStringListModel: public QStringListModel {
public:
};

void* QStringListModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QStringListModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QStringListModel_Flags(void* ptr, void* index){
	return static_cast<QStringListModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QStringListModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QStringListModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStringListModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QStringListModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStringListModel_RowCount(void* ptr, void* parent){
	return static_cast<QStringListModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QStringListModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QStringListModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QStringListModel_SetStringList(void* ptr, char* strin){
	static_cast<QStringListModel*>(ptr)->setStringList(QString(strin).split("|", QString::SkipEmptyParts));
}

void* QStringListModel_Sibling(void* ptr, int row, int column, void* idx){
	return static_cast<QStringListModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QStringListModel_Sort(void* ptr, int column, int order){
	static_cast<QStringListModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

char* QStringListModel_StringList(void* ptr){
	return static_cast<QStringListModel*>(ptr)->stringList().join("|").toUtf8().data();
}

int QStringListModel_SupportedDropActions(void* ptr){
	return static_cast<QStringListModel*>(ptr)->supportedDropActions();
}

