#include "qstringlistmodel.h"
#include <QList>
#include <QList>
#include <QStringList>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStringListModel>
#include "_cgo_export.h"

class MyQStringListModel: public QStringListModel {
public:
};

char* QStringListModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role){
	return static_cast<QStringListModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role).toString().toUtf8().data();
}

int QStringListModel_Flags(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QStringListModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QStringListModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QStringListModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStringListModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QStringListModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStringListModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QStringListModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QStringListModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role){
	return static_cast<QStringListModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), QVariant(value), role);
}

void QStringListModel_SetStringList(QtObjectPtr ptr, char* strin){
	static_cast<QStringListModel*>(ptr)->setStringList(QString(strin).split("|", QString::SkipEmptyParts));
}

QtObjectPtr QStringListModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx){
	return static_cast<QStringListModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QStringListModel_Sort(QtObjectPtr ptr, int column, int order){
	static_cast<QStringListModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

char* QStringListModel_StringList(QtObjectPtr ptr){
	return static_cast<QStringListModel*>(ptr)->stringList().join("|").toUtf8().data();
}

int QStringListModel_SupportedDropActions(QtObjectPtr ptr){
	return static_cast<QStringListModel*>(ptr)->supportedDropActions();
}

