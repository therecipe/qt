#include "qsequentialanimationgroup.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAbstractAnimation>
#include <QString>
#include <QSequentialAnimationGroup>
#include "_cgo_export.h"

class MyQSequentialAnimationGroup: public QSequentialAnimationGroup {
public:
void Signal_CurrentAnimationChanged(QAbstractAnimation * current){callbackQSequentialAnimationGroupCurrentAnimationChanged(this->objectName().toUtf8().data(), current);};
};

QtObjectPtr QSequentialAnimationGroup_CurrentAnimation(QtObjectPtr ptr){
	return static_cast<QSequentialAnimationGroup*>(ptr)->currentAnimation();
}

QtObjectPtr QSequentialAnimationGroup_AddPause(QtObjectPtr ptr, int msecs){
	return static_cast<QSequentialAnimationGroup*>(ptr)->addPause(msecs);
}

void QSequentialAnimationGroup_ConnectCurrentAnimationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSequentialAnimationGroup*>(ptr), static_cast<void (QSequentialAnimationGroup::*)(QAbstractAnimation *)>(&QSequentialAnimationGroup::currentAnimationChanged), static_cast<MyQSequentialAnimationGroup*>(ptr), static_cast<void (MyQSequentialAnimationGroup::*)(QAbstractAnimation *)>(&MyQSequentialAnimationGroup::Signal_CurrentAnimationChanged));;
}

void QSequentialAnimationGroup_DisconnectCurrentAnimationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSequentialAnimationGroup*>(ptr), static_cast<void (QSequentialAnimationGroup::*)(QAbstractAnimation *)>(&QSequentialAnimationGroup::currentAnimationChanged), static_cast<MyQSequentialAnimationGroup*>(ptr), static_cast<void (MyQSequentialAnimationGroup::*)(QAbstractAnimation *)>(&MyQSequentialAnimationGroup::Signal_CurrentAnimationChanged));;
}

int QSequentialAnimationGroup_Duration(QtObjectPtr ptr){
	return static_cast<QSequentialAnimationGroup*>(ptr)->duration();
}

QtObjectPtr QSequentialAnimationGroup_InsertPause(QtObjectPtr ptr, int index, int msecs){
	return static_cast<QSequentialAnimationGroup*>(ptr)->insertPause(index, msecs);
}

void QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(QtObjectPtr ptr){
	static_cast<QSequentialAnimationGroup*>(ptr)->~QSequentialAnimationGroup();
}

