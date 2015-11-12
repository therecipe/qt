#include "qlistview.h"
#include <QList>
#include <QString>
#include <QUrl>
#include <QAbstractItemView>
#include <QWidget>
#include <QVariant>
#include <QModelIndex>
#include <QPoint>
#include <QSize>
#include <QList>
#include <QListView>
#include "_cgo_export.h"

class MyQListView: public QListView {
public:
};

int QListView_BatchSize(void* ptr){
	return static_cast<QListView*>(ptr)->batchSize();
}

int QListView_Flow(void* ptr){
	return static_cast<QListView*>(ptr)->flow();
}

int QListView_IsSelectionRectVisible(void* ptr){
	return static_cast<QListView*>(ptr)->isSelectionRectVisible();
}

int QListView_IsWrapping(void* ptr){
	return static_cast<QListView*>(ptr)->isWrapping();
}

int QListView_LayoutMode(void* ptr){
	return static_cast<QListView*>(ptr)->layoutMode();
}

int QListView_ModelColumn(void* ptr){
	return static_cast<QListView*>(ptr)->modelColumn();
}

int QListView_Movement(void* ptr){
	return static_cast<QListView*>(ptr)->movement();
}

int QListView_ResizeMode(void* ptr){
	return static_cast<QListView*>(ptr)->resizeMode();
}

void QListView_SetBatchSize(void* ptr, int batchSize){
	static_cast<QListView*>(ptr)->setBatchSize(batchSize);
}

void QListView_SetFlow(void* ptr, int flow){
	static_cast<QListView*>(ptr)->setFlow(static_cast<QListView::Flow>(flow));
}

void QListView_SetGridSize(void* ptr, void* size){
	static_cast<QListView*>(ptr)->setGridSize(*static_cast<QSize*>(size));
}

void QListView_SetLayoutMode(void* ptr, int mode){
	static_cast<QListView*>(ptr)->setLayoutMode(static_cast<QListView::LayoutMode>(mode));
}

void QListView_SetModelColumn(void* ptr, int column){
	static_cast<QListView*>(ptr)->setModelColumn(column);
}

void QListView_SetMovement(void* ptr, int movement){
	static_cast<QListView*>(ptr)->setMovement(static_cast<QListView::Movement>(movement));
}

void QListView_SetResizeMode(void* ptr, int mode){
	static_cast<QListView*>(ptr)->setResizeMode(static_cast<QListView::ResizeMode>(mode));
}

void QListView_SetSelectionRectVisible(void* ptr, int show){
	static_cast<QListView*>(ptr)->setSelectionRectVisible(show != 0);
}

void QListView_SetSpacing(void* ptr, int space){
	static_cast<QListView*>(ptr)->setSpacing(space);
}

void QListView_SetUniformItemSizes(void* ptr, int enable){
	static_cast<QListView*>(ptr)->setUniformItemSizes(enable != 0);
}

void QListView_SetViewMode(void* ptr, int mode){
	static_cast<QListView*>(ptr)->setViewMode(static_cast<QListView::ViewMode>(mode));
}

void QListView_SetWordWrap(void* ptr, int on){
	static_cast<QListView*>(ptr)->setWordWrap(on != 0);
}

void QListView_SetWrapping(void* ptr, int enable){
	static_cast<QListView*>(ptr)->setWrapping(enable != 0);
}

int QListView_Spacing(void* ptr){
	return static_cast<QListView*>(ptr)->spacing();
}

int QListView_UniformItemSizes(void* ptr){
	return static_cast<QListView*>(ptr)->uniformItemSizes();
}

int QListView_ViewMode(void* ptr){
	return static_cast<QListView*>(ptr)->viewMode();
}

int QListView_WordWrap(void* ptr){
	return static_cast<QListView*>(ptr)->wordWrap();
}

void* QListView_NewQListView(void* parent){
	return new QListView(static_cast<QWidget*>(parent));
}

void QListView_ClearPropertyFlags(void* ptr){
	static_cast<QListView*>(ptr)->clearPropertyFlags();
}

void* QListView_IndexAt(void* ptr, void* p){
	return static_cast<QListView*>(ptr)->indexAt(*static_cast<QPoint*>(p)).internalPointer();
}

int QListView_IsRowHidden(void* ptr, int row){
	return static_cast<QListView*>(ptr)->isRowHidden(row);
}

void QListView_ScrollTo(void* ptr, void* index, int hint){
	static_cast<QListView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QListView_SetRowHidden(void* ptr, int row, int hide){
	static_cast<QListView*>(ptr)->setRowHidden(row, hide != 0);
}

void QListView_DestroyQListView(void* ptr){
	static_cast<QListView*>(ptr)->~QListView();
}

