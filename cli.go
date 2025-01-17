package main
import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"time"
)

func crearBaseDeDatos() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`drop database if exists lozano_moreno_schaab_vallejos_db1`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`create database lozano_moreno_schaab_vallejos_db1`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nBase de datos creada correctamente!\n")
}

func crearTablas() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=lozano_moreno_schaab_vallejos_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
	create table alumne(id_alumne int, nombre text, apellido text, dni int, fecha_nacimento date,	telefono char(12), email text);
	create table materia(id_materia int, nombre text);
	create table correlatividad(id_materia int, id_mat_correlativa int);
	create table comision(id_materia int, id_comision int, cupo int);
	create table cursada(id_materia int, id_alumne int, id_comision int, f_inscripcion timestamp, nota int,	estado char(12));
	create table periodo(semestre text,	estado char(15));
	create table historia_academica(id_alumne int, semestre text, id_materia int, id_comision int, estado char(15),	nota_regular int, nota_final int);
	create table error(id_error int, operacion char(15), semestre text,	id_alumne int,	id_materia int,	id_comision int, f_error timestamp,	motivo varchar(80));
	create table envio_email(id_email int, f_generacion timestamp, email_alumne text, asunto text, cuerpo text,	f_envio timestamp, estado char(10));
	create table entrada_trx(id_orden int, operacion char(15), año int,	nro_semestre int,	id_alumne int,	id_materia int,	id_comision int, nota int);
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nTablas creadas\n")
}

func crearPKsyFKs() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=lozano_moreno_schaab_vallejos_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
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
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPKs' y FK's agregadas\n")

}

func borrarPKsyFKs() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=lozano_moreno_schaab_vallejos_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
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
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPKs' y FK's borradas\n")

}

