package part2

import (
	"fmt"
	"sync"
	"time"
)

type User struct{
	ID int
	Name string
}

func fetchUser(id int, ch chan User, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Microsecond * 5)
	ch <- User{
		Name: fmt.Sprintf("User-%d", id),
		ID: id,
	}
}

func Prob9(){
	userIDs := []int{1,2,3}
	var wg sync.WaitGroup
	var userList []User
	ch := make(chan User, len(userList))

	for _, user := range userIDs{
		wg.Add(1)
		go fetchUser(user, ch, &wg)
	}

	go func ()  {
		wg.Wait()
		close(ch)
	}()

	for c := range ch{
		userList = append(userList, c)
	}
	fmt.Println(userList)
}