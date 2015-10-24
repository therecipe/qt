#include "qinputevent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QInputEvent>
#include "_cgo_export.h"

class MyQInputEvent: public QInputEvent {
public:
};

int QInputEvent_Modifiers(QtObjectPtr ptr){
	return static_cast<QInputEvent*>(ptr)->modifiers();
}

