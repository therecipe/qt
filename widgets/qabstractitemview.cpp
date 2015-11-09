#include "qabstractitemview.h"
#include <QString>
#include <QModelIndex>
#include <QWidget>
#include <QAbstractItemModel>
#include <QPoint>
#include <QAbstractItemDelegate>
#include <QVariant>
#include <QUrl>
#include <QObject>
#include <QSize>
#include <QItemSelection>
#include <QMetaObject>
#include <QItemSelectionModel>
#include <QAbstractItemView>
#include "_cgo_export.h"

class MyQAbstractItemView: public QAbstractItemView {
public:
void Signal_Activated(const QModelIndex & index){callbackQAbstractItemViewActivated(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_Clicked(const QModelIndex & index){callbackQAbstractItemViewClicked(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_DoubleClicked(const QModelIndex & index){callbackQAbstractItemViewDoubleClicked(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_Entered(const QModelIndex & index){callbackQAbstractItemViewEntered(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_Pressed(const QModelIndex & index){callbackQAbstractItemViewPressed(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_ViewportEntered(){callbackQAbstractItemViewViewportEntered(this->objectName().toUtf8().data());};
};

int QAbstractItemView_AlternatingRowColors(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->alternatingRowColors();
}

int QAbstractItemView_AutoScrollMargin(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->autoScrollMargin();
}

int QAbstractItemView_DefaultDropAction(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->defaultDropAction();
}

int QAbstractItemView_DragDropMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->dragDropMode();
}

int QAbstractItemView_DragDropOverwriteMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->dragDropOverwriteMode();
}

int QAbstractItemView_DragEnabled(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->dragEnabled();
}

int QAbstractItemView_EditTriggers(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->editTriggers();
}

int QAbstractItemView_HasAutoScroll(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->hasAutoScroll();
}

int QAbstractItemView_HorizontalScrollMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->horizontalScrollMode();
}

int QAbstractItemView_SelectionBehavior(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->selectionBehavior();
}

int QAbstractItemView_SelectionMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->selectionMode();
}

void QAbstractItemView_SetAlternatingRowColors(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setAlternatingRowColors(enable != 0);
}

void QAbstractItemView_SetAutoScroll(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setAutoScroll(enable != 0);
}

void QAbstractItemView_SetAutoScrollMargin(void* ptr, int margin){
	static_cast<QAbstractItemView*>(ptr)->setAutoScrollMargin(margin);
}

void QAbstractItemView_SetDefaultDropAction(void* ptr, int dropAction){
	static_cast<QAbstractItemView*>(ptr)->setDefaultDropAction(static_cast<Qt::DropAction>(dropAction));
}

void QAbstractItemView_SetDragDropMode(void* ptr, int behavior){
	static_cast<QAbstractItemView*>(ptr)->setDragDropMode(static_cast<QAbstractItemView::DragDropMode>(behavior));
}

void QAbstractItemView_SetDragDropOverwriteMode(void* ptr, int overwrite){
	static_cast<QAbstractItemView*>(ptr)->setDragDropOverwriteMode(overwrite != 0);
}

void QAbstractItemView_SetDragEnabled(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setDragEnabled(enable != 0);
}

void QAbstractItemView_SetDropIndicatorShown(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setDropIndicatorShown(enable != 0);
}

void QAbstractItemView_SetEditTriggers(void* ptr, int triggers){
	static_cast<QAbstractItemView*>(ptr)->setEditTriggers(static_cast<QAbstractItemView::EditTrigger>(triggers));
}

void QAbstractItemView_SetHorizontalScrollMode(void* ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setHorizontalScrollMode(static_cast<QAbstractItemView::ScrollMode>(mode));
}

void QAbstractItemView_SetIconSize(void* ptr, void* size){
	static_cast<QAbstractItemView*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QAbstractItemView_SetSelectionBehavior(void* ptr, int behavior){
	static_cast<QAbstractItemView*>(ptr)->setSelectionBehavior(static_cast<QAbstractItemView::SelectionBehavior>(behavior));
}

void QAbstractItemView_SetSelectionMode(void* ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setSelectionMode(static_cast<QAbstractItemView::SelectionMode>(mode));
}

void QAbstractItemView_SetTabKeyNavigation(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setTabKeyNavigation(enable != 0);
}

void QAbstractItemView_SetTextElideMode(void* ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setTextElideMode(static_cast<Qt::TextElideMode>(mode));
}

void QAbstractItemView_SetVerticalScrollMode(void* ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setVerticalScrollMode(static_cast<QAbstractItemView::ScrollMode>(mode));
}

int QAbstractItemView_ShowDropIndicator(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->showDropIndicator();
}

int QAbstractItemView_TabKeyNavigation(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->tabKeyNavigation();
}

int QAbstractItemView_TextElideMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->textElideMode();
}

int QAbstractItemView_VerticalScrollMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->verticalScrollMode();
}

void QAbstractItemView_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::activated), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Activated));;
}

void QAbstractItemView_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::activated), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Activated));;
}

