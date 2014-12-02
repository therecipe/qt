#include "qlistview.h"
#include <QListView>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QListView_New_QWidget(QtObjectPtr parent)
{
	return new QListView(((QWidget*)(parent)));
}

void QListView_Destroy(QtObjectPtr ptr)
{
	((QListView*)(ptr))->~QListView();
}

int QListView_BatchSize(QtObjectPtr ptr)
{
	return ((QListView*)(ptr))->batchSize();
}

void QListView_ClearPropertyFlags(QtObjectPtr ptr)
{
	((QListView*)(ptr))->clearPropertyFlags();
}

int QListView_IsRowHidden_Int(QtObjectPtr ptr, int row)
{
	return ((QListView*)(ptr))->isRowHidden(row);
}

int QListView_IsSelectionRectVisible(QtObjectPtr ptr)
{
	return ((QListView*)(ptr))->isSelectionRectVisible();
}

int QListView_IsWrapping(QtObjectPtr ptr)
{
	return ((QListView*)(ptr))->isWrapping();
}

int QListView_ModelColumn(QtObjectPtr ptr)
{
	return ((QListView*)(ptr))->modelColumn();
}

void QListView_SetBatchSize_Int(QtObjectPtr ptr, int batchSize)
{
	((QListView*)(ptr))->setBatchSize(batchSize);
}

void QListView_SetModelColumn_Int(QtObjectPtr ptr, int column)
{
	((QListView*)(ptr))->setModelColumn(column);
}

void QListView_SetRowHidden_Int_Bool(QtObjectPtr ptr, int row, int hide)
{
	((QListView*)(ptr))->setRowHidden(row, hide != 0);
}

void QListView_SetSelectionRectVisible_Bool(QtObjectPtr ptr, int show)
{
	((QListView*)(ptr))->setSelectionRectVisible(show != 0);
}

void QListView_SetSpacing_Int(QtObjectPtr ptr, int space)
{
	((QListView*)(ptr))->setSpacing(space);
}

void QListView_SetUniformItemSizes_Bool(QtObjectPtr ptr, int enable)
{
	((QListView*)(ptr))->setUniformItemSizes(enable != 0);
}

void QListView_SetWordWrap_Bool(QtObjectPtr ptr, int on)
{
	((QListView*)(ptr))->setWordWrap(on != 0);
}

void QListView_SetWrapping_Bool(QtObjectPtr ptr, int enable)
{
	((QListView*)(ptr))->setWrapping(enable != 0);
}

int QListView_Spacing(QtObjectPtr ptr)
{
	return ((QListView*)(ptr))->spacing();
}

int QListView_UniformItemSizes(QtObjectPtr ptr)
{
	return ((QListView*)(ptr))->uniformItemSizes();
}

int QListView_WordWrap(QtObjectPtr ptr)
{
	return ((QListView*)(ptr))->wordWrap();
}

