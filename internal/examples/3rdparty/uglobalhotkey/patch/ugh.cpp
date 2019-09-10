#include "ugh.h"

#include "_cgo_export.h"

#include <QString>
#include <QWidget>

#include "uglobalhotkeys.h"


class MyUGlobalHotkeys: public UGlobalHotkeys
{
public:
	MyUGlobalHotkeys(QWidget *parent) : UGlobalHotkeys(parent) {};
	void Signal_Activated(unsigned long id) { callbackUGlobalHotkeys_Activated(this, id); };
};

void* UGlobalHotkeys_NewUGlobalHotkeys(void* parent)
{
	return new MyUGlobalHotkeys(static_cast<QWidget*>(parent));
}

void UGlobalHotkeys_RegisterHotkey(void* ptr, struct UGlobalHotkeys_PackedString keySeq, unsigned long id)
{
	static_cast<UGlobalHotkeys*>(ptr)->registerHotkey(QString::fromUtf8(keySeq.data, keySeq.len), id);
}

void UGlobalHotkeys_UnregisterHotkey(void* ptr, unsigned long id)
{
	static_cast<UGlobalHotkeys*>(ptr)->unregisterHotkey(id);
}

void UGlobalHotkeys_UnregisterAllHotkeys(void* ptr)
{
	static_cast<UGlobalHotkeys*>(ptr)->unregisterAllHotkeys();
}

void UGlobalHotkeys_DestroyUGlobalHotkeys(void* ptr)
{
	static_cast<UGlobalHotkeys*>(ptr)->~UGlobalHotkeys();
}

void UGlobalHotkeys_ConnectActivated(void* ptr)
{
	QObject::connect(static_cast<UGlobalHotkeys*>(ptr), &UGlobalHotkeys::activated, static_cast<MyUGlobalHotkeys*>(ptr), &MyUGlobalHotkeys::Signal_Activated);
}

void UGlobalHotkeys_DisconnectActivated(void* ptr)
{
	QObject::disconnect(static_cast<UGlobalHotkeys*>(ptr), &UGlobalHotkeys::activated, static_cast<MyUGlobalHotkeys*>(ptr), &MyUGlobalHotkeys::Signal_Activated);
}

void UGlobalHotkeys_Activated(void* ptr, unsigned long id)
{
	static_cast<UGlobalHotkeys*>(ptr)->activated(id);
}
