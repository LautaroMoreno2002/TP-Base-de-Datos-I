package main

import (
    "encoding/json"
    "fmt"
    bolt "go.etcd.io/bbolt"
    "log"
    "strconv"
	"io/ioutil"
)

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
    // abre transacción de escritura
    tx, err := db.Begin(true)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

    err = b.Put(key, val)
    if err != nil {
        return err
    }

    // cierra transacción
    if err := tx.Commit(); err != nil {
        return err
    }

    return nil
}

func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
    var buf []byte

    // abre una transacción de lectura
    err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucketName))
        buf = b.Get(key)
        return nil
    })

    return buf, err
}

func cargarAlumnes(db *bolt.DB) {
	type Alumne struct {
		Id_alumne        int    `json:"id_alumne"`
		Nombre           string `json:"nombre"`
		Apellido         string `json:"apellido"`
		Dni              int    `json:"dni"`
		Fecha_nacimiento string `json:"fecha_nacimiento"`
		Telefono         string `json:"telefono"`
		Email            string `json:"email"`
	}


	archivo, err := ioutil.ReadFile("alumnes.json")
	if err != nil {
		log.Fatal(err)
	}
	var alumnes []Alumne

	err = json.Unmarshal(archivo, &alumnes)
	if err != nil {
		log.Fatal(err)
	}

	for _, Alumne := range alumnes {
		info, err := json.Marshal(Alumne)
		if err != nil {
			log.Fatal(err)
		}
		CreateUpdate(db, "alumne", []byte(strconv.Itoa(Alumne.Id_alumne)), info)

		//alumneEnBolt, err := ReadUnique(db, "alumne", []byte(strconv.Itoa(Alumne.Id_alumne)))

    	//fmt.Printf("%s\n", alumneEnBolt)
	}
	fmt.Println("Alumnes cargados exitosamente.")
}

func cargarMaterias(db *bolt.DB) {
	type Materia struct {
		Id_materia int  `json:"id_materia"`
		Nombre string	`json:"nombre"`
	}
	archivo, err := ioutil.ReadFile("materias.json")
	if err != nil {
		log.Fatal(err)
	}
	var materias []Materia

	err = json.Unmarshal(archivo, &materias)
	if err != nil {
		log.Fatal(err)
	}

	for _, Materia := range materias {
		info, err := json.Marshal(Materia)
		if err != nil {
			log.Fatal(err)
		}
		CreateUpdate(db, "materia", []byte(strconv.Itoa(Materia.Id_materia)), info)

		//materiaEnBolt, err := ReadUnique(db, "materia", []byte(strconv.Itoa(Materia.Id_materia)))

    	//fmt.Printf("%s\n", materiaEnBolt)
	}
	fmt.Println("Materias cargadas exitosamente.")
}

func cargarComision(db *bolt.DB) {
	type Comision struct {
		Id_materia int  `json:"id_materia"`
		Id_comision int	`json:"id_comision"`
		Cupo int  `json:"cupo"`
	}
	archivo, err := ioutil.ReadFile("comisiones.json")
	if err != nil {
		log.Fatal(err)
	}
	var comisiones []Comision

	err = json.Unmarshal(archivo, &comisiones)
	if err != nil {
		log.Fatal(err)
	}

	for _, Comision := range comisiones {
		info, err := json.Marshal(Comision)
		if err != nil {
			log.Fatal(err)
		}
		CreateUpdate(db, "comision", []byte(strconv.Itoa(Comision.Id_comision)), info)

		//comisionEnBolt, err := ReadUnique(db, "comision", []byte(strconv.Itoa(Comision.Id_comision)))

    	//fmt.Printf("%s\n", comisionEnBolt)
	}
	fmt.Println("Comisiones cargadas exitosamente.")
}

func cargarInscripcionesACursada(db *bolt.DB) {
	type Cursada struct {
		Id_cursada int
		Id_materia int
		Id_alumne int
		Id_comision int
		F_inscripcion string
		Nota int // inicialmente en null: no hay nota
		Estado string //`ingresade',`aceptade',`en espera',`dade de baja'
	}

	archivo, err := ioutil.ReadFile("inscripciones_cursada.json")
	if err != nil {
		log.Fatal(err)
	}
	var cursadas []Cursada

	err = json.Unmarshal(archivo, &cursadas)
	if err != nil {
		log.Fatal(err)
	}

	for _, Cursada := range cursadas {
		info, err := json.Marshal(Cursada)
		if err != nil {
			log.Fatal(err)
		}
		CreateUpdate(db, "cursada", []byte(strconv.Itoa(Cursada.Id_cursada)), info)

		//cursadaEnBolt, err := ReadUnique(db, "cursada", []byte(strconv.Itoa(Cursada.Id_cursada)))

    	//fmt.Printf("%s\n", cursadaEnBolt)
	}
	fmt.Println("Cursadas cargadas exitosamente.")
}

func cliBoltDB() {
	for {
		var n_operacion int

		fmt.Printf("\n¡Bienvenido!\n")
		fmt.Printf("Dadas las siguientes operaciones:\n")
		fmt.Printf("(1) Crear base de datos y cargarle datos\n")
		fmt.Printf("(2) Salir del CLI\n")
		fmt.Printf("Ingresa el número de operación que deseas realizar: ")

		fmt.Scanf("%d", &n_operacion)

		switch (n_operacion) {
		case 1: 
			db, err := bolt.Open("lozano_moreno_schaab_vallejos_db1.db", 0600, nil)
			if err != nil {
				log.Fatal(err)
			}

			cargarAlumnes(db)
			cargarMaterias(db)
			cargarComision(db)
			cargarInscripcionesACursada(db)

			defer db.Close()
			break
		case 2:
			return
		}
	}
}

func main() {
	cliBoltDB()
}	