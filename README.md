# OZON_COMMENT
OZON

Реализовать систему для добавления и чтения постов и комментариев с использованием GraphQL, аналогичную комментариям к постам на популярных платформах, таких как Хабр или Reddit.
Характеристики системы постов:

    Можно просмотреть список постов
    Можно просмотреть пост и комментарии под ним
    Пользователь, написавший пост, может запретить оставление комментариев к своему посту

Характеристики системы комментариев к постам:

    Система пагинации для получения списка комментариев
    Длина текста комментария ограничена до, например, 2000 символов
    Комментарии организованы иерархически, позволяя вложенность без ограничений

Дополнительные требования для реализации через GraphQL Subscriptions:

    Комментарии к постам должны доставляться асинхронно, т.е. клиенты, подписанные на определенный пост, должны получать уведомления о новых комментариях без необходимости повторного запроса.

Требования к реализации:

    Система должна быть написана на языке Go.
    Покрытие реализованного функционала unit-тестами.
    Использование Docker для распространения сервиса в виде Docker-образа.
    Хранение данных может быть как в памяти (in-memory), так и в PostgreSQL. Выбор хранилища должен быть определяемым параметром при запуске сервиса.

Критерии оценки:

    Обработка ошибок в различных сценариях использования.
    Удобство и логичность использования системы комментариев.
    Качество и чистота кода, структура проекта и распределение файлов по пакетам.
    Как хранятся комментарии и как организована таблица в базе данных/in-memory, включая механизм пагинации.
    Эффективность работы системы при множественном одновременном использовании, сравнимая с популярными сервисами, такими как Хабр.
    В реализации учитываются возможные проблемы с производительностью, такие как проблемы с n+1 запросами и большая вложенность комментариев.
