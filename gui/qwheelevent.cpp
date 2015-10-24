#include "qwheelevent.h"
#include <QPointF>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWheelEvent>
#include "_cgo_export.h"

class MyQWheelEvent: public QWheelEvent {
public:
};

QtObjectPtr QWheelEvent_NewQWheelEvent(QtObjectPtr pos, QtObjectPtr globalPos, QtObjectPtr pixelDelta, QtObjectPtr angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers){
	return new QWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

QtObjectPtr QWheelEvent_NewQWheelEvent4(QtObjectPtr pos, QtObjectPtr globalPos, QtObjectPtr pixelDelta, QtObjectPtr angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase){
	return new QWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase));
}

QtObjectPtr QWheelEvent_NewQWheelEvent5(QtObjectPtr pos, QtObjectPtr globalPos, QtObjectPtr pixelDelta, QtObjectPtr angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase, int source){
	return new QWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase), static_cast<Qt::MouseEventSource>(source));
}

int QWheelEvent_Buttons(QtObjectPtr ptr){
	return static_cast<QWheelEvent*>(ptr)->buttons();
}

int QWheelEvent_GlobalX(QtObjectPtr ptr){
	return static_cast<QWheelEvent*>(ptr)->globalX();
}

int QWheelEvent_GlobalY(QtObjectPtr ptr){
	return static_cast<QWheelEvent*>(ptr)->globalY();
}

int QWheelEvent_Phase(QtObjectPtr ptr){
	return static_cast<QWheelEvent*>(ptr)->phase();
}

int QWheelEvent_Source(QtObjectPtr ptr){
	return static_cast<QWheelEvent*>(ptr)->source();
}

int QWheelEvent_X(QtObjectPtr ptr){
	return static_cast<QWheelEvent*>(ptr)->x();
}

int QWheelEvent_Y(QtObjectPtr ptr){
	return static_cast<QWheelEvent*>(ptr)->y();
}

