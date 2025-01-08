--procedimientos

--función auxiliar para el conteo de errores
create or replace function conteo_de_errores() returns error.id_error%type as $$
declare
	cant_errores int;
	last_id_error error.id_error%type;
begin
	select count(*) into cant_errores from error;
	if cant_errores = 0 then
		last_id_error := 0;
	else
		select max(id_error) into last_id_error from error;
		last_id_error := last_id_error + 1;
	end if;
	return last_id_error;
end;
$$ language plpgsql;

--apertura de inscripcion
create or replace function apertura_de_inscripcion(_año int, _n_semestre int) returns boolean as $$
declare
	año_actual int;
	semestre_existente periodo%rowtype;
	semestre_actual periodo.semestre%type;
	last_id_error error.id_error%type;
begin

	select conteo_de_errores() into last_id_error;
	
	select into año_actual extract(year from current_timestamp);
	select p.semestre into semestre_actual from periodo p where (p.estado = 'inscripcion' or p.estado = 'cierre inscrip') and p.semestre != _año || '-' || _n_semestre;
	if exists(select * from periodo where semestre=_año || '-' || _n_semestre and estado!='cierre inscrip') then
		select * into semestre_existente from periodo p where p.semestre = _año || '-' || _n_semestre;
		insert into error values (last_id_error, 'apertura', null, null, null, null,current_date, '?no es posible reabrir la inscripción del período, estado actual:' || semestre_existente.estado);
		return false;
	elsif _n_semestre < 1 or _n_semestre > 2 then	
		insert into error values (last_id_error, 'apertura', null, null, null, null,current_date, '?número de semestre no válido.');
		return false;
	elsif exists(select * from periodo where semestre!=_año || '-' || _n_semestre and (estado='inscripcion' or estado='cierre inscrip')) then
		insert into error values (last_id_error, 'apertura', null, null, null, null,current_date, '?no es posible abrir otro período de inscripción, período actual:' || semestre_actual);
		return false;
	elsif _año < año_actual then
		insert into error values (last_id_error, 'apertura', null, null, null, null, current_date, '?no se permiten inscripciones para un período anterior.');
		return false;
	end if;

	if exists (select * from periodo where semestre = _año || '-' || _n_semestre and estado = 'cierre inscrip') then
		update periodo set estado = 'inscripcion' where semestre = _año || '-' || _n_semestre;
		return true; -- Actualizo el estado
	else
		insert into periodo values (_año || '-' || _n_semestre, 'inscripcion');
		return true; -- Logrò la inscripciòn al periodo
	end if;
end;
$$ language plpgsql;

--inscripción a materia
create or replace function inscripcion_a_materia(_id_alumne int, _id_materia int, _id_comision int) returns boolean as $$
declare
	last_id_error error.id_error%type;
	cumple_correlativas boolean; --acumulador booleano
	v1 correlatividad%rowtype; -- valor para el for
begin

	select conteo_de_errores() into last_id_error;
	
	cumple_correlativas:= true;
	for v1 in select * from correlatividad where id_materia=_id_materia loop
		cumple_correlativas:=cumple_correlativas and (exists(select id_materia from historia_academica where id_materia=v1.id_mat_correlativa and id_alumne=_id_alumne and(estado='aprobada' or estado='regular')));
	end loop;
	
	if not exists(select estado from periodo where estado='inscripcion')  then
		insert into error values (last_id_error, 'alta inscrip',null, _id_alumne, _id_materia, _id_comision,current_date, '?período de inscripción cerrado.');
		return false;
	elsif _id_alumne not in (select id_alumne from alumne) then 
		insert into error values (last_id_error, 'alta inscrip', null, _id_alumne, _id_materia, _id_comision,current_date, '?id de alumne no válido.');
		return false;
	elsif _id_materia not in (select id_materia from materia) then
		insert into error values (last_id_error, 'alta inscrip', null, _id_alumne, _id_materia, _id_comision,current_date, '?id  de materia no válido.');
		return false;
	elsif _id_comision not in(select id_comision from comision where id_materia=_id_materia)then
		insert into error values (last_id_error, 'alta inscrip', null, _id_alumne, _id_materia, _id_comision,current_date, '?id de comisión no válido para la materia.');
		return false;
	elsif _id_alumne in(select id_alumne from cursada where id_materia=_id_materia) then
		insert into error values (last_id_error, 'alta inscrip', null, _id_alumne, _id_materia, _id_comision,current_date, '?alumne ya inscripte en la materia.');
		return false;
	elsif not cumple_correlativas then
		insert into error values (last_id_error, 'alta inscrip', null, _id_alumne, _id_materia, _id_comision,current_date, '?alumne no cumple requisitos de correlatividad');
		return false;
	else
		insert into cursada values(_id_materia,_id_alumne,_id_comision,current_timestamp,null,'ingresade');
		return true;
	end if;
