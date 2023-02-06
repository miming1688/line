package line
 
import (
	"fmt"
	"runtime"
	"path/filepath"
)
 

func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}
func Print(vals...interface{}) {
	var skip int
	skip++
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		_, fileName := filepath.Split(file)
		fmt.Printf("%s %d ", fileName, line)
	}

	for _, val := range vals {
		if typeof(val)=="string"{
			fmt.Printf(" %s", val)
		}else if typeof(val)=="int"{
			fmt.Printf(" %d", val)
		}else if typeof(val)=="bool"{
			fmt.Printf(" %t", val)
		}else{
			fmt.Printf(" %s", val)
		}
	}
	fmt.Printf("%s \n", "")
}
 