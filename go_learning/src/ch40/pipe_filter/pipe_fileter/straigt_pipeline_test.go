package pipe

import "testing"

/**
go test -v ./filter.go ./straight_pipeline.go ./sum_filter.go ./split_filter.go ./straigt_pipeline_test.go ./to_int_filter.go
*/
func TestStraightPipeLine(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("p1", spliter, converter, sum)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("the expected is 6 , but the actual is %d", ret)
	}
	t.Log(ret)
}
