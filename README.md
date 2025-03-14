# SmartTable 
## Проблема
При посещении кафе или ресторанов, часто возникают неудобства, связанные с заказом и оплатой. Посетителям приходится долго ждать, чтобы получить меню или оформить заказ, а в часы пик — стоять в очереди или искать свободного официанта. Когда наконец удается заказать, возникает необходимость держать в уме сумму заказа, чтобы не выйти за рамки бюджета. Если посетители отдыхают компанией, в конце трапезы нужно вручную рассчитывать, кто сколько должен заплатить. Также часто бывает сложно подозвать официанта для дополнительного заказа, просьбы убрать со стола или принести счет. В условиях высокой загруженности персонала заведения зачастую не могут обеспечить своевременное обслуживание каждого посетителя, что снижает общий уровень комфорта и удобства.
## Идея
Проект представляет собой комплексную платформу для упрощения процесса заказа и обслуживания в ресторанах. 
Он состоит из двух основных компонентов:
- Веб-админка для ресторанов
- Telegram-бот для посетителей  
### Основные функции системы
#### Административная панель для ресторанов
- Регистрация и авторизация: Рестораны могут создавать аккаунты, регистрироваться в системе и настраивать профиль ресторана.  
- Создание и управление рестораном: В админке ресторан может добавить название, описание, адрес и контактные данные.  
- Загрузка меню: Меню загружается в систему с фотографиями блюд, ценами и описаниями. Меню может редактироваться в любой момент.  
- Управление заказами: В админке видны все текущие заказы с указанием столов и статусов. Заказы можно помечать как оплаченные, после чего сессия закрывается.  
#### Telegram-бот для посетителей
Посетитель открывает телеграмм мини-приложение для доступа к меню ресторана.  
- Просмотр меню и заказ: Через мини-приложение посетители могут просматривать меню, добавлять блюда в корзину и отправлять заказы.  
- Вызов официанта: В приложении предусмотрена функция вызова официанта, если требуется помощь.  
- Система комнат для групповых заказов: Несколько человек могут использовать один идентификатор для создания общей комнаты, где они смогут совместно просматривать меню, добавлять заказы и видеть, сколько каждый должен оплатить.  
- Разделение счета: В конце заказа приложение покажет общую сумму и предложит распределить оплату между участниками.  
### Основные преимущества проекта
- Упрощение процесса заказа: Использование приложения исключает необходимость предоставления бумажных меню и ускоряет процесс заказа.  
- Удобство для ресторанов: Админка позволяет ресторанам легко управлять меню, отслеживать заказы и вести учет.  
- Гибкость для посетителей: Возможность делать заказ без необходимости скачивания отдельных приложений или ожидания официанта.  
- Прозрачность оплаты: Система разделения счета позволяет удобно распределить расходы среди компании друзей.  
### Целевая аудитория
- Рестораны и кафе: Компании, которые хотят улучшить процесс обслуживания и ускорить обработку заказов.  
- Посетители ресторанов: Люди, которые ценят удобство и скорость заказа, а также возможность легко делить счет с друзьями.
##  Цель 
Автоматизировать процесс обслуживания, сделав его удобным как для ресторанов, так и для клиентов, минимизируя затраты на печатные меню и увеличивая удовлетворенность посетителей за счет быстрого и бесконтактного заказа.
## Задачи проекта
- Разработка административной панели для ресторанов. 
- Разработка Telegram-бота для посетителей. 
- Настроить взаимодействие между административной панелью и Telegram-ботом через REST API. 
- Разработать удобный и интуитивно понятный интерфейс админки для ресторанов. 
- Сделать Telegram-бот максимально простым и доступным для всех категорий пользователей.
##  Планируемые результаты
- Создание удобной платформы для ресторанов  
- Telegram-бот, который позволяет посетителям ресторанов легко делать заказ без необходимости скачивания отдельного приложения.  
- Мини-приложение в боте для просмотра меню, создания заказов и вызова официанта.  
- Реализованная функция разделения счета, упрощающая расчет для компаний друзей.  
- Уведомления о статусе заказа (подтвержден, в обработке, готов) через Telegram-бот.  
- Простая и понятная админка, разработанная с учетом всех потребностей ресторана.

## Идеи выходящие за рамки MVP
- Введение в админке аккаунтов сотрудников заведения для ведения статистики
- Бронирование столов
