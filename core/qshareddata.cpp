#include "qshareddata.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QSharedData>
#include "_cgo_export.h"

class MyQSharedData: public QSharedData {
public:
};

void* QSharedData_NewQSharedData(){
	return new QSharedData();
}

void* QSharedData_NewQSharedData2(void* other){
	return new QSharedData(*static_cast<QSharedData*>(other));
}

