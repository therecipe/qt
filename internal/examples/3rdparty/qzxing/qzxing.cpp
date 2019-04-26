#include "qzxing.h"

#include "qzxing/src/QZXing.h"

void QZXing_registerQMLTypes()
{
	QZXing::registerQMLTypes();
}

void QZXing_registerQMLImageProvider(void* ptr)
{
	QZXing::registerQMLImageProvider(*static_cast<QQmlEngine*>(ptr));
}
