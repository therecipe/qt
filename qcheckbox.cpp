#include "qcheckbox.h"
#include <QCheckBox>
#include "cgoexport.h"

//MyQCheckBox
class MyQCheckBox: public QCheckBox {
Q_OBJECT
public:
void Signal_StateChanged() { callbackQCheckBox(this, QString("stateChanged").toUtf8().data()); };

signals:

};
#include "qcheckbox.moc"


//Public Functions
QtObjectPtr QCheckBox_New_QWidget(QtObjectPtr parent)
{
	return new QCheckBox(((QWidget*)(parent)));
}

QtObjectPtr QCheckBox_New_String_QWidget(char* text, QtObjectPtr parent)
{
	return new QCheckBox(QString(text), ((QWidget*)(parent)));
}

void QCheckBox_Destroy(QtObjectPtr ptr)
{
	((QCheckBox*)(ptr))->~QCheckBox();
}

int QCheckBox_CheckState(QtObjectPtr ptr)
{
	return ((QCheckBox*)(ptr))->checkState();
}

int QCheckBox_IsTristate(QtObjectPtr ptr)
{
	return ((QCheckBox*)(ptr))->isTristate();
}

void QCheckBox_SetCheckState_CheckState(QtObjectPtr ptr, int state)
{
	((QCheckBox*)(ptr))->setCheckState(((Qt::CheckState)(state)));
}

void QCheckBox_SetTristate_Bool(QtObjectPtr ptr, int y)
{
	((QCheckBox*)(ptr))->setTristate(y != 0);
}

//Signals
void QCheckBox_ConnectSignalStateChanged(QtObjectPtr ptr)
{
	QObject::connect(((QCheckBox*)(ptr)), &QCheckBox::stateChanged, ((MyQCheckBox*)(ptr)), &MyQCheckBox::Signal_StateChanged, Qt::QueuedConnection);
}

void QCheckBox_DisconnectSignalStateChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QCheckBox*)(ptr)), &QCheckBox::stateChanged, ((MyQCheckBox*)(ptr)), &MyQCheckBox::Signal_StateChanged);
}

