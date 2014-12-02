#include "qaction.h"
#include <QAction>
#include "cgoexport.h"

//MyQAction
class MyQAction: public QAction {
Q_OBJECT
public:
void Signal_Changed() { callbackQAction(this, QString("changed").toUtf8().data()); };
void Signal_Hovered() { callbackQAction(this, QString("hovered").toUtf8().data()); };
void Signal_Toggled() { callbackQAction(this, QString("toggled").toUtf8().data()); };
void Signal_Triggered() { callbackQAction(this, QString("triggered").toUtf8().data()); };

signals:
void Slot_Hover();
void Slot_SetChecked(bool bo);
void Slot_SetDisabled(bool bo);
void Slot_SetEnabled(bool bo);
void Slot_SetVisible(bool bo);
void Slot_Toggle();
void Slot_Trigger();

};
#include "qaction.moc"


//Public Functions
QtObjectPtr QAction_New_QObject(QtObjectPtr parent)
{
	return new QAction(((QObject*)(parent)));
}

QtObjectPtr QAction_New_String_QObject(char* text, QtObjectPtr parent)
{
	return new QAction(QString(text), ((QObject*)(parent)));
}

void QAction_Destroy(QtObjectPtr ptr)
{
	((QAction*)(ptr))->~QAction();
}

int QAction_AutoRepeat(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->autoRepeat();
}

char* QAction_IconText(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->iconText().toUtf8().data();
}

int QAction_IsCheckable(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->isCheckable();
}

int QAction_IsChecked(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->isChecked();
}

int QAction_IsEnabled(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->isEnabled();
}

int QAction_IsIconVisibleInMenu(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->isIconVisibleInMenu();
}

int QAction_IsSeparator(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->isSeparator();
}

int QAction_IsVisible(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->isVisible();
}

QtObjectPtr QAction_Menu(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->menu();
}

QtObjectPtr QAction_ParentWidget(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->parentWidget();
}

void QAction_SetAutoRepeat_Bool(QtObjectPtr ptr, int autoRepeat)
{
	((QAction*)(ptr))->setAutoRepeat(autoRepeat != 0);
}

void QAction_SetCheckable_Bool(QtObjectPtr ptr, int checkable)
{
	((QAction*)(ptr))->setCheckable(checkable != 0);
}

void QAction_SetIconText_String(QtObjectPtr ptr, char* text)
{
	((QAction*)(ptr))->setIconText(QString(text));
}

void QAction_SetIconVisibleInMenu_Bool(QtObjectPtr ptr, int visible)
{
	((QAction*)(ptr))->setIconVisibleInMenu(visible != 0);
}

void QAction_SetMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu)
{
	((QAction*)(ptr))->setMenu(((QMenu*)(menu)));
}

void QAction_SetSeparator_Bool(QtObjectPtr ptr, int bo)
{
	((QAction*)(ptr))->setSeparator(bo != 0);
}

void QAction_SetShortcutContext_ShortcutContext(QtObjectPtr ptr, int context)
{
	((QAction*)(ptr))->setShortcutContext(((Qt::ShortcutContext)(context)));
}

void QAction_SetStatusTip_String(QtObjectPtr ptr, char* statusTip)
{
	((QAction*)(ptr))->setStatusTip(QString(statusTip));
}

void QAction_SetText_String(QtObjectPtr ptr, char* text)
{
	((QAction*)(ptr))->setText(QString(text));
}

void QAction_SetToolTip_String(QtObjectPtr ptr, char* tip)
{
	((QAction*)(ptr))->setToolTip(QString(tip));
}

void QAction_SetWhatsThis_String(QtObjectPtr ptr, char* what)
{
	((QAction*)(ptr))->setWhatsThis(QString(what));
}

int QAction_ShortcutContext(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->shortcutContext();
}

int QAction_ShowStatusText_QWidget(QtObjectPtr ptr, QtObjectPtr widget)
{
	return ((QAction*)(ptr))->showStatusText(((QWidget*)(widget)));
}

char* QAction_StatusTip(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->statusTip().toUtf8().data();
}

char* QAction_Text(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->text().toUtf8().data();
}

char* QAction_ToolTip(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->toolTip().toUtf8().data();
}

char* QAction_WhatsThis(QtObjectPtr ptr)
{
	return ((QAction*)(ptr))->whatsThis().toUtf8().data();
}

//Public Slots
void QAction_ConnectSlotHover(QtObjectPtr ptr)
{
	QObject::connect(((MyQAction*)(ptr)), &MyQAction::Slot_Hover, ((QAction*)(ptr)), &QAction::hover, Qt::QueuedConnection);
}

void QAction_DisconnectSlotHover(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAction*)(ptr)), &MyQAction::Slot_Hover, ((QAction*)(ptr)), &QAction::hover);
}

void QAction_Hover(QtObjectPtr ptr)
{
	((MyQAction*)(ptr))->Slot_Hover();
}

void QAction_ConnectSlotSetChecked(QtObjectPtr ptr)
{
	QObject::connect(((MyQAction*)(ptr)), &MyQAction::Slot_SetChecked, ((QAction*)(ptr)), &QAction::setChecked, Qt::QueuedConnection);
}

void QAction_DisconnectSlotSetChecked(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAction*)(ptr)), &MyQAction::Slot_SetChecked, ((QAction*)(ptr)), &QAction::setChecked);
}

