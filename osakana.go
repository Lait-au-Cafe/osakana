package main
import(
	"fmt"
	"reflect"
)

func main() {
	m ,e := simulate()
	fmt.Println("Score : (", m, ", ", e, ")")
	printAutomaton()
}

func simulate() (myScore int, enScore int){
	
	myInnerState := InnerVariables{}
	prevMyState := 's'
	prevEnemyState := 's'

	var myState, enState rune

	for i := 0; i < 100; i++ {
		myState, myInnerState = myInnerState.transit(prevEnemyState)
		enState = getEnemyState(prevMyState)
		if myState == 'g' && enState == 'g' {
			myScore--
			enScore--
		} else if myState == 'g' && enState == 'w' {
			myScore += 7
			enScore -= 3
		} else if myState == 'w' && enState == 'g' {
			myScore -= 3
			enScore += 7
		} else if myState == 'w' && enState == 'w' {
			myScore += 3
			enScore += 3
		} else {
			panic("Undefined State(" + string(myState) + "," + string(enState) + ")")
		}

		prevMyState = myState
		prevEnemyState = enState

		fmt.Println(i, ":(", myScore, ", ", enScore, ")")
	}

	return
}

func getEnemyState(prevMyState rune) rune {
	switch prevMyState {
	case 'g':
		return 'g'
	case 'w':
		return 'w'
	case 's':
		return 'g'
	default:
		panic("Undefined State")
	}
}


// InnerVariables 用いる内部変数はここで管理すること
// 内部状態を表す メンバ変数はグローバルスコープにすること
type InnerVariables struct {
	Count int
}

// !!!!!関数内部には絶対に変数を定義しないこと!!!!!
func (cur InnerVariables) transit(prevEnemyState rune) (rune, InnerVariables) {
	if cur.Count < 17 {
		cur.Count++
	}

	if cur.Count == 15 || cur.Count == 16 {
		return 'w', cur
	}
	
	switch prevEnemyState {
	case 'g':
		return 'g', cur
	case 'w':
		return 'w', cur
	case 's':
		return 'g', cur
	default:
		panic("Undefined State")
	}
}

func (cur InnerVariables) equals(other InnerVariables) bool {
	myVals := reflect.ValueOf(cur)
	otherVals := reflect.ValueOf(other)

	flg := true
	for i := 0; i < myVals.NumField(); i++ {
		flg = flg && (myVals.Field(i).Interface() == otherVals.Field(i).Interface())
	}

	return flg
}

// State 変換用の状態保持用構造体
type State struct {
	Output rune
	Vars InnerVariables
}
func (s State) equals(other State) bool {
	return  (s.Output == other.Output) && (s.Vars.equals(other.Vars))
}

func printAutomaton() {
	// 状態の集合
	var states []State

	// 開始前状態
	current := State{}

	// 初期状態の登録
	out, vars := current.Vars.transit('s')
	states = append(states, State{out, vars})

	// 入力の種類
	inputs := []rune{'g', 'w'}

	for index := 0; index < len(states); index++ {

		current := states[index]
		fmt.Print(index, ":", string(current.Output))
		
		for _, r := range inputs {
			out, vars = current.Vars.transit(r)
			next := State{out, vars}

			isFound := -1
			for i, v := range states {
				if v.equals(next) {
					isFound = i
					break
				}
			}

			// 存在しない状態の場合は新たに加える
			if isFound < 0 {
				fmt.Print(",", len(states))
				states = append(states, next)
			} else {
				fmt.Print(",", isFound)
			}
		}

		fmt.Println()
	}
}