#include "qtreeview.h"
#include <QTreeView>
#include "cgoexport.h"

//MyQTreeView
class MyQTreeView: public QTreeView {
Q_OBJECT
public:

signals:
void Slot_CollapseAll();
void Slot_ExpandAll();
void Slot_ExpandToDepth(int depth);
void Slot_HideColumn(int column);
void Slot_ResizeColumnToContents(int column);
void Slot_ShowColumn(int column);

};
#include "qtreeview.moc"


//Public Functions
QtObjectPtr QTreeView_New_QWidget(QtObjectPtr parent)
{
	return new QTreeView(((QWidget*)(parent)));
}

void QTreeView_Destroy(QtObjectPtr ptr)
{
	((QTreeView*)(ptr))->~QTreeView();
}

int QTreeView_AllColumnsShowFocus(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->allColumnsShowFocus();
}

int QTreeView_AutoExpandDelay(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->autoExpandDelay();
}

int QTreeView_ColumnAt_Int(QtObjectPtr ptr, int x)
{
	return ((QTreeView*)(ptr))->columnAt(x);
}

int QTreeView_ColumnViewportPosition_Int(QtObjectPtr ptr, int column)
{
	return ((QTreeView*)(ptr))->columnViewportPosition(column);
}

int QTreeView_ColumnWidth_Int(QtObjectPtr ptr, int column)
{
	return ((QTreeView*)(ptr))->columnWidth(column);
}

int QTreeView_ExpandsOnDoubleClick(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->expandsOnDoubleClick();
}

int QTreeView_Indentation(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->indentation();
}

int QTreeView_IsAnimated(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->isAnimated();
}

int QTreeView_IsColumnHidden_Int(QtObjectPtr ptr, int column)
{
	return ((QTreeView*)(ptr))->isColumnHidden(column);
}

int QTreeView_IsHeaderHidden(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->isHeaderHidden();
}

int QTreeView_IsSortingEnabled(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->isSortingEnabled();
}

int QTreeView_ItemsExpandable(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->itemsExpandable();
}

int QTreeView_RootIsDecorated(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->rootIsDecorated();
}

void QTreeView_SetAllColumnsShowFocus_Bool(QtObjectPtr ptr, int enable)
{
	((QTreeView*)(ptr))->setAllColumnsShowFocus(enable != 0);
}

void QTreeView_SetAnimated_Bool(QtObjectPtr ptr, int enable)
{
	((QTreeView*)(ptr))->setAnimated(enable != 0);
}

void QTreeView_SetAutoExpandDelay_Int(QtObjectPtr ptr, int delay)
{
	((QTreeView*)(ptr))->setAutoExpandDelay(delay);
}

void QTreeView_SetColumnHidden_Int_Bool(QtObjectPtr ptr, int column, int hide)
{
	((QTreeView*)(ptr))->setColumnHidden(column, hide != 0);
}

void QTreeView_SetColumnWidth_Int_Int(QtObjectPtr ptr, int column, int width)
{
	((QTreeView*)(ptr))->setColumnWidth(column, width);
}

void QTreeView_SetExpandsOnDoubleClick_Bool(QtObjectPtr ptr, int enable)
{
	((QTreeView*)(ptr))->setExpandsOnDoubleClick(enable != 0);
}

void QTreeView_SetHeaderHidden_Bool(QtObjectPtr ptr, int hide)
{
	((QTreeView*)(ptr))->setHeaderHidden(hide != 0);
}

void QTreeView_SetIndentation_Int(QtObjectPtr ptr, int i)
{
	((QTreeView*)(ptr))->setIndentation(i);
}

void QTreeView_SetItemsExpandable_Bool(QtObjectPtr ptr, int enable)
{
	((QTreeView*)(ptr))->setItemsExpandable(enable != 0);
}

void QTreeView_SetRootIsDecorated_Bool(QtObjectPtr ptr, int show)
{
	((QTreeView*)(ptr))->setRootIsDecorated(show != 0);
}

void QTreeView_SetSortingEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QTreeView*)(ptr))->setSortingEnabled(enable != 0);
}

void QTreeView_SetTreePosition_Int(QtObjectPtr ptr, int index)
{
	((QTreeView*)(ptr))->setTreePosition(index);
}

