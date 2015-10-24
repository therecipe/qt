#include "qflag.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QFlag>
#include "_cgo_export.h"

class MyQFlag: public QFlag {
public:
};

QtObjectPtr QFlag_NewQFlag(int value){
	return new QFlag(value);
}

