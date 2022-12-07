CREATE TABLE "role"
(
    id                      serial      not null unique,
    name                    varchar(50) not null
);
CREATE TABLE "user"
(
    id                    serial      not null unique,
    name                  varchar(25) not null,
    username              varchar(25) not null unique,
    password_hash         varchar(40) not null,
    numbers               varchar(11) ,
    role_id               int references role (id) on delete cascade not null,
    password              VARCHAR(40) NOT NULL,
    email                 VARCHAR(40),
    sex                   boolean NOT NULL,
    registration_datetime TIMESTAMP(0) NOT NULL
);
CREATE TABLE "house_type"
(
    id                      serial       not null unique,
    name                    varchar(100) not null
);
CREATE TABLE "place"
(
    id                      serial       not null unique,
    name                    varchar(255) not null,
    description             varchar(255),
    location_long           double precision NOT NULL,
    location_lat            double precision NOT NULL,
    min_price               int,
    address                 VARCHAR(100),
    numbers                  VARCHAR(11),
    pushkin                 boolean,
    house_price             INT,
    house_type_id           int references house_type (id) on delete cascade,
    count_room              INT,
    square                  INT
);
CREATE TABLE "favorite_place"
(
    id                      serial       not null unique,
    user_id                 int references "user" (id) on delete cascade not null,
    place_id                int references place (id) on delete cascade not null
);
CREATE TABLE "review"
(
    id                      serial       not null unique,
    user_id                 int references "user" (id) on delete cascade not null,
    rating                  int NOT NULL,
    review_text             TEXT,
    review_datetime         timestamp NOT NULL,
    place_id                int references place (id) on delete cascade, 
    guide_id                int references "user" (id) on delete cascade
);
CREATE TABLE "type"
(
    id                      serial       not null unique,
    name                    varchar(100) not null
);
CREATE TABLE "place_type"
(
    id                      serial       not null unique,
    place_id                int references place (id) on delete cascade not null,
    type_id                 int references type (id) on delete cascade not null
);
CREATE TABLE "achieve"
(
    id                      serial       not null unique,
    name                    varchar(100) not null,
    description             TEXT not null
);
CREATE TABLE "user_achieve"
(
    id                      serial       not null unique,
    user_id                 int references "user" (id) on delete cascade not null,
    achieve                 int references "achieve"(id) on delete cascade not null
);
CREATE TABLE "dayy"
(
    id                      serial not null unique,
    name                    varchar(14) not null
);
CREATE TABLE "week"
(
    id                      serial       not null unique,
    place_id                int references "place" (id) on delete cascade not null,
    day_id                  int references dayy(id) on delete cascade not null,
    start_work              time,
    end_work                time,
    start_timeout           time,
    end_timeout             time 
);

CREATE TABLE "tour"
(
    id                      serial       not null unique,
    name                    varchar(40) not null,
    description             TEXT not null,
    user_id                 int references "user" (id) on delete cascade not null
);
CREATE TABLE "tour_place"
(
    id                      serial       not null unique,
    tour_id                 int references "tour" (id) on delete cascade not null,
    place_id                int references "place" (id) on delete cascade not null,
    location_long           double precision NOT NULL,
    location_lat            double precision NOT NULL,
    address                 varchar(100),
    start_tour              boolean,
    end_tour                boolean
);