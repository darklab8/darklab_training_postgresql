package golang

import (
	"darklab_training_postgres/golang/settings"
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/testdb"
	"fmt"
	"testing"
)

func RunSubTests(task_number string, t *testing.T, test_func func(db_params testdb.DBParams)) {
	task_name := fmt.Sprintf("test%s", task_number)

	t.Run(task_name+"_unit_test", func(t *testing.T) {
		test_func(testdb.UnitTests)
	})

	if settings.ENABLED_PERFORMANCE_TESTS {
		t.Run(task_name+"_perf_with_index", func(t *testing.T) {
			shared.FixtureTimeMeasure(func() {
				test_func(testdb.PerformanceWithIndexes)
			}, fmt.Sprintf("%s_perf_with_index, params=%v", task_name, testdb.PerformanceWithIndexes.ToStr()))
		})
		t.Run(task_name+"_perf__without_index", func(t *testing.T) {
			shared.FixtureTimeMeasure(func() {
				test_func(testdb.PerformanceIndexless)
			}, fmt.Sprintf("%s_perf_without_index, params=%v", task_name, testdb.PerformanceIndexless.ToStr()))
		})
	}
}
