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

QtObjectPtr QSharedData_NewQSharedData(){
	return new QSharedData();
}

QtObjectPtr QSharedData_NewQSharedData2(QtObjectPtr other){
	return new QSharedData(*static_cast<QSharedData*>(other));
}

