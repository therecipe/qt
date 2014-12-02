#include "qtabwidget.h"
#include <QTabWidget>
#include "cgoexport.h"

//MyQTabWidget
class MyQTabWidget: public QTabWidget {
Q_OBJECT
public:
void Signal_CurrentChanged() { callbackQTabWidget(this, QString("currentChanged").toUtf8().data()); };
void Signal_TabBarClicked() { callbackQTabWidget(this, QString("tabBarClicked").toUtf8().data()); };
void Signal_TabBarDoubleClicked() { callbackQTabWidget(this, QString("tabBarDoubleClicked").toUtf8().data()); };
void Signal_TabCloseRequested() { callbackQTabWidget(this, QString("tabCloseRequested").toUtf8().data()); };

signals:
void Slot_SetCurrentIndex(int index);

};
#include "qtabwidget.moc"


//Public Functions
QtObjectPtr QTabWidget_New_QWidget(QtObjectPtr parent)
{
	return new QTabWidget(((QWidget*)(parent)));
}

void QTabWidget_Destroy(QtObjectPtr ptr)
{
	((QTabWidget*)(ptr))->~QTabWidget();
}

int QTabWidget_AddTab_QWidget_String(QtObjectPtr ptr, QtObjectPtr page, char* label)
{
	return ((QTabWidget*)(ptr))->addTab(((QWidget*)(page)), QString(label));
}

void QTabWidget_Clear(QtObjectPtr ptr)
{
	((QTabWidget*)(ptr))->clear();
}

QtObjectPtr QTabWidget_CornerWidget_Corner(QtObjectPtr ptr, int corner)
{
	return ((QTabWidget*)(ptr))->cornerWidget(((Qt::Corner)(corner)));
}

int QTabWidget_Count(QtObjectPtr ptr)
{
	return ((QTabWidget*)(ptr))->count();
}

int QTabWidget_CurrentIndex(QtObjectPtr ptr)
{
	return ((QTabWidget*)(ptr))->currentIndex();
}

QtObjectPtr QTabWidget_CurrentWidget(QtObjectPtr ptr)
{
	return ((QTabWidget*)(ptr))->currentWidget();
}

int QTabWidget_DocumentMode(QtObjectPtr ptr)
{
	return ((QTabWidget*)(ptr))->documentMode();
}

int QTabWidget_ElideMode(QtObjectPtr ptr)
{
	return ((QTabWidget*)(ptr))->elideMode();
}

int QTabWidget_IndexOf_QWidget(QtObjectPtr ptr, QtObjectPtr w)
{
	return ((QTabWidget*)(ptr))->indexOf(((QWidget*)(w)));
}

int QTabWidget_InsertTab_Int_QWidget_String(QtObjectPtr ptr, int index, QtObjectPtr page, char* label)
{
	return ((QTabWidget*)(ptr))->insertTab(index, ((QWidget*)(page)), QString(label));
}

int QTabWidget_IsMovable(QtObjectPtr ptr)
{
	return ((QTabWidget*)(ptr))->isMovable();
}

int QTabWidget_IsTabEnabled_Int(QtObjectPtr ptr, int index)
{
	return ((QTabWidget*)(ptr))->isTabEnabled(index);
}

void QTabWidget_RemoveTab_Int(QtObjectPtr ptr, int index)
{
	((QTabWidget*)(ptr))->removeTab(index);
}

void QTabWidget_SetCornerWidget_QWidget_Corner(QtObjectPtr ptr, QtObjectPtr widget, int corner)
{
	((QTabWidget*)(ptr))->setCornerWidget(((QWidget*)(widget)), ((Qt::Corner)(corner)));
}

void QTabWidget_SetDocumentMode_Bool(QtObjectPtr ptr, int set)
{
	((QTabWidget*)(ptr))->setDocumentMode(set != 0);
}

void QTabWidget_SetElideMode_TextElideMode(QtObjectPtr ptr, int TextElideMode)
{
	((QTabWidget*)(ptr))->setElideMode(((Qt::TextElideMode)(TextElideMode)));
}

void QTabWidget_SetMovable_Bool(QtObjectPtr ptr, int movable)
{
	((QTabWidget*)(ptr))->setMovable(movable != 0);
}

void QTabWidget_SetTabEnabled_Int_Bool(QtObjectPtr ptr, int index, int enable)
{
	((QTabWidget*)(ptr))->setTabEnabled(index, enable != 0);
}

void QTabWidget_SetTabText_Int_String(QtObjectPtr ptr, int index, char* label)
{
	((QTabWidget*)(ptr))->setTabText(index, QString(label));
}

void QTabWidget_SetTabToolTip_Int_String(QtObjectPtr ptr, int index, char* tip)
{
	((QTabWidget*)(ptr))->setTabToolTip(index, QString(tip));
}

