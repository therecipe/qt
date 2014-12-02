#include "qabstractitemview.h"
#include <QAbstractItemView>
#include "cgoexport.h"

//MyQAbstractItemView
class MyQAbstractItemView: public QAbstractItemView {
Q_OBJECT
public:
void Signal_Activated() { callbackQAbstractItemView(this, QString("activated").toUtf8().data()); };
void Signal_Clicked() { callbackQAbstractItemView(this, QString("clicked").toUtf8().data()); };
void Signal_DoubleClicked() { callbackQAbstractItemView(this, QString("doubleClicked").toUtf8().data()); };
void Signal_Entered() { callbackQAbstractItemView(this, QString("entered").toUtf8().data()); };
void Signal_Pressed() { callbackQAbstractItemView(this, QString("pressed").toUtf8().data()); };
void Signal_ViewportEntered() { callbackQAbstractItemView(this, QString("viewportEntered").toUtf8().data()); };

signals:
void Slot_ClearSelection();
void Slot_Reset();
void Slot_ScrollToBottom();
void Slot_ScrollToTop();
void Slot_SelectAll();

};
#include "qabstractitemview.moc"

//Public Types
int QAbstractItemView_NoDragDrop() { return QAbstractItemView::NoDragDrop; }
int QAbstractItemView_DragOnly() { return QAbstractItemView::DragOnly; }
int QAbstractItemView_DropOnly() { return QAbstractItemView::DropOnly; }
int QAbstractItemView_DragDrop() { return QAbstractItemView::DragDrop; }
int QAbstractItemView_InternalMove() { return QAbstractItemView::InternalMove; }
int QAbstractItemView_NoEditTriggers() { return QAbstractItemView::NoEditTriggers; }
int QAbstractItemView_CurrentChanged() { return QAbstractItemView::CurrentChanged; }
int QAbstractItemView_DoubleClicked() { return QAbstractItemView::DoubleClicked; }
int QAbstractItemView_SelectedClicked() { return QAbstractItemView::SelectedClicked; }
int QAbstractItemView_EditKeyPressed() { return QAbstractItemView::EditKeyPressed; }
int QAbstractItemView_AnyKeyPressed() { return QAbstractItemView::AnyKeyPressed; }
int QAbstractItemView_AllEditTriggers() { return QAbstractItemView::AllEditTriggers; }
int QAbstractItemView_EnsureVisible() { return QAbstractItemView::EnsureVisible; }
int QAbstractItemView_PositionAtTop() { return QAbstractItemView::PositionAtTop; }
int QAbstractItemView_PositionAtBottom() { return QAbstractItemView::PositionAtBottom; }
int QAbstractItemView_PositionAtCenter() { return QAbstractItemView::PositionAtCenter; }
int QAbstractItemView_ScrollPerItem() { return QAbstractItemView::ScrollPerItem; }
int QAbstractItemView_ScrollPerPixel() { return QAbstractItemView::ScrollPerPixel; }
int QAbstractItemView_SelectItems() { return QAbstractItemView::SelectItems; }
int QAbstractItemView_SelectRows() { return QAbstractItemView::SelectRows; }
int QAbstractItemView_SelectColumns() { return QAbstractItemView::SelectColumns; }
int QAbstractItemView_SingleSelection() { return QAbstractItemView::SingleSelection; }
int QAbstractItemView_ContiguousSelection() { return QAbstractItemView::ContiguousSelection; }
int QAbstractItemView_ExtendedSelection() { return QAbstractItemView::ExtendedSelection; }
int QAbstractItemView_MultiSelection() { return QAbstractItemView::MultiSelection; }
int QAbstractItemView_NoSelection() { return QAbstractItemView::NoSelection; }

//Public Functions
void QAbstractItemView_Destroy(QtObjectPtr ptr)
{
	((QAbstractItemView*)(ptr))->~QAbstractItemView();
}

int QAbstractItemView_AlternatingRowColors(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->alternatingRowColors();
}

int QAbstractItemView_AutoScrollMargin(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->autoScrollMargin();
}

int QAbstractItemView_DefaultDropAction(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->defaultDropAction();
}

int QAbstractItemView_DragDropOverwriteMode(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->dragDropOverwriteMode();
}

int QAbstractItemView_DragEnabled(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->dragEnabled();
}

int QAbstractItemView_HasAutoScroll(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->hasAutoScroll();
}

void QAbstractItemView_KeyboardSearch_String(QtObjectPtr ptr, char* search)
{
	((QAbstractItemView*)(ptr))->keyboardSearch(QString(search));
}

int QAbstractItemView_SelectionMode(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->selectionMode();
}

void QAbstractItemView_SetAlternatingRowColors_Bool(QtObjectPtr ptr, int enable)
{
	((QAbstractItemView*)(ptr))->setAlternatingRowColors(enable != 0);
}

void QAbstractItemView_SetAutoScroll_Bool(QtObjectPtr ptr, int enable)
{
	((QAbstractItemView*)(ptr))->setAutoScroll(enable != 0);
}

void QAbstractItemView_SetAutoScrollMargin_Int(QtObjectPtr ptr, int margin)
{
	((QAbstractItemView*)(ptr))->setAutoScrollMargin(margin);
}

void QAbstractItemView_SetDefaultDropAction_DropAction(QtObjectPtr ptr, int dropAction)
{
	((QAbstractItemView*)(ptr))->setDefaultDropAction(((Qt::DropAction)(dropAction)));
}

void QAbstractItemView_SetDragDropOverwriteMode_Bool(QtObjectPtr ptr, int overwrite)
{
	((QAbstractItemView*)(ptr))->setDragDropOverwriteMode(overwrite != 0);
}

void QAbstractItemView_SetDragEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QAbstractItemView*)(ptr))->setDragEnabled(enable != 0);
}

