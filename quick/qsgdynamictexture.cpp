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

int QSGDynamicTexture_UpdateTexture(QtObjectPtr ptr){
	return static_cast<QSGDynamicTexture*>(ptr)->updateTexture();
}

