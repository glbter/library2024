-- Add migration script here
create table if not exists comments
(
    id          uuid not null primary key,
    response_to uuid,
    user_id     bigint not null,
    book_id     bigint not null,
    text        text not null,

    foreign key (response_to) references comments (id) on update cascade on delete cascade,
    foreign key (user_id) references users (id) on update cascade on delete cascade,
    foreign key (book_id) references books (id) on update cascade on delete cascade
);