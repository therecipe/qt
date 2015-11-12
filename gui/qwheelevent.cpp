#include "qwheelevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QWheelEvent>
#include "_cgo_export.h"

class MyQWheelEvent: public QWheelEvent {
public:
};

void* QWheelEvent_NewQWheelEvent(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers){
	return new QWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QWheelEvent_NewQWheelEvent4(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase){
	return new QWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase));
}

void* QWheelEvent_NewQWheelEvent5(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase, int source){
	return new QWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase), static_cast<Qt::MouseEventSource>(source));
}

int QWheelEvent_Buttons(void* ptr){
	return static_cast<QWheelEvent*>(ptr)->buttons();
}

int QWheelEvent_GlobalX(void* ptr){
	return static_cast<QWheelEvent*>(ptr)->globalX();
}

int QWheelEvent_GlobalY(void* ptr){
	return static_cast<QWheelEvent*>(ptr)->globalY();
}

int QWheelEvent_Phase(void* ptr){
	return static_cast<QWheelEvent*>(ptr)->phase();
}

int QWheelEvent_Source(void* ptr){
	return static_cast<QWheelEvent*>(ptr)->source();
}

int QWheelEvent_X(void* ptr){
	return static_cast<QWheelEvent*>(ptr)->x();
}

int QWheelEvent_Y(void* ptr){
	return static_cast<QWheelEvent*>(ptr)->y();
}

