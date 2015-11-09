#include "qparallelanimationgroup.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QParallelAnimationGroup>
#include "_cgo_export.h"

class MyQParallelAnimationGroup: public QParallelAnimationGroup {
public:
};

int QParallelAnimationGroup_Duration(void* ptr){
	return static_cast<QParallelAnimationGroup*>(ptr)->duration();
}

void QParallelAnimationGroup_DestroyQParallelAnimationGroup(void* ptr){
	static_cast<QParallelAnimationGroup*>(ptr)->~QParallelAnimationGroup();
}

