#include "qmaskgenerator.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMaskGenerator>
#include "_cgo_export.h"

class MyQMaskGenerator: public QMaskGenerator {
public:
};

int QMaskGenerator_Seed(QtObjectPtr ptr){
	return static_cast<QMaskGenerator*>(ptr)->seed();
}

void QMaskGenerator_DestroyQMaskGenerator(QtObjectPtr ptr){
	static_cast<QMaskGenerator*>(ptr)->~QMaskGenerator();
}

