#include "qtreeview.h"
#include <QString>
#include <QItemSelectionModel>
#include <QPoint>
#include <QItemSelection>
#include <QAbstractItemView>
#include <QMetaObject>
#include <QHeaderView>
#include <QWidget>
#include <QAbstractItemModel>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QTreeView>
#include "_cgo_export.h"

class MyQTreeView: public QTreeView {
public:
void Signal_Collapsed(const QModelIndex & index){callbackQTreeViewCollapsed(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_Expanded(const QModelIndex & index){callbackQTreeViewExpanded(this->objectName().toUtf8().data(), index.internalPointer());};
};

int QTreeView_AllColumnsShowFocus(void* ptr){
	return static_cast<QTreeView*>(ptr)->allColumnsShowFocus();
}

int QTreeView_AutoExpandDelay(void* ptr){
	return static_cast<QTreeView*>(ptr)->autoExpandDelay();
}

void QTreeView_Collapse(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "collapse", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QTreeView_Expand(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expand", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

int QTreeView_ExpandsOnDoubleClick(void* ptr){
	return static_cast<QTreeView*>(ptr)->expandsOnDoubleClick();
}

int QTreeView_Indentation(void* ptr){
	return static_cast<QTreeView*>(ptr)->indentation();
}

int QTreeView_IsAnimated(void* ptr){
	return static_cast<QTreeView*>(ptr)->isAnimated();
}

int QTreeView_IsExpanded(void* ptr, void* index){
	return static_cast<QTreeView*>(ptr)->isExpanded(*static_cast<QModelIndex*>(index));
}

int QTreeView_IsHeaderHidden(void* ptr){
	return static_cast<QTreeView*>(ptr)->isHeaderHidden();
}

int QTreeView_IsSortingEnabled(void* ptr){
	return static_cast<QTreeView*>(ptr)->isSortingEnabled();
}

int QTreeView_ItemsExpandable(void* ptr){
	return static_cast<QTreeView*>(ptr)->itemsExpandable();
}

void QTreeView_ResetIndentation(void* ptr){
	static_cast<QTreeView*>(ptr)->resetIndentation();
}

int QTreeView_RootIsDecorated(void* ptr){
	return static_cast<QTreeView*>(ptr)->rootIsDecorated();
}

void QTreeView_SetAllColumnsShowFocus(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setAllColumnsShowFocus(enable != 0);
}

void QTreeView_SetAnimated(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setAnimated(enable != 0);
}

void QTreeView_SetAutoExpandDelay(void* ptr, int delay){
	static_cast<QTreeView*>(ptr)->setAutoExpandDelay(delay);
}

void QTreeView_SetExpandsOnDoubleClick(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setExpandsOnDoubleClick(enable != 0);
}

void QTreeView_SetHeaderHidden(void* ptr, int hide){
	static_cast<QTreeView*>(ptr)->setHeaderHidden(hide != 0);
}

void QTreeView_SetIndentation(void* ptr, int i){
	static_cast<QTreeView*>(ptr)->setIndentation(i);
}

void QTreeView_SetItemsExpandable(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setItemsExpandable(enable != 0);
}

void QTreeView_SetRootIsDecorated(void* ptr, int show){
	static_cast<QTreeView*>(ptr)->setRootIsDecorated(show != 0);
}

void QTreeView_SetSortingEnabled(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setSortingEnabled(enable != 0);
}

void QTreeView_SetUniformRowHeights(void* ptr, int uniform){
	static_cast<QTreeView*>(ptr)->setUniformRowHeights(uniform != 0);
}

void QTreeView_SetWordWrap(void* ptr, int on){
	static_cast<QTreeView*>(ptr)->setWordWrap(on != 0);
}

int QTreeView_UniformRowHeights(void* ptr){
	return static_cast<QTreeView*>(ptr)->uniformRowHeights();
}

int QTreeView_WordWrap(void* ptr){
	return static_cast<QTreeView*>(ptr)->wordWrap();
}

void* QTreeView_NewQTreeView(void* parent){
	return new QTreeView(static_cast<QWidget*>(parent));
}

void QTreeView_CollapseAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "collapseAll");
}

void QTreeView_ConnectCollapsed(void* ptr){
	QObject::connect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::collapsed), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Collapsed));;
}

void QTreeView_DisconnectCollapsed(void* ptr){
	QObject::disconnect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::collapsed), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Collapsed));;
}

int QTreeView_ColumnAt(void* ptr, int x){
	return static_cast<QTreeView*>(ptr)->columnAt(x);
}