func cargarDatos() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=lozano_moreno_schaab_vallejos_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	//Carga de alumnes

	type Alumne struct {
		Id_alumne        int    `json:"id_alumne"`
		Nombre           string `json:"nombre"`
		Apellido         string `json:"apellido"`
		Dni              int    `json:"dni"`
		Fecha_nacimiento string `json:"fecha_nacimiento"`
		Telefono         string `json:"telefono"`
		Email            string `json:"email"`
	}

	contenidoAl, err := ioutil.ReadFile("alumnes.json")
	if err != nil {
		log.Fatal(err)
	}

	var Alumnes []Alumne
	err2 := json.Unmarshal(contenidoAl, &Alumnes)
	if err2 != nil {
		log.Fatal(err)
	}

	formato:="2006-01-02"
	for _,x :=range Alumnes{
		fecha:=x.Fecha_nacimiento
		date, err := time.Parse(formato,fecha)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(`insert into alumne (id_alumne,nombre,apellido,dni,fecha_nacimento,telefono,email)values($1,$2,$3,$4,$5,$6,$7)`,x.Id_alumne,x.Nombre,x.Apellido,x.Dni,date,x.Telefono,x.Email)
		if err != nil {
		log.Fatal(err)
		}
	}
	
	//Carga de materias
	
	type Materia struct {
		Id_materia int  `json:"id_materia"`
		Nombre string	`json:"nombre"`
	}
	
	contenidoMat, err := ioutil.ReadFile("materias.json")
	if err != nil {
		log.Fatal(err)
	}
	
	var Materias []Materia
	err3 := json.Unmarshal(contenidoMat, &Materias)
	if err3 != nil {
		log.Fatal(err3)
	}
	
	for _,x :=range Materias{
		_, err = db.Exec(`insert into materia (id_materia,nombre) values ($1,$2)`,x.Id_materia,x.Nombre)
		if err != nil {
		log.Fatal(err)
		}
	}
	
	//Carga de comisiones
	
		type Comision struct {
		Id_materia int  `json:"id_materia"`
		Id_comision int	`json:"id_comision"`
		Cupo int  `json:"cupo"`
	}
	
	contenidoCom, err := ioutil.ReadFile("comisiones.json")
	if err != nil {
		log.Fatal(err)
	}
	
	var Comisiones []Comision
	err4 := json.Unmarshal(contenidoCom, &Comisiones)
	if err4 != nil {
		log.Fatal(err4)
	}
	
	for _,x :=range Comisiones{
		_, err = db.Exec(`insert into comision (id_materia,id_comision,cupo) values ($1,$2,$3)`,x.Id_materia,x.Id_comision,x.Cupo)
		if err != nil {
		log.Fatal(err)
		}
	}
	
	//Carga de correlatividades
	
		type Correlatividad struct {
		Id_materia int  `json:"id_materia"`
		Id_mat_correlativa int	`json:"id_mat_correlativa"`
	}
	
	contenidoCorr, err := ioutil.ReadFile("correlatividades.json")
	if err != nil {
		log.Fatal(err)
	}
	
	var Correlatividades []Correlatividad
	err5 := json.Unmarshal(contenidoCorr, &Correlatividades)
	if err5 != nil {
		log.Fatal(err5)
	}
	
	for _,x :=range Correlatividades{
		_, err = db.Exec(`insert into correlatividad (id_materia,id_mat_correlativa) values ($1,$2)`,x.Id_materia,x.Id_mat_correlativa)
		if err != nil {
		log.Fatal(err)
		}
	}
	
	//Carga de entradas
	
	type Entrada struct{
		Id_orden int  `json:"id_orden"`
		Operacion string `json:"operacion"`
		Año int  `json:"año"`
		Nro_semestre int  `json:"nro_semestre"`
		Id_alumne int `json:"id_alumne"`
		Id_materia int `json:"id_materia"`
		Id_comision int `json:"id_comision"`
		Nota int  `json:"nota"`
	}
	
	contenidoEn, err := ioutil.ReadFile("entradas_trx.json")
	if err != nil {
		log.Fatal(err)
	}
	
	var Entradas []Entrada
	err6 := json.Unmarshal(contenidoEn, &Entradas)
	if err6 != nil {
		log.Fatal(err6)
	}
	
	for _,x :=range Entradas{
		_, err = db.Exec(`insert into entrada_trx (id_orden,operacion,año,nro_semestre,id_alumne,id_materia,id_comision,nota) values ($1,$2,$3,$4,$5,$6,$7,$8)`,x.Id_orden,x.Operacion,x.Año,x.Nro_semestre,x.Id_alumne,x.Id_materia,x.Id_comision,x.Nota)
		if err != nil {
		log.Fatal(err)
		}
	}
	
	//Carga de periodos
	
	type Periodo struct{
		Semestre string  `json:"semestre"`
		Estado string `json:"estado"`
	}
	
	contenidoPe, err := ioutil.ReadFile("periodos.json")
	if err != nil {
		log.Fatal(err)
	}
	
	var Periodos []Periodo
	err7 := json.Unmarshal(contenidoPe, &Periodos)
	if err7 != nil {
		log.Fatal(err7)
	}
	
	for _,x :=range Periodos{
		_, err = db.Exec(`insert into periodo (semestre,estado) values ($1,$2)`,x.Semestre,x.Estado)
		if err != nil {
		log.Fatal(err)
		}
	}
	
	//Carga de historia academica
	
	type HistAcademica struct{
		Id_alumne int  `json:"id_alumne"`
		Semestre string `json:"semestre"`
		Id_materia int `json:"id_materia"`
		Id_comision int `json:"id_comision"`
		Estado string `json:"estado"`
		Nota_regular int `json:"nota_regular"`
		Nota_final int `json:"nota_final"`
	}
	
	contenidoHi, err := ioutil.ReadFile("historia_academica.json")
	if err != nil {
		log.Fatal(err)
	}
	
	var HistAcademicas []HistAcademica
	err8 := json.Unmarshal(contenidoHi, &HistAcademicas)
	if err8 != nil {
		log.Fatal(err8)
	}
	
	for _,x :=range HistAcademicas{
		_, err = db.Exec(`insert into historia_academica (id_alumne,semestre,id_materia,id_comision,estado,nota_regular,nota_final) values ($1,$2,$3,$4,$5,$6,$7)`,x.Id_alumne,x.Semestre,x.Id_materia,x.Id_comision,x.Estado,x.Nota_regular,x.Nota_final)
		if err != nil {
		log.Fatal(err)
		}
	}
	
	fmt.Printf("\nDatos cargados correctamente\n")

}

