CREATE TABLE payments(
                         id          serial not null unique ,
                         user_id     int not null ,
                         email  varchar(50) not null,
                         sum         numeric not null ,
                         value       varchar(50) not null,
                         create_date timestamp,
                         last_change timestamp,
                         status      varchar(50) not null

)