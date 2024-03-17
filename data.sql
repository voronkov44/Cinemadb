CREATE TABLE public.actor (
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              name VARCHAR(15) NOT NULL,
                              gender VARCHAR(15) NOT NULL,
                              date_of_birth VARCHAR(15) NOT NULL
);

CREATE TABLE public.film (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             name VARCHAR(150) NOT NULL,
                             description VARCHAR(1000) NOT NULL,
                             edition VARCHAR(15) NOT NULL,
                             rating VARCHAR(10) NOT NULL,
                             actor_id UUID NOT NULL,
                             CONSTRAINT actor_fk FOREIGN KEY (actor_id) REFERENCES public.actor(id)
);
INSERT INTO actor (name, gender, date_of_birth) VALUES ('Джонни Депп', 'Мужчина', '09.06.63'); -- 5d127af3-f9bd-47d9-8d06-6b3040751594
INSERT INTO actor (name, gender, date_of_birth) VALUES ('Джейсон Стейтем', 'Мужчина', '26.07.67'); --5737f939-8e16-4936-bd06-fc15877264b0
INSERT INTO actor (name, gender, date_of_birth) VALUES ('Анджелина Джоли', 'Женщина', '04.06.75'); -- d45b3434-0621-479a-a74f-7614769e19f0


INSERT INTO film (name, description, edition, rating, actor_id) VALUES ('Пираты Карибского моря', '"Пираты Карибского моря" - это захватывающий приключенческий фильм о капитане Джеке Воробье, смелом и хитроумном пирате.', '09.07.2003', '9', '5d127af3-f9bd-47d9-8d06-6b3040751594');
INSERT INTO film (name, description, edition, rating, actor_id) VALUES ('Перевозчик', '"Перевозчик" - это боевик о Френке Мартине, талантливом водителе, который специализируется на опасных и нелегальных перевозках.', '11.10.2002', '8', '5737f939-8e16-4936-bd06-fc15877264b0');
INSERT INTO film (name, description, edition, rating, actor_id) VALUES ('Мистер и Миссис Смит', '"Мистер и Миссис Смит" - это комедийный боевик о супружеской паре, которая случайно обнаруживает, что оба являются наемными убийцами, работающими на различные агентства.', '16.06.2005', '8', 'd45b3434-0621-479a-a74f-7614769e19f0');

TRUNCATE actor CASCADE;

