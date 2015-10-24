#include "qmouseevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QMouseEvent>
#include "_cgo_export.h"

class MyQMouseEvent: public QMouseEvent {
public:
};

QtObjectPtr QMouseEvent_NewQMouseEvent(int ty, QtObjectPtr localPos, int button, int buttons, int modifiers){
	return new QMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

QtObjectPtr QMouseEvent_NewQMouseEvent2(int ty, QtObjectPtr localPos, QtObjectPtr screenPos, int button, int buttons, int modifiers){
	return new QMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

QtObjectPtr QMouseEvent_NewQMouseEvent3(int ty, QtObjectPtr localPos, QtObjectPtr windowPos, QtObjectPtr screenPos, int button, int buttons, int modifiers){
	return new QMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

int QMouseEvent_Button(QtObjectPtr ptr){
	return static_cast<QMouseEvent*>(ptr)->button();
}

int QMouseEvent_Buttons(QtObjectPtr ptr){
	return static_cast<QMouseEvent*>(ptr)->buttons();
}

int QMouseEvent_Flags(QtObjectPtr ptr){
	return static_cast<QMouseEvent*>(ptr)->flags();
}

int QMouseEvent_GlobalX(QtObjectPtr ptr){
	return static_cast<QMouseEvent*>(ptr)->globalX();
}

int QMouseEvent_GlobalY(QtObjectPtr ptr){
	return static_cast<QMouseEvent*>(ptr)->globalY();
}

int QMouseEvent_Source(QtObjectPtr ptr){
	return static_cast<QMouseEvent*>(ptr)->source();
}

int QMouseEvent_X(QtObjectPtr ptr){
	return static_cast<QMouseEvent*>(ptr)->x();
}

int QMouseEvent_Y(QtObjectPtr ptr){
	return static_cast<QMouseEvent*>(ptr)->y();
}

