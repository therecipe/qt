#include "qtextblockuserdata.h"
#include <QModelIndex>
#include <QTextBlock>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTextBlockUserData>
#include "_cgo_export.h"

class MyQTextBlockUserData: public QTextBlockUserData {
public:
};

void QTextBlockUserData_DestroyQTextBlockUserData(QtObjectPtr ptr){
	static_cast<QTextBlockUserData*>(ptr)->~QTextBlockUserData();
}

