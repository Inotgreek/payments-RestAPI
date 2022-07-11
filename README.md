# REST API PAYMENTS
Для начала необходимо создать сервер в Postgres. В терминале директории проекта произвести миграцию при помощи:
(//migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:портсервера/postgres?sslmode=disable' up)
После этого можно делать запросы к серверу.
Для экономии времени я сделал скриншоты запросов и добавлю к ним описание.

## ЗАПРОСЫ

### Создание платежа. 
Передаю в post-запросе все необходимые данные, и получаю ответ о создании платежа.
По заданию, нужно было реализовать присвоение статусов "НОВЫЙ" и "ОШИБКА".
Для реализации я сделал проверку на сумму платежа. Если она больше нуля, то статус "НОВЫЙ". Если меньше или равна, то "ОШИБКА".

____  
####                                           1) Создание платежа со статусом "НОВЫЙ"
![post-positive](https://user-images.githubusercontent.com/98470862/174496650-2d639128-f294-49ad-9298-c996bdeda800.PNG)
![post-positiveDB](https://user-images.githubusercontent.com/98470862/174496657-4669336f-352d-459d-bbce-225ded98bfae.PNG)

____                                                                                        
####                                           2) Создание платежа со статусом "ОШИБКА" из-за нулевой суммы
                                               
![post-negative](https://user-images.githubusercontent.com/98470862/174496664-f0fa59b1-57c0-4bee-aa1d-2363b52b78a7.PNG)
![post-negativeDB](https://user-images.githubusercontent.com/98470862/174496673-11e08a2b-97e9-4b62-a951-3858940d033b.PNG)

____  

### Меняю статус платежа.
Поменять статус платежа возможно в случае, если изначальный статус "НОВЫЙ". Так же, при изменении статуса, меняется дата последнего изменения.
Использую платеж из предыдущего пункта.
____                                           
####                                            1)Изменение изначального статуса "НОВЫЙ"

![changestatus-positive](https://user-images.githubusercontent.com/98470862/174496818-95ed1e1a-a24b-43c8-af98-607cea9ee42e.PNG)
![changestatus-positiveDB](https://user-images.githubusercontent.com/98470862/174496824-aba51005-adaa-46db-8353-bf11b21d7e7c.PNG)
____
####                                            2)Неудачно изменение статуса


![changestatus-negative](https://user-images.githubusercontent.com/98470862/174496853-33d0a3e1-ccb6-4502-9ad5-eae871cf5ab6.PNG)
![changestatus-negativeDB](https://user-images.githubusercontent.com/98470862/174496861-edeff849-15cb-422a-baf8-2a9f31a498fc.PNG)
____
### Проверяю статус платежа по ID.
Get-запрос с указанием ID. В ответе получаю статус.
####                                           Get-запрос статуса     
![getstatus](https://user-images.githubusercontent.com/98470862/174497078-16af3710-79cb-4a19-bcd7-8bf177e90b49.PNG)
____
### Получение всех платежей по email.
Отправляю запрос с email, который был указан в POST-запросах.
####    Get-запрос платежей по email
![paymentsbyemail](https://user-images.githubusercontent.com/98470862/174497163-9cedf725-1e32-48ae-91db-a839fec43773.PNG)
____
### Отмена платежа
Отправляю DELETE-запрос для отмены платежей(удаления из бд). Невозможно отменить те платежи, который со статусом "ОШИБКА" или "УСПЕХ"

#### 1) Положительный запрос отмены
![cancel-positive](https://user-images.githubusercontent.com/98470862/174497291-b5660a17-03e6-4f8f-9739-fab6eac3a6b2.PNG)
____
#### 2) Отрицательный запрос отмены
![cancel-negative](https://user-images.githubusercontent.com/98470862/174497313-c4f75dfc-667c-4363-93ce-0223e3846fb4.PNG)
____
#### 3)Результат в БД
![cancel-result](https://user-images.githubusercontent.com/98470862/174497335-b7191e4e-5b00-4c6f-8fcb-38d80b0d90b5.PNG)
