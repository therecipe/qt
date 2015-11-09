#include "qfocusevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QFocusEvent>
#include "_cgo_export.h"

class MyQFocusEvent: public QFocusEvent {
public:
};

void* QFocusEvent_NewQFocusEvent(int ty, int reason){
	return new QFocusEvent(static_cast<QEvent::Type>(ty), static_cast<Qt::FocusReason>(reason));
}

int QFocusEvent_GotFocus(void* ptr){
	return static_cast<QFocusEvent*>(ptr)->gotFocus();
}

int QFocusEvent_LostFocus(void* ptr){
	return static_cast<QFocusEvent*>(ptr)->lostFocus();
}

int QFocusEvent_Reason(void* ptr){
	return static_cast<QFocusEvent*>(ptr)->reason();
}

