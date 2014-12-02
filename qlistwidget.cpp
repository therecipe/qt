#include "qlistwidget.h"
#include <QListWidget>
#include "cgoexport.h"

//MyQListWidget
class MyQListWidget: public QListWidget {
Q_OBJECT
public:
void Signal_CurrentItemChanged() { callbackQListWidget(this, QString("currentItemChanged").toUtf8().data()); };
void Signal_CurrentRowChanged() { callbackQListWidget(this, QString("currentRowChanged").toUtf8().data()); };
void Signal_CurrentTextChanged() { callbackQListWidget(this, QString("currentTextChanged").toUtf8().data()); };
void Signal_ItemActivated() { callbackQListWidget(this, QString("itemActivated").toUtf8().data()); };
void Signal_ItemChanged() { callbackQListWidget(this, QString("itemChanged").toUtf8().data()); };
void Signal_ItemClicked() { callbackQListWidget(this, QString("itemClicked").toUtf8().data()); };
void Signal_ItemDoubleClicked() { callbackQListWidget(this, QString("itemDoubleClicked").toUtf8().data()); };
void Signal_ItemEntered() { callbackQListWidget(this, QString("itemEntered").toUtf8().data()); };
void Signal_ItemPressed() { callbackQListWidget(this, QString("itemPressed").toUtf8().data()); };
void Signal_ItemSelectionChanged() { callbackQListWidget(this, QString("itemSelectionChanged").toUtf8().data()); };

signals:
void Slot_Clear();

};
#include "qlistwidget.moc"


//Public Functions
QtObjectPtr QListWidget_New_QWidget(QtObjectPtr parent)
{
	return new QListWidget(((QWidget*)(parent)));
}

void QListWidget_Destroy(QtObjectPtr ptr)
{
	((QListWidget*)(ptr))->~QListWidget();
}

void QListWidget_AddItem_String(QtObjectPtr ptr, char* label)
{
	((QListWidget*)(ptr))->addItem(QString(label));
}

void QListWidget_AddItem_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QListWidget*)(ptr))->addItem(((QListWidgetItem*)(item)));
}

void QListWidget_ClosePersistentEditor_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QListWidget*)(ptr))->closePersistentEditor(((QListWidgetItem*)(item)));
}

int QListWidget_Count(QtObjectPtr ptr)
{
	return ((QListWidget*)(ptr))->count();
}

QtObjectPtr QListWidget_CurrentItem(QtObjectPtr ptr)
{
	return ((QListWidget*)(ptr))->currentItem();
}

int QListWidget_CurrentRow(QtObjectPtr ptr)
{
	return ((QListWidget*)(ptr))->currentRow();
}

void QListWidget_EditItem_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QListWidget*)(ptr))->editItem(((QListWidgetItem*)(item)));
}

void QListWidget_InsertItem_Int_QListWidgetItem(QtObjectPtr ptr, int row, QtObjectPtr item)
{
	((QListWidget*)(ptr))->insertItem(row, ((QListWidgetItem*)(item)));
}

void QListWidget_InsertItem_Int_String(QtObjectPtr ptr, int row, char* label)
{
	((QListWidget*)(ptr))->insertItem(row, QString(label));
}

int QListWidget_IsSortingEnabled(QtObjectPtr ptr)
{
	return ((QListWidget*)(ptr))->isSortingEnabled();
}

QtObjectPtr QListWidget_Item_Int(QtObjectPtr ptr, int row)
{
	return ((QListWidget*)(ptr))->item(row);
}

QtObjectPtr QListWidget_ItemAt_Int_Int(QtObjectPtr ptr, int x, int y)
{
	return ((QListWidget*)(ptr))->itemAt(x, y);
}

QtObjectPtr QListWidget_ItemWidget_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	return ((QListWidget*)(ptr))->itemWidget(((QListWidgetItem*)(item)));
}

void QListWidget_OpenPersistentEditor_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QListWidget*)(ptr))->openPersistentEditor(((QListWidgetItem*)(item)));
}

