#include "qwindowstatechangeevent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWindow>
#include <QString>
#include <QWindowStateChangeEvent>
#include "_cgo_export.h"

class MyQWindowStateChangeEvent: public QWindowStateChangeEvent {
public:
};

int QWindowStateChangeEvent_OldState(void* ptr){
	return static_cast<QWindowStateChangeEvent*>(ptr)->oldState();
}

