#include "qdragenterevent.h"
#include <QPoint>
#include <QMimeData>
#include <QDrag>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDragEnterEvent>
#include "_cgo_export.h"

class MyQDragEnterEvent: public QDragEnterEvent {
public:
};

QtObjectPtr QDragEnterEvent_NewQDragEnterEvent(QtObjectPtr point, int actions, QtObjectPtr data, int buttons, int modifiers){
	return new QDragEnterEvent(*static_cast<QPoint*>(point), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

