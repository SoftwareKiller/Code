#include "QtPainterEvent.h"
#include <QtWidgets/QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QtPainterEvent w;
    w.show();
    return a.exec();
}
