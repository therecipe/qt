#include "qtreewidget.h"
#include <QTreeWidget>
#include "cgoexport.h"

//MyQTreeWidget
class MyQTreeWidget: public QTreeWidget {
Q_OBJECT
public:
void Signal_CurrentItemChanged() { callbackQTreeWidget(this, QString("currentItemChanged").toUtf8().data()); };
void Signal_ItemActivated() { callbackQTreeWidget(this, QString("itemActivated").toUtf8().data()); };
void Signal_ItemChanged() { callbackQTreeWidget(this, QString("itemChanged").toUtf8().data()); };
void Signal_ItemClicked() { callbackQTreeWidget(this, QString("itemClicked").toUtf8().data()); };
void Signal_ItemCollapsed() { callbackQTreeWidget(this, QString("itemCollapsed").toUtf8().data()); };
void Signal_ItemDoubleClicked() { callbackQTreeWidget(this, QString("itemDoubleClicked").toUtf8().data()); };
void Signal_ItemEntered() { callbackQTreeWidget(this, QString("itemEntered").toUtf8().data()); };
void Signal_ItemExpanded() { callbackQTreeWidget(this, QString("itemExpanded").toUtf8().data()); };
void Signal_ItemPressed() { callbackQTreeWidget(this, QString("itemPressed").toUtf8().data()); };
void Signal_ItemSelectionChanged() { callbackQTreeWidget(this, QString("itemSelectionChanged").toUtf8().data()); };

signals:
void Slot_Clear();

};
#include "qtreewidget.moc"


//Public Functions
QtObjectPtr QTreeWidget_New_QWidget(QtObjectPtr parent)
{
	return new QTreeWidget(((QWidget*)(parent)));
}

void QTreeWidget_Destroy(QtObjectPtr ptr)
{
	((QTreeWidget*)(ptr))->~QTreeWidget();
}

void QTreeWidget_AddTopLevelItem_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QTreeWidget*)(ptr))->addTopLevelItem(((QTreeWidgetItem*)(item)));
}

void QTreeWidget_ClosePersistentEditor_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column)
{
	((QTreeWidget*)(ptr))->closePersistentEditor(((QTreeWidgetItem*)(item)), column);
}

int QTreeWidget_ColumnCount(QtObjectPtr ptr)
{
	return ((QTreeWidget*)(ptr))->columnCount();
}

int QTreeWidget_CurrentColumn(QtObjectPtr ptr)
{
	return ((QTreeWidget*)(ptr))->currentColumn();
}

QtObjectPtr QTreeWidget_CurrentItem(QtObjectPtr ptr)
{
	return ((QTreeWidget*)(ptr))->currentItem();
}

void QTreeWidget_EditItem_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column)
{
	((QTreeWidget*)(ptr))->editItem(((QTreeWidgetItem*)(item)), column);
}

QtObjectPtr QTreeWidget_HeaderItem(QtObjectPtr ptr)
{
	return ((QTreeWidget*)(ptr))->headerItem();
}

int QTreeWidget_IndexOfTopLevelItem_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	return ((QTreeWidget*)(ptr))->indexOfTopLevelItem(((QTreeWidgetItem*)(item)));
}

void QTreeWidget_InsertTopLevelItem_Int_QTreeWidgetItem(QtObjectPtr ptr, int index, QtObjectPtr item)
{
	((QTreeWidget*)(ptr))->insertTopLevelItem(index, ((QTreeWidgetItem*)(item)));
}

QtObjectPtr QTreeWidget_InvisibleRootItem(QtObjectPtr ptr)
{
	return ((QTreeWidget*)(ptr))->invisibleRootItem();
}

int QTreeWidget_IsFirstItemColumnSpanned_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	return ((QTreeWidget*)(ptr))->isFirstItemColumnSpanned(((QTreeWidgetItem*)(item)));
}

QtObjectPtr QTreeWidget_ItemAbove_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	return ((QTreeWidget*)(ptr))->itemAbove(((QTreeWidgetItem*)(item)));
}

QtObjectPtr QTreeWidget_ItemAt_Int_Int(QtObjectPtr ptr, int x, int y)
{
	return ((QTreeWidget*)(ptr))->itemAt(x, y);
}

QtObjectPtr QTreeWidget_ItemBelow_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	return ((QTreeWidget*)(ptr))->itemBelow(((QTreeWidgetItem*)(item)));
}

QtObjectPtr QTreeWidget_ItemWidget_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column)
{
	return ((QTreeWidget*)(ptr))->itemWidget(((QTreeWidgetItem*)(item)), column);
}

void QTreeWidget_OpenPersistentEditor_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column)
{
	((QTreeWidget*)(ptr))->openPersistentEditor(((QTreeWidgetItem*)(item)), column);
}

void QTreeWidget_RemoveItemWidget_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column)
{
	((QTreeWidget*)(ptr))->removeItemWidget(((QTreeWidgetItem*)(item)), column);
}

void QTreeWidget_SetColumnCount_Int(QtObjectPtr ptr, int columns)
{
	((QTreeWidget*)(ptr))->setColumnCount(columns);
}

