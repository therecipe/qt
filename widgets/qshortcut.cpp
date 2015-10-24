#include "qshortcut.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QKeySequence>
#include <QString>
#include <QVariant>
#include <QShortcut>
#include "_cgo_export.h"

class MyQShortcut: public QShortcut {
public:
void Signal_Activated(){callbackQShortcutActivated(this->objectName().toUtf8().data());};
void Signal_ActivatedAmbiguously(){callbackQShortcutActivatedAmbiguously(this->objectName().toUtf8().data());};
};

int QShortcut_AutoRepeat(QtObjectPtr ptr){
	return static_cast<QShortcut*>(ptr)->autoRepeat();
}

int QShortcut_Context(QtObjectPtr ptr){
	return static_cast<QShortcut*>(ptr)->context();
}

int QShortcut_IsEnabled(QtObjectPtr ptr){
	return static_cast<QShortcut*>(ptr)->isEnabled();
}

void QShortcut_SetAutoRepeat(QtObjectPtr ptr, int on){
	static_cast<QShortcut*>(ptr)->setAutoRepeat(on != 0);
}

void QShortcut_SetContext(QtObjectPtr ptr, int context){
	static_cast<QShortcut*>(ptr)->setContext(static_cast<Qt::ShortcutContext>(context));
}

void QShortcut_SetEnabled(QtObjectPtr ptr, int enable){
	static_cast<QShortcut*>(ptr)->setEnabled(enable != 0);
}

void QShortcut_SetKey(QtObjectPtr ptr, QtObjectPtr key){
	static_cast<QShortcut*>(ptr)->setKey(*static_cast<QKeySequence*>(key));
}

void QShortcut_SetWhatsThis(QtObjectPtr ptr, char* text){
	static_cast<QShortcut*>(ptr)->setWhatsThis(QString(text));
}

char* QShortcut_WhatsThis(QtObjectPtr ptr){
	return static_cast<QShortcut*>(ptr)->whatsThis().toUtf8().data();
}

QtObjectPtr QShortcut_NewQShortcut(QtObjectPtr parent){
	return new QShortcut(static_cast<QWidget*>(parent));
}

QtObjectPtr QShortcut_NewQShortcut2(QtObjectPtr key, QtObjectPtr parent, char* member, char* ambiguousMember, int context){
	return new QShortcut(*static_cast<QKeySequence*>(key), static_cast<QWidget*>(parent), const_cast<const char*>(member), const_cast<const char*>(ambiguousMember), static_cast<Qt::ShortcutContext>(context));
}

void QShortcut_ConnectActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activated), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_Activated));;
}

void QShortcut_DisconnectActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activated), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_Activated));;
}

void QShortcut_ConnectActivatedAmbiguously(QtObjectPtr ptr){
	QObject::connect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activatedAmbiguously), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_ActivatedAmbiguously));;
}

void QShortcut_DisconnectActivatedAmbiguously(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activatedAmbiguously), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_ActivatedAmbiguously));;
}

int QShortcut_Id(QtObjectPtr ptr){
	return static_cast<QShortcut*>(ptr)->id();
}

QtObjectPtr QShortcut_ParentWidget(QtObjectPtr ptr){
	return static_cast<QShortcut*>(ptr)->parentWidget();
}

void QShortcut_DestroyQShortcut(QtObjectPtr ptr){
	static_cast<QShortcut*>(ptr)->~QShortcut();
}