void QListWidget_RemoveItemWidget_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QListWidget*)(ptr))->removeItemWidget(((QListWidgetItem*)(item)));
}

int QListWidget_Row_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	return ((QListWidget*)(ptr))->row(((QListWidgetItem*)(item)));
}

void QListWidget_SetCurrentItem_QListWidgetItem(QtObjectPtr ptr, QtObjectPtr item)
{
	((QListWidget*)(ptr))->setCurrentItem(((QListWidgetItem*)(item)));
}

void QListWidget_SetCurrentRow_Int(QtObjectPtr ptr, int row)
{
	((QListWidget*)(ptr))->setCurrentRow(row);
}

void QListWidget_SetItemWidget_QListWidgetItem_QWidget(QtObjectPtr ptr, QtObjectPtr item, QtObjectPtr widget)
{
	((QListWidget*)(ptr))->setItemWidget(((QListWidgetItem*)(item)), ((QWidget*)(widget)));
}

void QListWidget_SetSortingEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QListWidget*)(ptr))->setSortingEnabled(enable != 0);
}

void QListWidget_SortItems_SortOrder(QtObjectPtr ptr, int order)
{
	((QListWidget*)(ptr))->sortItems(((Qt::SortOrder)(order)));
}

QtObjectPtr QListWidget_TakeItem_Int(QtObjectPtr ptr, int row)
{
	return ((QListWidget*)(ptr))->takeItem(row);
}

//Public Slots
void QListWidget_ConnectSlotClear(QtObjectPtr ptr)
{
	QObject::connect(((MyQListWidget*)(ptr)), &MyQListWidget::Slot_Clear, ((QListWidget*)(ptr)), &QListWidget::clear, Qt::QueuedConnection);
}

void QListWidget_DisconnectSlotClear(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQListWidget*)(ptr)), &MyQListWidget::Slot_Clear, ((QListWidget*)(ptr)), &QListWidget::clear);
}

void QListWidget_Clear(QtObjectPtr ptr)
{
	((MyQListWidget*)(ptr))->Slot_Clear();
}

//Signals
void QListWidget_ConnectSignalCurrentItemChanged(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::currentItemChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_CurrentItemChanged, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalCurrentItemChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::currentItemChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_CurrentItemChanged);
}

void QListWidget_ConnectSignalCurrentRowChanged(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::currentRowChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_CurrentRowChanged, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalCurrentRowChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::currentRowChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_CurrentRowChanged);
}

void QListWidget_ConnectSignalCurrentTextChanged(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::currentTextChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_CurrentTextChanged, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalCurrentTextChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::currentTextChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_CurrentTextChanged);
}

void QListWidget_ConnectSignalItemActivated(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::itemActivated, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemActivated, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalItemActivated(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::itemActivated, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemActivated);
}

void QListWidget_ConnectSignalItemChanged(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::itemChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemChanged, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalItemChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::itemChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemChanged);
}

void QListWidget_ConnectSignalItemClicked(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::itemClicked, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemClicked, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalItemClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::itemClicked, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemClicked);
}

void QListWidget_ConnectSignalItemDoubleClicked(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::itemDoubleClicked, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemDoubleClicked, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalItemDoubleClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::itemDoubleClicked, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemDoubleClicked);
}

void QListWidget_ConnectSignalItemEntered(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::itemEntered, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemEntered, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalItemEntered(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::itemEntered, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemEntered);
}

void QListWidget_ConnectSignalItemPressed(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::itemPressed, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemPressed, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalItemPressed(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::itemPressed, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemPressed);
}

void QListWidget_ConnectSignalItemSelectionChanged(QtObjectPtr ptr)
{
	QObject::connect(((QListWidget*)(ptr)), &QListWidget::itemSelectionChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemSelectionChanged, Qt::QueuedConnection);
}

void QListWidget_DisconnectSignalItemSelectionChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QListWidget*)(ptr)), &QListWidget::itemSelectionChanged, ((MyQListWidget*)(ptr)), &MyQListWidget::Signal_ItemSelectionChanged);
}

