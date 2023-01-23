CREATE TABLE "role"
(
    id                      serial  PRIMARY KEY not null unique,
    name                    varchar(50) not null
);
INSERT INTO "role" (name) values('user'); 
INSERT INTO "role" (name) values('guide'); 
INSERT INTO "role" (name) values('admin'); 

CREATE TABLE "user"
(
    id                    serial PRIMARY KEY not null unique,
    name                  varchar(25) not null,
    last_name             varchar(25),
    password_hash         varchar(255) not null,
    numbers               varchar(11) unique,
    role_id               int references "role"(id) on delete cascade,
    email                 VARCHAR(40) unique,
    sex                   boolean DEFAULT TRUE,
    registration_datetime TIMESTAMP(0) NOT NULL,
    image_src             varchar(500)
);
CREATE TABLE "house_type"
(   
    id                      serial PRIMARY KEY not null unique,
    name                    varchar(100) not null 
);
INSERT INTO "house_type" (name) values('hotel'); 

CREATE TABLE "place"
(
    id                      serial PRIMARY KEY not null unique,
    name                    varchar(255) not null,
    description             text,
    location_long           double precision NOT NULL,
    location_lat            double precision NOT NULL,
    min_price               int,
    address                 VARCHAR(100),
    numbers                 VARCHAR(11),
    mail                    VARCHAR(30),
    site_url                VARCHAR(150),
    pushkin                 boolean DEFAULT FALSE,
    start_time              date,
    house_price             INT,
    house_type_id           int references "house_type" (id) on delete cascade,
    count_room              INT DEFAULT 0, 
    square                  INT DEFAULT 0
);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,mail,site_url,house_price,house_type_id) values('hotel1','vladik alooo',0.0,0.0,'jil','8928','jilnya','ublya',1400,1);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('place1','jil',0.0,0.0,'jil','8928',TRUE,2500);  
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event','jil',0.0,0.0,'jil','8928',TRUE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event2','jil',0.0,0.0,'jil','8928',TRUE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event3','jil',0.0,0.0,'jil','8928',TRUE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event4','jil',0.0,0.0,'jil','8928',TRUE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event5','jil',0.2,1.1255,'jil','8928',TRUE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event6','jil',0.0,0.0,'jil','8928',TRUE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event7','jil',13.7,0.0,'jil','8928',TRUE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event8','jil',0.0,0.0,'jil','8928',TRUE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event9','jil',0.0,-1552.14,'jil','8928',FALSE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event10','jil',0.0,0.0,'jil','8928',FALSE,350);
INSERT INTO "place" (name,description,location_long,location_lat,address,numbers,pushkin,min_price) values('event11','jil',0.0,0.0,'jil','8928',TRUE,350);
CREATE TABLE "place_src"
(
    id                      serial PRIMARY KEY not null unique,
    image_src               varchar(500) not null,
    place_id                int references "place" (id) on delete cascade not null
);
CREATE TABLE "favorite_place"
(
    id                      serial PRIMARY KEY not null unique,
    user_id                 int references "user" (id) on delete cascade not null,
    place_id                int references "place" (id) on delete cascade not null
);
CREATE TABLE "review"
(
    id                      serial PRIMARY KEY not null unique,
    user_id                 int references "user" (id) on delete cascade not null,
    rating                  int NOT NULL,
    review_text             TEXT,
    review_datetime         timestamp NOT NULL,
    place_id                int references "place" (id) on delete cascade, 
    guide_id                int references "user" (id) on delete cascade
);
CREATE TABLE "type"
(
    id                      serial  PRIMARY KEY  not null unique,
    name                    varchar(100) not null
);
INSERT INTO "type" (name) values('housing'); 
INSERT INTO "type" (name) values('event'); 
INSERT INTO "type" (name) values('cafe'); 
CREATE TABLE "place_type"
(
    id                      serial PRIMARY KEY not null unique,
    place_id                int references "place" (id) on delete cascade not null,
    type_id                 int references "type" (id) on delete cascade not null
);
INSERT INTO "place_type" (place_id,type_id) values(1,1); 
INSERT INTO "place_type" (place_id,type_id) values(2,3); 
INSERT INTO "place_type" (place_id,type_id) values(3,2); 
INSERT INTO "place_type" (place_id,type_id) values(4,2); 
INSERT INTO "place_type" (place_id,type_id) values(5,2); 
INSERT INTO "place_type" (place_id,type_id) values(6,2); 
INSERT INTO "place_type" (place_id,type_id) values(7,2); 
INSERT INTO "place_type" (place_id,type_id) values(8,2); 
INSERT INTO "place_type" (place_id,type_id) values(9,2); 
INSERT INTO "place_type" (place_id,type_id) values(10,2); 
INSERT INTO "place_type" (place_id,type_id) values(11,2); 
INSERT INTO "place_type" (place_id,type_id) values(12,2); 
INSERT INTO "place_type" (place_id,type_id) values(13,2); 
CREATE TABLE "achieve"
(
    id                      serial PRIMARY KEY not null unique,
    name                    varchar(100) not null,
    description             TEXT not null
);
CREATE TABLE "user_achieve"
(
    id                      serial PRIMARY KEY not null unique,
    user_id                 int references "user" (id) on delete cascade not null,
    achieve_id                 int references "achieve"(id) on delete cascade not null
);
CREATE TABLE "dayy"
(
    id                      serial PRIMARY KEY not null unique,
    name                    varchar(14) not null
);
INSERT INTO "dayy" (name) values('Понедельник'); 
INSERT INTO "dayy" (name) values('Вторник'); 
INSERT INTO "dayy" (name) values('Среда'); 
INSERT INTO "dayy" (name) values('Четверг'); 
INSERT INTO "dayy" (name) values('Пятница'); 
INSERT INTO "dayy" (name) values('Суббота'); 
INSERT INTO "dayy" (name) values('Воскресенье'); 

CREATE TABLE "week"
(
    id                      serial PRIMARY KEY  not null unique,
    place_id                int references "place" (id) on delete cascade not null,
    day_id                  int references "dayy"(id) on delete cascade not null unique,
    start_work              time,
    end_work                time,
    start_timeout           time,
    end_timeout             time 
);

CREATE TABLE "tour"
(
    id                      serial PRIMARY KEY  not null unique,
    name                    varchar(40) null,
    description             TEXT null,
    user_id                 int references "user" (id) on delete cascade not null
);
CREATE TABLE "tour_place"
(
    id                      serial PRIMARY KEY not null unique,
    tour_id                 int references "tour" (id) on delete cascade not null,
    place_id                int references "place" (id) on delete cascade not null,
    -- location_long           double precision NOT NULL,
    -- location_lat            double precision NOT NULL,
    -- address                 varchar(100),
    start_tour              boolean,
    end_tour                boolean
);