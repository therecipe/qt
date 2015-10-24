#include "qhoverevent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QEvent>
#include <QString>
#include <QHoverEvent>
#include "_cgo_export.h"

class MyQHoverEvent: public QHoverEvent {
public:
};

QtObjectPtr QHoverEvent_NewQHoverEvent(int ty, QtObjectPtr pos, QtObjectPtr oldPos, int modifiers){
	return new QHoverEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(pos), *static_cast<QPointF*>(oldPos), static_cast<Qt::KeyboardModifier>(modifiers));
}

