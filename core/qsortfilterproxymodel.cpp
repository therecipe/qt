#include "qsortfilterproxymodel.h"
#include <QUrl>
#include <QAbstractItemModel>
#include <QMetaObject>
#include <QObject>
#include <QRegExp>
#include <QMimeData>
#include <QVariant>
#include <QModelIndex>
#include <QString>
#include <QSortFilterProxyModel>
#include "_cgo_export.h"

class MyQSortFilterProxyModel: public QSortFilterProxyModel {
public:
};

int QSortFilterProxyModel_DynamicSortFilter(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->dynamicSortFilter();
}

int QSortFilterProxyModel_FilterCaseSensitivity(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterCaseSensitivity();
}

int QSortFilterProxyModel_FilterKeyColumn(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterKeyColumn();
}

int QSortFilterProxyModel_FilterRole(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterRole();
}

int QSortFilterProxyModel_IsSortLocaleAware(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->isSortLocaleAware();
}

void QSortFilterProxyModel_SetDynamicSortFilter(QtObjectPtr ptr, int enable){
	static_cast<QSortFilterProxyModel*>(ptr)->setDynamicSortFilter(enable != 0);
}

void QSortFilterProxyModel_SetFilterCaseSensitivity(QtObjectPtr ptr, int cs){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QSortFilterProxyModel_SetFilterKeyColumn(QtObjectPtr ptr, int column){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterKeyColumn(column);
}

void QSortFilterProxyModel_SetFilterRegExp(QtObjectPtr ptr, QtObjectPtr regExp){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterRegExp(*static_cast<QRegExp*>(regExp));
}

void QSortFilterProxyModel_SetFilterRole(QtObjectPtr ptr, int role){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterRole(role);
}

void QSortFilterProxyModel_SetSortCaseSensitivity(QtObjectPtr ptr, int cs){
	static_cast<QSortFilterProxyModel*>(ptr)->setSortCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QSortFilterProxyModel_SetSortLocaleAware(QtObjectPtr ptr, int on){
	static_cast<QSortFilterProxyModel*>(ptr)->setSortLocaleAware(on != 0);
}

void QSortFilterProxyModel_SetSortRole(QtObjectPtr ptr, int role){
	static_cast<QSortFilterProxyModel*>(ptr)->setSortRole(role);
}

int QSortFilterProxyModel_SortCaseSensitivity(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortCaseSensitivity();
}

int QSortFilterProxyModel_SortRole(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortRole();
}

QtObjectPtr QSortFilterProxyModel_NewQSortFilterProxyModel(QtObjectPtr parent){
	return new QSortFilterProxyModel(static_cast<QObject*>(parent));
}

QtObjectPtr QSortFilterProxyModel_Buddy(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QSortFilterProxyModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QSortFilterProxyModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

char* QSortFilterProxyModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role){
	return static_cast<QSortFilterProxyModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role).toString().toUtf8().data();
}

int QSortFilterProxyModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QSortFilterProxyModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QSortFilterProxyModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_Flags(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QSortFilterProxyModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QSortFilterProxyModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

char* QSortFilterProxyModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role){
	return static_cast<QSortFilterProxyModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role).toString().toUtf8().data();
}

QtObjectPtr QSortFilterProxyModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QSortFilterProxyModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSortFilterProxyModel_Invalidate(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "invalidate");
}

QtObjectPtr QSortFilterProxyModel_MapFromSource(QtObjectPtr ptr, QtObjectPtr sourceIndex){
	return static_cast<QSortFilterProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)).internalPointer();
}

QtObjectPtr QSortFilterProxyModel_MapToSource(QtObjectPtr ptr, QtObjectPtr proxyIndex){
	return static_cast<QSortFilterProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)).internalPointer();
}

char* QSortFilterProxyModel_MimeTypes(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

QtObjectPtr QSortFilterProxyModel_Parent(QtObjectPtr ptr, QtObjectPtr child){
	return static_cast<QSortFilterProxyModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)).internalPointer();
}

int QSortFilterProxyModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role){
	return static_cast<QSortFilterProxyModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), QVariant(value), role);
}

void QSortFilterProxyModel_SetFilterFixedString(QtObjectPtr ptr, char* pattern){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterFixedString", Q_ARG(QString, QString(pattern)));
}

void QSortFilterProxyModel_SetFilterRegExp2(QtObjectPtr ptr, char* pattern){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterRegExp", Q_ARG(QString, QString(pattern)));
}

void QSortFilterProxyModel_SetFilterWildcard(QtObjectPtr ptr, char* pattern){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterWildcard", Q_ARG(QString, QString(pattern)));
}

int QSortFilterProxyModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role){
	return static_cast<QSortFilterProxyModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), QVariant(value), role);
}

void QSortFilterProxyModel_SetSourceModel(QtObjectPtr ptr, QtObjectPtr sourceModel){
	static_cast<QSortFilterProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
}

QtObjectPtr QSortFilterProxyModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx){
	return static_cast<QSortFilterProxyModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QSortFilterProxyModel_Sort(QtObjectPtr ptr, int column, int order){
	static_cast<QSortFilterProxyModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QSortFilterProxyModel_SortColumn(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortColumn();
}

int QSortFilterProxyModel_SortOrder(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortOrder();
}

int QSortFilterProxyModel_SupportedDropActions(QtObjectPtr ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->supportedDropActions();
}

void QSortFilterProxyModel_DestroyQSortFilterProxyModel(QtObjectPtr ptr){
	static_cast<QSortFilterProxyModel*>(ptr)->~QSortFilterProxyModel();
}

