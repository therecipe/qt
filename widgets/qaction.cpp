#include "qaction.h"
#include <QModelIndex>
#include <QString>
#include <QUrl>
#include <QMenu>
#include <QFont>
#include <QMetaObject>
#include <QWidget>
#include <QActionGroup>
#include <QObject>
#include <QVariant>
#include <QIcon>
#include <QKeySequence>
#include <QAction>
#include "_cgo_export.h"

class MyQAction: public QAction {
public:
void Signal_Changed(){callbackQActionChanged(this->objectName().toUtf8().data());};
void Signal_Hovered(){callbackQActionHovered(this->objectName().toUtf8().data());};
void Signal_Toggled(bool checked){callbackQActionToggled(this->objectName().toUtf8().data(), checked);};
void Signal_Triggered(bool checked){callbackQActionTriggered(this->objectName().toUtf8().data(), checked);};
};

int QAction_AutoRepeat(void* ptr){
	return static_cast<QAction*>(ptr)->autoRepeat();
}

char* QAction_IconText(void* ptr){
	return static_cast<QAction*>(ptr)->iconText().toUtf8().data();
}

int QAction_IsCheckable(void* ptr){
	return static_cast<QAction*>(ptr)->isCheckable();
}

int QAction_IsChecked(void* ptr){
	return static_cast<QAction*>(ptr)->isChecked();
}

int QAction_IsEnabled(void* ptr){
	return static_cast<QAction*>(ptr)->isEnabled();
}

int QAction_IsIconVisibleInMenu(void* ptr){
	return static_cast<QAction*>(ptr)->isIconVisibleInMenu();
}

int QAction_IsVisible(void* ptr){
	return static_cast<QAction*>(ptr)->isVisible();
}

int QAction_MenuRole(void* ptr){
	return static_cast<QAction*>(ptr)->menuRole();
}

int QAction_Priority(void* ptr){
	return static_cast<QAction*>(ptr)->priority();
}

void QAction_SetAutoRepeat(void* ptr, int v){
	static_cast<QAction*>(ptr)->setAutoRepeat(v != 0);
}

void QAction_SetCheckable(void* ptr, int v){
	static_cast<QAction*>(ptr)->setCheckable(v != 0);
}

void QAction_SetChecked(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setChecked", Q_ARG(bool, v != 0));
}

void QAction_SetData(void* ptr, void* userData){
	static_cast<QAction*>(ptr)->setData(*static_cast<QVariant*>(userData));
}

void QAction_SetEnabled(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setEnabled", Q_ARG(bool, v != 0));
}

