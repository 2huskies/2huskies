INSERT INTO public.specialty(
	code, name)
	VALUES ('240801', 'Машины и аппараты химических производств'),
	('260601', 'Технология швейных изделий'),
	('260902', 'Конструирование швейных изделий');
	
INSERT INTO public.abiturient(
	first_name, last_last, birth_date, birth_place, address, phone_number, middle_name)
	VALUES ('Иван', 'Иванов', (to_date('1963-09-01', 'YYYY-MM-DD')), 'Москва', 'Московская обл., г.Балашиха, ул.Кирова, д.20, кв.70', '+78545896584', 'Ионович'),
	('Петров', 'Пётр', (to_date('1980-09-02', 'YYYY-MM-DD')), 'Иваново', 'г.Санкт Петербург, пр.Невский, д.1а, кв.2', '+78540000000', ''),
	('Алексеев', 'Николай', (to_date('1970-01-01', 'YYYY-MM-DD')), 'Нью Йорк', 'Нижегородская обл., г.Арзамас, ул.Иванова, д.1, кв.1', '+78545896522', 'Аглы');
	
