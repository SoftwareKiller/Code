package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// 日志记录结构
type LogEntry struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

// 模拟数据库表
type Database struct {
	mu         sync.Mutex
	data       map[string]int // 存储数据
	logFile    string         // 日志文件路径
	checkpoint string         // 检查点文件路径
}

// 事务结构
type Transaction struct {
	db        *Database
	dirtyData map[string]int // 事务中的脏数据
}

// 初始化数据库
func NewDatabase(logFile string, checkpointFile string) *Database {
	db := &Database{
		data:       make(map[string]int),
		logFile:    logFile,
		checkpoint: checkpointFile,
	}

	// 从检查点恢复数据
	db.recoverFromCheckpoint()
	return db
}

// 从检查点恢复数据
func (db *Database) recoverFromCheckpoint() {
	db.mu.Lock()
	defer db.mu.Unlock()

	// 读取检查点文件
	file, err := os.ReadFile(db.checkpoint)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果检查点文件不存在，初始化空数据库
			db.data = make(map[string]int)
			return
		}
		panic(fmt.Sprintf("Failed to read checkpoint file: %v", err))
	}

	// 解析检查点数据
	err = json.Unmarshal(file, &db.data)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse checkpoint file: %v", err))
	}

	// 重放日志（如果有）
	db.replayLogs()
}

// 重放日志
func (db *Database) replayLogs() {
	file, err := os.ReadFile(db.logFile)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果日志文件不存在，无需重放
			return
		}
		panic(fmt.Sprintf("Failed to read log file: %v", err))
	}

	// 解析日志
	var logs []LogEntry
	err = json.Unmarshal(file, &logs)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse log file: %v", err))
	}

	// 重放日志中的操作
	for _, logEntry := range logs {
		db.data[logEntry.Key] = logEntry.Value
	}

	// 清空日志文件
	err = os.WriteFile(db.logFile, []byte("[]"), 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to clear log file: %v", err))
	}
}

// 保存检查点
func (db *Database) saveCheckpoint() {
	db.mu.Lock()
	defer db.mu.Unlock()

	// 将当前数据保存到检查点文件
	data, err := json.Marshal(db.data)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal checkpoint data: %v", err))
	}

	err = os.WriteFile(db.checkpoint, data, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to write checkpoint file: %v", err))
	}
}

// 开始一个事务
func (db *Database) Begin() *Transaction {
	db.mu.Lock()
	defer db.mu.Unlock()

	return &Transaction{
		db:        db,
		dirtyData: make(map[string]int),
	}
}

// 事务中执行写操作
func (tx *Transaction) Write(key string, value int) {
	// 直接写入脏数据，不需要加锁
	tx.dirtyData[key] = value
}

// 提交事务
func (tx *Transaction) Commit() error {
	tx.db.mu.Lock()
	defer tx.db.mu.Unlock()

	// 将脏数据写入日志文件（预写日志）
	var logs []LogEntry
	for key, value := range tx.dirtyData {
		logs = append(logs, LogEntry{Key: key, Value: value})
	}

	logData, err := json.Marshal(logs)
	if err != nil {
		return fmt.Errorf("failed to marshal dirty data: %v", err)
	}

	// 读取现有日志
	file, err := os.ReadFile(tx.db.logFile)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read log file: %v", err)
	}

	var existingLogs []LogEntry
	if len(file) > 0 {
		err = json.Unmarshal(file, &existingLogs)
		if err != nil {
			return fmt.Errorf("failed to parse log file: %v", err)
		}
	}

	// 追加新日志
	existingLogs = append(existingLogs, logs...)

	// 写入日志文件
	logData, err = json.Marshal(existingLogs)
	if err != nil {
		return fmt.Errorf("failed to marshal logs: %v", err)
	}

	err = os.WriteFile(tx.db.logFile, logData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write log file: %v", err)
	}

	// 将脏数据写入实际数据
	for key, value := range tx.dirtyData {
		tx.db.data[key] = value
	}

	return nil
}

// 回滚事务
func (tx *Transaction) Rollback() {
	tx.db.mu.Lock()
	defer tx.db.mu.Unlock()

	// 清空脏数据
	tx.dirtyData = make(map[string]int)
}

// 读取数据
func (db *Database) Read(key string) (int, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	value, exists := db.data[key]
	return value, exists
}

// 打印日志
func (db *Database) PrintLogs() {
	db.mu.Lock()
	defer db.mu.Unlock()

	file, err := os.ReadFile(db.logFile)
	if err != nil {
		fmt.Println("No logs found.")
		return
	}

	var logs []LogEntry
	err = json.Unmarshal(file, &logs)
	if err != nil {
		fmt.Println("Failed to parse logs:", err)
		return
	}

	fmt.Println("Transaction Logs:")
	for _, log := range logs {
		fmt.Printf("Key: %s, Value: %d\n", log.Key, log.Value)
	}
}

func main() {
	// 初始化数据库
	db := NewDatabase("transaction_log.json", "checkpoint.json")

	// 设置初始数据
	db.data["accountA"] = 1000
	db.data["accountB"] = 500
	db.saveCheckpoint() // 保存初始检查点

	// 开始事务
	tx := db.Begin()

	// 执行写操作
	tx.Write("accountA", 900) // 从 accountA 扣除 100
	tx.Write("accountB", 600) // 向 accountB 增加 100
	tx.Write("accountC", 800) // 从 accountA 扣除 100
	tx.Write("accountD", 300) // 向 accountB 增加 100

	// 提交事务
	if err := tx.Commit(); err != nil {
		fmt.Println("Commit failed:", err)
	} else {
		fmt.Println("Commit succeeded!")
	}

	// 读取数据
	valueA, existsA := db.Read("accountA")
	valueB, existsB := db.Read("accountB")
	fmt.Printf("accountA: %d (exists: %v)\n", valueA, existsA)
	fmt.Printf("accountB: %d (exists: %v)\n", valueB, existsB)

	// 打印日志
	db.PrintLogs()

	// 保存检查点
	db.saveCheckpoint()
}
