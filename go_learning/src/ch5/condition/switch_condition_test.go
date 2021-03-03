package condition

import "testing"

/**
1:case 支持多值判断
2:自带 break
*/
func TestSwitchMultiCase(t *testing.T) {

	for i := 0; i < 5; i++ {

		switch i {

		case 1, 3:
			t.Logf("%d is Odd", i)
		case 0, 2:
			t.Logf("%d is Even", i)
		default:
			t.Logf("%d not 1 - 3 number", i)

		}

	}

}

/**
1:自动加 break
2:case 可以判断条件
*/
func TestSwitchCondition(t *testing.T) {

	for i := 0; i < 5; i++ {
		switch {

		case i%2 == 0:
			t.Logf("%d is Even", i)
		case i%2 == 1:
			t.Logf("%d is Odd", i)
		default:
			t.Logf("%d not Even/Odd", i)

		}
	}

}
