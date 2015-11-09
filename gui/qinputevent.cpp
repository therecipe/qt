#include "qinputevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QInputEvent>
#include "_cgo_export.h"

class MyQInputEvent: public QInputEvent {
public:
};

int QInputEvent_Modifiers(void* ptr){
	return static_cast<QInputEvent*>(ptr)->modifiers();
}

