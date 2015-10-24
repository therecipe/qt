#include "qabstractitemdelegate.h"
#include <QPainter>
#include <QObject>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QEvent>
#include <QString>
#include <QStyleOptionViewItem>
#include <QAbstractItemModel>
#include <QHelpEvent>
#include <QStyle>
#include <QStyleOption>
#include <QAbstractItemView>
#include <QAbstractItemDelegate>
#include "_cgo_export.h"

class MyQAbstractItemDelegate: public QAbstractItemDelegate {
public:
void Signal_CloseEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint){callbackQAbstractItemDelegateCloseEditor(this->objectName().toUtf8().data(), editor, hint);};
void Signal_CommitData(QWidget * editor){callbackQAbstractItemDelegateCommitData(this->objectName().toUtf8().data(), editor);};
void Signal_SizeHintChanged(const QModelIndex & index){callbackQAbstractItemDelegateSizeHintChanged(this->objectName().toUtf8().data(), index.internalPointer());};
};

void QAbstractItemDelegate_ConnectCloseEditor(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&QAbstractItemDelegate::closeEditor), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&MyQAbstractItemDelegate::Signal_CloseEditor));;
}

void QAbstractItemDelegate_DisconnectCloseEditor(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&QAbstractItemDelegate::closeEditor), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&MyQAbstractItemDelegate::Signal_CloseEditor));;
}

void QAbstractItemDelegate_ConnectCommitData(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *)>(&QAbstractItemDelegate::commitData), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *)>(&MyQAbstractItemDelegate::Signal_CommitData));;
}

void QAbstractItemDelegate_DisconnectCommitData(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *)>(&QAbstractItemDelegate::commitData), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *)>(&MyQAbstractItemDelegate::Signal_CommitData));;
}

QtObjectPtr QAbstractItemDelegate_CreateEditor(QtObjectPtr ptr, QtObjectPtr parent, QtObjectPtr option, QtObjectPtr index){
	return static_cast<QAbstractItemDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_DestroyEditor(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr index){
	static_cast<QAbstractItemDelegate*>(ptr)->destroyEditor(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

int QAbstractItemDelegate_EditorEvent(QtObjectPtr ptr, QtObjectPtr event, QtObjectPtr model, QtObjectPtr option, QtObjectPtr index){
	return static_cast<QAbstractItemDelegate*>(ptr)->editorEvent(static_cast<QEvent*>(event), static_cast<QAbstractItemModel*>(model), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

int QAbstractItemDelegate_HelpEvent(QtObjectPtr ptr, QtObjectPtr event, QtObjectPtr view, QtObjectPtr option, QtObjectPtr index){
	return static_cast<QAbstractItemDelegate*>(ptr)->helpEvent(static_cast<QHelpEvent*>(event), static_cast<QAbstractItemView*>(view), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr index){
	static_cast<QAbstractItemDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_SetEditorData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr index){
	static_cast<QAbstractItemDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_SetModelData(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr model, QtObjectPtr index){
	static_cast<QAbstractItemDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_ConnectSizeHintChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(const QModelIndex &)>(&QAbstractItemDelegate::sizeHintChanged), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(const QModelIndex &)>(&MyQAbstractItemDelegate::Signal_SizeHintChanged));;
}

void QAbstractItemDelegate_DisconnectSizeHintChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(const QModelIndex &)>(&QAbstractItemDelegate::sizeHintChanged), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(const QModelIndex &)>(&MyQAbstractItemDelegate::Signal_SizeHintChanged));;
}

void QAbstractItemDelegate_UpdateEditorGeometry(QtObjectPtr ptr, QtObjectPtr editor, QtObjectPtr option, QtObjectPtr index){
	static_cast<QAbstractItemDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_DestroyQAbstractItemDelegate(QtObjectPtr ptr){
	static_cast<QAbstractItemDelegate*>(ptr)->~QAbstractItemDelegate();
}

