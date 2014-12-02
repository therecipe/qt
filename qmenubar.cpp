#include "qmenubar.h"
#include <QMenuBar>
#include "cgoexport.h"

//MyQMenuBar
class MyQMenuBar: public QMenuBar {
Q_OBJECT
public:
void Signal_Hovered() { callbackQMenuBar(this, QString("hovered").toUtf8().data()); };
void Signal_Triggered() { callbackQMenuBar(this, QString("triggered").toUtf8().data()); };

signals:
void Slot_SetVisible(bool visible);

};
#include "qmenubar.moc"


//Public Functions
QtObjectPtr QMenuBar_New_QWidget(QtObjectPtr parent)
{
	return new QMenuBar(((QWidget*)(parent)));
}

void QMenuBar_Destroy(QtObjectPtr ptr)
{
	((QMenuBar*)(ptr))->~QMenuBar();
}

QtObjectPtr QMenuBar_ActiveAction(QtObjectPtr ptr)
{
	return ((QMenuBar*)(ptr))->activeAction();
}

QtObjectPtr QMenuBar_AddAction_String(QtObjectPtr ptr, char* text)
{
	return ((QMenuBar*)(ptr))->addAction(QString(text));
}

QtObjectPtr QMenuBar_AddAction_String_QObject_String(QtObjectPtr ptr, char* text, QtObjectPtr receiver, char* member)
{
	return ((QMenuBar*)(ptr))->addAction(QString(text), ((QObject*)(receiver)), member);
}

void QMenuBar_AddAction_QAction(QtObjectPtr ptr, QtObjectPtr action)
{
	((QMenuBar*)(ptr))->addAction(((QAction*)(action)));
}

QtObjectPtr QMenuBar_AddMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu)
{
	return ((QMenuBar*)(ptr))->addMenu(((QMenu*)(menu)));
}

QtObjectPtr QMenuBar_AddMenu_String(QtObjectPtr ptr, char* title)
{
	return ((QMenuBar*)(ptr))->addMenu(QString(title));
}

QtObjectPtr QMenuBar_AddSeparator(QtObjectPtr ptr)
{
	return ((QMenuBar*)(ptr))->addSeparator();
}

void QMenuBar_Clear(QtObjectPtr ptr)
{
	((QMenuBar*)(ptr))->clear();
}

QtObjectPtr QMenuBar_CornerWidget_Corner(QtObjectPtr ptr, int corner)
{
	return ((QMenuBar*)(ptr))->cornerWidget(((Qt::Corner)(corner)));
}

QtObjectPtr QMenuBar_InsertMenu_QAction_QMenu(QtObjectPtr ptr, QtObjectPtr before, QtObjectPtr menu)
{
	return ((QMenuBar*)(ptr))->insertMenu(((QAction*)(before)), ((QMenu*)(menu)));
}

QtObjectPtr QMenuBar_InsertSeparator_QAction(QtObjectPtr ptr, QtObjectPtr before)
{
	return ((QMenuBar*)(ptr))->insertSeparator(((QAction*)(before)));
}

int QMenuBar_IsDefaultUp(QtObjectPtr ptr)
{
	return ((QMenuBar*)(ptr))->isDefaultUp();
}

int QMenuBar_IsNativeMenuBar(QtObjectPtr ptr)
{
	return ((QMenuBar*)(ptr))->isNativeMenuBar();
}

void QMenuBar_SetActiveAction_QAction(QtObjectPtr ptr, QtObjectPtr act)
{
	((QMenuBar*)(ptr))->setActiveAction(((QAction*)(act)));
}

void QMenuBar_SetCornerWidget_QWidget_Corner(QtObjectPtr ptr, QtObjectPtr widget, int corner)
{
	((QMenuBar*)(ptr))->setCornerWidget(((QWidget*)(widget)), ((Qt::Corner)(corner)));
}

void QMenuBar_SetDefaultUp_Bool(QtObjectPtr ptr, int defaultAction)
{
	((QMenuBar*)(ptr))->setDefaultUp(defaultAction != 0);
}

void QMenuBar_SetNativeMenuBar_Bool(QtObjectPtr ptr, int nativeMenuBar)
{
	((QMenuBar*)(ptr))->setNativeMenuBar(nativeMenuBar != 0);
}

//Public Slots
void QMenuBar_ConnectSlotSetVisible(QtObjectPtr ptr)
{
	QObject::connect(((MyQMenuBar*)(ptr)), &MyQMenuBar::Slot_SetVisible, ((QMenuBar*)(ptr)), &QMenuBar::setVisible, Qt::QueuedConnection);
}

void QMenuBar_DisconnectSlotSetVisible(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQMenuBar*)(ptr)), &MyQMenuBar::Slot_SetVisible, ((QMenuBar*)(ptr)), &QMenuBar::setVisible);
}

void QMenuBar_SetVisible_Bool(QtObjectPtr ptr, int visible)
{
	((MyQMenuBar*)(ptr))->Slot_SetVisible(visible != 0);
}

//Signals
void QMenuBar_ConnectSignalHovered(QtObjectPtr ptr)
{
	QObject::connect(((QMenuBar*)(ptr)), &QMenuBar::hovered, ((MyQMenuBar*)(ptr)), &MyQMenuBar::Signal_Hovered, Qt::QueuedConnection);
}

void QMenuBar_DisconnectSignalHovered(QtObjectPtr ptr)
{
	QObject::disconnect(((QMenuBar*)(ptr)), &QMenuBar::hovered, ((MyQMenuBar*)(ptr)), &MyQMenuBar::Signal_Hovered);
}

void QMenuBar_ConnectSignalTriggered(QtObjectPtr ptr)
{
	QObject::connect(((QMenuBar*)(ptr)), &QMenuBar::triggered, ((MyQMenuBar*)(ptr)), &MyQMenuBar::Signal_Triggered, Qt::QueuedConnection);
}

void QMenuBar_DisconnectSignalTriggered(QtObjectPtr ptr)
{
	QObject::disconnect(((QMenuBar*)(ptr)), &QMenuBar::triggered, ((MyQMenuBar*)(ptr)), &MyQMenuBar::Signal_Triggered);
}