func cargarStoredProceduresTriggers() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=lozano_moreno_schaab_vallejos_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	conteoDeErrores(db)
	aperturaDeInscripcion(db)
	inscripcionAMateria(db)
	bajaDeInscripcion(db)
	cierreDeInscripcion(db)
	aplicacionDeCupos(db)
	ingresoDeNotaCursada(db)
	cierreDeCursada(db)
	envioDeEmailsAlumnes(db)
	prueba_entradas(db)
	fmt.Printf("\nStored Procedures cargados exitosamente.\n")
}

func conteoDeErrores(db *sql.DB) {
	_, err := db.Exec(`
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
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nConteo de errores cargada correctamente.\n")
	}
}

func aperturaDeInscripcion(db *sql.DB) {
	_, err := db.Exec(`
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
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nApertura de inscripcion cargada correctamente.\n")
	}
}

func inscripcionAMateria(db *sql.DB) {
	_, err := db.Exec(`
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
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nInscripcion a materia cargada correctamente.\n")
	}
}

func bajaDeInscripcion(db *sql.DB) {
	_, err := db.Exec(`
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
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nBaja de inscripcion cargada correctamente.\n")
	}
}

func cierreDeInscripcion(db *sql.DB) {
	_, err := db.Exec(`
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
	`)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nCierre de inscripcion cargada correctamente.\n")
	}
}

func aplicacionDeCupos(db *sql.DB) {
	_, err := db.Exec(`
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
	`)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nAplicacion de cupos cargada correctamente.\n")
	}
}

func ingresoDeNotaCursada(db *sql.DB) {
	_, err := db.Exec(`
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
	`)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nIngreso de nota cursada cargada correctamente.\n")
	}
}

func cierreDeCursada(db *sql.DB) {
	_, err := db.Exec(`
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
	`)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nCierre de cursada cargada correctamente.\n")
	}
}

func envioDeEmailsAlumnes(db *sql.DB) {
	_, err := db.Exec(`
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
	`)
		
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nEnvio de emails a alumnes y sus triggers cargados correctamente.\n")
	}
}
func prueba_entradas(db *sql.DB) {
	_, err := db.Exec(`
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
	
	`)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nFuncion para pruebas de la tabla entradas cargada correctamente.\n")
	}
}

func probarFuncionalidad() {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=lozano_moreno_schaab_vallejos_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	_, err9 := db.Exec(`select prueba_entradas()`)	
	if err9 != nil {
			log.Fatal(err9)
		} else {
			fmt.Printf("\nFuncionalidad probada correctamente.\n")
		}
}


func realizarOperacion(n_operacion string) {
	switch n_operacion {
	case "1":
		crearBaseDeDatos()
		break
	case "2":
		crearTablas()
		break
	case "3":
		crearPKsyFKs()
		break
	case "4":
		borrarPKsyFKs()
		break
	case "5":
		cargarDatos()
		break
	case "6":
		cargarStoredProceduresTriggers()
		break
	case "7":
		probarFuncionalidad()
		break
	default:
		fmt.Printf("\nOperación inválida\n")
	}
}

func iniciarCLI() {

	for {
		var n_operacion string

		fmt.Printf("\n¡Bienvenido!\n")
		fmt.Printf("Dadas las siguientes operaciones:\n")
		fmt.Printf("(1) Crear base de datos\n")
		fmt.Printf("(2) Crear las tablas\n")
		fmt.Printf("(3) Crear PK's y FK's\n")
		fmt.Printf("(4) Eliminar PK's y FK's\n")
		fmt.Printf("(5) Cargar datos en las tablas\n")
		fmt.Printf("(6) Crear los Stored Procedures y Triggers\n")
		fmt.Printf("(7) Probar funcionalidad con los datos cargados\n")
		fmt.Printf("(8) Salir del CLI\n")
		fmt.Printf("Ingresa el número de operación que deseas realizar: ")

		fmt.Scanf("%s", &n_operacion)

		if n_operacion == "8" {
			break
		} else {
			realizarOperacion(n_operacion)
		}
	}
}

func main() {
	iniciarCLI()
}
