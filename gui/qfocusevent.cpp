#include "qfocusevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QFocusEvent>
#include "_cgo_export.h"

class MyQFocusEvent: public QFocusEvent {
public:
};

QtObjectPtr QFocusEvent_NewQFocusEvent(int ty, int reason){
	return new QFocusEvent(static_cast<QEvent::Type>(ty), static_cast<Qt::FocusReason>(reason));
}

int QFocusEvent_GotFocus(QtObjectPtr ptr){
	return static_cast<QFocusEvent*>(ptr)->gotFocus();
}

int QFocusEvent_LostFocus(QtObjectPtr ptr){
	return static_cast<QFocusEvent*>(ptr)->lostFocus();
}

int QFocusEvent_Reason(QtObjectPtr ptr){
	return static_cast<QFocusEvent*>(ptr)->reason();
}

