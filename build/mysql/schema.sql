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

create table test_task_database.tokens
(
    id int auto_increment
        primary key,
    user_id int not null,
    token text not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    constraint tokens_users_id_fk
        foreign key (user_id) references test_task_database.users (id)
            on update cascade on delete cascade
);

