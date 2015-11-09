#include "qpauseanimation.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QPauseAnimation>
#include "_cgo_export.h"

class MyQPauseAnimation: public QPauseAnimation {
public:
};

int QPauseAnimation_Duration(void* ptr){
	return static_cast<QPauseAnimation*>(ptr)->duration();
}

void QPauseAnimation_SetDuration(void* ptr, int msecs){
	static_cast<QPauseAnimation*>(ptr)->setDuration(msecs);
}

void* QPauseAnimation_NewQPauseAnimation(void* parent){
	return new QPauseAnimation(static_cast<QObject*>(parent));
}

void* QPauseAnimation_NewQPauseAnimation2(int msecs, void* parent){
	return new QPauseAnimation(msecs, static_cast<QObject*>(parent));
}

void QPauseAnimation_DestroyQPauseAnimation(void* ptr){
	static_cast<QPauseAnimation*>(ptr)->~QPauseAnimation();
}

