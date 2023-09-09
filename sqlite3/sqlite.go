package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	_ "github.com/mattn/go-sqlite3"
)

func createDBIfNotExists(dbname string) (*sql.DB, error) {
	if _, err := os.Stat(dbname); os.IsNotExist(err) {
		db, err := sql.Open("sqlite3", dbname)
		if err != nil {
			return nil, err
		}
		return db, nil
	} else if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func toSnakeCase(s string) string {
	var res []string
	for i, c := range s {
		if i > 0 && c >= 'A' && c <= 'Z' {
			res = append(res, "_")
		}
		res = append(res, string(c))
	}
	return strings.ToLower(strings.Join(res, ""))
}

//go:generate genopts --function PopulateSqlite3Table dropIfExists primaryKey:string createDBIfNotExists lowerCaseColumnNames snakeCaseColumnNames removeInvalidCharsFromColumnNames
func PopulateSqlite3Table(dbname, tableName string, data []interface{}, optss ...PopulateSqlite3TableOption) error {
	opts := MakePopulateSqlite3TableOptions(optss...)

	primaryKey := opts.PrimaryKey()

	fieldName := func(structFieldName string) string {
		res := structFieldName
		if opts.SnakeCaseColumnNames() {
			res = toSnakeCase(res)
		}
		if opts.RemoveInvalidCharsFromColumnNames() {
			res = strings.ReplaceAll(res, " ", "_")
			res = strings.ReplaceAll(res, "-", "_")
			res = strings.ReplaceAll(res, ".", "_")
		}
		if opts.LowerCaseColumnNames() {
			res = strings.ToLower(res)
		}
		return res
	}

	if dbname == "" {
		return fmt.Errorf("no dbname provided")
	}
	if tableName == "" {
		return fmt.Errorf("no tableName provided")
	}

	if len(data) == 0 {
		return fmt.Errorf("no data provided")
	}

	var db *sql.DB
	if opts.CreateDBIfNotExists() {
		log.Printf("maybe creating database %s", dbname)
		d, err := createDBIfNotExists(dbname)
		if err != nil {
			return errors.Errorf("Could not create database %s: %v", dbname, err)
		}
		db = d
	} else {
		d, err := sql.Open("sqlite3", dbname)
		if err != nil {
			return errors.Errorf("Could not open database %s: %v", dbname, err)
		}
		db = d
	}

	if db == nil {
		return fmt.Errorf("no db")
	}

	defer db.Close()

	// Maybe drop the table
	if opts.DropIfExists() {
		log.Printf("maybe dropping database %s", dbname)
		sql := fmt.Sprintf(`DROP TABLE IF EXISTS "%s"`, tableName)
		_, err := db.Exec(sql)
		if err != nil {
			return errors.Errorf("Could not drop table %s: sql : %s, error: %v",
				tableName, sql, err)
		}
	}

	// Reflect over the first item to create a table schema
	typ := reflect.TypeOf(data[0])
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	fields := make([]string, 0, typ.NumField())
	fieldsToUse := map[string]bool{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		columnType := ""

		switch field.Type.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16, reflect.Int8:
			columnType = "INTEGER"
		case reflect.Float32, reflect.Float64:
			columnType = "REAL"
		case reflect.String:
			columnType = "TEXT"
		default:
			// TODO: More?
			continue
		}

		if field.Name == primaryKey {
			columnType += " PRIMARY KEY"
		}

		fieldsToUse[field.Name] = true
		fields = append(fields, fmt.Sprintf("%s %s", fieldName(field.Name), columnType))
	}
	createSQL := fmt.Sprintf(`CREATE TABLE "%s" (%s)`, tableName, strings.Join(fields, ", "))
	if _, err := db.Exec(createSQL); err != nil {
		return errors.Errorf("Could not create table %s: sql: %s, error: %v", tableName, createSQL, err)
	}

	// Insert the data
	for _, item := range data {
		vals := make([]interface{}, 0, typ.NumField())
		names := make([]string, 0, typ.NumField())
		placeholders := make([]string, 0, typ.NumField())

		val := reflect.ValueOf(item)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			if !fieldsToUse[field.Name] {
				continue
			}
			names = append(names, fieldName(field.Name))
			placeholders = append(placeholders, "?")
			v := val.Field(i).Interface()
			if val.Field(i).Type().Name() == "string" {
				v = fmt.Sprintf(`"%s"`, v)
			}
			vals = append(vals, v)
		}

		insertSQL := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s)`, tableName, strings.Join(names, ","), strings.Join(placeholders, ","))
		if _, err := db.Exec(insertSQL, vals...); err != nil {
			return errors.Errorf("Could not insert into table %s: item: %+v, sql: %s, error: %v", tableName, item, insertSQL, err)
		}
	}

	return nil
}
