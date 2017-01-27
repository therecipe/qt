// +build !minimal

#define protected public
#define private public

#include "testlib.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QList>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QPoint>
#include <QSignalSpy>
#include <QString>
#include <QTestEventList>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>

void* QSignalSpy_NewQSignalSpy(void* object, char* sign)
{
	return new QSignalSpy(static_cast<QObject*>(object), const_cast<const char*>(sign));
}

char QSignalSpy_IsValid(void* ptr)
{
		return static_cast<QSignalSpy*>(ptr)->isValid();
}

void* QSignalSpy_Signal(void* ptr)
{
		return new QByteArray(static_cast<QSignalSpy*>(ptr)->signal());
}

char QSignalSpy_Wait(void* ptr, int timeout)
{
		return static_cast<QSignalSpy*>(ptr)->wait(timeout);
}

void* QSignalSpy___children_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSignalSpy___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___children_newList(void* ptr)
{
		return new QList<QObject *>;
}

void* QSignalSpy___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSignalSpy___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSignalSpy___dynamicPropertyNames_newList(void* ptr)
{
		return new QList<QByteArray>;
}

void* QSignalSpy___findChildren_atList2(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___findChildren_newList2(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___findChildren_atList3(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___findChildren_newList3(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___findChildren_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___findChildren_newList(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___QList_other_atList3(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___QList_other_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___QList_other_newList3(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___QList_other_atList2(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___QList_other_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___QList_other_newList2(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___append_value_atList2(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___append_value_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___append_value_newList2(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___fromSet_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___fromSet_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___fromSet_newList(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___fromStdList_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___fromStdList_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___fromStdList_newList(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___fromVector_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___fromVector_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___fromVector_newList(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___fromVector_vector_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___fromVector_vector_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___fromVector_vector_newList(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___mid_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___mid_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___mid_newList(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___swap_other_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___swap_other_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___swap_other_newList(void* ptr)
{
		return new QList<QObject*>;
}

void* QSignalSpy___toVector_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSignalSpy___toVector_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QSignalSpy___toVector_newList(void* ptr)
{
		return new QList<QObject*>;
}

void QSignalSpy_TimerEvent(void* ptr, void* event)
{
		static_cast<QSignalSpy*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSignalSpy_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSignalSpy_ChildEvent(void* ptr, void* event)
{
		static_cast<QSignalSpy*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSignalSpy_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::childEvent(static_cast<QChildEvent*>(event));
}

void QSignalSpy_ConnectNotify(void* ptr, void* sign)
{
		static_cast<QSignalSpy*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSignalSpy_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSignalSpy_CustomEvent(void* ptr, void* event)
{
		static_cast<QSignalSpy*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSignalSpy_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::customEvent(static_cast<QEvent*>(event));
}

void QSignalSpy_DeleteLater(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QSignalSpy*>(ptr), "deleteLater");
}

void QSignalSpy_DeleteLaterDefault(void* ptr)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::deleteLater();
}

void QSignalSpy_DisconnectNotify(void* ptr, void* sign)
{
		static_cast<QSignalSpy*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSignalSpy_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSignalSpy*>(ptr)->QSignalSpy::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QSignalSpy_Event(void* ptr, void* e)
{
		return static_cast<QSignalSpy*>(ptr)->event(static_cast<QEvent*>(e));
}

char QSignalSpy_EventDefault(void* ptr, void* e)
{
		return static_cast<QSignalSpy*>(ptr)->QSignalSpy::event(static_cast<QEvent*>(e));
}

char QSignalSpy_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(watched))) {
		return static_cast<QSignalSpy*>(ptr)->eventFilter(static_cast<QSignalSpy*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QSignalSpy*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QSignalSpy_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(watched))) {
		return static_cast<QSignalSpy*>(ptr)->QSignalSpy::eventFilter(static_cast<QSignalSpy*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QSignalSpy*>(ptr)->QSignalSpy::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QSignalSpy_MetaObject(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSignalSpy*>(ptr)->metaObject());
}

void* QSignalSpy_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSignalSpy*>(ptr)->QSignalSpy::metaObject());
}

void* QTestEventList_NewQTestEventList()
{
	return new QTestEventList();
}

void* QTestEventList_NewQTestEventList2(void* other)
{
	return new QTestEventList(*static_cast<QTestEventList*>(other));
}

void QTestEventList_AddDelay(void* ptr, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addDelay(msecs);
}

void QTestEventList_AddKeyClick(void* ptr, long long qtKey, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyClick(static_cast<Qt::Key>(qtKey), static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyClick2(void* ptr, char* ascii, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyClick(*ascii, static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyClicks(void* ptr, char* keys, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyClicks(QString(keys), static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyPress(void* ptr, long long qtKey, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyPress(static_cast<Qt::Key>(qtKey), static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyPress2(void* ptr, char* ascii, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyPress(*ascii, static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyRelease(void* ptr, long long qtKey, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyRelease(static_cast<Qt::Key>(qtKey), static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddKeyRelease2(void* ptr, char* ascii, long long modifiers, int msecs)
{
	static_cast<QTestEventList*>(ptr)->addKeyRelease(*ascii, static_cast<Qt::KeyboardModifier>(modifiers), msecs);
}

void QTestEventList_AddMouseClick(void* ptr, long long button, long long modifiers, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMouseClick(static_cast<Qt::MouseButton>(button), static_cast<Qt::KeyboardModifier>(modifiers), *static_cast<QPoint*>(pos), delay);
}

void QTestEventList_AddMouseDClick(void* ptr, long long button, long long modifiers, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMouseDClick(static_cast<Qt::MouseButton>(button), static_cast<Qt::KeyboardModifier>(modifiers), *static_cast<QPoint*>(pos), delay);
}

void QTestEventList_AddMouseMove(void* ptr, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMouseMove(*static_cast<QPoint*>(pos), delay);
}

void QTestEventList_AddMousePress(void* ptr, long long button, long long modifiers, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMousePress(static_cast<Qt::MouseButton>(button), static_cast<Qt::KeyboardModifier>(modifiers), *static_cast<QPoint*>(pos), delay);
}

void QTestEventList_AddMouseRelease(void* ptr, long long button, long long modifiers, void* pos, int delay)
{
	static_cast<QTestEventList*>(ptr)->addMouseRelease(static_cast<Qt::MouseButton>(button), static_cast<Qt::KeyboardModifier>(modifiers), *static_cast<QPoint*>(pos), delay);
}

void QTestEventList_Clear(void* ptr)
{
	static_cast<QTestEventList*>(ptr)->clear();
}

void QTestEventList_Simulate(void* ptr, void* w)
{
	static_cast<QTestEventList*>(ptr)->simulate(static_cast<QWidget*>(w));
}

void QTestEventList_DestroyQTestEventList(void* ptr)
{
	static_cast<QTestEventList*>(ptr)->~QTestEventList();
}

void* QTestEventList___QList_other_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___QList_other_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___QList_other_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QTestEventList___QList_other_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___QList_other_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___QList_other_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QTestEventList___append_value_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___append_value_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___append_value_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QTestEventList___fromSet_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___fromSet_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___fromSet_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QTestEventList___fromStdList_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___fromStdList_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___fromStdList_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QTestEventList___fromVector_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___fromVector_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___fromVector_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QTestEventList___fromVector_vector_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___fromVector_vector_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___fromVector_vector_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QTestEventList___mid_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___mid_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___mid_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QTestEventList___swap_other_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___swap_other_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___swap_other_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QTestEventList___toVector_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QTestEventList___toVector_setList(void* ptr, void* i)
{
	if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QSignalSpy*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QTestEventList___toVector_newList(void* ptr)
{
	return new QList<QObject*>;
}

