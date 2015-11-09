#include "qdragenterevent.h"
#include <QMimeData>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDrag>
#include <QDragEnterEvent>
#include "_cgo_export.h"

class MyQDragEnterEvent: public QDragEnterEvent {
public:
};

void* QDragEnterEvent_NewQDragEnterEvent(void* point, int actions, void* data, int buttons, int modifiers){
	return new QDragEnterEvent(*static_cast<QPoint*>(point), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

