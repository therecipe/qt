#include "qapplication.h"
#include <QApplication>
#include "cgoexport.h"

//MyQApplication
class MyQApplication: public QApplication {
Q_OBJECT
public:

signals:
void Slot_AutoSipEnabled();

};
#include "qapplication.moc"


//Public Functions
QtObjectPtr QApplication_New_Int_String(int argc, char* argv)
{
	return new QApplication(argc, ((char**)(argv)));
}

void QApplication_Destroy(QtObjectPtr ptr)
{
	((QApplication*)(ptr))->~QApplication();
}

char* QApplication_StyleSheet(QtObjectPtr ptr)
{
	return ((QApplication*)(ptr))->styleSheet().toUtf8().data();
}

//Public Slots
void QApplication_ConnectSlotAutoSipEnabled(QtObjectPtr ptr)
{
	QObject::connect(((MyQApplication*)(ptr)), &MyQApplication::Slot_AutoSipEnabled, ((QApplication*)(ptr)), &QApplication::autoSipEnabled, Qt::QueuedConnection);
}

void QApplication_DisconnectSlotAutoSipEnabled(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQApplication*)(ptr)), &MyQApplication::Slot_AutoSipEnabled, ((QApplication*)(ptr)), &QApplication::autoSipEnabled);
}

void QApplication_AutoSipEnabled(QtObjectPtr ptr)
{
	((MyQApplication*)(ptr))->Slot_AutoSipEnabled();
}

//Static Public Members
QtObjectPtr QApplication_ActiveModalWidget()
{
	return QApplication::activeModalWidget();
}

QtObjectPtr QApplication_ActivePopupWidget()
{
	return QApplication::activePopupWidget();
}

QtObjectPtr QApplication_ActiveWindow()
{
	return QApplication::activeWindow();
}

void QApplication_Alert_QWidget_Int(QtObjectPtr widget, int msec)
{
	QApplication::alert(((QWidget*)(widget)), msec);
}

int QApplication_ColorSpec()
{
	return QApplication::colorSpec();
}

int QApplication_CursorFlashTime()
{
	return QApplication::cursorFlashTime();
}

int QApplication_DoubleClickInterval()
{
	return QApplication::doubleClickInterval();
}

int QApplication_Exec()
{
	return QApplication::exec();
}

QtObjectPtr QApplication_FocusWidget()
{
	return QApplication::focusWidget();
}

int QApplication_IsEffectEnabled_UIEffect(int effect)
{
	return QApplication::isEffectEnabled(((Qt::UIEffect)(effect)));
}

int QApplication_KeyboardInputInterval()
{
	return QApplication::keyboardInputInterval();
}

void QApplication_SetActiveWindow_QWidget(QtObjectPtr active)
{
	QApplication::setActiveWindow(((QWidget*)(active)));
}

void QApplication_SetColorSpec_Int(int spec)
{
	QApplication::setColorSpec(spec);
}

void QApplication_SetCursorFlashTime_Int(int flashTime)
{
	QApplication::setCursorFlashTime(flashTime);
}

void QApplication_SetDoubleClickInterval_Int(int doubleClickInterval)
{
	QApplication::setDoubleClickInterval(doubleClickInterval);
}

void QApplication_SetEffectEnabled_UIEffect_Bool(int effect, int enable)
{
	QApplication::setEffectEnabled(((Qt::UIEffect)(effect)), enable != 0);
}

void QApplication_SetKeyboardInputInterval_Int(int keyboardInputInterval)
{
	QApplication::setKeyboardInputInterval(keyboardInputInterval);
}

void QApplication_SetStartDragDistance_Int(int l)
{
	QApplication::setStartDragDistance(l);
}

void QApplication_SetStartDragTime_Int(int ms)
{
	QApplication::setStartDragTime(ms);
}

void QApplication_SetWheelScrollLines_Int(int wheelScrollLines)
{
	QApplication::setWheelScrollLines(wheelScrollLines);
}

int QApplication_StartDragDistance()
{
	return QApplication::startDragDistance();
}

int QApplication_StartDragTime()
{
	return QApplication::startDragTime();
}

QtObjectPtr QApplication_TopLevelAt_Int_Int(int x, int y)
{
	return QApplication::topLevelAt(x, y);
}

int QApplication_WheelScrollLines()
{
	return QApplication::wheelScrollLines();
}

QtObjectPtr QApplication_WidgetAt_Int_Int(int x, int y)
{
	return QApplication::widgetAt(x, y);
}

