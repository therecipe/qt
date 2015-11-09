#include "qinputmethodqueryevent.h"
#include <QModelIndex>
#include <QInputMethod>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QInputMethodQueryEvent>
#include "_cgo_export.h"

class MyQInputMethodQueryEvent: public QInputMethodQueryEvent {
public:
};

void* QInputMethodQueryEvent_NewQInputMethodQueryEvent(int queries){
	return new QInputMethodQueryEvent(static_cast<Qt::InputMethodQuery>(queries));
}

int QInputMethodQueryEvent_Queries(void* ptr){
	return static_cast<QInputMethodQueryEvent*>(ptr)->queries();
}

void QInputMethodQueryEvent_SetValue(void* ptr, int query, void* value){
	static_cast<QInputMethodQueryEvent*>(ptr)->setValue(static_cast<Qt::InputMethodQuery>(query), *static_cast<QVariant*>(value));
}

void* QInputMethodQueryEvent_Value(void* ptr, int query){
	return new QVariant(static_cast<QInputMethodQueryEvent*>(ptr)->value(static_cast<Qt::InputMethodQuery>(query)));
}

