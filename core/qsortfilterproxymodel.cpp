#include "qsortfilterproxymodel.h"
#include <QAbstractItemModel>
#include <QMetaObject>
#include <QObject>
#include <QRegExp>
#include <QMimeData>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSortFilterProxyModel>
#include "_cgo_export.h"

class MyQSortFilterProxyModel: public QSortFilterProxyModel {
public:
};

int QSortFilterProxyModel_DynamicSortFilter(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->dynamicSortFilter();
}

int QSortFilterProxyModel_FilterCaseSensitivity(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterCaseSensitivity();
}

int QSortFilterProxyModel_FilterKeyColumn(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterKeyColumn();
}

void* QSortFilterProxyModel_FilterRegExp(void* ptr){
	return new QRegExp(static_cast<QSortFilterProxyModel*>(ptr)->filterRegExp());
}

int QSortFilterProxyModel_FilterRole(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->filterRole();
}

int QSortFilterProxyModel_IsSortLocaleAware(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->isSortLocaleAware();
}

void QSortFilterProxyModel_SetDynamicSortFilter(void* ptr, int enable){
	static_cast<QSortFilterProxyModel*>(ptr)->setDynamicSortFilter(enable != 0);
}

void QSortFilterProxyModel_SetFilterCaseSensitivity(void* ptr, int cs){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QSortFilterProxyModel_SetFilterKeyColumn(void* ptr, int column){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterKeyColumn(column);
}

void QSortFilterProxyModel_SetFilterRegExp(void* ptr, void* regExp){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterRegExp(*static_cast<QRegExp*>(regExp));
}

void QSortFilterProxyModel_SetFilterRole(void* ptr, int role){
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterRole(role);
}

void QSortFilterProxyModel_SetSortCaseSensitivity(void* ptr, int cs){
	static_cast<QSortFilterProxyModel*>(ptr)->setSortCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QSortFilterProxyModel_SetSortLocaleAware(void* ptr, int on){
	static_cast<QSortFilterProxyModel*>(ptr)->setSortLocaleAware(on != 0);
}

void QSortFilterProxyModel_SetSortRole(void* ptr, int role){
	static_cast<QSortFilterProxyModel*>(ptr)->setSortRole(role);
}

int QSortFilterProxyModel_SortCaseSensitivity(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortCaseSensitivity();
}

int QSortFilterProxyModel_SortRole(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortRole();
}

void* QSortFilterProxyModel_NewQSortFilterProxyModel(void* parent){
	return new QSortFilterProxyModel(static_cast<QObject*>(parent));
}

void* QSortFilterProxyModel_Buddy(void* ptr, void* index){
	return static_cast<QSortFilterProxyModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QSortFilterProxyModel_CanFetchMore(void* ptr, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QSortFilterProxyModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QSortFilterProxyModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QSortFilterProxyModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QSortFilterProxyModel_FetchMore(void* ptr, void* parent){
	static_cast<QSortFilterProxyModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_Flags(void* ptr, void* index){
	return static_cast<QSortFilterProxyModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QSortFilterProxyModel_HasChildren(void* ptr, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QSortFilterProxyModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QSortFilterProxyModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QSortFilterProxyModel_Index(void* ptr, int row, int column, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QSortFilterProxyModel_InsertColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QSortFilterProxyModel_Invalidate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "invalidate");
}

void* QSortFilterProxyModel_MapFromSource(void* ptr, void* sourceIndex){
	return static_cast<QSortFilterProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)).internalPointer();
}

void* QSortFilterProxyModel_MapToSource(void* ptr, void* proxyIndex){
	return static_cast<QSortFilterProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)).internalPointer();
}

char* QSortFilterProxyModel_MimeTypes(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

void* QSortFilterProxyModel_Parent(void* ptr, void* child){
	return static_cast<QSortFilterProxyModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)).internalPointer();
}

int QSortFilterProxyModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_RowCount(void* ptr, void* parent){
	return static_cast<QSortFilterProxyModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QSortFilterProxyModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QSortFilterProxyModel_SetFilterFixedString(void* ptr, char* pattern){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterFixedString", Q_ARG(QString, QString(pattern)));
}

void QSortFilterProxyModel_SetFilterRegExp2(void* ptr, char* pattern){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterRegExp", Q_ARG(QString, QString(pattern)));
}

void QSortFilterProxyModel_SetFilterWildcard(void* ptr, char* pattern){
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterWildcard", Q_ARG(QString, QString(pattern)));
}

int QSortFilterProxyModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role){
	return static_cast<QSortFilterProxyModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QSortFilterProxyModel_SetSourceModel(void* ptr, void* sourceModel){
	static_cast<QSortFilterProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
}

void* QSortFilterProxyModel_Sibling(void* ptr, int row, int column, void* idx){
	return static_cast<QSortFilterProxyModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QSortFilterProxyModel_Sort(void* ptr, int column, int order){
	static_cast<QSortFilterProxyModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QSortFilterProxyModel_SortColumn(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortColumn();
}

int QSortFilterProxyModel_SortOrder(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->sortOrder();
}

int QSortFilterProxyModel_SupportedDropActions(void* ptr){
	return static_cast<QSortFilterProxyModel*>(ptr)->supportedDropActions();
}

void QSortFilterProxyModel_DestroyQSortFilterProxyModel(void* ptr){
	static_cast<QSortFilterProxyModel*>(ptr)->~QSortFilterProxyModel();
}

