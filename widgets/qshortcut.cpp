#include "qshortcut.h"
#include <QModelIndex>
#include <QWidget>
#include <QObject>
#include <QKeySequence>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QShortcut>
#include "_cgo_export.h"

class MyQShortcut: public QShortcut {
public:
void Signal_Activated(){callbackQShortcutActivated(this->objectName().toUtf8().data());};
void Signal_ActivatedAmbiguously(){callbackQShortcutActivatedAmbiguously(this->objectName().toUtf8().data());};
};

int QShortcut_AutoRepeat(void* ptr){
	return static_cast<QShortcut*>(ptr)->autoRepeat();
}

int QShortcut_Context(void* ptr){
	return static_cast<QShortcut*>(ptr)->context();
}

int QShortcut_IsEnabled(void* ptr){
	return static_cast<QShortcut*>(ptr)->isEnabled();
}

void QShortcut_SetAutoRepeat(void* ptr, int on){
	static_cast<QShortcut*>(ptr)->setAutoRepeat(on != 0);
}

void QShortcut_SetContext(void* ptr, int context){
	static_cast<QShortcut*>(ptr)->setContext(static_cast<Qt::ShortcutContext>(context));
}

void QShortcut_SetEnabled(void* ptr, int enable){
	static_cast<QShortcut*>(ptr)->setEnabled(enable != 0);
}

void QShortcut_SetKey(void* ptr, void* key){
	static_cast<QShortcut*>(ptr)->setKey(*static_cast<QKeySequence*>(key));
}

void QShortcut_SetWhatsThis(void* ptr, char* text){
	static_cast<QShortcut*>(ptr)->setWhatsThis(QString(text));
}

char* QShortcut_WhatsThis(void* ptr){
	return static_cast<QShortcut*>(ptr)->whatsThis().toUtf8().data();
}

void* QShortcut_NewQShortcut(void* parent){
	return new QShortcut(static_cast<QWidget*>(parent));
}

void* QShortcut_NewQShortcut2(void* key, void* parent, char* member, char* ambiguousMember, int context){
	return new QShortcut(*static_cast<QKeySequence*>(key), static_cast<QWidget*>(parent), const_cast<const char*>(member), const_cast<const char*>(ambiguousMember), static_cast<Qt::ShortcutContext>(context));
}

void QShortcut_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activated), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_Activated));;
}

void QShortcut_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activated), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_Activated));;
}

void QShortcut_ConnectActivatedAmbiguously(void* ptr){
	QObject::connect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activatedAmbiguously), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_ActivatedAmbiguously));;
}

void QShortcut_DisconnectActivatedAmbiguously(void* ptr){
	QObject::disconnect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activatedAmbiguously), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_ActivatedAmbiguously));;
}

int QShortcut_Id(void* ptr){
	return static_cast<QShortcut*>(ptr)->id();
}

void* QShortcut_ParentWidget(void* ptr){
	return static_cast<QShortcut*>(ptr)->parentWidget();
}

void QShortcut_DestroyQShortcut(void* ptr){
	static_cast<QShortcut*>(ptr)->~QShortcut();
}