void QAbstractItemView_SetDropIndicatorShown_Bool(QtObjectPtr ptr, int enable)
{
	((QAbstractItemView*)(ptr))->setDropIndicatorShown(enable != 0);
}

void QAbstractItemView_SetSelectionMode_SelectionMode(QtObjectPtr ptr, int mode)
{
	((QAbstractItemView*)(ptr))->setSelectionMode(((QAbstractItemView::SelectionMode)(mode)));
}

void QAbstractItemView_SetTabKeyNavigation_Bool(QtObjectPtr ptr, int enable)
{
	((QAbstractItemView*)(ptr))->setTabKeyNavigation(enable != 0);
}

void QAbstractItemView_SetTextElideMode_TextElideMode(QtObjectPtr ptr, int mode)
{
	((QAbstractItemView*)(ptr))->setTextElideMode(((Qt::TextElideMode)(mode)));
}

int QAbstractItemView_ShowDropIndicator(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->showDropIndicator();
}

int QAbstractItemView_SizeHintForColumn_Int(QtObjectPtr ptr, int column)
{
	return ((QAbstractItemView*)(ptr))->sizeHintForColumn(column);
}

int QAbstractItemView_SizeHintForRow_Int(QtObjectPtr ptr, int row)
{
	return ((QAbstractItemView*)(ptr))->sizeHintForRow(row);
}

int QAbstractItemView_TabKeyNavigation(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->tabKeyNavigation();
}

int QAbstractItemView_TextElideMode(QtObjectPtr ptr)
{
	return ((QAbstractItemView*)(ptr))->textElideMode();
}

//Public Slots
void QAbstractItemView_ConnectSlotClearSelection(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_ClearSelection, ((QAbstractItemView*)(ptr)), &QAbstractItemView::clearSelection, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSlotClearSelection(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_ClearSelection, ((QAbstractItemView*)(ptr)), &QAbstractItemView::clearSelection);
}

void QAbstractItemView_ClearSelection(QtObjectPtr ptr)
{
	((MyQAbstractItemView*)(ptr))->Slot_ClearSelection();
}

void QAbstractItemView_ConnectSlotReset(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_Reset, ((QAbstractItemView*)(ptr)), &QAbstractItemView::reset, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSlotReset(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_Reset, ((QAbstractItemView*)(ptr)), &QAbstractItemView::reset);
}

void QAbstractItemView_Reset(QtObjectPtr ptr)
{
	((MyQAbstractItemView*)(ptr))->Slot_Reset();
}

void QAbstractItemView_ConnectSlotScrollToBottom(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_ScrollToBottom, ((QAbstractItemView*)(ptr)), &QAbstractItemView::scrollToBottom, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSlotScrollToBottom(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_ScrollToBottom, ((QAbstractItemView*)(ptr)), &QAbstractItemView::scrollToBottom);
}

void QAbstractItemView_ScrollToBottom(QtObjectPtr ptr)
{
	((MyQAbstractItemView*)(ptr))->Slot_ScrollToBottom();
}

void QAbstractItemView_ConnectSlotScrollToTop(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_ScrollToTop, ((QAbstractItemView*)(ptr)), &QAbstractItemView::scrollToTop, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSlotScrollToTop(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_ScrollToTop, ((QAbstractItemView*)(ptr)), &QAbstractItemView::scrollToTop);
}

void QAbstractItemView_ScrollToTop(QtObjectPtr ptr)
{
	((MyQAbstractItemView*)(ptr))->Slot_ScrollToTop();
}

void QAbstractItemView_ConnectSlotSelectAll(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_SelectAll, ((QAbstractItemView*)(ptr)), &QAbstractItemView::selectAll, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSlotSelectAll(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Slot_SelectAll, ((QAbstractItemView*)(ptr)), &QAbstractItemView::selectAll);
}

void QAbstractItemView_SelectAll(QtObjectPtr ptr)
{
	((MyQAbstractItemView*)(ptr))->Slot_SelectAll();
}

//Signals
void QAbstractItemView_ConnectSignalActivated(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractItemView*)(ptr)), &QAbstractItemView::activated, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_Activated, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSignalActivated(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractItemView*)(ptr)), &QAbstractItemView::activated, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_Activated);
}

void QAbstractItemView_ConnectSignalClicked(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractItemView*)(ptr)), &QAbstractItemView::clicked, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_Clicked, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSignalClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractItemView*)(ptr)), &QAbstractItemView::clicked, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_Clicked);
}

void QAbstractItemView_ConnectSignalDoubleClicked(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractItemView*)(ptr)), &QAbstractItemView::doubleClicked, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_DoubleClicked, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSignalDoubleClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractItemView*)(ptr)), &QAbstractItemView::doubleClicked, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_DoubleClicked);
}

void QAbstractItemView_ConnectSignalEntered(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractItemView*)(ptr)), &QAbstractItemView::entered, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_Entered, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSignalEntered(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractItemView*)(ptr)), &QAbstractItemView::entered, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_Entered);
}

void QAbstractItemView_ConnectSignalPressed(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractItemView*)(ptr)), &QAbstractItemView::pressed, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_Pressed, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSignalPressed(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractItemView*)(ptr)), &QAbstractItemView::pressed, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_Pressed);
}

void QAbstractItemView_ConnectSignalViewportEntered(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractItemView*)(ptr)), &QAbstractItemView::viewportEntered, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_ViewportEntered, Qt::QueuedConnection);
}

void QAbstractItemView_DisconnectSignalViewportEntered(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractItemView*)(ptr)), &QAbstractItemView::viewportEntered, ((MyQAbstractItemView*)(ptr)), &MyQAbstractItemView::Signal_ViewportEntered);
}

