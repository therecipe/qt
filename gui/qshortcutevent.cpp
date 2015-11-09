#include "qshortcutevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QKeySequence>
#include <QString>
#include <QVariant>
#include <QShortcutEvent>
#include "_cgo_export.h"

class MyQShortcutEvent: public QShortcutEvent {
public:
};

void* QShortcutEvent_NewQShortcutEvent(void* key, int id, int ambiguous){
	return new QShortcutEvent(*static_cast<QKeySequence*>(key), id, ambiguous != 0);
}

int QShortcutEvent_IsAmbiguous(void* ptr){
	return static_cast<QShortcutEvent*>(ptr)->isAmbiguous();
}

int QShortcutEvent_ShortcutId(void* ptr){
	return static_cast<QShortcutEvent*>(ptr)->shortcutId();
}

void QShortcutEvent_DestroyQShortcutEvent(void* ptr){
	static_cast<QShortcutEvent*>(ptr)->~QShortcutEvent();
}

