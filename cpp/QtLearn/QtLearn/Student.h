#pragma once
#pragma execution_character_set("utf-8")

#include <QObject>

class Student : public QObject
{
	Q_OBJECT;
public:
	explicit Student(QObject* parent = 0);

public slots:
	void treat();

	void treat(QString foodName);
};