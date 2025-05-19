INSERT INTO "public"."users"("first_name","last_name","email","password","access_level","created_at","updated_at")
VALUES
(E'Trevor',E'Sawler',E'admin@admin.com',E'$2a$12$Wm8SHtNb7v9oRF6RmPP/c.PHE5tERA6mAfvShxcWJWT7i5nwXg94i',3,E'2020-12-04 00:00:00',E'2020-12-04 00:00:00');

INSERT INTO public.rooms (room_name,created_at,updated_at) VALUES
('General''s Quarters','2020-11-18 00:00:00.000','2020-11-18 00:00:00.000'),
('Major''s Suite','2020-11-19 00:00:00.000','0202-11-19 00:00:00.000');

INSERT INTO public.restrictions (restriction_name,created_at,updated_at) VALUES
('Reservation','2020-11-18 00:00:00.000','2020-11-18 00:00:00.000'),
('Owner Block','2020-11-19 00:00:00.000','2020-11-19 00:00:00.000');