#include "qitemdelegate.h"
#include <QWidget>
#include <QAbstractItemModel>
#include <QString>
#include <QModelIndex>
#include <QStyle>
#include <QObject>
#include <QStyleOption>
#include <QStyleOptionViewItem>
#include <QVariant>
#include <QUrl>
#include <QPainter>
#include <QItemEditorFactory>
#include <QItemDelegate>
#include "_cgo_export.h"

class MyQItemDelegate: public QItemDelegate {
public:
};

int QItemDelegate_HasClipping(QtObjectPtr ptr){
	return static_cast<QItemDelegate*>(ptr)->hasClipping();
}

void QItemDelegate_SetClipping(QtObjectPtr ptr, int clip){
	static_cast<QItemDelegate*>(ptr)->setClipping(clip != 0);
}

QtObjectPtr QItemDelegate_NewQItemDelegate(QtObjectPtr parent){
	return new QItemDelegate(static_cast<QObject*>(parent));
}

QtObjectPtr QItemDelegate_CreateEditor(QtObjectPtr ptr, QtObjectPtr parent, QtObjectPtr option, QtObjectPtr index){
	return static_cast<QItemDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

QtObjectPtr QItemDelegate_ItemEditorFactory(QtObjectPtr ptr){
	return static_cast<QItemDelegate*>(ptr)->itemEditorFactory();
}

void QItemDelegate_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr index){
	static_cast<QItemDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_SetEditorData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr index){
	static_cast<QItemDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_SetItemEditorFactory(QtObjectPtr ptr, QtObjectPtr factory){
	static_cast<QItemDelegate*>(ptr)->setItemEditorFactory(static_cast<QItemEditorFactory*>(factory));
}

void QItemDelegate_SetModelData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr model, QtObjectPtr index){
	static_cast<QItemDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_UpdateEditorGeometry(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr option, QtObjectPtr index){
	static_cast<QItemDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_DestroyQItemDelegate(QtObjectPtr ptr){
	static_cast<QItemDelegate*>(ptr)->~QItemDelegate();
}

