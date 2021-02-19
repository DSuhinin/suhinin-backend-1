create table test_task_database.users
(
    id int auto_increment
        primary key,
    email varchar(255) not null,
    password text null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    constraint users_email_uindex
        unique (email)
);

