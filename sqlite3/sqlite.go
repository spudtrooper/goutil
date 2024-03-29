package sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"unicode"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"

	_ "github.com/mattn/go-sqlite3"
)

func toSnakeCase(s string) string {
	var result strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if i > 0 && unicode.IsUpper(runes[i]) {
			if unicode.IsLower(runes[i-1]) || (i+1 < len(runes) && unicode.IsLower(runes[i+1])) {
				result.WriteRune('_')
			}
		}
		result.WriteRune(unicode.ToLower(runes[i]))
	}

	return result.String()
}

func mustFormat(res sql.Result) string {
	if res == nil {
		return "nil"
	}
	rows, err := res.RowsAffected()
	check.Err(err)
	id, err := res.LastInsertId()
	check.Err(err)
	return fmt.Sprintf("Result{RowsAffected: %d, LastInsertId: %d}", rows, id)
}

//go:generate genopts --function OpenDB createDBIfNotExists
func OpenDB(dbPath string, optss ...OpenDBOption) (*sql.DB, error) {
	if dbPath == "" {
		return nil, errors.Errorf("no dbname provided")
	}

	opts := MakeOpenDBOptions(optss...)
	if opts.CreateDBIfNotExists() {
		log.Printf("NOTE: We're not actually creating %s, I don't know why I had this as an option", dbPath)
	}

	return sql.Open("sqlite3", dbPath)
}

//go:generate genopts --function PopulateSqlite3Table --extends PopulateSqlite3TableFromDB,OpenDB
func PopulateSqlite3Table(dbPath, tableName string, data []interface{}, optss ...PopulateSqlite3TableOption) error {
	opts := MakePopulateSqlite3TableOptions(optss...)

	db, err := OpenDB(dbPath, opts.ToOpenDBOptions()...)
	if err != nil {
		return errors.Errorf("Could not open database %s: %v", dbPath, err)
	}
	if db == nil {
		return fmt.Errorf("no db")
	}
	defer db.Close()

	if err := PopulateSqlite3TableFromDB(db, tableName, data, opts.ToPopulateSqlite3TableFromDBOptions()...); err != nil {
		return errors.Errorf("Could not populate sqlite3 table from db: %v", err)
	}

	return nil
}

//go:generate genopts --function DropTableIfExists verbose
func DropTableIfExists(db *sql.DB, tableName string, optss ...DropTableIfExistsOption) error {
	opts := MakeDropTableIfExistsOptions(optss...)

	sql := fmt.Sprintf(`DROP TABLE IF EXISTS "%s"`, tableName)
	res, err := db.Exec(sql)
	if opts.Verbose() {
		log.Printf("result from dropping table %s: %s", tableName, mustFormat(res))
	}
	if err != nil {
		return errors.Errorf("Could not drop table %s: sql : %s, error: %v", tableName, sql, err)
	}

	return nil
}

//go:generate genopts --function PopulateSqlite3TableFromDB --extends OpenDB,DropTableIfExists dropIfExists primaryKey:string lowerCaseColumnNames snakeCaseColumnNames removeInvalidCharsFromColumnNames verbose deleteWhere:string
func PopulateSqlite3TableFromDB(db *sql.DB, tableName string, data []interface{}, optss ...PopulateSqlite3TableFromDBOption) error {
	if tableName == "" {
		return fmt.Errorf("no tableName provided")
	}

	if db == nil {
		return fmt.Errorf("no db")
	}

	if len(data) == 0 {
		return fmt.Errorf("no data provided")
	}

	opts := MakePopulateSqlite3TableFromDBOptions(optss...)

	primaryKey := opts.PrimaryKey()

	// Maybe drop the table
	if opts.DropIfExists() {
		log.Printf("dropping if exists table %s", tableName)
		if err := DropTableIfExists(db, tableName, opts.ToDropTableIfExistsOptions()...); err != nil {
			return err
		}
	}

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
	createSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s" (%s)`, tableName, strings.Join(fields, ", "))
	res, err := db.Exec(createSQL)
	if err != nil {
		return errors.Errorf("Could not create table %s: sql: %s, error: %v", tableName, createSQL, err)
	}
	if opts.Verbose() {
		log.Printf("result from creating table %s: %s", tableName, mustFormat(res))
	}

	// Maybe delete certain rows
	if opts.DeleteWhere() != "" {
		where := opts.DeleteWhere()
		log.Printf("deleting rows from %s where %s", tableName, where)
		sql := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, where)
		res, err := db.Exec(sql)
		if opts.Verbose() {
			log.Printf("result from deleting rows from table %s where %s: %s", tableName, where, mustFormat(res))
		}
		if err != nil {
			return errors.Errorf("Could not drop table %s where %s: sql : %s, error: %v", tableName, where, sql, err)
		}
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
				v = fmt.Sprintf(`%s`, v)
			}
			vals = append(vals, v)
		}

		insertSQL := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s)`, tableName, strings.Join(names, ","), strings.Join(placeholders, ","))
		if _, err := db.Exec(insertSQL, vals...); err != nil {
			return errors.Errorf("Could not insert into table %s: item: %+v, sql: %s, error: %v", tableName, item, insertSQL, err)
		}
	}

	if opts.Verbose() {
		log.Printf("inserted %d rows into table %s", len(data), tableName)
	}

	return nil
}
