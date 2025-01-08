-- Borrar foreign key's
alter table if exists correlatividad drop constraint if exists cor_materia_id_fk;
alter table if exists correlatividad drop constraint if exists cor2_materia_id_fk;
alter table if exists comision drop constraint if exists com_materia_id_fk;
alter table if exists cursada drop constraint if exists cur_materia_id_fk;
alter table if exists cursada drop constraint if exists cur_alumne_id_fk;
alter table if exists historia_academica drop constraint if exists his_alumne_id_fk;
alter table if exists historia_academica drop constraint if exists semestre_id_fk;
alter table if exists historia_academica drop constraint if exists his_materia_id_fk;

-- Borrar primary key's
alter table if exists alumne drop constraint if exists alumne_pk;
alter table if exists materia drop constraint if exists materia_pk;
alter table if exists correlatividad drop constraint if exists correlatividad_pk;
alter table if exists comision drop constraint if exists comision_pk;
alter table if exists cursada drop constraint if exists cursada_pk;
alter table if exists periodo drop constraint if exists periodo_pk;
alter table if exists historia_academica drop constraint if exists historia_academica_pk;
alter table if exists error drop constraint if exists error_pk;
alter table if exists envio_email drop constraint if exists envio_email_pk;
alter table if exists entrada_trx drop constraint if exists trx;
