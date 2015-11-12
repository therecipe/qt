#include "qtreewidget.h"
#include <QTreeWidgetItem>
#include <QItemSelection>
#include <QVariant>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QObject>
#include <QItemSelectionModel>
#include <QPoint>
#include <QAbstractItemView>
#include <QString>
#include <QUrl>
#include <QTreeWidget>
#include "_cgo_export.h"

class MyQTreeWidget: public QTreeWidget {
public:
void Signal_CurrentItemChanged(QTreeWidgetItem * current, QTreeWidgetItem * previous){callbackQTreeWidgetCurrentItemChanged(this->objectName().toUtf8().data(), current, previous);};
void Signal_ItemActivated(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemActivated(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemChanged(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemChanged(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemClicked(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemClicked(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemCollapsed(QTreeWidgetItem * item){callbackQTreeWidgetItemCollapsed(this->objectName().toUtf8().data(), item);};
void Signal_ItemDoubleClicked(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemDoubleClicked(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemEntered(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemEntered(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemExpanded(QTreeWidgetItem * item){callbackQTreeWidgetItemExpanded(this->objectName().toUtf8().data(), item);};
void Signal_ItemPressed(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemPressed(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemSelectionChanged(){callbackQTreeWidgetItemSelectionChanged(this->objectName().toUtf8().data());};
};

int QTreeWidget_ColumnCount(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->columnCount();
}

void QTreeWidget_SetColumnCount(void* ptr, int columns){
	static_cast<QTreeWidget*>(ptr)->setColumnCount(columns);
}

int QTreeWidget_TopLevelItemCount(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->topLevelItemCount();
}

void* QTreeWidget_NewQTreeWidget(void* parent){
	return new QTreeWidget(static_cast<QWidget*>(parent));
}

void QTreeWidget_AddTopLevelItem(void* ptr, void* item){
	static_cast<QTreeWidget*>(ptr)->addTopLevelItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "clear");
}

void QTreeWidget_ClosePersistentEditor(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->closePersistentEditor(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_CollapseItem(void* ptr, void* item){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "collapseItem", Q_ARG(QTreeWidgetItem*, static_cast<QTreeWidgetItem*>(item)));
}

int QTreeWidget_CurrentColumn(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->currentColumn();
}

void* QTreeWidget_CurrentItem(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->currentItem();
}

void QTreeWidget_ConnectCurrentItemChanged(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&QTreeWidget::currentItemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&MyQTreeWidget::Signal_CurrentItemChanged));;
}

void QTreeWidget_DisconnectCurrentItemChanged(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&QTreeWidget::currentItemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&MyQTreeWidget::Signal_CurrentItemChanged));;
}

void QTreeWidget_EditItem(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->editItem(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_ExpandItem(void* ptr, void* item){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "expandItem", Q_ARG(QTreeWidgetItem*, static_cast<QTreeWidgetItem*>(item)));
}

void* QTreeWidget_HeaderItem(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->headerItem();
}

int QTreeWidget_IndexOfTopLevelItem(void* ptr, void* item){
	return static_cast<QTreeWidget*>(ptr)->indexOfTopLevelItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_InsertTopLevelItem(void* ptr, int index, void* item){
	static_cast<QTreeWidget*>(ptr)->insertTopLevelItem(index, static_cast<QTreeWidgetItem*>(item));
}

void* QTreeWidget_InvisibleRootItem(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->invisibleRootItem();
}

int QTreeWidget_IsFirstItemColumnSpanned(void* ptr, void* item){
	return static_cast<QTreeWidget*>(ptr)->isFirstItemColumnSpanned(static_cast<QTreeWidgetItem*>(item));
}

void* QTreeWidget_ItemAbove(void* ptr, void* item){
	return static_cast<QTreeWidget*>(ptr)->itemAbove(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_ConnectItemActivated(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemActivated), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemActivated));;
}

void QTreeWidget_DisconnectItemActivated(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemActivated), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemActivated));;
}

void* QTreeWidget_ItemAt(void* ptr, void* p){
	return static_cast<QTreeWidget*>(ptr)->itemAt(*static_cast<QPoint*>(p));
}

void* QTreeWidget_ItemAt2(void* ptr, int x, int y){
	return static_cast<QTreeWidget*>(ptr)->itemAt(x, y);
}

void* QTreeWidget_ItemBelow(void* ptr, void* item){
	return static_cast<QTreeWidget*>(ptr)->itemBelow(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_ConnectItemChanged(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemChanged));;
}

void QTreeWidget_DisconnectItemChanged(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemChanged));;
}

void QTreeWidget_ConnectItemClicked(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemClicked));;
}

void QTreeWidget_DisconnectItemClicked(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemClicked));;
}

