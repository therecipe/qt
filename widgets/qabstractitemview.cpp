#include "qabstractitemview.h"
#include <QString>
#include <QAbstractItemDelegate>
#include <QSize>
#include <QPoint>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QWidget>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QMetaObject>
#include <QObject>
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

int QAbstractItemView_AlternatingRowColors(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->alternatingRowColors();
}

int QAbstractItemView_AutoScrollMargin(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->autoScrollMargin();
}

int QAbstractItemView_DefaultDropAction(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->defaultDropAction();
}

int QAbstractItemView_DragDropMode(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->dragDropMode();
}

int QAbstractItemView_DragDropOverwriteMode(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->dragDropOverwriteMode();
}

int QAbstractItemView_DragEnabled(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->dragEnabled();
}

int QAbstractItemView_EditTriggers(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->editTriggers();
}

int QAbstractItemView_HasAutoScroll(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->hasAutoScroll();
}

int QAbstractItemView_HorizontalScrollMode(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->horizontalScrollMode();
}

int QAbstractItemView_SelectionBehavior(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->selectionBehavior();
}

int QAbstractItemView_SelectionMode(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->selectionMode();
}

void QAbstractItemView_SetAlternatingRowColors(QtObjectPtr ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setAlternatingRowColors(enable != 0);
}

void QAbstractItemView_SetAutoScroll(QtObjectPtr ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setAutoScroll(enable != 0);
}

void QAbstractItemView_SetAutoScrollMargin(QtObjectPtr ptr, int margin){
	static_cast<QAbstractItemView*>(ptr)->setAutoScrollMargin(margin);
}

void QAbstractItemView_SetDefaultDropAction(QtObjectPtr ptr, int dropAction){
	static_cast<QAbstractItemView*>(ptr)->setDefaultDropAction(static_cast<Qt::DropAction>(dropAction));
}

void QAbstractItemView_SetDragDropMode(QtObjectPtr ptr, int behavior){
	static_cast<QAbstractItemView*>(ptr)->setDragDropMode(static_cast<QAbstractItemView::DragDropMode>(behavior));
}

void QAbstractItemView_SetDragDropOverwriteMode(QtObjectPtr ptr, int overwrite){
	static_cast<QAbstractItemView*>(ptr)->setDragDropOverwriteMode(overwrite != 0);
}

void QAbstractItemView_SetDragEnabled(QtObjectPtr ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setDragEnabled(enable != 0);
}

void QAbstractItemView_SetDropIndicatorShown(QtObjectPtr ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setDropIndicatorShown(enable != 0);
}

void QAbstractItemView_SetEditTriggers(QtObjectPtr ptr, int triggers){
	static_cast<QAbstractItemView*>(ptr)->setEditTriggers(static_cast<QAbstractItemView::EditTrigger>(triggers));
}

void QAbstractItemView_SetHorizontalScrollMode(QtObjectPtr ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setHorizontalScrollMode(static_cast<QAbstractItemView::ScrollMode>(mode));
}

void QAbstractItemView_SetIconSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QAbstractItemView*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QAbstractItemView_SetSelectionBehavior(QtObjectPtr ptr, int behavior){
	static_cast<QAbstractItemView*>(ptr)->setSelectionBehavior(static_cast<QAbstractItemView::SelectionBehavior>(behavior));
}

void QAbstractItemView_SetSelectionMode(QtObjectPtr ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setSelectionMode(static_cast<QAbstractItemView::SelectionMode>(mode));
}

void QAbstractItemView_SetTabKeyNavigation(QtObjectPtr ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setTabKeyNavigation(enable != 0);
}

void QAbstractItemView_SetTextElideMode(QtObjectPtr ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setTextElideMode(static_cast<Qt::TextElideMode>(mode));
}

void QAbstractItemView_SetVerticalScrollMode(QtObjectPtr ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setVerticalScrollMode(static_cast<QAbstractItemView::ScrollMode>(mode));
}

int QAbstractItemView_ShowDropIndicator(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->showDropIndicator();
}

int QAbstractItemView_TabKeyNavigation(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->tabKeyNavigation();
}

int QAbstractItemView_TextElideMode(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->textElideMode();
}

int QAbstractItemView_VerticalScrollMode(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->verticalScrollMode();
}

void QAbstractItemView_ConnectActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::activated), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Activated));;
}

void QAbstractItemView_DisconnectActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::activated), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Activated));;
}

void QAbstractItemView_ClearSelection(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "clearSelection");
}

void QAbstractItemView_ConnectClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::clicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Clicked));;
}

void QAbstractItemView_DisconnectClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::clicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Clicked));;
}

void QAbstractItemView_ClosePersistentEditor(QtObjectPtr ptr, QtObjectPtr index){
	static_cast<QAbstractItemView*>(ptr)->closePersistentEditor(*static_cast<QModelIndex*>(index));
}

QtObjectPtr QAbstractItemView_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->currentIndex().internalPointer();
}

void QAbstractItemView_ConnectDoubleClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::doubleClicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_DoubleClicked));;
}