end;
$$ language plpgsql;

--baja de inscripción
create or replace function baja_de_inscripcion(_id_alumne alumne.id_alumne%type, _id_materia materia.id_materia%type) returns boolean as $$
declare
    comision_materia cursada.id_comision%type;
    id_alumne_en_espera cursada.id_alumne%type;
	last_id_error error.id_error%type;
begin

   select conteo_de_errores() into last_id_error;

    if _id_alumne not in (select a.id_alumne from alumne a) then
        insert into error values (last_id_error, 'baja inscrip', null, _id_alumne, _id_materia, null, current_date, '?id de alumne no válido.');
        return false;
    elsif _id_materia not in (select m.id_materia from materia m) then
        insert into error values (last_id_error, 'baja inscrip', null, _id_alumne, _id_materia, null, current_date, '?id de materia no válido.');
        return false;
    end if;

    if not exists (select * from periodo p where p.estado = 'inscripcion' or p.estado = 'cursada') then
        insert into error values (last_id_error, 'baja inscrip', null, _id_alumne, _id_materia, null, current_date, '?no se permiten bajas en este período.');
        return false;
    end if;

    if _id_alumne not in (select c.id_alumne from cursada c where c.id_materia = _id_materia and c.id_alumne = _id_alumne and(c.estado= 'ingresade'or c.estado = 'aceptade')) then
        insert into error values (last_id_error, 'baja inscrip', null, _id_alumne, _id_materia, null, current_date, '?alumne no inscripte en la materia.');
        return false;
    end if;
    
    update cursada c set estado = 'dade de baja' where c.id_alumne = _id_alumne and c.id_materia = _id_materia;

    if exists(select from periodo where estado='cursada') then
        select id_comision into comision_materia from cursada c where c.id_alumne = _id_alumne and c.id_materia = _id_materia;
        if exists(select id_alumne from cursada where estado = 'en espera' and id_alumne != _id_alumne and id_materia = _id_materia and id_comision = comision_materia )then
            update cursada set estado = 'aceptade' where f_inscripcion in (select min(f_inscripcion) from cursada where id_alumne!=_id_alumne and id_materia=id_materia and id_comision=(select id_comision from cursada where id_alumne=_id_alumne and id_materia=id_materia) and estado='en espera');
       end if;
    end if;
    return true; -- Dar de baja inscripcion exitoso
end;
$$ language plpgsql;

--cierre de inscripción
create or replace function cierre_de_inscripcion(_año int, _n_semestre int) returns boolean as $$
declare 
	last_id_error error.id_error%type;
	semestre_ab periodo%rowtype;
	
begin

	select conteo_de_errores() into last_id_error;
	
	select * into semestre_ab from periodo where semestre = _año || '-' || _n_semestre and estado='inscripcion';
	if not found then
		insert into error values (last_id_error, 'cierre inscrip', _año || '-' || _n_semestre, null, null, null, current_date, '?el semestre no se encuentra en estado de inscripción.');
        return false;
    else
		update periodo set estado='cierre inscrip' where semestre=semestre_ab.semestre; 
		return true;
	end if;
		
end;
$$language plpgsql;

--aplicación de cupos
create or replace function aplicacion_de_cupos(_año int, _n_semestre int) returns boolean as $$
declare
	v record;
	v2 record;
	v3 record;
	last_id_error error.id_error%type;
begin

	select conteo_de_errores() into last_id_error;
	
	if not exists (select * from periodo where semestre = _año || '-' || _n_semestre and estado = 'cierre inscrip') then
       insert into error values (last_id_error, 'ingreso nota', _año|| '-' || _n_semestre, null, null, null, current_date, '?el semestre no se encuentra en un período válido para aplicar cupos.');
       return false; 
    end if;
	
	for v in select * from comision co,cursada cu where co.id_comision=cu.id_comision and co.id_materia=cu.id_materia loop
		for v2 in select * from cursada where id_materia=v.id_materia and id_comision=v.id_comision and estado!='dade de baja' order by f_inscripcion limit v.cupo loop
			update cursada c set estado='aceptade' where v2.id_alumne=c.id_alumne and c.estado='ingresade';
		end loop;
	
	end loop;
	for v3 in select * from cursada loop
		update cursada set estado='en espera' where v3.id_alumne=id_alumne and estado='ingresade';
	end loop;
	
	update periodo set estado='cursada' where semestre = _año || '-' || _n_semestre;
	return true;
end;
$$ language plpgsql;

