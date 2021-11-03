package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var dbPathName = "./DataBase/"

// IsDirFile to judge a dir or file is available or not
// Database functoins methods
func IsDirFile(pathName string) bool {
	_, err := os.Stat(pathName)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

// NewDBItem to new db unit line item, key and value
func NewDBItem(dbName string, tableName string, lineName string, lineUnits [][2]string) {
	// os.Stat(pth)
	if !(IsDirFile(dbPathName + dbName)) {
		// os.MkdirAll("./dir1", os.ModePerm)
		os.MkdirAll(dbPathName+dbName, os.ModePerm)
	}

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	var tableFile = dbPathName + dbName + "/" + tableName + ".json"

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// read the table file
	var fileRead []byte
	if IsDirFile(tableFile) {
		sonFileRead, err := ioutil.ReadFile(tableFile)
		fileRead = sonFileRead
		if err != nil {
			fmt.Println("tableFile ioutil.ReadFile err: ", err)
		}
		// fmt.Println("fileRead = ", string(fileRead))
	} else {
		fmt.Println("tableFile not exist")
		// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
		// write to save the data as a json file
		err := ioutil.WriteFile(tableFile, []byte(nil), 0644)
		if err != nil {
			fmt.Println("ioutil.WriteFile err: ", err)
		}
	}

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// Unmarshal the tableFilde JSON string
	var saveMap = make(map[string]map[string]string)
	if len(fileRead) != 0 {
		err := json.Unmarshal([]byte(string(fileRead)), &saveMap)
		if err != nil {
			fmt.Println("json.Unmarshal err: ", err)
		}
	} else {
		fmt.Println("fileRead is exist but nil")
	}

	// fmt.Println("saveMap = ", saveMap)

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// new a unit map to save the temple data
	var mapp = make(map[string]string)

	// Handle the saveMap
	for _, units := range lineUnits {
		mapp[units[0]] = units[1]
	}

	// Assignment for the big saveMap
	saveMap[lineName] = mapp

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	// write to save the data as a json file
	jsonMap, _ := json.Marshal(saveMap)
	strJSON := string(jsonMap)
	// fmt.Printf("strJSON = %s \n", strJSON)

	err := ioutil.WriteFile(tableFile, []byte(strJSON), 0644)
	if err != nil {
		fmt.Println("ioutil.WriteFile err: ", err)
	}

	// fmt.Println("saveMap = ", saveMap)
	// fmt.Println("mapp = ", mapp)

	// fmt.Println("=================================================================")

}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

// DelectDBItem to use the db line unit key name to delect it's corresponding value
// the key name of the DB item can not be the same
func DelectDBItem(dbName string, tableName string, lineName string, keyName string) {
	// tableFile
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	var tableFile = dbPathName + dbName + "/" + tableName + ".json"

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// read the table file
	fileRead, err := ioutil.ReadFile(tableFile)
	if err != nil {
		fmt.Println("tableFile ioutil.ReadFile err: ", err)
	}

	// fmt.Println("fileRead = ", string(fileRead))

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// Unmarshal the tableFilde JSON string
	var saveMap = make(map[string]map[string]string)
	if len(fileRead) != 0 {
		err := json.Unmarshal([]byte(string(fileRead)), &saveMap)
		if err != nil {
			fmt.Println("json.Unmarshal err: ", err)
		}
	} else {
		fmt.Println("fileRead is exist but nil")
	}

	// fmt.Println("saveMap = ", saveMap)

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// Main prpcessing
	// delect the specific item
	// delete(m map[Key]Type, key Key)
	delete(saveMap[lineName], keyName)

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	// write to save the data as a json file
	jsonMap, _ := json.Marshal(saveMap)
	strJSON := string(jsonMap)
	// fmt.Printf("strJSON = %s \n", strJSON)

	err = ioutil.WriteFile(tableFile, []byte(strJSON), 0644)
	if err != nil {
		fmt.Println("ioutil.WriteFile err: ", err)
	}

	// fmt.Println("after delect saveMap = ", saveMap)

	// fmt.Println("=================================================================")

	// fmt.Println(`saveMap["sonMap1"]["City"] = `, saveMap["Login1"]["City"])

}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

// SetDBItem to use the db line unit key name and it's corresponding value to set a line unit item
// the key name of the DB item can not be the same
func SetDBItem(dbName string, tableName string, lineName string, keyName string, value string) {
	// tableFile
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	var tableFile = dbPathName + dbName + "/" + tableName + ".json"

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// read the table file
	fileRead, err := ioutil.ReadFile(tableFile)
	if err != nil {
		fmt.Println("tableFile ioutil.ReadFile err: ", err)
	}

	// fmt.Println("fileRead = ", string(fileRead))

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// Unmarshal the tableFilde JSON string
	var saveMap = make(map[string]map[string]string)
	if len(fileRead) != 0 {
		err := json.Unmarshal([]byte(string(fileRead)), &saveMap)
		if err != nil {
			fmt.Println("json.Unmarshal err: ", err)
		}
	} else {
		fmt.Println("fileRead is exist but nil")
	}

	// fmt.Println("saveMap = ", saveMap)

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// Main prpcessing
	// saveMap[lineName], keyName
	// saveMap["Login1"]["City"]
	saveMap[lineName][keyName] = value

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	// write to save the data as a json file
	jsonMap, _ := json.Marshal(saveMap)
	strJSON := string(jsonMap)
	// fmt.Printf("strJSON = %s \n", strJSON)

	err = ioutil.WriteFile(tableFile, []byte(strJSON), 0644)
	if err != nil {
		fmt.Println("ioutil.WriteFile err: ", err)
	}

	// fmt.Println("After set saveMap = ", saveMap)
	// fmt.Println("=================================================================")
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

// GetDBItem to use the db line unit key name to get it's corresponding value
// the key name of the DB item can not be the same
func GetDBItem(dbName string, tableName string, lineName string, keyName string) string {
	// tableFile
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	var tableFile = dbPathName + dbName + "/" + tableName + ".json"

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// read the table file
	fileRead, err := ioutil.ReadFile(tableFile)
	if err != nil {
		fmt.Println("tableFile ioutil.ReadFile err: ", err)
	}

	// fmt.Println("fileRead = ", string(fileRead))

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// Unmarshal the tableFilde JSON string
	var saveMap = make(map[string]map[string]string)
	if len(fileRead) != 0 {
		err := json.Unmarshal([]byte(string(fileRead)), &saveMap)
		if err != nil {
			fmt.Println("json.Unmarshal err: ", err)
		}
	} else {
		fmt.Println("fileRead is exist but nil")
	}

	// fmt.Println("get saveMap = ", saveMap)

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// Main prpcessing
	// saveMap[lineName], keyName
	// saveMap["Login1"]["City"]
	value := saveMap[lineName][keyName]

	return value
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

// DelectDB to delect the whole specific
func DelectDB(dbName string) {
	// tableFile
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	var DB = dbPathName + dbName
	os.RemoveAll(DB)
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

// DelectTable to delect the whole specific DB Table
func DelectTable(dbName string, tableName string) {
	// tableFile
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	var DBTable = dbPathName + dbName + "/" + tableName + ".json"
	os.RemoveAll(DBTable)

}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

// SumTables to sum up the total number of the specific DB
func SumTables(dbName string) int {
	// tableFile
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	var DBTable = dbPathName + dbName + "/*"
	items, err := filepath.Glob(DBTable)
	if err != nil {
		return 0
	}
	// fmt.Println("items = ", items)
	// fmt.Println("len(items) = ", len(items))
	return len(items)
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

// SumDBs to sum up the total number of the specific DB
func SumDBs() int {
	// tableFile
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	var DBTable = dbPathName + "/*"
	items, err := filepath.Glob(DBTable)
	if err != nil {
		return 0
	}
	// fmt.Println("items = ", items)
	// fmt.Println("len(items) = ", len(items))
	return len(items)
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

// SumTableLines to sum up the total lines number of the specific DB table
func SumTableLines(dbName string, tableName string) int {
	// tableFile
	// err := ioutil.WriteFile(filename string, data []byte, perm os.FileMode)
	var tableFile = dbPathName + dbName + "/" + tableName + ".json"

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// read the table file
	fileRead, err := ioutil.ReadFile(tableFile)
	if err != nil {
		fmt.Println("tableFile ioutil.ReadFile err: ", err)
		return 0
	}

	// fmt.Println("fileRead = ", string(fileRead))

	// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
	// Unmarshal the tableFilde JSON string
	var saveMap = make(map[string]map[string]string)
	if len(fileRead) != 0 {
		err := json.Unmarshal([]byte(string(fileRead)), &saveMap)
		if err != nil {
			fmt.Println("json.Unmarshal err: ", err)
		}
	} else {
		fmt.Println("fileRead is exist but nil")
		return 0
	}

	return len(saveMap)
}