int QTreeView_ColumnViewportPosition(void* ptr, int column){
	return static_cast<QTreeView*>(ptr)->columnViewportPosition(column);
}

int QTreeView_ColumnWidth(void* ptr, int column){
	return static_cast<QTreeView*>(ptr)->columnWidth(column);
}

void QTreeView_ExpandAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expandAll");
}

void QTreeView_ExpandToDepth(void* ptr, int depth){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expandToDepth", Q_ARG(int, depth));
}

void QTreeView_ConnectExpanded(void* ptr){
	QObject::connect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::expanded), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Expanded));;
}

void QTreeView_DisconnectExpanded(void* ptr){
	QObject::disconnect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::expanded), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Expanded));;
}

void* QTreeView_Header(void* ptr){
	return static_cast<QTreeView*>(ptr)->header();
}

void QTreeView_HideColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "hideColumn", Q_ARG(int, column));
}

void* QTreeView_IndexAbove(void* ptr, void* index){
	return static_cast<QTreeView*>(ptr)->indexAbove(*static_cast<QModelIndex*>(index)).internalPointer();
}

void* QTreeView_IndexAt(void* ptr, void* point){
	return static_cast<QTreeView*>(ptr)->indexAt(*static_cast<QPoint*>(point)).internalPointer();
}

void* QTreeView_IndexBelow(void* ptr, void* index){
	return static_cast<QTreeView*>(ptr)->indexBelow(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QTreeView_IsColumnHidden(void* ptr, int column){
	return static_cast<QTreeView*>(ptr)->isColumnHidden(column);
}

int QTreeView_IsFirstColumnSpanned(void* ptr, int row, void* parent){
	return static_cast<QTreeView*>(ptr)->isFirstColumnSpanned(row, *static_cast<QModelIndex*>(parent));
}

int QTreeView_IsRowHidden(void* ptr, int row, void* parent){
	return static_cast<QTreeView*>(ptr)->isRowHidden(row, *static_cast<QModelIndex*>(parent));
}

void QTreeView_KeyboardSearch(void* ptr, char* search){
	static_cast<QTreeView*>(ptr)->keyboardSearch(QString(search));
}

void QTreeView_Reset(void* ptr){
	static_cast<QTreeView*>(ptr)->reset();
}

void QTreeView_ResizeColumnToContents(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "resizeColumnToContents", Q_ARG(int, column));
}

void QTreeView_ScrollTo(void* ptr, void* index, int hint){
	static_cast<QTreeView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QTreeView_SelectAll(void* ptr){
	static_cast<QTreeView*>(ptr)->selectAll();
}

void QTreeView_SetColumnHidden(void* ptr, int column, int hide){
	static_cast<QTreeView*>(ptr)->setColumnHidden(column, hide != 0);
}

void QTreeView_SetColumnWidth(void* ptr, int column, int width){
	static_cast<QTreeView*>(ptr)->setColumnWidth(column, width);
}

void QTreeView_SetExpanded(void* ptr, void* index, int expanded){
	static_cast<QTreeView*>(ptr)->setExpanded(*static_cast<QModelIndex*>(index), expanded != 0);
}

void QTreeView_SetFirstColumnSpanned(void* ptr, int row, void* parent, int span){
	static_cast<QTreeView*>(ptr)->setFirstColumnSpanned(row, *static_cast<QModelIndex*>(parent), span != 0);
}

void QTreeView_SetHeader(void* ptr, void* header){
	static_cast<QTreeView*>(ptr)->setHeader(static_cast<QHeaderView*>(header));
}

void QTreeView_SetModel(void* ptr, void* model){
	static_cast<QTreeView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QTreeView_SetRootIndex(void* ptr, void* index){
	static_cast<QTreeView*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QTreeView_SetRowHidden(void* ptr, int row, void* parent, int hide){
	static_cast<QTreeView*>(ptr)->setRowHidden(row, *static_cast<QModelIndex*>(parent), hide != 0);
}

void QTreeView_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<QTreeView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QTreeView_SetTreePosition(void* ptr, int index){
	static_cast<QTreeView*>(ptr)->setTreePosition(index);
}

void QTreeView_ShowColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "showColumn", Q_ARG(int, column));
}

void QTreeView_SortByColumn(void* ptr, int column, int order){
	static_cast<QTreeView*>(ptr)->sortByColumn(column, static_cast<Qt::SortOrder>(order));
}

int QTreeView_TreePosition(void* ptr){
	return static_cast<QTreeView*>(ptr)->treePosition();
}

void QTreeView_DestroyQTreeView(void* ptr){
	static_cast<QTreeView*>(ptr)->~QTreeView();
}