void QAbstractItemView_DisconnectDoubleClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::doubleClicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_DoubleClicked));;
}

void QAbstractItemView_Edit(QtObjectPtr ptr, QtObjectPtr index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "edit", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_ConnectEntered(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::entered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Entered));;
}

void QAbstractItemView_DisconnectEntered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::entered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Entered));;
}

QtObjectPtr QAbstractItemView_IndexAt(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QAbstractItemView*>(ptr)->indexAt(*static_cast<QPoint*>(point)).internalPointer();
}

QtObjectPtr QAbstractItemView_IndexWidget(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QAbstractItemView*>(ptr)->indexWidget(*static_cast<QModelIndex*>(index));
}

char* QAbstractItemView_InputMethodQuery(QtObjectPtr ptr, int query){
	return static_cast<QAbstractItemView*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)).toString().toUtf8().data();
}

QtObjectPtr QAbstractItemView_ItemDelegate(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegate();
}

QtObjectPtr QAbstractItemView_ItemDelegate2(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegate(*static_cast<QModelIndex*>(index));
}

QtObjectPtr QAbstractItemView_ItemDelegateForColumn(QtObjectPtr ptr, int column){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegateForColumn(column);
}

QtObjectPtr QAbstractItemView_ItemDelegateForRow(QtObjectPtr ptr, int row){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegateForRow(row);
}

void QAbstractItemView_KeyboardSearch(QtObjectPtr ptr, char* search){
	static_cast<QAbstractItemView*>(ptr)->keyboardSearch(QString(search));
}

QtObjectPtr QAbstractItemView_Model(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->model();
}

void QAbstractItemView_OpenPersistentEditor(QtObjectPtr ptr, QtObjectPtr index){
	static_cast<QAbstractItemView*>(ptr)->openPersistentEditor(*static_cast<QModelIndex*>(index));
}

void QAbstractItemView_ConnectPressed(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::pressed), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Pressed));;
}

void QAbstractItemView_DisconnectPressed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::pressed), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Pressed));;
}

void QAbstractItemView_Reset(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "reset");
}

QtObjectPtr QAbstractItemView_RootIndex(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->rootIndex().internalPointer();
}

void QAbstractItemView_ScrollTo(QtObjectPtr ptr, QtObjectPtr index, int hint){
	static_cast<QAbstractItemView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QAbstractItemView_ScrollToBottom(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "scrollToBottom");
}

void QAbstractItemView_ScrollToTop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "scrollToTop");
}

void QAbstractItemView_SelectAll(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "selectAll");
}

QtObjectPtr QAbstractItemView_SelectionModel(QtObjectPtr ptr){
	return static_cast<QAbstractItemView*>(ptr)->selectionModel();
}

void QAbstractItemView_SetCurrentIndex(QtObjectPtr ptr, QtObjectPtr index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "setCurrentIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_SetIndexWidget(QtObjectPtr ptr, QtObjectPtr index, QtObjectPtr widget){
	static_cast<QAbstractItemView*>(ptr)->setIndexWidget(*static_cast<QModelIndex*>(index), static_cast<QWidget*>(widget));
}

void QAbstractItemView_SetItemDelegate(QtObjectPtr ptr, QtObjectPtr delegate){
	static_cast<QAbstractItemView*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QAbstractItemView_SetItemDelegateForColumn(QtObjectPtr ptr, int column, QtObjectPtr delegate){
	static_cast<QAbstractItemView*>(ptr)->setItemDelegateForColumn(column, static_cast<QAbstractItemDelegate*>(delegate));
}

void QAbstractItemView_SetItemDelegateForRow(QtObjectPtr ptr, int row, QtObjectPtr delegate){
	static_cast<QAbstractItemView*>(ptr)->setItemDelegateForRow(row, static_cast<QAbstractItemDelegate*>(delegate));
}

void QAbstractItemView_SetModel(QtObjectPtr ptr, QtObjectPtr model){
	static_cast<QAbstractItemView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QAbstractItemView_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "setRootIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr selectionModel){
	static_cast<QAbstractItemView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

int QAbstractItemView_SizeHintForColumn(QtObjectPtr ptr, int column){
	return static_cast<QAbstractItemView*>(ptr)->sizeHintForColumn(column);
}

int QAbstractItemView_SizeHintForRow(QtObjectPtr ptr, int row){
	return static_cast<QAbstractItemView*>(ptr)->sizeHintForRow(row);
}

void QAbstractItemView_Update(QtObjectPtr ptr, QtObjectPtr index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "update", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_ConnectViewportEntered(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)()>(&QAbstractItemView::viewportEntered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)()>(&MyQAbstractItemView::Signal_ViewportEntered));;
}

void QAbstractItemView_DisconnectViewportEntered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)()>(&QAbstractItemView::viewportEntered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)()>(&MyQAbstractItemView::Signal_ViewportEntered));;
}

void QAbstractItemView_DestroyQAbstractItemView(QtObjectPtr ptr){
	static_cast<QAbstractItemView*>(ptr)->~QAbstractItemView();
}

