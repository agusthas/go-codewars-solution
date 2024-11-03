package src

import "testing"

// TODO: Write your test cases here

func TestFizzBuzzCuckooClock(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "13:34", args: args{time: "13:34"}, want: "tick"},
		{name: "21:00", args: args{time: "21:00"}, want: "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
		{name: "11:15", args: args{time: "11:15"}, want: "Fizz Buzz"},
		{name: "03:03", args: args{time: "03:03"}, want: "Fizz"},
		{name: "14:30", args: args{time: "14:30"}, want: "Cuckoo"},
		{name: "08:55", args: args{time: "08:55"}, want: "Buzz"},
		{name: "00:00", args: args{time: "00:00"}, want: "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
		{name: "12:00", args: args{time: "12:00"}, want: "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzzCuckooClock(tt.args.time); got != tt.want {
				t.Errorf("FizzBuzzCuckooClock() = %v, want %v", got, tt.want)
			}
		})
	}
}