void QAbstractItemView_ClearSelection(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "clearSelection");
}

void QAbstractItemView_ConnectClicked(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::clicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Clicked));;
}

void QAbstractItemView_DisconnectClicked(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::clicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Clicked));;
}

void QAbstractItemView_ClosePersistentEditor(void* ptr, void* index){
	static_cast<QAbstractItemView*>(ptr)->closePersistentEditor(*static_cast<QModelIndex*>(index));
}

void* QAbstractItemView_CurrentIndex(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->currentIndex().internalPointer();
}

void QAbstractItemView_ConnectDoubleClicked(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::doubleClicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_DoubleClicked));;
}

void QAbstractItemView_DisconnectDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::doubleClicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_DoubleClicked));;
}

void QAbstractItemView_Edit(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "edit", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_ConnectEntered(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::entered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Entered));;
}

void QAbstractItemView_DisconnectEntered(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::entered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Entered));;
}

void* QAbstractItemView_IndexAt(void* ptr, void* point){
	return static_cast<QAbstractItemView*>(ptr)->indexAt(*static_cast<QPoint*>(point)).internalPointer();
}

void* QAbstractItemView_IndexWidget(void* ptr, void* index){
	return static_cast<QAbstractItemView*>(ptr)->indexWidget(*static_cast<QModelIndex*>(index));
}

void* QAbstractItemView_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QAbstractItemView*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QAbstractItemView_ItemDelegate(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegate();
}

void* QAbstractItemView_ItemDelegate2(void* ptr, void* index){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegate(*static_cast<QModelIndex*>(index));
}

void* QAbstractItemView_ItemDelegateForColumn(void* ptr, int column){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegateForColumn(column);
}

void* QAbstractItemView_ItemDelegateForRow(void* ptr, int row){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegateForRow(row);
}

void QAbstractItemView_KeyboardSearch(void* ptr, char* search){
	static_cast<QAbstractItemView*>(ptr)->keyboardSearch(QString(search));
}

void* QAbstractItemView_Model(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->model();
}

void QAbstractItemView_OpenPersistentEditor(void* ptr, void* index){
	static_cast<QAbstractItemView*>(ptr)->openPersistentEditor(*static_cast<QModelIndex*>(index));
}

void QAbstractItemView_ConnectPressed(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::pressed), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Pressed));;
}

void QAbstractItemView_DisconnectPressed(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::pressed), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Pressed));;
}

void QAbstractItemView_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "reset");
}

void* QAbstractItemView_RootIndex(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->rootIndex().internalPointer();
}

void QAbstractItemView_ScrollTo(void* ptr, void* index, int hint){
	static_cast<QAbstractItemView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QAbstractItemView_ScrollToBottom(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "scrollToBottom");
}

void QAbstractItemView_ScrollToTop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "scrollToTop");
}

void QAbstractItemView_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "selectAll");
}

void* QAbstractItemView_SelectionModel(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->selectionModel();
}

void QAbstractItemView_SetCurrentIndex(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "setCurrentIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_SetIndexWidget(void* ptr, void* index, void* widget){
	static_cast<QAbstractItemView*>(ptr)->setIndexWidget(*static_cast<QModelIndex*>(index), static_cast<QWidget*>(widget));
}

void QAbstractItemView_SetItemDelegate(void* ptr, void* delegate){
	static_cast<QAbstractItemView*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QAbstractItemView_SetItemDelegateForColumn(void* ptr, int column, void* delegate){
	static_cast<QAbstractItemView*>(ptr)->setItemDelegateForColumn(column, static_cast<QAbstractItemDelegate*>(delegate));
}

void QAbstractItemView_SetItemDelegateForRow(void* ptr, int row, void* delegate){
	static_cast<QAbstractItemView*>(ptr)->setItemDelegateForRow(row, static_cast<QAbstractItemDelegate*>(delegate));
}

void QAbstractItemView_SetModel(void* ptr, void* model){
	static_cast<QAbstractItemView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QAbstractItemView_SetRootIndex(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "setRootIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<QAbstractItemView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

int QAbstractItemView_SizeHintForColumn(void* ptr, int column){
	return static_cast<QAbstractItemView*>(ptr)->sizeHintForColumn(column);
}

int QAbstractItemView_SizeHintForRow(void* ptr, int row){
	return static_cast<QAbstractItemView*>(ptr)->sizeHintForRow(row);
}

void QAbstractItemView_Update(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "update", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_ConnectViewportEntered(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)()>(&QAbstractItemView::viewportEntered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)()>(&MyQAbstractItemView::Signal_ViewportEntered));;
}

void QAbstractItemView_DisconnectViewportEntered(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)()>(&QAbstractItemView::viewportEntered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)()>(&MyQAbstractItemView::Signal_ViewportEntered));;
}

void QAbstractItemView_DestroyQAbstractItemView(void* ptr){
	static_cast<QAbstractItemView*>(ptr)->~QAbstractItemView();
}

