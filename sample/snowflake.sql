create table generated_images
(
    id      bigint auto_increment
        primary key,
    user_id bigint                               not null,
    prompt  mediumtext                           not null,
    hash    mediumtext                           not null,
    created datetime default current_timestamp() not null,
    updated datetime default current_timestamp() not null,
    deleted datetime                             null
);

create table lovers
(
    id       bigint auto_increment
        primary key,
    userId   bigint                                 not null,
    name     tinytext                               not null,
    nickname tinytext                               null,
    sex      tinyint(1) default 1                   not null,
    age      tinyint    default 18                  not null,
    hair     tinytext                               not null,
    face     tinytext                               not null,
    eyes     tinytext                               not null,
    nose     tinytext                               not null,
    mouth    tinytext                               not null,
    ears     tinytext                               not null,
    body     tinytext                               not null,
    breast   tinytext                               not null,
    `rank`   tinytext   default 'normal'            null,
    level    tinyint    default 1                   not null,
    exp      bigint     default 0                   not null,
    remarks  mediumtext                             null,
    isNft    tinyint(1) default 0                   not null,
    created  datetime   default current_timestamp() not null,
    updated  datetime   default current_timestamp() not null,
    deleted  datetime                               null,
    race     tinytext   default 'human'             not null
)
    comment '연인';

create table users
(
    id      bigint auto_increment
        primary key,
    uid     tinytext                               not null,
    name    tinytext                               not null,
    sex     tinyint(1) default 0                   not null,
    created datetime   default current_timestamp() not null,
    updated datetime   default current_timestamp() not null on update current_timestamp(),
    deleted datetime                               null
);
