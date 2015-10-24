#include "qshortcutevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QKeySequence>
#include <QShortcutEvent>
#include "_cgo_export.h"

class MyQShortcutEvent: public QShortcutEvent {
public:
};

QtObjectPtr QShortcutEvent_NewQShortcutEvent(QtObjectPtr key, int id, int ambiguous){
	return new QShortcutEvent(*static_cast<QKeySequence*>(key), id, ambiguous != 0);
}

int QShortcutEvent_IsAmbiguous(QtObjectPtr ptr){
	return static_cast<QShortcutEvent*>(ptr)->isAmbiguous();
}

int QShortcutEvent_ShortcutId(QtObjectPtr ptr){
	return static_cast<QShortcutEvent*>(ptr)->shortcutId();
}

void QShortcutEvent_DestroyQShortcutEvent(QtObjectPtr ptr){
	static_cast<QShortcutEvent*>(ptr)->~QShortcutEvent();
}

