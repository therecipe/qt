#include "qsgdynamictexture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSGDynamicTexture>
#include "_cgo_export.h"

class MyQSGDynamicTexture: public QSGDynamicTexture {
public:
};

int QSGDynamicTexture_UpdateTexture(void* ptr){
	return static_cast<QSGDynamicTexture*>(ptr)->updateTexture();
}

