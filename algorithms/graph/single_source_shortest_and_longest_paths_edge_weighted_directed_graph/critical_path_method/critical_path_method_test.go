package critical_path_method

import "testing"

type TestCase struct {
	numJobs           int
	overallFinishTime float64
	schedules         []JobSchedule
}

type JobSchedule struct {
	job    int
	start  float64
	finish float64
}

func TestCriticalPathMethod(t *testing.T) {
	tc := TestCase{
		numJobs:           10,
		overallFinishTime: 173.0,
		schedules: []JobSchedule{
			{
				job:    0,
				start:  0.0,
				finish: 41.0,
			},
			{
				job:    1,
				start:  41.0,
				finish: 92.0,
			},
			{
				job:    2,
				start:  123.0,
				finish: 173.0,
			},
			{
				job:    3,
				start:  91.0,
				finish: 127.0,
			},
			{
				job:    4,
				start:  70.0,
				finish: 108.0,
			},
			{
				job:    5,
				start:  0.0,
				finish: 45.0,
			},
			{
				job:    6,
				start:  70.0,
				finish: 91.0,
			},
			{
				job:    7,
				start:  41.0,
				finish: 73.0,
			},
			{
				job:    8,
				start:  91.0,
				finish: 123.0,
			},
			{
				job:    9,
				start:  41.0,
				finish: 70.0,
			},
		},
	}

	cpm := NewCPM("fixtures/jobs.txt")
	if cpm.OverallFinishTime() != tc.overallFinishTime {
		t.Errorf("Got overall finish time of %.1f, want %.1f", cpm.OverallFinishTime(), tc.overallFinishTime)
	}
	for i := 0; i < tc.numJobs; i++ {
		if cpm.StartTimeOf(i) != tc.schedules[i].start {
			t.Errorf("Got start time of %.1f for job %v, want %.1f", cpm.StartTimeOf(i), i, tc.schedules[i].start)
		}

		if cpm.FinishTimeOf(i) != tc.schedules[i].finish {
			t.Errorf("Got finish time of %.1f for job %v, want %.1f", cpm.FinishTimeOf(i), i, tc.schedules[i].finish)
		}
	}
}