void QTreeWidget_SetCurrentItem_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QTreeWidget*)(ptr))->setCurrentItem(((QTreeWidgetItem*)(item)));
}

void QTreeWidget_SetCurrentItem_QTreeWidgetItem_Int(QtObjectPtr ptr, QtObjectPtr item, int column)
{
	((QTreeWidget*)(ptr))->setCurrentItem(((QTreeWidgetItem*)(item)), column);
}

void QTreeWidget_SetFirstItemColumnSpanned_QTreeWidgetItem_Bool(QtObjectPtr ptr, QtObjectPtr item, int span)
{
	((QTreeWidget*)(ptr))->setFirstItemColumnSpanned(((QTreeWidgetItem*)(item)), span != 0);
}

void QTreeWidget_SetHeaderItem_QTreeWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QTreeWidget*)(ptr))->setHeaderItem(((QTreeWidgetItem*)(item)));
}

void QTreeWidget_SetHeaderLabel_String(QtObjectPtr ptr, char* label)
{
	((QTreeWidget*)(ptr))->setHeaderLabel(QString(label));
}

void QTreeWidget_SetHeaderLabels_QStringList(QtObjectPtr ptr, char* labels)
{
	((QTreeWidget*)(ptr))->setHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTreeWidget_SetItemWidget_QTreeWidgetItem_Int_QWidget(QtObjectPtr ptr, QtObjectPtr item, int column, QtObjectPtr widget)
{
	((QTreeWidget*)(ptr))->setItemWidget(((QTreeWidgetItem*)(item)), column, ((QWidget*)(widget)));
}

int QTreeWidget_SortColumn(QtObjectPtr ptr)
{
	return ((QTreeWidget*)(ptr))->sortColumn();
}

void QTreeWidget_SortItems_Int_SortOrder(QtObjectPtr ptr, int column, int order)
{
	((QTreeWidget*)(ptr))->sortItems(column, ((Qt::SortOrder)(order)));
}

QtObjectPtr QTreeWidget_TakeTopLevelItem_Int(QtObjectPtr ptr, int index)
{
	return ((QTreeWidget*)(ptr))->takeTopLevelItem(index);
}

QtObjectPtr QTreeWidget_TopLevelItem_Int(QtObjectPtr ptr, int index)
{
	return ((QTreeWidget*)(ptr))->topLevelItem(index);
}

int QTreeWidget_TopLevelItemCount(QtObjectPtr ptr)
{
	return ((QTreeWidget*)(ptr))->topLevelItemCount();
}

//Public Slots
void QTreeWidget_ConnectSlotClear(QtObjectPtr ptr)
{
	QObject::connect(((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Slot_Clear, ((QTreeWidget*)(ptr)), &QTreeWidget::clear, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSlotClear(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Slot_Clear, ((QTreeWidget*)(ptr)), &QTreeWidget::clear);
}

void QTreeWidget_Clear(QtObjectPtr ptr)
{
	((MyQTreeWidget*)(ptr))->Slot_Clear();
}

//Signals
void QTreeWidget_ConnectSignalCurrentItemChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::currentItemChanged, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_CurrentItemChanged, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalCurrentItemChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::currentItemChanged, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_CurrentItemChanged);
}

void QTreeWidget_ConnectSignalItemActivated(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::itemActivated, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemActivated, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalItemActivated(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::itemActivated, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemActivated);
}

void QTreeWidget_ConnectSignalItemChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::itemChanged, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemChanged, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalItemChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::itemChanged, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemChanged);
}

void QTreeWidget_ConnectSignalItemClicked(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::itemClicked, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemClicked, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalItemClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::itemClicked, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemClicked);
}

void QTreeWidget_ConnectSignalItemCollapsed(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::itemCollapsed, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemCollapsed, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalItemCollapsed(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::itemCollapsed, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemCollapsed);
}

void QTreeWidget_ConnectSignalItemDoubleClicked(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::itemDoubleClicked, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemDoubleClicked, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalItemDoubleClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::itemDoubleClicked, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemDoubleClicked);
}

void QTreeWidget_ConnectSignalItemEntered(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::itemEntered, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemEntered, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalItemEntered(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::itemEntered, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemEntered);
}

void QTreeWidget_ConnectSignalItemExpanded(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::itemExpanded, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemExpanded, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalItemExpanded(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::itemExpanded, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemExpanded);
}

void QTreeWidget_ConnectSignalItemPressed(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::itemPressed, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemPressed, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalItemPressed(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::itemPressed, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemPressed);
}

void QTreeWidget_ConnectSignalItemSelectionChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTreeWidget*)(ptr)), &QTreeWidget::itemSelectionChanged, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemSelectionChanged, Qt::QueuedConnection);
}

void QTreeWidget_DisconnectSignalItemSelectionChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTreeWidget*)(ptr)), &QTreeWidget::itemSelectionChanged, ((MyQTreeWidget*)(ptr)), &MyQTreeWidget::Signal_ItemSelectionChanged);
}

