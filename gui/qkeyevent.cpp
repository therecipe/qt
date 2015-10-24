#include "qkeyevent.h"
#include <QModelIndex>
#include <QKeySequence>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QKeyEvent>
#include "_cgo_export.h"

class MyQKeyEvent: public QKeyEvent {
public:
};

int QKeyEvent_Matches(QtObjectPtr ptr, int key){
	return static_cast<QKeyEvent*>(ptr)->matches(static_cast<QKeySequence::StandardKey>(key));
}

int QKeyEvent_Count(QtObjectPtr ptr){
	return static_cast<QKeyEvent*>(ptr)->count();
}

int QKeyEvent_IsAutoRepeat(QtObjectPtr ptr){
	return static_cast<QKeyEvent*>(ptr)->isAutoRepeat();
}

int QKeyEvent_Key(QtObjectPtr ptr){
	return static_cast<QKeyEvent*>(ptr)->key();
}

int QKeyEvent_Modifiers(QtObjectPtr ptr){
	return static_cast<QKeyEvent*>(ptr)->modifiers();
}

char* QKeyEvent_Text(QtObjectPtr ptr){
	return static_cast<QKeyEvent*>(ptr)->text().toUtf8().data();
}

