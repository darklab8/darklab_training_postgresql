package golang

import (
	"darklab_training_postgres/golang/shared"
	"fmt"
	"testing"
	"time"
)

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	fmt.Println()
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}

func TestFoo(t *testing.T) {
	N := 10
	function := func(t *testing.T) {
		fmt.Println(N)
	}
	// <setup code>
	t.Run("A=1", function)
	t.Run("A=2", function)
	t.Run("B=1", function)
	// <tear-down code>
}

func TestGroupedParallel(t *testing.T) {
	N := 10
	function := func(t *testing.T) {
		t.Parallel()
		time.Sleep(1 * time.Second)
		fmt.Println(N)
	}

	shared.FixtureTimeMeasure(func() {

		t.Run("parallel", func(t *testing.T) {
			t.Run("parallel1", function)
			t.Run("parallel2", function)
			t.Run("parallel3", function)
			fmt.Println("parallel finished")
		})
	}, "parallel test run finished")
}
