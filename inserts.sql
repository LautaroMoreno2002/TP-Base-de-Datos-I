-- inserts de alumne

insert into alumne values (1, 'Ken', 'Thompson', 5153057,
	'1995-05-05', '15-2889-7948', 'ken@thompson.org');

insert into alumne values (2, 'Dennis', 'Ritchie', 25610126,
	'1955-04-11', '15-7811-5045', 'dennis@ritchie.org');

insert into alumne values (3, 'Donald', 'Knuth', 9168297,
	'1984-04-05', '15-2780-6005', 'don@knuth.org' );

insert into alumne values (4, 'Rob', 'Pike', 4915593, 
	'1946-08-16', '15-1114-9719', 'rob@pike.org');

insert into alumne values (5,'Douglas','McIlroy',33187055,
	'1939-06-09','15-9625-0245','douglas@mcilroy.org');

insert into alumne values (6,'Brian','Kernighan',13897948,
	'1992-11-22','15-6410-6066','brian@kernighan.org');

insert into alumne values (7,'Bill','Joy',34115045,
	'1954-02-04','15-4215-8655','bill@joy.org');

insert into alumne values (8,'Marshal Kirk','McKusick',9806005,
	'1955.12.27','15-5197-4379','marshall_kirk@mckusick.org');

insert into alumne values (9,'Theo','de Raadt',5149719,
	'1950-02-07','15-6470-9444','theo@deraadt.org');

insert into alumne values (10,'Cristina','Kirchner',6250245,
	'1990-08-17','15-5291-0113','cfk@fpv.gov.ar');

insert into alumne values (11,'Diego','Maradona',19158655,
	'1985-02-27','15-3361-4854','diego@dios.com.ar');

insert into alumne values (12,'Martin','Palermo',5974379,
	'1918-06-09','15-9877-3169','martin@palermo.com.ar');

insert into alumne values (13,'Guillermo','Barros Schelotto'
	,3910113,'1982-05-03','15-5020-5695','guille@melli.com.ar');

insert into alumne values (14,'Susú','Pecoraro',7547862,
	'1935-04-03','15-6695-9505','susu@pecoraro.com.ar');

insert into alumne values (15,'Norma','Aleandro',26614854,
	'1992-03-18','15-9155-4115','norma@aleandro.com.ar');

insert into alumne values (16,'Soledad','Silveyra',7773169,
	'1957-07-28','15-9184-4522','sole@silveyra.com.ar');

insert into alumne values (17,'Libertad','Lamarque',32205695,
	'1971-03-07','15-6363-9690','libertad@lamarque.com.ar');

insert into alumne values (18,'Ana Maria','Picchio',19020903,
	'1946-08-06','15-4819-2117','ana.maria@picchio.com.ar');

insert into alumne values (19,'Niní','Marshall',10535508,
	'1951-09-07','15-9799-6045','nini@marshall.com');

insert into alumne values (20,'Claudia','Lapacó',30934609,
	'1961-08-03','15-2005-4879','claudia@lapaco.com.ar');

-- inserts de materia
insert into materia values (1, 'Taller Inicial Común: Taller de Lectura y Escritura');

insert into materia values (2, 'Taller Inicial Orientado: Ciencias Exactas');

insert into materia values (3, 'Taller Inicial Obligatorio del Área de Matemática');

insert into materia values (4, 'Introducción a la Programación');

insert into materia values (5, 'Taller de Lectura y Escritura en las Disciplinas');

insert into materia values (6, 'Introducción a la Matemática');

insert into materia values (7, 'Programación I');

insert into materia values (8, 'Organización del Computador');

insert into materia values (9, 'Inglés Lectocompresión I');

insert into materia values (10, 'Programación II');

insert into materia values (11, 'Sistemas Operativos y Redes');

insert into materia values (12, 'Lógica y Teoría de Números');

insert into materia values (13, 'Programación III');

insert into materia values (14, 'Problemas Socioeconómicos Contemporáneos');

insert into materia values (15, 'Inglés Lectocomprensión II');

insert into materia values (16, 'Gestión y Administración de Bases de Datos');

insert into materia values (17, 'Matemática Discreta');

insert into materia values (18, 'Inglés Lectocomprensión III');

insert into materia values (19, 'Ingeniería de Software');

insert into materia values (20, 'Laboratorio de Construcción de Software');

insert into materia values (21, 'Especificación de Software');

-- inserts correlatividad

insert into correlatividad values (4, 2);

insert into correlatividad values (4, 3);

insert into correlatividad values (5, 1);

insert into correlatividad values (6, 2);

insert into correlatividad values (6, 3);

insert into correlatividad values (7, 1);