void QAction_SetFont(void* ptr, void* font){
	static_cast<QAction*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QAction_SetIcon(void* ptr, void* icon){
	static_cast<QAction*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QAction_SetIconText(void* ptr, char* text){
	static_cast<QAction*>(ptr)->setIconText(QString(text));
}

void QAction_SetIconVisibleInMenu(void* ptr, int visible){
	static_cast<QAction*>(ptr)->setIconVisibleInMenu(visible != 0);
}

void QAction_SetMenuRole(void* ptr, int menuRole){
	static_cast<QAction*>(ptr)->setMenuRole(static_cast<QAction::MenuRole>(menuRole));
}

void QAction_SetPriority(void* ptr, int priority){
	static_cast<QAction*>(ptr)->setPriority(static_cast<QAction::Priority>(priority));
}

void QAction_SetShortcut(void* ptr, void* shortcut){
	static_cast<QAction*>(ptr)->setShortcut(*static_cast<QKeySequence*>(shortcut));
}

void QAction_SetShortcutContext(void* ptr, int context){
	static_cast<QAction*>(ptr)->setShortcutContext(static_cast<Qt::ShortcutContext>(context));
}

void QAction_SetStatusTip(void* ptr, char* statusTip){
	static_cast<QAction*>(ptr)->setStatusTip(QString(statusTip));
}

void QAction_SetText(void* ptr, char* text){
	static_cast<QAction*>(ptr)->setText(QString(text));
}

void QAction_SetToolTip(void* ptr, char* tip){
	static_cast<QAction*>(ptr)->setToolTip(QString(tip));
}

void QAction_SetVisible(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setVisible", Q_ARG(bool, v != 0));
}

void QAction_SetWhatsThis(void* ptr, char* what){
	static_cast<QAction*>(ptr)->setWhatsThis(QString(what));
}

int QAction_ShortcutContext(void* ptr){
	return static_cast<QAction*>(ptr)->shortcutContext();
}

char* QAction_StatusTip(void* ptr){
	return static_cast<QAction*>(ptr)->statusTip().toUtf8().data();
}

char* QAction_Text(void* ptr){
	return static_cast<QAction*>(ptr)->text().toUtf8().data();
}

void QAction_Toggle(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "toggle");
}

char* QAction_ToolTip(void* ptr){
	return static_cast<QAction*>(ptr)->toolTip().toUtf8().data();
}

char* QAction_WhatsThis(void* ptr){
	return static_cast<QAction*>(ptr)->whatsThis().toUtf8().data();
}

void* QAction_NewQAction(void* parent){
	return new QAction(static_cast<QObject*>(parent));
}

void* QAction_NewQAction3(void* icon, char* text, void* parent){
	return new QAction(*static_cast<QIcon*>(icon), QString(text), static_cast<QObject*>(parent));
}

void* QAction_NewQAction2(char* text, void* parent){
	return new QAction(QString(text), static_cast<QObject*>(parent));
}

void* QAction_ActionGroup(void* ptr){
	return static_cast<QAction*>(ptr)->actionGroup();
}

void QAction_Activate(void* ptr, int event){
	static_cast<QAction*>(ptr)->activate(static_cast<QAction::ActionEvent>(event));
}

void QAction_ConnectChanged(void* ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::changed), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Changed));;
}

void QAction_DisconnectChanged(void* ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::changed), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Changed));;
}

void* QAction_Data(void* ptr){
	return new QVariant(static_cast<QAction*>(ptr)->data());
}

void QAction_Hover(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "hover");
}

void QAction_ConnectHovered(void* ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::hovered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Hovered));;
}

void QAction_DisconnectHovered(void* ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::hovered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Hovered));;
}

int QAction_IsSeparator(void* ptr){
	return static_cast<QAction*>(ptr)->isSeparator();
}

void* QAction_Menu(void* ptr){
	return static_cast<QAction*>(ptr)->menu();
}

void* QAction_ParentWidget(void* ptr){
	return static_cast<QAction*>(ptr)->parentWidget();
}

void QAction_SetActionGroup(void* ptr, void* group){
	static_cast<QAction*>(ptr)->setActionGroup(static_cast<QActionGroup*>(group));
}

void QAction_SetDisabled(void* ptr, int b){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setDisabled", Q_ARG(bool, b != 0));
}

void QAction_SetMenu(void* ptr, void* menu){
	static_cast<QAction*>(ptr)->setMenu(static_cast<QMenu*>(menu));
}

void QAction_SetSeparator(void* ptr, int b){
	static_cast<QAction*>(ptr)->setSeparator(b != 0);
}

void QAction_SetShortcuts2(void* ptr, int key){
	static_cast<QAction*>(ptr)->setShortcuts(static_cast<QKeySequence::StandardKey>(key));
}

int QAction_ShowStatusText(void* ptr, void* widget){
	return static_cast<QAction*>(ptr)->showStatusText(static_cast<QWidget*>(widget));
}

void QAction_ConnectToggled(void* ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::toggled), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Toggled));;
}

void QAction_DisconnectToggled(void* ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::toggled), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Toggled));;
}

void QAction_Trigger(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "trigger");
}

void QAction_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::triggered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Triggered));;
}

void QAction_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::triggered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Triggered));;
}

void QAction_DestroyQAction(void* ptr){
	static_cast<QAction*>(ptr)->~QAction();
}

