#include "qcolumnview.h"
#include <QString>
#include <QVariant>
#include <QObject>
#include <QAbstractItemModel>
#include <QItemSelection>
#include <QAbstractItemView>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QItemSelectionModel>
#include <QPoint>
#include <QColumnView>
#include "_cgo_export.h"

class MyQColumnView: public QColumnView {
public:
void Signal_UpdatePreviewWidget(const QModelIndex & index){callbackQColumnViewUpdatePreviewWidget(this->objectName().toUtf8().data(), index.internalPointer());};
};

int QColumnView_ResizeGripsVisible(void* ptr){
	return static_cast<QColumnView*>(ptr)->resizeGripsVisible();
}

void QColumnView_SetResizeGripsVisible(void* ptr, int visible){
	static_cast<QColumnView*>(ptr)->setResizeGripsVisible(visible != 0);
}

void* QColumnView_NewQColumnView(void* parent){
	return new QColumnView(static_cast<QWidget*>(parent));
}

void* QColumnView_IndexAt(void* ptr, void* point){
	return static_cast<QColumnView*>(ptr)->indexAt(*static_cast<QPoint*>(point)).internalPointer();
}

void* QColumnView_PreviewWidget(void* ptr){
	return static_cast<QColumnView*>(ptr)->previewWidget();
}

void QColumnView_ScrollTo(void* ptr, void* index, int hint){
	static_cast<QColumnView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QColumnView_SelectAll(void* ptr){
	static_cast<QColumnView*>(ptr)->selectAll();
}

void QColumnView_SetModel(void* ptr, void* model){
	static_cast<QColumnView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QColumnView_SetPreviewWidget(void* ptr, void* widget){
	static_cast<QColumnView*>(ptr)->setPreviewWidget(static_cast<QWidget*>(widget));
}

void QColumnView_SetRootIndex(void* ptr, void* index){
	static_cast<QColumnView*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QColumnView_SetSelectionModel(void* ptr, void* newSelectionModel){
	static_cast<QColumnView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(newSelectionModel));
}

void QColumnView_ConnectUpdatePreviewWidget(void* ptr){
	QObject::connect(static_cast<QColumnView*>(ptr), static_cast<void (QColumnView::*)(const QModelIndex &)>(&QColumnView::updatePreviewWidget), static_cast<MyQColumnView*>(ptr), static_cast<void (MyQColumnView::*)(const QModelIndex &)>(&MyQColumnView::Signal_UpdatePreviewWidget));;
}

void QColumnView_DisconnectUpdatePreviewWidget(void* ptr){
	QObject::disconnect(static_cast<QColumnView*>(ptr), static_cast<void (QColumnView::*)(const QModelIndex &)>(&QColumnView::updatePreviewWidget), static_cast<MyQColumnView*>(ptr), static_cast<void (MyQColumnView::*)(const QModelIndex &)>(&MyQColumnView::Signal_UpdatePreviewWidget));;
}

void QColumnView_DestroyQColumnView(void* ptr){
	static_cast<QColumnView*>(ptr)->~QColumnView();
}

