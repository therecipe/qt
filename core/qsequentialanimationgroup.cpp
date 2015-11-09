#include "qsequentialanimationgroup.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractAnimation>
#include <QObject>
#include <QSequentialAnimationGroup>
#include "_cgo_export.h"

class MyQSequentialAnimationGroup: public QSequentialAnimationGroup {
public:
void Signal_CurrentAnimationChanged(QAbstractAnimation * current){callbackQSequentialAnimationGroupCurrentAnimationChanged(this->objectName().toUtf8().data(), current);};
};

void* QSequentialAnimationGroup_CurrentAnimation(void* ptr){
	return static_cast<QSequentialAnimationGroup*>(ptr)->currentAnimation();
}

void* QSequentialAnimationGroup_AddPause(void* ptr, int msecs){
	return static_cast<QSequentialAnimationGroup*>(ptr)->addPause(msecs);
}

void QSequentialAnimationGroup_ConnectCurrentAnimationChanged(void* ptr){
	QObject::connect(static_cast<QSequentialAnimationGroup*>(ptr), static_cast<void (QSequentialAnimationGroup::*)(QAbstractAnimation *)>(&QSequentialAnimationGroup::currentAnimationChanged), static_cast<MyQSequentialAnimationGroup*>(ptr), static_cast<void (MyQSequentialAnimationGroup::*)(QAbstractAnimation *)>(&MyQSequentialAnimationGroup::Signal_CurrentAnimationChanged));;
}

void QSequentialAnimationGroup_DisconnectCurrentAnimationChanged(void* ptr){
	QObject::disconnect(static_cast<QSequentialAnimationGroup*>(ptr), static_cast<void (QSequentialAnimationGroup::*)(QAbstractAnimation *)>(&QSequentialAnimationGroup::currentAnimationChanged), static_cast<MyQSequentialAnimationGroup*>(ptr), static_cast<void (MyQSequentialAnimationGroup::*)(QAbstractAnimation *)>(&MyQSequentialAnimationGroup::Signal_CurrentAnimationChanged));;
}

int QSequentialAnimationGroup_Duration(void* ptr){
	return static_cast<QSequentialAnimationGroup*>(ptr)->duration();
}

void* QSequentialAnimationGroup_InsertPause(void* ptr, int index, int msecs){
	return static_cast<QSequentialAnimationGroup*>(ptr)->insertPause(index, msecs);
}

void QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(void* ptr){
	static_cast<QSequentialAnimationGroup*>(ptr)->~QSequentialAnimationGroup();
}

