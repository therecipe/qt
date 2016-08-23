// +build sailfish

#define protected public
#define private public

#include "sailfish_sailfish.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QList>
#include <QString>
#include <QUrl>
#include <sailfishapp.h>

void* SailfishApp_SailfishApp_Application(int argc, char* argv)
{
	QList<QByteArray> aList = QByteArray(argv).split('|');
	char *argvs[argc];
	static int argcs = argc;
	for (int i = 0; i < argc; i++)
	argvs[i] = const_cast<char*>(aList[i].constData());

	return SailfishApp::application(argcs, argvs);
}

int SailfishApp_SailfishApp_Main(int argc, char* argv)
{
	QList<QByteArray> aList = QByteArray(argv).split('|');
	char *argvs[argc];
	static int argcs = argc;
	for (int i = 0; i < argc; i++)
	argvs[i] = const_cast<char*>(aList[i].constData());

	return SailfishApp::main(argcs, argvs);
}

void* SailfishApp_SailfishApp_CreateView()
{
	return SailfishApp::createView();
}

void* SailfishApp_SailfishApp_PathTo(char* filename)
{
	return new QUrl(SailfishApp::pathTo(*(new QString(filename))));
}

