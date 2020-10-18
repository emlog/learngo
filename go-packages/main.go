/*
包
*/
package gopkgs

import "fmt"

func RunTest() {
	//数组
	animals := [3]string{"cat", "dog", "monkey"}
	for i := 0; i <= 2; i++ {
		fmt.Printf("animal is %s \n", animals[i])
	}

}
