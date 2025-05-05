#include "QtEvent.h"
#include <QtWidgets/QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QtEvent w;
    w.show();
    return a.exec();
}