--ingreso de nota de cursada
create or replace function ingreso_de_nota_de_cursada(_id_alumne cursada.id_alumne%type, _id_materia cursada.id_materia%type, _id_comision cursada.id_comision%type, _nota cursada.nota%type) returns boolean as $$
declare
	last_id_error error.id_error%type;
begin

    select conteo_de_errores() into last_id_error;

    if not exists (select * from periodo where estado = 'cursada') then 
        insert into error values (last_id_error, 'ingreso nota', null, _id_alumne, _id_materia, _id_comision, current_date, '?período de cursada cerrado.');
        return false;
    elsif _id_alumne not in (select id_alumne from alumne) then
        insert into error values (last_id_error, 'ingreso nota', null, _id_alumne, _id_materia, _id_comision, current_date, '?id de alumne no válido.');
        return false;
    elsif _id_materia not in (select id_materia from materia) then
        insert into error values (last_id_error, 'ingreso nota', null, _id_alumne, _id_materia, _id_comision, current_date, '?id de materia no válido.');
        return false;
    elsif _id_comision not in (select id_comision from comision where id_materia = _id_materia) then
        insert into error values (last_id_error, 'ingreso nota', null, _id_alumne, _id_materia, _id_comision, current_date, '?id de comisión no válido para la materia.');
        return false;
    elsif not exists (select * from cursada c where c.id_alumne = _id_alumne and c.id_materia = _id_materia and c.id_comision = _id_comision and c.estado = 'aceptade') then
        insert into error values (last_id_error, 'ingreso nota', null, _id_alumne, _id_materia, _id_comision, current_date, '?alumne no cursa en la comisión.');
        return false;
    elsif _nota < 0 or _nota > 10 then 
        insert into error values (last_id_error, 'ingreso nota', null, _id_alumne, _id_materia, _id_comision, current_date, '?nota no válida:' || _nota);
        return false;
    else 
        update cursada set nota = _nota where id_alumne = _id_alumne and id_materia = _id_materia and id_comision = _id_comision;
        return true;
    end if;
end;
$$ language plpgsql;

--cierre de cursada
create or replace function cierre_de_cursada(_id_materia cursada.id_materia%type, _id_comision cursada.id_comision%type) returns boolean as $$
declare
	last_id_error error.id_error%type;
	semestre_actual periodo.semestre%type;
	v record;
begin

	select conteo_de_errores() into last_id_error;
	
	if not exists (select * from periodo where estado = 'cursada') then
		insert into error values(last_id_error, 'cierre cursada', null, null, _id_materia, _id_comision, current_date, '?periodo de cursada cerrado.');
		return false;
		
	elsif _id_materia not in (select id_materia from materia) then
		insert into error values(last_id_error, 'cierre cursada', null, null, _id_materia, _id_comision, current_date, '?id de materia no válido.');
		return false;
		
	elsif _id_comision not in (select id_comision from comision where id_materia = _id_materia) then
		insert into error values(last_id_error, 'cierre cursada', null, null, _id_materia, _id_comision, current_date, '?id de comisión no valido para la materia.');
		return false;
	
	elsif not exists (select 1 from cursada where id_materia = _id_materia and id_comision = _id_comision) then
		insert into error values(last_id_error, 'cierre cursada', null, null, _id_materia, _id_comision, current_date, '?comision sin alumnes inscriptes.');
		return false;
	
	elsif exists(select * from cursada where nota is null and id_materia=_id_materia and id_comision=_id_comision and estado='aceptade') then
		insert into error values (last_id_error, 'cierre cursada', null, null, _id_materia, _id_comision, current_date, '?la carga de notas no está completa.');
		return false;
	
	else 
	select semestre into semestre_actual from periodo where estado = 'cursada';
		for v in (select * from cursada where id_comision = _id_comision and id_materia = _id_materia and estado = 'aceptade') loop
			if v.nota = 0 then 
				insert into historia_academica values (v.id_alumne, semestre_actual, _id_materia, _id_comision, 'ausente', v.nota, null);
			 elsif v.nota >= 1 and v.nota <= 3 then 
				insert into historia_academica values (v.id_alumne, semestre_actual, _id_materia, _id_comision, 'reprobada', v.nota, null);
			 elsif v.nota >= 4 and v.nota <= 6 then 
				insert into historia_academica values (v.id_alumne, semestre_actual, _id_materia, _id_comision, 'regular', v.nota, null);
			 elsif v.nota >= 7 and v.nota <= 10 then 
					insert into historia_academica values (v.id_alumne, semestre_actual, _id_materia, _id_comision, 'aprobada', v.nota, v.nota);
			end if;
		end loop;
		delete from cursada where id_comision = _id_comision and id_materia = _id_materia;
		return true;
	end if;
end;
$$ language plpgsql;