void QAction_SetChecked_Bool(QtObjectPtr ptr, int bo)
{
	((MyQAction*)(ptr))->Slot_SetChecked(bo != 0);
}

void QAction_ConnectSlotSetDisabled(QtObjectPtr ptr)
{
	QObject::connect(((MyQAction*)(ptr)), &MyQAction::Slot_SetDisabled, ((QAction*)(ptr)), &QAction::setDisabled, Qt::QueuedConnection);
}

void QAction_DisconnectSlotSetDisabled(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAction*)(ptr)), &MyQAction::Slot_SetDisabled, ((QAction*)(ptr)), &QAction::setDisabled);
}

void QAction_SetDisabled_Bool(QtObjectPtr ptr, int bo)
{
	((MyQAction*)(ptr))->Slot_SetDisabled(bo != 0);
}

void QAction_ConnectSlotSetEnabled(QtObjectPtr ptr)
{
	QObject::connect(((MyQAction*)(ptr)), &MyQAction::Slot_SetEnabled, ((QAction*)(ptr)), &QAction::setEnabled, Qt::QueuedConnection);
}

void QAction_DisconnectSlotSetEnabled(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAction*)(ptr)), &MyQAction::Slot_SetEnabled, ((QAction*)(ptr)), &QAction::setEnabled);
}

void QAction_SetEnabled_Bool(QtObjectPtr ptr, int bo)
{
	((MyQAction*)(ptr))->Slot_SetEnabled(bo != 0);
}

void QAction_ConnectSlotSetVisible(QtObjectPtr ptr)
{
	QObject::connect(((MyQAction*)(ptr)), &MyQAction::Slot_SetVisible, ((QAction*)(ptr)), &QAction::setVisible, Qt::QueuedConnection);
}

void QAction_DisconnectSlotSetVisible(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAction*)(ptr)), &MyQAction::Slot_SetVisible, ((QAction*)(ptr)), &QAction::setVisible);
}

void QAction_SetVisible_Bool(QtObjectPtr ptr, int bo)
{
	((MyQAction*)(ptr))->Slot_SetVisible(bo != 0);
}

void QAction_ConnectSlotToggle(QtObjectPtr ptr)
{
	QObject::connect(((MyQAction*)(ptr)), &MyQAction::Slot_Toggle, ((QAction*)(ptr)), &QAction::toggle, Qt::QueuedConnection);
}

void QAction_DisconnectSlotToggle(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAction*)(ptr)), &MyQAction::Slot_Toggle, ((QAction*)(ptr)), &QAction::toggle);
}

void QAction_Toggle(QtObjectPtr ptr)
{
	((MyQAction*)(ptr))->Slot_Toggle();
}

void QAction_ConnectSlotTrigger(QtObjectPtr ptr)
{
	QObject::connect(((MyQAction*)(ptr)), &MyQAction::Slot_Trigger, ((QAction*)(ptr)), &QAction::trigger, Qt::QueuedConnection);
}

void QAction_DisconnectSlotTrigger(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAction*)(ptr)), &MyQAction::Slot_Trigger, ((QAction*)(ptr)), &QAction::trigger);
}

void QAction_Trigger(QtObjectPtr ptr)
{
	((MyQAction*)(ptr))->Slot_Trigger();
}

//Signals
void QAction_ConnectSignalChanged(QtObjectPtr ptr)
{
	QObject::connect(((QAction*)(ptr)), &QAction::changed, ((MyQAction*)(ptr)), &MyQAction::Signal_Changed, Qt::QueuedConnection);
}

void QAction_DisconnectSignalChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QAction*)(ptr)), &QAction::changed, ((MyQAction*)(ptr)), &MyQAction::Signal_Changed);
}

void QAction_ConnectSignalHovered(QtObjectPtr ptr)
{
	QObject::connect(((QAction*)(ptr)), &QAction::hovered, ((MyQAction*)(ptr)), &MyQAction::Signal_Hovered, Qt::QueuedConnection);
}

void QAction_DisconnectSignalHovered(QtObjectPtr ptr)
{
	QObject::disconnect(((QAction*)(ptr)), &QAction::hovered, ((MyQAction*)(ptr)), &MyQAction::Signal_Hovered);
}

void QAction_ConnectSignalToggled(QtObjectPtr ptr)
{
	QObject::connect(((QAction*)(ptr)), &QAction::toggled, ((MyQAction*)(ptr)), &MyQAction::Signal_Toggled, Qt::QueuedConnection);
}

void QAction_DisconnectSignalToggled(QtObjectPtr ptr)
{
	QObject::disconnect(((QAction*)(ptr)), &QAction::toggled, ((MyQAction*)(ptr)), &MyQAction::Signal_Toggled);
}

void QAction_ConnectSignalTriggered(QtObjectPtr ptr)
{
	QObject::connect(((QAction*)(ptr)), &QAction::triggered, ((MyQAction*)(ptr)), &MyQAction::Signal_Triggered, Qt::QueuedConnection);
}

void QAction_DisconnectSignalTriggered(QtObjectPtr ptr)
{
	QObject::disconnect(((QAction*)(ptr)), &QAction::triggered, ((MyQAction*)(ptr)), &MyQAction::Signal_Triggered);
}

