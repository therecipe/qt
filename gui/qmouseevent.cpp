#include "qmouseevent.h"
#include <QPoint>
#include <QPointF>
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMouseEvent>
#include "_cgo_export.h"

class MyQMouseEvent: public QMouseEvent {
public:
};

void* QMouseEvent_NewQMouseEvent(int ty, void* localPos, int button, int buttons, int modifiers){
	return new QMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent2(int ty, void* localPos, void* screenPos, int button, int buttons, int modifiers){
	return new QMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent3(int ty, void* localPos, void* windowPos, void* screenPos, int button, int buttons, int modifiers){
	return new QMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

int QMouseEvent_Button(void* ptr){
	return static_cast<QMouseEvent*>(ptr)->button();
}

int QMouseEvent_Buttons(void* ptr){
	return static_cast<QMouseEvent*>(ptr)->buttons();
}

int QMouseEvent_Flags(void* ptr){
	return static_cast<QMouseEvent*>(ptr)->flags();
}

int QMouseEvent_GlobalX(void* ptr){
	return static_cast<QMouseEvent*>(ptr)->globalX();
}

int QMouseEvent_GlobalY(void* ptr){
	return static_cast<QMouseEvent*>(ptr)->globalY();
}

int QMouseEvent_Source(void* ptr){
	return static_cast<QMouseEvent*>(ptr)->source();
}

int QMouseEvent_X(void* ptr){
	return static_cast<QMouseEvent*>(ptr)->x();
}

int QMouseEvent_Y(void* ptr){
	return static_cast<QMouseEvent*>(ptr)->y();
}

