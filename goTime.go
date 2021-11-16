
package main  
import (  
    "fmt"  
    "time"  
)  
func main() {  
    present := time.Now()  
    fmt.Println("Today : ", present.Format("Mon, Jan 2, 2006 at 3:04pm"))  
    someTime := time.Date(2017, time.March, 30, 11, 30, 55, 123456, time.Local)  
    // compare time with time.Equal()  
    sameTime := someTime.Equal(present)  
    fmt.Println("someTime equals to now ? : ", sameTime)  
    // calculate the time different between today  
    // and long time ago  
    diff := present.Sub(someTime)  
    // convert diff to days  
    days := int(diff.Hours() / 24)  
    fmt.Printf("30th March 2017 was %d days ago \n", days)  
}  

/*package main  
  
import "fmt"  
import "time"  
  
func main() {  
    p := fmt.Println  
  
    present := time.Now()// current time  
    p(present)  
  
    DOB := time.Date(1993, 02, 28, 9,04,39,213 ,time.Local)  
    fmt.Println(DOB)  
  
    fmt.Println(DOB.Year())  
    fmt.Println(DOB.Month())  
    fmt.Println(DOB.Day())  
    fmt.Println(DOB.Hour())  
    fmt.Println(DOB.Minute())  
    fmt.Println(DOB.Second())  
    fmt.Println(DOB.Nanosecond())  
    fmt.Println(DOB.Location())  
  
    fmt.Println(DOB.Weekday())  
  
    fmt.Println(DOB.Before(present))  
    fmt.Println(DOB.After(present))  
    fmt.Println(DOB.Equal(present))  
  
    diff := present.Sub(DOB)  
    fmt.Println(diff)  
    fmt.Println(diff.Hours())  
    fmt.Println(diff.Minutes())  
    fmt.Println(diff.Seconds())  
    fmt.Println(diff.Nanoseconds())  
    fmt.Println(DOB.Add(diff))  
    fmt.Println(DOB.Add(-diff))  
}  
*/
