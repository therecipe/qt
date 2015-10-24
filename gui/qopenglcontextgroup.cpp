#include "qopenglcontextgroup.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLContext>
#include <QOpenGLContextGroup>
#include "_cgo_export.h"

class MyQOpenGLContextGroup: public QOpenGLContextGroup {
public:
};

QtObjectPtr QOpenGLContextGroup_QOpenGLContextGroup_CurrentContextGroup(){
	return QOpenGLContextGroup::currentContextGroup();
}

