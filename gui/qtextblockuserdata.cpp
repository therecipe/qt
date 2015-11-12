#include "qtextblockuserdata.h"
#include <QTextBlock>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextBlockUserData>
#include "_cgo_export.h"

class MyQTextBlockUserData: public QTextBlockUserData {
public:
};

void QTextBlockUserData_DestroyQTextBlockUserData(void* ptr){
	static_cast<QTextBlockUserData*>(ptr)->~QTextBlockUserData();
}