void QTreeView_SetUniformRowHeights_Bool(QtObjectPtr ptr, int uniform)
{
	((QTreeView*)(ptr))->setUniformRowHeights(uniform != 0);
}

void QTreeView_SetWordWrap_Bool(QtObjectPtr ptr, int on)
{
	((QTreeView*)(ptr))->setWordWrap(on != 0);
}

void QTreeView_SortByColumn_Int_SortOrder(QtObjectPtr ptr, int column, int order)
{
	((QTreeView*)(ptr))->sortByColumn(column, ((Qt::SortOrder)(order)));
}

int QTreeView_TreePosition(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->treePosition();
}

int QTreeView_UniformRowHeights(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->uniformRowHeights();
}

int QTreeView_WordWrap(QtObjectPtr ptr)
{
	return ((QTreeView*)(ptr))->wordWrap();
}

//Public Slots
void QTreeView_ConnectSlotCollapseAll(QtObjectPtr ptr)
{
	QObject::connect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_CollapseAll, ((QTreeView*)(ptr)), &QTreeView::collapseAll, Qt::QueuedConnection);
}

void QTreeView_DisconnectSlotCollapseAll(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_CollapseAll, ((QTreeView*)(ptr)), &QTreeView::collapseAll);
}

void QTreeView_CollapseAll(QtObjectPtr ptr)
{
	((MyQTreeView*)(ptr))->Slot_CollapseAll();
}

void QTreeView_ConnectSlotExpandAll(QtObjectPtr ptr)
{
	QObject::connect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_ExpandAll, ((QTreeView*)(ptr)), &QTreeView::expandAll, Qt::QueuedConnection);
}

void QTreeView_DisconnectSlotExpandAll(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_ExpandAll, ((QTreeView*)(ptr)), &QTreeView::expandAll);
}

void QTreeView_ExpandAll(QtObjectPtr ptr)
{
	((MyQTreeView*)(ptr))->Slot_ExpandAll();
}

void QTreeView_ConnectSlotExpandToDepth(QtObjectPtr ptr)
{
	QObject::connect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_ExpandToDepth, ((QTreeView*)(ptr)), &QTreeView::expandToDepth, Qt::QueuedConnection);
}

void QTreeView_DisconnectSlotExpandToDepth(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_ExpandToDepth, ((QTreeView*)(ptr)), &QTreeView::expandToDepth);
}

void QTreeView_ExpandToDepth_Int(QtObjectPtr ptr, int depth)
{
	((MyQTreeView*)(ptr))->Slot_ExpandToDepth(depth);
}

void QTreeView_ConnectSlotHideColumn(QtObjectPtr ptr)
{
	QObject::connect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_HideColumn, ((QTreeView*)(ptr)), &QTreeView::hideColumn, Qt::QueuedConnection);
}

void QTreeView_DisconnectSlotHideColumn(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_HideColumn, ((QTreeView*)(ptr)), &QTreeView::hideColumn);
}

void QTreeView_HideColumn_Int(QtObjectPtr ptr, int column)
{
	((MyQTreeView*)(ptr))->Slot_HideColumn(column);
}

void QTreeView_ConnectSlotResizeColumnToContents(QtObjectPtr ptr)
{
	QObject::connect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_ResizeColumnToContents, ((QTreeView*)(ptr)), &QTreeView::resizeColumnToContents, Qt::QueuedConnection);
}

void QTreeView_DisconnectSlotResizeColumnToContents(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_ResizeColumnToContents, ((QTreeView*)(ptr)), &QTreeView::resizeColumnToContents);
}

void QTreeView_ResizeColumnToContents_Int(QtObjectPtr ptr, int column)
{
	((MyQTreeView*)(ptr))->Slot_ResizeColumnToContents(column);
}

void QTreeView_ConnectSlotShowColumn(QtObjectPtr ptr)
{
	QObject::connect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_ShowColumn, ((QTreeView*)(ptr)), &QTreeView::showColumn, Qt::QueuedConnection);
}

void QTreeView_DisconnectSlotShowColumn(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTreeView*)(ptr)), &MyQTreeView::Slot_ShowColumn, ((QTreeView*)(ptr)), &QTreeView::showColumn);
}

void QTreeView_ShowColumn_Int(QtObjectPtr ptr, int column)
{
	((MyQTreeView*)(ptr))->Slot_ShowColumn(column);
}

