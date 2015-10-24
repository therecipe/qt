#include "qaction.h"
#include <QModelIndex>
#include <QWidget>
#include <QFont>
#include <QMenu>
#include <QActionGroup>
#include <QVariant>
#include <QUrl>
#include <QObject>
#include <QIcon>
#include <QKeySequence>
#include <QString>
#include <QMetaObject>
#include <QAction>
#include "_cgo_export.h"

class MyQAction: public QAction {
public:
void Signal_Changed(){callbackQActionChanged(this->objectName().toUtf8().data());};
void Signal_Hovered(){callbackQActionHovered(this->objectName().toUtf8().data());};
void Signal_Toggled(bool checked){callbackQActionToggled(this->objectName().toUtf8().data(), checked);};
void Signal_Triggered(bool checked){callbackQActionTriggered(this->objectName().toUtf8().data(), checked);};
};

int QAction_AutoRepeat(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->autoRepeat();
}

char* QAction_IconText(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->iconText().toUtf8().data();
}

int QAction_IsCheckable(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->isCheckable();
}

int QAction_IsChecked(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->isChecked();
}

int QAction_IsEnabled(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->isEnabled();
}

int QAction_IsIconVisibleInMenu(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->isIconVisibleInMenu();
}

int QAction_IsVisible(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->isVisible();
}

int QAction_MenuRole(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->menuRole();
}

int QAction_Priority(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->priority();
}

void QAction_SetAutoRepeat(QtObjectPtr ptr, int v){
	static_cast<QAction*>(ptr)->setAutoRepeat(v != 0);
}

void QAction_SetCheckable(QtObjectPtr ptr, int v){
	static_cast<QAction*>(ptr)->setCheckable(v != 0);
}

void QAction_SetChecked(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setChecked", Q_ARG(bool, v != 0));
}

void QAction_SetData(QtObjectPtr ptr, char* userData){
	static_cast<QAction*>(ptr)->setData(QVariant(userData));
}

void QAction_SetEnabled(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setEnabled", Q_ARG(bool, v != 0));
}

void QAction_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QAction*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QAction_SetIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QAction*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QAction_SetIconText(QtObjectPtr ptr, char* text){
	static_cast<QAction*>(ptr)->setIconText(QString(text));
}

void QAction_SetIconVisibleInMenu(QtObjectPtr ptr, int visible){
	static_cast<QAction*>(ptr)->setIconVisibleInMenu(visible != 0);
}

void QAction_SetMenuRole(QtObjectPtr ptr, int menuRole){
	static_cast<QAction*>(ptr)->setMenuRole(static_cast<QAction::MenuRole>(menuRole));
}

void QAction_SetPriority(QtObjectPtr ptr, int priority){
	static_cast<QAction*>(ptr)->setPriority(static_cast<QAction::Priority>(priority));
}

void QAction_SetShortcut(QtObjectPtr ptr, QtObjectPtr shortcut){
	static_cast<QAction*>(ptr)->setShortcut(*static_cast<QKeySequence*>(shortcut));
}

void QAction_SetShortcutContext(QtObjectPtr ptr, int context){
	static_cast<QAction*>(ptr)->setShortcutContext(static_cast<Qt::ShortcutContext>(context));
}

void QAction_SetStatusTip(QtObjectPtr ptr, char* statusTip){
	static_cast<QAction*>(ptr)->setStatusTip(QString(statusTip));
}

void QAction_SetText(QtObjectPtr ptr, char* text){
	static_cast<QAction*>(ptr)->setText(QString(text));
}

void QAction_SetToolTip(QtObjectPtr ptr, char* tip){
	static_cast<QAction*>(ptr)->setToolTip(QString(tip));
}

void QAction_SetVisible(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setVisible", Q_ARG(bool, v != 0));
}

void QAction_SetWhatsThis(QtObjectPtr ptr, char* what){
	static_cast<QAction*>(ptr)->setWhatsThis(QString(what));
}

int QAction_ShortcutContext(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->shortcutContext();
}

char* QAction_StatusTip(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->statusTip().toUtf8().data();
}

char* QAction_Text(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->text().toUtf8().data();
}

void QAction_Toggle(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "toggle");
}

char* QAction_ToolTip(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->toolTip().toUtf8().data();
}

char* QAction_WhatsThis(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->whatsThis().toUtf8().data();
}

QtObjectPtr QAction_NewQAction(QtObjectPtr parent){
	return new QAction(static_cast<QObject*>(parent));
}

QtObjectPtr QAction_NewQAction3(QtObjectPtr icon, char* text, QtObjectPtr parent){
	return new QAction(*static_cast<QIcon*>(icon), QString(text), static_cast<QObject*>(parent));
}

QtObjectPtr QAction_NewQAction2(char* text, QtObjectPtr parent){
	return new QAction(QString(text), static_cast<QObject*>(parent));
}

QtObjectPtr QAction_ActionGroup(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->actionGroup();
}

void QAction_Activate(QtObjectPtr ptr, int event){
	static_cast<QAction*>(ptr)->activate(static_cast<QAction::ActionEvent>(event));
}

void QAction_ConnectChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::changed), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Changed));;
}

void QAction_DisconnectChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::changed), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Changed));;
}

char* QAction_Data(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->data().toString().toUtf8().data();
}

void QAction_Hover(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "hover");
}

void QAction_ConnectHovered(QtObjectPtr ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::hovered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Hovered));;
}

void QAction_DisconnectHovered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::hovered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Hovered));;
}

int QAction_IsSeparator(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->isSeparator();
}

QtObjectPtr QAction_Menu(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->menu();
}

QtObjectPtr QAction_ParentWidget(QtObjectPtr ptr){
	return static_cast<QAction*>(ptr)->parentWidget();
}

void QAction_SetActionGroup(QtObjectPtr ptr, QtObjectPtr group){
	static_cast<QAction*>(ptr)->setActionGroup(static_cast<QActionGroup*>(group));
}

void QAction_SetDisabled(QtObjectPtr ptr, int b){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setDisabled", Q_ARG(bool, b != 0));
}

void QAction_SetMenu(QtObjectPtr ptr, QtObjectPtr menu){
	static_cast<QAction*>(ptr)->setMenu(static_cast<QMenu*>(menu));
}

void QAction_SetSeparator(QtObjectPtr ptr, int b){
	static_cast<QAction*>(ptr)->setSeparator(b != 0);
}

void QAction_SetShortcuts2(QtObjectPtr ptr, int key){
	static_cast<QAction*>(ptr)->setShortcuts(static_cast<QKeySequence::StandardKey>(key));
}

int QAction_ShowStatusText(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QAction*>(ptr)->showStatusText(static_cast<QWidget*>(widget));
}

void QAction_ConnectToggled(QtObjectPtr ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::toggled), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Toggled));;
}

void QAction_DisconnectToggled(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::toggled), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Toggled));;
}

void QAction_Trigger(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "trigger");
}

void QAction_ConnectTriggered(QtObjectPtr ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::triggered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Triggered));;
}

void QAction_DisconnectTriggered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::triggered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Triggered));;
}

void QAction_DestroyQAction(QtObjectPtr ptr){
	static_cast<QAction*>(ptr)->~QAction();
}