--envio de emails
create or replace function envio_de_emails_a_alumnes() returns trigger as $$
declare
	cant_emails int;
	id_email_act envio_email.id_email%type; 
	email_alumne alumne.email%type;
	cuerpo_email text;
begin

	cuerpo_email:= 'Materia: ' || (select nombre from materia where id_materia = new.id_materia) || ', Comisión: ' || new.id_comision || ', Alumne: ' || (select nombre from alumne where id_alumne = new.id_alumne) || ' ' ||(select apellido from alumne where id_alumne = new.id_alumne);
	-- Obtengo el próximo id para el email a enviar
	
	select count(*) into cant_emails from envio_email;
	
    if cant_emails = 0 then
        id_email_act := 0;
    else
        select max(id_email) into id_email_act from envio_email;
        id_email_act := id_email_act + 1;
    end if;
	
	if new.estado = 'ingresade' then
		select email into email_alumne from alumne where id_alumne = new.id_alumne;
		insert into envio_email values (id_email_act, current_date, email_alumne, 'Inscripción registrada', cuerpo_email, null, 'pendiente');
		
	end if;
	
	if new.estado = 'dade de baja' then
		select email into email_alumne from alumne where id_alumne = new.id_alumne;
		insert into envio_email values (id_email_act, current_date, email_alumne, 'Inscripción dada de baja', cuerpo_email, null, 'pendiente');
		
	end if;

	if old.estado = 'ingresade' and new.estado = 'aceptade' then
		select email into email_alumne from alumne where id_alumne = new.id_alumne;
		insert into envio_email values (id_email_act, current_date, email_alumne, 'Inscripción aceptada', cuerpo_email || ' ; Felicidades, fuiste aceptado.', null, 'pendiente');
	end if;
	if old.estado = 'ingresade' and new.estado = 'en espera' then
		select email into email_alumne from alumne where id_alumne = new.id_alumne;
		insert into envio_email values (id_email_act, current_date, email_alumne, 'Inscripción en espera', cuerpo_email || ' ; Quedaste en espera.', null, 'pendiente');
	end if;
	
	if old.estado = 'en espera' and new.estado = 'aceptade' then
		select email into email_alumne from alumne where id_alumne = new.id_alumne;
		insert into envio_email values (id_email_act, current_date, email_alumne, 'Inscripción aceptada', cuerpo_email, null, 'pendiente');
	end if;
	
	if new.estado = 'aprobada' then
		select email into email_alumne from alumne where id_alumne = new.id_alumne;
		insert into envio_email values (id_email_act, current_date, email_alumne, 'Cierre de cursada', cuerpo_email || ', Estado: ' || new.estado || ', Nota regularidad: ' || new.nota_regular || ', Nota final: ' || new.nota_final, null, 'pendiente');
	elsif new.estado = 'regular' or new.estado = 'reprobada' or new.estado = 'ausente' then
		select email into email_alumne from alumne where id_alumne = new.id_alumne;
		insert into envio_email values (id_email_act, current_date, email_alumne, 'Cierre de cursada', cuerpo_email || ', Estado: ' || new.estado || ', Nota regularidad: ' || new.nota_regular, null, 'pendiente');
	end if;
	
	return new;
end;
$$ language plpgsql;

create trigger envio_de_emails_a_alumnes_historia_academica
after insert on historia_academica
for each row
execute procedure envio_de_emails_a_alumnes();

create trigger envio_de_emails_a_alumnes_cursada
after update or insert on cursada
for each row
execute procedure envio_de_emails_a_alumnes();

--funcion para probar entrada_trx
create or replace function prueba_entradas() returns boolean as $$
declare
	v record;
begin

	for v in select * from entrada_trx order by id_orden loop
		if v.operacion='apertura' then
			perform apertura_de_inscripcion(v.año,v.nro_semestre);
		elsif v.operacion='alta inscrip' then
			perform  inscripcion_a_materia(v.id_alumne,v.id_materia,v.id_comision);
		elsif v.operacion='baja inscrip' then
			perform baja_de_inscripcion(v.id_alumne,v.id_materia);
		elsif v.operacion='cierre inscrip' then
			perform cierre_de_inscripcion(v.año,v.nro_semestre);
		elsif v.operacion='aplicacion cupo' then
			perform aplicacion_de_cupos(v.año,v.nro_semestre);
		elsif v.operacion='ingreso nota' then
			perform ingreso_de_nota_de_cursada(v.id_alumne,v.id_materia,v.id_comision,v.nota);
		elsif v.operacion='cierre cursada' then
			perform cierre_de_cursada(v.id_materia,v.id_comision);
		end if;
	end loop;
	return true;
end;
$$ language plpgsql;
