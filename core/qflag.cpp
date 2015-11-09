#include "qflag.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFlag>
#include "_cgo_export.h"

class MyQFlag: public QFlag {
public:
};

void* QFlag_NewQFlag(int value){
	return new QFlag(value);
}