void QTabWidget_SetTabWhatsThis_Int_String(QtObjectPtr ptr, int index, char* text)
{
	((QTabWidget*)(ptr))->setTabWhatsThis(index, QString(text));
}

void QTabWidget_SetTabsClosable_Bool(QtObjectPtr ptr, int closeable)
{
	((QTabWidget*)(ptr))->setTabsClosable(closeable != 0);
}

void QTabWidget_SetUsesScrollButtons_Bool(QtObjectPtr ptr, int useButtons)
{
	((QTabWidget*)(ptr))->setUsesScrollButtons(useButtons != 0);
}

QtObjectPtr QTabWidget_TabBar(QtObjectPtr ptr)
{
	return ((QTabWidget*)(ptr))->tabBar();
}

char* QTabWidget_TabText_Int(QtObjectPtr ptr, int index)
{
	return ((QTabWidget*)(ptr))->tabText(index).toUtf8().data();
}

char* QTabWidget_TabToolTip_Int(QtObjectPtr ptr, int index)
{
	return ((QTabWidget*)(ptr))->tabToolTip(index).toUtf8().data();
}

char* QTabWidget_TabWhatsThis_Int(QtObjectPtr ptr, int index)
{
	return ((QTabWidget*)(ptr))->tabWhatsThis(index).toUtf8().data();
}

int QTabWidget_TabsClosable(QtObjectPtr ptr)
{
	return ((QTabWidget*)(ptr))->tabsClosable();
}

int QTabWidget_UsesScrollButtons(QtObjectPtr ptr)
{
	return ((QTabWidget*)(ptr))->usesScrollButtons();
}

QtObjectPtr QTabWidget_Widget_Int(QtObjectPtr ptr, int index)
{
	return ((QTabWidget*)(ptr))->widget(index);
}

//Public Slots
void QTabWidget_ConnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::connect(((MyQTabWidget*)(ptr)), &MyQTabWidget::Slot_SetCurrentIndex, ((QTabWidget*)(ptr)), &QTabWidget::setCurrentIndex, Qt::QueuedConnection);
}

void QTabWidget_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTabWidget*)(ptr)), &MyQTabWidget::Slot_SetCurrentIndex, ((QTabWidget*)(ptr)), &QTabWidget::setCurrentIndex);
}

void QTabWidget_SetCurrentIndex_Int(QtObjectPtr ptr, int index)
{
	((MyQTabWidget*)(ptr))->Slot_SetCurrentIndex(index);
}

//Signals
void QTabWidget_ConnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTabWidget*)(ptr)), &QTabWidget::currentChanged, ((MyQTabWidget*)(ptr)), &MyQTabWidget::Signal_CurrentChanged, Qt::QueuedConnection);
}

void QTabWidget_DisconnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTabWidget*)(ptr)), &QTabWidget::currentChanged, ((MyQTabWidget*)(ptr)), &MyQTabWidget::Signal_CurrentChanged);
}

void QTabWidget_ConnectSignalTabBarClicked(QtObjectPtr ptr)
{
	QObject::connect(((QTabWidget*)(ptr)), &QTabWidget::tabBarClicked, ((MyQTabWidget*)(ptr)), &MyQTabWidget::Signal_TabBarClicked, Qt::QueuedConnection);
}

void QTabWidget_DisconnectSignalTabBarClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QTabWidget*)(ptr)), &QTabWidget::tabBarClicked, ((MyQTabWidget*)(ptr)), &MyQTabWidget::Signal_TabBarClicked);
}

void QTabWidget_ConnectSignalTabBarDoubleClicked(QtObjectPtr ptr)
{
	QObject::connect(((QTabWidget*)(ptr)), &QTabWidget::tabBarDoubleClicked, ((MyQTabWidget*)(ptr)), &MyQTabWidget::Signal_TabBarDoubleClicked, Qt::QueuedConnection);
}

void QTabWidget_DisconnectSignalTabBarDoubleClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QTabWidget*)(ptr)), &QTabWidget::tabBarDoubleClicked, ((MyQTabWidget*)(ptr)), &MyQTabWidget::Signal_TabBarDoubleClicked);
}

void QTabWidget_ConnectSignalTabCloseRequested(QtObjectPtr ptr)
{
	QObject::connect(((QTabWidget*)(ptr)), &QTabWidget::tabCloseRequested, ((MyQTabWidget*)(ptr)), &MyQTabWidget::Signal_TabCloseRequested, Qt::QueuedConnection);
}

void QTabWidget_DisconnectSignalTabCloseRequested(QtObjectPtr ptr)
{
	QObject::disconnect(((QTabWidget*)(ptr)), &QTabWidget::tabCloseRequested, ((MyQTabWidget*)(ptr)), &MyQTabWidget::Signal_TabCloseRequested);
}

