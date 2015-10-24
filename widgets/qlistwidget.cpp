#include "qlistwidget.h"
#include <QList>
#include <QAbstractItemView>
#include <QWidget>
#include <QMetaObject>
#include <QUrl>
#include <QList>
#include <QString>
#include <QModelIndex>
#include <QPoint>
#include <QVariant>
#include <QDropEvent>
#include <QItemSelectionModel>
#include <QListWidgetItem>
#include <QObject>
#include <QItemSelection>
#include <QListWidget>
#include "_cgo_export.h"

class MyQListWidget: public QListWidget {
public:
void Signal_CurrentItemChanged(QListWidgetItem * current, QListWidgetItem * previous){callbackQListWidgetCurrentItemChanged(this->objectName().toUtf8().data(), current, previous);};
void Signal_CurrentRowChanged(int currentRow){callbackQListWidgetCurrentRowChanged(this->objectName().toUtf8().data(), currentRow);};
void Signal_CurrentTextChanged(const QString & currentText){callbackQListWidgetCurrentTextChanged(this->objectName().toUtf8().data(), currentText.toUtf8().data());};
void Signal_ItemActivated(QListWidgetItem * item){callbackQListWidgetItemActivated(this->objectName().toUtf8().data(), item);};
void Signal_ItemChanged(QListWidgetItem * item){callbackQListWidgetItemChanged(this->objectName().toUtf8().data(), item);};
void Signal_ItemClicked(QListWidgetItem * item){callbackQListWidgetItemClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemDoubleClicked(QListWidgetItem * item){callbackQListWidgetItemDoubleClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemEntered(QListWidgetItem * item){callbackQListWidgetItemEntered(this->objectName().toUtf8().data(), item);};
void Signal_ItemPressed(QListWidgetItem * item){callbackQListWidgetItemPressed(this->objectName().toUtf8().data(), item);};
void Signal_ItemSelectionChanged(){callbackQListWidgetItemSelectionChanged(this->objectName().toUtf8().data());};
};

int QListWidget_Count(QtObjectPtr ptr){
	return static_cast<QListWidget*>(ptr)->count();
}

int QListWidget_CurrentRow(QtObjectPtr ptr){
	return static_cast<QListWidget*>(ptr)->currentRow();
}

int QListWidget_IsSortingEnabled(QtObjectPtr ptr){
	return static_cast<QListWidget*>(ptr)->isSortingEnabled();
}

void QListWidget_SetCurrentRow(QtObjectPtr ptr, int row){
	static_cast<QListWidget*>(ptr)->setCurrentRow(row);
}

void QListWidget_SetSortingEnabled(QtObjectPtr ptr, int enable){
	static_cast<QListWidget*>(ptr)->setSortingEnabled(enable != 0);
}

QtObjectPtr QListWidget_NewQListWidget(QtObjectPtr parent){
	return new QListWidget(static_cast<QWidget*>(parent));
}

void QListWidget_AddItem2(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QListWidget*>(ptr)->addItem(static_cast<QListWidgetItem*>(item));
}

void QListWidget_AddItem(QtObjectPtr ptr, char* label){
	static_cast<QListWidget*>(ptr)->addItem(QString(label));
}

void QListWidget_AddItems(QtObjectPtr ptr, char* labels){
	static_cast<QListWidget*>(ptr)->addItems(QString(labels).split("|", QString::SkipEmptyParts));
}

void QListWidget_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QListWidget*>(ptr), "clear");
}

void QListWidget_ClosePersistentEditor(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QListWidget*>(ptr)->closePersistentEditor(static_cast<QListWidgetItem*>(item));
}

QtObjectPtr QListWidget_CurrentItem(QtObjectPtr ptr){
	return static_cast<QListWidget*>(ptr)->currentItem();
}

void QListWidget_ConnectCurrentItemChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&QListWidget::currentItemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&MyQListWidget::Signal_CurrentItemChanged));;
}

void QListWidget_DisconnectCurrentItemChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&QListWidget::currentItemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&MyQListWidget::Signal_CurrentItemChanged));;
}

void QListWidget_ConnectCurrentRowChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(int)>(&QListWidget::currentRowChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(int)>(&MyQListWidget::Signal_CurrentRowChanged));;
}

void QListWidget_DisconnectCurrentRowChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(int)>(&QListWidget::currentRowChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(int)>(&MyQListWidget::Signal_CurrentRowChanged));;
}

void QListWidget_ConnectCurrentTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(const QString &)>(&QListWidget::currentTextChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(const QString &)>(&MyQListWidget::Signal_CurrentTextChanged));;
}

void QListWidget_DisconnectCurrentTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(const QString &)>(&QListWidget::currentTextChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(const QString &)>(&MyQListWidget::Signal_CurrentTextChanged));;
}

