#include "qstackedlayout.h"
#include <QStackedLayout>
#include "cgoexport.h"

//MyQStackedLayout
class MyQStackedLayout: public QStackedLayout {
Q_OBJECT
public:
void Signal_CurrentChanged() { callbackQStackedLayout(this, QString("currentChanged").toUtf8().data()); };
void Signal_WidgetRemoved() { callbackQStackedLayout(this, QString("widgetRemoved").toUtf8().data()); };

signals:
void Slot_SetCurrentIndex(int index);

};
#include "qstackedlayout.moc"


//Public Functions
QtObjectPtr QStackedLayout_New()
{
	return new QStackedLayout();
}

QtObjectPtr QStackedLayout_New_QWidget(QtObjectPtr parent)
{
	return new QStackedLayout(((QWidget*)(parent)));
}

QtObjectPtr QStackedLayout_New_QLayout(QtObjectPtr parentLayout)
{
	return new QStackedLayout(((QLayout*)(parentLayout)));
}

void QStackedLayout_Destroy(QtObjectPtr ptr)
{
	((QStackedLayout*)(ptr))->~QStackedLayout();
}

int QStackedLayout_CurrentIndex(QtObjectPtr ptr)
{
	return ((QStackedLayout*)(ptr))->currentIndex();
}

QtObjectPtr QStackedLayout_CurrentWidget(QtObjectPtr ptr)
{
	return ((QStackedLayout*)(ptr))->currentWidget();
}

int QStackedLayout_InsertWidget_Int_QWidget(QtObjectPtr ptr, int index, QtObjectPtr widget)
{
	return ((QStackedLayout*)(ptr))->insertWidget(index, ((QWidget*)(widget)));
}

QtObjectPtr QStackedLayout_Widget_Int(QtObjectPtr ptr, int index)
{
	return ((QStackedLayout*)(ptr))->widget(index);
}

//Public Slots
void QStackedLayout_ConnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::connect(((MyQStackedLayout*)(ptr)), &MyQStackedLayout::Slot_SetCurrentIndex, ((QStackedLayout*)(ptr)), &QStackedLayout::setCurrentIndex, Qt::QueuedConnection);
}

void QStackedLayout_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQStackedLayout*)(ptr)), &MyQStackedLayout::Slot_SetCurrentIndex, ((QStackedLayout*)(ptr)), &QStackedLayout::setCurrentIndex);
}

void QStackedLayout_SetCurrentIndex_Int(QtObjectPtr ptr, int index)
{
	((MyQStackedLayout*)(ptr))->Slot_SetCurrentIndex(index);
}

//Signals
void QStackedLayout_ConnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::connect(((QStackedLayout*)(ptr)), &QStackedLayout::currentChanged, ((MyQStackedLayout*)(ptr)), &MyQStackedLayout::Signal_CurrentChanged, Qt::QueuedConnection);
}

void QStackedLayout_DisconnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QStackedLayout*)(ptr)), &QStackedLayout::currentChanged, ((MyQStackedLayout*)(ptr)), &MyQStackedLayout::Signal_CurrentChanged);
}

void QStackedLayout_ConnectSignalWidgetRemoved(QtObjectPtr ptr)
{
	QObject::connect(((QStackedLayout*)(ptr)), &QStackedLayout::widgetRemoved, ((MyQStackedLayout*)(ptr)), &MyQStackedLayout::Signal_WidgetRemoved, Qt::QueuedConnection);
}

void QStackedLayout_DisconnectSignalWidgetRemoved(QtObjectPtr ptr)
{
	QObject::disconnect(((QStackedLayout*)(ptr)), &QStackedLayout::widgetRemoved, ((MyQStackedLayout*)(ptr)), &MyQStackedLayout::Signal_WidgetRemoved);
}

