#include "qinputmethodqueryevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QInputMethod>
#include <QInputMethodQueryEvent>
#include "_cgo_export.h"

class MyQInputMethodQueryEvent: public QInputMethodQueryEvent {
public:
};

QtObjectPtr QInputMethodQueryEvent_NewQInputMethodQueryEvent(int queries){
	return new QInputMethodQueryEvent(static_cast<Qt::InputMethodQuery>(queries));
}

int QInputMethodQueryEvent_Queries(QtObjectPtr ptr){
	return static_cast<QInputMethodQueryEvent*>(ptr)->queries();
}

void QInputMethodQueryEvent_SetValue(QtObjectPtr ptr, int query, char* value){
	static_cast<QInputMethodQueryEvent*>(ptr)->setValue(static_cast<Qt::InputMethodQuery>(query), QVariant(value));
}

char* QInputMethodQueryEvent_Value(QtObjectPtr ptr, int query){
	return static_cast<QInputMethodQueryEvent*>(ptr)->value(static_cast<Qt::InputMethodQuery>(query)).toString().toUtf8().data();
}