void QListWidget_DropEvent(QtObjectPtr ptr, QtObjectPtr event){
	static_cast<QListWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QListWidget_EditItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QListWidget*>(ptr)->editItem(static_cast<QListWidgetItem*>(item));
}

void QListWidget_InsertItem(QtObjectPtr ptr, int row, QtObjectPtr item){
	static_cast<QListWidget*>(ptr)->insertItem(row, static_cast<QListWidgetItem*>(item));
}

void QListWidget_InsertItem2(QtObjectPtr ptr, int row, char* label){
	static_cast<QListWidget*>(ptr)->insertItem(row, QString(label));
}

void QListWidget_InsertItems(QtObjectPtr ptr, int row, char* labels){
	static_cast<QListWidget*>(ptr)->insertItems(row, QString(labels).split("|", QString::SkipEmptyParts));
}

QtObjectPtr QListWidget_Item(QtObjectPtr ptr, int row){
	return static_cast<QListWidget*>(ptr)->item(row);
}

void QListWidget_ConnectItemActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemActivated), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemActivated));;
}

void QListWidget_DisconnectItemActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemActivated), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemActivated));;
}

QtObjectPtr QListWidget_ItemAt(QtObjectPtr ptr, QtObjectPtr p){
	return static_cast<QListWidget*>(ptr)->itemAt(*static_cast<QPoint*>(p));
}

QtObjectPtr QListWidget_ItemAt2(QtObjectPtr ptr, int x, int y){
	return static_cast<QListWidget*>(ptr)->itemAt(x, y);
}

void QListWidget_ConnectItemChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemChanged));;
}

void QListWidget_DisconnectItemChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemChanged));;
}

void QListWidget_ConnectItemClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemClicked));;
}

void QListWidget_DisconnectItemClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemClicked));;
}

void QListWidget_ConnectItemDoubleClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemDoubleClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemDoubleClicked));;
}

void QListWidget_DisconnectItemDoubleClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemDoubleClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemDoubleClicked));;
}

void QListWidget_ConnectItemEntered(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemEntered), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemEntered));;
}

void QListWidget_DisconnectItemEntered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemEntered), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemEntered));;
}

void QListWidget_ConnectItemPressed(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemPressed), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemPressed));;
}

void QListWidget_DisconnectItemPressed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemPressed), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemPressed));;
}

void QListWidget_ConnectItemSelectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)()>(&QListWidget::itemSelectionChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)()>(&MyQListWidget::Signal_ItemSelectionChanged));;
}

void QListWidget_DisconnectItemSelectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)()>(&QListWidget::itemSelectionChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)()>(&MyQListWidget::Signal_ItemSelectionChanged));;
}

QtObjectPtr QListWidget_ItemWidget(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QListWidget*>(ptr)->itemWidget(static_cast<QListWidgetItem*>(item));
}

void QListWidget_OpenPersistentEditor(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QListWidget*>(ptr)->openPersistentEditor(static_cast<QListWidgetItem*>(item));
}

void QListWidget_RemoveItemWidget(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QListWidget*>(ptr)->removeItemWidget(static_cast<QListWidgetItem*>(item));
}

int QListWidget_Row(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QListWidget*>(ptr)->row(static_cast<QListWidgetItem*>(item));
}

void QListWidget_ScrollToItem(QtObjectPtr ptr, QtObjectPtr item, int hint){
	QMetaObject::invokeMethod(static_cast<QListWidget*>(ptr), "scrollToItem", Q_ARG(QListWidgetItem*, static_cast<QListWidgetItem*>(item)), Q_ARG(QAbstractItemView::ScrollHint, static_cast<QAbstractItemView::ScrollHint>(hint)));
}

void QListWidget_SetCurrentItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QListWidget*>(ptr)->setCurrentItem(static_cast<QListWidgetItem*>(item));
}

void QListWidget_SetCurrentItem2(QtObjectPtr ptr, QtObjectPtr item, int command){
	static_cast<QListWidget*>(ptr)->setCurrentItem(static_cast<QListWidgetItem*>(item), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QListWidget_SetCurrentRow2(QtObjectPtr ptr, int row, int command){
	static_cast<QListWidget*>(ptr)->setCurrentRow(row, static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QListWidget_SetItemWidget(QtObjectPtr ptr, QtObjectPtr item, QtObjectPtr widget){
	static_cast<QListWidget*>(ptr)->setItemWidget(static_cast<QListWidgetItem*>(item), static_cast<QWidget*>(widget));
}

void QListWidget_SortItems(QtObjectPtr ptr, int order){
	static_cast<QListWidget*>(ptr)->sortItems(static_cast<Qt::SortOrder>(order));
}

QtObjectPtr QListWidget_TakeItem(QtObjectPtr ptr, int row){
	return static_cast<QListWidget*>(ptr)->takeItem(row);
}

void QListWidget_DestroyQListWidget(QtObjectPtr ptr){
	static_cast<QListWidget*>(ptr)->~QListWidget();
}

