# yafu - yet another factorial utility in Go

- Скачайте последнюю версию [↓ YAFU.](http://files.matveynator.ru/yafu/latest/)

> Поддерживаемые операционные системы: [Linix](http://files.matveynator.ru/yafu/latest/linux), [Windows](http://files.matveynator.ru/yafu/latest/windows), [Android](http://files.matveynator.ru/yafu/latest/android), [Mac](http://files.matveynator.ru/yafu/latest/mac), [IOS](http://files.matveynator.ru/yafu/latest/ios), [FreeBSD](http://files.matveynator.ru/yafu/latest/freebsd), [DragonflyBSD](http://files.matveynator.ru/yafu/latest/dragonfly), [OpenBSD](http://files.matveynator.ru/yafu/latest/openbsd), [NetBSD](http://files.matveynator.ru/yafu/latest/netbsd), [Plan9](http://files.matveynator.ru/yafu/latest/plan9), [AIX](http://files.matveynator.ru/yafu/latest/aix), [Solaris](http://files.matveynator.ru/yafu/latest/solaris), [Illumos](http://files.matveynator.ru/yafu/latest/illumos)

- Download latest version of [↓ YAFU.](http://files.matveynator.ru/yafu/latest/)


### Вспомогательные конфигурационные опции:

```
yafu -h

-file string
    	путь к файлу с кандидатами
```


### Как работает программа:

1. **Чтение Кандидатов**
- Прочитать числа (кандидаты) из текстового файла, путь к которому передается программе через флаг `-file`.

2. **Работа с Большими Числами**
- Учесть, что все кандидаты являются очень большими числами, и использовать тип `big.Int` для их представления и обработки.

3. **Параллельные Вычисления**
- Использовать параллельные вычисления для обработки кандидатов.
- Использовать все доступные ядра процессора для выполнения задач параллельно.
- Разработать механизм, который будет раздавать задания различным исполнителям (workers) и синхронизировать их работу.

4. **Каналы и Select**
- Использовать каналы для передачи заданий и результатов между горутинами.
- Использовать оператор `select` для обработки нескольких каналов.

5. **Алгоритм Проверки Простых Чисел**
- Реализовать функцию, которая будет проверять, является ли число простым.
- Использовать эту функцию для вычисления и вывода простых чисел из списка кандидатов.

6. **Блокировки и Синхронизация**
- Использовать блокировки и синхронизацию для обеспечения корректной работы программы в условиях параллельного выполнения.
- Не использовать мьютексы; вместо этого применять каналы и `select` для синхронизации доступа к ресурсам.

7. **Логирование Результатов**
- Вывести результаты обработки каждого кандидата, указав, является ли он простым числом.

8. **Обработка Ошибок**
- Корректно обрабатывать возможные ошибки, такие как проблемы с чтением файла и некорректные входные данные.
