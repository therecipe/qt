#include "qtreeview.h"
#include <QHeaderView>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QAbstractItemModel>
#include <QMetaObject>
#include <QUrl>
#include <QPoint>
#include <QModelIndex>
#include <QObject>
#include <QAbstractItemView>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QTreeView>
#include "_cgo_export.h"

class MyQTreeView: public QTreeView {
public:
void Signal_Collapsed(const QModelIndex & index){callbackQTreeViewCollapsed(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_Expanded(const QModelIndex & index){callbackQTreeViewExpanded(this->objectName().toUtf8().data(), index.internalPointer());};
};

int QTreeView_AllColumnsShowFocus(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->allColumnsShowFocus();
}

int QTreeView_AutoExpandDelay(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->autoExpandDelay();
}

void QTreeView_Collapse(QtObjectPtr ptr, QtObjectPtr index){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "collapse", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QTreeView_Expand(QtObjectPtr ptr, QtObjectPtr index){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expand", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

int QTreeView_ExpandsOnDoubleClick(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->expandsOnDoubleClick();
}

int QTreeView_Indentation(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->indentation();
}

int QTreeView_IsAnimated(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->isAnimated();
}

int QTreeView_IsExpanded(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QTreeView*>(ptr)->isExpanded(*static_cast<QModelIndex*>(index));
}

int QTreeView_IsHeaderHidden(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->isHeaderHidden();
}

int QTreeView_IsSortingEnabled(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->isSortingEnabled();
}

int QTreeView_ItemsExpandable(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->itemsExpandable();
}

void QTreeView_ResetIndentation(QtObjectPtr ptr){
	static_cast<QTreeView*>(ptr)->resetIndentation();
}

int QTreeView_RootIsDecorated(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->rootIsDecorated();
}

void QTreeView_SetAllColumnsShowFocus(QtObjectPtr ptr, int enable){
	static_cast<QTreeView*>(ptr)->setAllColumnsShowFocus(enable != 0);
}

void QTreeView_SetAnimated(QtObjectPtr ptr, int enable){
	static_cast<QTreeView*>(ptr)->setAnimated(enable != 0);
}

void QTreeView_SetAutoExpandDelay(QtObjectPtr ptr, int delay){
	static_cast<QTreeView*>(ptr)->setAutoExpandDelay(delay);
}

void QTreeView_SetExpandsOnDoubleClick(QtObjectPtr ptr, int enable){
	static_cast<QTreeView*>(ptr)->setExpandsOnDoubleClick(enable != 0);
}

void QTreeView_SetHeaderHidden(QtObjectPtr ptr, int hide){
	static_cast<QTreeView*>(ptr)->setHeaderHidden(hide != 0);
}

void QTreeView_SetIndentation(QtObjectPtr ptr, int i){
	static_cast<QTreeView*>(ptr)->setIndentation(i);
}

void QTreeView_SetItemsExpandable(QtObjectPtr ptr, int enable){
	static_cast<QTreeView*>(ptr)->setItemsExpandable(enable != 0);
}

void QTreeView_SetRootIsDecorated(QtObjectPtr ptr, int show){
	static_cast<QTreeView*>(ptr)->setRootIsDecorated(show != 0);
}

void QTreeView_SetSortingEnabled(QtObjectPtr ptr, int enable){
	static_cast<QTreeView*>(ptr)->setSortingEnabled(enable != 0);
}

void QTreeView_SetUniformRowHeights(QtObjectPtr ptr, int uniform){
	static_cast<QTreeView*>(ptr)->setUniformRowHeights(uniform != 0);
}

void QTreeView_SetWordWrap(QtObjectPtr ptr, int on){
	static_cast<QTreeView*>(ptr)->setWordWrap(on != 0);
}

int QTreeView_UniformRowHeights(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->uniformRowHeights();
}

int QTreeView_WordWrap(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->wordWrap();
}

QtObjectPtr QTreeView_NewQTreeView(QtObjectPtr parent){
	return new QTreeView(static_cast<QWidget*>(parent));
}

void QTreeView_CollapseAll(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "collapseAll");
}

void QTreeView_ConnectCollapsed(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::collapsed), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Collapsed));;
}

void QTreeView_DisconnectCollapsed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::collapsed), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Collapsed));;
}

int QTreeView_ColumnAt(QtObjectPtr ptr, int x){
	return static_cast<QTreeView*>(ptr)->columnAt(x);
}

int QTreeView_ColumnViewportPosition(QtObjectPtr ptr, int column){
	return static_cast<QTreeView*>(ptr)->columnViewportPosition(column);
}

int QTreeView_ColumnWidth(QtObjectPtr ptr, int column){
	return static_cast<QTreeView*>(ptr)->columnWidth(column);
}