insert into correlatividad values (7, 4);

insert into correlatividad values (8, 1);

insert into correlatividad values (8, 4);

insert into correlatividad values (9, 1);

insert into correlatividad values (9, 2);

insert into correlatividad values (9, 3);

insert into correlatividad values (10, 6);

insert into correlatividad values (10, 7);

insert into correlatividad values (11, 7);

insert into correlatividad values (11, 8);

insert into correlatividad values (12, 6);

insert into correlatividad values (13, 10);

insert into correlatividad values (14, 1);

insert into correlatividad values (15, 5);

insert into correlatividad values (15, 9);

insert into correlatividad values (16, 8);

insert into correlatividad values (16, 10);

insert into correlatividad values (16, 12);

insert into correlatividad values (17, 12);

insert into correlatividad values (18, 15);

insert into correlatividad values (19, 13);

insert into correlatividad values (20, 5);

insert into correlatividad values (20, 14);

insert into correlatividad values (20, 16);

insert into correlatividad values (20, 19);

insert into correlatividad values (20, 21);

insert into correlatividad values (21, 12);

insert into correlatividad values (21, 13);

-- inserts comision

insert into comision values (1, 1, 5);

insert into comision values (2, 1, 5);

insert into comision values (3, 1, 15);

insert into comision values (4, 1, 3);

insert into comision values (4, 2, 4);

insert into comision values (4, 3, 5);

insert into comision values (5, 1, 5);

insert into comision values (6, 1, 8);

insert into comision values (7, 1, 3);

insert into comision values (7, 2, 5);

insert into comision values (8, 1, 10);

insert into comision values (9, 1, 7);

insert into comision values (10, 1, 9);

insert into comision values (11, 1, 5);

insert into comision values (12, 1, 15);

insert into comision values (13, 1, 13);

insert into comision values (14, 1, 12);

insert into comision values (15, 1, 8);

insert into comision values (16, 1, 5);

insert into comision values (17, 1, 4);

insert into comision values (18, 1, 8);

insert into comision values (19, 1 ,2);

insert into comision values (20, 1, 6);

insert into comision values (21, 1, 11);

-- inserts periodo

insert into periodo values ('2022-1', 'cerrado');

insert into periodo values ('2022-2', 'cerrado');

insert into periodo values ('2023-1', 'cerrado');

insert into periodo values ('2023-2', 'cerrado');

-- inserts historia_academia

insert into historia_academica values (1, '2023-1', 1, 1,
	'aprobada', 9, 9);

insert into historia_academica values (1, '2023-1', 2, 1,
	'aprobada', 10, 10);

insert into historia_academica values (1, '2023-1', 3, 1,
	'ausente', 0);

insert into historia_academica values (1, '2023-2', 3, 1,
	'regular', 5);

insert into historia_academica values (1, '2023-2', 5, 1,
	'aprobada', 7, 7);

insert into historia_academica values (2, '2023-1', 1, 1,
	'aprobada', 9, 9);

insert into historia_academica values (2, '2023-1', 2, 1, 
	'aprobada', 10, 10);

insert into historia_academica values (2, '2023-1', 3, 1,
	'ausente', 0);

insert into historia_academica values (2, '2023-2', 3, 1,
	'reprobada', 2);

insert into historia_academica values (2, '2023-2', 5, 1,
	'aprobada', 7, 7);

insert into historia_academica values (3, '2023-1', 1, 1,
	'aprobada', 9, 9);

insert into historia_academica values (3, '2023-1', 2, 1,
	'aprobada', 10, 10);

insert into historia_academica values (3, '2023-1', 3, 1,
	'aprobada', 10, 10);

insert into historia_academica values (3, '2023-2', 4, 2,
	'regular', 6);

insert into historia_academica values (3, '2023-2', 5, 1,
	'aprobada', 9, 9);

insert into historia_academica values (4, '2023-1', 1, 1,
	'aprobada', 9, 9);

insert into historia_academica values (4, '2023-1', 2, 1,
	'aprobada', 10, 10);

insert into historia_academica values (4, '2023-1', 3, 1,
	'aprobada', 10, 10);

insert into historia_academica values (4, '2023-2', 4, 2,
	'ausente', 0);

insert into historia_academica values (4, '2023-2', 5, 1,
	'aprobada', 9, 9);

insert into historia_academica values (5, '2022-1', 1, 1,
	'aprobada', 10, 10);

insert into historia_academica values (5, '2022-1', 2, 1,
	'aprobada', 10, 10);

insert into historia_academica values (5, '2022-1', 3, 1,
	'aprobada', 10, 10);

insert into historia_academica values (5, '2022-2', 4, 2,
	'aprobada', 9, 9);

