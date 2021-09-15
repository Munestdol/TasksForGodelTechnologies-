CREATE TABLE ads
(
    id              serial       not null unique,
    title           varchar(255) not null,
    price           int not null,
    main_photo      varchar(255),
    photo1          varchar(255),
    photo2          varchar(255),
    photo3          varchar(255),
    description     varchar(255),
    date            timestamp
);