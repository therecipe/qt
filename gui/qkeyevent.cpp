#include "qkeyevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QKeySequence>
#include <QKeyEvent>
#include "_cgo_export.h"

class MyQKeyEvent: public QKeyEvent {
public:
};

int QKeyEvent_Matches(void* ptr, int key){
	return static_cast<QKeyEvent*>(ptr)->matches(static_cast<QKeySequence::StandardKey>(key));
}

int QKeyEvent_Count(void* ptr){
	return static_cast<QKeyEvent*>(ptr)->count();
}

int QKeyEvent_IsAutoRepeat(void* ptr){
	return static_cast<QKeyEvent*>(ptr)->isAutoRepeat();
}

int QKeyEvent_Key(void* ptr){
	return static_cast<QKeyEvent*>(ptr)->key();
}

int QKeyEvent_Modifiers(void* ptr){
	return static_cast<QKeyEvent*>(ptr)->modifiers();
}

char* QKeyEvent_Text(void* ptr){
	return static_cast<QKeyEvent*>(ptr)->text().toUtf8().data();
}

