#include "qtabbar.h"
#include <QTabBar>
#include "cgoexport.h"

//MyQTabBar
class MyQTabBar: public QTabBar {
Q_OBJECT
public:
void Signal_CurrentChanged() { callbackQTabBar(this, QString("currentChanged").toUtf8().data()); };
void Signal_TabBarClicked() { callbackQTabBar(this, QString("tabBarClicked").toUtf8().data()); };
void Signal_TabBarDoubleClicked() { callbackQTabBar(this, QString("tabBarDoubleClicked").toUtf8().data()); };
void Signal_TabCloseRequested() { callbackQTabBar(this, QString("tabCloseRequested").toUtf8().data()); };
void Signal_TabMoved() { callbackQTabBar(this, QString("tabMoved").toUtf8().data()); };

signals:
void Slot_SetCurrentIndex(int index);

};
#include "qtabbar.moc"


//Public Functions
QtObjectPtr QTabBar_New_QWidget(QtObjectPtr parent)
{
	return new QTabBar(((QWidget*)(parent)));
}

void QTabBar_Destroy(QtObjectPtr ptr)
{
	((QTabBar*)(ptr))->~QTabBar();
}

int QTabBar_AddTab_String(QtObjectPtr ptr, char* text)
{
	return ((QTabBar*)(ptr))->addTab(QString(text));
}

int QTabBar_Count(QtObjectPtr ptr)
{
	return ((QTabBar*)(ptr))->count();
}

int QTabBar_CurrentIndex(QtObjectPtr ptr)
{
	return ((QTabBar*)(ptr))->currentIndex();
}

int QTabBar_DocumentMode(QtObjectPtr ptr)
{
	return ((QTabBar*)(ptr))->documentMode();
}

int QTabBar_DrawBase(QtObjectPtr ptr)
{
	return ((QTabBar*)(ptr))->drawBase();
}

int QTabBar_ElideMode(QtObjectPtr ptr)
{
	return ((QTabBar*)(ptr))->elideMode();
}

int QTabBar_Expanding(QtObjectPtr ptr)
{
	return ((QTabBar*)(ptr))->expanding();
}

int QTabBar_InsertTab_Int_String(QtObjectPtr ptr, int index, char* text)
{
	return ((QTabBar*)(ptr))->insertTab(index, QString(text));
}

int QTabBar_IsMovable(QtObjectPtr ptr)
{
	return ((QTabBar*)(ptr))->isMovable();
}

int QTabBar_IsTabEnabled_Int(QtObjectPtr ptr, int index)
{
	return ((QTabBar*)(ptr))->isTabEnabled(index);
}

void QTabBar_MoveTab_Int_Int(QtObjectPtr ptr, int from, int to)
{
	((QTabBar*)(ptr))->moveTab(from, to);
}

void QTabBar_RemoveTab_Int(QtObjectPtr ptr, int index)
{
	((QTabBar*)(ptr))->removeTab(index);
}

void QTabBar_SetDocumentMode_Bool(QtObjectPtr ptr, int set)
{
	((QTabBar*)(ptr))->setDocumentMode(set != 0);
}

void QTabBar_SetDrawBase_Bool(QtObjectPtr ptr, int drawTheBase)
{
	((QTabBar*)(ptr))->setDrawBase(drawTheBase != 0);
}

void QTabBar_SetElideMode_TextElideMode(QtObjectPtr ptr, int TextElideMode)
{
	((QTabBar*)(ptr))->setElideMode(((Qt::TextElideMode)(TextElideMode)));
}

void QTabBar_SetExpanding_Bool(QtObjectPtr ptr, int enabled)
{
	((QTabBar*)(ptr))->setExpanding(enabled != 0);
}

void QTabBar_SetMovable_Bool(QtObjectPtr ptr, int movable)
{
	((QTabBar*)(ptr))->setMovable(movable != 0);
}

void QTabBar_SetTabEnabled_Int_Bool(QtObjectPtr ptr, int index, int enabled)
{
	((QTabBar*)(ptr))->setTabEnabled(index, enabled != 0);
}

void QTabBar_SetTabText_Int_String(QtObjectPtr ptr, int index, char* text)
{
	((QTabBar*)(ptr))->setTabText(index, QString(text));
}

void QTabBar_SetTabToolTip_Int_String(QtObjectPtr ptr, int index, char* tip)
{
	((QTabBar*)(ptr))->setTabToolTip(index, QString(tip));
}

