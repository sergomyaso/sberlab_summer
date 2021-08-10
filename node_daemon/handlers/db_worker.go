package handlers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const PostgresUser = "postgres"
const PostgresPassword = "postgresbox"
const PostgresDb = "test_logs"
const PostgresMemoryTable = "memory_logs"
const PostgresCpuTable = "cpu_logs"
const SslMode = "disable"
const HostDb = "178.170.195.224"
const PortDb = "5432"
const ConnectTemplate = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
const TemplateInsertMemoryQuery = "insert into %s  (test_name, limits, max_usage) values ('%s', %d, %d)"
const TemplateInsertCpuQuery = "insert into %s  (test_name, usage, shares, cpu_period, cpu_quota, periods, throttling_periods, throttling_time) values ('%s', %d, %d, %d, %d, %d, %d, %d)"

type MemoryModel struct {
	TestName string
	Limits   int64
	MaxUsage int64
}

type CpuModel struct {
	TestName          string
	Usage             int64
	Shares            int64
	CpuPeriod         int64
	CpuQuota          int64
	Periods           int64
	ThrottlingPeriods int64
	ThrottlingTime    int64
}

func CreateConnection() (*sql.DB, error) {
	connectParams := fmt.Sprintf(ConnectTemplate, HostDb, PortDb, PostgresUser, PostgresPassword, PostgresDb, SslMode)
	db, err := sql.Open("postgres", connectParams)
	if err != nil {
		log.Printf("Connection to db was failed, error:%v\n", err)
		panic(err)
	}
	log.Println("Connection to db was succeed")
	return db, err
}

func InsertMemoryLogsToDb(model *MemoryModel) {
	sqlQuery := fmt.Sprintf(TemplateInsertMemoryQuery, PostgresMemoryTable, model.TestName, model.Limits, model.MaxUsage)
	InsertDataToDb(sqlQuery)
	log.Println("Adding memory logs was succeed")
}

func InsertCpuLogsToDb(model *CpuModel) {
	sqlQuery := fmt.Sprintf(
		TemplateInsertCpuQuery,
		PostgresCpuTable,
		model.TestName,
		model.Usage,
		model.Shares,
		model.CpuPeriod,
		model.CpuQuota,
		model.Periods,
		model.ThrottlingPeriods,
		model.ThrottlingTime,
	)
	InsertDataToDb(sqlQuery)
	log.Println("Adding cpu logs was succeed")
}

func InsertDataToDb(sqlQuery string) {
	db, err := CreateConnection()
	result, err := db.Exec(sqlQuery)
	if err != nil {
		log.Printf("Write data to db failed, error:%v\n", err)
		panic(err)
	}
	addedStrings, _ := result.RowsAffected()
	log.Printf("At db was added %d string", addedStrings)
	db.Close()
}
