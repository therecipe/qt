#include "qstyleditemdelegate.h"
#include <QStyleOption>
#include <QStyleOptionViewItem>
#include <QItemEditorFactory>
#include <QWidget>
#include <QAbstractItemModel>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QPainter>
#include <QObject>
#include <QUrl>
#include <QLocale>
#include <QStyle>
#include <QStyledItemDelegate>
#include "_cgo_export.h"

class MyQStyledItemDelegate: public QStyledItemDelegate {
public:
};

QtObjectPtr QStyledItemDelegate_NewQStyledItemDelegate(QtObjectPtr parent){
	return new QStyledItemDelegate(static_cast<QObject*>(parent));
}

QtObjectPtr QStyledItemDelegate_CreateEditor(QtObjectPtr ptr, QtObjectPtr parent, QtObjectPtr option, QtObjectPtr index){
	return static_cast<QStyledItemDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char* QStyledItemDelegate_DisplayText(QtObjectPtr ptr, char* value, QtObjectPtr locale){
	return static_cast<QStyledItemDelegate*>(ptr)->displayText(QVariant(value), *static_cast<QLocale*>(locale)).toUtf8().data();
}

QtObjectPtr QStyledItemDelegate_ItemEditorFactory(QtObjectPtr ptr){
	return static_cast<QStyledItemDelegate*>(ptr)->itemEditorFactory();
}

void QStyledItemDelegate_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr index){
	static_cast<QStyledItemDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_SetEditorData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr index){
	static_cast<QStyledItemDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_SetItemEditorFactory(QtObjectPtr ptr, QtObjectPtr factory){
	static_cast<QStyledItemDelegate*>(ptr)->setItemEditorFactory(static_cast<QItemEditorFactory*>(factory));
}

void QStyledItemDelegate_SetModelData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr model, QtObjectPtr index){
	static_cast<QStyledItemDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_UpdateEditorGeometry(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr option, QtObjectPtr index){
	static_cast<QStyledItemDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_DestroyQStyledItemDelegate(QtObjectPtr ptr){
	static_cast<QStyledItemDelegate*>(ptr)->~QStyledItemDelegate();
}