void QTreeView_ExpandAll(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expandAll");
}

void QTreeView_ExpandToDepth(QtObjectPtr ptr, int depth){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expandToDepth", Q_ARG(int, depth));
}

void QTreeView_ConnectExpanded(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::expanded), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Expanded));;
}

void QTreeView_DisconnectExpanded(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::expanded), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Expanded));;
}

QtObjectPtr QTreeView_Header(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->header();
}

void QTreeView_HideColumn(QtObjectPtr ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "hideColumn", Q_ARG(int, column));
}

QtObjectPtr QTreeView_IndexAbove(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QTreeView*>(ptr)->indexAbove(*static_cast<QModelIndex*>(index)).internalPointer();
}

QtObjectPtr QTreeView_IndexAt(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QTreeView*>(ptr)->indexAt(*static_cast<QPoint*>(point)).internalPointer();
}

QtObjectPtr QTreeView_IndexBelow(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QTreeView*>(ptr)->indexBelow(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QTreeView_IsColumnHidden(QtObjectPtr ptr, int column){
	return static_cast<QTreeView*>(ptr)->isColumnHidden(column);
}

int QTreeView_IsFirstColumnSpanned(QtObjectPtr ptr, int row, QtObjectPtr parent){
	return static_cast<QTreeView*>(ptr)->isFirstColumnSpanned(row, *static_cast<QModelIndex*>(parent));
}

int QTreeView_IsRowHidden(QtObjectPtr ptr, int row, QtObjectPtr parent){
	return static_cast<QTreeView*>(ptr)->isRowHidden(row, *static_cast<QModelIndex*>(parent));
}

void QTreeView_KeyboardSearch(QtObjectPtr ptr, char* search){
	static_cast<QTreeView*>(ptr)->keyboardSearch(QString(search));
}

void QTreeView_Reset(QtObjectPtr ptr){
	static_cast<QTreeView*>(ptr)->reset();
}

void QTreeView_ResizeColumnToContents(QtObjectPtr ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "resizeColumnToContents", Q_ARG(int, column));
}

void QTreeView_ScrollTo(QtObjectPtr ptr, QtObjectPtr index, int hint){
	static_cast<QTreeView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QTreeView_SelectAll(QtObjectPtr ptr){
	static_cast<QTreeView*>(ptr)->selectAll();
}

void QTreeView_SetColumnHidden(QtObjectPtr ptr, int column, int hide){
	static_cast<QTreeView*>(ptr)->setColumnHidden(column, hide != 0);
}

void QTreeView_SetColumnWidth(QtObjectPtr ptr, int column, int width){
	static_cast<QTreeView*>(ptr)->setColumnWidth(column, width);
}

void QTreeView_SetExpanded(QtObjectPtr ptr, QtObjectPtr index, int expanded){
	static_cast<QTreeView*>(ptr)->setExpanded(*static_cast<QModelIndex*>(index), expanded != 0);
}

void QTreeView_SetFirstColumnSpanned(QtObjectPtr ptr, int row, QtObjectPtr parent, int span){
	static_cast<QTreeView*>(ptr)->setFirstColumnSpanned(row, *static_cast<QModelIndex*>(parent), span != 0);
}

void QTreeView_SetHeader(QtObjectPtr ptr, QtObjectPtr header){
	static_cast<QTreeView*>(ptr)->setHeader(static_cast<QHeaderView*>(header));
}

void QTreeView_SetModel(QtObjectPtr ptr, QtObjectPtr model){
	static_cast<QTreeView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QTreeView_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index){
	static_cast<QTreeView*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QTreeView_SetRowHidden(QtObjectPtr ptr, int row, QtObjectPtr parent, int hide){
	static_cast<QTreeView*>(ptr)->setRowHidden(row, *static_cast<QModelIndex*>(parent), hide != 0);
}

void QTreeView_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr selectionModel){
	static_cast<QTreeView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QTreeView_SetTreePosition(QtObjectPtr ptr, int index){
	static_cast<QTreeView*>(ptr)->setTreePosition(index);
}

void QTreeView_ShowColumn(QtObjectPtr ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "showColumn", Q_ARG(int, column));
}

void QTreeView_SortByColumn(QtObjectPtr ptr, int column, int order){
	static_cast<QTreeView*>(ptr)->sortByColumn(column, static_cast<Qt::SortOrder>(order));
}

int QTreeView_TreePosition(QtObjectPtr ptr){
	return static_cast<QTreeView*>(ptr)->treePosition();
}

void QTreeView_DestroyQTreeView(QtObjectPtr ptr){
	static_cast<QTreeView*>(ptr)->~QTreeView();
}

