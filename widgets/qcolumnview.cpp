#include "qcolumnview.h"
#include <QString>
#include <QUrl>
#include <QItemSelection>
#include <QAbstractItemView>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QPoint>
#include <QItemSelectionModel>
#include <QWidget>
#include <QAbstractItemModel>
#include <QColumnView>
#include "_cgo_export.h"

class MyQColumnView: public QColumnView {
public:
void Signal_UpdatePreviewWidget(const QModelIndex & index){callbackQColumnViewUpdatePreviewWidget(this->objectName().toUtf8().data(), index.internalPointer());};
};

int QColumnView_ResizeGripsVisible(QtObjectPtr ptr){
	return static_cast<QColumnView*>(ptr)->resizeGripsVisible();
}

void QColumnView_SetResizeGripsVisible(QtObjectPtr ptr, int visible){
	static_cast<QColumnView*>(ptr)->setResizeGripsVisible(visible != 0);
}

QtObjectPtr QColumnView_NewQColumnView(QtObjectPtr parent){
	return new QColumnView(static_cast<QWidget*>(parent));
}

QtObjectPtr QColumnView_IndexAt(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QColumnView*>(ptr)->indexAt(*static_cast<QPoint*>(point)).internalPointer();
}

QtObjectPtr QColumnView_PreviewWidget(QtObjectPtr ptr){
	return static_cast<QColumnView*>(ptr)->previewWidget();
}

void QColumnView_ScrollTo(QtObjectPtr ptr, QtObjectPtr index, int hint){
	static_cast<QColumnView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QColumnView_SelectAll(QtObjectPtr ptr){
	static_cast<QColumnView*>(ptr)->selectAll();
}

void QColumnView_SetModel(QtObjectPtr ptr, QtObjectPtr model){
	static_cast<QColumnView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QColumnView_SetPreviewWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QColumnView*>(ptr)->setPreviewWidget(static_cast<QWidget*>(widget));
}

void QColumnView_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index){
	static_cast<QColumnView*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QColumnView_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr newSelectionModel){
	static_cast<QColumnView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(newSelectionModel));
}

void QColumnView_ConnectUpdatePreviewWidget(QtObjectPtr ptr){
	QObject::connect(static_cast<QColumnView*>(ptr), static_cast<void (QColumnView::*)(const QModelIndex &)>(&QColumnView::updatePreviewWidget), static_cast<MyQColumnView*>(ptr), static_cast<void (MyQColumnView::*)(const QModelIndex &)>(&MyQColumnView::Signal_UpdatePreviewWidget));;
}

void QColumnView_DisconnectUpdatePreviewWidget(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QColumnView*>(ptr), static_cast<void (QColumnView::*)(const QModelIndex &)>(&QColumnView::updatePreviewWidget), static_cast<MyQColumnView*>(ptr), static_cast<void (MyQColumnView::*)(const QModelIndex &)>(&MyQColumnView::Signal_UpdatePreviewWidget));;
}

void QColumnView_DestroyQColumnView(QtObjectPtr ptr){
	static_cast<QColumnView*>(ptr)->~QColumnView();
}

