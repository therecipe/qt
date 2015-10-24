#include "qlistview.h"
#include <QList>
#include <QWidget>
#include <QSize>
#include <QString>
#include <QPoint>
#include <QList>
#include <QAbstractItemView>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QListView>
#include "_cgo_export.h"

class MyQListView: public QListView {
public:
};

int QListView_BatchSize(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->batchSize();
}

int QListView_Flow(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->flow();
}

int QListView_IsSelectionRectVisible(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->isSelectionRectVisible();
}

int QListView_IsWrapping(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->isWrapping();
}

int QListView_LayoutMode(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->layoutMode();
}

int QListView_ModelColumn(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->modelColumn();
}

int QListView_Movement(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->movement();
}

int QListView_ResizeMode(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->resizeMode();
}

void QListView_SetBatchSize(QtObjectPtr ptr, int batchSize){
	static_cast<QListView*>(ptr)->setBatchSize(batchSize);
}

void QListView_SetFlow(QtObjectPtr ptr, int flow){
	static_cast<QListView*>(ptr)->setFlow(static_cast<QListView::Flow>(flow));
}

void QListView_SetGridSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QListView*>(ptr)->setGridSize(*static_cast<QSize*>(size));
}

void QListView_SetLayoutMode(QtObjectPtr ptr, int mode){
	static_cast<QListView*>(ptr)->setLayoutMode(static_cast<QListView::LayoutMode>(mode));
}

void QListView_SetModelColumn(QtObjectPtr ptr, int column){
	static_cast<QListView*>(ptr)->setModelColumn(column);
}

void QListView_SetMovement(QtObjectPtr ptr, int movement){
	static_cast<QListView*>(ptr)->setMovement(static_cast<QListView::Movement>(movement));
}

void QListView_SetResizeMode(QtObjectPtr ptr, int mode){
	static_cast<QListView*>(ptr)->setResizeMode(static_cast<QListView::ResizeMode>(mode));
}

void QListView_SetSelectionRectVisible(QtObjectPtr ptr, int show){
	static_cast<QListView*>(ptr)->setSelectionRectVisible(show != 0);
}

void QListView_SetSpacing(QtObjectPtr ptr, int space){
	static_cast<QListView*>(ptr)->setSpacing(space);
}

void QListView_SetUniformItemSizes(QtObjectPtr ptr, int enable){
	static_cast<QListView*>(ptr)->setUniformItemSizes(enable != 0);
}

void QListView_SetViewMode(QtObjectPtr ptr, int mode){
	static_cast<QListView*>(ptr)->setViewMode(static_cast<QListView::ViewMode>(mode));
}

void QListView_SetWordWrap(QtObjectPtr ptr, int on){
	static_cast<QListView*>(ptr)->setWordWrap(on != 0);
}

void QListView_SetWrapping(QtObjectPtr ptr, int enable){
	static_cast<QListView*>(ptr)->setWrapping(enable != 0);
}

int QListView_Spacing(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->spacing();
}

int QListView_UniformItemSizes(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->uniformItemSizes();
}

int QListView_ViewMode(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->viewMode();
}

int QListView_WordWrap(QtObjectPtr ptr){
	return static_cast<QListView*>(ptr)->wordWrap();
}

QtObjectPtr QListView_NewQListView(QtObjectPtr parent){
	return new QListView(static_cast<QWidget*>(parent));
}

void QListView_ClearPropertyFlags(QtObjectPtr ptr){
	static_cast<QListView*>(ptr)->clearPropertyFlags();
}

QtObjectPtr QListView_IndexAt(QtObjectPtr ptr, QtObjectPtr p){
	return static_cast<QListView*>(ptr)->indexAt(*static_cast<QPoint*>(p)).internalPointer();
}

int QListView_IsRowHidden(QtObjectPtr ptr, int row){
	return static_cast<QListView*>(ptr)->isRowHidden(row);
}

void QListView_ScrollTo(QtObjectPtr ptr, QtObjectPtr index, int hint){
	static_cast<QListView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QListView_SetRowHidden(QtObjectPtr ptr, int row, int hide){
	static_cast<QListView*>(ptr)->setRowHidden(row, hide != 0);
}

void QListView_DestroyQListView(QtObjectPtr ptr){
	static_cast<QListView*>(ptr)->~QListView();
}

