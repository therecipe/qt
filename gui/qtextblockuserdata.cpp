#include "qtextblockuserdata.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextBlock>
#include <QString>
#include <QTextBlockUserData>
#include "_cgo_export.h"

class MyQTextBlockUserData: public QTextBlockUserData {
public:
};

void QTextBlockUserData_DestroyQTextBlockUserData(void* ptr){
	static_cast<QTextBlockUserData*>(ptr)->~QTextBlockUserData();
}

