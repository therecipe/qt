#include "qvalidator.h"
#include <QValidator>
#include "cgoexport.h"

//MyQValidator
class MyQValidator: public QValidator {
Q_OBJECT
public:
void Signal_Changed() { callbackQValidator(this, QString("changed").toUtf8().data()); };

signals:

};
#include "qvalidator.moc"

//Public Types
int QValidator_Invalid() { return QValidator::Invalid; }
int QValidator_Intermediate() { return QValidator::Intermediate; }
int QValidator_Acceptable() { return QValidator::Acceptable; }

//Public Functions
void QValidator_Destroy(QtObjectPtr ptr)
{
	((QValidator*)(ptr))->~QValidator();
}

//Signals
void QValidator_ConnectSignalChanged(QtObjectPtr ptr)
{
	QObject::connect(((QValidator*)(ptr)), &QValidator::changed, ((MyQValidator*)(ptr)), &MyQValidator::Signal_Changed, Qt::QueuedConnection);
}

void QValidator_DisconnectSignalChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QValidator*)(ptr)), &QValidator::changed, ((MyQValidator*)(ptr)), &MyQValidator::Signal_Changed);
}

