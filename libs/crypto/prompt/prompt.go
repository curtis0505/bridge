package prompt

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/klaytn/klaytn/console"
)

type handler interface{}

type Prompt struct {
	handler handler
	batchs  []string
}

func NewPrompt(handler handler) *Prompt {
	r := &Prompt{
		handler: handler,
		batchs:  []string{},
	}
	r.initHistory()
	return r
}

// handler의 모든 메소드를 history에 넣는다.
func (p *Prompt) initHistory() {
	value := reflect.ValueOf(p.handler)
	t := value.Type()
	for m := 0; m < t.NumMethod(); m++ {
		method := value.MethodByName(t.Method(m).Name)
		cmd := t.Method(m).Name
		for i := 0; i < method.Type().NumIn(); i++ {
			cmd += fmt.Sprintf(" %v", method.Type().In(i).Name())
		}
		console.Stdin.SetHistory([]string{cmd})
	}
}

func (p *Prompt) SetBatch(batch ...string) {
	p.batchs = append(p.batchs, batch...)
}

func (p *Prompt) Run(fin chan<- struct{}) {
	for {
	top_loop:
		<-time.After(0)
		//command를 입력을 받음.
		args, err := func() ([]string, error) {
			var (
				command string
				err     error
			)

			if len(p.batchs) == 0 {
				command, err = console.Stdin.PromptInput(console.DefaultPrompt)
			} else {
				command = p.batchs[0]
				p.batchs = p.batchs[1:]
			}

			if err == nil {
				ret := []string{}
				//앞뒤 space 제거, 제거후 빈문자인경우 slice에서 제거.
				for _, e := range strings.Split(command, " ") {
					if trimed := strings.TrimSpace(e); trimed != "" {
						ret = append(ret, trimed)
					}
				}
				return ret, nil
			} else {
				return []string{}, err
			}
		}()

		//다음에 비슷한것을 작성할것을 대비해서 history에 넣는다.
		console.Stdin.SetHistory([]string{strings.Join(args, " ")})

		//에러인경우.. 대체로 Interrupt인경우
		if err != nil {
			if fin != nil {
				fin <- struct{}{}
			}
			return
		}

		//빈문자가 포함되어 있으면 args를 재정리함.
		args = func() []string {
			for i, a := range args {
				c := strings.TrimSuffix(strings.TrimSpace(a), "\n")
				if c == "" {
					args = append(args[:i], args[i+1:]...)
				}
			}
			return args
		}()
		if len(args) == 0 {
			goto top_loop
		} else if len(args) == 1 {
			if args[0] == "exit" {
				if fin != nil {
					fin <- struct{}{}
				}
				return
			} else if args[0] == "help" || args[0] == "h" { //커맨드 리스트를 출력
				t := reflect.TypeOf(p.handler)
				num := t.NumMethod()
				for i := 0; i < num; i++ {
					fmt.Println(t.Method(i).Name)
				}
				goto top_loop
			}
		}

		//args[0]은 메소드명임으로 맨앞문자를 대문자로 바꿔준다.
		methodName := strings.ToUpper(string([]byte(args[0])[0:1])) + string([]byte(args[0])[1:])
		inputs := args[1:]

		//handler객체에서 메소드를 찾는다.
		value := reflect.ValueOf(p.handler)
		method := value.MethodByName(methodName)
		if (method == reflect.Value{}) {
			fmt.Println("unknown method name", methodName)
			goto top_loop
		}

		//메소드 인자수 확인.
		methodType := method.Type()
		if methodType.NumIn() != len(inputs) {
			fmt.Println("mismatch", args[0], "input number")
			goto top_loop
		}

		//메소드 call
		input := []reflect.Value{}
		for i, e := range inputs {
			switch method.Type().In(i).Kind() {
			case reflect.String:
				input = append(input, reflect.ValueOf(e))
			case reflect.Uint:
				if u64, err := strconv.ParseUint(e, 10, 32); err != nil {
					fmt.Println(e, err)
					goto top_loop
				} else {
					input = append(input, reflect.ValueOf(uint(u64)))
				}
			default:
				fmt.Println(method.Type().In(i).Kind(), "type not supported")
				goto top_loop
			}
		}
		method.Call(input)
	}
}
