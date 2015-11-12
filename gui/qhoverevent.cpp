#include "qhoverevent.h"
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QPoint>
#include <QHoverEvent>
#include "_cgo_export.h"

class MyQHoverEvent: public QHoverEvent {
public:
};

void* QHoverEvent_NewQHoverEvent(int ty, void* pos, void* oldPos, int modifiers){
	return new QHoverEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(pos), *static_cast<QPointF*>(oldPos), static_cast<Qt::KeyboardModifier>(modifiers));
}

