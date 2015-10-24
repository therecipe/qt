#include "qwindowstatechangeevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWindow>
#include <QWindowStateChangeEvent>
#include "_cgo_export.h"

class MyQWindowStateChangeEvent: public QWindowStateChangeEvent {
public:
};

int QWindowStateChangeEvent_OldState(QtObjectPtr ptr){
	return static_cast<QWindowStateChangeEvent*>(ptr)->oldState();
}

