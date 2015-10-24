#include "qsystemtrayicon.h"
#include <QMetaObject>
#include <QObject>
#include <QIcon>
#include <QMenu>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSystemTrayIcon>
#include "_cgo_export.h"

class MyQSystemTrayIcon: public QSystemTrayIcon {
public:
void Signal_Activated(QSystemTrayIcon::ActivationReason reason){callbackQSystemTrayIconActivated(this->objectName().toUtf8().data(), reason);};
void Signal_MessageClicked(){callbackQSystemTrayIconMessageClicked(this->objectName().toUtf8().data());};
};

int QSystemTrayIcon_IsVisible(QtObjectPtr ptr){
	return static_cast<QSystemTrayIcon*>(ptr)->isVisible();
}

void QSystemTrayIcon_SetIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QSystemTrayIcon*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QSystemTrayIcon_SetToolTip(QtObjectPtr ptr, char* tip){
	static_cast<QSystemTrayIcon*>(ptr)->setToolTip(QString(tip));
}

void QSystemTrayIcon_SetVisible(QtObjectPtr ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QSystemTrayIcon*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QSystemTrayIcon_ShowMessage(QtObjectPtr ptr, char* title, char* message, int icon, int millisecondsTimeoutHint){
	QMetaObject::invokeMethod(static_cast<QSystemTrayIcon*>(ptr), "showMessage", Q_ARG(QString, QString(title)), Q_ARG(QString, QString(message)), Q_ARG(QSystemTrayIcon::MessageIcon, static_cast<QSystemTrayIcon::MessageIcon>(icon)), Q_ARG(int, millisecondsTimeoutHint));
}

char* QSystemTrayIcon_ToolTip(QtObjectPtr ptr){
	return static_cast<QSystemTrayIcon*>(ptr)->toolTip().toUtf8().data();
}

QtObjectPtr QSystemTrayIcon_NewQSystemTrayIcon(QtObjectPtr parent){
	return new QSystemTrayIcon(static_cast<QObject*>(parent));
}

QtObjectPtr QSystemTrayIcon_NewQSystemTrayIcon2(QtObjectPtr icon, QtObjectPtr parent){
	return new QSystemTrayIcon(*static_cast<QIcon*>(icon), static_cast<QObject*>(parent));
}

void QSystemTrayIcon_ConnectActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QSystemTrayIcon*>(ptr), static_cast<void (QSystemTrayIcon::*)(QSystemTrayIcon::ActivationReason)>(&QSystemTrayIcon::activated), static_cast<MyQSystemTrayIcon*>(ptr), static_cast<void (MyQSystemTrayIcon::*)(QSystemTrayIcon::ActivationReason)>(&MyQSystemTrayIcon::Signal_Activated));;
}

void QSystemTrayIcon_DisconnectActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSystemTrayIcon*>(ptr), static_cast<void (QSystemTrayIcon::*)(QSystemTrayIcon::ActivationReason)>(&QSystemTrayIcon::activated), static_cast<MyQSystemTrayIcon*>(ptr), static_cast<void (MyQSystemTrayIcon::*)(QSystemTrayIcon::ActivationReason)>(&MyQSystemTrayIcon::Signal_Activated));;
}

QtObjectPtr QSystemTrayIcon_ContextMenu(QtObjectPtr ptr){
	return static_cast<QSystemTrayIcon*>(ptr)->contextMenu();
}

void QSystemTrayIcon_Hide(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSystemTrayIcon*>(ptr), "hide");
}

int QSystemTrayIcon_QSystemTrayIcon_IsSystemTrayAvailable(){
	return QSystemTrayIcon::isSystemTrayAvailable();
}

void QSystemTrayIcon_ConnectMessageClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QSystemTrayIcon*>(ptr), static_cast<void (QSystemTrayIcon::*)()>(&QSystemTrayIcon::messageClicked), static_cast<MyQSystemTrayIcon*>(ptr), static_cast<void (MyQSystemTrayIcon::*)()>(&MyQSystemTrayIcon::Signal_MessageClicked));;
}

void QSystemTrayIcon_DisconnectMessageClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSystemTrayIcon*>(ptr), static_cast<void (QSystemTrayIcon::*)()>(&QSystemTrayIcon::messageClicked), static_cast<MyQSystemTrayIcon*>(ptr), static_cast<void (MyQSystemTrayIcon::*)()>(&MyQSystemTrayIcon::Signal_MessageClicked));;
}

void QSystemTrayIcon_SetContextMenu(QtObjectPtr ptr, QtObjectPtr menu){
	static_cast<QSystemTrayIcon*>(ptr)->setContextMenu(static_cast<QMenu*>(menu));
}

void QSystemTrayIcon_Show(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSystemTrayIcon*>(ptr), "show");
}

int QSystemTrayIcon_QSystemTrayIcon_SupportsMessages(){
	return QSystemTrayIcon::supportsMessages();
}

void QSystemTrayIcon_DestroyQSystemTrayIcon(QtObjectPtr ptr){
	static_cast<QSystemTrayIcon*>(ptr)->~QSystemTrayIcon();
}

