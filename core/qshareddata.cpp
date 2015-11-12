#include "qshareddata.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

