logger
=======

motain/logger is a general purpose logger library.

You can log to standard GOlang log, syslog, AMQP all at once or you can choose
which of them you want.

You don't need to change the anything in your code anymore in order to use one
or another logging facility.


Note
===

The AMQP facility is tested for AMQP only and it's built do use our own fork
of [streadway/amqp](https://github.com/streadway/amqp).
