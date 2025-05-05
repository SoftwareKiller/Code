#include "QtSmallWidget.h"
#include <QtWidgets/QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QtSmallWidget w;
    w.show();
    return a.exec();
}