void QTreeWidget_ConnectItemCollapsed(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemCollapsed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemCollapsed));;
}

void QTreeWidget_DisconnectItemCollapsed(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemCollapsed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemCollapsed));;
}

void QTreeWidget_ConnectItemDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemDoubleClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemDoubleClicked));;
}

void QTreeWidget_DisconnectItemDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemDoubleClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemDoubleClicked));;
}

void QTreeWidget_ConnectItemEntered(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemEntered), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemEntered));;
}

void QTreeWidget_DisconnectItemEntered(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemEntered), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemEntered));;
}

void QTreeWidget_ConnectItemExpanded(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemExpanded), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemExpanded));;
}

void QTreeWidget_DisconnectItemExpanded(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemExpanded), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemExpanded));;
}

void QTreeWidget_ConnectItemPressed(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemPressed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemPressed));;
}

void QTreeWidget_DisconnectItemPressed(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemPressed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemPressed));;
}

void QTreeWidget_ConnectItemSelectionChanged(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)()>(&QTreeWidget::itemSelectionChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)()>(&MyQTreeWidget::Signal_ItemSelectionChanged));;
}

void QTreeWidget_DisconnectItemSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)()>(&QTreeWidget::itemSelectionChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)()>(&MyQTreeWidget::Signal_ItemSelectionChanged));;
}

void* QTreeWidget_ItemWidget(void* ptr, void* item, int column){
	return static_cast<QTreeWidget*>(ptr)->itemWidget(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_OpenPersistentEditor(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->openPersistentEditor(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_RemoveItemWidget(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->removeItemWidget(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_ScrollToItem(void* ptr, void* item, int hint){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "scrollToItem", Q_ARG(QTreeWidgetItem*, static_cast<QTreeWidgetItem*>(item)), Q_ARG(QAbstractItemView::ScrollHint, static_cast<QAbstractItemView::ScrollHint>(hint)));
}

void QTreeWidget_SetCurrentItem(void* ptr, void* item){
	static_cast<QTreeWidget*>(ptr)->setCurrentItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_SetCurrentItem2(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->setCurrentItem(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_SetCurrentItem3(void* ptr, void* item, int column, int command){
	static_cast<QTreeWidget*>(ptr)->setCurrentItem(static_cast<QTreeWidgetItem*>(item), column, static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTreeWidget_SetFirstItemColumnSpanned(void* ptr, void* item, int span){
	static_cast<QTreeWidget*>(ptr)->setFirstItemColumnSpanned(static_cast<QTreeWidgetItem*>(item), span != 0);
}

void QTreeWidget_SetHeaderItem(void* ptr, void* item){
	static_cast<QTreeWidget*>(ptr)->setHeaderItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_SetHeaderLabel(void* ptr, char* label){
	static_cast<QTreeWidget*>(ptr)->setHeaderLabel(QString(label));
}

void QTreeWidget_SetHeaderLabels(void* ptr, char* labels){
	static_cast<QTreeWidget*>(ptr)->setHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTreeWidget_SetItemWidget(void* ptr, void* item, int column, void* widget){
	static_cast<QTreeWidget*>(ptr)->setItemWidget(static_cast<QTreeWidgetItem*>(item), column, static_cast<QWidget*>(widget));
}

void QTreeWidget_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<QTreeWidget*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

int QTreeWidget_SortColumn(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->sortColumn();
}

void QTreeWidget_SortItems(void* ptr, int column, int order){
	static_cast<QTreeWidget*>(ptr)->sortItems(column, static_cast<Qt::SortOrder>(order));
}

void* QTreeWidget_TakeTopLevelItem(void* ptr, int index){
	return static_cast<QTreeWidget*>(ptr)->takeTopLevelItem(index);
}

void* QTreeWidget_TopLevelItem(void* ptr, int index){
	return static_cast<QTreeWidget*>(ptr)->topLevelItem(index);
}

void QTreeWidget_DestroyQTreeWidget(void* ptr){
	static_cast<QTreeWidget*>(ptr)->~QTreeWidget();
}

