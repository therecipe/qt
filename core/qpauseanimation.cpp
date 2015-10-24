#include "qpauseanimation.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QPauseAnimation>
#include "_cgo_export.h"

class MyQPauseAnimation: public QPauseAnimation {
public:
};

int QPauseAnimation_Duration(QtObjectPtr ptr){
	return static_cast<QPauseAnimation*>(ptr)->duration();
}

void QPauseAnimation_SetDuration(QtObjectPtr ptr, int msecs){
	static_cast<QPauseAnimation*>(ptr)->setDuration(msecs);
}

QtObjectPtr QPauseAnimation_NewQPauseAnimation(QtObjectPtr parent){
	return new QPauseAnimation(static_cast<QObject*>(parent));
}

QtObjectPtr QPauseAnimation_NewQPauseAnimation2(int msecs, QtObjectPtr parent){
	return new QPauseAnimation(msecs, static_cast<QObject*>(parent));
}

void QPauseAnimation_DestroyQPauseAnimation(QtObjectPtr ptr){
	static_cast<QPauseAnimation*>(ptr)->~QPauseAnimation();
}

