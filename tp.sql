-- TP
drop database if exists lozano_moreno_schaab_vallejos_db1;
create database lozano_moreno_schaab_vallejos_db1;

-- Ingreso en la base de datos
\c lozano_moreno_schaab_vallejos_db1

-- Crear tablas
create table alumne(
	id_alumne int,
	nombre text,
	apellido text,
	dni int,
	fecha_nacimento date,
	telefono char(12),
	email text --válido
);

create table materia(
	id_materia int,
	nombre text
);

create table correlatividad(
	id_materia int,
	id_mat_correlativa int
);

create table comision(
	id_materia int,
	id_comision int, --1, 2, 3,... por cada materia
	cupo int --máxima cantidad de alumnes que pueden cursar
);

create table cursada(
	id_materia int,
	id_alumne int,
	id_comision int,
	f_inscripcion timestamp,
	nota int, --inicialmente en null: no hay nota
	estado char(12) --`ingresade',`aceptade',`en espera',`dade de baja'
);

create table periodo(
	semestre text, --ejemplo: `2024-1'
	estado char(15) --`inscripcion',`cierre inscrip',`cursada',`cerrado'
);

create table historia_academica(
	id_alumne int,
	semestre text,
	id_materia int,
	id_comision int,
	estado char(15), --`ausente',`reprobada',`regular',`aprobada'
	nota_regular int,
	nota_final int
);

create table error(
	id_error int,
	operacion char(15),
--`apertura',`alta inscrip',`baja inscrip',`cierre inscrip',
--`aplicacion cupo',`ingreso nota',`cierre cursada'
	semestre text,
	id_alumne int,
	id_materia int,
	id_comision int,
	f_error timestamp,
	motivo varchar(80)
);

create table envio_email(
	id_email int,
	f_generacion timestamp,
	email_alumne text,
	asunto text,
	cuerpo text,
	f_envio timestamp,
	estado char(10) --`pendiente', `enviado'
);

-- Esta tabla *no* es parte del modelo de datos, pero se incluye para
-- poder probar las funciones.
create table entrada_trx(
	id_orden int, --en qué orden se ejecutarán las transacciones
	operacion char(15),
--`apertura',`alta inscrip',`baja inscrip',`cierre inscrip',
--`aplicacion cupo',`ingreso nota',`cierre cursada'
	año int,
	nro_semestre int,
	id_alumne int,
	id_materia int,
	id_comision int,
	nota int
);

-- Creaciòn de primary key's
alter table alumne add constraint alumne_pk primary key (id_alumne);
alter table materia add constraint materia_pk primary key (id_materia);
alter table correlatividad add constraint correlatividad_pk primary key (id_materia,id_mat_correlativa);
alter table comision add constraint comision_pk primary key (id_materia,id_comision);
alter table cursada add constraint cursada_pk primary key (id_materia,id_alumne);
alter table periodo add constraint periodo_pk primary key (semestre);
alter table historia_academica add constraint historia_academica_pk primary key (id_alumne,id_materia,semestre);
alter table error add constraint error_pk primary key (id_error);
alter table envio_email add constraint envio_email_pk primary key (id_email);
alter table entrada_trx add constraint trx primary key (id_orden);
-- Creaciòn de foreign key's
alter table correlatividad add constraint cor_materia_id_fk foreign key (id_materia) references materia(id_materia);
alter table correlatividad add constraint cor2_materia_id_fk foreign key (id_mat_correlativa) references materia(id_materia);
alter table comision add constraint com_materia_id_fk foreign key (id_materia) references materia(id_materia);
alter table cursada add constraint cur_materia_id_fk foreign key (id_materia) references materia(id_materia);
alter table cursada add constraint cur_alumne_id_fk foreign key (id_alumne) references alumne(id_alumne);
alter table historia_academica add constraint his_alumne_id_fk foreign key (id_alumne) references alumne(id_alumne);
alter table historia_academica add constraint semestre_id_fk foreign key (semestre) references periodo(semestre);
alter table historia_academica add constraint his_materia_id_fk foreign key (id_materia) references materia(id_materia);