void QTabBar_SetTabWhatsThis_Int_String(QtObjectPtr ptr, int index, char* text)
{
	((QTabBar*)(ptr))->setTabWhatsThis(index, QString(text));
}

void QTabBar_SetTabsClosable_Bool(QtObjectPtr ptr, int closable)
{
	((QTabBar*)(ptr))->setTabsClosable(closable != 0);
}

void QTabBar_SetUsesScrollButtons_Bool(QtObjectPtr ptr, int useButtons)
{
	((QTabBar*)(ptr))->setUsesScrollButtons(useButtons != 0);
}

char* QTabBar_TabText_Int(QtObjectPtr ptr, int index)
{
	return ((QTabBar*)(ptr))->tabText(index).toUtf8().data();
}

char* QTabBar_TabToolTip_Int(QtObjectPtr ptr, int index)
{
	return ((QTabBar*)(ptr))->tabToolTip(index).toUtf8().data();
}

char* QTabBar_TabWhatsThis_Int(QtObjectPtr ptr, int index)
{
	return ((QTabBar*)(ptr))->tabWhatsThis(index).toUtf8().data();
}

int QTabBar_TabsClosable(QtObjectPtr ptr)
{
	return ((QTabBar*)(ptr))->tabsClosable();
}

int QTabBar_UsesScrollButtons(QtObjectPtr ptr)
{
	return ((QTabBar*)(ptr))->usesScrollButtons();
}

//Public Slots
void QTabBar_ConnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::connect(((MyQTabBar*)(ptr)), &MyQTabBar::Slot_SetCurrentIndex, ((QTabBar*)(ptr)), &QTabBar::setCurrentIndex, Qt::QueuedConnection);
}

void QTabBar_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTabBar*)(ptr)), &MyQTabBar::Slot_SetCurrentIndex, ((QTabBar*)(ptr)), &QTabBar::setCurrentIndex);
}

void QTabBar_SetCurrentIndex_Int(QtObjectPtr ptr, int index)
{
	((MyQTabBar*)(ptr))->Slot_SetCurrentIndex(index);
}

//Signals
void QTabBar_ConnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTabBar*)(ptr)), &QTabBar::currentChanged, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_CurrentChanged, Qt::QueuedConnection);
}

void QTabBar_DisconnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTabBar*)(ptr)), &QTabBar::currentChanged, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_CurrentChanged);
}

void QTabBar_ConnectSignalTabBarClicked(QtObjectPtr ptr)
{
	QObject::connect(((QTabBar*)(ptr)), &QTabBar::tabBarClicked, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_TabBarClicked, Qt::QueuedConnection);
}

void QTabBar_DisconnectSignalTabBarClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QTabBar*)(ptr)), &QTabBar::tabBarClicked, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_TabBarClicked);
}

void QTabBar_ConnectSignalTabBarDoubleClicked(QtObjectPtr ptr)
{
	QObject::connect(((QTabBar*)(ptr)), &QTabBar::tabBarDoubleClicked, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_TabBarDoubleClicked, Qt::QueuedConnection);
}

void QTabBar_DisconnectSignalTabBarDoubleClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QTabBar*)(ptr)), &QTabBar::tabBarDoubleClicked, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_TabBarDoubleClicked);
}

void QTabBar_ConnectSignalTabCloseRequested(QtObjectPtr ptr)
{
	QObject::connect(((QTabBar*)(ptr)), &QTabBar::tabCloseRequested, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_TabCloseRequested, Qt::QueuedConnection);
}

void QTabBar_DisconnectSignalTabCloseRequested(QtObjectPtr ptr)
{
	QObject::disconnect(((QTabBar*)(ptr)), &QTabBar::tabCloseRequested, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_TabCloseRequested);
}

void QTabBar_ConnectSignalTabMoved(QtObjectPtr ptr)
{
	QObject::connect(((QTabBar*)(ptr)), &QTabBar::tabMoved, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_TabMoved, Qt::QueuedConnection);
}

void QTabBar_DisconnectSignalTabMoved(QtObjectPtr ptr)
{
	QObject::disconnect(((QTabBar*)(ptr)), &QTabBar::tabMoved, ((MyQTabBar*)(ptr)), &MyQTabBar::Signal_TabMoved);
}

