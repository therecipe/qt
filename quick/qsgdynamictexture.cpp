#include "qsgdynamictexture.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QSGDynamicTexture>
#include "_cgo_export.h"

class MyQSGDynamicTexture: public QSGDynamicTexture {
public:
};

int QSGDynamicTexture_UpdateTexture(void* ptr){
	return static_cast<QSGDynamicTexture*>(ptr)->updateTexture();
}