insert into historia_academica values (5, '2022-2', 5, 1,
	'aprobada', 9, 9);

insert into historia_academica values (5, '2023-1', 6, 1,
	'regular', 5);

insert into historia_academica values (5, '2023-1', 7, 2,
	'aprobada', 8, 8);

insert into historia_academica values (5, '2023-1', 8, 1,
	'regular', 6);

insert into historia_academica values (5, '2023-2', 11, 1,
	'regular', 6);

insert into historia_academica values (6, '2023-1', 1, 1,
	'aprobada', 9, 9);

insert into historia_academica values (6, '2023-1', 2, 1,
	'aprobada', 10, 10);

insert into historia_academica values (6, '2023-1', 3, 1,
	'aprobada', 10, 10);

insert into historia_academica values (6, '2023-2', 4, 2,
	'aprobada', 10, 10);

insert into historia_academica values (6, '2023-2', 5, 1,
	'aprobada', 9, 9);

insert into historia_academica values (7, '2023-1', 1, 1,
	'aprobada', 9, 9);

insert into historia_academica values (7, '2023-1', 2, 1,
	'aprobada', 10, 10);

insert into historia_academica values (7, '2023-1', 3, 1,
	'aprobada', 10, 10);

insert into historia_academica values (7, '2023-2', 4, 2,
	'aprobada', 10, 10);

insert into historia_academica values (7, '2023-2', 5, 1,
	'regular', 6);

insert into historia_academica values (8, '2023-1', 1, 1,
	'aprobada', 9, 9);

insert into historia_academica values (8, '2023-1', 2, 1,
	'aprobada', 10, 10);

insert into historia_academica values (8, '2023-1', 3, 1,
	'aprobada', 10, 10);

insert into historia_academica values (8, '2023-2', 6, 1,
	'aprobada', 10, 10);

insert into historia_academica values (8, '2023-2', 9, 1,
	'aprobada', 8, 8);

-- inserts entrada_trx

insert into entrada_trx values (1, 'alta inscrip', null, null, 
	10, 3, 1, null);

insert into entrada_trx values (2, 'apertura', 2023, 2, null,
	null, null, null);

insert into entrada_trx values (3, 'apertura', 2024, 1, null,
	null, null, null);

insert into entrada_trx values (4, 'alta inscrip', null, null,
	1, 4, 2, null);

insert into entrada_trx values (5, 'alta inscrip', null, null,
	2, 4, 1, null);

insert into entrada_trx values (6, 'alta inscrip', null, null,
	8, 4, 2, null);

insert into entrada_trx values (7, 'alta inscrip', null, null,
	4, 7, 1, null);

insert into entrada_trx values (8, 'alta inscrip', null, null,
	21, 2, 1, null);

insert into entrada_trx values (9, 'alta inscrip', null, null,
	7, 7, 1, null);

insert into entrada_trx values (10, 'alta inscrip', null, null,
	17, 2, 1, null);

insert into entrada_trx values (11, 'alta inscrip', null, null,
	12, 2, 1, null);

insert into entrada_trx values (12, 'alta inscrip', null, null,
	13, 2, 1, null);

insert into entrada_trx values (13, 'alta inscrip', null, null,
	14, 2, 1, null);

insert into entrada_trx values (14, 'alta inscrip', null, null,
	15, 2, 1, null);

insert into entrada_trx values (15, 'alta inscrip', null, null,
	16, 2, 1, null);

insert into entrada_trx values (16, 'alta inscrip', null, null,
	11, 2, 1, null);

insert into entrada_trx values (17, 'baja inscrip', null, null,
	8, 2, null, null);

insert into entrada_trx values (18, 'baja inscrip', null, null,
	14, 2, null, null);

insert into entrada_trx values (19, 'cierre inscrip', 2024, 1,
	null, null, null, null);

insert into entrada_trx values (20, 'aplicacion cupo', 2024, 1,
	null, null, null, null);

insert into entrada_trx values (21, 'baja inscrip', null, null,
	13, 2, null, null);

insert into entrada_trx values (22, 'ingreso nota', null, null,
	1, 4, 2, 10);

insert into entrada_trx values (23, 'ingreso nota', null, null,
	8, 4, 1, 5);

insert into entrada_trx values (24, 'cierre cursada', null, null,
	null, 4, 2, null);

insert into entrada_trx values (25, 'cierre cursada', null, null,
	null, 16, 1, null);

insert into entrada_trx values (26, 'ingreso nota', null, null,
	8, 4, 2, 5);

insert into entrada_trx values (27, 'cierre cursada', null, null,
	null, 4, 2, null);

